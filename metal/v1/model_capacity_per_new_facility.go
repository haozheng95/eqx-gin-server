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

// CapacityPerNewFacility struct for CapacityPerNewFacility
type CapacityPerNewFacility struct {
	Baremetal1e *FindCapacityForFacility200ResponseCapacityAms1Baremetal0 `json:"baremetal_1e,omitempty"`
}

// NewCapacityPerNewFacility instantiates a new CapacityPerNewFacility object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapacityPerNewFacility() *CapacityPerNewFacility {
	this := CapacityPerNewFacility{}
	return &this
}

// NewCapacityPerNewFacilityWithDefaults instantiates a new CapacityPerNewFacility object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapacityPerNewFacilityWithDefaults() *CapacityPerNewFacility {
	this := CapacityPerNewFacility{}
	return &this
}

// GetBaremetal1e returns the Baremetal1e field value if set, zero value otherwise.
func (o *CapacityPerNewFacility) GetBaremetal1e() FindCapacityForFacility200ResponseCapacityAms1Baremetal0 {
	if o == nil || o.Baremetal1e == nil {
		var ret FindCapacityForFacility200ResponseCapacityAms1Baremetal0
		return ret
	}
	return *o.Baremetal1e
}

// GetBaremetal1eOk returns a tuple with the Baremetal1e field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapacityPerNewFacility) GetBaremetal1eOk() (*FindCapacityForFacility200ResponseCapacityAms1Baremetal0, bool) {
	if o == nil || o.Baremetal1e == nil {
		return nil, false
	}
	return o.Baremetal1e, true
}

// HasBaremetal1e returns a boolean if a field has been set.
func (o *CapacityPerNewFacility) HasBaremetal1e() bool {
	if o != nil && o.Baremetal1e != nil {
		return true
	}

	return false
}

// SetBaremetal1e gets a reference to the given FindCapacityForFacility200ResponseCapacityAms1Baremetal0 and assigns it to the Baremetal1e field.
func (o *CapacityPerNewFacility) SetBaremetal1e(v FindCapacityForFacility200ResponseCapacityAms1Baremetal0) {
	o.Baremetal1e = &v
}

func (o CapacityPerNewFacility) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Baremetal1e != nil {
		toSerialize["baremetal_1e"] = o.Baremetal1e
	}
	return json.Marshal(toSerialize)
}

type NullableCapacityPerNewFacility struct {
	value *CapacityPerNewFacility
	isSet bool
}

func (v NullableCapacityPerNewFacility) Get() *CapacityPerNewFacility {
	return v.value
}

func (v *NullableCapacityPerNewFacility) Set(val *CapacityPerNewFacility) {
	v.value = val
	v.isSet = true
}

func (v NullableCapacityPerNewFacility) IsSet() bool {
	return v.isSet
}

func (v *NullableCapacityPerNewFacility) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapacityPerNewFacility(val *CapacityPerNewFacility) *NullableCapacityPerNewFacility {
	return &NullableCapacityPerNewFacility{value: val, isSet: true}
}

func (v NullableCapacityPerNewFacility) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapacityPerNewFacility) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
