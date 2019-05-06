pragma solidity ^0.5.2;

interface IPairEx {
    function sell(uint256 _volatileTokenAmount, uint256 _stableTokenAmount, address maker, bytes32 index) external;
    function buy(uint256 _stableTokenAmount, uint256 _volatileTokenAmount, address payable maker, bytes32 index) external;
    function volatileTokenRegister(address _volatileTokenAddress) external;
    function stableTokenRegister(address _stableTokenAddress) external;
}