package marathon

import "testing"

var (
	client, _ = NewClient("http://localdocker:8080", nil)
)

func createApp() {
	client.CreateApp(&Application{
		ID:        appID,
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

func destroyApp() {
	client.DestroyApp(appID)
}

func TestNewClient(t *testing.T) {
	_, err := NewClient("http://localdocker:8080", nil)
	if err != nil {
		t.Error(err)
	}
}
