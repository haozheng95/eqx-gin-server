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

type FindInvitationById200Response struct {
	CreatedAt time.Time `json:"created_at,omitempty"`

	Href string `json:"href,omitempty"`

	Id string `json:"id,omitempty"`

	Invitation FindBatchById200ResponseDevicesInner `json:"invitation,omitempty"`

	InvitedBy FindBatchById200ResponseDevicesInner `json:"invited_by,omitempty"`

	Invitee string `json:"invitee,omitempty"`

	Nonce string `json:"nonce,omitempty"`

	Organization FindBatchById200ResponseDevicesInner `json:"organization,omitempty"`

	Projects []FindBatchById200ResponseDevicesInner `json:"projects,omitempty"`

	Roles []string `json:"roles,omitempty"`

	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
