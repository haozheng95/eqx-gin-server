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

// FindDeviceById200ResponsePlanSpecsFeatures struct for FindDeviceById200ResponsePlanSpecsFeatures
type FindDeviceById200ResponsePlanSpecsFeatures struct {
	Raid *bool `json:"raid,omitempty"`
	Txt  *bool `json:"txt,omitempty"`
}

// NewFindDeviceById200ResponsePlanSpecsFeatures instantiates a new FindDeviceById200ResponsePlanSpecsFeatures object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFindDeviceById200ResponsePlanSpecsFeatures() *FindDeviceById200ResponsePlanSpecsFeatures {
	this := FindDeviceById200ResponsePlanSpecsFeatures{}
	return &this
}

// NewFindDeviceById200ResponsePlanSpecsFeaturesWithDefaults instantiates a new FindDeviceById200ResponsePlanSpecsFeatures object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFindDeviceById200ResponsePlanSpecsFeaturesWithDefaults() *FindDeviceById200ResponsePlanSpecsFeatures {
	this := FindDeviceById200ResponsePlanSpecsFeatures{}
	return &this
}

// GetRaid returns the Raid field value if set, zero value otherwise.
func (o *FindDeviceById200ResponsePlanSpecsFeatures) GetRaid() bool {
	if o == nil || o.Raid == nil {
		var ret bool
		return ret
	}
	return *o.Raid
}

// GetRaidOk returns a tuple with the Raid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindDeviceById200ResponsePlanSpecsFeatures) GetRaidOk() (*bool, bool) {
	if o == nil || o.Raid == nil {
		return nil, false
	}
	return o.Raid, true
}

// HasRaid returns a boolean if a field has been set.
func (o *FindDeviceById200ResponsePlanSpecsFeatures) HasRaid() bool {
	if o != nil && o.Raid != nil {
		return true
	}

	return false
}

// SetRaid gets a reference to the given bool and assigns it to the Raid field.
func (o *FindDeviceById200ResponsePlanSpecsFeatures) SetRaid(v bool) {
	o.Raid = &v
}

// GetTxt returns the Txt field value if set, zero value otherwise.
func (o *FindDeviceById200ResponsePlanSpecsFeatures) GetTxt() bool {
	if o == nil || o.Txt == nil {
		var ret bool
		return ret
	}
	return *o.Txt
}

// GetTxtOk returns a tuple with the Txt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindDeviceById200ResponsePlanSpecsFeatures) GetTxtOk() (*bool, bool) {
	if o == nil || o.Txt == nil {
		return nil, false
	}
	return o.Txt, true
}

// HasTxt returns a boolean if a field has been set.
func (o *FindDeviceById200ResponsePlanSpecsFeatures) HasTxt() bool {
	if o != nil && o.Txt != nil {
		return true
	}

	return false
}

// SetTxt gets a reference to the given bool and assigns it to the Txt field.
func (o *FindDeviceById200ResponsePlanSpecsFeatures) SetTxt(v bool) {
	o.Txt = &v
}

func (o FindDeviceById200ResponsePlanSpecsFeatures) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Raid != nil {
		toSerialize["raid"] = o.Raid
	}
	if o.Txt != nil {
		toSerialize["txt"] = o.Txt
	}
	return json.Marshal(toSerialize)
}

type NullableFindDeviceById200ResponsePlanSpecsFeatures struct {
	value *FindDeviceById200ResponsePlanSpecsFeatures
	isSet bool
}

func (v NullableFindDeviceById200ResponsePlanSpecsFeatures) Get() *FindDeviceById200ResponsePlanSpecsFeatures {
	return v.value
}

func (v *NullableFindDeviceById200ResponsePlanSpecsFeatures) Set(val *FindDeviceById200ResponsePlanSpecsFeatures) {
	v.value = val
	v.isSet = true
}

func (v NullableFindDeviceById200ResponsePlanSpecsFeatures) IsSet() bool {
	return v.isSet
}

func (v *NullableFindDeviceById200ResponsePlanSpecsFeatures) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFindDeviceById200ResponsePlanSpecsFeatures(val *FindDeviceById200ResponsePlanSpecsFeatures) *NullableFindDeviceById200ResponsePlanSpecsFeatures {
	return &NullableFindDeviceById200ResponsePlanSpecsFeatures{value: val, isSet: true}
}

func (v NullableFindDeviceById200ResponsePlanSpecsFeatures) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFindDeviceById200ResponsePlanSpecsFeatures) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
