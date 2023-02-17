package helpers_email

/*
Author: Diego Morais
Data: 10/20/2023
Description: Send email
*/

import (
	"fmt"
	"github.com/Eco-Trading/internal/infra/models"
	"github.com/Eco-Trading/util"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"os"
	"sync"
)

type IEmailHelper interface {
	Send(event models.MessageRequest) (events.APIGatewayProxyResponse, error)
}

type emailHelper struct {
}

var emailHelp IEmailHelper
var responseUtil util.IResponseUtil

func New() IEmailHelper {
	lock := &sync.Mutex{}
	if emailHelp == nil {
		lock.Lock()
		defer lock.Unlock()
		emailHelp = &emailHelper{}
		responseUtil = util.NewResponseUtil()
	}
	return emailHelp
}

func (e emailHelper) Send(event models.MessageRequest) (events.APIGatewayProxyResponse, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION"))},
	)

	svc := ses.New(sess)

	result, err := svc.SendEmail(&ses.SendEmailInput{
		Destination: &ses.Destination{
			ToAddresses: []*string{
				aws.String(event.To),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Text: &ses.Content{
					Data: aws.String(event.Body),
				},
			},
			Subject: &ses.Content{
				Data: aws.String(event.Subject),
			},
		},
		Source: aws.String(event.From),
	})

	if err != nil {
		return responseUtil.Create(err.Error(), 409), err
	}

	return responseUtil.Create(fmt.Sprintf("Message retorn %v sucess", *result.MessageId), 409), err
}
