package main

import (
	"cloud.google.com/go/pubsub"
	log "github.com/sirupsen/logrus"
)

func CreateTopic(projectID string, topicID string, cl *log.Entry) error {
	client, err := pubsub.NewClient(cl.Context, projectID)
	if err != nil {
		return err
	}
	topic, err := client.CreateTopic(cl.Context, topicID)
	if err != nil {
		return err
	}
	log.Infof("topic [%s] in project [%s] successfully created", topic.ID(), projectID)
	return nil
}
