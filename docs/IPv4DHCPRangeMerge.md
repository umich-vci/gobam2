# IPv4DHCPRangeMerge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The resource type. | [optional] 
**Ranges** | Pointer to [**[]IPv4DHCPRangeMergeAllOfRanges**](IPv4DHCPRangeMergeAllOfRanges.md) |  | [optional] 
**Destination** | Pointer to [**InlinedIPv4DHCPRange**](InlinedIPv4DHCPRange.md) |  | [optional] 

## Methods

### NewIPv4DHCPRangeMerge

`func NewIPv4DHCPRangeMerge() *IPv4DHCPRangeMerge`

NewIPv4DHCPRangeMerge instantiates a new IPv4DHCPRangeMerge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPv4DHCPRangeMergeWithDefaults

`func NewIPv4DHCPRangeMergeWithDefaults() *IPv4DHCPRangeMerge`

NewIPv4DHCPRangeMergeWithDefaults instantiates a new IPv4DHCPRangeMerge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *IPv4DHCPRangeMerge) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IPv4DHCPRangeMerge) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IPv4DHCPRangeMerge) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IPv4DHCPRangeMerge) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRanges

`func (o *IPv4DHCPRangeMerge) GetRanges() []IPv4DHCPRangeMergeAllOfRanges`

GetRanges returns the Ranges field if non-nil, zero value otherwise.

### GetRangesOk

`func (o *IPv4DHCPRangeMerge) GetRangesOk() (*[]IPv4DHCPRangeMergeAllOfRanges, bool)`

GetRangesOk returns a tuple with the Ranges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanges

`func (o *IPv4DHCPRangeMerge) SetRanges(v []IPv4DHCPRangeMergeAllOfRanges)`

SetRanges sets Ranges field to given value.

### HasRanges

`func (o *IPv4DHCPRangeMerge) HasRanges() bool`

HasRanges returns a boolean if a field has been set.

### GetDestination

`func (o *IPv4DHCPRangeMerge) GetDestination() InlinedIPv4DHCPRange`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *IPv4DHCPRangeMerge) GetDestinationOk() (*InlinedIPv4DHCPRange, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *IPv4DHCPRangeMerge) SetDestination(v InlinedIPv4DHCPRange)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *IPv4DHCPRangeMerge) HasDestination() bool`

HasDestination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


