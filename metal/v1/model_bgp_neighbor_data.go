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
)

// BgpNeighborData struct for BgpNeighborData
type BgpNeighborData struct {
	// Address Family for IP Address. Accepted values are 4 or 6
	AddressFamily *float32 `json:"address_family,omitempty"`
	// The customer's ASN. In a local BGP deployment, this will be an internal ASN used to route within the data center. For a global BGP deployment, this will be the your own ASN, configured when you set up BGP for your project.
	CustomerAs *float32 `json:"customer_as,omitempty"`
	// The device's IP address. For an IPv4 BGP session, this is typically the private bond0 address for the device.
	CustomerIp *string `json:"customer_ip,omitempty"`
	// True if an MD5 password is configured for the project.
	Md5Enabled *bool `json:"md5_enabled,omitempty"`
	// The MD5 password configured for the project, if set.
	Md5Password *string `json:"md5_password,omitempty"`
	// True when the BGP session should be configured as multihop.
	Multihop *bool `json:"multihop,omitempty"`
	// The Peer ASN to use when configuring BGP on your device.
	PeerAs *float32 `json:"peer_as,omitempty"`
	// A list of one or more IP addresses to use for the Peer IP section of your BGP configuration. For non-multihop sessions, this will typically be a single gateway address for the device. For multihop sessions, it will be a list of IPs.
	PeerIps []string `json:"peer_ips,omitempty"`
	// A list of project subnets
	RoutesIn []GetBgpNeighborData200ResponseBgpNeighborsInnerRoutesInInner `json:"routes_in,omitempty"`
	// A list of outgoing routes. Only populated if the BGP session has default route enabled.
	RoutesOut []GetBgpNeighborData200ResponseBgpNeighborsInnerRoutesOutInner `json:"routes_out,omitempty"`
}

// NewBgpNeighborData instantiates a new BgpNeighborData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBgpNeighborData() *BgpNeighborData {
	this := BgpNeighborData{}
	return &this
}

// NewBgpNeighborDataWithDefaults instantiates a new BgpNeighborData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBgpNeighborDataWithDefaults() *BgpNeighborData {
	this := BgpNeighborData{}
	return &this
}

// GetAddressFamily returns the AddressFamily field value if set, zero value otherwise.
func (o *BgpNeighborData) GetAddressFamily() float32 {
	if o == nil || o.AddressFamily == nil {
		var ret float32
		return ret
	}
	return *o.AddressFamily
}

// GetAddressFamilyOk returns a tuple with the AddressFamily field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpNeighborData) GetAddressFamilyOk() (*float32, bool) {
	if o == nil || o.AddressFamily == nil {
		return nil, false
	}
	return o.AddressFamily, true
}

// HasAddressFamily returns a boolean if a field has been set.
func (o *BgpNeighborData) HasAddressFamily() bool {
	if o != nil && o.AddressFamily != nil {
		return true
	}

	return false
}

// SetAddressFamily gets a reference to the given float32 and assigns it to the AddressFamily field.
func (o *BgpNeighborData) SetAddressFamily(v float32) {
	o.AddressFamily = &v
}

// GetCustomerAs returns the CustomerAs field value if set, zero value otherwise.
func (o *BgpNeighborData) GetCustomerAs() float32 {
	if o == nil || o.CustomerAs == nil {
		var ret float32
		return ret
	}
	return *o.CustomerAs
}

// GetCustomerAsOk returns a tuple with the CustomerAs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpNeighborData) GetCustomerAsOk() (*float32, bool) {
	if o == nil || o.CustomerAs == nil {
		return nil, false
	}
	return o.CustomerAs, true
}

// HasCustomerAs returns a boolean if a field has been set.
func (o *BgpNeighborData) HasCustomerAs() bool {
	if o != nil && o.CustomerAs != nil {
		return true
	}

	return false
}

// SetCustomerAs gets a reference to the given float32 and assigns it to the CustomerAs field.
func (o *BgpNeighborData) SetCustomerAs(v float32) {
	o.CustomerAs = &v
}

// GetCustomerIp returns the CustomerIp field value if set, zero value otherwise.
func (o *BgpNeighborData) GetCustomerIp() string {
	if o == nil || o.CustomerIp == nil {
		var ret string
		return ret
	}
	return *o.CustomerIp
}

// GetCustomerIpOk returns a tuple with the CustomerIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpNeighborData) GetCustomerIpOk() (*string, bool) {
	if o == nil || o.CustomerIp == nil {
		return nil, false
	}
	return o.CustomerIp, true
}

