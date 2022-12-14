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

type VirtualNetworkCreateInput struct {
	Description string `json:"description,omitempty"`

	// The UUID (or facility code) for the Facility in which to create this Virtual network.
	Facility string `json:"facility,omitempty"`

	// The UUID (or metro code) for the Metro in which to create this Virtual Network.
	Metro string `json:"metro,omitempty"`

	ProjectId string `json:"project_id"`

	// VLAN ID between 2-3999. Must be unique for the project within the Metro in which this Virtual Network is being created. If no value is specified, the next-available VLAN ID in the range 1000-1999 will be automatically selected.
	Vxlan int32 `json:"vxlan,omitempty"`
}
