# PostAuthenticator201Response

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
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 

## Methods

### NewPostAuthenticator201Response

`func NewPostAuthenticator201Response() *PostAuthenticator201Response`

NewPostAuthenticator201Response instantiates a new PostAuthenticator201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostAuthenticator201ResponseWithDefaults

`func NewPostAuthenticator201ResponseWithDefaults() *PostAuthenticator201Response`

NewPostAuthenticator201ResponseWithDefaults instantiates a new PostAuthenticator201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PostAuthenticator201Response) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostAuthenticator201Response) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostAuthenticator201Response) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *PostAuthenticator201Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *PostAuthenticator201Response) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PostAuthenticator201Response) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PostAuthenticator201Response) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PostAuthenticator201Response) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *PostAuthenticator201Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostAuthenticator201Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostAuthenticator201Response) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PostAuthenticator201Response) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *PostAuthenticator201Response) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *PostAuthenticator201Response) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *PostAuthenticator201Response) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *PostAuthenticator201Response) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetSecondaryAuthenticator

`func (o *PostAuthenticator201Response) GetSecondaryAuthenticator() InlinedAuthenticator`

GetSecondaryAuthenticator returns the SecondaryAuthenticator field if non-nil, zero value otherwise.

### GetSecondaryAuthenticatorOk

`func (o *PostAuthenticator201Response) GetSecondaryAuthenticatorOk() (*InlinedAuthenticator, bool)`

GetSecondaryAuthenticatorOk returns a tuple with the SecondaryAuthenticator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryAuthenticator

`func (o *PostAuthenticator201Response) SetSecondaryAuthenticator(v InlinedAuthenticator)`

SetSecondaryAuthenticator sets SecondaryAuthenticator field to given value.

### HasSecondaryAuthenticator

`func (o *PostAuthenticator201Response) HasSecondaryAuthenticator() bool`

HasSecondaryAuthenticator returns a boolean if a field has been set.

### GetHostname

`func (o *PostAuthenticator201Response) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *PostAuthenticator201Response) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *PostAuthenticator201Response) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *PostAuthenticator201Response) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetRealm

`func (o *PostAuthenticator201Response) GetRealm() string`

GetRealm returns the Realm field if non-nil, zero value otherwise.

### GetRealmOk

`func (o *PostAuthenticator201Response) GetRealmOk() (*string, bool)`

GetRealmOk returns a tuple with the Realm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealm

`func (o *PostAuthenticator201Response) SetRealm(v string)`

SetRealm sets Realm field to given value.

### HasRealm

`func (o *PostAuthenticator201Response) HasRealm() bool`

HasRealm returns a boolean if a field has been set.

### GetPort

`func (o *PostAuthenticator201Response) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *PostAuthenticator201Response) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *PostAuthenticator201Response) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *PostAuthenticator201Response) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetAdminDn

`func (o *PostAuthenticator201Response) GetAdminDn() string`

GetAdminDn returns the AdminDn field if non-nil, zero value otherwise.

### GetAdminDnOk

`func (o *PostAuthenticator201Response) GetAdminDnOk() (*string, bool)`

GetAdminDnOk returns a tuple with the AdminDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminDn

`func (o *PostAuthenticator201Response) SetAdminDn(v string)`

SetAdminDn sets AdminDn field to given value.

### HasAdminDn

`func (o *PostAuthenticator201Response) HasAdminDn() bool`

HasAdminDn returns a boolean if a field has been set.

### GetAdminPassword

`func (o *PostAuthenticator201Response) GetAdminPassword() string`

GetAdminPassword returns the AdminPassword field if non-nil, zero value otherwise.

### GetAdminPasswordOk

`func (o *PostAuthenticator201Response) GetAdminPasswordOk() (*string, bool)`

GetAdminPasswordOk returns a tuple with the AdminPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminPassword

`func (o *PostAuthenticator201Response) SetAdminPassword(v string)`

