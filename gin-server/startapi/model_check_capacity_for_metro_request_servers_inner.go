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

type CheckCapacityForMetroRequestServersInner struct {

	// The metro ID or code to check the capacity in.
	Metro string `json:"metro,omitempty"`

	// The plan ID or slug to check the capacity of.
	Plan string `json:"plan,omitempty"`

	// The number of servers to check the capacity of.
	Quantity string `json:"quantity,omitempty"`
}
