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

// PipelineStepStateCompletedStopped struct for PipelineStepStateCompletedStopped
type PipelineStepStateCompletedStopped struct {
	PipelineStepStateCompletedResult
	// The name of the result (STOPPED)
	Name *string `json:"name,omitempty"`
}

// NewPipelineStepStateCompletedStopped instantiates a new PipelineStepStateCompletedStopped object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPipelineStepStateCompletedStopped(type_ string) *PipelineStepStateCompletedStopped {
	this := PipelineStepStateCompletedStopped{}
	this.Type = type_
	return &this
}

// NewPipelineStepStateCompletedStoppedWithDefaults instantiates a new PipelineStepStateCompletedStopped object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPipelineStepStateCompletedStoppedWithDefaults() *PipelineStepStateCompletedStopped {
	this := PipelineStepStateCompletedStopped{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PipelineStepStateCompletedStopped) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineStepStateCompletedStopped) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PipelineStepStateCompletedStopped) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PipelineStepStateCompletedStopped) SetName(v string) {
	o.Name = &v
}

func (o PipelineStepStateCompletedStopped) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPipelineStepStateCompletedResult, errPipelineStepStateCompletedResult := json.Marshal(o.PipelineStepStateCompletedResult)
	if errPipelineStepStateCompletedResult != nil {
		return []byte{}, errPipelineStepStateCompletedResult
	}
	errPipelineStepStateCompletedResult = json.Unmarshal([]byte(serializedPipelineStepStateCompletedResult), &toSerialize)
	if errPipelineStepStateCompletedResult != nil {
		return []byte{}, errPipelineStepStateCompletedResult
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullablePipelineStepStateCompletedStopped struct {
	value *PipelineStepStateCompletedStopped
	isSet bool
}

func (v NullablePipelineStepStateCompletedStopped) Get() *PipelineStepStateCompletedStopped {
	return v.value
}

func (v *NullablePipelineStepStateCompletedStopped) Set(val *PipelineStepStateCompletedStopped) {
	v.value = val
	v.isSet = true
}

func (v NullablePipelineStepStateCompletedStopped) IsSet() bool {
	return v.isSet
}

func (v *NullablePipelineStepStateCompletedStopped) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePipelineStepStateCompletedStopped(val *PipelineStepStateCompletedStopped) *NullablePipelineStepStateCompletedStopped {
	return &NullablePipelineStepStateCompletedStopped{value: val, isSet: true}
}

func (v NullablePipelineStepStateCompletedStopped) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePipelineStepStateCompletedStopped) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


