package gomarathon

import (
	"fmt"
	"net/http"
)

// CreateGroup create and start a new application group
// http://goo.gl/Q5n2OB
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

// UpdateGroup change parameters of a deployed application group
// http://goo.gl/Cmnpam
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

// GetGroups list all groups
// http://goo.gl/YR1KTx
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

// GetGroup gets the group with the specified ID
// http://goo.gl/DDf1AW
func (c *Client) GetGroup(groupID string) (group *Group, err error) {
	options := &RequestOptions{
		Path:   fmt.Sprintf("groups/%s", groupID),
		Method: "GET",
	}
	group = &Group{}
	err = c.unmarshalJSON(options, []int{http.StatusOK}, group)
	return
}

// DestroyGroup destroy a group
// http://goo.gl/bYQGu4
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
