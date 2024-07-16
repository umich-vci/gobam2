# SAMLIdentityProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of authenticator. | [optional] 
**Description** | Pointer to **string** | The description of the SAML IdP configuration. | [optional] 
**Enabled** | Pointer to **bool** | Indicates whether SAML IdP authentication is enabled. | [optional] 
**GroupAttribute** | Pointer to **string** | The attribute name for group in the SAML response. | [optional] 
**EmailAttribute** | Pointer to **string** | The attribute name for email in the SAML response. | [optional] 
**SignOnUrl** | Pointer to **string** | The sign-in URL of the SAMP IdP. | [optional] 
**LogoutUrl** | Pointer to **string** | The logout URL of the SAMP IdP. | [optional] 
**EntityId** | Pointer to **string** | The entity ID of the SAMP IdP. | [optional] 
**SigningCertificate** | Pointer to **string** | The SAML IdP server signing certificate. | [optional] 

## Methods

### NewSAMLIdentityProvider

`func NewSAMLIdentityProvider() *SAMLIdentityProvider`

NewSAMLIdentityProvider instantiates a new SAMLIdentityProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSAMLIdentityProviderWithDefaults

`func NewSAMLIdentityProviderWithDefaults() *SAMLIdentityProvider`

NewSAMLIdentityProviderWithDefaults instantiates a new SAMLIdentityProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SAMLIdentityProvider) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SAMLIdentityProvider) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SAMLIdentityProvider) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SAMLIdentityProvider) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDescription

`func (o *SAMLIdentityProvider) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SAMLIdentityProvider) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SAMLIdentityProvider) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SAMLIdentityProvider) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *SAMLIdentityProvider) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SAMLIdentityProvider) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SAMLIdentityProvider) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SAMLIdentityProvider) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetGroupAttribute

`func (o *SAMLIdentityProvider) GetGroupAttribute() string`

GetGroupAttribute returns the GroupAttribute field if non-nil, zero value otherwise.

### GetGroupAttributeOk

`func (o *SAMLIdentityProvider) GetGroupAttributeOk() (*string, bool)`

GetGroupAttributeOk returns a tuple with the GroupAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupAttribute

`func (o *SAMLIdentityProvider) SetGroupAttribute(v string)`

SetGroupAttribute sets GroupAttribute field to given value.

### HasGroupAttribute

`func (o *SAMLIdentityProvider) HasGroupAttribute() bool`

HasGroupAttribute returns a boolean if a field has been set.

### GetEmailAttribute

`func (o *SAMLIdentityProvider) GetEmailAttribute() string`

GetEmailAttribute returns the EmailAttribute field if non-nil, zero value otherwise.

### GetEmailAttributeOk

`func (o *SAMLIdentityProvider) GetEmailAttributeOk() (*string, bool)`

GetEmailAttributeOk returns a tuple with the EmailAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAttribute

`func (o *SAMLIdentityProvider) SetEmailAttribute(v string)`

SetEmailAttribute sets EmailAttribute field to given value.

### HasEmailAttribute

`func (o *SAMLIdentityProvider) HasEmailAttribute() bool`

HasEmailAttribute returns a boolean if a field has been set.

### GetSignOnUrl

`func (o *SAMLIdentityProvider) GetSignOnUrl() string`

GetSignOnUrl returns the SignOnUrl field if non-nil, zero value otherwise.

### GetSignOnUrlOk

`func (o *SAMLIdentityProvider) GetSignOnUrlOk() (*string, bool)`

GetSignOnUrlOk returns a tuple with the SignOnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignOnUrl

`func (o *SAMLIdentityProvider) SetSignOnUrl(v string)`

SetSignOnUrl sets SignOnUrl field to given value.

### HasSignOnUrl

`func (o *SAMLIdentityProvider) HasSignOnUrl() bool`

HasSignOnUrl returns a boolean if a field has been set.

### GetLogoutUrl

`func (o *SAMLIdentityProvider) GetLogoutUrl() string`

GetLogoutUrl returns the LogoutUrl field if non-nil, zero value otherwise.

### GetLogoutUrlOk

`func (o *SAMLIdentityProvider) GetLogoutUrlOk() (*string, bool)`

GetLogoutUrlOk returns a tuple with the LogoutUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoutUrl

`func (o *SAMLIdentityProvider) SetLogoutUrl(v string)`

SetLogoutUrl sets LogoutUrl field to given value.

### HasLogoutUrl

`func (o *SAMLIdentityProvider) HasLogoutUrl() bool`

HasLogoutUrl returns a boolean if a field has been set.

### GetEntityId

`func (o *SAMLIdentityProvider) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *SAMLIdentityProvider) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *SAMLIdentityProvider) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *SAMLIdentityProvider) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetSigningCertificate

`func (o *SAMLIdentityProvider) GetSigningCertificate() string`

GetSigningCertificate returns the SigningCertificate field if non-nil, zero value otherwise.

### GetSigningCertificateOk

`func (o *SAMLIdentityProvider) GetSigningCertificateOk() (*string, bool)`

GetSigningCertificateOk returns a tuple with the SigningCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningCertificate

`func (o *SAMLIdentityProvider) SetSigningCertificate(v string)`

SetSigningCertificate sets SigningCertificate field to given value.

### HasSigningCertificate

`func (o *SAMLIdentityProvider) HasSigningCertificate() bool`

HasSigningCertificate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


