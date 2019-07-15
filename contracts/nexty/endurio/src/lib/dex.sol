pragma solidity ^0.5.2;

import "openzeppelin-solidity/contracts/math/SafeMath.sol";
import "../interfaces/IToken.sol";

/**
 * Library for token pair exchange.
 *
 * Has no knownledge about what the token does. Any logic deal with stable or
 * volatile nature of the token must put in the contract level.
 */
library dex {
    using SafeMath for uint;

    bytes32 constant ZERO_ID = bytes32(0x0);
    bytes32 constant LAST_ID = bytes32(0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF);
    address constant ZERO_ADDRESS = address(0x0);
    uint constant INPUTS_MAX = 2 ** 128;

    function calcID(
        address maker,
        uint haveAmount,
        uint wantAmount
    )
        internal
        pure
        returns (bytes32)
    {
        return sha256(abi.encodePacked(maker, haveAmount, wantAmount));
    }

    struct Order {
        address maker;
        uint haveAmount;
        uint wantAmount;

        // linked list
        bytes32 prev;
        bytes32 next;
    }
    using dex for Order;

    function exists(Order storage order) internal view returns(bool)
    {
        // including meta orders [0] and [FF..FF]
        return order.maker != ZERO_ADDRESS;
    }

    function isEmpty(Order storage order) internal view returns(bool)
    {
        return order.haveAmount == 0 || order.wantAmount == 0;
    }

    function betterThan(Order storage o, Order storage p)
        internal
        view
        returns (bool)
    {
        return o.haveAmount * p.wantAmount > p.haveAmount * o.wantAmount;
    }

    // memory version of betterThan
    function m_betterThan(Order memory o, Order storage p)
        internal
        view
        returns (bool)
    {
        return o.haveAmount * p.wantAmount > p.haveAmount * o.wantAmount;
    }

    function fillableBy(Order memory o, Order storage p)
        internal
        view
        returns (bool)
    {
        return o.haveAmount * p.haveAmount >= p.wantAmount * o.wantAmount;
    }


    struct Book {
        IToken haveToken;
        IToken wantToken;
        mapping (bytes32 => Order) orders;
        // bytes32 top;	// the highest priority (lowest sell or highest buy)
        // bytes32 bottom;	// the lowest priority (highest sell or lowest buy)
    }
    using dex for Book;

    function init(
        Book storage book,
        IToken haveToken,
        IToken wantToken
    )
        internal
    {
        book.haveToken = haveToken;
        book.wantToken = wantToken;
        book.orders[ZERO_ID] = Order(address(this), 0, 0, ZERO_ID, LAST_ID); // [0] meta order
        book.orders[LAST_ID] = Order(address(this), 0, 1, ZERO_ID, LAST_ID); // worst order meta
    }

    // read functions
    function topID(
        Book storage book
    )
        internal
        view
        returns (bytes32)
    {
        return book.orders[ZERO_ID].next;
    }

    function bottomID(
        Book storage book
    )
        internal
        view
        returns (bytes32)
    {
        return book.orders[LAST_ID].prev;
    }

    function createOrder(
        Book storage book,
        address _maker,
        uint _haveAmount,
        uint _wantAmount
    )
        internal
        returns (bytes32)
    {
        // TODO move require check to API
        require(_haveAmount > 0 && _wantAmount > 0, "zero input");
        require(_haveAmount < INPUTS_MAX && _wantAmount < INPUTS_MAX, "input over limit");
        bytes32 id = calcID(_maker, _haveAmount, _wantAmount);
        Order storage order = book.orders[id];
        if (!order.exists()) {
            // create new order
            book.orders[id] = Order(_maker, _haveAmount, _wantAmount, 0, 0);
            return id;
        }
        require(order.maker == _maker, "hash collision");
        // duplicate orders, combine them into one
        order.haveAmount = order.haveAmount.add(_haveAmount);
        order.wantAmount = order.wantAmount.add(_wantAmount);
        require(order.haveAmount < INPUTS_MAX && order.wantAmount < INPUTS_MAX, "combined input over limit");
        // caller should take no further action
        return ZERO_ID;
    }

    // insert [id] as [prev].next
    function insertAfter(
        Book storage book,
        bytes32 id,
        bytes32 prev
    )
        internal
    {
        // prev => [id] => next
        bytes32 next = book.orders[prev].next;
        book.orders[id].prev = prev;
        book.orders[id].next = next;
        book.orders[next].prev = id;
        book.orders[prev].next = id;
    }

    // find the next id (position) to insertAfter
    function find(
        Book storage book,
        Order storage newOrder,
        bytes32 id // [id] => newOrder
    )
        internal
        view
 	    returns (bytes32)
    {
        // [junk] => [0] => order => [FF]
        Order storage order = book.orders[id];
        do {
            order = book.orders[id = order.next];
        } while(!newOrder.betterThan(order));

        // [0] <= order <= [FF]
        do {
            order = book.orders[id = order.prev];
        } while(newOrder.betterThan(order));

        return id;
    }

    // memory version of find
    function m_find(
        Book storage book,
        Order memory newOrder,
        bytes32 id // [id] => newOrder
    )
        internal
        view
 	    returns (bytes32)
    {
        // [junk] => [0] => order => [FF]
        Order storage order = book.orders[id];
        do {
            order = book.orders[id = order.next];
        } while(!newOrder.m_betterThan(order));

        // [0] <= order <= [FF]
        do {
            order = book.orders[id = order.prev];
        } while(newOrder.m_betterThan(order));

        return id;
    }

    // place the new order into its correct position
    function place(
        Book storage book,
        bytes32 newID,
        bytes32 assistingID
    )
        internal
 	    returns (bytes32)
    {
        Order storage newOrder = book.orders[newID];
        bytes32 id = book.find(newOrder, assistingID);
        book.insertAfter(newID, id);
        return id;
    }

    // NOTE: this function does not payout nor refund
    // Use payout/refund/fill instead
    function _remove(
        Book storage book,
        bytes32 id
    )
        internal
    {
        // TODO: handle order outside of the book, where next or prev is nil
        Order storage order = book.orders[id];
        // before: prev => order => next
        // after:  prev ==========> next
        book.orders[order.prev].next = order.next;
        book.orders[order.next].prev = order.prev;
        delete book.orders[id];
    }

    function payout(
        Book storage book,
        bytes32 id
    )
        internal
    {
        if (book.orders[id].wantAmount > 0) {
            book.wantToken.transfer(book.orders[id].maker, book.orders[id].wantAmount);
        }
        book._remove(id);
    }

    function refund(
        Book storage book,
        bytes32 id
    )
        internal
    {
        if (book.orders[id].haveAmount > 0) {
            book.haveToken.transfer(book.orders[id].maker, book.orders[id].haveAmount);
        }
        book._remove(id);
    }

    function payoutPartial(
        Book storage book,
        bytes32 id,
        uint fillableHave,
        uint fillableWant
    )
        internal
    {
        Order storage order = book.orders[id];
        require(order.exists(), "order not exist");
        require(fillableHave <= order.haveAmount, "fill more than have amount");
        order.haveAmount = order.haveAmount.sub(fillableHave);
        if (fillableWant < order.wantAmount) {
            // no need for SafeMath here
            order.wantAmount -= fillableWant;
        } else {
            // possibly profit from price diffirent
            order.wantAmount = 0;
        }
        book.wantToken.transfer(order.maker, fillableWant);
        // TODO: emit event for 'partial order filled'
        if (order.isEmpty()) {
            book.refund(id);
            // TODO: emit event for 'remain order rejected'
        }
    }

    function fill(
        Book storage orderBook,
        bytes32 orderID,
        Book storage redroBook
    )
        internal
        returns (bytes32 nextID)
    {
        Order storage order = orderBook.orders[orderID];
        bytes32 redroID = redroBook.topID();

        while (redroID != LAST_ID) {
            if (order.isEmpty()) {
                break;
            }
            Order storage redro = redroBook.orders[redroID];
            if (!order.fillableBy(redro)) {
                break;
            }
            if (order.haveAmount < redro.wantAmount) {
                uint fillable = order.haveAmount * redro.haveAmount / redro.wantAmount;
                require(fillable <= redro.haveAmount, "fillable > have");
                // partially payout the redro and stop
                redroBook.payoutPartial(redroID, fillable, order.haveAmount);
                orderBook.payoutPartial(orderID, order.haveAmount, fillable);
                break;
            }
            // fully payout the redro
            orderBook.payoutPartial(orderID, redro.wantAmount, redro.haveAmount);
            bytes32 next = redro.next;
            redroBook.payout(redroID);
            redroID = next;
        }
        return redroID;
    }

    // absorb and duplicate the effect to initiator
    function absorbPreemptive(
        Book storage book,
        bool useHaveAmount,
        uint target,
        address initiator
    )
        internal
        returns (uint totalBMT, uint totalAMT)
    {
        (totalBMT, totalAMT) = book.absorb(useHaveAmount, target);
        (uint haveAMT, uint wantAMT) = useHaveAmount ? (totalAMT, totalBMT) : (totalBMT, totalAMT);
        if (book.haveToken.allowance(initiator, address(this)) < haveAMT ||
            book.haveToken.balanceOf(initiator) < haveAMT) {
            // not enough alowance to side-absorb, halt the absorption for this call
            return (0, 0);
        }
        book.haveToken.transferFrom(initiator, book.haveToken.dex(), haveAMT);
        book.haveToken.dexBurn(haveAMT);
        book.wantToken.dexMint(wantAMT);
        book.wantToken.transfer(initiator, wantAMT);
        return (totalBMT, totalAMT);
    }

    function absorb(
        Book storage book,
        bool useHaveAmount,
        uint target
    )
        internal
        returns(uint totalBMT, uint totalAMT)
    {
        bytes32 id = book.topID();
        while(id != LAST_ID && totalAMT < target) {
            dex.Order storage order = book.orders[id];
            uint amt = useHaveAmount ? order.haveAmount : order.wantAmount;
            uint bmt = useHaveAmount ? order.wantAmount : order.haveAmount;
            if (totalAMT.add(amt) <= target) {
                // fill the order
                book.haveToken.dexBurn(order.haveAmount);
                book.wantToken.dexMint(order.wantAmount);
                bytes32 next = order.next;
                book.payout(id);
                id = next;
                // TODO: emit event for 'full order filled'
            } else {
                // partial order fill
                uint fillableAMT = target.sub(totalAMT);
                bmt = bmt * fillableAMT / amt;
                amt = fillableAMT;
                uint fillableHave = useHaveAmount ? amt : bmt;
                uint fillableWant = useHaveAmount ? bmt : amt;
                // fill the partial order
                book.haveToken.dexBurn(fillableHave);
                book.wantToken.dexMint(fillableWant);
                book.payoutPartial(id, fillableHave, fillableWant);
                // extra step to make sure the loop will stop after this
                id = LAST_ID;
            }
            totalAMT = totalAMT.add(amt);
            totalBMT = totalBMT.add(bmt);
        }
        // not enough order, return all we have
        return (totalBMT, totalAMT);
    }

    // amountToAbsorb calculates the amount will be absorbed by absorb()
    // always be updated with absorb() logic
    function amountToAbsorb(
        Book storage book,
        bool useHaveAmount,
        uint target
    )
        internal
        view
        returns(uint totalBMT, uint totalAMT)
    {
        bytes32 id = book.topID();
        while(id != LAST_ID && totalAMT < target) {
            dex.Order storage order = book.orders[id];
            uint amt = useHaveAmount ? order.haveAmount : order.wantAmount;
            uint bmt = useHaveAmount ? order.wantAmount : order.haveAmount;
            if (totalAMT.add(amt) <= target) {
                id = order.next;
            } else {
                // partial order fill
                uint fillableAMT = target.sub(totalAMT);
                bmt = bmt * fillableAMT / amt;
                amt = fillableAMT;
                // extra step to make sure the loop will stop after this
                id = LAST_ID;
            }
            totalAMT = totalAMT.add(amt);
            totalBMT = totalBMT.add(bmt);
        }
        // not enough order, return all we have
        return (totalBMT, totalAMT);
    }
}
