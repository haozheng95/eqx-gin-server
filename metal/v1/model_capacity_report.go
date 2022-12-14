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

// CapacityReport struct for CapacityReport
type CapacityReport struct {
	Ams1 *FindCapacityForFacility200ResponseCapacityAms1 `json:"ams1,omitempty"`
	Atl1 *FindCapacityForFacility200ResponseCapacityAtl1 `json:"atl1,omitempty"`
	Dfw1 *FindCapacityForFacility200ResponseCapacityAtl1 `json:"dfw1,omitempty"`
	Ewr1 *FindCapacityForFacility200ResponseCapacityAms1 `json:"ewr1,omitempty"`
	Fra1 *FindCapacityForFacility200ResponseCapacityAtl1 `json:"fra1,omitempty"`
	Iad1 *FindCapacityForFacility200ResponseCapacityAtl1 `json:"iad1,omitempty"`
	Lax1 *FindCapacityForFacility200ResponseCapacityAtl1 `json:"lax1,omitempty"`
	Nrt1 *FindCapacityForFacility200ResponseCapacityAms1 `json:"nrt1,omitempty"`
	Ord1 *FindCapacityForFacility200ResponseCapacityAtl1 `json:"ord1,omitempty"`
	Sea1 *FindCapacityForFacility200ResponseCapacityAtl1 `json:"sea1,omitempty"`
	Sin1 *FindCapacityForFacility200ResponseCapacityAtl1 `json:"sin1,omitempty"`
	Sjc1 *FindCapacityForFacility200ResponseCapacityAms1 `json:"sjc1,omitempty"`
	Syd1 *FindCapacityForFacility200ResponseCapacityAtl1 `json:"syd1,omitempty"`
	Yyz1 *FindCapacityForFacility200ResponseCapacityAtl1 `json:"yyz1,omitempty"`
}

// NewCapacityReport instantiates a new CapacityReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapacityReport() *CapacityReport {
	this := CapacityReport{}
	return &this
}

// NewCapacityReportWithDefaults instantiates a new CapacityReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapacityReportWithDefaults() *CapacityReport {
	this := CapacityReport{}
	return &this
}

// GetAms1 returns the Ams1 field value if set, zero value otherwise.
func (o *CapacityReport) GetAms1() FindCapacityForFacility200ResponseCapacityAms1 {
	if o == nil || o.Ams1 == nil {
		var ret FindCapacityForFacility200ResponseCapacityAms1
		return ret
	}
	return *o.Ams1
}

// GetAms1Ok returns a tuple with the Ams1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapacityReport) GetAms1Ok() (*FindCapacityForFacility200ResponseCapacityAms1, bool) {
	if o == nil || o.Ams1 == nil {
		return nil, false
	}
	return o.Ams1, true
}

// HasAms1 returns a boolean if a field has been set.
func (o *CapacityReport) HasAms1() bool {
	if o != nil && o.Ams1 != nil {
		return true
	}

	return false
}

// SetAms1 gets a reference to the given FindCapacityForFacility200ResponseCapacityAms1 and assigns it to the Ams1 field.
func (o *CapacityReport) SetAms1(v FindCapacityForFacility200ResponseCapacityAms1) {
	o.Ams1 = &v
}

// GetAtl1 returns the Atl1 field value if set, zero value otherwise.
func (o *CapacityReport) GetAtl1() FindCapacityForFacility200ResponseCapacityAtl1 {
	if o == nil || o.Atl1 == nil {
		var ret FindCapacityForFacility200ResponseCapacityAtl1
		return ret
	}
	return *o.Atl1
}

// GetAtl1Ok returns a tuple with the Atl1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapacityReport) GetAtl1Ok() (*FindCapacityForFacility200ResponseCapacityAtl1, bool) {
	if o == nil || o.Atl1 == nil {
		return nil, false
	}
	return o.Atl1, true
}

