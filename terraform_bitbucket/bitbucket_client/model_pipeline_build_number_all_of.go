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

// PipelineBuildNumberAllOf A Pipelines build number.
type PipelineBuildNumberAllOf struct {
	// The next number that will be used as build number.
	Next *int32 `json:"next,omitempty"`
}

// NewPipelineBuildNumberAllOf instantiates a new PipelineBuildNumberAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPipelineBuildNumberAllOf() *PipelineBuildNumberAllOf {
	this := PipelineBuildNumberAllOf{}
	return &this
}

// NewPipelineBuildNumberAllOfWithDefaults instantiates a new PipelineBuildNumberAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPipelineBuildNumberAllOfWithDefaults() *PipelineBuildNumberAllOf {
	this := PipelineBuildNumberAllOf{}
	return &this
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *PipelineBuildNumberAllOf) GetNext() int32 {
	if o == nil || o.Next == nil {
		var ret int32
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineBuildNumberAllOf) GetNextOk() (*int32, bool) {
	if o == nil || o.Next == nil {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *PipelineBuildNumberAllOf) HasNext() bool {
	if o != nil && o.Next != nil {
		return true
	}

	return false
}

// SetNext gets a reference to the given int32 and assigns it to the Next field.
func (o *PipelineBuildNumberAllOf) SetNext(v int32) {
	o.Next = &v
}

func (o PipelineBuildNumberAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Next != nil {
		toSerialize["next"] = o.Next
	}
	return json.Marshal(toSerialize)
}

type NullablePipelineBuildNumberAllOf struct {
	value *PipelineBuildNumberAllOf
	isSet bool
}

func (v NullablePipelineBuildNumberAllOf) Get() *PipelineBuildNumberAllOf {
	return v.value
}

func (v *NullablePipelineBuildNumberAllOf) Set(val *PipelineBuildNumberAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePipelineBuildNumberAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePipelineBuildNumberAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePipelineBuildNumberAllOf(val *PipelineBuildNumberAllOf) *NullablePipelineBuildNumberAllOf {
	return &NullablePipelineBuildNumberAllOf{value: val, isSet: true}
}

func (v NullablePipelineBuildNumberAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePipelineBuildNumberAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


