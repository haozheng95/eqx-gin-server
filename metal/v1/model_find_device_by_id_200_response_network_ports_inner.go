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

// FindDeviceById200ResponseNetworkPortsInner Port is a hardware port associated with a reserved or instantiated hardware device.
type FindDeviceById200ResponseNetworkPortsInner struct {
	Bond *FindDeviceById200ResponseNetworkPortsInnerBond `json:"bond,omitempty"`
	Data *FindDeviceById200ResponseNetworkPortsInnerData `json:"data,omitempty"`
	// Indicates whether or not the bond can be broken on the port (when applicable).
	DisbondOperationSupported *bool   `json:"disbond_operation_supported,omitempty"`
	Href                      *string `json:"href,omitempty"`
	Id                        *string `json:"id,omitempty"`
	Name                      *string `json:"name,omitempty"`
	// Type is either \"NetworkBondPort\" for bond ports or \"NetworkPort\" for bondable ethernet ports
	Type *string `json:"type,omitempty"`
	// Composite network type of the bond
	NetworkType          *string                                                         `json:"network_type,omitempty"`
	NativeVirtualNetwork *FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork `json:"native_virtual_network,omitempty"`
	VirtualNetworks      []FindBatchById200ResponseDevicesInner                          `json:"virtual_networks,omitempty"`
}

// NewFindDeviceById200ResponseNetworkPortsInner instantiates a new FindDeviceById200ResponseNetworkPortsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFindDeviceById200ResponseNetworkPortsInner() *FindDeviceById200ResponseNetworkPortsInner {
	this := FindDeviceById200ResponseNetworkPortsInner{}
	return &this
}

// NewFindDeviceById200ResponseNetworkPortsInnerWithDefaults instantiates a new FindDeviceById200ResponseNetworkPortsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFindDeviceById200ResponseNetworkPortsInnerWithDefaults() *FindDeviceById200ResponseNetworkPortsInner {
	this := FindDeviceById200ResponseNetworkPortsInner{}
	return &this
}

// GetBond returns the Bond field value if set, zero value otherwise.
func (o *FindDeviceById200ResponseNetworkPortsInner) GetBond() FindDeviceById200ResponseNetworkPortsInnerBond {
	if o == nil || o.Bond == nil {
		var ret FindDeviceById200ResponseNetworkPortsInnerBond
		return ret
	}
	return *o.Bond
}

// GetBondOk returns a tuple with the Bond field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindDeviceById200ResponseNetworkPortsInner) GetBondOk() (*FindDeviceById200ResponseNetworkPortsInnerBond, bool) {
	if o == nil || o.Bond == nil {
		return nil, false
	}
	return o.Bond, true
}

// HasBond returns a boolean if a field has been set.
func (o *FindDeviceById200ResponseNetworkPortsInner) HasBond() bool {
	if o != nil && o.Bond != nil {
		return true
	}

	return false
}

// SetBond gets a reference to the given FindDeviceById200ResponseNetworkPortsInnerBond and assigns it to the Bond field.
func (o *FindDeviceById200ResponseNetworkPortsInner) SetBond(v FindDeviceById200ResponseNetworkPortsInnerBond) {
	o.Bond = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *FindDeviceById200ResponseNetworkPortsInner) GetData() FindDeviceById200ResponseNetworkPortsInnerData {
	if o == nil || o.Data == nil {
		var ret FindDeviceById200ResponseNetworkPortsInnerData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindDeviceById200ResponseNetworkPortsInner) GetDataOk() (*FindDeviceById200ResponseNetworkPortsInnerData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *FindDeviceById200ResponseNetworkPortsInner) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given FindDeviceById200ResponseNetworkPortsInnerData and assigns it to the Data field.
func (o *FindDeviceById200ResponseNetworkPortsInner) SetData(v FindDeviceById200ResponseNetworkPortsInnerData) {
	o.Data = &v
}

// GetDisbondOperationSupported returns the DisbondOperationSupported field value if set, zero value otherwise.
func (o *FindDeviceById200ResponseNetworkPortsInner) GetDisbondOperationSupported() bool {
	if o == nil || o.DisbondOperationSupported == nil {
		var ret bool
		return ret
	}
	return *o.DisbondOperationSupported
}

// GetDisbondOperationSupportedOk returns a tuple with the DisbondOperationSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindDeviceById200ResponseNetworkPortsInner) GetDisbondOperationSupportedOk() (*bool, bool) {
	if o == nil || o.DisbondOperationSupported == nil {
		return nil, false
	}
	return o.DisbondOperationSupported, true
}

// HasDisbondOperationSupported returns a boolean if a field has been set.
func (o *FindDeviceById200ResponseNetworkPortsInner) HasDisbondOperationSupported() bool {
	if o != nil && o.DisbondOperationSupported != nil {
		return true
	}

	return false
}

// SetDisbondOperationSupported gets a reference to the given bool and assigns it to the DisbondOperationSupported field.
func (o *FindDeviceById200ResponseNetworkPortsInner) SetDisbondOperationSupported(v bool) {
	o.DisbondOperationSupported = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *FindDeviceById200ResponseNetworkPortsInner) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindDeviceById200ResponseNetworkPortsInner) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *FindDeviceById200ResponseNetworkPortsInner) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *FindDeviceById200ResponseNetworkPortsInner) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FindDeviceById200ResponseNetworkPortsInner) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindDeviceById200ResponseNetworkPortsInner) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FindDeviceById200ResponseNetworkPortsInner) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *FindDeviceById200ResponseNetworkPortsInner) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FindDeviceById200ResponseNetworkPortsInner) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindDeviceById200ResponseNetworkPortsInner) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FindDeviceById200ResponseNetworkPortsInner) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FindDeviceById200ResponseNetworkPortsInner) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *FindDeviceById200ResponseNetworkPortsInner) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindDeviceById200ResponseNetworkPortsInner) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *FindDeviceById200ResponseNetworkPortsInner) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *FindDeviceById200ResponseNetworkPortsInner) SetType(v string) {
	o.Type = &v
}

