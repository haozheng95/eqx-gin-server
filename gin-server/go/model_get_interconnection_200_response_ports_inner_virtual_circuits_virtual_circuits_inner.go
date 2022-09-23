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

type GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInner struct {

	// True if the Virtual Circuit is being billed. Currently, only Virtual Circuits of Fabric VCs (Metal Billed) will be billed. Usage will start the first time the Virtual Circuit becomes active, and will not stop until it is deleted from Metal.
	Bill bool `json:"bill"`

	Description string `json:"description"`

	Id string `json:"id"`

	Name string `json:"name"`

	NniVlan int32 `json:"nni_vlan"`

	Port FindBatchById200ResponseDevicesInner `json:"port"`

	Project FindBatchById200ResponseDevicesInner `json:"project"`

	// integer representing bps speed
	Speed int32 `json:"speed,omitempty"`

	Status string `json:"status"`

	Tags []string `json:"tags"`

	VirtualNetwork FindBatchById200ResponseDevicesInner `json:"virtual_network"`

	Vnid int32 `json:"vnid"`

	// An IP address from the subnet that will be used on the Customer side. This parameter is optional, but if supplied, we will use the other usable IP address in the subnet as the Metal IP. By default, the last usable IP address in the subnet will be used.
	CustomerIp string `json:"customer_ip,omitempty"`

	// The MD5 password for the BGP peering in plaintext (not a checksum).
	Md5 string `json:"md5,omitempty"`

	// An IP address from the subnet that will be used on the Metal side. This parameter is optional, but if supplied, we will use the other usable IP address in the subnet as the Customer IP. By default, the first usable IP address in the subnet will be used.
	MetalIp string `json:"metal_ip,omitempty"`

	// The peer ASN that will be used with the VRF on the Virtual Circuit.
	PeerAsn int32 `json:"peer_asn,omitempty"`

	// The /30 or /31 subnet of one of the VRF IP Blocks that will be used with the VRF for the Virtual Circuit. This subnet does not have to be an existing VRF IP reservation, as we will create the VRF IP reservation on creation if it does not exist. The Metal IP and Customer IP must be IPs from this subnet. For /30 subnets, the network and broadcast IPs cannot be used as the Metal or Customer IP.
	Subnet string `json:"subnet,omitempty"`

	Vrf GetInterconnection200ResponsePortsInnerVirtualCircuitsVirtualCircuitsInnerAnyOf1Vrf `json:"vrf,omitempty"`
}
