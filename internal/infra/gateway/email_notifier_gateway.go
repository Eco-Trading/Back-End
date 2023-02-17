package gateway_email_notifier

import (
	"github.com/Eco-Trading/internal/infra/enum"
	"github.com/Eco-Trading/internal/infra/models"
	"github.com/aws/aws-lambda-go/events"
	"sync"
)

type IEmailNotifierGateway interface {
	Invoke(request models.MessageRequest) (events.APIGatewayProxyResponse, error)
}

var iEmailNotifierGateway IEmailNotifierGateway

type emailNotifierGateway struct {
	GenericGateway[models.MessageRequest]
}

func New() IEmailNotifierGateway {
	lock := &sync.Mutex{}
	if iEmailNotifierGateway == nil {
		lock.Lock()
		defer lock.Unlock()
		iEmailNotifierGateway = emailNotifierGateway{}

	}
	return iEmailNotifierGateway
}

func (e emailNotifierGateway) Invoke(request models.MessageRequest) (events.APIGatewayProxyResponse, error) {
	return e.GenericGateway.Invoke(request, enum.EmailNotifier)
}
