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

// VirtualCircuit struct for VirtualCircuit
type VirtualCircuit struct {
	// True if the Virtual Circuit is being billed. Currently, only Virtual Circuits of Fabric VCs (Metal Billed) will be billed. Usage will start the first time the Virtual Circuit becomes active, and will not stop until it is deleted from Metal.
	Bill        bool                                 `json:"bill"`
	Description string                               `json:"description"`
	Id          string                               `json:"id"`
	Name        string                               `json:"name"`
	NniVlan     int32                                `json:"nni_vlan"`
	Port        FindBatchById200ResponseDevicesInner `json:"port"`
	Project     FindBatchById200ResponseDevicesInner `json:"project"`
	// For Virtual Circuits on shared and dedicated connections, this speed should match the one set on their Interconnection Ports. For Virtual Circuits on Fabric VCs (both Metal and Fabric Billed) that have found their corresponding Fabric connection, this is the actual speed of the interconnection that was configured when setting up the interconnection on the Fabric Portal. Details on Fabric VCs are included in the specification as a developer preview and is generally unavailable. Please contact our Support team for more details.
	Speed *int32 `json:"speed,omitempty"`
	// The status of a Virtual Circuit is always 'Pending' on creation. The status can turn to 'Waiting on Customer VLAN' if the VLAN was not set yet on the Virtual Circuit and is the last step needed for full activation. For Dedicated interconnections, as long as the Dedicated Port has been associated to the Virtual Circuit and a NNI VNID has been set, it will turn to 'Waiting on Customer VLAN'. For Fabric VCs, which are not generally available, it will only change to 'Waiting on Customer VLAN' once the corresponding Fabric connection has been found on the Fabric side. Once a VLAN is set on the Virtual Circuit (which for Fabric VCs, can be set on creation) and the necessary set up is done, it will turn into 'Activating' status as it tries to activate the Virtual Circuit. Once the Virtual Circuit fully activates and is configured on the switch, it will turn to staus 'Active'. For Fabric VCs (Metal Billed), we will start billing the moment the status of the Virtual Circuit turns to 'Active'. If there are any changes to the VLAN after the Virtual Circuit is in an 'Active' status, the status will show 'Changing VLAN' if a new VLAN has been provided, or 'Deactivating' if we are removing the VLAN. When a deletion request is issued for the Virtual Circuit, it will move to a 'deleting' status until it is fully deleted. If the Virtual Circuit is on a Fabric VC, it can also change into an 'Expired' status if the associated service token has expired. To get access to Fabric VCs, please contact our Support Team for more details.
	Status         string                               `json:"status"`
	Tags           []string                             `json:"tags"`
	VirtualNetwork FindBatchById200ResponseDevicesInner `json:"virtual_network"`
	Vnid           int32                                `json:"vnid"`
}

// NewVirtualCircuit instantiates a new VirtualCircuit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualCircuit(bill bool, description string, id string, name string, nniVlan int32, port FindBatchById200ResponseDevicesInner, project FindBatchById200ResponseDevicesInner, status string, tags []string, virtualNetwork FindBatchById200ResponseDevicesInner, vnid int32) *VirtualCircuit {
	this := VirtualCircuit{}
	this.Bill = bill
	this.Description = description
	this.Id = id
	this.Name = name
	this.NniVlan = nniVlan
	this.Port = port
	this.Project = project
	this.Status = status
	this.Tags = tags
	this.VirtualNetwork = virtualNetwork
	this.Vnid = vnid
	return &this
}

// NewVirtualCircuitWithDefaults instantiates a new VirtualCircuit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualCircuitWithDefaults() *VirtualCircuit {
	this := VirtualCircuit{}
	var bill bool = false
	this.Bill = bill
	return &this
}

// GetBill returns the Bill field value
func (o *VirtualCircuit) GetBill() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Bill
}

// GetBillOk returns a tuple with the Bill field value
// and a boolean to check if the value has been set.
func (o *VirtualCircuit) GetBillOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bill, true
}

// SetBill sets field value
func (o *VirtualCircuit) SetBill(v bool) {
	o.Bill = v
}

// GetDescription returns the Description field value
func (o *VirtualCircuit) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *VirtualCircuit) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *VirtualCircuit) SetDescription(v string) {
	o.Description = v
}

// GetId returns the Id field value
func (o *VirtualCircuit) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *VirtualCircuit) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *VirtualCircuit) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *VirtualCircuit) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *VirtualCircuit) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *VirtualCircuit) SetName(v string) {
	o.Name = v
}

// GetNniVlan returns the NniVlan field value
func (o *VirtualCircuit) GetNniVlan() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NniVlan
}

