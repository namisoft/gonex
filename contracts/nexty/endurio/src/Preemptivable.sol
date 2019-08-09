pragma solidity ^0.5.2;

import "./lib/util.sol";
import "./lib/map.sol";
import "./lib/absn.sol";
import "./Absorbable.sol";

/**
 * Pre-emptive absorption propsosal and voting logic.
 */
contract Preemptivable is Absorbable {
    using map for map.ProposalMap;
    using map for map.AddressBool;
    using absn for absn.Proposal;
    using absn for absn.Preemptive;

    address constant ZERO_ADDRESS = address(0x0);

    // adapting global default parameters, only used if proposal maker doesn't specify them
    uint internal globalLockdownExpiration = 2 weeks / 1 seconds;
    uint internal globalSlashingDuration = globalLockdownExpiration / 2;

    // adapting global requirement
    uint internal globalSuccessRank = 0;

    // proposal params must not lower than 1/3 of global params
    uint constant PARAM_TOLERANCE = 3;

    // proposal must have atleast globalLockdownExpiration/4 block to be voted
    // note: using globalLockdownExpiration instead of proposal's value for safety
    uint constant MIN_VOTING_DURATION = 4;

    // map (maker => Proposal)
    map.ProposalMap internal proposals;

    // revoked proposals' votes to clear by consensus
    // since the clearing job is too expensive for transaction to perform.
    map.AddressBool[] votesToClear;

    constructor (
        uint absorptionDuration,
        uint absorptionExpiration,
        uint initialSlashingDuration,
        uint initialLockdownExpiration
    )
        Absorbable(
            absorptionDuration,
            absorptionExpiration
        )
        public
    {
        if (initialLockdownExpiration > 0) {
            globalLockdownExpiration = initialLockdownExpiration;
        }
        globalSlashingDuration = initialSlashingDuration > 0 ?
            initialSlashingDuration : globalLockdownExpiration / 2;
    }

    // Token transfer's fallback
    // bytes _data = uint[2] = (wantAmount, assistingID)
    // RULE : delegateCall never used
    //
    // buy/sell order is created by sending token to this address,
    // with extra data = (wantAmount, assistingID)
    function tokenFallback(
        address maker,  // actual tx sender
        uint value,     // amount of ERC223(msg.sender) received
        bytes calldata data)
        external
    {
        // if MNTY is received and data contains 4 params
        if (data.length == 32*4 && msg.sender == address(VolatileToken)) {
            // pre-emptive absorption proposal
            require(!proposals.has(maker), "already has a proposal");

            (   int amount,
                uint slashingDuration,
                uint lockdownExpiration,
                bytes32 reserve // reserve params to distinguish proposal and trading request
            ) = abi.decode(data, (int, uint, uint, bytes32));

            // unused
            reserve = bytes32(0);

            propose(maker, value, amount, slashingDuration, lockdownExpiration);
            return;
        }

        // not a pre-emptive proposal, fallback to Orderbook trader order
        bytes32 index;
        uint wantAmount;
        bytes32 assistingID;
        if (data.length == 32*3) {
            (index, wantAmount, assistingID) = abi.decode(data, (bytes32, uint, bytes32));
        } else {
            (index, wantAmount) = abi.decode(data, (bytes32, uint));
        }

        super.trade(maker, index, value, wantAmount, assistingID);
    }

    function onBlockInitialized(uint target) public consensus {
        // cleaning up
        for (uint i = 0; i < votesToClear.length; i++) {
            votesToClear[i].clear();
        }
        delete votesToClear;

        checkAndTriggerPreemptive();
        super.onBlockInitialized(target);
    }

    /**
     * propose allows Preemptive initiator to lock their MNTY in and
     * introduces new proposal.
     */
    function propose(
        address maker,
        uint stake,
        int amount,
        uint slashingDuration,
        uint lockdownExpiration
    )
        internal
    {
        absn.Proposal memory proposal;
        proposal.maker = maker;
        proposal.stake = stake;
        proposal.amount = amount;
        proposal.number = block.number;

        if (slashingDuration > 0) {
            require(
                slashingDuration >
                globalSlashingDuration + globalSlashingDuration / PARAM_TOLERANCE,
                "slashing duration param too long");
        } else {
            proposal.slashingDuration = globalSlashingDuration;
        }

        if (lockdownExpiration > 0) {
            require(
                lockdownExpiration <
                globalLockdownExpiration - globalLockdownExpiration / PARAM_TOLERANCE,
                "lockdown duration param too short");
        } else {
            proposal.lockdownExpiration = globalLockdownExpiration;
        }

        proposals.push(proposal);
    }

    function revoke(address maker) external {
        absn.Proposal storage p = proposals.get(maker);
        require(maker == p.maker, "only maker can revoke proposal");
        votesToClear.push(p.votes); // leave the job for consensus
        VolatileToken.transfer(p.maker, p.stake);
        proposals.remove(maker);
    }

    function vote(address maker, bool up) external {
        require(proposals.has(maker), "no such proposal");
        absn.Proposal storage proposal = proposals.get(maker);
        proposal.vote(up);
    }

    // check and trigger a new Preemptive when one is eligible
    // return the true if a new preemptive is activated
    function checkAndTriggerPreemptive() internal returns (bool) {
        if (lockdown.isLocked()) {
            // there's current active or lockdown absorption
            return false;
        }
        address bestMaker = calcBestProposal();
        if (bestMaker == ZERO_ADDRESS) {
            // no eligible proposals
            return false;
        }
        triggerPreemptive(bestMaker);
        return true;
    }

    // trigger an absorption from a maker's proposal
    function triggerPreemptive(address maker) internal {
        absn.Proposal storage proposal = proposals.get(maker);
        proposal.votes.clear(); // clear the votes (consensus only)
        lockdown = absn.Preemptive(
            proposal.maker,
            proposal.amount,
            proposal.stake,
            proposal.slashingDuration,
            block.number + proposal.lockdownExpiration
        );
        proposals.remove(maker);
        triggerAbsorption(util.add(StablizeToken.totalSupply(), lockdown.amount), true);
    }

    // expensive calculation, only consensus can affort this
    function calcRank(absn.Proposal storage proposal) internal view returns (int) {
        int voteCount = countVote(proposal);
        if (voteCount <= 0) {
            return 0;
        }
        return int(proposal.stake) * countVote(proposal);
    }

    // expensive calculation, only consensus can affort this
    function countVote(absn.Proposal storage proposal) internal view returns(int) {
        int voteCount = 0;
        for (uint i = 0; i < proposal.votes.count(); ++i) {
            (address voter, bool up) = proposal.votes.get(i);
            int weight = int(voter.balance + VolatileToken.balanceOf(voter));
            if (up) {
                voteCount += weight;
            } else {
                voteCount -= weight;
            }
        }
        return voteCount;
    }

    // expensive calculation, only consensus can affort this
    function calcBestProposal() internal view returns(address) {
        int bestRank = 0;
        address bestMaker = ZERO_ADDRESS;
        for (uint i = 0; i < proposals.count(); ++i) {
            absn.Proposal storage proposal = proposals.get(i);
            if (block.number - proposal.number < globalLockdownExpiration / MIN_VOTING_DURATION) {
                // not enough time for voting
                continue;
            }
            int rank = calcRank(proposal);
            if (rank > bestRank) {
                bestRank = rank;
                bestMaker = proposal.maker;
            }
        }
        return bestMaker;
    }

    function totalVote(address maker) public view returns(int) {
        absn.Proposal storage p = proposals.get(maker);
        return countVote(p);
    }

    function getProposalCount() public view returns(uint) {
        return proposals.count();
    }

    function getProposal(uint idx) public view
        returns (
            address maker,
            uint stake,
            int amount,
            uint slashingDuration,
            uint lockdownExpiration,
            uint number
        )
    {
        absn.Proposal storage p = proposals.get(idx);
        return (p.maker, p.stake, p.amount, p.slashingDuration, p.lockdownExpiration, p.number);
    }
}