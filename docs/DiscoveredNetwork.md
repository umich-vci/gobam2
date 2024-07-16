# DiscoveredNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The resource type. | [optional] [readonly] 
**Device** | Pointer to [**DiscoveredDevice**](DiscoveredDevice.md) |  | [optional] [readonly] 
**Range** | Pointer to **string** |  | [optional] [readonly] 
**Interface** | Pointer to [**DiscoveredInterface**](DiscoveredInterface.md) |  | [optional] [readonly] 

## Methods

### NewDiscoveredNetwork

`func NewDiscoveredNetwork() *DiscoveredNetwork`

NewDiscoveredNetwork instantiates a new DiscoveredNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscoveredNetworkWithDefaults

`func NewDiscoveredNetworkWithDefaults() *DiscoveredNetwork`

NewDiscoveredNetworkWithDefaults instantiates a new DiscoveredNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DiscoveredNetwork) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DiscoveredNetwork) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DiscoveredNetwork) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DiscoveredNetwork) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDevice

`func (o *DiscoveredNetwork) GetDevice() DiscoveredDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *DiscoveredNetwork) GetDeviceOk() (*DiscoveredDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *DiscoveredNetwork) SetDevice(v DiscoveredDevice)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *DiscoveredNetwork) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetRange

`func (o *DiscoveredNetwork) GetRange() string`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *DiscoveredNetwork) GetRangeOk() (*string, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *DiscoveredNetwork) SetRange(v string)`

SetRange sets Range field to given value.

### HasRange

`func (o *DiscoveredNetwork) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetInterface

`func (o *DiscoveredNetwork) GetInterface() DiscoveredInterface`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *DiscoveredNetwork) GetInterfaceOk() (*DiscoveredInterface, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *DiscoveredNetwork) SetInterface(v DiscoveredInterface)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *DiscoveredNetwork) HasInterface() bool`

HasInterface returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


