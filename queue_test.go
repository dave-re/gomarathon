package gomarathon

import (
	"testing"

	log "github.com/Sirupsen/logrus"
)

func TestGetTaskQueue(t *testing.T) {
	createApp("taskqueuetest1")
	if queue, err := client.GetTaskQueue(); err != nil {
		t.Error(err)
	} else {
		for _, taskQueue := range queue {
			log.Debugf("TaskQueue: %#v\n", taskQueue)
		}
	}
	destroyApp("taskqueuetest1")
}
