package main

/*
Author: Diego Morais
Data: 10/20/2023
Description: AWS lambda main class create account
*/
import (
	"context"
	controller_account "github.com/Eco-Trading/internal/infra/controller"
	"github.com/Eco-Trading/internal/infra/models"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var accountController controller_account.IAccountController

func HandleRequest(_ context.Context, req models.AccountRequest) (events.APIGatewayProxyResponse, error) {
	return accountController.Create(req)
}

func main() {
	accountController = controller_account.New()
	lambda.Start(HandleRequest)
}
