# DiscoveredARPEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The resource type. | [optional] [readonly] 
**Device** | Pointer to [**DiscoveredDevice**](DiscoveredDevice.md) |  | [optional] [readonly] 
**Address** | Pointer to **string** |  | [optional] [readonly] 
**MacAddress** | Pointer to **string** |  | [optional] [readonly] 
**Interface** | Pointer to [**DiscoveredInterface**](DiscoveredInterface.md) |  | [optional] [readonly] 

## Methods

### NewDiscoveredARPEntry

`func NewDiscoveredARPEntry() *DiscoveredARPEntry`

NewDiscoveredARPEntry instantiates a new DiscoveredARPEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscoveredARPEntryWithDefaults

`func NewDiscoveredARPEntryWithDefaults() *DiscoveredARPEntry`

NewDiscoveredARPEntryWithDefaults instantiates a new DiscoveredARPEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DiscoveredARPEntry) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DiscoveredARPEntry) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DiscoveredARPEntry) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DiscoveredARPEntry) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDevice

`func (o *DiscoveredARPEntry) GetDevice() DiscoveredDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *DiscoveredARPEntry) GetDeviceOk() (*DiscoveredDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *DiscoveredARPEntry) SetDevice(v DiscoveredDevice)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *DiscoveredARPEntry) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetAddress

`func (o *DiscoveredARPEntry) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *DiscoveredARPEntry) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *DiscoveredARPEntry) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *DiscoveredARPEntry) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetMacAddress

`func (o *DiscoveredARPEntry) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *DiscoveredARPEntry) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *DiscoveredARPEntry) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *DiscoveredARPEntry) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetInterface

`func (o *DiscoveredARPEntry) GetInterface() DiscoveredInterface`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *DiscoveredARPEntry) GetInterfaceOk() (*DiscoveredInterface, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *DiscoveredARPEntry) SetInterface(v DiscoveredInterface)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *DiscoveredARPEntry) HasInterface() bool`

HasInterface returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


