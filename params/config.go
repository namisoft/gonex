// Copyright 2016 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

import (
	"encoding/binary"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	// BlockSeconds is the average block time
	BlockSeconds = 2

	// CanonicalDepth is the confirmation required for a block to be considered canonical.
	// Also the distant from the ThangLong checkpoint block to it's snapshot.
	CanonicalDepth = 7
)

// Genesis hashes to enforce below configs on.
var (
	MainnetGenesisHash = common.HexToHash("0x080eeb525df0e852343ba13afedf2b256f0991c1b18797e18863fd7b4ab3574b")
	TestnetGenesisHash = common.HexToHash("0xe53f0441b5f887499e49cb628262c4acc897492a4fe3236bb3c68023ea725f0d")
	RinkebyGenesisHash = common.HexToHash("0x6341fd3daf94b748c72ced5a5b26028f2474f5f00d824504e4fa37a75767e177")
	GoerliGenesisHash  = common.HexToHash("0xbf7e331f7f7c1dd2e05159666b3bf8bc7a8a3a9eb1d518969eab529dd9b88c1a")
	ZeroAddress        = common.HexToAddress("0x0000000000000000000000000000000000000000")
	TokenAddress       = common.HexToAddress("0x2c783ad80ff980ec75468477e3dd9f86123ecbda") // NTF token contract address
	// CoLoa contract addresses
	SeigniorageAddress   = common.HexToAddress("0x0000000000000000000000000000000000123456") // Seigniorage contract address
	VolatileTokenAddress = common.HexToAddress("0x0000000000000000000000000000000001234567") // MNTY token contract address
	StableTokenAddress   = common.HexToAddress("0x0000000000000000000000000000000012345678") // NUSD token contract address
)

// TrustedCheckpoints associates each known checkpoint with the genesis hash of
// the chain it belongs to.
var TrustedCheckpoints = map[common.Hash]*TrustedCheckpoint{
	MainnetGenesisHash: MainnetTrustedCheckpoint,
	TestnetGenesisHash: TestnetTrustedCheckpoint,
	RinkebyGenesisHash: RinkebyTrustedCheckpoint,
	GoerliGenesisHash:  GoerliTrustedCheckpoint,
}

// CheckpointOracles associates each known checkpoint oracles with the genesis hash of
// the chain it belongs to.
var CheckpointOracles = map[common.Hash]*CheckpointOracleConfig{
	MainnetGenesisHash: MainnetCheckpointOracle,
	TestnetGenesisHash: TestnetCheckpointOracle,
	RinkebyGenesisHash: RinkebyCheckpointOracle,
	GoerliGenesisHash:  GoerliCheckpointOracle,
}

