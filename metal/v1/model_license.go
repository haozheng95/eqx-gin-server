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

// License struct for License
type License struct {
	Description     *string                               `json:"description,omitempty"`
	Id              *string                               `json:"id,omitempty"`
	LicenseKey      *string                               `json:"license_key,omitempty"`
	LicenseeProduct *FindBatchById200ResponseDevicesInner `json:"licensee_product,omitempty"`
	Project         *FindBatchById200ResponseDevicesInner `json:"project,omitempty"`
	Size            *float32                              `json:"size,omitempty"`
}

// NewLicense instantiates a new License object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLicense() *License {
	this := License{}
	return &this
}

// NewLicenseWithDefaults instantiates a new License object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLicenseWithDefaults() *License {
	this := License{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *License) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *License) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *License) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *License) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *License) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *License) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *License) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *License) SetId(v string) {
	o.Id = &v
}

// GetLicenseKey returns the LicenseKey field value if set, zero value otherwise.
func (o *License) GetLicenseKey() string {
	if o == nil || o.LicenseKey == nil {
		var ret string
		return ret
	}
	return *o.LicenseKey
}

// GetLicenseKeyOk returns a tuple with the LicenseKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *License) GetLicenseKeyOk() (*string, bool) {
	if o == nil || o.LicenseKey == nil {
		return nil, false
	}
	return o.LicenseKey, true
}

// HasLicenseKey returns a boolean if a field has been set.
func (o *License) HasLicenseKey() bool {
	if o != nil && o.LicenseKey != nil {
		return true
	}

	return false
}

// SetLicenseKey gets a reference to the given string and assigns it to the LicenseKey field.
func (o *License) SetLicenseKey(v string) {
	o.LicenseKey = &v
}

// GetLicenseeProduct returns the LicenseeProduct field value if set, zero value otherwise.
func (o *License) GetLicenseeProduct() FindBatchById200ResponseDevicesInner {
	if o == nil || o.LicenseeProduct == nil {
		var ret FindBatchById200ResponseDevicesInner
		return ret
	}
	return *o.LicenseeProduct
}

// GetLicenseeProductOk returns a tuple with the LicenseeProduct field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *License) GetLicenseeProductOk() (*FindBatchById200ResponseDevicesInner, bool) {
	if o == nil || o.LicenseeProduct == nil {
		return nil, false
	}
	return o.LicenseeProduct, true
}

// HasLicenseeProduct returns a boolean if a field has been set.
func (o *License) HasLicenseeProduct() bool {
	if o != nil && o.LicenseeProduct != nil {
		return true
	}

	return false
}

// SetLicenseeProduct gets a reference to the given FindBatchById200ResponseDevicesInner and assigns it to the LicenseeProduct field.
func (o *License) SetLicenseeProduct(v FindBatchById200ResponseDevicesInner) {
	o.LicenseeProduct = &v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *License) GetProject() FindBatchById200ResponseDevicesInner {
	if o == nil || o.Project == nil {
		var ret FindBatchById200ResponseDevicesInner
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *License) GetProjectOk() (*FindBatchById200ResponseDevicesInner, bool) {
	if o == nil || o.Project == nil {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *License) HasProject() bool {
	if o != nil && o.Project != nil {
		return true
	}

	return false
}

// SetProject gets a reference to the given FindBatchById200ResponseDevicesInner and assigns it to the Project field.
func (o *License) SetProject(v FindBatchById200ResponseDevicesInner) {
	o.Project = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *License) GetSize() float32 {
	if o == nil || o.Size == nil {
		var ret float32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *License) GetSizeOk() (*float32, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *License) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given float32 and assigns it to the Size field.
func (o *License) SetSize(v float32) {
	o.Size = &v
}

func (o License) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LicenseKey != nil {
		toSerialize["license_key"] = o.LicenseKey
	}
	if o.LicenseeProduct != nil {
		toSerialize["licensee_product"] = o.LicenseeProduct
	}
	if o.Project != nil {
		toSerialize["project"] = o.Project
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	return json.Marshal(toSerialize)
}

type NullableLicense struct {
	value *License
	isSet bool
}

func (v NullableLicense) Get() *License {
	return v.value
}

func (v *NullableLicense) Set(val *License) {
	v.value = val
	v.isSet = true
}

func (v NullableLicense) IsSet() bool {
	return v.isSet
}

func (v *NullableLicense) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicense(val *License) *NullableLicense {
	return &NullableLicense{value: val, isSet: true}
}

func (v NullableLicense) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicense) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
