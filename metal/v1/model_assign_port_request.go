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

// AssignPortRequest struct for AssignPortRequest
type AssignPortRequest struct {
	// Virtual Network ID. May be the UUID of the Virtual Network record, or the VLAN value itself.
	Vnid *string `json:"vnid,omitempty"`
}

// NewAssignPortRequest instantiates a new AssignPortRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssignPortRequest() *AssignPortRequest {
	this := AssignPortRequest{}
	return &this
}

// NewAssignPortRequestWithDefaults instantiates a new AssignPortRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssignPortRequestWithDefaults() *AssignPortRequest {
	this := AssignPortRequest{}
	return &this
}

// GetVnid returns the Vnid field value if set, zero value otherwise.
func (o *AssignPortRequest) GetVnid() string {
	if o == nil || o.Vnid == nil {
		var ret string
		return ret
	}
	return *o.Vnid
}

// GetVnidOk returns a tuple with the Vnid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignPortRequest) GetVnidOk() (*string, bool) {
	if o == nil || o.Vnid == nil {
		return nil, false
	}
	return o.Vnid, true
}

// HasVnid returns a boolean if a field has been set.
func (o *AssignPortRequest) HasVnid() bool {
	if o != nil && o.Vnid != nil {
		return true
	}

	return false
}

// SetVnid gets a reference to the given string and assigns it to the Vnid field.
func (o *AssignPortRequest) SetVnid(v string) {
	o.Vnid = &v
}

func (o AssignPortRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Vnid != nil {
		toSerialize["vnid"] = o.Vnid
	}
	return json.Marshal(toSerialize)
}

type NullableAssignPortRequest struct {
	value *AssignPortRequest
	isSet bool
}

func (v NullableAssignPortRequest) Get() *AssignPortRequest {
	return v.value
}

func (v *NullableAssignPortRequest) Set(val *AssignPortRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAssignPortRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAssignPortRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssignPortRequest(val *AssignPortRequest) *NullableAssignPortRequest {
	return &NullableAssignPortRequest{value: val, isSet: true}
}

func (v NullableAssignPortRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssignPortRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}