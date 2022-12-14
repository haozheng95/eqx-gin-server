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

// FindDeviceById200ResponsePlan struct for FindDeviceById200ResponsePlan
type FindDeviceById200ResponsePlan struct {
	// Shows which facilities the plan is available in, and the facility-based price if it is different from the default price.
	AvailableIn []FindDeviceById200ResponsePlanAvailableInInner `json:"available_in,omitempty"`
	// Shows which metros the plan is available in, and the metro-based price if it is different from the default price.
	AvailableInMetros []FindDeviceById200ResponsePlanAvailableInMetrosInner `json:"available_in_metros,omitempty"`
	Class             *string                                               `json:"class,omitempty"`
	Description       *string                                               `json:"description,omitempty"`
	DeploymentTypes   []string                                              `json:"deployment_types,omitempty"`
	Id                *string                                               `json:"id,omitempty"`
	Legacy            *bool                                                 `json:"legacy,omitempty"`
	Line              *string                                               `json:"line,omitempty"`
	Name              *string                                               `json:"name,omitempty"`
	Pricing           map[string]interface{}                                `json:"pricing,omitempty"`
	Slug              *string                                               `json:"slug,omitempty"`
	Specs             *FindDeviceById200ResponsePlanSpecs                   `json:"specs,omitempty"`
	// The plan type
	Type *string `json:"type,omitempty"`
}

// NewFindDeviceById200ResponsePlan instantiates a new FindDeviceById200ResponsePlan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFindDeviceById200ResponsePlan() *FindDeviceById200ResponsePlan {
	this := FindDeviceById200ResponsePlan{}
	return &this
}

// NewFindDeviceById200ResponsePlanWithDefaults instantiates a new FindDeviceById200ResponsePlan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFindDeviceById200ResponsePlanWithDefaults() *FindDeviceById200ResponsePlan {
	this := FindDeviceById200ResponsePlan{}
	return &this
}

// GetAvailableIn returns the AvailableIn field value if set, zero value otherwise.
func (o *FindDeviceById200ResponsePlan) GetAvailableIn() []FindDeviceById200ResponsePlanAvailableInInner {
	if o == nil || o.AvailableIn == nil {
		var ret []FindDeviceById200ResponsePlanAvailableInInner
		return ret
	}
	return o.AvailableIn
}

// GetAvailableInOk returns a tuple with the AvailableIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindDeviceById200ResponsePlan) GetAvailableInOk() ([]FindDeviceById200ResponsePlanAvailableInInner, bool) {
	if o == nil || o.AvailableIn == nil {
		return nil, false
	}
	return o.AvailableIn, true
}

// HasAvailableIn returns a boolean if a field has been set.
func (o *FindDeviceById200ResponsePlan) HasAvailableIn() bool {
	if o != nil && o.AvailableIn != nil {
		return true
	}

	return false
}

// SetAvailableIn gets a reference to the given []FindDeviceById200ResponsePlanAvailableInInner and assigns it to the AvailableIn field.
func (o *FindDeviceById200ResponsePlan) SetAvailableIn(v []FindDeviceById200ResponsePlanAvailableInInner) {
	o.AvailableIn = v
}

// GetAvailableInMetros returns the AvailableInMetros field value if set, zero value otherwise.
func (o *FindDeviceById200ResponsePlan) GetAvailableInMetros() []FindDeviceById200ResponsePlanAvailableInMetrosInner {
	if o == nil || o.AvailableInMetros == nil {
		var ret []FindDeviceById200ResponsePlanAvailableInMetrosInner
		return ret
	}
	return o.AvailableInMetros
}

// GetAvailableInMetrosOk returns a tuple with the AvailableInMetros field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindDeviceById200ResponsePlan) GetAvailableInMetrosOk() ([]FindDeviceById200ResponsePlanAvailableInMetrosInner, bool) {
	if o == nil || o.AvailableInMetros == nil {
		return nil, false
	}
	return o.AvailableInMetros, true
}

// HasAvailableInMetros returns a boolean if a field has been set.
func (o *FindDeviceById200ResponsePlan) HasAvailableInMetros() bool {
	if o != nil && o.AvailableInMetros != nil {
		return true
	}

	return false
}

// SetAvailableInMetros gets a reference to the given []FindDeviceById200ResponsePlanAvailableInMetrosInner and assigns it to the AvailableInMetros field.
func (o *FindDeviceById200ResponsePlan) SetAvailableInMetros(v []FindDeviceById200ResponsePlanAvailableInMetrosInner) {
	o.AvailableInMetros = v
}

