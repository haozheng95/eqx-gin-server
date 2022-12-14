/*
Metal API

This is the API for Equinix Metal. The API allows you to programmatically interact with all of your Equinix Metal resources, including devices, networks, addresses, organizations, projects, and your user account.  The official API docs are hosted at <https://metal.equinix.com/developers/api>.

API version: 1.0.0
Contact: support@equinixmetal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
)

// NewPassword struct for NewPassword
type NewPassword struct {
	NewPassword *string `json:"new_password,omitempty"`
}

// NewNewPassword instantiates a new NewPassword object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewPassword() *NewPassword {
	this := NewPassword{}
	return &this
}

// NewNewPasswordWithDefaults instantiates a new NewPassword object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewPasswordWithDefaults() *NewPassword {
	this := NewPassword{}
	return &this
}

// GetNewPassword returns the NewPassword field value if set, zero value otherwise.
func (o *NewPassword) GetNewPassword() string {
	if o == nil || o.NewPassword == nil {
		var ret string
		return ret
	}
	return *o.NewPassword
}

// GetNewPasswordOk returns a tuple with the NewPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewPassword) GetNewPasswordOk() (*string, bool) {
	if o == nil || o.NewPassword == nil {
		return nil, false
	}
	return o.NewPassword, true
}

// HasNewPassword returns a boolean if a field has been set.
func (o *NewPassword) HasNewPassword() bool {
	if o != nil && o.NewPassword != nil {
		return true
	}

	return false
}

// SetNewPassword gets a reference to the given string and assigns it to the NewPassword field.
func (o *NewPassword) SetNewPassword(v string) {
	o.NewPassword = &v
}

func (o NewPassword) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NewPassword != nil {
		toSerialize["new_password"] = o.NewPassword
	}
	return json.Marshal(toSerialize)
}

type NullableNewPassword struct {
	value *NewPassword
	isSet bool
}

func (v NullableNewPassword) Get() *NewPassword {
	return v.value
}

func (v *NullableNewPassword) Set(val *NewPassword) {
	v.value = val
	v.isSet = true
}

func (v NullableNewPassword) IsSet() bool {
	return v.isSet
}

func (v *NullableNewPassword) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewPassword(val *NewPassword) *NullableNewPassword {
	return &NullableNewPassword{value: val, isSet: true}
}

func (v NullableNewPassword) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewPassword) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
