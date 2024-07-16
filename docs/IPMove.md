# IPMove

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The resource type. | [optional] 
**Address** | Pointer to **string** | The IP address of the destination resource where the IP resource is to be moved. | [optional] 
**MacAddress** | Pointer to [**IPMoveAllOfMacAddress**](IPMoveAllOfMacAddress.md) |  | [optional] 
**ServersUpdated** | Pointer to **bool** |  | [optional] 

## Methods

### NewIPMove

`func NewIPMove() *IPMove`

NewIPMove instantiates a new IPMove object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPMoveWithDefaults

`func NewIPMoveWithDefaults() *IPMove`

NewIPMoveWithDefaults instantiates a new IPMove object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *IPMove) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IPMove) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IPMove) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IPMove) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAddress

`func (o *IPMove) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *IPMove) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *IPMove) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *IPMove) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetMacAddress

`func (o *IPMove) GetMacAddress() IPMoveAllOfMacAddress`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *IPMove) GetMacAddressOk() (*IPMoveAllOfMacAddress, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *IPMove) SetMacAddress(v IPMoveAllOfMacAddress)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *IPMove) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetServersUpdated

`func (o *IPMove) GetServersUpdated() bool`

GetServersUpdated returns the ServersUpdated field if non-nil, zero value otherwise.

### GetServersUpdatedOk

`func (o *IPMove) GetServersUpdatedOk() (*bool, bool)`

GetServersUpdatedOk returns a tuple with the ServersUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServersUpdated

`func (o *IPMove) SetServersUpdated(v bool)`

SetServersUpdated sets ServersUpdated field to given value.

### HasServersUpdated

`func (o *IPMove) HasServersUpdated() bool`

HasServersUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


