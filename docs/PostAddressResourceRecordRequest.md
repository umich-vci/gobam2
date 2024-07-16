# PostAddressResourceRecordRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The resource identifier. | 
**Type** | Pointer to **string** | The resource type. | [optional] 
**AbsoluteName** | Pointer to **string** | The fully qualified domain name of the resource record. | [optional] 
**View** | Pointer to [**InlinedView**](InlinedView.md) |  | [optional] 

## Methods

### NewPostAddressResourceRecordRequest

`func NewPostAddressResourceRecordRequest(id int64, ) *PostAddressResourceRecordRequest`

NewPostAddressResourceRecordRequest instantiates a new PostAddressResourceRecordRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostAddressResourceRecordRequestWithDefaults

`func NewPostAddressResourceRecordRequestWithDefaults() *PostAddressResourceRecordRequest`

NewPostAddressResourceRecordRequestWithDefaults instantiates a new PostAddressResourceRecordRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PostAddressResourceRecordRequest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostAddressResourceRecordRequest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostAddressResourceRecordRequest) SetId(v int64)`

SetId sets Id field to given value.


### GetType

`func (o *PostAddressResourceRecordRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PostAddressResourceRecordRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PostAddressResourceRecordRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PostAddressResourceRecordRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAbsoluteName

`func (o *PostAddressResourceRecordRequest) GetAbsoluteName() string`

GetAbsoluteName returns the AbsoluteName field if non-nil, zero value otherwise.

### GetAbsoluteNameOk

`func (o *PostAddressResourceRecordRequest) GetAbsoluteNameOk() (*string, bool)`

GetAbsoluteNameOk returns a tuple with the AbsoluteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsoluteName

`func (o *PostAddressResourceRecordRequest) SetAbsoluteName(v string)`

SetAbsoluteName sets AbsoluteName field to given value.

### HasAbsoluteName

`func (o *PostAddressResourceRecordRequest) HasAbsoluteName() bool`

HasAbsoluteName returns a boolean if a field has been set.

### GetView

`func (o *PostAddressResourceRecordRequest) GetView() InlinedView`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *PostAddressResourceRecordRequest) GetViewOk() (*InlinedView, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *PostAddressResourceRecordRequest) SetView(v InlinedView)`

SetView sets View field to given value.

### HasView

`func (o *PostAddressResourceRecordRequest) HasView() bool`

HasView returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


