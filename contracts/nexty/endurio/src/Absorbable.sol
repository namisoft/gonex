pragma solidity ^0.5.2;

import "./lib/util.sol";
import "./lib/dex.sol";
import "./lib/absn.sol";
import "./Orderbook.sol";

/**
 * Mechanisms of absorption logic: active, passive and pre-emptive.
 */
contract Absorbable is Orderbook {
    using dex for dex.Book;
    using absn for absn.Absorption;
    using absn for absn.Preemptive;

    IToken VolatileToken;
    IToken StablizeToken; // intentional incorrect spelling

    // constants
    uint EXPIRATION = 1 weeks / 2 seconds;
    int DURATION = int(EXPIRATION / 2);

    // last absorption
    absn.Absorption internal last;
    // lockdown tracks the active and current being lockdown pre-emptive absorption
    absn.Preemptive internal lockdown;

    constructor (
        uint absorptionDuration,
        uint absorptionExpiration
    )
        public
    {
        if (absorptionExpiration > 0) EXPIRATION = absorptionExpiration;
        DURATION = int(absorptionDuration > 0 ? absorptionDuration : absorptionExpiration / 2);
    }

    function registerTokens(
        address volatileToken,
        address stablizeToken
    )
        public
    {
        require(address(VolatileToken) == address(0), "already set");
        require(address(StablizeToken) == address(0), "already set");
        VolatileToken = IToken(volatileToken);
        StablizeToken = IToken(stablizeToken);
        super.registerTokens(volatileToken, stablizeToken);
        // trigger the first blank absorption
        triggerAbsorption(StablizeToken.totalSupply(), false);
    }

    function toString(address x) internal pure returns (string memory) {
        bytes memory b = new bytes(20);
        for (uint i = 0; i < 20; i++)
            b[i] = byte(uint8(uint(x) / (2**(8*(19 - i)))));
        return string(b);
    }

    modifier consensus() {
        require(msg.sender == address(0x0), "consensus only");
        _;
    }

    // for ethstat
    function getRemainToAbsorb() public view returns (bool, int) {
        if (last.target == 0) {
            return (false, 0);
        }
        return (true, util.sub(last.target, StablizeToken.totalSupply()));
    }

    // called by the consensus on each block
    // median price = target / StablizeToken.totalSupply()
    // zero target is fed for no median price available
    function onBlockInitialized(uint target) public consensus {
        if (last.isExpired()) {
            // absorption takes no longer than one duration
            delete last;
        }
        if (lockdown.unlockable()) {
            unlock();
        }
        uint supply = StablizeToken.totalSupply();
        if (target > 0) { // absorption block
            if (shouldTriggerPassive() ||
                shouldTriggerActive(supply, target)) {
                triggerAbsorption(target, false);
            }
            if (lockdown.isLocked()) {
                // WIP: slash the pre-emptive maker if target goes wrong way
                int diviation = util.sub(target, supply);
                if (checkAndSlash(diviation) && last.isPreemptive) {
                    // lockdown violation, halt the preemptive absorption for this block
                    return;
                }
            }
        }
        if (last.isAbsorbing(supply)) {
            int nextAmount = calcNextAbsorption();
            absorb(nextAmount);
        }
    }

    function calcNextAbsorption() internal view returns(int) {
        int total = util.sub(last.target, last.supply);
        int remain = util.sub(last.target, StablizeToken.totalSupply());
        if (!util.inOrder(0, remain, total)) {
            // target reached or passed
            return 0;
        }
        int amount = total / DURATION;
        if (!util.inOrder(0, amount, remain)) {
            // don't over absorb
            return remain;
        }
        return amount;
    }

    // shouldTriggerPassive returns whether a new passive absorption can be activated
    // passive condition: 1 duration without any active absorption or absorption never occurs
    function shouldTriggerPassive() internal view returns (bool) {
        return last.isExpired();
    }

    // shouldTriggerActive returns whether the new target is sufficient to trigger a new active absorption
    // make things simple by compare only the (target-supply) instead (target-supply)/supply
    function shouldTriggerActive(uint supply, uint target) internal view returns (bool) {
        if (target == supply) {
            return false;
        }
        if (last.target == last.supply) {
            return true;
        }
        // int a = util.sub(target, supply);
        // int b = util.sub(last.target, last.supply);
        // return a/b >= 2 || util.inOrder(-2, b/a, 0);
        if (target > supply) {
            uint a = target - supply;
            if (last.target > last.supply) {
                uint b = last.target - last.supply;
                return a / b >= 2;
            } else {
                uint b = last.supply - last.target;
                return b / a <= 2;
            }
        } else {
            uint a = supply - target;
            if (last.target < last.supply) {
                uint b = last.supply - last.target;
                return a / b >= 2;
            } else {
                uint b = last.target - last.supply;
                return b / a <= 2;
            }
        }
     }

    function unlock() internal {
        if (!lockdown.exists()) {
            return;
        }
        if (lockdown.stake > 0) {
            VolatileToken.transfer(lockdown.maker, lockdown.stake);
        }
        delete lockdown;
    }

    function triggerAbsorption(uint target, bool isPreemptive) internal {
        last = absn.Absorption(block.number + EXPIRATION,
            StablizeToken.totalSupply(),
            target,
            isPreemptive);
    }

    function absorb(
        int stableTokenAmount
    )
        internal
        returns(uint totalVOL, uint totalSTB)
    {
        dex.Book storage book = books[stableTokenAmount > 0 ? Ask : Bid];
        bool useHaveAmount = book.haveToken == StablizeToken;
        if (last.isPreemptive) {
            return book.absorbPreemptive(useHaveAmount, util.abs(stableTokenAmount), lockdown.maker);
        }
        return book.absorb(useHaveAmount, util.abs(stableTokenAmount));
    }

    /**
     * @dev slash the initiator whenever the price is moving in
     * opposition direction with the initiator's direction,
     * the initiator's deposited balance will be minus by _amount
     *
     * _amount is the NTY value need to be burn, calculate in the consensus level
     * _amount = |d/D|/SlashingDuration
     * d = MedianPriceDeviation
     * D = X/S, X is the amount of NewSD will be absorbed, S is the current NewSD total supply
     * consensus need to update the balance to burn amount return from calling slash function
     *
     * @return true if the lockdown is violated and get slashed
     */
    function checkAndSlash(int diviation) internal returns (bool) {
        if (!util.inOrder(lockdown.amount, 0, diviation)) {
            // same direction, no slashing
            return false;
        }
        // lockdown violated
        uint slashed = uint(-diviation/lockdown.amount) / lockdown.slashingDuration;
        if (slashed == 0) {
            slashed = 1; // minimum 1 wei
        }
        if (lockdown.stake < slashed) {
            slashed = lockdown.stake;
            // there's nothing at stake anymore, clear the lockdown and its absorption
            delete last;
            unlock();
        }
        lockdown.stake -= slashed;
        VolatileToken.dexBurn(slashed);
        // this slashed NTY will be burnt by the consensus by calling setBalance
        return true;
    }
}