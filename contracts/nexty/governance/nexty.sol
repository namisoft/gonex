pragma solidity ^0.5.0;

import "node_modules/openzeppelin-solidity/contracts/math/SafeMath.sol";
import "node_modules/openzeppelin-solidity/contracts/token/ERC20/IERC20.sol";

/**
 * @title Nexty sealers management smart contract
 */
contract NextyGovernance {
    using SafeMath for uint256;

    // zero address
    address constant ZERO_ADDRESS = address(0x0);

    enum Status {
        PENDING_ACTIVE,     // Sealer deposited enough NTFs into registration contract successfully.

        ACTIVE,             // Sealer send request to become a sealer
                            // and added into activation sealer set successfully

        PENDING_WITHDRAW,   // Sealer send request to exit from activation sealer set successfully.
                            // Sealer casted out of activation sealer set

        WITHDRAWN,          // Sealer already withdrawn their deposit NTFs successfully.
                            // They can only make withdrawal after withdrawal period.

        PENALIZED           //Sealer marked as penalized node (update by consensus or voting result via dapp)
                            //and cannot become active sealer and cannot withdraw balance neither.
    }

    enum ProposalType {
        JOIN,               // request to join in activation sealers

        LEAVE,              // request to leave out of activation sealers

        SLASH               // propose slashing for the sealer
    }

    struct Account {
        Status status;
        // ntf amount deposited
        uint256 balance;
        // delegated address to seal blocks
        address signer;
        // withdrawable block number after leaving
        uint256 unlockHeight;
    }

    struct Proposal {
        // proposal type
        ProposalType proposalType;
        // delegated address to sign blocks
        address signer;
        // token holder and sealing reward beneficiary
        address beneficiary;
    }

    // Consensus variables

    // index = 0
    // signers array
    address[] public signers;

    // index = 1
    // signer => coinbase (beneficiary address) map
    mapping(address => address) public signerCoinbase;

    // End of consensus variables

    // coinbase => NTF Account map
    mapping(address => Account) public account;

    // minimum of deposited NTF to join
    uint256 public stakeRequire;
    // minimum number of blocks signer has to wait from leaving block to withdraw the fund
    uint256 public stakeLockHeight;

    // NTF token contract, unit used to join Nexty sealers
    IERC20 public token;

    // CoLoa hardfork consensus state need to come afterward
    // to keep the ancient ThangLong hardfork state

    // Pending proposals that sealer request to join/leave activation sealer list
    // or proposing slash to a sealer
    Proposal[] public proposals;

    event Deposited(address _sealer, uint _amount);
    event Joined(address _sealer, address _signer);
    event Left(address _sealer, address _signer);
    event Withdrawn(address _sealer, uint256 _amount);
    event Banned(address _sealer);
    event Unbanned(address _sealer);

    /**
    * Check if address is a valid destination to transfer tokens to
    * - must not be zero address
    * - must not be the token address
    * - must not be the sender's address
    */
    modifier validSigner(address _signer) {
        require(signerCoinbase[_signer] == ZERO_ADDRESS, "coinbase already used");
        require(_signer != ZERO_ADDRESS, "signer zero");
        require(_signer != address(this), "same contract's address");
        require(_signer != msg.sender, "same sender's address");
        _;
    }

    modifier notBanned() {
        require(account[msg.sender].status != Status.PENALIZED, "banned ");
        _;
    }

    modifier joinable() {
        require(account[msg.sender].status != Status.ACTIVE, "already joined ");
        require(account[msg.sender].balance >= stakeRequire, "not enough ntf");
        _;
    }

    modifier leaveable() {
        require(account[msg.sender].status == Status.ACTIVE, "not joined ");
        _;
    }

    modifier withdrawable() {
        require(isWithdrawable(msg.sender), "unable to withdraw at the moment");
        _;
    }

    modifier onlyConsensus() {
        require(msg.sender == ZERO_ADDRESS, "only can call from consensus level");
        _;
    }

    /**
    * contract initialize
    */
    constructor() public {
    }

    // Get ban status of a sealer's address
    function isBanned(address _address) public view returns(bool) {
        return (account[_address].status == Status.PENALIZED);
    }

    ////////////////////////////////

    function addSigner(address _signer) internal {
        signers.push(_signer);
    }

    function removeSigner(address _signer) internal {
        for (uint i = 0; i < signers.length; i++) {
            if (_signer == signers[i]) {
                signers[i] = signers[signers.length - 1];
                signers.length--;
                return;
            }
        }
    }

    /**
    * Transfer the NTF from token holder to registration contract.
    * Sealer might have to approve contract to transfer an amount of NTF before calling this function.
    * @param _amount NTF Tokens to deposit
    */
    function deposit(uint256 _amount) public returns (bool) {
        token.transferFrom(msg.sender, address(this), _amount);
        account[msg.sender].balance = (account[msg.sender].balance).add(_amount);
        emit Deposited(msg.sender, _amount);
        return true;
    }

    /**
    * To allow deposited NTF participate joining in as sealer.
    * Participate already must deposit enough NTF via Deposit function.
    * It takes signer as parameter.
    * @param _signer Destination address
    */
    function join(address _signer) public notBanned joinable validSigner(_signer) returns (bool) {
        account[msg.sender].signer = _signer;
        account[msg.sender].status = Status.ACTIVE;
        signerCoinbase[_signer] = msg.sender;
        addSigner(_signer);
        proposals.push(Proposal(ProposalType.JOIN, _signer, msg.sender));
        emit Joined(msg.sender, _signer);
        return true;
    }

    /**
     * @dev consensus can decide to slash a sealer
     */
    function slash(address _signer) external onlyConsensus validSigner(_signer) returns (bool) {
        proposals.push(Proposal(ProposalType.SLASH, _signer, signerCoinbase[_signer]));
        return true;
    }

    /**
     * @dev return the first proposal if any; otherwise, return -1
     */
    function getFirstProposal() external view returns (int256, address, address) {
        if (proposals.length > 0) {
            return (getProposalType(proposals[0].proposalType), proposals[0].signer, proposals[0].beneficiary);
        } else {
            return (-1, ZERO_ADDRESS, ZERO_ADDRESS);
        }
    }

    /**
     * @dev return all proposals
     */
    function getAllProposals() external view returns (int256[] memory, address[] memory, address[] memory) {
        int256[] memory _types = new int256[](proposals.length);
        address[] memory _signers = new address[](proposals.length);
        address[] memory _beneficiaries = new address[](proposals.length);
        for (uint i = 0; i < proposals.length; i++) {
            _types[i] = getProposalType(proposals[i].proposalType);
            _signers[i] = proposals[i].signer;
            _beneficiaries[i] = proposals[i].beneficiary;
        }
        return (_types, _signers, _beneficiaries);
    }

    /**
     * @dev remove all proposals
     */
    function cleanProposals() external onlyConsensus returns (bool) {
        delete proposals;
        return true;
    }

    /**
     * @dev pop proposal to block header preparing for activation (256 blocks later)
     * return (-1, 0x0, 0x0) if not found
     */
    function popProposal() external onlyConsensus returns (int256, address, address) {
        if (proposals.length > 0) {
            Proposal memory elem = proposals[0];
            for (uint i = 0; i < proposals.length - 1; i++) {
                proposals[i] = proposals[i + 1];
            }
            delete proposals[proposals.length - 1];
            proposals.length--;
            return (getProposalType(elem.proposalType), elem.signer, elem.beneficiary);
        } else {
            return (-1, ZERO_ADDRESS, ZERO_ADDRESS);
        }
    }

    /**
    * Request to exit out of activation sealer set
    */
    function leave() public notBanned leaveable returns (bool) {
        address _signer = account[msg.sender].signer;

        account[msg.sender].signer = ZERO_ADDRESS;
        account[msg.sender].status = Status.PENDING_WITHDRAW;
        account[msg.sender].unlockHeight = stakeLockHeight.add(block.number);
        delete signerCoinbase[_signer];
        removeSigner(_signer);
        proposals.push(Proposal(ProposalType.LEAVE, _signer, msg.sender));
        emit Left(msg.sender, _signer);
        return true;
    }

    /**
    * To withdraw sealer’s NTF balance when they already exited and after withdrawal period.
    */
    function withdraw() public notBanned withdrawable returns (bool) {
        uint256 amount = account[msg.sender].balance;
        account[msg.sender].balance = 0;
        account[msg.sender].status = Status.WITHDRAWN;
        token.transfer(msg.sender, amount);
        emit Withdrawn(msg.sender, amount);
        return true;
    }

    function getStatusCode(Status _status) private pure returns(uint256){
        if (_status == Status.PENDING_ACTIVE) return 0;
        if (_status == Status.ACTIVE) return 1;
        if (_status == Status.PENDING_WITHDRAW) return 2;
        if (_status == Status.WITHDRAWN) return 3;
        return 127;
    }

    function getProposalType(ProposalType _type) private pure returns (int256) {
        if (_type == ProposalType.JOIN) return 0;
        if (_type == ProposalType.LEAVE) return 1;
        if (_type == ProposalType.SLASH) return 2;
        return -1;
    }

    function getStatus(address _address) public view returns(uint256) {
        return getStatusCode(account[_address].status);
    }

    function getBalance(address _address) public view returns(uint256) {
        return account[_address].balance;
    }

    function getCoinbase(address _address) public view returns(address) {
        return account[_address].signer;
    }

    function getUnlockHeight(address _address) public view returns(uint256) {
        return account[_address].unlockHeight;
    }

    function isWithdrawable(address _address) public view returns(bool) {
        return
        (account[_address].status != Status.ACTIVE) &&
        (account[_address].status != Status.PENALIZED) &&
        (account[_address].unlockHeight < block.number);
    }

    /**
     * @dev get current sealer list (signer, beneficiary)
     */
    function getSealers() external view returns (address[] memory, address[] memory) {
        address[] memory _signers = new address[](signers.length);
        address[] memory _beneficiaries = new address[](signers.length);
        for (uint i = 0; i < signers.length; i++) {
            _signers[i] = signers[i];
            _beneficiaries[i] = signerCoinbase[_signers[i]];
        }
        return (_signers, _beneficiaries);
    }
}