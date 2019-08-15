pragma solidity ^0.5.2;

import "./util.sol";
import "./map.sol";

library absn {
    using map for map.AddressBool;

    address constant ZERO_ADDRESS = address(0x0);

    using absn for Absorption;
    struct Absorption {
        uint deadline; // the last block the absorption is valid
        uint supply;   // StablizeToken supply at the activation block
        uint target;   // StablizeToken target supply for the active
        bool isPreemptive;
    }

    function exists(Absorption storage this) internal view returns(bool) {
        return 0 < this.deadline;
    }

    function isExpired(Absorption storage this) internal view returns(bool) {
        return this.exists() && this.deadline < block.number;
    }

    function isAbsorbing(Absorption storage this, uint supply) internal view returns(bool) {
        return this.exists() &&
            supply != this.target &&                        // target not reached &&
            util.inOrder(this.supply, supply, this.target); // not over-absorbed
    }

    using absn for Proposal;
    struct Proposal {
        // address of the proposer
        address maker;

        // NTY amount staked for the preemptive proposal
        uint stake;

        // amount of StablizeToken to absorb, positive for inflation, negative for deflation
        int amount;

        // lockdown duration (in blocks from the activation)
        uint lockdownExpiration;

        // SlashingDuration = 1 / SlashingRate
        // slashed = |d/D|*SlashingRate
        //  d = MedianPriceDeviation
        //  D = X/S, X is the amount of StablizeToken will be absorbed, S is the current NewSD total supply
        uint slashingDuration;

        // block number the proposal is proposed
        uint number;

        // voters map
        map.AddressBool votes;
    }

    function vote(Proposal storage this, bool up) internal {
        this.votes.set(msg.sender, up);
    }

    function exists(Proposal storage this) internal view returns (bool) {
        return (this.maker != ZERO_ADDRESS);
    }

    function clear(Proposal storage this) internal {
        this.votes.clear();
    }

    struct Preemptive {
        address maker;
        int amount;
        uint stake;
        uint slashingDuration;

        // block number the lockdown will end
        uint unlockNumber;
    }
    using absn for Preemptive;

    function exists(Preemptive storage this) internal view returns (bool) {
        return this.maker != ZERO_ADDRESS;
    }

    function isLocked(Preemptive storage this) internal view returns (bool) {
        return this.exists() && block.number < this.unlockNumber;
    }

    function unlockable(Preemptive storage this) internal view returns (bool) {
        return this.exists() && this.unlockNumber <= block.number;
    }
}
