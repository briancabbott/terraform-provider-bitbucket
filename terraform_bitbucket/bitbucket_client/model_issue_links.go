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

// IssueLinks struct for IssueLinks
type IssueLinks struct {
	Self *Link `json:"self,omitempty"`
	Html *Link `json:"html,omitempty"`
	Comments *Link `json:"comments,omitempty"`
	Attachments *Link `json:"attachments,omitempty"`
	Watch *Link `json:"watch,omitempty"`
	Vote *Link `json:"vote,omitempty"`
}

// NewIssueLinks instantiates a new IssueLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIssueLinks() *IssueLinks {
	this := IssueLinks{}
	return &this
}

// NewIssueLinksWithDefaults instantiates a new IssueLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIssueLinksWithDefaults() *IssueLinks {
	this := IssueLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *IssueLinks) GetSelf() Link {
	if o == nil || o.Self == nil {
		var ret Link
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueLinks) GetSelfOk() (*Link, bool) {
	if o == nil || o.Self == nil {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *IssueLinks) HasSelf() bool {
	if o != nil && o.Self != nil {
		return true
	}

	return false
}

// SetSelf gets a reference to the given Link and assigns it to the Self field.
func (o *IssueLinks) SetSelf(v Link) {
	o.Self = &v
}

// GetHtml returns the Html field value if set, zero value otherwise.
func (o *IssueLinks) GetHtml() Link {
	if o == nil || o.Html == nil {
		var ret Link
		return ret
	}
	return *o.Html
}

// GetHtmlOk returns a tuple with the Html field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueLinks) GetHtmlOk() (*Link, bool) {
	if o == nil || o.Html == nil {
		return nil, false
	}
	return o.Html, true
}

// HasHtml returns a boolean if a field has been set.
func (o *IssueLinks) HasHtml() bool {
	if o != nil && o.Html != nil {
		return true
	}

	return false
}

// SetHtml gets a reference to the given Link and assigns it to the Html field.
func (o *IssueLinks) SetHtml(v Link) {
	o.Html = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *IssueLinks) GetComments() Link {
	if o == nil || o.Comments == nil {
		var ret Link
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueLinks) GetCommentsOk() (*Link, bool) {
	if o == nil || o.Comments == nil {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *IssueLinks) HasComments() bool {
	if o != nil && o.Comments != nil {
		return true
	}

	return false
}

// SetComments gets a reference to the given Link and assigns it to the Comments field.
func (o *IssueLinks) SetComments(v Link) {
	o.Comments = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *IssueLinks) GetAttachments() Link {
	if o == nil || o.Attachments == nil {
		var ret Link
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueLinks) GetAttachmentsOk() (*Link, bool) {
	if o == nil || o.Attachments == nil {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *IssueLinks) HasAttachments() bool {
	if o != nil && o.Attachments != nil {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given Link and assigns it to the Attachments field.
func (o *IssueLinks) SetAttachments(v Link) {
	o.Attachments = &v
}

// GetWatch returns the Watch field value if set, zero value otherwise.
func (o *IssueLinks) GetWatch() Link {
	if o == nil || o.Watch == nil {
		var ret Link
		return ret
	}
	return *o.Watch
}

// GetWatchOk returns a tuple with the Watch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueLinks) GetWatchOk() (*Link, bool) {
	if o == nil || o.Watch == nil {
		return nil, false
	}
	return o.Watch, true
}

// HasWatch returns a boolean if a field has been set.
func (o *IssueLinks) HasWatch() bool {
	if o != nil && o.Watch != nil {
		return true
	}

	return false
}

// SetWatch gets a reference to the given Link and assigns it to the Watch field.
func (o *IssueLinks) SetWatch(v Link) {
	o.Watch = &v
}

// GetVote returns the Vote field value if set, zero value otherwise.
func (o *IssueLinks) GetVote() Link {
	if o == nil || o.Vote == nil {
		var ret Link
		return ret
	}
	return *o.Vote
}

// GetVoteOk returns a tuple with the Vote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueLinks) GetVoteOk() (*Link, bool) {
	if o == nil || o.Vote == nil {
		return nil, false
	}
	return o.Vote, true
}

// HasVote returns a boolean if a field has been set.
func (o *IssueLinks) HasVote() bool {
	if o != nil && o.Vote != nil {
		return true
	}

	return false
}

// SetVote gets a reference to the given Link and assigns it to the Vote field.
func (o *IssueLinks) SetVote(v Link) {
	o.Vote = &v
}

func (o IssueLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Self != nil {
		toSerialize["self"] = o.Self
	}
	if o.Html != nil {
		toSerialize["html"] = o.Html
	}
	if o.Comments != nil {
		toSerialize["comments"] = o.Comments
	}
	if o.Attachments != nil {
		toSerialize["attachments"] = o.Attachments
	}
	if o.Watch != nil {
		toSerialize["watch"] = o.Watch
	}
	if o.Vote != nil {
		toSerialize["vote"] = o.Vote
	}
	return json.Marshal(toSerialize)
}

type NullableIssueLinks struct {
	value *IssueLinks
	isSet bool
}

func (v NullableIssueLinks) Get() *IssueLinks {
	return v.value
}

func (v *NullableIssueLinks) Set(val *IssueLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableIssueLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableIssueLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIssueLinks(val *IssueLinks) *NullableIssueLinks {
	return &NullableIssueLinks{value: val, isSet: true}
}

func (v NullableIssueLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIssueLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


