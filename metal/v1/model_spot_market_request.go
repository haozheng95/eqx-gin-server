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
	"time"
)

// SpotMarketRequest struct for SpotMarketRequest
type SpotMarketRequest struct {
	CreatedAt   *time.Time                                                     `json:"created_at,omitempty"`
	DevicesMax  *int32                                                         `json:"devices_max,omitempty"`
	DevicesMin  *int32                                                         `json:"devices_min,omitempty"`
	EndAt       *time.Time                                                     `json:"end_at,omitempty"`
	Facilities  *FindBatchById200ResponseDevicesInner                          `json:"facilities,omitempty"`
	Href        *string                                                        `json:"href,omitempty"`
	Id          *string                                                        `json:"id,omitempty"`
	Instances   *FindBatchById200ResponseDevicesInner                          `json:"instances,omitempty"`
	MaxBidPrice *float32                                                       `json:"max_bid_price,omitempty"`
	Metro       *ListSpotMarketRequests200ResponseSpotMarketRequestsInnerMetro `json:"metro,omitempty"`
	Project     *FindBatchById200ResponseDevicesInner                          `json:"project,omitempty"`
}

// NewSpotMarketRequest instantiates a new SpotMarketRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpotMarketRequest() *SpotMarketRequest {
	this := SpotMarketRequest{}
	return &this
}

// NewSpotMarketRequestWithDefaults instantiates a new SpotMarketRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpotMarketRequestWithDefaults() *SpotMarketRequest {
	this := SpotMarketRequest{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *SpotMarketRequest) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotMarketRequest) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *SpotMarketRequest) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *SpotMarketRequest) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDevicesMax returns the DevicesMax field value if set, zero value otherwise.
func (o *SpotMarketRequest) GetDevicesMax() int32 {
	if o == nil || o.DevicesMax == nil {
		var ret int32
		return ret
	}
	return *o.DevicesMax
}

// GetDevicesMaxOk returns a tuple with the DevicesMax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotMarketRequest) GetDevicesMaxOk() (*int32, bool) {
	if o == nil || o.DevicesMax == nil {
		return nil, false
	}
	return o.DevicesMax, true
}

// HasDevicesMax returns a boolean if a field has been set.
func (o *SpotMarketRequest) HasDevicesMax() bool {
	if o != nil && o.DevicesMax != nil {
		return true
	}

	return false
}

// SetDevicesMax gets a reference to the given int32 and assigns it to the DevicesMax field.
func (o *SpotMarketRequest) SetDevicesMax(v int32) {
	o.DevicesMax = &v
}

// GetDevicesMin returns the DevicesMin field value if set, zero value otherwise.
func (o *SpotMarketRequest) GetDevicesMin() int32 {
	if o == nil || o.DevicesMin == nil {
		var ret int32
		return ret
	}
	return *o.DevicesMin
}

// GetDevicesMinOk returns a tuple with the DevicesMin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotMarketRequest) GetDevicesMinOk() (*int32, bool) {
	if o == nil || o.DevicesMin == nil {
		return nil, false
	}
	return o.DevicesMin, true
}

// HasDevicesMin returns a boolean if a field has been set.
func (o *SpotMarketRequest) HasDevicesMin() bool {
	if o != nil && o.DevicesMin != nil {
		return true
	}

	return false
}

// SetDevicesMin gets a reference to the given int32 and assigns it to the DevicesMin field.
func (o *SpotMarketRequest) SetDevicesMin(v int32) {
	o.DevicesMin = &v
}

// GetEndAt returns the EndAt field value if set, zero value otherwise.
func (o *SpotMarketRequest) GetEndAt() time.Time {
	if o == nil || o.EndAt == nil {
		var ret time.Time
		return ret
	}
	return *o.EndAt
}

// GetEndAtOk returns a tuple with the EndAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotMarketRequest) GetEndAtOk() (*time.Time, bool) {
	if o == nil || o.EndAt == nil {
		return nil, false
	}
	return o.EndAt, true
}

// HasEndAt returns a boolean if a field has been set.
func (o *SpotMarketRequest) HasEndAt() bool {
	if o != nil && o.EndAt != nil {
		return true
	}

	return false
}

// SetEndAt gets a reference to the given time.Time and assigns it to the EndAt field.
func (o *SpotMarketRequest) SetEndAt(v time.Time) {
	o.EndAt = &v
}

// GetFacilities returns the Facilities field value if set, zero value otherwise.
func (o *SpotMarketRequest) GetFacilities() FindBatchById200ResponseDevicesInner {
	if o == nil || o.Facilities == nil {
		var ret FindBatchById200ResponseDevicesInner
		return ret
	}
	return *o.Facilities
}

// GetFacilitiesOk returns a tuple with the Facilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotMarketRequest) GetFacilitiesOk() (*FindBatchById200ResponseDevicesInner, bool) {
	if o == nil || o.Facilities == nil {
		return nil, false
	}
	return o.Facilities, true
}

// HasFacilities returns a boolean if a field has been set.
func (o *SpotMarketRequest) HasFacilities() bool {
	if o != nil && o.Facilities != nil {
		return true
	}

	return false
}

// SetFacilities gets a reference to the given FindBatchById200ResponseDevicesInner and assigns it to the Facilities field.
func (o *SpotMarketRequest) SetFacilities(v FindBatchById200ResponseDevicesInner) {
	o.Facilities = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *SpotMarketRequest) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotMarketRequest) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *SpotMarketRequest) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *SpotMarketRequest) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SpotMarketRequest) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotMarketRequest) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SpotMarketRequest) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SpotMarketRequest) SetId(v string) {
	o.Id = &v
}

