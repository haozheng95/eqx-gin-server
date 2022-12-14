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

type CapacityReport struct {
	Ams1 FindCapacityForFacility200ResponseCapacityAms1 `json:"ams1,omitempty"`

	Atl1 FindCapacityForFacility200ResponseCapacityAtl1 `json:"atl1,omitempty"`

	Dfw1 FindCapacityForFacility200ResponseCapacityAtl1 `json:"dfw1,omitempty"`

	Ewr1 FindCapacityForFacility200ResponseCapacityAms1 `json:"ewr1,omitempty"`

	Fra1 FindCapacityForFacility200ResponseCapacityAtl1 `json:"fra1,omitempty"`

	Iad1 FindCapacityForFacility200ResponseCapacityAtl1 `json:"iad1,omitempty"`

	Lax1 FindCapacityForFacility200ResponseCapacityAtl1 `json:"lax1,omitempty"`

	Nrt1 FindCapacityForFacility200ResponseCapacityAms1 `json:"nrt1,omitempty"`

	Ord1 FindCapacityForFacility200ResponseCapacityAtl1 `json:"ord1,omitempty"`

	Sea1 FindCapacityForFacility200ResponseCapacityAtl1 `json:"sea1,omitempty"`

	Sin1 FindCapacityForFacility200ResponseCapacityAtl1 `json:"sin1,omitempty"`

	Sjc1 FindCapacityForFacility200ResponseCapacityAms1 `json:"sjc1,omitempty"`

	Syd1 FindCapacityForFacility200ResponseCapacityAtl1 `json:"syd1,omitempty"`

	Yyz1 FindCapacityForFacility200ResponseCapacityAtl1 `json:"yyz1,omitempty"`
}
