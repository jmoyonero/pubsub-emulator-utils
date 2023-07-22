package main

import (
	"cloud.google.com/go/pubsub"
	log "github.com/sirupsen/logrus"
)

func CreateSubscription(projectID string, topicID string, subscriptionID string, cl *log.Entry) error {
	client, err := pubsub.NewClient(cl.Context, projectID)
	if err != nil {
		return err
	}

	subscription, err := client.CreateSubscription(cl.Context, subscriptionID, pubsub.SubscriptionConfig{
		Topic: client.Topic(topicID),
	})
	if err != nil {
		return err
	}

	log.Infof("subscription [%s] in topic [%s] in project [%s] successfully created", subscription.ID(), topicID, projectID)
	return nil
}
