pragma solidity ^0.5.2;

library ABI {
    function decode(
        bytes memory _data
    )
        public
        pure
        returns(uint256, bytes32)
    {
        require(_data.length == 64, "invalid data");
        return abi.decode(_data, (uint256, bytes32));
    }

    function encode(
        uint256 wantAmount,
        bytes32 assistingID
    )
        public
        pure
        returns(bytes memory)
    {
        return abi.encode(wantAmount, assistingID);
    }
}