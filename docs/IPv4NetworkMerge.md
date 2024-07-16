# IPv4NetworkMerge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The resource type. | [optional] 
**Networks** | Pointer to [**[]IPv4NetworkMergeAllOfNetworks**](IPv4NetworkMergeAllOfNetworks.md) |  | [optional] 
**Destination** | Pointer to [**InlinedIPv4Network**](InlinedIPv4Network.md) |  | [optional] 

## Methods

### NewIPv4NetworkMerge

`func NewIPv4NetworkMerge() *IPv4NetworkMerge`

NewIPv4NetworkMerge instantiates a new IPv4NetworkMerge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPv4NetworkMergeWithDefaults

`func NewIPv4NetworkMergeWithDefaults() *IPv4NetworkMerge`

NewIPv4NetworkMergeWithDefaults instantiates a new IPv4NetworkMerge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *IPv4NetworkMerge) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IPv4NetworkMerge) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IPv4NetworkMerge) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IPv4NetworkMerge) HasType() bool`

HasType returns a boolean if a field has been set.

### GetNetworks

`func (o *IPv4NetworkMerge) GetNetworks() []IPv4NetworkMergeAllOfNetworks`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *IPv4NetworkMerge) GetNetworksOk() (*[]IPv4NetworkMergeAllOfNetworks, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *IPv4NetworkMerge) SetNetworks(v []IPv4NetworkMergeAllOfNetworks)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *IPv4NetworkMerge) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetDestination

`func (o *IPv4NetworkMerge) GetDestination() InlinedIPv4Network`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *IPv4NetworkMerge) GetDestinationOk() (*InlinedIPv4Network, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *IPv4NetworkMerge) SetDestination(v InlinedIPv4Network)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *IPv4NetworkMerge) HasDestination() bool`

HasDestination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


