# UserGroupPostRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | **string** | The name for the user, LDAP, TACACS+, or SSO group. | 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**GroupType** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** | The description of the SSO group. | [optional] 
**AdministratorPrivilege** | Pointer to **bool** | Indicates whether the users within the group have administrative user privileges. | [optional] 
**Authenticator** | Pointer to [**InlinedAuthenticator**](InlinedAuthenticator.md) |  | [optional] 

## Methods

### NewUserGroupPostRequestBody

`func NewUserGroupPostRequestBody(name string, ) *UserGroupPostRequestBody`

NewUserGroupPostRequestBody instantiates a new UserGroupPostRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserGroupPostRequestBodyWithDefaults

`func NewUserGroupPostRequestBodyWithDefaults() *UserGroupPostRequestBody`

NewUserGroupPostRequestBodyWithDefaults instantiates a new UserGroupPostRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserGroupPostRequestBody) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserGroupPostRequestBody) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserGroupPostRequestBody) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UserGroupPostRequestBody) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *UserGroupPostRequestBody) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserGroupPostRequestBody) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserGroupPostRequestBody) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UserGroupPostRequestBody) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *UserGroupPostRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserGroupPostRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserGroupPostRequestBody) SetName(v string)`

SetName sets Name field to given value.


### GetUserDefinedFields

`func (o *UserGroupPostRequestBody) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *UserGroupPostRequestBody) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *UserGroupPostRequestBody) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *UserGroupPostRequestBody) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetGroupType

`func (o *UserGroupPostRequestBody) GetGroupType() string`

GetGroupType returns the GroupType field if non-nil, zero value otherwise.

### GetGroupTypeOk

`func (o *UserGroupPostRequestBody) GetGroupTypeOk() (*string, bool)`

GetGroupTypeOk returns a tuple with the GroupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupType

`func (o *UserGroupPostRequestBody) SetGroupType(v string)`

SetGroupType sets GroupType field to given value.

### HasGroupType

`func (o *UserGroupPostRequestBody) HasGroupType() bool`

HasGroupType returns a boolean if a field has been set.

### GetDescription

`func (o *UserGroupPostRequestBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UserGroupPostRequestBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UserGroupPostRequestBody) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UserGroupPostRequestBody) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAdministratorPrivilege

`func (o *UserGroupPostRequestBody) GetAdministratorPrivilege() bool`

GetAdministratorPrivilege returns the AdministratorPrivilege field if non-nil, zero value otherwise.

### GetAdministratorPrivilegeOk

`func (o *UserGroupPostRequestBody) GetAdministratorPrivilegeOk() (*bool, bool)`

GetAdministratorPrivilegeOk returns a tuple with the AdministratorPrivilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministratorPrivilege

`func (o *UserGroupPostRequestBody) SetAdministratorPrivilege(v bool)`

SetAdministratorPrivilege sets AdministratorPrivilege field to given value.

### HasAdministratorPrivilege

`func (o *UserGroupPostRequestBody) HasAdministratorPrivilege() bool`

HasAdministratorPrivilege returns a boolean if a field has been set.

### GetAuthenticator

`func (o *UserGroupPostRequestBody) GetAuthenticator() InlinedAuthenticator`

GetAuthenticator returns the Authenticator field if non-nil, zero value otherwise.

### GetAuthenticatorOk

`func (o *UserGroupPostRequestBody) GetAuthenticatorOk() (*InlinedAuthenticator, bool)`

GetAuthenticatorOk returns a tuple with the Authenticator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticator

`func (o *UserGroupPostRequestBody) SetAuthenticator(v InlinedAuthenticator)`

SetAuthenticator sets Authenticator field to given value.

### HasAuthenticator

`func (o *UserGroupPostRequestBody) HasAuthenticator() bool`

HasAuthenticator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


