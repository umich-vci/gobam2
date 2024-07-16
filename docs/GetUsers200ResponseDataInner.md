# GetUsers200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The user or user group resource identifier. | [optional] 
**Type** | Pointer to **string** | The user scope resource type. | [optional] 
**Name** | Pointer to **string** | The name of the user. | [optional] 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**Password** | Pointer to **string** | The password of the user. | [optional] 
**Authenticator** | Pointer to [**InlinedAuthenticator**](InlinedAuthenticator.md) |  | [optional] 
**Email** | Pointer to **string** | The email address of the user. | [optional] 
**PhoneNumber** | Pointer to **string** | The phone number of the user. | [optional] 
**SecurityPrivilege** | Pointer to **string** | The configured security privilege of the user. | [optional] 
**HistoryPrivilege** | Pointer to **string** | The configured history privilege of the user. | [optional] 
**AccessType** | Pointer to **string** | The access type of the user. The access type can be GUI, API, or both. | [optional] 
**PasswordResetRequired** | Pointer to **bool** | Indicates whether the user must reset their password upon their initial log in to Address Manager. | [optional] 
**AccountLocked** | Pointer to **bool** | Indicates whether the user account has been locked. | [optional] 
**X509Required** | Pointer to **bool** | Indicates whether the user must access Address Manager using X.509 authentication only. | [optional] 
**AdministrativeAccessLevels** | Pointer to **map[string]string** | The administrative access right levels that the user has for the events list, log management, and reporting. | [optional] [readonly] 
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 

## Methods

### NewGetUsers200ResponseDataInner

`func NewGetUsers200ResponseDataInner() *GetUsers200ResponseDataInner`

NewGetUsers200ResponseDataInner instantiates a new GetUsers200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUsers200ResponseDataInnerWithDefaults

`func NewGetUsers200ResponseDataInnerWithDefaults() *GetUsers200ResponseDataInner`

NewGetUsers200ResponseDataInnerWithDefaults instantiates a new GetUsers200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetUsers200ResponseDataInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetUsers200ResponseDataInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetUsers200ResponseDataInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetUsers200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GetUsers200ResponseDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetUsers200ResponseDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetUsers200ResponseDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetUsers200ResponseDataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *GetUsers200ResponseDataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetUsers200ResponseDataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetUsers200ResponseDataInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetUsers200ResponseDataInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *GetUsers200ResponseDataInner) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *GetUsers200ResponseDataInner) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *GetUsers200ResponseDataInner) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *GetUsers200ResponseDataInner) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetPassword

`func (o *GetUsers200ResponseDataInner) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *GetUsers200ResponseDataInner) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *GetUsers200ResponseDataInner) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *GetUsers200ResponseDataInner) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetAuthenticator

`func (o *GetUsers200ResponseDataInner) GetAuthenticator() InlinedAuthenticator`

GetAuthenticator returns the Authenticator field if non-nil, zero value otherwise.

### GetAuthenticatorOk

`func (o *GetUsers200ResponseDataInner) GetAuthenticatorOk() (*InlinedAuthenticator, bool)`

GetAuthenticatorOk returns a tuple with the Authenticator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticator

`func (o *GetUsers200ResponseDataInner) SetAuthenticator(v InlinedAuthenticator)`

SetAuthenticator sets Authenticator field to given value.

### HasAuthenticator

`func (o *GetUsers200ResponseDataInner) HasAuthenticator() bool`

HasAuthenticator returns a boolean if a field has been set.

### GetEmail

`func (o *GetUsers200ResponseDataInner) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GetUsers200ResponseDataInner) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GetUsers200ResponseDataInner) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *GetUsers200ResponseDataInner) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *GetUsers200ResponseDataInner) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *GetUsers200ResponseDataInner) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *GetUsers200ResponseDataInner) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *GetUsers200ResponseDataInner) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetSecurityPrivilege

`func (o *GetUsers200ResponseDataInner) GetSecurityPrivilege() string`

GetSecurityPrivilege returns the SecurityPrivilege field if non-nil, zero value otherwise.

### GetSecurityPrivilegeOk

`func (o *GetUsers200ResponseDataInner) GetSecurityPrivilegeOk() (*string, bool)`

GetSecurityPrivilegeOk returns a tuple with the SecurityPrivilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityPrivilege

`func (o *GetUsers200ResponseDataInner) SetSecurityPrivilege(v string)`

SetSecurityPrivilege sets SecurityPrivilege field to given value.

### HasSecurityPrivilege

`func (o *GetUsers200ResponseDataInner) HasSecurityPrivilege() bool`