var (
	// MainnetChainConfig is the chain parameters to run a node on the main network.
	MainnetChainConfig = &ChainConfig{
		ChainID:             big.NewInt(66666),
		HomesteadBlock:      big.NewInt(1),
		DAOForkBlock:        nil,
		DAOForkSupport:      true,
		EIP150Block:         big.NewInt(2),
		EIP150Hash:          common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
		EIP155Block:         big.NewInt(3),
		EIP158Block:         big.NewInt(3),
		ByzantiumBlock:      big.NewInt(4),
		ConstantinopleBlock: big.NewInt(15360000),
		PetersburgBlock:     big.NewInt(15360000),
		IstanbulBlock:       nil,
		Dccs: &DccsConfig{
			Period: BlockSeconds,
			Epoch:  30000,
			// Governance contract
			Contract: common.HexToAddress("0x0000000000000000000000000000000000012345"),
			// ThangLong hard-fork
			StakeRequire:    50000,
			StakeLockHeight: 30000,
			ThangLongBlock:  big.NewInt(15360000),
			ThangLongEpoch:  3000,
			// CoLoa hard-fork
			CoLoaBlock:            big.NewInt(30000000),
			PriceSamplingDuration: 7 * 24 * 60 * 60 / BlockSeconds,
			PriceSamplingInterval: 10*60/BlockSeconds - 7,
			AbsorptionDuration:    7 * 24 * 60 * 60 / BlockSeconds / 2,
			AbsorptionExpiration:  7 * 24 * 60 * 60 / BlockSeconds,
			SlashingDuration:      7 * 24 * 60 * 60 / BlockSeconds / 2,
			LockdownExpiration:    7 * 24 * 60 * 60 / BlockSeconds * 2,
		},
	}

	// MainnetTrustedCheckpoint contains the light client trusted checkpoint for the main network.
	MainnetTrustedCheckpoint = &TrustedCheckpoint{
		SectionIndex: 561,
		SectionHead:  common.HexToHash("0x6b87a31980a1fe764e5db97ea7c841b9d6db88fe8cd0506937532bae7eab783c"),
		CHTRoot:      common.Hash{}, // TODO: will calculate later when using LES
		BloomRoot:    common.Hash{}, // TODO: will calculate later when using LES
	}

	// MainnetCheckpointOracle contains a set of configs for the main network oracle.
	MainnetCheckpointOracle = &CheckpointOracleConfig{
		Address: common.HexToAddress("0x9a9070028361F7AAbeB3f2F2Dc07F82C4a98A02a"),
		Signers: []common.Address{
			common.HexToAddress("0x1b2C260efc720BE89101890E4Db589b44E950527"), // Peter
			common.HexToAddress("0x78d1aD571A1A09D60D9BBf25894b44e4C8859595"), // Martin
			common.HexToAddress("0x286834935f4A8Cfb4FF4C77D5770C2775aE2b0E7"), // Zsolt
			common.HexToAddress("0xb86e2B0Ab5A4B1373e40c51A7C712c70Ba2f9f8E"), // Gary
			common.HexToAddress("0x0DF8fa387C602AE62559cC4aFa4972A7045d6707"), // Guillaume
		},
		Threshold: 2,
	}

	// TestnetChainConfig contains the chain parameters to run a node on the Dccs test network.
	TestnetChainConfig = &ChainConfig{
		ChainID:             big.NewInt(111111),
		HomesteadBlock:      big.NewInt(1),
		DAOForkBlock:        nil,
		DAOForkSupport:      true,
		EIP150Block:         big.NewInt(2),
		EIP150Hash:          common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
		EIP155Block:         big.NewInt(3),
		EIP158Block:         big.NewInt(3),
		ByzantiumBlock:      big.NewInt(4),
		ConstantinopleBlock: big.NewInt(300000),
		PetersburgBlock:     big.NewInt(300000),
		IstanbulBlock:       nil,
		Dccs: &DccsConfig{
			Period: 2,
			Epoch:  30000,
			// Governance contract
			Contract: common.HexToAddress("0x0000000000000000000000000000000000012345"),
			// ThangLong hard-fork
			StakeRequire:    500,
			StakeLockHeight: 3000,
			ThangLongBlock:  big.NewInt(300000),
			ThangLongEpoch:  300,
		},
	}

	// TestnetTrustedCheckpoint contains the light client trusted checkpoint for the Ropsten test network.
	TestnetTrustedCheckpoint = &TrustedCheckpoint{
		SectionIndex: 19,
		SectionHead:  common.HexToHash("0x57a235095aba5f4c6cf7a1787c283003e93adadf68f46ecacbcc777d505d83b2"),
		CHTRoot:      common.Hash{}, // TODO: will calculate later when using LES
		BloomRoot:    common.Hash{}, // TODO: will calculate later when using LES
	}

	// TestnetCheckpointOracle contains a set of configs for the Ropsten test network oracle.
	TestnetCheckpointOracle = &CheckpointOracleConfig{
		Address: common.HexToAddress("0xEF79475013f154E6A65b54cB2742867791bf0B84"),
		Signers: []common.Address{
			common.HexToAddress("0x32162F3581E88a5f62e8A61892B42C46E2c18f7b"), // Peter
			common.HexToAddress("0x78d1aD571A1A09D60D9BBf25894b44e4C8859595"), // Martin
			common.HexToAddress("0x286834935f4A8Cfb4FF4C77D5770C2775aE2b0E7"), // Zsolt
			common.HexToAddress("0xb86e2B0Ab5A4B1373e40c51A7C712c70Ba2f9f8E"), // Gary
			common.HexToAddress("0x0DF8fa387C602AE62559cC4aFa4972A7045d6707"), // Guillaume
		},
		Threshold: 2,
	}

	// RinkebyChainConfig contains the chain parameters to run a node on the Rinkeby test network.
	RinkebyChainConfig = &ChainConfig{
		ChainID:             big.NewInt(4),
		HomesteadBlock:      big.NewInt(1),
		DAOForkBlock:        nil,
		DAOForkSupport:      true,
		EIP150Block:         big.NewInt(2),
		EIP150Hash:          common.HexToHash("0x9b095b36c15eaf13044373aef8ee0bd3a382a5abb92e402afa44b8249c3a90e9"),
		EIP155Block:         big.NewInt(3),
		EIP158Block:         big.NewInt(3),
		ByzantiumBlock:      big.NewInt(1035301),
		ConstantinopleBlock: big.NewInt(3660663),
		PetersburgBlock:     big.NewInt(4321234),
		IstanbulBlock:       nil,
		Clique: &CliqueConfig{
			Period: 15,
			Epoch:  30000,
		},
	}

	// DccsChainConfig contains the chain parameters to run a node on the Nexty test network.
	DccsChainConfig = &ChainConfig{
		ChainID:             big.NewInt(66666),
		HomesteadBlock:      big.NewInt(1),
		DAOForkBlock:        nil,
		DAOForkSupport:      true,
		EIP150Block:         big.NewInt(2),
		EIP150Hash:          common.HexToHash("0x9b095b36c15eaf13044373aef8ee0bd3a382a5abb92e402afa44b8249c3a90e9"),
		EIP155Block:         big.NewInt(3),
		EIP158Block:         big.NewInt(3),
		ByzantiumBlock:      big.NewInt(1035301),
		ConstantinopleBlock: nil,
		PetersburgBlock:     nil,
		Dccs: &DccsConfig{
			Period: 2,
			Epoch:  30000,
		},
	}

	// RinkebyTrustedCheckpoint contains the light client trusted checkpoint for the Rinkeby test network.
	RinkebyTrustedCheckpoint = &TrustedCheckpoint{
		SectionIndex: 148,
		SectionHead:  common.HexToHash("0x45918f4686732c2a3e80827e1bc39cdb6a27fa362ddfe1fdfb61c69a7f1df1a9"),
		CHTRoot:      common.HexToHash("0x8ac7046391fec14834a2a0183513937c0b5f696666545991477d24b067008961"),
		BloomRoot:    common.HexToHash("0xfe4b852517612d7da54bf7e9fc18861a83171a93c72583bb6a61893b74422168"),
	}

	// RinkebyCheckpointOracle contains a set of configs for the Rinkeby test network oracle.
	RinkebyCheckpointOracle = &CheckpointOracleConfig{
		Address: common.HexToAddress("0xebe8eFA441B9302A0d7eaECc277c09d20D684540"),
		Signers: []common.Address{
			common.HexToAddress("0xd9c9cd5f6779558b6e0ed4e6acf6b1947e7fa1f3"), // Peter
			common.HexToAddress("0x78d1aD571A1A09D60D9BBf25894b44e4C8859595"), // Martin
			common.HexToAddress("0x286834935f4A8Cfb4FF4C77D5770C2775aE2b0E7"), // Zsolt
			common.HexToAddress("0xb86e2B0Ab5A4B1373e40c51A7C712c70Ba2f9f8E"), // Gary
		},
		Threshold: 2,
	}

	// GoerliChainConfig contains the chain parameters to run a node on the Görli test network.
	GoerliChainConfig = &ChainConfig{
		ChainID:             big.NewInt(5),
		HomesteadBlock:      big.NewInt(0),
		DAOForkBlock:        nil,
		DAOForkSupport:      true,
		EIP150Block:         big.NewInt(0),
		EIP155Block:         big.NewInt(0),
		EIP158Block:         big.NewInt(0),
		ByzantiumBlock:      big.NewInt(0),
		ConstantinopleBlock: big.NewInt(0),
		PetersburgBlock:     big.NewInt(0),
		IstanbulBlock:       nil,
		Clique: &CliqueConfig{
			Period: 15,
			Epoch:  30000,
		},
	}

	// GoerliTrustedCheckpoint contains the light client trusted checkpoint for the Görli test network.
	GoerliTrustedCheckpoint = &TrustedCheckpoint{
		SectionIndex: 32,
		SectionHead:  common.HexToHash("0x50eaedd8361fa9edd0ac2dec410310b9bdf67b963b60f3b1dce47f84b30670f9"),
		CHTRoot:      common.HexToHash("0x6504db73139f75ffa9102ae980e41b361cf3d5b66cea06c79cde9f457368820c"),
		BloomRoot:    common.HexToHash("0x7551ae027bb776252a20ded51ee2ff0cbfbd1d8d57261b9161cc1f2f80237001"),
	}

	// GoerliCheckpointOracle contains a set of configs for the Goerli test network oracle.
	GoerliCheckpointOracle = &CheckpointOracleConfig{
		Address: common.HexToAddress("0x18CA0E045F0D772a851BC7e48357Bcaab0a0795D"),
		Signers: []common.Address{
			common.HexToAddress("0x4769bcaD07e3b938B7f43EB7D278Bc7Cb9efFb38"), // Peter
			common.HexToAddress("0x78d1aD571A1A09D60D9BBf25894b44e4C8859595"), // Martin
			common.HexToAddress("0x286834935f4A8Cfb4FF4C77D5770C2775aE2b0E7"), // Zsolt
			common.HexToAddress("0xb86e2B0Ab5A4B1373e40c51A7C712c70Ba2f9f8E"), // Gary
			common.HexToAddress("0x0DF8fa387C602AE62559cC4aFa4972A7045d6707"), // Guillaume
		},
		Threshold: 2,
	}

	// AllEthashProtocolChanges contains every protocol change (EIPs) introduced
	// and accepted by the Ethereum core developers into the Ethash consensus.
	//
	// This configuration is intentionally not using keyed fields to force anyone
	// adding flags to the config to also have to set these fields.
	AllEthashProtocolChanges = &ChainConfig{big.NewInt(1337), big.NewInt(0), nil, false, big.NewInt(0), common.Hash{}, big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), nil, nil, new(EthashConfig), nil, nil}

	// AllCliqueProtocolChanges contains every protocol change (EIPs) introduced
	// and accepted by the Ethereum core developers into the Clique consensus.
	//
	// This configuration is intentionally not using keyed fields to force anyone
	// adding flags to the config to also have to set these fields.
	AllCliqueProtocolChanges = &ChainConfig{big.NewInt(1337), big.NewInt(0), nil, false, big.NewInt(0), common.Hash{}, big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), nil, nil, nil, &CliqueConfig{Period: 0, Epoch: 30000}, nil}

	// AllDccsProtocolChanges contains every protocol change (EIPs) introduced
	// and accepted by the Ethereum core developers into the Dccs consensus.
	//
	// This configuration is intentionally not using keyed fields to force anyone
	// adding flags to the config to also have to set these fields.
	AllDccsProtocolChanges = &ChainConfig{big.NewInt(1337), big.NewInt(0), nil, false, big.NewInt(0), common.Hash{}, big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), nil, nil, nil, nil, &DccsConfig{Period: 0, Epoch: 30000, ThangLongBlock: big.NewInt(0), ThangLongEpoch: 3000, Contract: common.HexToAddress("0x0")}}

	TestChainConfig = &ChainConfig{big.NewInt(1), big.NewInt(0), nil, false, big.NewInt(0), common.Hash{}, big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), nil, nil, new(EthashConfig), nil, nil}
	TestRules       = TestChainConfig.Rules(new(big.Int))
)

