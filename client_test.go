package marathon

import "testing"

var (
	client, _ = NewClient("http://localdocker:8080", nil)
)

func TestNewClient(t *testing.T) {
	_, err := NewClient("http://localdocker:8080", nil)
	if err != nil {
		t.Error(err)
	}
}
