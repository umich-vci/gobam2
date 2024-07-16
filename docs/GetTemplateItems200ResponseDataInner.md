# GetTemplateItems200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**StartOffset** | Pointer to **int64** | The start offset of the IPv4 template item. | [optional] 
**Name** | Pointer to **string** | The name of the IPv4 group. | [optional] 
**RangeDefinitionFormat** | Pointer to **string** | The method that will be used to create the reserved address range of the IPv4 template. | [optional] 
**EndOffset** | Pointer to **int64** | The starting or ending position of the DHCP range. | [optional] 
**Size** | Pointer to **int64** | The number of addresses in the address block. | [optional] 
**SplitAroundStaticAddresses** | Pointer to **bool** | Indicates whether the DHCP range will automatically be split around any static, reserved, and gateway addresses in the network that are within the DHCP range. | [optional] 
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 

## Methods

### NewGetTemplateItems200ResponseDataInner

`func NewGetTemplateItems200ResponseDataInner() *GetTemplateItems200ResponseDataInner`

NewGetTemplateItems200ResponseDataInner instantiates a new GetTemplateItems200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTemplateItems200ResponseDataInnerWithDefaults

`func NewGetTemplateItems200ResponseDataInnerWithDefaults() *GetTemplateItems200ResponseDataInner`

NewGetTemplateItems200ResponseDataInnerWithDefaults instantiates a new GetTemplateItems200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetTemplateItems200ResponseDataInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetTemplateItems200ResponseDataInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetTemplateItems200ResponseDataInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetTemplateItems200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GetTemplateItems200ResponseDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetTemplateItems200ResponseDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetTemplateItems200ResponseDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetTemplateItems200ResponseDataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStartOffset

`func (o *GetTemplateItems200ResponseDataInner) GetStartOffset() int64`

GetStartOffset returns the StartOffset field if non-nil, zero value otherwise.

### GetStartOffsetOk

`func (o *GetTemplateItems200ResponseDataInner) GetStartOffsetOk() (*int64, bool)`

GetStartOffsetOk returns a tuple with the StartOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartOffset

`func (o *GetTemplateItems200ResponseDataInner) SetStartOffset(v int64)`

SetStartOffset sets StartOffset field to given value.

### HasStartOffset

`func (o *GetTemplateItems200ResponseDataInner) HasStartOffset() bool`

HasStartOffset returns a boolean if a field has been set.

### GetName

`func (o *GetTemplateItems200ResponseDataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetTemplateItems200ResponseDataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetTemplateItems200ResponseDataInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetTemplateItems200ResponseDataInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRangeDefinitionFormat

`func (o *GetTemplateItems200ResponseDataInner) GetRangeDefinitionFormat() string`

GetRangeDefinitionFormat returns the RangeDefinitionFormat field if non-nil, zero value otherwise.

### GetRangeDefinitionFormatOk

`func (o *GetTemplateItems200ResponseDataInner) GetRangeDefinitionFormatOk() (*string, bool)`

GetRangeDefinitionFormatOk returns a tuple with the RangeDefinitionFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeDefinitionFormat

`func (o *GetTemplateItems200ResponseDataInner) SetRangeDefinitionFormat(v string)`

SetRangeDefinitionFormat sets RangeDefinitionFormat field to given value.

### HasRangeDefinitionFormat

`func (o *GetTemplateItems200ResponseDataInner) HasRangeDefinitionFormat() bool`

HasRangeDefinitionFormat returns a boolean if a field has been set.

### GetEndOffset

`func (o *GetTemplateItems200ResponseDataInner) GetEndOffset() int64`

GetEndOffset returns the EndOffset field if non-nil, zero value otherwise.

### GetEndOffsetOk

`func (o *GetTemplateItems200ResponseDataInner) GetEndOffsetOk() (*int64, bool)`

GetEndOffsetOk returns a tuple with the EndOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndOffset

`func (o *GetTemplateItems200ResponseDataInner) SetEndOffset(v int64)`

SetEndOffset sets EndOffset field to given value.

### HasEndOffset

`func (o *GetTemplateItems200ResponseDataInner) HasEndOffset() bool`

HasEndOffset returns a boolean if a field has been set.

### GetSize

`func (o *GetTemplateItems200ResponseDataInner) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *GetTemplateItems200ResponseDataInner) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *GetTemplateItems200ResponseDataInner) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *GetTemplateItems200ResponseDataInner) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSplitAroundStaticAddresses

`func (o *GetTemplateItems200ResponseDataInner) GetSplitAroundStaticAddresses() bool`

GetSplitAroundStaticAddresses returns the SplitAroundStaticAddresses field if non-nil, zero value otherwise.

### GetSplitAroundStaticAddressesOk

`func (o *GetTemplateItems200ResponseDataInner) GetSplitAroundStaticAddressesOk() (*bool, bool)`

GetSplitAroundStaticAddressesOk returns a tuple with the SplitAroundStaticAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitAroundStaticAddresses

`func (o *GetTemplateItems200ResponseDataInner) SetSplitAroundStaticAddresses(v bool)`

SetSplitAroundStaticAddresses sets SplitAroundStaticAddresses field to given value.

### HasSplitAroundStaticAddresses

`func (o *GetTemplateItems200ResponseDataInner) HasSplitAroundStaticAddresses() bool`

HasSplitAroundStaticAddresses returns a boolean if a field has been set.

### GetLinks

`func (o *GetTemplateItems200ResponseDataInner) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetTemplateItems200ResponseDataInner) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetTemplateItems200ResponseDataInner) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetTemplateItems200ResponseDataInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


