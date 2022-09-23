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

// IPReservation struct for IPReservation
type IPReservation struct {
	Addon         *bool                                                                             `json:"addon,omitempty"`
	AddressFamily *int32                                                                            `json:"address_family,omitempty"`
	Assignments   []FindDeviceById200ResponseIpAddressesInner                                       `json:"assignments,omitempty"`
	Bill          *bool                                                                             `json:"bill,omitempty"`
	Cidr          *int32                                                                            `json:"cidr,omitempty"`
	CreatedAt     *time.Time                                                                        `json:"created_at,omitempty"`
	Enabled       *bool                                                                             `json:"enabled,omitempty"`
	Facility      *FindIPAddressById200ResponseOneOfFacility                                        `json:"facility,omitempty"`
	GlobalIp      *bool                                                                             `json:"global_ip,omitempty"`
	Href          *string                                                                           `json:"href,omitempty"`
	Id            *string                                                                           `json:"id,omitempty"`
	Manageable    *bool                                                                             `json:"manageable,omitempty"`
	Management    *bool                                                                             `json:"management,omitempty"`
	MetalGateway  *FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetworkMetalGatewaysInner `json:"metal_gateway,omitempty"`
	Metro         *FindIPAddressById200ResponseOneOfMetro                                           `json:"metro,omitempty"`
	Netmask       *string                                                                           `json:"netmask,omitempty"`
	Network       *string                                                                           `json:"network,omitempty"`
	Public        *bool                                                                             `json:"public,omitempty"`
	State         *string                                                                           `json:"state,omitempty"`
	Tags          []string                                                                          `json:"tags,omitempty"`
}

// NewIPReservation instantiates a new IPReservation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIPReservation() *IPReservation {
	this := IPReservation{}
	return &this
}

// NewIPReservationWithDefaults instantiates a new IPReservation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIPReservationWithDefaults() *IPReservation {
	this := IPReservation{}
	return &this
}

// GetAddon returns the Addon field value if set, zero value otherwise.
func (o *IPReservation) GetAddon() bool {
	if o == nil || o.Addon == nil {
		var ret bool
		return ret
	}
	return *o.Addon
}

// GetAddonOk returns a tuple with the Addon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPReservation) GetAddonOk() (*bool, bool) {
	if o == nil || o.Addon == nil {
		return nil, false
	}
	return o.Addon, true
}

// HasAddon returns a boolean if a field has been set.
func (o *IPReservation) HasAddon() bool {
	if o != nil && o.Addon != nil {
		return true
	}

	return false
}

// SetAddon gets a reference to the given bool and assigns it to the Addon field.
func (o *IPReservation) SetAddon(v bool) {
	o.Addon = &v
}

// GetAddressFamily returns the AddressFamily field value if set, zero value otherwise.
func (o *IPReservation) GetAddressFamily() int32 {
	if o == nil || o.AddressFamily == nil {
		var ret int32
		return ret
	}
	return *o.AddressFamily
}

// GetAddressFamilyOk returns a tuple with the AddressFamily field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPReservation) GetAddressFamilyOk() (*int32, bool) {
	if o == nil || o.AddressFamily == nil {
		return nil, false
	}
	return o.AddressFamily, true
}

// HasAddressFamily returns a boolean if a field has been set.
func (o *IPReservation) HasAddressFamily() bool {
	if o != nil && o.AddressFamily != nil {
		return true
	}

	return false
}

// SetAddressFamily gets a reference to the given int32 and assigns it to the AddressFamily field.
func (o *IPReservation) SetAddressFamily(v int32) {
	o.AddressFamily = &v
}

// GetAssignments returns the Assignments field value if set, zero value otherwise.
func (o *IPReservation) GetAssignments() []FindDeviceById200ResponseIpAddressesInner {
	if o == nil || o.Assignments == nil {
		var ret []FindDeviceById200ResponseIpAddressesInner
		return ret
	}
	return o.Assignments
}

