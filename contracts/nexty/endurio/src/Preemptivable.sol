pragma solidity ^0.5.2;

import "./lib/util.sol";
import "./lib/set.sol";
import "./lib/map.sol";
import "./lib/absn.sol";
import "./Absorbable.sol";

/**
 * Pre-emptive absorption propsosal and voting logic.
 */
contract Preemptivable is Absorbable {
    using set for set.AddressSet;
    using map for map.ProposalMap;
    using absn for absn.Proposal;
    using absn for absn.Preemptive;

    address constant ZERO_ADDRESS = address(0x0);

    // adapting global default parameters, only used if proposal maker doesn't specify them
    uint internal globalLockdownExpiration = 2 weeks / 1 seconds;
    uint internal globalSlashingDuration = globalLockdownExpiration / 2;

    // proposal params must not lower than 1/3 of global params
    uint constant PARAM_TOLERANCE = 3;

    // proposal must have atleast globalLockdownExpiration/4 block to be voted
    // note: using globalLockdownExpiration instead of proposal's value for safety
    uint constant MIN_VOTING_DURATION = 4;

    // map (maker => Proposal)
    map.ProposalMap internal proposals;

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
        // if MNTY is received and data contains 3 params
        if (data.length == 32*3 && msg.sender == address(VolatileToken)) {
            // pre-emptive absorption proposal
            require(!proposals.has(maker), "already has a proposal");

            (   int amount,
                uint lockdownExpiration,
                uint slashingDuration
            ) = abi.decode(data, (int, uint, uint));

            propose(maker, value, amount, lockdownExpiration, slashingDuration);
            return;
        }

        // not a pre-emptive proposal, fallback to Orderbook trader order
        (uint wantAmount, bytes32 assistingID) = (data.length == 32) ?
            (abi.decode(data, (uint         )), bytes32(0)) :
            (abi.decode(data, (uint, bytes32))            );

        super.trade(maker, value, wantAmount, assistingID);
    }

    function onBlockInitialized(uint target) public consensus {
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
        uint lockdownExpiration,
        uint slashingDuration
    )
        internal
    {
        absn.Proposal memory proposal;
        proposal.maker = maker;
        proposal.stake = stake;
        proposal.amount = amount;
        proposal.number = block.number;

        if (lockdownExpiration > 0) {
            require(
                lockdownExpiration <
                globalLockdownExpiration - globalLockdownExpiration / PARAM_TOLERANCE,
                "lockdown duration param too short");
        } else {
            proposal.lockdownExpiration = globalLockdownExpiration;
        }

        if (slashingDuration > 0) {
            require(
                slashingDuration >
                globalSlashingDuration + globalSlashingDuration / PARAM_TOLERANCE,
                "slashing duration param too long");
        } else {
            proposal.slashingDuration = globalSlashingDuration;
        }

        proposals.push(proposal);
    }

    function vote(address maker, bool up) external {
        require(proposals.has(maker), "no such proposal");
        absn.Proposal storage proposal = proposals.get(maker);
        if (up) {
            proposal.voteUp();
        } else {
            proposal.voteDown();
        }
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
    function calcRank(absn.Proposal storage proposal) internal view returns (uint) {
        uint voteCount = 0;
        for (uint i = 0; i < proposal.upVoters.count(); ++i) {
            address voter = proposal.upVoters.get(i);
            voteCount += voter.balance + VolatileToken.balanceOf(voter);
        }
        for (uint i = 0; i < proposal.downVoters.count(); ++i) {
            address voter = proposal.downVoters.get(i);
            voteCount -= voter.balance + VolatileToken.balanceOf(voter);
        }
        if (voteCount <= 0) {
            return 0;
        }
        return proposal.stake * voteCount;
    }

    // expensive calculation, only consensus can affort this
    function calcBestProposal() internal view returns (address) {
        uint bestRank = 0;
        address bestMaker = ZERO_ADDRESS;
        for (uint i = 0; i < proposals.count(); ++i) {
            absn.Proposal storage proposal = proposals.get(i);
            if (block.number - proposal.number < globalLockdownExpiration / MIN_VOTING_DURATION) {
                // not enough time for voting
                continue;
            }
            uint rank = calcRank(proposal);
            if (rank > bestRank) {
                bestRank = rank;
                bestMaker = proposal.maker;
            }
        }
        return bestMaker;
    }
}