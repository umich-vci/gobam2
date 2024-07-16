# PutAuthenticatorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | **string** | The type of authenticator. | 
**Name** | **string** | The name for the authenticator. | 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**SecondaryAuthenticator** | Pointer to [**InlinedAuthenticator**](InlinedAuthenticator.md) |  | [optional] 
**Hostname** | **string** | The fully qualified domain name or IP address for the authenticator. | 
**Port** | **int32** | The TCP port number used for communication between the client and server. | 
**AdminDn** | **string** | The distinguished name or relative distinguished name of a user with rights to search the LDAP directory. | 
**AdminPassword** | Pointer to **string** | The password for the distinguished name or relative distinguished name. | [optional] 
**SslEnabled** | **bool** | Determines whether SSL communication is enabled between Address Manager and the LDAP server. | 
**SchemaType** | **string** | The type of LDAP schema. | 
**SearchBase** | **string** | The search base distinguished name location from which the search for users on the LDAP server begins. | 
**UserObjectClass** | **string** | The user object class used to locate an LDAP user. | 
**GroupObjectClass** | **string** | The group object class used to indicate that a DN is a group. | 
**UserPrefix** | **string** | The user attribute for user accounts in the LDAP tree. | 
**EmailPrefix** | **string** | The variable to be used for the email prefix. | 
**MemberOfPrefix** | **string** | The attribute used to store user-group membership information. | 
**ReferralPolicy** | **string** | The environment property that indicates how to handle referrals from external resources to the service providers. | 
**AliasDereferencePolicy** | **string** | The environment property that indicates whether alias entries are dereferenced. | 
**Description** | Pointer to **string** | The description of the SAML IdP configuration. | [optional] 
**Enabled** | **bool** | Indicates whether SAML IdP authentication is enabled. | 
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
**Timeout** | **string** | The value that overrides the timeout setting used for authentication requests sent to the TACACS+ server. | 
**Retries** | **int32** | The number of times Address Manager attempts to retransmit a failed authentication request sent to the RADIUS server. | 
**SharedSecret** | Pointer to **string** | The shared secret used to encrypt and decrypt packets between the client and the server. | [optional] 
**AuthenticationProtocol** | **string** | The authentication protocol. | 
**GroupAttribute** | **string** | The special attribute used for the custom service in the TACACS+ server. This attribute is used to get the group name defined in the TACACS+ server. | 
**EmailAttribute** | **string** | The attribute name for email in the SAML response. | 
**SignOnUrl** | **string** | The sign-in URL of the SAMP IdP. | 
**LogoutUrl** | Pointer to **string** | The logout URL of the SAMP IdP. | [optional] 
**EntityId** | **string** | The entity ID of the SAMP IdP. | 
**Attributes** | [**[]NameValuePairBean**](NameValuePairBean.md) |  | 

## Methods

### NewPutAuthenticatorRequest

`func NewPutAuthenticatorRequest(type_ string, name string, hostname string, port int32, adminDn string, sslEnabled bool, schemaType string, searchBase string, userObjectClass string, groupObjectClass string, userPrefix string, emailPrefix string, memberOfPrefix string, referralPolicy string, aliasDereferencePolicy string, enabled bool, userClaim string, groupClaim string, emailClaim string, validationMethod string, signingCertificate string, timeout string, retries int32, authenticationProtocol string, groupAttribute string, emailAttribute string, signOnUrl string, entityId string, attributes []NameValuePairBean, ) *PutAuthenticatorRequest`

NewPutAuthenticatorRequest instantiates a new PutAuthenticatorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutAuthenticatorRequestWithDefaults

`func NewPutAuthenticatorRequestWithDefaults() *PutAuthenticatorRequest`

NewPutAuthenticatorRequestWithDefaults instantiates a new PutAuthenticatorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PutAuthenticatorRequest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PutAuthenticatorRequest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PutAuthenticatorRequest) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *PutAuthenticatorRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *PutAuthenticatorRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PutAuthenticatorRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PutAuthenticatorRequest) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *PutAuthenticatorRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PutAuthenticatorRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PutAuthenticatorRequest) SetName(v string)`

SetName sets Name field to given value.


### GetUserDefinedFields

`func (o *PutAuthenticatorRequest) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *PutAuthenticatorRequest) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *PutAuthenticatorRequest) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *PutAuthenticatorRequest) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetSecondaryAuthenticator

`func (o *PutAuthenticatorRequest) GetSecondaryAuthenticator() InlinedAuthenticator`

GetSecondaryAuthenticator returns the SecondaryAuthenticator field if non-nil, zero value otherwise.

