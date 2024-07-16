# UserGroupPutRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | **string** | The name for the user, LDAP, TACACS+, or SSO group. | 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**GroupType** | **string** |  | 
**Description** | Pointer to **string** | The description of the SSO group. | [optional] 
**AdministratorPrivilege** | **bool** | Indicates whether the users within the group have administrative user privileges. | 
**Authenticator** | Pointer to [**InlinedAuthenticator**](InlinedAuthenticator.md) |  | [optional] 

## Methods

### NewUserGroupPutRequestBody

`func NewUserGroupPutRequestBody(name string, groupType string, administratorPrivilege bool, ) *UserGroupPutRequestBody`

NewUserGroupPutRequestBody instantiates a new UserGroupPutRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserGroupPutRequestBodyWithDefaults

`func NewUserGroupPutRequestBodyWithDefaults() *UserGroupPutRequestBody`

NewUserGroupPutRequestBodyWithDefaults instantiates a new UserGroupPutRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserGroupPutRequestBody) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserGroupPutRequestBody) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserGroupPutRequestBody) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UserGroupPutRequestBody) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *UserGroupPutRequestBody) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserGroupPutRequestBody) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserGroupPutRequestBody) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UserGroupPutRequestBody) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *UserGroupPutRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserGroupPutRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserGroupPutRequestBody) SetName(v string)`

SetName sets Name field to given value.


### GetUserDefinedFields

`func (o *UserGroupPutRequestBody) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *UserGroupPutRequestBody) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *UserGroupPutRequestBody) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *UserGroupPutRequestBody) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetGroupType

`func (o *UserGroupPutRequestBody) GetGroupType() string`

GetGroupType returns the GroupType field if non-nil, zero value otherwise.

### GetGroupTypeOk

`func (o *UserGroupPutRequestBody) GetGroupTypeOk() (*string, bool)`

GetGroupTypeOk returns a tuple with the GroupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupType

`func (o *UserGroupPutRequestBody) SetGroupType(v string)`

SetGroupType sets GroupType field to given value.


### GetDescription

`func (o *UserGroupPutRequestBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UserGroupPutRequestBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UserGroupPutRequestBody) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UserGroupPutRequestBody) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAdministratorPrivilege

`func (o *UserGroupPutRequestBody) GetAdministratorPrivilege() bool`

GetAdministratorPrivilege returns the AdministratorPrivilege field if non-nil, zero value otherwise.

### GetAdministratorPrivilegeOk

`func (o *UserGroupPutRequestBody) GetAdministratorPrivilegeOk() (*bool, bool)`

GetAdministratorPrivilegeOk returns a tuple with the AdministratorPrivilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministratorPrivilege

`func (o *UserGroupPutRequestBody) SetAdministratorPrivilege(v bool)`

SetAdministratorPrivilege sets AdministratorPrivilege field to given value.


### GetAuthenticator

`func (o *UserGroupPutRequestBody) GetAuthenticator() InlinedAuthenticator`

GetAuthenticator returns the Authenticator field if non-nil, zero value otherwise.

### GetAuthenticatorOk

`func (o *UserGroupPutRequestBody) GetAuthenticatorOk() (*InlinedAuthenticator, bool)`

GetAuthenticatorOk returns a tuple with the Authenticator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticator

`func (o *UserGroupPutRequestBody) SetAuthenticator(v InlinedAuthenticator)`

SetAuthenticator sets Authenticator field to given value.

### HasAuthenticator

`func (o *UserGroupPutRequestBody) HasAuthenticator() bool`

HasAuthenticator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


