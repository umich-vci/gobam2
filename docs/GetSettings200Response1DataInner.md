# GetSettings200Response1DataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**MandatoryCommentsEnabled** | Pointer to **bool** | Indicates whether users must provide a change control comment when creating, editing, or deleting resources. | [optional] 
**RememberLastAddressEnabled** | Pointer to **bool** | Indicates whether the Address Manager UI Quick Actions widget will automatically provide the subsequent IPv4 address when adding host records. | [optional] 
**SessionTimeout** | Pointer to **string** | The maximum time period of user inactivity before a session is terminated. | [optional] 
**DisclaimerEnabled** | Pointer to **bool** | Indicates whether a disclaimer will display on the Address Manager UI login page with the contents of disclaimerText. | [optional] 
**DisclaimerText** | Pointer to **string** | The disclaimer text or HTML for display on the Address Manager UI login page. If adding or editing disclaimer HTML through JSON, ensure that reserved characters are escaped. | [optional] 
**CustomReverseZoneFormatsAllowed** | Pointer to **bool** | Indicates whether users can set a custom reverse zone name format in DNS deployment options. | [optional] 
**AddressManager** | Pointer to **string** |  | [optional] 
**RestV2** | Pointer to **string** |  | [optional] 
**ZoneImport** | Pointer to **string** |  | [optional] 
**ReconciliationService** | Pointer to **string** |  | [optional] 
**DiscoveryEngine** | Pointer to **string** |  | [optional] 
**SnmpClient** | Pointer to **string** |  | [optional] 
**MonitoringService** | Pointer to **string** |  | [optional] 
**RrdTool** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**FailureLimit** | Pointer to **int32** | The limit for consecutive failed login attempts that results in a login delay. | [optional] 
**FailureLimitPeriod** | Pointer to **string** | The timespan in which consecutive failed login attempts count towards the failure limit. | [optional] 
**DelayDuration** | Pointer to **string** | The amount of time that a user must wait before they can attempt to log in after exceeding the login policy failure conditions. | [optional] 
**AddressManagerMonitoring** | Pointer to [**AddressManagerMonitoringBean**](AddressManagerMonitoringBean.md) |  | [optional] 
**ServerMonitoring** | Pointer to [**ServerMonitoringBean**](ServerMonitoringBean.md) |  | [optional] 
**F5ServerMonitoring** | Pointer to [**F5ServerMonitoringBean**](F5ServerMonitoringBean.md) |  | [optional] 
**MinLength** | Pointer to **int32** | The minimum length that the password must be. | [optional] 
**MaxLength** | Pointer to **int32** | The maximum length that the password can be. | [optional] 
**DigitRequired** | Pointer to **bool** | Indicates whether the password must contain at least one digit. | [optional] 
**MixedCaseRequired** | Pointer to **bool** | Indicates whether the password must contain mixed case letters. | [optional] 
**SpecialCharacterRequired** | Pointer to **bool** | Indicates whether the password must contain at least one special character. | [optional] 
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
**SamlIdentityProviderEnabled** | Pointer to **bool** |  | [optional] [readonly] 
**NonSsoAuthenticatorCount** | Pointer to **int32** |  | [optional] [readonly] 
**NonSsoGroupCount** | Pointer to **int32** |  | [optional] [readonly] 
**LocalAdminUserCount** | Pointer to **int32** |  | [optional] [readonly] 
**LocalNonAdminUserCount** | Pointer to **int32** |  | [optional] [readonly] 
**Hostname** | Pointer to **string** | The hostname of the Address Manager server. | [optional] [readonly] 
**Version** | Pointer to **string** | The Address Manager server version. | [optional] [readonly] 
**Address** | Pointer to **string** | The management IP address of the Address Manager server. | [optional] [readonly] 
**InterfaceRedundancyEnabled** | Pointer to **bool** | Indicates whether Address Manager Active/Backup (eth0/eth1) bonding mode is active. | [optional] 
**ActiveSessions** | Pointer to **int32** | The current number of active Address Manager sessions. | [optional] [readonly] 

## Methods

### NewGetSettings200Response1DataInner

`func NewGetSettings200Response1DataInner() *GetSettings200Response1DataInner`

NewGetSettings200Response1DataInner instantiates a new GetSettings200Response1DataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSettings200Response1DataInnerWithDefaults

`func NewGetSettings200Response1DataInnerWithDefaults() *GetSettings200Response1DataInner`

NewGetSettings200Response1DataInnerWithDefaults instantiates a new GetSettings200Response1DataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetSettings200Response1DataInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetSettings200Response1DataInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetSettings200Response1DataInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetSettings200Response1DataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GetSettings200Response1DataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetSettings200Response1DataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetSettings200Response1DataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetSettings200Response1DataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMandatoryCommentsEnabled

`func (o *GetSettings200Response1DataInner) GetMandatoryCommentsEnabled() bool`

