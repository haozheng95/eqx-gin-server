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

// BatchesList struct for BatchesList
type BatchesList struct {
	Batches []FindBatchById200Response `json:"batches,omitempty"`
}

// NewBatchesList instantiates a new BatchesList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchesList() *BatchesList {
	this := BatchesList{}
	return &this
}

// NewBatchesListWithDefaults instantiates a new BatchesList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchesListWithDefaults() *BatchesList {
	this := BatchesList{}
	return &this
}

// GetBatches returns the Batches field value if set, zero value otherwise.
func (o *BatchesList) GetBatches() []FindBatchById200Response {
	if o == nil || o.Batches == nil {
		var ret []FindBatchById200Response
		return ret
	}
	return o.Batches
}

// GetBatchesOk returns a tuple with the Batches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchesList) GetBatchesOk() ([]FindBatchById200Response, bool) {
	if o == nil || o.Batches == nil {
		return nil, false
	}
	return o.Batches, true
}

// HasBatches returns a boolean if a field has been set.
func (o *BatchesList) HasBatches() bool {
	if o != nil && o.Batches != nil {
		return true
	}

	return false
}

// SetBatches gets a reference to the given []FindBatchById200Response and assigns it to the Batches field.
func (o *BatchesList) SetBatches(v []FindBatchById200Response) {
	o.Batches = v
}

func (o BatchesList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Batches != nil {
		toSerialize["batches"] = o.Batches
	}
	return json.Marshal(toSerialize)
}

type NullableBatchesList struct {
	value *BatchesList
	isSet bool
}

func (v NullableBatchesList) Get() *BatchesList {
	return v.value
}

func (v *NullableBatchesList) Set(val *BatchesList) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchesList) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchesList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchesList(val *BatchesList) *NullableBatchesList {
	return &NullableBatchesList{value: val, isSet: true}
}

func (v NullableBatchesList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchesList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