// GetNetworkType returns the NetworkType field value if set, zero value otherwise.
func (o *FindDeviceById200ResponseNetworkPortsInner) GetNetworkType() string {
	if o == nil || o.NetworkType == nil {
		var ret string
		return ret
	}
	return *o.NetworkType
}

// GetNetworkTypeOk returns a tuple with the NetworkType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindDeviceById200ResponseNetworkPortsInner) GetNetworkTypeOk() (*string, bool) {
	if o == nil || o.NetworkType == nil {
		return nil, false
	}
	return o.NetworkType, true
}

// HasNetworkType returns a boolean if a field has been set.
func (o *FindDeviceById200ResponseNetworkPortsInner) HasNetworkType() bool {
	if o != nil && o.NetworkType != nil {
		return true
	}

	return false
}

// SetNetworkType gets a reference to the given string and assigns it to the NetworkType field.
func (o *FindDeviceById200ResponseNetworkPortsInner) SetNetworkType(v string) {
	o.NetworkType = &v
}

// GetNativeVirtualNetwork returns the NativeVirtualNetwork field value if set, zero value otherwise.
func (o *FindDeviceById200ResponseNetworkPortsInner) GetNativeVirtualNetwork() FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork {
	if o == nil || o.NativeVirtualNetwork == nil {
		var ret FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork
		return ret
	}
	return *o.NativeVirtualNetwork
}

// GetNativeVirtualNetworkOk returns a tuple with the NativeVirtualNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindDeviceById200ResponseNetworkPortsInner) GetNativeVirtualNetworkOk() (*FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork, bool) {
	if o == nil || o.NativeVirtualNetwork == nil {
		return nil, false
	}
	return o.NativeVirtualNetwork, true
}

// HasNativeVirtualNetwork returns a boolean if a field has been set.
func (o *FindDeviceById200ResponseNetworkPortsInner) HasNativeVirtualNetwork() bool {
	if o != nil && o.NativeVirtualNetwork != nil {
		return true
	}

	return false
}

// SetNativeVirtualNetwork gets a reference to the given FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork and assigns it to the NativeVirtualNetwork field.
func (o *FindDeviceById200ResponseNetworkPortsInner) SetNativeVirtualNetwork(v FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork) {
	o.NativeVirtualNetwork = &v
}

// GetVirtualNetworks returns the VirtualNetworks field value if set, zero value otherwise.
func (o *FindDeviceById200ResponseNetworkPortsInner) GetVirtualNetworks() []FindBatchById200ResponseDevicesInner {
	if o == nil || o.VirtualNetworks == nil {
		var ret []FindBatchById200ResponseDevicesInner
		return ret
	}
	return o.VirtualNetworks
}

// GetVirtualNetworksOk returns a tuple with the VirtualNetworks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindDeviceById200ResponseNetworkPortsInner) GetVirtualNetworksOk() ([]FindBatchById200ResponseDevicesInner, bool) {
	if o == nil || o.VirtualNetworks == nil {
		return nil, false
	}
	return o.VirtualNetworks, true
}

// HasVirtualNetworks returns a boolean if a field has been set.
func (o *FindDeviceById200ResponseNetworkPortsInner) HasVirtualNetworks() bool {
	if o != nil && o.VirtualNetworks != nil {
		return true
	}

	return false
}

// SetVirtualNetworks gets a reference to the given []FindBatchById200ResponseDevicesInner and assigns it to the VirtualNetworks field.
func (o *FindDeviceById200ResponseNetworkPortsInner) SetVirtualNetworks(v []FindBatchById200ResponseDevicesInner) {
	o.VirtualNetworks = v
}

func (o FindDeviceById200ResponseNetworkPortsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Bond != nil {
		toSerialize["bond"] = o.Bond
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.DisbondOperationSupported != nil {
		toSerialize["disbond_operation_supported"] = o.DisbondOperationSupported
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.NetworkType != nil {
		toSerialize["network_type"] = o.NetworkType
	}
	if o.NativeVirtualNetwork != nil {
		toSerialize["native_virtual_network"] = o.NativeVirtualNetwork
	}
	if o.VirtualNetworks != nil {
		toSerialize["virtual_networks"] = o.VirtualNetworks
	}
	return json.Marshal(toSerialize)
}

type NullableFindDeviceById200ResponseNetworkPortsInner struct {
	value *FindDeviceById200ResponseNetworkPortsInner
	isSet bool
}

func (v NullableFindDeviceById200ResponseNetworkPortsInner) Get() *FindDeviceById200ResponseNetworkPortsInner {
	return v.value
}

func (v *NullableFindDeviceById200ResponseNetworkPortsInner) Set(val *FindDeviceById200ResponseNetworkPortsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableFindDeviceById200ResponseNetworkPortsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableFindDeviceById200ResponseNetworkPortsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFindDeviceById200ResponseNetworkPortsInner(val *FindDeviceById200ResponseNetworkPortsInner) *NullableFindDeviceById200ResponseNetworkPortsInner {
	return &NullableFindDeviceById200ResponseNetworkPortsInner{value: val, isSet: true}
}

func (v NullableFindDeviceById200ResponseNetworkPortsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFindDeviceById200ResponseNetworkPortsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
