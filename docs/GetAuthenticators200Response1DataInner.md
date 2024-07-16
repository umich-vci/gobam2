# GetAuthenticators200Response1DataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The type of authenticator. | [optional] 
**Name** | Pointer to **string** | The name for the authenticator. | [optional] 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**SecondaryAuthenticator** | Pointer to [**InlinedAuthenticator**](InlinedAuthenticator.md) |  | [optional] 
**Hostname** | Pointer to **string** | The fully qualified domain name or IP address for the authenticator. | [optional] 
**Realm** | Pointer to **string** | The administrative domain for the Kerberos server. | [optional] 
**Port** | Pointer to **int32** | The TCP port number used for communication between the client and server. | [optional] 
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
**Description** | Pointer to **string** | The description of the SAML IdP configuration. | [optional] 
**Enabled** | Pointer to **bool** | Indicates whether SAML IdP authentication is enabled. | [optional] 
**UserClaim** | Pointer to **string** | The user claim name of the authorization server. | [optional] 
**GroupClaim** | Pointer to **string** | The group claim name of the authorization server. | [optional] 
**EmailClaim** | Pointer to **string** | The email claim name of the authorization server. | [optional] 
**ValidationMethod** | Pointer to **string** | The method that indicates where the token validation occurs. | [optional] 
**ClientId** | Pointer to **string** | The public identifier of the application. | [optional] 
**ClientSecret** | Pointer to **string** | The client secret known only to the application and authorization server. | [optional] 
**IntrospectionUrl** | Pointer to **string** | The introspection endpoint that allows Address Manager to check the validity of access tokens. | [optional] 
**UserInfoUrl** | Pointer to **string** | The endpoint containing information about the user, including the group membership information and user ID. | [optional] 
**AuthorizationOption** | Pointer to **string** | The authorization method used when Address Manager sends the client ID and client secret to the authorization server. | [optional] 
**Audience** | Pointer to **string** | The name of the BAM API bearer token obtained from the authorization server. | [optional] 
**Issuer** | Pointer to **string** | The name of the token issuer. | [optional] 
**SigningCertificate** | Pointer to **string** | The SAML IdP server signing certificate. | [optional] 
**Timeout** | Pointer to **string** | The value that overrides the timeout setting used for authentication requests sent to the TACACS+ server. | [optional] 
**Retries** | Pointer to **int32** | The number of times Address Manager attempts to retransmit a failed authentication request sent to the RADIUS server. | [optional] 
**SharedSecret** | Pointer to **string** | The shared secret used to encrypt and decrypt packets between the client and the server. | [optional] 
**AuthenticationProtocol** | Pointer to **string** | The authentication protocol. | [optional] 
**GroupAttribute** | Pointer to **string** | The special attribute used for the custom service in the TACACS+ server. This attribute is used to get the group name defined in the TACACS+ server. | [optional] 
**EmailAttribute** | Pointer to **string** | The attribute name for email in the SAML response. | [optional] 
**SignOnUrl** | Pointer to **string** | The sign-in URL of the SAMP IdP. | [optional] 
**LogoutUrl** | Pointer to **string** | The logout URL of the SAMP IdP. | [optional] 
**EntityId** | Pointer to **string** | The entity ID of the SAMP IdP. | [optional] 
**Attributes** | Pointer to [**[]NameValuePairBean**](NameValuePairBean.md) |  | [optional] 

## Methods

### NewGetAuthenticators200Response1DataInner

`func NewGetAuthenticators200Response1DataInner() *GetAuthenticators200Response1DataInner`

NewGetAuthenticators200Response1DataInner instantiates a new GetAuthenticators200Response1DataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAuthenticators200Response1DataInnerWithDefaults

`func NewGetAuthenticators200Response1DataInnerWithDefaults() *GetAuthenticators200Response1DataInner`

