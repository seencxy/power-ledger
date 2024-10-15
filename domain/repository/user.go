package repository

import (
	"PowerLedgerGo/domain/entity"
)

type UserRepository interface {
	Save(*entity.UserInfo) error
	Query(id int64) (*entity.UserInfo, error)
	QueryByUserName(userName string) (*entity.UserInfo, error)
}
