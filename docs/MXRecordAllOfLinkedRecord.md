# MXRecordAllOfLinkedRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The resource identifier. | 
**Type** | Pointer to **string** | The resource type. | [optional] 
**AbsoluteName** | **string** | The fully qualified domain name of the resource record. | 

## Methods

### NewMXRecordAllOfLinkedRecord

`func NewMXRecordAllOfLinkedRecord(id int64, absoluteName string, ) *MXRecordAllOfLinkedRecord`

NewMXRecordAllOfLinkedRecord instantiates a new MXRecordAllOfLinkedRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMXRecordAllOfLinkedRecordWithDefaults

`func NewMXRecordAllOfLinkedRecordWithDefaults() *MXRecordAllOfLinkedRecord`

NewMXRecordAllOfLinkedRecordWithDefaults instantiates a new MXRecordAllOfLinkedRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MXRecordAllOfLinkedRecord) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MXRecordAllOfLinkedRecord) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MXRecordAllOfLinkedRecord) SetId(v int64)`

SetId sets Id field to given value.


### GetType

`func (o *MXRecordAllOfLinkedRecord) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MXRecordAllOfLinkedRecord) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MXRecordAllOfLinkedRecord) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MXRecordAllOfLinkedRecord) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAbsoluteName

`func (o *MXRecordAllOfLinkedRecord) GetAbsoluteName() string`

GetAbsoluteName returns the AbsoluteName field if non-nil, zero value otherwise.

### GetAbsoluteNameOk

`func (o *MXRecordAllOfLinkedRecord) GetAbsoluteNameOk() (*string, bool)`

GetAbsoluteNameOk returns a tuple with the AbsoluteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsoluteName

`func (o *MXRecordAllOfLinkedRecord) SetAbsoluteName(v string)`

SetAbsoluteName sets AbsoluteName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


