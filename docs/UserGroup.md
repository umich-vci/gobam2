# UserGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | Pointer to **string** | The name for the user, LDAP, TACACS+, or SSO group. | [optional] 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**GroupType** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** | The description of the SSO group. | [optional] 
**AdministratorPrivilege** | Pointer to **bool** | Indicates whether the users within the group have administrative user privileges. | [optional] 
**Authenticator** | Pointer to [**InlinedAuthenticator**](InlinedAuthenticator.md) |  | [optional] 

## Methods

### NewUserGroup

`func NewUserGroup() *UserGroup`

NewUserGroup instantiates a new UserGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserGroupWithDefaults

`func NewUserGroupWithDefaults() *UserGroup`

NewUserGroupWithDefaults instantiates a new UserGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserGroup) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserGroup) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserGroup) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UserGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *UserGroup) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserGroup) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserGroup) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UserGroup) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *UserGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *UserGroup) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *UserGroup) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *UserGroup) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *UserGroup) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetGroupType

`func (o *UserGroup) GetGroupType() string`

GetGroupType returns the GroupType field if non-nil, zero value otherwise.

### GetGroupTypeOk

`func (o *UserGroup) GetGroupTypeOk() (*string, bool)`

GetGroupTypeOk returns a tuple with the GroupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupType

`func (o *UserGroup) SetGroupType(v string)`

SetGroupType sets GroupType field to given value.

### HasGroupType

`func (o *UserGroup) HasGroupType() bool`

HasGroupType returns a boolean if a field has been set.

### GetDescription

`func (o *UserGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UserGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UserGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UserGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAdministratorPrivilege

`func (o *UserGroup) GetAdministratorPrivilege() bool`

GetAdministratorPrivilege returns the AdministratorPrivilege field if non-nil, zero value otherwise.

### GetAdministratorPrivilegeOk

`func (o *UserGroup) GetAdministratorPrivilegeOk() (*bool, bool)`

GetAdministratorPrivilegeOk returns a tuple with the AdministratorPrivilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministratorPrivilege

`func (o *UserGroup) SetAdministratorPrivilege(v bool)`

SetAdministratorPrivilege sets AdministratorPrivilege field to given value.

### HasAdministratorPrivilege

`func (o *UserGroup) HasAdministratorPrivilege() bool`

HasAdministratorPrivilege returns a boolean if a field has been set.

### GetAuthenticator

`func (o *UserGroup) GetAuthenticator() InlinedAuthenticator`

GetAuthenticator returns the Authenticator field if non-nil, zero value otherwise.

### GetAuthenticatorOk

`func (o *UserGroup) GetAuthenticatorOk() (*InlinedAuthenticator, bool)`

GetAuthenticatorOk returns a tuple with the Authenticator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticator

`func (o *UserGroup) SetAuthenticator(v InlinedAuthenticator)`

SetAuthenticator sets Authenticator field to given value.

### HasAuthenticator

`func (o *UserGroup) HasAuthenticator() bool`

HasAuthenticator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


