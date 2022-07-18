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

// UpdateInterconnectionRequest struct for UpdateInterconnectionRequest
type UpdateInterconnectionRequest struct {
	Tags         []string `json:"tags,omitempty"`
	ContactEmail *string  `json:"contact_email,omitempty"`
	Description  *string  `json:"description,omitempty"`
	// The mode of the connection (only relevant to dedicated connections). Shared connections won't have this field. Can be either 'standard' or 'tunnel'.   The default mode of a dedicated connection is 'standard'. The mode can only be changed when there are no associated virtual circuits on the connection.   In tunnel mode, an 802.1q tunnel is added to a port to send/receive double tagged packets from server instances.
	Mode *string `json:"mode,omitempty"`
	Name *string `json:"name,omitempty"`
	// Updating from 'redundant' to 'primary' will remove a secondary port, while updating from 'primary' to 'redundant' will add one.
	Redundancy *string `json:"redundancy,omitempty"`
}

// NewUpdateInterconnectionRequest instantiates a new UpdateInterconnectionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateInterconnectionRequest() *UpdateInterconnectionRequest {
	this := UpdateInterconnectionRequest{}
	return &this
}

// NewUpdateInterconnectionRequestWithDefaults instantiates a new UpdateInterconnectionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateInterconnectionRequestWithDefaults() *UpdateInterconnectionRequest {
	this := UpdateInterconnectionRequest{}
	return &this
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *UpdateInterconnectionRequest) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateInterconnectionRequest) GetTagsOk() ([]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *UpdateInterconnectionRequest) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *UpdateInterconnectionRequest) SetTags(v []string) {
	o.Tags = v
}

// GetContactEmail returns the ContactEmail field value if set, zero value otherwise.
func (o *UpdateInterconnectionRequest) GetContactEmail() string {
	if o == nil || o.ContactEmail == nil {
		var ret string
		return ret
	}
	return *o.ContactEmail
}

// GetContactEmailOk returns a tuple with the ContactEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateInterconnectionRequest) GetContactEmailOk() (*string, bool) {
	if o == nil || o.ContactEmail == nil {
		return nil, false
	}
	return o.ContactEmail, true
}

// HasContactEmail returns a boolean if a field has been set.
func (o *UpdateInterconnectionRequest) HasContactEmail() bool {
	if o != nil && o.ContactEmail != nil {
		return true
	}

	return false
}

// SetContactEmail gets a reference to the given string and assigns it to the ContactEmail field.
func (o *UpdateInterconnectionRequest) SetContactEmail(v string) {
	o.ContactEmail = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateInterconnectionRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateInterconnectionRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateInterconnectionRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateInterconnectionRequest) SetDescription(v string) {
	o.Description = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *UpdateInterconnectionRequest) GetMode() string {
	if o == nil || o.Mode == nil {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateInterconnectionRequest) GetModeOk() (*string, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *UpdateInterconnectionRequest) HasMode() bool {
	if o != nil && o.Mode != nil {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *UpdateInterconnectionRequest) SetMode(v string) {
	o.Mode = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateInterconnectionRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateInterconnectionRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateInterconnectionRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateInterconnectionRequest) SetName(v string) {
	o.Name = &v
}

// GetRedundancy returns the Redundancy field value if set, zero value otherwise.
func (o *UpdateInterconnectionRequest) GetRedundancy() string {
	if o == nil || o.Redundancy == nil {
		var ret string
		return ret
	}
	return *o.Redundancy
}

// GetRedundancyOk returns a tuple with the Redundancy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateInterconnectionRequest) GetRedundancyOk() (*string, bool) {
	if o == nil || o.Redundancy == nil {
		return nil, false
	}
	return o.Redundancy, true
}

// HasRedundancy returns a boolean if a field has been set.
func (o *UpdateInterconnectionRequest) HasRedundancy() bool {
	if o != nil && o.Redundancy != nil {
		return true
	}

	return false
}

// SetRedundancy gets a reference to the given string and assigns it to the Redundancy field.
func (o *UpdateInterconnectionRequest) SetRedundancy(v string) {
	o.Redundancy = &v
}

func (o UpdateInterconnectionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.ContactEmail != nil {
		toSerialize["contact_email"] = o.ContactEmail
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Mode != nil {
		toSerialize["mode"] = o.Mode
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Redundancy != nil {
		toSerialize["redundancy"] = o.Redundancy
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateInterconnectionRequest struct {
	value *UpdateInterconnectionRequest
	isSet bool
}

func (v NullableUpdateInterconnectionRequest) Get() *UpdateInterconnectionRequest {
	return v.value
}

func (v *NullableUpdateInterconnectionRequest) Set(val *UpdateInterconnectionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateInterconnectionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateInterconnectionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateInterconnectionRequest(val *UpdateInterconnectionRequest) *NullableUpdateInterconnectionRequest {
	return &NullableUpdateInterconnectionRequest{value: val, isSet: true}
}

func (v NullableUpdateInterconnectionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateInterconnectionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}