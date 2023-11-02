package main

import (
	"log"
	notifications "practice/main"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

//register and run workflow

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create the client", err)
	}
	defer c.Close()

	w := worker.New(c, "notification-queue-3", worker.Options{})
	w.RegisterWorkflow(notifications.ExecuteNotifications)
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}

}
