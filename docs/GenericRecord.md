# GenericRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The resource type. | [optional] 
**RecordType** | Pointer to **string** | The resource record type. | [optional] 
**Rdata** | Pointer to **string** | The resource record data. | [optional] 

## Methods

### NewGenericRecord

`func NewGenericRecord() *GenericRecord`

NewGenericRecord instantiates a new GenericRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenericRecordWithDefaults

`func NewGenericRecordWithDefaults() *GenericRecord`

NewGenericRecordWithDefaults instantiates a new GenericRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GenericRecord) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GenericRecord) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GenericRecord) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GenericRecord) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRecordType

`func (o *GenericRecord) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *GenericRecord) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *GenericRecord) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *GenericRecord) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetRdata

`func (o *GenericRecord) GetRdata() string`

GetRdata returns the Rdata field if non-nil, zero value otherwise.

### GetRdataOk

`func (o *GenericRecord) GetRdataOk() (*string, bool)`

GetRdataOk returns a tuple with the Rdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdata

`func (o *GenericRecord) SetRdata(v string)`

SetRdata sets Rdata field to given value.

### HasRdata

`func (o *GenericRecord) HasRdata() bool`

HasRdata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