// TrustedCheckpoint represents a set of post-processed trie roots (CHT and
// BloomTrie) associated with the appropriate section index and head hash. It is
// used to start light syncing from this checkpoint and avoid downloading the
// entire header chain while still being able to securely access old headers/logs.
type TrustedCheckpoint struct {
	SectionIndex uint64      `json:"sectionIndex"`
	SectionHead  common.Hash `json:"sectionHead"`
	CHTRoot      common.Hash `json:"chtRoot"`
	BloomRoot    common.Hash `json:"bloomRoot"`
}

// HashEqual returns an indicator comparing the itself hash with given one.
func (c *TrustedCheckpoint) HashEqual(hash common.Hash) bool {
	if c.Empty() {
		return hash == common.Hash{}
	}
	return c.Hash() == hash
}

// Hash returns the hash of checkpoint's four key fields(index, sectionHead, chtRoot and bloomTrieRoot).
func (c *TrustedCheckpoint) Hash() common.Hash {
	buf := make([]byte, 8+3*common.HashLength)
	binary.BigEndian.PutUint64(buf, c.SectionIndex)
	copy(buf[8:], c.SectionHead.Bytes())
	copy(buf[8+common.HashLength:], c.CHTRoot.Bytes())
	copy(buf[8+2*common.HashLength:], c.BloomRoot.Bytes())
	return crypto.Keccak256Hash(buf)
}

