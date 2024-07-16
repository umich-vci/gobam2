# ACL

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

## Methods

### NewACL

`func NewACL() *ACL`

NewACL instantiates a new ACL object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewACLWithDefaults

`func NewACLWithDefaults() *ACL`

NewACLWithDefaults instantiates a new ACL object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ACL) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ACL) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ACL) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ACL) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ACL) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ACL) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ACL) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ACL) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *ACL) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ACL) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ACL) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ACL) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *ACL) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *ACL) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *ACL) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *ACL) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetConfiguration

`func (o *ACL) GetConfiguration() InlinedConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *ACL) GetConfigurationOk() (*InlinedConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *ACL) SetConfiguration(v InlinedConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *ACL) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetPredefined

`func (o *ACL) GetPredefined() bool`

GetPredefined returns the Predefined field if non-nil, zero value otherwise.

### GetPredefinedOk

`func (o *ACL) GetPredefinedOk() (*bool, bool)`

GetPredefinedOk returns a tuple with the Predefined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredefined

`func (o *ACL) SetPredefined(v bool)`

SetPredefined sets Predefined field to given value.

### HasPredefined

`func (o *ACL) HasPredefined() bool`

HasPredefined returns a boolean if a field has been set.

### GetMatchElements

`func (o *ACL) GetMatchElements() []string`

GetMatchElements returns the MatchElements field if non-nil, zero value otherwise.

### GetMatchElementsOk

`func (o *ACL) GetMatchElementsOk() (*[]string, bool)`

GetMatchElementsOk returns a tuple with the MatchElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchElements

`func (o *ACL) SetMatchElements(v []string)`

SetMatchElements sets MatchElements field to given value.

### HasMatchElements

`func (o *ACL) HasMatchElements() bool`

HasMatchElements returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


