# AliasRecordAllOfLinkedRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The resource identifier. | 
**Type** | Pointer to **string** | The resource type. | [optional] 
**AbsoluteName** | **string** | The fully qualified domain name of the resource record. | 

## Methods

### NewAliasRecordAllOfLinkedRecord

`func NewAliasRecordAllOfLinkedRecord(id int64, absoluteName string, ) *AliasRecordAllOfLinkedRecord`

NewAliasRecordAllOfLinkedRecord instantiates a new AliasRecordAllOfLinkedRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAliasRecordAllOfLinkedRecordWithDefaults

`func NewAliasRecordAllOfLinkedRecordWithDefaults() *AliasRecordAllOfLinkedRecord`

NewAliasRecordAllOfLinkedRecordWithDefaults instantiates a new AliasRecordAllOfLinkedRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AliasRecordAllOfLinkedRecord) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AliasRecordAllOfLinkedRecord) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AliasRecordAllOfLinkedRecord) SetId(v int64)`

SetId sets Id field to given value.


### GetType

`func (o *AliasRecordAllOfLinkedRecord) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AliasRecordAllOfLinkedRecord) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AliasRecordAllOfLinkedRecord) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AliasRecordAllOfLinkedRecord) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAbsoluteName

`func (o *AliasRecordAllOfLinkedRecord) GetAbsoluteName() string`

GetAbsoluteName returns the AbsoluteName field if non-nil, zero value otherwise.

### GetAbsoluteNameOk

`func (o *AliasRecordAllOfLinkedRecord) GetAbsoluteNameOk() (*string, bool)`

GetAbsoluteNameOk returns a tuple with the AbsoluteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsoluteName

`func (o *AliasRecordAllOfLinkedRecord) SetAbsoluteName(v string)`

SetAbsoluteName sets AbsoluteName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


