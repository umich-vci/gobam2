# GetResourceRecords200Response1DataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | Pointer to **string** | The name of the resource. | [optional] 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**Configuration** | Pointer to [**InlinedConfiguration**](InlinedConfiguration.md) |  | [optional] [readonly] 
**Ttl** | Pointer to **int64** | The overridden TTL value of the resource record, in seconds. | [optional] 
**AbsoluteName** | Pointer to **string** | The fully qualified domain name of the resource record. | [optional] 
**Comment** | Pointer to **string** | The additional comments associated with the resource record. | [optional] 
**Dynamic** | Pointer to **bool** | Indicates whether the resource record is dynamic. | [optional] [readonly] 
**View** | Pointer to [**InlinedView**](InlinedView.md) |  | [optional] 
**RecordType** | Pointer to **string** | The resource record type. | [optional] 
**Rdata** | Pointer to **string** | The resource record data. | [optional] 
**Os** | Pointer to **string** | The string describing the server operating system. | [optional] 
**Cpu** | Pointer to **string** | The string describing the server CPU. | [optional] 
**Addresses** | Pointer to [**[]LinkedAddress**](LinkedAddress.md) |  | [optional] 
**ReverseRecord** | Pointer to **bool** | Indicates if a PTR record is created for the host record. | [optional] 
**Views** | Pointer to [**[]HostRecordAllOfViews**](HostRecordAllOfViews.md) |  | [optional] 
**LinkedRecord** | Pointer to [**SRVRecordAllOfLinkedRecord**](SRVRecordAllOfLinkedRecord.md) |  | [optional] 
**Priority** | Pointer to **int32** | The priority assigned to the service. A lower value indicates higher priority. | [optional] 
**Order** | Pointer to **int32** | The order in which NAPTR records are to be processed. | [optional] 
**Preference** | Pointer to **int32** | The preference assigned to the NAPTR record. The preference is referenced when NAPTR records have the same order. | [optional] 
**Service** | Pointer to **string** | The service or protocol to which the NAPTR record resolves. | [optional] 
**RegularExpression** | Pointer to **string** | The regular expression used to transform client data. | [optional] 
**Replacement** | Pointer to **string** | The replacement string used to replace client data. | [optional] 
**Flags** | Pointer to **string** | The control flags in the NAPTR record. Control flags influence how the replacement functions are performed when processing the record. | [optional] 
**Weight** | Pointer to **int32** | The weight assigned to the service. The weight is referenced when services have the same priority. | [optional] 
**Port** | Pointer to **int32** | The port on which the service is located. | [optional] 
**Text** | Pointer to **string** | The text data within the TXT resource record. | [optional] 

## Methods

### NewGetResourceRecords200Response1DataInner

`func NewGetResourceRecords200Response1DataInner() *GetResourceRecords200Response1DataInner`

NewGetResourceRecords200Response1DataInner instantiates a new GetResourceRecords200Response1DataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetResourceRecords200Response1DataInnerWithDefaults

`func NewGetResourceRecords200Response1DataInnerWithDefaults() *GetResourceRecords200Response1DataInner`

NewGetResourceRecords200Response1DataInnerWithDefaults instantiates a new GetResourceRecords200Response1DataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetResourceRecords200Response1DataInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetResourceRecords200Response1DataInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetResourceRecords200Response1DataInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetResourceRecords200Response1DataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GetResourceRecords200Response1DataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetResourceRecords200Response1DataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetResourceRecords200Response1DataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetResourceRecords200Response1DataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *GetResourceRecords200Response1DataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetResourceRecords200Response1DataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetResourceRecords200Response1DataInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetResourceRecords200Response1DataInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *GetResourceRecords200Response1DataInner) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *GetResourceRecords200Response1DataInner) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *GetResourceRecords200Response1DataInner) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *GetResourceRecords200Response1DataInner) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetConfiguration

`func (o *GetResourceRecords200Response1DataInner) GetConfiguration() InlinedConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *GetResourceRecords200Response1DataInner) GetConfigurationOk() (*InlinedConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *GetResourceRecords200Response1DataInner) SetConfiguration(v InlinedConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *GetResourceRecords200Response1DataInner) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetTtl

`func (o *GetResourceRecords200Response1DataInner) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *GetResourceRecords200Response1DataInner) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *GetResourceRecords200Response1DataInner) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *GetResourceRecords200Response1DataInner) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetAbsoluteName

`func (o *GetResourceRecords200Response1DataInner) GetAbsoluteName() string`

GetAbsoluteName returns the AbsoluteName field if non-nil, zero value otherwise.

### GetAbsoluteNameOk

