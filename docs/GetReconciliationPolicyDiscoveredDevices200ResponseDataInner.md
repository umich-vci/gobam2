# GetReconciliationPolicyDiscoveredDevices200ResponseDataInner

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
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 

## Methods

### NewGetReconciliationPolicyDiscoveredDevices200ResponseDataInner

`func NewGetReconciliationPolicyDiscoveredDevices200ResponseDataInner() *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner`

NewGetReconciliationPolicyDiscoveredDevices200ResponseDataInner instantiates a new GetReconciliationPolicyDiscoveredDevices200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetReconciliationPolicyDiscoveredDevices200ResponseDataInnerWithDefaults

`func NewGetReconciliationPolicyDiscoveredDevices200ResponseDataInnerWithDefaults() *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner`

NewGetReconciliationPolicyDiscoveredDevices200ResponseDataInnerWithDefaults instantiates a new GetReconciliationPolicyDiscoveredDevices200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDevice

`func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) GetDevice() DiscoveredDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) GetDeviceOk() (*DiscoveredDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) SetDevice(v DiscoveredDevice)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetAddress

`func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetMacAddress

`func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetInterface

`func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) GetInterface() DiscoveredInterface`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) GetInterfaceOk() (*DiscoveredInterface, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) SetInterface(v DiscoveredInterface)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### GetName

`func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInterfaceIndex

`func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) GetInterfaceIndex() int32`

GetInterfaceIndex returns the InterfaceIndex field if non-nil, zero value otherwise.

### GetInterfaceIndexOk

`func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) GetInterfaceIndexOk() (*int32, bool)`

GetInterfaceIndexOk returns a tuple with the InterfaceIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceIndex

`func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) SetInterfaceIndex(v int32)`

SetInterfaceIndex sets InterfaceIndex field to given value.

### HasInterfaceIndex

`func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) HasInterfaceIndex() bool`

HasInterfaceIndex returns a boolean if a field has been set.

### GetSpeed

`func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) GetSpeed() int32`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) GetSpeedOk() (*int32, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) SetSpeed(v int32)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetConnector

`func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) GetConnector() string`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) GetConnectorOk() (*string, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) SetConnector(v string)`

SetConnector sets Connector field to given value.

### HasConnector

`func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) HasConnector() bool`

HasConnector returns a boolean if a field has been set.

### GetAlias

`func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetPortMode

`func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) GetPortMode() string`

GetPortMode returns the PortMode field if non-nil, zero value otherwise.

### GetPortModeOk

`func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) GetPortModeOk() (*string, bool)`

GetPortModeOk returns a tuple with the PortMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortMode

`func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) SetPortMode(v string)`

SetPortMode sets PortMode field to given value.

### HasPortMode

`func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) HasPortMode() bool`

HasPortMode returns a boolean if a field has been set.

### GetRange

`func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) GetRange() string`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) GetRangeOk() (*string, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) SetRange(v string)`

SetRange sets Range field to given value.

### HasRange

`func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetLocation

`func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetVlanId

`func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### GetLinks

`func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


