# PutTemplateItemRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | **string** | The resource type. | 
**StartOffset** | **int64** | The start offset of the IPv4 template item. | 
**Name** | **string** | The name of the template item. | 
**RangeDefinitionFormat** | **string** | The method that will be used to create the reserved address range of the IPv4 template. | 
**Size** | **int64** | The number of addresses in the address block. | 

## Methods

### NewPutTemplateItemRequest

`func NewPutTemplateItemRequest(type_ string, startOffset int64, name string, rangeDefinitionFormat string, size int64, ) *PutTemplateItemRequest`

NewPutTemplateItemRequest instantiates a new PutTemplateItemRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutTemplateItemRequestWithDefaults

`func NewPutTemplateItemRequestWithDefaults() *PutTemplateItemRequest`

NewPutTemplateItemRequestWithDefaults instantiates a new PutTemplateItemRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PutTemplateItemRequest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PutTemplateItemRequest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PutTemplateItemRequest) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *PutTemplateItemRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *PutTemplateItemRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PutTemplateItemRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PutTemplateItemRequest) SetType(v string)`

SetType sets Type field to given value.


### GetStartOffset

`func (o *PutTemplateItemRequest) GetStartOffset() int64`

GetStartOffset returns the StartOffset field if non-nil, zero value otherwise.

### GetStartOffsetOk

`func (o *PutTemplateItemRequest) GetStartOffsetOk() (*int64, bool)`

GetStartOffsetOk returns a tuple with the StartOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartOffset

`func (o *PutTemplateItemRequest) SetStartOffset(v int64)`

SetStartOffset sets StartOffset field to given value.


### GetName

`func (o *PutTemplateItemRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PutTemplateItemRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PutTemplateItemRequest) SetName(v string)`

SetName sets Name field to given value.


### GetRangeDefinitionFormat

`func (o *PutTemplateItemRequest) GetRangeDefinitionFormat() string`

GetRangeDefinitionFormat returns the RangeDefinitionFormat field if non-nil, zero value otherwise.

### GetRangeDefinitionFormatOk

`func (o *PutTemplateItemRequest) GetRangeDefinitionFormatOk() (*string, bool)`

GetRangeDefinitionFormatOk returns a tuple with the RangeDefinitionFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeDefinitionFormat

`func (o *PutTemplateItemRequest) SetRangeDefinitionFormat(v string)`

SetRangeDefinitionFormat sets RangeDefinitionFormat field to given value.


### GetSize

`func (o *PutTemplateItemRequest) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PutTemplateItemRequest) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PutTemplateItemRequest) SetSize(v int64)`

SetSize sets Size field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


