# UserDefinedLinkDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | Pointer to **string** | The name of the user-defined link definition. | [optional] 
**DisplayName** | Pointer to **string** | The display name of the user-defined link definition. | [optional] 
**Description** | Pointer to **string** | The description of the user-defined link definition. | [optional] 
**SourceTypes** | Pointer to **[]string** | The source entity types of the user-defined link definition. | [optional] 
**DestinationTypes** | Pointer to **[]string** | The destination entity types of the user-defined link definition. | [optional] 

## Methods

### NewUserDefinedLinkDefinition

`func NewUserDefinedLinkDefinition() *UserDefinedLinkDefinition`

NewUserDefinedLinkDefinition instantiates a new UserDefinedLinkDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserDefinedLinkDefinitionWithDefaults

`func NewUserDefinedLinkDefinitionWithDefaults() *UserDefinedLinkDefinition`

NewUserDefinedLinkDefinitionWithDefaults instantiates a new UserDefinedLinkDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserDefinedLinkDefinition) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserDefinedLinkDefinition) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserDefinedLinkDefinition) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UserDefinedLinkDefinition) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *UserDefinedLinkDefinition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserDefinedLinkDefinition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserDefinedLinkDefinition) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UserDefinedLinkDefinition) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *UserDefinedLinkDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserDefinedLinkDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserDefinedLinkDefinition) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserDefinedLinkDefinition) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisplayName

`func (o *UserDefinedLinkDefinition) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UserDefinedLinkDefinition) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UserDefinedLinkDefinition) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *UserDefinedLinkDefinition) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *UserDefinedLinkDefinition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UserDefinedLinkDefinition) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UserDefinedLinkDefinition) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UserDefinedLinkDefinition) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSourceTypes

`func (o *UserDefinedLinkDefinition) GetSourceTypes() []string`

GetSourceTypes returns the SourceTypes field if non-nil, zero value otherwise.

### GetSourceTypesOk

`func (o *UserDefinedLinkDefinition) GetSourceTypesOk() (*[]string, bool)`

GetSourceTypesOk returns a tuple with the SourceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTypes

`func (o *UserDefinedLinkDefinition) SetSourceTypes(v []string)`

SetSourceTypes sets SourceTypes field to given value.

### HasSourceTypes

`func (o *UserDefinedLinkDefinition) HasSourceTypes() bool`

HasSourceTypes returns a boolean if a field has been set.

### GetDestinationTypes

`func (o *UserDefinedLinkDefinition) GetDestinationTypes() []string`

GetDestinationTypes returns the DestinationTypes field if non-nil, zero value otherwise.

### GetDestinationTypesOk

`func (o *UserDefinedLinkDefinition) GetDestinationTypesOk() (*[]string, bool)`

GetDestinationTypesOk returns a tuple with the DestinationTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationTypes

`func (o *UserDefinedLinkDefinition) SetDestinationTypes(v []string)`

SetDestinationTypes sets DestinationTypes field to given value.

### HasDestinationTypes

`func (o *UserDefinedLinkDefinition) HasDestinationTypes() bool`

HasDestinationTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


