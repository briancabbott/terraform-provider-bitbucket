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
	"time"
)

// DeploymentStateCompleted struct for DeploymentStateCompleted
type DeploymentStateCompleted struct {
	DeploymentState
	// The name of deployment state (COMPLETED).
	Name *string `json:"name,omitempty"`
	// Link to the deployment result.
	Url *string `json:"url,omitempty"`
	Deployer *Account `json:"deployer,omitempty"`
	Status *DeploymentStateCompletedStatus `json:"status,omitempty"`
	// The timestamp when the deployment was started.
	StartDate *time.Time `json:"start_date,omitempty"`
	// The timestamp when the deployment completed.
	CompletionDate *time.Time `json:"completion_date,omitempty"`
}

// NewDeploymentStateCompleted instantiates a new DeploymentStateCompleted object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentStateCompleted(type_ string) *DeploymentStateCompleted {
	this := DeploymentStateCompleted{}
	this.Type = type_
	return &this
}

// NewDeploymentStateCompletedWithDefaults instantiates a new DeploymentStateCompleted object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentStateCompletedWithDefaults() *DeploymentStateCompleted {
	this := DeploymentStateCompleted{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DeploymentStateCompleted) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentStateCompleted) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DeploymentStateCompleted) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DeploymentStateCompleted) SetName(v string) {
	o.Name = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *DeploymentStateCompleted) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentStateCompleted) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *DeploymentStateCompleted) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *DeploymentStateCompleted) SetUrl(v string) {
	o.Url = &v
}

// GetDeployer returns the Deployer field value if set, zero value otherwise.
func (o *DeploymentStateCompleted) GetDeployer() Account {
	if o == nil || o.Deployer == nil {
		var ret Account
		return ret
	}
	return *o.Deployer
}

// GetDeployerOk returns a tuple with the Deployer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentStateCompleted) GetDeployerOk() (*Account, bool) {
	if o == nil || o.Deployer == nil {
		return nil, false
	}
	return o.Deployer, true
}

// HasDeployer returns a boolean if a field has been set.
func (o *DeploymentStateCompleted) HasDeployer() bool {
	if o != nil && o.Deployer != nil {
		return true
	}

	return false
}

// SetDeployer gets a reference to the given Account and assigns it to the Deployer field.
func (o *DeploymentStateCompleted) SetDeployer(v Account) {
	o.Deployer = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DeploymentStateCompleted) GetStatus() DeploymentStateCompletedStatus {
	if o == nil || o.Status == nil {
		var ret DeploymentStateCompletedStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentStateCompleted) GetStatusOk() (*DeploymentStateCompletedStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DeploymentStateCompleted) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given DeploymentStateCompletedStatus and assigns it to the Status field.
func (o *DeploymentStateCompleted) SetStatus(v DeploymentStateCompletedStatus) {
	o.Status = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *DeploymentStateCompleted) GetStartDate() time.Time {
	if o == nil || o.StartDate == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentStateCompleted) GetStartDateOk() (*time.Time, bool) {
	if o == nil || o.StartDate == nil {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *DeploymentStateCompleted) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *DeploymentStateCompleted) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetCompletionDate returns the CompletionDate field value if set, zero value otherwise.
func (o *DeploymentStateCompleted) GetCompletionDate() time.Time {
	if o == nil || o.CompletionDate == nil {
		var ret time.Time
		return ret
	}
	return *o.CompletionDate
}

// GetCompletionDateOk returns a tuple with the CompletionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentStateCompleted) GetCompletionDateOk() (*time.Time, bool) {
	if o == nil || o.CompletionDate == nil {
		return nil, false
	}
	return o.CompletionDate, true
}

// HasCompletionDate returns a boolean if a field has been set.
func (o *DeploymentStateCompleted) HasCompletionDate() bool {
	if o != nil && o.CompletionDate != nil {
		return true
	}

	return false
}

// SetCompletionDate gets a reference to the given time.Time and assigns it to the CompletionDate field.
func (o *DeploymentStateCompleted) SetCompletionDate(v time.Time) {
	o.CompletionDate = &v
}

func (o DeploymentStateCompleted) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedDeploymentState, errDeploymentState := json.Marshal(o.DeploymentState)
	if errDeploymentState != nil {
		return []byte{}, errDeploymentState
	}
	errDeploymentState = json.Unmarshal([]byte(serializedDeploymentState), &toSerialize)
	if errDeploymentState != nil {
		return []byte{}, errDeploymentState
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.Deployer != nil {
		toSerialize["deployer"] = o.Deployer
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.StartDate != nil {
		toSerialize["start_date"] = o.StartDate
	}
	if o.CompletionDate != nil {
		toSerialize["completion_date"] = o.CompletionDate
	}
	return json.Marshal(toSerialize)
}

type NullableDeploymentStateCompleted struct {
	value *DeploymentStateCompleted
	isSet bool
}

func (v NullableDeploymentStateCompleted) Get() *DeploymentStateCompleted {
	return v.value
}

func (v *NullableDeploymentStateCompleted) Set(val *DeploymentStateCompleted) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentStateCompleted) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentStateCompleted) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentStateCompleted(val *DeploymentStateCompleted) *NullableDeploymentStateCompleted {
	return &NullableDeploymentStateCompleted{value: val, isSet: true}
}

func (v NullableDeploymentStateCompleted) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentStateCompleted) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


