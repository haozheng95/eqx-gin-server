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

// BgpSessionNeighbors struct for BgpSessionNeighbors
type BgpSessionNeighbors struct {
	// A list of BGP session neighbor data
	BgpNeighbors []GetBgpNeighborData200ResponseBgpNeighborsInner `json:"bgp_neighbors,omitempty"`
}

// NewBgpSessionNeighbors instantiates a new BgpSessionNeighbors object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBgpSessionNeighbors() *BgpSessionNeighbors {
	this := BgpSessionNeighbors{}
	return &this
}

// NewBgpSessionNeighborsWithDefaults instantiates a new BgpSessionNeighbors object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBgpSessionNeighborsWithDefaults() *BgpSessionNeighbors {
	this := BgpSessionNeighbors{}
	return &this
}

// GetBgpNeighbors returns the BgpNeighbors field value if set, zero value otherwise.
func (o *BgpSessionNeighbors) GetBgpNeighbors() []GetBgpNeighborData200ResponseBgpNeighborsInner {
	if o == nil || o.BgpNeighbors == nil {
		var ret []GetBgpNeighborData200ResponseBgpNeighborsInner
		return ret
	}
	return o.BgpNeighbors
}

// GetBgpNeighborsOk returns a tuple with the BgpNeighbors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpSessionNeighbors) GetBgpNeighborsOk() ([]GetBgpNeighborData200ResponseBgpNeighborsInner, bool) {
	if o == nil || o.BgpNeighbors == nil {
		return nil, false
	}
	return o.BgpNeighbors, true
}

// HasBgpNeighbors returns a boolean if a field has been set.
func (o *BgpSessionNeighbors) HasBgpNeighbors() bool {
	if o != nil && o.BgpNeighbors != nil {
		return true
	}

	return false
}

// SetBgpNeighbors gets a reference to the given []GetBgpNeighborData200ResponseBgpNeighborsInner and assigns it to the BgpNeighbors field.
func (o *BgpSessionNeighbors) SetBgpNeighbors(v []GetBgpNeighborData200ResponseBgpNeighborsInner) {
	o.BgpNeighbors = v
}

func (o BgpSessionNeighbors) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BgpNeighbors != nil {
		toSerialize["bgp_neighbors"] = o.BgpNeighbors
	}
	return json.Marshal(toSerialize)
}

type NullableBgpSessionNeighbors struct {
	value *BgpSessionNeighbors
	isSet bool
}

func (v NullableBgpSessionNeighbors) Get() *BgpSessionNeighbors {
	return v.value
}

func (v *NullableBgpSessionNeighbors) Set(val *BgpSessionNeighbors) {
	v.value = val
	v.isSet = true
}

func (v NullableBgpSessionNeighbors) IsSet() bool {
	return v.isSet
}

func (v *NullableBgpSessionNeighbors) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBgpSessionNeighbors(val *BgpSessionNeighbors) *NullableBgpSessionNeighbors {
	return &NullableBgpSessionNeighbors{value: val, isSet: true}
}

func (v NullableBgpSessionNeighbors) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBgpSessionNeighbors) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
