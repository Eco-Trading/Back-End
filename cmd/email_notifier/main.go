package main

/*
Author: Diego Morais
Data: 10/20/2023
Description: AWS lambda main class email notifier
*/
import (
	"context"
	helpers_email "github.com/Eco-Trading/internal/infra/helpers"
	"github.com/Eco-Trading/internal/infra/models"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var emailHelper helpers_email.IEmailHelper

func HandleRequest(_ context.Context, req models.MessageRequest) (events.APIGatewayProxyResponse, error) {
	return emailHelper.Send(req)
}

func main() {
	emailHelper = helpers_email.New()
	lambda.Start(HandleRequest)
}
