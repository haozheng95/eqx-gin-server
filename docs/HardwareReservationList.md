# HardwareReservationList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HardwareReservations** | Pointer to [**[]MoveHardwareReservation201Response**](MoveHardwareReservation201Response.md) |  | [optional] 
**Meta** | Pointer to [**FindDeviceEvents200ResponseMeta**](FindDeviceEvents200ResponseMeta.md) |  | [optional] 

## Methods

### NewHardwareReservationList

`func NewHardwareReservationList() *HardwareReservationList`

NewHardwareReservationList instantiates a new HardwareReservationList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHardwareReservationListWithDefaults

`func NewHardwareReservationListWithDefaults() *HardwareReservationList`

NewHardwareReservationListWithDefaults instantiates a new HardwareReservationList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHardwareReservations

`func (o *HardwareReservationList) GetHardwareReservations() []MoveHardwareReservation201Response`

GetHardwareReservations returns the HardwareReservations field if non-nil, zero value otherwise.

### GetHardwareReservationsOk

`func (o *HardwareReservationList) GetHardwareReservationsOk() (*[]MoveHardwareReservation201Response, bool)`

GetHardwareReservationsOk returns a tuple with the HardwareReservations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareReservations

`func (o *HardwareReservationList) SetHardwareReservations(v []MoveHardwareReservation201Response)`

SetHardwareReservations sets HardwareReservations field to given value.

### HasHardwareReservations

`func (o *HardwareReservationList) HasHardwareReservations() bool`

HasHardwareReservations returns a boolean if a field has been set.

### GetMeta

`func (o *HardwareReservationList) GetMeta() FindDeviceEvents200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *HardwareReservationList) GetMetaOk() (*FindDeviceEvents200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *HardwareReservationList) SetMeta(v FindDeviceEvents200ResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *HardwareReservationList) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