NewGetAuthenticators200Response1DataInnerWithDefaults instantiates a new GetAuthenticators200Response1DataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetAuthenticators200Response1DataInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetAuthenticators200Response1DataInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetAuthenticators200Response1DataInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetAuthenticators200Response1DataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GetAuthenticators200Response1DataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetAuthenticators200Response1DataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetAuthenticators200Response1DataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetAuthenticators200Response1DataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *GetAuthenticators200Response1DataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetAuthenticators200Response1DataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetAuthenticators200Response1DataInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetAuthenticators200Response1DataInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *GetAuthenticators200Response1DataInner) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *GetAuthenticators200Response1DataInner) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *GetAuthenticators200Response1DataInner) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *GetAuthenticators200Response1DataInner) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetSecondaryAuthenticator

`func (o *GetAuthenticators200Response1DataInner) GetSecondaryAuthenticator() InlinedAuthenticator`

GetSecondaryAuthenticator returns the SecondaryAuthenticator field if non-nil, zero value otherwise.

### GetSecondaryAuthenticatorOk

`func (o *GetAuthenticators200Response1DataInner) GetSecondaryAuthenticatorOk() (*InlinedAuthenticator, bool)`

GetSecondaryAuthenticatorOk returns a tuple with the SecondaryAuthenticator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryAuthenticator

`func (o *GetAuthenticators200Response1DataInner) SetSecondaryAuthenticator(v InlinedAuthenticator)`

SetSecondaryAuthenticator sets SecondaryAuthenticator field to given value.

### HasSecondaryAuthenticator

`func (o *GetAuthenticators200Response1DataInner) HasSecondaryAuthenticator() bool`

HasSecondaryAuthenticator returns a boolean if a field has been set.

### GetHostname

`func (o *GetAuthenticators200Response1DataInner) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *GetAuthenticators200Response1DataInner) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *GetAuthenticators200Response1DataInner) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *GetAuthenticators200Response1DataInner) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetRealm

`func (o *GetAuthenticators200Response1DataInner) GetRealm() string`

GetRealm returns the Realm field if non-nil, zero value otherwise.

### GetRealmOk

`func (o *GetAuthenticators200Response1DataInner) GetRealmOk() (*string, bool)`

GetRealmOk returns a tuple with the Realm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealm

`func (o *GetAuthenticators200Response1DataInner) SetRealm(v string)`

SetRealm sets Realm field to given value.

### HasRealm

`func (o *GetAuthenticators200Response1DataInner) HasRealm() bool`

HasRealm returns a boolean if a field has been set.

### GetPort

`func (o *GetAuthenticators200Response1DataInner) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *GetAuthenticators200Response1DataInner) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *GetAuthenticators200Response1DataInner) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *GetAuthenticators200Response1DataInner) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetAdminDn

`func (o *GetAuthenticators200Response1DataInner) GetAdminDn() string`

GetAdminDn returns the AdminDn field if non-nil, zero value otherwise.

### GetAdminDnOk

`func (o *GetAuthenticators200Response1DataInner) GetAdminDnOk() (*string, bool)`

GetAdminDnOk returns a tuple with the AdminDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminDn

`func (o *GetAuthenticators200Response1DataInner) SetAdminDn(v string)`

SetAdminDn sets AdminDn field to given value.

### HasAdminDn

`func (o *GetAuthenticators200Response1DataInner) HasAdminDn() bool`

HasAdminDn returns a boolean if a field has been set.

### GetAdminPassword

`func (o *GetAuthenticators200Response1DataInner) GetAdminPassword() string`

GetAdminPassword returns the AdminPassword field if non-nil, zero value otherwise.

### GetAdminPasswordOk

`func (o *GetAuthenticators200Response1DataInner) GetAdminPasswordOk() (*string, bool)`

GetAdminPasswordOk returns a tuple with the AdminPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminPassword

`func (o *GetAuthenticators200Response1DataInner) SetAdminPassword(v string)`

SetAdminPassword sets AdminPassword field to given value.

### HasAdminPassword

`func (o *GetAuthenticators200Response1DataInner) HasAdminPassword() bool`

HasAdminPassword returns a boolean if a field has been set.

### GetSslEnabled

`func (o *GetAuthenticators200Response1DataInner) GetSslEnabled() bool`

GetSslEnabled returns the SslEnabled field if non-nil, zero value otherwise.

### GetSslEnabledOk