GetMandatoryCommentsEnabled returns the MandatoryCommentsEnabled field if non-nil, zero value otherwise.

### GetMandatoryCommentsEnabledOk

`func (o *GetSettings200Response1DataInner) GetMandatoryCommentsEnabledOk() (*bool, bool)`

GetMandatoryCommentsEnabledOk returns a tuple with the MandatoryCommentsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandatoryCommentsEnabled

`func (o *GetSettings200Response1DataInner) SetMandatoryCommentsEnabled(v bool)`

SetMandatoryCommentsEnabled sets MandatoryCommentsEnabled field to given value.

### HasMandatoryCommentsEnabled

`func (o *GetSettings200Response1DataInner) HasMandatoryCommentsEnabled() bool`

HasMandatoryCommentsEnabled returns a boolean if a field has been set.

### GetRememberLastAddressEnabled

`func (o *GetSettings200Response1DataInner) GetRememberLastAddressEnabled() bool`

GetRememberLastAddressEnabled returns the RememberLastAddressEnabled field if non-nil, zero value otherwise.

### GetRememberLastAddressEnabledOk

`func (o *GetSettings200Response1DataInner) GetRememberLastAddressEnabledOk() (*bool, bool)`

GetRememberLastAddressEnabledOk returns a tuple with the RememberLastAddressEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRememberLastAddressEnabled

`func (o *GetSettings200Response1DataInner) SetRememberLastAddressEnabled(v bool)`

SetRememberLastAddressEnabled sets RememberLastAddressEnabled field to given value.

### HasRememberLastAddressEnabled

`func (o *GetSettings200Response1DataInner) HasRememberLastAddressEnabled() bool`

HasRememberLastAddressEnabled returns a boolean if a field has been set.

### GetSessionTimeout

`func (o *GetSettings200Response1DataInner) GetSessionTimeout() string`

GetSessionTimeout returns the SessionTimeout field if non-nil, zero value otherwise.

### GetSessionTimeoutOk

`func (o *GetSettings200Response1DataInner) GetSessionTimeoutOk() (*string, bool)`

GetSessionTimeoutOk returns a tuple with the SessionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionTimeout

`func (o *GetSettings200Response1DataInner) SetSessionTimeout(v string)`

SetSessionTimeout sets SessionTimeout field to given value.

### HasSessionTimeout

`func (o *GetSettings200Response1DataInner) HasSessionTimeout() bool`

HasSessionTimeout returns a boolean if a field has been set.

### GetDisclaimerEnabled

`func (o *GetSettings200Response1DataInner) GetDisclaimerEnabled() bool`

GetDisclaimerEnabled returns the DisclaimerEnabled field if non-nil, zero value otherwise.

### GetDisclaimerEnabledOk

`func (o *GetSettings200Response1DataInner) GetDisclaimerEnabledOk() (*bool, bool)`

GetDisclaimerEnabledOk returns a tuple with the DisclaimerEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisclaimerEnabled

`func (o *GetSettings200Response1DataInner) SetDisclaimerEnabled(v bool)`

SetDisclaimerEnabled sets DisclaimerEnabled field to given value.

### HasDisclaimerEnabled

`func (o *GetSettings200Response1DataInner) HasDisclaimerEnabled() bool`

HasDisclaimerEnabled returns a boolean if a field has been set.

### GetDisclaimerText

`func (o *GetSettings200Response1DataInner) GetDisclaimerText() string`

GetDisclaimerText returns the DisclaimerText field if non-nil, zero value otherwise.

### GetDisclaimerTextOk

`func (o *GetSettings200Response1DataInner) GetDisclaimerTextOk() (*string, bool)`

GetDisclaimerTextOk returns a tuple with the DisclaimerText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisclaimerText

`func (o *GetSettings200Response1DataInner) SetDisclaimerText(v string)`

SetDisclaimerText sets DisclaimerText field to given value.

### HasDisclaimerText

`func (o *GetSettings200Response1DataInner) HasDisclaimerText() bool`

HasDisclaimerText returns a boolean if a field has been set.

### GetCustomReverseZoneFormatsAllowed

`func (o *GetSettings200Response1DataInner) GetCustomReverseZoneFormatsAllowed() bool`

GetCustomReverseZoneFormatsAllowed returns the CustomReverseZoneFormatsAllowed field if non-nil, zero value otherwise.

### GetCustomReverseZoneFormatsAllowedOk

`func (o *GetSettings200Response1DataInner) GetCustomReverseZoneFormatsAllowedOk() (*bool, bool)`

GetCustomReverseZoneFormatsAllowedOk returns a tuple with the CustomReverseZoneFormatsAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomReverseZoneFormatsAllowed

`func (o *GetSettings200Response1DataInner) SetCustomReverseZoneFormatsAllowed(v bool)`

SetCustomReverseZoneFormatsAllowed sets CustomReverseZoneFormatsAllowed field to given value.

### HasCustomReverseZoneFormatsAllowed

