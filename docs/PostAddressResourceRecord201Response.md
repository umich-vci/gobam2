# PostAddressResourceRecord201Response

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
**View** | Pointer to [**InlinedView**](InlinedView.md) |  | [optional] 
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 

## Methods

### NewPostAddressResourceRecord201Response

`func NewPostAddressResourceRecord201Response() *PostAddressResourceRecord201Response`

NewPostAddressResourceRecord201Response instantiates a new PostAddressResourceRecord201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostAddressResourceRecord201ResponseWithDefaults

`func NewPostAddressResourceRecord201ResponseWithDefaults() *PostAddressResourceRecord201Response`

NewPostAddressResourceRecord201ResponseWithDefaults instantiates a new PostAddressResourceRecord201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PostAddressResourceRecord201Response) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostAddressResourceRecord201Response) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostAddressResourceRecord201Response) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *PostAddressResourceRecord201Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *PostAddressResourceRecord201Response) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PostAddressResourceRecord201Response) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PostAddressResourceRecord201Response) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PostAddressResourceRecord201Response) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *PostAddressResourceRecord201Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostAddressResourceRecord201Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostAddressResourceRecord201Response) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PostAddressResourceRecord201Response) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *PostAddressResourceRecord201Response) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *PostAddressResourceRecord201Response) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *PostAddressResourceRecord201Response) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *PostAddressResourceRecord201Response) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetConfiguration

`func (o *PostAddressResourceRecord201Response) GetConfiguration() InlinedConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *PostAddressResourceRecord201Response) GetConfigurationOk() (*InlinedConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *PostAddressResourceRecord201Response) SetConfiguration(v InlinedConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *PostAddressResourceRecord201Response) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetTtl

`func (o *PostAddressResourceRecord201Response) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *PostAddressResourceRecord201Response) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *PostAddressResourceRecord201Response) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *PostAddressResourceRecord201Response) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetAbsoluteName

`func (o *PostAddressResourceRecord201Response) GetAbsoluteName() string`

GetAbsoluteName returns the AbsoluteName field if non-nil, zero value otherwise.

### GetAbsoluteNameOk

`func (o *PostAddressResourceRecord201Response) GetAbsoluteNameOk() (*string, bool)`

GetAbsoluteNameOk returns a tuple with the AbsoluteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsoluteName

`func (o *PostAddressResourceRecord201Response) SetAbsoluteName(v string)`

SetAbsoluteName sets AbsoluteName field to given value.

### HasAbsoluteName

`func (o *PostAddressResourceRecord201Response) HasAbsoluteName() bool`

HasAbsoluteName returns a boolean if a field has been set.

### GetComment

`func (o *PostAddressResourceRecord201Response) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *PostAddressResourceRecord201Response) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *PostAddressResourceRecord201Response) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *PostAddressResourceRecord201Response) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDynamic

`func (o *PostAddressResourceRecord201Response) GetDynamic() bool`

GetDynamic returns the Dynamic field if non-nil, zero value otherwise.

### GetDynamicOk

`func (o *PostAddressResourceRecord201Response) GetDynamicOk() (*bool, bool)`

GetDynamicOk returns a tuple with the Dynamic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamic

`func (o *PostAddressResourceRecord201Response) SetDynamic(v bool)`

SetDynamic sets Dynamic field to given value.

### HasDynamic

`func (o *PostAddressResourceRecord201Response) HasDynamic() bool`

HasDynamic returns a boolean if a field has been set.

### GetAddresses

`func (o *PostAddressResourceRecord201Response) GetAddresses() []LinkedAddress`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *PostAddressResourceRecord201Response) GetAddressesOk() (*[]LinkedAddress, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *PostAddressResourceRecord201Response) SetAddresses(v []LinkedAddress)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *PostAddressResourceRecord201Response) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetReverseRecord

`func (o *PostAddressResourceRecord201Response) GetReverseRecord() bool`

GetReverseRecord returns the ReverseRecord field if non-nil, zero value otherwise.

### GetReverseRecordOk

`func (o *PostAddressResourceRecord201Response) GetReverseRecordOk() (*bool, bool)`

GetReverseRecordOk returns a tuple with the ReverseRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReverseRecord

`func (o *PostAddressResourceRecord201Response) SetReverseRecord(v bool)`

SetReverseRecord sets ReverseRecord field to given value.

### HasReverseRecord

`func (o *PostAddressResourceRecord201Response) HasReverseRecord() bool`

HasReverseRecord returns a boolean if a field has been set.

### GetViews

`func (o *PostAddressResourceRecord201Response) GetViews() []HostRecordAllOfViews`

GetViews returns the Views field if non-nil, zero value otherwise.

### GetViewsOk

`func (o *PostAddressResourceRecord201Response) GetViewsOk() (*[]HostRecordAllOfViews, bool)`

GetViewsOk returns a tuple with the Views field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViews

`func (o *PostAddressResourceRecord201Response) SetViews(v []HostRecordAllOfViews)`

SetViews sets Views field to given value.

### HasViews

`func (o *PostAddressResourceRecord201Response) HasViews() bool`

HasViews returns a boolean if a field has been set.

### GetView

`func (o *PostAddressResourceRecord201Response) GetView() InlinedView`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *PostAddressResourceRecord201Response) GetViewOk() (*InlinedView, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *PostAddressResourceRecord201Response) SetView(v InlinedView)`

SetView sets View field to given value.

### HasView

`func (o *PostAddressResourceRecord201Response) HasView() bool`

HasView returns a boolean if a field has been set.

### GetLinks

`func (o *PostAddressResourceRecord201Response) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PostAddressResourceRecord201Response) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PostAddressResourceRecord201Response) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PostAddressResourceRecord201Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