// HasAtl1 returns a boolean if a field has been set.
func (o *CapacityReport) HasAtl1() bool {
	if o != nil && o.Atl1 != nil {
		return true
	}

	return false
}

// SetAtl1 gets a reference to the given FindCapacityForFacility200ResponseCapacityAtl1 and assigns it to the Atl1 field.
func (o *CapacityReport) SetAtl1(v FindCapacityForFacility200ResponseCapacityAtl1) {
	o.Atl1 = &v
}

// GetDfw1 returns the Dfw1 field value if set, zero value otherwise.
func (o *CapacityReport) GetDfw1() FindCapacityForFacility200ResponseCapacityAtl1 {
	if o == nil || o.Dfw1 == nil {
		var ret FindCapacityForFacility200ResponseCapacityAtl1
		return ret
	}
	return *o.Dfw1
}

// GetDfw1Ok returns a tuple with the Dfw1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapacityReport) GetDfw1Ok() (*FindCapacityForFacility200ResponseCapacityAtl1, bool) {
	if o == nil || o.Dfw1 == nil {
		return nil, false
	}
	return o.Dfw1, true
}

// HasDfw1 returns a boolean if a field has been set.
func (o *CapacityReport) HasDfw1() bool {
	if o != nil && o.Dfw1 != nil {
		return true
	}

	return false
}

// SetDfw1 gets a reference to the given FindCapacityForFacility200ResponseCapacityAtl1 and assigns it to the Dfw1 field.
func (o *CapacityReport) SetDfw1(v FindCapacityForFacility200ResponseCapacityAtl1) {
	o.Dfw1 = &v
}

// GetEwr1 returns the Ewr1 field value if set, zero value otherwise.
func (o *CapacityReport) GetEwr1() FindCapacityForFacility200ResponseCapacityAms1 {
	if o == nil || o.Ewr1 == nil {
		var ret FindCapacityForFacility200ResponseCapacityAms1
		return ret
	}
	return *o.Ewr1
}

// GetEwr1Ok returns a tuple with the Ewr1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapacityReport) GetEwr1Ok() (*FindCapacityForFacility200ResponseCapacityAms1, bool) {
	if o == nil || o.Ewr1 == nil {
		return nil, false
	}
	return o.Ewr1, true
}

// HasEwr1 returns a boolean if a field has been set.
func (o *CapacityReport) HasEwr1() bool {
	if o != nil && o.Ewr1 != nil {
		return true
	}

	return false
}

// SetEwr1 gets a reference to the given FindCapacityForFacility200ResponseCapacityAms1 and assigns it to the Ewr1 field.
func (o *CapacityReport) SetEwr1(v FindCapacityForFacility200ResponseCapacityAms1) {
	o.Ewr1 = &v
}

// GetFra1 returns the Fra1 field value if set, zero value otherwise.
func (o *CapacityReport) GetFra1() FindCapacityForFacility200ResponseCapacityAtl1 {
	if o == nil || o.Fra1 == nil {
		var ret FindCapacityForFacility200ResponseCapacityAtl1
		return ret
	}
	return *o.Fra1
}

// GetFra1Ok returns a tuple with the Fra1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapacityReport) GetFra1Ok() (*FindCapacityForFacility200ResponseCapacityAtl1, bool) {
	if o == nil || o.Fra1 == nil {
		return nil, false
	}
	return o.Fra1, true
}

// HasFra1 returns a boolean if a field has been set.
func (o *CapacityReport) HasFra1() bool {
	if o != nil && o.Fra1 != nil {
		return true
	}

	return false
}

// SetFra1 gets a reference to the given FindCapacityForFacility200ResponseCapacityAtl1 and assigns it to the Fra1 field.
func (o *CapacityReport) SetFra1(v FindCapacityForFacility200ResponseCapacityAtl1) {
	o.Fra1 = &v
}

// GetIad1 returns the Iad1 field value if set, zero value otherwise.
func (o *CapacityReport) GetIad1() FindCapacityForFacility200ResponseCapacityAtl1 {
	if o == nil || o.Iad1 == nil {
		var ret FindCapacityForFacility200ResponseCapacityAtl1
		return ret
	}
	return *o.Iad1
}

