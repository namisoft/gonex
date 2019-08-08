pragma solidity ^0.5.2;

library set {
    using set for AddressSet;

    // Iterable set of address
    struct AddressSet {
        address[] vals;
        mapping (address => uint) idxs; // idxs an address to (its position + 1)
    }

    function count(AddressSet storage this) internal view returns (uint) {
        return this.vals.length;
    }

    function get(AddressSet storage this, uint index) internal view returns (address) {
        return this.vals[index];
    }

    function push(AddressSet storage this, address item) internal returns (bool success) {
        if (this.idxs[item] > 0) {
            return false;
        }
        this.idxs[item] = this.vals.push(item);
        return true;
    }

    function remove(AddressSet storage this, address item) internal returns (bool success) {
        uint position = this.idxs[item];
        if (position == 0) {
            return false;
        }
        this.remove(position-1);
        return true;
    }

    // position is the correct array position, which is (set.idxs[item]-1)
    function remove(AddressSet storage this, uint position) internal {
        delete this.idxs[this.vals[position]]; // delete from the idxs
        this.vals[position] = this.vals[this.vals.length-1]; // swap with the last item in the vals
        delete this.vals[this.vals.length-1]; // .. and delete the last item
        this.vals.length--;
    }
}