`func (o *GetResourceRecords200Response1DataInner) GetAbsoluteNameOk() (*string, bool)`

GetAbsoluteNameOk returns a tuple with the AbsoluteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsoluteName

`func (o *GetResourceRecords200Response1DataInner) SetAbsoluteName(v string)`

SetAbsoluteName sets AbsoluteName field to given value.

### HasAbsoluteName

`func (o *GetResourceRecords200Response1DataInner) HasAbsoluteName() bool`

HasAbsoluteName returns a boolean if a field has been set.

### GetComment

`func (o *GetResourceRecords200Response1DataInner) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetResourceRecords200Response1DataInner) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetResourceRecords200Response1DataInner) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetResourceRecords200Response1DataInner) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDynamic

`func (o *GetResourceRecords200Response1DataInner) GetDynamic() bool`

GetDynamic returns the Dynamic field if non-nil, zero value otherwise.

### GetDynamicOk

`func (o *GetResourceRecords200Response1DataInner) GetDynamicOk() (*bool, bool)`

GetDynamicOk returns a tuple with the Dynamic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamic

`func (o *GetResourceRecords200Response1DataInner) SetDynamic(v bool)`

SetDynamic sets Dynamic field to given value.

### HasDynamic

`func (o *GetResourceRecords200Response1DataInner) HasDynamic() bool`

HasDynamic returns a boolean if a field has been set.

### GetView

`func (o *GetResourceRecords200Response1DataInner) GetView() InlinedView`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *GetResourceRecords200Response1DataInner) GetViewOk() (*InlinedView, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *GetResourceRecords200Response1DataInner) SetView(v InlinedView)`

SetView sets View field to given value.

### HasView

`func (o *GetResourceRecords200Response1DataInner) HasView() bool`

HasView returns a boolean if a field has been set.

### GetRecordType

`func (o *GetResourceRecords200Response1DataInner) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *GetResourceRecords200Response1DataInner) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *GetResourceRecords200Response1DataInner) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *GetResourceRecords200Response1DataInner) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetRdata

`func (o *GetResourceRecords200Response1DataInner) GetRdata() string`

GetRdata returns the Rdata field if non-nil, zero value otherwise.

### GetRdataOk

`func (o *GetResourceRecords200Response1DataInner) GetRdataOk() (*string, bool)`

GetRdataOk returns a tuple with the Rdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdata

`func (o *GetResourceRecords200Response1DataInner) SetRdata(v string)`

SetRdata sets Rdata field to given value.

### HasRdata

`func (o *GetResourceRecords200Response1DataInner) HasRdata() bool`

HasRdata returns a boolean if a field has been set.

### GetOs

`func (o *GetResourceRecords200Response1DataInner) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *GetResourceRecords200Response1DataInner) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *GetResourceRecords200Response1DataInner) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *GetResourceRecords200Response1DataInner) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetCpu

`func (o *GetResourceRecords200Response1DataInner) GetCpu() string`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *GetResourceRecords200Response1DataInner) GetCpuOk() (*string, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *GetResourceRecords200Response1DataInner) SetCpu(v string)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *GetResourceRecords200Response1DataInner) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetAddresses

`func (o *GetResourceRecords200Response1DataInner) GetAddresses() []LinkedAddress`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *GetResourceRecords200Response1DataInner) GetAddressesOk() (*[]LinkedAddress, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *GetResourceRecords200Response1DataInner) SetAddresses(v []LinkedAddress)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *GetResourceRecords200Response1DataInner) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetReverseRecord

`func (o *GetResourceRecords200Response1DataInner) GetReverseRecord() bool`

GetReverseRecord returns the ReverseRecord field if non-nil, zero value otherwise.

### GetReverseRecordOk

`func (o *GetResourceRecords200Response1DataInner) GetReverseRecordOk() (*bool, bool)`

GetReverseRecordOk returns a tuple with the ReverseRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReverseRecord

`func (o *GetResourceRecords200Response1DataInner) SetReverseRecord(v bool)`

SetReverseRecord sets ReverseRecord field to given value.

### HasReverseRecord

`func (o *GetResourceRecords200Response1DataInner) HasReverseRecord() bool`

HasReverseRecord returns a boolean if a field has been set.

### GetViews

`func (o *GetResourceRecords200Response1DataInner) GetViews() []HostRecordAllOfViews`

GetViews returns the Views field if non-nil, zero value otherwise.

### GetViewsOk

`func (o *GetResourceRecords200Response1DataInner) GetViewsOk() (*[]HostRecordAllOfViews, bool)`

GetViewsOk returns a tuple with the Views field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViews

`func (o *GetResourceRecords200Response1DataInner) SetViews(v []HostRecordAllOfViews)`

SetViews sets Views field to given value.

