package usecase_account

import (
	"errors"
	"github.com/Eco-Trading/internal/core/use_case/mocked"
	"github.com/Eco-Trading/internal/infra/models"
	repository_account "github.com/Eco-Trading/internal/infra/repository/account"
	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
	"time"
)

func TestAccountInteractor_must_create_an_account_with_success(t *testing.T) {

	assertions := assert.New(t)

	iEmailNotifierGatewayMock := new(mocked.IEmailNotifierGatewayMock)
	iEmailNotifierGatewayMock.On("Invoke").Return(events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "email send",
	}, nil)

	var accountMapperNoFound *repository_account.AccountMapper
	iAccountRepositoryMock := new(mocked.IAccountRepositoryMock)
	iAccountRepositoryMock.On("FindByEmail").Return(accountMapperNoFound, nil)
	iAccountRepositoryMock.On("Create").Return("aadfc7fda9800543", nil)

	iAccountInteractor := New().(*accountInteractor)
	iAccountInteractor.emailNotifierGateway = iEmailNotifierGatewayMock
	iAccountInteractor.accountRepository = iAccountRepositoryMock

	req := models.AccountRequest{
		State:     "São paulo",
		Email:     "diego.morais@gmail.com",
		Phone:     "+55 (19) 9 9999-9999",
		LastName:  "Morais",
		Address:   "Rua Teste",
		FirstName: "Diego",
		City:      "São paulo",
	}

	response, err := iAccountInteractor.Create(req)

	assertions.NoError(err)
	assertions.Equal(200, response.StatusCode)
	assertions.Equal("{\"message\":\"account created success\"}", response.Body)

	iAccountRepositoryMock.AssertNumberOfCallsTime(t, "FindByEmail", 1)
	iAccountRepositoryMock.AssertNumberOfCallsTime(t, "Create", 1)
	iEmailNotifierGatewayMock.AssertNumberOfCallsTime(t, "Invoke", 1)

}

func TestAccountInteractor_must_create_an_account_data_nat_found(t *testing.T) {

	assertions := assert.New(t)

	iEmailNotifierGatewayMock := new(mocked.IEmailNotifierGatewayMock)
	iEmailNotifierGatewayMock.On("Invoke").Return(events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "email send",
	}, nil)

	var accountMapperNoFound *repository_account.AccountMapper
	iAccountRepositoryMock := new(mocked.IAccountRepositoryMock)
	iAccountRepositoryMock.On("FindByEmail").Return(accountMapperNoFound, nil)
	iAccountRepositoryMock.On("Create").Return("aadfc7fda9800543", nil)

	iAccountInteractor := New().(*accountInteractor)
	iAccountInteractor.emailNotifierGateway = iEmailNotifierGatewayMock
	iAccountInteractor.accountRepository = iAccountRepositoryMock

	req := models.AccountRequest{
		Email:     "diego.morais@gmail.com",
		Phone:     "+55 (19) 9 9999-9999",
		LastName:  "Morais",
		Address:   "Rua Teste",
		FirstName: "Diego",
		City:      "São paulo",
	}

	response, err := iAccountInteractor.Create(req)

	assertions.EqualError(err, "state is not filled")
	assertions.Equal(409, response.StatusCode)
	assertions.Equal("{\"message\":\"state is not filled\"}", response.Body)

}

func TestAccountInteractor_must_create_an_account_email_already(t *testing.T) {

	assertions := assert.New(t)

	iEmailNotifierGatewayMock := new(mocked.IEmailNotifierGatewayMock)
	iEmailNotifierGatewayMock.On("Invoke").Return(events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "email send",
	}, nil)

	accountMapperFound := &repository_account.AccountMapper{
		ID:            primitive.NewObjectID(),
		AccountNumber: "9999999",
		CreateAt:      time.Now(),
		Balance:       19989.90,
		State:         "São paulo",
		Email:         "diego.morais@gmail.com",
		Phone:         "+55 (19) 9 9999-9999",
		LastName:      "Morais",
		Address:       "Rua Teste",
		FirstName:     "Diego",
		City:          "São paulo",
	}
	iAccountRepositoryMock := new(mocked.IAccountRepositoryMock)
	iAccountRepositoryMock.On("FindByEmail").Return(accountMapperFound, nil)
	iAccountRepositoryMock.On("Create").Return("aadfc7fda9800543", nil)

	iAccountInteractor := New().(*accountInteractor)
	iAccountInteractor.emailNotifierGateway = iEmailNotifierGatewayMock
	iAccountInteractor.accountRepository = iAccountRepositoryMock

	req := models.AccountRequest{
		State:     "São paulo",
		Email:     "diego.morais@gmail.com",
		Phone:     "+55 (19) 9 9999-9999",
		LastName:  "Morais",
		Address:   "Rua Teste",
		FirstName: "Diego",
		City:      "São paulo",
	}

	response, err := iAccountInteractor.Create(req)

	assertions.EqualError(err, "email already exists")
	assertions.Equal(409, response.StatusCode)
	assertions.Equal("{\"message\":\"email already exists\"}", response.Body)

	iAccountRepositoryMock.AssertNumberOfCallsTime(t, "FindByEmail", 1)

}

