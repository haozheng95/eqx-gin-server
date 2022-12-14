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

// VrfVirtualCircuit struct for VrfVirtualCircuit
type VrfVirtualCircuit struct {
	// An IP address from the subnet that will be used on the Customer side. This parameter is optional, but if supplied, we will use the other usable IP address in the subnet as the Metal IP. By default, the last usable IP address in the subnet will be used.
	CustomerIp  *string `json:"customer_ip,omitempty"`
	Description *string `json:"description,omitempty"`
	Id          *string `json:"id,omitempty"`
	// The MD5 password for the BGP peering in plaintext (not a checksum).
	Md5 *string `json:"md5,omitempty"`
	// An IP address from the subnet that will be used on the Metal side. This parameter is optional, but if supplied, we will use the other usable IP address in the subnet as the Customer IP. By default, the first usable IP address in the subnet will be used.
	MetalIp *string                               `json:"metal_ip,omitempty"`
	Name    *string                               `json:"name,omitempty"`
	Port    *FindBatchById200ResponseDevicesInner `json:"port,omitempty"`
	NniVlan *int32                                `json:"nni_vlan,omitempty"`
	// The peer ASN that will be used with the VRF on the Virtual Circuit.
	PeerAsn *int32                                `json:"peer_asn,omitempty"`
	Project *FindBatchById200ResponseDevicesInner `json:"project,omitempty"`
	// integer representing bps speed
	Speed  *int32  `json:"speed,omitempty"`
	Status *string `json:"status,omitempty"`
	// The /30 or /31 subnet of one of the VRF IP Blocks that will be used with the VRF for the Virtual Circuit. This subnet does not have to be an existing VRF IP reservation, as we will create the VRF IP reservation on creation if it does not exist. The Metal IP and Customer IP must be IPs from this subnet. For /30 subnets, the network and broadcast IPs cannot be used as the Metal or Customer IP.
	Subnet *string                                                                              `json:"subnet,omitempty"`
	Tags   []string                                                                             `json:"tags,omitempty"`
	Vrf    *GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1Vrf `json:"vrf,omitempty"`
}

// NewVrfVirtualCircuit instantiates a new VrfVirtualCircuit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVrfVirtualCircuit() *VrfVirtualCircuit {
	this := VrfVirtualCircuit{}
	return &this
}

// NewVrfVirtualCircuitWithDefaults instantiates a new VrfVirtualCircuit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVrfVirtualCircuitWithDefaults() *VrfVirtualCircuit {
	this := VrfVirtualCircuit{}
	return &this
}

// GetCustomerIp returns the CustomerIp field value if set, zero value otherwise.
func (o *VrfVirtualCircuit) GetCustomerIp() string {
	if o == nil || o.CustomerIp == nil {
		var ret string
		return ret
	}
	return *o.CustomerIp
}

// GetCustomerIpOk returns a tuple with the CustomerIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfVirtualCircuit) GetCustomerIpOk() (*string, bool) {
	if o == nil || o.CustomerIp == nil {
		return nil, false
	}
	return o.CustomerIp, true
}

// HasCustomerIp returns a boolean if a field has been set.
func (o *VrfVirtualCircuit) HasCustomerIp() bool {
	if o != nil && o.CustomerIp != nil {
		return true
	}

	return false
}

// SetCustomerIp gets a reference to the given string and assigns it to the CustomerIp field.
func (o *VrfVirtualCircuit) SetCustomerIp(v string) {
	o.CustomerIp = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *VrfVirtualCircuit) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfVirtualCircuit) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *VrfVirtualCircuit) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *VrfVirtualCircuit) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VrfVirtualCircuit) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfVirtualCircuit) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VrfVirtualCircuit) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *VrfVirtualCircuit) SetId(v string) {
	o.Id = &v
}

// GetMd5 returns the Md5 field value if set, zero value otherwise.
func (o *VrfVirtualCircuit) GetMd5() string {
	if o == nil || o.Md5 == nil {
		var ret string
		return ret
	}
	return *o.Md5
}

// GetMd5Ok returns a tuple with the Md5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfVirtualCircuit) GetMd5Ok() (*string, bool) {
	if o == nil || o.Md5 == nil {
		return nil, false
	}
	return o.Md5, true
}

// HasMd5 returns a boolean if a field has been set.
func (o *VrfVirtualCircuit) HasMd5() bool {
	if o != nil && o.Md5 != nil {
		return true
	}

	return false
}