// GetClass returns the Class field value if set, zero value otherwise.
func (o *FindDeviceById200ResponsePlan) GetClass() string {
	if o == nil || o.Class == nil {
		var ret string
		return ret
	}
	return *o.Class
}

// GetClassOk returns a tuple with the Class field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindDeviceById200ResponsePlan) GetClassOk() (*string, bool) {
	if o == nil || o.Class == nil {
		return nil, false
	}
	return o.Class, true
}

// HasClass returns a boolean if a field has been set.
func (o *FindDeviceById200ResponsePlan) HasClass() bool {
	if o != nil && o.Class != nil {
		return true
	}

	return false
}

// SetClass gets a reference to the given string and assigns it to the Class field.
func (o *FindDeviceById200ResponsePlan) SetClass(v string) {
	o.Class = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *FindDeviceById200ResponsePlan) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindDeviceById200ResponsePlan) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FindDeviceById200ResponsePlan) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FindDeviceById200ResponsePlan) SetDescription(v string) {
	o.Description = &v
}

// GetDeploymentTypes returns the DeploymentTypes field value if set, zero value otherwise.
func (o *FindDeviceById200ResponsePlan) GetDeploymentTypes() []string {
	if o == nil || o.DeploymentTypes == nil {
		var ret []string
		return ret
	}
	return o.DeploymentTypes
}

// GetDeploymentTypesOk returns a tuple with the DeploymentTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindDeviceById200ResponsePlan) GetDeploymentTypesOk() ([]string, bool) {
	if o == nil || o.DeploymentTypes == nil {
		return nil, false
	}
	return o.DeploymentTypes, true
}

// HasDeploymentTypes returns a boolean if a field has been set.
func (o *FindDeviceById200ResponsePlan) HasDeploymentTypes() bool {
	if o != nil && o.DeploymentTypes != nil {
		return true
	}

	return false
}

// SetDeploymentTypes gets a reference to the given []string and assigns it to the DeploymentTypes field.
func (o *FindDeviceById200ResponsePlan) SetDeploymentTypes(v []string) {
	o.DeploymentTypes = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FindDeviceById200ResponsePlan) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindDeviceById200ResponsePlan) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FindDeviceById200ResponsePlan) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *FindDeviceById200ResponsePlan) SetId(v string) {
	o.Id = &v
}

// GetLegacy returns the Legacy field value if set, zero value otherwise.
func (o *FindDeviceById200ResponsePlan) GetLegacy() bool {
	if o == nil || o.Legacy == nil {
		var ret bool
		return ret
	}
	return *o.Legacy
}

// GetLegacyOk returns a tuple with the Legacy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindDeviceById200ResponsePlan) GetLegacyOk() (*bool, bool) {
	if o == nil || o.Legacy == nil {
		return nil, false
	}
	return o.Legacy, true
}

// HasLegacy returns a boolean if a field has been set.
func (o *FindDeviceById200ResponsePlan) HasLegacy() bool {
	if o != nil && o.Legacy != nil {
		return true
	}

	return false
}

// SetLegacy gets a reference to the given bool and assigns it to the Legacy field.
func (o *FindDeviceById200ResponsePlan) SetLegacy(v bool) {
	o.Legacy = &v
}

// GetLine returns the Line field value if set, zero value otherwise.
func (o *FindDeviceById200ResponsePlan) GetLine() string {
	if o == nil || o.Line == nil {
		var ret string
		return ret
	}
	return *o.Line
}

// GetLineOk returns a tuple with the Line field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindDeviceById200ResponsePlan) GetLineOk() (*string, bool) {
	if o == nil || o.Line == nil {
		return nil, false
	}
	return o.Line, true
}

// HasLine returns a boolean if a field has been set.
func (o *FindDeviceById200ResponsePlan) HasLine() bool {
	if o != nil && o.Line != nil {
		return true
	}

	return false
}

// SetLine gets a reference to the given string and assigns it to the Line field.
func (o *FindDeviceById200ResponsePlan) SetLine(v string) {
	o.Line = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FindDeviceById200ResponsePlan) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindDeviceById200ResponsePlan) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FindDeviceById200ResponsePlan) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FindDeviceById200ResponsePlan) SetName(v string) {
	o.Name = &v
}

