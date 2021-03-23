/*
 * Bitbucket API
 *
 * Code against the Bitbucket API to automate simple tasks, embed Bitbucket data into your own site, build mobile or desktop apps, or even add custom UI add-ons into Bitbucket itself using the Connect framework.
 *
 * API version: 2.0
 * Contact: support@bitbucket.org
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// WorkspaceMembership A Bitbucket workspace membership.             Links a user to a workspace.
type WorkspaceMembership struct {
	Links *BranchingModelSettingsLinks `json:"links,omitempty"`
	User *Account `json:"user,omitempty"`
	Workspace *Workspace `json:"workspace,omitempty"`
}

// NewWorkspaceMembership instantiates a new WorkspaceMembership object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkspaceMembership() *WorkspaceMembership {
	this := WorkspaceMembership{}
	return &this
}

// NewWorkspaceMembershipWithDefaults instantiates a new WorkspaceMembership object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkspaceMembershipWithDefaults() *WorkspaceMembership {
	this := WorkspaceMembership{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *WorkspaceMembership) GetLinks() BranchingModelSettingsLinks {
	if o == nil || o.Links == nil {
		var ret BranchingModelSettingsLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceMembership) GetLinksOk() (*BranchingModelSettingsLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *WorkspaceMembership) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given BranchingModelSettingsLinks and assigns it to the Links field.
func (o *WorkspaceMembership) SetLinks(v BranchingModelSettingsLinks) {
	o.Links = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *WorkspaceMembership) GetUser() Account {
	if o == nil || o.User == nil {
		var ret Account
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceMembership) GetUserOk() (*Account, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *WorkspaceMembership) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given Account and assigns it to the User field.
func (o *WorkspaceMembership) SetUser(v Account) {
	o.User = &v
}

// GetWorkspace returns the Workspace field value if set, zero value otherwise.
func (o *WorkspaceMembership) GetWorkspace() Workspace {
	if o == nil || o.Workspace == nil {
		var ret Workspace
		return ret
	}
	return *o.Workspace
}

// GetWorkspaceOk returns a tuple with the Workspace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceMembership) GetWorkspaceOk() (*Workspace, bool) {
	if o == nil || o.Workspace == nil {
		return nil, false
	}
	return o.Workspace, true
}

// HasWorkspace returns a boolean if a field has been set.
func (o *WorkspaceMembership) HasWorkspace() bool {
	if o != nil && o.Workspace != nil {
		return true
	}

	return false
}

// SetWorkspace gets a reference to the given Workspace and assigns it to the Workspace field.
func (o *WorkspaceMembership) SetWorkspace(v Workspace) {
	o.Workspace = &v
}

func (o WorkspaceMembership) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	if o.Workspace != nil {
		toSerialize["workspace"] = o.Workspace
	}
	return json.Marshal(toSerialize)
}

type NullableWorkspaceMembership struct {
	value *WorkspaceMembership
	isSet bool
}

func (v NullableWorkspaceMembership) Get() *WorkspaceMembership {
	return v.value
}

func (v *NullableWorkspaceMembership) Set(val *WorkspaceMembership) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkspaceMembership) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkspaceMembership) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkspaceMembership(val *WorkspaceMembership) *NullableWorkspaceMembership {
	return &NullableWorkspaceMembership{value: val, isSet: true}
}

func (v NullableWorkspaceMembership) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkspaceMembership) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