// Empty returns an indicator whether the checkpoint is regarded as empty.
func (c *TrustedCheckpoint) Empty() bool {
	return c.SectionHead == (common.Hash{}) || c.CHTRoot == (common.Hash{}) || c.BloomRoot == (common.Hash{})
}

// CheckpointOracleConfig represents a set of checkpoint contract(which acts as an oracle)
// config which used for light client checkpoint syncing.
type CheckpointOracleConfig struct {
	Address   common.Address   `json:"address"`
	Signers   []common.Address `json:"signers"`
	Threshold uint64           `json:"threshold"`
}

// ChainConfig is the core config which determines the blockchain settings.
//
// ChainConfig is stored in the database on a per block basis. This means
// that any network, identified by its genesis block, can have its own
// set of configuration options.
type ChainConfig struct {
	ChainID *big.Int `json:"chainId"` // chainId identifies the current chain and is used for replay protection

	HomesteadBlock *big.Int `json:"homesteadBlock,omitempty"` // Homestead switch block (nil = no fork, 0 = already homestead)

	DAOForkBlock   *big.Int `json:"daoForkBlock,omitempty"`   // TheDAO hard-fork switch block (nil = no fork)
	DAOForkSupport bool     `json:"daoForkSupport,omitempty"` // Whether the nodes supports or opposes the DAO hard-fork

	// EIP150 implements the Gas price changes (https://github.com/ethereum/EIPs/issues/150)
	EIP150Block *big.Int    `json:"eip150Block,omitempty"` // EIP150 HF block (nil = no fork)
	EIP150Hash  common.Hash `json:"eip150Hash,omitempty"`  // EIP150 HF hash (needed for header only clients as only gas pricing changed)

	EIP155Block *big.Int `json:"eip155Block,omitempty"` // EIP155 HF block
	EIP158Block *big.Int `json:"eip158Block,omitempty"` // EIP158 HF block

	ByzantiumBlock      *big.Int `json:"byzantiumBlock,omitempty"`      // Byzantium switch block (nil = no fork, 0 = already on byzantium)
	ConstantinopleBlock *big.Int `json:"constantinopleBlock,omitempty"` // Constantinople switch block (nil = no fork, 0 = already activated)
	PetersburgBlock     *big.Int `json:"petersburgBlock,omitempty"`     // Petersburg switch block (nil = same as Constantinople)
	IstanbulBlock       *big.Int `json:"istanbulBlock,omitempty"`       // Istanbul switch block (nil = no fork, 0 = already on istanbul)
	EWASMBlock          *big.Int `json:"ewasmBlock,omitempty"`          // EWASM switch block (nil = no fork, 0 = already activated)

	// Various consensus engines
	Ethash *EthashConfig `json:"ethash,omitempty"`
	Clique *CliqueConfig `json:"clique,omitempty"`
	Dccs   *DccsConfig   `json:"dccs,omitempty"`
}

