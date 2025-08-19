package solidity_base

/*
// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract Voting {

    mapping(address user => uint count) voteMapping;

    function vote(address somebody) public returns(bool) {
        voteMapping[somebody] += 1;
        return true;
    }

    function getVotes(address somebody) public view returns(uint) {
        return voteMapping[somebody];
    }

    function restVotes(address somebody) public returns (bool) {
        delete voteMapping[somebody];
        return true;
    }

    function reversalString(string calldata inputString) public pure returns(string memory) {
        bytes memory b = bytes(inputString);
        bytes memory returnBytes = new bytes(b.length);
        for (uint i = 0; i < b.length; i++) {
            returnBytes[i] = b[b.length - 1 - i];
        }
        return string(returnBytes);
    }


*/
/* 罗马数字转整数 */ /*

   mapping(bytes1 luomaCode => uint256 codeValue) public codeMapping;

   constructor() {
       codeMapping["I"] = 1;
       codeMapping["V"] = 5;
       codeMapping["X"] = 10;
       codeMapping["L"] = 50;
       codeMapping["C"] = 100;
       codeMapping["D"] = 500;
       codeMapping["M"] = 1000;
   }

   function romanToInt(string calldata romanNumber) public view returns(uint256) {
       bytes memory roma = bytes(romanNumber);
       uint256 total;
       for (uint i = 0;i < roma.length;i++) {
           uint256 currentValue = codeMapping[roma[i]];
           uint256 nextValue = i == roma.length - 1 ? 0 : codeMapping[roma[i + 1]];

           if(currentValue < nextValue) {
               total += nextValue - currentValue;
               i++;
           }else {
               total += currentValue;
           }
       }

       return total;
   }


   function intToRoman(uint256 intValue) public pure returns (string memory) {
       require(intValue >= 1 && intValue <= 3999);

       bytes memory romanString;

       // 千位
       uint256 thousand = intValue / 1000;
       if(thousand > 0) {
           for(uint i = 1; i <= thousand ; i++) {
               romanString = abi.encodePacked(romanString, "M");
           }
       }

       // 百位
       uint256 thousandRemainder = (intValue % 1000) / 100;
       if(thousandRemainder > 0) {
           if(thousandRemainder == 9 ) {
               romanString = abi.encodePacked(romanString, "CM");
           }else if (thousandRemainder == 4) {
               romanString = abi.encodePacked(romanString, "CD");
           }else if (thousandRemainder == 5) {
               romanString = abi.encodePacked(romanString, "D");
           }else if (thousandRemainder > 5){
               romanString = abi.encodePacked(romanString, "D");
               for(uint i = 1; i <= thousandRemainder-5 ; i++) {
                   romanString = abi.encodePacked(romanString, "C");
               }
           }else {
               for(uint i = 1; i <= thousandRemainder ; i++) {
                   romanString = abi.encodePacked(romanString, "C");
               }
           }
       }

       // 十位
       uint256 tenRemainder = (intValue % 1000 % 100) / 10;
       if(tenRemainder > 0) {
           if(tenRemainder == 9 ) {
               romanString = abi.encodePacked(romanString, "XC");
           }else if (tenRemainder == 4) {
               romanString = abi.encodePacked(romanString, "XL");
           }else if (tenRemainder == 5) {
               romanString = abi.encodePacked(romanString, "L");
           }else if (tenRemainder > 5){
               romanString = abi.encodePacked(romanString, "L");
               for(uint i = 1; i <= tenRemainder-5 ; i++) {
                   romanString = abi.encodePacked(romanString, "X");
               }
           }else {
               for(uint i = 1; i <= tenRemainder ; i++) {
                   romanString = abi.encodePacked(romanString, "X");
               }
           }
       }

       // 各位
       uint256 onesRemainder = intValue % 1000 % 100 % 10;
       if(onesRemainder > 0) {
           if(onesRemainder == 9 ) {
               romanString = abi.encodePacked(romanString, "IX");
           }else if (onesRemainder == 4) {
               romanString = abi.encodePacked(romanString, "IV");
           }else if (onesRemainder == 5) {
               romanString = abi.encodePacked(romanString, "V");
           }else if (onesRemainder > 5){
               romanString = abi.encodePacked(romanString, "V");
               for(uint i = 1; i <= onesRemainder-5 ; i++) {
                   romanString = abi.encodePacked(romanString, "I");
               }
           }else {
               for(uint i = 1; i <= onesRemainder ; i++) {
                   romanString = abi.encodePacked(romanString, "I");
               }
           }
       }

       return string(romanString);
   }


   function intToRoman2(uint256 num) public pure returns (string memory) {
       require(num >= 1 && num <= 3999, "Invalid range");

       string[13] memory romans = ["M","CM","D","CD","C","XC","L","XL","X","IX","V","IV","I"];
       uint16[13] memory values = [1000,900,500,400,100,90,50,40,10,9,5,4,1];

       bytes memory result;

       for(uint i = 0; i < values.length;i++) {
           while(num >= values[i]) {
               result = abi.encodePacked(result,romans[i]);
               num -= values[i];
           }
       }

       return string(result);
   }

   function testArrayMerge(uint[] calldata a, uint[] calldata b)public pure returns (uint[] memory result) {
       uint[] memory resultArray = new uint[](a.length + b.length);
       for (uint i = 0 ;i<a.length ;i++ ) {
           resultArray[i] = a[i];
       }

       for (uint i = 0 ;i<b.length ;i++ ) {
           resultArray[a.length + i] = b[i];
       }

       return resultArray;
   }

*/
/* 二分法查找数据 */ /*

    function searchMemoryIndex(uint[] memory array, uint target) public pure returns (int) {
        uint low = 0;
        uint high = array.length;

        while(high > low) {
            uint midIndex = (high + low) / 2;
            if(array[midIndex] == target) {
                return int(midIndex);
            }else if(array[midIndex] > target) {
                high = midIndex;
            }else {
                low = midIndex + 1;
            }
        }
        return -1;
    }

    // 内存数组二分查找（纯函数）
    function searchMemory(uint[] memory array, uint target) public pure returns (int) {
        uint low = 0;
        uint high = array.length;

        while (low < high) {
            uint mid = (low + high) / 2;
            if (array[mid] == target) {
                return int(mid);
            } else if (array[mid] < target) {
                low = mid + 1;
            } else {
                high = mid;
            }
        }
        return -1;
    }
}*/
