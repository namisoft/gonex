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

    function volatileTokenRegister()
        public
    {
        // SellType false
        require(address(token[false]) == address(0), "already set");
        token[false] = IOwnableERC223(msg.sender);
    }

    function stableTokenRegister()
        public
    {
        // BuyType true
        require(address(token[true]) == address(0), "already set");
        token[true] = IOwnableERC223(msg.sender);
    }
}