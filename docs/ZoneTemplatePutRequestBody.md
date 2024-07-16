# ZoneTemplatePutRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | **string** | The resource type. | 
**Name** | **string** | The name of the resource. | 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**Configuration** | Pointer to [**Configuration**](Configuration.md) |  | [optional] [readonly] 

## Methods

### NewZoneTemplatePutRequestBody

`func NewZoneTemplatePutRequestBody(type_ string, name string, ) *ZoneTemplatePutRequestBody`

NewZoneTemplatePutRequestBody instantiates a new ZoneTemplatePutRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneTemplatePutRequestBodyWithDefaults

`func NewZoneTemplatePutRequestBodyWithDefaults() *ZoneTemplatePutRequestBody`

NewZoneTemplatePutRequestBodyWithDefaults instantiates a new ZoneTemplatePutRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ZoneTemplatePutRequestBody) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ZoneTemplatePutRequestBody) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ZoneTemplatePutRequestBody) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ZoneTemplatePutRequestBody) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ZoneTemplatePutRequestBody) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ZoneTemplatePutRequestBody) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ZoneTemplatePutRequestBody) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *ZoneTemplatePutRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ZoneTemplatePutRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ZoneTemplatePutRequestBody) SetName(v string)`

SetName sets Name field to given value.


### GetUserDefinedFields

`func (o *ZoneTemplatePutRequestBody) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *ZoneTemplatePutRequestBody) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *ZoneTemplatePutRequestBody) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *ZoneTemplatePutRequestBody) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetConfiguration

`func (o *ZoneTemplatePutRequestBody) GetConfiguration() Configuration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *ZoneTemplatePutRequestBody) GetConfigurationOk() (*Configuration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *ZoneTemplatePutRequestBody) SetConfiguration(v Configuration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *ZoneTemplatePutRequestBody) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


