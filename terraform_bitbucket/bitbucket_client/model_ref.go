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

// Ref A ref object, representing a branch or tag in a repository.
type Ref struct {
	Type string `json:"type"`
	Links *RefLinks `json:"links,omitempty"`
	// The name of the ref.
	Name *string `json:"name,omitempty"`
	Target *Commit `json:"target,omitempty"`
}

// NewRef instantiates a new Ref object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRef(type_ string) *Ref {
	this := Ref{}
	this.Type = type_
	return &this
}

// NewRefWithDefaults instantiates a new Ref object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRefWithDefaults() *Ref {
	this := Ref{}
	return &this
}

// GetType returns the Type field value
func (o *Ref) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Ref) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Ref) SetType(v string) {
	o.Type = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *Ref) GetLinks() RefLinks {
	if o == nil || o.Links == nil {
		var ret RefLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ref) GetLinksOk() (*RefLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *Ref) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given RefLinks and assigns it to the Links field.
func (o *Ref) SetLinks(v RefLinks) {
	o.Links = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Ref) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ref) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Ref) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Ref) SetName(v string) {
	o.Name = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *Ref) GetTarget() Commit {
	if o == nil || o.Target == nil {
		var ret Commit
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ref) GetTargetOk() (*Commit, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *Ref) HasTarget() bool {
	if o != nil && o.Target != nil {
		return true
	}

	return false
}

// SetTarget gets a reference to the given Commit and assigns it to the Target field.
func (o *Ref) SetTarget(v Commit) {
	o.Target = &v
}

func (o Ref) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Target != nil {
		toSerialize["target"] = o.Target
	}
	return json.Marshal(toSerialize)
}

type NullableRef struct {
	value *Ref
	isSet bool
}

func (v NullableRef) Get() *Ref {
	return v.value
}

func (v *NullableRef) Set(val *Ref) {
	v.value = val
	v.isSet = true
}

func (v NullableRef) IsSet() bool {
	return v.isSet
}

func (v *NullableRef) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRef(val *Ref) *NullableRef {
	return &NullableRef{value: val, isSet: true}
}

func (v NullableRef) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRef) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


