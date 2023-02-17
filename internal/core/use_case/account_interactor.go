package usecase_account

/*
Author: Diego Morais
Data: 10/20/2023
Description: Class responsible for account application rules
*/

import (
	"errors"
	"github.com/Eco-Trading/internal/core/entities"
	gateway_email_notifier "github.com/Eco-Trading/internal/infra/gateway"
	"github.com/Eco-Trading/internal/infra/models"
	repository_account "github.com/Eco-Trading/internal/infra/repository/account"
	"github.com/Eco-Trading/util"
	"github.com/aws/aws-lambda-go/events"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"os"
	"sync"
	"time"
)

type IAccountInteractor interface {
	Create(req models.AccountRequest) (events.APIGatewayProxyResponse, error)
}

type accountInteractor struct {
	accountRepository    repository_account.IAccountRepository
	emailNotifierGateway gateway_email_notifier.IEmailNotifierGateway
}

var interactor IAccountInteractor
var responseUtil util.IResponseUtil

func New() IAccountInteractor {
	lock := &sync.Mutex{}
	if interactor == nil {
		lock.Lock()
		defer lock.Unlock()
		interactor = &accountInteractor{
			accountRepository:    repository_account.New(),
			emailNotifierGateway: gateway_email_notifier.New(),
		}
		responseUtil = util.NewResponseUtil()
	}
	return interactor
}

func (a accountInteractor) Create(req models.AccountRequest) (events.APIGatewayProxyResponse, error) {

	account := entities.Account{
		City:      req.City,
		FirstName: req.FirstName,
		Address:   req.Address,
		Email:     req.Email,
		LastName:  req.LastName,
		Phone:     req.Phone,
		State:     req.State,
	}

	account.GenerateAccountNumber()
	_, err := account.ValidateData()

	if err != nil {
		return responseUtil.Create(err.Error(), 409), err
	}

	accountFound, _ := a.accountRepository.FindByEmail(account.Email)

	if accountFound != nil {
		msg := "email already exists"
		return responseUtil.Create(msg, 409), errors.New(msg)
	}

	accountMapper := repository_account.AccountMapper{
		ID:            primitive.NewObjectID(),
		City:          account.City,
		FirstName:     account.FirstName,
		Address:       account.Address,
		Email:         account.Email,
		LastName:      account.LastName,
		Phone:         account.Phone,
		State:         account.State,
		Balance:       0,
		CreateAt:      time.Now(),
		AccountNumber: account.AccountNumber,
	}

	_, err3 := a.accountRepository.Create(accountMapper)
	if err3 != nil {
		return responseUtil.Create(err3.Error(), 500), err3
	}

	_, err4 := a.emailNotifierGateway.Invoke(models.MessageRequest{
		From:    os.Getenv("EMAIL_ADMIN"),
		To:      account.Email,
		Subject: os.Getenv("EMAIL_SUBJECT_CREATE_ACCOUNT"),
		Body:    "Your account has been successfully created",
	})

	if err4 != nil {
		return responseUtil.Create(err4.Error(), 500), err4
	}

	return responseUtil.Create("account created success", 200), nil
}
