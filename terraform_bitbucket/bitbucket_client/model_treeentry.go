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

// Treeentry Base type for most resource objects. It defines the common `type` element that identifies an object's type. It also identifies the element as Swagger's `discriminator`.
type Treeentry struct {
	Type string `json:"type"`
	// The path in the repository
	Path *string `json:"path,omitempty"`
	Commit *Commit `json:"commit,omitempty"`
}

// NewTreeentry instantiates a new Treeentry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTreeentry(type_ string) *Treeentry {
	this := Treeentry{}
	this.Type = type_
	return &this
}

// NewTreeentryWithDefaults instantiates a new Treeentry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTreeentryWithDefaults() *Treeentry {
	this := Treeentry{}
	return &this
}

// GetType returns the Type field value
func (o *Treeentry) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Treeentry) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Treeentry) SetType(v string) {
	o.Type = v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *Treeentry) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Treeentry) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *Treeentry) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *Treeentry) SetPath(v string) {
	o.Path = &v
}

// GetCommit returns the Commit field value if set, zero value otherwise.
func (o *Treeentry) GetCommit() Commit {
	if o == nil || o.Commit == nil {
		var ret Commit
		return ret
	}
	return *o.Commit
}

// GetCommitOk returns a tuple with the Commit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Treeentry) GetCommitOk() (*Commit, bool) {
	if o == nil || o.Commit == nil {
		return nil, false
	}
	return o.Commit, true
}

// HasCommit returns a boolean if a field has been set.
func (o *Treeentry) HasCommit() bool {
	if o != nil && o.Commit != nil {
		return true
	}

	return false
}

// SetCommit gets a reference to the given Commit and assigns it to the Commit field.
func (o *Treeentry) SetCommit(v Commit) {
	o.Commit = &v
}

func (o Treeentry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	if o.Commit != nil {
		toSerialize["commit"] = o.Commit
	}
	return json.Marshal(toSerialize)
}

type NullableTreeentry struct {
	value *Treeentry
	isSet bool
}

func (v NullableTreeentry) Get() *Treeentry {
	return v.value
}

func (v *NullableTreeentry) Set(val *Treeentry) {
	v.value = val
	v.isSet = true
}

func (v NullableTreeentry) IsSet() bool {
	return v.isSet
}

func (v *NullableTreeentry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTreeentry(val *Treeentry) *NullableTreeentry {
	return &NullableTreeentry{value: val, isSet: true}
}

func (v NullableTreeentry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTreeentry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


