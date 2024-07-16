# PostAuthenticatorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | **string** | The type of authenticator. | 
**Name** | **string** | The name for the authenticator. | 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**SecondaryAuthenticator** | Pointer to [**InlinedAuthenticator**](InlinedAuthenticator.md) |  | [optional] 
**Hostname** | **string** | The fully qualified domain name or IP address for the authenticator. | 
**Port** | Pointer to **int32** | The TCP port number used for communication between the client and server. | [optional] 
**AdminDn** | **string** | The distinguished name or relative distinguished name of a user with rights to search the LDAP directory. | 
**AdminPassword** | **string** | The password for the distinguished name or relative distinguished name. | 
**SslEnabled** | Pointer to **bool** | Determines whether SSL communication is enabled between Address Manager and the LDAP server. | [optional] 
**SchemaType** | **string** | The type of LDAP schema. | 
**SearchBase** | **string** | The search base distinguished name location from which the search for users on the LDAP server begins. | 
**UserObjectClass** | Pointer to **string** | The user object class used to locate an LDAP user. | [optional] 
**GroupObjectClass** | Pointer to **string** | The group object class used to indicate that a DN is a group. | [optional] 
**UserPrefix** | Pointer to **string** | The user attribute for user accounts in the LDAP tree. | [optional] 
**EmailPrefix** | Pointer to **string** | The variable to be used for the email prefix. | [optional] 
**MemberOfPrefix** | Pointer to **string** | The attribute used to store user-group membership information. | [optional] 
**ReferralPolicy** | **string** | The environment property that indicates how to handle referrals from external resources to the service providers. | 
**AliasDereferencePolicy** | **string** | The environment property that indicates whether alias entries are dereferenced. | 
**Description** | Pointer to **string** | The description of the SAML IdP configuration. | [optional] 
**Enabled** | Pointer to **bool** | Indicates whether SAML IdP authentication is enabled. | [optional] 
**UserClaim** | **string** | The user claim name of the authorization server. | 
**GroupClaim** | **string** | The group claim name of the authorization server. | 
**EmailClaim** | **string** | The email claim name of the authorization server. | 
**ValidationMethod** | **string** | The method that indicates where the token validation occurs. | 
**ClientId** | Pointer to **string** | The public identifier of the application. | [optional] 
**ClientSecret** | Pointer to **string** | The client secret known only to the application and authorization server. | [optional] 
**IntrospectionUrl** | Pointer to **string** | The introspection endpoint that allows Address Manager to check the validity of access tokens. | [optional] 
**UserInfoUrl** | Pointer to **string** | The endpoint containing information about the user, including the group membership information and user ID. | [optional] 
**AuthorizationOption** | Pointer to **string** | The authorization method used when Address Manager sends the client ID and client secret to the authorization server. | [optional] 
**Audience** | Pointer to **string** | The name of the BAM API bearer token obtained from the authorization server. | [optional] 
**Issuer** | Pointer to **string** | The name of the token issuer. | [optional] 
**SigningCertificate** | **string** | The SAML IdP server signing certificate. | 
**Timeout** | Pointer to **string** | The value that overrides the timeout setting used for authentication requests sent to the TACACS+ server. | [optional] 
**Retries** | Pointer to **int32** | The number of times Address Manager attempts to retransmit a failed authentication request sent to the RADIUS server. | [optional] 
**SharedSecret** | **string** | The shared secret used to encrypt and decrypt packets between the client and the server. | 
**AuthenticationProtocol** | **string** | The authentication protocol. | 
**GroupAttribute** | Pointer to **string** | The special attribute used for the custom service in the TACACS+ server. This attribute is used to get the group name defined in the TACACS+ server. | [optional] 
**EmailAttribute** | Pointer to **string** | The attribute name for email in the SAML response. | [optional] 
**SignOnUrl** | **string** | The sign-in URL of the SAMP IdP. | 
**LogoutUrl** | Pointer to **string** | The logout URL of the SAMP IdP. | [optional] 
**EntityId** | **string** | The entity ID of the SAMP IdP. | 
**Attributes** | Pointer to [**[]NameValuePairBean**](NameValuePairBean.md) |  | [optional] 

## Methods

