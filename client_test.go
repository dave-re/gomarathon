package gomarathon

import "testing"

var (
	client, _ = NewClient("http://localdocker:8080", nil)
)

func createApp(id string) {
	client.CreateApp(&Application{
		ID:        id,
		Instances: 1,
		CPUs:      0.1,
		Mem:       128.0,
		Container: &Container{
			Type: "DOCKER",
			Docker: &Docker{
				Image: "nginx:latest",
			},
		},
	})
}

func destroyApp(id string) {
	client.DestroyApp(id)
}

func TestNewClient(t *testing.T) {
	_, err := NewClient("http://localdocker:8080", nil)
	if err != nil {
		t.Error(err)
	}
}
