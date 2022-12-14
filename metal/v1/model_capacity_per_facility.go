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

// CapacityPerFacility struct for CapacityPerFacility
type CapacityPerFacility struct {
	Baremetal0   *FindCapacityForFacility200ResponseCapacityAms1Baremetal0 `json:"baremetal_0,omitempty"`
	Baremetal1   *FindCapacityForFacility200ResponseCapacityAms1Baremetal0 `json:"baremetal_1,omitempty"`
	Baremetal2   *FindCapacityForFacility200ResponseCapacityAms1Baremetal0 `json:"baremetal_2,omitempty"`
	Baremetal2a  *FindCapacityForFacility200ResponseCapacityAms1Baremetal0 `json:"baremetal_2a,omitempty"`
	Baremetal2a2 *FindCapacityForFacility200ResponseCapacityAms1Baremetal0 `json:"baremetal_2a2,omitempty"`
	Baremetal3   *FindCapacityForFacility200ResponseCapacityAms1Baremetal0 `json:"baremetal_3,omitempty"`
	BaremetalS   *FindCapacityForFacility200ResponseCapacityAms1Baremetal0 `json:"baremetal_s,omitempty"`
	C2MediumX86  *FindCapacityForFacility200ResponseCapacityAms1Baremetal0 `json:"c2.medium.x86,omitempty"`
	M2XlargeX86  *FindCapacityForFacility200ResponseCapacityAms1Baremetal0 `json:"m2.xlarge.x86,omitempty"`
}

// NewCapacityPerFacility instantiates a new CapacityPerFacility object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapacityPerFacility() *CapacityPerFacility {
	this := CapacityPerFacility{}
	return &this
}

// NewCapacityPerFacilityWithDefaults instantiates a new CapacityPerFacility object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapacityPerFacilityWithDefaults() *CapacityPerFacility {
	this := CapacityPerFacility{}
	return &this
}

// GetBaremetal0 returns the Baremetal0 field value if set, zero value otherwise.
func (o *CapacityPerFacility) GetBaremetal0() FindCapacityForFacility200ResponseCapacityAms1Baremetal0 {
	if o == nil || o.Baremetal0 == nil {
		var ret FindCapacityForFacility200ResponseCapacityAms1Baremetal0
		return ret
	}
	return *o.Baremetal0
}

// GetBaremetal0Ok returns a tuple with the Baremetal0 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapacityPerFacility) GetBaremetal0Ok() (*FindCapacityForFacility200ResponseCapacityAms1Baremetal0, bool) {
	if o == nil || o.Baremetal0 == nil {
		return nil, false
	}
	return o.Baremetal0, true
}

// HasBaremetal0 returns a boolean if a field has been set.
func (o *CapacityPerFacility) HasBaremetal0() bool {
	if o != nil && o.Baremetal0 != nil {
		return true
	}

	return false
}

// SetBaremetal0 gets a reference to the given FindCapacityForFacility200ResponseCapacityAms1Baremetal0 and assigns it to the Baremetal0 field.
func (o *CapacityPerFacility) SetBaremetal0(v FindCapacityForFacility200ResponseCapacityAms1Baremetal0) {
	o.Baremetal0 = &v
}

// GetBaremetal1 returns the Baremetal1 field value if set, zero value otherwise.
func (o *CapacityPerFacility) GetBaremetal1() FindCapacityForFacility200ResponseCapacityAms1Baremetal0 {
	if o == nil || o.Baremetal1 == nil {
		var ret FindCapacityForFacility200ResponseCapacityAms1Baremetal0
		return ret
	}
	return *o.Baremetal1
}

// GetBaremetal1Ok returns a tuple with the Baremetal1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapacityPerFacility) GetBaremetal1Ok() (*FindCapacityForFacility200ResponseCapacityAms1Baremetal0, bool) {
	if o == nil || o.Baremetal1 == nil {
		return nil, false
	}
	return o.Baremetal1, true
}

