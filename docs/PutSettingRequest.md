# PutSettingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | **string** | The resource type. | 
**MandatoryCommentsEnabled** | **bool** | Indicates whether users must provide a change control comment when creating, editing, or deleting resources. | 
**RememberLastAddressEnabled** | **bool** | Indicates whether the Address Manager UI Quick Actions widget will automatically provide the subsequent IPv4 address when adding host records. | 
**SessionTimeout** | **string** | The maximum time period of user inactivity before a session is terminated. | 
**DisclaimerEnabled** | **bool** | Indicates whether a disclaimer will display on the Address Manager UI login page with the contents of disclaimerText. | 
**DisclaimerText** | Pointer to **string** | The disclaimer text or HTML for display on the Address Manager UI login page. If adding or editing disclaimer HTML through JSON, ensure that reserved characters are escaped. | [optional] 
**CustomReverseZoneFormatsAllowed** | **bool** | Indicates whether users can set a custom reverse zone name format in DNS deployment options. | 
**AddressManager** | **string** |  | 
**RestV2** | **string** |  | 
**ZoneImport** | **string** |  | 
**ReconciliationService** | **string** |  | 
**DiscoveryEngine** | **string** |  | 
**SnmpClient** | **string** |  | 
**MonitoringService** | **string** |  | 
**RrdTool** | **string** |  | 
**Enabled** | **bool** |  | 
**FailureLimit** | **int32** | The limit for consecutive failed login attempts that results in a login delay. | 
**FailureLimitPeriod** | **string** | The timespan in which consecutive failed login attempts count towards the failure limit. | 
**DelayDuration** | **string** | The amount of time that a user must wait before they can attempt to log in after exceeding the login policy failure conditions. | 
**AddressManagerMonitoring** | [**AddressManagerMonitoringBean**](AddressManagerMonitoringBean.md) |  | 
**ServerMonitoring** | [**ServerMonitoringBean**](ServerMonitoringBean.md) |  | 
**F5ServerMonitoring** | Pointer to [**F5ServerMonitoringBean**](F5ServerMonitoringBean.md) |  | [optional] 
**MinLength** | Pointer to **int32** | The minimum length that the password must be. | [optional] 
**MaxLength** | Pointer to **int32** | The maximum length that the password can be. | [optional] 
**DigitRequired** | **bool** | Indicates whether the password must contain at least one digit. | 
**MixedCaseRequired** | **bool** | Indicates whether the password must contain mixed case letters. | 
**SpecialCharacterRequired** | **bool** | Indicates whether the password must contain at least one special character. | 
**AddressManagerFqdn** | **string** | The fully qualified domain name of the Address Manager server. | 
**MetadataUrl** | Pointer to **string** | The identifier of the service provider entity. | [optional] [readonly] 
**ConsumeUrl** | Pointer to **string** | The URL location where the response from the IdP will be returned to the service provider. | [optional] [readonly] 
**LogoutUrl** | Pointer to **string** | The URL location where the logout response message will be returned to the service provider. | [optional] [readonly] 
**NameIdFormat** | **string** | The specified constraints on the name identifier format used to represent the requested subject. | 
**SigningEnabled** | **bool** | Indicates whether the message sent by the service provider will be signed. | 
**EncryptionEnabled** | **bool** | Indicates the requirement for assertions received by the service provider to be encrypted. | 
**Pkcs12** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**OrganizationName** | **string** | The organization associated with the service provider. | 
**OrganizationUrl** | **string** | The website URL for the organization associated with the service provider. | 
**ContactName** | Pointer to **string** | The name of the contact for the organization associated with the service provider. | [optional] 
**ContactEmail** | Pointer to **string** | The contact email for the organization associated with the service provider. | [optional] 
**SamlIdentityProviderEnabled** | Pointer to **bool** |  | [optional] [readonly] 
**NonSsoAuthenticatorCount** | Pointer to **int32** |  | [optional] [readonly] 
**NonSsoGroupCount** | Pointer to **int32** |  | [optional] [readonly] 
**LocalAdminUserCount** | Pointer to **int32** |  | [optional] [readonly] 
**LocalNonAdminUserCount** | Pointer to **int32** |  | [optional] [readonly] 
**Hostname** | Pointer to **string** | The hostname of the Address Manager server. | [optional] [readonly] 
**Version** | Pointer to **string** | The Address Manager server version. | [optional] [readonly] 
**Address** | Pointer to **string** | The management IP address of the Address Manager server. | [optional] [readonly] 
**InterfaceRedundancyEnabled** | **bool** | Indicates whether Address Manager Active/Backup (eth0/eth1) bonding mode is active. | 
**ActiveSessions** | Pointer to **int32** | The current number of active Address Manager sessions. | [optional] [readonly] 

