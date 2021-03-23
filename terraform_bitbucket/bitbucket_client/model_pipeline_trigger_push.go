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

// PipelineTriggerPush struct for PipelineTriggerPush
type PipelineTriggerPush struct {
	PipelineTrigger
}

// NewPipelineTriggerPush instantiates a new PipelineTriggerPush object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPipelineTriggerPush(type_ string) *PipelineTriggerPush {
	this := PipelineTriggerPush{}
	this.Type = type_
	return &this
}

// NewPipelineTriggerPushWithDefaults instantiates a new PipelineTriggerPush object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPipelineTriggerPushWithDefaults() *PipelineTriggerPush {
	this := PipelineTriggerPush{}
	return &this
}

func (o PipelineTriggerPush) MarshalJSON() ([]byte, error) {
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

type NullablePipelineTriggerPush struct {
	value *PipelineTriggerPush
	isSet bool
}

func (v NullablePipelineTriggerPush) Get() *PipelineTriggerPush {
	return v.value
}

func (v *NullablePipelineTriggerPush) Set(val *PipelineTriggerPush) {
	v.value = val
	v.isSet = true
}

func (v NullablePipelineTriggerPush) IsSet() bool {
	return v.isSet
}

func (v *NullablePipelineTriggerPush) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePipelineTriggerPush(val *PipelineTriggerPush) *NullablePipelineTriggerPush {
	return &NullablePipelineTriggerPush{value: val, isSet: true}
}

func (v NullablePipelineTriggerPush) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePipelineTriggerPush) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

