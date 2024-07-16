# GetResourceRecordDependentRecords200ResponseDataInner

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
**Addresses** | Pointer to [**[]LinkedAddress**](LinkedAddress.md) |  | [optional] 
**ReverseRecord** | Pointer to **bool** | Indicates if a PTR record is created for the host record. | [optional] 
**Views** | Pointer to [**[]HostRecordAllOfViews**](HostRecordAllOfViews.md) |  | [optional] 
**LinkedRecord** | Pointer to [**AliasRecordAllOfLinkedRecord**](AliasRecordAllOfLinkedRecord.md) |  | [optional] 
**View** | Pointer to [**InlinedView**](InlinedView.md) |  | [optional] 
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 

## Methods

### NewGetResourceRecordDependentRecords200ResponseDataInner

`func NewGetResourceRecordDependentRecords200ResponseDataInner() *GetResourceRecordDependentRecords200ResponseDataInner`

NewGetResourceRecordDependentRecords200ResponseDataInner instantiates a new GetResourceRecordDependentRecords200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetResourceRecordDependentRecords200ResponseDataInnerWithDefaults

`func NewGetResourceRecordDependentRecords200ResponseDataInnerWithDefaults() *GetResourceRecordDependentRecords200ResponseDataInner`

NewGetResourceRecordDependentRecords200ResponseDataInnerWithDefaults instantiates a new GetResourceRecordDependentRecords200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetResourceRecordDependentRecords200ResponseDataInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetResourceRecordDependentRecords200ResponseDataInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetResourceRecordDependentRecords200ResponseDataInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetResourceRecordDependentRecords200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GetResourceRecordDependentRecords200ResponseDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetResourceRecordDependentRecords200ResponseDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetResourceRecordDependentRecords200ResponseDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetResourceRecordDependentRecords200ResponseDataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *GetResourceRecordDependentRecords200ResponseDataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetResourceRecordDependentRecords200ResponseDataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetResourceRecordDependentRecords200ResponseDataInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetResourceRecordDependentRecords200ResponseDataInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *GetResourceRecordDependentRecords200ResponseDataInner) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *GetResourceRecordDependentRecords200ResponseDataInner) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *GetResourceRecordDependentRecords200ResponseDataInner) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *GetResourceRecordDependentRecords200ResponseDataInner) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetConfiguration

`func (o *GetResourceRecordDependentRecords200ResponseDataInner) GetConfiguration() InlinedConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *GetResourceRecordDependentRecords200ResponseDataInner) GetConfigurationOk() (*InlinedConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *GetResourceRecordDependentRecords200ResponseDataInner) SetConfiguration(v InlinedConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *GetResourceRecordDependentRecords200ResponseDataInner) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetTtl

`func (o *GetResourceRecordDependentRecords200ResponseDataInner) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *GetResourceRecordDependentRecords200ResponseDataInner) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *GetResourceRecordDependentRecords200ResponseDataInner) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *GetResourceRecordDependentRecords200ResponseDataInner) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetAbsoluteName

`func (o *GetResourceRecordDependentRecords200ResponseDataInner) GetAbsoluteName() string`

GetAbsoluteName returns the AbsoluteName field if non-nil, zero value otherwise.

### GetAbsoluteNameOk

`func (o *GetResourceRecordDependentRecords200ResponseDataInner) GetAbsoluteNameOk() (*string, bool)`

GetAbsoluteNameOk returns a tuple with the AbsoluteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsoluteName

`func (o *GetResourceRecordDependentRecords200ResponseDataInner) SetAbsoluteName(v string)`

SetAbsoluteName sets AbsoluteName field to given value.

### HasAbsoluteName

`func (o *GetResourceRecordDependentRecords200ResponseDataInner) HasAbsoluteName() bool`

HasAbsoluteName returns a boolean if a field has been set.

### GetComment