`func (o *GetAuthenticators200Response1DataInner) GetSslEnabledOk() (*bool, bool)`

GetSslEnabledOk returns a tuple with the SslEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslEnabled

`func (o *GetAuthenticators200Response1DataInner) SetSslEnabled(v bool)`

SetSslEnabled sets SslEnabled field to given value.

### HasSslEnabled

`func (o *GetAuthenticators200Response1DataInner) HasSslEnabled() bool`

HasSslEnabled returns a boolean if a field has been set.

### GetSchemaType

`func (o *GetAuthenticators200Response1DataInner) GetSchemaType() string`

GetSchemaType returns the SchemaType field if non-nil, zero value otherwise.

### GetSchemaTypeOk

`func (o *GetAuthenticators200Response1DataInner) GetSchemaTypeOk() (*string, bool)`

GetSchemaTypeOk returns a tuple with the SchemaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaType

`func (o *GetAuthenticators200Response1DataInner) SetSchemaType(v string)`

SetSchemaType sets SchemaType field to given value.

### HasSchemaType

`func (o *GetAuthenticators200Response1DataInner) HasSchemaType() bool`

HasSchemaType returns a boolean if a field has been set.

### GetSearchBase

`func (o *GetAuthenticators200Response1DataInner) GetSearchBase() string`

GetSearchBase returns the SearchBase field if non-nil, zero value otherwise.

### GetSearchBaseOk

`func (o *GetAuthenticators200Response1DataInner) GetSearchBaseOk() (*string, bool)`

GetSearchBaseOk returns a tuple with the SearchBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchBase

`func (o *GetAuthenticators200Response1DataInner) SetSearchBase(v string)`

SetSearchBase sets SearchBase field to given value.

### HasSearchBase

`func (o *GetAuthenticators200Response1DataInner) HasSearchBase() bool`

HasSearchBase returns a boolean if a field has been set.

### GetUserObjectClass

`func (o *GetAuthenticators200Response1DataInner) GetUserObjectClass() string`

GetUserObjectClass returns the UserObjectClass field if non-nil, zero value otherwise.

### GetUserObjectClassOk

`func (o *GetAuthenticators200Response1DataInner) GetUserObjectClassOk() (*string, bool)`

GetUserObjectClassOk returns a tuple with the UserObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserObjectClass

`func (o *GetAuthenticators200Response1DataInner) SetUserObjectClass(v string)`

SetUserObjectClass sets UserObjectClass field to given value.

### HasUserObjectClass

`func (o *GetAuthenticators200Response1DataInner) HasUserObjectClass() bool`

HasUserObjectClass returns a boolean if a field has been set.

### GetGroupObjectClass

`func (o *GetAuthenticators200Response1DataInner) GetGroupObjectClass() string`

GetGroupObjectClass returns the GroupObjectClass field if non-nil, zero value otherwise.

### GetGroupObjectClassOk

`func (o *GetAuthenticators200Response1DataInner) GetGroupObjectClassOk() (*string, bool)`

GetGroupObjectClassOk returns a tuple with the GroupObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupObjectClass

`func (o *GetAuthenticators200Response1DataInner) SetGroupObjectClass(v string)`

SetGroupObjectClass sets GroupObjectClass field to given value.

### HasGroupObjectClass

`func (o *GetAuthenticators200Response1DataInner) HasGroupObjectClass() bool`

HasGroupObjectClass returns a boolean if a field has been set.

### GetUserPrefix

`func (o *GetAuthenticators200Response1DataInner) GetUserPrefix() string`

GetUserPrefix returns the UserPrefix field if non-nil, zero value otherwise.

### GetUserPrefixOk

`func (o *GetAuthenticators200Response1DataInner) GetUserPrefixOk() (*string, bool)`

GetUserPrefixOk returns a tuple with the UserPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPrefix

`func (o *GetAuthenticators200Response1DataInner) SetUserPrefix(v string)`

SetUserPrefix sets UserPrefix field to given value.

### HasUserPrefix

`func (o *GetAuthenticators200Response1DataInner) HasUserPrefix() bool`

HasUserPrefix returns a boolean if a field has been set.

