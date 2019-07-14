pragma solidity ^0.5.2;

import "./interfaces/IOwnableERC223.sol";

contract Initializer {
    mapping(bool => IOwnableERC223) token;

    bool public constant Volatile = false;
    bool public constant Stable = true;
    bool public constant Sell = false;
    bool public constant Buy = true;

    constructor ()
        public
    {

    }

    function volatileTokenRegister(address _address)
        public
    {
        // SellType false
        require(address(token[Volatile]) == address(0), "already set");
        token[Volatile] = IOwnableERC223(_address);
    }

    function stableTokenRegister(address _address)
        public
    {
        // BuyType true
        require(address(token[Stable]) == address(0), "already set");
        token[Stable] = IOwnableERC223(_address);
    }
}