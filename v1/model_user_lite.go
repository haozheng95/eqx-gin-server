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

// UserLite struct for UserLite
type UserLite struct {
	// ID of the User
	Id string `json:"id"`
	// Short ID of the User
	ShortId string `json:"short_id"`
	// First name of the User
	FirstName *string `json:"first_name,omitempty"`
	// Last name of the User
	LastName *string `json:"last_name,omitempty"`
	// Full name of the User
	FullName *string `json:"full_name,omitempty"`
	// Primary email address of the User
	Email *string `json:"email,omitempty"`
	// Avatar thumbnail URL of the User
	AvatarThumbUrl *string `json:"avatar_thumb_url,omitempty"`
	// When the user was created
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// When the user details were last updated
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// API URL uniquely representing the User
	Href *string `json:"href,omitempty"`
}

// NewUserLite instantiates a new UserLite object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserLite(id string, shortId string) *UserLite {
	this := UserLite{}
	this.Id = id
	this.ShortId = shortId
	return &this
}

// NewUserLiteWithDefaults instantiates a new UserLite object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserLiteWithDefaults() *UserLite {
	this := UserLite{}
	return &this
}

// GetId returns the Id field value
func (o *UserLite) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UserLite) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UserLite) SetId(v string) {
	o.Id = v
}

// GetShortId returns the ShortId field value
func (o *UserLite) GetShortId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ShortId
}

// GetShortIdOk returns a tuple with the ShortId field value
// and a boolean to check if the value has been set.
func (o *UserLite) GetShortIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ShortId, true
}

// SetShortId sets field value
func (o *UserLite) SetShortId(v string) {
	o.ShortId = v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *UserLite) GetFirstName() string {
	if o == nil || o.FirstName == nil {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLite) GetFirstNameOk() (*string, bool) {
	if o == nil || o.FirstName == nil {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *UserLite) HasFirstName() bool {
	if o != nil && o.FirstName != nil {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *UserLite) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *UserLite) GetLastName() string {
	if o == nil || o.LastName == nil {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLite) GetLastNameOk() (*string, bool) {
	if o == nil || o.LastName == nil {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *UserLite) HasLastName() bool {
	if o != nil && o.LastName != nil {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *UserLite) SetLastName(v string) {
	o.LastName = &v
}

// GetFullName returns the FullName field value if set, zero value otherwise.
func (o *UserLite) GetFullName() string {
	if o == nil || o.FullName == nil {
		var ret string
		return ret
	}
	return *o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLite) GetFullNameOk() (*string, bool) {
	if o == nil || o.FullName == nil {
		return nil, false
	}
	return o.FullName, true
}

// HasFullName returns a boolean if a field has been set.
func (o *UserLite) HasFullName() bool {
	if o != nil && o.FullName != nil {
		return true
	}

	return false
}

// SetFullName gets a reference to the given string and assigns it to the FullName field.
func (o *UserLite) SetFullName(v string) {
	o.FullName = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UserLite) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLite) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UserLite) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *UserLite) SetEmail(v string) {
	o.Email = &v
}

// GetAvatarThumbUrl returns the AvatarThumbUrl field value if set, zero value otherwise.
func (o *UserLite) GetAvatarThumbUrl() string {
	if o == nil || o.AvatarThumbUrl == nil {
		var ret string
		return ret
	}
	return *o.AvatarThumbUrl
}

// GetAvatarThumbUrlOk returns a tuple with the AvatarThumbUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLite) GetAvatarThumbUrlOk() (*string, bool) {
	if o == nil || o.AvatarThumbUrl == nil {
		return nil, false
	}
	return o.AvatarThumbUrl, true
}

// HasAvatarThumbUrl returns a boolean if a field has been set.
func (o *UserLite) HasAvatarThumbUrl() bool {
	if o != nil && o.AvatarThumbUrl != nil {
		return true
	}

	return false
}

// SetAvatarThumbUrl gets a reference to the given string and assigns it to the AvatarThumbUrl field.
func (o *UserLite) SetAvatarThumbUrl(v string) {
	o.AvatarThumbUrl = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *UserLite) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLite) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *UserLite) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *UserLite) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *UserLite) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLite) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *UserLite) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *UserLite) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *UserLite) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLite) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *UserLite) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *UserLite) SetHref(v string) {
	o.Href = &v
}

func (o UserLite) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["short_id"] = o.ShortId
	}
	if o.FirstName != nil {
		toSerialize["first_name"] = o.FirstName
	}
	if o.LastName != nil {
		toSerialize["last_name"] = o.LastName
	}
	if o.FullName != nil {
		toSerialize["full_name"] = o.FullName
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.AvatarThumbUrl != nil {
		toSerialize["avatar_thumb_url"] = o.AvatarThumbUrl
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	return json.Marshal(toSerialize)
}

type NullableUserLite struct {
	value *UserLite
	isSet bool
}

func (v NullableUserLite) Get() *UserLite {
	return v.value
}

func (v *NullableUserLite) Set(val *UserLite) {
	v.value = val
	v.isSet = true
}

func (v NullableUserLite) IsSet() bool {
	return v.isSet
}

func (v *NullableUserLite) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserLite(val *UserLite) *NullableUserLite {
	return &NullableUserLite{value: val, isSet: true}
}

func (v NullableUserLite) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserLite) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


