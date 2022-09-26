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

// FindDeviceById200ResponsePlanSpecsDrivesInner struct for FindDeviceById200ResponsePlanSpecsDrivesInner
type FindDeviceById200ResponsePlanSpecsDrivesInner struct {
	Count *int32  `json:"count,omitempty"`
	Type  *string `json:"type,omitempty"`
	Size  *string `json:"size,omitempty"`
}

// NewFindDeviceById200ResponsePlanSpecsDrivesInner instantiates a new FindDeviceById200ResponsePlanSpecsDrivesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFindDeviceById200ResponsePlanSpecsDrivesInner() *FindDeviceById200ResponsePlanSpecsDrivesInner {
	this := FindDeviceById200ResponsePlanSpecsDrivesInner{}
	return &this
}

// NewFindDeviceById200ResponsePlanSpecsDrivesInnerWithDefaults instantiates a new FindDeviceById200ResponsePlanSpecsDrivesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFindDeviceById200ResponsePlanSpecsDrivesInnerWithDefaults() *FindDeviceById200ResponsePlanSpecsDrivesInner {
	this := FindDeviceById200ResponsePlanSpecsDrivesInner{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *FindDeviceById200ResponsePlanSpecsDrivesInner) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindDeviceById200ResponsePlanSpecsDrivesInner) GetCountOk() (*int32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *FindDeviceById200ResponsePlanSpecsDrivesInner) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *FindDeviceById200ResponsePlanSpecsDrivesInner) SetCount(v int32) {
	o.Count = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *FindDeviceById200ResponsePlanSpecsDrivesInner) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindDeviceById200ResponsePlanSpecsDrivesInner) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *FindDeviceById200ResponsePlanSpecsDrivesInner) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *FindDeviceById200ResponsePlanSpecsDrivesInner) SetType(v string) {
	o.Type = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *FindDeviceById200ResponsePlanSpecsDrivesInner) GetSize() string {
	if o == nil || o.Size == nil {
		var ret string
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindDeviceById200ResponsePlanSpecsDrivesInner) GetSizeOk() (*string, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *FindDeviceById200ResponsePlanSpecsDrivesInner) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given string and assigns it to the Size field.
func (o *FindDeviceById200ResponsePlanSpecsDrivesInner) SetSize(v string) {
	o.Size = &v
}

func (o FindDeviceById200ResponsePlanSpecsDrivesInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	return json.Marshal(toSerialize)
}

type NullableFindDeviceById200ResponsePlanSpecsDrivesInner struct {
	value *FindDeviceById200ResponsePlanSpecsDrivesInner
	isSet bool
}

func (v NullableFindDeviceById200ResponsePlanSpecsDrivesInner) Get() *FindDeviceById200ResponsePlanSpecsDrivesInner {
	return v.value
}

func (v *NullableFindDeviceById200ResponsePlanSpecsDrivesInner) Set(val *FindDeviceById200ResponsePlanSpecsDrivesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableFindDeviceById200ResponsePlanSpecsDrivesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableFindDeviceById200ResponsePlanSpecsDrivesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFindDeviceById200ResponsePlanSpecsDrivesInner(val *FindDeviceById200ResponsePlanSpecsDrivesInner) *NullableFindDeviceById200ResponsePlanSpecsDrivesInner {
	return &NullableFindDeviceById200ResponsePlanSpecsDrivesInner{value: val, isSet: true}
}

func (v NullableFindDeviceById200ResponsePlanSpecsDrivesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFindDeviceById200ResponsePlanSpecsDrivesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}