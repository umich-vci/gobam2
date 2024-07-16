# PostAddressResourceRecord201Response1

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

## Methods

### NewPostAddressResourceRecord201Response1

`func NewPostAddressResourceRecord201Response1() *PostAddressResourceRecord201Response1`

NewPostAddressResourceRecord201Response1 instantiates a new PostAddressResourceRecord201Response1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostAddressResourceRecord201Response1WithDefaults

`func NewPostAddressResourceRecord201Response1WithDefaults() *PostAddressResourceRecord201Response1`

NewPostAddressResourceRecord201Response1WithDefaults instantiates a new PostAddressResourceRecord201Response1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PostAddressResourceRecord201Response1) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostAddressResourceRecord201Response1) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostAddressResourceRecord201Response1) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *PostAddressResourceRecord201Response1) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *PostAddressResourceRecord201Response1) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PostAddressResourceRecord201Response1) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PostAddressResourceRecord201Response1) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PostAddressResourceRecord201Response1) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *PostAddressResourceRecord201Response1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostAddressResourceRecord201Response1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostAddressResourceRecord201Response1) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PostAddressResourceRecord201Response1) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *PostAddressResourceRecord201Response1) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *PostAddressResourceRecord201Response1) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *PostAddressResourceRecord201Response1) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *PostAddressResourceRecord201Response1) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetConfiguration

`func (o *PostAddressResourceRecord201Response1) GetConfiguration() InlinedConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *PostAddressResourceRecord201Response1) GetConfigurationOk() (*InlinedConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *PostAddressResourceRecord201Response1) SetConfiguration(v InlinedConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *PostAddressResourceRecord201Response1) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetTtl

`func (o *PostAddressResourceRecord201Response1) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *PostAddressResourceRecord201Response1) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *PostAddressResourceRecord201Response1) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *PostAddressResourceRecord201Response1) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetAbsoluteName

`func (o *PostAddressResourceRecord201Response1) GetAbsoluteName() string`

GetAbsoluteName returns the AbsoluteName field if non-nil, zero value otherwise.

### GetAbsoluteNameOk

`func (o *PostAddressResourceRecord201Response1) GetAbsoluteNameOk() (*string, bool)`

GetAbsoluteNameOk returns a tuple with the AbsoluteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsoluteName

`func (o *PostAddressResourceRecord201Response1) SetAbsoluteName(v string)`

SetAbsoluteName sets AbsoluteName field to given value.

### HasAbsoluteName

`func (o *PostAddressResourceRecord201Response1) HasAbsoluteName() bool`

HasAbsoluteName returns a boolean if a field has been set.

### GetComment

`func (o *PostAddressResourceRecord201Response1) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *PostAddressResourceRecord201Response1) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *PostAddressResourceRecord201Response1) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *PostAddressResourceRecord201Response1) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDynamic

`func (o *PostAddressResourceRecord201Response1) GetDynamic() bool`

GetDynamic returns the Dynamic field if non-nil, zero value otherwise.

### GetDynamicOk

`func (o *PostAddressResourceRecord201Response1) GetDynamicOk() (*bool, bool)`

GetDynamicOk returns a tuple with the Dynamic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamic

`func (o *PostAddressResourceRecord201Response1) SetDynamic(v bool)`

SetDynamic sets Dynamic field to given value.

### HasDynamic

`func (o *PostAddressResourceRecord201Response1) HasDynamic() bool`

HasDynamic returns a boolean if a field has been set.

### GetView

`func (o *PostAddressResourceRecord201Response1) GetView() InlinedView`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *PostAddressResourceRecord201Response1) GetViewOk() (*InlinedView, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *PostAddressResourceRecord201Response1) SetView(v InlinedView)`

SetView sets View field to given value.

### HasView

`func (o *PostAddressResourceRecord201Response1) HasView() bool`

HasView returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


