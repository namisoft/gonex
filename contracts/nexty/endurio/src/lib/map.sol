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
        uint idx = this.idxs[proposal.maker];
        require (idx == 0, "maker already has a proposal");
        idx = this.idxs[proposal.maker] = this.vals.push(proposal);
        return this.vals[idx-1];
    }

    function remove(ProposalMap storage this, address maker) internal returns (bool success) {
        uint idx = this.idxs[maker];
        if (idx == 0) {
            return false;
        }
        this.remove(idx-1);
        return true;
    }

    // position is the correct array position, which is (set.idxs[item]-1)
    function remove(ProposalMap storage this, uint position) internal {
        delete this.idxs[this.vals[position].maker];
        //this.vals[position].clear(); // should be handled differently for user and consensus

        if (this.vals.length-1 != position) {
            // swap with the last item in the vals
            this.vals[position] = this.vals[this.vals.length-1];
            this.idxs[this.vals[position].maker] = position + 1;
        }
        // remove the last item from the array
        delete this.vals[this.vals.length-1];
        this.vals.length--;
    }

    function clear(ProposalMap storage this) internal {
        for (uint i = 0; i < this.vals.length; i++) {
            delete this.idxs[this.vals[i].maker];
        }
        delete this.vals;
    }

    ///////////////////////////////////////////////////////////////////////

    // Iterable map of (address => Proposal)
    using map for AddressBool;
    struct AddressBool {
        address[] keys;
        mapping (address => uint) idxs; // map the voter's adress to (its position + 1)
        mapping (address => bool) vals;
    }

    function count(AddressBool storage this) internal view returns (uint) {
        return this.keys.length;
    }

    function getKey(AddressBool storage this, uint index) internal view returns (address) {
        return this.keys[index];
    }

    function get(AddressBool storage this, uint index) internal view returns (address, bool) {
        address key = this.getKey(index);
        return (key, this.vals[key]);
    }

    function get(AddressBool storage this, address key) internal view returns (bool) {
        return this.vals[key];
    }

    function has(AddressBool storage this, address key) internal view returns (bool) {
        return this.idxs[key] > 0;
    }

    /**
     * @return true if new (key,val) is added, false if old key is map to a new value
     */
    function set(AddressBool storage this, address key, bool val) internal returns (bool) {
        this.vals[key] = val;
        uint idx = this.idxs[key];
        if (idx == 0) {
            this.idxs[key] = this.keys.push(key);
            return true;
        }
        if (idx > this.keys.length || this.keys[idx-1] != key) {
            // storage inconsistency due to deleting proposal without clearing proposal.votes
            this.idxs[key] = this.keys.push(key);
            return true;
        }
        // key already has a proposal
        return false;
    }

    function remove(AddressBool storage this, address key) internal {
        uint idx = this.idxs[key];
        require(idx > 0, "key not exist");
        this.remove(idx-1, key);
    }

    function remove(AddressBool storage this, uint position) internal {
        address key = this.keys[position];
        require(key != address(0x0), "position not exist");
        this.remove(position, key);
    }

    // unsafe internal function
    function remove(AddressBool storage this, uint position, address key) internal {
        delete this.idxs[key];
        delete this.vals[key];

        if (this.keys.length-1 != position) {
            // swap with the last item in the keys and delete it
            this.keys[position] = this.keys[this.keys.length-1];
            this.idxs[this.keys[position]] = position + 1;
        }
        // remove the last item from the array
        delete this.keys[this.keys.length-1];
        this.keys.length--;
    }

    function clear(AddressBool storage this) internal {
        for (uint i = 0; i < this.keys.length; i++) {
            address key = this.keys[i];
            delete this.idxs[key];
            delete this.vals[key];
        }
        delete this.keys;
    }
}