SetAdminPassword sets AdminPassword field to given value.

### HasAdminPassword

`func (o *PostAuthenticator201Response) HasAdminPassword() bool`

HasAdminPassword returns a boolean if a field has been set.

### GetSslEnabled

`func (o *PostAuthenticator201Response) GetSslEnabled() bool`

GetSslEnabled returns the SslEnabled field if non-nil, zero value otherwise.

### GetSslEnabledOk

`func (o *PostAuthenticator201Response) GetSslEnabledOk() (*bool, bool)`

GetSslEnabledOk returns a tuple with the SslEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslEnabled

`func (o *PostAuthenticator201Response) SetSslEnabled(v bool)`

SetSslEnabled sets SslEnabled field to given value.

### HasSslEnabled

`func (o *PostAuthenticator201Response) HasSslEnabled() bool`

HasSslEnabled returns a boolean if a field has been set.

### GetSchemaType

`func (o *PostAuthenticator201Response) GetSchemaType() string`

GetSchemaType returns the SchemaType field if non-nil, zero value otherwise.

### GetSchemaTypeOk

`func (o *PostAuthenticator201Response) GetSchemaTypeOk() (*string, bool)`

GetSchemaTypeOk returns a tuple with the SchemaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaType

`func (o *PostAuthenticator201Response) SetSchemaType(v string)`

SetSchemaType sets SchemaType field to given value.

### HasSchemaType

`func (o *PostAuthenticator201Response) HasSchemaType() bool`

HasSchemaType returns a boolean if a field has been set.

### GetSearchBase

`func (o *PostAuthenticator201Response) GetSearchBase() string`

GetSearchBase returns the SearchBase field if non-nil, zero value otherwise.

### GetSearchBaseOk

`func (o *PostAuthenticator201Response) GetSearchBaseOk() (*string, bool)`

GetSearchBaseOk returns a tuple with the SearchBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchBase

`func (o *PostAuthenticator201Response) SetSearchBase(v string)`

SetSearchBase sets SearchBase field to given value.

### HasSearchBase

`func (o *PostAuthenticator201Response) HasSearchBase() bool`

HasSearchBase returns a boolean if a field has been set.

### GetUserObjectClass

`func (o *PostAuthenticator201Response) GetUserObjectClass() string`

GetUserObjectClass returns the UserObjectClass field if non-nil, zero value otherwise.

### GetUserObjectClassOk

`func (o *PostAuthenticator201Response) GetUserObjectClassOk() (*string, bool)`

GetUserObjectClassOk returns a tuple with the UserObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserObjectClass

`func (o *PostAuthenticator201Response) SetUserObjectClass(v string)`

SetUserObjectClass sets UserObjectClass field to given value.

### HasUserObjectClass

`func (o *PostAuthenticator201Response) HasUserObjectClass() bool`

HasUserObjectClass returns a boolean if a field has been set.

### GetGroupObjectClass

`func (o *PostAuthenticator201Response) GetGroupObjectClass() string`

GetGroupObjectClass returns the GroupObjectClass field if non-nil, zero value otherwise.

### GetGroupObjectClassOk

`func (o *PostAuthenticator201Response) GetGroupObjectClassOk() (*string, bool)`

GetGroupObjectClassOk returns a tuple with the GroupObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupObjectClass

`func (o *PostAuthenticator201Response) SetGroupObjectClass(v string)`

SetGroupObjectClass sets GroupObjectClass field to given value.

### HasGroupObjectClass

`func (o *PostAuthenticator201Response) HasGroupObjectClass() bool`

HasGroupObjectClass returns a boolean if a field has been set.

### GetUserPrefix

`func (o *PostAuthenticator201Response) GetUserPrefix() string`

GetUserPrefix returns the UserPrefix field if non-nil, zero value otherwise.

### GetUserPrefixOk

`func (o *PostAuthenticator201Response) GetUserPrefixOk() (*string, bool)`

