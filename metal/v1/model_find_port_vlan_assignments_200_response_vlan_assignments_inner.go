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

// FindPortVlanAssignments200ResponseVlanAssignmentsInner struct for FindPortVlanAssignments200ResponseVlanAssignmentsInner
type FindPortVlanAssignments200ResponseVlanAssignmentsInner struct {
	CreatedAt      *time.Time                            `json:"created_at,omitempty"`
	Id             *string                               `json:"id,omitempty"`
	Native         *bool                                 `json:"native,omitempty"`
	Port           *FindBatchById200ResponseDevicesInner `json:"port,omitempty"`
	State          *string                               `json:"state,omitempty"`
	UpdatedAt      *time.Time                            `json:"updated_at,omitempty"`
	VirtualNetwork *FindBatchById200ResponseDevicesInner `json:"virtual_network,omitempty"`
	Vlan           *int32                                `json:"vlan,omitempty"`
}

// NewFindPortVlanAssignments200ResponseVlanAssignmentsInner instantiates a new FindPortVlanAssignments200ResponseVlanAssignmentsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFindPortVlanAssignments200ResponseVlanAssignmentsInner() *FindPortVlanAssignments200ResponseVlanAssignmentsInner {
	this := FindPortVlanAssignments200ResponseVlanAssignmentsInner{}
	return &this
}

// NewFindPortVlanAssignments200ResponseVlanAssignmentsInnerWithDefaults instantiates a new FindPortVlanAssignments200ResponseVlanAssignmentsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFindPortVlanAssignments200ResponseVlanAssignmentsInnerWithDefaults() *FindPortVlanAssignments200ResponseVlanAssignmentsInner {
	this := FindPortVlanAssignments200ResponseVlanAssignmentsInner{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *FindPortVlanAssignments200ResponseVlanAssignmentsInner) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindPortVlanAssignments200ResponseVlanAssignmentsInner) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *FindPortVlanAssignments200ResponseVlanAssignmentsInner) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *FindPortVlanAssignments200ResponseVlanAssignmentsInner) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FindPortVlanAssignments200ResponseVlanAssignmentsInner) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindPortVlanAssignments200ResponseVlanAssignmentsInner) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FindPortVlanAssignments200ResponseVlanAssignmentsInner) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *FindPortVlanAssignments200ResponseVlanAssignmentsInner) SetId(v string) {
	o.Id = &v
}

// GetNative returns the Native field value if set, zero value otherwise.
func (o *FindPortVlanAssignments200ResponseVlanAssignmentsInner) GetNative() bool {
	if o == nil || o.Native == nil {
		var ret bool
		return ret
	}
	return *o.Native
}

// GetNativeOk returns a tuple with the Native field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindPortVlanAssignments200ResponseVlanAssignmentsInner) GetNativeOk() (*bool, bool) {
	if o == nil || o.Native == nil {
		return nil, false
	}
	return o.Native, true
}

// HasNative returns a boolean if a field has been set.
func (o *FindPortVlanAssignments200ResponseVlanAssignmentsInner) HasNative() bool {
	if o != nil && o.Native != nil {
		return true
	}

	return false
}

// SetNative gets a reference to the given bool and assigns it to the Native field.
func (o *FindPortVlanAssignments200ResponseVlanAssignmentsInner) SetNative(v bool) {
	o.Native = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *FindPortVlanAssignments200ResponseVlanAssignmentsInner) GetPort() FindBatchById200ResponseDevicesInner {
	if o == nil || o.Port == nil {
		var ret FindBatchById200ResponseDevicesInner
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindPortVlanAssignments200ResponseVlanAssignmentsInner) GetPortOk() (*FindBatchById200ResponseDevicesInner, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *FindPortVlanAssignments200ResponseVlanAssignmentsInner) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given FindBatchById200ResponseDevicesInner and assigns it to the Port field.
func (o *FindPortVlanAssignments200ResponseVlanAssignmentsInner) SetPort(v FindBatchById200ResponseDevicesInner) {
	o.Port = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *FindPortVlanAssignments200ResponseVlanAssignmentsInner) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindPortVlanAssignments200ResponseVlanAssignmentsInner) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *FindPortVlanAssignments200ResponseVlanAssignmentsInner) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *FindPortVlanAssignments200ResponseVlanAssignmentsInner) SetState(v string) {
	o.State = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *FindPortVlanAssignments200ResponseVlanAssignmentsInner) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindPortVlanAssignments200ResponseVlanAssignmentsInner) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *FindPortVlanAssignments200ResponseVlanAssignmentsInner) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *FindPortVlanAssignments200ResponseVlanAssignmentsInner) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetVirtualNetwork returns the VirtualNetwork field value if set, zero value otherwise.
