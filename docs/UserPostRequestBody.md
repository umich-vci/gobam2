# UserPostRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The user or user group resource identifier. | [optional] 
**Type** | Pointer to **string** | The user scope resource type. | [optional] 
**Name** | **string** | The name of the user. | 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**Password** | Pointer to **string** | The password of the user. | [optional] 
**Authenticator** | Pointer to [**InlinedAuthenticator**](InlinedAuthenticator.md) |  | [optional] 
**Email** | **string** | The email address of the user. | 
**PhoneNumber** | Pointer to **string** | The phone number of the user. | [optional] 
**SecurityPrivilege** | **string** | The configured security privilege of the user. | 
**HistoryPrivilege** | Pointer to **string** | The configured history privilege of the user. | [optional] 
**AccessType** | **string** | The access type of the user. The access type can be GUI, API, or both. | 
**PasswordResetRequired** | Pointer to **bool** | Indicates whether the user must reset their password upon their initial log in to Address Manager. | [optional] 
**AccountLocked** | Pointer to **bool** | Indicates whether the user account has been locked. | [optional] 
**X509Required** | Pointer to **bool** | Indicates whether the user must access Address Manager using X.509 authentication only. | [optional] 
**AdministrativeAccessLevels** | Pointer to **map[string]string** | The administrative access right levels that the user has for the events list, log management, and reporting. | [optional] [readonly] 

## Methods

### NewUserPostRequestBody

`func NewUserPostRequestBody(name string, email string, securityPrivilege string, accessType string, ) *UserPostRequestBody`

NewUserPostRequestBody instantiates a new UserPostRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserPostRequestBodyWithDefaults

`func NewUserPostRequestBodyWithDefaults() *UserPostRequestBody`

NewUserPostRequestBodyWithDefaults instantiates a new UserPostRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserPostRequestBody) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserPostRequestBody) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserPostRequestBody) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UserPostRequestBody) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *UserPostRequestBody) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserPostRequestBody) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserPostRequestBody) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UserPostRequestBody) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *UserPostRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserPostRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserPostRequestBody) SetName(v string)`

SetName sets Name field to given value.


### GetUserDefinedFields

`func (o *UserPostRequestBody) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *UserPostRequestBody) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *UserPostRequestBody) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *UserPostRequestBody) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetPassword

`func (o *UserPostRequestBody) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UserPostRequestBody) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UserPostRequestBody) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UserPostRequestBody) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetAuthenticator

`func (o *UserPostRequestBody) GetAuthenticator() InlinedAuthenticator`

GetAuthenticator returns the Authenticator field if non-nil, zero value otherwise.

### GetAuthenticatorOk

`func (o *UserPostRequestBody) GetAuthenticatorOk() (*InlinedAuthenticator, bool)`

GetAuthenticatorOk returns a tuple with the Authenticator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticator

`func (o *UserPostRequestBody) SetAuthenticator(v InlinedAuthenticator)`

SetAuthenticator sets Authenticator field to given value.

### HasAuthenticator

`func (o *UserPostRequestBody) HasAuthenticator() bool`

HasAuthenticator returns a boolean if a field has been set.

### GetEmail

`func (o *UserPostRequestBody) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserPostRequestBody) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserPostRequestBody) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPhoneNumber

