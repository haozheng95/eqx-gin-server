// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/t0mk/gometal/client/batches"
	"github.com/t0mk/gometal/client/bgp"
	"github.com/t0mk/gometal/client/capacity"
	"github.com/t0mk/gometal/client/connections"
	"github.com/t0mk/gometal/client/devices"
	"github.com/t0mk/gometal/client/emails"
	"github.com/t0mk/gometal/client/events"
	"github.com/t0mk/gometal/client/facilities"
	"github.com/t0mk/gometal/client/hardware_reservations"
	"github.com/t0mk/gometal/client/incidents"
	"github.com/t0mk/gometal/client/internet_gateways"
	"github.com/t0mk/gometal/client/invitations"
	"github.com/t0mk/gometal/client/ip_addresses"
	"github.com/t0mk/gometal/client/licenses"
	"github.com/t0mk/gometal/client/market"
	"github.com/t0mk/gometal/client/memberships"
	"github.com/t0mk/gometal/client/operating_system_versions"
	"github.com/t0mk/gometal/client/operating_systems"
	"github.com/t0mk/gometal/client/organizations"
	"github.com/t0mk/gometal/client/otps"
	"github.com/t0mk/gometal/client/password_reset_tokens"
	"github.com/t0mk/gometal/client/payment_methods"
	"github.com/t0mk/gometal/client/plans"
	"github.com/t0mk/gometal/client/ports"
	"github.com/t0mk/gometal/client/projects"
	"github.com/t0mk/gometal/client/regions"
	"github.com/t0mk/gometal/client/spot_market_request"
	"github.com/t0mk/gometal/client/ssh_keys"
	"github.com/t0mk/gometal/client/transfer_requests"
	"github.com/t0mk/gometal/client/two_factor_auth"
	"github.com/t0mk/gometal/client/user_verification_tokens"
	"github.com/t0mk/gometal/client/userdata"
	"github.com/t0mk/gometal/client/users"
	"github.com/t0mk/gometal/client/v_l_a_ns"
	"github.com/t0mk/gometal/client/volumes"
	"github.com/t0mk/gometal/client/vpn"
)

// Default metal API HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "api.equinix.com"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/metal/v1"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

