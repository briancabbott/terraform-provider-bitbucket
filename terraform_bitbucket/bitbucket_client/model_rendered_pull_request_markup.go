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

// RenderedPullRequestMarkup User provided pull request text, interpreted in a markup language and rendered in HTML
type RenderedPullRequestMarkup struct {
	Title *IssueContent `json:"title,omitempty"`
	Description *IssueContent `json:"description,omitempty"`
	Reason *IssueContent `json:"reason,omitempty"`
}

// NewRenderedPullRequestMarkup instantiates a new RenderedPullRequestMarkup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRenderedPullRequestMarkup() *RenderedPullRequestMarkup {
	this := RenderedPullRequestMarkup{}
	return &this
}

// NewRenderedPullRequestMarkupWithDefaults instantiates a new RenderedPullRequestMarkup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRenderedPullRequestMarkupWithDefaults() *RenderedPullRequestMarkup {
	this := RenderedPullRequestMarkup{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *RenderedPullRequestMarkup) GetTitle() IssueContent {
	if o == nil || o.Title == nil {
		var ret IssueContent
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenderedPullRequestMarkup) GetTitleOk() (*IssueContent, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *RenderedPullRequestMarkup) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given IssueContent and assigns it to the Title field.
func (o *RenderedPullRequestMarkup) SetTitle(v IssueContent) {
	o.Title = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RenderedPullRequestMarkup) GetDescription() IssueContent {
	if o == nil || o.Description == nil {
		var ret IssueContent
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenderedPullRequestMarkup) GetDescriptionOk() (*IssueContent, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RenderedPullRequestMarkup) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given IssueContent and assigns it to the Description field.
func (o *RenderedPullRequestMarkup) SetDescription(v IssueContent) {
	o.Description = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *RenderedPullRequestMarkup) GetReason() IssueContent {
	if o == nil || o.Reason == nil {
		var ret IssueContent
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenderedPullRequestMarkup) GetReasonOk() (*IssueContent, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *RenderedPullRequestMarkup) HasReason() bool {
	if o != nil && o.Reason != nil {
		return true
	}

	return false
}

// SetReason gets a reference to the given IssueContent and assigns it to the Reason field.
func (o *RenderedPullRequestMarkup) SetReason(v IssueContent) {
	o.Reason = &v
}

func (o RenderedPullRequestMarkup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}
	return json.Marshal(toSerialize)
}

type NullableRenderedPullRequestMarkup struct {
	value *RenderedPullRequestMarkup
	isSet bool
}

func (v NullableRenderedPullRequestMarkup) Get() *RenderedPullRequestMarkup {
	return v.value
}

func (v *NullableRenderedPullRequestMarkup) Set(val *RenderedPullRequestMarkup) {
	v.value = val
	v.isSet = true
}

func (v NullableRenderedPullRequestMarkup) IsSet() bool {
	return v.isSet
}

func (v *NullableRenderedPullRequestMarkup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRenderedPullRequestMarkup(val *RenderedPullRequestMarkup) *NullableRenderedPullRequestMarkup {
	return &NullableRenderedPullRequestMarkup{value: val, isSet: true}
}

func (v NullableRenderedPullRequestMarkup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRenderedPullRequestMarkup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


