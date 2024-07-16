# DiscoveredHost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The resource type. | [optional] [readonly] 
**Device** | Pointer to [**DiscoveredDevice**](DiscoveredDevice.md) |  | [optional] [readonly] 
**Address** | Pointer to **string** |  | [optional] [readonly] 
**MacAddress** | Pointer to **string** |  | [optional] [readonly] 
**Interface** | Pointer to [**DiscoveredInterface**](DiscoveredInterface.md) |  | [optional] [readonly] 

## Methods

### NewDiscoveredHost

`func NewDiscoveredHost() *DiscoveredHost`

NewDiscoveredHost instantiates a new DiscoveredHost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscoveredHostWithDefaults

`func NewDiscoveredHostWithDefaults() *DiscoveredHost`

NewDiscoveredHostWithDefaults instantiates a new DiscoveredHost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DiscoveredHost) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DiscoveredHost) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DiscoveredHost) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DiscoveredHost) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDevice

`func (o *DiscoveredHost) GetDevice() DiscoveredDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *DiscoveredHost) GetDeviceOk() (*DiscoveredDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *DiscoveredHost) SetDevice(v DiscoveredDevice)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *DiscoveredHost) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetAddress

`func (o *DiscoveredHost) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *DiscoveredHost) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *DiscoveredHost) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *DiscoveredHost) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetMacAddress

`func (o *DiscoveredHost) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *DiscoveredHost) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *DiscoveredHost) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *DiscoveredHost) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetInterface

`func (o *DiscoveredHost) GetInterface() DiscoveredInterface`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *DiscoveredHost) GetInterfaceOk() (*DiscoveredInterface, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *DiscoveredHost) SetInterface(v DiscoveredInterface)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *DiscoveredHost) HasInterface() bool`

HasInterface returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