func TestAccountInteractor_must_create_an_account_erro_when_creating_account(t *testing.T) {

	assertions := assert.New(t)

	iEmailNotifierGatewayMock := new(mocked.IEmailNotifierGatewayMock)
	iEmailNotifierGatewayMock.On("Invoke").Return(events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "email send",
	}, nil)

	var accountMapperNoFound *repository_account.AccountMapper
	iAccountRepositoryMock := new(mocked.IAccountRepositoryMock)
	iAccountRepositoryMock.On("FindByEmail").Return(accountMapperNoFound, nil)
	iAccountRepositoryMock.On("Create").Return("aadfc7fda9800543", errors.New("connect not found"))

	iAccountInteractor := New().(*accountInteractor)
	iAccountInteractor.emailNotifierGateway = iEmailNotifierGatewayMock
	iAccountInteractor.accountRepository = iAccountRepositoryMock

	req := models.AccountRequest{
		State:     "São paulo",
		Email:     "diego.morais@gmail.com",
		Phone:     "+55 (19) 9 9999-9999",
		LastName:  "Morais",
		Address:   "Rua Teste",
		FirstName: "Diego",
		City:      "São paulo",
	}

	response, err := iAccountInteractor.Create(req)

	assertions.EqualError(err, "connect not found")
	assertions.Equal(500, response.StatusCode)
	assertions.Equal("{\"message\":\"connect not found\"}", response.Body)

	iAccountRepositoryMock.AssertNumberOfCallsTime(t, "FindByEmail", 1)
	iAccountRepositoryMock.AssertNumberOfCallsTime(t, "Create", 1)

}

func TestAccountInteractor_must_create_an_account_erro_when_sending_email(t *testing.T) {

	assertions := assert.New(t)

	iEmailNotifierGatewayMock := new(mocked.IEmailNotifierGatewayMock)
	iEmailNotifierGatewayMock.On("Invoke").Return(events.APIGatewayProxyResponse{
		StatusCode: 500,
		Body:       "connect not found",
	}, errors.New("connect not found"))

	var accountMapperNoFound *repository_account.AccountMapper
	iAccountRepositoryMock := new(mocked.IAccountRepositoryMock)
	iAccountRepositoryMock.On("FindByEmail").Return(accountMapperNoFound, nil)
	iAccountRepositoryMock.On("Create").Return("aadfc7fda9800543", nil)

	iAccountInteractor := New().(*accountInteractor)
	iAccountInteractor.emailNotifierGateway = iEmailNotifierGatewayMock
	iAccountInteractor.accountRepository = iAccountRepositoryMock

	req := models.AccountRequest{
		State:     "São paulo",
		Email:     "diego.morais@gmail.com",
		Phone:     "+55 (19) 9 9999-9999",
		LastName:  "Morais",
		Address:   "Rua Teste",
		FirstName: "Diego",
		City:      "São paulo",
	}

	response, err := iAccountInteractor.Create(req)

	assertions.EqualError(err, "connect not found")
	assertions.Equal(500, response.StatusCode)
	assertions.Equal("{\"message\":\"connect not found\"}", response.Body)

	iAccountRepositoryMock.AssertNumberOfCallsTime(t, "FindByEmail", 1)
	iAccountRepositoryMock.AssertNumberOfCallsTime(t, "Create", 1)
	iEmailNotifierGatewayMock.AssertNumberOfCallsTime(t, "Invoke", 1)

}
