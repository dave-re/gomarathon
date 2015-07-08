package gomarathon

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
	Status      Status
}

// Application is marathon application
// http://goo.gl/MNP22N
type Application struct {
	ID                    string            `json:"id,omitempty"`
	Cmd                   string            `json:"cmd,omitempty"`
	Args                  []string          `json:"args,omitempty"`
	User                  string            `json:"user,omitempty"`
	Container             *Container        `json:"container,omitempty"`
	Env                   map[string]string `json:"env,omitempty"`
	Instances             int               `json:"instances,omitempty"`
	CPUs                  float64           `json:"cpus,omitempty"`
	Mem                   float64           `json:"mem,omitempty"`
	Disk                  float64           `json:"disk,omitempty"`
	Executor              string            `json:"executor,omitempty"`
	Constraints           [][]string        `json:"constraints,omitempty"`
	Uris                  []string          `json:"uris,omitempty"`
	StoreUrls             []string          `json:"storeUrls,omitempty"`
	Ports                 []int             `json:"ports,omitempty"`
	RequirePorts          bool              `json:"requirePorts,omitempty"`
	BackoffSeconds        int               `json:"backoffSeconds,omitempty"`
	BackoffFactor         float64           `json:"backoffFactor,omitempty"`
	MaxLaunchDelaySeconds int               `json:"maxLaunchDelaySeconds,omitempty"`
	TasksRunning          int               `json:"tasksRunning,omitempty"`
	TasksHealthy          int               `json:"tasksHealthy,omitempty"`
	TasksUnhealthy        int               `json:"tasksUnhealthy,omitempty"`
	TasksStaged           int               `json:"tasksStaged,omitempty"`
	HealthChecks          []*HealthCheck    `json:"healthChecks,omitempty"`
	Dependencies          []string          `json:"dependencies,omitempty"`
	UpgradeStrategy       *UpgradeStrategy  `json:"upgradeStrategy,omitempty"`
	Labels                map[string]string `json:"labels,omitempty"`
	Deployments           []*Deployment     `json:"deployments,omitempty"`
	Version               string            `json:"version,omitempty"`
	Tasks                 []*Task           `json:"tasks,omitempty"`
	LastTaskFailure       *TaskFailure      `json:"lastTaskFailure,omitempty"`
}

// GetStatus returns current status of an app
func (app *Application) GetStatus() AppStatus {
	if app.Instances == 0 {
		return AppStatusNone
	}
	switch {
	case (app.Instances == app.TasksHealthy) && (app.TasksHealthy == app.TasksRunning):
		return AppStatusHealthy
	case (app.TasksUnhealthy > 0):
		return AppStatusUnHealthy
	case (app.TasksHealthy == 0):
		return AppStatusCreating
	default:
		return AppStatusUpdating
	}
}

// Group is marathon group
// http://goo.gl/6n6FSI
type Group struct {
	ID           string         `json:"id,omitempty"`
	Dependencies []string       `json:"dependencies,omitempty"`
	Apps         []*Application `json:"apps,omitempty"`
	Groups       []*Group       `json:"groups,omitempty"`
	Version      string         `json:"version,omitempty"`
	ScaleBy      float64        `json:"scaleBy,omitempty"`
}

// Task is described here:
// http://goo.gl/gw6dTA
type Task struct {
	AppID        string `json:"appId,omitempty"`
	ID           string `json:"id,omitempty"`
	Host         string `json:"host,omitempty"`
	Ports        []int  `json:"ports,omitempty"`
	StagedAt     string `json:"stagedAt,omitempty"`
	StartedAt    string `json:"startedAt,omitempty"`
	Version      string `json:"version,omitempty"`
	ServicePorts []int  `json:"servicePorts,omitempty"`
}

// Container is container parameters
type Container struct {
	Type    string    `json:"type,omitempty"`
	Docker  *Docker   `json:"docker,omitempty"`
	Volumes []*Volume `json:"volumes,omitempty"`
}

