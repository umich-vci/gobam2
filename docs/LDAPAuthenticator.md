# LDAPAuthenticator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of authenticator. | [optional] 
**Hostname** | Pointer to **string** | The fully qualified domain name or IP address for the authenticator. | [optional] 
**Port** | Pointer to **int32** | The TCP port used for communication between Address Manager and the default LDAP server. | [optional] 
**AdminDn** | Pointer to **string** | The distinguished name or relative distinguished name of a user with rights to search the LDAP directory. | [optional] 
**AdminPassword** | Pointer to **string** | The password for the distinguished name or relative distinguished name. | [optional] 
**SslEnabled** | Pointer to **bool** | Determines whether SSL communication is enabled between Address Manager and the LDAP server. | [optional] 
**SchemaType** | Pointer to **string** | The type of LDAP schema. | [optional] 
**SearchBase** | Pointer to **string** | The search base distinguished name location from which the search for users on the LDAP server begins. | [optional] 
**UserObjectClass** | Pointer to **string** | The user object class used to locate an LDAP user. | [optional] 
**GroupObjectClass** | Pointer to **string** | The group object class used to indicate that a DN is a group. | [optional] 
**UserPrefix** | Pointer to **string** | The user attribute for user accounts in the LDAP tree. | [optional] 
**EmailPrefix** | Pointer to **string** | The variable to be used for the email prefix. | [optional] 
**MemberOfPrefix** | Pointer to **string** | The attribute used to store user-group membership information. | [optional] 
**ReferralPolicy** | Pointer to **string** | The environment property that indicates how to handle referrals from external resources to the service providers. | [optional] 
**AliasDereferencePolicy** | Pointer to **string** | The environment property that indicates whether alias entries are dereferenced. | [optional] 

## Methods

### NewLDAPAuthenticator

`func NewLDAPAuthenticator() *LDAPAuthenticator`

NewLDAPAuthenticator instantiates a new LDAPAuthenticator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLDAPAuthenticatorWithDefaults

`func NewLDAPAuthenticatorWithDefaults() *LDAPAuthenticator`

NewLDAPAuthenticatorWithDefaults instantiates a new LDAPAuthenticator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *LDAPAuthenticator) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LDAPAuthenticator) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LDAPAuthenticator) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *LDAPAuthenticator) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHostname

`func (o *LDAPAuthenticator) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *LDAPAuthenticator) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *LDAPAuthenticator) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *LDAPAuthenticator) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetPort

`func (o *LDAPAuthenticator) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *LDAPAuthenticator) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *LDAPAuthenticator) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *LDAPAuthenticator) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetAdminDn

`func (o *LDAPAuthenticator) GetAdminDn() string`

GetAdminDn returns the AdminDn field if non-nil, zero value otherwise.

### GetAdminDnOk

`func (o *LDAPAuthenticator) GetAdminDnOk() (*string, bool)`

GetAdminDnOk returns a tuple with the AdminDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminDn

`func (o *LDAPAuthenticator) SetAdminDn(v string)`

SetAdminDn sets AdminDn field to given value.

### HasAdminDn

`func (o *LDAPAuthenticator) HasAdminDn() bool`

HasAdminDn returns a boolean if a field has been set.

### GetAdminPassword

`func (o *LDAPAuthenticator) GetAdminPassword() string`

GetAdminPassword returns the AdminPassword field if non-nil, zero value otherwise.

### GetAdminPasswordOk

`func (o *LDAPAuthenticator) GetAdminPasswordOk() (*string, bool)`

GetAdminPasswordOk returns a tuple with the AdminPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminPassword

`func (o *LDAPAuthenticator) SetAdminPassword(v string)`

SetAdminPassword sets AdminPassword field to given value.

### HasAdminPassword

`func (o *LDAPAuthenticator) HasAdminPassword() bool`

HasAdminPassword returns a boolean if a field has been set.

### GetSslEnabled

`func (o *LDAPAuthenticator) GetSslEnabled() bool`

GetSslEnabled returns the SslEnabled field if non-nil, zero value otherwise.

### GetSslEnabledOk

`func (o *LDAPAuthenticator) GetSslEnabledOk() (*bool, bool)`

GetSslEnabledOk returns a tuple with the SslEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslEnabled

`func (o *LDAPAuthenticator) SetSslEnabled(v bool)`

SetSslEnabled sets SslEnabled field to given value.

### HasSslEnabled

`func (o *LDAPAuthenticator) HasSslEnabled() bool`

HasSslEnabled returns a boolean if a field has been set.

### GetSchemaType

`func (o *LDAPAuthenticator) GetSchemaType() string`

GetSchemaType returns the SchemaType field if non-nil, zero value otherwise.

### GetSchemaTypeOk

`func (o *LDAPAuthenticator) GetSchemaTypeOk() (*string, bool)`

GetSchemaTypeOk returns a tuple with the SchemaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaType

`func (o *LDAPAuthenticator) SetSchemaType(v string)`

SetSchemaType sets SchemaType field to given value.

### HasSchemaType

`func (o *LDAPAuthenticator) HasSchemaType() bool`

HasSchemaType returns a boolean if a field has been set.

### GetSearchBase

`func (o *LDAPAuthenticator) GetSearchBase() string`

GetSearchBase returns the SearchBase field if non-nil, zero value otherwise.

### GetSearchBaseOk

`func (o *LDAPAuthenticator) GetSearchBaseOk() (*string, bool)`

