package gomarathon

import "strings"

// AppStatus is app's status
type AppStatus int

// AppStatus const
const (
	AppStatusNone AppStatus = iota
	AppStatusHealthy
	AppStatusUnHealthy
	AppStatusCreating
	AppStatusUpdating
)

func (s AppStatus) String() string {
	return strings.ToLower(_AppStatusValueToName[s])
}

// IsScaling represents whether an app is scaling
func (s AppStatus) IsScaling() bool {
	return s == AppStatusCreating || s == AppStatusUpdating
}
