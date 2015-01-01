package marathon

// RequestOptions passed for query api
type RequestOptions struct {
	Method string
	Path   string
	Datas  interface{}
	Params *Parameters
}

// Parameters to build url query
type Parameters struct {
	Cmd         string
	Host        string
	Scale       bool
	CallbackURL string
	Embed       Embed
}

// Application marathon application see :
// https://mesosphere.github.io/marathon/docs/rest-api.html#apps
type Application struct {
	ID              string            `json:"id,omitempty"`
	Cmd             string            `json:"cmd,omitempty"`
	Args            []string          `json:"args,omitempty"`
	Container       *Container        `json:"container,omitempty"`
	CPUs            float64           `json:"cpus,omitempty"`
	Mem             float64           `json:"mem,omitempty"`
	Deployments     []*Deployment     `json:"deployments,omitempty"`
	Env             map[string]string `json:"env,omitempty"`
	Executor        string            `json:"executor,omitempty"`
	Constraints     [][]string        `json:"constraints,omitempty"`
	HealthChecks    []*HealthCheck    `json:"healthChecks,omitempty"`
	Instances       int               `json:"instances,omitemptys"`
	Ports           []int             `json:"ports,omitempty"`
	BackoffSeconds  int               `json:"backoffSeconds,omitempty"`
	BackoffFactor   float64           `json:"backoffFactor,omitempty"`
	TasksRunning    int               `json:"tasksRunning,omitempty"`
	TasksStaged     int               `json:"tasksStaged,omitempty"`
	Uris            []string          `json:"uris,omitempty"`
	Dependencies    []string          `json:"dependencies,omitempty"`
	UpgradeStrategy *UpgradeStrategy  `json:"upgradeStrategy,omitempty"`
	Version         string            `json:"version,omitempty"`
	Tasks           []*Task           `json:"tasks,omitempty"`
}

// Task is described here:
// https://mesosphere.github.io/marathon/docs/rest-api.html#tasks
type Task struct {
	AppID     string `json:"appId"`
	Host      string `json:"host"`
	ID        string `json:"id"`
	Ports     []int  `json:"ports"`
	StagedAt  string `json:"stagedAt"`
	StartedAt string `json:"startedAt"`
	Version   string `json:"version"`
}

// Container is docker parameters
type Container struct {
	Type         string         `json:"type"`
	Image        string         `json:"image,omitempty"`
	Network      string         `json:"network,omitempty"`
	PortMappings []*PortMapping `json:"portMappings,omitempty"`
	Volumes      []*Volume      `json:"volumes,omitempty"`
}

type PortMapping struct {
	ContainerPort int    `json:"containerPort,omitempty"`
	HostPort      int    `json:"hostPort,omitempty"`
	ServicePort   int    `json:"servicePort,omitempty"`
	Protocol      string `json:"protocol,omitempty"`
}

type Volume struct {
	ContainerPath string `json:"containerPath,omitempty"`
	HostPath      string `json:"hostPath,omitempty"`
	Mode          string `json:"mode,omitempty"`
}

type Deployment struct {
	ID string `json:"id"`
}

// HealthCheck is described here:
// https://mesosphere.github.io/marathon/docs/health-checks.html
type HealthCheck struct {
	Protocol               string   `json:"protocol,omitempty"`
	Path                   string   `json:"path,omitempty"`
	GracePeriodSeconds     int      `json:"gracePeriodSeconds,omitempty"`
	IntervalSeconds        int      `json:"intervalSeconds,omitempty"`
	PortIndex              int      `json:"portIndex,omitempty"`
	TimeoutSeconds         int      `json:"timeoutSeconds,omitempty"`
	MaxConsecutiveFailures int      `json:"maxConsecutiveFailures,omitempty"`
	Command                *Command `json:"command,omitempty"`
}

type Command struct {
	Value string `json:"value"`
}

type UpgradeStrategy struct {
	MinimumHealthCapacity float64 `json:"minimumHealthCapacity"`
}

type Embed int

const (
	None Embed = iota
	AppsTasks
	AppsFailures
)

func (e Embed) String() string {
	switch e {
	case None:
		return "none"
	case AppsTasks:
		return "apps.tasks"
	case AppsFailures:
		return "apps.failures"
	}
	return ""
}

type ResponseApps struct {
	Apps []*Application `json:"apps"`
}

type ResponseApp struct {
	App *Application `json:"app"`
}

type ResponseVersions struct {
	Versions []string `json:"versions"`
}

type ResponseAppInfo struct {
	DeploymentID string   `json:"deploymentId"`
	Versions     []string `json:"versions"`
}
