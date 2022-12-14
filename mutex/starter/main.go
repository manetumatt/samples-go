package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/pborman/uuid"
	"go.temporal.io/sdk/client"

	"github.com/temporalio/samples-go/mutex"
)

func main() {
	// The client is a heavyweight object that should be created once per process.
	c, err := client.Dial(client.Options{
		HostPort: client.DefaultHostPort,
	})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	// This workflow ID can be user business logic identifier as well.
	resourceID := "MTRID"

	ctx := context.Background()

	for x := 1; x <= 5; x++ {

		id := fmt.Sprintf("SWF%d_%v", x, uuid.New())
		workflow1Options := client.StartWorkflowOptions{
			ID:        id,
			TaskQueue: "mutex",
		}

		we, err := c.ExecuteWorkflow(ctx, workflow1Options, mutex.SampleWorkflowWithMutex, resourceID)
		if err != nil {
			log.Fatalln("Unable to execute workflow", err)
		} else {
			log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())
		}

		time.Sleep(time.Millisecond * 50)
	}

	log.Println("Sleep 10 seconds before round 2")
	time.Sleep(10 * time.Second)
	resourceID = "MTRID2"

	for x := 1; x <= 5; x++ {

		id := fmt.Sprintf("SWF%d_%v", x, uuid.New())
		workflow1Options := client.StartWorkflowOptions{
			ID:        id,
			TaskQueue: "mutex",
		}

		we, err := c.ExecuteWorkflow(ctx, workflow1Options, mutex.SampleWorkflowWithMutex, resourceID)
		if err != nil {
			log.Fatalln("Unable to execute workflow", err)
		} else {
			log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())
		}

		time.Sleep(time.Millisecond * 50)
	}

}
