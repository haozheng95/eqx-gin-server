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

// FindCapacityForMetro200Response struct for FindCapacityForMetro200Response
type FindCapacityForMetro200Response struct {
	Capacity *FindCapacityForMetro200ResponseCapacity `json:"capacity,omitempty"`
}

// NewFindCapacityForMetro200Response instantiates a new FindCapacityForMetro200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFindCapacityForMetro200Response() *FindCapacityForMetro200Response {
	this := FindCapacityForMetro200Response{}
	return &this
}

// NewFindCapacityForMetro200ResponseWithDefaults instantiates a new FindCapacityForMetro200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFindCapacityForMetro200ResponseWithDefaults() *FindCapacityForMetro200Response {
	this := FindCapacityForMetro200Response{}
	return &this
}

// GetCapacity returns the Capacity field value if set, zero value otherwise.
func (o *FindCapacityForMetro200Response) GetCapacity() FindCapacityForMetro200ResponseCapacity {
	if o == nil || o.Capacity == nil {
		var ret FindCapacityForMetro200ResponseCapacity
		return ret
	}
	return *o.Capacity
}

// GetCapacityOk returns a tuple with the Capacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindCapacityForMetro200Response) GetCapacityOk() (*FindCapacityForMetro200ResponseCapacity, bool) {
	if o == nil || o.Capacity == nil {
		return nil, false
	}
	return o.Capacity, true
}

// HasCapacity returns a boolean if a field has been set.
func (o *FindCapacityForMetro200Response) HasCapacity() bool {
	if o != nil && o.Capacity != nil {
		return true
	}

	return false
}

// SetCapacity gets a reference to the given FindCapacityForMetro200ResponseCapacity and assigns it to the Capacity field.
func (o *FindCapacityForMetro200Response) SetCapacity(v FindCapacityForMetro200ResponseCapacity) {
	o.Capacity = &v
}

func (o FindCapacityForMetro200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Capacity != nil {
		toSerialize["capacity"] = o.Capacity
	}
	return json.Marshal(toSerialize)
}

type NullableFindCapacityForMetro200Response struct {
	value *FindCapacityForMetro200Response
	isSet bool
}

func (v NullableFindCapacityForMetro200Response) Get() *FindCapacityForMetro200Response {
	return v.value
}

func (v *NullableFindCapacityForMetro200Response) Set(val *FindCapacityForMetro200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableFindCapacityForMetro200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableFindCapacityForMetro200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFindCapacityForMetro200Response(val *FindCapacityForMetro200Response) *NullableFindCapacityForMetro200Response {
	return &NullableFindCapacityForMetro200Response{value: val, isSet: true}
}

func (v NullableFindCapacityForMetro200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFindCapacityForMetro200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}