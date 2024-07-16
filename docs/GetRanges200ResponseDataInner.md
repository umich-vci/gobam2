# GetRanges200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | Pointer to **string** | The name of the resource. | [optional] 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**Configuration** | Pointer to [**InlinedConfiguration**](InlinedConfiguration.md) |  | [optional] [readonly] 
**Range** | Pointer to **string** | The address range for the DHCP range. | [optional] 
**Usage** | Pointer to **map[string]int64** | Values indicating the current amount of DHCP allocated, DHCP abandoned, DHCP reserved, and unassigned IP addresses within the range. | [optional] [readonly] 
**InheritedFields** | Pointer to **[]string** |  | [optional] [readonly] 
**Template** | Pointer to [**InlinedIPv4Template**](InlinedIPv4Template.md) |  | [optional] 
**SplitAroundStaticAddresses** | Pointer to **bool** | Indicates whether the DHCP range will automatically be split around any static, reserved, and gateway addresses in the network that are within the DHCP range. | [optional] 
**ExclusionRanges** | Pointer to **[]string** | The exclusion ranges for the DHCP range. Addresses in each exclusion range will not be dynamically assigned. | [optional] 
**LowWaterMark** | Pointer to **int32** | A DHCP alert is triggered when usage falls below this percentage (when too few addresses are in use). | [optional] 
**HighWaterMark** | Pointer to **int32** | A DHCP alert is triggered when usage exceeds this percentage (when too many addresses are in use). | [optional] 
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 

## Methods

### NewGetRanges200ResponseDataInner

`func NewGetRanges200ResponseDataInner() *GetRanges200ResponseDataInner`

NewGetRanges200ResponseDataInner instantiates a new GetRanges200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRanges200ResponseDataInnerWithDefaults

`func NewGetRanges200ResponseDataInnerWithDefaults() *GetRanges200ResponseDataInner`

NewGetRanges200ResponseDataInnerWithDefaults instantiates a new GetRanges200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetRanges200ResponseDataInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetRanges200ResponseDataInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetRanges200ResponseDataInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetRanges200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GetRanges200ResponseDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetRanges200ResponseDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetRanges200ResponseDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetRanges200ResponseDataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *GetRanges200ResponseDataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetRanges200ResponseDataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetRanges200ResponseDataInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetRanges200ResponseDataInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *GetRanges200ResponseDataInner) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *GetRanges200ResponseDataInner) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *GetRanges200ResponseDataInner) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *GetRanges200ResponseDataInner) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetConfiguration

`func (o *GetRanges200ResponseDataInner) GetConfiguration() InlinedConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *GetRanges200ResponseDataInner) GetConfigurationOk() (*InlinedConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *GetRanges200ResponseDataInner) SetConfiguration(v InlinedConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *GetRanges200ResponseDataInner) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetRange

`func (o *GetRanges200ResponseDataInner) GetRange() string`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *GetRanges200ResponseDataInner) GetRangeOk() (*string, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *GetRanges200ResponseDataInner) SetRange(v string)`

SetRange sets Range field to given value.

### HasRange

