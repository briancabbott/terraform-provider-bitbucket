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

// PipelineStepStateCompletedAllOf A Bitbucket Pipelines COMPLETED pipeline step state.
type PipelineStepStateCompletedAllOf struct {
	// The name of pipeline step state (COMPLETED).
	Name *string `json:"name,omitempty"`
	Result *PipelineStepStateCompletedResult `json:"result,omitempty"`
}

// NewPipelineStepStateCompletedAllOf instantiates a new PipelineStepStateCompletedAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPipelineStepStateCompletedAllOf() *PipelineStepStateCompletedAllOf {
	this := PipelineStepStateCompletedAllOf{}
	return &this
}

// NewPipelineStepStateCompletedAllOfWithDefaults instantiates a new PipelineStepStateCompletedAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPipelineStepStateCompletedAllOfWithDefaults() *PipelineStepStateCompletedAllOf {
	this := PipelineStepStateCompletedAllOf{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PipelineStepStateCompletedAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineStepStateCompletedAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PipelineStepStateCompletedAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PipelineStepStateCompletedAllOf) SetName(v string) {
	o.Name = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *PipelineStepStateCompletedAllOf) GetResult() PipelineStepStateCompletedResult {
	if o == nil || o.Result == nil {
		var ret PipelineStepStateCompletedResult
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineStepStateCompletedAllOf) GetResultOk() (*PipelineStepStateCompletedResult, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *PipelineStepStateCompletedAllOf) HasResult() bool {
	if o != nil && o.Result != nil {
		return true
	}

	return false
}

// SetResult gets a reference to the given PipelineStepStateCompletedResult and assigns it to the Result field.
func (o *PipelineStepStateCompletedAllOf) SetResult(v PipelineStepStateCompletedResult) {
	o.Result = &v
}

func (o PipelineStepStateCompletedAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}
	return json.Marshal(toSerialize)
}

type NullablePipelineStepStateCompletedAllOf struct {
	value *PipelineStepStateCompletedAllOf
	isSet bool
}

func (v NullablePipelineStepStateCompletedAllOf) Get() *PipelineStepStateCompletedAllOf {
	return v.value
}

func (v *NullablePipelineStepStateCompletedAllOf) Set(val *PipelineStepStateCompletedAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePipelineStepStateCompletedAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePipelineStepStateCompletedAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePipelineStepStateCompletedAllOf(val *PipelineStepStateCompletedAllOf) *NullablePipelineStepStateCompletedAllOf {
	return &NullablePipelineStepStateCompletedAllOf{value: val, isSet: true}
}

func (v NullablePipelineStepStateCompletedAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePipelineStepStateCompletedAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


