// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract PrecompileTester {
    // ecrecover (address 0x01)
    function testEcrecover(bytes memory data) public view returns (bytes memory) {
        (bool success, bytes memory result) = address(0x01).staticcall(data);
        require(success, "ecrecover call failed");
        return result;
    }
    
    // sha256 (address 0x02)
    function testSha256(bytes memory data) public view returns (bytes memory) {
        (bool success, bytes memory result) = address(0x02).staticcall(data);
        require(success, "sha256 call failed");
        return result;
    }
    
    // ripemd160 (address 0x03)
    function testRipemd160(bytes memory data) public view returns (bytes memory) {
        (bool success, bytes memory result) = address(0x03).staticcall(data);
        require(success, "ripemd160 call failed");
        return result;
    }
    
    // identity (address 0x04)
    function testIdentity(bytes memory data) public view returns (bytes memory) {
        (bool success, bytes memory result) = address(0x04).staticcall(data);
        require(success, "identity precompile failed");
        return result;
    }

    // modexp (address 0x05)
    function testModexp(bytes memory data) public view returns (bytes memory) {
        (bool success, bytes memory result) = address(0x05).staticcall(data);
        require(success, "modexp call failed");
        return result;
    }
    
    // bn256Add (address 0x06)
    function testBn256Add(bytes memory data) public view returns (bytes memory) {
        (bool success, bytes memory result) = address(0x06).staticcall(data);
        require(success, "bn256Add call failed");
        return result;
    }
    
    // bn256ScalarMul (address 0x07)
    function testBn256ScalarMul(bytes memory data) public view returns (bytes memory) {
        (bool success, bytes memory result) = address(0x07).staticcall(data);
        require(success, "bn256ScalarMul call failed");
        return result;
    }
    
    // bn256Pairing (address 0x08)
    function testBn256Pairing(bytes memory data) public view returns (bool) {
        (bool success, bytes memory result) = address(0x08).staticcall(data);
        require(success, "bn256Pairing call failed");
        return result.length > 0 && result[0] != 0;
    }

    // blake2 (eip-152) (address 0x09)
    function testBlake2(bytes memory data) public view returns (bytes memory) {
        (bool success, bytes memory result) = address(0x09).staticcall(data);
        require(success, "blake2 call failed");
        return result;
    }

    // pointEvaluation (address 0x0A)
    function testPointEvaluation(bytes memory data) public view returns (bytes memory) {
        (bool success, bytes memory result) = address(0x0A).staticcall(data);
        require(success, "pointEvaluation call failed");
        return result;
    }

    // PECTRA (EIP-2537)
    // Precompiles for BLS12-381 curve operations
    
    // BLS12_G1ADD (address 0x0B)
    function testBLS12_G1Add(bytes memory data) public view returns (bytes memory) {
        (bool success, bytes memory result) = address(0x0B).staticcall(data);
        require(success, "BLS12_G1Add call failed");
        return result;
    }

    // BLS12_G1MSM (address 0x0C)
    function testBLS12_G1MSM(bytes memory data) public view returns (bytes memory) {
        (bool success, bytes memory result) = address(0x0C).staticcall(data);
        require(success, "BLS12_G1MSM call failed");
        return result;
    }

    // BLS12_G2ADD (address 0x0D)
    function testBLS12_G2Add(bytes memory data) public view returns (bytes memory) {
        (bool success, bytes memory result) = address(0x0D).staticcall(data);
        require(success, "BLS12_G2Add call failed");
        return result;
    }
    
    // BLS12_G2MSM (address 0x0E)
    function testBLS12_G2MSM(bytes memory data) public view returns (bytes memory) {
        (bool success, bytes memory result) = address(0x0E).staticcall(data);
        require(success, "BLS12_G2MSM call failed");
        return result;
    }

    // BLS12_PAIRING_CHECK (address 0x0F)
    function testBLS12_PairingCheck(bytes memory data) public view returns (bool) {
        (bool success, bytes memory result) = address(0x0F).staticcall(data);
        require(success, "BLS12_PairingCheck call failed");
        return result.length > 0 && result[0] != 0;
    }

    // BLS12_MAP_FP_TO_G1 (address 0x10)
    function testBLS12_MapFPtoG1(bytes memory data) public view returns (bytes memory) {
        (bool success, bytes memory result) = address(0x10).staticcall(data);
        require(success, "BLS12_MapFPtoG1 call failed");
        return result;
    }

    //BLSS12_MAP_FP2_TO_G2 (address 0x11)
    function testBLSS12_MapFP2toG2(bytes memory data) public view returns (bytes memory) {
        (bool success, bytes memory result) = address(0x11).staticcall(data);
        require(success, "BLSS12_MapFP2toG2 call failed");
        return result;
    }
}
