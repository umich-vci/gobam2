# GetCollectionTemplates200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | Pointer to **string** | The name of the resource. | [optional] 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**Configuration** | Pointer to [**InlinedConfiguration**](InlinedConfiguration.md) |  | [optional] [readonly] 
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 

## Methods

### NewGetCollectionTemplates200ResponseDataInner

`func NewGetCollectionTemplates200ResponseDataInner() *GetCollectionTemplates200ResponseDataInner`

NewGetCollectionTemplates200ResponseDataInner instantiates a new GetCollectionTemplates200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCollectionTemplates200ResponseDataInnerWithDefaults

`func NewGetCollectionTemplates200ResponseDataInnerWithDefaults() *GetCollectionTemplates200ResponseDataInner`

NewGetCollectionTemplates200ResponseDataInnerWithDefaults instantiates a new GetCollectionTemplates200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetCollectionTemplates200ResponseDataInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetCollectionTemplates200ResponseDataInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetCollectionTemplates200ResponseDataInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetCollectionTemplates200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GetCollectionTemplates200ResponseDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetCollectionTemplates200ResponseDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetCollectionTemplates200ResponseDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetCollectionTemplates200ResponseDataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *GetCollectionTemplates200ResponseDataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetCollectionTemplates200ResponseDataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetCollectionTemplates200ResponseDataInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetCollectionTemplates200ResponseDataInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *GetCollectionTemplates200ResponseDataInner) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *GetCollectionTemplates200ResponseDataInner) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *GetCollectionTemplates200ResponseDataInner) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *GetCollectionTemplates200ResponseDataInner) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetConfiguration

`func (o *GetCollectionTemplates200ResponseDataInner) GetConfiguration() InlinedConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *GetCollectionTemplates200ResponseDataInner) GetConfigurationOk() (*InlinedConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *GetCollectionTemplates200ResponseDataInner) SetConfiguration(v InlinedConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *GetCollectionTemplates200ResponseDataInner) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetLinks

`func (o *GetCollectionTemplates200ResponseDataInner) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetCollectionTemplates200ResponseDataInner) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetCollectionTemplates200ResponseDataInner) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetCollectionTemplates200ResponseDataInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

