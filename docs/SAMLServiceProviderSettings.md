# SAMLServiceProviderSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The resource type. | [optional] 
**AddressManagerFqdn** | Pointer to **string** | The fully qualified domain name of the Address Manager server. | [optional] 
**MetadataUrl** | Pointer to **string** | The identifier of the service provider entity. | [optional] [readonly] 
**ConsumeUrl** | Pointer to **string** | The URL location where the response from the IdP will be returned to the service provider. | [optional] [readonly] 
**LogoutUrl** | Pointer to **string** | The URL location where the logout response message will be returned to the service provider. | [optional] [readonly] 
**NameIdFormat** | Pointer to **string** | The specified constraints on the name identifier format used to represent the requested subject. | [optional] 
**SigningEnabled** | Pointer to **bool** | Indicates whether the message sent by the service provider will be signed. | [optional] 
**EncryptionEnabled** | Pointer to **bool** | Indicates the requirement for assertions received by the service provider to be encrypted. | [optional] 
**Pkcs12** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**OrganizationName** | Pointer to **string** | The organization associated with the service provider. | [optional] 
**OrganizationUrl** | Pointer to **string** | The website URL for the organization associated with the service provider. | [optional] 
**ContactName** | Pointer to **string** | The name of the contact for the organization associated with the service provider. | [optional] 
**ContactEmail** | Pointer to **string** | The contact email for the organization associated with the service provider. | [optional] 

## Methods

### NewSAMLServiceProviderSettings

`func NewSAMLServiceProviderSettings() *SAMLServiceProviderSettings`

NewSAMLServiceProviderSettings instantiates a new SAMLServiceProviderSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSAMLServiceProviderSettingsWithDefaults

`func NewSAMLServiceProviderSettingsWithDefaults() *SAMLServiceProviderSettings`

NewSAMLServiceProviderSettingsWithDefaults instantiates a new SAMLServiceProviderSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SAMLServiceProviderSettings) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SAMLServiceProviderSettings) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SAMLServiceProviderSettings) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SAMLServiceProviderSettings) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAddressManagerFqdn

`func (o *SAMLServiceProviderSettings) GetAddressManagerFqdn() string`

GetAddressManagerFqdn returns the AddressManagerFqdn field if non-nil, zero value otherwise.

### GetAddressManagerFqdnOk

`func (o *SAMLServiceProviderSettings) GetAddressManagerFqdnOk() (*string, bool)`

GetAddressManagerFqdnOk returns a tuple with the AddressManagerFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressManagerFqdn

`func (o *SAMLServiceProviderSettings) SetAddressManagerFqdn(v string)`

SetAddressManagerFqdn sets AddressManagerFqdn field to given value.

### HasAddressManagerFqdn

`func (o *SAMLServiceProviderSettings) HasAddressManagerFqdn() bool`

HasAddressManagerFqdn returns a boolean if a field has been set.

### GetMetadataUrl

`func (o *SAMLServiceProviderSettings) GetMetadataUrl() string`

GetMetadataUrl returns the MetadataUrl field if non-nil, zero value otherwise.

### GetMetadataUrlOk

`func (o *SAMLServiceProviderSettings) GetMetadataUrlOk() (*string, bool)`

GetMetadataUrlOk returns a tuple with the MetadataUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataUrl

`func (o *SAMLServiceProviderSettings) SetMetadataUrl(v string)`

SetMetadataUrl sets MetadataUrl field to given value.

### HasMetadataUrl

`func (o *SAMLServiceProviderSettings) HasMetadataUrl() bool`

HasMetadataUrl returns a boolean if a field has been set.

### GetConsumeUrl

`func (o *SAMLServiceProviderSettings) GetConsumeUrl() string`

GetConsumeUrl returns the ConsumeUrl field if non-nil, zero value otherwise.

### GetConsumeUrlOk

`func (o *SAMLServiceProviderSettings) GetConsumeUrlOk() (*string, bool)`

GetConsumeUrlOk returns a tuple with the ConsumeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumeUrl

`func (o *SAMLServiceProviderSettings) SetConsumeUrl(v string)`

SetConsumeUrl sets ConsumeUrl field to given value.

### HasConsumeUrl

`func (o *SAMLServiceProviderSettings) HasConsumeUrl() bool`

HasConsumeUrl returns a boolean if a field has been set.

### GetLogoutUrl

`func (o *SAMLServiceProviderSettings) GetLogoutUrl() string`

GetLogoutUrl returns the LogoutUrl field if non-nil, zero value otherwise.

### GetLogoutUrlOk

`func (o *SAMLServiceProviderSettings) GetLogoutUrlOk() (*string, bool)`

GetLogoutUrlOk returns a tuple with the LogoutUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoutUrl

`func (o *SAMLServiceProviderSettings) SetLogoutUrl(v string)`

SetLogoutUrl sets LogoutUrl field to given value.

### HasLogoutUrl

`func (o *SAMLServiceProviderSettings) HasLogoutUrl() bool`

HasLogoutUrl returns a boolean if a field has been set.

### GetNameIdFormat

`func (o *SAMLServiceProviderSettings) GetNameIdFormat() string`

