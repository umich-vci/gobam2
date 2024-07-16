# PostCollectionAddressRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The resource identifier. | 
**Type** | Pointer to **string** | The type of IP address. | [optional] 
**Address** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** | The state of the IP address. | [optional] 

## Methods

### NewPostCollectionAddressRequest

`func NewPostCollectionAddressRequest(id int64, ) *PostCollectionAddressRequest`

NewPostCollectionAddressRequest instantiates a new PostCollectionAddressRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostCollectionAddressRequestWithDefaults

`func NewPostCollectionAddressRequestWithDefaults() *PostCollectionAddressRequest`

NewPostCollectionAddressRequestWithDefaults instantiates a new PostCollectionAddressRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PostCollectionAddressRequest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostCollectionAddressRequest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostCollectionAddressRequest) SetId(v int64)`

SetId sets Id field to given value.


### GetType

`func (o *PostCollectionAddressRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PostCollectionAddressRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PostCollectionAddressRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PostCollectionAddressRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAddress

`func (o *PostCollectionAddressRequest) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *PostCollectionAddressRequest) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *PostCollectionAddressRequest) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *PostCollectionAddressRequest) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetState

`func (o *PostCollectionAddressRequest) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PostCollectionAddressRequest) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PostCollectionAddressRequest) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *PostCollectionAddressRequest) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


