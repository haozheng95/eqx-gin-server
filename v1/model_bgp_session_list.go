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

// BgpSessionList struct for BgpSessionList
type BgpSessionList struct {
	BgpSessions []BgpSession `json:"bgp_sessions,omitempty"`
}

// NewBgpSessionList instantiates a new BgpSessionList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBgpSessionList() *BgpSessionList {
	this := BgpSessionList{}
	return &this
}

// NewBgpSessionListWithDefaults instantiates a new BgpSessionList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBgpSessionListWithDefaults() *BgpSessionList {
	this := BgpSessionList{}
	return &this
}

// GetBgpSessions returns the BgpSessions field value if set, zero value otherwise.
func (o *BgpSessionList) GetBgpSessions() []BgpSession {
	if o == nil || o.BgpSessions == nil {
		var ret []BgpSession
		return ret
	}
	return o.BgpSessions
}

// GetBgpSessionsOk returns a tuple with the BgpSessions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpSessionList) GetBgpSessionsOk() ([]BgpSession, bool) {
	if o == nil || o.BgpSessions == nil {
		return nil, false
	}
	return o.BgpSessions, true
}

// HasBgpSessions returns a boolean if a field has been set.
func (o *BgpSessionList) HasBgpSessions() bool {
	if o != nil && o.BgpSessions != nil {
		return true
	}

	return false
}

// SetBgpSessions gets a reference to the given []BgpSession and assigns it to the BgpSessions field.
func (o *BgpSessionList) SetBgpSessions(v []BgpSession) {
	o.BgpSessions = v
}

func (o BgpSessionList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BgpSessions != nil {
		toSerialize["bgp_sessions"] = o.BgpSessions
	}
	return json.Marshal(toSerialize)
}

type NullableBgpSessionList struct {
	value *BgpSessionList
	isSet bool
}

func (v NullableBgpSessionList) Get() *BgpSessionList {
	return v.value
}

func (v *NullableBgpSessionList) Set(val *BgpSessionList) {
	v.value = val
	v.isSet = true
}

func (v NullableBgpSessionList) IsSet() bool {
	return v.isSet
}

func (v *NullableBgpSessionList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBgpSessionList(val *BgpSessionList) *NullableBgpSessionList {
	return &NullableBgpSessionList{value: val, isSet: true}
}

func (v NullableBgpSessionList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBgpSessionList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