func (o *FindPortVlanAssignments200ResponseVlanAssignmentsInner) GetVirtualNetwork() FindBatchById200ResponseDevicesInner {
	if o == nil || o.VirtualNetwork == nil {
		var ret FindBatchById200ResponseDevicesInner
		return ret
	}
	return *o.VirtualNetwork
}

// GetVirtualNetworkOk returns a tuple with the VirtualNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindPortVlanAssignments200ResponseVlanAssignmentsInner) GetVirtualNetworkOk() (*FindBatchById200ResponseDevicesInner, bool) {
	if o == nil || o.VirtualNetwork == nil {
		return nil, false
	}
	return o.VirtualNetwork, true
}

// HasVirtualNetwork returns a boolean if a field has been set.
func (o *FindPortVlanAssignments200ResponseVlanAssignmentsInner) HasVirtualNetwork() bool {
	if o != nil && o.VirtualNetwork != nil {
		return true
	}

	return false
}

// SetVirtualNetwork gets a reference to the given FindBatchById200ResponseDevicesInner and assigns it to the VirtualNetwork field.
func (o *FindPortVlanAssignments200ResponseVlanAssignmentsInner) SetVirtualNetwork(v FindBatchById200ResponseDevicesInner) {
	o.VirtualNetwork = &v
}

// GetVlan returns the Vlan field value if set, zero value otherwise.
func (o *FindPortVlanAssignments200ResponseVlanAssignmentsInner) GetVlan() int32 {
	if o == nil || o.Vlan == nil {
		var ret int32
		return ret
	}
	return *o.Vlan
}

// GetVlanOk returns a tuple with the Vlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindPortVlanAssignments200ResponseVlanAssignmentsInner) GetVlanOk() (*int32, bool) {
	if o == nil || o.Vlan == nil {
		return nil, false
	}
	return o.Vlan, true
}

// HasVlan returns a boolean if a field has been set.
func (o *FindPortVlanAssignments200ResponseVlanAssignmentsInner) HasVlan() bool {
	if o != nil && o.Vlan != nil {
		return true
	}

	return false
}

// SetVlan gets a reference to the given int32 and assigns it to the Vlan field.
func (o *FindPortVlanAssignments200ResponseVlanAssignmentsInner) SetVlan(v int32) {
	o.Vlan = &v
}

func (o FindPortVlanAssignments200ResponseVlanAssignmentsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Native != nil {
		toSerialize["native"] = o.Native
	}
	if o.Port != nil {
		toSerialize["port"] = o.Port
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.VirtualNetwork != nil {
		toSerialize["virtual_network"] = o.VirtualNetwork
	}
	if o.Vlan != nil {
		toSerialize["vlan"] = o.Vlan
	}
	return json.Marshal(toSerialize)
}

type NullableFindPortVlanAssignments200ResponseVlanAssignmentsInner struct {
	value *FindPortVlanAssignments200ResponseVlanAssignmentsInner
	isSet bool
}

func (v NullableFindPortVlanAssignments200ResponseVlanAssignmentsInner) Get() *FindPortVlanAssignments200ResponseVlanAssignmentsInner {
	return v.value
}

func (v *NullableFindPortVlanAssignments200ResponseVlanAssignmentsInner) Set(val *FindPortVlanAssignments200ResponseVlanAssignmentsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableFindPortVlanAssignments200ResponseVlanAssignmentsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableFindPortVlanAssignments200ResponseVlanAssignmentsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFindPortVlanAssignments200ResponseVlanAssignmentsInner(val *FindPortVlanAssignments200ResponseVlanAssignmentsInner) *NullableFindPortVlanAssignments200ResponseVlanAssignmentsInner {
	return &NullableFindPortVlanAssignments200ResponseVlanAssignmentsInner{value: val, isSet: true}
}

func (v NullableFindPortVlanAssignments200ResponseVlanAssignmentsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFindPortVlanAssignments200ResponseVlanAssignmentsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}