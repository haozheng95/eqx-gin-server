/*
Metal API

This is the API for Equinix Metal. The API allows you to programmatically interact with all of your Equinix Metal resources, including devices, networks, addresses, organizations, projects, and your user account.  The official API docs are hosted at <https://metal.equinix.com/developers/api>.

API version: 1.0.0
Contact: support@equinixmetal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
	"time"
)

// CreateDeviceRequestOneOfAllOf1 struct for CreateDeviceRequestOneOfAllOf1
type CreateDeviceRequestOneOfAllOf1 struct {
	// When true, devices with a `custom_ipxe` OS will always boot to iPXE. The default setting of false ensures that iPXE will be used on only the first boot.
	AlwaysPxe *bool `json:"always_pxe,omitempty"`
	// The billing cycle of the device.
	BillingCycle *string `json:"billing_cycle,omitempty"`
	// Customdata is an arbitrary JSON value that can be accessed via the metadata service.
	Customdata interface{} `json:"customdata,omitempty"`
	// Any description of the device or how it will be used. This may be used to inform other API consumers with project access.
	Description *string `json:"description,omitempty"`
	// The features attribute allows you to optionally specify what features your server should have.  In the API shorthand syntax, all features listed are `required`:  ``` { \"features\": [\"tpm\"] } ```  Alternatively, if you do not require a certain feature, but would prefer to be assigned a server with that feature if there are any available, you may specify that feature with a `preferred` value. The request will not fail if we have no servers with that feature in our inventory. The API offers an alternative syntax for mixing preferred and required features:  ``` { \"features\": { \"tpm\": \"required\", \"raid\": \"preferred\" } } ```  The request will only fail if there are no available servers matching the required `tpm` criteria.
	Features []string `json:"features,omitempty"`
	// The Hardware Reservation UUID to provision. Alternatively, `next-available` can be specified to select from any of the available hardware reservations. An error will be returned if the requested reservation option is not available.  See [Reserved Hardware](https://metal.equinix.com/developers/docs/deploy/reserved/) for more details.
	HardwareReservationId *string `json:"hardware_reservation_id,omitempty"`
	// The hostname to use within the operating system. The same hostname may be used on multiple devices within a project.
	Hostname *string `json:"hostname,omitempty"`
	// The `ip_addresses attribute will allow you to specify the addresses you want created with your device.  The default value configures public IPv4, public IPv6, and private IPv4.  Private IPv4 address is required. When specifying `ip_addresses`, one of the array items must enable private IPv4.  Some operating systems require public IPv4 address. In those cases you will receive an error message if public IPv4 is not enabled.  For example, to only configure your server with a private IPv4 address, you can send `{ \"ip_addresses\": [{ \"address_family\": 4, \"public\": false }] }`.  It is possible to request a subnet size larger than a `/30` by assigning addresses using the UUID(s) of ip_reservations in your project.  For example, `{ \"ip_addresses\": [..., {\"address_family\": 4, \"public\": true, \"ip_reservations\": [\"uuid1\", \"uuid2\"]}] }`  To access a server without public IPs, you can use our Out-of-Band console access (SOS) or proxy through another server in the project with public IPs enabled.
	IpAddresses []CreateDeviceRequestOneOfAllOf1IpAddressesInner `json:"ip_addresses,omitempty"`
	// When set, the device will chainload an iPXE Script at boot fetched from the supplied URL.  See [Custom iPXE](https://metal.equinix.com/developers/docs/operating-systems/custom-ipxe/) for more details.
	IpxeScriptUrl *string `json:"ipxe_script_url,omitempty"`
	// Whether the device should be locked, preventing accidental deletion.
	Locked *bool `json:"locked,omitempty"`
	// Overrides default behaviour of attaching all of the organization members ssh keys and project ssh keys to device if no specific keys specified
	NoSshKeys *bool `json:"no_ssh_keys,omitempty"`
	// The slug of the operating system to provision. Check the Equinix Metal operating system documentation for rules that may be imposed per operating system, including restrictions on IP address options and device plans.
	OperatingSystem string `json:"operating_system"`
	// The slug of the device plan to provision.
	Plan string `json:"plan"`
	// Deprecated. Use ip_addresses. Subnet range for addresses allocated to this device.
	PrivateIpv4SubnetSize *float32 `json:"private_ipv4_subnet_size,omitempty"`
	// A list of UUIDs identifying the device parent project that should be authorized to access this device (typically via /root/.ssh/authorized_keys). These keys will also appear in the device metadata.  If no SSH keys are specified (`user_ssh_keys`, `project_ssh_keys`, and `ssh_keys` are all empty lists or omitted), all parent project keys, parent project members keys and organization members keys will be included. This behaviour can be changed with 'no_ssh_keys' option to omit any SSH key being added.
	ProjectSshKeys []string `json:"project_ssh_keys,omitempty"`
	// Deprecated. Use ip_addresses. Subnet range for addresses allocated to this device. Your project must have addresses available for a non-default request.
	PublicIpv4SubnetSize *float32 `json:"public_ipv4_subnet_size,omitempty"`
	SpotInstance         *bool    `json:"spot_instance,omitempty"`
	SpotPriceMax         *float32 `json:"spot_price_max,omitempty"`
	// A list of new or existing project ssh_keys that should be authorized to access this device (typically via /root/.ssh/authorized_keys). These keys will also appear in the device metadata.  These keys are added in addition to any keys defined by   `project_ssh_keys` and `user_ssh_keys`.
	SshKeys         []CreateDeviceRequestOneOfAllOf1SshKeysInner `json:"ssh_keys,omitempty"`
	Tags            []string                                     `json:"tags,omitempty"`
	TerminationTime *time.Time                                   `json:"termination_time,omitempty"`
	// A list of UUIDs identifying the users that should be authorized to access this device (typically via /root/.ssh/authorized_keys).  These keys will also appear in the device metadata.  The users must be members of the project or organization.  If no SSH keys are specified (`user_ssh_keys`, `project_ssh_keys`, and `ssh_keys` are all empty lists or omitted), all parent project keys, parent project members keys and organization members keys will be included. This behaviour can be changed with 'no_ssh_keys' option to omit any SSH key being added.
	UserSshKeys []string `json:"user_ssh_keys,omitempty"`
	// The userdata presented in the metadata service for this device.  Userdata is fetched and interpreted by the operating system installed on the device. Acceptable formats are determined by the operating system, with the exception of a special iPXE enabling syntax which is handled before the operating system starts.  See [Server User Data](https://metal.equinix.com/developers/docs/servers/user-data/) and [Provisioning with Custom iPXE](https://metal.equinix.com/developers/docs/operating-systems/custom-ipxe/#provisioning-with-custom-ipxe) for more details.
	Userdata *string `json:"userdata,omitempty"`
}

// NewCreateDeviceRequestOneOfAllOf1 instantiates a new CreateDeviceRequestOneOfAllOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDeviceRequestOneOfAllOf1(operatingSystem string, plan string) *CreateDeviceRequestOneOfAllOf1 {
	this := CreateDeviceRequestOneOfAllOf1{}
	var alwaysPxe bool = false
	this.AlwaysPxe = &alwaysPxe
	var hardwareReservationId string = ""
	this.HardwareReservationId = &hardwareReservationId
	var locked bool = false
	this.Locked = &locked
	var noSshKeys bool = false
	this.NoSshKeys = &noSshKeys
	this.OperatingSystem = operatingSystem
	this.Plan = plan
	var privateIpv4SubnetSize float32 = 28
	this.PrivateIpv4SubnetSize = &privateIpv4SubnetSize
	var publicIpv4SubnetSize float32 = 31
	this.PublicIpv4SubnetSize = &publicIpv4SubnetSize
	return &this
}

// NewCreateDeviceRequestOneOfAllOf1WithDefaults instantiates a new CreateDeviceRequestOneOfAllOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDeviceRequestOneOfAllOf1WithDefaults() *CreateDeviceRequestOneOfAllOf1 {
	this := CreateDeviceRequestOneOfAllOf1{}
	var alwaysPxe bool = false
	this.AlwaysPxe = &alwaysPxe
	var hardwareReservationId string = ""
	this.HardwareReservationId = &hardwareReservationId
	var locked bool = false
	this.Locked = &locked
	var noSshKeys bool = false
	this.NoSshKeys = &noSshKeys
	var privateIpv4SubnetSize float32 = 28
	this.PrivateIpv4SubnetSize = &privateIpv4SubnetSize
	var publicIpv4SubnetSize float32 = 31
	this.PublicIpv4SubnetSize = &publicIpv4SubnetSize
	return &this
}

// GetAlwaysPxe returns the AlwaysPxe field value if set, zero value otherwise.
func (o *CreateDeviceRequestOneOfAllOf1) GetAlwaysPxe() bool {
	if o == nil || o.AlwaysPxe == nil {
		var ret bool
		return ret
	}
	return *o.AlwaysPxe
}

// GetAlwaysPxeOk returns a tuple with the AlwaysPxe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceRequestOneOfAllOf1) GetAlwaysPxeOk() (*bool, bool) {
	if o == nil || o.AlwaysPxe == nil {
		return nil, false
	}
	return o.AlwaysPxe, true
}

// HasAlwaysPxe returns a boolean if a field has been set.
func (o *CreateDeviceRequestOneOfAllOf1) HasAlwaysPxe() bool {
	if o != nil && o.AlwaysPxe != nil {
		return true
	}

	return false
}

// SetAlwaysPxe gets a reference to the given bool and assigns it to the AlwaysPxe field.
func (o *CreateDeviceRequestOneOfAllOf1) SetAlwaysPxe(v bool) {
	o.AlwaysPxe = &v
}

// GetBillingCycle returns the BillingCycle field value if set, zero value otherwise.
func (o *CreateDeviceRequestOneOfAllOf1) GetBillingCycle() string {
	if o == nil || o.BillingCycle == nil {
		var ret string
		return ret
	}
	return *o.BillingCycle
}

// GetBillingCycleOk returns a tuple with the BillingCycle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceRequestOneOfAllOf1) GetBillingCycleOk() (*string, bool) {
	if o == nil || o.BillingCycle == nil {
		return nil, false
	}
	return o.BillingCycle, true
}

// HasBillingCycle returns a boolean if a field has been set.
func (o *CreateDeviceRequestOneOfAllOf1) HasBillingCycle() bool {
	if o != nil && o.BillingCycle != nil {
		return true
	}

	return false
}

// SetBillingCycle gets a reference to the given string and assigns it to the BillingCycle field.
func (o *CreateDeviceRequestOneOfAllOf1) SetBillingCycle(v string) {
	o.BillingCycle = &v
}

// GetCustomdata returns the Customdata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateDeviceRequestOneOfAllOf1) GetCustomdata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Customdata
}

// GetCustomdataOk returns a tuple with the Customdata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateDeviceRequestOneOfAllOf1) GetCustomdataOk() (*interface{}, bool) {
	if o == nil || o.Customdata == nil {
		return nil, false
	}
	return &o.Customdata, true
}

// HasCustomdata returns a boolean if a field has been set.
func (o *CreateDeviceRequestOneOfAllOf1) HasCustomdata() bool {
	if o != nil && o.Customdata != nil {
		return true
	}

	return false
}

// SetCustomdata gets a reference to the given interface{} and assigns it to the Customdata field.
func (o *CreateDeviceRequestOneOfAllOf1) SetCustomdata(v interface{}) {
	o.Customdata = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateDeviceRequestOneOfAllOf1) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceRequestOneOfAllOf1) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateDeviceRequestOneOfAllOf1) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateDeviceRequestOneOfAllOf1) SetDescription(v string) {
	o.Description = &v
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *CreateDeviceRequestOneOfAllOf1) GetFeatures() []string {
	if o == nil || o.Features == nil {
		var ret []string
		return ret
	}
	return o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceRequestOneOfAllOf1) GetFeaturesOk() ([]string, bool) {
	if o == nil || o.Features == nil {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *CreateDeviceRequestOneOfAllOf1) HasFeatures() bool {
	if o != nil && o.Features != nil {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given []string and assigns it to the Features field.
func (o *CreateDeviceRequestOneOfAllOf1) SetFeatures(v []string) {
	o.Features = v
}

// GetHardwareReservationId returns the HardwareReservationId field value if set, zero value otherwise.
func (o *CreateDeviceRequestOneOfAllOf1) GetHardwareReservationId() string {
	if o == nil || o.HardwareReservationId == nil {
		var ret string
		return ret
	}
	return *o.HardwareReservationId
}

// GetHardwareReservationIdOk returns a tuple with the HardwareReservationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceRequestOneOfAllOf1) GetHardwareReservationIdOk() (*string, bool) {
	if o == nil || o.HardwareReservationId == nil {
		return nil, false
	}
	return o.HardwareReservationId, true
}

// HasHardwareReservationId returns a boolean if a field has been set.
func (o *CreateDeviceRequestOneOfAllOf1) HasHardwareReservationId() bool {
	if o != nil && o.HardwareReservationId != nil {
		return true
	}

	return false
}

// SetHardwareReservationId gets a reference to the given string and assigns it to the HardwareReservationId field.
func (o *CreateDeviceRequestOneOfAllOf1) SetHardwareReservationId(v string) {
	o.HardwareReservationId = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *CreateDeviceRequestOneOfAllOf1) GetHostname() string {
	if o == nil || o.Hostname == nil {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceRequestOneOfAllOf1) GetHostnameOk() (*string, bool) {
	if o == nil || o.Hostname == nil {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *CreateDeviceRequestOneOfAllOf1) HasHostname() bool {
	if o != nil && o.Hostname != nil {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *CreateDeviceRequestOneOfAllOf1) SetHostname(v string) {
	o.Hostname = &v
}

// GetIpAddresses returns the IpAddresses field value if set, zero value otherwise.
func (o *CreateDeviceRequestOneOfAllOf1) GetIpAddresses() []CreateDeviceRequestOneOfAllOf1IpAddressesInner {
	if o == nil || o.IpAddresses == nil {
		var ret []CreateDeviceRequestOneOfAllOf1IpAddressesInner
		return ret
	}
	return o.IpAddresses
}

// GetIpAddressesOk returns a tuple with the IpAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceRequestOneOfAllOf1) GetIpAddressesOk() ([]CreateDeviceRequestOneOfAllOf1IpAddressesInner, bool) {
	if o == nil || o.IpAddresses == nil {
		return nil, false
	}
	return o.IpAddresses, true
}

// HasIpAddresses returns a boolean if a field has been set.
func (o *CreateDeviceRequestOneOfAllOf1) HasIpAddresses() bool {
	if o != nil && o.IpAddresses != nil {
		return true
	}

	return false
}

// SetIpAddresses gets a reference to the given []CreateDeviceRequestOneOfAllOf1IpAddressesInner and assigns it to the IpAddresses field.
func (o *CreateDeviceRequestOneOfAllOf1) SetIpAddresses(v []CreateDeviceRequestOneOfAllOf1IpAddressesInner) {
	o.IpAddresses = v
}

// GetIpxeScriptUrl returns the IpxeScriptUrl field value if set, zero value otherwise.
func (o *CreateDeviceRequestOneOfAllOf1) GetIpxeScriptUrl() string {
	if o == nil || o.IpxeScriptUrl == nil {
		var ret string
		return ret
	}
	return *o.IpxeScriptUrl
}

// GetIpxeScriptUrlOk returns a tuple with the IpxeScriptUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceRequestOneOfAllOf1) GetIpxeScriptUrlOk() (*string, bool) {
	if o == nil || o.IpxeScriptUrl == nil {
		return nil, false
	}
	return o.IpxeScriptUrl, true
}

// HasIpxeScriptUrl returns a boolean if a field has been set.
func (o *CreateDeviceRequestOneOfAllOf1) HasIpxeScriptUrl() bool {
	if o != nil && o.IpxeScriptUrl != nil {
		return true
	}

	return false
}

// SetIpxeScriptUrl gets a reference to the given string and assigns it to the IpxeScriptUrl field.
func (o *CreateDeviceRequestOneOfAllOf1) SetIpxeScriptUrl(v string) {
	o.IpxeScriptUrl = &v
}

// GetLocked returns the Locked field value if set, zero value otherwise.
func (o *CreateDeviceRequestOneOfAllOf1) GetLocked() bool {
	if o == nil || o.Locked == nil {
		var ret bool
		return ret
	}
	return *o.Locked
}

// GetLockedOk returns a tuple with the Locked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceRequestOneOfAllOf1) GetLockedOk() (*bool, bool) {
	if o == nil || o.Locked == nil {
		return nil, false
	}
	return o.Locked, true
}

// HasLocked returns a boolean if a field has been set.
func (o *CreateDeviceRequestOneOfAllOf1) HasLocked() bool {
	if o != nil && o.Locked != nil {
		return true
	}

	return false
}

// SetLocked gets a reference to the given bool and assigns it to the Locked field.
func (o *CreateDeviceRequestOneOfAllOf1) SetLocked(v bool) {
	o.Locked = &v
}

// GetNoSshKeys returns the NoSshKeys field value if set, zero value otherwise.
func (o *CreateDeviceRequestOneOfAllOf1) GetNoSshKeys() bool {
	if o == nil || o.NoSshKeys == nil {
		var ret bool
		return ret
	}
	return *o.NoSshKeys
}

// GetNoSshKeysOk returns a tuple with the NoSshKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceRequestOneOfAllOf1) GetNoSshKeysOk() (*bool, bool) {
	if o == nil || o.NoSshKeys == nil {
		return nil, false
	}
	return o.NoSshKeys, true
}

// HasNoSshKeys returns a boolean if a field has been set.
func (o *CreateDeviceRequestOneOfAllOf1) HasNoSshKeys() bool {
	if o != nil && o.NoSshKeys != nil {
		return true
	}

	return false
}

// SetNoSshKeys gets a reference to the given bool and assigns it to the NoSshKeys field.
func (o *CreateDeviceRequestOneOfAllOf1) SetNoSshKeys(v bool) {
	o.NoSshKeys = &v
}

// GetOperatingSystem returns the OperatingSystem field value
func (o *CreateDeviceRequestOneOfAllOf1) GetOperatingSystem() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OperatingSystem
}

// GetOperatingSystemOk returns a tuple with the OperatingSystem field value
// and a boolean to check if the value has been set.
func (o *CreateDeviceRequestOneOfAllOf1) GetOperatingSystemOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OperatingSystem, true
}

// SetOperatingSystem sets field value
func (o *CreateDeviceRequestOneOfAllOf1) SetOperatingSystem(v string) {
	o.OperatingSystem = v
}

// GetPlan returns the Plan field value
func (o *CreateDeviceRequestOneOfAllOf1) GetPlan() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Plan
}

// GetPlanOk returns a tuple with the Plan field value
// and a boolean to check if the value has been set.
func (o *CreateDeviceRequestOneOfAllOf1) GetPlanOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Plan, true
}

// SetPlan sets field value
func (o *CreateDeviceRequestOneOfAllOf1) SetPlan(v string) {
	o.Plan = v
}

// GetPrivateIpv4SubnetSize returns the PrivateIpv4SubnetSize field value if set, zero value otherwise.
func (o *CreateDeviceRequestOneOfAllOf1) GetPrivateIpv4SubnetSize() float32 {
	if o == nil || o.PrivateIpv4SubnetSize == nil {
		var ret float32
		return ret
	}
	return *o.PrivateIpv4SubnetSize
}

// GetPrivateIpv4SubnetSizeOk returns a tuple with the PrivateIpv4SubnetSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceRequestOneOfAllOf1) GetPrivateIpv4SubnetSizeOk() (*float32, bool) {
	if o == nil || o.PrivateIpv4SubnetSize == nil {
		return nil, false
	}
	return o.PrivateIpv4SubnetSize, true
}

// HasPrivateIpv4SubnetSize returns a boolean if a field has been set.
func (o *CreateDeviceRequestOneOfAllOf1) HasPrivateIpv4SubnetSize() bool {
	if o != nil && o.PrivateIpv4SubnetSize != nil {
		return true
	}

	return false
}

// SetPrivateIpv4SubnetSize gets a reference to the given float32 and assigns it to the PrivateIpv4SubnetSize field.
func (o *CreateDeviceRequestOneOfAllOf1) SetPrivateIpv4SubnetSize(v float32) {
	o.PrivateIpv4SubnetSize = &v
}

// GetProjectSshKeys returns the ProjectSshKeys field value if set, zero value otherwise.
func (o *CreateDeviceRequestOneOfAllOf1) GetProjectSshKeys() []string {
	if o == nil || o.ProjectSshKeys == nil {
		var ret []string
		return ret
	}
	return o.ProjectSshKeys
}

// GetProjectSshKeysOk returns a tuple with the ProjectSshKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceRequestOneOfAllOf1) GetProjectSshKeysOk() ([]string, bool) {
	if o == nil || o.ProjectSshKeys == nil {
		return nil, false
	}
	return o.ProjectSshKeys, true
}

// HasProjectSshKeys returns a boolean if a field has been set.
func (o *CreateDeviceRequestOneOfAllOf1) HasProjectSshKeys() bool {
	if o != nil && o.ProjectSshKeys != nil {
		return true
	}

	return false
}

// SetProjectSshKeys gets a reference to the given []string and assigns it to the ProjectSshKeys field.
func (o *CreateDeviceRequestOneOfAllOf1) SetProjectSshKeys(v []string) {
	o.ProjectSshKeys = v
}

// GetPublicIpv4SubnetSize returns the PublicIpv4SubnetSize field value if set, zero value otherwise.
func (o *CreateDeviceRequestOneOfAllOf1) GetPublicIpv4SubnetSize() float32 {
	if o == nil || o.PublicIpv4SubnetSize == nil {
		var ret float32
		return ret
	}
	return *o.PublicIpv4SubnetSize
}

// GetPublicIpv4SubnetSizeOk returns a tuple with the PublicIpv4SubnetSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceRequestOneOfAllOf1) GetPublicIpv4SubnetSizeOk() (*float32, bool) {
	if o == nil || o.PublicIpv4SubnetSize == nil {
		return nil, false
	}
	return o.PublicIpv4SubnetSize, true
}

// HasPublicIpv4SubnetSize returns a boolean if a field has been set.
func (o *CreateDeviceRequestOneOfAllOf1) HasPublicIpv4SubnetSize() bool {
	if o != nil && o.PublicIpv4SubnetSize != nil {
		return true
	}

	return false
}

// SetPublicIpv4SubnetSize gets a reference to the given float32 and assigns it to the PublicIpv4SubnetSize field.
func (o *CreateDeviceRequestOneOfAllOf1) SetPublicIpv4SubnetSize(v float32) {
	o.PublicIpv4SubnetSize = &v
}

// GetSpotInstance returns the SpotInstance field value if set, zero value otherwise.
func (o *CreateDeviceRequestOneOfAllOf1) GetSpotInstance() bool {
	if o == nil || o.SpotInstance == nil {
		var ret bool
		return ret
	}
	return *o.SpotInstance
}

// GetSpotInstanceOk returns a tuple with the SpotInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceRequestOneOfAllOf1) GetSpotInstanceOk() (*bool, bool) {
	if o == nil || o.SpotInstance == nil {
		return nil, false
	}
	return o.SpotInstance, true
}

// HasSpotInstance returns a boolean if a field has been set.
func (o *CreateDeviceRequestOneOfAllOf1) HasSpotInstance() bool {
	if o != nil && o.SpotInstance != nil {
		return true
	}

	return false
}

// SetSpotInstance gets a reference to the given bool and assigns it to the SpotInstance field.
func (o *CreateDeviceRequestOneOfAllOf1) SetSpotInstance(v bool) {
	o.SpotInstance = &v
}

// GetSpotPriceMax returns the SpotPriceMax field value if set, zero value otherwise.
func (o *CreateDeviceRequestOneOfAllOf1) GetSpotPriceMax() float32 {
	if o == nil || o.SpotPriceMax == nil {
		var ret float32
		return ret
	}
	return *o.SpotPriceMax
}

// GetSpotPriceMaxOk returns a tuple with the SpotPriceMax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceRequestOneOfAllOf1) GetSpotPriceMaxOk() (*float32, bool) {
	if o == nil || o.SpotPriceMax == nil {
		return nil, false
	}
	return o.SpotPriceMax, true
}

// HasSpotPriceMax returns a boolean if a field has been set.
func (o *CreateDeviceRequestOneOfAllOf1) HasSpotPriceMax() bool {
	if o != nil && o.SpotPriceMax != nil {
		return true
	}

	return false
}

// SetSpotPriceMax gets a reference to the given float32 and assigns it to the SpotPriceMax field.
func (o *CreateDeviceRequestOneOfAllOf1) SetSpotPriceMax(v float32) {
	o.SpotPriceMax = &v
}

// GetSshKeys returns the SshKeys field value if set, zero value otherwise.
func (o *CreateDeviceRequestOneOfAllOf1) GetSshKeys() []CreateDeviceRequestOneOfAllOf1SshKeysInner {
	if o == nil || o.SshKeys == nil {
		var ret []CreateDeviceRequestOneOfAllOf1SshKeysInner
		return ret
	}
	return o.SshKeys
}

// GetSshKeysOk returns a tuple with the SshKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceRequestOneOfAllOf1) GetSshKeysOk() ([]CreateDeviceRequestOneOfAllOf1SshKeysInner, bool) {
	if o == nil || o.SshKeys == nil {
		return nil, false
	}
	return o.SshKeys, true
}

// HasSshKeys returns a boolean if a field has been set.
func (o *CreateDeviceRequestOneOfAllOf1) HasSshKeys() bool {
	if o != nil && o.SshKeys != nil {
		return true
	}

	return false
}

// SetSshKeys gets a reference to the given []CreateDeviceRequestOneOfAllOf1SshKeysInner and assigns it to the SshKeys field.
func (o *CreateDeviceRequestOneOfAllOf1) SetSshKeys(v []CreateDeviceRequestOneOfAllOf1SshKeysInner) {
	o.SshKeys = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *CreateDeviceRequestOneOfAllOf1) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceRequestOneOfAllOf1) GetTagsOk() ([]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *CreateDeviceRequestOneOfAllOf1) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *CreateDeviceRequestOneOfAllOf1) SetTags(v []string) {
	o.Tags = v
}

// GetTerminationTime returns the TerminationTime field value if set, zero value otherwise.
func (o *CreateDeviceRequestOneOfAllOf1) GetTerminationTime() time.Time {
	if o == nil || o.TerminationTime == nil {
		var ret time.Time
		return ret
	}
	return *o.TerminationTime
}

// GetTerminationTimeOk returns a tuple with the TerminationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceRequestOneOfAllOf1) GetTerminationTimeOk() (*time.Time, bool) {
	if o == nil || o.TerminationTime == nil {
		return nil, false
	}
	return o.TerminationTime, true
}

// HasTerminationTime returns a boolean if a field has been set.
func (o *CreateDeviceRequestOneOfAllOf1) HasTerminationTime() bool {
	if o != nil && o.TerminationTime != nil {
		return true
	}

	return false
}

// SetTerminationTime gets a reference to the given time.Time and assigns it to the TerminationTime field.
func (o *CreateDeviceRequestOneOfAllOf1) SetTerminationTime(v time.Time) {
	o.TerminationTime = &v
}

// GetUserSshKeys returns the UserSshKeys field value if set, zero value otherwise.
func (o *CreateDeviceRequestOneOfAllOf1) GetUserSshKeys() []string {
	if o == nil || o.UserSshKeys == nil {
		var ret []string
		return ret
	}
	return o.UserSshKeys
}

// GetUserSshKeysOk returns a tuple with the UserSshKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceRequestOneOfAllOf1) GetUserSshKeysOk() ([]string, bool) {
	if o == nil || o.UserSshKeys == nil {
		return nil, false
	}
	return o.UserSshKeys, true
}

// HasUserSshKeys returns a boolean if a field has been set.
func (o *CreateDeviceRequestOneOfAllOf1) HasUserSshKeys() bool {
	if o != nil && o.UserSshKeys != nil {
		return true
	}

	return false
}

// SetUserSshKeys gets a reference to the given []string and assigns it to the UserSshKeys field.
func (o *CreateDeviceRequestOneOfAllOf1) SetUserSshKeys(v []string) {
	o.UserSshKeys = v
}

// GetUserdata returns the Userdata field value if set, zero value otherwise.
func (o *CreateDeviceRequestOneOfAllOf1) GetUserdata() string {
	if o == nil || o.Userdata == nil {
		var ret string
		return ret
	}
	return *o.Userdata
}

// GetUserdataOk returns a tuple with the Userdata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceRequestOneOfAllOf1) GetUserdataOk() (*string, bool) {
	if o == nil || o.Userdata == nil {
		return nil, false
	}
	return o.Userdata, true
}

// HasUserdata returns a boolean if a field has been set.
func (o *CreateDeviceRequestOneOfAllOf1) HasUserdata() bool {
	if o != nil && o.Userdata != nil {
		return true
	}

	return false
}

// SetUserdata gets a reference to the given string and assigns it to the Userdata field.
func (o *CreateDeviceRequestOneOfAllOf1) SetUserdata(v string) {
	o.Userdata = &v
}

func (o CreateDeviceRequestOneOfAllOf1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AlwaysPxe != nil {
		toSerialize["always_pxe"] = o.AlwaysPxe
	}
	if o.BillingCycle != nil {
		toSerialize["billing_cycle"] = o.BillingCycle
	}
	if o.Customdata != nil {
		toSerialize["customdata"] = o.Customdata
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Features != nil {
		toSerialize["features"] = o.Features
	}
	if o.HardwareReservationId != nil {
		toSerialize["hardware_reservation_id"] = o.HardwareReservationId
	}
	if o.Hostname != nil {
		toSerialize["hostname"] = o.Hostname
	}
	if o.IpAddresses != nil {
		toSerialize["ip_addresses"] = o.IpAddresses
	}
	if o.IpxeScriptUrl != nil {
		toSerialize["ipxe_script_url"] = o.IpxeScriptUrl
	}
	if o.Locked != nil {
		toSerialize["locked"] = o.Locked
	}
	if o.NoSshKeys != nil {
		toSerialize["no_ssh_keys"] = o.NoSshKeys
	}
	if true {
		toSerialize["operating_system"] = o.OperatingSystem
	}
	if true {
		toSerialize["plan"] = o.Plan
	}
	if o.PrivateIpv4SubnetSize != nil {
		toSerialize["private_ipv4_subnet_size"] = o.PrivateIpv4SubnetSize
	}
	if o.ProjectSshKeys != nil {
		toSerialize["project_ssh_keys"] = o.ProjectSshKeys
	}
	if o.PublicIpv4SubnetSize != nil {
		toSerialize["public_ipv4_subnet_size"] = o.PublicIpv4SubnetSize
	}
	if o.SpotInstance != nil {
		toSerialize["spot_instance"] = o.SpotInstance
	}
	if o.SpotPriceMax != nil {
		toSerialize["spot_price_max"] = o.SpotPriceMax
	}
	if o.SshKeys != nil {
		toSerialize["ssh_keys"] = o.SshKeys
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.TerminationTime != nil {
		toSerialize["termination_time"] = o.TerminationTime
	}
	if o.UserSshKeys != nil {
		toSerialize["user_ssh_keys"] = o.UserSshKeys
	}
	if o.Userdata != nil {
		toSerialize["userdata"] = o.Userdata
	}
	return json.Marshal(toSerialize)
}

type NullableCreateDeviceRequestOneOfAllOf1 struct {
	value *CreateDeviceRequestOneOfAllOf1
	isSet bool
}

func (v NullableCreateDeviceRequestOneOfAllOf1) Get() *CreateDeviceRequestOneOfAllOf1 {
	return v.value
}

func (v *NullableCreateDeviceRequestOneOfAllOf1) Set(val *CreateDeviceRequestOneOfAllOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDeviceRequestOneOfAllOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDeviceRequestOneOfAllOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDeviceRequestOneOfAllOf1(val *CreateDeviceRequestOneOfAllOf1) *NullableCreateDeviceRequestOneOfAllOf1 {
	return &NullableCreateDeviceRequestOneOfAllOf1{value: val, isSet: true}
}

func (v NullableCreateDeviceRequestOneOfAllOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDeviceRequestOneOfAllOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