// GetAssignmentsOk returns a tuple with the Assignments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPReservation) GetAssignmentsOk() ([]FindDeviceById200ResponseIpAddressesInner, bool) {
	if o == nil || o.Assignments == nil {
		return nil, false
	}
	return o.Assignments, true
}

// HasAssignments returns a boolean if a field has been set.
func (o *IPReservation) HasAssignments() bool {
	if o != nil && o.Assignments != nil {
		return true
	}

	return false
}

// SetAssignments gets a reference to the given []FindDeviceById200ResponseIpAddressesInner and assigns it to the Assignments field.
func (o *IPReservation) SetAssignments(v []FindDeviceById200ResponseIpAddressesInner) {
	o.Assignments = v
}

// GetBill returns the Bill field value if set, zero value otherwise.
func (o *IPReservation) GetBill() bool {
	if o == nil || o.Bill == nil {
		var ret bool
		return ret
	}
	return *o.Bill
}

// GetBillOk returns a tuple with the Bill field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPReservation) GetBillOk() (*bool, bool) {
	if o == nil || o.Bill == nil {
		return nil, false
	}
	return o.Bill, true
}

// HasBill returns a boolean if a field has been set.
func (o *IPReservation) HasBill() bool {
	if o != nil && o.Bill != nil {
		return true
	}

	return false
}

// SetBill gets a reference to the given bool and assigns it to the Bill field.
func (o *IPReservation) SetBill(v bool) {
	o.Bill = &v
}

// GetCidr returns the Cidr field value if set, zero value otherwise.
func (o *IPReservation) GetCidr() int32 {
	if o == nil || o.Cidr == nil {
		var ret int32
		return ret
	}
	return *o.Cidr
}

// GetCidrOk returns a tuple with the Cidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPReservation) GetCidrOk() (*int32, bool) {
	if o == nil || o.Cidr == nil {
		return nil, false
	}
	return o.Cidr, true
}

// HasCidr returns a boolean if a field has been set.
func (o *IPReservation) HasCidr() bool {
	if o != nil && o.Cidr != nil {
		return true
	}

	return false
}

// SetCidr gets a reference to the given int32 and assigns it to the Cidr field.
func (o *IPReservation) SetCidr(v int32) {
	o.Cidr = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *IPReservation) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPReservation) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *IPReservation) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *IPReservation) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *IPReservation) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPReservation) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *IPReservation) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *IPReservation) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetFacility returns the Facility field value if set, zero value otherwise.
func (o *IPReservation) GetFacility() FindIPAddressById200ResponseOneOfFacility {
	if o == nil || o.Facility == nil {
		var ret FindIPAddressById200ResponseOneOfFacility
		return ret
	}
	return *o.Facility
}

// GetFacilityOk returns a tuple with the Facility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPReservation) GetFacilityOk() (*FindIPAddressById200ResponseOneOfFacility, bool) {
	if o == nil || o.Facility == nil {
		return nil, false
	}
	return o.Facility, true
}

// HasFacility returns a boolean if a field has been set.
func (o *IPReservation) HasFacility() bool {
	if o != nil && o.Facility != nil {
		return true
	}

	return false
}

// SetFacility gets a reference to the given FindIPAddressById200ResponseOneOfFacility and assigns it to the Facility field.
func (o *IPReservation) SetFacility(v FindIPAddressById200ResponseOneOfFacility) {
	o.Facility = &v
}

// GetGlobalIp returns the GlobalIp field value if set, zero value otherwise.
func (o *IPReservation) GetGlobalIp() bool {
	if o == nil || o.GlobalIp == nil {
		var ret bool
		return ret
	}
	return *o.GlobalIp
}

// GetGlobalIpOk returns a tuple with the GlobalIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPReservation) GetGlobalIpOk() (*bool, bool) {
	if o == nil || o.GlobalIp == nil {
		return nil, false
	}
	return o.GlobalIp, true
}

// HasGlobalIp returns a boolean if a field has been set.
func (o *IPReservation) HasGlobalIp() bool {
	if o != nil && o.GlobalIp != nil {
		return true
	}

	return false
}

