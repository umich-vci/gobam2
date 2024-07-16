# GetAccessControlLists200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | Pointer to **string** | The name of the resource. | [optional] 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**Configuration** | Pointer to [**InlinedConfiguration**](InlinedConfiguration.md) |  | [optional] [readonly] 
**Predefined** | Pointer to **bool** | Indicates whether the access control list is predefined by the system (true) or user-created (false). | [optional] [readonly] 
**MatchElements** | Pointer to **[]string** |  | [optional] 
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 

## Methods

### NewGetAccessControlLists200ResponseDataInner

`func NewGetAccessControlLists200ResponseDataInner() *GetAccessControlLists200ResponseDataInner`

NewGetAccessControlLists200ResponseDataInner instantiates a new GetAccessControlLists200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccessControlLists200ResponseDataInnerWithDefaults

`func NewGetAccessControlLists200ResponseDataInnerWithDefaults() *GetAccessControlLists200ResponseDataInner`

NewGetAccessControlLists200ResponseDataInnerWithDefaults instantiates a new GetAccessControlLists200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetAccessControlLists200ResponseDataInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetAccessControlLists200ResponseDataInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetAccessControlLists200ResponseDataInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetAccessControlLists200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GetAccessControlLists200ResponseDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetAccessControlLists200ResponseDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetAccessControlLists200ResponseDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetAccessControlLists200ResponseDataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *GetAccessControlLists200ResponseDataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetAccessControlLists200ResponseDataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetAccessControlLists200ResponseDataInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetAccessControlLists200ResponseDataInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *GetAccessControlLists200ResponseDataInner) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *GetAccessControlLists200ResponseDataInner) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *GetAccessControlLists200ResponseDataInner) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *GetAccessControlLists200ResponseDataInner) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetConfiguration

`func (o *GetAccessControlLists200ResponseDataInner) GetConfiguration() InlinedConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *GetAccessControlLists200ResponseDataInner) GetConfigurationOk() (*InlinedConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *GetAccessControlLists200ResponseDataInner) SetConfiguration(v InlinedConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *GetAccessControlLists200ResponseDataInner) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetPredefined

`func (o *GetAccessControlLists200ResponseDataInner) GetPredefined() bool`

GetPredefined returns the Predefined field if non-nil, zero value otherwise.

### GetPredefinedOk

`func (o *GetAccessControlLists200ResponseDataInner) GetPredefinedOk() (*bool, bool)`

GetPredefinedOk returns a tuple with the Predefined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredefined

`func (o *GetAccessControlLists200ResponseDataInner) SetPredefined(v bool)`

SetPredefined sets Predefined field to given value.

### HasPredefined

`func (o *GetAccessControlLists200ResponseDataInner) HasPredefined() bool`

HasPredefined returns a boolean if a field has been set.

### GetMatchElements

`func (o *GetAccessControlLists200ResponseDataInner) GetMatchElements() []string`

GetMatchElements returns the MatchElements field if non-nil, zero value otherwise.

### GetMatchElementsOk

`func (o *GetAccessControlLists200ResponseDataInner) GetMatchElementsOk() (*[]string, bool)`

GetMatchElementsOk returns a tuple with the MatchElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchElements

`func (o *GetAccessControlLists200ResponseDataInner) SetMatchElements(v []string)`

SetMatchElements sets MatchElements field to given value.

### HasMatchElements

`func (o *GetAccessControlLists200ResponseDataInner) HasMatchElements() bool`

HasMatchElements returns a boolean if a field has been set.

### GetLinks

`func (o *GetAccessControlLists200ResponseDataInner) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetAccessControlLists200ResponseDataInner) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetAccessControlLists200ResponseDataInner) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetAccessControlLists200ResponseDataInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


