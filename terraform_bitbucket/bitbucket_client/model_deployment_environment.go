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

// DeploymentEnvironment struct for DeploymentEnvironment
type DeploymentEnvironment struct {
	Object
	// The UUID identifying the environment.
	Uuid *string `json:"uuid,omitempty"`
	// The name of the environment.
	Name *string `json:"name,omitempty"`
}

// NewDeploymentEnvironment instantiates a new DeploymentEnvironment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentEnvironment(type_ string) *DeploymentEnvironment {
	this := DeploymentEnvironment{}
	this.Type = type_
	return &this
}

// NewDeploymentEnvironmentWithDefaults instantiates a new DeploymentEnvironment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentEnvironmentWithDefaults() *DeploymentEnvironment {
	this := DeploymentEnvironment{}
	return &this
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *DeploymentEnvironment) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentEnvironment) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *DeploymentEnvironment) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *DeploymentEnvironment) SetUuid(v string) {
	o.Uuid = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DeploymentEnvironment) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentEnvironment) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DeploymentEnvironment) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DeploymentEnvironment) SetName(v string) {
	o.Name = &v
}

func (o DeploymentEnvironment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedObject, errObject := json.Marshal(o.Object)
	if errObject != nil {
		return []byte{}, errObject
	}
	errObject = json.Unmarshal([]byte(serializedObject), &toSerialize)
	if errObject != nil {
		return []byte{}, errObject
	}
	if o.Uuid != nil {
		toSerialize["uuid"] = o.Uuid
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableDeploymentEnvironment struct {
	value *DeploymentEnvironment
	isSet bool
}

func (v NullableDeploymentEnvironment) Get() *DeploymentEnvironment {
	return v.value
}

func (v *NullableDeploymentEnvironment) Set(val *DeploymentEnvironment) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentEnvironment) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentEnvironment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentEnvironment(val *DeploymentEnvironment) *NullableDeploymentEnvironment {
	return &NullableDeploymentEnvironment{value: val, isSet: true}
}

func (v NullableDeploymentEnvironment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentEnvironment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

