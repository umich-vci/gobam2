# ZoneTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | Pointer to **string** | The name of the resource. | [optional] 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**Configuration** | Pointer to [**Configuration**](Configuration.md) |  | [optional] [readonly] 

## Methods

### NewZoneTemplate

`func NewZoneTemplate() *ZoneTemplate`

NewZoneTemplate instantiates a new ZoneTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneTemplateWithDefaults

`func NewZoneTemplateWithDefaults() *ZoneTemplate`

NewZoneTemplateWithDefaults instantiates a new ZoneTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ZoneTemplate) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ZoneTemplate) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ZoneTemplate) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ZoneTemplate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ZoneTemplate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ZoneTemplate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ZoneTemplate) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ZoneTemplate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *ZoneTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ZoneTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ZoneTemplate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ZoneTemplate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *ZoneTemplate) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *ZoneTemplate) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *ZoneTemplate) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *ZoneTemplate) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetConfiguration

`func (o *ZoneTemplate) GetConfiguration() Configuration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *ZoneTemplate) GetConfigurationOk() (*Configuration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *ZoneTemplate) SetConfiguration(v Configuration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *ZoneTemplate) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