// HasBaremetal1 returns a boolean if a field has been set.
func (o *CapacityPerFacility) HasBaremetal1() bool {
	if o != nil && o.Baremetal1 != nil {
		return true
	}

	return false
}

// SetBaremetal1 gets a reference to the given FindCapacityForFacility200ResponseCapacityAms1Baremetal0 and assigns it to the Baremetal1 field.
func (o *CapacityPerFacility) SetBaremetal1(v FindCapacityForFacility200ResponseCapacityAms1Baremetal0) {
	o.Baremetal1 = &v
}

// GetBaremetal2 returns the Baremetal2 field value if set, zero value otherwise.
func (o *CapacityPerFacility) GetBaremetal2() FindCapacityForFacility200ResponseCapacityAms1Baremetal0 {
	if o == nil || o.Baremetal2 == nil {
		var ret FindCapacityForFacility200ResponseCapacityAms1Baremetal0
		return ret
	}
	return *o.Baremetal2
}

// GetBaremetal2Ok returns a tuple with the Baremetal2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapacityPerFacility) GetBaremetal2Ok() (*FindCapacityForFacility200ResponseCapacityAms1Baremetal0, bool) {
	if o == nil || o.Baremetal2 == nil {
		return nil, false
	}
	return o.Baremetal2, true
}

// HasBaremetal2 returns a boolean if a field has been set.
func (o *CapacityPerFacility) HasBaremetal2() bool {
	if o != nil && o.Baremetal2 != nil {
		return true
	}

	return false
}

// SetBaremetal2 gets a reference to the given FindCapacityForFacility200ResponseCapacityAms1Baremetal0 and assigns it to the Baremetal2 field.
func (o *CapacityPerFacility) SetBaremetal2(v FindCapacityForFacility200ResponseCapacityAms1Baremetal0) {
	o.Baremetal2 = &v
}

// GetBaremetal2a returns the Baremetal2a field value if set, zero value otherwise.
func (o *CapacityPerFacility) GetBaremetal2a() FindCapacityForFacility200ResponseCapacityAms1Baremetal0 {
	if o == nil || o.Baremetal2a == nil {
		var ret FindCapacityForFacility200ResponseCapacityAms1Baremetal0
		return ret
	}
	return *o.Baremetal2a
}

// GetBaremetal2aOk returns a tuple with the Baremetal2a field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapacityPerFacility) GetBaremetal2aOk() (*FindCapacityForFacility200ResponseCapacityAms1Baremetal0, bool) {
	if o == nil || o.Baremetal2a == nil {
		return nil, false
	}
	return o.Baremetal2a, true
}

// HasBaremetal2a returns a boolean if a field has been set.
func (o *CapacityPerFacility) HasBaremetal2a() bool {
	if o != nil && o.Baremetal2a != nil {
		return true
	}

	return false
}

// SetBaremetal2a gets a reference to the given FindCapacityForFacility200ResponseCapacityAms1Baremetal0 and assigns it to the Baremetal2a field.
func (o *CapacityPerFacility) SetBaremetal2a(v FindCapacityForFacility200ResponseCapacityAms1Baremetal0) {
	o.Baremetal2a = &v
}

// GetBaremetal2a2 returns the Baremetal2a2 field value if set, zero value otherwise.
func (o *CapacityPerFacility) GetBaremetal2a2() FindCapacityForFacility200ResponseCapacityAms1Baremetal0 {
	if o == nil || o.Baremetal2a2 == nil {
		var ret FindCapacityForFacility200ResponseCapacityAms1Baremetal0
		return ret
	}
	return *o.Baremetal2a2
}

// GetBaremetal2a2Ok returns a tuple with the Baremetal2a2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapacityPerFacility) GetBaremetal2a2Ok() (*FindCapacityForFacility200ResponseCapacityAms1Baremetal0, bool) {
	if o == nil || o.Baremetal2a2 == nil {
		return nil, false
	}
	return o.Baremetal2a2, true
}

