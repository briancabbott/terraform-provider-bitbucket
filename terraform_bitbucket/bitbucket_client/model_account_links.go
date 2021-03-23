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

// AccountLinks struct for AccountLinks
type AccountLinks struct {
	Self *Link `json:"self,omitempty"`
	Html *Link `json:"html,omitempty"`
	Avatar *Link `json:"avatar,omitempty"`
	Followers *Link `json:"followers,omitempty"`
	Following *Link `json:"following,omitempty"`
	Repositories *Link `json:"repositories,omitempty"`
}

// NewAccountLinks instantiates a new AccountLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountLinks() *AccountLinks {
	this := AccountLinks{}
	return &this
}

// NewAccountLinksWithDefaults instantiates a new AccountLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountLinksWithDefaults() *AccountLinks {
	this := AccountLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *AccountLinks) GetSelf() Link {
	if o == nil || o.Self == nil {
		var ret Link
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountLinks) GetSelfOk() (*Link, bool) {
	if o == nil || o.Self == nil {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *AccountLinks) HasSelf() bool {
	if o != nil && o.Self != nil {
		return true
	}

	return false
}

// SetSelf gets a reference to the given Link and assigns it to the Self field.
func (o *AccountLinks) SetSelf(v Link) {
	o.Self = &v
}

// GetHtml returns the Html field value if set, zero value otherwise.
func (o *AccountLinks) GetHtml() Link {
	if o == nil || o.Html == nil {
		var ret Link
		return ret
	}
	return *o.Html
}

// GetHtmlOk returns a tuple with the Html field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountLinks) GetHtmlOk() (*Link, bool) {
	if o == nil || o.Html == nil {
		return nil, false
	}
	return o.Html, true
}

// HasHtml returns a boolean if a field has been set.
func (o *AccountLinks) HasHtml() bool {
	if o != nil && o.Html != nil {
		return true
	}

	return false
}

// SetHtml gets a reference to the given Link and assigns it to the Html field.
func (o *AccountLinks) SetHtml(v Link) {
	o.Html = &v
}

// GetAvatar returns the Avatar field value if set, zero value otherwise.
func (o *AccountLinks) GetAvatar() Link {
	if o == nil || o.Avatar == nil {
		var ret Link
		return ret
	}
	return *o.Avatar
}

// GetAvatarOk returns a tuple with the Avatar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountLinks) GetAvatarOk() (*Link, bool) {
	if o == nil || o.Avatar == nil {
		return nil, false
	}
	return o.Avatar, true
}

// HasAvatar returns a boolean if a field has been set.
func (o *AccountLinks) HasAvatar() bool {
	if o != nil && o.Avatar != nil {
		return true
	}

	return false
}

// SetAvatar gets a reference to the given Link and assigns it to the Avatar field.
func (o *AccountLinks) SetAvatar(v Link) {
	o.Avatar = &v
}

// GetFollowers returns the Followers field value if set, zero value otherwise.
func (o *AccountLinks) GetFollowers() Link {
	if o == nil || o.Followers == nil {
		var ret Link
		return ret
	}
	return *o.Followers
}

// GetFollowersOk returns a tuple with the Followers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountLinks) GetFollowersOk() (*Link, bool) {
	if o == nil || o.Followers == nil {
		return nil, false
	}
	return o.Followers, true
}

// HasFollowers returns a boolean if a field has been set.
func (o *AccountLinks) HasFollowers() bool {
	if o != nil && o.Followers != nil {
		return true
	}

	return false
}

// SetFollowers gets a reference to the given Link and assigns it to the Followers field.
func (o *AccountLinks) SetFollowers(v Link) {
	o.Followers = &v
}

// GetFollowing returns the Following field value if set, zero value otherwise.
func (o *AccountLinks) GetFollowing() Link {
	if o == nil || o.Following == nil {
		var ret Link
		return ret
	}
	return *o.Following
}

// GetFollowingOk returns a tuple with the Following field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountLinks) GetFollowingOk() (*Link, bool) {
	if o == nil || o.Following == nil {
		return nil, false
	}
	return o.Following, true
}

// HasFollowing returns a boolean if a field has been set.
func (o *AccountLinks) HasFollowing() bool {
	if o != nil && o.Following != nil {
		return true
	}

	return false
}

// SetFollowing gets a reference to the given Link and assigns it to the Following field.
func (o *AccountLinks) SetFollowing(v Link) {
	o.Following = &v
}

// GetRepositories returns the Repositories field value if set, zero value otherwise.
func (o *AccountLinks) GetRepositories() Link {
	if o == nil || o.Repositories == nil {
		var ret Link
		return ret
	}
	return *o.Repositories
}

// GetRepositoriesOk returns a tuple with the Repositories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountLinks) GetRepositoriesOk() (*Link, bool) {
	if o == nil || o.Repositories == nil {
		return nil, false
	}
	return o.Repositories, true
}

// HasRepositories returns a boolean if a field has been set.
func (o *AccountLinks) HasRepositories() bool {
	if o != nil && o.Repositories != nil {
		return true
	}

	return false
}

// SetRepositories gets a reference to the given Link and assigns it to the Repositories field.
func (o *AccountLinks) SetRepositories(v Link) {
	o.Repositories = &v
}

func (o AccountLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Self != nil {
		toSerialize["self"] = o.Self
	}
	if o.Html != nil {
		toSerialize["html"] = o.Html
	}
	if o.Avatar != nil {
		toSerialize["avatar"] = o.Avatar
	}
	if o.Followers != nil {
		toSerialize["followers"] = o.Followers
	}
	if o.Following != nil {
		toSerialize["following"] = o.Following
	}
	if o.Repositories != nil {
		toSerialize["repositories"] = o.Repositories
	}
	return json.Marshal(toSerialize)
}

type NullableAccountLinks struct {
	value *AccountLinks
	isSet bool
}

func (v NullableAccountLinks) Get() *AccountLinks {
	return v.value
}

func (v *NullableAccountLinks) Set(val *AccountLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountLinks(val *AccountLinks) *NullableAccountLinks {
	return &NullableAccountLinks{value: val, isSet: true}
}

func (v NullableAccountLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


