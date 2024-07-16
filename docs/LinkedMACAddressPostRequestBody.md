# LinkedMACAddressPostRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The resource identifier. | 
**Type** | Pointer to **string** | The resource type. | [optional] 

## Methods

### NewLinkedMACAddressPostRequestBody

`func NewLinkedMACAddressPostRequestBody(id int64, ) *LinkedMACAddressPostRequestBody`

NewLinkedMACAddressPostRequestBody instantiates a new LinkedMACAddressPostRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkedMACAddressPostRequestBodyWithDefaults

`func NewLinkedMACAddressPostRequestBodyWithDefaults() *LinkedMACAddressPostRequestBody`

NewLinkedMACAddressPostRequestBodyWithDefaults instantiates a new LinkedMACAddressPostRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LinkedMACAddressPostRequestBody) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LinkedMACAddressPostRequestBody) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LinkedMACAddressPostRequestBody) SetId(v int64)`

SetId sets Id field to given value.


### GetType

`func (o *LinkedMACAddressPostRequestBody) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LinkedMACAddressPostRequestBody) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LinkedMACAddressPostRequestBody) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *LinkedMACAddressPostRequestBody) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


