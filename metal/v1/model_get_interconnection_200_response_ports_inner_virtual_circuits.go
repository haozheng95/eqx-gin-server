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

// GetInterconnection200ResponsePortsInnerVirtualCircuits struct for GetInterconnection200ResponsePortsInnerVirtualCircuits
type GetInterconnection200ResponsePortsInnerVirtualCircuits struct {
	VirtualCircuits []GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInner `json:"virtual_circuits,omitempty"`
}

// NewGetInterconnection200ResponsePortsInnerVirtualCircuits instantiates a new GetInterconnection200ResponsePortsInnerVirtualCircuits object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetInterconnection200ResponsePortsInnerVirtualCircuits() *GetInterconnection200ResponsePortsInnerVirtualCircuits {
	this := GetInterconnection200ResponsePortsInnerVirtualCircuits{}
	return &this
}

// NewGetInterconnection200ResponsePortsInnerVirtualCircuitsWithDefaults instantiates a new GetInterconnection200ResponsePortsInnerVirtualCircuits object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetInterconnection200ResponsePortsInnerVirtualCircuitsWithDefaults() *GetInterconnection200ResponsePortsInnerVirtualCircuits {
	this := GetInterconnection200ResponsePortsInnerVirtualCircuits{}
	return &this
}

// GetVirtualCircuits returns the VirtualCircuits field value if set, zero value otherwise.
func (o *GetInterconnection200ResponsePortsInnerVirtualCircuits) GetVirtualCircuits() []GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInner {
	if o == nil || o.VirtualCircuits == nil {
		var ret []GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInner
		return ret
	}
	return o.VirtualCircuits
}

// GetVirtualCircuitsOk returns a tuple with the VirtualCircuits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInterconnection200ResponsePortsInnerVirtualCircuits) GetVirtualCircuitsOk() ([]GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInner, bool) {
	if o == nil || o.VirtualCircuits == nil {
		return nil, false
	}
	return o.VirtualCircuits, true
}

// HasVirtualCircuits returns a boolean if a field has been set.
func (o *GetInterconnection200ResponsePortsInnerVirtualCircuits) HasVirtualCircuits() bool {
	if o != nil && o.VirtualCircuits != nil {
		return true
	}

	return false
}

// SetVirtualCircuits gets a reference to the given []GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInner and assigns it to the VirtualCircuits field.
func (o *GetInterconnection200ResponsePortsInnerVirtualCircuits) SetVirtualCircuits(v []GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInner) {
	o.VirtualCircuits = v
}

func (o GetInterconnection200ResponsePortsInnerVirtualCircuits) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.VirtualCircuits != nil {
		toSerialize["virtual_circuits"] = o.VirtualCircuits
	}
	return json.Marshal(toSerialize)
}

type NullableGetInterconnection200ResponsePortsInnerVirtualCircuits struct {
	value *GetInterconnection200ResponsePortsInnerVirtualCircuits
	isSet bool
}

func (v NullableGetInterconnection200ResponsePortsInnerVirtualCircuits) Get() *GetInterconnection200ResponsePortsInnerVirtualCircuits {
	return v.value
}

func (v *NullableGetInterconnection200ResponsePortsInnerVirtualCircuits) Set(val *GetInterconnection200ResponsePortsInnerVirtualCircuits) {
	v.value = val
	v.isSet = true
}

func (v NullableGetInterconnection200ResponsePortsInnerVirtualCircuits) IsSet() bool {
	return v.isSet
}

func (v *NullableGetInterconnection200ResponsePortsInnerVirtualCircuits) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetInterconnection200ResponsePortsInnerVirtualCircuits(val *GetInterconnection200ResponsePortsInnerVirtualCircuits) *NullableGetInterconnection200ResponsePortsInnerVirtualCircuits {
	return &NullableGetInterconnection200ResponsePortsInnerVirtualCircuits{value: val, isSet: true}
}

func (v NullableGetInterconnection200ResponsePortsInnerVirtualCircuits) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetInterconnection200ResponsePortsInnerVirtualCircuits) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