### NewPostAuthenticatorRequest

`func NewPostAuthenticatorRequest(type_ string, name string, hostname string, adminDn string, adminPassword string, schemaType string, searchBase string, referralPolicy string, aliasDereferencePolicy string, userClaim string, groupClaim string, emailClaim string, validationMethod string, signingCertificate string, sharedSecret string, authenticationProtocol string, signOnUrl string, entityId string, ) *PostAuthenticatorRequest`

NewPostAuthenticatorRequest instantiates a new PostAuthenticatorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostAuthenticatorRequestWithDefaults

`func NewPostAuthenticatorRequestWithDefaults() *PostAuthenticatorRequest`

NewPostAuthenticatorRequestWithDefaults instantiates a new PostAuthenticatorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PostAuthenticatorRequest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostAuthenticatorRequest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostAuthenticatorRequest) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *PostAuthenticatorRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *PostAuthenticatorRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PostAuthenticatorRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PostAuthenticatorRequest) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *PostAuthenticatorRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostAuthenticatorRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostAuthenticatorRequest) SetName(v string)`

SetName sets Name field to given value.


### GetUserDefinedFields

`func (o *PostAuthenticatorRequest) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *PostAuthenticatorRequest) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *PostAuthenticatorRequest) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *PostAuthenticatorRequest) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetSecondaryAuthenticator

`func (o *PostAuthenticatorRequest) GetSecondaryAuthenticator() InlinedAuthenticator`

GetSecondaryAuthenticator returns the SecondaryAuthenticator field if non-nil, zero value otherwise.

### GetSecondaryAuthenticatorOk

`func (o *PostAuthenticatorRequest) GetSecondaryAuthenticatorOk() (*InlinedAuthenticator, bool)`

GetSecondaryAuthenticatorOk returns a tuple with the SecondaryAuthenticator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryAuthenticator

`func (o *PostAuthenticatorRequest) SetSecondaryAuthenticator(v InlinedAuthenticator)`

SetSecondaryAuthenticator sets SecondaryAuthenticator field to given value.

### HasSecondaryAuthenticator

`func (o *PostAuthenticatorRequest) HasSecondaryAuthenticator() bool`

HasSecondaryAuthenticator returns a boolean if a field has been set.

### GetHostname

`func (o *PostAuthenticatorRequest) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *PostAuthenticatorRequest) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *PostAuthenticatorRequest) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetPort

`func (o *PostAuthenticatorRequest) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *PostAuthenticatorRequest) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *PostAuthenticatorRequest) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *PostAuthenticatorRequest) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetAdminDn

`func (o *PostAuthenticatorRequest) GetAdminDn() string`

GetAdminDn returns the AdminDn field if non-nil, zero value otherwise.

### GetAdminDnOk

`func (o *PostAuthenticatorRequest) GetAdminDnOk() (*string, bool)`

GetAdminDnOk returns a tuple with the AdminDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminDn

`func (o *PostAuthenticatorRequest) SetAdminDn(v string)`

SetAdminDn sets AdminDn field to given value.


### GetAdminPassword

`func (o *PostAuthenticatorRequest) GetAdminPassword() string`

GetAdminPassword returns the AdminPassword field if non-nil, zero value otherwise.

### GetAdminPasswordOk

`func (o *PostAuthenticatorRequest) GetAdminPasswordOk() (*string, bool)`

GetAdminPasswordOk returns a tuple with the AdminPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminPassword

`func (o *PostAuthenticatorRequest) SetAdminPassword(v string)`

SetAdminPassword sets AdminPassword field to given value.


### GetSslEnabled

`func (o *PostAuthenticatorRequest) GetSslEnabled() bool`

GetSslEnabled returns the SslEnabled field if non-nil, zero value otherwise.

### GetSslEnabledOk

`func (o *PostAuthenticatorRequest) GetSslEnabledOk() (*bool, bool)`

GetSslEnabledOk returns a tuple with the SslEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslEnabled

`func (o *PostAuthenticatorRequest) SetSslEnabled(v bool)`

SetSslEnabled sets SslEnabled field to given value.

### HasSslEnabled

`func (o *PostAuthenticatorRequest) HasSslEnabled() bool`

