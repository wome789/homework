// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract myCount {
    mapping(address => uint) count;

    function increment(address userAddress) public {
        count[userAddress]++;
    }

    function getCount(address userAddress) public view returns (uint) {
        return count[userAddress];
    }
}
