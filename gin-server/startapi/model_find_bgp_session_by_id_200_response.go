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

type FindBgpSessionById200Response struct {
	AddressFamily string `json:"address_family"`

	CreatedAt time.Time `json:"created_at,omitempty"`

	DefaultRoute bool `json:"default_route,omitempty"`

	Device FindBatchById200ResponseDevicesInner `json:"device,omitempty"`

	Href string `json:"href,omitempty"`

	Id string `json:"id,omitempty"`

	LearnedRoutes []string `json:"learned_routes,omitempty"`

	//  The status of the BGP Session. Multiple status values may be reported when the device is connected to multiple switches, one value per switch. Each status will start with \"unknown\" and progress to \"up\" or \"down\" depending on the connected device. Subsequent \"unknown\" values indicate a problem acquiring status from the switch.
	Status string `json:"status,omitempty"`

	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