// HasCustomerIp returns a boolean if a field has been set.
func (o *BgpNeighborData) HasCustomerIp() bool {
	if o != nil && o.CustomerIp != nil {
		return true
	}

	return false
}

// SetCustomerIp gets a reference to the given string and assigns it to the CustomerIp field.
func (o *BgpNeighborData) SetCustomerIp(v string) {
	o.CustomerIp = &v
}

// GetMd5Enabled returns the Md5Enabled field value if set, zero value otherwise.
func (o *BgpNeighborData) GetMd5Enabled() bool {
	if o == nil || o.Md5Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Md5Enabled
}

// GetMd5EnabledOk returns a tuple with the Md5Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpNeighborData) GetMd5EnabledOk() (*bool, bool) {
	if o == nil || o.Md5Enabled == nil {
		return nil, false
	}
	return o.Md5Enabled, true
}

// HasMd5Enabled returns a boolean if a field has been set.
func (o *BgpNeighborData) HasMd5Enabled() bool {
	if o != nil && o.Md5Enabled != nil {
		return true
	}

	return false
}

// SetMd5Enabled gets a reference to the given bool and assigns it to the Md5Enabled field.
func (o *BgpNeighborData) SetMd5Enabled(v bool) {
	o.Md5Enabled = &v
}

// GetMd5Password returns the Md5Password field value if set, zero value otherwise.
func (o *BgpNeighborData) GetMd5Password() string {
	if o == nil || o.Md5Password == nil {
		var ret string
		return ret
	}
	return *o.Md5Password
}

// GetMd5PasswordOk returns a tuple with the Md5Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpNeighborData) GetMd5PasswordOk() (*string, bool) {
	if o == nil || o.Md5Password == nil {
		return nil, false
	}
	return o.Md5Password, true
}

// HasMd5Password returns a boolean if a field has been set.
func (o *BgpNeighborData) HasMd5Password() bool {
	if o != nil && o.Md5Password != nil {
		return true
	}

	return false
}

// SetMd5Password gets a reference to the given string and assigns it to the Md5Password field.
func (o *BgpNeighborData) SetMd5Password(v string) {
	o.Md5Password = &v
}

// GetMultihop returns the Multihop field value if set, zero value otherwise.
func (o *BgpNeighborData) GetMultihop() bool {
	if o == nil || o.Multihop == nil {
		var ret bool
		return ret
	}
	return *o.Multihop
}

// GetMultihopOk returns a tuple with the Multihop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpNeighborData) GetMultihopOk() (*bool, bool) {
	if o == nil || o.Multihop == nil {
		return nil, false
	}
	return o.Multihop, true
}

// HasMultihop returns a boolean if a field has been set.
func (o *BgpNeighborData) HasMultihop() bool {
	if o != nil && o.Multihop != nil {
		return true
	}

	return false
}

// SetMultihop gets a reference to the given bool and assigns it to the Multihop field.
func (o *BgpNeighborData) SetMultihop(v bool) {
	o.Multihop = &v
}

// GetPeerAs returns the PeerAs field value if set, zero value otherwise.
func (o *BgpNeighborData) GetPeerAs() float32 {
	if o == nil || o.PeerAs == nil {
		var ret float32
		return ret
	}
	return *o.PeerAs
}

// GetPeerAsOk returns a tuple with the PeerAs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpNeighborData) GetPeerAsOk() (*float32, bool) {
	if o == nil || o.PeerAs == nil {
		return nil, false
	}
	return o.PeerAs, true
}

// HasPeerAs returns a boolean if a field has been set.
func (o *BgpNeighborData) HasPeerAs() bool {
	if o != nil && o.PeerAs != nil {
		return true
	}

	return false
}

// SetPeerAs gets a reference to the given float32 and assigns it to the PeerAs field.
func (o *BgpNeighborData) SetPeerAs(v float32) {
	o.PeerAs = &v
}

// GetPeerIps returns the PeerIps field value if set, zero value otherwise.
func (o *BgpNeighborData) GetPeerIps() []string {
	if o == nil || o.PeerIps == nil {
		var ret []string
		return ret
	}
	return o.PeerIps
}

// GetPeerIpsOk returns a tuple with the PeerIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpNeighborData) GetPeerIpsOk() ([]string, bool) {
	if o == nil || o.PeerIps == nil {
		return nil, false
	}
	return o.PeerIps, true
}

// HasPeerIps returns a boolean if a field has been set.
func (o *BgpNeighborData) HasPeerIps() bool {
	if o != nil && o.PeerIps != nil {
		return true
	}

	return false
}

