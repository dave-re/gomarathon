package marathon

import (
	"testing"

	log "github.com/Sirupsen/logrus"
)

func TestRegisterCallbackURL(t *testing.T) {
	if response, err := client.RegisterCallbackURL("http://testgo.com"); err != nil {
		t.Error(err)
	} else {
		log.Debugf("Response: %#v\n", response)
	}
}

func TestGetCallbackURLs(t *testing.T) {
	if callbackURLs, err := client.GetCallbackURLs(); err != nil {
		t.Error(err)
	} else {
		for _, callbackURL := range callbackURLs {
			log.Debugf("callbackURL: %s\n", callbackURL)
		}
	}
}

func TestUnregisterCallbackURL(t *testing.T) {
	if response, err := client.UnregisterCallbackURL("http://testgo.com"); err != nil {
		t.Error(err)
	} else {
		log.Debugf("Response: %#v\n", response)
	}
}
