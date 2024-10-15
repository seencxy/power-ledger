package application

import (
	"PowerLedgerGo/domain/entity"
	"PowerLedgerGo/domain/repository"
	"PowerLedgerGo/infrastructure/persistence"
	"github.com/ethereum/go-ethereum/common"
	"gorm.io/gorm"
	"math/big"
)

type MainService struct {
	DB           *gorm.DB                      `inject:""`
	UserRepo     repository.UserRepository     `inject:""`
	ContractRepo repository.ContractRepository `inject:""`
	TraceRepo    repository.TradeRepository    `inject:""`
}

func (u *MainService) SubmitBid(price, mount, id int64) (string, error) {
	user, err := u.UserRepo.Query(id)
	if err != nil {
		return "", err
	}

	encryptPrice, err := u.ContractRepo.Encrypt(price, user.Address)
	if err != nil {
		return "", err
	}

	return u.ContractRepo.SubmitBid(mount, encryptPrice, user.Address)
}

func (u *MainService) SubmitOffer(price, mount, id int64) (string, error) {
	user, err := u.UserRepo.Query(id)
	if err != nil {
		return "", err
	}

	encryptPrice, err := u.ContractRepo.Encrypt(price, user.Address)
	if err != nil {
		return "", err
	}

	return u.ContractRepo.SubmitOffer(mount, encryptPrice, user.Address)
}

func (u *MainService) StartWatchTradeExecuted() {
	_ = u.ContractRepo.WatchTradeExecuted(func(traceId *big.Int, seller common.Address, buyer common.Address, amount *big.Int) {
		traceInfo := entity.Trade{
			TradeId: traceId.Int64(),
			Seller:  seller.String(),
			Buyer:   buyer.String(),
			Amount:  amount.Int64(),
		}
		_ = u.TraceRepo.Save(&traceInfo)
	})
}

func (u *MainService) CheckTraceExpired() error {
	expiredTrace, err := u.TraceRepo.QueryExpired()
	if err != nil {
		return err
	}
	for _, element := range expiredTrace {
		if err = u.TraceRepo.UpdateStatus(element.TradeId, persistence.TradeStatusCancel); err != nil {
			return err
		}
		if err = u.ContractRepo.CancelTrade(element.TradeId); err != nil {
			return err
		}
	}

	return nil
}

func (u *MainService) QueryTradeBySelf(id int64, mode int64) ([]*entity.Trade, error) {
	userInfo, err := u.UserRepo.Query(id)
	if err != nil {
		return nil, err
	}
	return u.TraceRepo.QueryByBuyer(userInfo.Address, mode)
}

func (u *MainService) SettleTradePayments(id int64, tradeId int64) error {
	user, err := u.UserRepo.Query(id)
	if err != nil {
		return err
	}

	return u.ContractRepo.TradePayment(user.Address, tradeId)
}

func (u *MainService) StartWatchTradePayment() {
	_ = u.ContractRepo.WatchTradeSettled(func(traceId *big.Int) {
		_ = u.TraceRepo.UpdateStatus(traceId.Int64(), persistence.TradeStatusSuccess)
	})
}

func (u *MainService) StartWatchUserWithdraw() {
	_ = u.ContractRepo.WatchUserWithdraw()
}