GetUserPrefixOk returns a tuple with the UserPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPrefix

`func (o *PostAuthenticator201Response) SetUserPrefix(v string)`

SetUserPrefix sets UserPrefix field to given value.

### HasUserPrefix

`func (o *PostAuthenticator201Response) HasUserPrefix() bool`

HasUserPrefix returns a boolean if a field has been set.

### GetEmailPrefix

`func (o *PostAuthenticator201Response) GetEmailPrefix() string`

GetEmailPrefix returns the EmailPrefix field if non-nil, zero value otherwise.

### GetEmailPrefixOk

`func (o *PostAuthenticator201Response) GetEmailPrefixOk() (*string, bool)`

GetEmailPrefixOk returns a tuple with the EmailPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailPrefix

`func (o *PostAuthenticator201Response) SetEmailPrefix(v string)`

SetEmailPrefix sets EmailPrefix field to given value.

### HasEmailPrefix

`func (o *PostAuthenticator201Response) HasEmailPrefix() bool`

HasEmailPrefix returns a boolean if a field has been set.

### GetMemberOfPrefix

`func (o *PostAuthenticator201Response) GetMemberOfPrefix() string`

GetMemberOfPrefix returns the MemberOfPrefix field if non-nil, zero value otherwise.

### GetMemberOfPrefixOk

`func (o *PostAuthenticator201Response) GetMemberOfPrefixOk() (*string, bool)`

GetMemberOfPrefixOk returns a tuple with the MemberOfPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberOfPrefix

`func (o *PostAuthenticator201Response) SetMemberOfPrefix(v string)`

SetMemberOfPrefix sets MemberOfPrefix field to given value.

### HasMemberOfPrefix

`func (o *PostAuthenticator201Response) HasMemberOfPrefix() bool`

HasMemberOfPrefix returns a boolean if a field has been set.

### GetReferralPolicy

`func (o *PostAuthenticator201Response) GetReferralPolicy() string`

GetReferralPolicy returns the ReferralPolicy field if non-nil, zero value otherwise.

### GetReferralPolicyOk

`func (o *PostAuthenticator201Response) GetReferralPolicyOk() (*string, bool)`

GetReferralPolicyOk returns a tuple with the ReferralPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferralPolicy

`func (o *PostAuthenticator201Response) SetReferralPolicy(v string)`

SetReferralPolicy sets ReferralPolicy field to given value.

### HasReferralPolicy

`func (o *PostAuthenticator201Response) HasReferralPolicy() bool`

HasReferralPolicy returns a boolean if a field has been set.

### GetAliasDereferencePolicy

`func (o *PostAuthenticator201Response) GetAliasDereferencePolicy() string`

GetAliasDereferencePolicy returns the AliasDereferencePolicy field if non-nil, zero value otherwise.

### GetAliasDereferencePolicyOk

`func (o *PostAuthenticator201Response) GetAliasDereferencePolicyOk() (*string, bool)`

GetAliasDereferencePolicyOk returns a tuple with the AliasDereferencePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasDereferencePolicy

`func (o *PostAuthenticator201Response) SetAliasDereferencePolicy(v string)`

SetAliasDereferencePolicy sets AliasDereferencePolicy field to given value.

### HasAliasDereferencePolicy

`func (o *PostAuthenticator201Response) HasAliasDereferencePolicy() bool`

HasAliasDereferencePolicy returns a boolean if a field has been set.

### GetDescription