// NewHTTPClient creates a new metal API HTTP client.
func NewHTTPClient(formats strfmt.Registry) *MetalAPI {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new metal API HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *MetalAPI {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new metal API client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *MetalAPI {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(MetalAPI)
	cli.Transport = transport
	cli.Batches = batches.New(transport, formats)
	cli.BGP = bgp.New(transport, formats)
	cli.Capacity = capacity.New(transport, formats)
	cli.Connections = connections.New(transport, formats)
	cli.Devices = devices.New(transport, formats)
	cli.Emails = emails.New(transport, formats)
	cli.Events = events.New(transport, formats)
	cli.Facilities = facilities.New(transport, formats)
	cli.HardwareReservations = hardware_reservations.New(transport, formats)
	cli.Incidents = incidents.New(transport, formats)
	cli.InternetGateways = internet_gateways.New(transport, formats)
	cli.Invitations = invitations.New(transport, formats)
	cli.IPAddresses = ip_addresses.New(transport, formats)
	cli.Licenses = licenses.New(transport, formats)
	cli.Market = market.New(transport, formats)
	cli.Memberships = memberships.New(transport, formats)
	cli.OperatingSystemVersions = operating_system_versions.New(transport, formats)
	cli.OperatingSystems = operating_systems.New(transport, formats)
	cli.Organizations = organizations.New(transport, formats)
	cli.Otps = otps.New(transport, formats)
	cli.PasswordResetTokens = password_reset_tokens.New(transport, formats)
	cli.PaymentMethods = payment_methods.New(transport, formats)
	cli.Plans = plans.New(transport, formats)
	cli.Ports = ports.New(transport, formats)
	cli.Projects = projects.New(transport, formats)
	cli.Regions = regions.New(transport, formats)
	cli.SpotMarketRequest = spot_market_request.New(transport, formats)
	cli.SSHKeys = ssh_keys.New(transport, formats)
	cli.TransferRequests = transfer_requests.New(transport, formats)
	cli.TwoFactorAuth = two_factor_auth.New(transport, formats)
	cli.UserVerificationTokens = user_verification_tokens.New(transport, formats)
	cli.Userdata = userdata.New(transport, formats)
	cli.Users = users.New(transport, formats)
	cli.VlaNs = v_l_a_ns.New(transport, formats)
	cli.Volumes = volumes.New(transport, formats)
	cli.VPN = vpn.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// MetalAPI is a client for metal API
type MetalAPI struct {
	Batches batches.ClientService

	BGP bgp.ClientService

	Capacity capacity.ClientService

	Connections connections.ClientService

	Devices devices.ClientService

	Emails emails.ClientService

	Events events.ClientService

	Facilities facilities.ClientService

	HardwareReservations hardware_reservations.ClientService

	Incidents incidents.ClientService

	InternetGateways internet_gateways.ClientService

	Invitations invitations.ClientService

	IPAddresses ip_addresses.ClientService

	Licenses licenses.ClientService

	Market market.ClientService

	Memberships memberships.ClientService

	OperatingSystemVersions operating_system_versions.ClientService

	OperatingSystems operating_systems.ClientService

	Organizations organizations.ClientService

	Otps otps.ClientService

	PasswordResetTokens password_reset_tokens.ClientService

	PaymentMethods payment_methods.ClientService

	Plans plans.ClientService

	Ports ports.ClientService

	Projects projects.ClientService

	Regions regions.ClientService

	SpotMarketRequest spot_market_request.ClientService

	SSHKeys ssh_keys.ClientService

	TransferRequests transfer_requests.ClientService

	TwoFactorAuth two_factor_auth.ClientService

	UserVerificationTokens user_verification_tokens.ClientService

	Userdata userdata.ClientService

	Users users.ClientService

	VlaNs v_l_a_ns.ClientService

	Volumes volumes.ClientService

	VPN vpn.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *MetalAPI) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.Batches.SetTransport(transport)
	c.BGP.SetTransport(transport)
	c.Capacity.SetTransport(transport)
	c.Connections.SetTransport(transport)
	c.Devices.SetTransport(transport)
	c.Emails.SetTransport(transport)
	c.Events.SetTransport(transport)
	c.Facilities.SetTransport(transport)
	c.HardwareReservations.SetTransport(transport)
	c.Incidents.SetTransport(transport)
	c.InternetGateways.SetTransport(transport)
	c.Invitations.SetTransport(transport)
	c.IPAddresses.SetTransport(transport)
	c.Licenses.SetTransport(transport)
	c.Market.SetTransport(transport)
	c.Memberships.SetTransport(transport)
	c.OperatingSystemVersions.SetTransport(transport)
	c.OperatingSystems.SetTransport(transport)
	c.Organizations.SetTransport(transport)
	c.Otps.SetTransport(transport)
	c.PasswordResetTokens.SetTransport(transport)
	c.PaymentMethods.SetTransport(transport)
	c.Plans.SetTransport(transport)
	c.Ports.SetTransport(transport)
	c.Projects.SetTransport(transport)
	c.Regions.SetTransport(transport)
	c.SpotMarketRequest.SetTransport(transport)
	c.SSHKeys.SetTransport(transport)
	c.TransferRequests.SetTransport(transport)
	c.TwoFactorAuth.SetTransport(transport)
	c.UserVerificationTokens.SetTransport(transport)
	c.Userdata.SetTransport(transport)
	c.Users.SetTransport(transport)
	c.VlaNs.SetTransport(transport)
	c.Volumes.SetTransport(transport)
	c.VPN.SetTransport(transport)
}