// SetMd5 gets a reference to the given string and assigns it to the Md5 field.
func (o *VrfVirtualCircuit) SetMd5(v string) {
	o.Md5 = &v
}

// GetMetalIp returns the MetalIp field value if set, zero value otherwise.
func (o *VrfVirtualCircuit) GetMetalIp() string {
	if o == nil || o.MetalIp == nil {
		var ret string
		return ret
	}
	return *o.MetalIp
}

// GetMetalIpOk returns a tuple with the MetalIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfVirtualCircuit) GetMetalIpOk() (*string, bool) {
	if o == nil || o.MetalIp == nil {
		return nil, false
	}
	return o.MetalIp, true
}

// HasMetalIp returns a boolean if a field has been set.
func (o *VrfVirtualCircuit) HasMetalIp() bool {
	if o != nil && o.MetalIp != nil {
		return true
	}

	return false
}

// SetMetalIp gets a reference to the given string and assigns it to the MetalIp field.
func (o *VrfVirtualCircuit) SetMetalIp(v string) {
	o.MetalIp = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VrfVirtualCircuit) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfVirtualCircuit) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VrfVirtualCircuit) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VrfVirtualCircuit) SetName(v string) {
	o.Name = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *VrfVirtualCircuit) GetPort() FindBatchById200ResponseDevicesInner {
	if o == nil || o.Port == nil {
		var ret FindBatchById200ResponseDevicesInner
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfVirtualCircuit) GetPortOk() (*FindBatchById200ResponseDevicesInner, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *VrfVirtualCircuit) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given FindBatchById200ResponseDevicesInner and assigns it to the Port field.
func (o *VrfVirtualCircuit) SetPort(v FindBatchById200ResponseDevicesInner) {
	o.Port = &v
}

// GetNniVlan returns the NniVlan field value if set, zero value otherwise.
func (o *VrfVirtualCircuit) GetNniVlan() int32 {
	if o == nil || o.NniVlan == nil {
		var ret int32
		return ret
	}
	return *o.NniVlan
}

// GetNniVlanOk returns a tuple with the NniVlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfVirtualCircuit) GetNniVlanOk() (*int32, bool) {
	if o == nil || o.NniVlan == nil {
		return nil, false
	}
	return o.NniVlan, true
}

// HasNniVlan returns a boolean if a field has been set.
func (o *VrfVirtualCircuit) HasNniVlan() bool {
	if o != nil && o.NniVlan != nil {
		return true
	}

	return false
}

// SetNniVlan gets a reference to the given int32 and assigns it to the NniVlan field.
func (o *VrfVirtualCircuit) SetNniVlan(v int32) {
	o.NniVlan = &v
}

// GetPeerAsn returns the PeerAsn field value if set, zero value otherwise.
func (o *VrfVirtualCircuit) GetPeerAsn() int32 {
	if o == nil || o.PeerAsn == nil {
		var ret int32
		return ret
	}
	return *o.PeerAsn
}

// GetPeerAsnOk returns a tuple with the PeerAsn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfVirtualCircuit) GetPeerAsnOk() (*int32, bool) {
	if o == nil || o.PeerAsn == nil {
		return nil, false
	}
	return o.PeerAsn, true
}

// HasPeerAsn returns a boolean if a field has been set.
func (o *VrfVirtualCircuit) HasPeerAsn() bool {
	if o != nil && o.PeerAsn != nil {
		return true
	}

	return false
}

// SetPeerAsn gets a reference to the given int32 and assigns it to the PeerAsn field.
func (o *VrfVirtualCircuit) SetPeerAsn(v int32) {
	o.PeerAsn = &v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *VrfVirtualCircuit) GetProject() FindBatchById200ResponseDevicesInner {
	if o == nil || o.Project == nil {
		var ret FindBatchById200ResponseDevicesInner
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfVirtualCircuit) GetProjectOk() (*FindBatchById200ResponseDevicesInner, bool) {
	if o == nil || o.Project == nil {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *VrfVirtualCircuit) HasProject() bool {
	if o != nil && o.Project != nil {
		return true
	}

	return false
}

// SetProject gets a reference to the given FindBatchById200ResponseDevicesInner and assigns it to the Project field.
func (o *VrfVirtualCircuit) SetProject(v FindBatchById200ResponseDevicesInner) {
	o.Project = &v
}

// GetSpeed returns the Speed field value if set, zero value otherwise.
func (o *VrfVirtualCircuit) GetSpeed() int32 {
	if o == nil || o.Speed == nil {
		var ret int32
		return ret
	}
	return *o.Speed
}

// GetSpeedOk returns a tuple with the Speed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfVirtualCircuit) GetSpeedOk() (*int32, bool) {
	if o == nil || o.Speed == nil {
		return nil, false
	}
	return o.Speed, true
}

// HasSpeed returns a boolean if a field has been set.
func (o *VrfVirtualCircuit) HasSpeed() bool {
	if o != nil && o.Speed != nil {
		return true
	}

	return false
}

// SetSpeed gets a reference to the given int32 and assigns it to the Speed field.
func (o *VrfVirtualCircuit) SetSpeed(v int32) {
	o.Speed = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *VrfVirtualCircuit) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfVirtualCircuit) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *VrfVirtualCircuit) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *VrfVirtualCircuit) SetStatus(v string) {
	o.Status = &v
}

