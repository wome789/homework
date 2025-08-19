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
contract AuctionFactoryContractCopy is UUPSUpgradeable, OwnableUpgradeable {

    // tokenId 拍卖合约实现类
    mapping (uint256  => address) auctionAddress;

    // 创建拍卖合约日志事件
    event AuctionCreated(address indexed auctionAddress,uint256 tokenId);

    // 禁用构造函数
    // constructor(){ _disableInitializers(); }

    // 初始化合约信息
    function initialize() public initializer {
        __Ownable_init(msg.sender);
        __UUPSUpgradeable_init(); // UUPS初始化
    }

    // 设置拍卖合约的实现类
    function setAuctionImplementation(uint256 _tokenId,address _auctionImplementation) public onlyOwner {
        auctionAddress[_tokenId] = _auctionImplementation;
    }

    // 实现升级接口
    function _authorizeUpgrade(
        address newImplementation
    ) internal virtual override onlyOwner {
        // Only allow the owner to upgrade the implementation
    }

    // 创建拍卖合约 调用合约的createAuction方法
    /* function createAuction(uint256 _tokenId,address _tokenAddress,address _owner,uint256 _startPrice,uint256 _endPrice,uint256 _startTime,uint256 _endTime) public {
        AuctionContract auction = AuctionContract(auctionAddress[_tokenId]);
        // auction.createAuction(_tokenAddress,_owner,_startPrice,_endPrice,_startTime,_endTime); // 调用合约的createAuction方法
        // emit AuctionCreated(address(auction),_tokenId);
    }*/
} 
