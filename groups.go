package marathon

import (
	"fmt"
	"net/http"
)

func (c *Client) GetGroups() (groups []*Group, err error) {
	options := &RequestOptions{
		Path:   "groups",
		Method: "GET",
	}
	group := &Group{}
	err = c.unmarshalJSON(options, []int{http.StatusOK}, group)
	groups = group.Groups
	return
}

func (c *Client) GetGroup(groupID string) (group *Group, err error) {
	options := &RequestOptions{
		Path:   fmt.Sprintf("groups/%s", groupID),
		Method: "GET",
	}
	err = c.unmarshalJSON(options, []int{http.StatusOK}, group)
	return
}

func (c *Client) CreateGroup(group *Group) error {
	options := &RequestOptions{
		Path:   "groups",
		Datas:  group,
		Method: "POST",
	}
	return c.requestAndCheckSucc(options, []int{http.StatusCreated})
}

func (c *Client) UpdateGroup(groupID string, group *Group) (deploymentID, version string, err error) {
	options := &RequestOptions{
		Path:   fmt.Sprintf("groups/%s", groupID),
		Datas:  group,
		Method: "PUT",
	}
	resp := &response{}
	err = c.unmarshalJSON(options, []int{http.StatusOK}, resp)
	deploymentID = resp.DeploymentID
	version = resp.Version
	return
}

func (c *Client) RollbackGroup(groupID, version string, force bool) (deploymentID, respVersion string, err error) {
	options := &RequestOptions{
		Path:   fmt.Sprintf("groups/%s/versions/%s", groupID, version),
		Method: "PUT",
		Params: &Parameters{Force: force},
	}
	resp := &response{}
	err = c.unmarshalJSON(options, []int{http.StatusOK}, resp)
	deploymentID = resp.DeploymentID
	respVersion = resp.Version
	return
}

func (c *Client) DestroyGroup(groupID string) (version string, err error) {
	options := &RequestOptions{
		Path:   fmt.Sprintf("groups/%s", groupID),
		Method: "DELETE",
	}
	resp := &response{}
	err = c.unmarshalJSON(options, []int{http.StatusOK}, resp)
	version = resp.Version
	return
}
