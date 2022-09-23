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

type GlobalBgpRange struct {
	AddressFamily int32 `json:"address_family,omitempty"`

	Href string `json:"href,omitempty"`

	Id string `json:"id,omitempty"`

	Project FindBatchById200ResponseDevicesInner `json:"project,omitempty"`

	Range string `json:"range,omitempty"`
}