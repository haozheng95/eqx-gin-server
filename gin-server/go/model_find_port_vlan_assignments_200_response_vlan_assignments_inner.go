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

type FindPortVlanAssignments200ResponseVlanAssignmentsInner struct {
	CreatedAt time.Time `json:"created_at,omitempty"`

	Id string `json:"id,omitempty"`

	Native bool `json:"native,omitempty"`

	Port FindBatchById200ResponseDevicesInner `json:"port,omitempty"`

	State string `json:"state,omitempty"`

	UpdatedAt time.Time `json:"updated_at,omitempty"`

	VirtualNetwork FindBatchById200ResponseDevicesInner `json:"virtual_network,omitempty"`

	Vlan int32 `json:"vlan,omitempty"`
}
