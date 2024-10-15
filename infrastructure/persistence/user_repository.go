package persistence

import (
	"PowerLedgerGo/domain/entity"
	"PowerLedgerGo/domain/repository"
	"gorm.io/gorm"
)

var _ repository.UserRepository = (*UserRepositoryImpl)(nil)

// UserRepositoryImpl Implements repository.UserRepository
type UserRepositoryImpl struct {
	DB *gorm.DB `inject:""`
}

func (u *UserRepositoryImpl) Save(user *entity.UserInfo) error {
	return u.DB.Save(user).Error
}

func (u *UserRepositoryImpl) Query(id int64) (*entity.UserInfo, error) {
	var user entity.UserInfo
	tx := u.DB.First(&user, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &user, nil
}

func (u *UserRepositoryImpl) QueryByUserName(userName string) (*entity.UserInfo, error) {
	var user entity.UserInfo
	tx := u.DB.Where("user_name = ?", userName).First(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &user, nil
}