// Docker options
type Docker struct {
	Image          string         `json:"image,omitempty"`
	Network        string         `json:"network,omitempty"`
	PortMappings   []*PortMapping `json:"portMappings,omitempty"`
	Privileged     bool           `json:"privileged,omitempty"`
	Parameters     []*Parameter   `json:"parameters,omitempty"`
	ForcePullImage bool           `json:"forcePullImage,omitempty"`
}

// Parameter is a generic (key, value) pair used in various places for parameters.
type Parameter struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

// PortMapping is port mapping for docker container
type PortMapping struct {
	ContainerPort int    `json:"containerPort,omitempty"`
	HostPort      int    `json:"hostPort,omitempty"`
	ServicePort   int    `json:"servicePort,omitempty"`
	Protocol      string `json:"protocol,omitempty"`
}

// Volume is used for mounting a host directory as a container volume
type Volume struct {
	ContainerPath string `json:"containerPath,omitempty"`
	HostPath      string `json:"hostPath,omitempty"`
	Mode          string `json:"mode,omitempty"`
}

// Deployment is marathon deployment
// http://goo.gl/iJludX
type Deployment struct {
	ID             string    `json:"id,omitempty"`
	AffectedApps   []string  `json:"affectedApps,omitempty"`
	Steps          []*Step   `json:"steps,omitempty"`
	CurrentActions []*Action `json:"currentActions,omitempty"`
	Version        string    `json:"version,omitempty"`
	CurrentStep    int       `json:"currentStep,omitempty"`
	TotalSteps     int       `json:"totalSteps,omitempty"`
}

// Step is step for deployment
type Step []*Action

// Action is Action for deployment
type Action struct {
	Action string `json:"action,omitempty"`
	App    string `json:"app,omitempty"`
}

// TaskQueue is Action for queue
type TaskQueue struct {
	App   *Application `json:"app,omitempty"`
	Delay *Delay       `json:"delay,omitempty"`
}

// Delay is delay for task queue
type Delay struct {
	Overdue bool `json:"overdue,omitempty"`
}

// HealthCheck is described here:
// http://goo.gl/0GVD6o
type HealthCheck struct {
	Protocol               string   `json:"protocol,omitempty"`
	PortIndex              int      `json:"portIndex,omitempty"`
	GracePeriodSeconds     int      `json:"gracePeriodSeconds,omitempty"`
	IntervalSeconds        int      `json:"intervalSeconds,omitempty"`
	TimeoutSeconds         int      `json:"timeoutSeconds,omitempty"`
	Path                   string   `json:"path,omitempty"`
	MaxConsecutiveFailures int      `json:"maxConsecutiveFailures,omitempty"`
	Command                *Command `json:"command,omitempty"`
	IgnoreHTTP1xx          bool     `json:"ignoreHttp1xx,omitempty"`
}

// Command is command for health check
type Command struct {
	Value string `json:"value,omitempty"`
}

// UpgradeStrategy has a minimumHealthCapacity which defines the minimum number of healty nodes
type UpgradeStrategy struct {
	MinimumHealthCapacity float64 `json:"minimumHealthCapacity,omitempty"`
}

// ServerInfo is info about the Marathon Instance
// http://goo.gl/RSPWrw
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

// HTTPConfig is http config for server info
type HTTPConfig struct {
	AssetsPath string `json:"assets_path,omitempty"`
	HTTPPort   int    `json:"http_port,omitempty"`
	HTTPSPort  int    `json:"https_port,omitempty"`
}

// EventSubscriber is described here:
// http://goo.gl/02SuRT
type EventSubscriber struct {
	Type          string   `json:"type,omitempty"`
	HTTPEndpoints []string `json:"http_endpoints,omitempty"`
}

