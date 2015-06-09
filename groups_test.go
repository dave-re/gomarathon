package gomarathon

import (
	"testing"

	log "github.com/Sirupsen/logrus"
)

var (
	groupID = "group1"
)

func TestCreateGroup(t *testing.T) {
	if deploymentID, version, err := client.CreateGroup(&Group{
		ID: groupID,
		Apps: []*Application{
			&Application{
				ID:        appID,
				Instances: 1,
				CPUs:      0.1,
				Mem:       128.0,
				Container: &Container{
					Type: "DOCKER",
					Docker: &Docker{
						Image: dockerAppImage,
					},
				},
			},
		},
	}); err != nil {
		t.Error(err)
	} else {
		log.Debugf("DeploymentID: %s\n", deploymentID)
		log.Debugf("Version: %s\n", version)
	}
}

func TestUpdateGroup(t *testing.T) {
	if deploymentID, version, err := client.UpdateGroup(groupID, &Group{
		Apps: []*Application{
			&Application{
				ID:        appID,
				Instances: 2,
				CPUs:      0.2,
				Mem:       256.0,
				Container: &Container{
					Type: "DOCKER",
					Docker: &Docker{
						Image: dockerAppImage,
					},
				},
			},
		},
	}); err != nil {
		t.Error(err)
	} else {
		log.Debugf("DeploymentID: %s\n", deploymentID)
		log.Debugf("Version: %s\n", version)
	}
}

func TestGetGroups(t *testing.T) {
	if groups, err := client.GetGroups(); err != nil {
		t.Error(err)
	} else {
		for _, group := range groups {
			log.Debugf("Group: %#v\n", group)
		}
	}
}

func TestGetGroup(t *testing.T) {
	if group, err := client.GetGroup(groupID); err != nil {
		t.Error(err)
	} else {
		log.Debugf("Group: %#v\n", group)
	}
}

func TestDestroyGroup(t *testing.T) {
	if deploymentID, version, err := client.DestroyGroup(groupID); err != nil {
		t.Error(err)
	} else {
		log.Debugf("DeploymentID: %s\n", deploymentID)
		log.Debugf("Version: %s\n", version)
	}
}
