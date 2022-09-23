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

type InterconnectionPort struct {
	Id string `json:"id,omitempty"`

	Organization FindBatchById200ResponseDevicesInner `json:"organization,omitempty"`

	// Either 'primary' or 'secondary'.
	Role string `json:"role,omitempty"`

	// For both Fabric VCs and Dedicated Ports, this will be 'requested' on creation and 'deleting' on deletion. Once the Fabric VC has found its corresponding Fabric connection, this will turn to 'active'. For Dedicated Ports, once the dedicated port is associated, this will also turn to 'active'. For Fabric VCs, this can turn into an 'expired' state if the service token associated is expired.
	Status string `json:"status,omitempty"`

	// A switch 'short ID'
	SwitchId string `json:"switch_id,omitempty"`

	VirtualCircuits GetInterconnection200ResponsePortsInnerVirtualCircuits `json:"virtual_circuits,omitempty"`
}
