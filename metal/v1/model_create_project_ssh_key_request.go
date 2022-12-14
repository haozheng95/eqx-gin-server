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

// CreateProjectSSHKeyRequest struct for CreateProjectSSHKeyRequest
type CreateProjectSSHKeyRequest struct {
	// List of instance UUIDs to associate SSH key with, when empty array is sent all instances belonging       to entity will be included
	InstancesIds []string `json:"instances_ids,omitempty"`
	Key          *string  `json:"key,omitempty"`
	Label        *string  `json:"label,omitempty"`
}

// NewCreateProjectSSHKeyRequest instantiates a new CreateProjectSSHKeyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateProjectSSHKeyRequest() *CreateProjectSSHKeyRequest {
	this := CreateProjectSSHKeyRequest{}
	return &this
}

// NewCreateProjectSSHKeyRequestWithDefaults instantiates a new CreateProjectSSHKeyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateProjectSSHKeyRequestWithDefaults() *CreateProjectSSHKeyRequest {
	this := CreateProjectSSHKeyRequest{}
	return &this
}

// GetInstancesIds returns the InstancesIds field value if set, zero value otherwise.
func (o *CreateProjectSSHKeyRequest) GetInstancesIds() []string {
	if o == nil || o.InstancesIds == nil {
		var ret []string
		return ret
	}
	return o.InstancesIds
}

// GetInstancesIdsOk returns a tuple with the InstancesIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProjectSSHKeyRequest) GetInstancesIdsOk() ([]string, bool) {
	if o == nil || o.InstancesIds == nil {
		return nil, false
	}
	return o.InstancesIds, true
}

// HasInstancesIds returns a boolean if a field has been set.
func (o *CreateProjectSSHKeyRequest) HasInstancesIds() bool {
	if o != nil && o.InstancesIds != nil {
		return true
	}

	return false
}

// SetInstancesIds gets a reference to the given []string and assigns it to the InstancesIds field.
func (o *CreateProjectSSHKeyRequest) SetInstancesIds(v []string) {
	o.InstancesIds = v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *CreateProjectSSHKeyRequest) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProjectSSHKeyRequest) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *CreateProjectSSHKeyRequest) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *CreateProjectSSHKeyRequest) SetKey(v string) {
	o.Key = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *CreateProjectSSHKeyRequest) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProjectSSHKeyRequest) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *CreateProjectSSHKeyRequest) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *CreateProjectSSHKeyRequest) SetLabel(v string) {
	o.Label = &v
}

func (o CreateProjectSSHKeyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.InstancesIds != nil {
		toSerialize["instances_ids"] = o.InstancesIds
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	return json.Marshal(toSerialize)
}

type NullableCreateProjectSSHKeyRequest struct {
	value *CreateProjectSSHKeyRequest
	isSet bool
}

func (v NullableCreateProjectSSHKeyRequest) Get() *CreateProjectSSHKeyRequest {
	return v.value
}

func (v *NullableCreateProjectSSHKeyRequest) Set(val *CreateProjectSSHKeyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateProjectSSHKeyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateProjectSSHKeyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateProjectSSHKeyRequest(val *CreateProjectSSHKeyRequest) *NullableCreateProjectSSHKeyRequest {
	return &NullableCreateProjectSSHKeyRequest{value: val, isSet: true}
}

func (v NullableCreateProjectSSHKeyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateProjectSSHKeyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