// GetPricing returns the Pricing field value if set, zero value otherwise.
func (o *FindDeviceById200ResponsePlan) GetPricing() map[string]interface{} {
	if o == nil || o.Pricing == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Pricing
}

// GetPricingOk returns a tuple with the Pricing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindDeviceById200ResponsePlan) GetPricingOk() (map[string]interface{}, bool) {
	if o == nil || o.Pricing == nil {
		return nil, false
	}
	return o.Pricing, true
}

// HasPricing returns a boolean if a field has been set.
func (o *FindDeviceById200ResponsePlan) HasPricing() bool {
	if o != nil && o.Pricing != nil {
		return true
	}

	return false
}

// SetPricing gets a reference to the given map[string]interface{} and assigns it to the Pricing field.
func (o *FindDeviceById200ResponsePlan) SetPricing(v map[string]interface{}) {
	o.Pricing = v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *FindDeviceById200ResponsePlan) GetSlug() string {
	if o == nil || o.Slug == nil {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindDeviceById200ResponsePlan) GetSlugOk() (*string, bool) {
	if o == nil || o.Slug == nil {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *FindDeviceById200ResponsePlan) HasSlug() bool {
	if o != nil && o.Slug != nil {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *FindDeviceById200ResponsePlan) SetSlug(v string) {
	o.Slug = &v
}

// GetSpecs returns the Specs field value if set, zero value otherwise.
func (o *FindDeviceById200ResponsePlan) GetSpecs() FindDeviceById200ResponsePlanSpecs {
	if o == nil || o.Specs == nil {
		var ret FindDeviceById200ResponsePlanSpecs
		return ret
	}
	return *o.Specs
}

// GetSpecsOk returns a tuple with the Specs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindDeviceById200ResponsePlan) GetSpecsOk() (*FindDeviceById200ResponsePlanSpecs, bool) {
	if o == nil || o.Specs == nil {
		return nil, false
	}
	return o.Specs, true
}

// HasSpecs returns a boolean if a field has been set.
func (o *FindDeviceById200ResponsePlan) HasSpecs() bool {
	if o != nil && o.Specs != nil {
		return true
	}

	return false
}

// SetSpecs gets a reference to the given FindDeviceById200ResponsePlanSpecs and assigns it to the Specs field.
func (o *FindDeviceById200ResponsePlan) SetSpecs(v FindDeviceById200ResponsePlanSpecs) {
	o.Specs = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *FindDeviceById200ResponsePlan) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindDeviceById200ResponsePlan) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *FindDeviceById200ResponsePlan) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *FindDeviceById200ResponsePlan) SetType(v string) {
	o.Type = &v
}

func (o FindDeviceById200ResponsePlan) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AvailableIn != nil {
		toSerialize["available_in"] = o.AvailableIn
	}
	if o.AvailableInMetros != nil {
		toSerialize["available_in_metros"] = o.AvailableInMetros
	}
	if o.Class != nil {
		toSerialize["class"] = o.Class
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DeploymentTypes != nil {
		toSerialize["deployment_types"] = o.DeploymentTypes
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Legacy != nil {
		toSerialize["legacy"] = o.Legacy
	}
	if o.Line != nil {
		toSerialize["line"] = o.Line
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Pricing != nil {
		toSerialize["pricing"] = o.Pricing
	}
	if o.Slug != nil {
		toSerialize["slug"] = o.Slug
	}
	if o.Specs != nil {
		toSerialize["specs"] = o.Specs
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableFindDeviceById200ResponsePlan struct {
	value *FindDeviceById200ResponsePlan
	isSet bool
}

func (v NullableFindDeviceById200ResponsePlan) Get() *FindDeviceById200ResponsePlan {
	return v.value
}

func (v *NullableFindDeviceById200ResponsePlan) Set(val *FindDeviceById200ResponsePlan) {
	v.value = val
	v.isSet = true
}

func (v NullableFindDeviceById200ResponsePlan) IsSet() bool {
	return v.isSet
}

func (v *NullableFindDeviceById200ResponsePlan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFindDeviceById200ResponsePlan(val *FindDeviceById200ResponsePlan) *NullableFindDeviceById200ResponsePlan {
	return &NullableFindDeviceById200ResponsePlan{value: val, isSet: true}
}

func (v NullableFindDeviceById200ResponsePlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFindDeviceById200ResponsePlan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
