/*
 * Bitbucket API
 *
 * Code against the Bitbucket API to automate simple tasks, embed Bitbucket data into your own site, build mobile or desktop apps, or even add custom UI add-ons into Bitbucket itself using the Connect framework.
 *
 * API version: 2.0
 * Contact: support@bitbucket.org
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package bitbucket_client

import (
	"encoding/json"
)

// DeploymentStateCompletedStatusSuccessful struct for DeploymentStateCompletedStatusSuccessful
type DeploymentStateCompletedStatusSuccessful struct {
	DeploymentStateCompletedStatus
	// The name of the completed deployment status (SUCCESSFUL).
	Name *string `json:"name,omitempty"`
}

// NewDeploymentStateCompletedStatusSuccessful instantiates a new DeploymentStateCompletedStatusSuccessful object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentStateCompletedStatusSuccessful(type_ string) *DeploymentStateCompletedStatusSuccessful {
	this := DeploymentStateCompletedStatusSuccessful{}
	this.Type = type_
	return &this
}

// NewDeploymentStateCompletedStatusSuccessfulWithDefaults instantiates a new DeploymentStateCompletedStatusSuccessful object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentStateCompletedStatusSuccessfulWithDefaults() *DeploymentStateCompletedStatusSuccessful {
	this := DeploymentStateCompletedStatusSuccessful{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DeploymentStateCompletedStatusSuccessful) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentStateCompletedStatusSuccessful) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DeploymentStateCompletedStatusSuccessful) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DeploymentStateCompletedStatusSuccessful) SetName(v string) {
	o.Name = &v
}

func (o DeploymentStateCompletedStatusSuccessful) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedDeploymentStateCompletedStatus, errDeploymentStateCompletedStatus := json.Marshal(o.DeploymentStateCompletedStatus)
	if errDeploymentStateCompletedStatus != nil {
		return []byte{}, errDeploymentStateCompletedStatus
	}
	errDeploymentStateCompletedStatus = json.Unmarshal([]byte(serializedDeploymentStateCompletedStatus), &toSerialize)
	if errDeploymentStateCompletedStatus != nil {
		return []byte{}, errDeploymentStateCompletedStatus
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableDeploymentStateCompletedStatusSuccessful struct {
	value *DeploymentStateCompletedStatusSuccessful
	isSet bool
}

func (v NullableDeploymentStateCompletedStatusSuccessful) Get() *DeploymentStateCompletedStatusSuccessful {
	return v.value
}

func (v *NullableDeploymentStateCompletedStatusSuccessful) Set(val *DeploymentStateCompletedStatusSuccessful) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentStateCompletedStatusSuccessful) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentStateCompletedStatusSuccessful) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentStateCompletedStatusSuccessful(val *DeploymentStateCompletedStatusSuccessful) *NullableDeploymentStateCompletedStatusSuccessful {
	return &NullableDeploymentStateCompletedStatusSuccessful{value: val, isSet: true}
}

func (v NullableDeploymentStateCompletedStatusSuccessful) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentStateCompletedStatusSuccessful) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


