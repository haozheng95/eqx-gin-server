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

// FindMetalGatewaysByProject200Response struct for FindMetalGatewaysByProject200Response
type FindMetalGatewaysByProject200Response struct {
	MetalGateways []FindMetalGatewaysByProject200ResponseMetalGatewaysInner `json:"MetalGateways,omitempty"`
}

// NewFindMetalGatewaysByProject200Response instantiates a new FindMetalGatewaysByProject200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFindMetalGatewaysByProject200Response() *FindMetalGatewaysByProject200Response {
	this := FindMetalGatewaysByProject200Response{}
	return &this
}

// NewFindMetalGatewaysByProject200ResponseWithDefaults instantiates a new FindMetalGatewaysByProject200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFindMetalGatewaysByProject200ResponseWithDefaults() *FindMetalGatewaysByProject200Response {
	this := FindMetalGatewaysByProject200Response{}
	return &this
}

// GetMetalGateways returns the MetalGateways field value if set, zero value otherwise.
func (o *FindMetalGatewaysByProject200Response) GetMetalGateways() []FindMetalGatewaysByProject200ResponseMetalGatewaysInner {
	if o == nil || o.MetalGateways == nil {
		var ret []FindMetalGatewaysByProject200ResponseMetalGatewaysInner
		return ret
	}
	return o.MetalGateways
}

// GetMetalGatewaysOk returns a tuple with the MetalGateways field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindMetalGatewaysByProject200Response) GetMetalGatewaysOk() ([]FindMetalGatewaysByProject200ResponseMetalGatewaysInner, bool) {
	if o == nil || o.MetalGateways == nil {
		return nil, false
	}
	return o.MetalGateways, true
}

// HasMetalGateways returns a boolean if a field has been set.
func (o *FindMetalGatewaysByProject200Response) HasMetalGateways() bool {
	if o != nil && o.MetalGateways != nil {
		return true
	}

	return false
}

// SetMetalGateways gets a reference to the given []FindMetalGatewaysByProject200ResponseMetalGatewaysInner and assigns it to the MetalGateways field.
func (o *FindMetalGatewaysByProject200Response) SetMetalGateways(v []FindMetalGatewaysByProject200ResponseMetalGatewaysInner) {
	o.MetalGateways = v
}

func (o FindMetalGatewaysByProject200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MetalGateways != nil {
		toSerialize["MetalGateways"] = o.MetalGateways
	}
	return json.Marshal(toSerialize)
}

type NullableFindMetalGatewaysByProject200Response struct {
	value *FindMetalGatewaysByProject200Response
	isSet bool
}

func (v NullableFindMetalGatewaysByProject200Response) Get() *FindMetalGatewaysByProject200Response {
	return v.value
}

func (v *NullableFindMetalGatewaysByProject200Response) Set(val *FindMetalGatewaysByProject200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableFindMetalGatewaysByProject200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableFindMetalGatewaysByProject200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFindMetalGatewaysByProject200Response(val *FindMetalGatewaysByProject200Response) *NullableFindMetalGatewaysByProject200Response {
	return &NullableFindMetalGatewaysByProject200Response{value: val, isSet: true}
}

func (v NullableFindMetalGatewaysByProject200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFindMetalGatewaysByProject200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
