package gomarathon

import (
	"os"
	"testing"

	log "github.com/Sirupsen/logrus"
)

const (
	dockerAppImage = "mdock.daumkakao.io/nginx:latest"
)

var (
	client *Client
)

func init() {
	SetLogLevel(log.DebugLevel)
	client, _ = NewClient(os.Getenv("MARATHON_URL"), nil)
}

func createApp(id string) {
	client.CreateApp(&Application{
		ID:        id,
		Instances: 1,
		CPUs:      0.1,
		Mem:       128.0,
		Container: &Container{
			Type: "DOCKER",
			Docker: &Docker{
				Image: dockerAppImage,
			},
		},
	})
}

func destroyApp(id string) {
	client.DestroyApp(id)
}

func TestNewClient(t *testing.T) {
	_, err := NewClient(os.Getenv("MARATHON_URL"), nil)
	if err != nil {
		t.Error(err)
	}
}
