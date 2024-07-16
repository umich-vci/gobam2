# AddressMacAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The resource identifier. | 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | Pointer to **string** | The name of the resource. | [optional] 
**Address** | **string** | The MAC address. | 

## Methods

### NewAddressMacAddress

`func NewAddressMacAddress(id int64, address string, ) *AddressMacAddress`

NewAddressMacAddress instantiates a new AddressMacAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressMacAddressWithDefaults

`func NewAddressMacAddressWithDefaults() *AddressMacAddress`

NewAddressMacAddressWithDefaults instantiates a new AddressMacAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddressMacAddress) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddressMacAddress) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddressMacAddress) SetId(v int64)`

SetId sets Id field to given value.


### GetType

`func (o *AddressMacAddress) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddressMacAddress) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddressMacAddress) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AddressMacAddress) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *AddressMacAddress) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddressMacAddress) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddressMacAddress) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddressMacAddress) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAddress

`func (o *AddressMacAddress) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AddressMacAddress) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AddressMacAddress) SetAddress(v string)`

SetAddress sets Address field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


