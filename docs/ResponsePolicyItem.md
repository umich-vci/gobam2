# ResponsePolicyItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | Pointer to **string** | The fully qualified domain name to be blocked or redirected by the responsepolicy. | [optional] 

## Methods

### NewResponsePolicyItem

`func NewResponsePolicyItem() *ResponsePolicyItem`

NewResponsePolicyItem instantiates a new ResponsePolicyItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsePolicyItemWithDefaults

`func NewResponsePolicyItemWithDefaults() *ResponsePolicyItem`

NewResponsePolicyItemWithDefaults instantiates a new ResponsePolicyItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResponsePolicyItem) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponsePolicyItem) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponsePolicyItem) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ResponsePolicyItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ResponsePolicyItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResponsePolicyItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResponsePolicyItem) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ResponsePolicyItem) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *ResponsePolicyItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResponsePolicyItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResponsePolicyItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ResponsePolicyItem) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


