package entity

import "gorm.io/gorm"

// Trade 存储交易信息表
type Trade struct {
	gorm.Model
	TradeId     int64  `json:"tradeId" gorm:"unique;not null"`
	Seller      string `json:"seller"`
	Buyer       string `json:"buyer"`
	Amount      int64  `json:"amount"`
	TradeStatus int64  `json:"tradeStatus"` // 交易状态: 0 等待买家确认 1 交易成功 2 交易取消
}

func (Trade) TableName() string {
	return "trades"
}
