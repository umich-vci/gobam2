# PostCollectionResourceRecordRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | **string** | The resource type. | 
**Name** | **string** | The name of the resource. | 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**Configuration** | Pointer to [**InlinedConfiguration**](InlinedConfiguration.md) |  | [optional] [readonly] 
**Ttl** | Pointer to **int64** | The overridden TTL value of the resource record, in seconds. | [optional] 
**AbsoluteName** | Pointer to **string** | The fully qualified domain name of the resource record. | [optional] 
**Comment** | Pointer to **string** | The additional comments associated with the resource record. | [optional] 
**Dynamic** | Pointer to **bool** | Indicates whether the resource record is dynamic. | [optional] [readonly] 
**View** | Pointer to [**InlinedView**](InlinedView.md) |  | [optional] 
**RecordType** | **string** | The resource record type. | 
**Rdata** | **string** | The resource record data. | 
**Os** | Pointer to **string** | The string describing the server operating system. | [optional] 
**Cpu** | Pointer to **string** | The string describing the server CPU. | [optional] 
**Addresses** | [**[]LinkedAddress**](LinkedAddress.md) |  | 
**ReverseRecord** | Pointer to **bool** | Indicates if a PTR record is created for the host record. | [optional] 
**Views** | Pointer to [**[]HostRecordAllOfViews**](HostRecordAllOfViews.md) |  | [optional] 
**LinkedRecord** | [**SRVRecordAllOfLinkedRecord**](SRVRecordAllOfLinkedRecord.md) |  | 
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

### NewPostCollectionResourceRecordRequest

`func NewPostCollectionResourceRecordRequest(type_ string, name string, recordType string, rdata string, addresses []LinkedAddress, linkedRecord SRVRecordAllOfLinkedRecord, ) *PostCollectionResourceRecordRequest`

NewPostCollectionResourceRecordRequest instantiates a new PostCollectionResourceRecordRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostCollectionResourceRecordRequestWithDefaults

`func NewPostCollectionResourceRecordRequestWithDefaults() *PostCollectionResourceRecordRequest`

NewPostCollectionResourceRecordRequestWithDefaults instantiates a new PostCollectionResourceRecordRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PostCollectionResourceRecordRequest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostCollectionResourceRecordRequest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostCollectionResourceRecordRequest) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *PostCollectionResourceRecordRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *PostCollectionResourceRecordRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PostCollectionResourceRecordRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PostCollectionResourceRecordRequest) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *PostCollectionResourceRecordRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostCollectionResourceRecordRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostCollectionResourceRecordRequest) SetName(v string)`

SetName sets Name field to given value.


### GetUserDefinedFields

`func (o *PostCollectionResourceRecordRequest) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *PostCollectionResourceRecordRequest) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *PostCollectionResourceRecordRequest) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *PostCollectionResourceRecordRequest) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetConfiguration

`func (o *PostCollectionResourceRecordRequest) GetConfiguration() InlinedConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *PostCollectionResourceRecordRequest) GetConfigurationOk() (*InlinedConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *PostCollectionResourceRecordRequest) SetConfiguration(v InlinedConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *PostCollectionResourceRecordRequest) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetTtl

`func (o *PostCollectionResourceRecordRequest) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *PostCollectionResourceRecordRequest) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *PostCollectionResourceRecordRequest) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *PostCollectionResourceRecordRequest) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetAbsoluteName

`func (o *PostCollectionResourceRecordRequest) GetAbsoluteName() string`

GetAbsoluteName returns the AbsoluteName field if non-nil, zero value otherwise.

### GetAbsoluteNameOk

`func (o *PostCollectionResourceRecordRequest) GetAbsoluteNameOk() (*string, bool)`

GetAbsoluteNameOk returns a tuple with the AbsoluteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsoluteName

`func (o *PostCollectionResourceRecordRequest) SetAbsoluteName(v string)`

SetAbsoluteName sets AbsoluteName field to given value.

### HasAbsoluteName

`func (o *PostCollectionResourceRecordRequest) HasAbsoluteName() bool`

HasAbsoluteName returns a boolean if a field has been set.

### GetComment

