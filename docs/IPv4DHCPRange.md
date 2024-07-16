# IPv4DHCPRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The resource type. | [optional] 
**Template** | Pointer to [**InlinedIPv4Template**](InlinedIPv4Template.md) |  | [optional] 
**SplitAroundStaticAddresses** | Pointer to **bool** | Indicates whether the DHCP range will automatically be split around any static, reserved, and gateway addresses in the network that are within the DHCP range. | [optional] 
**ExclusionRanges** | Pointer to **[]string** | The exclusion ranges for the DHCP range. Addresses in each exclusion range will not be dynamically assigned. | [optional] 
**LowWaterMark** | Pointer to **int32** | A DHCP alert is triggered when usage falls below this percentage (when too few addresses are in use). | [optional] 
**HighWaterMark** | Pointer to **int32** | A DHCP alert is triggered when usage exceeds this percentage (when too many addresses are in use). | [optional] 

## Methods

### NewIPv4DHCPRange

`func NewIPv4DHCPRange() *IPv4DHCPRange`

NewIPv4DHCPRange instantiates a new IPv4DHCPRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPv4DHCPRangeWithDefaults

`func NewIPv4DHCPRangeWithDefaults() *IPv4DHCPRange`

NewIPv4DHCPRangeWithDefaults instantiates a new IPv4DHCPRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *IPv4DHCPRange) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IPv4DHCPRange) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IPv4DHCPRange) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IPv4DHCPRange) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTemplate

`func (o *IPv4DHCPRange) GetTemplate() InlinedIPv4Template`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *IPv4DHCPRange) GetTemplateOk() (*InlinedIPv4Template, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *IPv4DHCPRange) SetTemplate(v InlinedIPv4Template)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *IPv4DHCPRange) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetSplitAroundStaticAddresses

`func (o *IPv4DHCPRange) GetSplitAroundStaticAddresses() bool`

GetSplitAroundStaticAddresses returns the SplitAroundStaticAddresses field if non-nil, zero value otherwise.

### GetSplitAroundStaticAddressesOk

`func (o *IPv4DHCPRange) GetSplitAroundStaticAddressesOk() (*bool, bool)`

GetSplitAroundStaticAddressesOk returns a tuple with the SplitAroundStaticAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitAroundStaticAddresses

`func (o *IPv4DHCPRange) SetSplitAroundStaticAddresses(v bool)`

SetSplitAroundStaticAddresses sets SplitAroundStaticAddresses field to given value.

### HasSplitAroundStaticAddresses

`func (o *IPv4DHCPRange) HasSplitAroundStaticAddresses() bool`

HasSplitAroundStaticAddresses returns a boolean if a field has been set.

### GetExclusionRanges

`func (o *IPv4DHCPRange) GetExclusionRanges() []string`

GetExclusionRanges returns the ExclusionRanges field if non-nil, zero value otherwise.

### GetExclusionRangesOk

`func (o *IPv4DHCPRange) GetExclusionRangesOk() (*[]string, bool)`

GetExclusionRangesOk returns a tuple with the ExclusionRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusionRanges

`func (o *IPv4DHCPRange) SetExclusionRanges(v []string)`

SetExclusionRanges sets ExclusionRanges field to given value.

### HasExclusionRanges

`func (o *IPv4DHCPRange) HasExclusionRanges() bool`

HasExclusionRanges returns a boolean if a field has been set.

### GetLowWaterMark

`func (o *IPv4DHCPRange) GetLowWaterMark() int32`

GetLowWaterMark returns the LowWaterMark field if non-nil, zero value otherwise.

### GetLowWaterMarkOk

`func (o *IPv4DHCPRange) GetLowWaterMarkOk() (*int32, bool)`

GetLowWaterMarkOk returns a tuple with the LowWaterMark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowWaterMark

`func (o *IPv4DHCPRange) SetLowWaterMark(v int32)`

SetLowWaterMark sets LowWaterMark field to given value.

### HasLowWaterMark

`func (o *IPv4DHCPRange) HasLowWaterMark() bool`

HasLowWaterMark returns a boolean if a field has been set.

### GetHighWaterMark

`func (o *IPv4DHCPRange) GetHighWaterMark() int32`

GetHighWaterMark returns the HighWaterMark field if non-nil, zero value otherwise.

### GetHighWaterMarkOk

`func (o *IPv4DHCPRange) GetHighWaterMarkOk() (*int32, bool)`

GetHighWaterMarkOk returns a tuple with the HighWaterMark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighWaterMark

`func (o *IPv4DHCPRange) SetHighWaterMark(v int32)`

SetHighWaterMark sets HighWaterMark field to given value.

### HasHighWaterMark

`func (o *IPv4DHCPRange) HasHighWaterMark() bool`

HasHighWaterMark returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


