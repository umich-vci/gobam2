# GetMerges200Response1DataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Networks** | Pointer to [**[]IPv4NetworkMergeAllOfNetworks**](IPv4NetworkMergeAllOfNetworks.md) |  | [optional] 
**Destination** | Pointer to [**InlinedIPv4DHCPRange**](InlinedIPv4DHCPRange.md) |  | [optional] 
**Ranges** | Pointer to [**[]IPv4DHCPRangeMergeAllOfRanges**](IPv4DHCPRangeMergeAllOfRanges.md) |  | [optional] 

## Methods

### NewGetMerges200Response1DataInner

`func NewGetMerges200Response1DataInner() *GetMerges200Response1DataInner`

NewGetMerges200Response1DataInner instantiates a new GetMerges200Response1DataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMerges200Response1DataInnerWithDefaults

`func NewGetMerges200Response1DataInnerWithDefaults() *GetMerges200Response1DataInner`

NewGetMerges200Response1DataInnerWithDefaults instantiates a new GetMerges200Response1DataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetMerges200Response1DataInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetMerges200Response1DataInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetMerges200Response1DataInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetMerges200Response1DataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GetMerges200Response1DataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetMerges200Response1DataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetMerges200Response1DataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetMerges200Response1DataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetNetworks

`func (o *GetMerges200Response1DataInner) GetNetworks() []IPv4NetworkMergeAllOfNetworks`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *GetMerges200Response1DataInner) GetNetworksOk() (*[]IPv4NetworkMergeAllOfNetworks, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *GetMerges200Response1DataInner) SetNetworks(v []IPv4NetworkMergeAllOfNetworks)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *GetMerges200Response1DataInner) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetDestination

`func (o *GetMerges200Response1DataInner) GetDestination() InlinedIPv4DHCPRange`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *GetMerges200Response1DataInner) GetDestinationOk() (*InlinedIPv4DHCPRange, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *GetMerges200Response1DataInner) SetDestination(v InlinedIPv4DHCPRange)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *GetMerges200Response1DataInner) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetRanges

`func (o *GetMerges200Response1DataInner) GetRanges() []IPv4DHCPRangeMergeAllOfRanges`

GetRanges returns the Ranges field if non-nil, zero value otherwise.

### GetRangesOk

`func (o *GetMerges200Response1DataInner) GetRangesOk() (*[]IPv4DHCPRangeMergeAllOfRanges, bool)`

GetRangesOk returns a tuple with the Ranges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanges

`func (o *GetMerges200Response1DataInner) SetRanges(v []IPv4DHCPRangeMergeAllOfRanges)`

SetRanges sets Ranges field to given value.

### HasRanges

`func (o *GetMerges200Response1DataInner) HasRanges() bool`

HasRanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


