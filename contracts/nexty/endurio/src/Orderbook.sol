pragma solidity ^0.5.2;

import "openzeppelin-solidity/contracts/math/Math.sol";
import "openzeppelin-solidity/contracts/math/SafeMath.sol";
import "./lib/map.sol";
import "./lib/dex.sol";
import "./lib/absn.sol";

/**
 * Contract code for Volatile/Stablize exchange
 */
contract Orderbook {
    using SafeMath for uint;
    using dex for dex.Order;
    using dex for dex.Book;

    bool public constant Ask = false;
    bool public constant Bid = true;

    // TODO: mapping (hash(haveTokenAddres,wantTokenAddress) => dex.Book)
    mapping(bool => dex.Book) internal books;

    function registerTokens(
        address volatileToken,
        address stablizeToken
    )
        public
    {
        books[Ask].init(IToken(volatileToken), IToken(stablizeToken));
        books[Bid].init(IToken(stablizeToken), IToken(volatileToken));
    }

    function trade(
        address maker,
        bytes32 index,
        uint haveAmount,
        uint wantAmount,
        bytes32 assistingID
    )
        internal
    {
        // TODO: get order type from ripem160(haveToken)[:16] + ripem160(wantToken)[:16]
        dex.Book storage book = bookHave(msg.sender);

        bytes32 newID = book.createOrder(maker, index, haveAmount, wantAmount);
        book.place(newID, assistingID);
        book.fill(newID, bookWant(msg.sender));
    }

    // iterator
    function top(
        bool orderType
    )
        public
        view
        returns (bytes32)
    {
        dex.Book storage book = books[orderType];
        return book.topID();
    }

    // iterator
    function next(
        bool orderType,
        bytes32 id
    )
        public
        view
        returns (bytes32)
    {
        dex.Book storage book = books[orderType];
        return book.orders[id].next;
    }

    // iterator
    function prev(
        bool orderType,
        bytes32 id
    )
        public
        view
        returns (bytes32)
    {
        dex.Book storage book = books[orderType];
        return book.orders[id].prev;
    }

    function calcOrderID(
        address maker,
        bytes32 index
    )
        external
        pure
        returns (bytes32)
    {
        return dex._calcID(maker, index);
    }

    function getOrder(
        bool _orderType,
        bytes32 _id
    )
        public
        view
        returns (
            address,
            uint,
            uint,
            bytes32,
            bytes32
        )
    {
        dex.Book storage book = books[_orderType];
        dex.Order storage order = book.orders[_id];
        return (order.maker, order.haveAmount, order.wantAmount, order.prev, order.next);
    }

    // find the next assisting id for an order
    function findAssistingID(
        bool orderType,
        address maker,
        uint haveAmount,
        uint wantAmount,
        bytes32 assistingID
    )
        public
        view
 	    returns (bytes32)
    {
        dex.Book storage book = books[orderType];
        dex.Order memory newOrder = dex.Order(
            maker,
            haveAmount,
            wantAmount,
            bytes32(0),
            bytes32(0));
        return book.m_find(newOrder, assistingID);
    }

    function bookHave(
        address haveToken
    )
        internal
        view
        returns(dex.Book storage)
    {
        if (haveToken == address(books[false].haveToken)) {
            return books[false];
        }
        if (haveToken == address(books[true].haveToken)) {
            return books[true];
        }
        revert("no order book for token");
    }

    function bookWant(
        address wantToken
    )
        internal
        view
        returns(dex.Book storage)
    {
        if (wantToken == address(books[false].wantToken)) {
            return books[false];
        }
        if (wantToken == address(books[true].wantToken)) {
            return books[true];
        }
        revert("no order book for token");
    }

    // Cancel and refund the remaining order.haveAmount
    function cancel(bool orderType, bytes32 id) public {
        dex.Book storage book = books[orderType];
        dex.Order storage order = book.orders[id];
        require(msg.sender == order.maker, "only order maker");
        book.refund(id);
    }
}