## Methods

### NewPutSettingRequest

`func NewPutSettingRequest(type_ string, mandatoryCommentsEnabled bool, rememberLastAddressEnabled bool, sessionTimeout string, disclaimerEnabled bool, customReverseZoneFormatsAllowed bool, addressManager string, restV2 string, zoneImport string, reconciliationService string, discoveryEngine string, snmpClient string, monitoringService string, rrdTool string, enabled bool, failureLimit int32, failureLimitPeriod string, delayDuration string, addressManagerMonitoring AddressManagerMonitoringBean, serverMonitoring ServerMonitoringBean, digitRequired bool, mixedCaseRequired bool, specialCharacterRequired bool, addressManagerFqdn string, nameIdFormat string, signingEnabled bool, encryptionEnabled bool, organizationName string, organizationUrl string, interfaceRedundancyEnabled bool, ) *PutSettingRequest`

NewPutSettingRequest instantiates a new PutSettingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutSettingRequestWithDefaults

`func NewPutSettingRequestWithDefaults() *PutSettingRequest`

NewPutSettingRequestWithDefaults instantiates a new PutSettingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PutSettingRequest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PutSettingRequest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PutSettingRequest) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *PutSettingRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *PutSettingRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PutSettingRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PutSettingRequest) SetType(v string)`

SetType sets Type field to given value.


### GetMandatoryCommentsEnabled

`func (o *PutSettingRequest) GetMandatoryCommentsEnabled() bool`

GetMandatoryCommentsEnabled returns the MandatoryCommentsEnabled field if non-nil, zero value otherwise.

### GetMandatoryCommentsEnabledOk

`func (o *PutSettingRequest) GetMandatoryCommentsEnabledOk() (*bool, bool)`

GetMandatoryCommentsEnabledOk returns a tuple with the MandatoryCommentsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandatoryCommentsEnabled

`func (o *PutSettingRequest) SetMandatoryCommentsEnabled(v bool)`

SetMandatoryCommentsEnabled sets MandatoryCommentsEnabled field to given value.


### GetRememberLastAddressEnabled

`func (o *PutSettingRequest) GetRememberLastAddressEnabled() bool`

GetRememberLastAddressEnabled returns the RememberLastAddressEnabled field if non-nil, zero value otherwise.

### GetRememberLastAddressEnabledOk

`func (o *PutSettingRequest) GetRememberLastAddressEnabledOk() (*bool, bool)`

GetRememberLastAddressEnabledOk returns a tuple with the RememberLastAddressEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRememberLastAddressEnabled

`func (o *PutSettingRequest) SetRememberLastAddressEnabled(v bool)`

SetRememberLastAddressEnabled sets RememberLastAddressEnabled field to given value.


### GetSessionTimeout

`func (o *PutSettingRequest) GetSessionTimeout() string`

GetSessionTimeout returns the SessionTimeout field if non-nil, zero value otherwise.

### GetSessionTimeoutOk

`func (o *PutSettingRequest) GetSessionTimeoutOk() (*string, bool)`

GetSessionTimeoutOk returns a tuple with the SessionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionTimeout

`func (o *PutSettingRequest) SetSessionTimeout(v string)`

SetSessionTimeout sets SessionTimeout field to given value.


### GetDisclaimerEnabled

`func (o *PutSettingRequest) GetDisclaimerEnabled() bool`

GetDisclaimerEnabled returns the DisclaimerEnabled field if non-nil, zero value otherwise.

### GetDisclaimerEnabledOk

`func (o *PutSettingRequest) GetDisclaimerEnabledOk() (*bool, bool)`

GetDisclaimerEnabledOk returns a tuple with the DisclaimerEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisclaimerEnabled

`func (o *PutSettingRequest) SetDisclaimerEnabled(v bool)`

SetDisclaimerEnabled sets DisclaimerEnabled field to given value.


### GetDisclaimerText

`func (o *PutSettingRequest) GetDisclaimerText() string`

GetDisclaimerText returns the DisclaimerText field if non-nil, zero value otherwise.

### GetDisclaimerTextOk

`func (o *PutSettingRequest) GetDisclaimerTextOk() (*string, bool)`

GetDisclaimerTextOk returns a tuple with the DisclaimerText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisclaimerText

`func (o *PutSettingRequest) SetDisclaimerText(v string)`

SetDisclaimerText sets DisclaimerText field to given value.

### HasDisclaimerText

`func (o *PutSettingRequest) HasDisclaimerText() bool`

HasDisclaimerText returns a boolean if a field has been set.

### GetCustomReverseZoneFormatsAllowed

`func (o *PutSettingRequest) GetCustomReverseZoneFormatsAllowed() bool`

GetCustomReverseZoneFormatsAllowed returns the CustomReverseZoneFormatsAllowed field if non-nil, zero value otherwise.

### GetCustomReverseZoneFormatsAllowedOk

`func (o *PutSettingRequest) GetCustomReverseZoneFormatsAllowedOk() (*bool, bool)`

GetCustomReverseZoneFormatsAllowedOk returns a tuple with the CustomReverseZoneFormatsAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomReverseZoneFormatsAllowed

`func (o *PutSettingRequest) SetCustomReverseZoneFormatsAllowed(v bool)`

SetCustomReverseZoneFormatsAllowed sets CustomReverseZoneFormatsAllowed field to given value.


### GetAddressManager

`func (o *PutSettingRequest) GetAddressManager() string`

GetAddressManager returns the AddressManager field if non-nil, zero value otherwise.

### GetAddressManagerOk

`func (o *PutSettingRequest) GetAddressManagerOk() (*string, bool)`

GetAddressManagerOk returns a tuple with the AddressManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressManager

`func (o *PutSettingRequest) SetAddressManager(v string)`

SetAddressManager sets AddressManager field to given value.


### GetRestV2

`func (o *PutSettingRequest) GetRestV2() string`

GetRestV2 returns the RestV2 field if non-nil, zero value otherwise.

### GetRestV2Ok

`func (o *PutSettingRequest) GetRestV2Ok() (*string, bool)`

GetRestV2Ok returns a tuple with the RestV2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestV2

`func (o *PutSettingRequest) SetRestV2(v string)`

SetRestV2 sets RestV2 field to given value.


### GetZoneImport

`func (o *PutSettingRequest) GetZoneImport() string`

GetZoneImport returns the ZoneImport field if non-nil, zero value otherwise.

### GetZoneImportOk

`func (o *PutSettingRequest) GetZoneImportOk() (*string, bool)`

GetZoneImportOk returns a tuple with the ZoneImport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneImport

`func (o *PutSettingRequest) SetZoneImport(v string)`

SetZoneImport sets ZoneImport field to given value.


### GetReconciliationService

`func (o *PutSettingRequest) GetReconciliationService() string`

GetReconciliationService returns the ReconciliationService field if non-nil, zero value otherwise.

### GetReconciliationServiceOk

`func (o *PutSettingRequest) GetReconciliationServiceOk() (*string, bool)`

GetReconciliationServiceOk returns a tuple with the ReconciliationService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReconciliationService

`func (o *PutSettingRequest) SetReconciliationService(v string)`

SetReconciliationService sets ReconciliationService field to given value.


### GetDiscoveryEngine

`func (o *PutSettingRequest) GetDiscoveryEngine() string`

GetDiscoveryEngine returns the DiscoveryEngine field if non-nil, zero value otherwise.

### GetDiscoveryEngineOk

`func (o *PutSettingRequest) GetDiscoveryEngineOk() (*string, bool)`

GetDiscoveryEngineOk returns a tuple with the DiscoveryEngine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryEngine

`func (o *PutSettingRequest) SetDiscoveryEngine(v string)`

SetDiscoveryEngine sets DiscoveryEngine field to given value.


### GetSnmpClient

`func (o *PutSettingRequest) GetSnmpClient() string`

GetSnmpClient returns the SnmpClient field if non-nil, zero value otherwise.

### GetSnmpClientOk

`func (o *PutSettingRequest) GetSnmpClientOk() (*string, bool)`

GetSnmpClientOk returns a tuple with the SnmpClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpClient

`func (o *PutSettingRequest) SetSnmpClient(v string)`

SetSnmpClient sets SnmpClient field to given value.


### GetMonitoringService

`func (o *PutSettingRequest) GetMonitoringService() string`

GetMonitoringService returns the MonitoringService field if non-nil, zero value otherwise.

### GetMonitoringServiceOk

`func (o *PutSettingRequest) GetMonitoringServiceOk() (*string, bool)`

GetMonitoringServiceOk returns a tuple with the MonitoringService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringService

`func (o *PutSettingRequest) SetMonitoringService(v string)`

SetMonitoringService sets MonitoringService field to given value.


### GetRrdTool

`func (o *PutSettingRequest) GetRrdTool() string`

GetRrdTool returns the RrdTool field if non-nil, zero value otherwise.

### GetRrdToolOk

`func (o *PutSettingRequest) GetRrdToolOk() (*string, bool)`

GetRrdToolOk returns a tuple with the RrdTool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrdTool

`func (o *PutSettingRequest) SetRrdTool(v string)`

SetRrdTool sets RrdTool field to given value.


### GetEnabled

`func (o *PutSettingRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PutSettingRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PutSettingRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetFailureLimit

`func (o *PutSettingRequest) GetFailureLimit() int32`

GetFailureLimit returns the FailureLimit field if non-nil, zero value otherwise.

### GetFailureLimitOk

`func (o *PutSettingRequest) GetFailureLimitOk() (*int32, bool)`

GetFailureLimitOk returns a tuple with the FailureLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureLimit

`func (o *PutSettingRequest) SetFailureLimit(v int32)`

SetFailureLimit sets FailureLimit field to given value.


### GetFailureLimitPeriod

`func (o *PutSettingRequest) GetFailureLimitPeriod() string`

GetFailureLimitPeriod returns the FailureLimitPeriod field if non-nil, zero value otherwise.

### GetFailureLimitPeriodOk

`func (o *PutSettingRequest) GetFailureLimitPeriodOk() (*string, bool)`

GetFailureLimitPeriodOk returns a tuple with the FailureLimitPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureLimitPeriod

`func (o *PutSettingRequest) SetFailureLimitPeriod(v string)`

SetFailureLimitPeriod sets FailureLimitPeriod field to given value.


### GetDelayDuration

`func (o *PutSettingRequest) GetDelayDuration() string`

GetDelayDuration returns the DelayDuration field if non-nil, zero value otherwise.

### GetDelayDurationOk

`func (o *PutSettingRequest) GetDelayDurationOk() (*string, bool)`

GetDelayDurationOk returns a tuple with the DelayDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayDuration

`func (o *PutSettingRequest) SetDelayDuration(v string)`

SetDelayDuration sets DelayDuration field to given value.


### GetAddressManagerMonitoring

`func (o *PutSettingRequest) GetAddressManagerMonitoring() AddressManagerMonitoringBean`

GetAddressManagerMonitoring returns the AddressManagerMonitoring field if non-nil, zero value otherwise.

### GetAddressManagerMonitoringOk

`func (o *PutSettingRequest) GetAddressManagerMonitoringOk() (*AddressManagerMonitoringBean, bool)`

GetAddressManagerMonitoringOk returns a tuple with the AddressManagerMonitoring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressManagerMonitoring

`func (o *PutSettingRequest) SetAddressManagerMonitoring(v AddressManagerMonitoringBean)`

SetAddressManagerMonitoring sets AddressManagerMonitoring field to given value.


### GetServerMonitoring

`func (o *PutSettingRequest) GetServerMonitoring() ServerMonitoringBean`

GetServerMonitoring returns the ServerMonitoring field if non-nil, zero value otherwise.

### GetServerMonitoringOk

`func (o *PutSettingRequest) GetServerMonitoringOk() (*ServerMonitoringBean, bool)`

GetServerMonitoringOk returns a tuple with the ServerMonitoring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerMonitoring

`func (o *PutSettingRequest) SetServerMonitoring(v ServerMonitoringBean)`

SetServerMonitoring sets ServerMonitoring field to given value.


### GetF5ServerMonitoring

`func (o *PutSettingRequest) GetF5ServerMonitoring() F5ServerMonitoringBean`

GetF5ServerMonitoring returns the F5ServerMonitoring field if non-nil, zero value otherwise.

### GetF5ServerMonitoringOk

`func (o *PutSettingRequest) GetF5ServerMonitoringOk() (*F5ServerMonitoringBean, bool)`

GetF5ServerMonitoringOk returns a tuple with the F5ServerMonitoring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetF5ServerMonitoring

`func (o *PutSettingRequest) SetF5ServerMonitoring(v F5ServerMonitoringBean)`

SetF5ServerMonitoring sets F5ServerMonitoring field to given value.

### HasF5ServerMonitoring

`func (o *PutSettingRequest) HasF5ServerMonitoring() bool`

HasF5ServerMonitoring returns a boolean if a field has been set.

### GetMinLength

`func (o *PutSettingRequest) GetMinLength() int32`

GetMinLength returns the MinLength field if non-nil, zero value otherwise.

### GetMinLengthOk

`func (o *PutSettingRequest) GetMinLengthOk() (*int32, bool)`

GetMinLengthOk returns a tuple with the MinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLength

`func (o *PutSettingRequest) SetMinLength(v int32)`

SetMinLength sets MinLength field to given value.

### HasMinLength

`func (o *PutSettingRequest) HasMinLength() bool`

HasMinLength returns a boolean if a field has been set.

### GetMaxLength

`func (o *PutSettingRequest) GetMaxLength() int32`

GetMaxLength returns the MaxLength field if non-nil, zero value otherwise.

### GetMaxLengthOk

`func (o *PutSettingRequest) GetMaxLengthOk() (*int32, bool)`

GetMaxLengthOk returns a tuple with the MaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLength

`func (o *PutSettingRequest) SetMaxLength(v int32)`

SetMaxLength sets MaxLength field to given value.

### HasMaxLength

`func (o *PutSettingRequest) HasMaxLength() bool`

HasMaxLength returns a boolean if a field has been set.

### GetDigitRequired

`func (o *PutSettingRequest) GetDigitRequired() bool`

GetDigitRequired returns the DigitRequired field if non-nil, zero value otherwise.

### GetDigitRequiredOk

`func (o *PutSettingRequest) GetDigitRequiredOk() (*bool, bool)`

GetDigitRequiredOk returns a tuple with the DigitRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitRequired

`func (o *PutSettingRequest) SetDigitRequired(v bool)`

SetDigitRequired sets DigitRequired field to given value.


### GetMixedCaseRequired

`func (o *PutSettingRequest) GetMixedCaseRequired() bool`

GetMixedCaseRequired returns the MixedCaseRequired field if non-nil, zero value otherwise.

### GetMixedCaseRequiredOk

`func (o *PutSettingRequest) GetMixedCaseRequiredOk() (*bool, bool)`

GetMixedCaseRequiredOk returns a tuple with the MixedCaseRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMixedCaseRequired

`func (o *PutSettingRequest) SetMixedCaseRequired(v bool)`

SetMixedCaseRequired sets MixedCaseRequired field to given value.


### GetSpecialCharacterRequired

`func (o *PutSettingRequest) GetSpecialCharacterRequired() bool`

GetSpecialCharacterRequired returns the SpecialCharacterRequired field if non-nil, zero value otherwise.

### GetSpecialCharacterRequiredOk

`func (o *PutSettingRequest) GetSpecialCharacterRequiredOk() (*bool, bool)`

GetSpecialCharacterRequiredOk returns a tuple with the SpecialCharacterRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialCharacterRequired

`func (o *PutSettingRequest) SetSpecialCharacterRequired(v bool)`

SetSpecialCharacterRequired sets SpecialCharacterRequired field to given value.


### GetAddressManagerFqdn

`func (o *PutSettingRequest) GetAddressManagerFqdn() string`

GetAddressManagerFqdn returns the AddressManagerFqdn field if non-nil, zero value otherwise.

### GetAddressManagerFqdnOk

`func (o *PutSettingRequest) GetAddressManagerFqdnOk() (*string, bool)`

GetAddressManagerFqdnOk returns a tuple with the AddressManagerFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressManagerFqdn

`func (o *PutSettingRequest) SetAddressManagerFqdn(v string)`

SetAddressManagerFqdn sets AddressManagerFqdn field to given value.


### GetMetadataUrl

`func (o *PutSettingRequest) GetMetadataUrl() string`

GetMetadataUrl returns the MetadataUrl field if non-nil, zero value otherwise.

### GetMetadataUrlOk

`func (o *PutSettingRequest) GetMetadataUrlOk() (*string, bool)`

GetMetadataUrlOk returns a tuple with the MetadataUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataUrl

`func (o *PutSettingRequest) SetMetadataUrl(v string)`

SetMetadataUrl sets MetadataUrl field to given value.

### HasMetadataUrl

`func (o *PutSettingRequest) HasMetadataUrl() bool`

HasMetadataUrl returns a boolean if a field has been set.

### GetConsumeUrl

`func (o *PutSettingRequest) GetConsumeUrl() string`

GetConsumeUrl returns the ConsumeUrl field if non-nil, zero value otherwise.

### GetConsumeUrlOk

`func (o *PutSettingRequest) GetConsumeUrlOk() (*string, bool)`

GetConsumeUrlOk returns a tuple with the ConsumeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumeUrl

`func (o *PutSettingRequest) SetConsumeUrl(v string)`

SetConsumeUrl sets ConsumeUrl field to given value.

### HasConsumeUrl

`func (o *PutSettingRequest) HasConsumeUrl() bool`

HasConsumeUrl returns a boolean if a field has been set.

### GetLogoutUrl

`func (o *PutSettingRequest) GetLogoutUrl() string`

GetLogoutUrl returns the LogoutUrl field if non-nil, zero value otherwise.

### GetLogoutUrlOk

`func (o *PutSettingRequest) GetLogoutUrlOk() (*string, bool)`

GetLogoutUrlOk returns a tuple with the LogoutUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoutUrl

`func (o *PutSettingRequest) SetLogoutUrl(v string)`

SetLogoutUrl sets LogoutUrl field to given value.

### HasLogoutUrl

`func (o *PutSettingRequest) HasLogoutUrl() bool`

HasLogoutUrl returns a boolean if a field has been set.

### GetNameIdFormat

`func (o *PutSettingRequest) GetNameIdFormat() string`

GetNameIdFormat returns the NameIdFormat field if non-nil, zero value otherwise.

### GetNameIdFormatOk

`func (o *PutSettingRequest) GetNameIdFormatOk() (*string, bool)`

GetNameIdFormatOk returns a tuple with the NameIdFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameIdFormat

`func (o *PutSettingRequest) SetNameIdFormat(v string)`

SetNameIdFormat sets NameIdFormat field to given value.


### GetSigningEnabled

`func (o *PutSettingRequest) GetSigningEnabled() bool`

GetSigningEnabled returns the SigningEnabled field if non-nil, zero value otherwise.

### GetSigningEnabledOk

`func (o *PutSettingRequest) GetSigningEnabledOk() (*bool, bool)`

GetSigningEnabledOk returns a tuple with the SigningEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningEnabled

`func (o *PutSettingRequest) SetSigningEnabled(v bool)`

SetSigningEnabled sets SigningEnabled field to given value.


### GetEncryptionEnabled

`func (o *PutSettingRequest) GetEncryptionEnabled() bool`

GetEncryptionEnabled returns the EncryptionEnabled field if non-nil, zero value otherwise.

### GetEncryptionEnabledOk

`func (o *PutSettingRequest) GetEncryptionEnabledOk() (*bool, bool)`

GetEncryptionEnabledOk returns a tuple with the EncryptionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionEnabled

`func (o *PutSettingRequest) SetEncryptionEnabled(v bool)`

SetEncryptionEnabled sets EncryptionEnabled field to given value.


### GetPkcs12

`func (o *PutSettingRequest) GetPkcs12() string`

GetPkcs12 returns the Pkcs12 field if non-nil, zero value otherwise.

### GetPkcs12Ok

`func (o *PutSettingRequest) GetPkcs12Ok() (*string, bool)`

GetPkcs12Ok returns a tuple with the Pkcs12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkcs12

`func (o *PutSettingRequest) SetPkcs12(v string)`

SetPkcs12 sets Pkcs12 field to given value.

### HasPkcs12

`func (o *PutSettingRequest) HasPkcs12() bool`

HasPkcs12 returns a boolean if a field has been set.

### GetPassword

`func (o *PutSettingRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PutSettingRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PutSettingRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *PutSettingRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetOrganizationName

`func (o *PutSettingRequest) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *PutSettingRequest) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *PutSettingRequest) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.


