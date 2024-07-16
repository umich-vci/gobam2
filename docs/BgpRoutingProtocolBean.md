# BgpRoutingProtocolBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutonomousSystemNumber** | Pointer to **int64** |  | [optional] 
**RouterId** | Pointer to **string** |  | [optional] 
**CommandLineInterfaceEnabled** | Pointer to **bool** |  | [optional] 
**CommandLineInterfacePassword** | Pointer to **string** |  | [optional] 
**KeepAliveInterval** | Pointer to **string** |  | [optional] 
**HoldTimeInterval** | Pointer to **string** |  | [optional] 
**Neighbors** | Pointer to [**[]BgpNeighborBean**](BgpNeighborBean.md) |  | [optional] 
**PrefixFilters** | Pointer to [**[]BgpPrefixFilterBean**](BgpPrefixFilterBean.md) |  | [optional] 

## Methods

### NewBgpRoutingProtocolBean

`func NewBgpRoutingProtocolBean() *BgpRoutingProtocolBean`

NewBgpRoutingProtocolBean instantiates a new BgpRoutingProtocolBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBgpRoutingProtocolBeanWithDefaults

`func NewBgpRoutingProtocolBeanWithDefaults() *BgpRoutingProtocolBean`

NewBgpRoutingProtocolBeanWithDefaults instantiates a new BgpRoutingProtocolBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutonomousSystemNumber

`func (o *BgpRoutingProtocolBean) GetAutonomousSystemNumber() int64`

GetAutonomousSystemNumber returns the AutonomousSystemNumber field if non-nil, zero value otherwise.

### GetAutonomousSystemNumberOk

`func (o *BgpRoutingProtocolBean) GetAutonomousSystemNumberOk() (*int64, bool)`

GetAutonomousSystemNumberOk returns a tuple with the AutonomousSystemNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutonomousSystemNumber

`func (o *BgpRoutingProtocolBean) SetAutonomousSystemNumber(v int64)`

SetAutonomousSystemNumber sets AutonomousSystemNumber field to given value.

### HasAutonomousSystemNumber

`func (o *BgpRoutingProtocolBean) HasAutonomousSystemNumber() bool`

HasAutonomousSystemNumber returns a boolean if a field has been set.

### GetRouterId

`func (o *BgpRoutingProtocolBean) GetRouterId() string`

GetRouterId returns the RouterId field if non-nil, zero value otherwise.

### GetRouterIdOk

`func (o *BgpRoutingProtocolBean) GetRouterIdOk() (*string, bool)`

GetRouterIdOk returns a tuple with the RouterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouterId

`func (o *BgpRoutingProtocolBean) SetRouterId(v string)`

SetRouterId sets RouterId field to given value.

### HasRouterId

`func (o *BgpRoutingProtocolBean) HasRouterId() bool`

HasRouterId returns a boolean if a field has been set.

### GetCommandLineInterfaceEnabled

`func (o *BgpRoutingProtocolBean) GetCommandLineInterfaceEnabled() bool`

GetCommandLineInterfaceEnabled returns the CommandLineInterfaceEnabled field if non-nil, zero value otherwise.

### GetCommandLineInterfaceEnabledOk

`func (o *BgpRoutingProtocolBean) GetCommandLineInterfaceEnabledOk() (*bool, bool)`

GetCommandLineInterfaceEnabledOk returns a tuple with the CommandLineInterfaceEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandLineInterfaceEnabled

`func (o *BgpRoutingProtocolBean) SetCommandLineInterfaceEnabled(v bool)`

SetCommandLineInterfaceEnabled sets CommandLineInterfaceEnabled field to given value.

### HasCommandLineInterfaceEnabled

`func (o *BgpRoutingProtocolBean) HasCommandLineInterfaceEnabled() bool`

HasCommandLineInterfaceEnabled returns a boolean if a field has been set.

### GetCommandLineInterfacePassword

`func (o *BgpRoutingProtocolBean) GetCommandLineInterfacePassword() string`