`func (o *GetSettings200Response1DataInner) HasCustomReverseZoneFormatsAllowed() bool`

HasCustomReverseZoneFormatsAllowed returns a boolean if a field has been set.

### GetAddressManager

`func (o *GetSettings200Response1DataInner) GetAddressManager() string`

GetAddressManager returns the AddressManager field if non-nil, zero value otherwise.

### GetAddressManagerOk

`func (o *GetSettings200Response1DataInner) GetAddressManagerOk() (*string, bool)`

GetAddressManagerOk returns a tuple with the AddressManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressManager

`func (o *GetSettings200Response1DataInner) SetAddressManager(v string)`

SetAddressManager sets AddressManager field to given value.

### HasAddressManager

`func (o *GetSettings200Response1DataInner) HasAddressManager() bool`

HasAddressManager returns a boolean if a field has been set.

### GetRestV2

`func (o *GetSettings200Response1DataInner) GetRestV2() string`

GetRestV2 returns the RestV2 field if non-nil, zero value otherwise.

### GetRestV2Ok

`func (o *GetSettings200Response1DataInner) GetRestV2Ok() (*string, bool)`

GetRestV2Ok returns a tuple with the RestV2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestV2

`func (o *GetSettings200Response1DataInner) SetRestV2(v string)`

SetRestV2 sets RestV2 field to given value.

### HasRestV2

`func (o *GetSettings200Response1DataInner) HasRestV2() bool`

HasRestV2 returns a boolean if a field has been set.

### GetZoneImport

`func (o *GetSettings200Response1DataInner) GetZoneImport() string`

GetZoneImport returns the ZoneImport field if non-nil, zero value otherwise.

### GetZoneImportOk

`func (o *GetSettings200Response1DataInner) GetZoneImportOk() (*string, bool)`

GetZoneImportOk returns a tuple with the ZoneImport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneImport

`func (o *GetSettings200Response1DataInner) SetZoneImport(v string)`

SetZoneImport sets ZoneImport field to given value.

### HasZoneImport

`func (o *GetSettings200Response1DataInner) HasZoneImport() bool`

HasZoneImport returns a boolean if a field has been set.

### GetReconciliationService

`func (o *GetSettings200Response1DataInner) GetReconciliationService() string`

GetReconciliationService returns the ReconciliationService field if non-nil, zero value otherwise.

### GetReconciliationServiceOk

`func (o *GetSettings200Response1DataInner) GetReconciliationServiceOk() (*string, bool)`

GetReconciliationServiceOk returns a tuple with the ReconciliationService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReconciliationService

`func (o *GetSettings200Response1DataInner) SetReconciliationService(v string)`

SetReconciliationService sets ReconciliationService field to given value.

### HasReconciliationService

`func (o *GetSettings200Response1DataInner) HasReconciliationService() bool`

HasReconciliationService returns a boolean if a field has been set.

### GetDiscoveryEngine

`func (o *GetSettings200Response1DataInner) GetDiscoveryEngine() string`

GetDiscoveryEngine returns the DiscoveryEngine field if non-nil, zero value otherwise.

### GetDiscoveryEngineOk

`func (o *GetSettings200Response1DataInner) GetDiscoveryEngineOk() (*string, bool)`

GetDiscoveryEngineOk returns a tuple with the DiscoveryEngine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryEngine

`func (o *GetSettings200Response1DataInner) SetDiscoveryEngine(v string)`

SetDiscoveryEngine sets DiscoveryEngine field to given value.

### HasDiscoveryEngine

`func (o *GetSettings200Response1DataInner) HasDiscoveryEngine() bool`

HasDiscoveryEngine returns a boolean if a field has been set.

### GetSnmpClient

`func (o *GetSettings200Response1DataInner) GetSnmpClient() string`

GetSnmpClient returns the SnmpClient field if non-nil, zero value otherwise.

### GetSnmpClientOk

`func (o *GetSettings200Response1DataInner) GetSnmpClientOk() (*string, bool)`

GetSnmpClientOk returns a tuple with the SnmpClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpClient

`func (o *GetSettings200Response1DataInner) SetSnmpClient(v string)`

SetSnmpClient sets SnmpClient field to given value.

### HasSnmpClient

`func (o *GetSettings200Response1DataInner) HasSnmpClient() bool`

HasSnmpClient returns a boolean if a field has been set.

### GetMonitoringService

`func (o *GetSettings200Response1DataInner) GetMonitoringService() string`

GetMonitoringService returns the MonitoringService field if non-nil, zero value otherwise.

### GetMonitoringServiceOk

`func (o *GetSettings200Response1DataInner) GetMonitoringServiceOk() (*string, bool)`

GetMonitoringServiceOk returns a tuple with the MonitoringService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringService

`func (o *GetSettings200Response1DataInner) SetMonitoringService(v string)`

SetMonitoringService sets MonitoringService field to given value.

### HasMonitoringService

