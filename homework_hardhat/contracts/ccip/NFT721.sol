// SPDX-License-Identifier: SEE LICENSE IN LICENSE
pragma solidity ^0.8;
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";

contract NFT721 is ERC721URIStorage {

    string public constant BASE_URI = "https://ipfs.io/ipfs/QmYuKY45Aq87LeL1R5dhb1hqHLp6ZFbJaCP8jxqKM1MX6y/babe_ruth_1.json";
    uint256 public tokenId = 1;

    constructor() ERC721("NFT721", "NFT721") {

    }

    // 定义safeMint函数
    function safeMint(address to) public {
        _safeMint(to, tokenId);
        _setTokenURI(tokenId, BASE_URI);
        tokenId++;
    }
}