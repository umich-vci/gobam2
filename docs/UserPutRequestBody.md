# UserPutRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The user or user group resource identifier. | [optional] 
**Type** | Pointer to **string** | The user scope resource type. | [optional] 
**Name** | Pointer to **string** | The name of the user. | [optional] 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**Password** | Pointer to **string** | The password of the user. | [optional] 
**Authenticator** | [**InlinedAuthenticator**](InlinedAuthenticator.md) |  | 
**Email** | **string** | The email address of the user. | 
**PhoneNumber** | Pointer to **string** | The phone number of the user. | [optional] 
**SecurityPrivilege** | **string** | The configured security privilege of the user. | 
**HistoryPrivilege** | **string** | The configured history privilege of the user. | 
**AccessType** | **string** | The access type of the user. The access type can be GUI, API, or both. | 
**PasswordResetRequired** | **bool** | Indicates whether the user must reset their password upon their initial log in to Address Manager. | 
**AccountLocked** | **bool** | Indicates whether the user account has been locked. | 
**X509Required** | **bool** | Indicates whether the user must access Address Manager using X.509 authentication only. | 
**AdministrativeAccessLevels** | Pointer to **map[string]string** | The administrative access right levels that the user has for the events list, log management, and reporting. | [optional] [readonly] 

## Methods

### NewUserPutRequestBody

`func NewUserPutRequestBody(authenticator InlinedAuthenticator, email string, securityPrivilege string, historyPrivilege string, accessType string, passwordResetRequired bool, accountLocked bool, x509Required bool, ) *UserPutRequestBody`

NewUserPutRequestBody instantiates a new UserPutRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserPutRequestBodyWithDefaults

`func NewUserPutRequestBodyWithDefaults() *UserPutRequestBody`

NewUserPutRequestBodyWithDefaults instantiates a new UserPutRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserPutRequestBody) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserPutRequestBody) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserPutRequestBody) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UserPutRequestBody) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *UserPutRequestBody) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserPutRequestBody) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserPutRequestBody) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UserPutRequestBody) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *UserPutRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserPutRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserPutRequestBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserPutRequestBody) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *UserPutRequestBody) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *UserPutRequestBody) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *UserPutRequestBody) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *UserPutRequestBody) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetPassword

`func (o *UserPutRequestBody) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UserPutRequestBody) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UserPutRequestBody) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UserPutRequestBody) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetAuthenticator

`func (o *UserPutRequestBody) GetAuthenticator() InlinedAuthenticator`

GetAuthenticator returns the Authenticator field if non-nil, zero value otherwise.

### GetAuthenticatorOk

`func (o *UserPutRequestBody) GetAuthenticatorOk() (*InlinedAuthenticator, bool)`

GetAuthenticatorOk returns a tuple with the Authenticator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticator

`func (o *UserPutRequestBody) SetAuthenticator(v InlinedAuthenticator)`

SetAuthenticator sets Authenticator field to given value.


### GetEmail

`func (o *UserPutRequestBody) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserPutRequestBody) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserPutRequestBody) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPhoneNumber

`func (o *UserPutRequestBody) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *UserPutRequestBody) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *UserPutRequestBody) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *UserPutRequestBody) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetSecurityPrivilege

`func (o *UserPutRequestBody) GetSecurityPrivilege() string`

GetSecurityPrivilege returns the SecurityPrivilege field if non-nil, zero value otherwise.

### GetSecurityPrivilegeOk

`func (o *UserPutRequestBody) GetSecurityPrivilegeOk() (*string, bool)`

GetSecurityPrivilegeOk returns a tuple with the SecurityPrivilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityPrivilege

`func (o *UserPutRequestBody) SetSecurityPrivilege(v string)`

SetSecurityPrivilege sets SecurityPrivilege field to given value.


### GetHistoryPrivilege

`func (o *UserPutRequestBody) GetHistoryPrivilege() string`

GetHistoryPrivilege returns the HistoryPrivilege field if non-nil, zero value otherwise.

### GetHistoryPrivilegeOk

`func (o *UserPutRequestBody) GetHistoryPrivilegeOk() (*string, bool)`

GetHistoryPrivilegeOk returns a tuple with the HistoryPrivilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoryPrivilege

`func (o *UserPutRequestBody) SetHistoryPrivilege(v string)`

SetHistoryPrivilege sets HistoryPrivilege field to given value.


### GetAccessType

`func (o *UserPutRequestBody) GetAccessType() string`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *UserPutRequestBody) GetAccessTypeOk() (*string, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *UserPutRequestBody) SetAccessType(v string)`

SetAccessType sets AccessType field to given value.


### GetPasswordResetRequired

`func (o *UserPutRequestBody) GetPasswordResetRequired() bool`

GetPasswordResetRequired returns the PasswordResetRequired field if non-nil, zero value otherwise.

### GetPasswordResetRequiredOk

`func (o *UserPutRequestBody) GetPasswordResetRequiredOk() (*bool, bool)`

GetPasswordResetRequiredOk returns a tuple with the PasswordResetRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordResetRequired

`func (o *UserPutRequestBody) SetPasswordResetRequired(v bool)`

SetPasswordResetRequired sets PasswordResetRequired field to given value.


### GetAccountLocked

`func (o *UserPutRequestBody) GetAccountLocked() bool`

GetAccountLocked returns the AccountLocked field if non-nil, zero value otherwise.

### GetAccountLockedOk

`func (o *UserPutRequestBody) GetAccountLockedOk() (*bool, bool)`

GetAccountLockedOk returns a tuple with the AccountLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountLocked

`func (o *UserPutRequestBody) SetAccountLocked(v bool)`

SetAccountLocked sets AccountLocked field to given value.


### GetX509Required

`func (o *UserPutRequestBody) GetX509Required() bool`

GetX509Required returns the X509Required field if non-nil, zero value otherwise.

### GetX509RequiredOk

`func (o *UserPutRequestBody) GetX509RequiredOk() (*bool, bool)`

GetX509RequiredOk returns a tuple with the X509Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX509Required

`func (o *UserPutRequestBody) SetX509Required(v bool)`

SetX509Required sets X509Required field to given value.


### GetAdministrativeAccessLevels

`func (o *UserPutRequestBody) GetAdministrativeAccessLevels() map[string]string`

GetAdministrativeAccessLevels returns the AdministrativeAccessLevels field if non-nil, zero value otherwise.

### GetAdministrativeAccessLevelsOk

`func (o *UserPutRequestBody) GetAdministrativeAccessLevelsOk() (*map[string]string, bool)`

GetAdministrativeAccessLevelsOk returns a tuple with the AdministrativeAccessLevels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrativeAccessLevels

`func (o *UserPutRequestBody) SetAdministrativeAccessLevels(v map[string]string)`

SetAdministrativeAccessLevels sets AdministrativeAccessLevels field to given value.

### HasAdministrativeAccessLevels

`func (o *UserPutRequestBody) HasAdministrativeAccessLevels() bool`

HasAdministrativeAccessLevels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


