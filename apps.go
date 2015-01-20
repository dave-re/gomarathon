package gomarathon

import (
	"fmt"
	"net/http"
)

// GetAppsParams is parameters for GetAppsWithParams function
type GetAppsParams struct {
	Cmd   string
	Embed Embed
}

// KillTasksParams is parameters for KillTasksWithParams function
type KillTasksParams struct {
	Host  string
	Scale bool
}

// CreateApp create and start a new application.
// http://goo.gl/fM0CLu
func (c *Client) CreateApp(app *Application) (resApp *Application, err error) {
	options := &RequestOptions{
		Path:   "apps",
		Datas:  app,
		Method: "POST",
	}
	resApp = &Application{}
	err = c.unmarshalJSON(options, []int{http.StatusCreated}, resApp)
	return
}

// GetApps gets all running applications
// http://goo.gl/Sh617O
func (c *Client) GetApps() ([]*Application, error) {
	return c.GetAppsWithParams(nil)
}

// GetAppsWithParams gets all running applications with parameters
// http://goo.gl/i2EO8b
func (c *Client) GetAppsWithParams(params *GetAppsParams) (apps []*Application, err error) {
	options := &RequestOptions{
		Path:   "apps",
		Method: "GET",
	}
	if params != nil {
		options.Params = &Parameters{
			Cmd:   params.Cmd,
			Embed: params.Embed,
		}
	}
	resp := &response{}
	err = c.unmarshalJSON(options, []int{http.StatusOK}, resp)
	apps = resp.Apps
	return
}

// GetApp gets the application with appID
// http://goo.gl/4pbxGV
func (c *Client) GetApp(appID string) (app *Application, err error) {
	options := &RequestOptions{
		Path:   fmt.Sprintf("apps/%s", appID),
		Method: "GET",
	}
	resp := &response{}
	err = c.unmarshalJSON(options, []int{http.StatusOK}, resp)
	app = resp.App
	return
}

// GetAppVersions gets the versions of the application with appID
// http://goo.gl/hdBz1v
func (c *Client) GetAppVersions(appID string) (versions []string, err error) {
	options := &RequestOptions{
		Path:   fmt.Sprintf("apps/%s/versions", appID),
		Method: "GET",
	}
	resp := &response{}
	err = c.unmarshalJSON(options, []int{http.StatusOK}, resp)
	versions = resp.Versions
	return
}

// GetAppByVersion gets the configuration of the application with appID at version
// http://goo.gl/yUtHNn
func (c *Client) GetAppByVersion(appID, version string) (app *Application, err error) {
	options := &RequestOptions{
		Path:   fmt.Sprintf("apps/%s/versions/%s", appID, version),
		Method: "GET",
	}
	app = &Application{}
	err = c.unmarshalJSON(options, []int{http.StatusOK}, app)
	return
}

// UpdateApp change parameters of a running application
// http://goo.gl/LVbV33
func (c *Client) UpdateApp(appID string, app *Application) (deploymentID, version string, err error) {
	return c.UpdateAppWithParams(appID, app, false)
}

// UpdateAppWithParams change parameters of a running application with parameters
// http://goo.gl/LVbV33
func (c *Client) UpdateAppWithParams(appID string, app *Application, force bool) (deploymentID, version string, err error) {
	options := &RequestOptions{
		Path:   fmt.Sprintf("apps/%s", appID),
		Datas:  app,
		Method: "PUT",
		Params: &Parameters{
			Force: force,
		},
	}
	resp := &response{}
	err = c.unmarshalJSON(options, []int{http.StatusOK}, resp)
	deploymentID = resp.DeploymentID
	version = resp.Version
	return
}

// RestartAllTasks initiates a rolling restart of all running tasks of the given app
// http://goo.gl/aL7ndP
func (c *Client) RestartAllTasks(appID string, force bool) (deploymentID, version string, err error) {
	return c.RestartAllTasksWithParams(appID, false)
}

// RestartAllTasksWithParams initiates a rolling restart of all running tasks of the given app with parameters
// http://goo.gl/aL7ndP
func (c *Client) RestartAllTasksWithParams(appID string, force bool) (deploymentID, version string, err error) {
	options := &RequestOptions{
		Path:   fmt.Sprintf("apps/%s/restart", appID),
		Method: "POST",
		Params: &Parameters{
			Force: force,
		},
	}
	resp := &response{}
	err = c.unmarshalJSON(options, []int{http.StatusOK}, resp)
	deploymentID = resp.DeploymentID
	version = resp.Version
	return
}

// DestroyApp destroy an application
// http://goo.gl/70pGRO
func (c *Client) DestroyApp(appID string) (deploymentID, version string, err error) {
	options := &RequestOptions{
		Path:   fmt.Sprintf("apps/%s", appID),
		Method: "DELETE",
	}
	resp := &response{}
	err = c.unmarshalJSON(options, []int{http.StatusOK}, resp)
	deploymentID = resp.DeploymentID
	version = resp.Version
	return
}

// GetAppTasks gets all running tasks for application with appID
// http://goo.gl/Q2CbyV
func (c *Client) GetAppTasks(appID string) (tasks []*Task, err error) {
	options := &RequestOptions{
		Path:   fmt.Sprintf("apps/%s/tasks", appID),
		Method: "GET",
	}
	resp := &response{}
	err = c.unmarshalJSON(options, []int{http.StatusOK}, resp)
	tasks = resp.Tasks
	return
}

// KillTasks kill tasks that belong to the application with appID
// http://goo.gl/Czr77g
func (c *Client) KillTasks(appID string) (tasks []*Task, err error) {
	return c.KillTasksWithParams(appID, nil)
}

// KillTasksWithParams kill tasks that belong to the application with appID and parameters
// http://goo.gl/Czr77g
func (c *Client) KillTasksWithParams(appID string, params *KillTasksParams) (tasks []*Task, err error) {
	options := &RequestOptions{
		Path:   fmt.Sprintf("apps/%s/tasks", appID),
		Method: "DELETE",
	}
	if params != nil {
		options.Params = &Parameters{
			Host:  params.Host,
			Scale: params.Scale,
		}
	}
	resp := &response{}
	err = c.unmarshalJSON(options, []int{http.StatusOK}, resp)
	tasks = resp.Tasks
	return
}

// KillTask kill the task with taskID that belongs to the application with appID
// http://goo.gl/SURgNZ
func (c *Client) KillTask(appID, taskID string) (task *Task, err error) {
	return c.KillTaskWithParams(appID, taskID, false)
}

// KillTaskWithParams kill the task with taskID that belongs to the application with appID and parameters
// http://goo.gl/SURgNZ
func (c *Client) KillTaskWithParams(appID, taskID string, scale bool) (task *Task, err error) {
	options := &RequestOptions{
		Path:   fmt.Sprintf("apps/%s/tasks/%s", appID, taskID),
		Method: "DELETE",
		Params: &Parameters{
			Scale: scale,
		},
	}
	resp := &response{}
	err = c.unmarshalJSON(options, []int{http.StatusOK}, resp)
	task = resp.Task
	return
}
