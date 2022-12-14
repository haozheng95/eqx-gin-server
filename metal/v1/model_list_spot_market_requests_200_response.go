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

// ListSpotMarketRequests200Response struct for ListSpotMarketRequests200Response
type ListSpotMarketRequests200Response struct {
	SpotMarketRequests []ListSpotMarketRequests200ResponseSpotMarketRequestsInner `json:"spot_market_requests,omitempty"`
}

// NewListSpotMarketRequests200Response instantiates a new ListSpotMarketRequests200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListSpotMarketRequests200Response() *ListSpotMarketRequests200Response {
	this := ListSpotMarketRequests200Response{}
	return &this
}

// NewListSpotMarketRequests200ResponseWithDefaults instantiates a new ListSpotMarketRequests200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListSpotMarketRequests200ResponseWithDefaults() *ListSpotMarketRequests200Response {
	this := ListSpotMarketRequests200Response{}
	return &this
}

// GetSpotMarketRequests returns the SpotMarketRequests field value if set, zero value otherwise.
func (o *ListSpotMarketRequests200Response) GetSpotMarketRequests() []ListSpotMarketRequests200ResponseSpotMarketRequestsInner {
	if o == nil || o.SpotMarketRequests == nil {
		var ret []ListSpotMarketRequests200ResponseSpotMarketRequestsInner
		return ret
	}
	return o.SpotMarketRequests
}

// GetSpotMarketRequestsOk returns a tuple with the SpotMarketRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSpotMarketRequests200Response) GetSpotMarketRequestsOk() ([]ListSpotMarketRequests200ResponseSpotMarketRequestsInner, bool) {
	if o == nil || o.SpotMarketRequests == nil {
		return nil, false
	}
	return o.SpotMarketRequests, true
}

// HasSpotMarketRequests returns a boolean if a field has been set.
func (o *ListSpotMarketRequests200Response) HasSpotMarketRequests() bool {
	if o != nil && o.SpotMarketRequests != nil {
		return true
	}

	return false
}

// SetSpotMarketRequests gets a reference to the given []ListSpotMarketRequests200ResponseSpotMarketRequestsInner and assigns it to the SpotMarketRequests field.
func (o *ListSpotMarketRequests200Response) SetSpotMarketRequests(v []ListSpotMarketRequests200ResponseSpotMarketRequestsInner) {
	o.SpotMarketRequests = v
}

func (o ListSpotMarketRequests200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SpotMarketRequests != nil {
		toSerialize["spot_market_requests"] = o.SpotMarketRequests
	}
	return json.Marshal(toSerialize)
}

type NullableListSpotMarketRequests200Response struct {
	value *ListSpotMarketRequests200Response
	isSet bool
}

func (v NullableListSpotMarketRequests200Response) Get() *ListSpotMarketRequests200Response {
	return v.value
}

func (v *NullableListSpotMarketRequests200Response) Set(val *ListSpotMarketRequests200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListSpotMarketRequests200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListSpotMarketRequests200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListSpotMarketRequests200Response(val *ListSpotMarketRequests200Response) *NullableListSpotMarketRequests200Response {
	return &NullableListSpotMarketRequests200Response{value: val, isSet: true}
}

func (v NullableListSpotMarketRequests200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListSpotMarketRequests200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