// EthashConfig is the consensus engine configs for proof-of-work based sealing.
type EthashConfig struct{}

// String implements the stringer interface, returning the consensus engine details.
func (c *EthashConfig) String() string {
	return "ethash"
}

// CliqueConfig is the consensus engine configs for proof-of-authority based sealing.
type CliqueConfig struct {
	Period uint64 `json:"period"` // Number of seconds between blocks to enforce
	Epoch  uint64 `json:"epoch"`  // Epoch length to reset votes and checkpoint
}

// String implements the stringer interface, returning the consensus engine details.
func (c *CliqueConfig) String() string {
	return "clique"
}

// DccsConfig is the consensus engine configs for proof-of-foundation based sealing.
// All epoch changing hardforks must have the activate block number divisible by both old and new epoch.
type DccsConfig struct {
	Period uint64 `json:"period"` // Number of seconds between blocks to enforce
	Epoch  uint64 `json:"epoch"`  // Epoch length to reset votes and checkpoint
	// governance smart contract address
	Contract common.Address `json:"contract,omitempty"`
	// Governance contract stake params
	StakeRequire    uint64 `json:"stakeRequire"`    // stake requirement
	StakeLockHeight uint64 `json:"stakeLockHeight"` // lock time (in blocks) after leaving
	// ThangLong hardfork
	ThangLongBlock *big.Int `json:"thangLongBlock,omitempty"` // ThangLong switch block (nil = no fork, 0 = already activated)
	ThangLongEpoch uint64   `json:"thangLongEpoch"`           // Epoch length to reset votes and checkpoint
	// CoLoa hardfork
	CoLoaBlock            *big.Int `json:"coLoaBlock,omitempty"`
	PriceSamplingDuration uint64   `json:"priceSamplingDuration"` // number of blocks to take price samples (a week)
	PriceSamplingInterval uint64   `json:"priceSamplingInterval"` // the largest prime number of blocks in 10 minutes
	AbsorptionDuration    uint64   `json:"absorptionDuration"`    // each block can absorb a maximum of targetAbsorption/absorptionDuration (half a week)
	AbsorptionExpiration  uint64   `json:"absorptionExpiration"`  // number of blocks that the absorption will be expired (a week)
	SlashingDuration      uint64   `json:"slashingDuration"`      // each block can slash a maximum value of d/D/slashingDuration (half a week)
	LockdownExpiration    uint64   `json:"lockdownExpiration"`    // number of blocks that the lockdown will be expired (2 weeks)
}

// IsPriceBlock returns whether a block could include a price
func (c *DccsConfig) IsPriceBlock(number uint64) bool {
	if c.IsCoLoa(new(big.Int).SetUint64(number)) {
		return number%c.PriceSamplingInterval == 0
	}
	return false
}

// PositionInEpoch returns the offset of a block from the start of an epoch
func (c *DccsConfig) PositionInEpoch(number uint64) uint64 {
	if c.IsThangLong(new(big.Int).SetUint64(number)) {
		return number % c.ThangLongEpoch
	}
	return number % c.Epoch
}

