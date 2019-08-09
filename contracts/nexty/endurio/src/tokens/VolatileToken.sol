pragma solidity ^0.5.2;

import "openzeppelin-solidity/contracts/math/SafeMath.sol";
import "./ERC223.sol";

/*
    . Exchanged with NTY with rate 1 MegaNTY = 1000000 NTY
    . Mint. / burn. able(free) by owner = orderbook contract
*/

contract VolatileToken is ERC223 {
    string public constant symbol = "MNTY";
    string public constant name = "Mega NTY";
    uint public constant decimals = 24;

    constructor (
        address orderbook,      // mandatory
        address prefundAddress, // optional
        uint prefundAmount      // optional
    )
        public
    {
        if (prefundAmount > 0 ) {
            _mint(prefundAddress, prefundAmount * 10**decimals);
        }
        initialize(orderbook);
    }

    // deposit (MNTY <- NTY)
    function deposit()
        external
        payable
        returns(bool)
    {
        depositTo(msg.sender);
    }

    // withdraw (MNTY -> NTY)
    function withdraw(uint _amount)
        external
        returns(bool)
    {
        withdrawTo(_amount, msg.sender);
    }

    // withdrawTo (MNTY -> NTY -> address)
    function withdrawTo(uint _amount, address payable _to)
        public
        returns(bool)
    {
        address _sender = msg.sender;
        _burn(_sender, _amount);

        /************************************************************************/
        /* concensus garantures, this contract always got enough NTY to withdraw */
        /************************************************************************/

        _to.transfer(_amount);
    }

    // depositTo (addresss <- MNTY <- NTY)
    function depositTo(
        address _to
    )
        public
        payable
        returns(bool)
    {
        uint _amount = msg.value;
        _mint(_to, _amount);
        return true;
    }

    // deposit and order (NTY -> MNTY -> USD)
    function depositAndTrade(
        bytes32 index,
        uint haveAmount,
        uint wantAmount,
        bytes32 assistingID
    )
        public
        payable
    {
        depositTo(msg.sender);
        trade(index, haveAmount, wantAmount, assistingID);
    }

    // create selling order (MNTY -> USD)
    // with verbose data = (wantAmount, assistingID)
    function trade(
        bytes32 index,
        uint haveAmount,
        uint wantAmount,
        bytes32 assistingID
    )
        public
    {
        bytes memory data = abi.encode(index, wantAmount, assistingID);
        transfer(dex(), haveAmount, data);
    }

    // deposit and propose()
    function depositAndPropose(
        uint stake,              // staked amount of VolatileToken
        int amount,             // absorption amount of StablizeToken
        uint slashingDuration,
        uint lockdownExpiration
    )
        public
        payable
    {
        depositTo(msg.sender);
        propose(stake, amount, slashingDuration, lockdownExpiration);
    }

    // propose a new pre-emptive absorption
    // with verbose data = (amount, slashingDuration, lockdownExpiration);
    function propose(
        uint stake,              // staked amount of VolatileToken
        int amount,             // absorption amount of StablizeToken
        uint slashingDuration,
        uint lockdownExpiration
    )
        public
    {
        bytes memory data = abi.encode(amount, slashingDuration, lockdownExpiration, bytes32(0));
        transfer(dex(), stake, data);
    }
}