### GetSecondaryAuthenticatorOk

`func (o *PutAuthenticatorRequest) GetSecondaryAuthenticatorOk() (*InlinedAuthenticator, bool)`

GetSecondaryAuthenticatorOk returns a tuple with the SecondaryAuthenticator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryAuthenticator

`func (o *PutAuthenticatorRequest) SetSecondaryAuthenticator(v InlinedAuthenticator)`

SetSecondaryAuthenticator sets SecondaryAuthenticator field to given value.

### HasSecondaryAuthenticator

`func (o *PutAuthenticatorRequest) HasSecondaryAuthenticator() bool`

HasSecondaryAuthenticator returns a boolean if a field has been set.

### GetHostname

`func (o *PutAuthenticatorRequest) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *PutAuthenticatorRequest) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *PutAuthenticatorRequest) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetPort

`func (o *PutAuthenticatorRequest) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *PutAuthenticatorRequest) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *PutAuthenticatorRequest) SetPort(v int32)`

SetPort sets Port field to given value.


### GetAdminDn

`func (o *PutAuthenticatorRequest) GetAdminDn() string`

GetAdminDn returns the AdminDn field if non-nil, zero value otherwise.

### GetAdminDnOk

`func (o *PutAuthenticatorRequest) GetAdminDnOk() (*string, bool)`

GetAdminDnOk returns a tuple with the AdminDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminDn

`func (o *PutAuthenticatorRequest) SetAdminDn(v string)`

SetAdminDn sets AdminDn field to given value.


### GetAdminPassword

`func (o *PutAuthenticatorRequest) GetAdminPassword() string`

GetAdminPassword returns the AdminPassword field if non-nil, zero value otherwise.

### GetAdminPasswordOk

`func (o *PutAuthenticatorRequest) GetAdminPasswordOk() (*string, bool)`

GetAdminPasswordOk returns a tuple with the AdminPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminPassword

`func (o *PutAuthenticatorRequest) SetAdminPassword(v string)`

SetAdminPassword sets AdminPassword field to given value.

### HasAdminPassword

`func (o *PutAuthenticatorRequest) HasAdminPassword() bool`

HasAdminPassword returns a boolean if a field has been set.

### GetSslEnabled

`func (o *PutAuthenticatorRequest) GetSslEnabled() bool`

GetSslEnabled returns the SslEnabled field if non-nil, zero value otherwise.

### GetSslEnabledOk

`func (o *PutAuthenticatorRequest) GetSslEnabledOk() (*bool, bool)`

GetSslEnabledOk returns a tuple with the SslEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslEnabled

`func (o *PutAuthenticatorRequest) SetSslEnabled(v bool)`

SetSslEnabled sets SslEnabled field to given value.


### GetSchemaType

`func (o *PutAuthenticatorRequest) GetSchemaType() string`

GetSchemaType returns the SchemaType field if non-nil, zero value otherwise.

### GetSchemaTypeOk

`func (o *PutAuthenticatorRequest) GetSchemaTypeOk() (*string, bool)`

GetSchemaTypeOk returns a tuple with the SchemaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaType

`func (o *PutAuthenticatorRequest) SetSchemaType(v string)`

SetSchemaType sets SchemaType field to given value.


### GetSearchBase

`func (o *PutAuthenticatorRequest) GetSearchBase() string`

GetSearchBase returns the SearchBase field if non-nil, zero value otherwise.

### GetSearchBaseOk

`func (o *PutAuthenticatorRequest) GetSearchBaseOk() (*string, bool)`

GetSearchBaseOk returns a tuple with the SearchBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchBase

`func (o *PutAuthenticatorRequest) SetSearchBase(v string)`

SetSearchBase sets SearchBase field to given value.


### GetUserObjectClass

`func (o *PutAuthenticatorRequest) GetUserObjectClass() string`

GetUserObjectClass returns the UserObjectClass field if non-nil, zero value otherwise.

### GetUserObjectClassOk

`func (o *PutAuthenticatorRequest) GetUserObjectClassOk() (*string, bool)`

GetUserObjectClassOk returns a tuple with the UserObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserObjectClass

`func (o *PutAuthenticatorRequest) SetUserObjectClass(v string)`

SetUserObjectClass sets UserObjectClass field to given value.


### GetGroupObjectClass

`func (o *PutAuthenticatorRequest) GetGroupObjectClass() string`

GetGroupObjectClass returns the GroupObjectClass field if non-nil, zero value otherwise.

### GetGroupObjectClassOk

`func (o *PutAuthenticatorRequest) GetGroupObjectClassOk() (*string, bool)`

GetGroupObjectClassOk returns a tuple with the GroupObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupObjectClass

`func (o *PutAuthenticatorRequest) SetGroupObjectClass(v string)`

SetGroupObjectClass sets GroupObjectClass field to given value.


### GetUserPrefix

`func (o *PutAuthenticatorRequest) GetUserPrefix() string`

GetUserPrefix returns the UserPrefix field if non-nil, zero value otherwise.

### GetUserPrefixOk

`func (o *PutAuthenticatorRequest) GetUserPrefixOk() (*string, bool)`

GetUserPrefixOk returns a tuple with the UserPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPrefix

`func (o *PutAuthenticatorRequest) SetUserPrefix(v string)`

SetUserPrefix sets UserPrefix field to given value.


### GetEmailPrefix

`func (o *PutAuthenticatorRequest) GetEmailPrefix() string`

GetEmailPrefix returns the EmailPrefix field if non-nil, zero value otherwise.

### GetEmailPrefixOk

`func (o *PutAuthenticatorRequest) GetEmailPrefixOk() (*string, bool)`

GetEmailPrefixOk returns a tuple with the EmailPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailPrefix

`func (o *PutAuthenticatorRequest) SetEmailPrefix(v string)`

SetEmailPrefix sets EmailPrefix field to given value.


### GetMemberOfPrefix

`func (o *PutAuthenticatorRequest) GetMemberOfPrefix() string`

GetMemberOfPrefix returns the MemberOfPrefix field if non-nil, zero value otherwise.

### GetMemberOfPrefixOk

`func (o *PutAuthenticatorRequest) GetMemberOfPrefixOk() (*string, bool)`

GetMemberOfPrefixOk returns a tuple with the MemberOfPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberOfPrefix

`func (o *PutAuthenticatorRequest) SetMemberOfPrefix(v string)`

SetMemberOfPrefix sets MemberOfPrefix field to given value.


### GetReferralPolicy

`func (o *PutAuthenticatorRequest) GetReferralPolicy() string`

GetReferralPolicy returns the ReferralPolicy field if non-nil, zero value otherwise.

### GetReferralPolicyOk

`func (o *PutAuthenticatorRequest) GetReferralPolicyOk() (*string, bool)`

GetReferralPolicyOk returns a tuple with the ReferralPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferralPolicy

`func (o *PutAuthenticatorRequest) SetReferralPolicy(v string)`

SetReferralPolicy sets ReferralPolicy field to given value.


### GetAliasDereferencePolicy

`func (o *PutAuthenticatorRequest) GetAliasDereferencePolicy() string`

GetAliasDereferencePolicy returns the AliasDereferencePolicy field if non-nil, zero value otherwise.

### GetAliasDereferencePolicyOk

`func (o *PutAuthenticatorRequest) GetAliasDereferencePolicyOk() (*string, bool)`

GetAliasDereferencePolicyOk returns a tuple with the AliasDereferencePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasDereferencePolicy

`func (o *PutAuthenticatorRequest) SetAliasDereferencePolicy(v string)`

SetAliasDereferencePolicy sets AliasDereferencePolicy field to given value.


### GetDescription

`func (o *PutAuthenticatorRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PutAuthenticatorRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PutAuthenticatorRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PutAuthenticatorRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *PutAuthenticatorRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PutAuthenticatorRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PutAuthenticatorRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetUserClaim

`func (o *PutAuthenticatorRequest) GetUserClaim() string`

GetUserClaim returns the UserClaim field if non-nil, zero value otherwise.

### GetUserClaimOk

`func (o *PutAuthenticatorRequest) GetUserClaimOk() (*string, bool)`

GetUserClaimOk returns a tuple with the UserClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserClaim

`func (o *PutAuthenticatorRequest) SetUserClaim(v string)`

SetUserClaim sets UserClaim field to given value.


### GetGroupClaim

`func (o *PutAuthenticatorRequest) GetGroupClaim() string`

GetGroupClaim returns the GroupClaim field if non-nil, zero value otherwise.

### GetGroupClaimOk

`func (o *PutAuthenticatorRequest) GetGroupClaimOk() (*string, bool)`

GetGroupClaimOk returns a tuple with the GroupClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupClaim

`func (o *PutAuthenticatorRequest) SetGroupClaim(v string)`

SetGroupClaim sets GroupClaim field to given value.


### GetEmailClaim

`func (o *PutAuthenticatorRequest) GetEmailClaim() string`

GetEmailClaim returns the EmailClaim field if non-nil, zero value otherwise.

### GetEmailClaimOk

`func (o *PutAuthenticatorRequest) GetEmailClaimOk() (*string, bool)`

GetEmailClaimOk returns a tuple with the EmailClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailClaim

`func (o *PutAuthenticatorRequest) SetEmailClaim(v string)`

SetEmailClaim sets EmailClaim field to given value.


### GetValidationMethod

`func (o *PutAuthenticatorRequest) GetValidationMethod() string`

GetValidationMethod returns the ValidationMethod field if non-nil, zero value otherwise.

### GetValidationMethodOk

`func (o *PutAuthenticatorRequest) GetValidationMethodOk() (*string, bool)`

GetValidationMethodOk returns a tuple with the ValidationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationMethod

`func (o *PutAuthenticatorRequest) SetValidationMethod(v string)`

SetValidationMethod sets ValidationMethod field to given value.


### GetClientId

`func (o *PutAuthenticatorRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *PutAuthenticatorRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *PutAuthenticatorRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *PutAuthenticatorRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *PutAuthenticatorRequest) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *PutAuthenticatorRequest) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *PutAuthenticatorRequest) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *PutAuthenticatorRequest) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetIntrospectionUrl

`func (o *PutAuthenticatorRequest) GetIntrospectionUrl() string`

GetIntrospectionUrl returns the IntrospectionUrl field if non-nil, zero value otherwise.

### GetIntrospectionUrlOk

`func (o *PutAuthenticatorRequest) GetIntrospectionUrlOk() (*string, bool)`

GetIntrospectionUrlOk returns a tuple with the IntrospectionUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntrospectionUrl

`func (o *PutAuthenticatorRequest) SetIntrospectionUrl(v string)`

SetIntrospectionUrl sets IntrospectionUrl field to given value.

### HasIntrospectionUrl

`func (o *PutAuthenticatorRequest) HasIntrospectionUrl() bool`

HasIntrospectionUrl returns a boolean if a field has been set.

### GetUserInfoUrl

`func (o *PutAuthenticatorRequest) GetUserInfoUrl() string`

GetUserInfoUrl returns the UserInfoUrl field if non-nil, zero value otherwise.

### GetUserInfoUrlOk

`func (o *PutAuthenticatorRequest) GetUserInfoUrlOk() (*string, bool)`

GetUserInfoUrlOk returns a tuple with the UserInfoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInfoUrl

`func (o *PutAuthenticatorRequest) SetUserInfoUrl(v string)`

SetUserInfoUrl sets UserInfoUrl field to given value.

### HasUserInfoUrl

`func (o *PutAuthenticatorRequest) HasUserInfoUrl() bool`

HasUserInfoUrl returns a boolean if a field has been set.

### GetAuthorizationOption

`func (o *PutAuthenticatorRequest) GetAuthorizationOption() string`

GetAuthorizationOption returns the AuthorizationOption field if non-nil, zero value otherwise.

### GetAuthorizationOptionOk

`func (o *PutAuthenticatorRequest) GetAuthorizationOptionOk() (*string, bool)`

GetAuthorizationOptionOk returns a tuple with the AuthorizationOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationOption

`func (o *PutAuthenticatorRequest) SetAuthorizationOption(v string)`

SetAuthorizationOption sets AuthorizationOption field to given value.

### HasAuthorizationOption

`func (o *PutAuthenticatorRequest) HasAuthorizationOption() bool`

HasAuthorizationOption returns a boolean if a field has been set.

### GetAudience

`func (o *PutAuthenticatorRequest) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *PutAuthenticatorRequest) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *PutAuthenticatorRequest) SetAudience(v string)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *PutAuthenticatorRequest) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### GetIssuer