`func (o *GetSettings200Response1DataInner) HasMonitoringService() bool`

HasMonitoringService returns a boolean if a field has been set.

### GetRrdTool

`func (o *GetSettings200Response1DataInner) GetRrdTool() string`

GetRrdTool returns the RrdTool field if non-nil, zero value otherwise.

### GetRrdToolOk

`func (o *GetSettings200Response1DataInner) GetRrdToolOk() (*string, bool)`

GetRrdToolOk returns a tuple with the RrdTool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrdTool

`func (o *GetSettings200Response1DataInner) SetRrdTool(v string)`

SetRrdTool sets RrdTool field to given value.

### HasRrdTool

`func (o *GetSettings200Response1DataInner) HasRrdTool() bool`

HasRrdTool returns a boolean if a field has been set.

### GetEnabled

`func (o *GetSettings200Response1DataInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetSettings200Response1DataInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetSettings200Response1DataInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetSettings200Response1DataInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFailureLimit

`func (o *GetSettings200Response1DataInner) GetFailureLimit() int32`

GetFailureLimit returns the FailureLimit field if non-nil, zero value otherwise.

### GetFailureLimitOk

`func (o *GetSettings200Response1DataInner) GetFailureLimitOk() (*int32, bool)`

GetFailureLimitOk returns a tuple with the FailureLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureLimit

`func (o *GetSettings200Response1DataInner) SetFailureLimit(v int32)`

SetFailureLimit sets FailureLimit field to given value.

### HasFailureLimit

`func (o *GetSettings200Response1DataInner) HasFailureLimit() bool`

HasFailureLimit returns a boolean if a field has been set.

### GetFailureLimitPeriod

`func (o *GetSettings200Response1DataInner) GetFailureLimitPeriod() string`

GetFailureLimitPeriod returns the FailureLimitPeriod field if non-nil, zero value otherwise.

### GetFailureLimitPeriodOk

`func (o *GetSettings200Response1DataInner) GetFailureLimitPeriodOk() (*string, bool)`

GetFailureLimitPeriodOk returns a tuple with the FailureLimitPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureLimitPeriod

`func (o *GetSettings200Response1DataInner) SetFailureLimitPeriod(v string)`

SetFailureLimitPeriod sets FailureLimitPeriod field to given value.

### HasFailureLimitPeriod

`func (o *GetSettings200Response1DataInner) HasFailureLimitPeriod() bool`

HasFailureLimitPeriod returns a boolean if a field has been set.

### GetDelayDuration

`func (o *GetSettings200Response1DataInner) GetDelayDuration() string`

GetDelayDuration returns the DelayDuration field if non-nil, zero value otherwise.

### GetDelayDurationOk

`func (o *GetSettings200Response1DataInner) GetDelayDurationOk() (*string, bool)`

GetDelayDurationOk returns a tuple with the DelayDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayDuration

`func (o *GetSettings200Response1DataInner) SetDelayDuration(v string)`

SetDelayDuration sets DelayDuration field to given value.

### HasDelayDuration

`func (o *GetSettings200Response1DataInner) HasDelayDuration() bool`

HasDelayDuration returns a boolean if a field has been set.

### GetAddressManagerMonitoring

`func (o *GetSettings200Response1DataInner) GetAddressManagerMonitoring() AddressManagerMonitoringBean`

GetAddressManagerMonitoring returns the AddressManagerMonitoring field if non-nil, zero value otherwise.

### GetAddressManagerMonitoringOk

`func (o *GetSettings200Response1DataInner) GetAddressManagerMonitoringOk() (*AddressManagerMonitoringBean, bool)`

GetAddressManagerMonitoringOk returns a tuple with the AddressManagerMonitoring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressManagerMonitoring

`func (o *GetSettings200Response1DataInner) SetAddressManagerMonitoring(v AddressManagerMonitoringBean)`

SetAddressManagerMonitoring sets AddressManagerMonitoring field to given value.

### HasAddressManagerMonitoring

`func (o *GetSettings200Response1DataInner) HasAddressManagerMonitoring() bool`

HasAddressManagerMonitoring returns a boolean if a field has been set.

### GetServerMonitoring

`func (o *GetSettings200Response1DataInner) GetServerMonitoring() ServerMonitoringBean`

GetServerMonitoring returns the ServerMonitoring field if non-nil, zero value otherwise.

### GetServerMonitoringOk

`func (o *GetSettings200Response1DataInner) GetServerMonitoringOk() (*ServerMonitoringBean, bool)`

GetServerMonitoringOk returns a tuple with the ServerMonitoring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerMonitoring

`func (o *GetSettings200Response1DataInner) SetServerMonitoring(v ServerMonitoringBean)`

SetServerMonitoring sets ServerMonitoring field to given value.

### HasServerMonitoring

`func (o *GetSettings200Response1DataInner) HasServerMonitoring() bool`

HasServerMonitoring returns a boolean if a field has been set.

### GetF5ServerMonitoring

`func (o *GetSettings200Response1DataInner) GetF5ServerMonitoring() F5ServerMonitoringBean`

GetF5ServerMonitoring returns the F5ServerMonitoring field if non-nil, zero value otherwise.

### GetF5ServerMonitoringOk

`func (o *GetSettings200Response1DataInner) GetF5ServerMonitoringOk() (*F5ServerMonitoringBean, bool)`

GetF5ServerMonitoringOk returns a tuple with the F5ServerMonitoring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetF5ServerMonitoring

`func (o *GetSettings200Response1DataInner) SetF5ServerMonitoring(v F5ServerMonitoringBean)`

SetF5ServerMonitoring sets F5ServerMonitoring field to given value.

### HasF5ServerMonitoring

`func (o *GetSettings200Response1DataInner) HasF5ServerMonitoring() bool`

HasF5ServerMonitoring returns a boolean if a field has been set.

### GetMinLength

`func (o *GetSettings200Response1DataInner) GetMinLength() int32`

GetMinLength returns the MinLength field if non-nil, zero value otherwise.

### GetMinLengthOk

`func (o *GetSettings200Response1DataInner) GetMinLengthOk() (*int32, bool)`

GetMinLengthOk returns a tuple with the MinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLength

`func (o *GetSettings200Response1DataInner) SetMinLength(v int32)`

SetMinLength sets MinLength field to given value.

### HasMinLength

`func (o *GetSettings200Response1DataInner) HasMinLength() bool`

HasMinLength returns a boolean if a field has been set.

### GetMaxLength

`func (o *GetSettings200Response1DataInner) GetMaxLength() int32`

GetMaxLength returns the MaxLength field if non-nil, zero value otherwise.

### GetMaxLengthOk

`func (o *GetSettings200Response1DataInner) GetMaxLengthOk() (*int32, bool)`

GetMaxLengthOk returns a tuple with the MaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLength

`func (o *GetSettings200Response1DataInner) SetMaxLength(v int32)`

SetMaxLength sets MaxLength field to given value.

### HasMaxLength

`func (o *GetSettings200Response1DataInner) HasMaxLength() bool`

HasMaxLength returns a boolean if a field has been set.

### GetDigitRequired

`func (o *GetSettings200Response1DataInner) GetDigitRequired() bool`

GetDigitRequired returns the DigitRequired field if non-nil, zero value otherwise.

### GetDigitRequiredOk

`func (o *GetSettings200Response1DataInner) GetDigitRequiredOk() (*bool, bool)`

GetDigitRequiredOk returns a tuple with the DigitRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitRequired

`func (o *GetSettings200Response1DataInner) SetDigitRequired(v bool)`

SetDigitRequired sets DigitRequired field to given value.

### HasDigitRequired

`func (o *GetSettings200Response1DataInner) HasDigitRequired() bool`

HasDigitRequired returns a boolean if a field has been set.

### GetMixedCaseRequired

`func (o *GetSettings200Response1DataInner) GetMixedCaseRequired() bool`

GetMixedCaseRequired returns the MixedCaseRequired field if non-nil, zero value otherwise.

### GetMixedCaseRequiredOk

`func (o *GetSettings200Response1DataInner) GetMixedCaseRequiredOk() (*bool, bool)`

GetMixedCaseRequiredOk returns a tuple with the MixedCaseRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMixedCaseRequired

`func (o *GetSettings200Response1DataInner) SetMixedCaseRequired(v bool)`

SetMixedCaseRequired sets MixedCaseRequired field to given value.

### HasMixedCaseRequired

`func (o *GetSettings200Response1DataInner) HasMixedCaseRequired() bool`

HasMixedCaseRequired returns a boolean if a field has been set.

### GetSpecialCharacterRequired

`func (o *GetSettings200Response1DataInner) GetSpecialCharacterRequired() bool`

GetSpecialCharacterRequired returns the SpecialCharacterRequired field if non-nil, zero value otherwise.

### GetSpecialCharacterRequiredOk

`func (o *GetSettings200Response1DataInner) GetSpecialCharacterRequiredOk() (*bool, bool)`

GetSpecialCharacterRequiredOk returns a tuple with the SpecialCharacterRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialCharacterRequired

`func (o *GetSettings200Response1DataInner) SetSpecialCharacterRequired(v bool)`

SetSpecialCharacterRequired sets SpecialCharacterRequired field to given value.

### HasSpecialCharacterRequired

`func (o *GetSettings200Response1DataInner) HasSpecialCharacterRequired() bool`

HasSpecialCharacterRequired returns a boolean if a field has been set.

### GetAddressManagerFqdn

`func (o *GetSettings200Response1DataInner) GetAddressManagerFqdn() string`

GetAddressManagerFqdn returns the AddressManagerFqdn field if non-nil, zero value otherwise.

### GetAddressManagerFqdnOk

`func (o *GetSettings200Response1DataInner) GetAddressManagerFqdnOk() (*string, bool)`

GetAddressManagerFqdnOk returns a tuple with the AddressManagerFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressManagerFqdn

`func (o *GetSettings200Response1DataInner) SetAddressManagerFqdn(v string)`

SetAddressManagerFqdn sets AddressManagerFqdn field to given value.

### HasAddressManagerFqdn

`func (o *GetSettings200Response1DataInner) HasAddressManagerFqdn() bool`

HasAddressManagerFqdn returns a boolean if a field has been set.

### GetMetadataUrl

`func (o *GetSettings200Response1DataInner) GetMetadataUrl() string`

GetMetadataUrl returns the MetadataUrl field if non-nil, zero value otherwise.

### GetMetadataUrlOk

`func (o *GetSettings200Response1DataInner) GetMetadataUrlOk() (*string, bool)`

GetMetadataUrlOk returns a tuple with the MetadataUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataUrl

`func (o *GetSettings200Response1DataInner) SetMetadataUrl(v string)`

SetMetadataUrl sets MetadataUrl field to given value.

### HasMetadataUrl

`func (o *GetSettings200Response1DataInner) HasMetadataUrl() bool`

HasMetadataUrl returns a boolean if a field has been set.

### GetConsumeUrl

`func (o *GetSettings200Response1DataInner) GetConsumeUrl() string`

GetConsumeUrl returns the ConsumeUrl field if non-nil, zero value otherwise.

### GetConsumeUrlOk

`func (o *GetSettings200Response1DataInner) GetConsumeUrlOk() (*string, bool)`

GetConsumeUrlOk returns a tuple with the ConsumeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumeUrl

`func (o *GetSettings200Response1DataInner) SetConsumeUrl(v string)`

SetConsumeUrl sets ConsumeUrl field to given value.

### HasConsumeUrl

`func (o *GetSettings200Response1DataInner) HasConsumeUrl() bool`

HasConsumeUrl returns a boolean if a field has been set.

### GetLogoutUrl

`func (o *GetSettings200Response1DataInner) GetLogoutUrl() string`

GetLogoutUrl returns the LogoutUrl field if non-nil, zero value otherwise.

### GetLogoutUrlOk

`func (o *GetSettings200Response1DataInner) GetLogoutUrlOk() (*string, bool)`

GetLogoutUrlOk returns a tuple with the LogoutUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoutUrl

`func (o *GetSettings200Response1DataInner) SetLogoutUrl(v string)`

SetLogoutUrl sets LogoutUrl field to given value.

### HasLogoutUrl

`func (o *GetSettings200Response1DataInner) HasLogoutUrl() bool`

HasLogoutUrl returns a boolean if a field has been set.

### GetNameIdFormat

`func (o *GetSettings200Response1DataInner) GetNameIdFormat() string`

GetNameIdFormat returns the NameIdFormat field if non-nil, zero value otherwise.

### GetNameIdFormatOk

`func (o *GetSettings200Response1DataInner) GetNameIdFormatOk() (*string, bool)`

GetNameIdFormatOk returns a tuple with the NameIdFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameIdFormat

`func (o *GetSettings200Response1DataInner) SetNameIdFormat(v string)`

SetNameIdFormat sets NameIdFormat field to given value.

### HasNameIdFormat

`func (o *GetSettings200Response1DataInner) HasNameIdFormat() bool`

HasNameIdFormat returns a boolean if a field has been set.

### GetSigningEnabled

`func (o *GetSettings200Response1DataInner) GetSigningEnabled() bool`

GetSigningEnabled returns the SigningEnabled field if non-nil, zero value otherwise.

### GetSigningEnabledOk

`func (o *GetSettings200Response1DataInner) GetSigningEnabledOk() (*bool, bool)`

GetSigningEnabledOk returns a tuple with the SigningEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningEnabled

`func (o *GetSettings200Response1DataInner) SetSigningEnabled(v bool)`

SetSigningEnabled sets SigningEnabled field to given value.

### HasSigningEnabled

`func (o *GetSettings200Response1DataInner) HasSigningEnabled() bool`

HasSigningEnabled returns a boolean if a field has been set.

### GetEncryptionEnabled

`func (o *GetSettings200Response1DataInner) GetEncryptionEnabled() bool`

GetEncryptionEnabled returns the EncryptionEnabled field if non-nil, zero value otherwise.

### GetEncryptionEnabledOk

`func (o *GetSettings200Response1DataInner) GetEncryptionEnabledOk() (*bool, bool)`

GetEncryptionEnabledOk returns a tuple with the EncryptionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionEnabled

`func (o *GetSettings200Response1DataInner) SetEncryptionEnabled(v bool)`

SetEncryptionEnabled sets EncryptionEnabled field to given value.

### HasEncryptionEnabled

`func (o *GetSettings200Response1DataInner) HasEncryptionEnabled() bool`

HasEncryptionEnabled returns a boolean if a field has been set.

### GetPkcs12

`func (o *GetSettings200Response1DataInner) GetPkcs12() string`

GetPkcs12 returns the Pkcs12 field if non-nil, zero value otherwise.

### GetPkcs12Ok

`func (o *GetSettings200Response1DataInner) GetPkcs12Ok() (*string, bool)`

GetPkcs12Ok returns a tuple with the Pkcs12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkcs12

`func (o *GetSettings200Response1DataInner) SetPkcs12(v string)`

SetPkcs12 sets Pkcs12 field to given value.

### HasPkcs12

`func (o *GetSettings200Response1DataInner) HasPkcs12() bool`

HasPkcs12 returns a boolean if a field has been set.

### GetPassword

`func (o *GetSettings200Response1DataInner) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *GetSettings200Response1DataInner) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *GetSettings200Response1DataInner) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *GetSettings200Response1DataInner) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetOrganizationName

