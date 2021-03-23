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

// DeploymentStateCompletedStatusFailed struct for DeploymentStateCompletedStatusFailed
type DeploymentStateCompletedStatusFailed struct {
	DeploymentStateCompletedStatus
	// The name of the completed deployment status (FAILED).
	Name *string `json:"name,omitempty"`
}

// NewDeploymentStateCompletedStatusFailed instantiates a new DeploymentStateCompletedStatusFailed object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentStateCompletedStatusFailed(type_ string) *DeploymentStateCompletedStatusFailed {
	this := DeploymentStateCompletedStatusFailed{}
	this.Type = type_
	return &this
}

// NewDeploymentStateCompletedStatusFailedWithDefaults instantiates a new DeploymentStateCompletedStatusFailed object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentStateCompletedStatusFailedWithDefaults() *DeploymentStateCompletedStatusFailed {
	this := DeploymentStateCompletedStatusFailed{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DeploymentStateCompletedStatusFailed) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentStateCompletedStatusFailed) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DeploymentStateCompletedStatusFailed) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DeploymentStateCompletedStatusFailed) SetName(v string) {
	o.Name = &v
}

func (o DeploymentStateCompletedStatusFailed) MarshalJSON() ([]byte, error) {
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

type NullableDeploymentStateCompletedStatusFailed struct {
	value *DeploymentStateCompletedStatusFailed
	isSet bool
}

func (v NullableDeploymentStateCompletedStatusFailed) Get() *DeploymentStateCompletedStatusFailed {
	return v.value
}

func (v *NullableDeploymentStateCompletedStatusFailed) Set(val *DeploymentStateCompletedStatusFailed) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentStateCompletedStatusFailed) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentStateCompletedStatusFailed) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentStateCompletedStatusFailed(val *DeploymentStateCompletedStatusFailed) *NullableDeploymentStateCompletedStatusFailed {
	return &NullableDeploymentStateCompletedStatusFailed{value: val, isSet: true}
}

func (v NullableDeploymentStateCompletedStatusFailed) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentStateCompletedStatusFailed) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


