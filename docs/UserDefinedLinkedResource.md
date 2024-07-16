# UserDefinedLinkedResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**LinkDescription** | Pointer to **string** | The description for the user-defined link. | [optional] 
**LinkDefinition** | Pointer to [**UserDefinedLinkedResourceLinkDefinition**](UserDefinedLinkedResourceLinkDefinition.md) |  | [optional] 

## Methods

### NewUserDefinedLinkedResource

`func NewUserDefinedLinkedResource() *UserDefinedLinkedResource`

NewUserDefinedLinkedResource instantiates a new UserDefinedLinkedResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserDefinedLinkedResourceWithDefaults

`func NewUserDefinedLinkedResourceWithDefaults() *UserDefinedLinkedResource`

NewUserDefinedLinkedResourceWithDefaults instantiates a new UserDefinedLinkedResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserDefinedLinkedResource) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserDefinedLinkedResource) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserDefinedLinkedResource) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UserDefinedLinkedResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *UserDefinedLinkedResource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserDefinedLinkedResource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserDefinedLinkedResource) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UserDefinedLinkedResource) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinkDescription

`func (o *UserDefinedLinkedResource) GetLinkDescription() string`

GetLinkDescription returns the LinkDescription field if non-nil, zero value otherwise.

### GetLinkDescriptionOk

`func (o *UserDefinedLinkedResource) GetLinkDescriptionOk() (*string, bool)`

GetLinkDescriptionOk returns a tuple with the LinkDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkDescription

`func (o *UserDefinedLinkedResource) SetLinkDescription(v string)`

SetLinkDescription sets LinkDescription field to given value.

### HasLinkDescription

`func (o *UserDefinedLinkedResource) HasLinkDescription() bool`

HasLinkDescription returns a boolean if a field has been set.

### GetLinkDefinition

`func (o *UserDefinedLinkedResource) GetLinkDefinition() UserDefinedLinkedResourceLinkDefinition`

GetLinkDefinition returns the LinkDefinition field if non-nil, zero value otherwise.

### GetLinkDefinitionOk

`func (o *UserDefinedLinkedResource) GetLinkDefinitionOk() (*UserDefinedLinkedResourceLinkDefinition, bool)`

GetLinkDefinitionOk returns a tuple with the LinkDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkDefinition

`func (o *UserDefinedLinkedResource) SetLinkDefinition(v UserDefinedLinkedResourceLinkDefinition)`

SetLinkDefinition sets LinkDefinition field to given value.

### HasLinkDefinition

`func (o *UserDefinedLinkedResource) HasLinkDefinition() bool`

HasLinkDefinition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


