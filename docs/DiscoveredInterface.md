# DiscoveredInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The resource type. | [optional] [readonly] 
**Device** | Pointer to [**DiscoveredDevice**](DiscoveredDevice.md) |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] [readonly] 
**Description** | Pointer to **string** |  | [optional] [readonly] 
**InterfaceIndex** | Pointer to **int32** |  | [optional] [readonly] 
**MacAddress** | Pointer to **string** |  | [optional] [readonly] 
**Speed** | Pointer to **int32** |  | [optional] [readonly] 
**Connector** | Pointer to **string** |  | [optional] [readonly] 
**Alias** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewDiscoveredInterface

`func NewDiscoveredInterface() *DiscoveredInterface`

NewDiscoveredInterface instantiates a new DiscoveredInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscoveredInterfaceWithDefaults

`func NewDiscoveredInterfaceWithDefaults() *DiscoveredInterface`

NewDiscoveredInterfaceWithDefaults instantiates a new DiscoveredInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DiscoveredInterface) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DiscoveredInterface) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DiscoveredInterface) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DiscoveredInterface) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDevice

`func (o *DiscoveredInterface) GetDevice() DiscoveredDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *DiscoveredInterface) GetDeviceOk() (*DiscoveredDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *DiscoveredInterface) SetDevice(v DiscoveredDevice)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *DiscoveredInterface) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetName

`func (o *DiscoveredInterface) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DiscoveredInterface) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DiscoveredInterface) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DiscoveredInterface) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *DiscoveredInterface) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DiscoveredInterface) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DiscoveredInterface) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DiscoveredInterface) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInterfaceIndex

`func (o *DiscoveredInterface) GetInterfaceIndex() int32`

GetInterfaceIndex returns the InterfaceIndex field if non-nil, zero value otherwise.

### GetInterfaceIndexOk

`func (o *DiscoveredInterface) GetInterfaceIndexOk() (*int32, bool)`

GetInterfaceIndexOk returns a tuple with the InterfaceIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceIndex

`func (o *DiscoveredInterface) SetInterfaceIndex(v int32)`

SetInterfaceIndex sets InterfaceIndex field to given value.

### HasInterfaceIndex

`func (o *DiscoveredInterface) HasInterfaceIndex() bool`

HasInterfaceIndex returns a boolean if a field has been set.

### GetMacAddress

`func (o *DiscoveredInterface) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *DiscoveredInterface) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *DiscoveredInterface) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *DiscoveredInterface) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetSpeed

`func (o *DiscoveredInterface) GetSpeed() int32`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *DiscoveredInterface) GetSpeedOk() (*int32, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *DiscoveredInterface) SetSpeed(v int32)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *DiscoveredInterface) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetConnector

`func (o *DiscoveredInterface) GetConnector() string`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *DiscoveredInterface) GetConnectorOk() (*string, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *DiscoveredInterface) SetConnector(v string)`

SetConnector sets Connector field to given value.

### HasConnector

`func (o *DiscoveredInterface) HasConnector() bool`

HasConnector returns a boolean if a field has been set.

### GetAlias

`func (o *DiscoveredInterface) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *DiscoveredInterface) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *DiscoveredInterface) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *DiscoveredInterface) HasAlias() bool`

HasAlias returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


