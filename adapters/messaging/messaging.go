package messaging

import (
	"api.default.marincor.pt/clients/google/pubsub"
)

type Messaging struct{}

func New() *Messaging {
	return &Messaging{}
}

func (messaging *Messaging) Publish(queueName string, message interface{}) {
	pubsub.Publish(queueName, message)
}