### GetEmailPrefix

`func (o *GetAuthenticators200Response1DataInner) GetEmailPrefix() string`

GetEmailPrefix returns the EmailPrefix field if non-nil, zero value otherwise.

### GetEmailPrefixOk

`func (o *GetAuthenticators200Response1DataInner) GetEmailPrefixOk() (*string, bool)`

GetEmailPrefixOk returns a tuple with the EmailPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailPrefix

`func (o *GetAuthenticators200Response1DataInner) SetEmailPrefix(v string)`

SetEmailPrefix sets EmailPrefix field to given value.

### HasEmailPrefix

`func (o *GetAuthenticators200Response1DataInner) HasEmailPrefix() bool`

HasEmailPrefix returns a boolean if a field has been set.

### GetMemberOfPrefix

`func (o *GetAuthenticators200Response1DataInner) GetMemberOfPrefix() string`

GetMemberOfPrefix returns the MemberOfPrefix field if non-nil, zero value otherwise.

### GetMemberOfPrefixOk

`func (o *GetAuthenticators200Response1DataInner) GetMemberOfPrefixOk() (*string, bool)`

GetMemberOfPrefixOk returns a tuple with the MemberOfPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberOfPrefix

`func (o *GetAuthenticators200Response1DataInner) SetMemberOfPrefix(v string)`

SetMemberOfPrefix sets MemberOfPrefix field to given value.

### HasMemberOfPrefix

`func (o *GetAuthenticators200Response1DataInner) HasMemberOfPrefix() bool`

HasMemberOfPrefix returns a boolean if a field has been set.

### GetReferralPolicy

`func (o *GetAuthenticators200Response1DataInner) GetReferralPolicy() string`

GetReferralPolicy returns the ReferralPolicy field if non-nil, zero value otherwise.

### GetReferralPolicyOk

`func (o *GetAuthenticators200Response1DataInner) GetReferralPolicyOk() (*string, bool)`

GetReferralPolicyOk returns a tuple with the ReferralPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferralPolicy

`func (o *GetAuthenticators200Response1DataInner) SetReferralPolicy(v string)`

SetReferralPolicy sets ReferralPolicy field to given value.

### HasReferralPolicy

`func (o *GetAuthenticators200Response1DataInner) HasReferralPolicy() bool`

HasReferralPolicy returns a boolean if a field has been set.

### GetAliasDereferencePolicy

`func (o *GetAuthenticators200Response1DataInner) GetAliasDereferencePolicy() string`

GetAliasDereferencePolicy returns the AliasDereferencePolicy field if non-nil, zero value otherwise.

### GetAliasDereferencePolicyOk

`func (o *GetAuthenticators200Response1DataInner) GetAliasDereferencePolicyOk() (*string, bool)`

GetAliasDereferencePolicyOk returns a tuple with the AliasDereferencePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasDereferencePolicy

`func (o *GetAuthenticators200Response1DataInner) SetAliasDereferencePolicy(v string)`

SetAliasDereferencePolicy sets AliasDereferencePolicy field to given value.

### HasAliasDereferencePolicy

`func (o *GetAuthenticators200Response1DataInner) HasAliasDereferencePolicy() bool`

HasAliasDereferencePolicy returns a boolean if a field has been set.

### GetDescription

