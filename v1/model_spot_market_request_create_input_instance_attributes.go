/*
 * Metal API
 *
 * This is the API for Equinix Metal. The API allows you to programmatically interact with all of your Equinix Metal resources, including devices, networks, addresses, organizations, projects, and your user account.  The official API docs are hosted at <https://metal.equinix.com/developers/api>. 
 *
 * API version: 1.0.0
 * Contact: support@equinixmetal.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
	"time"
)

// SpotMarketRequestCreateInputInstanceAttributes struct for SpotMarketRequestCreateInputInstanceAttributes
type SpotMarketRequestCreateInputInstanceAttributes struct {
	Plan *string `json:"plan,omitempty"`
	Hostname *string `json:"hostname,omitempty"`
	Hostnames *[]string `json:"hostnames,omitempty"`
	Description *string `json:"description,omitempty"`
	BillingCycle *string `json:"billing_cycle,omitempty"`
	OperatingSystem *string `json:"operating_system,omitempty"`
	AlwaysPxe *bool `json:"always_pxe,omitempty"`
	Userdata *string `json:"userdata,omitempty"`
	Locked *bool `json:"locked,omitempty"`
	TerminationTime *time.Time `json:"termination_time,omitempty"`
	Tags *[]string `json:"tags,omitempty"`
	ProjectSshKeys *[]string `json:"project_ssh_keys,omitempty"`
	// The UUIDs of users whose SSH keys should be included on the provisioned device.
	UserSshKeys *[]string `json:"user_ssh_keys,omitempty"`
	NoSshKeys *bool `json:"no_ssh_keys,omitempty"`
	Features *[]string `json:"features,omitempty"`
	Customdata *map[string]interface{} `json:"customdata,omitempty"`
	PublicIpv4SubnetSize *int32 `json:"public_ipv4_subnet_size,omitempty"`
	PrivateIpv4SubnetSize *int32 `json:"private_ipv4_subnet_size,omitempty"`
}

// NewSpotMarketRequestCreateInputInstanceAttributes instantiates a new SpotMarketRequestCreateInputInstanceAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpotMarketRequestCreateInputInstanceAttributes() *SpotMarketRequestCreateInputInstanceAttributes {
	this := SpotMarketRequestCreateInputInstanceAttributes{}
	return &this
}

// NewSpotMarketRequestCreateInputInstanceAttributesWithDefaults instantiates a new SpotMarketRequestCreateInputInstanceAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpotMarketRequestCreateInputInstanceAttributesWithDefaults() *SpotMarketRequestCreateInputInstanceAttributes {
	this := SpotMarketRequestCreateInputInstanceAttributes{}
	return &this
}

// GetPlan returns the Plan field value if set, zero value otherwise.
func (o *SpotMarketRequestCreateInputInstanceAttributes) GetPlan() string {
	if o == nil || o.Plan == nil {
		var ret string
		return ret
	}
	return *o.Plan
}

// GetPlanOk returns a tuple with the Plan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotMarketRequestCreateInputInstanceAttributes) GetPlanOk() (*string, bool) {
	if o == nil || o.Plan == nil {
		return nil, false
	}
	return o.Plan, true
}

// HasPlan returns a boolean if a field has been set.
func (o *SpotMarketRequestCreateInputInstanceAttributes) HasPlan() bool {
	if o != nil && o.Plan != nil {
		return true
	}

	return false
}

// SetPlan gets a reference to the given string and assigns it to the Plan field.
func (o *SpotMarketRequestCreateInputInstanceAttributes) SetPlan(v string) {
	o.Plan = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *SpotMarketRequestCreateInputInstanceAttributes) GetHostname() string {
	if o == nil || o.Hostname == nil {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotMarketRequestCreateInputInstanceAttributes) GetHostnameOk() (*string, bool) {
	if o == nil || o.Hostname == nil {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *SpotMarketRequestCreateInputInstanceAttributes) HasHostname() bool {
	if o != nil && o.Hostname != nil {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *SpotMarketRequestCreateInputInstanceAttributes) SetHostname(v string) {
	o.Hostname = &v
}

// GetHostnames returns the Hostnames field value if set, zero value otherwise.
func (o *SpotMarketRequestCreateInputInstanceAttributes) GetHostnames() []string {
	if o == nil || o.Hostnames == nil {
		var ret []string
		return ret
	}
	return *o.Hostnames
}

// GetHostnamesOk returns a tuple with the Hostnames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotMarketRequestCreateInputInstanceAttributes) GetHostnamesOk() (*[]string, bool) {
	if o == nil || o.Hostnames == nil {
		return nil, false
	}
	return o.Hostnames, true
}

// HasHostnames returns a boolean if a field has been set.
func (o *SpotMarketRequestCreateInputInstanceAttributes) HasHostnames() bool {
	if o != nil && o.Hostnames != nil {
		return true
	}

	return false
}

// SetHostnames gets a reference to the given []string and assigns it to the Hostnames field.
func (o *SpotMarketRequestCreateInputInstanceAttributes) SetHostnames(v []string) {
	o.Hostnames = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SpotMarketRequestCreateInputInstanceAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotMarketRequestCreateInputInstanceAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SpotMarketRequestCreateInputInstanceAttributes) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SpotMarketRequestCreateInputInstanceAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetBillingCycle returns the BillingCycle field value if set, zero value otherwise.
func (o *SpotMarketRequestCreateInputInstanceAttributes) GetBillingCycle() string {
	if o == nil || o.BillingCycle == nil {
		var ret string
		return ret
	}
	return *o.BillingCycle
}

// GetBillingCycleOk returns a tuple with the BillingCycle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotMarketRequestCreateInputInstanceAttributes) GetBillingCycleOk() (*string, bool) {
	if o == nil || o.BillingCycle == nil {
		return nil, false
	}
	return o.BillingCycle, true
}

// HasBillingCycle returns a boolean if a field has been set.
func (o *SpotMarketRequestCreateInputInstanceAttributes) HasBillingCycle() bool {
	if o != nil && o.BillingCycle != nil {
		return true
	}

	return false
}

// SetBillingCycle gets a reference to the given string and assigns it to the BillingCycle field.
func (o *SpotMarketRequestCreateInputInstanceAttributes) SetBillingCycle(v string) {
	o.BillingCycle = &v
}

// GetOperatingSystem returns the OperatingSystem field value if set, zero value otherwise.
func (o *SpotMarketRequestCreateInputInstanceAttributes) GetOperatingSystem() string {
	if o == nil || o.OperatingSystem == nil {
		var ret string
		return ret
	}
	return *o.OperatingSystem
}

// GetOperatingSystemOk returns a tuple with the OperatingSystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotMarketRequestCreateInputInstanceAttributes) GetOperatingSystemOk() (*string, bool) {
	if o == nil || o.OperatingSystem == nil {
		return nil, false
	}
	return o.OperatingSystem, true
}

// HasOperatingSystem returns a boolean if a field has been set.
func (o *SpotMarketRequestCreateInputInstanceAttributes) HasOperatingSystem() bool {
	if o != nil && o.OperatingSystem != nil {
		return true
	}

	return false
}

// SetOperatingSystem gets a reference to the given string and assigns it to the OperatingSystem field.
func (o *SpotMarketRequestCreateInputInstanceAttributes) SetOperatingSystem(v string) {
	o.OperatingSystem = &v
}

// GetAlwaysPxe returns the AlwaysPxe field value if set, zero value otherwise.
func (o *SpotMarketRequestCreateInputInstanceAttributes) GetAlwaysPxe() bool {
	if o == nil || o.AlwaysPxe == nil {
		var ret bool
		return ret
	}
	return *o.AlwaysPxe
}

// GetAlwaysPxeOk returns a tuple with the AlwaysPxe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotMarketRequestCreateInputInstanceAttributes) GetAlwaysPxeOk() (*bool, bool) {
	if o == nil || o.AlwaysPxe == nil {
		return nil, false
	}
	return o.AlwaysPxe, true
}

// HasAlwaysPxe returns a boolean if a field has been set.
func (o *SpotMarketRequestCreateInputInstanceAttributes) HasAlwaysPxe() bool {
	if o != nil && o.AlwaysPxe != nil {
		return true
	}

	return false
}

// SetAlwaysPxe gets a reference to the given bool and assigns it to the AlwaysPxe field.
func (o *SpotMarketRequestCreateInputInstanceAttributes) SetAlwaysPxe(v bool) {
	o.AlwaysPxe = &v
}

// GetUserdata returns the Userdata field value if set, zero value otherwise.
func (o *SpotMarketRequestCreateInputInstanceAttributes) GetUserdata() string {
	if o == nil || o.Userdata == nil {
		var ret string
		return ret
	}
	return *o.Userdata
}

// GetUserdataOk returns a tuple with the Userdata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotMarketRequestCreateInputInstanceAttributes) GetUserdataOk() (*string, bool) {
	if o == nil || o.Userdata == nil {
		return nil, false
	}
	return o.Userdata, true
}

// HasUserdata returns a boolean if a field has been set.
func (o *SpotMarketRequestCreateInputInstanceAttributes) HasUserdata() bool {
	if o != nil && o.Userdata != nil {
		return true
	}

	return false
}

// SetUserdata gets a reference to the given string and assigns it to the Userdata field.
func (o *SpotMarketRequestCreateInputInstanceAttributes) SetUserdata(v string) {
	o.Userdata = &v
}

// GetLocked returns the Locked field value if set, zero value otherwise.
func (o *SpotMarketRequestCreateInputInstanceAttributes) GetLocked() bool {
	if o == nil || o.Locked == nil {
		var ret bool
		return ret
	}
	return *o.Locked
}

// GetLockedOk returns a tuple with the Locked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotMarketRequestCreateInputInstanceAttributes) GetLockedOk() (*bool, bool) {
	if o == nil || o.Locked == nil {
		return nil, false
	}
	return o.Locked, true
}

// HasLocked returns a boolean if a field has been set.
func (o *SpotMarketRequestCreateInputInstanceAttributes) HasLocked() bool {
	if o != nil && o.Locked != nil {
		return true
	}

	return false
}

// SetLocked gets a reference to the given bool and assigns it to the Locked field.
func (o *SpotMarketRequestCreateInputInstanceAttributes) SetLocked(v bool) {
	o.Locked = &v
}

// GetTerminationTime returns the TerminationTime field value if set, zero value otherwise.
func (o *SpotMarketRequestCreateInputInstanceAttributes) GetTerminationTime() time.Time {
	if o == nil || o.TerminationTime == nil {
		var ret time.Time
		return ret
	}
	return *o.TerminationTime
}

// GetTerminationTimeOk returns a tuple with the TerminationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotMarketRequestCreateInputInstanceAttributes) GetTerminationTimeOk() (*time.Time, bool) {
	if o == nil || o.TerminationTime == nil {
		return nil, false
	}
	return o.TerminationTime, true
}

// HasTerminationTime returns a boolean if a field has been set.
func (o *SpotMarketRequestCreateInputInstanceAttributes) HasTerminationTime() bool {
	if o != nil && o.TerminationTime != nil {
		return true
	}

	return false
}

// SetTerminationTime gets a reference to the given time.Time and assigns it to the TerminationTime field.
func (o *SpotMarketRequestCreateInputInstanceAttributes) SetTerminationTime(v time.Time) {
	o.TerminationTime = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *SpotMarketRequestCreateInputInstanceAttributes) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotMarketRequestCreateInputInstanceAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *SpotMarketRequestCreateInputInstanceAttributes) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *SpotMarketRequestCreateInputInstanceAttributes) SetTags(v []string) {
	o.Tags = &v
}

// GetProjectSshKeys returns the ProjectSshKeys field value if set, zero value otherwise.
func (o *SpotMarketRequestCreateInputInstanceAttributes) GetProjectSshKeys() []string {
	if o == nil || o.ProjectSshKeys == nil {
		var ret []string
		return ret
	}
	return *o.ProjectSshKeys
}

// GetProjectSshKeysOk returns a tuple with the ProjectSshKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotMarketRequestCreateInputInstanceAttributes) GetProjectSshKeysOk() (*[]string, bool) {
	if o == nil || o.ProjectSshKeys == nil {
		return nil, false
	}
	return o.ProjectSshKeys, true
}

// HasProjectSshKeys returns a boolean if a field has been set.
func (o *SpotMarketRequestCreateInputInstanceAttributes) HasProjectSshKeys() bool {
	if o != nil && o.ProjectSshKeys != nil {
		return true
	}

	return false
}

// SetProjectSshKeys gets a reference to the given []string and assigns it to the ProjectSshKeys field.
func (o *SpotMarketRequestCreateInputInstanceAttributes) SetProjectSshKeys(v []string) {
	o.ProjectSshKeys = &v
}

// GetUserSshKeys returns the UserSshKeys field value if set, zero value otherwise.
func (o *SpotMarketRequestCreateInputInstanceAttributes) GetUserSshKeys() []string {
	if o == nil || o.UserSshKeys == nil {
		var ret []string
		return ret
	}
	return *o.UserSshKeys
}

// GetUserSshKeysOk returns a tuple with the UserSshKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotMarketRequestCreateInputInstanceAttributes) GetUserSshKeysOk() (*[]string, bool) {
	if o == nil || o.UserSshKeys == nil {
		return nil, false
	}
	return o.UserSshKeys, true
}

// HasUserSshKeys returns a boolean if a field has been set.
func (o *SpotMarketRequestCreateInputInstanceAttributes) HasUserSshKeys() bool {
	if o != nil && o.UserSshKeys != nil {
		return true
	}

	return false
}

// SetUserSshKeys gets a reference to the given []string and assigns it to the UserSshKeys field.
func (o *SpotMarketRequestCreateInputInstanceAttributes) SetUserSshKeys(v []string) {
	o.UserSshKeys = &v
}

// GetNoSshKeys returns the NoSshKeys field value if set, zero value otherwise.
func (o *SpotMarketRequestCreateInputInstanceAttributes) GetNoSshKeys() bool {
	if o == nil || o.NoSshKeys == nil {
		var ret bool
		return ret
	}
	return *o.NoSshKeys
}

// GetNoSshKeysOk returns a tuple with the NoSshKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotMarketRequestCreateInputInstanceAttributes) GetNoSshKeysOk() (*bool, bool) {
	if o == nil || o.NoSshKeys == nil {
		return nil, false
	}
	return o.NoSshKeys, true
}

// HasNoSshKeys returns a boolean if a field has been set.
func (o *SpotMarketRequestCreateInputInstanceAttributes) HasNoSshKeys() bool {
	if o != nil && o.NoSshKeys != nil {
		return true
	}

	return false
}

// SetNoSshKeys gets a reference to the given bool and assigns it to the NoSshKeys field.
func (o *SpotMarketRequestCreateInputInstanceAttributes) SetNoSshKeys(v bool) {
	o.NoSshKeys = &v
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *SpotMarketRequestCreateInputInstanceAttributes) GetFeatures() []string {
	if o == nil || o.Features == nil {
		var ret []string
		return ret
	}
	return *o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotMarketRequestCreateInputInstanceAttributes) GetFeaturesOk() (*[]string, bool) {
	if o == nil || o.Features == nil {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *SpotMarketRequestCreateInputInstanceAttributes) HasFeatures() bool {
	if o != nil && o.Features != nil {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given []string and assigns it to the Features field.
func (o *SpotMarketRequestCreateInputInstanceAttributes) SetFeatures(v []string) {
	o.Features = &v
}

// GetCustomdata returns the Customdata field value if set, zero value otherwise.
func (o *SpotMarketRequestCreateInputInstanceAttributes) GetCustomdata() map[string]interface{} {
	if o == nil || o.Customdata == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Customdata
}

// GetCustomdataOk returns a tuple with the Customdata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotMarketRequestCreateInputInstanceAttributes) GetCustomdataOk() (*map[string]interface{}, bool) {
	if o == nil || o.Customdata == nil {
		return nil, false
	}
	return o.Customdata, true
}

// HasCustomdata returns a boolean if a field has been set.
func (o *SpotMarketRequestCreateInputInstanceAttributes) HasCustomdata() bool {
	if o != nil && o.Customdata != nil {
		return true
	}

	return false
}

// SetCustomdata gets a reference to the given map[string]interface{} and assigns it to the Customdata field.
func (o *SpotMarketRequestCreateInputInstanceAttributes) SetCustomdata(v map[string]interface{}) {
	o.Customdata = &v
}

// GetPublicIpv4SubnetSize returns the PublicIpv4SubnetSize field value if set, zero value otherwise.
func (o *SpotMarketRequestCreateInputInstanceAttributes) GetPublicIpv4SubnetSize() int32 {
	if o == nil || o.PublicIpv4SubnetSize == nil {
		var ret int32
		return ret
	}
	return *o.PublicIpv4SubnetSize
}

// GetPublicIpv4SubnetSizeOk returns a tuple with the PublicIpv4SubnetSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotMarketRequestCreateInputInstanceAttributes) GetPublicIpv4SubnetSizeOk() (*int32, bool) {
	if o == nil || o.PublicIpv4SubnetSize == nil {
		return nil, false
	}
	return o.PublicIpv4SubnetSize, true
}

// HasPublicIpv4SubnetSize returns a boolean if a field has been set.
func (o *SpotMarketRequestCreateInputInstanceAttributes) HasPublicIpv4SubnetSize() bool {
	if o != nil && o.PublicIpv4SubnetSize != nil {
		return true
	}

	return false
}

// SetPublicIpv4SubnetSize gets a reference to the given int32 and assigns it to the PublicIpv4SubnetSize field.
func (o *SpotMarketRequestCreateInputInstanceAttributes) SetPublicIpv4SubnetSize(v int32) {
	o.PublicIpv4SubnetSize = &v
}

// GetPrivateIpv4SubnetSize returns the PrivateIpv4SubnetSize field value if set, zero value otherwise.
func (o *SpotMarketRequestCreateInputInstanceAttributes) GetPrivateIpv4SubnetSize() int32 {
	if o == nil || o.PrivateIpv4SubnetSize == nil {
		var ret int32
		return ret
	}
	return *o.PrivateIpv4SubnetSize
}

// GetPrivateIpv4SubnetSizeOk returns a tuple with the PrivateIpv4SubnetSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotMarketRequestCreateInputInstanceAttributes) GetPrivateIpv4SubnetSizeOk() (*int32, bool) {
	if o == nil || o.PrivateIpv4SubnetSize == nil {
		return nil, false
	}
	return o.PrivateIpv4SubnetSize, true
}

// HasPrivateIpv4SubnetSize returns a boolean if a field has been set.
func (o *SpotMarketRequestCreateInputInstanceAttributes) HasPrivateIpv4SubnetSize() bool {
	if o != nil && o.PrivateIpv4SubnetSize != nil {
		return true
	}

	return false
}

// SetPrivateIpv4SubnetSize gets a reference to the given int32 and assigns it to the PrivateIpv4SubnetSize field.
func (o *SpotMarketRequestCreateInputInstanceAttributes) SetPrivateIpv4SubnetSize(v int32) {
	o.PrivateIpv4SubnetSize = &v
}

func (o SpotMarketRequestCreateInputInstanceAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Plan != nil {
		toSerialize["plan"] = o.Plan
	}
	if o.Hostname != nil {
		toSerialize["hostname"] = o.Hostname
	}
	if o.Hostnames != nil {
		toSerialize["hostnames"] = o.Hostnames
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.BillingCycle != nil {
		toSerialize["billing_cycle"] = o.BillingCycle
	}
	if o.OperatingSystem != nil {
		toSerialize["operating_system"] = o.OperatingSystem
	}
	if o.AlwaysPxe != nil {
		toSerialize["always_pxe"] = o.AlwaysPxe
	}
	if o.Userdata != nil {
		toSerialize["userdata"] = o.Userdata
	}
	if o.Locked != nil {
		toSerialize["locked"] = o.Locked
	}
	if o.TerminationTime != nil {
		toSerialize["termination_time"] = o.TerminationTime
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.ProjectSshKeys != nil {
		toSerialize["project_ssh_keys"] = o.ProjectSshKeys
	}
	if o.UserSshKeys != nil {
		toSerialize["user_ssh_keys"] = o.UserSshKeys
	}
	if o.NoSshKeys != nil {
		toSerialize["no_ssh_keys"] = o.NoSshKeys
	}
	if o.Features != nil {
		toSerialize["features"] = o.Features
	}
	if o.Customdata != nil {
		toSerialize["customdata"] = o.Customdata
	}
	if o.PublicIpv4SubnetSize != nil {
		toSerialize["public_ipv4_subnet_size"] = o.PublicIpv4SubnetSize
	}
	if o.PrivateIpv4SubnetSize != nil {
		toSerialize["private_ipv4_subnet_size"] = o.PrivateIpv4SubnetSize
	}
	return json.Marshal(toSerialize)
}

type NullableSpotMarketRequestCreateInputInstanceAttributes struct {
	value *SpotMarketRequestCreateInputInstanceAttributes
	isSet bool
}

func (v NullableSpotMarketRequestCreateInputInstanceAttributes) Get() *SpotMarketRequestCreateInputInstanceAttributes {
	return v.value
}

func (v *NullableSpotMarketRequestCreateInputInstanceAttributes) Set(val *SpotMarketRequestCreateInputInstanceAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableSpotMarketRequestCreateInputInstanceAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableSpotMarketRequestCreateInputInstanceAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpotMarketRequestCreateInputInstanceAttributes(val *SpotMarketRequestCreateInputInstanceAttributes) *NullableSpotMarketRequestCreateInputInstanceAttributes {
	return &NullableSpotMarketRequestCreateInputInstanceAttributes{value: val, isSet: true}
}

func (v NullableSpotMarketRequestCreateInputInstanceAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpotMarketRequestCreateInputInstanceAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

