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

// FindIPAddressById200ResponseOneOf1 struct for FindIPAddressById200ResponseOneOf1
type FindIPAddressById200ResponseOneOf1 struct {
	AddressFamily *int32                                                                                      `json:"address_family,omitempty"`
	Cidr          *int32                                                                                      `json:"cidr,omitempty"`
	CreatedAt     *time.Time                                                                                  `json:"created_at,omitempty"`
	CreatedBy     *FindBatchById200ResponseDevicesInner                                                       `json:"created_by,omitempty"`
	Details       *string                                                                                     `json:"details,omitempty"`
	Href          *string                                                                                     `json:"href,omitempty"`
	Id            *string                                                                                     `json:"id,omitempty"`
	MetalGateway  *FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetworkMetalGatewaysInner           `json:"metal_gateway,omitempty"`
	Netmask       *string                                                                                     `json:"netmask,omitempty"`
	Network       *string                                                                                     `json:"network,omitempty"`
	Project       *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1VrfProject `json:"project,omitempty"`
	State         *string                                                                                     `json:"state,omitempty"`
	Tags          []string                                                                                    `json:"tags,omitempty"`
	Type          *string                                                                                     `json:"type,omitempty"`
	Vrf           GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1Vrf         `json:"vrf"`
}

// NewFindIPAddressById200ResponseOneOf1 instantiates a new FindIPAddressById200ResponseOneOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFindIPAddressById200ResponseOneOf1(vrf GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1Vrf) *FindIPAddressById200ResponseOneOf1 {
	this := FindIPAddressById200ResponseOneOf1{}
	this.Vrf = vrf
	return &this
}

// NewFindIPAddressById200ResponseOneOf1WithDefaults instantiates a new FindIPAddressById200ResponseOneOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFindIPAddressById200ResponseOneOf1WithDefaults() *FindIPAddressById200ResponseOneOf1 {
	this := FindIPAddressById200ResponseOneOf1{}
	return &this
}

// GetAddressFamily returns the AddressFamily field value if set, zero value otherwise.
func (o *FindIPAddressById200ResponseOneOf1) GetAddressFamily() int32 {
	if o == nil || o.AddressFamily == nil {
		var ret int32
		return ret
	}
	return *o.AddressFamily
}

// GetAddressFamilyOk returns a tuple with the AddressFamily field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindIPAddressById200ResponseOneOf1) GetAddressFamilyOk() (*int32, bool) {
	if o == nil || o.AddressFamily == nil {
		return nil, false
	}
	return o.AddressFamily, true
}

// HasAddressFamily returns a boolean if a field has been set.
func (o *FindIPAddressById200ResponseOneOf1) HasAddressFamily() bool {
	if o != nil && o.AddressFamily != nil {
		return true
	}

	return false
}

// SetAddressFamily gets a reference to the given int32 and assigns it to the AddressFamily field.
func (o *FindIPAddressById200ResponseOneOf1) SetAddressFamily(v int32) {
	o.AddressFamily = &v
}

// GetCidr returns the Cidr field value if set, zero value otherwise.
func (o *FindIPAddressById200ResponseOneOf1) GetCidr() int32 {
	if o == nil || o.Cidr == nil {
		var ret int32
		return ret
	}
	return *o.Cidr
}

// GetCidrOk returns a tuple with the Cidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindIPAddressById200ResponseOneOf1) GetCidrOk() (*int32, bool) {
	if o == nil || o.Cidr == nil {
		return nil, false
	}
	return o.Cidr, true
}

// HasCidr returns a boolean if a field has been set.
func (o *FindIPAddressById200ResponseOneOf1) HasCidr() bool {
	if o != nil && o.Cidr != nil {
		return true
	}

	return false
}

