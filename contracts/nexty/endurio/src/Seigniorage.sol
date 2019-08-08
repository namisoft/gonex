pragma solidity ^0.5.2;

import "./Preemptivable.sol";

/**
 * Seigniorage Shares Stablecoin
 */
contract Seigniorage is Preemptivable {
    constructor (
        uint absorptionDuration,
        uint absorptionExpiration,
        uint initialSlashingDuration,
        uint initialLockdownExpiration
    )
        Preemptivable(
            absorptionDuration,
            absorptionExpiration,
            initialSlashingDuration,
            initialLockdownExpiration
        )
        public
    {
    }
}