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

type InterconnectionUpdateInput struct {
	ContactEmail string `json:"contact_email,omitempty"`

	Description string `json:"description,omitempty"`

	// The mode of the interconnection (only relevant to Dedicated Ports). Shared connections won't have this field. Can be either 'standard' or 'tunnel'.   The default mode of an interconnection on a Dedicated Port is 'standard'. The mode can only be changed when there are no associated virtual circuits on the interconnection.   In tunnel mode, an 802.1q tunnel is added to a port to send/receive double tagged packets from server instances.
	Mode string `json:"mode,omitempty"`

	Name string `json:"name,omitempty"`

	// Updating from 'redundant' to 'primary' will remove a secondary port, while updating from 'primary' to 'redundant' will add one.
	Redundancy string `json:"redundancy,omitempty"`

	Tags []string `json:"tags,omitempty"`
}
