# LinkedUserPostRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The resource identifier. | 
**Type** | Pointer to **string** | The resource type. | [optional] 

## Methods

### NewLinkedUserPostRequestBody

`func NewLinkedUserPostRequestBody(id int64, ) *LinkedUserPostRequestBody`

NewLinkedUserPostRequestBody instantiates a new LinkedUserPostRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkedUserPostRequestBodyWithDefaults

`func NewLinkedUserPostRequestBodyWithDefaults() *LinkedUserPostRequestBody`

NewLinkedUserPostRequestBodyWithDefaults instantiates a new LinkedUserPostRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LinkedUserPostRequestBody) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LinkedUserPostRequestBody) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LinkedUserPostRequestBody) SetId(v int64)`

SetId sets Id field to given value.


### GetType

`func (o *LinkedUserPostRequestBody) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LinkedUserPostRequestBody) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LinkedUserPostRequestBody) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *LinkedUserPostRequestBody) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