`func (o *PostCollectionResourceRecordRequest) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *PostCollectionResourceRecordRequest) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *PostCollectionResourceRecordRequest) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *PostCollectionResourceRecordRequest) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDynamic

`func (o *PostCollectionResourceRecordRequest) GetDynamic() bool`

GetDynamic returns the Dynamic field if non-nil, zero value otherwise.

### GetDynamicOk

`func (o *PostCollectionResourceRecordRequest) GetDynamicOk() (*bool, bool)`

GetDynamicOk returns a tuple with the Dynamic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamic

`func (o *PostCollectionResourceRecordRequest) SetDynamic(v bool)`

SetDynamic sets Dynamic field to given value.

### HasDynamic

`func (o *PostCollectionResourceRecordRequest) HasDynamic() bool`

HasDynamic returns a boolean if a field has been set.

### GetView

`func (o *PostCollectionResourceRecordRequest) GetView() InlinedView`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *PostCollectionResourceRecordRequest) GetViewOk() (*InlinedView, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *PostCollectionResourceRecordRequest) SetView(v InlinedView)`

SetView sets View field to given value.

### HasView

`func (o *PostCollectionResourceRecordRequest) HasView() bool`

HasView returns a boolean if a field has been set.

### GetRecordType

`func (o *PostCollectionResourceRecordRequest) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *PostCollectionResourceRecordRequest) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *PostCollectionResourceRecordRequest) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.


### GetRdata

`func (o *PostCollectionResourceRecordRequest) GetRdata() string`

GetRdata returns the Rdata field if non-nil, zero value otherwise.

### GetRdataOk

`func (o *PostCollectionResourceRecordRequest) GetRdataOk() (*string, bool)`

GetRdataOk returns a tuple with the Rdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdata

`func (o *PostCollectionResourceRecordRequest) SetRdata(v string)`

SetRdata sets Rdata field to given value.


### GetOs

`func (o *PostCollectionResourceRecordRequest) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *PostCollectionResourceRecordRequest) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *PostCollectionResourceRecordRequest) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *PostCollectionResourceRecordRequest) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetCpu

`func (o *PostCollectionResourceRecordRequest) GetCpu() string`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *PostCollectionResourceRecordRequest) GetCpuOk() (*string, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *PostCollectionResourceRecordRequest) SetCpu(v string)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *PostCollectionResourceRecordRequest) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetAddresses

`func (o *PostCollectionResourceRecordRequest) GetAddresses() []LinkedAddress`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *PostCollectionResourceRecordRequest) GetAddressesOk() (*[]LinkedAddress, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *PostCollectionResourceRecordRequest) SetAddresses(v []LinkedAddress)`

SetAddresses sets Addresses field to given value.


### GetReverseRecord

`func (o *PostCollectionResourceRecordRequest) GetReverseRecord() bool`

GetReverseRecord returns the ReverseRecord field if non-nil, zero value otherwise.

### GetReverseRecordOk

`func (o *PostCollectionResourceRecordRequest) GetReverseRecordOk() (*bool, bool)`

GetReverseRecordOk returns a tuple with the ReverseRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReverseRecord

`func (o *PostCollectionResourceRecordRequest) SetReverseRecord(v bool)`

SetReverseRecord sets ReverseRecord field to given value.

### HasReverseRecord

`func (o *PostCollectionResourceRecordRequest) HasReverseRecord() bool`

HasReverseRecord returns a boolean if a field has been set.

### GetViews

`func (o *PostCollectionResourceRecordRequest) GetViews() []HostRecordAllOfViews`

GetViews returns the Views field if non-nil, zero value otherwise.

### GetViewsOk

`func (o *PostCollectionResourceRecordRequest) GetViewsOk() (*[]HostRecordAllOfViews, bool)`

GetViewsOk returns a tuple with the Views field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViews

`func (o *PostCollectionResourceRecordRequest) SetViews(v []HostRecordAllOfViews)`

SetViews sets Views field to given value.

### HasViews

`func (o *PostCollectionResourceRecordRequest) HasViews() bool`

HasViews returns a boolean if a field has been set.

### GetLinkedRecord

`func (o *PostCollectionResourceRecordRequest) GetLinkedRecord() SRVRecordAllOfLinkedRecord`

GetLinkedRecord returns the LinkedRecord field if non-nil, zero value otherwise.

### GetLinkedRecordOk

`func (o *PostCollectionResourceRecordRequest) GetLinkedRecordOk() (*SRVRecordAllOfLinkedRecord, bool)`

GetLinkedRecordOk returns a tuple with the LinkedRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedRecord

`func (o *PostCollectionResourceRecordRequest) SetLinkedRecord(v SRVRecordAllOfLinkedRecord)`

SetLinkedRecord sets LinkedRecord field to given value.


### GetPriority

`func (o *PostCollectionResourceRecordRequest) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *PostCollectionResourceRecordRequest) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *PostCollectionResourceRecordRequest) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *PostCollectionResourceRecordRequest) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetOrder

`func (o *PostCollectionResourceRecordRequest) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *PostCollectionResourceRecordRequest) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *PostCollectionResourceRecordRequest) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *PostCollectionResourceRecordRequest) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetPreference

`func (o *PostCollectionResourceRecordRequest) GetPreference() int32`

GetPreference returns the Preference field if non-nil, zero value otherwise.

### GetPreferenceOk

`func (o *PostCollectionResourceRecordRequest) GetPreferenceOk() (*int32, bool)`

GetPreferenceOk returns a tuple with the Preference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreference

`func (o *PostCollectionResourceRecordRequest) SetPreference(v int32)`

SetPreference sets Preference field to given value.

### HasPreference

`func (o *PostCollectionResourceRecordRequest) HasPreference() bool`

HasPreference returns a boolean if a field has been set.

### GetService

`func (o *PostCollectionResourceRecordRequest) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *PostCollectionResourceRecordRequest) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *PostCollectionResourceRecordRequest) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *PostCollectionResourceRecordRequest) HasService() bool`

HasService returns a boolean if a field has been set.

### GetRegularExpression

`func (o *PostCollectionResourceRecordRequest) GetRegularExpression() string`

GetRegularExpression returns the RegularExpression field if non-nil, zero value otherwise.

### GetRegularExpressionOk

`func (o *PostCollectionResourceRecordRequest) GetRegularExpressionOk() (*string, bool)`

GetRegularExpressionOk returns a tuple with the RegularExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegularExpression

`func (o *PostCollectionResourceRecordRequest) SetRegularExpression(v string)`

SetRegularExpression sets RegularExpression field to given value.

### HasRegularExpression

`func (o *PostCollectionResourceRecordRequest) HasRegularExpression() bool`

HasRegularExpression returns a boolean if a field has been set.

### GetReplacement

`func (o *PostCollectionResourceRecordRequest) GetReplacement() string`

GetReplacement returns the Replacement field if non-nil, zero value otherwise.

### GetReplacementOk

`func (o *PostCollectionResourceRecordRequest) GetReplacementOk() (*string, bool)`

GetReplacementOk returns a tuple with the Replacement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacement

`func (o *PostCollectionResourceRecordRequest) SetReplacement(v string)`

SetReplacement sets Replacement field to given value.

### HasReplacement

`func (o *PostCollectionResourceRecordRequest) HasReplacement() bool`

HasReplacement returns a boolean if a field has been set.

### GetFlags

`func (o *PostCollectionResourceRecordRequest) GetFlags() string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *PostCollectionResourceRecordRequest) GetFlagsOk() (*string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *PostCollectionResourceRecordRequest) SetFlags(v string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *PostCollectionResourceRecordRequest) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetWeight

`func (o *PostCollectionResourceRecordRequest) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *PostCollectionResourceRecordRequest) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *PostCollectionResourceRecordRequest) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *PostCollectionResourceRecordRequest) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetPort

`func (o *PostCollectionResourceRecordRequest) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *PostCollectionResourceRecordRequest) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *PostCollectionResourceRecordRequest) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *PostCollectionResourceRecordRequest) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetText

`func (o *PostCollectionResourceRecordRequest) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *PostCollectionResourceRecordRequest) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *PostCollectionResourceRecordRequest) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *PostCollectionResourceRecordRequest) HasText() bool`

HasText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