`func (o *GetAuthenticators200Response1DataInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetAuthenticators200Response1DataInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetAuthenticators200Response1DataInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetAuthenticators200Response1DataInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *GetAuthenticators200Response1DataInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetAuthenticators200Response1DataInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetAuthenticators200Response1DataInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetAuthenticators200Response1DataInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetUserClaim

`func (o *GetAuthenticators200Response1DataInner) GetUserClaim() string`

GetUserClaim returns the UserClaim field if non-nil, zero value otherwise.

### GetUserClaimOk

`func (o *GetAuthenticators200Response1DataInner) GetUserClaimOk() (*string, bool)`

GetUserClaimOk returns a tuple with the UserClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserClaim

`func (o *GetAuthenticators200Response1DataInner) SetUserClaim(v string)`

SetUserClaim sets UserClaim field to given value.

### HasUserClaim

`func (o *GetAuthenticators200Response1DataInner) HasUserClaim() bool`

HasUserClaim returns a boolean if a field has been set.

### GetGroupClaim

`func (o *GetAuthenticators200Response1DataInner) GetGroupClaim() string`

GetGroupClaim returns the GroupClaim field if non-nil, zero value otherwise.

### GetGroupClaimOk

`func (o *GetAuthenticators200Response1DataInner) GetGroupClaimOk() (*string, bool)`

GetGroupClaimOk returns a tuple with the GroupClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupClaim

`func (o *GetAuthenticators200Response1DataInner) SetGroupClaim(v string)`

SetGroupClaim sets GroupClaim field to given value.

### HasGroupClaim

`func (o *GetAuthenticators200Response1DataInner) HasGroupClaim() bool`

HasGroupClaim returns a boolean if a field has been set.

### GetEmailClaim

`func (o *GetAuthenticators200Response1DataInner) GetEmailClaim() string`

GetEmailClaim returns the EmailClaim field if non-nil, zero value otherwise.

### GetEmailClaimOk

`func (o *GetAuthenticators200Response1DataInner) GetEmailClaimOk() (*string, bool)`

GetEmailClaimOk returns a tuple with the EmailClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailClaim

`func (o *GetAuthenticators200Response1DataInner) SetEmailClaim(v string)`

SetEmailClaim sets EmailClaim field to given value.

### HasEmailClaim

`func (o *GetAuthenticators200Response1DataInner) HasEmailClaim() bool`

HasEmailClaim returns a boolean if a field has been set.

### GetValidationMethod

`func (o *GetAuthenticators200Response1DataInner) GetValidationMethod() string`

GetValidationMethod returns the ValidationMethod field if non-nil, zero value otherwise.

### GetValidationMethodOk

`func (o *GetAuthenticators200Response1DataInner) GetValidationMethodOk() (*string, bool)`

GetValidationMethodOk returns a tuple with the ValidationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationMethod

`func (o *GetAuthenticators200Response1DataInner) SetValidationMethod(v string)`

SetValidationMethod sets ValidationMethod field to given value.

### HasValidationMethod

`func (o *GetAuthenticators200Response1DataInner) HasValidationMethod() bool`

HasValidationMethod returns a boolean if a field has been set.

### GetClientId

`func (o *GetAuthenticators200Response1DataInner) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *GetAuthenticators200Response1DataInner) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *GetAuthenticators200Response1DataInner) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *GetAuthenticators200Response1DataInner) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *GetAuthenticators200Response1DataInner) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *GetAuthenticators200Response1DataInner) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *GetAuthenticators200Response1DataInner) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *GetAuthenticators200Response1DataInner) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetIntrospectionUrl

`func (o *GetAuthenticators200Response1DataInner) GetIntrospectionUrl() string`

GetIntrospectionUrl returns the IntrospectionUrl field if non-nil, zero value otherwise.

### GetIntrospectionUrlOk

`func (o *GetAuthenticators200Response1DataInner) GetIntrospectionUrlOk() (*string, bool)`

GetIntrospectionUrlOk returns a tuple with the IntrospectionUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntrospectionUrl

`func (o *GetAuthenticators200Response1DataInner) SetIntrospectionUrl(v string)`

SetIntrospectionUrl sets IntrospectionUrl field to given value.

### HasIntrospectionUrl

`func (o *GetAuthenticators200Response1DataInner) HasIntrospectionUrl() bool`

HasIntrospectionUrl returns a boolean if a field has been set.

### GetUserInfoUrl

`func (o *GetAuthenticators200Response1DataInner) GetUserInfoUrl() string`

GetUserInfoUrl returns the UserInfoUrl field if non-nil, zero value otherwise.

### GetUserInfoUrlOk

`func (o *GetAuthenticators200Response1DataInner) GetUserInfoUrlOk() (*string, bool)`

GetUserInfoUrlOk returns a tuple with the UserInfoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInfoUrl

`func (o *GetAuthenticators200Response1DataInner) SetUserInfoUrl(v string)`

