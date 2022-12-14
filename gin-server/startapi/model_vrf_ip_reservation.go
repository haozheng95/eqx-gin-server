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

type VrfIpReservation struct {
	AddressFamily int32 `json:"address_family,omitempty"`

	Cidr int32 `json:"cidr,omitempty"`

	CreatedAt time.Time `json:"created_at,omitempty"`

	CreatedBy FindBatchById200ResponseDevicesInner `json:"created_by,omitempty"`

	Details string `json:"details,omitempty"`

	Href string `json:"href,omitempty"`

	Id string `json:"id,omitempty"`

	MetalGateway FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetworkMetalGatewaysInner `json:"metal_gateway,omitempty"`

	Netmask string `json:"netmask,omitempty"`

	Network string `json:"network,omitempty"`

	Project GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1VrfProject `json:"project,omitempty"`

	State string `json:"state,omitempty"`

	Tags []string `json:"tags,omitempty"`

	Type string `json:"type,omitempty"`

	Vrf GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1Vrf `json:"vrf"`
}
