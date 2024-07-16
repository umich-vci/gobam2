# GetResourceRecordDependentRecords200Response1DataInner

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
**LinkedRecord** | Pointer to [**AliasRecordAllOfLinkedRecord**](AliasRecordAllOfLinkedRecord.md) |  | [optional] 
**View** | Pointer to [**InlinedView**](InlinedView.md) |  | [optional] 

## Methods

### NewGetResourceRecordDependentRecords200Response1DataInner

`func NewGetResourceRecordDependentRecords200Response1DataInner() *GetResourceRecordDependentRecords200Response1DataInner`

NewGetResourceRecordDependentRecords200Response1DataInner instantiates a new GetResourceRecordDependentRecords200Response1DataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetResourceRecordDependentRecords200Response1DataInnerWithDefaults

`func NewGetResourceRecordDependentRecords200Response1DataInnerWithDefaults() *GetResourceRecordDependentRecords200Response1DataInner`

NewGetResourceRecordDependentRecords200Response1DataInnerWithDefaults instantiates a new GetResourceRecordDependentRecords200Response1DataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetResourceRecordDependentRecords200Response1DataInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetResourceRecordDependentRecords200Response1DataInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetResourceRecordDependentRecords200Response1DataInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetResourceRecordDependentRecords200Response1DataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GetResourceRecordDependentRecords200Response1DataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetResourceRecordDependentRecords200Response1DataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetResourceRecordDependentRecords200Response1DataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetResourceRecordDependentRecords200Response1DataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *GetResourceRecordDependentRecords200Response1DataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetResourceRecordDependentRecords200Response1DataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetResourceRecordDependentRecords200Response1DataInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetResourceRecordDependentRecords200Response1DataInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *GetResourceRecordDependentRecords200Response1DataInner) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *GetResourceRecordDependentRecords200Response1DataInner) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *GetResourceRecordDependentRecords200Response1DataInner) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *GetResourceRecordDependentRecords200Response1DataInner) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetConfiguration

`func (o *GetResourceRecordDependentRecords200Response1DataInner) GetConfiguration() InlinedConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *GetResourceRecordDependentRecords200Response1DataInner) GetConfigurationOk() (*InlinedConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *GetResourceRecordDependentRecords200Response1DataInner) SetConfiguration(v InlinedConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *GetResourceRecordDependentRecords200Response1DataInner) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetTtl

`func (o *GetResourceRecordDependentRecords200Response1DataInner) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *GetResourceRecordDependentRecords200Response1DataInner) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *GetResourceRecordDependentRecords200Response1DataInner) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *GetResourceRecordDependentRecords200Response1DataInner) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetAbsoluteName

`func (o *GetResourceRecordDependentRecords200Response1DataInner) GetAbsoluteName() string`

GetAbsoluteName returns the AbsoluteName field if non-nil, zero value otherwise.

### GetAbsoluteNameOk

`func (o *GetResourceRecordDependentRecords200Response1DataInner) GetAbsoluteNameOk() (*string, bool)`

GetAbsoluteNameOk returns a tuple with the AbsoluteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsoluteName

`func (o *GetResourceRecordDependentRecords200Response1DataInner) SetAbsoluteName(v string)`

SetAbsoluteName sets AbsoluteName field to given value.

### HasAbsoluteName

`func (o *GetResourceRecordDependentRecords200Response1DataInner) HasAbsoluteName() bool`

HasAbsoluteName returns a boolean if a field has been set.

### GetComment

`func (o *GetResourceRecordDependentRecords200Response1DataInner) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetResourceRecordDependentRecords200Response1DataInner) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetResourceRecordDependentRecords200Response1DataInner) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetResourceRecordDependentRecords200Response1DataInner) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDynamic

`func (o *GetResourceRecordDependentRecords200Response1DataInner) GetDynamic() bool`

GetDynamic returns the Dynamic field if non-nil, zero value otherwise.

### GetDynamicOk

`func (o *GetResourceRecordDependentRecords200Response1DataInner) GetDynamicOk() (*bool, bool)`

GetDynamicOk returns a tuple with the Dynamic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamic

`func (o *GetResourceRecordDependentRecords200Response1DataInner) SetDynamic(v bool)`

SetDynamic sets Dynamic field to given value.

### HasDynamic

`func (o *GetResourceRecordDependentRecords200Response1DataInner) HasDynamic() bool`

HasDynamic returns a boolean if a field has been set.

### GetLinkedRecord

`func (o *GetResourceRecordDependentRecords200Response1DataInner) GetLinkedRecord() AliasRecordAllOfLinkedRecord`

GetLinkedRecord returns the LinkedRecord field if non-nil, zero value otherwise.

### GetLinkedRecordOk

`func (o *GetResourceRecordDependentRecords200Response1DataInner) GetLinkedRecordOk() (*AliasRecordAllOfLinkedRecord, bool)`

GetLinkedRecordOk returns a tuple with the LinkedRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedRecord

`func (o *GetResourceRecordDependentRecords200Response1DataInner) SetLinkedRecord(v AliasRecordAllOfLinkedRecord)`

SetLinkedRecord sets LinkedRecord field to given value.

### HasLinkedRecord

`func (o *GetResourceRecordDependentRecords200Response1DataInner) HasLinkedRecord() bool`

HasLinkedRecord returns a boolean if a field has been set.

### GetView

`func (o *GetResourceRecordDependentRecords200Response1DataInner) GetView() InlinedView`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *GetResourceRecordDependentRecords200Response1DataInner) GetViewOk() (*InlinedView, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *GetResourceRecordDependentRecords200Response1DataInner) SetView(v InlinedView)`

SetView sets View field to given value.

### HasView

`func (o *GetResourceRecordDependentRecords200Response1DataInner) HasView() bool`

HasView returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