// SetCidr gets a reference to the given int32 and assigns it to the Cidr field.
func (o *FindIPAddressById200ResponseOneOf1) SetCidr(v int32) {
	o.Cidr = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *FindIPAddressById200ResponseOneOf1) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindIPAddressById200ResponseOneOf1) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *FindIPAddressById200ResponseOneOf1) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *FindIPAddressById200ResponseOneOf1) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *FindIPAddressById200ResponseOneOf1) GetCreatedBy() FindBatchById200ResponseDevicesInner {
	if o == nil || o.CreatedBy == nil {
		var ret FindBatchById200ResponseDevicesInner
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindIPAddressById200ResponseOneOf1) GetCreatedByOk() (*FindBatchById200ResponseDevicesInner, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *FindIPAddressById200ResponseOneOf1) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given FindBatchById200ResponseDevicesInner and assigns it to the CreatedBy field.
func (o *FindIPAddressById200ResponseOneOf1) SetCreatedBy(v FindBatchById200ResponseDevicesInner) {
	o.CreatedBy = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *FindIPAddressById200ResponseOneOf1) GetDetails() string {
	if o == nil || o.Details == nil {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindIPAddressById200ResponseOneOf1) GetDetailsOk() (*string, bool) {
	if o == nil || o.Details == nil {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *FindIPAddressById200ResponseOneOf1) HasDetails() bool {
	if o != nil && o.Details != nil {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *FindIPAddressById200ResponseOneOf1) SetDetails(v string) {
	o.Details = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *FindIPAddressById200ResponseOneOf1) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindIPAddressById200ResponseOneOf1) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *FindIPAddressById200ResponseOneOf1) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *FindIPAddressById200ResponseOneOf1) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FindIPAddressById200ResponseOneOf1) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindIPAddressById200ResponseOneOf1) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FindIPAddressById200ResponseOneOf1) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *FindIPAddressById200ResponseOneOf1) SetId(v string) {
	o.Id = &v
}

// GetMetalGateway returns the MetalGateway field value if set, zero value otherwise.
func (o *FindIPAddressById200ResponseOneOf1) GetMetalGateway() FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetworkMetalGatewaysInner {
	if o == nil || o.MetalGateway == nil {
		var ret FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetworkMetalGatewaysInner
		return ret
	}
	return *o.MetalGateway
}

// GetMetalGatewayOk returns a tuple with the MetalGateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindIPAddressById200ResponseOneOf1) GetMetalGatewayOk() (*FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetworkMetalGatewaysInner, bool) {
	if o == nil || o.MetalGateway == nil {
		return nil, false
	}
	return o.MetalGateway, true
}

// HasMetalGateway returns a boolean if a field has been set.
func (o *FindIPAddressById200ResponseOneOf1) HasMetalGateway() bool {
	if o != nil && o.MetalGateway != nil {
		return true
	}

	return false
}

// SetMetalGateway gets a reference to the given FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetworkMetalGatewaysInner and assigns it to the MetalGateway field.
func (o *FindIPAddressById200ResponseOneOf1) SetMetalGateway(v FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetworkMetalGatewaysInner) {
	o.MetalGateway = &v
}

// GetNetmask returns the Netmask field value if set, zero value otherwise.
func (o *FindIPAddressById200ResponseOneOf1) GetNetmask() string {
	if o == nil || o.Netmask == nil {
		var ret string
		return ret
	}
	return *o.Netmask
}

// GetNetmaskOk returns a tuple with the Netmask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindIPAddressById200ResponseOneOf1) GetNetmaskOk() (*string, bool) {
	if o == nil || o.Netmask == nil {
		return nil, false
	}
	return o.Netmask, true
}

// HasNetmask returns a boolean if a field has been set.
func (o *FindIPAddressById200ResponseOneOf1) HasNetmask() bool {
	if o != nil && o.Netmask != nil {
		return true
	}

	return false
}

