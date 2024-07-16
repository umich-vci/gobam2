# LinkedHostRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The resource type. | [optional] 
**ReverseRecord** | Pointer to **bool** |  | [optional] 
**Views** | Pointer to [**[]LinkedHostRecordAllOfViews**](LinkedHostRecordAllOfViews.md) |  | [optional] 

## Methods

### NewLinkedHostRecord

`func NewLinkedHostRecord() *LinkedHostRecord`

NewLinkedHostRecord instantiates a new LinkedHostRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkedHostRecordWithDefaults

`func NewLinkedHostRecordWithDefaults() *LinkedHostRecord`

NewLinkedHostRecordWithDefaults instantiates a new LinkedHostRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *LinkedHostRecord) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LinkedHostRecord) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LinkedHostRecord) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *LinkedHostRecord) HasType() bool`

HasType returns a boolean if a field has been set.

### GetReverseRecord

`func (o *LinkedHostRecord) GetReverseRecord() bool`

GetReverseRecord returns the ReverseRecord field if non-nil, zero value otherwise.

### GetReverseRecordOk

`func (o *LinkedHostRecord) GetReverseRecordOk() (*bool, bool)`

GetReverseRecordOk returns a tuple with the ReverseRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReverseRecord

`func (o *LinkedHostRecord) SetReverseRecord(v bool)`

SetReverseRecord sets ReverseRecord field to given value.

### HasReverseRecord

`func (o *LinkedHostRecord) HasReverseRecord() bool`

HasReverseRecord returns a boolean if a field has been set.

### GetViews

`func (o *LinkedHostRecord) GetViews() []LinkedHostRecordAllOfViews`

GetViews returns the Views field if non-nil, zero value otherwise.

### GetViewsOk

`func (o *LinkedHostRecord) GetViewsOk() (*[]LinkedHostRecordAllOfViews, bool)`

GetViewsOk returns a tuple with the Views field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViews

`func (o *LinkedHostRecord) SetViews(v []LinkedHostRecordAllOfViews)`

SetViews sets Views field to given value.

### HasViews

`func (o *LinkedHostRecord) HasViews() bool`

HasViews returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


