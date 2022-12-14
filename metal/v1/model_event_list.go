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

// EventList struct for EventList
type EventList struct {
	Events []FindInterconnectionEvents200Response `json:"events,omitempty"`
	Meta   *FindDeviceEvents200ResponseMeta       `json:"meta,omitempty"`
}

// NewEventList instantiates a new EventList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventList() *EventList {
	this := EventList{}
	return &this
}

// NewEventListWithDefaults instantiates a new EventList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventListWithDefaults() *EventList {
	this := EventList{}
	return &this
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *EventList) GetEvents() []FindInterconnectionEvents200Response {
	if o == nil || o.Events == nil {
		var ret []FindInterconnectionEvents200Response
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventList) GetEventsOk() ([]FindInterconnectionEvents200Response, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *EventList) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []FindInterconnectionEvents200Response and assigns it to the Events field.
func (o *EventList) SetEvents(v []FindInterconnectionEvents200Response) {
	o.Events = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *EventList) GetMeta() FindDeviceEvents200ResponseMeta {
	if o == nil || o.Meta == nil {
		var ret FindDeviceEvents200ResponseMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventList) GetMetaOk() (*FindDeviceEvents200ResponseMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *EventList) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given FindDeviceEvents200ResponseMeta and assigns it to the Meta field.
func (o *EventList) SetMeta(v FindDeviceEvents200ResponseMeta) {
	o.Meta = &v
}

func (o EventList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableEventList struct {
	value *EventList
	isSet bool
}

func (v NullableEventList) Get() *EventList {
	return v.value
}

func (v *NullableEventList) Set(val *EventList) {
	v.value = val
	v.isSet = true
}

func (v NullableEventList) IsSet() bool {
	return v.isSet
}

func (v *NullableEventList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventList(val *EventList) *NullableEventList {
	return &NullableEventList{value: val, isSet: true}
}

func (v NullableEventList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
