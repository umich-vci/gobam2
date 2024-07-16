# PutResourceRecordRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | **string** | The name of the resource. | 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**Configuration** | Pointer to [**InlinedConfiguration**](InlinedConfiguration.md) |  | [optional] [readonly] 
**Ttl** | Pointer to **int64** | The overridden TTL value of the resource record, in seconds. | [optional] 
**AbsoluteName** | Pointer to **string** | The fully qualified domain name of the resource record. | [optional] 
**Comment** | Pointer to **string** | The additional comments associated with the resource record. | [optional] 
**Dynamic** | Pointer to **bool** | Indicates whether the resource record is dynamic. | [optional] [readonly] 
**View** | Pointer to [**InlinedView**](InlinedView.md) |  | [optional] 
**RecordType** | Pointer to **string** | The resource record type. | [optional] 
**Rdata** | **string** | The resource record data. | 
**Os** | Pointer to **string** | The string describing the server operating system. | [optional] 
**Cpu** | Pointer to **string** | The string describing the server CPU. | [optional] 
**Addresses** | Pointer to [**[]LinkedAddress**](LinkedAddress.md) |  | [optional] 
**ReverseRecord** | **bool** | Indicates if a PTR record is created for the host record. | 
**Views** | Pointer to [**[]HostRecordAllOfViews**](HostRecordAllOfViews.md) |  | [optional] 
**LinkedRecord** | [**SRVRecordAllOfLinkedRecord**](SRVRecordAllOfLinkedRecord.md) |  | 
**Priority** | **int32** | The priority assigned to the service. A lower value indicates higher priority. | 
**Order** | **int32** | The order in which NAPTR records are to be processed. | 
**Preference** | **int32** | The preference assigned to the NAPTR record. The preference is referenced when NAPTR records have the same order. | 
**Service** | Pointer to **string** | The service or protocol to which the NAPTR record resolves. | [optional] 
**RegularExpression** | Pointer to **string** | The regular expression used to transform client data. | [optional] 
**Replacement** | Pointer to **string** | The replacement string used to replace client data. | [optional] 
**Flags** | Pointer to **string** | The control flags in the NAPTR record. Control flags influence how the replacement functions are performed when processing the record. | [optional] 
**Weight** | **int32** | The weight assigned to the service. The weight is referenced when services have the same priority. | 
**Port** | **int32** | The port on which the service is located. | 
**Text** | Pointer to **string** | The text data within the TXT resource record. | [optional] 

## Methods

### NewPutResourceRecordRequest

`func NewPutResourceRecordRequest(name string, rdata string, reverseRecord bool, linkedRecord SRVRecordAllOfLinkedRecord, priority int32, order int32, preference int32, weight int32, port int32, ) *PutResourceRecordRequest`

NewPutResourceRecordRequest instantiates a new PutResourceRecordRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutResourceRecordRequestWithDefaults

`func NewPutResourceRecordRequestWithDefaults() *PutResourceRecordRequest`

NewPutResourceRecordRequestWithDefaults instantiates a new PutResourceRecordRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PutResourceRecordRequest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PutResourceRecordRequest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PutResourceRecordRequest) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *PutResourceRecordRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *PutResourceRecordRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PutResourceRecordRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PutResourceRecordRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PutResourceRecordRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *PutResourceRecordRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PutResourceRecordRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PutResourceRecordRequest) SetName(v string)`

SetName sets Name field to given value.


### GetUserDefinedFields

`func (o *PutResourceRecordRequest) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *PutResourceRecordRequest) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *PutResourceRecordRequest) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *PutResourceRecordRequest) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetConfiguration

`func (o *PutResourceRecordRequest) GetConfiguration() InlinedConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *PutResourceRecordRequest) GetConfigurationOk() (*InlinedConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *PutResourceRecordRequest) SetConfiguration(v InlinedConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *PutResourceRecordRequest) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetTtl

`func (o *PutResourceRecordRequest) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *PutResourceRecordRequest) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *PutResourceRecordRequest) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *PutResourceRecordRequest) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetAbsoluteName

