# UserDefinedLinkedResourcePostRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The resource identifier. | 
**Type** | Pointer to **string** | The resource type. | [optional] 
**LinkDescription** | Pointer to **string** | The description for the user-defined link. | [optional] 
**LinkDefinition** | [**UserDefinedLinkedResourceLinkDefinition**](UserDefinedLinkedResourceLinkDefinition.md) |  | 

## Methods

### NewUserDefinedLinkedResourcePostRequestBody

`func NewUserDefinedLinkedResourcePostRequestBody(id int64, linkDefinition UserDefinedLinkedResourceLinkDefinition, ) *UserDefinedLinkedResourcePostRequestBody`

NewUserDefinedLinkedResourcePostRequestBody instantiates a new UserDefinedLinkedResourcePostRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserDefinedLinkedResourcePostRequestBodyWithDefaults

`func NewUserDefinedLinkedResourcePostRequestBodyWithDefaults() *UserDefinedLinkedResourcePostRequestBody`

NewUserDefinedLinkedResourcePostRequestBodyWithDefaults instantiates a new UserDefinedLinkedResourcePostRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserDefinedLinkedResourcePostRequestBody) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserDefinedLinkedResourcePostRequestBody) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserDefinedLinkedResourcePostRequestBody) SetId(v int64)`

SetId sets Id field to given value.


### GetType

`func (o *UserDefinedLinkedResourcePostRequestBody) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserDefinedLinkedResourcePostRequestBody) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserDefinedLinkedResourcePostRequestBody) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UserDefinedLinkedResourcePostRequestBody) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinkDescription

`func (o *UserDefinedLinkedResourcePostRequestBody) GetLinkDescription() string`

GetLinkDescription returns the LinkDescription field if non-nil, zero value otherwise.

### GetLinkDescriptionOk

`func (o *UserDefinedLinkedResourcePostRequestBody) GetLinkDescriptionOk() (*string, bool)`

GetLinkDescriptionOk returns a tuple with the LinkDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkDescription

`func (o *UserDefinedLinkedResourcePostRequestBody) SetLinkDescription(v string)`

SetLinkDescription sets LinkDescription field to given value.

### HasLinkDescription

`func (o *UserDefinedLinkedResourcePostRequestBody) HasLinkDescription() bool`

HasLinkDescription returns a boolean if a field has been set.

### GetLinkDefinition

`func (o *UserDefinedLinkedResourcePostRequestBody) GetLinkDefinition() UserDefinedLinkedResourceLinkDefinition`

GetLinkDefinition returns the LinkDefinition field if non-nil, zero value otherwise.

### GetLinkDefinitionOk

`func (o *UserDefinedLinkedResourcePostRequestBody) GetLinkDefinitionOk() (*UserDefinedLinkedResourceLinkDefinition, bool)`

GetLinkDefinitionOk returns a tuple with the LinkDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkDefinition

`func (o *UserDefinedLinkedResourcePostRequestBody) SetLinkDefinition(v UserDefinedLinkedResourceLinkDefinition)`

SetLinkDefinition sets LinkDefinition field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


