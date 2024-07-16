# UserDefinedLinkedResourceLinkDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The resource identifier. | 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | **string** | The name of the user-defined link definition. | 

## Methods

### NewUserDefinedLinkedResourceLinkDefinition

`func NewUserDefinedLinkedResourceLinkDefinition(id int64, name string, ) *UserDefinedLinkedResourceLinkDefinition`

NewUserDefinedLinkedResourceLinkDefinition instantiates a new UserDefinedLinkedResourceLinkDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserDefinedLinkedResourceLinkDefinitionWithDefaults

`func NewUserDefinedLinkedResourceLinkDefinitionWithDefaults() *UserDefinedLinkedResourceLinkDefinition`

NewUserDefinedLinkedResourceLinkDefinitionWithDefaults instantiates a new UserDefinedLinkedResourceLinkDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserDefinedLinkedResourceLinkDefinition) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserDefinedLinkedResourceLinkDefinition) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserDefinedLinkedResourceLinkDefinition) SetId(v int64)`

SetId sets Id field to given value.


### GetType

`func (o *UserDefinedLinkedResourceLinkDefinition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserDefinedLinkedResourceLinkDefinition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserDefinedLinkedResourceLinkDefinition) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UserDefinedLinkedResourceLinkDefinition) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *UserDefinedLinkedResourceLinkDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserDefinedLinkedResourceLinkDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserDefinedLinkedResourceLinkDefinition) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