// HasBaremetal2a2 returns a boolean if a field has been set.
func (o *CapacityPerFacility) HasBaremetal2a2() bool {
	if o != nil && o.Baremetal2a2 != nil {
		return true
	}

	return false
}

// SetBaremetal2a2 gets a reference to the given FindCapacityForFacility200ResponseCapacityAms1Baremetal0 and assigns it to the Baremetal2a2 field.
func (o *CapacityPerFacility) SetBaremetal2a2(v FindCapacityForFacility200ResponseCapacityAms1Baremetal0) {
	o.Baremetal2a2 = &v
}

// GetBaremetal3 returns the Baremetal3 field value if set, zero value otherwise.
func (o *CapacityPerFacility) GetBaremetal3() FindCapacityForFacility200ResponseCapacityAms1Baremetal0 {
	if o == nil || o.Baremetal3 == nil {
		var ret FindCapacityForFacility200ResponseCapacityAms1Baremetal0
		return ret
	}
	return *o.Baremetal3
}

// GetBaremetal3Ok returns a tuple with the Baremetal3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapacityPerFacility) GetBaremetal3Ok() (*FindCapacityForFacility200ResponseCapacityAms1Baremetal0, bool) {
	if o == nil || o.Baremetal3 == nil {
		return nil, false
	}
	return o.Baremetal3, true
}

// HasBaremetal3 returns a boolean if a field has been set.
func (o *CapacityPerFacility) HasBaremetal3() bool {
	if o != nil && o.Baremetal3 != nil {
		return true
	}

	return false
}

// SetBaremetal3 gets a reference to the given FindCapacityForFacility200ResponseCapacityAms1Baremetal0 and assigns it to the Baremetal3 field.
func (o *CapacityPerFacility) SetBaremetal3(v FindCapacityForFacility200ResponseCapacityAms1Baremetal0) {
	o.Baremetal3 = &v
}

// GetBaremetalS returns the BaremetalS field value if set, zero value otherwise.
func (o *CapacityPerFacility) GetBaremetalS() FindCapacityForFacility200ResponseCapacityAms1Baremetal0 {
	if o == nil || o.BaremetalS == nil {
		var ret FindCapacityForFacility200ResponseCapacityAms1Baremetal0
		return ret
	}
	return *o.BaremetalS
}

// GetBaremetalSOk returns a tuple with the BaremetalS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapacityPerFacility) GetBaremetalSOk() (*FindCapacityForFacility200ResponseCapacityAms1Baremetal0, bool) {
	if o == nil || o.BaremetalS == nil {
		return nil, false
	}
	return o.BaremetalS, true
}

// HasBaremetalS returns a boolean if a field has been set.
func (o *CapacityPerFacility) HasBaremetalS() bool {
	if o != nil && o.BaremetalS != nil {
		return true
	}

	return false
}

// SetBaremetalS gets a reference to the given FindCapacityForFacility200ResponseCapacityAms1Baremetal0 and assigns it to the BaremetalS field.
func (o *CapacityPerFacility) SetBaremetalS(v FindCapacityForFacility200ResponseCapacityAms1Baremetal0) {
	o.BaremetalS = &v
}

// GetC2MediumX86 returns the C2MediumX86 field value if set, zero value otherwise.
func (o *CapacityPerFacility) GetC2MediumX86() FindCapacityForFacility200ResponseCapacityAms1Baremetal0 {
	if o == nil || o.C2MediumX86 == nil {
		var ret FindCapacityForFacility200ResponseCapacityAms1Baremetal0
		return ret
	}
	return *o.C2MediumX86
}

// GetC2MediumX86Ok returns a tuple with the C2MediumX86 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapacityPerFacility) GetC2MediumX86Ok() (*FindCapacityForFacility200ResponseCapacityAms1Baremetal0, bool) {
	if o == nil || o.C2MediumX86 == nil {
		return nil, false
	}
	return o.C2MediumX86, true
}