HasSecurityPrivilege returns a boolean if a field has been set.

### GetHistoryPrivilege

`func (o *GetUsers200ResponseDataInner) GetHistoryPrivilege() string`

GetHistoryPrivilege returns the HistoryPrivilege field if non-nil, zero value otherwise.

### GetHistoryPrivilegeOk

`func (o *GetUsers200ResponseDataInner) GetHistoryPrivilegeOk() (*string, bool)`

GetHistoryPrivilegeOk returns a tuple with the HistoryPrivilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoryPrivilege

`func (o *GetUsers200ResponseDataInner) SetHistoryPrivilege(v string)`

SetHistoryPrivilege sets HistoryPrivilege field to given value.

### HasHistoryPrivilege

`func (o *GetUsers200ResponseDataInner) HasHistoryPrivilege() bool`

HasHistoryPrivilege returns a boolean if a field has been set.

### GetAccessType

`func (o *GetUsers200ResponseDataInner) GetAccessType() string`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *GetUsers200ResponseDataInner) GetAccessTypeOk() (*string, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *GetUsers200ResponseDataInner) SetAccessType(v string)`

SetAccessType sets AccessType field to given value.

### HasAccessType

`func (o *GetUsers200ResponseDataInner) HasAccessType() bool`

HasAccessType returns a boolean if a field has been set.

### GetPasswordResetRequired

`func (o *GetUsers200ResponseDataInner) GetPasswordResetRequired() bool`

GetPasswordResetRequired returns the PasswordResetRequired field if non-nil, zero value otherwise.

### GetPasswordResetRequiredOk

`func (o *GetUsers200ResponseDataInner) GetPasswordResetRequiredOk() (*bool, bool)`

GetPasswordResetRequiredOk returns a tuple with the PasswordResetRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordResetRequired

`func (o *GetUsers200ResponseDataInner) SetPasswordResetRequired(v bool)`

SetPasswordResetRequired sets PasswordResetRequired field to given value.

### HasPasswordResetRequired

`func (o *GetUsers200ResponseDataInner) HasPasswordResetRequired() bool`

HasPasswordResetRequired returns a boolean if a field has been set.

### GetAccountLocked

`func (o *GetUsers200ResponseDataInner) GetAccountLocked() bool`

GetAccountLocked returns the AccountLocked field if non-nil, zero value otherwise.

### GetAccountLockedOk

`func (o *GetUsers200ResponseDataInner) GetAccountLockedOk() (*bool, bool)`

GetAccountLockedOk returns a tuple with the AccountLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountLocked

`func (o *GetUsers200ResponseDataInner) SetAccountLocked(v bool)`

SetAccountLocked sets AccountLocked field to given value.

### HasAccountLocked

`func (o *GetUsers200ResponseDataInner) HasAccountLocked() bool`

HasAccountLocked returns a boolean if a field has been set.

### GetX509Required

`func (o *GetUsers200ResponseDataInner) GetX509Required() bool`

GetX509Required returns the X509Required field if non-nil, zero value otherwise.

### GetX509RequiredOk

`func (o *GetUsers200ResponseDataInner) GetX509RequiredOk() (*bool, bool)`

GetX509RequiredOk returns a tuple with the X509Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX509Required

`func (o *GetUsers200ResponseDataInner) SetX509Required(v bool)`

SetX509Required sets X509Required field to given value.

### HasX509Required

`func (o *GetUsers200ResponseDataInner) HasX509Required() bool`

HasX509Required returns a boolean if a field has been set.

### GetAdministrativeAccessLevels

`func (o *GetUsers200ResponseDataInner) GetAdministrativeAccessLevels() map[string]string`

GetAdministrativeAccessLevels returns the AdministrativeAccessLevels field if non-nil, zero value otherwise.

### GetAdministrativeAccessLevelsOk

`func (o *GetUsers200ResponseDataInner) GetAdministrativeAccessLevelsOk() (*map[string]string, bool)`

GetAdministrativeAccessLevelsOk returns a tuple with the AdministrativeAccessLevels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrativeAccessLevels

`func (o *GetUsers200ResponseDataInner) SetAdministrativeAccessLevels(v map[string]string)`

SetAdministrativeAccessLevels sets AdministrativeAccessLevels field to given value.

### HasAdministrativeAccessLevels

`func (o *GetUsers200ResponseDataInner) HasAdministrativeAccessLevels() bool`

HasAdministrativeAccessLevels returns a boolean if a field has been set.

### GetLinks

`func (o *GetUsers200ResponseDataInner) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetUsers200ResponseDataInner) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetUsers200ResponseDataInner) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetUsers200ResponseDataInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