`func (o *PostAuthenticator201Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PostAuthenticator201Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PostAuthenticator201Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PostAuthenticator201Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *PostAuthenticator201Response) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PostAuthenticator201Response) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PostAuthenticator201Response) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PostAuthenticator201Response) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetUserClaim

`func (o *PostAuthenticator201Response) GetUserClaim() string`

GetUserClaim returns the UserClaim field if non-nil, zero value otherwise.

### GetUserClaimOk

`func (o *PostAuthenticator201Response) GetUserClaimOk() (*string, bool)`

GetUserClaimOk returns a tuple with the UserClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserClaim

`func (o *PostAuthenticator201Response) SetUserClaim(v string)`

SetUserClaim sets UserClaim field to given value.

### HasUserClaim

`func (o *PostAuthenticator201Response) HasUserClaim() bool`

HasUserClaim returns a boolean if a field has been set.

### GetGroupClaim

`func (o *PostAuthenticator201Response) GetGroupClaim() string`

GetGroupClaim returns the GroupClaim field if non-nil, zero value otherwise.

### GetGroupClaimOk

`func (o *PostAuthenticator201Response) GetGroupClaimOk() (*string, bool)`

GetGroupClaimOk returns a tuple with the GroupClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupClaim

`func (o *PostAuthenticator201Response) SetGroupClaim(v string)`

SetGroupClaim sets GroupClaim field to given value.

### HasGroupClaim

`func (o *PostAuthenticator201Response) HasGroupClaim() bool`

HasGroupClaim returns a boolean if a field has been set.

### GetEmailClaim

`func (o *PostAuthenticator201Response) GetEmailClaim() string`

GetEmailClaim returns the EmailClaim field if non-nil, zero value otherwise.

### GetEmailClaimOk

`func (o *PostAuthenticator201Response) GetEmailClaimOk() (*string, bool)`

GetEmailClaimOk returns a tuple with the EmailClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailClaim

`func (o *PostAuthenticator201Response) SetEmailClaim(v string)`

SetEmailClaim sets EmailClaim field to given value.

### HasEmailClaim

`func (o *PostAuthenticator201Response) HasEmailClaim() bool`

HasEmailClaim returns a boolean if a field has been set.

### GetValidationMethod

`func (o *PostAuthenticator201Response) GetValidationMethod() string`

GetValidationMethod returns the ValidationMethod field if non-nil, zero value otherwise.

### GetValidationMethodOk

`func (o *PostAuthenticator201Response) GetValidationMethodOk() (*string, bool)`

GetValidationMethodOk returns a tuple with the ValidationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationMethod

`func (o *PostAuthenticator201Response) SetValidationMethod(v string)`

SetValidationMethod sets ValidationMethod field to given value.

### HasValidationMethod

`func (o *PostAuthenticator201Response) HasValidationMethod() bool`

HasValidationMethod returns a boolean if a field has been set.

### GetClientId

`func (o *PostAuthenticator201Response) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *PostAuthenticator201Response) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *PostAuthenticator201Response) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *PostAuthenticator201Response) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *PostAuthenticator201Response) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *PostAuthenticator201Response) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *PostAuthenticator201Response) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *PostAuthenticator201Response) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetIntrospectionUrl

`func (o *PostAuthenticator201Response) GetIntrospectionUrl() string`

GetIntrospectionUrl returns the IntrospectionUrl field if non-nil, zero value otherwise.

### GetIntrospectionUrlOk

`func (o *PostAuthenticator201Response) GetIntrospectionUrlOk() (*string, bool)`

GetIntrospectionUrlOk returns a tuple with the IntrospectionUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntrospectionUrl

`func (o *PostAuthenticator201Response) SetIntrospectionUrl(v string)`

SetIntrospectionUrl sets IntrospectionUrl field to given value.

### HasIntrospectionUrl

`func (o *PostAuthenticator201Response) HasIntrospectionUrl() bool`

HasIntrospectionUrl returns a boolean if a field has been set.

### GetUserInfoUrl

`func (o *PostAuthenticator201Response) GetUserInfoUrl() string`

GetUserInfoUrl returns the UserInfoUrl field if non-nil, zero value otherwise.

### GetUserInfoUrlOk

`func (o *PostAuthenticator201Response) GetUserInfoUrlOk() (*string, bool)`

GetUserInfoUrlOk returns a tuple with the UserInfoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInfoUrl

`func (o *PostAuthenticator201Response) SetUserInfoUrl(v string)`

SetUserInfoUrl sets UserInfoUrl field to given value.

### HasUserInfoUrl

`func (o *PostAuthenticator201Response) HasUserInfoUrl() bool`

HasUserInfoUrl returns a boolean if a field has been set.

### GetAuthorizationOption

`func (o *PostAuthenticator201Response) GetAuthorizationOption() string`

GetAuthorizationOption returns the AuthorizationOption field if non-nil, zero value otherwise.

### GetAuthorizationOptionOk

`func (o *PostAuthenticator201Response) GetAuthorizationOptionOk() (*string, bool)`

GetAuthorizationOptionOk returns a tuple with the AuthorizationOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationOption

`func (o *PostAuthenticator201Response) SetAuthorizationOption(v string)`

SetAuthorizationOption sets AuthorizationOption field to given value.

### HasAuthorizationOption

`func (o *PostAuthenticator201Response) HasAuthorizationOption() bool`

HasAuthorizationOption returns a boolean if a field has been set.

### GetAudience

`func (o *PostAuthenticator201Response) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *PostAuthenticator201Response) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *PostAuthenticator201Response) SetAudience(v string)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *PostAuthenticator201Response) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### GetIssuer

`func (o *PostAuthenticator201Response) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *PostAuthenticator201Response) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *PostAuthenticator201Response) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *PostAuthenticator201Response) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetSigningCertificate

`func (o *PostAuthenticator201Response) GetSigningCertificate() string`

GetSigningCertificate returns the SigningCertificate field if non-nil, zero value otherwise.

### GetSigningCertificateOk

`func (o *PostAuthenticator201Response) GetSigningCertificateOk() (*string, bool)`

GetSigningCertificateOk returns a tuple with the SigningCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningCertificate

`func (o *PostAuthenticator201Response) SetSigningCertificate(v string)`

SetSigningCertificate sets SigningCertificate field to given value.

### HasSigningCertificate

`func (o *PostAuthenticator201Response) HasSigningCertificate() bool`

HasSigningCertificate returns a boolean if a field has been set.

### GetTimeout

