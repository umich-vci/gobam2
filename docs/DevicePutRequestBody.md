# DevicePutRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | **string** | The name of the resource. | 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**Configuration** | Pointer to [**InlinedConfiguration**](InlinedConfiguration.md) |  | [optional] [readonly] 
**Addresses** | Pointer to [**[]DeviceAddressesInner**](DeviceAddressesInner.md) |  | [optional] 
**DeviceType** | Pointer to [**InlinedDeviceType**](InlinedDeviceType.md) |  | [optional] 
**DeviceSubtype** | Pointer to [**InlinedDeviceSubtype**](InlinedDeviceSubtype.md) |  | [optional] 

## Methods

### NewDevicePutRequestBody

`func NewDevicePutRequestBody(name string, ) *DevicePutRequestBody`

NewDevicePutRequestBody instantiates a new DevicePutRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevicePutRequestBodyWithDefaults

`func NewDevicePutRequestBodyWithDefaults() *DevicePutRequestBody`

NewDevicePutRequestBodyWithDefaults instantiates a new DevicePutRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DevicePutRequestBody) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DevicePutRequestBody) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DevicePutRequestBody) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DevicePutRequestBody) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *DevicePutRequestBody) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DevicePutRequestBody) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DevicePutRequestBody) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DevicePutRequestBody) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *DevicePutRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DevicePutRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DevicePutRequestBody) SetName(v string)`

SetName sets Name field to given value.


### GetUserDefinedFields

`func (o *DevicePutRequestBody) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *DevicePutRequestBody) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *DevicePutRequestBody) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *DevicePutRequestBody) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetConfiguration

`func (o *DevicePutRequestBody) GetConfiguration() InlinedConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *DevicePutRequestBody) GetConfigurationOk() (*InlinedConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *DevicePutRequestBody) SetConfiguration(v InlinedConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *DevicePutRequestBody) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetAddresses

`func (o *DevicePutRequestBody) GetAddresses() []DeviceAddressesInner`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *DevicePutRequestBody) GetAddressesOk() (*[]DeviceAddressesInner, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *DevicePutRequestBody) SetAddresses(v []DeviceAddressesInner)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *DevicePutRequestBody) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetDeviceType

`func (o *DevicePutRequestBody) GetDeviceType() InlinedDeviceType`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *DevicePutRequestBody) GetDeviceTypeOk() (*InlinedDeviceType, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *DevicePutRequestBody) SetDeviceType(v InlinedDeviceType)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *DevicePutRequestBody) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetDeviceSubtype

`func (o *DevicePutRequestBody) GetDeviceSubtype() InlinedDeviceSubtype`

GetDeviceSubtype returns the DeviceSubtype field if non-nil, zero value otherwise.

### GetDeviceSubtypeOk

`func (o *DevicePutRequestBody) GetDeviceSubtypeOk() (*InlinedDeviceSubtype, bool)`

GetDeviceSubtypeOk returns a tuple with the DeviceSubtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceSubtype

`func (o *DevicePutRequestBody) SetDeviceSubtype(v InlinedDeviceSubtype)`

SetDeviceSubtype sets DeviceSubtype field to given value.

### HasDeviceSubtype

`func (o *DevicePutRequestBody) HasDeviceSubtype() bool`

HasDeviceSubtype returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


