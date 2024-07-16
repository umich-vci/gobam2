# AliasRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The resource type. | [optional] 
**LinkedRecord** | Pointer to [**AliasRecordAllOfLinkedRecord**](AliasRecordAllOfLinkedRecord.md) |  | [optional] 

## Methods

### NewAliasRecord

`func NewAliasRecord() *AliasRecord`

NewAliasRecord instantiates a new AliasRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAliasRecordWithDefaults

`func NewAliasRecordWithDefaults() *AliasRecord`

NewAliasRecordWithDefaults instantiates a new AliasRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AliasRecord) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AliasRecord) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AliasRecord) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AliasRecord) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinkedRecord

`func (o *AliasRecord) GetLinkedRecord() AliasRecordAllOfLinkedRecord`

GetLinkedRecord returns the LinkedRecord field if non-nil, zero value otherwise.

### GetLinkedRecordOk

`func (o *AliasRecord) GetLinkedRecordOk() (*AliasRecordAllOfLinkedRecord, bool)`

GetLinkedRecordOk returns a tuple with the LinkedRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedRecord

`func (o *AliasRecord) SetLinkedRecord(v AliasRecordAllOfLinkedRecord)`

SetLinkedRecord sets LinkedRecord field to given value.

### HasLinkedRecord

`func (o *AliasRecord) HasLinkedRecord() bool`

HasLinkedRecord returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


