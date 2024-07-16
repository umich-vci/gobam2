# MXRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The resource type. | [optional] 
**LinkedRecord** | Pointer to [**MXRecordAllOfLinkedRecord**](MXRecordAllOfLinkedRecord.md) |  | [optional] 
**Priority** | Pointer to **int32** | The priority assigned to the mail exchanger. A lower value indicates a higher priotiry. | [optional] 

## Methods

### NewMXRecord

`func NewMXRecord() *MXRecord`

NewMXRecord instantiates a new MXRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMXRecordWithDefaults

`func NewMXRecordWithDefaults() *MXRecord`

NewMXRecordWithDefaults instantiates a new MXRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MXRecord) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MXRecord) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MXRecord) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MXRecord) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinkedRecord

`func (o *MXRecord) GetLinkedRecord() MXRecordAllOfLinkedRecord`

GetLinkedRecord returns the LinkedRecord field if non-nil, zero value otherwise.

### GetLinkedRecordOk

`func (o *MXRecord) GetLinkedRecordOk() (*MXRecordAllOfLinkedRecord, bool)`

GetLinkedRecordOk returns a tuple with the LinkedRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedRecord

`func (o *MXRecord) SetLinkedRecord(v MXRecordAllOfLinkedRecord)`

SetLinkedRecord sets LinkedRecord field to given value.

### HasLinkedRecord

`func (o *MXRecord) HasLinkedRecord() bool`

HasLinkedRecord returns a boolean if a field has been set.

### GetPriority

`func (o *MXRecord) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *MXRecord) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *MXRecord) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *MXRecord) HasPriority() bool`

HasPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


