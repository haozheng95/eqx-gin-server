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

// Port - Port is a hardware port associated with a reserved or instantiated hardware device.
type Port struct {
	Bond FindDeviceById200ResponseNetworkPortsInnerBond `json:"bond,omitempty"`

	Data FindDeviceById200ResponseNetworkPortsInnerData `json:"data,omitempty"`

	// Indicates whether or not the bond can be broken on the port (when applicable).
	DisbondOperationSupported bool `json:"disbond_operation_supported,omitempty"`

	Href string `json:"href,omitempty"`

	Id string `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	// Type is either \"NetworkBondPort\" for bond ports or \"NetworkPort\" for bondable ethernet ports
	Type string `json:"type,omitempty"`

	// Composite network type of the bond
	NetworkType string `json:"network_type,omitempty"`

	NativeVirtualNetwork FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork `json:"native_virtual_network,omitempty"`

	VirtualNetworks []FindBatchById200ResponseDevicesInner `json:"virtual_networks,omitempty"`
}