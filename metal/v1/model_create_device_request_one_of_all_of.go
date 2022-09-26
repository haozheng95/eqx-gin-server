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

// CreateDeviceRequestOneOfAllOf struct for CreateDeviceRequestOneOfAllOf
type CreateDeviceRequestOneOfAllOf struct {
	// Metro code or ID of where the instance should be provisioned in.  Either metro or facility must be provided.
	Metro string `json:"metro"`
}

// NewCreateDeviceRequestOneOfAllOf instantiates a new CreateDeviceRequestOneOfAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDeviceRequestOneOfAllOf(metro string) *CreateDeviceRequestOneOfAllOf {
	this := CreateDeviceRequestOneOfAllOf{}
	this.Metro = metro
	return &this
}

// NewCreateDeviceRequestOneOfAllOfWithDefaults instantiates a new CreateDeviceRequestOneOfAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDeviceRequestOneOfAllOfWithDefaults() *CreateDeviceRequestOneOfAllOf {
	this := CreateDeviceRequestOneOfAllOf{}
	return &this
}

// GetMetro returns the Metro field value
func (o *CreateDeviceRequestOneOfAllOf) GetMetro() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Metro
}

// GetMetroOk returns a tuple with the Metro field value
// and a boolean to check if the value has been set.
func (o *CreateDeviceRequestOneOfAllOf) GetMetroOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metro, true
}

// SetMetro sets field value
func (o *CreateDeviceRequestOneOfAllOf) SetMetro(v string) {
	o.Metro = v
}

func (o CreateDeviceRequestOneOfAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["metro"] = o.Metro
	}
	return json.Marshal(toSerialize)
}

type NullableCreateDeviceRequestOneOfAllOf struct {
	value *CreateDeviceRequestOneOfAllOf
	isSet bool
}

func (v NullableCreateDeviceRequestOneOfAllOf) Get() *CreateDeviceRequestOneOfAllOf {
	return v.value
}

func (v *NullableCreateDeviceRequestOneOfAllOf) Set(val *CreateDeviceRequestOneOfAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDeviceRequestOneOfAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDeviceRequestOneOfAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDeviceRequestOneOfAllOf(val *CreateDeviceRequestOneOfAllOf) *NullableCreateDeviceRequestOneOfAllOf {
	return &NullableCreateDeviceRequestOneOfAllOf{value: val, isSet: true}
}

func (v NullableCreateDeviceRequestOneOfAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDeviceRequestOneOfAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}