HasSslEnabled returns a boolean if a field has been set.

### GetSchemaType

`func (o *PostAuthenticatorRequest) GetSchemaType() string`

GetSchemaType returns the SchemaType field if non-nil, zero value otherwise.

### GetSchemaTypeOk

`func (o *PostAuthenticatorRequest) GetSchemaTypeOk() (*string, bool)`

GetSchemaTypeOk returns a tuple with the SchemaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaType

`func (o *PostAuthenticatorRequest) SetSchemaType(v string)`

SetSchemaType sets SchemaType field to given value.


### GetSearchBase

`func (o *PostAuthenticatorRequest) GetSearchBase() string`

GetSearchBase returns the SearchBase field if non-nil, zero value otherwise.

### GetSearchBaseOk

`func (o *PostAuthenticatorRequest) GetSearchBaseOk() (*string, bool)`

GetSearchBaseOk returns a tuple with the SearchBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchBase

`func (o *PostAuthenticatorRequest) SetSearchBase(v string)`

SetSearchBase sets SearchBase field to given value.


### GetUserObjectClass

`func (o *PostAuthenticatorRequest) GetUserObjectClass() string`

GetUserObjectClass returns the UserObjectClass field if non-nil, zero value otherwise.

### GetUserObjectClassOk

`func (o *PostAuthenticatorRequest) GetUserObjectClassOk() (*string, bool)`

GetUserObjectClassOk returns a tuple with the UserObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserObjectClass

`func (o *PostAuthenticatorRequest) SetUserObjectClass(v string)`

SetUserObjectClass sets UserObjectClass field to given value.

### HasUserObjectClass

`func (o *PostAuthenticatorRequest) HasUserObjectClass() bool`

HasUserObjectClass returns a boolean if a field has been set.

### GetGroupObjectClass

`func (o *PostAuthenticatorRequest) GetGroupObjectClass() string`

GetGroupObjectClass returns the GroupObjectClass field if non-nil, zero value otherwise.

### GetGroupObjectClassOk

`func (o *PostAuthenticatorRequest) GetGroupObjectClassOk() (*string, bool)`

GetGroupObjectClassOk returns a tuple with the GroupObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupObjectClass

`func (o *PostAuthenticatorRequest) SetGroupObjectClass(v string)`

SetGroupObjectClass sets GroupObjectClass field to given value.

### HasGroupObjectClass

`func (o *PostAuthenticatorRequest) HasGroupObjectClass() bool`

HasGroupObjectClass returns a boolean if a field has been set.

### GetUserPrefix

`func (o *PostAuthenticatorRequest) GetUserPrefix() string`

GetUserPrefix returns the UserPrefix field if non-nil, zero value otherwise.

### GetUserPrefixOk

`func (o *PostAuthenticatorRequest) GetUserPrefixOk() (*string, bool)`

GetUserPrefixOk returns a tuple with the UserPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPrefix

`func (o *PostAuthenticatorRequest) SetUserPrefix(v string)`

SetUserPrefix sets UserPrefix field to given value.

### HasUserPrefix

`func (o *PostAuthenticatorRequest) HasUserPrefix() bool`

HasUserPrefix returns a boolean if a field has been set.

### GetEmailPrefix

`func (o *PostAuthenticatorRequest) GetEmailPrefix() string`

GetEmailPrefix returns the EmailPrefix field if non-nil, zero value otherwise.

### GetEmailPrefixOk

`func (o *PostAuthenticatorRequest) GetEmailPrefixOk() (*string, bool)`

GetEmailPrefixOk returns a tuple with the EmailPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailPrefix

`func (o *PostAuthenticatorRequest) SetEmailPrefix(v string)`

SetEmailPrefix sets EmailPrefix field to given value.

### HasEmailPrefix

`func (o *PostAuthenticatorRequest) HasEmailPrefix() bool`

HasEmailPrefix returns a boolean if a field has been set.

### GetMemberOfPrefix

`func (o *PostAuthenticatorRequest) GetMemberOfPrefix() string`

GetMemberOfPrefix returns the MemberOfPrefix field if non-nil, zero value otherwise.

### GetMemberOfPrefixOk

`func (o *PostAuthenticatorRequest) GetMemberOfPrefixOk() (*string, bool)`

