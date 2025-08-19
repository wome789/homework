// SPDX-License-Identifier: SEE LICENSE IN LICENSE
pragma solidity ^0.8;
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "../auction/AuctionContract.sol";

/* 
    拍卖工厂合约
    1. 创建拍卖合约
    2. 获取所有拍卖合约
 */
contract AuctionFactoryContract is UUPSUpgradeable, OwnableUpgradeable {

    address[] public auctions;

    mapping (uint256 tokenId => AuctionContract) public auctionMap;

    // 创建拍卖合约日志事件
    event AuctionCreated(address indexed auctionAddress,uint256 tokenId);

    // 禁用构造函数
    // constructor(){ _disableInitializers(); }

    // 初始化合约信息
    function initialize() public initializer {
        __Ownable_init(msg.sender);
        __UUPSUpgradeable_init(); // UUPS初始化
    }

    // 实现升级接口
    function _authorizeUpgrade(
        address newImplementation
    ) internal virtual override onlyOwner {
        // Only allow the owner to upgrade the implementation
    }

        // Create a new auction
    function createAuction(
        uint256 _startPrice,
        uint256 _startTime,
        uint256 _endTime,
        address _nftContract,
        uint256 tokenId
    ) external onlyOwner returns (address) {
        AuctionContract auction = new AuctionContract();
        auction.initialize(
            _startPrice,
            _startTime,
            _endTime,
            _nftContract,
            tokenId,
            msg.sender
        );
        auctions.push(address(auction));
        auctionMap[tokenId] = auction;

        emit AuctionCreated(address(auction), tokenId);
        return address(auction);
    }

    function getAuctions() external view returns (address[] memory) {
        return auctions;
    }

    function getAuction(uint256 tokenId) external view returns (AuctionContract) {
        return auctionMap[tokenId];
    }
} 
