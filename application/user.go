package application

import (
	"PowerLedgerGo/domain/entity"
	"PowerLedgerGo/domain/repository"
	"gorm.io/gorm"
)

type UserService struct {
	DB           *gorm.DB                      `inject:""`
	UserRepo     repository.UserRepository     `inject:""`
	ContractRepo repository.ContractRepository `inject:""`
}

func (u *UserService) CreateUser(user *entity.UserInfo) error {
	if err := u.UserRepo.Save(user); err != nil {
		return err
	}
	return u.ContractRepo.Register(user.Address, user.Mode, user.Identify)
}

func (u *UserService) QueryUserByUsername(userName string) (*entity.UserInfo, error) {
	user, err := u.UserRepo.QueryByUserName(userName)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserService) QueryBalance(id int64) (int64, error) {
	user, err := u.UserRepo.Query(id)
	if err != nil {
		return 0, err
	}

	balance, err := u.ContractRepo.QueryBalance(user.Address)
	if err != nil {
		return 0, err
	}
	return balance, nil
}

func (u *UserService) Recharge(id int64, amount int64) error {
	user, err := u.UserRepo.Query(id)
	if err != nil {
		return err
	}
	return u.ContractRepo.Recharge(user.Address, amount)
}

func (u *UserService) Withdraw(id int64) error {
	user, err := u.UserRepo.Query(id)
	if err != nil {
		return err
	}
	return u.ContractRepo.Withdraw(user.Address)
}
