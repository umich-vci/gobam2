# IPv6Address

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of IP address. | [optional] 
**State** | Pointer to **string** | The state of the IP address. | [optional] 
**ReservedUsing** | Pointer to **string** | The type of DHCP reservation. | [optional] 
**IdentityAssociationIdentifier** | Pointer to **string** | The Identity Association Identifier for the network interface of the client system. | [optional] [readonly] 
**ClientIdentifier** | Pointer to [**InlinedDHCPUniqueIdentifier**](InlinedDHCPUniqueIdentifier.md) |  | [optional] 
**InterfaceId** | Pointer to **string** | The ID of the interface from which the request came. | [optional] [readonly] 

## Methods

### NewIPv6Address

`func NewIPv6Address() *IPv6Address`

NewIPv6Address instantiates a new IPv6Address object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPv6AddressWithDefaults

`func NewIPv6AddressWithDefaults() *IPv6Address`

NewIPv6AddressWithDefaults instantiates a new IPv6Address object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *IPv6Address) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IPv6Address) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IPv6Address) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IPv6Address) HasType() bool`

HasType returns a boolean if a field has been set.

### GetState

`func (o *IPv6Address) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *IPv6Address) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *IPv6Address) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *IPv6Address) HasState() bool`

HasState returns a boolean if a field has been set.

### GetReservedUsing

`func (o *IPv6Address) GetReservedUsing() string`

GetReservedUsing returns the ReservedUsing field if non-nil, zero value otherwise.

### GetReservedUsingOk

`func (o *IPv6Address) GetReservedUsingOk() (*string, bool)`

GetReservedUsingOk returns a tuple with the ReservedUsing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedUsing

`func (o *IPv6Address) SetReservedUsing(v string)`

SetReservedUsing sets ReservedUsing field to given value.

### HasReservedUsing

`func (o *IPv6Address) HasReservedUsing() bool`

HasReservedUsing returns a boolean if a field has been set.

### GetIdentityAssociationIdentifier

`func (o *IPv6Address) GetIdentityAssociationIdentifier() string`

GetIdentityAssociationIdentifier returns the IdentityAssociationIdentifier field if non-nil, zero value otherwise.

### GetIdentityAssociationIdentifierOk

`func (o *IPv6Address) GetIdentityAssociationIdentifierOk() (*string, bool)`

GetIdentityAssociationIdentifierOk returns a tuple with the IdentityAssociationIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityAssociationIdentifier

`func (o *IPv6Address) SetIdentityAssociationIdentifier(v string)`

SetIdentityAssociationIdentifier sets IdentityAssociationIdentifier field to given value.

### HasIdentityAssociationIdentifier

`func (o *IPv6Address) HasIdentityAssociationIdentifier() bool`

HasIdentityAssociationIdentifier returns a boolean if a field has been set.

### GetClientIdentifier

`func (o *IPv6Address) GetClientIdentifier() InlinedDHCPUniqueIdentifier`

GetClientIdentifier returns the ClientIdentifier field if non-nil, zero value otherwise.

### GetClientIdentifierOk

`func (o *IPv6Address) GetClientIdentifierOk() (*InlinedDHCPUniqueIdentifier, bool)`

GetClientIdentifierOk returns a tuple with the ClientIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIdentifier

`func (o *IPv6Address) SetClientIdentifier(v InlinedDHCPUniqueIdentifier)`

SetClientIdentifier sets ClientIdentifier field to given value.

### HasClientIdentifier

`func (o *IPv6Address) HasClientIdentifier() bool`

HasClientIdentifier returns a boolean if a field has been set.

### GetInterfaceId

`func (o *IPv6Address) GetInterfaceId() string`

GetInterfaceId returns the InterfaceId field if non-nil, zero value otherwise.

### GetInterfaceIdOk

`func (o *IPv6Address) GetInterfaceIdOk() (*string, bool)`

GetInterfaceIdOk returns a tuple with the InterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceId

`func (o *IPv6Address) SetInterfaceId(v string)`

SetInterfaceId sets InterfaceId field to given value.

### HasInterfaceId

`func (o *IPv6Address) HasInterfaceId() bool`

HasInterfaceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


