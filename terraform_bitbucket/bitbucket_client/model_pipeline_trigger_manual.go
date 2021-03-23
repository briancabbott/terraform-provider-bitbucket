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

// PipelineTriggerManual struct for PipelineTriggerManual
type PipelineTriggerManual struct {
	PipelineTrigger
}

// NewPipelineTriggerManual instantiates a new PipelineTriggerManual object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPipelineTriggerManual(type_ string) *PipelineTriggerManual {
	this := PipelineTriggerManual{}
	this.Type = type_
	return &this
}

// NewPipelineTriggerManualWithDefaults instantiates a new PipelineTriggerManual object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPipelineTriggerManualWithDefaults() *PipelineTriggerManual {
	this := PipelineTriggerManual{}
	return &this
}

func (o PipelineTriggerManual) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPipelineTrigger, errPipelineTrigger := json.Marshal(o.PipelineTrigger)
	if errPipelineTrigger != nil {
		return []byte{}, errPipelineTrigger
	}
	errPipelineTrigger = json.Unmarshal([]byte(serializedPipelineTrigger), &toSerialize)
	if errPipelineTrigger != nil {
		return []byte{}, errPipelineTrigger
	}
	return json.Marshal(toSerialize)
}

type NullablePipelineTriggerManual struct {
	value *PipelineTriggerManual
	isSet bool
}

func (v NullablePipelineTriggerManual) Get() *PipelineTriggerManual {
	return v.value
}

func (v *NullablePipelineTriggerManual) Set(val *PipelineTriggerManual) {
	v.value = val
	v.isSet = true
}

func (v NullablePipelineTriggerManual) IsSet() bool {
	return v.isSet
}

func (v *NullablePipelineTriggerManual) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePipelineTriggerManual(val *PipelineTriggerManual) *NullablePipelineTriggerManual {
	return &NullablePipelineTriggerManual{value: val, isSet: true}
}

func (v NullablePipelineTriggerManual) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePipelineTriggerManual) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


