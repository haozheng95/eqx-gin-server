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

// FindMetroSpotMarketPrices200Response struct for FindMetroSpotMarketPrices200Response
type FindMetroSpotMarketPrices200Response struct {
	SpotMarketPrices *FindMetroSpotMarketPrices200ResponseSpotMarketPrices `json:"spot_market_prices,omitempty"`
}

// NewFindMetroSpotMarketPrices200Response instantiates a new FindMetroSpotMarketPrices200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFindMetroSpotMarketPrices200Response() *FindMetroSpotMarketPrices200Response {
	this := FindMetroSpotMarketPrices200Response{}
	return &this
}

// NewFindMetroSpotMarketPrices200ResponseWithDefaults instantiates a new FindMetroSpotMarketPrices200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFindMetroSpotMarketPrices200ResponseWithDefaults() *FindMetroSpotMarketPrices200Response {
	this := FindMetroSpotMarketPrices200Response{}
	return &this
}

// GetSpotMarketPrices returns the SpotMarketPrices field value if set, zero value otherwise.
func (o *FindMetroSpotMarketPrices200Response) GetSpotMarketPrices() FindMetroSpotMarketPrices200ResponseSpotMarketPrices {
	if o == nil || o.SpotMarketPrices == nil {
		var ret FindMetroSpotMarketPrices200ResponseSpotMarketPrices
		return ret
	}
	return *o.SpotMarketPrices
}

// GetSpotMarketPricesOk returns a tuple with the SpotMarketPrices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindMetroSpotMarketPrices200Response) GetSpotMarketPricesOk() (*FindMetroSpotMarketPrices200ResponseSpotMarketPrices, bool) {
	if o == nil || o.SpotMarketPrices == nil {
		return nil, false
	}
	return o.SpotMarketPrices, true
}

// HasSpotMarketPrices returns a boolean if a field has been set.
func (o *FindMetroSpotMarketPrices200Response) HasSpotMarketPrices() bool {
	if o != nil && o.SpotMarketPrices != nil {
		return true
	}

	return false
}

// SetSpotMarketPrices gets a reference to the given FindMetroSpotMarketPrices200ResponseSpotMarketPrices and assigns it to the SpotMarketPrices field.
func (o *FindMetroSpotMarketPrices200Response) SetSpotMarketPrices(v FindMetroSpotMarketPrices200ResponseSpotMarketPrices) {
	o.SpotMarketPrices = &v
}

func (o FindMetroSpotMarketPrices200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SpotMarketPrices != nil {
		toSerialize["spot_market_prices"] = o.SpotMarketPrices
	}
	return json.Marshal(toSerialize)
}

type NullableFindMetroSpotMarketPrices200Response struct {
	value *FindMetroSpotMarketPrices200Response
	isSet bool
}

func (v NullableFindMetroSpotMarketPrices200Response) Get() *FindMetroSpotMarketPrices200Response {
	return v.value
}

func (v *NullableFindMetroSpotMarketPrices200Response) Set(val *FindMetroSpotMarketPrices200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableFindMetroSpotMarketPrices200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableFindMetroSpotMarketPrices200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFindMetroSpotMarketPrices200Response(val *FindMetroSpotMarketPrices200Response) *NullableFindMetroSpotMarketPrices200Response {
	return &NullableFindMetroSpotMarketPrices200Response{value: val, isSet: true}
}

func (v NullableFindMetroSpotMarketPrices200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFindMetroSpotMarketPrices200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}