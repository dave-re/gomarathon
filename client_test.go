package marathon

import "testing"

func TestNewClient(t *testing.T) {
	_, err := NewClient("http://localdocker:8080", nil)
	if err != nil {
		t.Error(err)
	}
}
