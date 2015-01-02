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
	Force       bool
}

// Application marathon application see :
// https://mesosphere.github.io/marathon/docs/rest-api.html#apps
type Application struct {
	ID              string            `json:"id,omitempty"`
	Cmd             string            `json:"cmd,omitempty"`
	Args            []string          `json:"args,omitempty"`
	User            string            `json:"user,omitempty"`
	Container       *Container        `json:"container,omitempty"`
	CPUs            float64           `json:"cpus,omitempty"`
	Mem             float64           `json:"mem,omitempty"`
	Disk            float64           `json:"disk,omitempty"`
	Deployments     []*Deployment     `json:"deployments,omitempty"`
	Env             map[string]string `json:"env,omitempty"`
	Executor        string            `json:"executor,omitempty"`
	Constraints     [][]string        `json:"constraints,omitempty"`
	HealthChecks    []*HealthCheck    `json:"healthChecks,omitempty"`
	Instances       int               `json:"instances,omitemptys"`
	Ports           []int             `json:"ports,omitempty"`
	RequirePorts    bool              `json:"requirePorts,omitempty"`
	BackoffSeconds  int               `json:"backoffSeconds,omitempty"`
	BackoffFactor   float64           `json:"backoffFactor,omitempty"`
	TasksRunning    int               `json:"tasksRunning,omitempty"`
	TasksStaged     int               `json:"tasksStaged,omitempty"`
	Uris            []string          `json:"uris,omitempty"`
	StoreUris       []string          `json:"storeUris,omitempty"`
	Dependencies    []string          `json:"dependencies,omitempty"`
	UpgradeStrategy *UpgradeStrategy  `json:"upgradeStrategy,omitempty"`
	Version         string            `json:"version,omitempty"`
	Tasks           []*Task           `json:"tasks,omitempty"`
}

type Group struct {
	ID           string         `json:"id,omitempty"`
	Dependencies []string       `json:"dependencies,omitempty"`
	Apps         []*Application `json:"apps,omitempty"`
	Groups       []*Group       `json:"groups,omitempty"`
	Version      string         `json:"version,omitempty"`
	ScaleBy      float64        `json:"scaleBy,omitempty"`
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
	Type   string  `json:"type"`
	Docker *Docker `json:"docker,omitempty"`
}

type Docker struct {
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
	ID             string    `json:"id,omitempty"`
	AffectedApps   []string  `json:"affectedApps,omitempty"`
	Steps          []*Step   `json:"steps,omitempty"`
	CurrentActions []*Action `json:"currentActions,omitempty"`
	Version        string    `json:"version,omitempty"`
	CurrentStep    int       `json:"currentStep,omitempty"`
	TotalSteps     int       `json:"totalSteps,omitempty"`
}

type Step []*Action

type Action struct {
	Action string `json:"action"`
	App    string `json:"app"`
}

type TaskQueue struct {
	App   string `json:"app"`
	Delay *Delay `json:"delay"`
}

type Delay struct {
	Overdue bool `json:"overdue"`
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
	MinimumHealthCapacity float64 `json:"minimumHealthCapacity,omitempty"`
}

type ServerInfo struct {
	FrameworkID     string           `json:"frameworkId,omitempty"`
	Leader          string           `json:"leader,omitempty"`
	HTTPConfig      *HTTPConfig      `json:"http_config,omitempty"`
	EventSubscriber *EventSubscriber `json:"event_subscriber,omitempty"`
	MarathonConfig  *MarathonConfig  `json:"marathon_config,omitempty"`
	Name            string           `json:"name,omitempty"`
	Version         string           `json:"version,omitempty"`
	ZookeeperConfig *ZookeeperConfig `json:"zookeeper_config,omitempty"`
}

type HTTPConfig struct {
	AssetsPath string `json:"assets_path,omitempty"`
	HTTPPort   int    `json:"http_port,omitempty"`
	HTTPSPort  int    `json:"https_port,omitempty"`
}

type EventSubscriber struct {
	Type          string   `json:"type,omitempty"`
	HTTPEndpoints []string `json:"http_endpoints,omitempty"`
}

type MarathonConfig struct {
	Checkpoint                 bool   `json:"checkpoint,omitempty"`
	Executor                   string `json:"executor,omitempty"`
	FailoverTimeout            int    `json:"failover_timeout,omitempty"`
	HA                         bool   `json:"ha,omitempty"`
	Hostname                   string `json:"hostname,omitempty"`
	LocalPortMax               int    `json:"local_port_max,omitempty"`
	LocalPortMin               int    `json:"local_port_min,omitempty"`
	Master                     string `json:"master,omitempty"`
	MesosRole                  string `json:"mesos_role,omitempty"`
	MesosUser                  string `json:"mesos_user,omitempty"`
	ReconciliationInitialDelay int    `json:"reconciliation_initial_delay,omitempty"`
	ReconciliationInterval     int    `json:"reconciliation_interval,omitempty"`
	TaskLaunchTimeout          int    `json:"task_launch_timeout,omitempty"`
}

type ZookeeperConfig struct {
	ZK              string           `json:"zk,omitempty"`
	ZKFutureTimeout *ZKFutureTimeout `json:"zk_future_timeout,omitempty"`
	ZKHosts         string           `json:"zk_hosts,omitempty"`
	ZKPath          string           `json:"zk_path,omitempty"`
	ZKState         string           `json:"zk_state,omitempty"`
	ZKTimeout       int              `json:"zk_timeout,omitempty"`
}

type ZKFutureTimeout struct {
	Duration int `json:"duration,omitempty"`
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

type response struct {
	Apps         []*Application `json:"apps,omitempty"`
	App          *Application   `json:"app,omitempty"`
	Versions     []string       `json:"versions,omitempty"`
	DeploymentID string         `json:"deploymentId,omitempty"`
	Version      string         `json:"version,omitempty"`
	Tasks        []*Task        `json:"tasks,omitempty"`
	Task         *Task          `json:"task,omitempty"`
	CallbackUrls []string       `json:"callbackUrls,omitempty"`
	Queue        []*TaskQueue   `json:"queue,omitempty"`
}