`func (o *PutAuthenticatorRequest) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *PutAuthenticatorRequest) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *PutAuthenticatorRequest) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *PutAuthenticatorRequest) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetSigningCertificate

`func (o *PutAuthenticatorRequest) GetSigningCertificate() string`

GetSigningCertificate returns the SigningCertificate field if non-nil, zero value otherwise.

### GetSigningCertificateOk

`func (o *PutAuthenticatorRequest) GetSigningCertificateOk() (*string, bool)`

GetSigningCertificateOk returns a tuple with the SigningCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningCertificate

`func (o *PutAuthenticatorRequest) SetSigningCertificate(v string)`

SetSigningCertificate sets SigningCertificate field to given value.


### GetTimeout

`func (o *PutAuthenticatorRequest) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *PutAuthenticatorRequest) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *PutAuthenticatorRequest) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.


### GetRetries

`func (o *PutAuthenticatorRequest) GetRetries() int32`

GetRetries returns the Retries field if non-nil, zero value otherwise.

### GetRetriesOk

`func (o *PutAuthenticatorRequest) GetRetriesOk() (*int32, bool)`

GetRetriesOk returns a tuple with the Retries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetries

`func (o *PutAuthenticatorRequest) SetRetries(v int32)`

SetRetries sets Retries field to given value.


