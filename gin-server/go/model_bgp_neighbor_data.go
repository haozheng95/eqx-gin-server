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

type BgpNeighborData struct {

	// Address Family for IP Address. Accepted values are 4 or 6
	AddressFamily float32 `json:"address_family,omitempty"`

	// The customer's ASN. In a local BGP deployment, this will be an internal ASN used to route within the data center. For a global BGP deployment, this will be the your own ASN, configured when you set up BGP for your project.
	CustomerAs float32 `json:"customer_as,omitempty"`

	// The device's IP address. For an IPv4 BGP session, this is typically the private bond0 address for the device.
	CustomerIp string `json:"customer_ip,omitempty"`

	// True if an MD5 password is configured for the project.
	Md5Enabled bool `json:"md5_enabled,omitempty"`

	// The MD5 password configured for the project, if set.
	Md5Password string `json:"md5_password,omitempty"`

	// True when the BGP session should be configured as multihop.
	Multihop bool `json:"multihop,omitempty"`

	// The Peer ASN to use when configuring BGP on your device.
	PeerAs float32 `json:"peer_as,omitempty"`

	// A list of one or more IP addresses to use for the Peer IP section of your BGP configuration. For non-multihop sessions, this will typically be a single gateway address for the device. For multihop sessions, it will be a list of IPs.
	PeerIps []string `json:"peer_ips,omitempty"`

	// A list of project subnets
	RoutesIn []GetBgpNeighborData200ResponseBgpNeighborsInnerRoutesInInner `json:"routes_in,omitempty"`

	// A list of outgoing routes. Only populated if the BGP session has default route enabled.
	RoutesOut []GetBgpNeighborData200ResponseBgpNeighborsInnerRoutesOutInner `json:"routes_out,omitempty"`
}
