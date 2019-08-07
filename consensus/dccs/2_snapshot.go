package dccs

import (
	"bytes"
	"encoding/json"
	"sort"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rlp"
	lru "github.com/hashicorp/golang-lru"
	"golang.org/x/crypto/sha3"
)

type Proposal int

const (
	Join  Proposal = 0
	Leave Proposal = 1
	Slash Proposal = 2
)

// Snapshot2 is the state of the authorization voting at a given point in time.
type Snapshot2 struct {
	config   *params.DccsConfig // Consensus engine parameters to fine tune behavior
	sigcache *lru.ARCCache      // Cache of recent block signatures to speed up ecrecover

	Number     uint64                            `json:"number"`     // Block number where the snapshot was created
	Hash       common.Hash                       `json:"hash"`       // Block hash where the snapshot was created
	SealerHash common.Hash                       `json:"sealerHash"` // Block hash to order the sealer list
	Sealers    map[common.Address]common.Address `json:"sealers"`    // Set of authorized sealers at this moment
	Recents    map[uint64]common.Address         `json:"recents"`    // Set of recent signers for spam protections
}

// Sealer contain data of sealer's (signer, beneficiary) address and checkpoint block hash
type Sealer struct {
	Hash        common.Hash
	Signer      common.Address
	Beneficiary common.Address
}

// sealersAscendingByHash implements the sort interface to allow sorting a list of signers by hash
type sealersAscendingByHash []Sealer

func (sn sealersAscendingByHash) Len() int { return len(sn) }
func (sn sealersAscendingByHash) Less(i, j int) bool {
	h1 := rlpHash2(sn[i])
	h2 := rlpHash2(sn[j])
	return bytes.Compare(h1[:], h2[:]) < 0
}
func (sn sealersAscendingByHash) Swap(i, j int) { sn[i], sn[j] = sn[j], sn[i] }

// newSnapshot2 creates a new snapshot with the specified startup parameters. This
// method does not initialize the set of recent signers, so only ever use if for
// the genesis block.
func newSnapshot2(config *params.DccsConfig, sigcache *lru.ARCCache, number uint64, hash, sealerHash common.Hash, signers, beneficiaries []common.Address) *Snapshot2 {
	snap := &Snapshot2{
		config:     config,
		sigcache:   sigcache,
		Number:     number,
		Hash:       hash,
		SealerHash: sealerHash,
		Sealers:    make(map[common.Address]common.Address),
		Recents:    make(map[uint64]common.Address),
	}
	for i, signer := range signers {
		snap.Sealers[signer] = beneficiaries[i]
	}
	return snap
}

// loadSnapshot2 loads an existing snapshot from the database.
func loadSnapshot2(config *params.DccsConfig, sigcache *lru.ARCCache, db ethdb.Database, hash common.Hash) (*Snapshot2, error) {
	blob, err := db.Get(append([]byte("dccs2-"), hash[:]...))
	if err != nil {
		return nil, err
	}
	snap := new(Snapshot2)
	if err := json.Unmarshal(blob, snap); err != nil {
		return nil, err
	}
	snap.config = config
	snap.sigcache = sigcache

	return snap, nil
}

// store inserts the snapshot into the database.
func (s *Snapshot2) store(db ethdb.Database) error {
	blob, err := json.Marshal(s)
	if err != nil {
		return err
	}
	return db.Put(append([]byte("dccs2-"), s.Hash[:]...), blob)
}

// copy creates a deep copy of the snapshot, though not the individual votes.
func (s *Snapshot2) copy() *Snapshot2 {
	cpy := &Snapshot2{
		config:     s.config,
		sigcache:   s.sigcache,
		Number:     s.Number,
		Hash:       s.Hash,
		SealerHash: s.SealerHash,
		Sealers:    make(map[common.Address]common.Address),
		Recents:    make(map[uint64]common.Address),
	}
	for sealer := range s.Sealers {
		cpy.Sealers[sealer] = s.Sealers[sealer]
	}
	for block, signer := range s.Recents {
		cpy.Recents[block] = signer
	}

	return cpy
}

