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

// FindDeviceById200ResponsePlanAvailableInInner struct for FindDeviceById200ResponsePlanAvailableInInner
type FindDeviceById200ResponsePlanAvailableInInner struct {
	// href to the Facility
	Href  *string                                             `json:"href,omitempty"`
	Price *FindDeviceById200ResponsePlanAvailableInInnerPrice `json:"price,omitempty"`
}

// NewFindDeviceById200ResponsePlanAvailableInInner instantiates a new FindDeviceById200ResponsePlanAvailableInInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFindDeviceById200ResponsePlanAvailableInInner() *FindDeviceById200ResponsePlanAvailableInInner {
	this := FindDeviceById200ResponsePlanAvailableInInner{}
	return &this
}

// NewFindDeviceById200ResponsePlanAvailableInInnerWithDefaults instantiates a new FindDeviceById200ResponsePlanAvailableInInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFindDeviceById200ResponsePlanAvailableInInnerWithDefaults() *FindDeviceById200ResponsePlanAvailableInInner {
	this := FindDeviceById200ResponsePlanAvailableInInner{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *FindDeviceById200ResponsePlanAvailableInInner) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindDeviceById200ResponsePlanAvailableInInner) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *FindDeviceById200ResponsePlanAvailableInInner) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *FindDeviceById200ResponsePlanAvailableInInner) SetHref(v string) {
	o.Href = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *FindDeviceById200ResponsePlanAvailableInInner) GetPrice() FindDeviceById200ResponsePlanAvailableInInnerPrice {
	if o == nil || o.Price == nil {
		var ret FindDeviceById200ResponsePlanAvailableInInnerPrice
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindDeviceById200ResponsePlanAvailableInInner) GetPriceOk() (*FindDeviceById200ResponsePlanAvailableInInnerPrice, bool) {
	if o == nil || o.Price == nil {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *FindDeviceById200ResponsePlanAvailableInInner) HasPrice() bool {
	if o != nil && o.Price != nil {
		return true
	}

	return false
}

// SetPrice gets a reference to the given FindDeviceById200ResponsePlanAvailableInInnerPrice and assigns it to the Price field.
func (o *FindDeviceById200ResponsePlanAvailableInInner) SetPrice(v FindDeviceById200ResponsePlanAvailableInInnerPrice) {
	o.Price = &v
}

func (o FindDeviceById200ResponsePlanAvailableInInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Price != nil {
		toSerialize["price"] = o.Price
	}
	return json.Marshal(toSerialize)
}

type NullableFindDeviceById200ResponsePlanAvailableInInner struct {
	value *FindDeviceById200ResponsePlanAvailableInInner
	isSet bool
}

func (v NullableFindDeviceById200ResponsePlanAvailableInInner) Get() *FindDeviceById200ResponsePlanAvailableInInner {
	return v.value
}

func (v *NullableFindDeviceById200ResponsePlanAvailableInInner) Set(val *FindDeviceById200ResponsePlanAvailableInInner) {
	v.value = val
	v.isSet = true
}

func (v NullableFindDeviceById200ResponsePlanAvailableInInner) IsSet() bool {
	return v.isSet
}

func (v *NullableFindDeviceById200ResponsePlanAvailableInInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFindDeviceById200ResponsePlanAvailableInInner(val *FindDeviceById200ResponsePlanAvailableInInner) *NullableFindDeviceById200ResponsePlanAvailableInInner {
	return &NullableFindDeviceById200ResponsePlanAvailableInInner{value: val, isSet: true}
}

func (v NullableFindDeviceById200ResponsePlanAvailableInInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFindDeviceById200ResponsePlanAvailableInInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}