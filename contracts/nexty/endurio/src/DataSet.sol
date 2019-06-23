pragma solidity ^0.5.2;

contract DataSet {

    struct Proposal {
        // address who make the preemptive initization
        address payable initiator;

        // amount of stable coin will be absorbed
        // positive value mean upper anchor price prediction
        // negative value mean lower anchor price prediction
        int256 amount;

        // NTY amount need to be locked for the preemptive initization
        uint256 value;

        // slashing rate
        uint256 slashingRate;

        // lockdown duration
        uint256 lockdownDuration;
    }

    struct Premptive {
        // address of who initilize the premptive
        // and also the address to receive premptive absorb
        address payable initiator;

        // unlock block number of current active premetive
        // equal to 0 mean non-active premptive
        uint256 unlockBlockNumber;

        // amount of stable coin will be absorbed
        // positive value mean upper anchor price prediction
        // negative value mean lower anchor price prediction
        int256 amount;

        // the final balance that will be released to initiator at unlock time
        // before activing a new premptive
        uint256 balance;

        // slashing rate
        uint256 slashingRate;

        // lockdown duration
        uint256 lockdownDuration;
    }

    struct Order {
        address maker;
        uint256 haveAmount;
        uint256 wantAmount;

        // linked list
        bytes32 prev;
        bytes32 next;
    }

    struct OrderList {
        mapping (bytes32 => Order) orders;
        // bytes32 top;	// the highest priority (lowest sell or highest buy)
        // bytes32 bottom;	// the lowest priority (highest sell or lowest buy)
    }

    // current best premptive proposal
    Proposal internal currentProposal;

    // current active preemptive
    Premptive internal currentPremptive;

    mapping(bool => OrderList) internal books;
    // contract private nonce ot generate unique ids
    mapping(address => uint256) internal pNonce;
    bytes32 constant zeroBytes32 = bytes32(0);
    address payable constant ZERO_ADDRESS = address(0x0);

    function abs(int256 _value) internal pure returns (uint256) {
        if (_value < 0) {
            return uint256(_value*-1);
        }
        return uint256(_value);
    }
}