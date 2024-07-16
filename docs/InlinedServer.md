# InlinedServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | Pointer to **string** | The name of the resource. | [optional] 
**Password** | Pointer to **string** | The password used to authenticate with the server. | [optional] 

## Methods

### NewInlinedServer

`func NewInlinedServer() *InlinedServer`

NewInlinedServer instantiates a new InlinedServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlinedServerWithDefaults

`func NewInlinedServerWithDefaults() *InlinedServer`

NewInlinedServerWithDefaults instantiates a new InlinedServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlinedServer) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlinedServer) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlinedServer) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *InlinedServer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *InlinedServer) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlinedServer) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlinedServer) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InlinedServer) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *InlinedServer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlinedServer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlinedServer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlinedServer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPassword

`func (o *InlinedServer) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *InlinedServer) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *InlinedServer) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *InlinedServer) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


