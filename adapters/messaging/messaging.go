package messaging

import (
	"api.default.marincor/clients/google/pubsub"
)

func Publish(queueName string, message interface{}) {
	pubsub.Publish(queueName, message)
}