GetNameIdFormat returns the NameIdFormat field if non-nil, zero value otherwise.

### GetNameIdFormatOk

`func (o *SAMLServiceProviderSettings) GetNameIdFormatOk() (*string, bool)`

GetNameIdFormatOk returns a tuple with the NameIdFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameIdFormat

`func (o *SAMLServiceProviderSettings) SetNameIdFormat(v string)`

SetNameIdFormat sets NameIdFormat field to given value.

### HasNameIdFormat

`func (o *SAMLServiceProviderSettings) HasNameIdFormat() bool`

HasNameIdFormat returns a boolean if a field has been set.

### GetSigningEnabled

`func (o *SAMLServiceProviderSettings) GetSigningEnabled() bool`

GetSigningEnabled returns the SigningEnabled field if non-nil, zero value otherwise.

### GetSigningEnabledOk

`func (o *SAMLServiceProviderSettings) GetSigningEnabledOk() (*bool, bool)`

GetSigningEnabledOk returns a tuple with the SigningEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningEnabled

`func (o *SAMLServiceProviderSettings) SetSigningEnabled(v bool)`

SetSigningEnabled sets SigningEnabled field to given value.

### HasSigningEnabled

`func (o *SAMLServiceProviderSettings) HasSigningEnabled() bool`

HasSigningEnabled returns a boolean if a field has been set.

### GetEncryptionEnabled

`func (o *SAMLServiceProviderSettings) GetEncryptionEnabled() bool`

GetEncryptionEnabled returns the EncryptionEnabled field if non-nil, zero value otherwise.

### GetEncryptionEnabledOk

`func (o *SAMLServiceProviderSettings) GetEncryptionEnabledOk() (*bool, bool)`

GetEncryptionEnabledOk returns a tuple with the EncryptionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionEnabled

`func (o *SAMLServiceProviderSettings) SetEncryptionEnabled(v bool)`

SetEncryptionEnabled sets EncryptionEnabled field to given value.

### HasEncryptionEnabled

`func (o *SAMLServiceProviderSettings) HasEncryptionEnabled() bool`

HasEncryptionEnabled returns a boolean if a field has been set.

### GetPkcs12

`func (o *SAMLServiceProviderSettings) GetPkcs12() string`

GetPkcs12 returns the Pkcs12 field if non-nil, zero value otherwise.

### GetPkcs12Ok

`func (o *SAMLServiceProviderSettings) GetPkcs12Ok() (*string, bool)`

GetPkcs12Ok returns a tuple with the Pkcs12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkcs12

`func (o *SAMLServiceProviderSettings) SetPkcs12(v string)`

SetPkcs12 sets Pkcs12 field to given value.

### HasPkcs12

`func (o *SAMLServiceProviderSettings) HasPkcs12() bool`

HasPkcs12 returns a boolean if a field has been set.

### GetPassword

`func (o *SAMLServiceProviderSettings) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SAMLServiceProviderSettings) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SAMLServiceProviderSettings) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *SAMLServiceProviderSettings) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetOrganizationName

`func (o *SAMLServiceProviderSettings) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *SAMLServiceProviderSettings) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *SAMLServiceProviderSettings) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *SAMLServiceProviderSettings) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### GetOrganizationUrl

`func (o *SAMLServiceProviderSettings) GetOrganizationUrl() string`

GetOrganizationUrl returns the OrganizationUrl field if non-nil, zero value otherwise.

### GetOrganizationUrlOk

`func (o *SAMLServiceProviderSettings) GetOrganizationUrlOk() (*string, bool)`

GetOrganizationUrlOk returns a tuple with the OrganizationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationUrl

`func (o *SAMLServiceProviderSettings) SetOrganizationUrl(v string)`

SetOrganizationUrl sets OrganizationUrl field to given value.

### HasOrganizationUrl

`func (o *SAMLServiceProviderSettings) HasOrganizationUrl() bool`

HasOrganizationUrl returns a boolean if a field has been set.

### GetContactName

`func (o *SAMLServiceProviderSettings) GetContactName() string`

GetContactName returns the ContactName field if non-nil, zero value otherwise.

### GetContactNameOk

`func (o *SAMLServiceProviderSettings) GetContactNameOk() (*string, bool)`

GetContactNameOk returns a tuple with the ContactName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactName

`func (o *SAMLServiceProviderSettings) SetContactName(v string)`

SetContactName sets ContactName field to given value.

### HasContactName

`func (o *SAMLServiceProviderSettings) HasContactName() bool`

HasContactName returns a boolean if a field has been set.

### GetContactEmail

`func (o *SAMLServiceProviderSettings) GetContactEmail() string`

GetContactEmail returns the ContactEmail field if non-nil, zero value otherwise.

### GetContactEmailOk

`func (o *SAMLServiceProviderSettings) GetContactEmailOk() (*string, bool)`

GetContactEmailOk returns a tuple with the ContactEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactEmail

`func (o *SAMLServiceProviderSettings) SetContactEmail(v string)`

SetContactEmail sets ContactEmail field to given value.

### HasContactEmail

`func (o *SAMLServiceProviderSettings) HasContactEmail() bool`

HasContactEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


