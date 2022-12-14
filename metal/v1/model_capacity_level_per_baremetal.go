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

// CapacityLevelPerBaremetal struct for CapacityLevelPerBaremetal
type CapacityLevelPerBaremetal struct {
	Level *string `json:"level,omitempty"`
}

// NewCapacityLevelPerBaremetal instantiates a new CapacityLevelPerBaremetal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapacityLevelPerBaremetal() *CapacityLevelPerBaremetal {
	this := CapacityLevelPerBaremetal{}
	return &this
}

// NewCapacityLevelPerBaremetalWithDefaults instantiates a new CapacityLevelPerBaremetal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapacityLevelPerBaremetalWithDefaults() *CapacityLevelPerBaremetal {
	this := CapacityLevelPerBaremetal{}
	return &this
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *CapacityLevelPerBaremetal) GetLevel() string {
	if o == nil || o.Level == nil {
		var ret string
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapacityLevelPerBaremetal) GetLevelOk() (*string, bool) {
	if o == nil || o.Level == nil {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *CapacityLevelPerBaremetal) HasLevel() bool {
	if o != nil && o.Level != nil {
		return true
	}

	return false
}

// SetLevel gets a reference to the given string and assigns it to the Level field.
func (o *CapacityLevelPerBaremetal) SetLevel(v string) {
	o.Level = &v
}

func (o CapacityLevelPerBaremetal) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Level != nil {
		toSerialize["level"] = o.Level
	}
	return json.Marshal(toSerialize)
}

type NullableCapacityLevelPerBaremetal struct {
	value *CapacityLevelPerBaremetal
	isSet bool
}

func (v NullableCapacityLevelPerBaremetal) Get() *CapacityLevelPerBaremetal {
	return v.value
}

func (v *NullableCapacityLevelPerBaremetal) Set(val *CapacityLevelPerBaremetal) {
	v.value = val
	v.isSet = true
}

func (v NullableCapacityLevelPerBaremetal) IsSet() bool {
	return v.isSet
}

func (v *NullableCapacityLevelPerBaremetal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapacityLevelPerBaremetal(val *CapacityLevelPerBaremetal) *NullableCapacityLevelPerBaremetal {
	return &NullableCapacityLevelPerBaremetal{value: val, isSet: true}
}

func (v NullableCapacityLevelPerBaremetal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapacityLevelPerBaremetal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
