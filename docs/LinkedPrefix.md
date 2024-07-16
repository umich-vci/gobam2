# LinkedPrefix

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Range** | Pointer to **string** | The address range for the IP block or network. | [optional] 

## Methods

### NewLinkedPrefix

`func NewLinkedPrefix() *LinkedPrefix`

NewLinkedPrefix instantiates a new LinkedPrefix object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkedPrefixWithDefaults

`func NewLinkedPrefixWithDefaults() *LinkedPrefix`

NewLinkedPrefixWithDefaults instantiates a new LinkedPrefix object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LinkedPrefix) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LinkedPrefix) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LinkedPrefix) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *LinkedPrefix) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *LinkedPrefix) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LinkedPrefix) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LinkedPrefix) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *LinkedPrefix) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRange

`func (o *LinkedPrefix) GetRange() string`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *LinkedPrefix) GetRangeOk() (*string, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *LinkedPrefix) SetRange(v string)`

SetRange sets Range field to given value.

### HasRange

`func (o *LinkedPrefix) HasRange() bool`

HasRange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


