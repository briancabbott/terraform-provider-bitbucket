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

// PipelineStepStateCompletedResult struct for PipelineStepStateCompletedResult
type PipelineStepStateCompletedResult struct {
	Object
}

// NewPipelineStepStateCompletedResult instantiates a new PipelineStepStateCompletedResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPipelineStepStateCompletedResult(type_ string) *PipelineStepStateCompletedResult {
	this := PipelineStepStateCompletedResult{}
	this.Type = type_
	return &this
}

// NewPipelineStepStateCompletedResultWithDefaults instantiates a new PipelineStepStateCompletedResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPipelineStepStateCompletedResultWithDefaults() *PipelineStepStateCompletedResult {
	this := PipelineStepStateCompletedResult{}
	return &this
}

func (o PipelineStepStateCompletedResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedObject, errObject := json.Marshal(o.Object)
	if errObject != nil {
		return []byte{}, errObject
	}
	errObject = json.Unmarshal([]byte(serializedObject), &toSerialize)
	if errObject != nil {
		return []byte{}, errObject
	}
	return json.Marshal(toSerialize)
}

type NullablePipelineStepStateCompletedResult struct {
	value *PipelineStepStateCompletedResult
	isSet bool
}

func (v NullablePipelineStepStateCompletedResult) Get() *PipelineStepStateCompletedResult {
	return v.value
}

func (v *NullablePipelineStepStateCompletedResult) Set(val *PipelineStepStateCompletedResult) {
	v.value = val
	v.isSet = true
}

func (v NullablePipelineStepStateCompletedResult) IsSet() bool {
	return v.isSet
}

func (v *NullablePipelineStepStateCompletedResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePipelineStepStateCompletedResult(val *PipelineStepStateCompletedResult) *NullablePipelineStepStateCompletedResult {
	return &NullablePipelineStepStateCompletedResult{value: val, isSet: true}
}

func (v NullablePipelineStepStateCompletedResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePipelineStepStateCompletedResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