GetMemberOfPrefixOk returns a tuple with the MemberOfPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberOfPrefix

`func (o *PostAuthenticatorRequest) SetMemberOfPrefix(v string)`

SetMemberOfPrefix sets MemberOfPrefix field to given value.

### HasMemberOfPrefix

`func (o *PostAuthenticatorRequest) HasMemberOfPrefix() bool`

HasMemberOfPrefix returns a boolean if a field has been set.

### GetReferralPolicy

`func (o *PostAuthenticatorRequest) GetReferralPolicy() string`

GetReferralPolicy returns the ReferralPolicy field if non-nil, zero value otherwise.

### GetReferralPolicyOk

`func (o *PostAuthenticatorRequest) GetReferralPolicyOk() (*string, bool)`

GetReferralPolicyOk returns a tuple with the ReferralPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferralPolicy

`func (o *PostAuthenticatorRequest) SetReferralPolicy(v string)`

SetReferralPolicy sets ReferralPolicy field to given value.


### GetAliasDereferencePolicy

`func (o *PostAuthenticatorRequest) GetAliasDereferencePolicy() string`

GetAliasDereferencePolicy returns the AliasDereferencePolicy field if non-nil, zero value otherwise.

### GetAliasDereferencePolicyOk

`func (o *PostAuthenticatorRequest) GetAliasDereferencePolicyOk() (*string, bool)`

GetAliasDereferencePolicyOk returns a tuple with the AliasDereferencePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasDereferencePolicy

`func (o *PostAuthenticatorRequest) SetAliasDereferencePolicy(v string)`

SetAliasDereferencePolicy sets AliasDereferencePolicy field to given value.


### GetDescription

`func (o *PostAuthenticatorRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PostAuthenticatorRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PostAuthenticatorRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PostAuthenticatorRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *PostAuthenticatorRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PostAuthenticatorRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PostAuthenticatorRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PostAuthenticatorRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetUserClaim

`func (o *PostAuthenticatorRequest) GetUserClaim() string`

GetUserClaim returns the UserClaim field if non-nil, zero value otherwise.

### GetUserClaimOk

`func (o *PostAuthenticatorRequest) GetUserClaimOk() (*string, bool)`

GetUserClaimOk returns a tuple with the UserClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserClaim

`func (o *PostAuthenticatorRequest) SetUserClaim(v string)`

SetUserClaim sets UserClaim field to given value.


### GetGroupClaim

`func (o *PostAuthenticatorRequest) GetGroupClaim() string`

GetGroupClaim returns the GroupClaim field if non-nil, zero value otherwise.

### GetGroupClaimOk

`func (o *PostAuthenticatorRequest) GetGroupClaimOk() (*string, bool)`

GetGroupClaimOk returns a tuple with the GroupClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupClaim

`func (o *PostAuthenticatorRequest) SetGroupClaim(v string)`

SetGroupClaim sets GroupClaim field to given value.


### GetEmailClaim

`func (o *PostAuthenticatorRequest) GetEmailClaim() string`

GetEmailClaim returns the EmailClaim field if non-nil, zero value otherwise.

### GetEmailClaimOk

`func (o *PostAuthenticatorRequest) GetEmailClaimOk() (*string, bool)`

GetEmailClaimOk returns a tuple with the EmailClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailClaim

`func (o *PostAuthenticatorRequest) SetEmailClaim(v string)`

SetEmailClaim sets EmailClaim field to given value.


### GetValidationMethod

`func (o *PostAuthenticatorRequest) GetValidationMethod() string`

GetValidationMethod returns the ValidationMethod field if non-nil, zero value otherwise.

### GetValidationMethodOk

`func (o *PostAuthenticatorRequest) GetValidationMethodOk() (*string, bool)`

GetValidationMethodOk returns a tuple with the ValidationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationMethod

`func (o *PostAuthenticatorRequest) SetValidationMethod(v string)`

SetValidationMethod sets ValidationMethod field to given value.


### GetClientId

`func (o *PostAuthenticatorRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *PostAuthenticatorRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *PostAuthenticatorRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *PostAuthenticatorRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *PostAuthenticatorRequest) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *PostAuthenticatorRequest) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *PostAuthenticatorRequest) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *PostAuthenticatorRequest) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetIntrospectionUrl

