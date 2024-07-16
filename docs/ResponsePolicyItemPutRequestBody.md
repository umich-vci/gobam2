# ResponsePolicyItemPutRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | **string** | The fully qualified domain name to be blocked or redirected by the responsepolicy. | 

## Methods

### NewResponsePolicyItemPutRequestBody

`func NewResponsePolicyItemPutRequestBody(name string, ) *ResponsePolicyItemPutRequestBody`

NewResponsePolicyItemPutRequestBody instantiates a new ResponsePolicyItemPutRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsePolicyItemPutRequestBodyWithDefaults

`func NewResponsePolicyItemPutRequestBodyWithDefaults() *ResponsePolicyItemPutRequestBody`

NewResponsePolicyItemPutRequestBodyWithDefaults instantiates a new ResponsePolicyItemPutRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResponsePolicyItemPutRequestBody) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponsePolicyItemPutRequestBody) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponsePolicyItemPutRequestBody) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ResponsePolicyItemPutRequestBody) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ResponsePolicyItemPutRequestBody) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResponsePolicyItemPutRequestBody) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResponsePolicyItemPutRequestBody) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ResponsePolicyItemPutRequestBody) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *ResponsePolicyItemPutRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResponsePolicyItemPutRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResponsePolicyItemPutRequestBody) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


