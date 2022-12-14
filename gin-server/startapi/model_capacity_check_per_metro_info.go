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

type CapacityCheckPerMetroInfo struct {

	// Returns true if there is enough capacity in the metro to fulfill the quantity set. Returns false if there is not enough.
	Available bool `json:"available,omitempty"`

	// The metro ID or code sent to check capacity.
	Metro string `json:"metro,omitempty"`

	// The plan ID or slug sent to check capacity.
	Plan string `json:"plan,omitempty"`

	// The number of servers sent to check capacity.
	Quantity string `json:"quantity,omitempty"`
}
