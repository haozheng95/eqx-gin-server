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

// BgpConfigRequestInput struct for BgpConfigRequestInput
type BgpConfigRequestInput struct {
	Asn            int32   `json:"asn"`
	DeploymentType string  `json:"deployment_type"`
	Md5            *string `json:"md5,omitempty"`
	UseCase        *string `json:"use_case,omitempty"`
}

// NewBgpConfigRequestInput instantiates a new BgpConfigRequestInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBgpConfigRequestInput(asn int32, deploymentType string) *BgpConfigRequestInput {
	this := BgpConfigRequestInput{}
	this.Asn = asn
	this.DeploymentType = deploymentType
	return &this
}

// NewBgpConfigRequestInputWithDefaults instantiates a new BgpConfigRequestInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBgpConfigRequestInputWithDefaults() *BgpConfigRequestInput {
	this := BgpConfigRequestInput{}
	return &this
}

// GetAsn returns the Asn field value
func (o *BgpConfigRequestInput) GetAsn() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Asn
}

// GetAsnOk returns a tuple with the Asn field value
// and a boolean to check if the value has been set.
func (o *BgpConfigRequestInput) GetAsnOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Asn, true
}

// SetAsn sets field value
func (o *BgpConfigRequestInput) SetAsn(v int32) {
	o.Asn = v
}

// GetDeploymentType returns the DeploymentType field value
func (o *BgpConfigRequestInput) GetDeploymentType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeploymentType
}

// GetDeploymentTypeOk returns a tuple with the DeploymentType field value
// and a boolean to check if the value has been set.
func (o *BgpConfigRequestInput) GetDeploymentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeploymentType, true
}

// SetDeploymentType sets field value
func (o *BgpConfigRequestInput) SetDeploymentType(v string) {
	o.DeploymentType = v
}

// GetMd5 returns the Md5 field value if set, zero value otherwise.
func (o *BgpConfigRequestInput) GetMd5() string {
	if o == nil || o.Md5 == nil {
		var ret string
		return ret
	}
	return *o.Md5
}

// GetMd5Ok returns a tuple with the Md5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpConfigRequestInput) GetMd5Ok() (*string, bool) {
	if o == nil || o.Md5 == nil {
		return nil, false
	}
	return o.Md5, true
}

// HasMd5 returns a boolean if a field has been set.
func (o *BgpConfigRequestInput) HasMd5() bool {
	if o != nil && o.Md5 != nil {
		return true
	}

	return false
}

// SetMd5 gets a reference to the given string and assigns it to the Md5 field.
func (o *BgpConfigRequestInput) SetMd5(v string) {
	o.Md5 = &v
}

// GetUseCase returns the UseCase field value if set, zero value otherwise.
func (o *BgpConfigRequestInput) GetUseCase() string {
	if o == nil || o.UseCase == nil {
		var ret string
		return ret
	}
	return *o.UseCase
}

// GetUseCaseOk returns a tuple with the UseCase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpConfigRequestInput) GetUseCaseOk() (*string, bool) {
	if o == nil || o.UseCase == nil {
		return nil, false
	}
	return o.UseCase, true
}

// HasUseCase returns a boolean if a field has been set.
func (o *BgpConfigRequestInput) HasUseCase() bool {
	if o != nil && o.UseCase != nil {
		return true
	}

	return false
}

// SetUseCase gets a reference to the given string and assigns it to the UseCase field.
func (o *BgpConfigRequestInput) SetUseCase(v string) {
	o.UseCase = &v
}

func (o BgpConfigRequestInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["asn"] = o.Asn
	}
	if true {
		toSerialize["deployment_type"] = o.DeploymentType
	}
	if o.Md5 != nil {
		toSerialize["md5"] = o.Md5
	}
	if o.UseCase != nil {
		toSerialize["use_case"] = o.UseCase
	}
	return json.Marshal(toSerialize)
}

type NullableBgpConfigRequestInput struct {
	value *BgpConfigRequestInput
	isSet bool
}

func (v NullableBgpConfigRequestInput) Get() *BgpConfigRequestInput {
	return v.value
}

func (v *NullableBgpConfigRequestInput) Set(val *BgpConfigRequestInput) {
	v.value = val
	v.isSet = true
}

func (v NullableBgpConfigRequestInput) IsSet() bool {
	return v.isSet
}

func (v *NullableBgpConfigRequestInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBgpConfigRequestInput(val *BgpConfigRequestInput) *NullableBgpConfigRequestInput {
	return &NullableBgpConfigRequestInput{value: val, isSet: true}
}

func (v NullableBgpConfigRequestInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBgpConfigRequestInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
