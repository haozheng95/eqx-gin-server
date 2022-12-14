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

// Membership struct for Membership
type Membership struct {
	CreatedAt *time.Time                            `json:"created_at,omitempty"`
	Href      *string                               `json:"href,omitempty"`
	Id        *string                               `json:"id,omitempty"`
	Project   *FindBatchById200ResponseDevicesInner `json:"project,omitempty"`
	Roles     []string                              `json:"roles,omitempty"`
	UpdatedAt *time.Time                            `json:"updated_at,omitempty"`
	User      *FindBatchById200ResponseDevicesInner `json:"user,omitempty"`
}

// NewMembership instantiates a new Membership object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMembership() *Membership {
	this := Membership{}
	return &this
}

// NewMembershipWithDefaults instantiates a new Membership object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMembershipWithDefaults() *Membership {
	this := Membership{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Membership) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Membership) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Membership) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Membership) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *Membership) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Membership) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *Membership) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *Membership) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Membership) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Membership) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Membership) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Membership) SetId(v string) {
	o.Id = &v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *Membership) GetProject() FindBatchById200ResponseDevicesInner {
	if o == nil || o.Project == nil {
		var ret FindBatchById200ResponseDevicesInner
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Membership) GetProjectOk() (*FindBatchById200ResponseDevicesInner, bool) {
	if o == nil || o.Project == nil {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *Membership) HasProject() bool {
	if o != nil && o.Project != nil {
		return true
	}

	return false
}

// SetProject gets a reference to the given FindBatchById200ResponseDevicesInner and assigns it to the Project field.
func (o *Membership) SetProject(v FindBatchById200ResponseDevicesInner) {
	o.Project = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *Membership) GetRoles() []string {
	if o == nil || o.Roles == nil {
		var ret []string
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Membership) GetRolesOk() ([]string, bool) {
	if o == nil || o.Roles == nil {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *Membership) HasRoles() bool {
	if o != nil && o.Roles != nil {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *Membership) SetRoles(v []string) {
	o.Roles = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Membership) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Membership) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Membership) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Membership) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *Membership) GetUser() FindBatchById200ResponseDevicesInner {
	if o == nil || o.User == nil {
		var ret FindBatchById200ResponseDevicesInner
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Membership) GetUserOk() (*FindBatchById200ResponseDevicesInner, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *Membership) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given FindBatchById200ResponseDevicesInner and assigns it to the User field.
func (o *Membership) SetUser(v FindBatchById200ResponseDevicesInner) {
	o.User = &v
}

func (o Membership) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
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
	if o.Roles != nil {
		toSerialize["roles"] = o.Roles
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	return json.Marshal(toSerialize)
}

type NullableMembership struct {
	value *Membership
	isSet bool
}

func (v NullableMembership) Get() *Membership {
	return v.value
}

func (v *NullableMembership) Set(val *Membership) {
	v.value = val
	v.isSet = true
}

func (v NullableMembership) IsSet() bool {
	return v.isSet
}

func (v *NullableMembership) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMembership(val *Membership) *NullableMembership {
	return &NullableMembership{value: val, isSet: true}
}

func (v NullableMembership) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMembership) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
