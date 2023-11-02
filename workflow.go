package notifications

import (
	"log"

	"go.temporal.io/sdk/workflow"
)

func ExecuteNotifications(ctx workflow.Context) (int, error) {

	executionChannel := workflow.GetSignalChannel(ctx, "EXECUTE_NOTIFICATIONS")
	var signal int

	flag := 0
	for {

		selector := workflow.NewSelector(ctx)
		selector.AddReceive(executionChannel, func(c workflow.ReceiveChannel, _ bool) {
			c.Receive(ctx, &signal)

			workflow.Go(ctx, func(ctx workflow.Context) {
				err := workflow.ExecuteActivity(ctx, SendNotification).Get(ctx, nil)
				if err != nil {
					log.Println("---------------Activity Error", err)
					return
				}
				log.Println("Executed ====> ", signal)
			})

		})

		selector.Select(ctx)

		if flag == 1 {
			break
		}

	}

	return signal, nil
}
