package marathon

import (
	"testing"

	log "github.com/Sirupsen/logrus"
)

func TestGetDeployments(t *testing.T) {
	createApp()
	if deployments, err := client.GetDeployments(); err != nil {
		t.Error(err)
	} else {
		for _, deployment := range deployments {
			log.Debugf("Deployment: %#v\n", deployment)
		}
	}
	destroyApp()
}

func TestCancelDeployment(t *testing.T) {
	createApp()
	deployments, err := client.GetDeployments()
	if err != nil {
		t.Error(err)
	}

	if deploymentID, version, err := client.CancelDeployment(deployments[0].ID, false); err != nil {
		t.Error(err)
	} else {
		log.Debugf("DeploymentID: %s\n", deploymentID)
		log.Debugf("Version: %s\n", version)
	}
	destroyApp()
}
