// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "dependency/Ownable.sol";
import "dependency/ReentrancyGuard.sol";
import "HomomorphicEncryption.sol";

contract AdvancedVirtualPowerPlantDAO is Ownable, ReentrancyGuard {
    HomomorphicEncryption private immutable encryptionLib;

    enum SettlementMode {
        BILL_DEDUCTION,
        CAPITAL_SETTLEMENT
    }

    enum Identifier {
        SELLER,
        BUYER
    }

    struct Participant {
        address addr;
        SettlementMode mode;
        uint256 balance;
        uint256 reputationScore;
        Identifier identifier;
    }

    struct EncryptedOffer {
        address seller;
        uint256 amount; // in kWh
        uint256 encryptedPrice;
        bool active;
    }

    struct EncryptedBid {
        address buyer;
        uint256 amount; // in kWh
        uint256 encryptedPrice;
        bool active;
    }

    struct PrivateTrade {
        address seller;
        address buyer;
        uint256 amount;
        uint256 encryptedPrice;
        bool settled;
        bool canceled;
    }

    mapping(address => Participant) public participants;
    EncryptedOffer[] private offers;
    EncryptedBid[] private bids;
    PrivateTrade[] private trades;

    uint256 public constant MIN_REPUTATION_SCORE = 80;

    event ParticipantRegistered(
        address indexed participant,
        SettlementMode mode
    );
    event OfferSubmitted(
        uint256 indexed offerId,
        address indexed seller,
        uint256 amount
    );
    event BidSubmitted(
        uint256 indexed bidId,
        address indexed buyer,
        uint256 amount
    );
    event TradeExecuted(
        uint256 indexed tradeId,
        address indexed seller,
        address indexed buyer,
        uint256 amount
    );
    event TradeSettled(uint256 indexed tradeId, SettlementMode mode);
    event BalanceUpdated(address indexed participant, uint256 newBalance);
    event BalanceWithdrawn(
        address indexed participant,
        uint256 withdrawnAmount
    );
    event TradeCanceled(uint256 indexed tradeId);

    constructor(address _encryptionLib, address initialOwner)
    Ownable(initialOwner)
    {
        encryptionLib = HomomorphicEncryption(_encryptionLib);
    }

    function registerParticipant(
        SettlementMode mode,
        address addr,
        Identifier identifier
    ) external {
        require(participants[addr].addr == address(0), "Already registered");
        participants[addr] = Participant(addr, mode, 0, 100, identifier);
        emit ParticipantRegistered(addr, mode);
    }

    function submitOffer(
        uint256 amount,
        uint256 encryptedPrice,
        address addr
    ) external {
        require(
            participants[addr].reputationScore >= MIN_REPUTATION_SCORE,
            "Insufficient reputation score"
        );
        require(
            participants[addr].identifier == Identifier.SELLER,
            "not seller"
        );

        uint256 offerId = offers.length;
        offers.push(EncryptedOffer(addr, amount, encryptedPrice, true));

        _matchTrades(true);

        emit OfferSubmitted(offerId, addr, amount);
    }

    function submitBid(
        uint256 amount,
        uint256 encryptedPrice,
        address addr
    ) external {
        require(
            participants[addr].reputationScore >= MIN_REPUTATION_SCORE,
            "Insufficient reputation score"
        );

        require(participants[addr].identifier == Identifier.BUYER, "not buyer");

        uint256 bidId = bids.length;
        bids.push(EncryptedBid(addr, amount, encryptedPrice, true));

        _matchTrades(false);

        emit BidSubmitted(bidId, addr, amount);
    }

    function _matchTrades(bool isNewOffer) internal {
        if (isNewOffer) {
            _sortBids();
            for (uint256 i = 0; i < offers.length; i++) {
                if (offers[i].active) {
                    _matchSingleOffer(i);
                }
            }
        } else {
            _sortOffers();
            for (uint256 i = 0; i < bids.length; i++) {
                if (bids[i].active) {
                    _matchSingleBid(i);
                }
            }
        }
    }

    function _matchSingleOffer(uint256 offerId) internal {
        EncryptedOffer storage offer = offers[offerId];
        for (uint256 i = 0; i < bids.length && offer.active; i++) {
            EncryptedBid storage bid = bids[i];
            if (
                bid.active &&
                encryptionLib.isGreaterThanOrEqual(
                    bid.encryptedPrice,
                    bid.buyer,
                    offer.encryptedPrice,
                    offer.seller
                )
            ) {
                _executeTrade(offer, bid);
            }
        }
    }

    function _matchSingleBid(uint256 bidId) internal {
        EncryptedBid storage bid = bids[bidId];
        for (uint256 i = 0; i < offers.length && bid.active; i++) {
            EncryptedOffer storage offer = offers[i];
            if (
                offer.active &&
                encryptionLib.isLessThanOrEqual(
                    offer.encryptedPrice,
                    offer.seller,
                    bid.encryptedPrice,
                    bid.buyer
                )
            ) {
                _executeTrade(offer, bid);
            }
        }
    }

    function _sortOffers() internal {
        for (uint256 i = 0; i < offers.length; i++) {
            for (uint256 j = i + 1; j < offers.length; j++) {
                if (
                    encryptionLib.isGreaterThan(
                    offers[i].encryptedPrice,
                    offers[i].seller,
                    offers[j].encryptedPrice,
                    offers[j].seller
                )
                ) {
                    EncryptedOffer memory temp = offers[i];
                    offers[i] = offers[j];
                    offers[j] = temp;
                }
            }
        }
    }

    function _sortBids() internal {
        for (uint256 i = 0; i < bids.length; i++) {
            for (uint256 j = i + 1; j < bids.length; j++) {
                if (
                    encryptionLib.isLessThan(
                    bids[i].encryptedPrice,
                    bids[i].buyer,
                    bids[j].encryptedPrice,
                    bids[j].buyer
                )
                ) {
                    EncryptedBid memory temp = bids[i];
                    bids[i] = bids[j];
                    bids[j] = temp;
                }
            }
        }
    }

    function _executeTrade(
        EncryptedOffer storage offer,
        EncryptedBid storage bid
    ) internal {
        uint256 tradeAmount = _min(offer.amount, bid.amount);
        uint256 encryptedTradePrice = encryptionLib.average(
            bid.encryptedPrice,
            bid.buyer,
            offer.encryptedPrice,
            offer.seller
        );

        trades.push(
            PrivateTrade(
                offer.seller,
                bid.buyer,
                tradeAmount,
                encryptedTradePrice,
                false,
                false
            )
        );
        emit TradeExecuted(
            trades.length - 1,
            offer.seller,
            bid.buyer,
            tradeAmount
        );

        offer.amount -= tradeAmount;
        bid.amount -= tradeAmount;

        if (offer.amount == 0) offer.active = false;
        if (bid.amount == 0) bid.active = false;
    }

    function settleTradePayments(uint256 tradeId, address addr)
    external
    nonReentrant
    {
        PrivateTrade storage trade = trades[tradeId];
        require(!trade.settled, "Trade already settled");
        require(addr == trade.buyer || addr == owner(), "Unauthorized");
        uint256 payment = trade.amount *
                            encryptionLib.decrypt(
                trade.encryptedPrice,
                uint256(uint160(trade.buyer))
            );
        Participant storage buyer = participants[trade.buyer];
        Participant storage seller = participants[trade.seller];
        require(buyer.balance >= payment, "Insufficient balance");

        if (seller.mode == SettlementMode.BILL_DEDUCTION) {
            buyer.balance -= trade.amount;
            emit BalanceWithdrawn(seller.addr, trade.amount);
        } else {
            buyer.balance -= payment;
            seller.balance += payment;
        }

        trade.settled = true;
        _updateReputationScore(trade.buyer, 1);

        emit TradeSettled(tradeId, buyer.mode);
        emit BalanceUpdated(trade.seller, seller.balance);
        emit BalanceUpdated(trade.buyer, buyer.balance);
    }

    function withdrawBalance(address addr) external nonReentrant {
        Participant storage participant = participants[addr];
        require(participant.balance > 0, "No balance to withdraw");
        uint256 amount = participant.balance;
        participant.balance = 0;
        emit BalanceUpdated(addr, 0);
        emit BalanceWithdrawn(addr, amount);
    }

    function deposit(uint256 amount, address addr)
    external
    payable
    nonReentrant
    {
        participants[addr].balance += amount;
        emit BalanceUpdated(addr, participants[addr].balance);
    }

    function _updateReputationScore(address participant, int256 change)
    internal
    {
        Participant storage p = participants[participant];
        if (
            change > 0 &&
            p.reputationScore < type(uint256).max - uint256(change)
        ) {
            p.reputationScore += uint256(change);
        } else if (change < 0 && p.reputationScore > uint256(-change)) {
            p.reputationScore -= uint256(-change);
        }
    }

    function getTradeAmount(uint256 tradeId) external view returns (uint256) {
        return trades[tradeId].amount;
    }

    function _min(uint256 a, uint256 b) internal pure returns (uint256) {
        return a < b ? a : b;
    }

    function cancelTrade(uint256 tradeId) external onlyOwner {
        PrivateTrade storage trade = trades[tradeId];
        require(!trade.settled, "Trade already settled");
        require(!trade.canceled, "Trade already canceled");

        _restoreBidOffer(trade);

        trade.canceled = true;
        _updateReputationScore(trade.buyer, -5);

        emit TradeCanceled(tradeId);
    }

    function _restoreBidOffer(PrivateTrade storage trade) internal {
        for (uint256 i = 0; i < bids.length; i++) {
            if (bids[i].buyer == trade.buyer && !bids[i].active) {
                bids[i].active = true;
                bids[i].amount += trade.amount;
                break;
            }
        }

        for (uint256 i = 0; i < offers.length; i++) {
            if (offers[i].seller == trade.seller && !offers[i].active) {
                offers[i].active = true;
                offers[i].amount += trade.amount;
                break;
            }
        }
    }
}