`func (o *UserPostRequestBody) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *UserPostRequestBody) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *UserPostRequestBody) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *UserPostRequestBody) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetSecurityPrivilege

`func (o *UserPostRequestBody) GetSecurityPrivilege() string`

GetSecurityPrivilege returns the SecurityPrivilege field if non-nil, zero value otherwise.

### GetSecurityPrivilegeOk

`func (o *UserPostRequestBody) GetSecurityPrivilegeOk() (*string, bool)`

GetSecurityPrivilegeOk returns a tuple with the SecurityPrivilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityPrivilege

`func (o *UserPostRequestBody) SetSecurityPrivilege(v string)`

SetSecurityPrivilege sets SecurityPrivilege field to given value.


### GetHistoryPrivilege

`func (o *UserPostRequestBody) GetHistoryPrivilege() string`

GetHistoryPrivilege returns the HistoryPrivilege field if non-nil, zero value otherwise.

### GetHistoryPrivilegeOk

`func (o *UserPostRequestBody) GetHistoryPrivilegeOk() (*string, bool)`

GetHistoryPrivilegeOk returns a tuple with the HistoryPrivilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoryPrivilege

`func (o *UserPostRequestBody) SetHistoryPrivilege(v string)`

SetHistoryPrivilege sets HistoryPrivilege field to given value.

### HasHistoryPrivilege

`func (o *UserPostRequestBody) HasHistoryPrivilege() bool`

HasHistoryPrivilege returns a boolean if a field has been set.

### GetAccessType

`func (o *UserPostRequestBody) GetAccessType() string`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *UserPostRequestBody) GetAccessTypeOk() (*string, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *UserPostRequestBody) SetAccessType(v string)`

SetAccessType sets AccessType field to given value.


### GetPasswordResetRequired

`func (o *UserPostRequestBody) GetPasswordResetRequired() bool`

GetPasswordResetRequired returns the PasswordResetRequired field if non-nil, zero value otherwise.

### GetPasswordResetRequiredOk

`func (o *UserPostRequestBody) GetPasswordResetRequiredOk() (*bool, bool)`

GetPasswordResetRequiredOk returns a tuple with the PasswordResetRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordResetRequired

`func (o *UserPostRequestBody) SetPasswordResetRequired(v bool)`

SetPasswordResetRequired sets PasswordResetRequired field to given value.

### HasPasswordResetRequired

`func (o *UserPostRequestBody) HasPasswordResetRequired() bool`

HasPasswordResetRequired returns a boolean if a field has been set.

### GetAccountLocked

`func (o *UserPostRequestBody) GetAccountLocked() bool`

GetAccountLocked returns the AccountLocked field if non-nil, zero value otherwise.

### GetAccountLockedOk

`func (o *UserPostRequestBody) GetAccountLockedOk() (*bool, bool)`

GetAccountLockedOk returns a tuple with the AccountLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountLocked

`func (o *UserPostRequestBody) SetAccountLocked(v bool)`

SetAccountLocked sets AccountLocked field to given value.

### HasAccountLocked

`func (o *UserPostRequestBody) HasAccountLocked() bool`

HasAccountLocked returns a boolean if a field has been set.

### GetX509Required

`func (o *UserPostRequestBody) GetX509Required() bool`

GetX509Required returns the X509Required field if non-nil, zero value otherwise.

### GetX509RequiredOk

`func (o *UserPostRequestBody) GetX509RequiredOk() (*bool, bool)`

GetX509RequiredOk returns a tuple with the X509Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX509Required

`func (o *UserPostRequestBody) SetX509Required(v bool)`

SetX509Required sets X509Required field to given value.

### HasX509Required

`func (o *UserPostRequestBody) HasX509Required() bool`

HasX509Required returns a boolean if a field has been set.

### GetAdministrativeAccessLevels

`func (o *UserPostRequestBody) GetAdministrativeAccessLevels() map[string]string`

GetAdministrativeAccessLevels returns the AdministrativeAccessLevels field if non-nil, zero value otherwise.

### GetAdministrativeAccessLevelsOk

`func (o *UserPostRequestBody) GetAdministrativeAccessLevelsOk() (*map[string]string, bool)`

GetAdministrativeAccessLevelsOk returns a tuple with the AdministrativeAccessLevels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrativeAccessLevels

`func (o *UserPostRequestBody) SetAdministrativeAccessLevels(v map[string]string)`

SetAdministrativeAccessLevels sets AdministrativeAccessLevels field to given value.

### HasAdministrativeAccessLevels

`func (o *UserPostRequestBody) HasAdministrativeAccessLevels() bool`

HasAdministrativeAccessLevels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


