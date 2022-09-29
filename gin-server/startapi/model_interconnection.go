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

type Interconnection struct {
	ContactEmail string `json:"contact_email,omitempty"`

	Description string `json:"description,omitempty"`

	Facility FindBatchById200ResponseDevicesInner `json:"facility,omitempty"`

	Id string `json:"id,omitempty"`

	Metro GetInterconnection200ResponseMetro `json:"metro,omitempty"`

	// The mode of the interconnection (only relevant to Dedicated Ports). Shared connections won't have this field. Can be either 'standard' or 'tunnel'.   The default mode of an interconnection on a Dedicated Port is 'standard'. The mode can only be changed when there are no associated virtual circuits on the interconnection.   In tunnel mode, an 802.1q tunnel is added to a port to send/receive double tagged packets from server instances.
	Mode string `json:"mode,omitempty"`

	Name string `json:"name,omitempty"`

	Organization FindBatchById200ResponseDevicesInner `json:"organization,omitempty"`

	// For Fabric VCs, these represent Virtual Port(s) created for the interconnection. For dedicated interconnections, these represent the Dedicated Port(s).
	Ports []GetInterconnection200ResponsePortsInner `json:"ports,omitempty"`

	// Either 'primary', meaning a single interconnection, or 'redundant', meaning a redundant interconnection.
	Redundancy string `json:"redundancy,omitempty"`

	// For Fabric VCs (Metal Billed), this will show details of the A-Side service tokens issued for the interconnection. For Fabric VCs (Fabric Billed), this will show the details of the Z-Side service tokens issued for the interconnection. Dedicated interconnections will not have any service tokens issued. There will be one per interconnection, so for redundant interconnections, there should be two service tokens issued. For access to Fabric VCs, which are not generally available, please contact our Support Team for more details.
	ServiceTokens []GetInterconnection200ResponseServiceTokensInner `json:"service_tokens,omitempty"`

	// For interconnections on Dedicated Ports and shared connections, this represents the interconnection's speed in bps. For Fabric VCs, this field refers to the maximum speed of the interconnection in bps. This value will default to 10Gbps for Fabric VCs (Fabric Billed). For access to Fabric VCs, which are not generally available, please contact our Support Team for more details.
	Speed int32 `json:"speed,omitempty"`

	Status string `json:"status,omitempty"`

	Tags []string `json:"tags,omitempty"`

	// This token is used for shared interconnections to be used as the Fabric Token. This field will be deprecated when we release Fabric VCs. With the release of Fabric VCs that use A-Side and Z-Side service tokens, we will no longer issue these tokens for any shared interconnections created after the release of Fabric VCs. This token will also never be issued for dedicated interconnections. For access to Fabric VCs, which are not generally available, please contact our Support Team for more details.
	Token string `json:"token,omitempty"`

	// The 'shared' type of interconnection refers to shared connections, or later also known as Fabric Virtual Connections (or Fabric VCs). The 'dedicated' type of interconnection refers to interconnections created with Dedicated Ports. For access to Fabric VCs, which are not generally available, please contact our Support Team for more details.
	Type string `json:"type,omitempty"`
}