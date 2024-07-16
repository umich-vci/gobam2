# LinkedTag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 

## Methods

### NewLinkedTag

`func NewLinkedTag() *LinkedTag`

NewLinkedTag instantiates a new LinkedTag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkedTagWithDefaults

`func NewLinkedTagWithDefaults() *LinkedTag`

NewLinkedTagWithDefaults instantiates a new LinkedTag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LinkedTag) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LinkedTag) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LinkedTag) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *LinkedTag) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *LinkedTag) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LinkedTag) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LinkedTag) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *LinkedTag) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


