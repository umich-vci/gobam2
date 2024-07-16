# OAuthAuthorizationServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of authenticator. | [optional] 
**Description** | Pointer to **string** | The description of the authorization server. | [optional] 
**Enabled** | Pointer to **bool** | Indicates whether OAuth authentication is enabled. | [optional] 
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
**SigningCertificate** | Pointer to **string** | The authoritative server signing certificate. | [optional] 

## Methods

### NewOAuthAuthorizationServer

`func NewOAuthAuthorizationServer() *OAuthAuthorizationServer`

NewOAuthAuthorizationServer instantiates a new OAuthAuthorizationServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuthAuthorizationServerWithDefaults

`func NewOAuthAuthorizationServerWithDefaults() *OAuthAuthorizationServer`

NewOAuthAuthorizationServerWithDefaults instantiates a new OAuthAuthorizationServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OAuthAuthorizationServer) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OAuthAuthorizationServer) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OAuthAuthorizationServer) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OAuthAuthorizationServer) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDescription

`func (o *OAuthAuthorizationServer) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OAuthAuthorizationServer) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OAuthAuthorizationServer) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OAuthAuthorizationServer) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *OAuthAuthorizationServer) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *OAuthAuthorizationServer) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *OAuthAuthorizationServer) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *OAuthAuthorizationServer) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetUserClaim

`func (o *OAuthAuthorizationServer) GetUserClaim() string`

GetUserClaim returns the UserClaim field if non-nil, zero value otherwise.

### GetUserClaimOk

`func (o *OAuthAuthorizationServer) GetUserClaimOk() (*string, bool)`

GetUserClaimOk returns a tuple with the UserClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserClaim

`func (o *OAuthAuthorizationServer) SetUserClaim(v string)`

SetUserClaim sets UserClaim field to given value.

### HasUserClaim

`func (o *OAuthAuthorizationServer) HasUserClaim() bool`

HasUserClaim returns a boolean if a field has been set.

### GetGroupClaim

`func (o *OAuthAuthorizationServer) GetGroupClaim() string`

GetGroupClaim returns the GroupClaim field if non-nil, zero value otherwise.

### GetGroupClaimOk

`func (o *OAuthAuthorizationServer) GetGroupClaimOk() (*string, bool)`

GetGroupClaimOk returns a tuple with the GroupClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupClaim

`func (o *OAuthAuthorizationServer) SetGroupClaim(v string)`

SetGroupClaim sets GroupClaim field to given value.

### HasGroupClaim

`func (o *OAuthAuthorizationServer) HasGroupClaim() bool`

HasGroupClaim returns a boolean if a field has been set.

### GetEmailClaim

`func (o *OAuthAuthorizationServer) GetEmailClaim() string`

GetEmailClaim returns the EmailClaim field if non-nil, zero value otherwise.

### GetEmailClaimOk

`func (o *OAuthAuthorizationServer) GetEmailClaimOk() (*string, bool)`

GetEmailClaimOk returns a tuple with the EmailClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailClaim

`func (o *OAuthAuthorizationServer) SetEmailClaim(v string)`

SetEmailClaim sets EmailClaim field to given value.

### HasEmailClaim

`func (o *OAuthAuthorizationServer) HasEmailClaim() bool`

HasEmailClaim returns a boolean if a field has been set.

### GetValidationMethod

`func (o *OAuthAuthorizationServer) GetValidationMethod() string`

GetValidationMethod returns the ValidationMethod field if non-nil, zero value otherwise.

### GetValidationMethodOk

`func (o *OAuthAuthorizationServer) GetValidationMethodOk() (*string, bool)`

GetValidationMethodOk returns a tuple with the ValidationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationMethod

`func (o *OAuthAuthorizationServer) SetValidationMethod(v string)`

SetValidationMethod sets ValidationMethod field to given value.

### HasValidationMethod

`func (o *OAuthAuthorizationServer) HasValidationMethod() bool`

HasValidationMethod returns a boolean if a field has been set.

### GetClientId

`func (o *OAuthAuthorizationServer) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *OAuthAuthorizationServer) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *OAuthAuthorizationServer) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *OAuthAuthorizationServer) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *OAuthAuthorizationServer) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *OAuthAuthorizationServer) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *OAuthAuthorizationServer) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *OAuthAuthorizationServer) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetIntrospectionUrl

`func (o *OAuthAuthorizationServer) GetIntrospectionUrl() string`

GetIntrospectionUrl returns the IntrospectionUrl field if non-nil, zero value otherwise.

### GetIntrospectionUrlOk

`func (o *OAuthAuthorizationServer) GetIntrospectionUrlOk() (*string, bool)`

GetIntrospectionUrlOk returns a tuple with the IntrospectionUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntrospectionUrl

`func (o *OAuthAuthorizationServer) SetIntrospectionUrl(v string)`

SetIntrospectionUrl sets IntrospectionUrl field to given value.

### HasIntrospectionUrl

`func (o *OAuthAuthorizationServer) HasIntrospectionUrl() bool`

HasIntrospectionUrl returns a boolean if a field has been set.

### GetUserInfoUrl

`func (o *OAuthAuthorizationServer) GetUserInfoUrl() string`

GetUserInfoUrl returns the UserInfoUrl field if non-nil, zero value otherwise.

### GetUserInfoUrlOk

`func (o *OAuthAuthorizationServer) GetUserInfoUrlOk() (*string, bool)`

GetUserInfoUrlOk returns a tuple with the UserInfoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInfoUrl

`func (o *OAuthAuthorizationServer) SetUserInfoUrl(v string)`

SetUserInfoUrl sets UserInfoUrl field to given value.

### HasUserInfoUrl

`func (o *OAuthAuthorizationServer) HasUserInfoUrl() bool`

HasUserInfoUrl returns a boolean if a field has been set.

### GetAuthorizationOption

`func (o *OAuthAuthorizationServer) GetAuthorizationOption() string`

GetAuthorizationOption returns the AuthorizationOption field if non-nil, zero value otherwise.

### GetAuthorizationOptionOk

`func (o *OAuthAuthorizationServer) GetAuthorizationOptionOk() (*string, bool)`

GetAuthorizationOptionOk returns a tuple with the AuthorizationOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationOption

`func (o *OAuthAuthorizationServer) SetAuthorizationOption(v string)`

SetAuthorizationOption sets AuthorizationOption field to given value.

### HasAuthorizationOption

`func (o *OAuthAuthorizationServer) HasAuthorizationOption() bool`

HasAuthorizationOption returns a boolean if a field has been set.

### GetAudience

`func (o *OAuthAuthorizationServer) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *OAuthAuthorizationServer) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *OAuthAuthorizationServer) SetAudience(v string)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *OAuthAuthorizationServer) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### GetIssuer

`func (o *OAuthAuthorizationServer) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *OAuthAuthorizationServer) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *OAuthAuthorizationServer) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *OAuthAuthorizationServer) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetSigningCertificate

`func (o *OAuthAuthorizationServer) GetSigningCertificate() string`

GetSigningCertificate returns the SigningCertificate field if non-nil, zero value otherwise.

### GetSigningCertificateOk

`func (o *OAuthAuthorizationServer) GetSigningCertificateOk() (*string, bool)`

GetSigningCertificateOk returns a tuple with the SigningCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningCertificate

`func (o *OAuthAuthorizationServer) SetSigningCertificate(v string)`

SetSigningCertificate sets SigningCertificate field to given value.

### HasSigningCertificate

`func (o *OAuthAuthorizationServer) HasSigningCertificate() bool`

HasSigningCertificate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


