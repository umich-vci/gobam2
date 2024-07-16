# PostTemplateTemplateItemRequest

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

### NewPostTemplateTemplateItemRequest

`func NewPostTemplateTemplateItemRequest(type_ string, startOffset int64, name string, rangeDefinitionFormat string, size int64, ) *PostTemplateTemplateItemRequest`

NewPostTemplateTemplateItemRequest instantiates a new PostTemplateTemplateItemRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostTemplateTemplateItemRequestWithDefaults

`func NewPostTemplateTemplateItemRequestWithDefaults() *PostTemplateTemplateItemRequest`

NewPostTemplateTemplateItemRequestWithDefaults instantiates a new PostTemplateTemplateItemRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PostTemplateTemplateItemRequest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostTemplateTemplateItemRequest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostTemplateTemplateItemRequest) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *PostTemplateTemplateItemRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *PostTemplateTemplateItemRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PostTemplateTemplateItemRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PostTemplateTemplateItemRequest) SetType(v string)`

SetType sets Type field to given value.


### GetStartOffset

`func (o *PostTemplateTemplateItemRequest) GetStartOffset() int64`

GetStartOffset returns the StartOffset field if non-nil, zero value otherwise.

### GetStartOffsetOk

`func (o *PostTemplateTemplateItemRequest) GetStartOffsetOk() (*int64, bool)`

GetStartOffsetOk returns a tuple with the StartOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartOffset

`func (o *PostTemplateTemplateItemRequest) SetStartOffset(v int64)`

SetStartOffset sets StartOffset field to given value.


### GetName

`func (o *PostTemplateTemplateItemRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostTemplateTemplateItemRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostTemplateTemplateItemRequest) SetName(v string)`

SetName sets Name field to given value.


### GetRangeDefinitionFormat

`func (o *PostTemplateTemplateItemRequest) GetRangeDefinitionFormat() string`

GetRangeDefinitionFormat returns the RangeDefinitionFormat field if non-nil, zero value otherwise.

### GetRangeDefinitionFormatOk

`func (o *PostTemplateTemplateItemRequest) GetRangeDefinitionFormatOk() (*string, bool)`

GetRangeDefinitionFormatOk returns a tuple with the RangeDefinitionFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeDefinitionFormat

`func (o *PostTemplateTemplateItemRequest) SetRangeDefinitionFormat(v string)`

SetRangeDefinitionFormat sets RangeDefinitionFormat field to given value.


### GetSize

`func (o *PostTemplateTemplateItemRequest) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PostTemplateTemplateItemRequest) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PostTemplateTemplateItemRequest) SetSize(v int64)`

SetSize sets Size field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


