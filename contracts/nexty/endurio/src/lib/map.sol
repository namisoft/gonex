pragma solidity ^0.5.2;

import "./absn.sol";

library map {
    using absn for absn.Proposal;

    // Iterable map of (address => Proposal)
    using map for ProposalMap;
    struct ProposalMap {
        absn.Proposal[] vals;
        mapping (address => uint) idxs; // map the proposal's maker to (its position + 1)
    }

    function count(ProposalMap storage this) internal view returns (uint) {
        return this.vals.length;
    }

    function get(ProposalMap storage this, uint index) internal view returns (absn.Proposal storage) {
        return this.vals[index];
    }

    function get(ProposalMap storage this, address maker) internal view returns (absn.Proposal storage) {
        return this.vals[this.idxs[maker]-1];
    }

    function has(ProposalMap storage this, address maker) internal view returns (bool) {
        return this.idxs[maker] > 0;
    }

    function push(ProposalMap storage this, absn.Proposal memory proposal) internal returns (absn.Proposal storage) {
        uint position = this.idxs[proposal.maker];
        require (position == 0, "maker already has a proposal");
        position = this.idxs[proposal.maker] = this.vals.push(proposal);
        return this.vals[position-1];
    }

    function remove(ProposalMap storage this, address maker) internal returns (bool success) {
        uint position = this.idxs[maker];
        if (position == 0) {
            return false;
        }
        this.remove(position-1);
        return true;
    }

    // position is the correct array position, which is (set.idxs[item]-1)
    function remove(ProposalMap storage this, uint position) internal {
        delete this.idxs[this.vals[position].maker]; // delete from the idxs
        this.vals[position] = this.vals[this.vals.length-1]; // swap with the last item in the vals
        delete this.vals[this.vals.length-1]; // .. and delete the last item
        this.vals.length--;
    }
}