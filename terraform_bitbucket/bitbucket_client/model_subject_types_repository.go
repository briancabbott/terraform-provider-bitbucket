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

// SubjectTypesRepository struct for SubjectTypesRepository
type SubjectTypesRepository struct {
	Events *Link `json:"events,omitempty"`
}

// NewSubjectTypesRepository instantiates a new SubjectTypesRepository object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubjectTypesRepository() *SubjectTypesRepository {
	this := SubjectTypesRepository{}
	return &this
}

// NewSubjectTypesRepositoryWithDefaults instantiates a new SubjectTypesRepository object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubjectTypesRepositoryWithDefaults() *SubjectTypesRepository {
	this := SubjectTypesRepository{}
	return &this
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *SubjectTypesRepository) GetEvents() Link {
	if o == nil || o.Events == nil {
		var ret Link
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubjectTypesRepository) GetEventsOk() (*Link, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *SubjectTypesRepository) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given Link and assigns it to the Events field.
func (o *SubjectTypesRepository) SetEvents(v Link) {
	o.Events = &v
}

func (o SubjectTypesRepository) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	return json.Marshal(toSerialize)
}

type NullableSubjectTypesRepository struct {
	value *SubjectTypesRepository
	isSet bool
}

func (v NullableSubjectTypesRepository) Get() *SubjectTypesRepository {
	return v.value
}

func (v *NullableSubjectTypesRepository) Set(val *SubjectTypesRepository) {
	v.value = val
	v.isSet = true
}

func (v NullableSubjectTypesRepository) IsSet() bool {
	return v.isSet
}

func (v *NullableSubjectTypesRepository) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubjectTypesRepository(val *SubjectTypesRepository) *NullableSubjectTypesRepository {
	return &NullableSubjectTypesRepository{value: val, isSet: true}
}

func (v NullableSubjectTypesRepository) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubjectTypesRepository) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


