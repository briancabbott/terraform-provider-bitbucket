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
	"time"
)

// Account An account object.
type Account struct {
	Links *AccountLinks `json:"links,omitempty"`
	Username *string `json:"username,omitempty"`
	// Account name defined by the owner. Should be used instead of the \"username\" field. Note that \"nickname\" cannot be used in place of \"username\" in URLs and queries, as \"nickname\" is not guaranteed to be unique.
	Nickname *string `json:"nickname,omitempty"`
	// The status of the account. Currently the only possible value is \"active\", but more values may be added in the future.
	AccountStatus *string `json:"account_status,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`
	Website *string `json:"website,omitempty"`
	CreatedOn *time.Time `json:"created_on,omitempty"`
	Uuid *string `json:"uuid,omitempty"`
	Has2faEnabled *bool `json:"has_2fa_enabled,omitempty"`
}

// NewAccount instantiates a new Account object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccount() *Account {
	this := Account{}
	return &this
}

// NewAccountWithDefaults instantiates a new Account object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountWithDefaults() *Account {
	this := Account{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *Account) GetLinks() AccountLinks {
	if o == nil || o.Links == nil {
		var ret AccountLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetLinksOk() (*AccountLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *Account) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AccountLinks and assigns it to the Links field.
func (o *Account) SetLinks(v AccountLinks) {
	o.Links = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *Account) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *Account) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *Account) SetUsername(v string) {
	o.Username = &v
}

// GetNickname returns the Nickname field value if set, zero value otherwise.
func (o *Account) GetNickname() string {
	if o == nil || o.Nickname == nil {
		var ret string
		return ret
	}
	return *o.Nickname
}

// GetNicknameOk returns a tuple with the Nickname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetNicknameOk() (*string, bool) {
	if o == nil || o.Nickname == nil {
		return nil, false
	}
	return o.Nickname, true
}

// HasNickname returns a boolean if a field has been set.
func (o *Account) HasNickname() bool {
	if o != nil && o.Nickname != nil {
		return true
	}

	return false
}

// SetNickname gets a reference to the given string and assigns it to the Nickname field.
func (o *Account) SetNickname(v string) {
	o.Nickname = &v
}

// GetAccountStatus returns the AccountStatus field value if set, zero value otherwise.
func (o *Account) GetAccountStatus() string {
	if o == nil || o.AccountStatus == nil {
		var ret string
		return ret
	}
	return *o.AccountStatus
}

// GetAccountStatusOk returns a tuple with the AccountStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetAccountStatusOk() (*string, bool) {
	if o == nil || o.AccountStatus == nil {
		return nil, false
	}
	return o.AccountStatus, true
}

// HasAccountStatus returns a boolean if a field has been set.
func (o *Account) HasAccountStatus() bool {
	if o != nil && o.AccountStatus != nil {
		return true
	}

	return false
}

// SetAccountStatus gets a reference to the given string and assigns it to the AccountStatus field.
func (o *Account) SetAccountStatus(v string) {
	o.AccountStatus = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *Account) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *Account) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *Account) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetWebsite returns the Website field value if set, zero value otherwise.
func (o *Account) GetWebsite() string {
	if o == nil || o.Website == nil {
		var ret string
		return ret
	}
	return *o.Website
}

// GetWebsiteOk returns a tuple with the Website field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetWebsiteOk() (*string, bool) {
	if o == nil || o.Website == nil {
		return nil, false
	}
	return o.Website, true
}

// HasWebsite returns a boolean if a field has been set.
func (o *Account) HasWebsite() bool {
	if o != nil && o.Website != nil {
		return true
	}

	return false
}

// SetWebsite gets a reference to the given string and assigns it to the Website field.
func (o *Account) SetWebsite(v string) {
	o.Website = &v
}

// GetCreatedOn returns the CreatedOn field value if set, zero value otherwise.
func (o *Account) GetCreatedOn() time.Time {
	if o == nil || o.CreatedOn == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetCreatedOnOk() (*time.Time, bool) {
	if o == nil || o.CreatedOn == nil {
		return nil, false
	}
	return o.CreatedOn, true
}

// HasCreatedOn returns a boolean if a field has been set.
func (o *Account) HasCreatedOn() bool {
	if o != nil && o.CreatedOn != nil {
		return true
	}

	return false
}

// SetCreatedOn gets a reference to the given time.Time and assigns it to the CreatedOn field.
func (o *Account) SetCreatedOn(v time.Time) {
	o.CreatedOn = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *Account) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *Account) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *Account) SetUuid(v string) {
	o.Uuid = &v
}

// GetHas2faEnabled returns the Has2faEnabled field value if set, zero value otherwise.
func (o *Account) GetHas2faEnabled() bool {
	if o == nil || o.Has2faEnabled == nil {
		var ret bool
		return ret
	}
	return *o.Has2faEnabled
}

// GetHas2faEnabledOk returns a tuple with the Has2faEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetHas2faEnabledOk() (*bool, bool) {
	if o == nil || o.Has2faEnabled == nil {
		return nil, false
	}
	return o.Has2faEnabled, true
}

// HasHas2faEnabled returns a boolean if a field has been set.
func (o *Account) HasHas2faEnabled() bool {
	if o != nil && o.Has2faEnabled != nil {
		return true
	}

	return false
}

// SetHas2faEnabled gets a reference to the given bool and assigns it to the Has2faEnabled field.
func (o *Account) SetHas2faEnabled(v bool) {
	o.Has2faEnabled = &v
}

func (o Account) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	if o.Nickname != nil {
		toSerialize["nickname"] = o.Nickname
	}
	if o.AccountStatus != nil {
		toSerialize["account_status"] = o.AccountStatus
	}
	if o.DisplayName != nil {
		toSerialize["display_name"] = o.DisplayName
	}
	if o.Website != nil {
		toSerialize["website"] = o.Website
	}
	if o.CreatedOn != nil {
		toSerialize["created_on"] = o.CreatedOn
	}
	if o.Uuid != nil {
		toSerialize["uuid"] = o.Uuid
	}
	if o.Has2faEnabled != nil {
		toSerialize["has_2fa_enabled"] = o.Has2faEnabled
	}
	return json.Marshal(toSerialize)
}

type NullableAccount struct {
	value *Account
	isSet bool
}

func (v NullableAccount) Get() *Account {
	return v.value
}

func (v *NullableAccount) Set(val *Account) {
	v.value = val
	v.isSet = true
}

func (v NullableAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccount(val *Account) *NullableAccount {
	return &NullableAccount{value: val, isSet: true}
}

func (v NullableAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