// IsCheckpoint returns whether a block is at the start of an epoch
func (c *DccsConfig) IsCheckpoint(number uint64) bool {
	return c.PositionInEpoch(number) == 0
}

// Checkpoint returns the epoch block for this block
func (c *DccsConfig) Checkpoint(number uint64) uint64 {
	return number - c.PositionInEpoch(number)
}

// Snapshot returns the snapshot block for this block's epoch.
// Snapshot is (Checkpoint - CanonicalDepth) for ThangLong consensus.
func (c *DccsConfig) Snapshot(number uint64) uint64 {
	// Get the checkpoint first
	cp := c.Checkpoint(number)
	// Get genesis block as checkpoint for 1st epoch
	if cp <= CanonicalDepth {
		return 0
	}
	// Get the state from canonical chain to ensure the chain and state are not in sidefork
	return cp - CanonicalDepth
}

// String implements the stringer interface, returning the consensus engine details.
func (c *DccsConfig) String() string {
	return fmt.Sprintf("dccs {ThangLong: %v Epoch: %v Contract: %v CoLoa: %v PriceDuration: %v PriceInterval: %v AbsorptionDuration: %v AbsorptionExpiration: %v SlashingDuration: %v LockdownExpiration: %v}",
		c.ThangLongBlock,
		c.ThangLongEpoch,
		c.Contract.String(),
		c.CoLoaBlock,
		c.PriceSamplingDuration,
		c.PriceSamplingInterval,
		c.AbsorptionDuration,
		c.AbsorptionExpiration,
		c.SlashingDuration,
		c.LockdownExpiration,
	)
}

// String implements the fmt.Stringer interface.
func (c *ChainConfig) String() string {
	var engine interface{}
	switch {
	case c.Ethash != nil:
		engine = c.Ethash
	case c.Clique != nil:
		engine = c.Clique
	case c.Dccs != nil:
		engine = c.Dccs
	default:
		engine = "unknown"
	}
	return fmt.Sprintf("{ChainID: %v Homestead: %v DAO: %v DAOSupport: %v EIP150: %v EIP155: %v EIP158: %v Byzantium: %v Constantinople: %v Petersburg: %v Istanbul: %v Engine: %v}",
		c.ChainID,
		c.HomesteadBlock,
		c.DAOForkBlock,
		c.DAOForkSupport,
		c.EIP150Block,
		c.EIP155Block,
		c.EIP158Block,
		c.ByzantiumBlock,
		c.ConstantinopleBlock,
		c.PetersburgBlock,
		c.IstanbulBlock,
		engine,
	)
}

// IsHomestead returns whether num is either equal to the homestead block or greater.
func (c *ChainConfig) IsHomestead(num *big.Int) bool {
	return isForked(c.HomesteadBlock, num)
}

// IsDAOFork returns whether num is either equal to the DAO fork block or greater.
func (c *ChainConfig) IsDAOFork(num *big.Int) bool {
	return isForked(c.DAOForkBlock, num)
}

// IsEIP150 returns whether num is either equal to the EIP150 fork block or greater.
func (c *ChainConfig) IsEIP150(num *big.Int) bool {
	return isForked(c.EIP150Block, num)
}

// IsEIP155 returns whether num is either equal to the EIP155 fork block or greater.
func (c *ChainConfig) IsEIP155(num *big.Int) bool {
	return isForked(c.EIP155Block, num)
}

// IsEIP158 returns whether num is either equal to the EIP158 fork block or greater.
func (c *ChainConfig) IsEIP158(num *big.Int) bool {
	return isForked(c.EIP158Block, num)
}

// IsByzantium returns whether num is either equal to the Byzantium fork block or greater.
func (c *ChainConfig) IsByzantium(num *big.Int) bool {
	return isForked(c.ByzantiumBlock, num)
}

// IsConstantinople returns whether num is either equal to the Constantinople fork block or greater.
func (c *ChainConfig) IsConstantinople(num *big.Int) bool {
	return isForked(c.ConstantinopleBlock, num)
}

// IsPetersburg returns whether num is either
// - equal to or greater than the PetersburgBlock fork block,
// - OR is nil, and Constantinople is active
func (c *ChainConfig) IsPetersburg(num *big.Int) bool {
	return isForked(c.PetersburgBlock, num) || c.PetersburgBlock == nil && isForked(c.ConstantinopleBlock, num)
}

// IsIstanbul returns whether num is either equal to the Istanbul fork block or greater.
func (c *ChainConfig) IsIstanbul(num *big.Int) bool {
	return isForked(c.IstanbulBlock, num)
}