`func (o *GetSettings200Response1DataInner) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *GetSettings200Response1DataInner) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *GetSettings200Response1DataInner) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *GetSettings200Response1DataInner) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### GetOrganizationUrl

`func (o *GetSettings200Response1DataInner) GetOrganizationUrl() string`

GetOrganizationUrl returns the OrganizationUrl field if non-nil, zero value otherwise.

### GetOrganizationUrlOk

`func (o *GetSettings200Response1DataInner) GetOrganizationUrlOk() (*string, bool)`

GetOrganizationUrlOk returns a tuple with the OrganizationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationUrl

`func (o *GetSettings200Response1DataInner) SetOrganizationUrl(v string)`

SetOrganizationUrl sets OrganizationUrl field to given value.

### HasOrganizationUrl

`func (o *GetSettings200Response1DataInner) HasOrganizationUrl() bool`

HasOrganizationUrl returns a boolean if a field has been set.

### GetContactName

`func (o *GetSettings200Response1DataInner) GetContactName() string`

GetContactName returns the ContactName field if non-nil, zero value otherwise.

### GetContactNameOk

`func (o *GetSettings200Response1DataInner) GetContactNameOk() (*string, bool)`

GetContactNameOk returns a tuple with the ContactName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactName

`func (o *GetSettings200Response1DataInner) SetContactName(v string)`

SetContactName sets ContactName field to given value.

### HasContactName

`func (o *GetSettings200Response1DataInner) HasContactName() bool`

HasContactName returns a boolean if a field has been set.

### GetContactEmail

`func (o *GetSettings200Response1DataInner) GetContactEmail() string`

GetContactEmail returns the ContactEmail field if non-nil, zero value otherwise.

### GetContactEmailOk

`func (o *GetSettings200Response1DataInner) GetContactEmailOk() (*string, bool)`

GetContactEmailOk returns a tuple with the ContactEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactEmail

`func (o *GetSettings200Response1DataInner) SetContactEmail(v string)`

SetContactEmail sets ContactEmail field to given value.

### HasContactEmail

`func (o *GetSettings200Response1DataInner) HasContactEmail() bool`

HasContactEmail returns a boolean if a field has been set.

### GetSamlIdentityProviderEnabled

`func (o *GetSettings200Response1DataInner) GetSamlIdentityProviderEnabled() bool`

GetSamlIdentityProviderEnabled returns the SamlIdentityProviderEnabled field if non-nil, zero value otherwise.

### GetSamlIdentityProviderEnabledOk

`func (o *GetSettings200Response1DataInner) GetSamlIdentityProviderEnabledOk() (*bool, bool)`

GetSamlIdentityProviderEnabledOk returns a tuple with the SamlIdentityProviderEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlIdentityProviderEnabled

`func (o *GetSettings200Response1DataInner) SetSamlIdentityProviderEnabled(v bool)`

SetSamlIdentityProviderEnabled sets SamlIdentityProviderEnabled field to given value.

### HasSamlIdentityProviderEnabled

`func (o *GetSettings200Response1DataInner) HasSamlIdentityProviderEnabled() bool`

HasSamlIdentityProviderEnabled returns a boolean if a field has been set.

### GetNonSsoAuthenticatorCount

`func (o *GetSettings200Response1DataInner) GetNonSsoAuthenticatorCount() int32`

GetNonSsoAuthenticatorCount returns the NonSsoAuthenticatorCount field if non-nil, zero value otherwise.

### GetNonSsoAuthenticatorCountOk

`func (o *GetSettings200Response1DataInner) GetNonSsoAuthenticatorCountOk() (*int32, bool)`

GetNonSsoAuthenticatorCountOk returns a tuple with the NonSsoAuthenticatorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonSsoAuthenticatorCount

`func (o *GetSettings200Response1DataInner) SetNonSsoAuthenticatorCount(v int32)`

SetNonSsoAuthenticatorCount sets NonSsoAuthenticatorCount field to given value.

### HasNonSsoAuthenticatorCount

`func (o *GetSettings200Response1DataInner) HasNonSsoAuthenticatorCount() bool`

HasNonSsoAuthenticatorCount returns a boolean if a field has been set.

### GetNonSsoGroupCount

`func (o *GetSettings200Response1DataInner) GetNonSsoGroupCount() int32`

GetNonSsoGroupCount returns the NonSsoGroupCount field if non-nil, zero value otherwise.

### GetNonSsoGroupCountOk

`func (o *GetSettings200Response1DataInner) GetNonSsoGroupCountOk() (*int32, bool)`

GetNonSsoGroupCountOk returns a tuple with the NonSsoGroupCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonSsoGroupCount

`func (o *GetSettings200Response1DataInner) SetNonSsoGroupCount(v int32)`

SetNonSsoGroupCount sets NonSsoGroupCount field to given value.

### HasNonSsoGroupCount

`func (o *GetSettings200Response1DataInner) HasNonSsoGroupCount() bool`

HasNonSsoGroupCount returns a boolean if a field has been set.

### GetLocalAdminUserCount

`func (o *GetSettings200Response1DataInner) GetLocalAdminUserCount() int32`

GetLocalAdminUserCount returns the LocalAdminUserCount field if non-nil, zero value otherwise.

### GetLocalAdminUserCountOk

`func (o *GetSettings200Response1DataInner) GetLocalAdminUserCountOk() (*int32, bool)`

GetLocalAdminUserCountOk returns a tuple with the LocalAdminUserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAdminUserCount

`func (o *GetSettings200Response1DataInner) SetLocalAdminUserCount(v int32)`

SetLocalAdminUserCount sets LocalAdminUserCount field to given value.

### HasLocalAdminUserCount

`func (o *GetSettings200Response1DataInner) HasLocalAdminUserCount() bool`

HasLocalAdminUserCount returns a boolean if a field has been set.

### GetLocalNonAdminUserCount

`func (o *GetSettings200Response1DataInner) GetLocalNonAdminUserCount() int32`

GetLocalNonAdminUserCount returns the LocalNonAdminUserCount field if non-nil, zero value otherwise.

### GetLocalNonAdminUserCountOk

`func (o *GetSettings200Response1DataInner) GetLocalNonAdminUserCountOk() (*int32, bool)`

GetLocalNonAdminUserCountOk returns a tuple with the LocalNonAdminUserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalNonAdminUserCount

`func (o *GetSettings200Response1DataInner) SetLocalNonAdminUserCount(v int32)`

SetLocalNonAdminUserCount sets LocalNonAdminUserCount field to given value.

### HasLocalNonAdminUserCount

`func (o *GetSettings200Response1DataInner) HasLocalNonAdminUserCount() bool`

HasLocalNonAdminUserCount returns a boolean if a field has been set.

### GetHostname

`func (o *GetSettings200Response1DataInner) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *GetSettings200Response1DataInner) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *GetSettings200Response1DataInner) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *GetSettings200Response1DataInner) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetVersion

