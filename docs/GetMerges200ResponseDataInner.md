# GetMerges200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Blocks** | Pointer to [**[]IPv4BlockMergeAllOfBlocks**](IPv4BlockMergeAllOfBlocks.md) |  | [optional] 
**Destination** | Pointer to [**InlinedIPv4DHCPRange**](InlinedIPv4DHCPRange.md) |  | [optional] 
**Networks** | Pointer to [**[]IPv4NetworkMergeAllOfNetworks**](IPv4NetworkMergeAllOfNetworks.md) |  | [optional] 
**Ranges** | Pointer to [**[]IPv4DHCPRangeMergeAllOfRanges**](IPv4DHCPRangeMergeAllOfRanges.md) |  | [optional] 
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 

## Methods

### NewGetMerges200ResponseDataInner

`func NewGetMerges200ResponseDataInner() *GetMerges200ResponseDataInner`

NewGetMerges200ResponseDataInner instantiates a new GetMerges200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMerges200ResponseDataInnerWithDefaults

`func NewGetMerges200ResponseDataInnerWithDefaults() *GetMerges200ResponseDataInner`

NewGetMerges200ResponseDataInnerWithDefaults instantiates a new GetMerges200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetMerges200ResponseDataInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetMerges200ResponseDataInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetMerges200ResponseDataInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetMerges200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GetMerges200ResponseDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetMerges200ResponseDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetMerges200ResponseDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetMerges200ResponseDataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetBlocks

`func (o *GetMerges200ResponseDataInner) GetBlocks() []IPv4BlockMergeAllOfBlocks`

GetBlocks returns the Blocks field if non-nil, zero value otherwise.

### GetBlocksOk

`func (o *GetMerges200ResponseDataInner) GetBlocksOk() (*[]IPv4BlockMergeAllOfBlocks, bool)`

GetBlocksOk returns a tuple with the Blocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocks

`func (o *GetMerges200ResponseDataInner) SetBlocks(v []IPv4BlockMergeAllOfBlocks)`

SetBlocks sets Blocks field to given value.

### HasBlocks

`func (o *GetMerges200ResponseDataInner) HasBlocks() bool`

HasBlocks returns a boolean if a field has been set.

### GetDestination

`func (o *GetMerges200ResponseDataInner) GetDestination() InlinedIPv4DHCPRange`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *GetMerges200ResponseDataInner) GetDestinationOk() (*InlinedIPv4DHCPRange, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *GetMerges200ResponseDataInner) SetDestination(v InlinedIPv4DHCPRange)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *GetMerges200ResponseDataInner) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetNetworks

`func (o *GetMerges200ResponseDataInner) GetNetworks() []IPv4NetworkMergeAllOfNetworks`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *GetMerges200ResponseDataInner) GetNetworksOk() (*[]IPv4NetworkMergeAllOfNetworks, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *GetMerges200ResponseDataInner) SetNetworks(v []IPv4NetworkMergeAllOfNetworks)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *GetMerges200ResponseDataInner) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetRanges

`func (o *GetMerges200ResponseDataInner) GetRanges() []IPv4DHCPRangeMergeAllOfRanges`

GetRanges returns the Ranges field if non-nil, zero value otherwise.

### GetRangesOk

`func (o *GetMerges200ResponseDataInner) GetRangesOk() (*[]IPv4DHCPRangeMergeAllOfRanges, bool)`

GetRangesOk returns a tuple with the Ranges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanges

`func (o *GetMerges200ResponseDataInner) SetRanges(v []IPv4DHCPRangeMergeAllOfRanges)`

SetRanges sets Ranges field to given value.

### HasRanges

`func (o *GetMerges200ResponseDataInner) HasRanges() bool`

HasRanges returns a boolean if a field has been set.

### GetLinks

`func (o *GetMerges200ResponseDataInner) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetMerges200ResponseDataInner) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetMerges200ResponseDataInner) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetMerges200ResponseDataInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