`func (o *GetResourceRecordDependentRecords200ResponseDataInner) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetResourceRecordDependentRecords200ResponseDataInner) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetResourceRecordDependentRecords200ResponseDataInner) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetResourceRecordDependentRecords200ResponseDataInner) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDynamic

`func (o *GetResourceRecordDependentRecords200ResponseDataInner) GetDynamic() bool`

GetDynamic returns the Dynamic field if non-nil, zero value otherwise.

### GetDynamicOk

`func (o *GetResourceRecordDependentRecords200ResponseDataInner) GetDynamicOk() (*bool, bool)`

GetDynamicOk returns a tuple with the Dynamic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamic

`func (o *GetResourceRecordDependentRecords200ResponseDataInner) SetDynamic(v bool)`

SetDynamic sets Dynamic field to given value.

### HasDynamic

`func (o *GetResourceRecordDependentRecords200ResponseDataInner) HasDynamic() bool`

HasDynamic returns a boolean if a field has been set.

### GetAddresses

`func (o *GetResourceRecordDependentRecords200ResponseDataInner) GetAddresses() []LinkedAddress`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *GetResourceRecordDependentRecords200ResponseDataInner) GetAddressesOk() (*[]LinkedAddress, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *GetResourceRecordDependentRecords200ResponseDataInner) SetAddresses(v []LinkedAddress)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *GetResourceRecordDependentRecords200ResponseDataInner) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetReverseRecord

`func (o *GetResourceRecordDependentRecords200ResponseDataInner) GetReverseRecord() bool`

GetReverseRecord returns the ReverseRecord field if non-nil, zero value otherwise.

### GetReverseRecordOk

`func (o *GetResourceRecordDependentRecords200ResponseDataInner) GetReverseRecordOk() (*bool, bool)`

GetReverseRecordOk returns a tuple with the ReverseRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReverseRecord

`func (o *GetResourceRecordDependentRecords200ResponseDataInner) SetReverseRecord(v bool)`

SetReverseRecord sets ReverseRecord field to given value.

### HasReverseRecord

`func (o *GetResourceRecordDependentRecords200ResponseDataInner) HasReverseRecord() bool`

HasReverseRecord returns a boolean if a field has been set.

### GetViews

`func (o *GetResourceRecordDependentRecords200ResponseDataInner) GetViews() []HostRecordAllOfViews`

GetViews returns the Views field if non-nil, zero value otherwise.

### GetViewsOk

`func (o *GetResourceRecordDependentRecords200ResponseDataInner) GetViewsOk() (*[]HostRecordAllOfViews, bool)`

GetViewsOk returns a tuple with the Views field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViews

`func (o *GetResourceRecordDependentRecords200ResponseDataInner) SetViews(v []HostRecordAllOfViews)`

SetViews sets Views field to given value.

### HasViews

`func (o *GetResourceRecordDependentRecords200ResponseDataInner) HasViews() bool`

HasViews returns a boolean if a field has been set.

### GetLinkedRecord

`func (o *GetResourceRecordDependentRecords200ResponseDataInner) GetLinkedRecord() AliasRecordAllOfLinkedRecord`

GetLinkedRecord returns the LinkedRecord field if non-nil, zero value otherwise.

### GetLinkedRecordOk

`func (o *GetResourceRecordDependentRecords200ResponseDataInner) GetLinkedRecordOk() (*AliasRecordAllOfLinkedRecord, bool)`

GetLinkedRecordOk returns a tuple with the LinkedRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedRecord

`func (o *GetResourceRecordDependentRecords200ResponseDataInner) SetLinkedRecord(v AliasRecordAllOfLinkedRecord)`

SetLinkedRecord sets LinkedRecord field to given value.

### HasLinkedRecord

`func (o *GetResourceRecordDependentRecords200ResponseDataInner) HasLinkedRecord() bool`

HasLinkedRecord returns a boolean if a field has been set.

### GetView

`func (o *GetResourceRecordDependentRecords200ResponseDataInner) GetView() InlinedView`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *GetResourceRecordDependentRecords200ResponseDataInner) GetViewOk() (*InlinedView, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *GetResourceRecordDependentRecords200ResponseDataInner) SetView(v InlinedView)`

SetView sets View field to given value.

### HasView

`func (o *GetResourceRecordDependentRecords200ResponseDataInner) HasView() bool`

HasView returns a boolean if a field has been set.

### GetLinks

`func (o *GetResourceRecordDependentRecords200ResponseDataInner) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetResourceRecordDependentRecords200ResponseDataInner) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetResourceRecordDependentRecords200ResponseDataInner) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetResourceRecordDependentRecords200ResponseDataInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


