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

// FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway struct for FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway
type FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The gateway address with subnet CIDR value for this Metal Gateway. For example, a Metal Gateway using an IP reservation with block 10.1.2.0/27 would have a gateway address of 10.1.2.1/27.
	GatewayAddress *string `json:"gateway_address,omitempty"`
	Href           *string `json:"href,omitempty"`
	Id             *string `json:"id,omitempty"`
	// The current state of the Metal Gateway. 'Ready' indicates the gateway record has been configured, but is currently not active on the network. 'Active' indicates the gateway has been configured on the network. 'Deleting' is a temporary state used to indicate that the gateway is in the process of being un-configured from the network, after which the gateway record will be deleted.
	State     *string    `json:"state,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// The VLAN id of the Virtual Network record associated to this Metal Gateway. Example: 1001.
	Vlan *float32 `json:"vlan,omitempty"`
}

// NewFindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway instantiates a new FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway() *FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway {
	this := FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway{}
	return &this
}

// NewFindVirtualNetworks200ResponseVirtualNetworksInnerMetalGatewayWithDefaults instantiates a new FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFindVirtualNetworks200ResponseVirtualNetworksInnerMetalGatewayWithDefaults() *FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway {
	this := FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetGatewayAddress returns the GatewayAddress field value if set, zero value otherwise.
func (o *FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway) GetGatewayAddress() string {
	if o == nil || o.GatewayAddress == nil {
		var ret string
		return ret
	}
	return *o.GatewayAddress
}

// GetGatewayAddressOk returns a tuple with the GatewayAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway) GetGatewayAddressOk() (*string, bool) {
	if o == nil || o.GatewayAddress == nil {
		return nil, false
	}
	return o.GatewayAddress, true
}

// HasGatewayAddress returns a boolean if a field has been set.
func (o *FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway) HasGatewayAddress() bool {
	if o != nil && o.GatewayAddress != nil {
		return true
	}

	return false
}

// SetGatewayAddress gets a reference to the given string and assigns it to the GatewayAddress field.
func (o *FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway) SetGatewayAddress(v string) {
	o.GatewayAddress = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway) SetId(v string) {
	o.Id = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway) SetState(v string) {
	o.State = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetVlan returns the Vlan field value if set, zero value otherwise.
func (o *FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway) GetVlan() float32 {
	if o == nil || o.Vlan == nil {
		var ret float32
		return ret
	}
	return *o.Vlan
}

// GetVlanOk returns a tuple with the Vlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway) GetVlanOk() (*float32, bool) {
	if o == nil || o.Vlan == nil {
		return nil, false
	}
	return o.Vlan, true
}

// HasVlan returns a boolean if a field has been set.
func (o *FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway) HasVlan() bool {
	if o != nil && o.Vlan != nil {
		return true
	}

	return false
}

// SetVlan gets a reference to the given float32 and assigns it to the Vlan field.
func (o *FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway) SetVlan(v float32) {
	o.Vlan = &v
}

func (o FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.GatewayAddress != nil {
		toSerialize["gateway_address"] = o.GatewayAddress
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.Vlan != nil {
		toSerialize["vlan"] = o.Vlan
	}
	return json.Marshal(toSerialize)
}

type NullableFindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway struct {
	value *FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway
	isSet bool
}

func (v NullableFindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway) Get() *FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway {
	return v.value
}

func (v *NullableFindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway) Set(val *FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway) {
	v.value = val
	v.isSet = true
}

func (v NullableFindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway) IsSet() bool {
	return v.isSet
}

func (v *NullableFindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway(val *FindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway) *NullableFindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway {
	return &NullableFindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway{value: val, isSet: true}
}

func (v NullableFindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFindVirtualNetworks200ResponseVirtualNetworksInnerMetalGateway) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}