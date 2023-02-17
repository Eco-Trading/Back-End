package mocked

import (
	"github.com/Eco-Trading/internal/infra/models"
	"github.com/Eco-Trading/util/mocked"
	"github.com/aws/aws-lambda-go/events"
)

type IEmailNotifierGatewayMock struct {
	mocked.GenericMock[IEmailNotifierGatewayMock]
}

func (e *IEmailNotifierGatewayMock) Invoke(request models.MessageRequest) (events.APIGatewayProxyResponse, error) {
	args := e.Called()
	return args.Get(0).(events.APIGatewayProxyResponse), args.Error(1)
}
