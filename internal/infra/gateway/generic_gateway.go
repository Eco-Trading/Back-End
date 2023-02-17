package gateway_email_notifier

import (
	"encoding/json"
	"github.com/Eco-Trading/util"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	service "github.com/aws/aws-sdk-go/service/lambda"
)

var responseUtil util.IResponseUtil

type GenericGateway[T any] struct {
}

func (e GenericGateway[T]) Invoke(request T, lambdaInvoke string) (events.APIGatewayProxyResponse, error) {
	responseUtil = util.NewResponseUtil()

	newSession := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	client := service.New(newSession)

	personJSON, _ := json.Marshal(request)

	input := &service.InvokeInput{
		FunctionName: aws.String(lambdaInvoke),
		Payload:      personJSON,
	}

	result, err := client.Invoke(input)

	if err != nil {
		return responseUtil.Create(err.Error(), 500), err
	}
	return responseUtil.Create(string(result.Payload), 200), nil
}
