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

// AuthTokenList struct for AuthTokenList
type AuthTokenList struct {
	ApiKeys []FindProjectAPIKeys200ResponseApiKeysInner `json:"api_keys,omitempty"`
}

// NewAuthTokenList instantiates a new AuthTokenList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthTokenList() *AuthTokenList {
	this := AuthTokenList{}
	return &this
}

// NewAuthTokenListWithDefaults instantiates a new AuthTokenList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthTokenListWithDefaults() *AuthTokenList {
	this := AuthTokenList{}
	return &this
}

// GetApiKeys returns the ApiKeys field value if set, zero value otherwise.
func (o *AuthTokenList) GetApiKeys() []FindProjectAPIKeys200ResponseApiKeysInner {
	if o == nil || o.ApiKeys == nil {
		var ret []FindProjectAPIKeys200ResponseApiKeysInner
		return ret
	}
	return o.ApiKeys
}

// GetApiKeysOk returns a tuple with the ApiKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthTokenList) GetApiKeysOk() ([]FindProjectAPIKeys200ResponseApiKeysInner, bool) {
	if o == nil || o.ApiKeys == nil {
		return nil, false
	}
	return o.ApiKeys, true
}

// HasApiKeys returns a boolean if a field has been set.
func (o *AuthTokenList) HasApiKeys() bool {
	if o != nil && o.ApiKeys != nil {
		return true
	}

	return false
}

// SetApiKeys gets a reference to the given []FindProjectAPIKeys200ResponseApiKeysInner and assigns it to the ApiKeys field.
func (o *AuthTokenList) SetApiKeys(v []FindProjectAPIKeys200ResponseApiKeysInner) {
	o.ApiKeys = v
}

func (o AuthTokenList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApiKeys != nil {
		toSerialize["api_keys"] = o.ApiKeys
	}
	return json.Marshal(toSerialize)
}

type NullableAuthTokenList struct {
	value *AuthTokenList
	isSet bool
}

func (v NullableAuthTokenList) Get() *AuthTokenList {
	return v.value
}

func (v *NullableAuthTokenList) Set(val *AuthTokenList) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthTokenList) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthTokenList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthTokenList(val *AuthTokenList) *NullableAuthTokenList {
	return &NullableAuthTokenList{value: val, isSet: true}
}

func (v NullableAuthTokenList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthTokenList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
