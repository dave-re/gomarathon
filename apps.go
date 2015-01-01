package marathon

import (
	"fmt"
	"net/http"
)

const (
	rootPath = "apps"
)

func (c *Client) CreateApp(app *Application) (resApp *Application, err error) {
	options := &RequestOptions{
		Path:   rootPath,
		Datas:  app,
		Method: "POST",
	}
	err = c.unmarshalJson(options, http.StatusCreated, resApp)
	return
}

func (c *Client) GetApps() (*ResponseApps, error) {
	return c.GetAppsWithParams(nil)
}

func (c *Client) GetAppsWithParams(params *Parameters) (resApps *ResponseApps, err error) {
	options := &RequestOptions{
		Path:   rootPath,
		Method: "GET",
		Params: params,
	}
	err = c.unmarshalJson(options, http.StatusOK, resApps)
	return
}

func (c *Client) GetApp(appID string) (resApp *ResponseApp, err error) {
	options := &RequestOptions{
		Path:   fmt.Sprintf("%s/%s", rootPath, appID),
		Method: "GET",
	}
	err = c.unmarshalJson(options, http.StatusOK, resApp)
	return
}

func (c *Client) GetAppVersions(appID string) (resVersions *ResponseVersions, err error) {
	options := &RequestOptions{
		Path:   fmt.Sprintf("%s/%s/versions", rootPath, appID),
		Method: "GET",
	}
	err = c.unmarshalJson(options, http.StatusOK, resVersions)
	return
}

func (c *Client) GetAppByVersion(appID, version string) (resApp *Application, err error) {
	options := &RequestOptions{
		Path:   fmt.Sprintf("%s/%s/versions/%s", rootPath, appID, version),
		Method: "GET",
	}
	err = c.unmarshalJson(options, http.StatusOK, resApp)
	return
}

func (c *Client) UpdateApp(appID string, app *Application) (resAppInfo *ResponseAppInfo, err error) {
	return c.UpdateAppWithParams(appID, app, nil)
}

func (c *Client) UpdateAppWithParams(appID string, app *Application, params *Parameters) (resAppInfo *ResponseAppInfo, err error) {
	options := &RequestOptions{
		Path:   fmt.Sprintf("%s/%s", rootPath, appID),
		Method: "PUT",
		Params: params,
	}
	err = c.unmarshalJson(options, http.StatusOK, resAppInfo)
	return
}

func (c *Client) DestroyApp(appID string) error {
	options := &RequestOptions{
		Path:   fmt.Sprintf("%s/%s", rootPath, appID),
		Method: "DELETE",
	}
	return c.requestAndCheckSucc(options, http.StatusNoContent)
}
