# TFTPGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | Pointer to **string** | The name of the TFTP group. | [optional] 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**Configuration** | Pointer to [**InlinedConfiguration**](InlinedConfiguration.md) |  | [optional] [readonly] 

## Methods

### NewTFTPGroup

`func NewTFTPGroup() *TFTPGroup`

NewTFTPGroup instantiates a new TFTPGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTFTPGroupWithDefaults

`func NewTFTPGroupWithDefaults() *TFTPGroup`

NewTFTPGroupWithDefaults instantiates a new TFTPGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TFTPGroup) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TFTPGroup) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TFTPGroup) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *TFTPGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *TFTPGroup) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TFTPGroup) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TFTPGroup) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TFTPGroup) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *TFTPGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TFTPGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TFTPGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TFTPGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *TFTPGroup) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *TFTPGroup) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *TFTPGroup) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *TFTPGroup) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetConfiguration

`func (o *TFTPGroup) GetConfiguration() InlinedConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *TFTPGroup) GetConfigurationOk() (*InlinedConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *TFTPGroup) SetConfiguration(v InlinedConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *TFTPGroup) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

