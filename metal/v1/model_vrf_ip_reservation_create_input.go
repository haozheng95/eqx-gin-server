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

// VrfIpReservationCreateInput struct for VrfIpReservationCreateInput
type VrfIpReservationCreateInput struct {
	// The size of the VRF IP Reservation's subnet
	Cidr       int32                  `json:"cidr"`
	Customdata map[string]interface{} `json:"customdata,omitempty"`
	Details    *string                `json:"details,omitempty"`
	// The starting address for this VRF IP Reservation's subnet
	Network string   `json:"network"`
	Tags    []string `json:"tags,omitempty"`
	// Must be set to 'vrf'
	Type string `json:"type"`
	// The ID of the VRF in which this VRF IP Reservation is created. The VRF must have an existing IP Range that contains the requested subnet. This field may be aliased as just 'vrf'.
	VrfId string `json:"vrf_id"`
}

// NewVrfIpReservationCreateInput instantiates a new VrfIpReservationCreateInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVrfIpReservationCreateInput(cidr int32, network string, type_ string, vrfId string) *VrfIpReservationCreateInput {
	this := VrfIpReservationCreateInput{}
	this.Cidr = cidr
	this.Network = network
	this.Type = type_
	this.VrfId = vrfId
	return &this
}

// NewVrfIpReservationCreateInputWithDefaults instantiates a new VrfIpReservationCreateInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVrfIpReservationCreateInputWithDefaults() *VrfIpReservationCreateInput {
	this := VrfIpReservationCreateInput{}
	return &this
}

// GetCidr returns the Cidr field value
func (o *VrfIpReservationCreateInput) GetCidr() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Cidr
}

// GetCidrOk returns a tuple with the Cidr field value
// and a boolean to check if the value has been set.
func (o *VrfIpReservationCreateInput) GetCidrOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cidr, true
}

// SetCidr sets field value
func (o *VrfIpReservationCreateInput) SetCidr(v int32) {
	o.Cidr = v
}

// GetCustomdata returns the Customdata field value if set, zero value otherwise.
func (o *VrfIpReservationCreateInput) GetCustomdata() map[string]interface{} {
	if o == nil || o.Customdata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Customdata
}

// GetCustomdataOk returns a tuple with the Customdata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfIpReservationCreateInput) GetCustomdataOk() (map[string]interface{}, bool) {
	if o == nil || o.Customdata == nil {
		return nil, false
	}
	return o.Customdata, true
}

// HasCustomdata returns a boolean if a field has been set.
func (o *VrfIpReservationCreateInput) HasCustomdata() bool {
	if o != nil && o.Customdata != nil {
		return true
	}

	return false
}

// SetCustomdata gets a reference to the given map[string]interface{} and assigns it to the Customdata field.
func (o *VrfIpReservationCreateInput) SetCustomdata(v map[string]interface{}) {
	o.Customdata = v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *VrfIpReservationCreateInput) GetDetails() string {
	if o == nil || o.Details == nil {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfIpReservationCreateInput) GetDetailsOk() (*string, bool) {
	if o == nil || o.Details == nil {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *VrfIpReservationCreateInput) HasDetails() bool {
	if o != nil && o.Details != nil {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *VrfIpReservationCreateInput) SetDetails(v string) {
	o.Details = &v
}

// GetNetwork returns the Network field value
func (o *VrfIpReservationCreateInput) GetNetwork() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Network
}

// GetNetworkOk returns a tuple with the Network field value
// and a boolean to check if the value has been set.
func (o *VrfIpReservationCreateInput) GetNetworkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Network, true
}

// SetNetwork sets field value
func (o *VrfIpReservationCreateInput) SetNetwork(v string) {
	o.Network = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *VrfIpReservationCreateInput) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfIpReservationCreateInput) GetTagsOk() ([]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *VrfIpReservationCreateInput) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *VrfIpReservationCreateInput) SetTags(v []string) {
	o.Tags = v
}

// GetType returns the Type field value
func (o *VrfIpReservationCreateInput) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *VrfIpReservationCreateInput) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *VrfIpReservationCreateInput) SetType(v string) {
	o.Type = v
}

// GetVrfId returns the VrfId field value
func (o *VrfIpReservationCreateInput) GetVrfId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VrfId
}

// GetVrfIdOk returns a tuple with the VrfId field value
// and a boolean to check if the value has been set.
func (o *VrfIpReservationCreateInput) GetVrfIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VrfId, true
}

// SetVrfId sets field value
func (o *VrfIpReservationCreateInput) SetVrfId(v string) {
	o.VrfId = v
}

func (o VrfIpReservationCreateInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cidr"] = o.Cidr
	}
	if o.Customdata != nil {
		toSerialize["customdata"] = o.Customdata
	}
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}
	if true {
		toSerialize["network"] = o.Network
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["vrf_id"] = o.VrfId
	}
	return json.Marshal(toSerialize)
}

type NullableVrfIpReservationCreateInput struct {
	value *VrfIpReservationCreateInput
	isSet bool
}

func (v NullableVrfIpReservationCreateInput) Get() *VrfIpReservationCreateInput {
	return v.value
}

func (v *NullableVrfIpReservationCreateInput) Set(val *VrfIpReservationCreateInput) {
	v.value = val
	v.isSet = true
}

func (v NullableVrfIpReservationCreateInput) IsSet() bool {
	return v.isSet
}

func (v *NullableVrfIpReservationCreateInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVrfIpReservationCreateInput(val *VrfIpReservationCreateInput) *NullableVrfIpReservationCreateInput {
	return &NullableVrfIpReservationCreateInput{value: val, isSet: true}
}

func (v NullableVrfIpReservationCreateInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVrfIpReservationCreateInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
