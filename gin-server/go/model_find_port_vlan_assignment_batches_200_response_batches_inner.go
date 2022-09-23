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

type FindPortVlanAssignmentBatches200ResponseBatchesInner struct {
	CreatedAt time.Time `json:"created_at,omitempty"`

	ErrorMessages []string `json:"error_messages,omitempty"`

	Id string `json:"id,omitempty"`

	Port FindDeviceById200ResponseNetworkPortsInner `json:"port,omitempty"`

	Quantity int32 `json:"quantity,omitempty"`

	State string `json:"state,omitempty"`

	UpdatedAt time.Time `json:"updated_at,omitempty"`

	VlanAssignments []FindPortVlanAssignmentBatches200ResponseBatchesInnerVlanAssignmentsInner `json:"vlan_assignments,omitempty"`
}