// SetNetmask gets a reference to the given string and assigns it to the Netmask field.
func (o *FindIPAddressById200ResponseOneOf1) SetNetmask(v string) {
	o.Netmask = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *FindIPAddressById200ResponseOneOf1) GetNetwork() string {
	if o == nil || o.Network == nil {
		var ret string
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindIPAddressById200ResponseOneOf1) GetNetworkOk() (*string, bool) {
	if o == nil || o.Network == nil {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *FindIPAddressById200ResponseOneOf1) HasNetwork() bool {
	if o != nil && o.Network != nil {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given string and assigns it to the Network field.
func (o *FindIPAddressById200ResponseOneOf1) SetNetwork(v string) {
	o.Network = &v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *FindIPAddressById200ResponseOneOf1) GetProject() GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1VrfProject {
	if o == nil || o.Project == nil {
		var ret GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1VrfProject
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindIPAddressById200ResponseOneOf1) GetProjectOk() (*GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1VrfProject, bool) {
	if o == nil || o.Project == nil {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *FindIPAddressById200ResponseOneOf1) HasProject() bool {
	if o != nil && o.Project != nil {
		return true
	}

	return false
}

// SetProject gets a reference to the given GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1VrfProject and assigns it to the Project field.
func (o *FindIPAddressById200ResponseOneOf1) SetProject(v GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1VrfProject) {
	o.Project = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *FindIPAddressById200ResponseOneOf1) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindIPAddressById200ResponseOneOf1) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *FindIPAddressById200ResponseOneOf1) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *FindIPAddressById200ResponseOneOf1) SetState(v string) {
	o.State = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *FindIPAddressById200ResponseOneOf1) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindIPAddressById200ResponseOneOf1) GetTagsOk() ([]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *FindIPAddressById200ResponseOneOf1) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *FindIPAddressById200ResponseOneOf1) SetTags(v []string) {
	o.Tags = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *FindIPAddressById200ResponseOneOf1) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindIPAddressById200ResponseOneOf1) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *FindIPAddressById200ResponseOneOf1) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *FindIPAddressById200ResponseOneOf1) SetType(v string) {
	o.Type = &v
}

// GetVrf returns the Vrf field value
func (o *FindIPAddressById200ResponseOneOf1) GetVrf() GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1Vrf {
	if o == nil {
		var ret GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1Vrf
		return ret
	}

	return o.Vrf
}

// GetVrfOk returns a tuple with the Vrf field value
// and a boolean to check if the value has been set.
func (o *FindIPAddressById200ResponseOneOf1) GetVrfOk() (*GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1Vrf, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vrf, true
}

// SetVrf sets field value
func (o *FindIPAddressById200ResponseOneOf1) SetVrf(v GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1Vrf) {
	o.Vrf = v
}

func (o FindIPAddressById200ResponseOneOf1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AddressFamily != nil {
		toSerialize["address_family"] = o.AddressFamily
	}
	if o.Cidr != nil {
		toSerialize["cidr"] = o.Cidr
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.CreatedBy != nil {
		toSerialize["created_by"] = o.CreatedBy
	}
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.MetalGateway != nil {
		toSerialize["metal_gateway"] = o.MetalGateway
	}
	if o.Netmask != nil {
		toSerialize["netmask"] = o.Netmask
	}
	if o.Network != nil {
		toSerialize["network"] = o.Network
	}
	if o.Project != nil {
		toSerialize["project"] = o.Project
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["vrf"] = o.Vrf
	}
	return json.Marshal(toSerialize)
}

type NullableFindIPAddressById200ResponseOneOf1 struct {
	value *FindIPAddressById200ResponseOneOf1
	isSet bool
}

func (v NullableFindIPAddressById200ResponseOneOf1) Get() *FindIPAddressById200ResponseOneOf1 {
	return v.value
}

func (v *NullableFindIPAddressById200ResponseOneOf1) Set(val *FindIPAddressById200ResponseOneOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableFindIPAddressById200ResponseOneOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableFindIPAddressById200ResponseOneOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFindIPAddressById200ResponseOneOf1(val *FindIPAddressById200ResponseOneOf1) *NullableFindIPAddressById200ResponseOneOf1 {
	return &NullableFindIPAddressById200ResponseOneOf1{value: val, isSet: true}
}

func (v NullableFindIPAddressById200ResponseOneOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFindIPAddressById200ResponseOneOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}