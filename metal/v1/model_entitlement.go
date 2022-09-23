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

// Entitlement struct for Entitlement
type Entitlement struct {
	Description   *string                `json:"description,omitempty"`
	FeatureAccess map[string]interface{} `json:"feature_access,omitempty"`
	Href          *string                `json:"href,omitempty"`
	Id            string                 `json:"id"`
	InstanceQuota map[string]interface{} `json:"instance_quota,omitempty"`
	IpQuota       map[string]interface{} `json:"ip_quota,omitempty"`
	Name          *string                `json:"name,omitempty"`
	ProjectQuota  *int32                 `json:"project_quota,omitempty"`
	Slug          string                 `json:"slug"`
	VolumeLimits  map[string]interface{} `json:"volume_limits,omitempty"`
	VolumeQuota   map[string]interface{} `json:"volume_quota,omitempty"`
	Weight        int32                  `json:"weight"`
}

// NewEntitlement instantiates a new Entitlement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlement(id string, slug string, weight int32) *Entitlement {
	this := Entitlement{}
	this.Id = id
	var projectQuota int32 = 0
	this.ProjectQuota = &projectQuota
	this.Slug = slug
	this.Weight = weight
	return &this
}

// NewEntitlementWithDefaults instantiates a new Entitlement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementWithDefaults() *Entitlement {
	this := Entitlement{}
	var projectQuota int32 = 0
	this.ProjectQuota = &projectQuota
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Entitlement) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Entitlement) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Entitlement) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Entitlement) SetDescription(v string) {
	o.Description = &v
}

// GetFeatureAccess returns the FeatureAccess field value if set, zero value otherwise.
func (o *Entitlement) GetFeatureAccess() map[string]interface{} {
	if o == nil || o.FeatureAccess == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.FeatureAccess
}

// GetFeatureAccessOk returns a tuple with the FeatureAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Entitlement) GetFeatureAccessOk() (map[string]interface{}, bool) {
	if o == nil || o.FeatureAccess == nil {
		return nil, false
	}
	return o.FeatureAccess, true
}

// HasFeatureAccess returns a boolean if a field has been set.
func (o *Entitlement) HasFeatureAccess() bool {
	if o != nil && o.FeatureAccess != nil {
		return true
	}

	return false
}

// SetFeatureAccess gets a reference to the given map[string]interface{} and assigns it to the FeatureAccess field.
func (o *Entitlement) SetFeatureAccess(v map[string]interface{}) {
	o.FeatureAccess = v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *Entitlement) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Entitlement) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *Entitlement) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *Entitlement) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value
func (o *Entitlement) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Entitlement) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Entitlement) SetId(v string) {
	o.Id = v
}

// GetInstanceQuota returns the InstanceQuota field value if set, zero value otherwise.
func (o *Entitlement) GetInstanceQuota() map[string]interface{} {
	if o == nil || o.InstanceQuota == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.InstanceQuota
}

// GetInstanceQuotaOk returns a tuple with the InstanceQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Entitlement) GetInstanceQuotaOk() (map[string]interface{}, bool) {
	if o == nil || o.InstanceQuota == nil {
		return nil, false
	}
	return o.InstanceQuota, true
}

// HasInstanceQuota returns a boolean if a field has been set.
func (o *Entitlement) HasInstanceQuota() bool {
	if o != nil && o.InstanceQuota != nil {
		return true
	}

	return false
}

// SetInstanceQuota gets a reference to the given map[string]interface{} and assigns it to the InstanceQuota field.
func (o *Entitlement) SetInstanceQuota(v map[string]interface{}) {
	o.InstanceQuota = v
}

// GetIpQuota returns the IpQuota field value if set, zero value otherwise.
func (o *Entitlement) GetIpQuota() map[string]interface{} {
	if o == nil || o.IpQuota == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.IpQuota
}

// GetIpQuotaOk returns a tuple with the IpQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Entitlement) GetIpQuotaOk() (map[string]interface{}, bool) {
	if o == nil || o.IpQuota == nil {
		return nil, false
	}
	return o.IpQuota, true
}

// HasIpQuota returns a boolean if a field has been set.
func (o *Entitlement) HasIpQuota() bool {
	if o != nil && o.IpQuota != nil {
		return true
	}

	return false
}

// SetIpQuota gets a reference to the given map[string]interface{} and assigns it to the IpQuota field.
func (o *Entitlement) SetIpQuota(v map[string]interface{}) {
	o.IpQuota = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Entitlement) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Entitlement) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Entitlement) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Entitlement) SetName(v string) {
	o.Name = &v
}