// SetGlobalIp gets a reference to the given bool and assigns it to the GlobalIp field.
func (o *IPReservation) SetGlobalIp(v bool) {
	o.GlobalIp = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *IPReservation) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPReservation) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *IPReservation) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *IPReservation) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IPReservation) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPReservation) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IPReservation) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IPReservation) SetId(v string) {
	o.Id = &v
}

// GetManageable returns the Manageable field value if set, zero value otherwise.
func (o *IPReservation) GetManageable() bool {
	if o == nil || o.Manageable == nil {
		var ret bool
		return ret
	}
	return *o.Manageable
}

// GetManageableOk returns a tuple with the Manageable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPReservation) GetManageableOk() (*bool, bool) {
	if o == nil || o.Manageable == nil {
		return nil, false
	}
	return o.Manageable, true
}

// HasManageable returns a boolean if a field has been set.
func (o *IPReservation) HasManageable() bool {
	if o != nil && o.Manageable != nil {
		return true
	}

	return false
}

// SetManageable gets a reference to the given bool and assigns it to the Manageable field.
func (o *IPReservation) SetManageable(v bool) {
	o.Manageable = &v
}

// GetManagement returns the Management field value if set, zero value otherwise.
func (o *IPReservation) GetManagement() bool {
	if o == nil || o.Management == nil {
		var ret bool
		return ret
	}
	return *o.Management
}

// GetManagementOk returns a tuple with the Management field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPReservation) GetManagementOk() (*bool, bool) {
	if o == nil || o.Management == nil {
		return nil, false
	}
	return o.Management, true
}

// HasManagement returns a boolean if a field has been set.
func (o *IPReservation) HasManagement() bool {
	if o != nil && o.Management != nil {
		return true
	}

	return false
}

// SetManagement gets a reference to the given bool and assigns it to the Management field.
func (o *IPReservation) SetManagement(v bool) {
	o.Management = &v
}

// GetMetalGateway returns the MetalGateway field value if set, zero value otherwise.
func (o *IPReservation) GetMetalGateway() FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetworkMetalGatewaysInner {
	if o == nil || o.MetalGateway == nil {
		var ret FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetworkMetalGatewaysInner
		return ret
	}
	return *o.MetalGateway
}

// GetMetalGatewayOk returns a tuple with the MetalGateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPReservation) GetMetalGatewayOk() (*FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetworkMetalGatewaysInner, bool) {
	if o == nil || o.MetalGateway == nil {
		return nil, false
	}
	return o.MetalGateway, true
}

// HasMetalGateway returns a boolean if a field has been set.
func (o *IPReservation) HasMetalGateway() bool {
	if o != nil && o.MetalGateway != nil {
		return true
	}

	return false
}

// SetMetalGateway gets a reference to the given FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetworkMetalGatewaysInner and assigns it to the MetalGateway field.
func (o *IPReservation) SetMetalGateway(v FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetworkMetalGatewaysInner) {
	o.MetalGateway = &v
}

// GetMetro returns the Metro field value if set, zero value otherwise.
func (o *IPReservation) GetMetro() FindIPAddressById200ResponseOneOfMetro {
	if o == nil || o.Metro == nil {
		var ret FindIPAddressById200ResponseOneOfMetro
		return ret
	}
	return *o.Metro
}

// GetMetroOk returns a tuple with the Metro field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPReservation) GetMetroOk() (*FindIPAddressById200ResponseOneOfMetro, bool) {
	if o == nil || o.Metro == nil {
		return nil, false
	}
	return o.Metro, true
}

// HasMetro returns a boolean if a field has been set.
func (o *IPReservation) HasMetro() bool {
	if o != nil && o.Metro != nil {
		return true
	}

	return false
}

// SetMetro gets a reference to the given FindIPAddressById200ResponseOneOfMetro and assigns it to the Metro field.
func (o *IPReservation) SetMetro(v FindIPAddressById200ResponseOneOfMetro) {
	o.Metro = &v
}

