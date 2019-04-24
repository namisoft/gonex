pragma solidity ^0.5.2;

import "openzeppelin-solidity/contracts/math/SafeMath.sol";
import "./ERC223.sol";
import "../lib/BytesConvert.sol";
import "../interfaces/IPairEx.sol";

/*
    ...
*/

contract StableToken is ERC223{
    string public constant name = "Nexty USD";
    string public constant symbol = "NUSD";
    uint8 public constant decimals = 18;

    using BytesConvert for *;

    IPairEx internal orderbook;

    constructor (address _orderbook)
        public
    {
        orderbook = IPairEx(_orderbook);
        orderbook.stableTokenRegister();
        initialize(address(_orderbook));
        _mint(msg.sender, 10**24);
    }

    function simpleBuy(
        uint256  _value,
        uint256 _wantAmount,
        bytes32 _assistingID
    ) 
        public 
        payable 
    {
        bytes memory data = abi.encode(_wantAmount, _assistingID);
        transfer(owner(), _value, data);
    }
}