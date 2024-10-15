package persistence

import (
	"PowerLedgerGo/domain/entity"
	"PowerLedgerGo/domain/repository"
	"errors"
	"gorm.io/gorm"
	"time"
)

var _ repository.TradeRepository = (*TradeRepositoryImpl)(nil)

type TradeRepositoryImpl struct {
	DB *gorm.DB `inject:""`
}

const (
	TradeStatusWaitBuyerConfirm = iota
	TradeStatusSuccess
	TradeStatusCancel
)

func (t *TradeRepositoryImpl) QueryExpired() ([]*entity.Trade, error) {
	var trades []*entity.Trade
	oneDayAgo := time.Now().Add(-24 * time.Hour)

	if err := t.DB.Where("created_at < ? AND trade_status = ?", oneDayAgo, 0).Find(&trades).Error; err != nil {
		return nil, err
	}

	return trades, nil
}

func (t *TradeRepositoryImpl) Save(trade *entity.Trade) error {
	return t.DB.Save(trade).Error
}

func (t *TradeRepositoryImpl) UpdateStatus(traceId int64, traceStatus int64) error {
	return t.DB.Model(&entity.Trade{}).Where("trade_id = ?", traceId).Update("trade_status", traceStatus).Error
}

func (t *TradeRepositoryImpl) QueryByBuyer(addr string, mode int64) ([]*entity.Trade, error) {
	var trades []*entity.Trade
	if mode == 1 {
		if err := t.DB.Where("buyer = ?", addr).Find(&trades).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	} else {
		if err := t.DB.Where("seller = ?", addr).Find(&trades).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}
	return trades, nil
}
