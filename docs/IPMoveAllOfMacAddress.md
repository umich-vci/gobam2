# IPMoveAllOfMacAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The resource identifier. | 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | Pointer to **string** | The name of the resource. | [optional] 
**Address** | **string** | The MAC address. | 

## Methods

### NewIPMoveAllOfMacAddress

`func NewIPMoveAllOfMacAddress(id int64, address string, ) *IPMoveAllOfMacAddress`

NewIPMoveAllOfMacAddress instantiates a new IPMoveAllOfMacAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPMoveAllOfMacAddressWithDefaults

`func NewIPMoveAllOfMacAddressWithDefaults() *IPMoveAllOfMacAddress`

NewIPMoveAllOfMacAddressWithDefaults instantiates a new IPMoveAllOfMacAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IPMoveAllOfMacAddress) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IPMoveAllOfMacAddress) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IPMoveAllOfMacAddress) SetId(v int64)`

SetId sets Id field to given value.


### GetType

`func (o *IPMoveAllOfMacAddress) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IPMoveAllOfMacAddress) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IPMoveAllOfMacAddress) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IPMoveAllOfMacAddress) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *IPMoveAllOfMacAddress) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IPMoveAllOfMacAddress) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IPMoveAllOfMacAddress) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IPMoveAllOfMacAddress) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAddress

`func (o *IPMoveAllOfMacAddress) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *IPMoveAllOfMacAddress) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *IPMoveAllOfMacAddress) SetAddress(v string)`

SetAddress sets Address field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


