package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	notifications "practice/main"
	notificationtypes "practice/main/types"
	"strconv"
	"strings"
	"time"

	"go.temporal.io/sdk/client"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	//execute workflow

	startTime := time.Now()
	users, _ := strconv.Atoi(os.Args[1])
	metaInfo := map[int]notificationtypes.WorkflowInfo{}

	for i := 1; i <= users; i++ {
		options := client.StartWorkflowOptions{
			ID:        "notification-" + fmt.Sprint(i),
			TaskQueue: "notification-queue-3",
		}
		we, err := c.ExecuteWorkflow(context.Background(), options, notifications.ExecuteNotifications)
		log.Println("workflow Run ID : ", we.GetRunID())
		if err != nil {
			log.Fatalln("unable to execute workflow", err)
		}
		metaInfo[i] = notificationtypes.WorkflowInfo{WorkflowID: options.ID, RunID: we.GetRunID()}
	}
	reader := bufio.NewReader(os.Stdin)

	for i := 1; i <= users; i++ {
		log.Println(metaInfo[i])
	}

	for {
		input, _ := reader.ReadString('\n')
		arr := strings.Split(input, " ")
		log.Println(arr)
		choice, _ := strconv.Atoi(arr[0])
		x, _ := strconv.Atoi(arr[1])
		y, _ := strconv.Atoi(strings.Trim(arr[2], "\n"))

		switch choice {
		case 1:
			for i := x; i <= y; i++ {
				log.Println("**********Sending signal : ", i)
				c.SignalWorkflow(context.Background(), metaInfo[i].WorkflowID, metaInfo[i].RunID, "EXECUTE_NOTIFICATIONS", i)
			}
		}

	}

	log.Println("************ Message is :", time.Since(startTime))
}
