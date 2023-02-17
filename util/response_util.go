package util

import (
	"encoding/json"
	"github.com/Eco-Trading/internal/infra/models"
	"github.com/aws/aws-lambda-go/events"
	"sync"
)

type IResponseUtil interface {
	Create(message string, code int) events.APIGatewayProxyResponse
}

type responseUtil struct {
}

var response IResponseUtil

func NewResponseUtil() IResponseUtil {
	lock := &sync.Mutex{}
	if response == nil {
		lock.Lock()
		defer lock.Unlock()
		response = &responseUtil{}
	}
	return response
}

func (r responseUtil) Create(message string, code int) events.APIGatewayProxyResponse {
	value, _ := json.Marshal(models.GenericResponse{
		Message: message,
	})
	return events.APIGatewayProxyResponse{
		StatusCode: code,
		Body:       string(value),
	}
}
