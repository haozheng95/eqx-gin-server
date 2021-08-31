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
	"time"
)

// PortVlanAssignment struct for PortVlanAssignment
type PortVlanAssignment struct {
	Id *string `json:"id,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	Native *bool `json:"native,omitempty"`
	State *string `json:"state,omitempty"`
	Vlan *int32 `json:"vlan,omitempty"`
	Port *Href `json:"port,omitempty"`
	VirtualNetwork *Href `json:"virtual_network,omitempty"`
}

// NewPortVlanAssignment instantiates a new PortVlanAssignment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortVlanAssignment() *PortVlanAssignment {
	this := PortVlanAssignment{}
	return &this
}

// NewPortVlanAssignmentWithDefaults instantiates a new PortVlanAssignment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortVlanAssignmentWithDefaults() *PortVlanAssignment {
	this := PortVlanAssignment{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PortVlanAssignment) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortVlanAssignment) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PortVlanAssignment) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PortVlanAssignment) SetId(v string) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *PortVlanAssignment) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortVlanAssignment) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PortVlanAssignment) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *PortVlanAssignment) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *PortVlanAssignment) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortVlanAssignment) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *PortVlanAssignment) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *PortVlanAssignment) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetNative returns the Native field value if set, zero value otherwise.
func (o *PortVlanAssignment) GetNative() bool {
	if o == nil || o.Native == nil {
		var ret bool
		return ret
	}
	return *o.Native
}

// GetNativeOk returns a tuple with the Native field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortVlanAssignment) GetNativeOk() (*bool, bool) {
	if o == nil || o.Native == nil {
		return nil, false
	}
	return o.Native, true
}

// HasNative returns a boolean if a field has been set.
func (o *PortVlanAssignment) HasNative() bool {
	if o != nil && o.Native != nil {
		return true
	}

	return false
}

// SetNative gets a reference to the given bool and assigns it to the Native field.
func (o *PortVlanAssignment) SetNative(v bool) {
	o.Native = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *PortVlanAssignment) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortVlanAssignment) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *PortVlanAssignment) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *PortVlanAssignment) SetState(v string) {
	o.State = &v
}

// GetVlan returns the Vlan field value if set, zero value otherwise.
func (o *PortVlanAssignment) GetVlan() int32 {
	if o == nil || o.Vlan == nil {
		var ret int32
		return ret
	}
	return *o.Vlan
}

// GetVlanOk returns a tuple with the Vlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortVlanAssignment) GetVlanOk() (*int32, bool) {
	if o == nil || o.Vlan == nil {
		return nil, false
	}
	return o.Vlan, true
}

// HasVlan returns a boolean if a field has been set.
func (o *PortVlanAssignment) HasVlan() bool {
	if o != nil && o.Vlan != nil {
		return true
	}

	return false
}

// SetVlan gets a reference to the given int32 and assigns it to the Vlan field.
func (o *PortVlanAssignment) SetVlan(v int32) {
	o.Vlan = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *PortVlanAssignment) GetPort() Href {
	if o == nil || o.Port == nil {
		var ret Href
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortVlanAssignment) GetPortOk() (*Href, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *PortVlanAssignment) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given Href and assigns it to the Port field.
func (o *PortVlanAssignment) SetPort(v Href) {
	o.Port = &v
}

// GetVirtualNetwork returns the VirtualNetwork field value if set, zero value otherwise.
func (o *PortVlanAssignment) GetVirtualNetwork() Href {
	if o == nil || o.VirtualNetwork == nil {
		var ret Href
		return ret
	}
	return *o.VirtualNetwork
}

// GetVirtualNetworkOk returns a tuple with the VirtualNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortVlanAssignment) GetVirtualNetworkOk() (*Href, bool) {
	if o == nil || o.VirtualNetwork == nil {
		return nil, false
	}
	return o.VirtualNetwork, true
}

// HasVirtualNetwork returns a boolean if a field has been set.
func (o *PortVlanAssignment) HasVirtualNetwork() bool {
	if o != nil && o.VirtualNetwork != nil {
		return true
	}

	return false
}

// SetVirtualNetwork gets a reference to the given Href and assigns it to the VirtualNetwork field.
func (o *PortVlanAssignment) SetVirtualNetwork(v Href) {
	o.VirtualNetwork = &v
}

func (o PortVlanAssignment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.Native != nil {
		toSerialize["native"] = o.Native
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.Vlan != nil {
		toSerialize["vlan"] = o.Vlan
	}
	if o.Port != nil {
		toSerialize["port"] = o.Port
	}
	if o.VirtualNetwork != nil {
		toSerialize["virtual_network"] = o.VirtualNetwork
	}
	return json.Marshal(toSerialize)
}

type NullablePortVlanAssignment struct {
	value *PortVlanAssignment
	isSet bool
}

func (v NullablePortVlanAssignment) Get() *PortVlanAssignment {
	return v.value
}

func (v *NullablePortVlanAssignment) Set(val *PortVlanAssignment) {
	v.value = val
	v.isSet = true
}

func (v NullablePortVlanAssignment) IsSet() bool {
	return v.isSet
}

func (v *NullablePortVlanAssignment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortVlanAssignment(val *PortVlanAssignment) *NullablePortVlanAssignment {
	return &NullablePortVlanAssignment{value: val, isSet: true}
}

func (v NullablePortVlanAssignment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortVlanAssignment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

