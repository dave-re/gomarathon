package marathon

import (
	"fmt"
	"net/http"
)

func (c *Client) CreateGroup(group *Group) (deploymentID, version string, err error) {
	options := &RequestOptions{
		Path:   "groups",
		Datas:  group,
		Method: "POST",
	}
	resp := &response{}
	err = c.unmarshalJSON(options, []int{http.StatusCreated}, resp)
	deploymentID = resp.DeploymentID
	version = resp.Version
	return
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
	group = &Group{}
	err = c.unmarshalJSON(options, []int{http.StatusOK}, group)
	return
}

func (c *Client) DestroyGroup(groupID string) (deploymentID, version string, err error) {
	options := &RequestOptions{
		Path:   fmt.Sprintf("groups/%s", groupID),
		Method: "DELETE",
	}
	resp := &response{}
	err = c.unmarshalJSON(options, []int{http.StatusOK}, resp)
	deploymentID = resp.DeploymentID
	version = resp.Version
	return
}
