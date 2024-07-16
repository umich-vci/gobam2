# GetGroups200ResponseDataInner

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
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 

## Methods

### NewGetGroups200ResponseDataInner

`func NewGetGroups200ResponseDataInner() *GetGroups200ResponseDataInner`

NewGetGroups200ResponseDataInner instantiates a new GetGroups200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGroups200ResponseDataInnerWithDefaults

`func NewGetGroups200ResponseDataInnerWithDefaults() *GetGroups200ResponseDataInner`

NewGetGroups200ResponseDataInnerWithDefaults instantiates a new GetGroups200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetGroups200ResponseDataInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetGroups200ResponseDataInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetGroups200ResponseDataInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetGroups200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GetGroups200ResponseDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetGroups200ResponseDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetGroups200ResponseDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetGroups200ResponseDataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *GetGroups200ResponseDataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetGroups200ResponseDataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetGroups200ResponseDataInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetGroups200ResponseDataInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *GetGroups200ResponseDataInner) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *GetGroups200ResponseDataInner) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *GetGroups200ResponseDataInner) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *GetGroups200ResponseDataInner) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetGroupType

`func (o *GetGroups200ResponseDataInner) GetGroupType() string`

GetGroupType returns the GroupType field if non-nil, zero value otherwise.

### GetGroupTypeOk

`func (o *GetGroups200ResponseDataInner) GetGroupTypeOk() (*string, bool)`

GetGroupTypeOk returns a tuple with the GroupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupType

`func (o *GetGroups200ResponseDataInner) SetGroupType(v string)`

SetGroupType sets GroupType field to given value.

### HasGroupType

`func (o *GetGroups200ResponseDataInner) HasGroupType() bool`

HasGroupType returns a boolean if a field has been set.

### GetDescription

`func (o *GetGroups200ResponseDataInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetGroups200ResponseDataInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetGroups200ResponseDataInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetGroups200ResponseDataInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAdministratorPrivilege

`func (o *GetGroups200ResponseDataInner) GetAdministratorPrivilege() bool`

GetAdministratorPrivilege returns the AdministratorPrivilege field if non-nil, zero value otherwise.

### GetAdministratorPrivilegeOk

`func (o *GetGroups200ResponseDataInner) GetAdministratorPrivilegeOk() (*bool, bool)`

GetAdministratorPrivilegeOk returns a tuple with the AdministratorPrivilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministratorPrivilege

`func (o *GetGroups200ResponseDataInner) SetAdministratorPrivilege(v bool)`

SetAdministratorPrivilege sets AdministratorPrivilege field to given value.

### HasAdministratorPrivilege

`func (o *GetGroups200ResponseDataInner) HasAdministratorPrivilege() bool`

HasAdministratorPrivilege returns a boolean if a field has been set.

### GetAuthenticator

`func (o *GetGroups200ResponseDataInner) GetAuthenticator() InlinedAuthenticator`

GetAuthenticator returns the Authenticator field if non-nil, zero value otherwise.

### GetAuthenticatorOk

`func (o *GetGroups200ResponseDataInner) GetAuthenticatorOk() (*InlinedAuthenticator, bool)`

GetAuthenticatorOk returns a tuple with the Authenticator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticator

`func (o *GetGroups200ResponseDataInner) SetAuthenticator(v InlinedAuthenticator)`

SetAuthenticator sets Authenticator field to given value.

### HasAuthenticator

`func (o *GetGroups200ResponseDataInner) HasAuthenticator() bool`

HasAuthenticator returns a boolean if a field has been set.

### GetLinks

`func (o *GetGroups200ResponseDataInner) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetGroups200ResponseDataInner) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetGroups200ResponseDataInner) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetGroups200ResponseDataInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