// IsEWASM returns whether num represents a block number after the EWASM fork
func (c *ChainConfig) IsEWASM(num *big.Int) bool {
	return isForked(c.EWASMBlock, num)
}

// IsThangLongPreparationBlock returns whether num represents a block number exactly at the ThangLong Preparation
func (c *ChainConfig) IsThangLongPreparationBlock(num *big.Int) bool {
	return c.Dccs != nil && c.Dccs.ThangLongBlock != nil && c.Dccs.IsThangLongPreparationBlock(num)
}

// IsSnapshotBlock returns whether num represents a block number exactly at the snapshot block of an epoch.
func (c *ChainConfig) IsSnapshotBlock(num *big.Int) bool {
	if c.Dccs == nil || c.Dccs.ThangLongBlock == nil {
		return false
	}
	forBlock := new(big.Int).SetUint64(CanonicalDepth)
	forBlock.Add(num, forBlock)
	return c.Dccs.IsThangLong(forBlock) && c.Dccs.IsCheckpoint(forBlock.Uint64())
}

// IsThangLongPreparationBlock returns whether num represents a block number exactly at the ThangLong Preparation
// ThangLong preparation block is hard-coded to 32 blocks before the ThangLong hard-fork
func (c *DccsConfig) IsThangLongPreparationBlock(num *big.Int) bool {
	if c.ThangLongBlock == nil {
		return false
	}
	preparationBlock := big.NewInt(1)
	if c.ThangLongBlock.Cmp(common.Big32) > 0 {
		preparationBlock = preparationBlock.Sub(c.ThangLongBlock, common.Big32)
	}
	return preparationBlock.Cmp(num) == 0
}

// IsThangLong returns whether num represents a block number after the ThangLong fork
func (c *ChainConfig) IsThangLong(num *big.Int) bool {
	return c.Dccs != nil && c.Dccs.IsThangLong(num)
}

// IsThangLong returns whether num represents a block number after the ThangLong fork
func (c *DccsConfig) IsThangLong(num *big.Int) bool {
	return isForked(c.ThangLongBlock, num)
}

// IsCoLoa returns whether num represents a block number after the CoLoa fork
func (c *ChainConfig) IsCoLoa(num *big.Int) bool {
	return c.Dccs != nil && c.Dccs.IsCoLoa(num)
}

// IsCoLoa returns whether num represents a block number after the CoLoa fork
func (c *DccsConfig) IsCoLoa(num *big.Int) bool {
	return isForked(c.CoLoaBlock, num)
}

// CheckCompatible checks whether scheduled fork transitions have been imported
// with a mismatching chain configuration.
func (c *ChainConfig) CheckCompatible(newcfg *ChainConfig, height uint64) *ConfigCompatError {
	bhead := new(big.Int).SetUint64(height)

	// Iterate checkCompatible to find the lowest conflict.
	var lasterr *ConfigCompatError
	for {
		err := c.checkCompatible(newcfg, bhead)
		if err == nil || (lasterr != nil && err.RewindTo == lasterr.RewindTo) {
			break
		}
		lasterr = err
		bhead.SetUint64(err.RewindTo)
	}
	return lasterr
}

