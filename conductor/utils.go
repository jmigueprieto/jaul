// package with functions to work with Conductor
package conductor_utils

import (
	"os"

	"github.com/conductor-sdk/conductor-go/sdk/client"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/settings"
	"github.com/conductor-sdk/conductor-go/sdk/workflow/executor"
	log "github.com/sirupsen/logrus"
)

var (
	apiClient = client.NewAPIClient(
		authSettings(),
		httpSettings(),
	)
	workflowExecutor = executor.NewWorkflowExecutor(apiClient)
)

func RunWorkflow(name string, version int32, input map[string]interface{}) (workflowId string, err error) {
	startRequest := model.StartWorkflowRequest{
		Name:    name,
		Version: version,
		Input:   input,
	}

	return workflowExecutor.StartWorkflow(&startRequest)
}

func authSettings() *settings.AuthenticationSettings {
	key := os.Getenv("KEY")
	secret := os.Getenv("SECRET")
	if key != "" && secret != "" {
		return settings.NewAuthenticationSettings(
			key,
			secret,
		)
	}

	return nil
}

func httpSettings() *settings.HttpSettings {
	url := os.Getenv("CONDUCTOR_SERVER_URL")
	if url == "" {
		log.Error("Error: CONDUCTOR_SERVER_URL env variable is not set")
		os.Exit(1)
	}

	return settings.NewHttpSettings(url)
}
