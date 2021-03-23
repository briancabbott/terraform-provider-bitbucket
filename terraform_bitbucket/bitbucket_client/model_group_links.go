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

// GroupLinks struct for GroupLinks
type GroupLinks struct {
	Self *Link `json:"self,omitempty"`
	Html *Link `json:"html,omitempty"`
}

// NewGroupLinks instantiates a new GroupLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupLinks() *GroupLinks {
	this := GroupLinks{}
	return &this
}

// NewGroupLinksWithDefaults instantiates a new GroupLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupLinksWithDefaults() *GroupLinks {
	this := GroupLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *GroupLinks) GetSelf() Link {
	if o == nil || o.Self == nil {
		var ret Link
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupLinks) GetSelfOk() (*Link, bool) {
	if o == nil || o.Self == nil {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *GroupLinks) HasSelf() bool {
	if o != nil && o.Self != nil {
		return true
	}

	return false
}

// SetSelf gets a reference to the given Link and assigns it to the Self field.
func (o *GroupLinks) SetSelf(v Link) {
	o.Self = &v
}

// GetHtml returns the Html field value if set, zero value otherwise.
func (o *GroupLinks) GetHtml() Link {
	if o == nil || o.Html == nil {
		var ret Link
		return ret
	}
	return *o.Html
}

// GetHtmlOk returns a tuple with the Html field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupLinks) GetHtmlOk() (*Link, bool) {
	if o == nil || o.Html == nil {
		return nil, false
	}
	return o.Html, true
}

// HasHtml returns a boolean if a field has been set.
func (o *GroupLinks) HasHtml() bool {
	if o != nil && o.Html != nil {
		return true
	}

	return false
}

// SetHtml gets a reference to the given Link and assigns it to the Html field.
func (o *GroupLinks) SetHtml(v Link) {
	o.Html = &v
}

func (o GroupLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Self != nil {
		toSerialize["self"] = o.Self
	}
	if o.Html != nil {
		toSerialize["html"] = o.Html
	}
	return json.Marshal(toSerialize)
}

type NullableGroupLinks struct {
	value *GroupLinks
	isSet bool
}

func (v NullableGroupLinks) Get() *GroupLinks {
	return v.value
}

func (v *NullableGroupLinks) Set(val *GroupLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupLinks(val *GroupLinks) *NullableGroupLinks {
	return &NullableGroupLinks{value: val, isSet: true}
}

func (v NullableGroupLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


