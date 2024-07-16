# DiscoveredMACAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The resource type. | [optional] [readonly] 
**Device** | Pointer to [**DiscoveredDevice**](DiscoveredDevice.md) |  | [optional] [readonly] 
**MacAddress** | Pointer to **string** |  | [optional] [readonly] 
**PortMode** | Pointer to **string** |  | [optional] [readonly] 
**Interface** | Pointer to [**DiscoveredInterface**](DiscoveredInterface.md) |  | [optional] [readonly] 

## Methods

### NewDiscoveredMACAddress

`func NewDiscoveredMACAddress() *DiscoveredMACAddress`

NewDiscoveredMACAddress instantiates a new DiscoveredMACAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscoveredMACAddressWithDefaults

`func NewDiscoveredMACAddressWithDefaults() *DiscoveredMACAddress`

NewDiscoveredMACAddressWithDefaults instantiates a new DiscoveredMACAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DiscoveredMACAddress) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DiscoveredMACAddress) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DiscoveredMACAddress) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DiscoveredMACAddress) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDevice

`func (o *DiscoveredMACAddress) GetDevice() DiscoveredDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *DiscoveredMACAddress) GetDeviceOk() (*DiscoveredDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *DiscoveredMACAddress) SetDevice(v DiscoveredDevice)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *DiscoveredMACAddress) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetMacAddress

`func (o *DiscoveredMACAddress) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *DiscoveredMACAddress) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *DiscoveredMACAddress) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *DiscoveredMACAddress) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetPortMode

`func (o *DiscoveredMACAddress) GetPortMode() string`

GetPortMode returns the PortMode field if non-nil, zero value otherwise.

### GetPortModeOk

`func (o *DiscoveredMACAddress) GetPortModeOk() (*string, bool)`

GetPortModeOk returns a tuple with the PortMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortMode

`func (o *DiscoveredMACAddress) SetPortMode(v string)`

SetPortMode sets PortMode field to given value.

### HasPortMode

`func (o *DiscoveredMACAddress) HasPortMode() bool`

HasPortMode returns a boolean if a field has been set.

### GetInterface

`func (o *DiscoveredMACAddress) GetInterface() DiscoveredInterface`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *DiscoveredMACAddress) GetInterfaceOk() (*DiscoveredInterface, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *DiscoveredMACAddress) SetInterface(v DiscoveredInterface)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *DiscoveredMACAddress) HasInterface() bool`

HasInterface returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


