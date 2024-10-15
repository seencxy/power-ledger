package entity

import "gorm.io/gorm"

// UserInfo 用户信息
type UserInfo struct {
	gorm.Model
	Address  string `json:"address"`
	Prv      string `json:"prv"`
	UserName string `json:"userName"`
	PassWord string `json:"passWord"`
	Mode     uint8  `json:"mode"`
	Identify uint8  `json:"identify"`
}

// TableName 用户信息表
func (UserInfo) TableName() string {
	return "userInfos" // 自定义表名
}