// GetNetmask returns the Netmask field value if set, zero value otherwise.
func (o *IPReservation) GetNetmask() string {
	if o == nil || o.Netmask == nil {
		var ret string
		return ret
	}
	return *o.Netmask
}

// GetNetmaskOk returns a tuple with the Netmask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPReservation) GetNetmaskOk() (*string, bool) {
	if o == nil || o.Netmask == nil {
		return nil, false
	}
	return o.Netmask, true
}

// HasNetmask returns a boolean if a field has been set.
func (o *IPReservation) HasNetmask() bool {
	if o != nil && o.Netmask != nil {
		return true
	}

	return false
}

// SetNetmask gets a reference to the given string and assigns it to the Netmask field.
func (o *IPReservation) SetNetmask(v string) {
	o.Netmask = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *IPReservation) GetNetwork() string {
	if o == nil || o.Network == nil {
		var ret string
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPReservation) GetNetworkOk() (*string, bool) {
	if o == nil || o.Network == nil {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *IPReservation) HasNetwork() bool {
	if o != nil && o.Network != nil {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given string and assigns it to the Network field.
func (o *IPReservation) SetNetwork(v string) {
	o.Network = &v
}

// GetPublic returns the Public field value if set, zero value otherwise.
func (o *IPReservation) GetPublic() bool {
	if o == nil || o.Public == nil {
		var ret bool
		return ret
	}
	return *o.Public
}

// GetPublicOk returns a tuple with the Public field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPReservation) GetPublicOk() (*bool, bool) {
	if o == nil || o.Public == nil {
		return nil, false
	}
	return o.Public, true
}

// HasPublic returns a boolean if a field has been set.
func (o *IPReservation) HasPublic() bool {
	if o != nil && o.Public != nil {
		return true
	}

	return false
}

// SetPublic gets a reference to the given bool and assigns it to the Public field.
func (o *IPReservation) SetPublic(v bool) {
	o.Public = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *IPReservation) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPReservation) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *IPReservation) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *IPReservation) SetState(v string) {
	o.State = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *IPReservation) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPReservation) GetTagsOk() ([]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *IPReservation) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *IPReservation) SetTags(v []string) {
	o.Tags = v
}

func (o IPReservation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Addon != nil {
		toSerialize["addon"] = o.Addon
	}
	if o.AddressFamily != nil {
		toSerialize["address_family"] = o.AddressFamily
	}
	if o.Assignments != nil {
		toSerialize["assignments"] = o.Assignments
	}
	if o.Bill != nil {
		toSerialize["bill"] = o.Bill
	}
	if o.Cidr != nil {
		toSerialize["cidr"] = o.Cidr
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Facility != nil {
		toSerialize["facility"] = o.Facility
	}
	if o.GlobalIp != nil {
		toSerialize["global_ip"] = o.GlobalIp
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Manageable != nil {
		toSerialize["manageable"] = o.Manageable
	}
	if o.Management != nil {
		toSerialize["management"] = o.Management
	}
	if o.MetalGateway != nil {
		toSerialize["metal_gateway"] = o.MetalGateway
	}
	if o.Metro != nil {
		toSerialize["metro"] = o.Metro
	}
	if o.Netmask != nil {
		toSerialize["netmask"] = o.Netmask
	}
	if o.Network != nil {
		toSerialize["network"] = o.Network
	}
	if o.Public != nil {
		toSerialize["public"] = o.Public
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableIPReservation struct {
	value *IPReservation
	isSet bool
}

func (v NullableIPReservation) Get() *IPReservation {
	return v.value
}

func (v *NullableIPReservation) Set(val *IPReservation) {
	v.value = val
	v.isSet = true
}

func (v NullableIPReservation) IsSet() bool {
	return v.isSet
}

func (v *NullableIPReservation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIPReservation(val *IPReservation) *NullableIPReservation {
	return &NullableIPReservation{value: val, isSet: true}
}

func (v NullableIPReservation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIPReservation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
