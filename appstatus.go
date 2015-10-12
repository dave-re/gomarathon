package gomarathon

import "strings"

// AppStatus is app's status
type AppStatus int

// AppStatus const
const (
	AppStatusNone AppStatus = iota
	AppStatusHealthy
	AppStatusUnHealthy
	AppStatusScaling
	AppStatusRunning
)

func (s AppStatus) String() string {
	return strings.ToLower(_AppStatusValueToName[s])
}
