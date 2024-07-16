# GetReconciliationPolicyDiscoveredDevices200Response1DataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] [readonly] 
**Device** | Pointer to [**DiscoveredDevice**](DiscoveredDevice.md) |  | [optional] [readonly] 
**Address** | Pointer to **string** |  | [optional] [readonly] 
**MacAddress** | Pointer to **string** |  | [optional] [readonly] 
**Interface** | Pointer to [**DiscoveredInterface**](DiscoveredInterface.md) |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] [readonly] 
**Description** | Pointer to **string** |  | [optional] [readonly] 
**InterfaceIndex** | Pointer to **int32** |  | [optional] [readonly] 
**Speed** | Pointer to **int32** |  | [optional] [readonly] 
**Connector** | Pointer to **string** |  | [optional] [readonly] 
**Alias** | Pointer to **string** |  | [optional] [readonly] 
**PortMode** | Pointer to **string** |  | [optional] [readonly] 
**Range** | Pointer to **string** |  | [optional] [readonly] 
**Location** | Pointer to **string** |  | [optional] [readonly] 
**VlanId** | Pointer to **int32** |  | [optional] [readonly] 

## Methods

### NewGetReconciliationPolicyDiscoveredDevices200Response1DataInner

`func NewGetReconciliationPolicyDiscoveredDevices200Response1DataInner() *GetReconciliationPolicyDiscoveredDevices200Response1DataInner`

NewGetReconciliationPolicyDiscoveredDevices200Response1DataInner instantiates a new GetReconciliationPolicyDiscoveredDevices200Response1DataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetReconciliationPolicyDiscoveredDevices200Response1DataInnerWithDefaults

`func NewGetReconciliationPolicyDiscoveredDevices200Response1DataInnerWithDefaults() *GetReconciliationPolicyDiscoveredDevices200Response1DataInner`

NewGetReconciliationPolicyDiscoveredDevices200Response1DataInnerWithDefaults instantiates a new GetReconciliationPolicyDiscoveredDevices200Response1DataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetReconciliationPolicyDiscoveredDevices200Response1DataInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetReconciliationPolicyDiscoveredDevices200Response1DataInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetReconciliationPolicyDiscoveredDevices200Response1DataInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetReconciliationPolicyDiscoveredDevices200Response1DataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GetReconciliationPolicyDiscoveredDevices200Response1DataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetReconciliationPolicyDiscoveredDevices200Response1DataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetReconciliationPolicyDiscoveredDevices200Response1DataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetReconciliationPolicyDiscoveredDevices200Response1DataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDevice

`func (o *GetReconciliationPolicyDiscoveredDevices200Response1DataInner) GetDevice() DiscoveredDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *GetReconciliationPolicyDiscoveredDevices200Response1DataInner) GetDeviceOk() (*DiscoveredDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *GetReconciliationPolicyDiscoveredDevices200Response1DataInner) SetDevice(v DiscoveredDevice)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *GetReconciliationPolicyDiscoveredDevices200Response1DataInner) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetAddress

`func (o *GetReconciliationPolicyDiscoveredDevices200Response1DataInner) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *GetReconciliationPolicyDiscoveredDevices200Response1DataInner) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *GetReconciliationPolicyDiscoveredDevices200Response1DataInner) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *GetReconciliationPolicyDiscoveredDevices200Response1DataInner) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetMacAddress

