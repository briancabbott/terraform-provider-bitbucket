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

// PipelineStepStateCompletedExpiredAllOf A Bitbucket Pipelines EXPIRED pipeline step result.
type PipelineStepStateCompletedExpiredAllOf struct {
	// The name of the result (EXPIRED)
	Name *string `json:"name,omitempty"`
}

// NewPipelineStepStateCompletedExpiredAllOf instantiates a new PipelineStepStateCompletedExpiredAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPipelineStepStateCompletedExpiredAllOf() *PipelineStepStateCompletedExpiredAllOf {
	this := PipelineStepStateCompletedExpiredAllOf{}
	return &this
}

// NewPipelineStepStateCompletedExpiredAllOfWithDefaults instantiates a new PipelineStepStateCompletedExpiredAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPipelineStepStateCompletedExpiredAllOfWithDefaults() *PipelineStepStateCompletedExpiredAllOf {
	this := PipelineStepStateCompletedExpiredAllOf{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PipelineStepStateCompletedExpiredAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineStepStateCompletedExpiredAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PipelineStepStateCompletedExpiredAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PipelineStepStateCompletedExpiredAllOf) SetName(v string) {
	o.Name = &v
}

func (o PipelineStepStateCompletedExpiredAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullablePipelineStepStateCompletedExpiredAllOf struct {
	value *PipelineStepStateCompletedExpiredAllOf
	isSet bool
}

func (v NullablePipelineStepStateCompletedExpiredAllOf) Get() *PipelineStepStateCompletedExpiredAllOf {
	return v.value
}

func (v *NullablePipelineStepStateCompletedExpiredAllOf) Set(val *PipelineStepStateCompletedExpiredAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePipelineStepStateCompletedExpiredAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePipelineStepStateCompletedExpiredAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePipelineStepStateCompletedExpiredAllOf(val *PipelineStepStateCompletedExpiredAllOf) *NullablePipelineStepStateCompletedExpiredAllOf {
	return &NullablePipelineStepStateCompletedExpiredAllOf{value: val, isSet: true}
}

func (v NullablePipelineStepStateCompletedExpiredAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePipelineStepStateCompletedExpiredAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


