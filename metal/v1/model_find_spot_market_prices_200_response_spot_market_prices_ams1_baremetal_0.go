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

// FindSpotMarketPrices200ResponseSpotMarketPricesAms1Baremetal0 struct for FindSpotMarketPrices200ResponseSpotMarketPricesAms1Baremetal0
type FindSpotMarketPrices200ResponseSpotMarketPricesAms1Baremetal0 struct {
	Price *float32 `json:"price,omitempty"`
}

// NewFindSpotMarketPrices200ResponseSpotMarketPricesAms1Baremetal0 instantiates a new FindSpotMarketPrices200ResponseSpotMarketPricesAms1Baremetal0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFindSpotMarketPrices200ResponseSpotMarketPricesAms1Baremetal0() *FindSpotMarketPrices200ResponseSpotMarketPricesAms1Baremetal0 {
	this := FindSpotMarketPrices200ResponseSpotMarketPricesAms1Baremetal0{}
	return &this
}

// NewFindSpotMarketPrices200ResponseSpotMarketPricesAms1Baremetal0WithDefaults instantiates a new FindSpotMarketPrices200ResponseSpotMarketPricesAms1Baremetal0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFindSpotMarketPrices200ResponseSpotMarketPricesAms1Baremetal0WithDefaults() *FindSpotMarketPrices200ResponseSpotMarketPricesAms1Baremetal0 {
	this := FindSpotMarketPrices200ResponseSpotMarketPricesAms1Baremetal0{}
	return &this
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *FindSpotMarketPrices200ResponseSpotMarketPricesAms1Baremetal0) GetPrice() float32 {
	if o == nil || o.Price == nil {
		var ret float32
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindSpotMarketPrices200ResponseSpotMarketPricesAms1Baremetal0) GetPriceOk() (*float32, bool) {
	if o == nil || o.Price == nil {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *FindSpotMarketPrices200ResponseSpotMarketPricesAms1Baremetal0) HasPrice() bool {
	if o != nil && o.Price != nil {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float32 and assigns it to the Price field.
func (o *FindSpotMarketPrices200ResponseSpotMarketPricesAms1Baremetal0) SetPrice(v float32) {
	o.Price = &v
}

func (o FindSpotMarketPrices200ResponseSpotMarketPricesAms1Baremetal0) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Price != nil {
		toSerialize["price"] = o.Price
	}
	return json.Marshal(toSerialize)
}

type NullableFindSpotMarketPrices200ResponseSpotMarketPricesAms1Baremetal0 struct {
	value *FindSpotMarketPrices200ResponseSpotMarketPricesAms1Baremetal0
	isSet bool
}

func (v NullableFindSpotMarketPrices200ResponseSpotMarketPricesAms1Baremetal0) Get() *FindSpotMarketPrices200ResponseSpotMarketPricesAms1Baremetal0 {
	return v.value
}

func (v *NullableFindSpotMarketPrices200ResponseSpotMarketPricesAms1Baremetal0) Set(val *FindSpotMarketPrices200ResponseSpotMarketPricesAms1Baremetal0) {
	v.value = val
	v.isSet = true
}

func (v NullableFindSpotMarketPrices200ResponseSpotMarketPricesAms1Baremetal0) IsSet() bool {
	return v.isSet
}

func (v *NullableFindSpotMarketPrices200ResponseSpotMarketPricesAms1Baremetal0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFindSpotMarketPrices200ResponseSpotMarketPricesAms1Baremetal0(val *FindSpotMarketPrices200ResponseSpotMarketPricesAms1Baremetal0) *NullableFindSpotMarketPrices200ResponseSpotMarketPricesAms1Baremetal0 {
	return &NullableFindSpotMarketPrices200ResponseSpotMarketPricesAms1Baremetal0{value: val, isSet: true}
}

func (v NullableFindSpotMarketPrices200ResponseSpotMarketPricesAms1Baremetal0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFindSpotMarketPrices200ResponseSpotMarketPricesAms1Baremetal0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
