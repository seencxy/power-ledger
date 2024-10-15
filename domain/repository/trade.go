package repository

import "PowerLedgerGo/domain/entity"

type TradeRepository interface {
	Save(trade *entity.Trade) error
	QueryExpired() ([]*entity.Trade, error)
	UpdateStatus(int64, int64) error
	QueryByBuyer(string, int64) ([]*entity.Trade, error)
}
