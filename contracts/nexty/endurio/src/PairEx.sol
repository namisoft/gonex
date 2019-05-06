pragma solidity ^0.5.2;

import "openzeppelin-solidity/contracts/math/Math.sol";
import "openzeppelin-solidity/contracts/math/SafeMath.sol";
import "./OrderBook.sol";

contract PairEx is OrderBook {

    constructor (
        address _volatileTokenAddress,
        address _stableTokenAddress
    )
        public
    {
        Order memory order;
        order.wantAmount = 1;
        books[false].orders[bytes32(0)] = order;
        books[true].orders[bytes32(0)] = order;
        volatileTokenRegister(_volatileTokenAddress);
        stableTokenRegister(_stableTokenAddress);
    }

    function getOrderType()
        public
        view
        returns(bool)
    {
        address _sender = msg.sender;
        require(_sender == address(token[true]) || _sender == address(token[false]), "only VolatileToken and StableToken accepted");
        return _sender == address(token[true]);
    }

    // Token transfer's fallback
    // bytes _data = uint256[2] = (wantAmount, assistingID)
    // RULE : delegateCall never used
    function tokenFallback(
        address _from,
        uint _value,
        bytes calldata _data) 
        external
    {
        address maker = _from;
        bool orderType = getOrderType();
        uint256 haveAmount = _value;
        uint256 wantAmount;
        bytes32 assistingID;
        (wantAmount, assistingID) = _data.length == 32 ? (abi.decode(_data, (uint256)), bytes32(0)) : abi.decode(_data, (uint256, bytes32));
        bytes32 _orderID = insert(
            orderType,
            haveAmount,
            wantAmount,
            maker,
            assistingID
        );
        pairing(orderType, _orderID);
    }

    function pairing(
        bool _orderType,
        bytes32 _orderID
    )
        private
    {
        bool _redroType = !_orderType;
        OrderList storage orderBook = books[_orderType];
        Order storage order = orderBook.orders[_orderID];
        OrderList storage redroBook = books[_redroType];
        bytes32 redroTopID = top(_redroType);

        while (redroTopID[0] != 0) {
            Order storage redro = redroBook.orders[redroTopID];
            if (order.haveAmount.mul(redro.haveAmount) < order.wantAmount.mul(redro.wantAmount)) {
                // not pairable
                return;
            }
            uint256 orderPairableAmount = Math.min(order.haveAmount, redro.wantAmount);
            order.wantAmount = order.wantAmount.mul(order.haveAmount.sub(orderPairableAmount)).div(order.haveAmount);
            order.haveAmount = order.haveAmount.sub(orderPairableAmount);

            uint256 redroPairableAmount = redro.haveAmount.mul(orderPairableAmount).div(redro.wantAmount);
            redro.wantAmount = redro.wantAmount.mul(redro.haveAmount.sub(redroPairableAmount)).div(redro.haveAmount);
            redro.haveAmount = redro.haveAmount.sub(redroPairableAmount);

            token[_redroType].transfer(order.maker, redroPairableAmount);
            token[_orderType].transfer(redro.maker, orderPairableAmount);
            if (redro.haveAmount == 0 || redro.wantAmount == 0) {
                _remove(_redroType, redroTopID);
            }
            redroTopID = top(_redroType);
            if (order.haveAmount == 0 || order.wantAmount == 0) {
                _remove(_orderType, _orderID);
                return;
            }
        }
        return;
    }

    // orderToFill(BuyType, 123) returns an SellType order to fill enough buying orders to burn as much as 123 StableToken.
    // orderToFill(SellType, 456) returns an BuyType order to fill enough selling orders to create as much as 456 StableToken.
    function orderToFill(
        bool _orderType,
        uint256 _stableTokenTarget
    )
        public
        view
        returns(uint256, uint256)
    {
        OrderList storage book = books[_orderType];
        uint256 totalSTB;
        uint256 totalVOL;
        bytes32 cursor = top(_orderType);
        while(cursor[0] != 0 && totalSTB < _stableTokenTarget) {
            Order storage order = book.orders[cursor];
            uint256 stb = _orderType ? order.haveAmount : order.wantAmount;
            uint256 vol = _orderType ? order.wantAmount : order.haveAmount;
            // break-point
            if (totalSTB.add(stb) > _stableTokenTarget) {
                uint256 remainSTB = _stableTokenTarget.sub(totalSTB);
                uint256 remainVOL = vol.mul(remainSTB).div(stb);
                return (totalVOL.add(remainVOL), totalSTB.add(remainSTB));
            }
            totalVOL = totalVOL.add(vol);
            totalSTB = totalSTB.add(stb);
            cursor = order.next;
        }
        //  Not enough order, return all we have
        return (totalVOL, totalSTB);
    }

}