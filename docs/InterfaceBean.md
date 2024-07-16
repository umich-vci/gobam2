# InterfaceBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InterfaceType** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] [readonly] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Addresses** | Pointer to [**[]IpAddressAndPrefixBean**](IpAddressAndPrefixBean.md) |  | [optional] 
**VlanId** | Pointer to **int32** |  | [optional] 
**RedundancyScenario** | Pointer to **string** |  | [optional] 

## Methods

### NewInterfaceBean

`func NewInterfaceBean() *InterfaceBean`

NewInterfaceBean instantiates a new InterfaceBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInterfaceBeanWithDefaults

`func NewInterfaceBeanWithDefaults() *InterfaceBean`

NewInterfaceBeanWithDefaults instantiates a new InterfaceBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterfaceType

`func (o *InterfaceBean) GetInterfaceType() string`

GetInterfaceType returns the InterfaceType field if non-nil, zero value otherwise.

### GetInterfaceTypeOk

`func (o *InterfaceBean) GetInterfaceTypeOk() (*string, bool)`

GetInterfaceTypeOk returns a tuple with the InterfaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceType

`func (o *InterfaceBean) SetInterfaceType(v string)`

SetInterfaceType sets InterfaceType field to given value.

### HasInterfaceType

`func (o *InterfaceBean) HasInterfaceType() bool`

HasInterfaceType returns a boolean if a field has been set.

### GetName

`func (o *InterfaceBean) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InterfaceBean) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InterfaceBean) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InterfaceBean) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnabled

`func (o *InterfaceBean) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InterfaceBean) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InterfaceBean) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InterfaceBean) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDescription

`func (o *InterfaceBean) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InterfaceBean) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InterfaceBean) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InterfaceBean) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAddresses

`func (o *InterfaceBean) GetAddresses() []IpAddressAndPrefixBean`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *InterfaceBean) GetAddressesOk() (*[]IpAddressAndPrefixBean, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *InterfaceBean) SetAddresses(v []IpAddressAndPrefixBean)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *InterfaceBean) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetVlanId

`func (o *InterfaceBean) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *InterfaceBean) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *InterfaceBean) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *InterfaceBean) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### GetRedundancyScenario

`func (o *InterfaceBean) GetRedundancyScenario() string`

GetRedundancyScenario returns the RedundancyScenario field if non-nil, zero value otherwise.

### GetRedundancyScenarioOk

`func (o *InterfaceBean) GetRedundancyScenarioOk() (*string, bool)`

GetRedundancyScenarioOk returns a tuple with the RedundancyScenario field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancyScenario

`func (o *InterfaceBean) SetRedundancyScenario(v string)`

SetRedundancyScenario sets RedundancyScenario field to given value.

### HasRedundancyScenario

`func (o *InterfaceBean) HasRedundancyScenario() bool`

HasRedundancyScenario returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