`func (o *GetRanges200ResponseDataInner) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetUsage

`func (o *GetRanges200ResponseDataInner) GetUsage() map[string]int64`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *GetRanges200ResponseDataInner) GetUsageOk() (*map[string]int64, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *GetRanges200ResponseDataInner) SetUsage(v map[string]int64)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *GetRanges200ResponseDataInner) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetInheritedFields

`func (o *GetRanges200ResponseDataInner) GetInheritedFields() []string`

GetInheritedFields returns the InheritedFields field if non-nil, zero value otherwise.

### GetInheritedFieldsOk

`func (o *GetRanges200ResponseDataInner) GetInheritedFieldsOk() (*[]string, bool)`

GetInheritedFieldsOk returns a tuple with the InheritedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritedFields

`func (o *GetRanges200ResponseDataInner) SetInheritedFields(v []string)`

SetInheritedFields sets InheritedFields field to given value.

### HasInheritedFields

`func (o *GetRanges200ResponseDataInner) HasInheritedFields() bool`

HasInheritedFields returns a boolean if a field has been set.

### GetTemplate

`func (o *GetRanges200ResponseDataInner) GetTemplate() InlinedIPv4Template`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *GetRanges200ResponseDataInner) GetTemplateOk() (*InlinedIPv4Template, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *GetRanges200ResponseDataInner) SetTemplate(v InlinedIPv4Template)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *GetRanges200ResponseDataInner) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetSplitAroundStaticAddresses

`func (o *GetRanges200ResponseDataInner) GetSplitAroundStaticAddresses() bool`

GetSplitAroundStaticAddresses returns the SplitAroundStaticAddresses field if non-nil, zero value otherwise.

### GetSplitAroundStaticAddressesOk

`func (o *GetRanges200ResponseDataInner) GetSplitAroundStaticAddressesOk() (*bool, bool)`

GetSplitAroundStaticAddressesOk returns a tuple with the SplitAroundStaticAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitAroundStaticAddresses

`func (o *GetRanges200ResponseDataInner) SetSplitAroundStaticAddresses(v bool)`

SetSplitAroundStaticAddresses sets SplitAroundStaticAddresses field to given value.

### HasSplitAroundStaticAddresses

`func (o *GetRanges200ResponseDataInner) HasSplitAroundStaticAddresses() bool`

HasSplitAroundStaticAddresses returns a boolean if a field has been set.

### GetExclusionRanges

`func (o *GetRanges200ResponseDataInner) GetExclusionRanges() []string`

GetExclusionRanges returns the ExclusionRanges field if non-nil, zero value otherwise.

### GetExclusionRangesOk

`func (o *GetRanges200ResponseDataInner) GetExclusionRangesOk() (*[]string, bool)`

GetExclusionRangesOk returns a tuple with the ExclusionRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusionRanges

`func (o *GetRanges200ResponseDataInner) SetExclusionRanges(v []string)`

SetExclusionRanges sets ExclusionRanges field to given value.

### HasExclusionRanges

`func (o *GetRanges200ResponseDataInner) HasExclusionRanges() bool`

HasExclusionRanges returns a boolean if a field has been set.

### GetLowWaterMark

`func (o *GetRanges200ResponseDataInner) GetLowWaterMark() int32`

GetLowWaterMark returns the LowWaterMark field if non-nil, zero value otherwise.

### GetLowWaterMarkOk

`func (o *GetRanges200ResponseDataInner) GetLowWaterMarkOk() (*int32, bool)`

GetLowWaterMarkOk returns a tuple with the LowWaterMark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowWaterMark

`func (o *GetRanges200ResponseDataInner) SetLowWaterMark(v int32)`

SetLowWaterMark sets LowWaterMark field to given value.

### HasLowWaterMark

`func (o *GetRanges200ResponseDataInner) HasLowWaterMark() bool`

HasLowWaterMark returns a boolean if a field has been set.

### GetHighWaterMark

`func (o *GetRanges200ResponseDataInner) GetHighWaterMark() int32`

GetHighWaterMark returns the HighWaterMark field if non-nil, zero value otherwise.

### GetHighWaterMarkOk

`func (o *GetRanges200ResponseDataInner) GetHighWaterMarkOk() (*int32, bool)`

GetHighWaterMarkOk returns a tuple with the HighWaterMark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighWaterMark

`func (o *GetRanges200ResponseDataInner) SetHighWaterMark(v int32)`

SetHighWaterMark sets HighWaterMark field to given value.

### HasHighWaterMark

`func (o *GetRanges200ResponseDataInner) HasHighWaterMark() bool`

HasHighWaterMark returns a boolean if a field has been set.

### GetLinks

`func (o *GetRanges200ResponseDataInner) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetRanges200ResponseDataInner) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetRanges200ResponseDataInner) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetRanges200ResponseDataInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


