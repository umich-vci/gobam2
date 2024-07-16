# GetUserDefinedLinkDefinitions200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | Pointer to **string** | The name of the user-defined link definition. | [optional] 
**DisplayName** | Pointer to **string** | The display name of the user-defined link definition. | [optional] 
**Description** | Pointer to **string** | The description of the user-defined link definition. | [optional] 
**SourceTypes** | Pointer to **[]string** | The source entity types of the user-defined link definition. | [optional] 
**DestinationTypes** | Pointer to **[]string** | The destination entity types of the user-defined link definition. | [optional] 
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 

## Methods

### NewGetUserDefinedLinkDefinitions200ResponseDataInner

`func NewGetUserDefinedLinkDefinitions200ResponseDataInner() *GetUserDefinedLinkDefinitions200ResponseDataInner`

NewGetUserDefinedLinkDefinitions200ResponseDataInner instantiates a new GetUserDefinedLinkDefinitions200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserDefinedLinkDefinitions200ResponseDataInnerWithDefaults

`func NewGetUserDefinedLinkDefinitions200ResponseDataInnerWithDefaults() *GetUserDefinedLinkDefinitions200ResponseDataInner`

NewGetUserDefinedLinkDefinitions200ResponseDataInnerWithDefaults instantiates a new GetUserDefinedLinkDefinitions200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetUserDefinedLinkDefinitions200ResponseDataInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetUserDefinedLinkDefinitions200ResponseDataInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetUserDefinedLinkDefinitions200ResponseDataInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetUserDefinedLinkDefinitions200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GetUserDefinedLinkDefinitions200ResponseDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetUserDefinedLinkDefinitions200ResponseDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetUserDefinedLinkDefinitions200ResponseDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetUserDefinedLinkDefinitions200ResponseDataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *GetUserDefinedLinkDefinitions200ResponseDataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetUserDefinedLinkDefinitions200ResponseDataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetUserDefinedLinkDefinitions200ResponseDataInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetUserDefinedLinkDefinitions200ResponseDataInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisplayName

`func (o *GetUserDefinedLinkDefinitions200ResponseDataInner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GetUserDefinedLinkDefinitions200ResponseDataInner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GetUserDefinedLinkDefinitions200ResponseDataInner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *GetUserDefinedLinkDefinitions200ResponseDataInner) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *GetUserDefinedLinkDefinitions200ResponseDataInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetUserDefinedLinkDefinitions200ResponseDataInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetUserDefinedLinkDefinitions200ResponseDataInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetUserDefinedLinkDefinitions200ResponseDataInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSourceTypes

`func (o *GetUserDefinedLinkDefinitions200ResponseDataInner) GetSourceTypes() []string`

GetSourceTypes returns the SourceTypes field if non-nil, zero value otherwise.

### GetSourceTypesOk

`func (o *GetUserDefinedLinkDefinitions200ResponseDataInner) GetSourceTypesOk() (*[]string, bool)`

GetSourceTypesOk returns a tuple with the SourceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTypes

`func (o *GetUserDefinedLinkDefinitions200ResponseDataInner) SetSourceTypes(v []string)`

SetSourceTypes sets SourceTypes field to given value.

### HasSourceTypes

`func (o *GetUserDefinedLinkDefinitions200ResponseDataInner) HasSourceTypes() bool`

HasSourceTypes returns a boolean if a field has been set.

### GetDestinationTypes

`func (o *GetUserDefinedLinkDefinitions200ResponseDataInner) GetDestinationTypes() []string`

GetDestinationTypes returns the DestinationTypes field if non-nil, zero value otherwise.

### GetDestinationTypesOk

`func (o *GetUserDefinedLinkDefinitions200ResponseDataInner) GetDestinationTypesOk() (*[]string, bool)`

GetDestinationTypesOk returns a tuple with the DestinationTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationTypes

`func (o *GetUserDefinedLinkDefinitions200ResponseDataInner) SetDestinationTypes(v []string)`

SetDestinationTypes sets DestinationTypes field to given value.

### HasDestinationTypes

`func (o *GetUserDefinedLinkDefinitions200ResponseDataInner) HasDestinationTypes() bool`

HasDestinationTypes returns a boolean if a field has been set.

### GetLinks

`func (o *GetUserDefinedLinkDefinitions200ResponseDataInner) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetUserDefinedLinkDefinitions200ResponseDataInner) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetUserDefinedLinkDefinitions200ResponseDataInner) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetUserDefinedLinkDefinitions200ResponseDataInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


