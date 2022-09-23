/*
 * Metal API
 *
 * This is the API for Equinix Metal. The API allows you to programmatically interact with all of your Equinix Metal resources, including devices, networks, addresses, organizations, projects, and your user account.  The official API docs are hosted at <https://metal.equinix.com/developers/api>.
 *
 * API version: 1.0.0
 * Contact: support@equinixmetal.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package startapi

import (
	"time"
)

type FindDeviceById200ResponseCreatedBy struct {

	// Avatar thumbnail URL of the User
	AvatarThumbUrl string `json:"avatar_thumb_url,omitempty"`

	// When the user was created
	CreatedAt time.Time `json:"created_at,omitempty"`

	// Primary email address of the User
	Email string `json:"email,omitempty"`

	// First name of the User
	FirstName string `json:"first_name,omitempty"`

	// Full name of the User
	FullName string `json:"full_name,omitempty"`

	// API URL uniquely representing the User
	Href string `json:"href,omitempty"`

	// ID of the User
	Id string `json:"id"`

	// Last name of the User
	LastName string `json:"last_name,omitempty"`

	// Short ID of the User
	ShortId string `json:"short_id"`

	// When the user details were last updated
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
