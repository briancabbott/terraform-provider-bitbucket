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

// DeploymentStateCompletedStatusStoppedAllOf A STOPPED completed deployment status.
type DeploymentStateCompletedStatusStoppedAllOf struct {
	// The name of the completed deployment status (STOPPED).
	Name *string `json:"name,omitempty"`
}

// NewDeploymentStateCompletedStatusStoppedAllOf instantiates a new DeploymentStateCompletedStatusStoppedAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentStateCompletedStatusStoppedAllOf() *DeploymentStateCompletedStatusStoppedAllOf {
	this := DeploymentStateCompletedStatusStoppedAllOf{}
	return &this
}

// NewDeploymentStateCompletedStatusStoppedAllOfWithDefaults instantiates a new DeploymentStateCompletedStatusStoppedAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentStateCompletedStatusStoppedAllOfWithDefaults() *DeploymentStateCompletedStatusStoppedAllOf {
	this := DeploymentStateCompletedStatusStoppedAllOf{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DeploymentStateCompletedStatusStoppedAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentStateCompletedStatusStoppedAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DeploymentStateCompletedStatusStoppedAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DeploymentStateCompletedStatusStoppedAllOf) SetName(v string) {
	o.Name = &v
}

func (o DeploymentStateCompletedStatusStoppedAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableDeploymentStateCompletedStatusStoppedAllOf struct {
	value *DeploymentStateCompletedStatusStoppedAllOf
	isSet bool
}

func (v NullableDeploymentStateCompletedStatusStoppedAllOf) Get() *DeploymentStateCompletedStatusStoppedAllOf {
	return v.value
}

func (v *NullableDeploymentStateCompletedStatusStoppedAllOf) Set(val *DeploymentStateCompletedStatusStoppedAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentStateCompletedStatusStoppedAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentStateCompletedStatusStoppedAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentStateCompletedStatusStoppedAllOf(val *DeploymentStateCompletedStatusStoppedAllOf) *NullableDeploymentStateCompletedStatusStoppedAllOf {
	return &NullableDeploymentStateCompletedStatusStoppedAllOf{value: val, isSet: true}
}

func (v NullableDeploymentStateCompletedStatusStoppedAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentStateCompletedStatusStoppedAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