### GetOrganizationUrl

`func (o *PutSettingRequest) GetOrganizationUrl() string`

GetOrganizationUrl returns the OrganizationUrl field if non-nil, zero value otherwise.

### GetOrganizationUrlOk

`func (o *PutSettingRequest) GetOrganizationUrlOk() (*string, bool)`

GetOrganizationUrlOk returns a tuple with the OrganizationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationUrl

`func (o *PutSettingRequest) SetOrganizationUrl(v string)`

SetOrganizationUrl sets OrganizationUrl field to given value.


### GetContactName

`func (o *PutSettingRequest) GetContactName() string`

GetContactName returns the ContactName field if non-nil, zero value otherwise.

### GetContactNameOk

`func (o *PutSettingRequest) GetContactNameOk() (*string, bool)`

GetContactNameOk returns a tuple with the ContactName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactName

`func (o *PutSettingRequest) SetContactName(v string)`

SetContactName sets ContactName field to given value.

### HasContactName

`func (o *PutSettingRequest) HasContactName() bool`

HasContactName returns a boolean if a field has been set.

### GetContactEmail

`func (o *PutSettingRequest) GetContactEmail() string`

GetContactEmail returns the ContactEmail field if non-nil, zero value otherwise.

### GetContactEmailOk

`func (o *PutSettingRequest) GetContactEmailOk() (*string, bool)`