SetUserInfoUrl sets UserInfoUrl field to given value.

### HasUserInfoUrl

`func (o *GetAuthenticators200Response1DataInner) HasUserInfoUrl() bool`

HasUserInfoUrl returns a boolean if a field has been set.

### GetAuthorizationOption

`func (o *GetAuthenticators200Response1DataInner) GetAuthorizationOption() string`

GetAuthorizationOption returns the AuthorizationOption field if non-nil, zero value otherwise.

### GetAuthorizationOptionOk

`func (o *GetAuthenticators200Response1DataInner) GetAuthorizationOptionOk() (*string, bool)`

GetAuthorizationOptionOk returns a tuple with the AuthorizationOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationOption

`func (o *GetAuthenticators200Response1DataInner) SetAuthorizationOption(v string)`

SetAuthorizationOption sets AuthorizationOption field to given value.

### HasAuthorizationOption

`func (o *GetAuthenticators200Response1DataInner) HasAuthorizationOption() bool`

HasAuthorizationOption returns a boolean if a field has been set.

### GetAudience

`func (o *GetAuthenticators200Response1DataInner) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *GetAuthenticators200Response1DataInner) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *GetAuthenticators200Response1DataInner) SetAudience(v string)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *GetAuthenticators200Response1DataInner) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### GetIssuer

`func (o *GetAuthenticators200Response1DataInner) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *GetAuthenticators200Response1DataInner) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *GetAuthenticators200Response1DataInner) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *GetAuthenticators200Response1DataInner) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetSigningCertificate

`func (o *GetAuthenticators200Response1DataInner) GetSigningCertificate() string`

GetSigningCertificate returns the SigningCertificate field if non-nil, zero value otherwise.

### GetSigningCertificateOk

`func (o *GetAuthenticators200Response1DataInner) GetSigningCertificateOk() (*string, bool)`

GetSigningCertificateOk returns a tuple with the SigningCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningCertificate

`func (o *GetAuthenticators200Response1DataInner) SetSigningCertificate(v string)`

SetSigningCertificate sets SigningCertificate field to given value.

### HasSigningCertificate

`func (o *GetAuthenticators200Response1DataInner) HasSigningCertificate() bool`

HasSigningCertificate returns a boolean if a field has been set.

### GetTimeout

