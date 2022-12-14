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

// FindProjectLicenses200Response struct for FindProjectLicenses200Response
type FindProjectLicenses200Response struct {
	Licenses []FindLicenseById200Response `json:"licenses,omitempty"`
}

// NewFindProjectLicenses200Response instantiates a new FindProjectLicenses200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFindProjectLicenses200Response() *FindProjectLicenses200Response {
	this := FindProjectLicenses200Response{}
	return &this
}

// NewFindProjectLicenses200ResponseWithDefaults instantiates a new FindProjectLicenses200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFindProjectLicenses200ResponseWithDefaults() *FindProjectLicenses200Response {
	this := FindProjectLicenses200Response{}
	return &this
}

// GetLicenses returns the Licenses field value if set, zero value otherwise.
func (o *FindProjectLicenses200Response) GetLicenses() []FindLicenseById200Response {
	if o == nil || o.Licenses == nil {
		var ret []FindLicenseById200Response
		return ret
	}
	return o.Licenses
}

// GetLicensesOk returns a tuple with the Licenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindProjectLicenses200Response) GetLicensesOk() ([]FindLicenseById200Response, bool) {
	if o == nil || o.Licenses == nil {
		return nil, false
	}
	return o.Licenses, true
}

// HasLicenses returns a boolean if a field has been set.
func (o *FindProjectLicenses200Response) HasLicenses() bool {
	if o != nil && o.Licenses != nil {
		return true
	}

	return false
}

// SetLicenses gets a reference to the given []FindLicenseById200Response and assigns it to the Licenses field.
func (o *FindProjectLicenses200Response) SetLicenses(v []FindLicenseById200Response) {
	o.Licenses = v
}

func (o FindProjectLicenses200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Licenses != nil {
		toSerialize["licenses"] = o.Licenses
	}
	return json.Marshal(toSerialize)
}

type NullableFindProjectLicenses200Response struct {
	value *FindProjectLicenses200Response
	isSet bool
}

func (v NullableFindProjectLicenses200Response) Get() *FindProjectLicenses200Response {
	return v.value
}

func (v *NullableFindProjectLicenses200Response) Set(val *FindProjectLicenses200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableFindProjectLicenses200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableFindProjectLicenses200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFindProjectLicenses200Response(val *FindProjectLicenses200Response) *NullableFindProjectLicenses200Response {
	return &NullableFindProjectLicenses200Response{value: val, isSet: true}
}

func (v NullableFindProjectLicenses200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFindProjectLicenses200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