GetContactEmailOk returns a tuple with the ContactEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactEmail

`func (o *PutSettingRequest) SetContactEmail(v string)`

SetContactEmail sets ContactEmail field to given value.

### HasContactEmail

`func (o *PutSettingRequest) HasContactEmail() bool`

HasContactEmail returns a boolean if a field has been set.

### GetSamlIdentityProviderEnabled

`func (o *PutSettingRequest) GetSamlIdentityProviderEnabled() bool`

GetSamlIdentityProviderEnabled returns the SamlIdentityProviderEnabled field if non-nil, zero value otherwise.

### GetSamlIdentityProviderEnabledOk

`func (o *PutSettingRequest) GetSamlIdentityProviderEnabledOk() (*bool, bool)`

GetSamlIdentityProviderEnabledOk returns a tuple with the SamlIdentityProviderEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlIdentityProviderEnabled

`func (o *PutSettingRequest) SetSamlIdentityProviderEnabled(v bool)`

SetSamlIdentityProviderEnabled sets SamlIdentityProviderEnabled field to given value.

### HasSamlIdentityProviderEnabled

`func (o *PutSettingRequest) HasSamlIdentityProviderEnabled() bool`

HasSamlIdentityProviderEnabled returns a boolean if a field has been set.

### GetNonSsoAuthenticatorCount

