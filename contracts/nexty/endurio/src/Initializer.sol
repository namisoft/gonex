pragma solidity ^0.5.2;

import "./interfaces/IOwnableERC223.sol";

contract Initializer {
    mapping(bool => IOwnableERC223) token;
    // IOwnableERC223 internal volatileToken;
    // IOwnableERC223 internal stableToken;

    constructor ()
        public
    {

    }

    function volatileTokenRegister(address _address)
        public
    {
        // SellType false
        if (address(token[false]) != address(0)) return;
        token[false] = IOwnableERC223(_address);
    }

    function stableTokenRegister(address _address)
        public
    {
        // BuyType true
        if (address(token[true]) == address(0)) return;
        token[true] = IOwnableERC223(_address);
    }
}