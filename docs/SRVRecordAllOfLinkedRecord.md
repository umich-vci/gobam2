# SRVRecordAllOfLinkedRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The resource identifier. | 
**Type** | Pointer to **string** | The resource type. | [optional] 
**AbsoluteName** | **string** | The fully qualified domain name of the resource record. | 

## Methods

### NewSRVRecordAllOfLinkedRecord

`func NewSRVRecordAllOfLinkedRecord(id int64, absoluteName string, ) *SRVRecordAllOfLinkedRecord`

NewSRVRecordAllOfLinkedRecord instantiates a new SRVRecordAllOfLinkedRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSRVRecordAllOfLinkedRecordWithDefaults

`func NewSRVRecordAllOfLinkedRecordWithDefaults() *SRVRecordAllOfLinkedRecord`

NewSRVRecordAllOfLinkedRecordWithDefaults instantiates a new SRVRecordAllOfLinkedRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SRVRecordAllOfLinkedRecord) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SRVRecordAllOfLinkedRecord) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SRVRecordAllOfLinkedRecord) SetId(v int64)`

SetId sets Id field to given value.


### GetType

`func (o *SRVRecordAllOfLinkedRecord) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SRVRecordAllOfLinkedRecord) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SRVRecordAllOfLinkedRecord) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SRVRecordAllOfLinkedRecord) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAbsoluteName

`func (o *SRVRecordAllOfLinkedRecord) GetAbsoluteName() string`

GetAbsoluteName returns the AbsoluteName field if non-nil, zero value otherwise.

### GetAbsoluteNameOk

`func (o *SRVRecordAllOfLinkedRecord) GetAbsoluteNameOk() (*string, bool)`

GetAbsoluteNameOk returns a tuple with the AbsoluteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsoluteName

`func (o *SRVRecordAllOfLinkedRecord) SetAbsoluteName(v string)`

SetAbsoluteName sets AbsoluteName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


