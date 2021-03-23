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

// PipelineImage The definition of a Docker image that can be used for a Bitbucket Pipelines step execution context.
type PipelineImage struct {
	// The name of the image. If the image is hosted on DockerHub the short name can be used, otherwise the fully qualified name is required here.
	Name *string `json:"name,omitempty"`
	// The username needed to authenticate with the Docker registry. Only required when using a private Docker image.
	Username *string `json:"username,omitempty"`
	// The password needed to authenticate with the Docker registry. Only required when using a private Docker image.
	Password *string `json:"password,omitempty"`
	// The email needed to authenticate with the Docker registry. Only required when using a private Docker image.
	Email *string `json:"email,omitempty"`
}

// NewPipelineImage instantiates a new PipelineImage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPipelineImage() *PipelineImage {
	this := PipelineImage{}
	return &this
}

// NewPipelineImageWithDefaults instantiates a new PipelineImage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPipelineImageWithDefaults() *PipelineImage {
	this := PipelineImage{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PipelineImage) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineImage) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PipelineImage) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PipelineImage) SetName(v string) {
	o.Name = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *PipelineImage) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineImage) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *PipelineImage) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *PipelineImage) SetUsername(v string) {
	o.Username = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *PipelineImage) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineImage) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *PipelineImage) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *PipelineImage) SetPassword(v string) {
	o.Password = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *PipelineImage) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineImage) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *PipelineImage) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *PipelineImage) SetEmail(v string) {
	o.Email = &v
}

func (o PipelineImage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	return json.Marshal(toSerialize)
}

type NullablePipelineImage struct {
	value *PipelineImage
	isSet bool
}

func (v NullablePipelineImage) Get() *PipelineImage {
	return v.value
}

func (v *NullablePipelineImage) Set(val *PipelineImage) {
	v.value = val
	v.isSet = true
}

func (v NullablePipelineImage) IsSet() bool {
	return v.isSet
}

func (v *NullablePipelineImage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePipelineImage(val *PipelineImage) *NullablePipelineImage {
	return &NullablePipelineImage{value: val, isSet: true}
}

func (v NullablePipelineImage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePipelineImage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