### HasViews

`func (o *GetResourceRecords200Response1DataInner) HasViews() bool`

HasViews returns a boolean if a field has been set.

### GetLinkedRecord

`func (o *GetResourceRecords200Response1DataInner) GetLinkedRecord() SRVRecordAllOfLinkedRecord`

GetLinkedRecord returns the LinkedRecord field if non-nil, zero value otherwise.

### GetLinkedRecordOk

`func (o *GetResourceRecords200Response1DataInner) GetLinkedRecordOk() (*SRVRecordAllOfLinkedRecord, bool)`

GetLinkedRecordOk returns a tuple with the LinkedRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedRecord

`func (o *GetResourceRecords200Response1DataInner) SetLinkedRecord(v SRVRecordAllOfLinkedRecord)`

SetLinkedRecord sets LinkedRecord field to given value.

### HasLinkedRecord

`func (o *GetResourceRecords200Response1DataInner) HasLinkedRecord() bool`

HasLinkedRecord returns a boolean if a field has been set.

### GetPriority

`func (o *GetResourceRecords200Response1DataInner) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *GetResourceRecords200Response1DataInner) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *GetResourceRecords200Response1DataInner) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *GetResourceRecords200Response1DataInner) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetOrder

`func (o *GetResourceRecords200Response1DataInner) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *GetResourceRecords200Response1DataInner) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *GetResourceRecords200Response1DataInner) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *GetResourceRecords200Response1DataInner) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetPreference

`func (o *GetResourceRecords200Response1DataInner) GetPreference() int32`

GetPreference returns the Preference field if non-nil, zero value otherwise.

### GetPreferenceOk

`func (o *GetResourceRecords200Response1DataInner) GetPreferenceOk() (*int32, bool)`

GetPreferenceOk returns a tuple with the Preference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreference

`func (o *GetResourceRecords200Response1DataInner) SetPreference(v int32)`

SetPreference sets Preference field to given value.

### HasPreference

`func (o *GetResourceRecords200Response1DataInner) HasPreference() bool`

HasPreference returns a boolean if a field has been set.

### GetService

`func (o *GetResourceRecords200Response1DataInner) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *GetResourceRecords200Response1DataInner) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *GetResourceRecords200Response1DataInner) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *GetResourceRecords200Response1DataInner) HasService() bool`

HasService returns a boolean if a field has been set.

### GetRegularExpression

`func (o *GetResourceRecords200Response1DataInner) GetRegularExpression() string`

GetRegularExpression returns the RegularExpression field if non-nil, zero value otherwise.

### GetRegularExpressionOk

`func (o *GetResourceRecords200Response1DataInner) GetRegularExpressionOk() (*string, bool)`

GetRegularExpressionOk returns a tuple with the RegularExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegularExpression

`func (o *GetResourceRecords200Response1DataInner) SetRegularExpression(v string)`

SetRegularExpression sets RegularExpression field to given value.

### HasRegularExpression

`func (o *GetResourceRecords200Response1DataInner) HasRegularExpression() bool`

HasRegularExpression returns a boolean if a field has been set.

### GetReplacement

`func (o *GetResourceRecords200Response1DataInner) GetReplacement() string`

GetReplacement returns the Replacement field if non-nil, zero value otherwise.

### GetReplacementOk

`func (o *GetResourceRecords200Response1DataInner) GetReplacementOk() (*string, bool)`

GetReplacementOk returns a tuple with the Replacement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacement

`func (o *GetResourceRecords200Response1DataInner) SetReplacement(v string)`

SetReplacement sets Replacement field to given value.

### HasReplacement

`func (o *GetResourceRecords200Response1DataInner) HasReplacement() bool`

HasReplacement returns a boolean if a field has been set.

### GetFlags

`func (o *GetResourceRecords200Response1DataInner) GetFlags() string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *GetResourceRecords200Response1DataInner) GetFlagsOk() (*string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *GetResourceRecords200Response1DataInner) SetFlags(v string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *GetResourceRecords200Response1DataInner) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetWeight

`func (o *GetResourceRecords200Response1DataInner) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *GetResourceRecords200Response1DataInner) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *GetResourceRecords200Response1DataInner) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *GetResourceRecords200Response1DataInner) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetPort

`func (o *GetResourceRecords200Response1DataInner) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *GetResourceRecords200Response1DataInner) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *GetResourceRecords200Response1DataInner) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *GetResourceRecords200Response1DataInner) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetText

`func (o *GetResourceRecords200Response1DataInner) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *GetResourceRecords200Response1DataInner) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *GetResourceRecords200Response1DataInner) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *GetResourceRecords200Response1DataInner) HasText() bool`

HasText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


