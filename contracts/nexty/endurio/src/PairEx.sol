pragma solidity ^0.5.2;

import "openzeppelin-solidity/contracts/math/Math.sol";
import "openzeppelin-solidity/contracts/math/SafeMath.sol";
import "./OrderBook.sol";

contract PairEx is OrderBook {

    // Minimum amount bump percentage to overwrite the current non-active premptive
    uint256 constant AMOUNT_BUMP = 10;

    uint256 constant LOCK_DURATION = 2 weeks;

    uint256 constant SLASHING_RATE = 10 minutes;

    modifier unlockable() {
        require(block.number > currentPremptive.unlockBlockNumber, "is in lock duration of current premptive");
        _;
    }

    modifier activable() {
        require(currentPremptive.initiator == ZERO_ADDRESS && currentPremptive.unlockBlockNumber == 0, "current premptive is not finished yet");
        _;
    }

    event PremptiveActivation(address indexed _initator, int256 _amount, uint256 _value);
    event PremptiveUnlock(address indexed _initiator, uint256 _amount);
    event PremptiveAbsorption(address indexed _initiator, uint256 _stb, uint256 _vol);

    constructor (
        address _volatileTokenAddress,
        address _stableTokenAddress
    )
        public
    {
        if (_volatileTokenAddress != ZERO_ADDRESS) {
            volatileTokenRegister(_volatileTokenAddress);
        }
        if (_stableTokenAddress != ZERO_ADDRESS) {
            stableTokenRegister(_stableTokenAddress);
        }
        initBooks();
    }

    function initBooks()
        private
    {
        Order memory order;
        order.wantAmount = 1;
        // Selling Book
        books[Sell].orders[zeroBytes32] = order;
        // Buying Book
        books[Buy].orders[zeroBytes32] = order;
    }

    function setup(
        address _volatileTokenAddress,
        address _stableTokenAddress
    )
        public
    {
        volatileTokenRegister(_volatileTokenAddress);
        stableTokenRegister(_stableTokenAddress);
        token[Volatile].setup(address(this));
        token[Stable].setup(address(this));
    }

    function getOrderType()
        public
        view
        returns (bool)
    {
        address _sender = msg.sender;
        require(_sender == address(token[Stable]) || _sender == address(token[Volatile]), "only VolatileToken and StableToken accepted");
        return _sender == address(token[Stable]);
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

        while (redroTopID != zeroBytes32) {
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
                _remove(_redroType, redroTopID, false);
            }
            redroTopID = top(_redroType);
            if (order.haveAmount == 0 || order.wantAmount == 0) {
                _remove(_orderType, _orderID, false);
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
        returns (uint256, uint256)
    {
        OrderList storage book = books[_orderType];
        uint256 totalSTB;
        uint256 totalVOL;
        bytes32 cursor = top(_orderType);
        while (cursor != zeroBytes32 && totalSTB < _stableTokenTarget) {
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

    function absorb(
        bool _inflate,
        uint256 _stableTokenTarget
    )
        public
        returns (uint256, uint256)
    {
        require(msg.sender == address(this), "consensus only");
        bool orderType = _inflate ? Sell : Buy; // inflate by filling NTY sell order
        OrderList storage book = books[orderType];
        uint256 totalVOL;
        uint256 totalSTB;
        bytes32 cursor = top(orderType);
        while (cursor != zeroBytes32 && totalSTB < _stableTokenTarget) {
            Order storage order = book.orders[cursor];
            uint256 vol = _inflate ? order.haveAmount : order.wantAmount;
            uint256 stb = _inflate ? order.wantAmount : order.haveAmount;
            if (totalSTB.add(stb) <= _stableTokenTarget) {
                totalVOL = totalVOL.add(vol);
                totalSTB = totalSTB.add(stb);
                // fill the order
                token[!_inflate].burnFromOwner(order.haveAmount);
                token[_inflate].mintToOwner(order.wantAmount);
                cursor = _remove(orderType, cursor, true);
                // TODO: emit event for 'full order filled'
                continue;
            }
            // partial order fill
            uint256 fillableHave;
            uint256 fillableWant;

            // temporary scope
            {
                uint256 fillableSTB = _stableTokenTarget.sub(totalSTB);
                uint256 fillableVOL = vol.mul(fillableSTB).div(stb);
                totalVOL = totalVOL.add(fillableVOL);
                totalSTB = totalSTB.add(fillableSTB);
                fillableHave = _inflate ? fillableVOL : fillableSTB;
                fillableWant = _inflate ? fillableSTB : fillableVOL;
            }

            // fill the partial order
            token[!_inflate].burnFromOwner(fillableHave);
            token[_inflate].mintToOwner(fillableWant);
            token[_inflate].transfer(order.maker, fillableWant);

            // TODO: emit event for 'partial order filled'
            order.haveAmount = order.haveAmount.sub(fillableHave);
            order.wantAmount = order.wantAmount.sub(fillableWant);
            if (order.haveAmount == 0 || order.wantAmount == 0) {
                _remove(orderType, cursor, false);
                // TODO: emit event for 'remain order rejected'
            }
            break; // stop the absorption
        }

        return (totalVOL, totalSTB);
    }

    // Premptive absorb

    /**
     * @dev anyone can propose a preemptive absorption at anytime
     */
    function propose(int256 amount) external payable {
        uint256 _amount = abs(amount);
        address payable _initiator = currentProposal.initiator;
        uint256 _value = currentProposal.value;
        uint256 _currentAmount = abs(currentProposal.amount);
        if (_amount < _currentAmount.add(_currentAmount.mul(AMOUNT_BUMP).div(100))) {
            revert("your premptive is not good enough compare to current premptive");
        }

        // make new proposal as current proposal
        currentProposal.initiator = msg.sender;
        currentProposal.amount = amount;
        currentProposal.value = msg.value;
        currentProposal.slashingRate = SLASHING_RATE;
        currentProposal.lockdownDuration = LOCK_DURATION;

        // refund for the previous initiator
        if (_value > 0) {
            _initiator.transfer(_value);
        }
    }

    /**
     * @dev cleanup the current premptive lock data,
     * and refund remaining locked amount to premptive initiator after slashing if any.
     */
    function unlockPremptive() external unlockable {
        uint256 _value = currentPremptive.balance;
        address payable _receiver = currentPremptive.initiator;
        currentPremptive = Premptive(ZERO_ADDRESS, 0, 0, 0, 0, 0);
        if (_value > 0) {
            _receiver.transfer(_value);
        }
        emit PremptiveUnlock(_receiver, _value);
    }

    /**
     * @dev move current best proposal as current active premptive
     * then clean up the current proposal to accept new proposal for next premptive.
     *
     * Only consensus can call this function.
     * At consensus level, blochchain core will check some conditions
     * (a.k.a initiator deposited enough NTY for slashing) before calling this method.
     */
    function activatePremptive() external activable returns (bool) {
        require(msg.sender == address(this), "consensus only");
        if (currentProposal.initiator == ZERO_ADDRESS || currentProposal.amount == 0) {
            return false;
        }
        currentPremptive.initiator = currentProposal.initiator;
        currentPremptive.amount = currentProposal.amount;
        currentPremptive.balance = currentProposal.value;
        currentPremptive.unlockBlockNumber = block.number + currentProposal.lockdownDuration;
        currentPremptive.slashingRate = currentProposal.slashingRate;
        currentPremptive.lockdownDuration = currentProposal.lockdownDuration;

        // clean up current proposal to receive proposals for new premtive
        currentProposal = Proposal(ZERO_ADDRESS, 0, 0, 0, 0);

        emit PremptiveActivation(currentPremptive.initiator, currentPremptive.amount, currentPremptive.balance);
        return true;
    }

    /**
     * @dev return true if there's an active premptive; otherwise, return false
     */
    function hasActivePremptive() public view returns (bool) {
        if (currentPremptive.initiator != ZERO_ADDRESS || currentPremptive.unlockBlockNumber > 0) {
            return true;
        }
        return false;
    }

    /**
     * @dev the current atmost premptive proposal
     */
    function getCurrentProposal() public view returns (address, int256, uint256) {
        return (currentProposal.initiator, currentProposal.amount, currentProposal.value);
    }

    /**
     * @dev get current active premptive
     */
    function getCurrentPremptive() public view returns (address, int256, uint256, uint256, uint256, uint256) {
        return (currentPremptive.initiator, currentPremptive.amount,
            currentPremptive.balance, currentPremptive.unlockBlockNumber, currentPremptive.slashingRate, currentPremptive.lockdownDuration);
    }

    /**
     * @dev perform an equivalent premptive absorption (if any) whenever an active/passive absorbs occurs.
     */
    function premptiveAbsorb(
        bool _inflate,
        uint256 _totalSTB,
        uint256 _totalVOL
    ) external {
        require(msg.sender == address(this), "consensus only");
        if (!hasActivePremptive()) {
            revert("no on-going premptive");
        }
        token[_inflate].mintToOwner(_totalSTB);
        token[!_inflate].burnFromOwner(_totalVOL);
        token[_inflate].transfer(currentPremptive.initiator, _totalSTB);
        emit PremptiveAbsorption(currentPremptive.initiator, _totalSTB, _totalVOL);
    }

    /**
     * @dev slash the initiator whenever the price is moving in
     * opposition direction with the initiator's direction,
     * the initiator's deposited balance will be minus by _amount
     *
     * _amount is the NTY value need to be burn, calculate in the consensus level
     * _amount = |d/D|*SlashingRate/LockdownDuration
     * d = MedianPriceDeviation
     * D = X/S, X is the amount of NewSD will be absorbed, S is the current NewSD total supply
     * consensus need to update the balance to burn amount return from calling slash function
     *
     * @return the actual NTY amount will be burn
     */
    function slash(uint256 _amount) external returns (uint256) {
        require(msg.sender == address(this), "consensus only");
        if (!hasActivePremptive()) {
            revert("no on-going premptive");
        }
        uint256 _balance = currentPremptive.balance;
        if (_balance < _amount) {
            currentPremptive.balance = 0;
            return _balance;
        }
        currentPremptive.balance -= _amount;
        return _amount;
    }
}