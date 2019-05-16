pragma solidity ^0.5.2;

contract DataSet {
    struct Order {
        address maker;
        uint256 haveAmount;
        uint256 wantAmount;

        // linked list
        bytes32 prev;
        bytes32 next;
    }

    struct OrderList {
        mapping (bytes32 => Order) orders;
        // bytes32 top;	// the highest priority (lowest sell or highest buy)
        // bytes32 bottom;	// the lowest priority (highest sell or lowest buy)
    }

    mapping(bool => OrderList) internal books;
    // contract private nonce ot generate unique ids
    mapping(address => uint256) internal pNonce;
    bytes32 constant zeroBytes32 = bytes32(0);
}