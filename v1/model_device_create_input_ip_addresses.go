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

// DeviceCreateInputIpAddresses struct for DeviceCreateInputIpAddresses
type DeviceCreateInputIpAddresses struct {
	// Address Family for IP Address
	AddressFamily *float32 `json:"address_family,omitempty"`
	// Address Type for IP Address
	Public *bool `json:"public,omitempty"`
	// Cidr Size for the IP Block created. Valid values depends on the operating system being provisioned. (28..32 for IPv4 addresses, 124..127 for IPv6 addresses)
	Cidr *float32 `json:"cidr,omitempty"`
	// UUIDs of any IP reservations to use when assigning IPs
	IpReservations []string `json:"ip_reservations,omitempty"`
}

// NewDeviceCreateInputIpAddresses instantiates a new DeviceCreateInputIpAddresses object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceCreateInputIpAddresses() *DeviceCreateInputIpAddresses {
	this := DeviceCreateInputIpAddresses{}
	var public bool = true
	this.Public = &public
	return &this
}

// NewDeviceCreateInputIpAddressesWithDefaults instantiates a new DeviceCreateInputIpAddresses object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceCreateInputIpAddressesWithDefaults() *DeviceCreateInputIpAddresses {
	this := DeviceCreateInputIpAddresses{}
	var public bool = true
	this.Public = &public
	return &this
}

// GetAddressFamily returns the AddressFamily field value if set, zero value otherwise.
func (o *DeviceCreateInputIpAddresses) GetAddressFamily() float32 {
	if o == nil || o.AddressFamily == nil {
		var ret float32
		return ret
	}
	return *o.AddressFamily
}

// GetAddressFamilyOk returns a tuple with the AddressFamily field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceCreateInputIpAddresses) GetAddressFamilyOk() (*float32, bool) {
	if o == nil || o.AddressFamily == nil {
		return nil, false
	}
	return o.AddressFamily, true
}

// HasAddressFamily returns a boolean if a field has been set.
func (o *DeviceCreateInputIpAddresses) HasAddressFamily() bool {
	if o != nil && o.AddressFamily != nil {
		return true
	}

	return false
}

// SetAddressFamily gets a reference to the given float32 and assigns it to the AddressFamily field.
func (o *DeviceCreateInputIpAddresses) SetAddressFamily(v float32) {
	o.AddressFamily = &v
}

// GetPublic returns the Public field value if set, zero value otherwise.
func (o *DeviceCreateInputIpAddresses) GetPublic() bool {
	if o == nil || o.Public == nil {
		var ret bool
		return ret
	}
	return *o.Public
}

// GetPublicOk returns a tuple with the Public field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceCreateInputIpAddresses) GetPublicOk() (*bool, bool) {
	if o == nil || o.Public == nil {
		return nil, false
	}
	return o.Public, true
}

// HasPublic returns a boolean if a field has been set.
func (o *DeviceCreateInputIpAddresses) HasPublic() bool {
	if o != nil && o.Public != nil {
		return true
	}

	return false
}

// SetPublic gets a reference to the given bool and assigns it to the Public field.
func (o *DeviceCreateInputIpAddresses) SetPublic(v bool) {
	o.Public = &v
}

// GetCidr returns the Cidr field value if set, zero value otherwise.
func (o *DeviceCreateInputIpAddresses) GetCidr() float32 {
	if o == nil || o.Cidr == nil {
		var ret float32
		return ret
	}
	return *o.Cidr
}

// GetCidrOk returns a tuple with the Cidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceCreateInputIpAddresses) GetCidrOk() (*float32, bool) {
	if o == nil || o.Cidr == nil {
		return nil, false
	}
	return o.Cidr, true
}

// HasCidr returns a boolean if a field has been set.
func (o *DeviceCreateInputIpAddresses) HasCidr() bool {
	if o != nil && o.Cidr != nil {
		return true
	}

	return false
}

// SetCidr gets a reference to the given float32 and assigns it to the Cidr field.
func (o *DeviceCreateInputIpAddresses) SetCidr(v float32) {
	o.Cidr = &v
}

// GetIpReservations returns the IpReservations field value if set, zero value otherwise.
func (o *DeviceCreateInputIpAddresses) GetIpReservations() []string {
	if o == nil || o.IpReservations == nil {
		var ret []string
		return ret
	}
	return o.IpReservations
}

// GetIpReservationsOk returns a tuple with the IpReservations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceCreateInputIpAddresses) GetIpReservationsOk() ([]string, bool) {
	if o == nil || o.IpReservations == nil {
		return nil, false
	}
	return o.IpReservations, true
}

// HasIpReservations returns a boolean if a field has been set.
func (o *DeviceCreateInputIpAddresses) HasIpReservations() bool {
	if o != nil && o.IpReservations != nil {
		return true
	}

	return false
}

// SetIpReservations gets a reference to the given []string and assigns it to the IpReservations field.
func (o *DeviceCreateInputIpAddresses) SetIpReservations(v []string) {
	o.IpReservations = v
}

func (o DeviceCreateInputIpAddresses) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AddressFamily != nil {
		toSerialize["address_family"] = o.AddressFamily
	}
	if o.Public != nil {
		toSerialize["public"] = o.Public
	}
	if o.Cidr != nil {
		toSerialize["cidr"] = o.Cidr
	}
	if o.IpReservations != nil {
		toSerialize["ip_reservations"] = o.IpReservations
	}
	return json.Marshal(toSerialize)
}

type NullableDeviceCreateInputIpAddresses struct {
	value *DeviceCreateInputIpAddresses
	isSet bool
}

func (v NullableDeviceCreateInputIpAddresses) Get() *DeviceCreateInputIpAddresses {
	return v.value
}

func (v *NullableDeviceCreateInputIpAddresses) Set(val *DeviceCreateInputIpAddresses) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceCreateInputIpAddresses) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceCreateInputIpAddresses) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceCreateInputIpAddresses(val *DeviceCreateInputIpAddresses) *NullableDeviceCreateInputIpAddresses {
	return &NullableDeviceCreateInputIpAddresses{value: val, isSet: true}
}

func (v NullableDeviceCreateInputIpAddresses) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceCreateInputIpAddresses) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


