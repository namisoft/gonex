pragma solidity ^0.5.2;

interface IOwnableERC223 {
    function initialize(address _owner) external;
    function transfer(address to, uint256 value) external returns (bool);
    function mintToOwner(uint256 _amount) external;
    function burnFromOwner(uint256 _amount) external;
}