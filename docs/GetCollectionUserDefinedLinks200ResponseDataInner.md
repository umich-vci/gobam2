# GetCollectionUserDefinedLinks200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**LinkDescription** | Pointer to **string** | The description for the user-defined link. | [optional] 
**LinkDefinition** | Pointer to [**UserDefinedLinkedResourceLinkDefinition**](UserDefinedLinkedResourceLinkDefinition.md) |  | [optional] 
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 

## Methods

### NewGetCollectionUserDefinedLinks200ResponseDataInner

`func NewGetCollectionUserDefinedLinks200ResponseDataInner() *GetCollectionUserDefinedLinks200ResponseDataInner`

NewGetCollectionUserDefinedLinks200ResponseDataInner instantiates a new GetCollectionUserDefinedLinks200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCollectionUserDefinedLinks200ResponseDataInnerWithDefaults

`func NewGetCollectionUserDefinedLinks200ResponseDataInnerWithDefaults() *GetCollectionUserDefinedLinks200ResponseDataInner`

NewGetCollectionUserDefinedLinks200ResponseDataInnerWithDefaults instantiates a new GetCollectionUserDefinedLinks200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetCollectionUserDefinedLinks200ResponseDataInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetCollectionUserDefinedLinks200ResponseDataInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetCollectionUserDefinedLinks200ResponseDataInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetCollectionUserDefinedLinks200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GetCollectionUserDefinedLinks200ResponseDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetCollectionUserDefinedLinks200ResponseDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetCollectionUserDefinedLinks200ResponseDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetCollectionUserDefinedLinks200ResponseDataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinkDescription

`func (o *GetCollectionUserDefinedLinks200ResponseDataInner) GetLinkDescription() string`

GetLinkDescription returns the LinkDescription field if non-nil, zero value otherwise.

### GetLinkDescriptionOk

`func (o *GetCollectionUserDefinedLinks200ResponseDataInner) GetLinkDescriptionOk() (*string, bool)`

GetLinkDescriptionOk returns a tuple with the LinkDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkDescription

`func (o *GetCollectionUserDefinedLinks200ResponseDataInner) SetLinkDescription(v string)`

SetLinkDescription sets LinkDescription field to given value.

### HasLinkDescription

`func (o *GetCollectionUserDefinedLinks200ResponseDataInner) HasLinkDescription() bool`

HasLinkDescription returns a boolean if a field has been set.

### GetLinkDefinition

`func (o *GetCollectionUserDefinedLinks200ResponseDataInner) GetLinkDefinition() UserDefinedLinkedResourceLinkDefinition`

GetLinkDefinition returns the LinkDefinition field if non-nil, zero value otherwise.

### GetLinkDefinitionOk

`func (o *GetCollectionUserDefinedLinks200ResponseDataInner) GetLinkDefinitionOk() (*UserDefinedLinkedResourceLinkDefinition, bool)`

GetLinkDefinitionOk returns a tuple with the LinkDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkDefinition

`func (o *GetCollectionUserDefinedLinks200ResponseDataInner) SetLinkDefinition(v UserDefinedLinkedResourceLinkDefinition)`

SetLinkDefinition sets LinkDefinition field to given value.

### HasLinkDefinition

`func (o *GetCollectionUserDefinedLinks200ResponseDataInner) HasLinkDefinition() bool`

HasLinkDefinition returns a boolean if a field has been set.

### GetLinks

`func (o *GetCollectionUserDefinedLinks200ResponseDataInner) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetCollectionUserDefinedLinks200ResponseDataInner) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetCollectionUserDefinedLinks200ResponseDataInner) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetCollectionUserDefinedLinks200ResponseDataInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