// GetIad1Ok returns a tuple with the Iad1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapacityReport) GetIad1Ok() (*FindCapacityForFacility200ResponseCapacityAtl1, bool) {
	if o == nil || o.Iad1 == nil {
		return nil, false
	}
	return o.Iad1, true
}

// HasIad1 returns a boolean if a field has been set.
func (o *CapacityReport) HasIad1() bool {
	if o != nil && o.Iad1 != nil {
		return true
	}

	return false
}

// SetIad1 gets a reference to the given FindCapacityForFacility200ResponseCapacityAtl1 and assigns it to the Iad1 field.
func (o *CapacityReport) SetIad1(v FindCapacityForFacility200ResponseCapacityAtl1) {
	o.Iad1 = &v
}

// GetLax1 returns the Lax1 field value if set, zero value otherwise.
func (o *CapacityReport) GetLax1() FindCapacityForFacility200ResponseCapacityAtl1 {
	if o == nil || o.Lax1 == nil {
		var ret FindCapacityForFacility200ResponseCapacityAtl1
		return ret
	}
	return *o.Lax1
}

// GetLax1Ok returns a tuple with the Lax1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapacityReport) GetLax1Ok() (*FindCapacityForFacility200ResponseCapacityAtl1, bool) {
	if o == nil || o.Lax1 == nil {
		return nil, false
	}
	return o.Lax1, true
}

// HasLax1 returns a boolean if a field has been set.
func (o *CapacityReport) HasLax1() bool {
	if o != nil && o.Lax1 != nil {
		return true
	}

	return false
}

// SetLax1 gets a reference to the given FindCapacityForFacility200ResponseCapacityAtl1 and assigns it to the Lax1 field.
func (o *CapacityReport) SetLax1(v FindCapacityForFacility200ResponseCapacityAtl1) {
	o.Lax1 = &v
}

// GetNrt1 returns the Nrt1 field value if set, zero value otherwise.
func (o *CapacityReport) GetNrt1() FindCapacityForFacility200ResponseCapacityAms1 {
	if o == nil || o.Nrt1 == nil {
		var ret FindCapacityForFacility200ResponseCapacityAms1
		return ret
	}
	return *o.Nrt1
}

// GetNrt1Ok returns a tuple with the Nrt1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapacityReport) GetNrt1Ok() (*FindCapacityForFacility200ResponseCapacityAms1, bool) {
	if o == nil || o.Nrt1 == nil {
		return nil, false
	}
	return o.Nrt1, true
}

// HasNrt1 returns a boolean if a field has been set.
func (o *CapacityReport) HasNrt1() bool {
	if o != nil && o.Nrt1 != nil {
		return true
	}

	return false
}

// SetNrt1 gets a reference to the given FindCapacityForFacility200ResponseCapacityAms1 and assigns it to the Nrt1 field.
func (o *CapacityReport) SetNrt1(v FindCapacityForFacility200ResponseCapacityAms1) {
	o.Nrt1 = &v
}

// GetOrd1 returns the Ord1 field value if set, zero value otherwise.
func (o *CapacityReport) GetOrd1() FindCapacityForFacility200ResponseCapacityAtl1 {
	if o == nil || o.Ord1 == nil {
		var ret FindCapacityForFacility200ResponseCapacityAtl1
		return ret
	}
	return *o.Ord1
}

// GetOrd1Ok returns a tuple with the Ord1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapacityReport) GetOrd1Ok() (*FindCapacityForFacility200ResponseCapacityAtl1, bool) {
	if o == nil || o.Ord1 == nil {
		return nil, false
	}
	return o.Ord1, true
}

// HasOrd1 returns a boolean if a field has been set.
func (o *CapacityReport) HasOrd1() bool {
	if o != nil && o.Ord1 != nil {
		return true
	}

	return false
}

