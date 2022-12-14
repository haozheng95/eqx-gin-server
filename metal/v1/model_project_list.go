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

// ProjectList struct for ProjectList
type ProjectList struct {
	Meta     *FindDeviceEvents200ResponseMeta                                                             `json:"meta,omitempty"`
	Projects []GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1VrfProject `json:"projects,omitempty"`
}

// NewProjectList instantiates a new ProjectList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectList() *ProjectList {
	this := ProjectList{}
	return &this
}

// NewProjectListWithDefaults instantiates a new ProjectList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectListWithDefaults() *ProjectList {
	this := ProjectList{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ProjectList) GetMeta() FindDeviceEvents200ResponseMeta {
	if o == nil || o.Meta == nil {
		var ret FindDeviceEvents200ResponseMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectList) GetMetaOk() (*FindDeviceEvents200ResponseMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ProjectList) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given FindDeviceEvents200ResponseMeta and assigns it to the Meta field.
func (o *ProjectList) SetMeta(v FindDeviceEvents200ResponseMeta) {
	o.Meta = &v
}

// GetProjects returns the Projects field value if set, zero value otherwise.
func (o *ProjectList) GetProjects() []GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1VrfProject {
	if o == nil || o.Projects == nil {
		var ret []GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1VrfProject
		return ret
	}
	return o.Projects
}

// GetProjectsOk returns a tuple with the Projects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectList) GetProjectsOk() ([]GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1VrfProject, bool) {
	if o == nil || o.Projects == nil {
		return nil, false
	}
	return o.Projects, true
}

// HasProjects returns a boolean if a field has been set.
func (o *ProjectList) HasProjects() bool {
	if o != nil && o.Projects != nil {
		return true
	}

	return false
}

// SetProjects gets a reference to the given []GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1VrfProject and assigns it to the Projects field.
func (o *ProjectList) SetProjects(v []GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1VrfProject) {
	o.Projects = v
}

func (o ProjectList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	if o.Projects != nil {
		toSerialize["projects"] = o.Projects
	}
	return json.Marshal(toSerialize)
}

type NullableProjectList struct {
	value *ProjectList
	isSet bool
}

func (v NullableProjectList) Get() *ProjectList {
	return v.value
}

func (v *NullableProjectList) Set(val *ProjectList) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectList) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectList(val *ProjectList) *NullableProjectList {
	return &NullableProjectList{value: val, isSet: true}
}

func (v NullableProjectList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
