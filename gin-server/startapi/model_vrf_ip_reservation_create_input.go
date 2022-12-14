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

type VrfIpReservationCreateInput struct {

	// The size of the VRF IP Reservation's subnet
	Cidr int32 `json:"cidr"`

	Customdata map[string]interface{} `json:"customdata,omitempty"`

	Details string `json:"details,omitempty"`

	// The starting address for this VRF IP Reservation's subnet
	Network string `json:"network"`

	Tags []string `json:"tags,omitempty"`

	// Must be set to 'vrf'
	Type string `json:"type"`

	// The ID of the VRF in which this VRF IP Reservation is created. The VRF must have an existing IP Range that contains the requested subnet. This field may be aliased as just 'vrf'.
	VrfId string `json:"vrf_id"`
}
