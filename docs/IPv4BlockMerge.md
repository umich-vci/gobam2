# IPv4BlockMerge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The resource type. | [optional] 
**Blocks** | Pointer to [**[]IPv4BlockMergeAllOfBlocks**](IPv4BlockMergeAllOfBlocks.md) |  | [optional] 
**Destination** | Pointer to [**InlinedIpv4BlockMergeDestination**](InlinedIpv4BlockMergeDestination.md) |  | [optional] 

## Methods

### NewIPv4BlockMerge

`func NewIPv4BlockMerge() *IPv4BlockMerge`

NewIPv4BlockMerge instantiates a new IPv4BlockMerge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPv4BlockMergeWithDefaults

`func NewIPv4BlockMergeWithDefaults() *IPv4BlockMerge`

NewIPv4BlockMergeWithDefaults instantiates a new IPv4BlockMerge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *IPv4BlockMerge) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IPv4BlockMerge) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IPv4BlockMerge) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IPv4BlockMerge) HasType() bool`

HasType returns a boolean if a field has been set.

### GetBlocks

`func (o *IPv4BlockMerge) GetBlocks() []IPv4BlockMergeAllOfBlocks`

GetBlocks returns the Blocks field if non-nil, zero value otherwise.

### GetBlocksOk

`func (o *IPv4BlockMerge) GetBlocksOk() (*[]IPv4BlockMergeAllOfBlocks, bool)`

GetBlocksOk returns a tuple with the Blocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocks

`func (o *IPv4BlockMerge) SetBlocks(v []IPv4BlockMergeAllOfBlocks)`

SetBlocks sets Blocks field to given value.

### HasBlocks

`func (o *IPv4BlockMerge) HasBlocks() bool`

HasBlocks returns a boolean if a field has been set.

### GetDestination

`func (o *IPv4BlockMerge) GetDestination() InlinedIpv4BlockMergeDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *IPv4BlockMerge) GetDestinationOk() (*InlinedIpv4BlockMergeDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *IPv4BlockMerge) SetDestination(v InlinedIpv4BlockMergeDestination)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *IPv4BlockMerge) HasDestination() bool`

HasDestination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