// GetSubnet returns the Subnet field value if set, zero value otherwise.
func (o *VrfVirtualCircuit) GetSubnet() string {
	if o == nil || o.Subnet == nil {
		var ret string
		return ret
	}
	return *o.Subnet
}

// GetSubnetOk returns a tuple with the Subnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfVirtualCircuit) GetSubnetOk() (*string, bool) {
	if o == nil || o.Subnet == nil {
		return nil, false
	}
	return o.Subnet, true
}

// HasSubnet returns a boolean if a field has been set.
func (o *VrfVirtualCircuit) HasSubnet() bool {
	if o != nil && o.Subnet != nil {
		return true
	}

	return false
}

// SetSubnet gets a reference to the given string and assigns it to the Subnet field.
func (o *VrfVirtualCircuit) SetSubnet(v string) {
	o.Subnet = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *VrfVirtualCircuit) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfVirtualCircuit) GetTagsOk() ([]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *VrfVirtualCircuit) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *VrfVirtualCircuit) SetTags(v []string) {
	o.Tags = v
}

// GetVrf returns the Vrf field value if set, zero value otherwise.
func (o *VrfVirtualCircuit) GetVrf() GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1Vrf {
	if o == nil || o.Vrf == nil {
		var ret GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1Vrf
		return ret
	}
	return *o.Vrf
}

// GetVrfOk returns a tuple with the Vrf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfVirtualCircuit) GetVrfOk() (*GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1Vrf, bool) {
	if o == nil || o.Vrf == nil {
		return nil, false
	}
	return o.Vrf, true
}

// HasVrf returns a boolean if a field has been set.
func (o *VrfVirtualCircuit) HasVrf() bool {
	if o != nil && o.Vrf != nil {
		return true
	}

	return false
}

// SetVrf gets a reference to the given GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1Vrf and assigns it to the Vrf field.
func (o *VrfVirtualCircuit) SetVrf(v GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1Vrf) {
	o.Vrf = &v
}

func (o VrfVirtualCircuit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CustomerIp != nil {
		toSerialize["customer_ip"] = o.CustomerIp
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Md5 != nil {
		toSerialize["md5"] = o.Md5
	}
	if o.MetalIp != nil {
		toSerialize["metal_ip"] = o.MetalIp
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Port != nil {
		toSerialize["port"] = o.Port
	}
	if o.NniVlan != nil {
		toSerialize["nni_vlan"] = o.NniVlan
	}
	if o.PeerAsn != nil {
		toSerialize["peer_asn"] = o.PeerAsn
	}
	if o.Project != nil {
		toSerialize["project"] = o.Project
	}
	if o.Speed != nil {
		toSerialize["speed"] = o.Speed
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Subnet != nil {
		toSerialize["subnet"] = o.Subnet
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Vrf != nil {
		toSerialize["vrf"] = o.Vrf
	}
	return json.Marshal(toSerialize)
}

type NullableVrfVirtualCircuit struct {
	value *VrfVirtualCircuit
	isSet bool
}

func (v NullableVrfVirtualCircuit) Get() *VrfVirtualCircuit {
	return v.value
}

func (v *NullableVrfVirtualCircuit) Set(val *VrfVirtualCircuit) {
	v.value = val
	v.isSet = true
}

func (v NullableVrfVirtualCircuit) IsSet() bool {
	return v.isSet
}

func (v *NullableVrfVirtualCircuit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVrfVirtualCircuit(val *VrfVirtualCircuit) *NullableVrfVirtualCircuit {
	return &NullableVrfVirtualCircuit{value: val, isSet: true}
}

func (v NullableVrfVirtualCircuit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVrfVirtualCircuit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
