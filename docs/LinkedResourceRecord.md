# LinkedResourceRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**AbsoluteName** | Pointer to **string** | The fully qualified domain name of the resource record. | [optional] 

## Methods

### NewLinkedResourceRecord

`func NewLinkedResourceRecord() *LinkedResourceRecord`

NewLinkedResourceRecord instantiates a new LinkedResourceRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkedResourceRecordWithDefaults

`func NewLinkedResourceRecordWithDefaults() *LinkedResourceRecord`

NewLinkedResourceRecordWithDefaults instantiates a new LinkedResourceRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LinkedResourceRecord) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LinkedResourceRecord) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LinkedResourceRecord) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *LinkedResourceRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *LinkedResourceRecord) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LinkedResourceRecord) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LinkedResourceRecord) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *LinkedResourceRecord) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAbsoluteName

`func (o *LinkedResourceRecord) GetAbsoluteName() string`

GetAbsoluteName returns the AbsoluteName field if non-nil, zero value otherwise.

### GetAbsoluteNameOk

`func (o *LinkedResourceRecord) GetAbsoluteNameOk() (*string, bool)`

GetAbsoluteNameOk returns a tuple with the AbsoluteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsoluteName

`func (o *LinkedResourceRecord) SetAbsoluteName(v string)`

SetAbsoluteName sets AbsoluteName field to given value.

### HasAbsoluteName

`func (o *LinkedResourceRecord) HasAbsoluteName() bool`

HasAbsoluteName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


