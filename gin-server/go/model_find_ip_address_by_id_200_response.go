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

type FindIpAddressById200Response struct {
	Address string `json:"address,omitempty"`

	AddressFamily int32 `json:"address_family,omitempty"`

	AssignedTo FindBatchById200ResponseDevicesInner `json:"assigned_to,omitempty"`

	Cidr int32 `json:"cidr,omitempty"`

	CreatedAt time.Time `json:"created_at,omitempty"`

	Enabled bool `json:"enabled,omitempty"`

	Gateway string `json:"gateway,omitempty"`

	GlobalIp bool `json:"global_ip,omitempty"`

	Href string `json:"href,omitempty"`

	Id string `json:"id,omitempty"`

	Manageable bool `json:"manageable,omitempty"`

	Management bool `json:"management,omitempty"`

	Metro FindIpAddressById200ResponseOneOfMetro `json:"metro,omitempty"`

	Netmask string `json:"netmask,omitempty"`

	Network string `json:"network,omitempty"`

	ParentBlock FindDeviceById200ResponseIpAddressesInnerParentBlock `json:"parent_block,omitempty"`

	Public bool `json:"public,omitempty"`

	Addon bool `json:"addon,omitempty"`

	Assignments []FindDeviceById200ResponseIpAddressesInner `json:"assignments,omitempty"`

	Bill bool `json:"bill,omitempty"`

	Facility FindIpAddressById200ResponseOneOfFacility `json:"facility,omitempty"`

	MetalGateway FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetworkMetalGatewaysInner `json:"metal_gateway,omitempty"`

	State string `json:"state,omitempty"`

	Tags []string `json:"tags,omitempty"`

	CreatedBy FindBatchById200ResponseDevicesInner `json:"created_by,omitempty"`

	Details string `json:"details,omitempty"`

	Project GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1VrfProject `json:"project,omitempty"`

	Type string `json:"type,omitempty"`

	Vrf GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1Vrf `json:"vrf"`
}
