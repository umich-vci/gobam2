# PostCollectionMergeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Networks** | [**[]IPv4NetworkMergeAllOfNetworks**](IPv4NetworkMergeAllOfNetworks.md) |  | 
**Destination** | [**InlinedIPv4DHCPRange**](InlinedIPv4DHCPRange.md) |  | 
**Ranges** | [**[]IPv4DHCPRangeMergeAllOfRanges**](IPv4DHCPRangeMergeAllOfRanges.md) |  | 

## Methods

### NewPostCollectionMergeRequest

`func NewPostCollectionMergeRequest(networks []IPv4NetworkMergeAllOfNetworks, destination InlinedIPv4DHCPRange, ranges []IPv4DHCPRangeMergeAllOfRanges, ) *PostCollectionMergeRequest`

NewPostCollectionMergeRequest instantiates a new PostCollectionMergeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostCollectionMergeRequestWithDefaults

`func NewPostCollectionMergeRequestWithDefaults() *PostCollectionMergeRequest`

NewPostCollectionMergeRequestWithDefaults instantiates a new PostCollectionMergeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PostCollectionMergeRequest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostCollectionMergeRequest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostCollectionMergeRequest) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *PostCollectionMergeRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *PostCollectionMergeRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PostCollectionMergeRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PostCollectionMergeRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PostCollectionMergeRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetNetworks

`func (o *PostCollectionMergeRequest) GetNetworks() []IPv4NetworkMergeAllOfNetworks`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *PostCollectionMergeRequest) GetNetworksOk() (*[]IPv4NetworkMergeAllOfNetworks, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *PostCollectionMergeRequest) SetNetworks(v []IPv4NetworkMergeAllOfNetworks)`

SetNetworks sets Networks field to given value.


### GetDestination

`func (o *PostCollectionMergeRequest) GetDestination() InlinedIPv4DHCPRange`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *PostCollectionMergeRequest) GetDestinationOk() (*InlinedIPv4DHCPRange, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *PostCollectionMergeRequest) SetDestination(v InlinedIPv4DHCPRange)`

SetDestination sets Destination field to given value.


### GetRanges

`func (o *PostCollectionMergeRequest) GetRanges() []IPv4DHCPRangeMergeAllOfRanges`

GetRanges returns the Ranges field if non-nil, zero value otherwise.

### GetRangesOk

`func (o *PostCollectionMergeRequest) GetRangesOk() (*[]IPv4DHCPRangeMergeAllOfRanges, bool)`

GetRangesOk returns a tuple with the Ranges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanges

`func (o *PostCollectionMergeRequest) SetRanges(v []IPv4DHCPRangeMergeAllOfRanges)`

SetRanges sets Ranges field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


