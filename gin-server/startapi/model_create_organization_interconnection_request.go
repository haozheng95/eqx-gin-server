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

type CreateOrganizationInterconnectionRequest struct {
	ContactEmail string `json:"contact_email,omitempty"`

	Description string `json:"description,omitempty"`

	// A Metro ID or code. For interconnections with Dedicated Ports, this will be the location of the issued Dedicated Ports. When creating Fabric VCs (Metal Billed), this is where interconnection will be originating from, as we pre-authorize the use of one of our shared ports as the origin of the interconnection using A-Side service tokens. We only allow local connections for Fabric VCs (Metal Billed), so the destination location must be the same as the origin. For Fabric VCs (Fabric Billed), or shared connections, this will be the destination of the interconnection. We allow remote connections for Fabric VCs (Fabric Billed), so the origin of the interconnection can be a different metro set here. For access to Fabric VCs, which are not generally available, please contact our Support Team for more details.
	Metro string `json:"metro"`

	// The mode of the interconnection (only relevant to Dedicated Ports). Fabric VCs won't have this field. Can be either 'standard' or 'tunnel'.   The default mode of an interconnection on a Dedicated Port is 'standard'. The mode can only be changed when there are no associated virtual circuits on the interconnection.   In tunnel mode, an 802.1q tunnel is added to a port to send/receive double tagged packets from server instances.
	Mode string `json:"mode,omitempty"`

	Name string `json:"name"`

	Project string `json:"project,omitempty"`

	// Either 'primary' or 'redundant'.
	Redundancy string `json:"redundancy"`

	// Either 'a_side' or 'z_side'. Setting this field to 'a_side' will create an interconnection with Fabric VCs (Metal Billed). Setting this field to 'z_side' will create an interconnection with Fabric VCs (Fabric Billed). This is required when the 'type' is 'shared', but this is not applicable when the 'type' is 'dedicated'. This parameter is included in the specification as a developer preview and is generally unavailable. Please contact our Support team for more details.
	ServiceTokenType string `json:"service_token_type,omitempty"`

	// A interconnection speed, in bps, mbps, or gbps. For Dedicated Ports, this can be 10Gbps or 100Gbps. For Fabric VCs, this represents the maximum speed of the interconnection. For Fabric VCs (Metal Billed), this can only be one of the following:  ''50mbps'', ''200mbps'', ''500mbps'', ''1gbps'', ''2gbps'', ''5gbps'' or ''10gbps'', and is required for creation. For Fabric VCs (Fabric Billed), this field will always default to ''10gbps'' even if it is not provided. For example, ''500000000'', ''50m'', or' ''500mbps'' will all work as valid inputs.
	Speed int32 `json:"speed,omitempty"`

	Tags []string `json:"tags,omitempty"`

	// Either 'shared' or 'dedicated'. The 'shared' type represents shared interconnections, or also known as Fabric VCs. The 'dedicated' type represents dedicated interconnections, or also known as Dedicated Ports.
	Type string `json:"type"`

	// A list of one or two metro-based VLANs that will be set on the virtual circuits of primary and/or secondary (if redundant) interconnections respectively when creating Fabric VCs. VLANs can also be set after the interconnection is created, but are required to fully activate the interconnection. This parameter is included in the specification as a developer preview and is generally unavailable. Please contact our Support team for more details.
	Vlans []int32 `json:"vlans,omitempty"`
}
