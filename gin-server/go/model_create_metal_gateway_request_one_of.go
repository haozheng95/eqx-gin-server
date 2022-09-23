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

type CreateMetalGatewayRequestOneOf struct {

	// The UUID of an IP reservation that belongs to the same project as where the metal gateway will be created in. This field is required unless the private IPv4 subnet size is specified.
	IpReservationId string `json:"ip_reservation_id,omitempty"`

	// The subnet size (8, 16, 32, 64, or 128) of the private IPv4 reservation that will be created for the metal gateway. This field is required unless a project IP reservation was specified.           Please keep in mind that the number of private metal gateway ranges are limited per project. If you would like to increase the limit per project, please contact support for assistance.
	PrivateIpv4SubnetSize int32 `json:"private_ipv4_subnet_size,omitempty"`

	// The UUID of a metro virtual network that belongs to the same project as where the metal gateway will be created in.
	VirtualNetworkId string `json:"virtual_network_id"`
}
