# IPv4ReservedRangeTemplateItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The resource type. | [optional] 
**RangeDefinitionFormat** | Pointer to **string** | The method that will be used to create the reserved address range of the IPv4 template. | [optional] 
**Size** | Pointer to **int64** | The number of addresses in the address block. | [optional] 

## Methods

### NewIPv4ReservedRangeTemplateItem

`func NewIPv4ReservedRangeTemplateItem() *IPv4ReservedRangeTemplateItem`

NewIPv4ReservedRangeTemplateItem instantiates a new IPv4ReservedRangeTemplateItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPv4ReservedRangeTemplateItemWithDefaults

`func NewIPv4ReservedRangeTemplateItemWithDefaults() *IPv4ReservedRangeTemplateItem`

NewIPv4ReservedRangeTemplateItemWithDefaults instantiates a new IPv4ReservedRangeTemplateItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *IPv4ReservedRangeTemplateItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IPv4ReservedRangeTemplateItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IPv4ReservedRangeTemplateItem) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IPv4ReservedRangeTemplateItem) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRangeDefinitionFormat

`func (o *IPv4ReservedRangeTemplateItem) GetRangeDefinitionFormat() string`

GetRangeDefinitionFormat returns the RangeDefinitionFormat field if non-nil, zero value otherwise.

### GetRangeDefinitionFormatOk

`func (o *IPv4ReservedRangeTemplateItem) GetRangeDefinitionFormatOk() (*string, bool)`

GetRangeDefinitionFormatOk returns a tuple with the RangeDefinitionFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeDefinitionFormat

`func (o *IPv4ReservedRangeTemplateItem) SetRangeDefinitionFormat(v string)`

SetRangeDefinitionFormat sets RangeDefinitionFormat field to given value.

### HasRangeDefinitionFormat

`func (o *IPv4ReservedRangeTemplateItem) HasRangeDefinitionFormat() bool`

HasRangeDefinitionFormat returns a boolean if a field has been set.

### GetSize

`func (o *IPv4ReservedRangeTemplateItem) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *IPv4ReservedRangeTemplateItem) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *IPv4ReservedRangeTemplateItem) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *IPv4ReservedRangeTemplateItem) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