// GetInstances returns the Instances field value if set, zero value otherwise.
func (o *SpotMarketRequest) GetInstances() FindBatchById200ResponseDevicesInner {
	if o == nil || o.Instances == nil {
		var ret FindBatchById200ResponseDevicesInner
		return ret
	}
	return *o.Instances
}

// GetInstancesOk returns a tuple with the Instances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotMarketRequest) GetInstancesOk() (*FindBatchById200ResponseDevicesInner, bool) {
	if o == nil || o.Instances == nil {
		return nil, false
	}
	return o.Instances, true
}

// HasInstances returns a boolean if a field has been set.
func (o *SpotMarketRequest) HasInstances() bool {
	if o != nil && o.Instances != nil {
		return true
	}

	return false
}

// SetInstances gets a reference to the given FindBatchById200ResponseDevicesInner and assigns it to the Instances field.
func (o *SpotMarketRequest) SetInstances(v FindBatchById200ResponseDevicesInner) {
	o.Instances = &v
}

// GetMaxBidPrice returns the MaxBidPrice field value if set, zero value otherwise.
func (o *SpotMarketRequest) GetMaxBidPrice() float32 {
	if o == nil || o.MaxBidPrice == nil {
		var ret float32
		return ret
	}
	return *o.MaxBidPrice
}

// GetMaxBidPriceOk returns a tuple with the MaxBidPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotMarketRequest) GetMaxBidPriceOk() (*float32, bool) {
	if o == nil || o.MaxBidPrice == nil {
		return nil, false
	}
	return o.MaxBidPrice, true
}

// HasMaxBidPrice returns a boolean if a field has been set.
func (o *SpotMarketRequest) HasMaxBidPrice() bool {
	if o != nil && o.MaxBidPrice != nil {
		return true
	}

	return false
}

// SetMaxBidPrice gets a reference to the given float32 and assigns it to the MaxBidPrice field.
func (o *SpotMarketRequest) SetMaxBidPrice(v float32) {
	o.MaxBidPrice = &v
}

// GetMetro returns the Metro field value if set, zero value otherwise.
func (o *SpotMarketRequest) GetMetro() ListSpotMarketRequests200ResponseSpotMarketRequestsInnerMetro {
	if o == nil || o.Metro == nil {
		var ret ListSpotMarketRequests200ResponseSpotMarketRequestsInnerMetro
		return ret
	}
	return *o.Metro
}

// GetMetroOk returns a tuple with the Metro field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotMarketRequest) GetMetroOk() (*ListSpotMarketRequests200ResponseSpotMarketRequestsInnerMetro, bool) {
	if o == nil || o.Metro == nil {
		return nil, false
	}
	return o.Metro, true
}

// HasMetro returns a boolean if a field has been set.
func (o *SpotMarketRequest) HasMetro() bool {
	if o != nil && o.Metro != nil {
		return true
	}

	return false
}

// SetMetro gets a reference to the given ListSpotMarketRequests200ResponseSpotMarketRequestsInnerMetro and assigns it to the Metro field.
func (o *SpotMarketRequest) SetMetro(v ListSpotMarketRequests200ResponseSpotMarketRequestsInnerMetro) {
	o.Metro = &v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *SpotMarketRequest) GetProject() FindBatchById200ResponseDevicesInner {
	if o == nil || o.Project == nil {
		var ret FindBatchById200ResponseDevicesInner
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotMarketRequest) GetProjectOk() (*FindBatchById200ResponseDevicesInner, bool) {
	if o == nil || o.Project == nil {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *SpotMarketRequest) HasProject() bool {
	if o != nil && o.Project != nil {
		return true
	}

	return false
}

// SetProject gets a reference to the given FindBatchById200ResponseDevicesInner and assigns it to the Project field.
func (o *SpotMarketRequest) SetProject(v FindBatchById200ResponseDevicesInner) {
	o.Project = &v
}

func (o SpotMarketRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.DevicesMax != nil {
		toSerialize["devices_max"] = o.DevicesMax
	}
	if o.DevicesMin != nil {
		toSerialize["devices_min"] = o.DevicesMin
	}
	if o.EndAt != nil {
		toSerialize["end_at"] = o.EndAt
	}
	if o.Facilities != nil {
		toSerialize["facilities"] = o.Facilities
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Instances != nil {
		toSerialize["instances"] = o.Instances
	}
	if o.MaxBidPrice != nil {
		toSerialize["max_bid_price"] = o.MaxBidPrice
	}
	if o.Metro != nil {
		toSerialize["metro"] = o.Metro
	}
	if o.Project != nil {
		toSerialize["project"] = o.Project
	}
	return json.Marshal(toSerialize)
}

type NullableSpotMarketRequest struct {
	value *SpotMarketRequest
	isSet bool
}

func (v NullableSpotMarketRequest) Get() *SpotMarketRequest {
	return v.value
}

func (v *NullableSpotMarketRequest) Set(val *SpotMarketRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSpotMarketRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSpotMarketRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpotMarketRequest(val *SpotMarketRequest) *NullableSpotMarketRequest {
	return &NullableSpotMarketRequest{value: val, isSet: true}
}

func (v NullableSpotMarketRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpotMarketRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
