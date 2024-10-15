package persistence

import (
	"PowerLedgerGo/domain/repository"
	"PowerLedgerGo/infrastructure/config"
	"context"
	"fmt"
	"github.com/FISCO-BCOS/go-sdk/v3/types"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"math/big"
)

var _ repository.ContractRepository = (*ContractRepositoryImpl)(nil)

type ContractRepositoryImpl struct {
	Instance *config.ContractInstance `inject:""`
}

func (u *ContractRepositoryImpl) Register(addr string, mode uint8, identify uint8) error {
	tx, receipt, err := u.Instance.AdvancedVirtualPowerPlantDAOSession.RegisterParticipant(mode, common.HexToAddress(addr), identify)
	if err != nil {
		return fmt.Errorf("transaction failed: %w", err)
	}

	// 检查交易收据
	if receipt.Status != 0 {
		return fmt.Errorf("transaction reverted: %s,%s", receipt.TransactionHash, tx.Hash())
	}

	return nil
}

func (u *ContractRepositoryImpl) QueryBalance(addr string) (int64, error) {
	participants, err := u.Instance.AdvancedVirtualPowerPlantDAOSession.Participants(common.HexToAddress(addr))
	if err != nil {
		return 0, err
	}
	return participants.Balance.Int64(), nil
}

func (u *ContractRepositoryImpl) Recharge(addr string, amount int64) error {
	tx, receipt, err := u.Instance.AdvancedVirtualPowerPlantDAOSession.Deposit(big.NewInt(amount), common.HexToAddress(addr))
	if err != nil {
		return err
	}
	// 检查交易收据
	if receipt.Status != 0 {
		return fmt.Errorf("transaction reverted: %s,%s", receipt.TransactionHash, tx.Hash())
	}

	return nil
}

func (u *ContractRepositoryImpl) Withdraw(addr string) error {
	tx, receipt, err := u.Instance.AdvancedVirtualPowerPlantDAOSession.WithdrawBalance(common.HexToAddress(addr))
	if err != nil {
		return err
	}
	// 检查交易收据
	if receipt.Status != 0 {
		return fmt.Errorf("transaction reverted: %s,%s", receipt.TransactionHash, tx.Hash())
	}

	return nil
}

func (u *ContractRepositoryImpl) Encrypt(mount int64, addr string) (*big.Int, error) {
	return u.Instance.HomomorphicEncryptionSession.Encrypt(big.NewInt(mount), common.HexToAddress(addr).Big())
}

func (u *ContractRepositoryImpl) SubmitBid(mount int64, encryptPrice *big.Int, addr string) (string, error) {
	_, receipt, err := u.Instance.AdvancedVirtualPowerPlantDAOSession.SubmitBid(big.NewInt(mount), encryptPrice, common.HexToAddress(addr))
	if err != nil {
		return "", err
	}

	// 检查交易收据
	if receipt.Status != 0 {
		return "", fmt.Errorf("transaction reverted: %s,%s", receipt.TransactionHash)
	}

	return receipt.TransactionHash, nil
}

func (u *ContractRepositoryImpl) SubmitOffer(mount int64, encryptPrice *big.Int, addr string) (string, error) {
	_, receipt, err := u.Instance.AdvancedVirtualPowerPlantDAOSession.SubmitOffer(big.NewInt(mount), encryptPrice, common.HexToAddress(addr))
	if err != nil {
		return "", err
	}

	if receipt.Status != 0 {
		return "", fmt.Errorf("transaction reverted: %s", receipt.TransactionHash)
	}

	return receipt.TransactionHash, nil
}

// WatchTradeExecuted 监听交易执行
func (u *ContractRepositoryImpl) WatchTradeExecuted(nextOperate func(*big.Int, common.Address, common.Address, *big.Int)) error {
	signal := make(chan error)
	blockNumber, err := u.Instance.Client.GetBlockNumber(context.Background())
	if err != nil {
		return err
	}
	if _, err = u.Instance.AdvancedVirtualPowerPlantDAOSession.WatchAllTradeExecuted(&blockNumber, func(ret int, logs []types.Log) {
		for _, element := range logs {
			tradeInfo, err := u.Instance.AdvancedVirtualPowerPlantDAOSession.ParseTradeExecuted(element)
			if err != nil {
				signal <- err
			}
			fmt.Printf("TradeExecuted(%d, %s, %s, %d) \n",
				tradeInfo.TradeId, tradeInfo.Seller, tradeInfo.Buyer, tradeInfo.Amount)
			nextOperate(tradeInfo.TradeId, tradeInfo.Seller, tradeInfo.Buyer, tradeInfo.Amount)
		}
	}); err != nil {
		return err
	}
	return <-signal
}

func (u *ContractRepositoryImpl) CancelTrade(traceId int64) error {
	_, receipt, err := u.Instance.AdvancedVirtualPowerPlantDAOSession.CancelTrade(big.NewInt(traceId))
	if err != nil {
		return err
	}
	// 检查交易收据
	if receipt.Status != 0 {
		return fmt.Errorf("transaction reverted: %s,%s", receipt.TransactionHash)
	}

	return nil
}

func (u *ContractRepositoryImpl) TradePayment(addr string, traceId int64) error {
	_, receipt, err := u.Instance.AdvancedVirtualPowerPlantDAOSession.SettleTradePayments(big.NewInt(traceId), common.HexToAddress(addr))
	if err != nil {
		return err
	}
	// 检查交易收据
	if receipt.Status != 0 {
		return fmt.Errorf("transaction reverted: %s", receipt.TransactionHash)
	}

	return nil
}

// WatchTradeSettled 监听交易成功
func (u *ContractRepositoryImpl) WatchTradeSettled(nextOperate func(*big.Int)) error {
	signal := make(chan error)
	blockNumber, err := u.Instance.Client.GetBlockNumber(context.Background())
	if err != nil {
		return err
	}
	if _, err = u.Instance.AdvancedVirtualPowerPlantDAOSession.WatchAllTradeSettled(&blockNumber, func(ret int, logs []types.Log) {
		tradeInfo, err := u.Instance.AdvancedVirtualPowerPlantDAOSession.ParseTradeSettled(logs[0])
		if err != nil {
			signal <- err
		}
		fmt.Printf("TradeSettled(%d, %d) \n",
			tradeInfo.TradeId, tradeInfo.Mode)
		nextOperate(tradeInfo.TradeId)
	}); err != nil {
		return err
	}
	return <-signal
}

// WatchUserWithdraw 监听用户转账
func (u *ContractRepositoryImpl) WatchUserWithdraw() error {
	signal := make(chan error)
	blockNumber, err := u.Instance.Client.GetBlockNumber(context.Background())
	if err != nil {
		return err
	}
	if _, err = u.Instance.AdvancedVirtualPowerPlantDAOSession.WatchAllBalanceWithdrawn(&blockNumber, func(ret int, logs []types.Log) {
		withdrawInfo, err := u.Instance.AdvancedVirtualPowerPlantDAOSession.ParseBalanceWithdrawn(logs[0])
		if err != nil {
			signal <- err
		}
		// todo 处理提取余额操作
		log.Printf("提现地址: %s, 提现金额: %d", withdrawInfo.Participant.String(), withdrawInfo.WithdrawnAmount)
	}); err != nil {
		return err
	}
	return <-signal
}
