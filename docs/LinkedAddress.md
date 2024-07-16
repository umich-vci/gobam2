# LinkedAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The type of IP address. | [optional] 
**Address** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** | The state of the IP address. | [optional] 

## Methods

### NewLinkedAddress

`func NewLinkedAddress() *LinkedAddress`

NewLinkedAddress instantiates a new LinkedAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkedAddressWithDefaults

`func NewLinkedAddressWithDefaults() *LinkedAddress`

NewLinkedAddressWithDefaults instantiates a new LinkedAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LinkedAddress) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LinkedAddress) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LinkedAddress) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *LinkedAddress) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *LinkedAddress) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LinkedAddress) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LinkedAddress) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *LinkedAddress) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAddress

`func (o *LinkedAddress) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *LinkedAddress) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *LinkedAddress) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *LinkedAddress) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetState

`func (o *LinkedAddress) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *LinkedAddress) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *LinkedAddress) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *LinkedAddress) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


