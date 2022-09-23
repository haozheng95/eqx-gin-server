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

// VirtualCircuitUpdateInput struct for VirtualCircuitUpdateInput
type VirtualCircuitUpdateInput struct {
	Description *string `json:"description,omitempty"`
	Name        *string `json:"name,omitempty"`
	// Speed can be changed only if it is an interconnection on a Dedicated Port
	Speed *string  `json:"speed,omitempty"`
	Tags  []string `json:"tags,omitempty"`
	// A Virtual Network record UUID or the VNID of a Virtual Network in your project.
	Vnid *string `json:"vnid,omitempty"`
}

// NewVirtualCircuitUpdateInput instantiates a new VirtualCircuitUpdateInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualCircuitUpdateInput() *VirtualCircuitUpdateInput {
	this := VirtualCircuitUpdateInput{}
	return &this
}

// NewVirtualCircuitUpdateInputWithDefaults instantiates a new VirtualCircuitUpdateInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualCircuitUpdateInputWithDefaults() *VirtualCircuitUpdateInput {
	this := VirtualCircuitUpdateInput{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *VirtualCircuitUpdateInput) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualCircuitUpdateInput) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *VirtualCircuitUpdateInput) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *VirtualCircuitUpdateInput) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VirtualCircuitUpdateInput) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualCircuitUpdateInput) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VirtualCircuitUpdateInput) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VirtualCircuitUpdateInput) SetName(v string) {
	o.Name = &v
}

// GetSpeed returns the Speed field value if set, zero value otherwise.
func (o *VirtualCircuitUpdateInput) GetSpeed() string {
	if o == nil || o.Speed == nil {
		var ret string
		return ret
	}
	return *o.Speed
}

// GetSpeedOk returns a tuple with the Speed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualCircuitUpdateInput) GetSpeedOk() (*string, bool) {
	if o == nil || o.Speed == nil {
		return nil, false
	}
	return o.Speed, true
}

// HasSpeed returns a boolean if a field has been set.
func (o *VirtualCircuitUpdateInput) HasSpeed() bool {
	if o != nil && o.Speed != nil {
		return true
	}

	return false
}

// SetSpeed gets a reference to the given string and assigns it to the Speed field.
func (o *VirtualCircuitUpdateInput) SetSpeed(v string) {
	o.Speed = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *VirtualCircuitUpdateInput) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualCircuitUpdateInput) GetTagsOk() ([]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *VirtualCircuitUpdateInput) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *VirtualCircuitUpdateInput) SetTags(v []string) {
	o.Tags = v
}

// GetVnid returns the Vnid field value if set, zero value otherwise.
func (o *VirtualCircuitUpdateInput) GetVnid() string {
	if o == nil || o.Vnid == nil {
		var ret string
		return ret
	}
	return *o.Vnid
}

// GetVnidOk returns a tuple with the Vnid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualCircuitUpdateInput) GetVnidOk() (*string, bool) {
	if o == nil || o.Vnid == nil {
		return nil, false
	}
	return o.Vnid, true
}

// HasVnid returns a boolean if a field has been set.
func (o *VirtualCircuitUpdateInput) HasVnid() bool {
	if o != nil && o.Vnid != nil {
		return true
	}

	return false
}

// SetVnid gets a reference to the given string and assigns it to the Vnid field.
func (o *VirtualCircuitUpdateInput) SetVnid(v string) {
	o.Vnid = &v
}

func (o VirtualCircuitUpdateInput) MarshalJSON() ([]byte, error) {
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
	if o.Vnid != nil {
		toSerialize["vnid"] = o.Vnid
	}
	return json.Marshal(toSerialize)
}

type NullableVirtualCircuitUpdateInput struct {
	value *VirtualCircuitUpdateInput
	isSet bool
}

func (v NullableVirtualCircuitUpdateInput) Get() *VirtualCircuitUpdateInput {
	return v.value
}

func (v *NullableVirtualCircuitUpdateInput) Set(val *VirtualCircuitUpdateInput) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualCircuitUpdateInput) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualCircuitUpdateInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualCircuitUpdateInput(val *VirtualCircuitUpdateInput) *NullableVirtualCircuitUpdateInput {
	return &NullableVirtualCircuitUpdateInput{value: val, isSet: true}
}

func (v NullableVirtualCircuitUpdateInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualCircuitUpdateInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