// GetNniVlanOk returns a tuple with the NniVlan field value
// and a boolean to check if the value has been set.
func (o *VirtualCircuit) GetNniVlanOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NniVlan, true
}

// SetNniVlan sets field value
func (o *VirtualCircuit) SetNniVlan(v int32) {
	o.NniVlan = v
}

// GetPort returns the Port field value
func (o *VirtualCircuit) GetPort() FindBatchById200ResponseDevicesInner {
	if o == nil {
		var ret FindBatchById200ResponseDevicesInner
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *VirtualCircuit) GetPortOk() (*FindBatchById200ResponseDevicesInner, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *VirtualCircuit) SetPort(v FindBatchById200ResponseDevicesInner) {
	o.Port = v
}

// GetProject returns the Project field value
func (o *VirtualCircuit) GetProject() FindBatchById200ResponseDevicesInner {
	if o == nil {
		var ret FindBatchById200ResponseDevicesInner
		return ret
	}

	return o.Project
}

// GetProjectOk returns a tuple with the Project field value
// and a boolean to check if the value has been set.
func (o *VirtualCircuit) GetProjectOk() (*FindBatchById200ResponseDevicesInner, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Project, true
}

// SetProject sets field value
func (o *VirtualCircuit) SetProject(v FindBatchById200ResponseDevicesInner) {
	o.Project = v
}

// GetSpeed returns the Speed field value if set, zero value otherwise.
func (o *VirtualCircuit) GetSpeed() int32 {
	if o == nil || o.Speed == nil {
		var ret int32
		return ret
	}
	return *o.Speed
}

// GetSpeedOk returns a tuple with the Speed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualCircuit) GetSpeedOk() (*int32, bool) {
	if o == nil || o.Speed == nil {
		return nil, false
	}
	return o.Speed, true
}

// HasSpeed returns a boolean if a field has been set.
func (o *VirtualCircuit) HasSpeed() bool {
	if o != nil && o.Speed != nil {
		return true
	}

	return false
}

// SetSpeed gets a reference to the given int32 and assigns it to the Speed field.
func (o *VirtualCircuit) SetSpeed(v int32) {
	o.Speed = &v
}

// GetStatus returns the Status field value
func (o *VirtualCircuit) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *VirtualCircuit) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *VirtualCircuit) SetStatus(v string) {
	o.Status = v
}

// GetTags returns the Tags field value
func (o *VirtualCircuit) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *VirtualCircuit) GetTagsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *VirtualCircuit) SetTags(v []string) {
	o.Tags = v
}

// GetVirtualNetwork returns the VirtualNetwork field value
func (o *VirtualCircuit) GetVirtualNetwork() FindBatchById200ResponseDevicesInner {
	if o == nil {
		var ret FindBatchById200ResponseDevicesInner
		return ret
	}

	return o.VirtualNetwork
}

// GetVirtualNetworkOk returns a tuple with the VirtualNetwork field value
// and a boolean to check if the value has been set.
func (o *VirtualCircuit) GetVirtualNetworkOk() (*FindBatchById200ResponseDevicesInner, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VirtualNetwork, true
}

// SetVirtualNetwork sets field value
func (o *VirtualCircuit) SetVirtualNetwork(v FindBatchById200ResponseDevicesInner) {
	o.VirtualNetwork = v
}

// GetVnid returns the Vnid field value
func (o *VirtualCircuit) GetVnid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Vnid
}

// GetVnidOk returns a tuple with the Vnid field value
// and a boolean to check if the value has been set.
func (o *VirtualCircuit) GetVnidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vnid, true
}

// SetVnid sets field value
func (o *VirtualCircuit) SetVnid(v int32) {
	o.Vnid = v
}

func (o VirtualCircuit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["bill"] = o.Bill
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["nni_vlan"] = o.NniVlan
	}
	if true {
		toSerialize["port"] = o.Port
	}
	if true {
		toSerialize["project"] = o.Project
	}
	if o.Speed != nil {
		toSerialize["speed"] = o.Speed
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["tags"] = o.Tags
	}
	if true {
		toSerialize["virtual_network"] = o.VirtualNetwork
	}
	if true {
		toSerialize["vnid"] = o.Vnid
	}
	return json.Marshal(toSerialize)
}

type NullableVirtualCircuit struct {
	value *VirtualCircuit
	isSet bool
}

func (v NullableVirtualCircuit) Get() *VirtualCircuit {
	return v.value
}

func (v *NullableVirtualCircuit) Set(val *VirtualCircuit) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualCircuit) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualCircuit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualCircuit(val *VirtualCircuit) *NullableVirtualCircuit {
	return &NullableVirtualCircuit{value: val, isSet: true}
}

func (v NullableVirtualCircuit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualCircuit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
