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

// FindVirtualNetworks200Response struct for FindVirtualNetworks200Response
type FindVirtualNetworks200Response struct {
	VirtualNetworks []FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork `json:"virtual_networks,omitempty"`
}

// NewFindVirtualNetworks200Response instantiates a new FindVirtualNetworks200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFindVirtualNetworks200Response() *FindVirtualNetworks200Response {
	this := FindVirtualNetworks200Response{}
	return &this
}

// NewFindVirtualNetworks200ResponseWithDefaults instantiates a new FindVirtualNetworks200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFindVirtualNetworks200ResponseWithDefaults() *FindVirtualNetworks200Response {
	this := FindVirtualNetworks200Response{}
	return &this
}

// GetVirtualNetworks returns the VirtualNetworks field value if set, zero value otherwise.
func (o *FindVirtualNetworks200Response) GetVirtualNetworks() []FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork {
	if o == nil || o.VirtualNetworks == nil {
		var ret []FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork
		return ret
	}
	return o.VirtualNetworks
}

// GetVirtualNetworksOk returns a tuple with the VirtualNetworks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindVirtualNetworks200Response) GetVirtualNetworksOk() ([]FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork, bool) {
	if o == nil || o.VirtualNetworks == nil {
		return nil, false
	}
	return o.VirtualNetworks, true
}

// HasVirtualNetworks returns a boolean if a field has been set.
func (o *FindVirtualNetworks200Response) HasVirtualNetworks() bool {
	if o != nil && o.VirtualNetworks != nil {
		return true
	}

	return false
}

// SetVirtualNetworks gets a reference to the given []FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork and assigns it to the VirtualNetworks field.
func (o *FindVirtualNetworks200Response) SetVirtualNetworks(v []FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork) {
	o.VirtualNetworks = v
}

func (o FindVirtualNetworks200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.VirtualNetworks != nil {
		toSerialize["virtual_networks"] = o.VirtualNetworks
	}
	return json.Marshal(toSerialize)
}

type NullableFindVirtualNetworks200Response struct {
	value *FindVirtualNetworks200Response
	isSet bool
}

func (v NullableFindVirtualNetworks200Response) Get() *FindVirtualNetworks200Response {
	return v.value
}

func (v *NullableFindVirtualNetworks200Response) Set(val *FindVirtualNetworks200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableFindVirtualNetworks200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableFindVirtualNetworks200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFindVirtualNetworks200Response(val *FindVirtualNetworks200Response) *NullableFindVirtualNetworks200Response {
	return &NullableFindVirtualNetworks200Response{value: val, isSet: true}
}

func (v NullableFindVirtualNetworks200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFindVirtualNetworks200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
