# SRVRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The resource type. | [optional] 
**LinkedRecord** | Pointer to [**SRVRecordAllOfLinkedRecord**](SRVRecordAllOfLinkedRecord.md) |  | [optional] 
**Priority** | Pointer to **int32** | The priority assigned to the service. A lower value indicates higher priority. | [optional] 
**Weight** | Pointer to **int32** | The weight assigned to the service. The weight is referenced when services have the same priority. | [optional] 
**Port** | Pointer to **int32** | The port on which the service is located. | [optional] 

## Methods

### NewSRVRecord

`func NewSRVRecord() *SRVRecord`

NewSRVRecord instantiates a new SRVRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSRVRecordWithDefaults

`func NewSRVRecordWithDefaults() *SRVRecord`

NewSRVRecordWithDefaults instantiates a new SRVRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SRVRecord) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SRVRecord) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SRVRecord) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SRVRecord) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinkedRecord

`func (o *SRVRecord) GetLinkedRecord() SRVRecordAllOfLinkedRecord`

GetLinkedRecord returns the LinkedRecord field if non-nil, zero value otherwise.

### GetLinkedRecordOk

`func (o *SRVRecord) GetLinkedRecordOk() (*SRVRecordAllOfLinkedRecord, bool)`

GetLinkedRecordOk returns a tuple with the LinkedRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedRecord

`func (o *SRVRecord) SetLinkedRecord(v SRVRecordAllOfLinkedRecord)`

SetLinkedRecord sets LinkedRecord field to given value.

### HasLinkedRecord

`func (o *SRVRecord) HasLinkedRecord() bool`

HasLinkedRecord returns a boolean if a field has been set.

### GetPriority

`func (o *SRVRecord) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *SRVRecord) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *SRVRecord) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *SRVRecord) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetWeight

`func (o *SRVRecord) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *SRVRecord) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *SRVRecord) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *SRVRecord) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetPort

`func (o *SRVRecord) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SRVRecord) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *SRVRecord) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *SRVRecord) HasPort() bool`

HasPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


