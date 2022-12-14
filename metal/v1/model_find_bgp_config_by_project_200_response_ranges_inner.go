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

// FindBgpConfigByProject200ResponseRangesInner struct for FindBgpConfigByProject200ResponseRangesInner
type FindBgpConfigByProject200ResponseRangesInner struct {
	AddressFamily *int32                                `json:"address_family,omitempty"`
	Href          *string                               `json:"href,omitempty"`
	Id            *string                               `json:"id,omitempty"`
	Project       *FindBatchById200ResponseDevicesInner `json:"project,omitempty"`
	Range         *string                               `json:"range,omitempty"`
}

// NewFindBgpConfigByProject200ResponseRangesInner instantiates a new FindBgpConfigByProject200ResponseRangesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFindBgpConfigByProject200ResponseRangesInner() *FindBgpConfigByProject200ResponseRangesInner {
	this := FindBgpConfigByProject200ResponseRangesInner{}
	return &this
}

// NewFindBgpConfigByProject200ResponseRangesInnerWithDefaults instantiates a new FindBgpConfigByProject200ResponseRangesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFindBgpConfigByProject200ResponseRangesInnerWithDefaults() *FindBgpConfigByProject200ResponseRangesInner {
	this := FindBgpConfigByProject200ResponseRangesInner{}
	return &this
}

// GetAddressFamily returns the AddressFamily field value if set, zero value otherwise.
func (o *FindBgpConfigByProject200ResponseRangesInner) GetAddressFamily() int32 {
	if o == nil || o.AddressFamily == nil {
		var ret int32
		return ret
	}
	return *o.AddressFamily
}

// GetAddressFamilyOk returns a tuple with the AddressFamily field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindBgpConfigByProject200ResponseRangesInner) GetAddressFamilyOk() (*int32, bool) {
	if o == nil || o.AddressFamily == nil {
		return nil, false
	}
	return o.AddressFamily, true
}

// HasAddressFamily returns a boolean if a field has been set.
func (o *FindBgpConfigByProject200ResponseRangesInner) HasAddressFamily() bool {
	if o != nil && o.AddressFamily != nil {
		return true
	}

	return false
}

// SetAddressFamily gets a reference to the given int32 and assigns it to the AddressFamily field.
func (o *FindBgpConfigByProject200ResponseRangesInner) SetAddressFamily(v int32) {
	o.AddressFamily = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *FindBgpConfigByProject200ResponseRangesInner) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindBgpConfigByProject200ResponseRangesInner) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *FindBgpConfigByProject200ResponseRangesInner) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *FindBgpConfigByProject200ResponseRangesInner) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FindBgpConfigByProject200ResponseRangesInner) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindBgpConfigByProject200ResponseRangesInner) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FindBgpConfigByProject200ResponseRangesInner) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *FindBgpConfigByProject200ResponseRangesInner) SetId(v string) {
	o.Id = &v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *FindBgpConfigByProject200ResponseRangesInner) GetProject() FindBatchById200ResponseDevicesInner {
	if o == nil || o.Project == nil {
		var ret FindBatchById200ResponseDevicesInner
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindBgpConfigByProject200ResponseRangesInner) GetProjectOk() (*FindBatchById200ResponseDevicesInner, bool) {
	if o == nil || o.Project == nil {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *FindBgpConfigByProject200ResponseRangesInner) HasProject() bool {
	if o != nil && o.Project != nil {
		return true
	}

	return false
}

// SetProject gets a reference to the given FindBatchById200ResponseDevicesInner and assigns it to the Project field.
func (o *FindBgpConfigByProject200ResponseRangesInner) SetProject(v FindBatchById200ResponseDevicesInner) {
	o.Project = &v
}

// GetRange returns the Range field value if set, zero value otherwise.
func (o *FindBgpConfigByProject200ResponseRangesInner) GetRange() string {
	if o == nil || o.Range == nil {
		var ret string
		return ret
	}
	return *o.Range
}

// GetRangeOk returns a tuple with the Range field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindBgpConfigByProject200ResponseRangesInner) GetRangeOk() (*string, bool) {
	if o == nil || o.Range == nil {
		return nil, false
	}
	return o.Range, true
}

// HasRange returns a boolean if a field has been set.
func (o *FindBgpConfigByProject200ResponseRangesInner) HasRange() bool {
	if o != nil && o.Range != nil {
		return true
	}

	return false
}

// SetRange gets a reference to the given string and assigns it to the Range field.
func (o *FindBgpConfigByProject200ResponseRangesInner) SetRange(v string) {
	o.Range = &v
}

func (o FindBgpConfigByProject200ResponseRangesInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AddressFamily != nil {
		toSerialize["address_family"] = o.AddressFamily
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Project != nil {
		toSerialize["project"] = o.Project
	}
	if o.Range != nil {
		toSerialize["range"] = o.Range
	}
	return json.Marshal(toSerialize)
}

type NullableFindBgpConfigByProject200ResponseRangesInner struct {
	value *FindBgpConfigByProject200ResponseRangesInner
	isSet bool
}

func (v NullableFindBgpConfigByProject200ResponseRangesInner) Get() *FindBgpConfigByProject200ResponseRangesInner {
	return v.value
}

func (v *NullableFindBgpConfigByProject200ResponseRangesInner) Set(val *FindBgpConfigByProject200ResponseRangesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableFindBgpConfigByProject200ResponseRangesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableFindBgpConfigByProject200ResponseRangesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFindBgpConfigByProject200ResponseRangesInner(val *FindBgpConfigByProject200ResponseRangesInner) *NullableFindBgpConfigByProject200ResponseRangesInner {
	return &NullableFindBgpConfigByProject200ResponseRangesInner{value: val, isSet: true}
}

func (v NullableFindBgpConfigByProject200ResponseRangesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFindBgpConfigByProject200ResponseRangesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
