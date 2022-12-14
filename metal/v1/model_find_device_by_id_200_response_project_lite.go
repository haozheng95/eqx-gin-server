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

// FindDeviceById200ResponseProjectLite struct for FindDeviceById200ResponseProjectLite
type FindDeviceById200ResponseProjectLite struct {
	Href string `json:"href"`
}

// NewFindDeviceById200ResponseProjectLite instantiates a new FindDeviceById200ResponseProjectLite object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFindDeviceById200ResponseProjectLite(href string) *FindDeviceById200ResponseProjectLite {
	this := FindDeviceById200ResponseProjectLite{}
	this.Href = href
	return &this
}

// NewFindDeviceById200ResponseProjectLiteWithDefaults instantiates a new FindDeviceById200ResponseProjectLite object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFindDeviceById200ResponseProjectLiteWithDefaults() *FindDeviceById200ResponseProjectLite {
	this := FindDeviceById200ResponseProjectLite{}
	return &this
}

// GetHref returns the Href field value
func (o *FindDeviceById200ResponseProjectLite) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *FindDeviceById200ResponseProjectLite) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *FindDeviceById200ResponseProjectLite) SetHref(v string) {
	o.Href = v
}

func (o FindDeviceById200ResponseProjectLite) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["href"] = o.Href
	}
	return json.Marshal(toSerialize)
}

type NullableFindDeviceById200ResponseProjectLite struct {
	value *FindDeviceById200ResponseProjectLite
	isSet bool
}

func (v NullableFindDeviceById200ResponseProjectLite) Get() *FindDeviceById200ResponseProjectLite {
	return v.value
}

func (v *NullableFindDeviceById200ResponseProjectLite) Set(val *FindDeviceById200ResponseProjectLite) {
	v.value = val
	v.isSet = true
}

func (v NullableFindDeviceById200ResponseProjectLite) IsSet() bool {
	return v.isSet
}

func (v *NullableFindDeviceById200ResponseProjectLite) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFindDeviceById200ResponseProjectLite(val *FindDeviceById200ResponseProjectLite) *NullableFindDeviceById200ResponseProjectLite {
	return &NullableFindDeviceById200ResponseProjectLite{value: val, isSet: true}
}

func (v NullableFindDeviceById200ResponseProjectLite) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFindDeviceById200ResponseProjectLite) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
