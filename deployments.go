package gomarathon

import (
	"fmt"
	"net/http"
)

// GetDeployments gets running deployments
// http://goo.gl/SNyzga
func (c *Client) GetDeployments() (deployments []*Deployment, err error) {
	options := &RequestOptions{
		Path:   "deployments",
		Method: "GET",
	}
	deployments = []*Deployment{}
	err = c.unmarshalJSON(options, []int{http.StatusOK}, &deployments)
	return
}

// CancelDeployment cancel the deployment with deploymentID
// http://goo.gl/3px8pX
func (c *Client) CancelDeployment(deploymentID string) (respDeploymentID, version string, err error) {
	return c.CancelDeploymentWithParams(deploymentID, false)
}

// CancelDeploymentWithParams cancel the deployment with deploymentID and parameters
// http://goo.gl/3px8pX
func (c *Client) CancelDeploymentWithParams(deploymentID string, force bool) (respDeploymentID, version string, err error) {
	options := &RequestOptions{
		Path:   fmt.Sprintf("deployments/%s", deploymentID),
		Method: "DELETE",
		Params: &Parameters{Force: force},
	}
	resp := &response{}
	err = c.unmarshalJSON(options, []int{http.StatusOK, http.StatusAccepted}, resp)
	respDeploymentID = resp.DeploymentID
	version = resp.Version
	return
}
