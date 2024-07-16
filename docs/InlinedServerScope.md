# InlinedServerScope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The server scope resource type. | [optional] 
**Name** | Pointer to **string** | The name of the resource. | [optional] [readonly] 

## Methods

### NewInlinedServerScope

`func NewInlinedServerScope() *InlinedServerScope`

NewInlinedServerScope instantiates a new InlinedServerScope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlinedServerScopeWithDefaults

`func NewInlinedServerScopeWithDefaults() *InlinedServerScope`

NewInlinedServerScopeWithDefaults instantiates a new InlinedServerScope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlinedServerScope) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlinedServerScope) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlinedServerScope) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *InlinedServerScope) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *InlinedServerScope) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlinedServerScope) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlinedServerScope) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InlinedServerScope) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *InlinedServerScope) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlinedServerScope) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlinedServerScope) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlinedServerScope) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

