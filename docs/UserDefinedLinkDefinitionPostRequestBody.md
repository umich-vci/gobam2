# UserDefinedLinkDefinitionPostRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | **string** | The name of the user-defined link definition. | 
**DisplayName** | **string** | The display name of the user-defined link definition. | 
**Description** | Pointer to **string** | The description of the user-defined link definition. | [optional] 
**SourceTypes** | **[]string** | The source entity types of the user-defined link definition. | 
**DestinationTypes** | **[]string** | The destination entity types of the user-defined link definition. | 

## Methods

### NewUserDefinedLinkDefinitionPostRequestBody

`func NewUserDefinedLinkDefinitionPostRequestBody(name string, displayName string, sourceTypes []string, destinationTypes []string, ) *UserDefinedLinkDefinitionPostRequestBody`

NewUserDefinedLinkDefinitionPostRequestBody instantiates a new UserDefinedLinkDefinitionPostRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserDefinedLinkDefinitionPostRequestBodyWithDefaults

`func NewUserDefinedLinkDefinitionPostRequestBodyWithDefaults() *UserDefinedLinkDefinitionPostRequestBody`

NewUserDefinedLinkDefinitionPostRequestBodyWithDefaults instantiates a new UserDefinedLinkDefinitionPostRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserDefinedLinkDefinitionPostRequestBody) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserDefinedLinkDefinitionPostRequestBody) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserDefinedLinkDefinitionPostRequestBody) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UserDefinedLinkDefinitionPostRequestBody) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *UserDefinedLinkDefinitionPostRequestBody) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserDefinedLinkDefinitionPostRequestBody) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserDefinedLinkDefinitionPostRequestBody) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UserDefinedLinkDefinitionPostRequestBody) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *UserDefinedLinkDefinitionPostRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserDefinedLinkDefinitionPostRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserDefinedLinkDefinitionPostRequestBody) SetName(v string)`

SetName sets Name field to given value.


### GetDisplayName

`func (o *UserDefinedLinkDefinitionPostRequestBody) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UserDefinedLinkDefinitionPostRequestBody) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UserDefinedLinkDefinitionPostRequestBody) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetDescription

`func (o *UserDefinedLinkDefinitionPostRequestBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UserDefinedLinkDefinitionPostRequestBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UserDefinedLinkDefinitionPostRequestBody) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UserDefinedLinkDefinitionPostRequestBody) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSourceTypes

`func (o *UserDefinedLinkDefinitionPostRequestBody) GetSourceTypes() []string`

GetSourceTypes returns the SourceTypes field if non-nil, zero value otherwise.

### GetSourceTypesOk

`func (o *UserDefinedLinkDefinitionPostRequestBody) GetSourceTypesOk() (*[]string, bool)`

GetSourceTypesOk returns a tuple with the SourceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTypes

`func (o *UserDefinedLinkDefinitionPostRequestBody) SetSourceTypes(v []string)`

SetSourceTypes sets SourceTypes field to given value.


### GetDestinationTypes

`func (o *UserDefinedLinkDefinitionPostRequestBody) GetDestinationTypes() []string`

GetDestinationTypes returns the DestinationTypes field if non-nil, zero value otherwise.

### GetDestinationTypesOk

`func (o *UserDefinedLinkDefinitionPostRequestBody) GetDestinationTypesOk() (*[]string, bool)`

GetDestinationTypesOk returns a tuple with the DestinationTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationTypes

`func (o *UserDefinedLinkDefinitionPostRequestBody) SetDestinationTypes(v []string)`

SetDestinationTypes sets DestinationTypes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


