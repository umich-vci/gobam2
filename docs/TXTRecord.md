# TXTRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The resource type. | [optional] 
**Text** | Pointer to **string** | The text data within the TXT resource record. | [optional] 

## Methods

### NewTXTRecord

`func NewTXTRecord() *TXTRecord`

NewTXTRecord instantiates a new TXTRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTXTRecordWithDefaults

`func NewTXTRecordWithDefaults() *TXTRecord`

NewTXTRecordWithDefaults instantiates a new TXTRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TXTRecord) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TXTRecord) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TXTRecord) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TXTRecord) HasType() bool`

HasType returns a boolean if a field has been set.

### GetText

`func (o *TXTRecord) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *TXTRecord) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *TXTRecord) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *TXTRecord) HasText() bool`

HasText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


