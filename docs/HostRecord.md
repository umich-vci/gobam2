# HostRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The resource type. | [optional] 
**Addresses** | Pointer to [**[]LinkedAddress**](LinkedAddress.md) |  | [optional] 
**ReverseRecord** | Pointer to **bool** | Indicates if a PTR record is created for the host record. | [optional] 
**Views** | Pointer to [**[]HostRecordAllOfViews**](HostRecordAllOfViews.md) |  | [optional] 

## Methods

### NewHostRecord

`func NewHostRecord() *HostRecord`

NewHostRecord instantiates a new HostRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostRecordWithDefaults

`func NewHostRecordWithDefaults() *HostRecord`

NewHostRecordWithDefaults instantiates a new HostRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *HostRecord) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HostRecord) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HostRecord) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *HostRecord) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAddresses

`func (o *HostRecord) GetAddresses() []LinkedAddress`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *HostRecord) GetAddressesOk() (*[]LinkedAddress, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *HostRecord) SetAddresses(v []LinkedAddress)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *HostRecord) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetReverseRecord

`func (o *HostRecord) GetReverseRecord() bool`

GetReverseRecord returns the ReverseRecord field if non-nil, zero value otherwise.

### GetReverseRecordOk

`func (o *HostRecord) GetReverseRecordOk() (*bool, bool)`

GetReverseRecordOk returns a tuple with the ReverseRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReverseRecord

`func (o *HostRecord) SetReverseRecord(v bool)`

SetReverseRecord sets ReverseRecord field to given value.

### HasReverseRecord

`func (o *HostRecord) HasReverseRecord() bool`

HasReverseRecord returns a boolean if a field has been set.

### GetViews

`func (o *HostRecord) GetViews() []HostRecordAllOfViews`

GetViews returns the Views field if non-nil, zero value otherwise.

### GetViewsOk

`func (o *HostRecord) GetViewsOk() (*[]HostRecordAllOfViews, bool)`

GetViewsOk returns a tuple with the Views field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViews

`func (o *HostRecord) SetViews(v []HostRecordAllOfViews)`

SetViews sets Views field to given value.

### HasViews

`func (o *HostRecord) HasViews() bool`

HasViews returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