GetCommandLineInterfacePassword returns the CommandLineInterfacePassword field if non-nil, zero value otherwise.

### GetCommandLineInterfacePasswordOk

`func (o *BgpRoutingProtocolBean) GetCommandLineInterfacePasswordOk() (*string, bool)`

GetCommandLineInterfacePasswordOk returns a tuple with the CommandLineInterfacePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandLineInterfacePassword

`func (o *BgpRoutingProtocolBean) SetCommandLineInterfacePassword(v string)`

SetCommandLineInterfacePassword sets CommandLineInterfacePassword field to given value.

### HasCommandLineInterfacePassword

`func (o *BgpRoutingProtocolBean) HasCommandLineInterfacePassword() bool`

HasCommandLineInterfacePassword returns a boolean if a field has been set.

### GetKeepAliveInterval

`func (o *BgpRoutingProtocolBean) GetKeepAliveInterval() string`

GetKeepAliveInterval returns the KeepAliveInterval field if non-nil, zero value otherwise.

### GetKeepAliveIntervalOk

`func (o *BgpRoutingProtocolBean) GetKeepAliveIntervalOk() (*string, bool)`

GetKeepAliveIntervalOk returns a tuple with the KeepAliveInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepAliveInterval

`func (o *BgpRoutingProtocolBean) SetKeepAliveInterval(v string)`

SetKeepAliveInterval sets KeepAliveInterval field to given value.

### HasKeepAliveInterval

`func (o *BgpRoutingProtocolBean) HasKeepAliveInterval() bool`

HasKeepAliveInterval returns a boolean if a field has been set.

### GetHoldTimeInterval

`func (o *BgpRoutingProtocolBean) GetHoldTimeInterval() string`

GetHoldTimeInterval returns the HoldTimeInterval field if non-nil, zero value otherwise.

### GetHoldTimeIntervalOk

`func (o *BgpRoutingProtocolBean) GetHoldTimeIntervalOk() (*string, bool)`

GetHoldTimeIntervalOk returns a tuple with the HoldTimeInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldTimeInterval

`func (o *BgpRoutingProtocolBean) SetHoldTimeInterval(v string)`

SetHoldTimeInterval sets HoldTimeInterval field to given value.

### HasHoldTimeInterval

`func (o *BgpRoutingProtocolBean) HasHoldTimeInterval() bool`

HasHoldTimeInterval returns a boolean if a field has been set.

### GetNeighbors

`func (o *BgpRoutingProtocolBean) GetNeighbors() []BgpNeighborBean`

GetNeighbors returns the Neighbors field if non-nil, zero value otherwise.

### GetNeighborsOk

`func (o *BgpRoutingProtocolBean) GetNeighborsOk() (*[]BgpNeighborBean, bool)`

GetNeighborsOk returns a tuple with the Neighbors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeighbors

`func (o *BgpRoutingProtocolBean) SetNeighbors(v []BgpNeighborBean)`

SetNeighbors sets Neighbors field to given value.

### HasNeighbors

`func (o *BgpRoutingProtocolBean) HasNeighbors() bool`

HasNeighbors returns a boolean if a field has been set.

### GetPrefixFilters

`func (o *BgpRoutingProtocolBean) GetPrefixFilters() []BgpPrefixFilterBean`

GetPrefixFilters returns the PrefixFilters field if non-nil, zero value otherwise.

### GetPrefixFiltersOk

`func (o *BgpRoutingProtocolBean) GetPrefixFiltersOk() (*[]BgpPrefixFilterBean, bool)`

GetPrefixFiltersOk returns a tuple with the PrefixFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixFilters

`func (o *BgpRoutingProtocolBean) SetPrefixFilters(v []BgpPrefixFilterBean)`

SetPrefixFilters sets PrefixFilters field to given value.

### HasPrefixFilters

`func (o *BgpRoutingProtocolBean) HasPrefixFilters() bool`

HasPrefixFilters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


