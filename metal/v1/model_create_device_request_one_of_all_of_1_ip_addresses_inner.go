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

// CreateDeviceRequestOneOfAllOf1IpAddressesInner struct for CreateDeviceRequestOneOfAllOf1IpAddressesInner
type CreateDeviceRequestOneOfAllOf1IpAddressesInner struct {
	// Address Family for IP Address
	AddressFamily *float32 `json:"address_family,omitempty"`
	// Cidr Size for the IP Block created. Valid values depends on the operating system being provisioned. (28..32 for IPv4 addresses, 124..127 for IPv6 addresses)
	Cidr *float32 `json:"cidr,omitempty"`
	// UUIDs of any IP reservations to use when assigning IPs
	IpReservations []string `json:"ip_reservations,omitempty"`
	// Address Type for IP Address
	Public *bool `json:"public,omitempty"`
}

// NewCreateDeviceRequestOneOfAllOf1IpAddressesInner instantiates a new CreateDeviceRequestOneOfAllOf1IpAddressesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDeviceRequestOneOfAllOf1IpAddressesInner() *CreateDeviceRequestOneOfAllOf1IpAddressesInner {
	this := CreateDeviceRequestOneOfAllOf1IpAddressesInner{}
	var public bool = true
	this.Public = &public
	return &this
}

// NewCreateDeviceRequestOneOfAllOf1IpAddressesInnerWithDefaults instantiates a new CreateDeviceRequestOneOfAllOf1IpAddressesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDeviceRequestOneOfAllOf1IpAddressesInnerWithDefaults() *CreateDeviceRequestOneOfAllOf1IpAddressesInner {
	this := CreateDeviceRequestOneOfAllOf1IpAddressesInner{}
	var public bool = true
	this.Public = &public
	return &this
}

// GetAddressFamily returns the AddressFamily field value if set, zero value otherwise.
func (o *CreateDeviceRequestOneOfAllOf1IpAddressesInner) GetAddressFamily() float32 {
	if o == nil || o.AddressFamily == nil {
		var ret float32
		return ret
	}
	return *o.AddressFamily
}

// GetAddressFamilyOk returns a tuple with the AddressFamily field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceRequestOneOfAllOf1IpAddressesInner) GetAddressFamilyOk() (*float32, bool) {
	if o == nil || o.AddressFamily == nil {
		return nil, false
	}
	return o.AddressFamily, true
}

// HasAddressFamily returns a boolean if a field has been set.
func (o *CreateDeviceRequestOneOfAllOf1IpAddressesInner) HasAddressFamily() bool {
	if o != nil && o.AddressFamily != nil {
		return true
	}

	return false
}

// SetAddressFamily gets a reference to the given float32 and assigns it to the AddressFamily field.
func (o *CreateDeviceRequestOneOfAllOf1IpAddressesInner) SetAddressFamily(v float32) {
	o.AddressFamily = &v
}

// GetCidr returns the Cidr field value if set, zero value otherwise.
func (o *CreateDeviceRequestOneOfAllOf1IpAddressesInner) GetCidr() float32 {
	if o == nil || o.Cidr == nil {
		var ret float32
		return ret
	}
	return *o.Cidr
}

// GetCidrOk returns a tuple with the Cidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceRequestOneOfAllOf1IpAddressesInner) GetCidrOk() (*float32, bool) {
	if o == nil || o.Cidr == nil {
		return nil, false
	}
	return o.Cidr, true
}

// HasCidr returns a boolean if a field has been set.
func (o *CreateDeviceRequestOneOfAllOf1IpAddressesInner) HasCidr() bool {
	if o != nil && o.Cidr != nil {
		return true
	}

	return false
}

// SetCidr gets a reference to the given float32 and assigns it to the Cidr field.
func (o *CreateDeviceRequestOneOfAllOf1IpAddressesInner) SetCidr(v float32) {
	o.Cidr = &v
}

// GetIpReservations returns the IpReservations field value if set, zero value otherwise.
func (o *CreateDeviceRequestOneOfAllOf1IpAddressesInner) GetIpReservations() []string {
	if o == nil || o.IpReservations == nil {
		var ret []string
		return ret
	}
	return o.IpReservations
}

// GetIpReservationsOk returns a tuple with the IpReservations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceRequestOneOfAllOf1IpAddressesInner) GetIpReservationsOk() ([]string, bool) {
	if o == nil || o.IpReservations == nil {
		return nil, false
	}
	return o.IpReservations, true
}

// HasIpReservations returns a boolean if a field has been set.
func (o *CreateDeviceRequestOneOfAllOf1IpAddressesInner) HasIpReservations() bool {
	if o != nil && o.IpReservations != nil {
		return true
	}

	return false
}

// SetIpReservations gets a reference to the given []string and assigns it to the IpReservations field.
func (o *CreateDeviceRequestOneOfAllOf1IpAddressesInner) SetIpReservations(v []string) {
	o.IpReservations = v
}

// GetPublic returns the Public field value if set, zero value otherwise.
func (o *CreateDeviceRequestOneOfAllOf1IpAddressesInner) GetPublic() bool {
	if o == nil || o.Public == nil {
		var ret bool
		return ret
	}
	return *o.Public
}

// GetPublicOk returns a tuple with the Public field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceRequestOneOfAllOf1IpAddressesInner) GetPublicOk() (*bool, bool) {
	if o == nil || o.Public == nil {
		return nil, false
	}
	return o.Public, true
}

// HasPublic returns a boolean if a field has been set.
func (o *CreateDeviceRequestOneOfAllOf1IpAddressesInner) HasPublic() bool {
	if o != nil && o.Public != nil {
		return true
	}

	return false
}

// SetPublic gets a reference to the given bool and assigns it to the Public field.
func (o *CreateDeviceRequestOneOfAllOf1IpAddressesInner) SetPublic(v bool) {
	o.Public = &v
}

func (o CreateDeviceRequestOneOfAllOf1IpAddressesInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AddressFamily != nil {
		toSerialize["address_family"] = o.AddressFamily
	}
	if o.Cidr != nil {
		toSerialize["cidr"] = o.Cidr
	}
	if o.IpReservations != nil {
		toSerialize["ip_reservations"] = o.IpReservations
	}
	if o.Public != nil {
		toSerialize["public"] = o.Public
	}
	return json.Marshal(toSerialize)
}

type NullableCreateDeviceRequestOneOfAllOf1IpAddressesInner struct {
	value *CreateDeviceRequestOneOfAllOf1IpAddressesInner
	isSet bool
}

func (v NullableCreateDeviceRequestOneOfAllOf1IpAddressesInner) Get() *CreateDeviceRequestOneOfAllOf1IpAddressesInner {
	return v.value
}

func (v *NullableCreateDeviceRequestOneOfAllOf1IpAddressesInner) Set(val *CreateDeviceRequestOneOfAllOf1IpAddressesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDeviceRequestOneOfAllOf1IpAddressesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDeviceRequestOneOfAllOf1IpAddressesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDeviceRequestOneOfAllOf1IpAddressesInner(val *CreateDeviceRequestOneOfAllOf1IpAddressesInner) *NullableCreateDeviceRequestOneOfAllOf1IpAddressesInner {
	return &NullableCreateDeviceRequestOneOfAllOf1IpAddressesInner{value: val, isSet: true}
}

func (v NullableCreateDeviceRequestOneOfAllOf1IpAddressesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDeviceRequestOneOfAllOf1IpAddressesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}