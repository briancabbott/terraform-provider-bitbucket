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

// PipelinesConfig struct for PipelinesConfig
type PipelinesConfig struct {
	Object
	// Whether Pipelines is enabled for the repository.
	Enabled *bool `json:"enabled,omitempty"`
	Repository *Repository `json:"repository,omitempty"`
}

// NewPipelinesConfig instantiates a new PipelinesConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPipelinesConfig(type_ string) *PipelinesConfig {
	this := PipelinesConfig{}
	this.Type = type_
	return &this
}

// NewPipelinesConfigWithDefaults instantiates a new PipelinesConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPipelinesConfigWithDefaults() *PipelinesConfig {
	this := PipelinesConfig{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *PipelinesConfig) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelinesConfig) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *PipelinesConfig) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *PipelinesConfig) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetRepository returns the Repository field value if set, zero value otherwise.
func (o *PipelinesConfig) GetRepository() Repository {
	if o == nil || o.Repository == nil {
		var ret Repository
		return ret
	}
	return *o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelinesConfig) GetRepositoryOk() (*Repository, bool) {
	if o == nil || o.Repository == nil {
		return nil, false
	}
	return o.Repository, true
}

// HasRepository returns a boolean if a field has been set.
func (o *PipelinesConfig) HasRepository() bool {
	if o != nil && o.Repository != nil {
		return true
	}

	return false
}

// SetRepository gets a reference to the given Repository and assigns it to the Repository field.
func (o *PipelinesConfig) SetRepository(v Repository) {
	o.Repository = &v
}

func (o PipelinesConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedObject, errObject := json.Marshal(o.Object)
	if errObject != nil {
		return []byte{}, errObject
	}
	errObject = json.Unmarshal([]byte(serializedObject), &toSerialize)
	if errObject != nil {
		return []byte{}, errObject
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Repository != nil {
		toSerialize["repository"] = o.Repository
	}
	return json.Marshal(toSerialize)
}

type NullablePipelinesConfig struct {
	value *PipelinesConfig
	isSet bool
}

func (v NullablePipelinesConfig) Get() *PipelinesConfig {
	return v.value
}

func (v *NullablePipelinesConfig) Set(val *PipelinesConfig) {
	v.value = val
	v.isSet = true
}

func (v NullablePipelinesConfig) IsSet() bool {
	return v.isSet
}

func (v *NullablePipelinesConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePipelinesConfig(val *PipelinesConfig) *NullablePipelinesConfig {
	return &NullablePipelinesConfig{value: val, isSet: true}
}

func (v NullablePipelinesConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePipelinesConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

