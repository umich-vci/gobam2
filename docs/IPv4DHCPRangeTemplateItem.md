# IPv4DHCPRangeTemplateItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The resource type. | [optional] 
**RangeDefinitionFormat** | Pointer to **string** | The type of method that will be used to create a DHCP range. | [optional] 
**EndOffset** | Pointer to **int64** | The starting or ending position of the DHCP range. | [optional] 
**Size** | Pointer to **int64** | The size of the DHCP range. When the method is OFFSETS_AND_SIZE, the value is the size of the DHCP range. When the method is OFFSETS_AND_PERCENTAGE, the value is the size of the DHCP range proportional to the parent network size as a percentage. | [optional] 
**SplitAroundStaticAddresses** | Pointer to **bool** | Indicates whether the DHCP range will automatically be split around any static, reserved, and gateway addresses in the network that are within the DHCP range. | [optional] 

## Methods

### NewIPv4DHCPRangeTemplateItem

`func NewIPv4DHCPRangeTemplateItem() *IPv4DHCPRangeTemplateItem`

NewIPv4DHCPRangeTemplateItem instantiates a new IPv4DHCPRangeTemplateItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPv4DHCPRangeTemplateItemWithDefaults

`func NewIPv4DHCPRangeTemplateItemWithDefaults() *IPv4DHCPRangeTemplateItem`

NewIPv4DHCPRangeTemplateItemWithDefaults instantiates a new IPv4DHCPRangeTemplateItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *IPv4DHCPRangeTemplateItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IPv4DHCPRangeTemplateItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IPv4DHCPRangeTemplateItem) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IPv4DHCPRangeTemplateItem) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRangeDefinitionFormat

`func (o *IPv4DHCPRangeTemplateItem) GetRangeDefinitionFormat() string`

GetRangeDefinitionFormat returns the RangeDefinitionFormat field if non-nil, zero value otherwise.

### GetRangeDefinitionFormatOk

`func (o *IPv4DHCPRangeTemplateItem) GetRangeDefinitionFormatOk() (*string, bool)`

GetRangeDefinitionFormatOk returns a tuple with the RangeDefinitionFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeDefinitionFormat

`func (o *IPv4DHCPRangeTemplateItem) SetRangeDefinitionFormat(v string)`

SetRangeDefinitionFormat sets RangeDefinitionFormat field to given value.

### HasRangeDefinitionFormat

`func (o *IPv4DHCPRangeTemplateItem) HasRangeDefinitionFormat() bool`

HasRangeDefinitionFormat returns a boolean if a field has been set.

### GetEndOffset

`func (o *IPv4DHCPRangeTemplateItem) GetEndOffset() int64`

GetEndOffset returns the EndOffset field if non-nil, zero value otherwise.

### GetEndOffsetOk

`func (o *IPv4DHCPRangeTemplateItem) GetEndOffsetOk() (*int64, bool)`

GetEndOffsetOk returns a tuple with the EndOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndOffset

`func (o *IPv4DHCPRangeTemplateItem) SetEndOffset(v int64)`

SetEndOffset sets EndOffset field to given value.

### HasEndOffset

`func (o *IPv4DHCPRangeTemplateItem) HasEndOffset() bool`

HasEndOffset returns a boolean if a field has been set.

### GetSize

`func (o *IPv4DHCPRangeTemplateItem) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *IPv4DHCPRangeTemplateItem) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *IPv4DHCPRangeTemplateItem) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *IPv4DHCPRangeTemplateItem) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSplitAroundStaticAddresses

`func (o *IPv4DHCPRangeTemplateItem) GetSplitAroundStaticAddresses() bool`

GetSplitAroundStaticAddresses returns the SplitAroundStaticAddresses field if non-nil, zero value otherwise.

### GetSplitAroundStaticAddressesOk

`func (o *IPv4DHCPRangeTemplateItem) GetSplitAroundStaticAddressesOk() (*bool, bool)`

GetSplitAroundStaticAddressesOk returns a tuple with the SplitAroundStaticAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitAroundStaticAddresses

`func (o *IPv4DHCPRangeTemplateItem) SetSplitAroundStaticAddresses(v bool)`

SetSplitAroundStaticAddresses sets SplitAroundStaticAddresses field to given value.

### HasSplitAroundStaticAddresses

`func (o *IPv4DHCPRangeTemplateItem) HasSplitAroundStaticAddresses() bool`

HasSplitAroundStaticAddresses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