GetSearchBaseOk returns a tuple with the SearchBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchBase

`func (o *LDAPAuthenticator) SetSearchBase(v string)`

SetSearchBase sets SearchBase field to given value.

### HasSearchBase

`func (o *LDAPAuthenticator) HasSearchBase() bool`

HasSearchBase returns a boolean if a field has been set.

### GetUserObjectClass

`func (o *LDAPAuthenticator) GetUserObjectClass() string`

GetUserObjectClass returns the UserObjectClass field if non-nil, zero value otherwise.

### GetUserObjectClassOk

`func (o *LDAPAuthenticator) GetUserObjectClassOk() (*string, bool)`

GetUserObjectClassOk returns a tuple with the UserObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserObjectClass

`func (o *LDAPAuthenticator) SetUserObjectClass(v string)`

SetUserObjectClass sets UserObjectClass field to given value.

### HasUserObjectClass

`func (o *LDAPAuthenticator) HasUserObjectClass() bool`

HasUserObjectClass returns a boolean if a field has been set.

### GetGroupObjectClass

`func (o *LDAPAuthenticator) GetGroupObjectClass() string`

GetGroupObjectClass returns the GroupObjectClass field if non-nil, zero value otherwise.

### GetGroupObjectClassOk

`func (o *LDAPAuthenticator) GetGroupObjectClassOk() (*string, bool)`

GetGroupObjectClassOk returns a tuple with the GroupObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupObjectClass

`func (o *LDAPAuthenticator) SetGroupObjectClass(v string)`

SetGroupObjectClass sets GroupObjectClass field to given value.

### HasGroupObjectClass

`func (o *LDAPAuthenticator) HasGroupObjectClass() bool`

HasGroupObjectClass returns a boolean if a field has been set.

### GetUserPrefix

`func (o *LDAPAuthenticator) GetUserPrefix() string`

GetUserPrefix returns the UserPrefix field if non-nil, zero value otherwise.

### GetUserPrefixOk

`func (o *LDAPAuthenticator) GetUserPrefixOk() (*string, bool)`

GetUserPrefixOk returns a tuple with the UserPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPrefix

`func (o *LDAPAuthenticator) SetUserPrefix(v string)`

SetUserPrefix sets UserPrefix field to given value.

### HasUserPrefix

`func (o *LDAPAuthenticator) HasUserPrefix() bool`

HasUserPrefix returns a boolean if a field has been set.

### GetEmailPrefix

`func (o *LDAPAuthenticator) GetEmailPrefix() string`

GetEmailPrefix returns the EmailPrefix field if non-nil, zero value otherwise.

### GetEmailPrefixOk

`func (o *LDAPAuthenticator) GetEmailPrefixOk() (*string, bool)`

GetEmailPrefixOk returns a tuple with the EmailPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailPrefix

`func (o *LDAPAuthenticator) SetEmailPrefix(v string)`

SetEmailPrefix sets EmailPrefix field to given value.

### HasEmailPrefix

`func (o *LDAPAuthenticator) HasEmailPrefix() bool`

HasEmailPrefix returns a boolean if a field has been set.

### GetMemberOfPrefix

`func (o *LDAPAuthenticator) GetMemberOfPrefix() string`

GetMemberOfPrefix returns the MemberOfPrefix field if non-nil, zero value otherwise.

### GetMemberOfPrefixOk

`func (o *LDAPAuthenticator) GetMemberOfPrefixOk() (*string, bool)`

GetMemberOfPrefixOk returns a tuple with the MemberOfPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberOfPrefix

`func (o *LDAPAuthenticator) SetMemberOfPrefix(v string)`

SetMemberOfPrefix sets MemberOfPrefix field to given value.

### HasMemberOfPrefix

`func (o *LDAPAuthenticator) HasMemberOfPrefix() bool`

HasMemberOfPrefix returns a boolean if a field has been set.

### GetReferralPolicy

`func (o *LDAPAuthenticator) GetReferralPolicy() string`

GetReferralPolicy returns the ReferralPolicy field if non-nil, zero value otherwise.

### GetReferralPolicyOk

`func (o *LDAPAuthenticator) GetReferralPolicyOk() (*string, bool)`

GetReferralPolicyOk returns a tuple with the ReferralPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferralPolicy

`func (o *LDAPAuthenticator) SetReferralPolicy(v string)`

SetReferralPolicy sets ReferralPolicy field to given value.

### HasReferralPolicy

`func (o *LDAPAuthenticator) HasReferralPolicy() bool`

HasReferralPolicy returns a boolean if a field has been set.

### GetAliasDereferencePolicy

`func (o *LDAPAuthenticator) GetAliasDereferencePolicy() string`

GetAliasDereferencePolicy returns the AliasDereferencePolicy field if non-nil, zero value otherwise.

### GetAliasDereferencePolicyOk

`func (o *LDAPAuthenticator) GetAliasDereferencePolicyOk() (*string, bool)`

GetAliasDereferencePolicyOk returns a tuple with the AliasDereferencePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasDereferencePolicy

`func (o *LDAPAuthenticator) SetAliasDereferencePolicy(v string)`

SetAliasDereferencePolicy sets AliasDereferencePolicy field to given value.

### HasAliasDereferencePolicy

`func (o *LDAPAuthenticator) HasAliasDereferencePolicy() bool`

HasAliasDereferencePolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


