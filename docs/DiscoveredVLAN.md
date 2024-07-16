# DiscoveredVLAN

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The resource type. | [optional] [readonly] 
**Device** | Pointer to [**DiscoveredDevice**](DiscoveredDevice.md) |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] [readonly] 
**VlanId** | Pointer to **int32** |  | [optional] [readonly] 
**Interface** | Pointer to [**DiscoveredInterface**](DiscoveredInterface.md) |  | [optional] [readonly] 

## Methods

### NewDiscoveredVLAN

`func NewDiscoveredVLAN() *DiscoveredVLAN`

NewDiscoveredVLAN instantiates a new DiscoveredVLAN object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscoveredVLANWithDefaults

`func NewDiscoveredVLANWithDefaults() *DiscoveredVLAN`

NewDiscoveredVLANWithDefaults instantiates a new DiscoveredVLAN object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DiscoveredVLAN) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DiscoveredVLAN) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DiscoveredVLAN) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DiscoveredVLAN) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDevice

`func (o *DiscoveredVLAN) GetDevice() DiscoveredDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *DiscoveredVLAN) GetDeviceOk() (*DiscoveredDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *DiscoveredVLAN) SetDevice(v DiscoveredDevice)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *DiscoveredVLAN) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetName

`func (o *DiscoveredVLAN) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DiscoveredVLAN) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DiscoveredVLAN) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DiscoveredVLAN) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVlanId

`func (o *DiscoveredVLAN) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *DiscoveredVLAN) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *DiscoveredVLAN) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *DiscoveredVLAN) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### GetInterface

`func (o *DiscoveredVLAN) GetInterface() DiscoveredInterface`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *DiscoveredVLAN) GetInterfaceOk() (*DiscoveredInterface, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *DiscoveredVLAN) SetInterface(v DiscoveredInterface)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *DiscoveredVLAN) HasInterface() bool`

HasInterface returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