// SetOrd1 gets a reference to the given FindCapacityForFacility200ResponseCapacityAtl1 and assigns it to the Ord1 field.
func (o *CapacityReport) SetOrd1(v FindCapacityForFacility200ResponseCapacityAtl1) {
	o.Ord1 = &v
}

// GetSea1 returns the Sea1 field value if set, zero value otherwise.
func (o *CapacityReport) GetSea1() FindCapacityForFacility200ResponseCapacityAtl1 {
	if o == nil || o.Sea1 == nil {
		var ret FindCapacityForFacility200ResponseCapacityAtl1
		return ret
	}
	return *o.Sea1
}

// GetSea1Ok returns a tuple with the Sea1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapacityReport) GetSea1Ok() (*FindCapacityForFacility200ResponseCapacityAtl1, bool) {
	if o == nil || o.Sea1 == nil {
		return nil, false
	}
	return o.Sea1, true
}

// HasSea1 returns a boolean if a field has been set.
func (o *CapacityReport) HasSea1() bool {
	if o != nil && o.Sea1 != nil {
		return true
	}

	return false
}

// SetSea1 gets a reference to the given FindCapacityForFacility200ResponseCapacityAtl1 and assigns it to the Sea1 field.
func (o *CapacityReport) SetSea1(v FindCapacityForFacility200ResponseCapacityAtl1) {
	o.Sea1 = &v
}

// GetSin1 returns the Sin1 field value if set, zero value otherwise.
func (o *CapacityReport) GetSin1() FindCapacityForFacility200ResponseCapacityAtl1 {
	if o == nil || o.Sin1 == nil {
		var ret FindCapacityForFacility200ResponseCapacityAtl1
		return ret
	}
	return *o.Sin1
}

// GetSin1Ok returns a tuple with the Sin1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapacityReport) GetSin1Ok() (*FindCapacityForFacility200ResponseCapacityAtl1, bool) {
	if o == nil || o.Sin1 == nil {
		return nil, false
	}
	return o.Sin1, true
}

// HasSin1 returns a boolean if a field has been set.
func (o *CapacityReport) HasSin1() bool {
	if o != nil && o.Sin1 != nil {
		return true
	}

	return false
}

// SetSin1 gets a reference to the given FindCapacityForFacility200ResponseCapacityAtl1 and assigns it to the Sin1 field.
func (o *CapacityReport) SetSin1(v FindCapacityForFacility200ResponseCapacityAtl1) {
	o.Sin1 = &v
}

// GetSjc1 returns the Sjc1 field value if set, zero value otherwise.
func (o *CapacityReport) GetSjc1() FindCapacityForFacility200ResponseCapacityAms1 {
	if o == nil || o.Sjc1 == nil {
		var ret FindCapacityForFacility200ResponseCapacityAms1
		return ret
	}
	return *o.Sjc1
}

// GetSjc1Ok returns a tuple with the Sjc1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapacityReport) GetSjc1Ok() (*FindCapacityForFacility200ResponseCapacityAms1, bool) {
	if o == nil || o.Sjc1 == nil {
		return nil, false
	}
	return o.Sjc1, true
}

// HasSjc1 returns a boolean if a field has been set.
func (o *CapacityReport) HasSjc1() bool {
	if o != nil && o.Sjc1 != nil {
		return true
	}

	return false
}

// SetSjc1 gets a reference to the given FindCapacityForFacility200ResponseCapacityAms1 and assigns it to the Sjc1 field.
func (o *CapacityReport) SetSjc1(v FindCapacityForFacility200ResponseCapacityAms1) {
	o.Sjc1 = &v
}

// GetSyd1 returns the Syd1 field value if set, zero value otherwise.
func (o *CapacityReport) GetSyd1() FindCapacityForFacility200ResponseCapacityAtl1 {
	if o == nil || o.Syd1 == nil {
		var ret FindCapacityForFacility200ResponseCapacityAtl1
		return ret
	}
	return *o.Syd1
}