`func (o *GetAuthenticators200Response1DataInner) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *GetAuthenticators200Response1DataInner) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *GetAuthenticators200Response1DataInner) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *GetAuthenticators200Response1DataInner) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetRetries

`func (o *GetAuthenticators200Response1DataInner) GetRetries() int32`

GetRetries returns the Retries field if non-nil, zero value otherwise.

### GetRetriesOk

`func (o *GetAuthenticators200Response1DataInner) GetRetriesOk() (*int32, bool)`

GetRetriesOk returns a tuple with the Retries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetries

`func (o *GetAuthenticators200Response1DataInner) SetRetries(v int32)`

SetRetries sets Retries field to given value.

### HasRetries

`func (o *GetAuthenticators200Response1DataInner) HasRetries() bool`

HasRetries returns a boolean if a field has been set.

### GetSharedSecret

`func (o *GetAuthenticators200Response1DataInner) GetSharedSecret() string`

GetSharedSecret returns the SharedSecret field if non-nil, zero value otherwise.

### GetSharedSecretOk

`func (o *GetAuthenticators200Response1DataInner) GetSharedSecretOk() (*string, bool)`

GetSharedSecretOk returns a tuple with the SharedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSecret

`func (o *GetAuthenticators200Response1DataInner) SetSharedSecret(v string)`

SetSharedSecret sets SharedSecret field to given value.

### HasSharedSecret

`func (o *GetAuthenticators200Response1DataInner) HasSharedSecret() bool`

HasSharedSecret returns a boolean if a field has been set.

### GetAuthenticationProtocol

`func (o *GetAuthenticators200Response1DataInner) GetAuthenticationProtocol() string`

GetAuthenticationProtocol returns the AuthenticationProtocol field if non-nil, zero value otherwise.

### GetAuthenticationProtocolOk

`func (o *GetAuthenticators200Response1DataInner) GetAuthenticationProtocolOk() (*string, bool)`

GetAuthenticationProtocolOk returns a tuple with the AuthenticationProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationProtocol

`func (o *GetAuthenticators200Response1DataInner) SetAuthenticationProtocol(v string)`

SetAuthenticationProtocol sets AuthenticationProtocol field to given value.

### HasAuthenticationProtocol

`func (o *GetAuthenticators200Response1DataInner) HasAuthenticationProtocol() bool`

HasAuthenticationProtocol returns a boolean if a field has been set.

### GetGroupAttribute

`func (o *GetAuthenticators200Response1DataInner) GetGroupAttribute() string`

GetGroupAttribute returns the GroupAttribute field if non-nil, zero value otherwise.

### GetGroupAttributeOk

`func (o *GetAuthenticators200Response1DataInner) GetGroupAttributeOk() (*string, bool)`

GetGroupAttributeOk returns a tuple with the GroupAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupAttribute

`func (o *GetAuthenticators200Response1DataInner) SetGroupAttribute(v string)`

SetGroupAttribute sets GroupAttribute field to given value.

### HasGroupAttribute

`func (o *GetAuthenticators200Response1DataInner) HasGroupAttribute() bool`

HasGroupAttribute returns a boolean if a field has been set.

### GetEmailAttribute

`func (o *GetAuthenticators200Response1DataInner) GetEmailAttribute() string`

GetEmailAttribute returns the EmailAttribute field if non-nil, zero value otherwise.

### GetEmailAttributeOk

`func (o *GetAuthenticators200Response1DataInner) GetEmailAttributeOk() (*string, bool)`

GetEmailAttributeOk returns a tuple with the EmailAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAttribute

`func (o *GetAuthenticators200Response1DataInner) SetEmailAttribute(v string)`

SetEmailAttribute sets EmailAttribute field to given value.

### HasEmailAttribute

`func (o *GetAuthenticators200Response1DataInner) HasEmailAttribute() bool`

HasEmailAttribute returns a boolean if a field has been set.

### GetSignOnUrl

`func (o *GetAuthenticators200Response1DataInner) GetSignOnUrl() string`

GetSignOnUrl returns the SignOnUrl field if non-nil, zero value otherwise.

### GetSignOnUrlOk

`func (o *GetAuthenticators200Response1DataInner) GetSignOnUrlOk() (*string, bool)`

GetSignOnUrlOk returns a tuple with the SignOnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignOnUrl

`func (o *GetAuthenticators200Response1DataInner) SetSignOnUrl(v string)`

SetSignOnUrl sets SignOnUrl field to given value.

### HasSignOnUrl

`func (o *GetAuthenticators200Response1DataInner) HasSignOnUrl() bool`

HasSignOnUrl returns a boolean if a field has been set.

### GetLogoutUrl

`func (o *GetAuthenticators200Response1DataInner) GetLogoutUrl() string`

GetLogoutUrl returns the LogoutUrl field if non-nil, zero value otherwise.

### GetLogoutUrlOk

`func (o *GetAuthenticators200Response1DataInner) GetLogoutUrlOk() (*string, bool)`

GetLogoutUrlOk returns a tuple with the LogoutUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoutUrl

`func (o *GetAuthenticators200Response1DataInner) SetLogoutUrl(v string)`

SetLogoutUrl sets LogoutUrl field to given value.

### HasLogoutUrl

`func (o *GetAuthenticators200Response1DataInner) HasLogoutUrl() bool`

HasLogoutUrl returns a boolean if a field has been set.

### GetEntityId

`func (o *GetAuthenticators200Response1DataInner) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *GetAuthenticators200Response1DataInner) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *GetAuthenticators200Response1DataInner) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *GetAuthenticators200Response1DataInner) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetAttributes

`func (o *GetAuthenticators200Response1DataInner) GetAttributes() []NameValuePairBean`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GetAuthenticators200Response1DataInner) GetAttributesOk() (*[]NameValuePairBean, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GetAuthenticators200Response1DataInner) SetAttributes(v []NameValuePairBean)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *GetAuthenticators200Response1DataInner) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