### GetSharedSecret

`func (o *PutAuthenticatorRequest) GetSharedSecret() string`

GetSharedSecret returns the SharedSecret field if non-nil, zero value otherwise.

### GetSharedSecretOk

`func (o *PutAuthenticatorRequest) GetSharedSecretOk() (*string, bool)`

GetSharedSecretOk returns a tuple with the SharedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSecret

`func (o *PutAuthenticatorRequest) SetSharedSecret(v string)`

SetSharedSecret sets SharedSecret field to given value.

### HasSharedSecret

`func (o *PutAuthenticatorRequest) HasSharedSecret() bool`

HasSharedSecret returns a boolean if a field has been set.

### GetAuthenticationProtocol

`func (o *PutAuthenticatorRequest) GetAuthenticationProtocol() string`

GetAuthenticationProtocol returns the AuthenticationProtocol field if non-nil, zero value otherwise.

### GetAuthenticationProtocolOk

`func (o *PutAuthenticatorRequest) GetAuthenticationProtocolOk() (*string, bool)`

GetAuthenticationProtocolOk returns a tuple with the AuthenticationProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationProtocol

`func (o *PutAuthenticatorRequest) SetAuthenticationProtocol(v string)`

SetAuthenticationProtocol sets AuthenticationProtocol field to given value.