// HasC2MediumX86 returns a boolean if a field has been set.
func (o *CapacityPerFacility) HasC2MediumX86() bool {
	if o != nil && o.C2MediumX86 != nil {
		return true
	}

	return false
}

// SetC2MediumX86 gets a reference to the given FindCapacityForFacility200ResponseCapacityAms1Baremetal0 and assigns it to the C2MediumX86 field.
func (o *CapacityPerFacility) SetC2MediumX86(v FindCapacityForFacility200ResponseCapacityAms1Baremetal0) {
	o.C2MediumX86 = &v
}

// GetM2XlargeX86 returns the M2XlargeX86 field value if set, zero value otherwise.
func (o *CapacityPerFacility) GetM2XlargeX86() FindCapacityForFacility200ResponseCapacityAms1Baremetal0 {
	if o == nil || o.M2XlargeX86 == nil {
		var ret FindCapacityForFacility200ResponseCapacityAms1Baremetal0
		return ret
	}
	return *o.M2XlargeX86
}

// GetM2XlargeX86Ok returns a tuple with the M2XlargeX86 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapacityPerFacility) GetM2XlargeX86Ok() (*FindCapacityForFacility200ResponseCapacityAms1Baremetal0, bool) {
	if o == nil || o.M2XlargeX86 == nil {
		return nil, false
	}
	return o.M2XlargeX86, true
}

// HasM2XlargeX86 returns a boolean if a field has been set.
func (o *CapacityPerFacility) HasM2XlargeX86() bool {
	if o != nil && o.M2XlargeX86 != nil {
		return true
	}

	return false
}

// SetM2XlargeX86 gets a reference to the given FindCapacityForFacility200ResponseCapacityAms1Baremetal0 and assigns it to the M2XlargeX86 field.
func (o *CapacityPerFacility) SetM2XlargeX86(v FindCapacityForFacility200ResponseCapacityAms1Baremetal0) {
	o.M2XlargeX86 = &v
}

func (o CapacityPerFacility) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Baremetal0 != nil {
		toSerialize["baremetal_0"] = o.Baremetal0
	}
	if o.Baremetal1 != nil {
		toSerialize["baremetal_1"] = o.Baremetal1
	}
	if o.Baremetal2 != nil {
		toSerialize["baremetal_2"] = o.Baremetal2
	}
	if o.Baremetal2a != nil {
		toSerialize["baremetal_2a"] = o.Baremetal2a
	}
	if o.Baremetal2a2 != nil {
		toSerialize["baremetal_2a2"] = o.Baremetal2a2
	}
	if o.Baremetal3 != nil {
		toSerialize["baremetal_3"] = o.Baremetal3
	}
	if o.BaremetalS != nil {
		toSerialize["baremetal_s"] = o.BaremetalS
	}
	if o.C2MediumX86 != nil {
		toSerialize["c2.medium.x86"] = o.C2MediumX86
	}
	if o.M2XlargeX86 != nil {
		toSerialize["m2.xlarge.x86"] = o.M2XlargeX86
	}
	return json.Marshal(toSerialize)
}

type NullableCapacityPerFacility struct {
	value *CapacityPerFacility
	isSet bool
}

func (v NullableCapacityPerFacility) Get() *CapacityPerFacility {
	return v.value
}

func (v *NullableCapacityPerFacility) Set(val *CapacityPerFacility) {
	v.value = val
	v.isSet = true
}

func (v NullableCapacityPerFacility) IsSet() bool {
	return v.isSet
}

func (v *NullableCapacityPerFacility) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapacityPerFacility(val *CapacityPerFacility) *NullableCapacityPerFacility {
	return &NullableCapacityPerFacility{value: val, isSet: true}
}

func (v NullableCapacityPerFacility) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapacityPerFacility) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
