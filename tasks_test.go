package gomarathon

import (
	"testing"
	"time"

	"github.com/Sirupsen/logrus"
)

func TestTasks(t *testing.T) {
	createApp("tasktest")
	time.Sleep(5 * time.Second)
	if tasks, err := client.GetTasks(); err != nil {
		t.Error(err)
	} else {
		for _, task := range tasks {
			logrus.Debugf("Task: %#v\n", task)
		}
	}
}
