# PostTemplateTemplateItem201Response

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

### NewPostTemplateTemplateItem201Response

`func NewPostTemplateTemplateItem201Response() *PostTemplateTemplateItem201Response`

NewPostTemplateTemplateItem201Response instantiates a new PostTemplateTemplateItem201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostTemplateTemplateItem201ResponseWithDefaults

`func NewPostTemplateTemplateItem201ResponseWithDefaults() *PostTemplateTemplateItem201Response`

NewPostTemplateTemplateItem201ResponseWithDefaults instantiates a new PostTemplateTemplateItem201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PostTemplateTemplateItem201Response) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostTemplateTemplateItem201Response) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostTemplateTemplateItem201Response) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *PostTemplateTemplateItem201Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *PostTemplateTemplateItem201Response) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PostTemplateTemplateItem201Response) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PostTemplateTemplateItem201Response) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PostTemplateTemplateItem201Response) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStartOffset

`func (o *PostTemplateTemplateItem201Response) GetStartOffset() int64`

GetStartOffset returns the StartOffset field if non-nil, zero value otherwise.

### GetStartOffsetOk

`func (o *PostTemplateTemplateItem201Response) GetStartOffsetOk() (*int64, bool)`

GetStartOffsetOk returns a tuple with the StartOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartOffset

`func (o *PostTemplateTemplateItem201Response) SetStartOffset(v int64)`

SetStartOffset sets StartOffset field to given value.

### HasStartOffset

`func (o *PostTemplateTemplateItem201Response) HasStartOffset() bool`

HasStartOffset returns a boolean if a field has been set.

### GetName

`func (o *PostTemplateTemplateItem201Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostTemplateTemplateItem201Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostTemplateTemplateItem201Response) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PostTemplateTemplateItem201Response) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRangeDefinitionFormat

`func (o *PostTemplateTemplateItem201Response) GetRangeDefinitionFormat() string`

GetRangeDefinitionFormat returns the RangeDefinitionFormat field if non-nil, zero value otherwise.

### GetRangeDefinitionFormatOk

`func (o *PostTemplateTemplateItem201Response) GetRangeDefinitionFormatOk() (*string, bool)`

GetRangeDefinitionFormatOk returns a tuple with the RangeDefinitionFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeDefinitionFormat

`func (o *PostTemplateTemplateItem201Response) SetRangeDefinitionFormat(v string)`

SetRangeDefinitionFormat sets RangeDefinitionFormat field to given value.

### HasRangeDefinitionFormat

`func (o *PostTemplateTemplateItem201Response) HasRangeDefinitionFormat() bool`

HasRangeDefinitionFormat returns a boolean if a field has been set.

### GetEndOffset

`func (o *PostTemplateTemplateItem201Response) GetEndOffset() int64`

GetEndOffset returns the EndOffset field if non-nil, zero value otherwise.

### GetEndOffsetOk

`func (o *PostTemplateTemplateItem201Response) GetEndOffsetOk() (*int64, bool)`

GetEndOffsetOk returns a tuple with the EndOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndOffset

`func (o *PostTemplateTemplateItem201Response) SetEndOffset(v int64)`

SetEndOffset sets EndOffset field to given value.

### HasEndOffset

`func (o *PostTemplateTemplateItem201Response) HasEndOffset() bool`

HasEndOffset returns a boolean if a field has been set.

### GetSize

`func (o *PostTemplateTemplateItem201Response) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PostTemplateTemplateItem201Response) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PostTemplateTemplateItem201Response) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *PostTemplateTemplateItem201Response) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSplitAroundStaticAddresses

`func (o *PostTemplateTemplateItem201Response) GetSplitAroundStaticAddresses() bool`

GetSplitAroundStaticAddresses returns the SplitAroundStaticAddresses field if non-nil, zero value otherwise.

### GetSplitAroundStaticAddressesOk

`func (o *PostTemplateTemplateItem201Response) GetSplitAroundStaticAddressesOk() (*bool, bool)`

GetSplitAroundStaticAddressesOk returns a tuple with the SplitAroundStaticAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitAroundStaticAddresses

`func (o *PostTemplateTemplateItem201Response) SetSplitAroundStaticAddresses(v bool)`

SetSplitAroundStaticAddresses sets SplitAroundStaticAddresses field to given value.

### HasSplitAroundStaticAddresses

`func (o *PostTemplateTemplateItem201Response) HasSplitAroundStaticAddresses() bool`

HasSplitAroundStaticAddresses returns a boolean if a field has been set.

### GetLinks

`func (o *PostTemplateTemplateItem201Response) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PostTemplateTemplateItem201Response) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PostTemplateTemplateItem201Response) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PostTemplateTemplateItem201Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


