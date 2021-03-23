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

// PipelineCommitTarget struct for PipelineCommitTarget
type PipelineCommitTarget struct {
	PipelineTarget
	Commit *Commit `json:"commit,omitempty"`
	Selector *PipelineSelector `json:"selector,omitempty"`
}

// NewPipelineCommitTarget instantiates a new PipelineCommitTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPipelineCommitTarget(type_ string) *PipelineCommitTarget {
	this := PipelineCommitTarget{}
	this.Type = type_
	return &this
}

// NewPipelineCommitTargetWithDefaults instantiates a new PipelineCommitTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPipelineCommitTargetWithDefaults() *PipelineCommitTarget {
	this := PipelineCommitTarget{}
	return &this
}

// GetCommit returns the Commit field value if set, zero value otherwise.
func (o *PipelineCommitTarget) GetCommit() Commit {
	if o == nil || o.Commit == nil {
		var ret Commit
		return ret
	}
	return *o.Commit
}

// GetCommitOk returns a tuple with the Commit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineCommitTarget) GetCommitOk() (*Commit, bool) {
	if o == nil || o.Commit == nil {
		return nil, false
	}
	return o.Commit, true
}

// HasCommit returns a boolean if a field has been set.
func (o *PipelineCommitTarget) HasCommit() bool {
	if o != nil && o.Commit != nil {
		return true
	}

	return false
}

// SetCommit gets a reference to the given Commit and assigns it to the Commit field.
func (o *PipelineCommitTarget) SetCommit(v Commit) {
	o.Commit = &v
}

// GetSelector returns the Selector field value if set, zero value otherwise.
func (o *PipelineCommitTarget) GetSelector() PipelineSelector {
	if o == nil || o.Selector == nil {
		var ret PipelineSelector
		return ret
	}
	return *o.Selector
}

// GetSelectorOk returns a tuple with the Selector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineCommitTarget) GetSelectorOk() (*PipelineSelector, bool) {
	if o == nil || o.Selector == nil {
		return nil, false
	}
	return o.Selector, true
}

// HasSelector returns a boolean if a field has been set.
func (o *PipelineCommitTarget) HasSelector() bool {
	if o != nil && o.Selector != nil {
		return true
	}

	return false
}

// SetSelector gets a reference to the given PipelineSelector and assigns it to the Selector field.
func (o *PipelineCommitTarget) SetSelector(v PipelineSelector) {
	o.Selector = &v
}

func (o PipelineCommitTarget) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPipelineTarget, errPipelineTarget := json.Marshal(o.PipelineTarget)
	if errPipelineTarget != nil {
		return []byte{}, errPipelineTarget
	}
	errPipelineTarget = json.Unmarshal([]byte(serializedPipelineTarget), &toSerialize)
	if errPipelineTarget != nil {
		return []byte{}, errPipelineTarget
	}
	if o.Commit != nil {
		toSerialize["commit"] = o.Commit
	}
	if o.Selector != nil {
		toSerialize["selector"] = o.Selector
	}
	return json.Marshal(toSerialize)
}

type NullablePipelineCommitTarget struct {
	value *PipelineCommitTarget
	isSet bool
}

func (v NullablePipelineCommitTarget) Get() *PipelineCommitTarget {
	return v.value
}

func (v *NullablePipelineCommitTarget) Set(val *PipelineCommitTarget) {
	v.value = val
	v.isSet = true
}

func (v NullablePipelineCommitTarget) IsSet() bool {
	return v.isSet
}

func (v *NullablePipelineCommitTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePipelineCommitTarget(val *PipelineCommitTarget) *NullablePipelineCommitTarget {
	return &NullablePipelineCommitTarget{value: val, isSet: true}
}

func (v NullablePipelineCommitTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePipelineCommitTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


