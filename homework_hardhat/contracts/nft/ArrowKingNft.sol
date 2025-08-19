// SPDX-License-Identifier: SEE LICENSE IN LICENSE
pragma solidity ^0.8;
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721Enumerable.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract ArrowKingNft is ERC721Enumerable, Ownable {
    constructor() ERC721("ArrowKingNft", "AKN") Ownable(msg.sender) {}

    uint256 [] private _tokenIds;

    function mint(address to, uint256 tokenId) external onlyOwner {
        _mint(to, tokenId);
        _tokenIds.push(tokenId);
    }
    
    function getTokenIds() external view returns (uint256[] memory) {
        return _tokenIds;
    }
}
