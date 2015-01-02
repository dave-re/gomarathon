package marathon

import (
	"fmt"
	"net/http"
)

type GetAppsParams struct {
	Cmd   string
	Embed Embed
}

type KillTasksParams struct {
	Host  string
	Scale bool
}

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

func (c *Client) GetApps() ([]*Application, error) {
	return c.GetAppsWithParams(nil)
}

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

func (c *Client) GetAppByVersion(appID, version string) (app *Application, err error) {
	options := &RequestOptions{
		Path:   fmt.Sprintf("apps/%s/versions/%s", appID, version),
		Method: "GET",
	}
	app = &Application{}
	err = c.unmarshalJSON(options, []int{http.StatusOK}, app)
	return
}

func (c *Client) UpdateApp(appID string, app *Application) (deploymentID, version string, err error) {
	return c.UpdateAppWithParams(appID, app, false)
}

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

func (c *Client) DestroyApp(appID string) error {
	options := &RequestOptions{
		Path:   fmt.Sprintf("apps/%s", appID),
		Method: "DELETE",
	}
	return c.requestAndCheckSucc(options, []int{http.StatusOK})
}

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

func (c *Client) KillTasks(appID string) (tasks []*Task, err error) {
	return c.KillTasksWithParams(appID, nil)
}

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

func (c *Client) KillTask(appID, taskID string) (task *Task, err error) {
	return c.KillTaskWithParams(appID, taskID, false)
}

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
