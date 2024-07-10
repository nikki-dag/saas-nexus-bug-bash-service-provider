package app

import (
	"context"
	"fmt"

	"github.com/nexus-rpc/sdk-go/nexus"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/temporalnexus"
	"go.temporal.io/sdk/workflow"

	"github.com/nikki-dag/saas-nexus-bug-bash-api/service"
)

var (
	EchoOperation = temporalnexus.NewSyncOperation(service.EchoOperationName, func(ctx context.Context, c client.Client, input service.EchoInput, o nexus.StartOperationOptions) (service.EchoOutput, error) {
		return service.EchoOutput(input), nil
	})
	HelloOperation = temporalnexus.NewWorkflowRunOperation(service.HelloOperationName, HelloHandlerWorkflow, func(ctx context.Context, input service.HelloInput, soo nexus.StartOperationOptions) (client.StartWorkflowOptions, error) {
		return client.StartWorkflowOptions{
			ID: soo.RequestID,
		}, nil
	})
)

func HelloHandlerWorkflow(_ workflow.Context, input service.HelloInput) (service.HelloOutput, error) {
	switch input.Language {
	case service.EN:
		return service.HelloOutput{Message: "Hello " + input.Name + " 👋"}, nil
	case service.FR:
		return service.HelloOutput{Message: "Bonjour " + input.Name + " 👋"}, nil
	case service.DE:
		return service.HelloOutput{Message: "Hallo " + input.Name + " 👋"}, nil
	case service.ES:
		return service.HelloOutput{Message: "¡Hola! " + input.Name + " 👋"}, nil
	case service.TR:
		return service.HelloOutput{Message: "Merhaba " + input.Name + " 👋"}, nil
	}
	return service.HelloOutput{}, fmt.Errorf("unsupported language %q", input.Language)
}