// apply creates a new authorization snapshot by applying the given headers to
// the original one.
func (s *Snapshot2) apply(headers []*types.Header) (*Snapshot2, error) {
	// Allow passing in no headers for cleaner code
	if len(headers) == 0 {
		return s, nil
	}
	// Sanity check that the headers can be applied
	for i := 0; i < len(headers)-1; i++ {
		if headers[i+1].Number.Uint64() != headers[i].Number.Uint64()+1 {
			return nil, errInvalidVotingChain
		}
	}
	if headers[0].Number.Uint64() != s.Number+1 {
		return nil, errInvalidVotingChain
	}
	// Iterate through the headers and create a new snapshot
	snap := s.copy()

	var (
		start  = time.Now()
		logged = time.Now()
	)
loop:
	for i, header := range headers {
		number := header.Number.Uint64()
		// Delete the oldest signer from the recent list to allow it signing again
		if limit := uint64(len(snap.Sealers)/2 + 1); number >= limit {
			delete(snap.Recents, number-limit)
		}
		// Resolve the authorization key and check against signers
		signer, err := ecrecover(header, s.sigcache)
		if err != nil {
			return nil, err
		}
		if _, ok := snap.Sealers[signer]; !ok {
			return nil, errUnauthorizedSigner
		}
		for _, recent := range snap.Recents {
			if recent == signer {
				return nil, errRecentlySigned
			}
		}
		snap.Recents[number] = signer

		// Apply proposal from block header
		var proposal Proposal
		switch {
		case bytes.Equal(header.Nonce[:], nonceSealerNone):
			continue loop
		case bytes.Equal(header.Nonce[:], nonceSealerJoin):
			proposal = Join
		case bytes.Equal(header.Nonce[:], nonceSealerLeave):
			proposal = Leave
		case bytes.Equal(header.Nonce[:], nonceSealerSlash):
			proposal = Slash
		default:
			return nil, errInvalidVote
		}

		// Get (signer, beneficiary) from header block
		var beneficiary common.Address
		copy(signer[:], header.Extra[extraVanity:])
		copy(beneficiary[:], header.Extra[extraVanity+common.AddressLength:])

		switch proposal {
		case Join:
			snap.Sealers[signer] = beneficiary
		default:
			delete(snap.Sealers, signer)
			// Signer list shrunk, delete any leftover recent caches
			if limit := uint64(len(snap.Sealers)/2 + 1); number >= limit {
				delete(snap.Recents, number-limit)
			}
		}
		// If we're taking too much time (ecrecover), notify the user once a while
		if time.Since(logged) > 8*time.Second {
			log.Info("Reconstructing voting history", "processed", i, "total", len(headers), "elapsed", common.PrettyDuration(time.Since(start)))
			logged = time.Now()
		}
	}
	if time.Since(start) > 8*time.Second {
		log.Info("Reconstructed voting history", "processed", len(headers), "elapsed", common.PrettyDuration(time.Since(start)))
	}
	snap.Number += uint64(len(headers))
	snap.Hash = headers[len(headers)-1].Hash()

	return snap, nil
}

// signers retrieves the list of authorized signers in hash ascending order.
func (s *Snapshot2) signers() []Sealer {
	sealers := make([]Sealer, 0, len(s.Sealers))
	for signer := range s.Sealers {
		sealers = append(sealers, Sealer{Hash: s.SealerHash, Signer: signer, Beneficiary: s.Sealers[signer]})
	}
	sort.Sort(sealersAscendingByHash(sealers))
	return sealers
}

// inturn returns if a signer at a given block height is in-turn or not.
func (s *Snapshot2) inturn(signer common.Address, parent *types.Header) bool {
	offset, err := s.offset(signer, parent)
	if err != nil {
		return false
	}
	log.Trace("inturn", "offset", offset)
	return offset == 0
}

func signerPosition2(signer common.Address, sealers []Sealer) (int, bool) {
	for i, sig := range sealers {
		if sig.Signer == signer {
			return i, true
		}
	}
	return -1, false
}

func (s *Snapshot2) offset(signer common.Address, parent *types.Header) (int, error) {
	signers := s.signers()
	n := len(signers)
	if n <= 1 {
		// no competition
		return 0, nil
	}

	pos, ok := signerPosition2(signer, signers)
	if !ok {
		// unable to find the signer position
		return n, errUnauthorizedSigner
	}

	if parent == nil || s.config.IsCheckpoint(parent.Number.Uint64()+1) {
		// first block of an epoch, just return the rightful order
		log.Debug("the first block of an epoch")
		return pos, nil
	}

	// Resolve the last authorization key and check against signer
	prevSigner, err := ecrecover(parent, s.sigcache)
	if err != nil {
		return 0, err
	}
	prevPos, ok := signerPosition2(prevSigner, signers)
	if !ok {
		// unable to find the previous signer position
		return n, errUnauthorizedSigner
	}

	offset := pos - prevPos - 1
	if offset < 0 {
		offset += n
	}

	log.Debug("offset", "signer position", pos, "previous signer position", prevPos, "len(signers)", n, "offset", offset)

	return offset, nil
}

// difficulty returns the block weight at a given block height for a signer.
// Turn-ness is the directional distant from a signer to the previous one,
// following a circular order of the signers list.
// @return maximum value = len(signers) if signer is right after the prevSigner (circularly)
// @return minimum value = 1 if the signer is right before the prevSigner (circularly)
// @return invalid value = 0 if the signer or parent signer is not on the sealer list
func (s *Snapshot2) difficulty(signer common.Address, parent *types.Header) uint64 {
	offset, err := s.offset(signer, parent)
	if err != nil {
		return 0
	}

	signers := s.signers()
	n := len(signers)

	return uint64(n - offset)
}

// rlpHash2 return hash of an input
func rlpHash2(s Sealer) (h common.Hash) {
	hasher := sha3.NewLegacyKeccak256()
	rlp.Encode(hasher, []interface{}{
		s.Signer,
		s.Hash,
	})
	hasher.Sum(h[:0])
	return h
}