### GetGroupAttribute

`func (o *PutAuthenticatorRequest) GetGroupAttribute() string`

GetGroupAttribute returns the GroupAttribute field if non-nil, zero value otherwise.

### GetGroupAttributeOk

`func (o *PutAuthenticatorRequest) GetGroupAttributeOk() (*string, bool)`

GetGroupAttributeOk returns a tuple with the GroupAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupAttribute

`func (o *PutAuthenticatorRequest) SetGroupAttribute(v string)`

SetGroupAttribute sets GroupAttribute field to given value.


### GetEmailAttribute

`func (o *PutAuthenticatorRequest) GetEmailAttribute() string`

GetEmailAttribute returns the EmailAttribute field if non-nil, zero value otherwise.

### GetEmailAttributeOk

`func (o *PutAuthenticatorRequest) GetEmailAttributeOk() (*string, bool)`

GetEmailAttributeOk returns a tuple with the EmailAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAttribute

`func (o *PutAuthenticatorRequest) SetEmailAttribute(v string)`

SetEmailAttribute sets EmailAttribute field to given value.


### GetSignOnUrl

`func (o *PutAuthenticatorRequest) GetSignOnUrl() string`

GetSignOnUrl returns the SignOnUrl field if non-nil, zero value otherwise.

### GetSignOnUrlOk

`func (o *PutAuthenticatorRequest) GetSignOnUrlOk() (*string, bool)`

GetSignOnUrlOk returns a tuple with the SignOnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignOnUrl

`func (o *PutAuthenticatorRequest) SetSignOnUrl(v string)`

SetSignOnUrl sets SignOnUrl field to given value.


### GetLogoutUrl

`func (o *PutAuthenticatorRequest) GetLogoutUrl() string`

GetLogoutUrl returns the LogoutUrl field if non-nil, zero value otherwise.

### GetLogoutUrlOk

`func (o *PutAuthenticatorRequest) GetLogoutUrlOk() (*string, bool)`

GetLogoutUrlOk returns a tuple with the LogoutUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoutUrl

`func (o *PutAuthenticatorRequest) SetLogoutUrl(v string)`

SetLogoutUrl sets LogoutUrl field to given value.

### HasLogoutUrl

`func (o *PutAuthenticatorRequest) HasLogoutUrl() bool`

HasLogoutUrl returns a boolean if a field has been set.

### GetEntityId

`func (o *PutAuthenticatorRequest) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *PutAuthenticatorRequest) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *PutAuthenticatorRequest) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.


### GetAttributes

`func (o *PutAuthenticatorRequest) GetAttributes() []NameValuePairBean`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PutAuthenticatorRequest) GetAttributesOk() (*[]NameValuePairBean, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PutAuthenticatorRequest) SetAttributes(v []NameValuePairBean)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


