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

// CreateMetalGatewayRequestOneOf struct for CreateMetalGatewayRequestOneOf
type CreateMetalGatewayRequestOneOf struct {
	// The UUID of an IP reservation that belongs to the same project as where the metal gateway will be created in. This field is required unless the private IPv4 subnet size is specified.
	IpReservationId *string `json:"ip_reservation_id,omitempty"`
	// The subnet size (8, 16, 32, 64, or 128) of the private IPv4 reservation that will be created for the metal gateway. This field is required unless a project IP reservation was specified.           Please keep in mind that the number of private metal gateway ranges are limited per project. If you would like to increase the limit per project, please contact support for assistance.
	PrivateIpv4SubnetSize *int32 `json:"private_ipv4_subnet_size,omitempty"`
	// The UUID of a metro virtual network that belongs to the same project as where the metal gateway will be created in.
	VirtualNetworkId string `json:"virtual_network_id"`
}

// NewCreateMetalGatewayRequestOneOf instantiates a new CreateMetalGatewayRequestOneOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateMetalGatewayRequestOneOf(virtualNetworkId string) *CreateMetalGatewayRequestOneOf {
	this := CreateMetalGatewayRequestOneOf{}
	this.VirtualNetworkId = virtualNetworkId
	return &this
}

// NewCreateMetalGatewayRequestOneOfWithDefaults instantiates a new CreateMetalGatewayRequestOneOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateMetalGatewayRequestOneOfWithDefaults() *CreateMetalGatewayRequestOneOf {
	this := CreateMetalGatewayRequestOneOf{}
	return &this
}

// GetIpReservationId returns the IpReservationId field value if set, zero value otherwise.
func (o *CreateMetalGatewayRequestOneOf) GetIpReservationId() string {
	if o == nil || o.IpReservationId == nil {
		var ret string
		return ret
	}
	return *o.IpReservationId
}

// GetIpReservationIdOk returns a tuple with the IpReservationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMetalGatewayRequestOneOf) GetIpReservationIdOk() (*string, bool) {
	if o == nil || o.IpReservationId == nil {
		return nil, false
	}
	return o.IpReservationId, true
}

// HasIpReservationId returns a boolean if a field has been set.
func (o *CreateMetalGatewayRequestOneOf) HasIpReservationId() bool {
	if o != nil && o.IpReservationId != nil {
		return true
	}

	return false
}

// SetIpReservationId gets a reference to the given string and assigns it to the IpReservationId field.
func (o *CreateMetalGatewayRequestOneOf) SetIpReservationId(v string) {
	o.IpReservationId = &v
}

// GetPrivateIpv4SubnetSize returns the PrivateIpv4SubnetSize field value if set, zero value otherwise.
func (o *CreateMetalGatewayRequestOneOf) GetPrivateIpv4SubnetSize() int32 {
	if o == nil || o.PrivateIpv4SubnetSize == nil {
		var ret int32
		return ret
	}
	return *o.PrivateIpv4SubnetSize
}

// GetPrivateIpv4SubnetSizeOk returns a tuple with the PrivateIpv4SubnetSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMetalGatewayRequestOneOf) GetPrivateIpv4SubnetSizeOk() (*int32, bool) {
	if o == nil || o.PrivateIpv4SubnetSize == nil {
		return nil, false
	}
	return o.PrivateIpv4SubnetSize, true
}

// HasPrivateIpv4SubnetSize returns a boolean if a field has been set.
func (o *CreateMetalGatewayRequestOneOf) HasPrivateIpv4SubnetSize() bool {
	if o != nil && o.PrivateIpv4SubnetSize != nil {
		return true
	}

	return false
}

// SetPrivateIpv4SubnetSize gets a reference to the given int32 and assigns it to the PrivateIpv4SubnetSize field.
func (o *CreateMetalGatewayRequestOneOf) SetPrivateIpv4SubnetSize(v int32) {
	o.PrivateIpv4SubnetSize = &v
}

// GetVirtualNetworkId returns the VirtualNetworkId field value
func (o *CreateMetalGatewayRequestOneOf) GetVirtualNetworkId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VirtualNetworkId
}

// GetVirtualNetworkIdOk returns a tuple with the VirtualNetworkId field value
// and a boolean to check if the value has been set.
func (o *CreateMetalGatewayRequestOneOf) GetVirtualNetworkIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VirtualNetworkId, true
}

// SetVirtualNetworkId sets field value
func (o *CreateMetalGatewayRequestOneOf) SetVirtualNetworkId(v string) {
	o.VirtualNetworkId = v
}

func (o CreateMetalGatewayRequestOneOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IpReservationId != nil {
		toSerialize["ip_reservation_id"] = o.IpReservationId
	}
	if o.PrivateIpv4SubnetSize != nil {
		toSerialize["private_ipv4_subnet_size"] = o.PrivateIpv4SubnetSize
	}
	if true {
		toSerialize["virtual_network_id"] = o.VirtualNetworkId
	}
	return json.Marshal(toSerialize)
}

type NullableCreateMetalGatewayRequestOneOf struct {
	value *CreateMetalGatewayRequestOneOf
	isSet bool
}

func (v NullableCreateMetalGatewayRequestOneOf) Get() *CreateMetalGatewayRequestOneOf {
	return v.value
}

func (v *NullableCreateMetalGatewayRequestOneOf) Set(val *CreateMetalGatewayRequestOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateMetalGatewayRequestOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateMetalGatewayRequestOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateMetalGatewayRequestOneOf(val *CreateMetalGatewayRequestOneOf) *NullableCreateMetalGatewayRequestOneOf {
	return &NullableCreateMetalGatewayRequestOneOf{value: val, isSet: true}
}

func (v NullableCreateMetalGatewayRequestOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateMetalGatewayRequestOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
