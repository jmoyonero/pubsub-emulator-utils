package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	log "github.com/sirupsen/logrus"
)

func main() {

	fmt.Println("----------------------------------")
	fmt.Println("Welcome To PubSub emulator utils")
	fmt.Println("----------------------------------")

	cl := log.WithContext(context.Background())

	var port int
	fmt.Print("Enter the emulator port: ")
	fmt.Scan(&port)

	os.Setenv("PUBSUB_EMULATOR_HOST", fmt.Sprintf("0.0.0.0:%d", port))

	var project string
	fmt.Print("Enter the project name: ")
	fmt.Scan(&project)

	for {
		fmt.Println("\nPlease choose an option:")

		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"#", "Available Command"})
		t.AppendRows([]table.Row{
			{1, "Create Topic"},
			{2, "Create Subscription"},
			{3, "Exit"},
		})
		t.Render()

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var topic string
			fmt.Print("Enter the topic name: ")
			fmt.Scan(&topic)
			err := CreateTopic(project, topic, cl)
			if err != nil {
				cl.Errorf("Error creating the topic [%s] in project [%s] error: %s", topic, project, err)
			}
		case 2:
			var topic string
			fmt.Print("Enter the topic name: ")
			fmt.Scan(&topic)

			var subscription string
			fmt.Print("Enter the subscription name: ")
			fmt.Scan(&subscription)

			err := CreateSubscription(project, topic, subscription, cl)
			if err != nil {
				cl.Errorf("Error creating the subscription [%s] in topic [%s] in project [%s] error: %s", topic, project, err)
			}
		case 3:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice!")
		}
	}
}
