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

type Device struct {
	AlwaysPxe bool `json:"always_pxe,omitempty"`

	BillingCycle string `json:"billing_cycle,omitempty"`

	BondingMode int32 `json:"bonding_mode,omitempty"`

	CreatedAt time.Time `json:"created_at,omitempty"`

	CreatedBy FindDeviceById200ResponseCreatedBy `json:"created_by,omitempty"`

	Customdata map[string]interface{} `json:"customdata,omitempty"`

	Description string `json:"description,omitempty"`

	Facility FindDeviceById200ResponseFacility `json:"facility,omitempty"`

	HardwareReservation FindBatchById200ResponseDevicesInner `json:"hardware_reservation,omitempty"`

	Hostname string `json:"hostname,omitempty"`

	Href string `json:"href,omitempty"`

	Id string `json:"id,omitempty"`

	ImageUrl string `json:"image_url,omitempty"`

	IpAddresses []FindDeviceById200ResponseIpAddressesInner `json:"ip_addresses,omitempty"`

	IpxeScriptUrl string `json:"ipxe_script_url,omitempty"`

	Iqn string `json:"iqn,omitempty"`

	Locked bool `json:"locked,omitempty"`

	Metro FindDeviceById200ResponseFacilityMetro `json:"metro,omitempty"`

	// By default, servers at Equinix Metal are configured in a “bonded” mode using LACP (Link Aggregation Control Protocol). Each 2-NIC server is configured with a single bond (namely bond0) with both interfaces eth0 and eth1 as members of the bond in a default Layer 3 mode. Some device plans may have a different number of ports and bonds available.
	NetworkPorts []FindDeviceById200ResponseNetworkPortsInner `json:"network_ports,omitempty"`

	OperatingSystem FindDeviceById200ResponseOperatingSystem `json:"operating_system,omitempty"`

	// Actions supported by the device instance.
	Actions []FindDeviceById200ResponseActionsInner `json:"actions,omitempty"`

	Plan FindDeviceById200ResponsePlan `json:"plan,omitempty"`

	Project FindDeviceById200ResponseProject `json:"project,omitempty"`

	ProjectLite FindDeviceById200ResponseProjectLite `json:"project_lite,omitempty"`

	ProvisioningEvents []FindInterconnectionEvents200Response `json:"provisioning_events,omitempty"`

	// Only visible while device provisioning
	ProvisioningPercentage float32 `json:"provisioning_percentage,omitempty"`

	// Root password is automatically generated when server is provisioned and it is removed after 24 hours
	RootPassword string `json:"root_password,omitempty"`

	ShortId string `json:"short_id,omitempty"`

	// Whether or not the device is a spot instance.
	SpotInstance bool `json:"spot_instance,omitempty"`

	// The maximum price per hour you are willing to pay to keep this spot instance.  If you are outbid, the termination will be set allowing two minutes before shutdown.
	SpotPriceMax float32 `json:"spot_price_max,omitempty"`

	SshKeys []FindBatchById200ResponseDevicesInner `json:"ssh_keys,omitempty"`

	State string `json:"state,omitempty"`

	// Switch short id. This can be used to determine if two devices are connected to the same switch, for example.
	SwitchUuid string `json:"switch_uuid,omitempty"`

	Tags []string `json:"tags,omitempty"`

	// When the device will be terminated. This is commonly set in advance for ephemeral spot market instances but this field may also be set with on-demand and reservation instances to automatically delete the resource at a given time. The termination time can also be used to release a hardware reservation instance at a given time, keeping the reservation open for other uses.  On a spot market device, the termination time will be set automatically when outbid.
	TerminationTime time.Time `json:"termination_time,omitempty"`

	UpdatedAt time.Time `json:"updated_at,omitempty"`

	User string `json:"user,omitempty"`

	Userdata string `json:"userdata,omitempty"`

	Volumes []FindBatchById200ResponseDevicesInner `json:"volumes,omitempty"`
}