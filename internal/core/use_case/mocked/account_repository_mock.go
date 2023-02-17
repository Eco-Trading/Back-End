package mocked

import (
	repository_account "github.com/Eco-Trading/internal/infra/repository/account"
	"github.com/Eco-Trading/util/mocked"
)

type IAccountRepositoryMock struct {
	mocked.GenericMock[IAccountRepositoryMock]
}

func (a *IAccountRepositoryMock) Create(account repository_account.AccountMapper) (string, error) {
	args := a.Called()
	return args.String(0), args.Error(1)
}

func (a *IAccountRepositoryMock) FindByEmail(email string) (*repository_account.AccountMapper, error) {
	args := a.Called()
	return args.Get(0).(*repository_account.AccountMapper), args.Error(1)
}
