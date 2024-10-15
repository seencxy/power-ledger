package repository

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type ContractRepository interface {
	Register(addr string, mode uint8, identify uint8) error
	QueryBalance(addr string) (int64, error)
	Recharge(addr string, amount int64) error
	Withdraw(addr string) error
	Encrypt(mount int64, addr string) (*big.Int, error)
	SubmitBid(mount int64, encryptPrice *big.Int, addr string) (string, error)
	SubmitOffer(mount int64, encryptPrice *big.Int, addr string) (string, error)
	WatchTradeExecuted(nextOperate func(*big.Int, common.Address, common.Address, *big.Int)) error
	CancelTrade(traceId int64) error
	TradePayment(addr string, traceId int64) error
	WatchTradeSettled(nextOperate func(*big.Int)) error
	WatchUserWithdraw() error
}
