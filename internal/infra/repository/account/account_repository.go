package repository_account

/*
Author: Diego Morais
Data: 10/20/2023
Description: Class mapper database
*/
import (
	database_mongo "github.com/Eco-Trading/internal/infra/database"
	"github.com/Eco-Trading/internal/infra/repository"
	"sync"
)

type IAccountRepository interface {
	Create(account AccountMapper) (string, error)
	FindByEmail(email string) (*AccountMapper, error)
}

var iAccountRepository IAccountRepository

type accountRepository struct {
	repository.GenericCommandRepository[AccountMapper]
	repository.GenericQueryRepository[AccountMapper]
}

func New() IAccountRepository {
	lock := &sync.Mutex{}
	if iAccountRepository == nil {
		lock.Lock()
		defer lock.Unlock()
		iAccountRepository = &accountRepository{
			repository.GenericCommandRepository[AccountMapper]{
				MongoDataSource: database_mongo.New(),
				DatabaseName:    "ecotrading",
				TableName:       "account",
			},
			repository.GenericQueryRepository[AccountMapper]{
				MongoDataSource: database_mongo.New(),
				DatabaseName:    "ecotrading",
				TableName:       "account",
			},
		}
	}
	return iAccountRepository
}

func (a accountRepository) Create(account AccountMapper) (string, error) {
	return a.GenericCommandRepository.Save(account)
}

func (a accountRepository) FindByEmail(email string) (*AccountMapper, error) {
	return a.GenericQueryRepository.FindByCustomized("email", email, "email", "desc")
}