`func (o *GetSettings200Response1DataInner) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetSettings200Response1DataInner) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetSettings200Response1DataInner) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *GetSettings200Response1DataInner) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetAddress

`func (o *GetSettings200Response1DataInner) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *GetSettings200Response1DataInner) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *GetSettings200Response1DataInner) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *GetSettings200Response1DataInner) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetInterfaceRedundancyEnabled

`func (o *GetSettings200Response1DataInner) GetInterfaceRedundancyEnabled() bool`

GetInterfaceRedundancyEnabled returns the InterfaceRedundancyEnabled field if non-nil, zero value otherwise.

### GetInterfaceRedundancyEnabledOk

`func (o *GetSettings200Response1DataInner) GetInterfaceRedundancyEnabledOk() (*bool, bool)`

GetInterfaceRedundancyEnabledOk returns a tuple with the InterfaceRedundancyEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceRedundancyEnabled

`func (o *GetSettings200Response1DataInner) SetInterfaceRedundancyEnabled(v bool)`

SetInterfaceRedundancyEnabled sets InterfaceRedundancyEnabled field to given value.

### HasInterfaceRedundancyEnabled

`func (o *GetSettings200Response1DataInner) HasInterfaceRedundancyEnabled() bool`

HasInterfaceRedundancyEnabled returns a boolean if a field has been set.

### GetActiveSessions

`func (o *GetSettings200Response1DataInner) GetActiveSessions() int32`

GetActiveSessions returns the ActiveSessions field if non-nil, zero value otherwise.

### GetActiveSessionsOk

`func (o *GetSettings200Response1DataInner) GetActiveSessionsOk() (*int32, bool)`

GetActiveSessionsOk returns a tuple with the ActiveSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveSessions

`func (o *GetSettings200Response1DataInner) SetActiveSessions(v int32)`

SetActiveSessions sets ActiveSessions field to given value.

### HasActiveSessions

`func (o *GetSettings200Response1DataInner) HasActiveSessions() bool`

HasActiveSessions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