// MarathonEvent is base event of the marathon
// http://goo.gl/vxuutX
type MarathonEvent struct {
	EventType string `json:"eventType,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
}

// StatusUpdateEvent is status update event of the mesos
// http://goo.gl/vxuutX
type StatusUpdateEvent struct {
	MarathonEvent
	SlaveID    string `json:"slaveId,omitempty"`
	TaskID     string `json:"taskId,omitempty"`
	TaskStatus string `json:"taskStatus,omitempty"`
	AppID      string `json:"appId,omitempty"`
	Host       string `json:"host,omitempty"`
	Ports      []int  `json:"ports,omitempty"`
	Version    string `json:"version,omitempty"`
}

// FrameworkMessageEvent is framework message event of the mesos
// http://goo.gl/vxuutX
type FrameworkMessageEvent struct {
	MarathonEvent
	SlaveID    string `json:"slaveId,omitempty"`
	ExecutorID string `json:"executorId,omitempty"`
	Message    string `json:"message,omitempty"`
}

// SubscriptionEvent is subscription event of the marathon
// http://goo.gl/vxuutX
type SubscriptionEvent struct {
	MarathonEvent
	ClientIP    string `json:"clientIp,omitempty"`
	CallbackURL string `json:"callbackUrl,omitempty"`
}

// HealthCheckEvent is health check event of the marathon
// http://goo.gl/vxuutX
type HealthCheckEvent struct {
	MarathonEvent
	AppID       string       `json:"appId,omitempty"`
	TaskID      string       `json:"taskId,omitempty"`
	Version     string       `json:"version,omitempty"`
	Alive       bool         `json:"alive,omitempty"`
	HealthCheck *HealthCheck `json:"healthCheck,omitempty"`
}

// DeploymentsEvent is deployments event of the marathon
// http://goo.gl/vxuutX
type DeploymentsEvent struct {
	MarathonEvent
	ID          string          `json:"id,omitempty"`
	GroupID     string          `json:"groupId,omitempty"`
	Version     string          `json:"version,omitempty"`
	Reason      string          `json:"reason,omitempty"`
	Plan        *DeploymentPlan `json:"plan,omitempty"`
	CurrentStep *Action         `json:"currentStep,omitempty"`
}

type DeploymentPlan struct {
	ID       string    `json:"id,omitempty"`
	Original *Group    `json:"original,omitempty"`
	Target   *Group    `json:"target,omitempty"`
	Steps    []*Action `json:"steps,omitempty"`
	Version  string    `json:"version,omitempty"`
}

// MarathonConfig is config about the marathon
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

// ZookeeperConfig is config about the zookeeper
type ZookeeperConfig struct {
	ZK              string           `json:"zk,omitempty"`
	ZKFutureTimeout *ZKFutureTimeout `json:"zk_future_timeout,omitempty"`
	ZKHosts         string           `json:"zk_hosts,omitempty"`
	ZKPath          string           `json:"zk_path,omitempty"`
	ZKState         string           `json:"zk_state,omitempty"`
	ZKTimeout       int              `json:"zk_timeout,omitempty"`
}

// ZKFutureTimeout is future timeout for zookeeper
type ZKFutureTimeout struct {
	Duration int `json:"duration,omitempty"`
}

// TaskFailure is task that has failed
type TaskFailure struct {
	AppID     string `json:"appId,omitempty"`
	Host      string `json:"host,omitempty"`
	Message   string `json:"message,omitempty"`
	State     string `json:"state,omitempty"`
	TaskID    string `json:"taskId,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
	Version   string `json:"version,omitempty"`
}

// Embed is embed parameter
type Embed int

// Embed const
const (
	NoneEmbed Embed = iota
	AppsTasks
	AppsFailures
)

func (e Embed) String() string {
	switch e {
	case NoneEmbed:
		return "none"
	case AppsTasks:
		return "apps.tasks"
	case AppsFailures:
		return "apps.failures"
	}
	return ""
}

// Status is status parameter
type Status int

// Status const
const (
	StatusNone Status = iota
	StatusRunning
	StatusStaging
)

func (s Status) String() string {
	switch s {
	case StatusNone:
		return "none"
	case StatusRunning:
		return "running"
	case StatusStaging:
		return "staging"
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