// SetPeerIps gets a reference to the given []string and assigns it to the PeerIps field.
func (o *BgpNeighborData) SetPeerIps(v []string) {
	o.PeerIps = v
}

// GetRoutesIn returns the RoutesIn field value if set, zero value otherwise.
func (o *BgpNeighborData) GetRoutesIn() []GetBgpNeighborData200ResponseBgpNeighborsInnerRoutesInInner {
	if o == nil || o.RoutesIn == nil {
		var ret []GetBgpNeighborData200ResponseBgpNeighborsInnerRoutesInInner
		return ret
	}
	return o.RoutesIn
}

// GetRoutesInOk returns a tuple with the RoutesIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpNeighborData) GetRoutesInOk() ([]GetBgpNeighborData200ResponseBgpNeighborsInnerRoutesInInner, bool) {
	if o == nil || o.RoutesIn == nil {
		return nil, false
	}
	return o.RoutesIn, true
}

// HasRoutesIn returns a boolean if a field has been set.
func (o *BgpNeighborData) HasRoutesIn() bool {
	if o != nil && o.RoutesIn != nil {
		return true
	}

	return false
}

// SetRoutesIn gets a reference to the given []GetBgpNeighborData200ResponseBgpNeighborsInnerRoutesInInner and assigns it to the RoutesIn field.
func (o *BgpNeighborData) SetRoutesIn(v []GetBgpNeighborData200ResponseBgpNeighborsInnerRoutesInInner) {
	o.RoutesIn = v
}

// GetRoutesOut returns the RoutesOut field value if set, zero value otherwise.
func (o *BgpNeighborData) GetRoutesOut() []GetBgpNeighborData200ResponseBgpNeighborsInnerRoutesOutInner {
	if o == nil || o.RoutesOut == nil {
		var ret []GetBgpNeighborData200ResponseBgpNeighborsInnerRoutesOutInner
		return ret
	}
	return o.RoutesOut
}

// GetRoutesOutOk returns a tuple with the RoutesOut field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpNeighborData) GetRoutesOutOk() ([]GetBgpNeighborData200ResponseBgpNeighborsInnerRoutesOutInner, bool) {
	if o == nil || o.RoutesOut == nil {
		return nil, false
	}
	return o.RoutesOut, true
}

// HasRoutesOut returns a boolean if a field has been set.
func (o *BgpNeighborData) HasRoutesOut() bool {
	if o != nil && o.RoutesOut != nil {
		return true
	}

	return false
}

// SetRoutesOut gets a reference to the given []GetBgpNeighborData200ResponseBgpNeighborsInnerRoutesOutInner and assigns it to the RoutesOut field.
func (o *BgpNeighborData) SetRoutesOut(v []GetBgpNeighborData200ResponseBgpNeighborsInnerRoutesOutInner) {
	o.RoutesOut = v
}

func (o BgpNeighborData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AddressFamily != nil {
		toSerialize["address_family"] = o.AddressFamily
	}
	if o.CustomerAs != nil {
		toSerialize["customer_as"] = o.CustomerAs
	}
	if o.CustomerIp != nil {
		toSerialize["customer_ip"] = o.CustomerIp
	}
	if o.Md5Enabled != nil {
		toSerialize["md5_enabled"] = o.Md5Enabled
	}
	if o.Md5Password != nil {
		toSerialize["md5_password"] = o.Md5Password
	}
	if o.Multihop != nil {
		toSerialize["multihop"] = o.Multihop
	}
	if o.PeerAs != nil {
		toSerialize["peer_as"] = o.PeerAs
	}
	if o.PeerIps != nil {
		toSerialize["peer_ips"] = o.PeerIps
	}
	if o.RoutesIn != nil {
		toSerialize["routes_in"] = o.RoutesIn
	}
	if o.RoutesOut != nil {
		toSerialize["routes_out"] = o.RoutesOut
	}
	return json.Marshal(toSerialize)
}

type NullableBgpNeighborData struct {
	value *BgpNeighborData
	isSet bool
}

func (v NullableBgpNeighborData) Get() *BgpNeighborData {
	return v.value
}

func (v *NullableBgpNeighborData) Set(val *BgpNeighborData) {
	v.value = val
	v.isSet = true
}

func (v NullableBgpNeighborData) IsSet() bool {
	return v.isSet
}

func (v *NullableBgpNeighborData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBgpNeighborData(val *BgpNeighborData) *NullableBgpNeighborData {
	return &NullableBgpNeighborData{value: val, isSet: true}
}

func (v NullableBgpNeighborData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBgpNeighborData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
