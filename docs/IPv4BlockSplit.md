# IPv4BlockSplit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The resource type. | [optional] 
**Block** | Pointer to [**InlinedIPv4Block**](InlinedIPv4Block.md) |  | [optional] 
**SplitPointAddress** | Pointer to **string** | The IP address of the point at which you would like to split the block. | [optional] 

## Methods

### NewIPv4BlockSplit

`func NewIPv4BlockSplit() *IPv4BlockSplit`

NewIPv4BlockSplit instantiates a new IPv4BlockSplit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPv4BlockSplitWithDefaults

`func NewIPv4BlockSplitWithDefaults() *IPv4BlockSplit`

NewIPv4BlockSplitWithDefaults instantiates a new IPv4BlockSplit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *IPv4BlockSplit) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IPv4BlockSplit) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IPv4BlockSplit) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IPv4BlockSplit) HasType() bool`

HasType returns a boolean if a field has been set.

### GetBlock

`func (o *IPv4BlockSplit) GetBlock() InlinedIPv4Block`

GetBlock returns the Block field if non-nil, zero value otherwise.

### GetBlockOk

`func (o *IPv4BlockSplit) GetBlockOk() (*InlinedIPv4Block, bool)`

GetBlockOk returns a tuple with the Block field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlock

`func (o *IPv4BlockSplit) SetBlock(v InlinedIPv4Block)`

SetBlock sets Block field to given value.

### HasBlock

`func (o *IPv4BlockSplit) HasBlock() bool`

HasBlock returns a boolean if a field has been set.

### GetSplitPointAddress

`func (o *IPv4BlockSplit) GetSplitPointAddress() string`

GetSplitPointAddress returns the SplitPointAddress field if non-nil, zero value otherwise.

### GetSplitPointAddressOk

`func (o *IPv4BlockSplit) GetSplitPointAddressOk() (*string, bool)`

GetSplitPointAddressOk returns a tuple with the SplitPointAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitPointAddress

`func (o *IPv4BlockSplit) SetSplitPointAddress(v string)`

SetSplitPointAddress sets SplitPointAddress field to given value.

### HasSplitPointAddress

`func (o *IPv4BlockSplit) HasSplitPointAddress() bool`

HasSplitPointAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


