# BgpNeighborBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutonomousSystemNumber** | Pointer to **int64** |  | [optional] 
**Address** | Pointer to **string** |  | [optional] 
**AuthenticationPassword** | Pointer to **string** |  | [optional] 
**HopLimit** | Pointer to **int32** |  | [optional] 
**AnnounceNextHopSelfEnabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewBgpNeighborBean

`func NewBgpNeighborBean() *BgpNeighborBean`

NewBgpNeighborBean instantiates a new BgpNeighborBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBgpNeighborBeanWithDefaults

`func NewBgpNeighborBeanWithDefaults() *BgpNeighborBean`

NewBgpNeighborBeanWithDefaults instantiates a new BgpNeighborBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutonomousSystemNumber

`func (o *BgpNeighborBean) GetAutonomousSystemNumber() int64`

GetAutonomousSystemNumber returns the AutonomousSystemNumber field if non-nil, zero value otherwise.

### GetAutonomousSystemNumberOk

`func (o *BgpNeighborBean) GetAutonomousSystemNumberOk() (*int64, bool)`

GetAutonomousSystemNumberOk returns a tuple with the AutonomousSystemNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutonomousSystemNumber

`func (o *BgpNeighborBean) SetAutonomousSystemNumber(v int64)`

SetAutonomousSystemNumber sets AutonomousSystemNumber field to given value.

### HasAutonomousSystemNumber

`func (o *BgpNeighborBean) HasAutonomousSystemNumber() bool`

HasAutonomousSystemNumber returns a boolean if a field has been set.

### GetAddress

`func (o *BgpNeighborBean) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *BgpNeighborBean) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *BgpNeighborBean) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *BgpNeighborBean) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAuthenticationPassword

`func (o *BgpNeighborBean) GetAuthenticationPassword() string`

GetAuthenticationPassword returns the AuthenticationPassword field if non-nil, zero value otherwise.

### GetAuthenticationPasswordOk

`func (o *BgpNeighborBean) GetAuthenticationPasswordOk() (*string, bool)`

GetAuthenticationPasswordOk returns a tuple with the AuthenticationPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationPassword

`func (o *BgpNeighborBean) SetAuthenticationPassword(v string)`

SetAuthenticationPassword sets AuthenticationPassword field to given value.

### HasAuthenticationPassword

`func (o *BgpNeighborBean) HasAuthenticationPassword() bool`

HasAuthenticationPassword returns a boolean if a field has been set.

### GetHopLimit

`func (o *BgpNeighborBean) GetHopLimit() int32`

GetHopLimit returns the HopLimit field if non-nil, zero value otherwise.

### GetHopLimitOk

`func (o *BgpNeighborBean) GetHopLimitOk() (*int32, bool)`

GetHopLimitOk returns a tuple with the HopLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHopLimit

`func (o *BgpNeighborBean) SetHopLimit(v int32)`

SetHopLimit sets HopLimit field to given value.

### HasHopLimit

`func (o *BgpNeighborBean) HasHopLimit() bool`

HasHopLimit returns a boolean if a field has been set.

### GetAnnounceNextHopSelfEnabled

`func (o *BgpNeighborBean) GetAnnounceNextHopSelfEnabled() bool`

GetAnnounceNextHopSelfEnabled returns the AnnounceNextHopSelfEnabled field if non-nil, zero value otherwise.

### GetAnnounceNextHopSelfEnabledOk

`func (o *BgpNeighborBean) GetAnnounceNextHopSelfEnabledOk() (*bool, bool)`

GetAnnounceNextHopSelfEnabledOk returns a tuple with the AnnounceNextHopSelfEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnounceNextHopSelfEnabled

`func (o *BgpNeighborBean) SetAnnounceNextHopSelfEnabled(v bool)`

SetAnnounceNextHopSelfEnabled sets AnnounceNextHopSelfEnabled field to given value.

### HasAnnounceNextHopSelfEnabled

`func (o *BgpNeighborBean) HasAnnounceNextHopSelfEnabled() bool`

HasAnnounceNextHopSelfEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


