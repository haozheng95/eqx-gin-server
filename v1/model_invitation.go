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

// Invitation struct for Invitation
type Invitation struct {
	Id *string `json:"id,omitempty"`
	Roles []string `json:"roles,omitempty"`
	Invitee *string `json:"invitee,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	InvitedBy *Href `json:"invited_by,omitempty"`
	Organization *Href `json:"organization,omitempty"`
	ProjectsIds []string `json:"projects_ids,omitempty"`
	Invitation *Href `json:"invitation,omitempty"`
	Href *string `json:"href,omitempty"`
}

// NewInvitation instantiates a new Invitation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvitation() *Invitation {
	this := Invitation{}
	return &this
}

// NewInvitationWithDefaults instantiates a new Invitation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvitationWithDefaults() *Invitation {
	this := Invitation{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Invitation) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invitation) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Invitation) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Invitation) SetId(v string) {
	o.Id = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *Invitation) GetRoles() []string {
	if o == nil || o.Roles == nil {
		var ret []string
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invitation) GetRolesOk() ([]string, bool) {
	if o == nil || o.Roles == nil {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *Invitation) HasRoles() bool {
	if o != nil && o.Roles != nil {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *Invitation) SetRoles(v []string) {
	o.Roles = v
}

// GetInvitee returns the Invitee field value if set, zero value otherwise.
func (o *Invitation) GetInvitee() string {
	if o == nil || o.Invitee == nil {
		var ret string
		return ret
	}
	return *o.Invitee
}

// GetInviteeOk returns a tuple with the Invitee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invitation) GetInviteeOk() (*string, bool) {
	if o == nil || o.Invitee == nil {
		return nil, false
	}
	return o.Invitee, true
}

// HasInvitee returns a boolean if a field has been set.
func (o *Invitation) HasInvitee() bool {
	if o != nil && o.Invitee != nil {
		return true
	}

	return false
}

// SetInvitee gets a reference to the given string and assigns it to the Invitee field.
func (o *Invitation) SetInvitee(v string) {
	o.Invitee = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Invitation) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invitation) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Invitation) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Invitation) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Invitation) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invitation) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Invitation) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Invitation) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetInvitedBy returns the InvitedBy field value if set, zero value otherwise.
func (o *Invitation) GetInvitedBy() Href {
	if o == nil || o.InvitedBy == nil {
		var ret Href
		return ret
	}
	return *o.InvitedBy
}

// GetInvitedByOk returns a tuple with the InvitedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invitation) GetInvitedByOk() (*Href, bool) {
	if o == nil || o.InvitedBy == nil {
		return nil, false
	}
	return o.InvitedBy, true
}

// HasInvitedBy returns a boolean if a field has been set.
func (o *Invitation) HasInvitedBy() bool {
	if o != nil && o.InvitedBy != nil {
		return true
	}

	return false
}

// SetInvitedBy gets a reference to the given Href and assigns it to the InvitedBy field.
func (o *Invitation) SetInvitedBy(v Href) {
	o.InvitedBy = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *Invitation) GetOrganization() Href {
	if o == nil || o.Organization == nil {
		var ret Href
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invitation) GetOrganizationOk() (*Href, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *Invitation) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given Href and assigns it to the Organization field.
func (o *Invitation) SetOrganization(v Href) {
	o.Organization = &v
}

// GetProjectsIds returns the ProjectsIds field value if set, zero value otherwise.
func (o *Invitation) GetProjectsIds() []string {
	if o == nil || o.ProjectsIds == nil {
		var ret []string
		return ret
	}
	return o.ProjectsIds
}

// GetProjectsIdsOk returns a tuple with the ProjectsIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invitation) GetProjectsIdsOk() ([]string, bool) {
	if o == nil || o.ProjectsIds == nil {
		return nil, false
	}
	return o.ProjectsIds, true
}

// HasProjectsIds returns a boolean if a field has been set.
func (o *Invitation) HasProjectsIds() bool {
	if o != nil && o.ProjectsIds != nil {
		return true
	}

	return false
}

// SetProjectsIds gets a reference to the given []string and assigns it to the ProjectsIds field.
func (o *Invitation) SetProjectsIds(v []string) {
	o.ProjectsIds = v
}

// GetInvitation returns the Invitation field value if set, zero value otherwise.
func (o *Invitation) GetInvitation() Href {
	if o == nil || o.Invitation == nil {
		var ret Href
		return ret
	}
	return *o.Invitation
}

// GetInvitationOk returns a tuple with the Invitation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invitation) GetInvitationOk() (*Href, bool) {
	if o == nil || o.Invitation == nil {
		return nil, false
	}
	return o.Invitation, true
}

// HasInvitation returns a boolean if a field has been set.
func (o *Invitation) HasInvitation() bool {
	if o != nil && o.Invitation != nil {
		return true
	}

	return false
}

// SetInvitation gets a reference to the given Href and assigns it to the Invitation field.
func (o *Invitation) SetInvitation(v Href) {
	o.Invitation = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *Invitation) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invitation) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *Invitation) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *Invitation) SetHref(v string) {
	o.Href = &v
}

func (o Invitation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Roles != nil {
		toSerialize["roles"] = o.Roles
	}
	if o.Invitee != nil {
		toSerialize["invitee"] = o.Invitee
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.InvitedBy != nil {
		toSerialize["invited_by"] = o.InvitedBy
	}
	if o.Organization != nil {
		toSerialize["organization"] = o.Organization
	}
	if o.ProjectsIds != nil {
		toSerialize["projects_ids"] = o.ProjectsIds
	}
	if o.Invitation != nil {
		toSerialize["invitation"] = o.Invitation
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	return json.Marshal(toSerialize)
}

type NullableInvitation struct {
	value *Invitation
	isSet bool
}

func (v NullableInvitation) Get() *Invitation {
	return v.value
}

func (v *NullableInvitation) Set(val *Invitation) {
	v.value = val
	v.isSet = true
}

func (v NullableInvitation) IsSet() bool {
	return v.isSet
}

func (v *NullableInvitation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvitation(val *Invitation) *NullableInvitation {
	return &NullableInvitation{value: val, isSet: true}
}

func (v NullableInvitation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvitation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