// GetProjectQuota returns the ProjectQuota field value if set, zero value otherwise.
func (o *Entitlement) GetProjectQuota() int32 {
	if o == nil || o.ProjectQuota == nil {
		var ret int32
		return ret
	}
	return *o.ProjectQuota
}

// GetProjectQuotaOk returns a tuple with the ProjectQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Entitlement) GetProjectQuotaOk() (*int32, bool) {
	if o == nil || o.ProjectQuota == nil {
		return nil, false
	}
	return o.ProjectQuota, true
}

// HasProjectQuota returns a boolean if a field has been set.
func (o *Entitlement) HasProjectQuota() bool {
	if o != nil && o.ProjectQuota != nil {
		return true
	}

	return false
}

// SetProjectQuota gets a reference to the given int32 and assigns it to the ProjectQuota field.
func (o *Entitlement) SetProjectQuota(v int32) {
	o.ProjectQuota = &v
}

// GetSlug returns the Slug field value
func (o *Entitlement) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *Entitlement) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *Entitlement) SetSlug(v string) {
	o.Slug = v
}

// GetVolumeLimits returns the VolumeLimits field value if set, zero value otherwise.
func (o *Entitlement) GetVolumeLimits() map[string]interface{} {
	if o == nil || o.VolumeLimits == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.VolumeLimits
}

// GetVolumeLimitsOk returns a tuple with the VolumeLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Entitlement) GetVolumeLimitsOk() (map[string]interface{}, bool) {
	if o == nil || o.VolumeLimits == nil {
		return nil, false
	}
	return o.VolumeLimits, true
}

// HasVolumeLimits returns a boolean if a field has been set.
func (o *Entitlement) HasVolumeLimits() bool {
	if o != nil && o.VolumeLimits != nil {
		return true
	}

	return false
}

// SetVolumeLimits gets a reference to the given map[string]interface{} and assigns it to the VolumeLimits field.
func (o *Entitlement) SetVolumeLimits(v map[string]interface{}) {
	o.VolumeLimits = v
}

// GetVolumeQuota returns the VolumeQuota field value if set, zero value otherwise.
func (o *Entitlement) GetVolumeQuota() map[string]interface{} {
	if o == nil || o.VolumeQuota == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.VolumeQuota
}

// GetVolumeQuotaOk returns a tuple with the VolumeQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Entitlement) GetVolumeQuotaOk() (map[string]interface{}, bool) {
	if o == nil || o.VolumeQuota == nil {
		return nil, false
	}
	return o.VolumeQuota, true
}

// HasVolumeQuota returns a boolean if a field has been set.
func (o *Entitlement) HasVolumeQuota() bool {
	if o != nil && o.VolumeQuota != nil {
		return true
	}

	return false
}

// SetVolumeQuota gets a reference to the given map[string]interface{} and assigns it to the VolumeQuota field.
func (o *Entitlement) SetVolumeQuota(v map[string]interface{}) {
	o.VolumeQuota = v
}

// GetWeight returns the Weight field value
func (o *Entitlement) GetWeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Weight
}

// GetWeightOk returns a tuple with the Weight field value
// and a boolean to check if the value has been set.
func (o *Entitlement) GetWeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Weight, true
}

// SetWeight sets field value
func (o *Entitlement) SetWeight(v int32) {
	o.Weight = v
}

func (o Entitlement) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.FeatureAccess != nil {
		toSerialize["feature_access"] = o.FeatureAccess
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.InstanceQuota != nil {
		toSerialize["instance_quota"] = o.InstanceQuota
	}
	if o.IpQuota != nil {
		toSerialize["ip_quota"] = o.IpQuota
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ProjectQuota != nil {
		toSerialize["project_quota"] = o.ProjectQuota
	}
	if true {
		toSerialize["slug"] = o.Slug
	}
	if o.VolumeLimits != nil {
		toSerialize["volume_limits"] = o.VolumeLimits
	}
	if o.VolumeQuota != nil {
		toSerialize["volume_quota"] = o.VolumeQuota
	}
	if true {
		toSerialize["weight"] = o.Weight
	}
	return json.Marshal(toSerialize)
}

type NullableEntitlement struct {
	value *Entitlement
	isSet bool
}

func (v NullableEntitlement) Get() *Entitlement {
	return v.value
}

func (v *NullableEntitlement) Set(val *Entitlement) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlement) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlement(val *Entitlement) *NullableEntitlement {
	return &NullableEntitlement{value: val, isSet: true}
}

func (v NullableEntitlement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
