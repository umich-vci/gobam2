# GetZoneGroups200Response1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** |  | [optional] [readonly] 
**TotalCount** | Pointer to **int32** |  | [optional] [readonly] 
**Data** | Pointer to [**[]DHCPZoneGroup**](DHCPZoneGroup.md) |  | [optional] [readonly] 

## Methods

### NewGetZoneGroups200Response1

`func NewGetZoneGroups200Response1() *GetZoneGroups200Response1`

NewGetZoneGroups200Response1 instantiates a new GetZoneGroups200Response1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetZoneGroups200Response1WithDefaults

`func NewGetZoneGroups200Response1WithDefaults() *GetZoneGroups200Response1`

NewGetZoneGroups200Response1WithDefaults instantiates a new GetZoneGroups200Response1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *GetZoneGroups200Response1) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GetZoneGroups200Response1) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GetZoneGroups200Response1) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *GetZoneGroups200Response1) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetTotalCount

`func (o *GetZoneGroups200Response1) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *GetZoneGroups200Response1) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *GetZoneGroups200Response1) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *GetZoneGroups200Response1) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetData

`func (o *GetZoneGroups200Response1) GetData() []DHCPZoneGroup`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetZoneGroups200Response1) GetDataOk() (*[]DHCPZoneGroup, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetZoneGroups200Response1) SetData(v []DHCPZoneGroup)`

SetData sets Data field to given value.

### HasData

`func (o *GetZoneGroups200Response1) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