`func (o *PutResourceRecordRequest) GetAbsoluteName() string`

GetAbsoluteName returns the AbsoluteName field if non-nil, zero value otherwise.

### GetAbsoluteNameOk

`func (o *PutResourceRecordRequest) GetAbsoluteNameOk() (*string, bool)`

GetAbsoluteNameOk returns a tuple with the AbsoluteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsoluteName

`func (o *PutResourceRecordRequest) SetAbsoluteName(v string)`

SetAbsoluteName sets AbsoluteName field to given value.

### HasAbsoluteName

`func (o *PutResourceRecordRequest) HasAbsoluteName() bool`

HasAbsoluteName returns a boolean if a field has been set.

### GetComment

`func (o *PutResourceRecordRequest) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *PutResourceRecordRequest) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *PutResourceRecordRequest) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *PutResourceRecordRequest) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDynamic

`func (o *PutResourceRecordRequest) GetDynamic() bool`

GetDynamic returns the Dynamic field if non-nil, zero value otherwise.

### GetDynamicOk

`func (o *PutResourceRecordRequest) GetDynamicOk() (*bool, bool)`

GetDynamicOk returns a tuple with the Dynamic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamic

`func (o *PutResourceRecordRequest) SetDynamic(v bool)`

SetDynamic sets Dynamic field to given value.

### HasDynamic

`func (o *PutResourceRecordRequest) HasDynamic() bool`

HasDynamic returns a boolean if a field has been set.

### GetView

`func (o *PutResourceRecordRequest) GetView() InlinedView`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *PutResourceRecordRequest) GetViewOk() (*InlinedView, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *PutResourceRecordRequest) SetView(v InlinedView)`

SetView sets View field to given value.

### HasView

`func (o *PutResourceRecordRequest) HasView() bool`

HasView returns a boolean if a field has been set.

### GetRecordType

`func (o *PutResourceRecordRequest) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *PutResourceRecordRequest) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *PutResourceRecordRequest) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *PutResourceRecordRequest) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetRdata

`func (o *PutResourceRecordRequest) GetRdata() string`

GetRdata returns the Rdata field if non-nil, zero value otherwise.

### GetRdataOk

`func (o *PutResourceRecordRequest) GetRdataOk() (*string, bool)`

GetRdataOk returns a tuple with the Rdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdata

`func (o *PutResourceRecordRequest) SetRdata(v string)`

SetRdata sets Rdata field to given value.


### GetOs

`func (o *PutResourceRecordRequest) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *PutResourceRecordRequest) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *PutResourceRecordRequest) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *PutResourceRecordRequest) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetCpu

`func (o *PutResourceRecordRequest) GetCpu() string`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *PutResourceRecordRequest) GetCpuOk() (*string, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *PutResourceRecordRequest) SetCpu(v string)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *PutResourceRecordRequest) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetAddresses

`func (o *PutResourceRecordRequest) GetAddresses() []LinkedAddress`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *PutResourceRecordRequest) GetAddressesOk() (*[]LinkedAddress, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *PutResourceRecordRequest) SetAddresses(v []LinkedAddress)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *PutResourceRecordRequest) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetReverseRecord

`func (o *PutResourceRecordRequest) GetReverseRecord() bool`

GetReverseRecord returns the ReverseRecord field if non-nil, zero value otherwise.

### GetReverseRecordOk

`func (o *PutResourceRecordRequest) GetReverseRecordOk() (*bool, bool)`

GetReverseRecordOk returns a tuple with the ReverseRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReverseRecord

`func (o *PutResourceRecordRequest) SetReverseRecord(v bool)`

SetReverseRecord sets ReverseRecord field to given value.


### GetViews

`func (o *PutResourceRecordRequest) GetViews() []HostRecordAllOfViews`

GetViews returns the Views field if non-nil, zero value otherwise.

### GetViewsOk

`func (o *PutResourceRecordRequest) GetViewsOk() (*[]HostRecordAllOfViews, bool)`

GetViewsOk returns a tuple with the Views field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViews

`func (o *PutResourceRecordRequest) SetViews(v []HostRecordAllOfViews)`

SetViews sets Views field to given value.

### HasViews

`func (o *PutResourceRecordRequest) HasViews() bool`

HasViews returns a boolean if a field has been set.

### GetLinkedRecord

`func (o *PutResourceRecordRequest) GetLinkedRecord() SRVRecordAllOfLinkedRecord`

GetLinkedRecord returns the LinkedRecord field if non-nil, zero value otherwise.

### GetLinkedRecordOk

`func (o *PutResourceRecordRequest) GetLinkedRecordOk() (*SRVRecordAllOfLinkedRecord, bool)`

GetLinkedRecordOk returns a tuple with the LinkedRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedRecord

`func (o *PutResourceRecordRequest) SetLinkedRecord(v SRVRecordAllOfLinkedRecord)`

SetLinkedRecord sets LinkedRecord field to given value.


### GetPriority

`func (o *PutResourceRecordRequest) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *PutResourceRecordRequest) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *PutResourceRecordRequest) SetPriority(v int32)`

SetPriority sets Priority field to given value.


### GetOrder

`func (o *PutResourceRecordRequest) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *PutResourceRecordRequest) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *PutResourceRecordRequest) SetOrder(v int32)`

SetOrder sets Order field to given value.


### GetPreference

`func (o *PutResourceRecordRequest) GetPreference() int32`

GetPreference returns the Preference field if non-nil, zero value otherwise.

### GetPreferenceOk

`func (o *PutResourceRecordRequest) GetPreferenceOk() (*int32, bool)`

GetPreferenceOk returns a tuple with the Preference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreference

`func (o *PutResourceRecordRequest) SetPreference(v int32)`

SetPreference sets Preference field to given value.


### GetService

`func (o *PutResourceRecordRequest) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *PutResourceRecordRequest) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *PutResourceRecordRequest) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *PutResourceRecordRequest) HasService() bool`

HasService returns a boolean if a field has been set.

### GetRegularExpression

`func (o *PutResourceRecordRequest) GetRegularExpression() string`

GetRegularExpression returns the RegularExpression field if non-nil, zero value otherwise.

### GetRegularExpressionOk

`func (o *PutResourceRecordRequest) GetRegularExpressionOk() (*string, bool)`

GetRegularExpressionOk returns a tuple with the RegularExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegularExpression

`func (o *PutResourceRecordRequest) SetRegularExpression(v string)`

SetRegularExpression sets RegularExpression field to given value.

### HasRegularExpression

`func (o *PutResourceRecordRequest) HasRegularExpression() bool`

HasRegularExpression returns a boolean if a field has been set.

### GetReplacement

`func (o *PutResourceRecordRequest) GetReplacement() string`

GetReplacement returns the Replacement field if non-nil, zero value otherwise.

### GetReplacementOk

`func (o *PutResourceRecordRequest) GetReplacementOk() (*string, bool)`

GetReplacementOk returns a tuple with the Replacement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacement

`func (o *PutResourceRecordRequest) SetReplacement(v string)`

SetReplacement sets Replacement field to given value.

### HasReplacement

`func (o *PutResourceRecordRequest) HasReplacement() bool`

HasReplacement returns a boolean if a field has been set.

### GetFlags

`func (o *PutResourceRecordRequest) GetFlags() string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *PutResourceRecordRequest) GetFlagsOk() (*string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *PutResourceRecordRequest) SetFlags(v string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *PutResourceRecordRequest) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetWeight

`func (o *PutResourceRecordRequest) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *PutResourceRecordRequest) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *PutResourceRecordRequest) SetWeight(v int32)`

SetWeight sets Weight field to given value.


### GetPort

`func (o *PutResourceRecordRequest) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *PutResourceRecordRequest) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *PutResourceRecordRequest) SetPort(v int32)`

SetPort sets Port field to given value.


### GetText

`func (o *PutResourceRecordRequest) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *PutResourceRecordRequest) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *PutResourceRecordRequest) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *PutResourceRecordRequest) HasText() bool`

HasText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


