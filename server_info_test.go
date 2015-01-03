package marathon

import (
	"testing"

	log "github.com/Sirupsen/logrus"
)

func TestGetInfo(t *testing.T) {
	if serverInfo, err := client.GetInfo(); err != nil {
		t.Error(err)
	} else {
		log.Debugf("ServerInfo: %#v\n", serverInfo)
	}
}