`func (o *PostAuthenticatorRequest) GetIntrospectionUrl() string`

GetIntrospectionUrl returns the IntrospectionUrl field if non-nil, zero value otherwise.

### GetIntrospectionUrlOk

`func (o *PostAuthenticatorRequest) GetIntrospectionUrlOk() (*string, bool)`

GetIntrospectionUrlOk returns a tuple with the IntrospectionUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntrospectionUrl

`func (o *PostAuthenticatorRequest) SetIntrospectionUrl(v string)`

SetIntrospectionUrl sets IntrospectionUrl field to given value.

### HasIntrospectionUrl

`func (o *PostAuthenticatorRequest) HasIntrospectionUrl() bool`

HasIntrospectionUrl returns a boolean if a field has been set.

### GetUserInfoUrl

`func (o *PostAuthenticatorRequest) GetUserInfoUrl() string`

GetUserInfoUrl returns the UserInfoUrl field if non-nil, zero value otherwise.

### GetUserInfoUrlOk

`func (o *PostAuthenticatorRequest) GetUserInfoUrlOk() (*string, bool)`

GetUserInfoUrlOk returns a tuple with the UserInfoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInfoUrl

`func (o *PostAuthenticatorRequest) SetUserInfoUrl(v string)`

SetUserInfoUrl sets UserInfoUrl field to given value.

### HasUserInfoUrl

`func (o *PostAuthenticatorRequest) HasUserInfoUrl() bool`

HasUserInfoUrl returns a boolean if a field has been set.

### GetAuthorizationOption

`func (o *PostAuthenticatorRequest) GetAuthorizationOption() string`

GetAuthorizationOption returns the AuthorizationOption field if non-nil, zero value otherwise.

### GetAuthorizationOptionOk

`func (o *PostAuthenticatorRequest) GetAuthorizationOptionOk() (*string, bool)`

GetAuthorizationOptionOk returns a tuple with the AuthorizationOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationOption

`func (o *PostAuthenticatorRequest) SetAuthorizationOption(v string)`

SetAuthorizationOption sets AuthorizationOption field to given value.

### HasAuthorizationOption

`func (o *PostAuthenticatorRequest) HasAuthorizationOption() bool`

HasAuthorizationOption returns a boolean if a field has been set.

### GetAudience

`func (o *PostAuthenticatorRequest) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *PostAuthenticatorRequest) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *PostAuthenticatorRequest) SetAudience(v string)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *PostAuthenticatorRequest) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### GetIssuer

`func (o *PostAuthenticatorRequest) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *PostAuthenticatorRequest) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *PostAuthenticatorRequest) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *PostAuthenticatorRequest) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetSigningCertificate

`func (o *PostAuthenticatorRequest) GetSigningCertificate() string`

GetSigningCertificate returns the SigningCertificate field if non-nil, zero value otherwise.

### GetSigningCertificateOk

`func (o *PostAuthenticatorRequest) GetSigningCertificateOk() (*string, bool)`

GetSigningCertificateOk returns a tuple with the SigningCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningCertificate

`func (o *PostAuthenticatorRequest) SetSigningCertificate(v string)`

SetSigningCertificate sets SigningCertificate field to given value.


### GetTimeout