`func (o *PutSettingRequest) GetNonSsoAuthenticatorCount() int32`

GetNonSsoAuthenticatorCount returns the NonSsoAuthenticatorCount field if non-nil, zero value otherwise.

### GetNonSsoAuthenticatorCountOk

`func (o *PutSettingRequest) GetNonSsoAuthenticatorCountOk() (*int32, bool)`

GetNonSsoAuthenticatorCountOk returns a tuple with the NonSsoAuthenticatorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonSsoAuthenticatorCount

`func (o *PutSettingRequest) SetNonSsoAuthenticatorCount(v int32)`

SetNonSsoAuthenticatorCount sets NonSsoAuthenticatorCount field to given value.

### HasNonSsoAuthenticatorCount

`func (o *PutSettingRequest) HasNonSsoAuthenticatorCount() bool`

HasNonSsoAuthenticatorCount returns a boolean if a field has been set.

### GetNonSsoGroupCount

`func (o *PutSettingRequest) GetNonSsoGroupCount() int32`

GetNonSsoGroupCount returns the NonSsoGroupCount field if non-nil, zero value otherwise.

### GetNonSsoGroupCountOk

`func (o *PutSettingRequest) GetNonSsoGroupCountOk() (*int32, bool)`

GetNonSsoGroupCountOk returns a tuple with the NonSsoGroupCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonSsoGroupCount

`func (o *PutSettingRequest) SetNonSsoGroupCount(v int32)`

SetNonSsoGroupCount sets NonSsoGroupCount field to given value.

### HasNonSsoGroupCount

`func (o *PutSettingRequest) HasNonSsoGroupCount() bool`

HasNonSsoGroupCount returns a boolean if a field has been set.

### GetLocalAdminUserCount

`func (o *PutSettingRequest) GetLocalAdminUserCount() int32`

GetLocalAdminUserCount returns the LocalAdminUserCount field if non-nil, zero value otherwise.

### GetLocalAdminUserCountOk

`func (o *PutSettingRequest) GetLocalAdminUserCountOk() (*int32, bool)`

GetLocalAdminUserCountOk returns a tuple with the LocalAdminUserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAdminUserCount

`func (o *PutSettingRequest) SetLocalAdminUserCount(v int32)`

SetLocalAdminUserCount sets LocalAdminUserCount field to given value.

### HasLocalAdminUserCount

`func (o *PutSettingRequest) HasLocalAdminUserCount() bool`

HasLocalAdminUserCount returns a boolean if a field has been set.

### GetLocalNonAdminUserCount

`func (o *PutSettingRequest) GetLocalNonAdminUserCount() int32`

GetLocalNonAdminUserCount returns the LocalNonAdminUserCount field if non-nil, zero value otherwise.

