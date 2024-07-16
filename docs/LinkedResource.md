# LinkedResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 

## Methods

### NewLinkedResource

`func NewLinkedResource() *LinkedResource`

NewLinkedResource instantiates a new LinkedResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkedResourceWithDefaults

`func NewLinkedResourceWithDefaults() *LinkedResource`

NewLinkedResourceWithDefaults instantiates a new LinkedResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LinkedResource) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LinkedResource) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LinkedResource) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *LinkedResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *LinkedResource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LinkedResource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LinkedResource) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *LinkedResource) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


