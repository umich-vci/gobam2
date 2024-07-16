# MACAddressPostRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | Pointer to **string** | The name of the resource. | [optional] 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**Configuration** | Pointer to [**InlinedConfiguration**](InlinedConfiguration.md) |  | [optional] [readonly] 
**Address** | **string** | The MAC address. | 
**MacPool** | Pointer to [**InlinedMacPool**](InlinedMacPool.md) |  | [optional] 

## Methods

### NewMACAddressPostRequestBody

`func NewMACAddressPostRequestBody(address string, ) *MACAddressPostRequestBody`

NewMACAddressPostRequestBody instantiates a new MACAddressPostRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMACAddressPostRequestBodyWithDefaults

`func NewMACAddressPostRequestBodyWithDefaults() *MACAddressPostRequestBody`

NewMACAddressPostRequestBodyWithDefaults instantiates a new MACAddressPostRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MACAddressPostRequestBody) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MACAddressPostRequestBody) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MACAddressPostRequestBody) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *MACAddressPostRequestBody) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *MACAddressPostRequestBody) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MACAddressPostRequestBody) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MACAddressPostRequestBody) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MACAddressPostRequestBody) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *MACAddressPostRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MACAddressPostRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MACAddressPostRequestBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MACAddressPostRequestBody) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *MACAddressPostRequestBody) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *MACAddressPostRequestBody) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *MACAddressPostRequestBody) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *MACAddressPostRequestBody) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetConfiguration

`func (o *MACAddressPostRequestBody) GetConfiguration() InlinedConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *MACAddressPostRequestBody) GetConfigurationOk() (*InlinedConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *MACAddressPostRequestBody) SetConfiguration(v InlinedConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *MACAddressPostRequestBody) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetAddress

`func (o *MACAddressPostRequestBody) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *MACAddressPostRequestBody) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *MACAddressPostRequestBody) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetMacPool

`func (o *MACAddressPostRequestBody) GetMacPool() InlinedMacPool`

GetMacPool returns the MacPool field if non-nil, zero value otherwise.

### GetMacPoolOk

`func (o *MACAddressPostRequestBody) GetMacPoolOk() (*InlinedMacPool, bool)`

GetMacPoolOk returns a tuple with the MacPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacPool

`func (o *MACAddressPostRequestBody) SetMacPool(v InlinedMacPool)`

SetMacPool sets MacPool field to given value.

### HasMacPool

`func (o *MACAddressPostRequestBody) HasMacPool() bool`

HasMacPool returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


