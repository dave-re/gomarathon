package gomarathon

import (
	"testing"

	"github.com/Sirupsen/logrus"
)

func TestGetDeployments(t *testing.T) {
	createApp("deploytest")
	if deployments, err := client.GetDeployments(); err != nil {
		t.Error(err)
	} else {
		for _, deployment := range deployments {
			logrus.Debugf("Deployment: %#v\n", deployment)
		}
	}
	destroyApp("deploytest")
}

func TestCancelDeployment(t *testing.T) {
	createApp("deploytest")
	deployments, err := client.GetDeployments()
	if err != nil {
		t.Error(err)
	}

	if deploymentID, version, err := client.CancelDeployment(deployments[0].ID); err != nil {
		t.Error(err)
	} else {
		logrus.Debugf("DeploymentID: %s\n", deploymentID)
		logrus.Debugf("Version: %s\n", version)
	}
	destroyApp("deploytest")
}
