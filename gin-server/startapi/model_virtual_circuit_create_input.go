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

type VirtualCircuitCreateInput struct {
	Description string `json:"description,omitempty"`

	Name string `json:"name,omitempty"`

	NniVlan int32 `json:"nni_vlan,omitempty"`

	Project string `json:"project,omitempty"`

	// speed can be passed as integer number representing bps speed or string (e.g. '52m' or '100g' or '4 gbps')
	Speed int32 `json:"speed,omitempty"`

	Tags []string `json:"tags,omitempty"`

	// A Virtual Network record UUID or the VNID of a Virtual Network in your project (sent as integer).
	Vnid string `json:"vnid,omitempty"`
}
