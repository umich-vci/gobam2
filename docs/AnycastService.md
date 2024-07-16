# AnycastService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The resource type. | [optional] 
**Enabled** | Pointer to **bool** | Indicates whether the Anycast service is enabled. | [optional] 
**ManualOverrideEnabled** | Pointer to **bool** | Indicates whether the Anycast service configuration has been manually overridden. | [optional] [readonly] 
**Addresses** | Pointer to **[]string** | The Anycast IP addresses. | [optional] 
**Routing** | Pointer to [**RoutingProtocolBean**](RoutingProtocolBean.md) |  | [optional] 

## Methods

### NewAnycastService

`func NewAnycastService() *AnycastService`

NewAnycastService instantiates a new AnycastService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnycastServiceWithDefaults

`func NewAnycastServiceWithDefaults() *AnycastService`

NewAnycastServiceWithDefaults instantiates a new AnycastService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AnycastService) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AnycastService) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AnycastService) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AnycastService) HasType() bool`

HasType returns a boolean if a field has been set.

### GetEnabled

`func (o *AnycastService) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AnycastService) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AnycastService) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AnycastService) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetManualOverrideEnabled

`func (o *AnycastService) GetManualOverrideEnabled() bool`

GetManualOverrideEnabled returns the ManualOverrideEnabled field if non-nil, zero value otherwise.

### GetManualOverrideEnabledOk

`func (o *AnycastService) GetManualOverrideEnabledOk() (*bool, bool)`

GetManualOverrideEnabledOk returns a tuple with the ManualOverrideEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualOverrideEnabled

`func (o *AnycastService) SetManualOverrideEnabled(v bool)`

SetManualOverrideEnabled sets ManualOverrideEnabled field to given value.

### HasManualOverrideEnabled

`func (o *AnycastService) HasManualOverrideEnabled() bool`

HasManualOverrideEnabled returns a boolean if a field has been set.

### GetAddresses

`func (o *AnycastService) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *AnycastService) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *AnycastService) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *AnycastService) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetRouting

`func (o *AnycastService) GetRouting() RoutingProtocolBean`

GetRouting returns the Routing field if non-nil, zero value otherwise.

### GetRoutingOk

`func (o *AnycastService) GetRoutingOk() (*RoutingProtocolBean, bool)`

GetRoutingOk returns a tuple with the Routing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouting

`func (o *AnycastService) SetRouting(v RoutingProtocolBean)`

SetRouting sets Routing field to given value.

### HasRouting

`func (o *AnycastService) HasRouting() bool`

HasRouting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


