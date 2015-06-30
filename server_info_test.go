package gomarathon

import (
	"testing"

	"github.com/Sirupsen/logrus"
)

func TestGetInfo(t *testing.T) {
	if serverInfo, err := client.GetInfo(); err != nil {
		t.Error(err)
	} else {
		logrus.Debugf("ServerInfo: %#v\n", serverInfo)
	}
}
