pragma solidity ^0.5.2;

library BytesConvert {
    function uint256ToBytes(uint256 x) public pure returns (bytes memory b) {
        b = new bytes(32);
        assembly { mstore(add(b, 32), x) }
    }

    function bytesToUint256(bytes memory b) public pure returns (uint256){
        uint256 number;
        for(uint i = 0; i < b.length; i++){
            number = number + uint256(uint8(b[i]))*(2**(8*(b.length-(i+1))));
        }
        return number;
    }

    function bytesToBytes32(bytes memory b, uint offset) public pure returns (bytes32) {
        bytes32 out;

        for (uint i = 0; i < 32; i++) {
            out |= bytes32(b[offset + i] & 0xFF) >> (i * 8);
        }
        return out;
    }
}