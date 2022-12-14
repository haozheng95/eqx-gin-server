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

// PortAssignInput struct for PortAssignInput
type PortAssignInput struct {
	// Virtual Network ID. May be the UUID of the Virtual Network record, or the VLAN value itself.
	Vnid *string `json:"vnid,omitempty"`
}

// NewPortAssignInput instantiates a new PortAssignInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortAssignInput() *PortAssignInput {
	this := PortAssignInput{}
	return &this
}

// NewPortAssignInputWithDefaults instantiates a new PortAssignInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortAssignInputWithDefaults() *PortAssignInput {
	this := PortAssignInput{}
	return &this
}

// GetVnid returns the Vnid field value if set, zero value otherwise.
func (o *PortAssignInput) GetVnid() string {
	if o == nil || o.Vnid == nil {
		var ret string
		return ret
	}
	return *o.Vnid
}

// GetVnidOk returns a tuple with the Vnid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortAssignInput) GetVnidOk() (*string, bool) {
	if o == nil || o.Vnid == nil {
		return nil, false
	}
	return o.Vnid, true
}

// HasVnid returns a boolean if a field has been set.
func (o *PortAssignInput) HasVnid() bool {
	if o != nil && o.Vnid != nil {
		return true
	}

	return false
}

// SetVnid gets a reference to the given string and assigns it to the Vnid field.
func (o *PortAssignInput) SetVnid(v string) {
	o.Vnid = &v
}

func (o PortAssignInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Vnid != nil {
		toSerialize["vnid"] = o.Vnid
	}
	return json.Marshal(toSerialize)
}

type NullablePortAssignInput struct {
	value *PortAssignInput
	isSet bool
}

func (v NullablePortAssignInput) Get() *PortAssignInput {
	return v.value
}

func (v *NullablePortAssignInput) Set(val *PortAssignInput) {
	v.value = val
	v.isSet = true
}

func (v NullablePortAssignInput) IsSet() bool {
	return v.isSet
}

func (v *NullablePortAssignInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortAssignInput(val *PortAssignInput) *NullablePortAssignInput {
	return &NullablePortAssignInput{value: val, isSet: true}
}

func (v NullablePortAssignInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortAssignInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
