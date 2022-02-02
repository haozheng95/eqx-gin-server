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

// FacilityList struct for FacilityList
type FacilityList struct {
	Facilities []Facility `json:"facilities,omitempty"`
}

// NewFacilityList instantiates a new FacilityList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFacilityList() *FacilityList {
	this := FacilityList{}
	return &this
}

// NewFacilityListWithDefaults instantiates a new FacilityList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFacilityListWithDefaults() *FacilityList {
	this := FacilityList{}
	return &this
}

// GetFacilities returns the Facilities field value if set, zero value otherwise.
func (o *FacilityList) GetFacilities() []Facility {
	if o == nil || o.Facilities == nil {
		var ret []Facility
		return ret
	}
	return o.Facilities
}

// GetFacilitiesOk returns a tuple with the Facilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FacilityList) GetFacilitiesOk() ([]Facility, bool) {
	if o == nil || o.Facilities == nil {
		return nil, false
	}
	return o.Facilities, true
}

// HasFacilities returns a boolean if a field has been set.
func (o *FacilityList) HasFacilities() bool {
	if o != nil && o.Facilities != nil {
		return true
	}

	return false
}

// SetFacilities gets a reference to the given []Facility and assigns it to the Facilities field.
func (o *FacilityList) SetFacilities(v []Facility) {
	o.Facilities = v
}

func (o FacilityList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Facilities != nil {
		toSerialize["facilities"] = o.Facilities
	}
	return json.Marshal(toSerialize)
}

type NullableFacilityList struct {
	value *FacilityList
	isSet bool
}

func (v NullableFacilityList) Get() *FacilityList {
	return v.value
}

func (v *NullableFacilityList) Set(val *FacilityList) {
	v.value = val
	v.isSet = true
}

func (v NullableFacilityList) IsSet() bool {
	return v.isSet
}

func (v *NullableFacilityList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFacilityList(val *FacilityList) *NullableFacilityList {
	return &NullableFacilityList{value: val, isSet: true}
}

func (v NullableFacilityList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFacilityList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


