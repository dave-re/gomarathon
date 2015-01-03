package marathon

import (
	"fmt"
	"net/http"
)

func (c *Client) GetDeployments() (deployments []*Deployment, err error) {
	options := &RequestOptions{
		Path:   "deployments",
		Method: "GET",
	}
	deployments = []*Deployment{}
	err = c.unmarshalJSON(options, []int{http.StatusOK}, &deployments)
	return
}

func (c *Client) CancelDeployment(deploymentID string, force bool) (respDeploymentID, version string, err error) {
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
