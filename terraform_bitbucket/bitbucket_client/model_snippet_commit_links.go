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

// SnippetCommitLinks struct for SnippetCommitLinks
type SnippetCommitLinks struct {
	Self *Link `json:"self,omitempty"`
	Html *Link `json:"html,omitempty"`
	Diff *Link `json:"diff,omitempty"`
}

// NewSnippetCommitLinks instantiates a new SnippetCommitLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnippetCommitLinks() *SnippetCommitLinks {
	this := SnippetCommitLinks{}
	return &this
}

// NewSnippetCommitLinksWithDefaults instantiates a new SnippetCommitLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnippetCommitLinksWithDefaults() *SnippetCommitLinks {
	this := SnippetCommitLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *SnippetCommitLinks) GetSelf() Link {
	if o == nil || o.Self == nil {
		var ret Link
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnippetCommitLinks) GetSelfOk() (*Link, bool) {
	if o == nil || o.Self == nil {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *SnippetCommitLinks) HasSelf() bool {
	if o != nil && o.Self != nil {
		return true
	}

	return false
}

// SetSelf gets a reference to the given Link and assigns it to the Self field.
func (o *SnippetCommitLinks) SetSelf(v Link) {
	o.Self = &v
}

// GetHtml returns the Html field value if set, zero value otherwise.
func (o *SnippetCommitLinks) GetHtml() Link {
	if o == nil || o.Html == nil {
		var ret Link
		return ret
	}
	return *o.Html
}

// GetHtmlOk returns a tuple with the Html field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnippetCommitLinks) GetHtmlOk() (*Link, bool) {
	if o == nil || o.Html == nil {
		return nil, false
	}
	return o.Html, true
}

// HasHtml returns a boolean if a field has been set.
func (o *SnippetCommitLinks) HasHtml() bool {
	if o != nil && o.Html != nil {
		return true
	}

	return false
}

// SetHtml gets a reference to the given Link and assigns it to the Html field.
func (o *SnippetCommitLinks) SetHtml(v Link) {
	o.Html = &v
}

// GetDiff returns the Diff field value if set, zero value otherwise.
func (o *SnippetCommitLinks) GetDiff() Link {
	if o == nil || o.Diff == nil {
		var ret Link
		return ret
	}
	return *o.Diff
}

// GetDiffOk returns a tuple with the Diff field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnippetCommitLinks) GetDiffOk() (*Link, bool) {
	if o == nil || o.Diff == nil {
		return nil, false
	}
	return o.Diff, true
}

// HasDiff returns a boolean if a field has been set.
func (o *SnippetCommitLinks) HasDiff() bool {
	if o != nil && o.Diff != nil {
		return true
	}

	return false
}

// SetDiff gets a reference to the given Link and assigns it to the Diff field.
func (o *SnippetCommitLinks) SetDiff(v Link) {
	o.Diff = &v
}

func (o SnippetCommitLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Self != nil {
		toSerialize["self"] = o.Self
	}
	if o.Html != nil {
		toSerialize["html"] = o.Html
	}
	if o.Diff != nil {
		toSerialize["diff"] = o.Diff
	}
	return json.Marshal(toSerialize)
}

type NullableSnippetCommitLinks struct {
	value *SnippetCommitLinks
	isSet bool
}

func (v NullableSnippetCommitLinks) Get() *SnippetCommitLinks {
	return v.value
}

func (v *NullableSnippetCommitLinks) Set(val *SnippetCommitLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableSnippetCommitLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableSnippetCommitLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnippetCommitLinks(val *SnippetCommitLinks) *NullableSnippetCommitLinks {
	return &NullableSnippetCommitLinks{value: val, isSet: true}
}

func (v NullableSnippetCommitLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnippetCommitLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


