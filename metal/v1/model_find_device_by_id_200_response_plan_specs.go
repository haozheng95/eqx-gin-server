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

// FindDeviceById200ResponsePlanSpecs struct for FindDeviceById200ResponsePlanSpecs
type FindDeviceById200ResponsePlanSpecs struct {
	Cpus     []FindDeviceById200ResponsePlanSpecsCpusInner   `json:"cpus,omitempty"`
	Drives   []FindDeviceById200ResponsePlanSpecsDrivesInner `json:"drives,omitempty"`
	Nics     []FindDeviceById200ResponsePlanSpecsNicsInner   `json:"nics,omitempty"`
	Features *FindDeviceById200ResponsePlanSpecsFeatures     `json:"features,omitempty"`
}

// NewFindDeviceById200ResponsePlanSpecs instantiates a new FindDeviceById200ResponsePlanSpecs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFindDeviceById200ResponsePlanSpecs() *FindDeviceById200ResponsePlanSpecs {
	this := FindDeviceById200ResponsePlanSpecs{}
	return &this
}

// NewFindDeviceById200ResponsePlanSpecsWithDefaults instantiates a new FindDeviceById200ResponsePlanSpecs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFindDeviceById200ResponsePlanSpecsWithDefaults() *FindDeviceById200ResponsePlanSpecs {
	this := FindDeviceById200ResponsePlanSpecs{}
	return &this
}

// GetCpus returns the Cpus field value if set, zero value otherwise.
func (o *FindDeviceById200ResponsePlanSpecs) GetCpus() []FindDeviceById200ResponsePlanSpecsCpusInner {
	if o == nil || o.Cpus == nil {
		var ret []FindDeviceById200ResponsePlanSpecsCpusInner
		return ret
	}
	return o.Cpus
}

// GetCpusOk returns a tuple with the Cpus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindDeviceById200ResponsePlanSpecs) GetCpusOk() ([]FindDeviceById200ResponsePlanSpecsCpusInner, bool) {
	if o == nil || o.Cpus == nil {
		return nil, false
	}
	return o.Cpus, true
}

// HasCpus returns a boolean if a field has been set.
func (o *FindDeviceById200ResponsePlanSpecs) HasCpus() bool {
	if o != nil && o.Cpus != nil {
		return true
	}

	return false
}

// SetCpus gets a reference to the given []FindDeviceById200ResponsePlanSpecsCpusInner and assigns it to the Cpus field.
func (o *FindDeviceById200ResponsePlanSpecs) SetCpus(v []FindDeviceById200ResponsePlanSpecsCpusInner) {
	o.Cpus = v
}

// GetDrives returns the Drives field value if set, zero value otherwise.
func (o *FindDeviceById200ResponsePlanSpecs) GetDrives() []FindDeviceById200ResponsePlanSpecsDrivesInner {
	if o == nil || o.Drives == nil {
		var ret []FindDeviceById200ResponsePlanSpecsDrivesInner
		return ret
	}
	return o.Drives
}

// GetDrivesOk returns a tuple with the Drives field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindDeviceById200ResponsePlanSpecs) GetDrivesOk() ([]FindDeviceById200ResponsePlanSpecsDrivesInner, bool) {
	if o == nil || o.Drives == nil {
		return nil, false
	}
	return o.Drives, true
}

// HasDrives returns a boolean if a field has been set.
func (o *FindDeviceById200ResponsePlanSpecs) HasDrives() bool {
	if o != nil && o.Drives != nil {
		return true
	}

	return false
}

// SetDrives gets a reference to the given []FindDeviceById200ResponsePlanSpecsDrivesInner and assigns it to the Drives field.
func (o *FindDeviceById200ResponsePlanSpecs) SetDrives(v []FindDeviceById200ResponsePlanSpecsDrivesInner) {
	o.Drives = v
}

// GetNics returns the Nics field value if set, zero value otherwise.
func (o *FindDeviceById200ResponsePlanSpecs) GetNics() []FindDeviceById200ResponsePlanSpecsNicsInner {
	if o == nil || o.Nics == nil {
		var ret []FindDeviceById200ResponsePlanSpecsNicsInner
		return ret
	}
	return o.Nics
}

// GetNicsOk returns a tuple with the Nics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindDeviceById200ResponsePlanSpecs) GetNicsOk() ([]FindDeviceById200ResponsePlanSpecsNicsInner, bool) {
	if o == nil || o.Nics == nil {
		return nil, false
	}
	return o.Nics, true
}

// HasNics returns a boolean if a field has been set.
func (o *FindDeviceById200ResponsePlanSpecs) HasNics() bool {
	if o != nil && o.Nics != nil {
		return true
	}

	return false
}

// SetNics gets a reference to the given []FindDeviceById200ResponsePlanSpecsNicsInner and assigns it to the Nics field.
func (o *FindDeviceById200ResponsePlanSpecs) SetNics(v []FindDeviceById200ResponsePlanSpecsNicsInner) {
	o.Nics = v
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *FindDeviceById200ResponsePlanSpecs) GetFeatures() FindDeviceById200ResponsePlanSpecsFeatures {
	if o == nil || o.Features == nil {
		var ret FindDeviceById200ResponsePlanSpecsFeatures
		return ret
	}
	return *o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindDeviceById200ResponsePlanSpecs) GetFeaturesOk() (*FindDeviceById200ResponsePlanSpecsFeatures, bool) {
	if o == nil || o.Features == nil {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *FindDeviceById200ResponsePlanSpecs) HasFeatures() bool {
	if o != nil && o.Features != nil {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given FindDeviceById200ResponsePlanSpecsFeatures and assigns it to the Features field.
func (o *FindDeviceById200ResponsePlanSpecs) SetFeatures(v FindDeviceById200ResponsePlanSpecsFeatures) {
	o.Features = &v
}

func (o FindDeviceById200ResponsePlanSpecs) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Cpus != nil {
		toSerialize["cpus"] = o.Cpus
	}
	if o.Drives != nil {
		toSerialize["drives"] = o.Drives
	}
	if o.Nics != nil {
		toSerialize["nics"] = o.Nics
	}
	if o.Features != nil {
		toSerialize["features"] = o.Features
	}
	return json.Marshal(toSerialize)
}

type NullableFindDeviceById200ResponsePlanSpecs struct {
	value *FindDeviceById200ResponsePlanSpecs
	isSet bool
}

func (v NullableFindDeviceById200ResponsePlanSpecs) Get() *FindDeviceById200ResponsePlanSpecs {
	return v.value
}

func (v *NullableFindDeviceById200ResponsePlanSpecs) Set(val *FindDeviceById200ResponsePlanSpecs) {
	v.value = val
	v.isSet = true
}

func (v NullableFindDeviceById200ResponsePlanSpecs) IsSet() bool {
	return v.isSet
}

func (v *NullableFindDeviceById200ResponsePlanSpecs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFindDeviceById200ResponsePlanSpecs(val *FindDeviceById200ResponsePlanSpecs) *NullableFindDeviceById200ResponsePlanSpecs {
	return &NullableFindDeviceById200ResponsePlanSpecs{value: val, isSet: true}
}

func (v NullableFindDeviceById200ResponsePlanSpecs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFindDeviceById200ResponsePlanSpecs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