func (c *ChainConfig) checkCompatible(newcfg *ChainConfig, head *big.Int) *ConfigCompatError {
	if isForkIncompatible(c.HomesteadBlock, newcfg.HomesteadBlock, head) {
		return newCompatError("Homestead fork block", c.HomesteadBlock, newcfg.HomesteadBlock)
	}
	if isForkIncompatible(c.DAOForkBlock, newcfg.DAOForkBlock, head) {
		return newCompatError("DAO fork block", c.DAOForkBlock, newcfg.DAOForkBlock)
	}
	if c.IsDAOFork(head) && c.DAOForkSupport != newcfg.DAOForkSupport {
		return newCompatError("DAO fork support flag", c.DAOForkBlock, newcfg.DAOForkBlock)
	}
	if isForkIncompatible(c.EIP150Block, newcfg.EIP150Block, head) {
		return newCompatError("EIP150 fork block", c.EIP150Block, newcfg.EIP150Block)
	}
	if isForkIncompatible(c.EIP155Block, newcfg.EIP155Block, head) {
		return newCompatError("EIP155 fork block", c.EIP155Block, newcfg.EIP155Block)
	}
	if isForkIncompatible(c.EIP158Block, newcfg.EIP158Block, head) {
		return newCompatError("EIP158 fork block", c.EIP158Block, newcfg.EIP158Block)
	}
	if c.IsEIP158(head) && !configNumEqual(c.ChainID, newcfg.ChainID) {
		return newCompatError("EIP158 chain ID", c.EIP158Block, newcfg.EIP158Block)
	}
	if isForkIncompatible(c.ByzantiumBlock, newcfg.ByzantiumBlock, head) {
		return newCompatError("Byzantium fork block", c.ByzantiumBlock, newcfg.ByzantiumBlock)
	}
	if isForkIncompatible(c.ConstantinopleBlock, newcfg.ConstantinopleBlock, head) {
		return newCompatError("Constantinople fork block", c.ConstantinopleBlock, newcfg.ConstantinopleBlock)
	}
	if isForkIncompatible(c.PetersburgBlock, newcfg.PetersburgBlock, head) {
		return newCompatError("Petersburg fork block", c.PetersburgBlock, newcfg.PetersburgBlock)
	}
	if isForkIncompatible(c.IstanbulBlock, newcfg.IstanbulBlock, head) {
		return newCompatError("Istanbul fork block", c.IstanbulBlock, newcfg.IstanbulBlock)
	}
	if isForkIncompatible(c.EWASMBlock, newcfg.EWASMBlock, head) {
		return newCompatError("ewasm fork block", c.EWASMBlock, newcfg.EWASMBlock)
	}
	if c.Dccs != nil && newcfg.Dccs != nil {
		if isForkIncompatible(c.Dccs.ThangLongBlock, newcfg.Dccs.ThangLongBlock, head) {
			return newCompatError("Thang Long fork block", c.Dccs.ThangLongBlock, newcfg.Dccs.ThangLongBlock)
		}
	}
	return nil
}

// isForkIncompatible returns true if a fork scheduled at s1 cannot be rescheduled to
// block s2 because head is already past the fork.
func isForkIncompatible(s1, s2, head *big.Int) bool {
	return (isForked(s1, head) || isForked(s2, head)) && !configNumEqual(s1, s2)
}

// isForked returns whether a fork scheduled at block s is active at the given head block.
func isForked(s, head *big.Int) bool {
	if s == nil || head == nil {
		return false
	}
	return s.Cmp(head) <= 0
}

func configNumEqual(x, y *big.Int) bool {
	if x == nil {
		return y == nil
	}
	if y == nil {
		return x == nil
	}
	return x.Cmp(y) == 0
}

// ConfigCompatError is raised if the locally-stored blockchain is initialised with a
// ChainConfig that would alter the past.
type ConfigCompatError struct {
	What string
	// block numbers of the stored and new configurations
	StoredConfig, NewConfig *big.Int
	// the block number to which the local chain must be rewound to correct the error
	RewindTo uint64
}

func newCompatError(what string, storedblock, newblock *big.Int) *ConfigCompatError {
	var rew *big.Int
	switch {
	case storedblock == nil:
		rew = newblock
	case newblock == nil || storedblock.Cmp(newblock) < 0:
		rew = storedblock
	default:
		rew = newblock
	}
	err := &ConfigCompatError{what, storedblock, newblock, 0}
	if rew != nil && rew.Sign() > 0 {
		err.RewindTo = rew.Uint64() - 1
	}
	return err
}

func (err *ConfigCompatError) Error() string {
	return fmt.Sprintf("mismatching %s in database (have %d, want %d, rewindto %d)", err.What, err.StoredConfig, err.NewConfig, err.RewindTo)
}

// Rules wraps ChainConfig and is merely syntactic sugar or can be used for functions
// that do not have or require information about the block.
//
// Rules is a one time interface meaning that it shouldn't be used in between transition
// phases.
type Rules struct {
	ChainID                                                 *big.Int
	IsHomestead, IsEIP150, IsEIP155, IsEIP158               bool
	IsByzantium, IsConstantinople, IsPetersburg, IsIstanbul bool
	IsThangLong, IsCoLoa                                    bool
}

// Rules ensures c's ChainID is not nil.
func (c *ChainConfig) Rules(num *big.Int) Rules {
	chainID := c.ChainID
	if chainID == nil {
		chainID = new(big.Int)
	}
	return Rules{
		ChainID:          new(big.Int).Set(chainID),
		IsHomestead:      c.IsHomestead(num),
		IsEIP150:         c.IsEIP150(num),
		IsEIP155:         c.IsEIP155(num),
		IsEIP158:         c.IsEIP158(num),
		IsByzantium:      c.IsByzantium(num),
		IsConstantinople: c.IsConstantinople(num),
		IsPetersburg:     c.IsPetersburg(num),
		IsIstanbul:       c.IsIstanbul(num),
		IsThangLong:      c.IsThangLong(num),
		IsCoLoa:          c.IsCoLoa(num),
	}
}
