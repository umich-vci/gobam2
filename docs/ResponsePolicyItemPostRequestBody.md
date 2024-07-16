# ResponsePolicyItemPostRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | **string** | The fully qualified domain name to be blocked or redirected by the responsepolicy. | 

## Methods

### NewResponsePolicyItemPostRequestBody

`func NewResponsePolicyItemPostRequestBody(name string, ) *ResponsePolicyItemPostRequestBody`

NewResponsePolicyItemPostRequestBody instantiates a new ResponsePolicyItemPostRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsePolicyItemPostRequestBodyWithDefaults

`func NewResponsePolicyItemPostRequestBodyWithDefaults() *ResponsePolicyItemPostRequestBody`

NewResponsePolicyItemPostRequestBodyWithDefaults instantiates a new ResponsePolicyItemPostRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResponsePolicyItemPostRequestBody) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponsePolicyItemPostRequestBody) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponsePolicyItemPostRequestBody) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ResponsePolicyItemPostRequestBody) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ResponsePolicyItemPostRequestBody) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResponsePolicyItemPostRequestBody) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResponsePolicyItemPostRequestBody) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ResponsePolicyItemPostRequestBody) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *ResponsePolicyItemPostRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResponsePolicyItemPostRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResponsePolicyItemPostRequestBody) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


