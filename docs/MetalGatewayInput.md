# MetalGatewayInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpReservationId** | Pointer to **string** | The UUID of an IP reservation that belongs to the same project as where the metal gateway will be created in. This field is required unless the private IPv4 subnet size is specified. | [optional] 
**PrivateIpv4SubnetSize** | Pointer to **int32** | The subnet size (8, 16, 32, 64, or 128) of the private IPv4 reservation that will be created for the metal gateway. This field is required unless a project IP reservation was specified.           Please keep in mind that the number of private metal gateway ranges are limited per project. If you would like to increase the limit per project, please contact support for assistance. | [optional] 
**VirtualNetworkId** | **string** | The UUID of a metro virtual network that belongs to the same project as where the metal gateway will be created in. | 

## Methods

### NewMetalGatewayInput

`func NewMetalGatewayInput(virtualNetworkId string, ) *MetalGatewayInput`

NewMetalGatewayInput instantiates a new MetalGatewayInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetalGatewayInputWithDefaults

`func NewMetalGatewayInputWithDefaults() *MetalGatewayInput`

NewMetalGatewayInputWithDefaults instantiates a new MetalGatewayInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpReservationId

`func (o *MetalGatewayInput) GetIpReservationId() string`

GetIpReservationId returns the IpReservationId field if non-nil, zero value otherwise.

### GetIpReservationIdOk

`func (o *MetalGatewayInput) GetIpReservationIdOk() (*string, bool)`

GetIpReservationIdOk returns a tuple with the IpReservationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpReservationId

`func (o *MetalGatewayInput) SetIpReservationId(v string)`

SetIpReservationId sets IpReservationId field to given value.

### HasIpReservationId

`func (o *MetalGatewayInput) HasIpReservationId() bool`

HasIpReservationId returns a boolean if a field has been set.

### GetPrivateIpv4SubnetSize

`func (o *MetalGatewayInput) GetPrivateIpv4SubnetSize() int32`

GetPrivateIpv4SubnetSize returns the PrivateIpv4SubnetSize field if non-nil, zero value otherwise.

### GetPrivateIpv4SubnetSizeOk

`func (o *MetalGatewayInput) GetPrivateIpv4SubnetSizeOk() (*int32, bool)`

GetPrivateIpv4SubnetSizeOk returns a tuple with the PrivateIpv4SubnetSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIpv4SubnetSize

`func (o *MetalGatewayInput) SetPrivateIpv4SubnetSize(v int32)`

SetPrivateIpv4SubnetSize sets PrivateIpv4SubnetSize field to given value.

### HasPrivateIpv4SubnetSize

`func (o *MetalGatewayInput) HasPrivateIpv4SubnetSize() bool`

HasPrivateIpv4SubnetSize returns a boolean if a field has been set.

### GetVirtualNetworkId

`func (o *MetalGatewayInput) GetVirtualNetworkId() string`

GetVirtualNetworkId returns the VirtualNetworkId field if non-nil, zero value otherwise.

### GetVirtualNetworkIdOk

`func (o *MetalGatewayInput) GetVirtualNetworkIdOk() (*string, bool)`

GetVirtualNetworkIdOk returns a tuple with the VirtualNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualNetworkId

`func (o *MetalGatewayInput) SetVirtualNetworkId(v string)`

SetVirtualNetworkId sets VirtualNetworkId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