`func (o *PostAuthenticator201Response) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *PostAuthenticator201Response) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *PostAuthenticator201Response) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *PostAuthenticator201Response) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetRetries

`func (o *PostAuthenticator201Response) GetRetries() int32`

GetRetries returns the Retries field if non-nil, zero value otherwise.

### GetRetriesOk

`func (o *PostAuthenticator201Response) GetRetriesOk() (*int32, bool)`

GetRetriesOk returns a tuple with the Retries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetries

`func (o *PostAuthenticator201Response) SetRetries(v int32)`

SetRetries sets Retries field to given value.

### HasRetries

`func (o *PostAuthenticator201Response) HasRetries() bool`

HasRetries returns a boolean if a field has been set.

### GetSharedSecret

`func (o *PostAuthenticator201Response) GetSharedSecret() string`

GetSharedSecret returns the SharedSecret field if non-nil, zero value otherwise.

### GetSharedSecretOk

`func (o *PostAuthenticator201Response) GetSharedSecretOk() (*string, bool)`

GetSharedSecretOk returns a tuple with the SharedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSecret

`func (o *PostAuthenticator201Response) SetSharedSecret(v string)`

SetSharedSecret sets SharedSecret field to given value.

### HasSharedSecret

`func (o *PostAuthenticator201Response) HasSharedSecret() bool`

HasSharedSecret returns a boolean if a field has been set.

### GetAuthenticationProtocol

`func (o *PostAuthenticator201Response) GetAuthenticationProtocol() string`

GetAuthenticationProtocol returns the AuthenticationProtocol field if non-nil, zero value otherwise.

### GetAuthenticationProtocolOk

`func (o *PostAuthenticator201Response) GetAuthenticationProtocolOk() (*string, bool)`

GetAuthenticationProtocolOk returns a tuple with the AuthenticationProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationProtocol

`func (o *PostAuthenticator201Response) SetAuthenticationProtocol(v string)`

SetAuthenticationProtocol sets AuthenticationProtocol field to given value.

### HasAuthenticationProtocol

`func (o *PostAuthenticator201Response) HasAuthenticationProtocol() bool`

HasAuthenticationProtocol returns a boolean if a field has been set.

### GetGroupAttribute

`func (o *PostAuthenticator201Response) GetGroupAttribute() string`

GetGroupAttribute returns the GroupAttribute field if non-nil, zero value otherwise.

### GetGroupAttributeOk

`func (o *PostAuthenticator201Response) GetGroupAttributeOk() (*string, bool)`

GetGroupAttributeOk returns a tuple with the GroupAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupAttribute

`func (o *PostAuthenticator201Response) SetGroupAttribute(v string)`

SetGroupAttribute sets GroupAttribute field to given value.

### HasGroupAttribute

`func (o *PostAuthenticator201Response) HasGroupAttribute() bool`

HasGroupAttribute returns a boolean if a field has been set.

### GetEmailAttribute

`func (o *PostAuthenticator201Response) GetEmailAttribute() string`

GetEmailAttribute returns the EmailAttribute field if non-nil, zero value otherwise.

### GetEmailAttributeOk

`func (o *PostAuthenticator201Response) GetEmailAttributeOk() (*string, bool)`

GetEmailAttributeOk returns a tuple with the EmailAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAttribute

`func (o *PostAuthenticator201Response) SetEmailAttribute(v string)`

SetEmailAttribute sets EmailAttribute field to given value.

### HasEmailAttribute

`func (o *PostAuthenticator201Response) HasEmailAttribute() bool`

HasEmailAttribute returns a boolean if a field has been set.

### GetSignOnUrl

`func (o *PostAuthenticator201Response) GetSignOnUrl() string`

GetSignOnUrl returns the SignOnUrl field if non-nil, zero value otherwise.

### GetSignOnUrlOk

`func (o *PostAuthenticator201Response) GetSignOnUrlOk() (*string, bool)`

GetSignOnUrlOk returns a tuple with the SignOnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignOnUrl

`func (o *PostAuthenticator201Response) SetSignOnUrl(v string)`

SetSignOnUrl sets SignOnUrl field to given value.

### HasSignOnUrl

`func (o *PostAuthenticator201Response) HasSignOnUrl() bool`

HasSignOnUrl returns a boolean if a field has been set.

### GetLogoutUrl

`func (o *PostAuthenticator201Response) GetLogoutUrl() string`

GetLogoutUrl returns the LogoutUrl field if non-nil, zero value otherwise.

### GetLogoutUrlOk

`func (o *PostAuthenticator201Response) GetLogoutUrlOk() (*string, bool)`

GetLogoutUrlOk returns a tuple with the LogoutUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoutUrl

`func (o *PostAuthenticator201Response) SetLogoutUrl(v string)`

SetLogoutUrl sets LogoutUrl field to given value.

### HasLogoutUrl

`func (o *PostAuthenticator201Response) HasLogoutUrl() bool`

HasLogoutUrl returns a boolean if a field has been set.

### GetEntityId

`func (o *PostAuthenticator201Response) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *PostAuthenticator201Response) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *PostAuthenticator201Response) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *PostAuthenticator201Response) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetAttributes

`func (o *PostAuthenticator201Response) GetAttributes() []NameValuePairBean`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PostAuthenticator201Response) GetAttributesOk() (*[]NameValuePairBean, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PostAuthenticator201Response) SetAttributes(v []NameValuePairBean)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *PostAuthenticator201Response) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetLinks

`func (o *PostAuthenticator201Response) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PostAuthenticator201Response) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PostAuthenticator201Response) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PostAuthenticator201Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


