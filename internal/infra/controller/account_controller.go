package controller_account

/*
Author: Diego Morais
Data: 10/20/2023
Description: Class responsible for the account controller layer
*/
import (
	usecase_account "github.com/Eco-Trading/internal/core/use_case"
	"github.com/Eco-Trading/internal/infra/models"
	"github.com/aws/aws-lambda-go/events"
	"sync"
)

type IAccountController interface {
	Create(req models.AccountRequest) (events.APIGatewayProxyResponse, error)
}

var controller IAccountController

type accountController struct {
	accountInteractor usecase_account.IAccountInteractor
}

func New() IAccountController {
	lock := &sync.Mutex{}
	if controller == nil {
		lock.Lock()
		defer lock.Unlock()
		controller = accountController{
			accountInteractor: usecase_account.New(),
		}
	}
	return controller
}

func (a accountController) Create(req models.AccountRequest) (events.APIGatewayProxyResponse, error) {
	return a.accountInteractor.Create(req)
}
