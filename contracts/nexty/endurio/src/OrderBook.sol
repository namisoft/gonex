pragma solidity ^0.5.2;

import "openzeppelin-solidity/contracts/math/SafeMath.sol";
import "./DataSet.sol";
import "./Initializer.sol";

contract OrderBook is Initializer, DataSet {
    using SafeMath for uint256;
    uint256 constant public INPUTS_MAX = 2 ** 120;
    // Stepping price param
    uint256 internal StepDividend = 0;
    uint256 internal StepDivisor = 1;

    function insert(
        bool _orderType,
        uint256 _haveAmount,
        uint256 _wantAmount,
        address _maker,
        bytes32 _assistingID
        )
        public
 	    returns (bytes32)
    {
        require(validOrder(_orderType, _assistingID), "save your gas");
        OrderList storage book = books[_orderType];
        bytes32 newID = initOrder(_maker, _orderType, _haveAmount, _wantAmount);

        // direction to bottom, search first order, that new order better than
        bytes32 id = _assistingID == zeroBytes32 ? top(_orderType) : _assistingID;
        if (!betterOrder(_orderType, newID, id)) {
            // order[newID] always better than order[bytes32(0)] with price = 0
            // if price of order[newID] = 0 => throw cause infinite loop
            while (!betterOrder(_orderType, newID, id)) {
                id = book.orders[id].next;
            }
            insertBefore(_orderType, newID, id);
            return newID;
        }
        id = book.orders[id].prev;
        // direction to top, search first order that new order not better than
        // this part triggered only if new order not better than assistingID order
        while (id != zeroBytes32 && betterOrder(_orderType, newID, id)) {
            id = book.orders[id].prev;
        }
        insertBefore(_orderType, newID, book.orders[id].next);
        return newID;
    }

    // inseter _id as prev element of _next
    function insertBefore(
        bool _orderType,
        bytes32 _id,
        bytes32 _next
    )
        private
    {
        // _prev => _id => _next
        OrderList storage book = books[_orderType];
        bytes32 _prev = book.orders[_next].prev;
        book.orders[_id].prev = _prev;
        book.orders[_id].next = _next;
        book.orders[_next].prev = _id;
        book.orders[_prev].next = _id;
    }

    function validOrder(
        bool _orderType,
        bytes32 _id
    )
        public
        view
        returns(bool)
    {
        OrderList storage book = books[_orderType];
        Order storage _order = book.orders[_id];
        return _order.wantAmount > 0;
    }

    function initOrder(
        address _maker,
        bool _orderType,
        uint256 _haveAmount,
        uint256 _wantAmount
    )
        public
        returns (bytes32)
    {
        require(_haveAmount > 0 && _wantAmount > 0, "save your time");
        require((_haveAmount < INPUTS_MAX) && (_wantAmount < INPUTS_MAX), "greater than supply?");
        pNonce[_maker]++;
        OrderList storage book = books[_orderType];
        bytes32 id = sha256(abi.encodePacked(_maker, pNonce[_maker], _haveAmount, _wantAmount));
        book.orders[id] = Order(_maker, _haveAmount, _wantAmount, 0, 0);
        return id;
    }


    // Remove order and payout or refund
    function _remove(
        bool _orderType,
        bytes32 _id,
        bool payout
    )
        internal
        returns (bytes32)
    {
        OrderList storage book = books[_orderType];
        Order storage order = book.orders[_id];
        // before: prev => order => next
        // after:  prev ==========> next
        book.orders[order.prev].next = order.next;
        book.orders[order.next].prev = order.prev;
        if (payout) { // order is filled
            token[!_orderType].transfer(order.maker, order.wantAmount);
        } else { // order is refunded
            token[_orderType].transfer(order.maker, order.haveAmount);
        }
        bytes32 next = order.next;
        delete book.orders[_id];
        return next;
    }

    // Cancel and refund the remaining order.haveAmount
    function cancel(bool _orderType, bytes32 _id) public {
        OrderList storage book = books[_orderType];
        Order storage order = book.orders[_id];
        require(msg.sender == order.maker, "only order owner");
        _remove(_orderType, _id, false);
    }

    function _setStep(uint256 dividend, uint256 divisor) private {
        StepDividend = dividend;
        StepDivisor = divisor;
    }

    // read functions
    function top(
        bool _orderType
    )
        public
        view
        returns (bytes32)
    {
        OrderList storage book = books[_orderType];
        return book.orders[bytes32(0)].next;
    }

    function bottom(
        bool _orderType
    )
        public
        view
        returns (bytes32)
    {
        OrderList storage book = books[_orderType];
        return book.orders[bytes32(0)].prev;
    }

    function betterOrder(
        bool orderType,
        bytes32 _newId,
        bytes32 _oldId
    )
        public
        view
        returns (bool)
    {
        OrderList storage book = books[orderType];
        Order storage _new = book.orders[_newId];
        Order storage _old = book.orders[_oldId];
        // stepping price
        // newWant / newHave < (oldWant / oldHave) * (10000 / (10000 + T))
        // uint256 a = _new.haveAmount.mul(_old.wantAmount).div(StepDivisor + StepDividend);
        // uint256 b = _old.haveAmount.mul(_new.wantAmount).div(StepDivisor);
        uint256 a = _new.haveAmount.mul(_old.wantAmount).mul(StepDivisor);
        uint256 b = _old.haveAmount.mul(_new.wantAmount).mul(StepDivisor + StepDividend);
        return a > b;
    }

    function getNext(
        bool _orderType,
        bytes32 _id
    )
        public
        view
        returns (bytes32)
    {
        OrderList storage book = books[_orderType];
        return book.orders[_id].next;
    }

    function getPrev(
        bool _orderType,
        bytes32 _id
    )
        public
        view
        returns (bytes32)
    {
        OrderList storage book = books[_orderType];
        return book.orders[_id].prev;
    }

    function getOrder(
        bool _orderType,
        bytes32 _id
    )
        public
        view
        returns (
            address,
            uint256,
            uint256,
            bytes32,
            bytes32
        )
    {
        OrderList storage book = books[_orderType];
        Order storage order = book.orders[_id];
        return (order.maker, order.haveAmount, order.wantAmount, order.prev, order.next);
    }
}