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

// FindDeviceById200ResponseNetworkPortsInnerData struct for FindDeviceById200ResponseNetworkPortsInnerData
type FindDeviceById200ResponseNetworkPortsInnerData struct {
	// MAC address is set for NetworkPort ports
	Mac *string `json:"mac,omitempty"`
	// Bonded is true for NetworkPort ports in a bond and NetworkBondPort ports that are active
	Bonded *bool `json:"bonded,omitempty"`
}

// NewFindDeviceById200ResponseNetworkPortsInnerData instantiates a new FindDeviceById200ResponseNetworkPortsInnerData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFindDeviceById200ResponseNetworkPortsInnerData() *FindDeviceById200ResponseNetworkPortsInnerData {
	this := FindDeviceById200ResponseNetworkPortsInnerData{}
	return &this
}

// NewFindDeviceById200ResponseNetworkPortsInnerDataWithDefaults instantiates a new FindDeviceById200ResponseNetworkPortsInnerData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFindDeviceById200ResponseNetworkPortsInnerDataWithDefaults() *FindDeviceById200ResponseNetworkPortsInnerData {
	this := FindDeviceById200ResponseNetworkPortsInnerData{}
	return &this
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *FindDeviceById200ResponseNetworkPortsInnerData) GetMac() string {
	if o == nil || o.Mac == nil {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindDeviceById200ResponseNetworkPortsInnerData) GetMacOk() (*string, bool) {
	if o == nil || o.Mac == nil {
		return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *FindDeviceById200ResponseNetworkPortsInnerData) HasMac() bool {
	if o != nil && o.Mac != nil {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *FindDeviceById200ResponseNetworkPortsInnerData) SetMac(v string) {
	o.Mac = &v
}

// GetBonded returns the Bonded field value if set, zero value otherwise.
func (o *FindDeviceById200ResponseNetworkPortsInnerData) GetBonded() bool {
	if o == nil || o.Bonded == nil {
		var ret bool
		return ret
	}
	return *o.Bonded
}

// GetBondedOk returns a tuple with the Bonded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindDeviceById200ResponseNetworkPortsInnerData) GetBondedOk() (*bool, bool) {
	if o == nil || o.Bonded == nil {
		return nil, false
	}
	return o.Bonded, true
}

// HasBonded returns a boolean if a field has been set.
func (o *FindDeviceById200ResponseNetworkPortsInnerData) HasBonded() bool {
	if o != nil && o.Bonded != nil {
		return true
	}

	return false
}

// SetBonded gets a reference to the given bool and assigns it to the Bonded field.
func (o *FindDeviceById200ResponseNetworkPortsInnerData) SetBonded(v bool) {
	o.Bonded = &v
}

func (o FindDeviceById200ResponseNetworkPortsInnerData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Mac != nil {
		toSerialize["mac"] = o.Mac
	}
	if o.Bonded != nil {
		toSerialize["bonded"] = o.Bonded
	}
	return json.Marshal(toSerialize)
}

type NullableFindDeviceById200ResponseNetworkPortsInnerData struct {
	value *FindDeviceById200ResponseNetworkPortsInnerData
	isSet bool
}

func (v NullableFindDeviceById200ResponseNetworkPortsInnerData) Get() *FindDeviceById200ResponseNetworkPortsInnerData {
	return v.value
}

func (v *NullableFindDeviceById200ResponseNetworkPortsInnerData) Set(val *FindDeviceById200ResponseNetworkPortsInnerData) {
	v.value = val
	v.isSet = true
}

func (v NullableFindDeviceById200ResponseNetworkPortsInnerData) IsSet() bool {
	return v.isSet
}

func (v *NullableFindDeviceById200ResponseNetworkPortsInnerData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFindDeviceById200ResponseNetworkPortsInnerData(val *FindDeviceById200ResponseNetworkPortsInnerData) *NullableFindDeviceById200ResponseNetworkPortsInnerData {
	return &NullableFindDeviceById200ResponseNetworkPortsInnerData{value: val, isSet: true}
}

func (v NullableFindDeviceById200ResponseNetworkPortsInnerData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFindDeviceById200ResponseNetworkPortsInnerData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}