`func (o *GetReconciliationPolicyDiscoveredDevices200Response1DataInner) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *GetReconciliationPolicyDiscoveredDevices200Response1DataInner) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *GetReconciliationPolicyDiscoveredDevices200Response1DataInner) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *GetReconciliationPolicyDiscoveredDevices200Response1DataInner) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetInterface

`func (o *GetReconciliationPolicyDiscoveredDevices200Response1DataInner) GetInterface() DiscoveredInterface`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *GetReconciliationPolicyDiscoveredDevices200Response1DataInner) GetInterfaceOk() (*DiscoveredInterface, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *GetReconciliationPolicyDiscoveredDevices200Response1DataInner) SetInterface(v DiscoveredInterface)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *GetReconciliationPolicyDiscoveredDevices200Response1DataInner) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### GetName

`func (o *GetReconciliationPolicyDiscoveredDevices200Response1DataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetReconciliationPolicyDiscoveredDevices200Response1DataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetReconciliationPolicyDiscoveredDevices200Response1DataInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetReconciliationPolicyDiscoveredDevices200Response1DataInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GetReconciliationPolicyDiscoveredDevices200Response1DataInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetReconciliationPolicyDiscoveredDevices200Response1DataInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetReconciliationPolicyDiscoveredDevices200Response1DataInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetReconciliationPolicyDiscoveredDevices200Response1DataInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInterfaceIndex

`func (o *GetReconciliationPolicyDiscoveredDevices200Response1DataInner) GetInterfaceIndex() int32`

GetInterfaceIndex returns the InterfaceIndex field if non-nil, zero value otherwise.

### GetInterfaceIndexOk

`func (o *GetReconciliationPolicyDiscoveredDevices200Response1DataInner) GetInterfaceIndexOk() (*int32, bool)`

GetInterfaceIndexOk returns a tuple with the InterfaceIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceIndex

`func (o *GetReconciliationPolicyDiscoveredDevices200Response1DataInner) SetInterfaceIndex(v int32)`

SetInterfaceIndex sets InterfaceIndex field to given value.

### HasInterfaceIndex

`func (o *GetReconciliationPolicyDiscoveredDevices200Response1DataInner) HasInterfaceIndex() bool`

HasInterfaceIndex returns a boolean if a field has been set.

### GetSpeed

`func (o *GetReconciliationPolicyDiscoveredDevices200Response1DataInner) GetSpeed() int32`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *GetReconciliationPolicyDiscoveredDevices200Response1DataInner) GetSpeedOk() (*int32, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *GetReconciliationPolicyDiscoveredDevices200Response1DataInner) SetSpeed(v int32)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *GetReconciliationPolicyDiscoveredDevices200Response1DataInner) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetConnector

`func (o *GetReconciliationPolicyDiscoveredDevices200Response1DataInner) GetConnector() string`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *GetReconciliationPolicyDiscoveredDevices200Response1DataInner) GetConnectorOk() (*string, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *GetReconciliationPolicyDiscoveredDevices200Response1DataInner) SetConnector(v string)`

SetConnector sets Connector field to given value.

### HasConnector

`func (o *GetReconciliationPolicyDiscoveredDevices200Response1DataInner) HasConnector() bool`

HasConnector returns a boolean if a field has been set.

### GetAlias

`func (o *GetReconciliationPolicyDiscoveredDevices200Response1DataInner) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *GetReconciliationPolicyDiscoveredDevices200Response1DataInner) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *GetReconciliationPolicyDiscoveredDevices200Response1DataInner) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *GetReconciliationPolicyDiscoveredDevices200Response1DataInner) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetPortMode

`func (o *GetReconciliationPolicyDiscoveredDevices200Response1DataInner) GetPortMode() string`

GetPortMode returns the PortMode field if non-nil, zero value otherwise.

### GetPortModeOk

`func (o *GetReconciliationPolicyDiscoveredDevices200Response1DataInner) GetPortModeOk() (*string, bool)`

GetPortModeOk returns a tuple with the PortMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortMode

`func (o *GetReconciliationPolicyDiscoveredDevices200Response1DataInner) SetPortMode(v string)`

SetPortMode sets PortMode field to given value.

### HasPortMode

`func (o *GetReconciliationPolicyDiscoveredDevices200Response1DataInner) HasPortMode() bool`

HasPortMode returns a boolean if a field has been set.

### GetRange

`func (o *GetReconciliationPolicyDiscoveredDevices200Response1DataInner) GetRange() string`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *GetReconciliationPolicyDiscoveredDevices200Response1DataInner) GetRangeOk() (*string, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *GetReconciliationPolicyDiscoveredDevices200Response1DataInner) SetRange(v string)`

SetRange sets Range field to given value.

### HasRange

`func (o *GetReconciliationPolicyDiscoveredDevices200Response1DataInner) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetLocation

`func (o *GetReconciliationPolicyDiscoveredDevices200Response1DataInner) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *GetReconciliationPolicyDiscoveredDevices200Response1DataInner) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *GetReconciliationPolicyDiscoveredDevices200Response1DataInner) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *GetReconciliationPolicyDiscoveredDevices200Response1DataInner) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetVlanId

`func (o *GetReconciliationPolicyDiscoveredDevices200Response1DataInner) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *GetReconciliationPolicyDiscoveredDevices200Response1DataInner) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *GetReconciliationPolicyDiscoveredDevices200Response1DataInner) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *GetReconciliationPolicyDiscoveredDevices200Response1DataInner) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