// GetSyd1Ok returns a tuple with the Syd1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapacityReport) GetSyd1Ok() (*FindCapacityForFacility200ResponseCapacityAtl1, bool) {
	if o == nil || o.Syd1 == nil {
		return nil, false
	}
	return o.Syd1, true
}

// HasSyd1 returns a boolean if a field has been set.
func (o *CapacityReport) HasSyd1() bool {
	if o != nil && o.Syd1 != nil {
		return true
	}

	return false
}

// SetSyd1 gets a reference to the given FindCapacityForFacility200ResponseCapacityAtl1 and assigns it to the Syd1 field.
func (o *CapacityReport) SetSyd1(v FindCapacityForFacility200ResponseCapacityAtl1) {
	o.Syd1 = &v
}

// GetYyz1 returns the Yyz1 field value if set, zero value otherwise.
func (o *CapacityReport) GetYyz1() FindCapacityForFacility200ResponseCapacityAtl1 {
	if o == nil || o.Yyz1 == nil {
		var ret FindCapacityForFacility200ResponseCapacityAtl1
		return ret
	}
	return *o.Yyz1
}

// GetYyz1Ok returns a tuple with the Yyz1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapacityReport) GetYyz1Ok() (*FindCapacityForFacility200ResponseCapacityAtl1, bool) {
	if o == nil || o.Yyz1 == nil {
		return nil, false
	}
	return o.Yyz1, true
}

// HasYyz1 returns a boolean if a field has been set.
func (o *CapacityReport) HasYyz1() bool {
	if o != nil && o.Yyz1 != nil {
		return true
	}

	return false
}

// SetYyz1 gets a reference to the given FindCapacityForFacility200ResponseCapacityAtl1 and assigns it to the Yyz1 field.
func (o *CapacityReport) SetYyz1(v FindCapacityForFacility200ResponseCapacityAtl1) {
	o.Yyz1 = &v
}

func (o CapacityReport) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Ams1 != nil {
		toSerialize["ams1"] = o.Ams1
	}
	if o.Atl1 != nil {
		toSerialize["atl1"] = o.Atl1
	}
	if o.Dfw1 != nil {
		toSerialize["dfw1"] = o.Dfw1
	}
	if o.Ewr1 != nil {
		toSerialize["ewr1"] = o.Ewr1
	}
	if o.Fra1 != nil {
		toSerialize["fra1"] = o.Fra1
	}
	if o.Iad1 != nil {
		toSerialize["iad1"] = o.Iad1
	}
	if o.Lax1 != nil {
		toSerialize["lax1"] = o.Lax1
	}
	if o.Nrt1 != nil {
		toSerialize["nrt1"] = o.Nrt1
	}
	if o.Ord1 != nil {
		toSerialize["ord1"] = o.Ord1
	}
	if o.Sea1 != nil {
		toSerialize["sea1"] = o.Sea1
	}
	if o.Sin1 != nil {
		toSerialize["sin1"] = o.Sin1
	}
	if o.Sjc1 != nil {
		toSerialize["sjc1"] = o.Sjc1
	}
	if o.Syd1 != nil {
		toSerialize["syd1"] = o.Syd1
	}
	if o.Yyz1 != nil {
		toSerialize["yyz1"] = o.Yyz1
	}
	return json.Marshal(toSerialize)
}

type NullableCapacityReport struct {
	value *CapacityReport
	isSet bool
}

func (v NullableCapacityReport) Get() *CapacityReport {
	return v.value
}

func (v *NullableCapacityReport) Set(val *CapacityReport) {
	v.value = val
	v.isSet = true
}

func (v NullableCapacityReport) IsSet() bool {
	return v.isSet
}

func (v *NullableCapacityReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapacityReport(val *CapacityReport) *NullableCapacityReport {
	return &NullableCapacityReport{value: val, isSet: true}
}

func (v NullableCapacityReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapacityReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
