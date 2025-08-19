// SPDX-License-Identifier: SEE LICENSE IN LICENSE
pragma solidity ^0.8;

import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {AggregatorV3Interface} from "@chainlink/contracts/src/v0.8/shared/interfaces/AggregatorV3Interface.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721.sol";
import "hardhat/console.sol";

contract AuctionContract is Initializable, OwnableUpgradeable {
    struct Auction {
        // 卖家
        address seller;
        // 起拍价
        uint256 startPrice;
        // 拍卖开始时间
        uint256 startTime;
        // 拍卖结束时间
        uint256 endTime;
        // 拍卖是否结束
        bool ended;
        // 是否有人参与竞拍
        bool hasBid;
        // 最高出价
        uint highestBid;
        // 最高出价者
        address highestBidder;
        // 最高出价者货币地址
        address highestBidderCurrency;
        // NFT合约地址
        address nftContract;
        // NFT ID
        uint256 nftId;
    }
    Auction auction;
    address ethAdderss = 0x694AA1769357215DE4FAC081bf1f309aDC325306;

    // 初始化合约信息
    function initialize(
        uint256 _startPrice,
        uint256 _startTime,
        uint256 _endTime,
        address _nftContract,
        uint256 _nftId,
        address _owner
    ) public initializer {
        __Ownable_init(_owner);

        auction = Auction({
            seller: _owner,
            startPrice: _startPrice,
            startTime: _startTime,
            endTime: _endTime,
            ended: false,
            hasBid: false,
            highestBid: 0,
            highestBidder: address(0),
            highestBidderCurrency: address(0),
            nftContract: _nftContract,
            nftId: _nftId
        });
    }

    function getAuction() public view returns (Auction memory) {
        return auction;
    }

    // chainlink 喂价
    AggregatorV3Interface internal priceFeed;

    mapping(address => AggregatorV3Interface) priceFeeds;

    function setPriceFeed(address coinAddress) external {
        priceFeeds[coinAddress] = AggregatorV3Interface(coinAddress);
    }

    // ETH->USD     3787    50390000
    // BTC->USD     118078  17051200
    // DAI->USD     0       999 84557
    // USDC->USD    0       99985930
    // function getLatestPrice(address coinAddress) public view returns (uint) {
    //     (
    //         /* uint80 roundId */ int256 answer 
    //         /*uint256 startedAt*/
    //         /*uint256 updatedAt*/ 
    //         /*uint80 answeredInRound*/
    //     ) = priceFeeds[coinAddress].latestRoundData();
    //     return uint(answer);
    // }

    // 假设价格都是1
    function getLatestPrice(address coinAddress) public view returns (uint) {
        if (coinAddress == ethAdderss) {
            return 2;
        }else {
            return 1;
        }
    }

    // 用户参加拍卖 接收多种渠道的ERC20标准货币 如果是ETH就直接取Value 否则取要传入货币地址和货币数量。 设计预言机 转换为USD 进行计算价格
    function bid(uint _amount, address _currency) external payable {
        require(
            msg.sender != auction.seller,
            "Seller cannot bid on their own auction"
        );
        require(
            block.timestamp >= auction.startTime,
            "Auction has not started yet"
        );
        require(block.timestamp <= auction.endTime, "Auction has ended");

        // 使用预言机计算货币价格
        uint price;
        if (_currency != address(0) && _currency != ethAdderss) {
            price = _amount * getLatestPrice(_currency);
        } else {
            _amount = msg.value;
            price = _amount * getLatestPrice(ethAdderss);
        }
        uint startPrice = auction.startPrice * getLatestPrice(ethAdderss);

        // 是否有人出价了 如果没有人出价 只要判断传入的货币数量是否大于起拍价
        if (auction.hasBid) {
            // 计算最高出价金额
            uint highestBidPrice;
            if (auction.hasBid) {
                highestBidPrice =
                    auction.highestBid *
                    getLatestPrice(auction.highestBidderCurrency);
            }
            require(
                price >= startPrice && price >= highestBidPrice,
                "Bid amount is too low"
            );
        } else {
            require(price >= startPrice, "Bid amount is too low");
        }

        // 如果有人出价，退还最高出价者的货币
        if (auction.hasBid) {
            if (auction.highestBidderCurrency != ethAdderss) {
                IERC20(auction.highestBidderCurrency).transfer(
                    auction.highestBidder,
                    auction.highestBid
                );
            } else {
                payable(auction.highestBidder).transfer(auction.highestBid);
            }
        }

        // 更新最高出价和最高出价者
        auction.highestBid = _amount;
        auction.highestBidder = msg.sender;
        auction.highestBidderCurrency = _currency;
        auction.hasBid = true;
    }

    // 拍卖结束
    function endAuction() external onlyOwner {
        // require(block.timestamp > auction.endTime, "Auction has not ended yet");
        require(!auction.ended, "Auction has already ended");

        // 如果有人出价，将NFT转移到最高出价者
        if (auction.hasBid) {
            IERC721(auction.nftContract).transferFrom(auction.seller,auction.highestBidder,auction.nftId);
        }

        // 将合约的余额转移到管理员
        if (address(this).balance > 0) {
            payable(msg.sender).transfer(auction.highestBid);
        }

        // 将拍卖标记为已结束
        auction.ended = true;
    }
}