### GetLocalNonAdminUserCountOk

`func (o *PutSettingRequest) GetLocalNonAdminUserCountOk() (*int32, bool)`

GetLocalNonAdminUserCountOk returns a tuple with the LocalNonAdminUserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalNonAdminUserCount

`func (o *PutSettingRequest) SetLocalNonAdminUserCount(v int32)`

SetLocalNonAdminUserCount sets LocalNonAdminUserCount field to given value.

### HasLocalNonAdminUserCount

`func (o *PutSettingRequest) HasLocalNonAdminUserCount() bool`

HasLocalNonAdminUserCount returns a boolean if a field has been set.

### GetHostname

`func (o *PutSettingRequest) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *PutSettingRequest) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *PutSettingRequest) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *PutSettingRequest) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetVersion

`func (o *PutSettingRequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PutSettingRequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PutSettingRequest) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PutSettingRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetAddress

`func (o *PutSettingRequest) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *PutSettingRequest) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *PutSettingRequest) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *PutSettingRequest) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetInterfaceRedundancyEnabled

`func (o *PutSettingRequest) GetInterfaceRedundancyEnabled() bool`

GetInterfaceRedundancyEnabled returns the InterfaceRedundancyEnabled field if non-nil, zero value otherwise.

### GetInterfaceRedundancyEnabledOk

`func (o *PutSettingRequest) GetInterfaceRedundancyEnabledOk() (*bool, bool)`

GetInterfaceRedundancyEnabledOk returns a tuple with the InterfaceRedundancyEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceRedundancyEnabled

`func (o *PutSettingRequest) SetInterfaceRedundancyEnabled(v bool)`

SetInterfaceRedundancyEnabled sets InterfaceRedundancyEnabled field to given value.


### GetActiveSessions

`func (o *PutSettingRequest) GetActiveSessions() int32`

GetActiveSessions returns the ActiveSessions field if non-nil, zero value otherwise.

### GetActiveSessionsOk

`func (o *PutSettingRequest) GetActiveSessionsOk() (*int32, bool)`

GetActiveSessionsOk returns a tuple with the ActiveSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveSessions

`func (o *PutSettingRequest) SetActiveSessions(v int32)`

SetActiveSessions sets ActiveSessions field to given value.

### HasActiveSessions

`func (o *PutSettingRequest) HasActiveSessions() bool`

HasActiveSessions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


