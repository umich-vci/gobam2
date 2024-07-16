# UserDefinedLinkDefinitionPutRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | Pointer to **string** | The name of the user-defined link definition. | [optional] 
**DisplayName** | **string** | The display name of the user-defined link definition. | 
**Description** | Pointer to **string** | The description of the user-defined link definition. | [optional] 
**SourceTypes** | **[]string** | The source entity types of the user-defined link definition. | 
**DestinationTypes** | **[]string** | The destination entity types of the user-defined link definition. | 

## Methods

### NewUserDefinedLinkDefinitionPutRequestBody

`func NewUserDefinedLinkDefinitionPutRequestBody(displayName string, sourceTypes []string, destinationTypes []string, ) *UserDefinedLinkDefinitionPutRequestBody`

NewUserDefinedLinkDefinitionPutRequestBody instantiates a new UserDefinedLinkDefinitionPutRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserDefinedLinkDefinitionPutRequestBodyWithDefaults

`func NewUserDefinedLinkDefinitionPutRequestBodyWithDefaults() *UserDefinedLinkDefinitionPutRequestBody`

NewUserDefinedLinkDefinitionPutRequestBodyWithDefaults instantiates a new UserDefinedLinkDefinitionPutRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserDefinedLinkDefinitionPutRequestBody) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserDefinedLinkDefinitionPutRequestBody) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserDefinedLinkDefinitionPutRequestBody) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UserDefinedLinkDefinitionPutRequestBody) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *UserDefinedLinkDefinitionPutRequestBody) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserDefinedLinkDefinitionPutRequestBody) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserDefinedLinkDefinitionPutRequestBody) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UserDefinedLinkDefinitionPutRequestBody) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *UserDefinedLinkDefinitionPutRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserDefinedLinkDefinitionPutRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserDefinedLinkDefinitionPutRequestBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserDefinedLinkDefinitionPutRequestBody) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisplayName

`func (o *UserDefinedLinkDefinitionPutRequestBody) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UserDefinedLinkDefinitionPutRequestBody) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UserDefinedLinkDefinitionPutRequestBody) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetDescription

`func (o *UserDefinedLinkDefinitionPutRequestBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UserDefinedLinkDefinitionPutRequestBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UserDefinedLinkDefinitionPutRequestBody) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UserDefinedLinkDefinitionPutRequestBody) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSourceTypes

`func (o *UserDefinedLinkDefinitionPutRequestBody) GetSourceTypes() []string`

GetSourceTypes returns the SourceTypes field if non-nil, zero value otherwise.

### GetSourceTypesOk

`func (o *UserDefinedLinkDefinitionPutRequestBody) GetSourceTypesOk() (*[]string, bool)`

GetSourceTypesOk returns a tuple with the SourceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTypes

`func (o *UserDefinedLinkDefinitionPutRequestBody) SetSourceTypes(v []string)`

SetSourceTypes sets SourceTypes field to given value.


### GetDestinationTypes

`func (o *UserDefinedLinkDefinitionPutRequestBody) GetDestinationTypes() []string`

GetDestinationTypes returns the DestinationTypes field if non-nil, zero value otherwise.

### GetDestinationTypesOk

`func (o *UserDefinedLinkDefinitionPutRequestBody) GetDestinationTypesOk() (*[]string, bool)`

GetDestinationTypesOk returns a tuple with the DestinationTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationTypes

`func (o *UserDefinedLinkDefinitionPutRequestBody) SetDestinationTypes(v []string)`

SetDestinationTypes sets DestinationTypes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