`func (o *PostAuthenticatorRequest) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *PostAuthenticatorRequest) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *PostAuthenticatorRequest) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *PostAuthenticatorRequest) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetRetries

`func (o *PostAuthenticatorRequest) GetRetries() int32`

GetRetries returns the Retries field if non-nil, zero value otherwise.

### GetRetriesOk

`func (o *PostAuthenticatorRequest) GetRetriesOk() (*int32, bool)`

GetRetriesOk returns a tuple with the Retries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetries

`func (o *PostAuthenticatorRequest) SetRetries(v int32)`

SetRetries sets Retries field to given value.

### HasRetries

`func (o *PostAuthenticatorRequest) HasRetries() bool`

HasRetries returns a boolean if a field has been set.

### GetSharedSecret

`func (o *PostAuthenticatorRequest) GetSharedSecret() string`

GetSharedSecret returns the SharedSecret field if non-nil, zero value otherwise.

### GetSharedSecretOk

`func (o *PostAuthenticatorRequest) GetSharedSecretOk() (*string, bool)`

GetSharedSecretOk returns a tuple with the SharedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSecret

`func (o *PostAuthenticatorRequest) SetSharedSecret(v string)`

SetSharedSecret sets SharedSecret field to given value.


### GetAuthenticationProtocol

`func (o *PostAuthenticatorRequest) GetAuthenticationProtocol() string`

GetAuthenticationProtocol returns the AuthenticationProtocol field if non-nil, zero value otherwise.

### GetAuthenticationProtocolOk

`func (o *PostAuthenticatorRequest) GetAuthenticationProtocolOk() (*string, bool)`

GetAuthenticationProtocolOk returns a tuple with the AuthenticationProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationProtocol

`func (o *PostAuthenticatorRequest) SetAuthenticationProtocol(v string)`

SetAuthenticationProtocol sets AuthenticationProtocol field to given value.


### GetGroupAttribute

`func (o *PostAuthenticatorRequest) GetGroupAttribute() string`

GetGroupAttribute returns the GroupAttribute field if non-nil, zero value otherwise.

### GetGroupAttributeOk

`func (o *PostAuthenticatorRequest) GetGroupAttributeOk() (*string, bool)`

GetGroupAttributeOk returns a tuple with the GroupAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupAttribute

`func (o *PostAuthenticatorRequest) SetGroupAttribute(v string)`

SetGroupAttribute sets GroupAttribute field to given value.

### HasGroupAttribute

`func (o *PostAuthenticatorRequest) HasGroupAttribute() bool`

HasGroupAttribute returns a boolean if a field has been set.

### GetEmailAttribute

`func (o *PostAuthenticatorRequest) GetEmailAttribute() string`

GetEmailAttribute returns the EmailAttribute field if non-nil, zero value otherwise.

### GetEmailAttributeOk

`func (o *PostAuthenticatorRequest) GetEmailAttributeOk() (*string, bool)`

GetEmailAttributeOk returns a tuple with the EmailAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAttribute

`func (o *PostAuthenticatorRequest) SetEmailAttribute(v string)`

SetEmailAttribute sets EmailAttribute field to given value.

### HasEmailAttribute

`func (o *PostAuthenticatorRequest) HasEmailAttribute() bool`

HasEmailAttribute returns a boolean if a field has been set.

### GetSignOnUrl

`func (o *PostAuthenticatorRequest) GetSignOnUrl() string`

GetSignOnUrl returns the SignOnUrl field if non-nil, zero value otherwise.

### GetSignOnUrlOk

`func (o *PostAuthenticatorRequest) GetSignOnUrlOk() (*string, bool)`

GetSignOnUrlOk returns a tuple with the SignOnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignOnUrl

`func (o *PostAuthenticatorRequest) SetSignOnUrl(v string)`

SetSignOnUrl sets SignOnUrl field to given value.


### GetLogoutUrl

`func (o *PostAuthenticatorRequest) GetLogoutUrl() string`

GetLogoutUrl returns the LogoutUrl field if non-nil, zero value otherwise.

### GetLogoutUrlOk

`func (o *PostAuthenticatorRequest) GetLogoutUrlOk() (*string, bool)`

GetLogoutUrlOk returns a tuple with the LogoutUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoutUrl

`func (o *PostAuthenticatorRequest) SetLogoutUrl(v string)`

SetLogoutUrl sets LogoutUrl field to given value.

### HasLogoutUrl

`func (o *PostAuthenticatorRequest) HasLogoutUrl() bool`

HasLogoutUrl returns a boolean if a field has been set.

### GetEntityId

`func (o *PostAuthenticatorRequest) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *PostAuthenticatorRequest) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *PostAuthenticatorRequest) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.


### GetAttributes

`func (o *PostAuthenticatorRequest) GetAttributes() []NameValuePairBean`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PostAuthenticatorRequest) GetAttributesOk() (*[]NameValuePairBean, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PostAuthenticatorRequest) SetAttributes(v []NameValuePairBean)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *PostAuthenticatorRequest) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


