# GetClientIdentifiers200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | Pointer to **string** | The name of the resource. | [optional] 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**Configuration** | Pointer to [**InlinedConfiguration**](InlinedConfiguration.md) |  | [optional] [readonly] 
**Uid** | Pointer to [**UID**](UID.md) |  | [optional] 
**Duid** | Pointer to **string** | The DHCP unique identifier. | [optional] 
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 

## Methods

### NewGetClientIdentifiers200ResponseDataInner

`func NewGetClientIdentifiers200ResponseDataInner() *GetClientIdentifiers200ResponseDataInner`

NewGetClientIdentifiers200ResponseDataInner instantiates a new GetClientIdentifiers200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetClientIdentifiers200ResponseDataInnerWithDefaults

`func NewGetClientIdentifiers200ResponseDataInnerWithDefaults() *GetClientIdentifiers200ResponseDataInner`

NewGetClientIdentifiers200ResponseDataInnerWithDefaults instantiates a new GetClientIdentifiers200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetClientIdentifiers200ResponseDataInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetClientIdentifiers200ResponseDataInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetClientIdentifiers200ResponseDataInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetClientIdentifiers200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GetClientIdentifiers200ResponseDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetClientIdentifiers200ResponseDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetClientIdentifiers200ResponseDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetClientIdentifiers200ResponseDataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *GetClientIdentifiers200ResponseDataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetClientIdentifiers200ResponseDataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetClientIdentifiers200ResponseDataInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetClientIdentifiers200ResponseDataInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *GetClientIdentifiers200ResponseDataInner) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *GetClientIdentifiers200ResponseDataInner) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *GetClientIdentifiers200ResponseDataInner) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *GetClientIdentifiers200ResponseDataInner) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetConfiguration

`func (o *GetClientIdentifiers200ResponseDataInner) GetConfiguration() InlinedConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *GetClientIdentifiers200ResponseDataInner) GetConfigurationOk() (*InlinedConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *GetClientIdentifiers200ResponseDataInner) SetConfiguration(v InlinedConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *GetClientIdentifiers200ResponseDataInner) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetUid

`func (o *GetClientIdentifiers200ResponseDataInner) GetUid() UID`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *GetClientIdentifiers200ResponseDataInner) GetUidOk() (*UID, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *GetClientIdentifiers200ResponseDataInner) SetUid(v UID)`

SetUid sets Uid field to given value.

### HasUid

`func (o *GetClientIdentifiers200ResponseDataInner) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetDuid

`func (o *GetClientIdentifiers200ResponseDataInner) GetDuid() string`

GetDuid returns the Duid field if non-nil, zero value otherwise.

### GetDuidOk

`func (o *GetClientIdentifiers200ResponseDataInner) GetDuidOk() (*string, bool)`

GetDuidOk returns a tuple with the Duid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuid

`func (o *GetClientIdentifiers200ResponseDataInner) SetDuid(v string)`

SetDuid sets Duid field to given value.

### HasDuid

`func (o *GetClientIdentifiers200ResponseDataInner) HasDuid() bool`

HasDuid returns a boolean if a field has been set.

### GetLinks

`func (o *GetClientIdentifiers200ResponseDataInner) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetClientIdentifiers200ResponseDataInner) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetClientIdentifiers200ResponseDataInner) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetClientIdentifiers200ResponseDataInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


