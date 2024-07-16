# GetUsers200Response1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** |  | [optional] [readonly] 
**TotalCount** | Pointer to **int32** |  | [optional] [readonly] 
**Data** | Pointer to [**[]User**](User.md) |  | [optional] [readonly] 

## Methods

### NewGetUsers200Response1

`func NewGetUsers200Response1() *GetUsers200Response1`

NewGetUsers200Response1 instantiates a new GetUsers200Response1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUsers200Response1WithDefaults

`func NewGetUsers200Response1WithDefaults() *GetUsers200Response1`

NewGetUsers200Response1WithDefaults instantiates a new GetUsers200Response1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *GetUsers200Response1) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GetUsers200Response1) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GetUsers200Response1) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *GetUsers200Response1) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetTotalCount

`func (o *GetUsers200Response1) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *GetUsers200Response1) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *GetUsers200Response1) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *GetUsers200Response1) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetData

`func (o *GetUsers200Response1) GetData() []User`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetUsers200Response1) GetDataOk() (*[]User, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetUsers200Response1) SetData(v []User)`

SetData sets Data field to given value.

### HasData

`func (o *GetUsers200Response1) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


