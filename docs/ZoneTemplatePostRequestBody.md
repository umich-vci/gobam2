# ZoneTemplatePostRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | **string** | The resource type. | 
**Name** | **string** | The name of the resource. | 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**Configuration** | Pointer to [**Configuration**](Configuration.md) |  | [optional] [readonly] 

## Methods

### NewZoneTemplatePostRequestBody

`func NewZoneTemplatePostRequestBody(type_ string, name string, ) *ZoneTemplatePostRequestBody`

NewZoneTemplatePostRequestBody instantiates a new ZoneTemplatePostRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneTemplatePostRequestBodyWithDefaults

`func NewZoneTemplatePostRequestBodyWithDefaults() *ZoneTemplatePostRequestBody`

NewZoneTemplatePostRequestBodyWithDefaults instantiates a new ZoneTemplatePostRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ZoneTemplatePostRequestBody) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ZoneTemplatePostRequestBody) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ZoneTemplatePostRequestBody) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ZoneTemplatePostRequestBody) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ZoneTemplatePostRequestBody) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ZoneTemplatePostRequestBody) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ZoneTemplatePostRequestBody) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *ZoneTemplatePostRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ZoneTemplatePostRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ZoneTemplatePostRequestBody) SetName(v string)`

SetName sets Name field to given value.


### GetUserDefinedFields

`func (o *ZoneTemplatePostRequestBody) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *ZoneTemplatePostRequestBody) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *ZoneTemplatePostRequestBody) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *ZoneTemplatePostRequestBody) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetConfiguration

`func (o *ZoneTemplatePostRequestBody) GetConfiguration() Configuration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *ZoneTemplatePostRequestBody) GetConfigurationOk() (*Configuration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *ZoneTemplatePostRequestBody) SetConfiguration(v Configuration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *ZoneTemplatePostRequestBody) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


