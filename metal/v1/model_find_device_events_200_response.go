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

// FindDeviceEvents200Response struct for FindDeviceEvents200Response
type FindDeviceEvents200Response struct {
	Events []FindInterconnectionEvents200Response `json:"events,omitempty"`
	Meta   *FindDeviceEvents200ResponseMeta       `json:"meta,omitempty"`
}

// NewFindDeviceEvents200Response instantiates a new FindDeviceEvents200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFindDeviceEvents200Response() *FindDeviceEvents200Response {
	this := FindDeviceEvents200Response{}
	return &this
}

// NewFindDeviceEvents200ResponseWithDefaults instantiates a new FindDeviceEvents200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFindDeviceEvents200ResponseWithDefaults() *FindDeviceEvents200Response {
	this := FindDeviceEvents200Response{}
	return &this
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *FindDeviceEvents200Response) GetEvents() []FindInterconnectionEvents200Response {
	if o == nil || o.Events == nil {
		var ret []FindInterconnectionEvents200Response
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindDeviceEvents200Response) GetEventsOk() ([]FindInterconnectionEvents200Response, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *FindDeviceEvents200Response) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []FindInterconnectionEvents200Response and assigns it to the Events field.
func (o *FindDeviceEvents200Response) SetEvents(v []FindInterconnectionEvents200Response) {
	o.Events = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *FindDeviceEvents200Response) GetMeta() FindDeviceEvents200ResponseMeta {
	if o == nil || o.Meta == nil {
		var ret FindDeviceEvents200ResponseMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindDeviceEvents200Response) GetMetaOk() (*FindDeviceEvents200ResponseMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *FindDeviceEvents200Response) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given FindDeviceEvents200ResponseMeta and assigns it to the Meta field.
func (o *FindDeviceEvents200Response) SetMeta(v FindDeviceEvents200ResponseMeta) {
	o.Meta = &v
}

func (o FindDeviceEvents200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableFindDeviceEvents200Response struct {
	value *FindDeviceEvents200Response
	isSet bool
}

func (v NullableFindDeviceEvents200Response) Get() *FindDeviceEvents200Response {
	return v.value
}

func (v *NullableFindDeviceEvents200Response) Set(val *FindDeviceEvents200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableFindDeviceEvents200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableFindDeviceEvents200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFindDeviceEvents200Response(val *FindDeviceEvents200Response) *NullableFindDeviceEvents200Response {
	return &NullableFindDeviceEvents200Response{value: val, isSet: true}
}

func (v NullableFindDeviceEvents200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFindDeviceEvents200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
