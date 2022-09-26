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

// UpdateVirtualCircuitRequestOneOf1 struct for UpdateVirtualCircuitRequestOneOf1
type UpdateVirtualCircuitRequestOneOf1 struct {
	Description *string `json:"description,omitempty"`
	Name        *string `json:"name,omitempty"`
	// Speed can be changed only if it is an interconnection on a Dedicated Port
	Speed *string  `json:"speed,omitempty"`
	Tags  []string `json:"tags,omitempty"`
}

// NewUpdateVirtualCircuitRequestOneOf1 instantiates a new UpdateVirtualCircuitRequestOneOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateVirtualCircuitRequestOneOf1() *UpdateVirtualCircuitRequestOneOf1 {
	this := UpdateVirtualCircuitRequestOneOf1{}
	return &this
}

// NewUpdateVirtualCircuitRequestOneOf1WithDefaults instantiates a new UpdateVirtualCircuitRequestOneOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateVirtualCircuitRequestOneOf1WithDefaults() *UpdateVirtualCircuitRequestOneOf1 {
	this := UpdateVirtualCircuitRequestOneOf1{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateVirtualCircuitRequestOneOf1) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVirtualCircuitRequestOneOf1) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateVirtualCircuitRequestOneOf1) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateVirtualCircuitRequestOneOf1) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateVirtualCircuitRequestOneOf1) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVirtualCircuitRequestOneOf1) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateVirtualCircuitRequestOneOf1) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateVirtualCircuitRequestOneOf1) SetName(v string) {
	o.Name = &v
}

// GetSpeed returns the Speed field value if set, zero value otherwise.
func (o *UpdateVirtualCircuitRequestOneOf1) GetSpeed() string {
	if o == nil || o.Speed == nil {
		var ret string
		return ret
	}
	return *o.Speed
}

// GetSpeedOk returns a tuple with the Speed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVirtualCircuitRequestOneOf1) GetSpeedOk() (*string, bool) {
	if o == nil || o.Speed == nil {
		return nil, false
	}
	return o.Speed, true
}

// HasSpeed returns a boolean if a field has been set.
func (o *UpdateVirtualCircuitRequestOneOf1) HasSpeed() bool {
	if o != nil && o.Speed != nil {
		return true
	}

	return false
}

// SetSpeed gets a reference to the given string and assigns it to the Speed field.
func (o *UpdateVirtualCircuitRequestOneOf1) SetSpeed(v string) {
	o.Speed = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *UpdateVirtualCircuitRequestOneOf1) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVirtualCircuitRequestOneOf1) GetTagsOk() ([]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *UpdateVirtualCircuitRequestOneOf1) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *UpdateVirtualCircuitRequestOneOf1) SetTags(v []string) {
	o.Tags = v
}

func (o UpdateVirtualCircuitRequestOneOf1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Speed != nil {
		toSerialize["speed"] = o.Speed
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateVirtualCircuitRequestOneOf1 struct {
	value *UpdateVirtualCircuitRequestOneOf1
	isSet bool
}

func (v NullableUpdateVirtualCircuitRequestOneOf1) Get() *UpdateVirtualCircuitRequestOneOf1 {
	return v.value
}

func (v *NullableUpdateVirtualCircuitRequestOneOf1) Set(val *UpdateVirtualCircuitRequestOneOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateVirtualCircuitRequestOneOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateVirtualCircuitRequestOneOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateVirtualCircuitRequestOneOf1(val *UpdateVirtualCircuitRequestOneOf1) *NullableUpdateVirtualCircuitRequestOneOf1 {
	return &NullableUpdateVirtualCircuitRequestOneOf1{value: val, isSet: true}
}

func (v NullableUpdateVirtualCircuitRequestOneOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateVirtualCircuitRequestOneOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}