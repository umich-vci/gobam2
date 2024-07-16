# GetSettings200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**AdministrativeHistoryRetentionPeriod** | Pointer to **string** | The number of days to retain the administrative history records. A null retention period indicates that the data will be retained indefinitely. | [optional] 
**SessionAndEventsRetentionPeriod** | Pointer to **string** | The number of days to retain the session and event history records. A null retention period indicates that the data will be retained indefinitely. | [optional] 
**DhcpTransactionRetentionPeriod** | Pointer to **string** | The number of days to retain the DHCP transaction history records. A null retention period indicates that the data will be retained indefinitely. A retention period of 0 seconds indicates that the data wil not be retained. | [optional] 
**DdnsTransactionRetentionPeriod** | Pointer to **string** | The number of days to retain the DDNS transaction history records. A null retention period indicates that the data will be retained indefinitely. A retention period of 0 seconds indicates that the data wil not be retained. | [optional] 
**ExportEnabled** | Pointer to **bool** | Indicates if the audit data export service is enabled. | [optional] 
**Destination** | Pointer to [**SinkBean**](SinkBean.md) |  | [optional] 
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
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 

## Methods

### NewGetSettings200ResponseDataInner

`func NewGetSettings200ResponseDataInner() *GetSettings200ResponseDataInner`

NewGetSettings200ResponseDataInner instantiates a new GetSettings200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSettings200ResponseDataInnerWithDefaults

`func NewGetSettings200ResponseDataInnerWithDefaults() *GetSettings200ResponseDataInner`

NewGetSettings200ResponseDataInnerWithDefaults instantiates a new GetSettings200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetSettings200ResponseDataInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetSettings200ResponseDataInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetSettings200ResponseDataInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetSettings200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GetSettings200ResponseDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetSettings200ResponseDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetSettings200ResponseDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetSettings200ResponseDataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAdministrativeHistoryRetentionPeriod

`func (o *GetSettings200ResponseDataInner) GetAdministrativeHistoryRetentionPeriod() string`

GetAdministrativeHistoryRetentionPeriod returns the AdministrativeHistoryRetentionPeriod field if non-nil, zero value otherwise.

### GetAdministrativeHistoryRetentionPeriodOk

`func (o *GetSettings200ResponseDataInner) GetAdministrativeHistoryRetentionPeriodOk() (*string, bool)`

GetAdministrativeHistoryRetentionPeriodOk returns a tuple with the AdministrativeHistoryRetentionPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrativeHistoryRetentionPeriod

`func (o *GetSettings200ResponseDataInner) SetAdministrativeHistoryRetentionPeriod(v string)`

SetAdministrativeHistoryRetentionPeriod sets AdministrativeHistoryRetentionPeriod field to given value.

### HasAdministrativeHistoryRetentionPeriod

`func (o *GetSettings200ResponseDataInner) HasAdministrativeHistoryRetentionPeriod() bool`

HasAdministrativeHistoryRetentionPeriod returns a boolean if a field has been set.

### GetSessionAndEventsRetentionPeriod

`func (o *GetSettings200ResponseDataInner) GetSessionAndEventsRetentionPeriod() string`

GetSessionAndEventsRetentionPeriod returns the SessionAndEventsRetentionPeriod field if non-nil, zero value otherwise.

### GetSessionAndEventsRetentionPeriodOk

`func (o *GetSettings200ResponseDataInner) GetSessionAndEventsRetentionPeriodOk() (*string, bool)`

GetSessionAndEventsRetentionPeriodOk returns a tuple with the SessionAndEventsRetentionPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionAndEventsRetentionPeriod

`func (o *GetSettings200ResponseDataInner) SetSessionAndEventsRetentionPeriod(v string)`

SetSessionAndEventsRetentionPeriod sets SessionAndEventsRetentionPeriod field to given value.

### HasSessionAndEventsRetentionPeriod

`func (o *GetSettings200ResponseDataInner) HasSessionAndEventsRetentionPeriod() bool`

HasSessionAndEventsRetentionPeriod returns a boolean if a field has been set.

### GetDhcpTransactionRetentionPeriod

`func (o *GetSettings200ResponseDataInner) GetDhcpTransactionRetentionPeriod() string`

GetDhcpTransactionRetentionPeriod returns the DhcpTransactionRetentionPeriod field if non-nil, zero value otherwise.

### GetDhcpTransactionRetentionPeriodOk

`func (o *GetSettings200ResponseDataInner) GetDhcpTransactionRetentionPeriodOk() (*string, bool)`

GetDhcpTransactionRetentionPeriodOk returns a tuple with the DhcpTransactionRetentionPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpTransactionRetentionPeriod

`func (o *GetSettings200ResponseDataInner) SetDhcpTransactionRetentionPeriod(v string)`

SetDhcpTransactionRetentionPeriod sets DhcpTransactionRetentionPeriod field to given value.

### HasDhcpTransactionRetentionPeriod

`func (o *GetSettings200ResponseDataInner) HasDhcpTransactionRetentionPeriod() bool`

HasDhcpTransactionRetentionPeriod returns a boolean if a field has been set.

### GetDdnsTransactionRetentionPeriod

`func (o *GetSettings200ResponseDataInner) GetDdnsTransactionRetentionPeriod() string`

GetDdnsTransactionRetentionPeriod returns the DdnsTransactionRetentionPeriod field if non-nil, zero value otherwise.

### GetDdnsTransactionRetentionPeriodOk

`func (o *GetSettings200ResponseDataInner) GetDdnsTransactionRetentionPeriodOk() (*string, bool)`

GetDdnsTransactionRetentionPeriodOk returns a tuple with the DdnsTransactionRetentionPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsTransactionRetentionPeriod

`func (o *GetSettings200ResponseDataInner) SetDdnsTransactionRetentionPeriod(v string)`

SetDdnsTransactionRetentionPeriod sets DdnsTransactionRetentionPeriod field to given value.

### HasDdnsTransactionRetentionPeriod

`func (o *GetSettings200ResponseDataInner) HasDdnsTransactionRetentionPeriod() bool`

HasDdnsTransactionRetentionPeriod returns a boolean if a field has been set.

### GetExportEnabled

`func (o *GetSettings200ResponseDataInner) GetExportEnabled() bool`

GetExportEnabled returns the ExportEnabled field if non-nil, zero value otherwise.

### GetExportEnabledOk

`func (o *GetSettings200ResponseDataInner) GetExportEnabledOk() (*bool, bool)`

GetExportEnabledOk returns a tuple with the ExportEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportEnabled

`func (o *GetSettings200ResponseDataInner) SetExportEnabled(v bool)`

SetExportEnabled sets ExportEnabled field to given value.

### HasExportEnabled

`func (o *GetSettings200ResponseDataInner) HasExportEnabled() bool`

HasExportEnabled returns a boolean if a field has been set.

### GetDestination

`func (o *GetSettings200ResponseDataInner) GetDestination() SinkBean`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *GetSettings200ResponseDataInner) GetDestinationOk() (*SinkBean, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *GetSettings200ResponseDataInner) SetDestination(v SinkBean)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *GetSettings200ResponseDataInner) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetMandatoryCommentsEnabled

`func (o *GetSettings200ResponseDataInner) GetMandatoryCommentsEnabled() bool`

GetMandatoryCommentsEnabled returns the MandatoryCommentsEnabled field if non-nil, zero value otherwise.

### GetMandatoryCommentsEnabledOk

`func (o *GetSettings200ResponseDataInner) GetMandatoryCommentsEnabledOk() (*bool, bool)`

GetMandatoryCommentsEnabledOk returns a tuple with the MandatoryCommentsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandatoryCommentsEnabled

`func (o *GetSettings200ResponseDataInner) SetMandatoryCommentsEnabled(v bool)`

SetMandatoryCommentsEnabled sets MandatoryCommentsEnabled field to given value.

### HasMandatoryCommentsEnabled

`func (o *GetSettings200ResponseDataInner) HasMandatoryCommentsEnabled() bool`

HasMandatoryCommentsEnabled returns a boolean if a field has been set.

### GetRememberLastAddressEnabled

`func (o *GetSettings200ResponseDataInner) GetRememberLastAddressEnabled() bool`

GetRememberLastAddressEnabled returns the RememberLastAddressEnabled field if non-nil, zero value otherwise.

### GetRememberLastAddressEnabledOk

`func (o *GetSettings200ResponseDataInner) GetRememberLastAddressEnabledOk() (*bool, bool)`

GetRememberLastAddressEnabledOk returns a tuple with the RememberLastAddressEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRememberLastAddressEnabled

`func (o *GetSettings200ResponseDataInner) SetRememberLastAddressEnabled(v bool)`

SetRememberLastAddressEnabled sets RememberLastAddressEnabled field to given value.

### HasRememberLastAddressEnabled

`func (o *GetSettings200ResponseDataInner) HasRememberLastAddressEnabled() bool`

HasRememberLastAddressEnabled returns a boolean if a field has been set.

### GetSessionTimeout

`func (o *GetSettings200ResponseDataInner) GetSessionTimeout() string`

GetSessionTimeout returns the SessionTimeout field if non-nil, zero value otherwise.

### GetSessionTimeoutOk

`func (o *GetSettings200ResponseDataInner) GetSessionTimeoutOk() (*string, bool)`

GetSessionTimeoutOk returns a tuple with the SessionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionTimeout

`func (o *GetSettings200ResponseDataInner) SetSessionTimeout(v string)`

SetSessionTimeout sets SessionTimeout field to given value.

### HasSessionTimeout

`func (o *GetSettings200ResponseDataInner) HasSessionTimeout() bool`

HasSessionTimeout returns a boolean if a field has been set.

### GetDisclaimerEnabled

`func (o *GetSettings200ResponseDataInner) GetDisclaimerEnabled() bool`

GetDisclaimerEnabled returns the DisclaimerEnabled field if non-nil, zero value otherwise.

### GetDisclaimerEnabledOk

`func (o *GetSettings200ResponseDataInner) GetDisclaimerEnabledOk() (*bool, bool)`

GetDisclaimerEnabledOk returns a tuple with the DisclaimerEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisclaimerEnabled

`func (o *GetSettings200ResponseDataInner) SetDisclaimerEnabled(v bool)`

SetDisclaimerEnabled sets DisclaimerEnabled field to given value.

### HasDisclaimerEnabled

`func (o *GetSettings200ResponseDataInner) HasDisclaimerEnabled() bool`

HasDisclaimerEnabled returns a boolean if a field has been set.

### GetDisclaimerText

`func (o *GetSettings200ResponseDataInner) GetDisclaimerText() string`

GetDisclaimerText returns the DisclaimerText field if non-nil, zero value otherwise.

### GetDisclaimerTextOk

`func (o *GetSettings200ResponseDataInner) GetDisclaimerTextOk() (*string, bool)`

GetDisclaimerTextOk returns a tuple with the DisclaimerText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisclaimerText

`func (o *GetSettings200ResponseDataInner) SetDisclaimerText(v string)`

SetDisclaimerText sets DisclaimerText field to given value.

### HasDisclaimerText

`func (o *GetSettings200ResponseDataInner) HasDisclaimerText() bool`

HasDisclaimerText returns a boolean if a field has been set.

### GetCustomReverseZoneFormatsAllowed

`func (o *GetSettings200ResponseDataInner) GetCustomReverseZoneFormatsAllowed() bool`

GetCustomReverseZoneFormatsAllowed returns the CustomReverseZoneFormatsAllowed field if non-nil, zero value otherwise.

### GetCustomReverseZoneFormatsAllowedOk

`func (o *GetSettings200ResponseDataInner) GetCustomReverseZoneFormatsAllowedOk() (*bool, bool)`

GetCustomReverseZoneFormatsAllowedOk returns a tuple with the CustomReverseZoneFormatsAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomReverseZoneFormatsAllowed

`func (o *GetSettings200ResponseDataInner) SetCustomReverseZoneFormatsAllowed(v bool)`

SetCustomReverseZoneFormatsAllowed sets CustomReverseZoneFormatsAllowed field to given value.

### HasCustomReverseZoneFormatsAllowed

`func (o *GetSettings200ResponseDataInner) HasCustomReverseZoneFormatsAllowed() bool`

HasCustomReverseZoneFormatsAllowed returns a boolean if a field has been set.

### GetAddressManager

`func (o *GetSettings200ResponseDataInner) GetAddressManager() string`

GetAddressManager returns the AddressManager field if non-nil, zero value otherwise.

### GetAddressManagerOk

`func (o *GetSettings200ResponseDataInner) GetAddressManagerOk() (*string, bool)`

GetAddressManagerOk returns a tuple with the AddressManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressManager

`func (o *GetSettings200ResponseDataInner) SetAddressManager(v string)`

SetAddressManager sets AddressManager field to given value.

### HasAddressManager

`func (o *GetSettings200ResponseDataInner) HasAddressManager() bool`

HasAddressManager returns a boolean if a field has been set.

### GetRestV2

`func (o *GetSettings200ResponseDataInner) GetRestV2() string`

GetRestV2 returns the RestV2 field if non-nil, zero value otherwise.

### GetRestV2Ok

`func (o *GetSettings200ResponseDataInner) GetRestV2Ok() (*string, bool)`

GetRestV2Ok returns a tuple with the RestV2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestV2

`func (o *GetSettings200ResponseDataInner) SetRestV2(v string)`

SetRestV2 sets RestV2 field to given value.

### HasRestV2

`func (o *GetSettings200ResponseDataInner) HasRestV2() bool`

HasRestV2 returns a boolean if a field has been set.

### GetZoneImport

`func (o *GetSettings200ResponseDataInner) GetZoneImport() string`

GetZoneImport returns the ZoneImport field if non-nil, zero value otherwise.

### GetZoneImportOk

`func (o *GetSettings200ResponseDataInner) GetZoneImportOk() (*string, bool)`

GetZoneImportOk returns a tuple with the ZoneImport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneImport

`func (o *GetSettings200ResponseDataInner) SetZoneImport(v string)`

SetZoneImport sets ZoneImport field to given value.

### HasZoneImport

`func (o *GetSettings200ResponseDataInner) HasZoneImport() bool`

HasZoneImport returns a boolean if a field has been set.

### GetReconciliationService

`func (o *GetSettings200ResponseDataInner) GetReconciliationService() string`

GetReconciliationService returns the ReconciliationService field if non-nil, zero value otherwise.

### GetReconciliationServiceOk

`func (o *GetSettings200ResponseDataInner) GetReconciliationServiceOk() (*string, bool)`

GetReconciliationServiceOk returns a tuple with the ReconciliationService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReconciliationService

`func (o *GetSettings200ResponseDataInner) SetReconciliationService(v string)`

SetReconciliationService sets ReconciliationService field to given value.

### HasReconciliationService

`func (o *GetSettings200ResponseDataInner) HasReconciliationService() bool`

HasReconciliationService returns a boolean if a field has been set.

### GetDiscoveryEngine

`func (o *GetSettings200ResponseDataInner) GetDiscoveryEngine() string`

GetDiscoveryEngine returns the DiscoveryEngine field if non-nil, zero value otherwise.

### GetDiscoveryEngineOk

`func (o *GetSettings200ResponseDataInner) GetDiscoveryEngineOk() (*string, bool)`

GetDiscoveryEngineOk returns a tuple with the DiscoveryEngine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryEngine

`func (o *GetSettings200ResponseDataInner) SetDiscoveryEngine(v string)`

SetDiscoveryEngine sets DiscoveryEngine field to given value.

### HasDiscoveryEngine

`func (o *GetSettings200ResponseDataInner) HasDiscoveryEngine() bool`

HasDiscoveryEngine returns a boolean if a field has been set.

### GetSnmpClient

`func (o *GetSettings200ResponseDataInner) GetSnmpClient() string`

GetSnmpClient returns the SnmpClient field if non-nil, zero value otherwise.

### GetSnmpClientOk

`func (o *GetSettings200ResponseDataInner) GetSnmpClientOk() (*string, bool)`

GetSnmpClientOk returns a tuple with the SnmpClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpClient

`func (o *GetSettings200ResponseDataInner) SetSnmpClient(v string)`

SetSnmpClient sets SnmpClient field to given value.

### HasSnmpClient

`func (o *GetSettings200ResponseDataInner) HasSnmpClient() bool`

HasSnmpClient returns a boolean if a field has been set.

### GetMonitoringService

`func (o *GetSettings200ResponseDataInner) GetMonitoringService() string`

GetMonitoringService returns the MonitoringService field if non-nil, zero value otherwise.

### GetMonitoringServiceOk

`func (o *GetSettings200ResponseDataInner) GetMonitoringServiceOk() (*string, bool)`

GetMonitoringServiceOk returns a tuple with the MonitoringService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringService

`func (o *GetSettings200ResponseDataInner) SetMonitoringService(v string)`

SetMonitoringService sets MonitoringService field to given value.

### HasMonitoringService

`func (o *GetSettings200ResponseDataInner) HasMonitoringService() bool`

HasMonitoringService returns a boolean if a field has been set.

### GetRrdTool

`func (o *GetSettings200ResponseDataInner) GetRrdTool() string`

GetRrdTool returns the RrdTool field if non-nil, zero value otherwise.

### GetRrdToolOk

`func (o *GetSettings200ResponseDataInner) GetRrdToolOk() (*string, bool)`

GetRrdToolOk returns a tuple with the RrdTool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrdTool

`func (o *GetSettings200ResponseDataInner) SetRrdTool(v string)`

SetRrdTool sets RrdTool field to given value.

### HasRrdTool

`func (o *GetSettings200ResponseDataInner) HasRrdTool() bool`

HasRrdTool returns a boolean if a field has been set.

### GetEnabled

`func (o *GetSettings200ResponseDataInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetSettings200ResponseDataInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetSettings200ResponseDataInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetSettings200ResponseDataInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFailureLimit

`func (o *GetSettings200ResponseDataInner) GetFailureLimit() int32`

GetFailureLimit returns the FailureLimit field if non-nil, zero value otherwise.

### GetFailureLimitOk

`func (o *GetSettings200ResponseDataInner) GetFailureLimitOk() (*int32, bool)`

GetFailureLimitOk returns a tuple with the FailureLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureLimit

`func (o *GetSettings200ResponseDataInner) SetFailureLimit(v int32)`

SetFailureLimit sets FailureLimit field to given value.

### HasFailureLimit

`func (o *GetSettings200ResponseDataInner) HasFailureLimit() bool`

HasFailureLimit returns a boolean if a field has been set.

### GetFailureLimitPeriod

`func (o *GetSettings200ResponseDataInner) GetFailureLimitPeriod() string`

GetFailureLimitPeriod returns the FailureLimitPeriod field if non-nil, zero value otherwise.

### GetFailureLimitPeriodOk

`func (o *GetSettings200ResponseDataInner) GetFailureLimitPeriodOk() (*string, bool)`

GetFailureLimitPeriodOk returns a tuple with the FailureLimitPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureLimitPeriod

`func (o *GetSettings200ResponseDataInner) SetFailureLimitPeriod(v string)`

SetFailureLimitPeriod sets FailureLimitPeriod field to given value.

### HasFailureLimitPeriod

`func (o *GetSettings200ResponseDataInner) HasFailureLimitPeriod() bool`

HasFailureLimitPeriod returns a boolean if a field has been set.

### GetDelayDuration

`func (o *GetSettings200ResponseDataInner) GetDelayDuration() string`

GetDelayDuration returns the DelayDuration field if non-nil, zero value otherwise.

### GetDelayDurationOk

`func (o *GetSettings200ResponseDataInner) GetDelayDurationOk() (*string, bool)`

GetDelayDurationOk returns a tuple with the DelayDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayDuration

`func (o *GetSettings200ResponseDataInner) SetDelayDuration(v string)`

SetDelayDuration sets DelayDuration field to given value.

### HasDelayDuration

`func (o *GetSettings200ResponseDataInner) HasDelayDuration() bool`

HasDelayDuration returns a boolean if a field has been set.

### GetAddressManagerMonitoring

`func (o *GetSettings200ResponseDataInner) GetAddressManagerMonitoring() AddressManagerMonitoringBean`

GetAddressManagerMonitoring returns the AddressManagerMonitoring field if non-nil, zero value otherwise.

### GetAddressManagerMonitoringOk

`func (o *GetSettings200ResponseDataInner) GetAddressManagerMonitoringOk() (*AddressManagerMonitoringBean, bool)`

GetAddressManagerMonitoringOk returns a tuple with the AddressManagerMonitoring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressManagerMonitoring

`func (o *GetSettings200ResponseDataInner) SetAddressManagerMonitoring(v AddressManagerMonitoringBean)`

SetAddressManagerMonitoring sets AddressManagerMonitoring field to given value.

### HasAddressManagerMonitoring

`func (o *GetSettings200ResponseDataInner) HasAddressManagerMonitoring() bool`

HasAddressManagerMonitoring returns a boolean if a field has been set.

### GetServerMonitoring

`func (o *GetSettings200ResponseDataInner) GetServerMonitoring() ServerMonitoringBean`

GetServerMonitoring returns the ServerMonitoring field if non-nil, zero value otherwise.

### GetServerMonitoringOk

`func (o *GetSettings200ResponseDataInner) GetServerMonitoringOk() (*ServerMonitoringBean, bool)`

GetServerMonitoringOk returns a tuple with the ServerMonitoring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerMonitoring

`func (o *GetSettings200ResponseDataInner) SetServerMonitoring(v ServerMonitoringBean)`

SetServerMonitoring sets ServerMonitoring field to given value.

### HasServerMonitoring

`func (o *GetSettings200ResponseDataInner) HasServerMonitoring() bool`

HasServerMonitoring returns a boolean if a field has been set.

### GetF5ServerMonitoring

`func (o *GetSettings200ResponseDataInner) GetF5ServerMonitoring() F5ServerMonitoringBean`

GetF5ServerMonitoring returns the F5ServerMonitoring field if non-nil, zero value otherwise.

### GetF5ServerMonitoringOk

`func (o *GetSettings200ResponseDataInner) GetF5ServerMonitoringOk() (*F5ServerMonitoringBean, bool)`

GetF5ServerMonitoringOk returns a tuple with the F5ServerMonitoring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetF5ServerMonitoring

`func (o *GetSettings200ResponseDataInner) SetF5ServerMonitoring(v F5ServerMonitoringBean)`

SetF5ServerMonitoring sets F5ServerMonitoring field to given value.

### HasF5ServerMonitoring

`func (o *GetSettings200ResponseDataInner) HasF5ServerMonitoring() bool`

HasF5ServerMonitoring returns a boolean if a field has been set.

### GetMinLength

`func (o *GetSettings200ResponseDataInner) GetMinLength() int32`

GetMinLength returns the MinLength field if non-nil, zero value otherwise.

### GetMinLengthOk

`func (o *GetSettings200ResponseDataInner) GetMinLengthOk() (*int32, bool)`

GetMinLengthOk returns a tuple with the MinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLength

`func (o *GetSettings200ResponseDataInner) SetMinLength(v int32)`

SetMinLength sets MinLength field to given value.

### HasMinLength

`func (o *GetSettings200ResponseDataInner) HasMinLength() bool`

HasMinLength returns a boolean if a field has been set.

### GetMaxLength

`func (o *GetSettings200ResponseDataInner) GetMaxLength() int32`

GetMaxLength returns the MaxLength field if non-nil, zero value otherwise.

### GetMaxLengthOk

`func (o *GetSettings200ResponseDataInner) GetMaxLengthOk() (*int32, bool)`

GetMaxLengthOk returns a tuple with the MaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLength

`func (o *GetSettings200ResponseDataInner) SetMaxLength(v int32)`

SetMaxLength sets MaxLength field to given value.

### HasMaxLength

`func (o *GetSettings200ResponseDataInner) HasMaxLength() bool`

HasMaxLength returns a boolean if a field has been set.

### GetDigitRequired

`func (o *GetSettings200ResponseDataInner) GetDigitRequired() bool`

GetDigitRequired returns the DigitRequired field if non-nil, zero value otherwise.

### GetDigitRequiredOk

`func (o *GetSettings200ResponseDataInner) GetDigitRequiredOk() (*bool, bool)`

GetDigitRequiredOk returns a tuple with the DigitRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitRequired

`func (o *GetSettings200ResponseDataInner) SetDigitRequired(v bool)`

SetDigitRequired sets DigitRequired field to given value.

### HasDigitRequired

`func (o *GetSettings200ResponseDataInner) HasDigitRequired() bool`

HasDigitRequired returns a boolean if a field has been set.

### GetMixedCaseRequired

`func (o *GetSettings200ResponseDataInner) GetMixedCaseRequired() bool`

GetMixedCaseRequired returns the MixedCaseRequired field if non-nil, zero value otherwise.

### GetMixedCaseRequiredOk

`func (o *GetSettings200ResponseDataInner) GetMixedCaseRequiredOk() (*bool, bool)`

GetMixedCaseRequiredOk returns a tuple with the MixedCaseRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMixedCaseRequired

`func (o *GetSettings200ResponseDataInner) SetMixedCaseRequired(v bool)`

SetMixedCaseRequired sets MixedCaseRequired field to given value.

### HasMixedCaseRequired

`func (o *GetSettings200ResponseDataInner) HasMixedCaseRequired() bool`

HasMixedCaseRequired returns a boolean if a field has been set.

### GetSpecialCharacterRequired

`func (o *GetSettings200ResponseDataInner) GetSpecialCharacterRequired() bool`

GetSpecialCharacterRequired returns the SpecialCharacterRequired field if non-nil, zero value otherwise.

### GetSpecialCharacterRequiredOk

`func (o *GetSettings200ResponseDataInner) GetSpecialCharacterRequiredOk() (*bool, bool)`

GetSpecialCharacterRequiredOk returns a tuple with the SpecialCharacterRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialCharacterRequired

`func (o *GetSettings200ResponseDataInner) SetSpecialCharacterRequired(v bool)`

SetSpecialCharacterRequired sets SpecialCharacterRequired field to given value.

### HasSpecialCharacterRequired

`func (o *GetSettings200ResponseDataInner) HasSpecialCharacterRequired() bool`

HasSpecialCharacterRequired returns a boolean if a field has been set.

### GetAddressManagerFqdn

`func (o *GetSettings200ResponseDataInner) GetAddressManagerFqdn() string`

GetAddressManagerFqdn returns the AddressManagerFqdn field if non-nil, zero value otherwise.

### GetAddressManagerFqdnOk

`func (o *GetSettings200ResponseDataInner) GetAddressManagerFqdnOk() (*string, bool)`

GetAddressManagerFqdnOk returns a tuple with the AddressManagerFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressManagerFqdn

`func (o *GetSettings200ResponseDataInner) SetAddressManagerFqdn(v string)`

SetAddressManagerFqdn sets AddressManagerFqdn field to given value.

### HasAddressManagerFqdn

`func (o *GetSettings200ResponseDataInner) HasAddressManagerFqdn() bool`

HasAddressManagerFqdn returns a boolean if a field has been set.

### GetMetadataUrl

`func (o *GetSettings200ResponseDataInner) GetMetadataUrl() string`

GetMetadataUrl returns the MetadataUrl field if non-nil, zero value otherwise.

### GetMetadataUrlOk

`func (o *GetSettings200ResponseDataInner) GetMetadataUrlOk() (*string, bool)`

GetMetadataUrlOk returns a tuple with the MetadataUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataUrl

`func (o *GetSettings200ResponseDataInner) SetMetadataUrl(v string)`

SetMetadataUrl sets MetadataUrl field to given value.

### HasMetadataUrl

`func (o *GetSettings200ResponseDataInner) HasMetadataUrl() bool`

HasMetadataUrl returns a boolean if a field has been set.

### GetConsumeUrl

`func (o *GetSettings200ResponseDataInner) GetConsumeUrl() string`

GetConsumeUrl returns the ConsumeUrl field if non-nil, zero value otherwise.

### GetConsumeUrlOk

`func (o *GetSettings200ResponseDataInner) GetConsumeUrlOk() (*string, bool)`

GetConsumeUrlOk returns a tuple with the ConsumeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumeUrl

`func (o *GetSettings200ResponseDataInner) SetConsumeUrl(v string)`

SetConsumeUrl sets ConsumeUrl field to given value.

### HasConsumeUrl

`func (o *GetSettings200ResponseDataInner) HasConsumeUrl() bool`

HasConsumeUrl returns a boolean if a field has been set.

### GetLogoutUrl

`func (o *GetSettings200ResponseDataInner) GetLogoutUrl() string`

GetLogoutUrl returns the LogoutUrl field if non-nil, zero value otherwise.

### GetLogoutUrlOk

`func (o *GetSettings200ResponseDataInner) GetLogoutUrlOk() (*string, bool)`

GetLogoutUrlOk returns a tuple with the LogoutUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoutUrl

`func (o *GetSettings200ResponseDataInner) SetLogoutUrl(v string)`

SetLogoutUrl sets LogoutUrl field to given value.

### HasLogoutUrl

`func (o *GetSettings200ResponseDataInner) HasLogoutUrl() bool`

HasLogoutUrl returns a boolean if a field has been set.

### GetNameIdFormat

`func (o *GetSettings200ResponseDataInner) GetNameIdFormat() string`

GetNameIdFormat returns the NameIdFormat field if non-nil, zero value otherwise.

### GetNameIdFormatOk

`func (o *GetSettings200ResponseDataInner) GetNameIdFormatOk() (*string, bool)`

GetNameIdFormatOk returns a tuple with the NameIdFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameIdFormat

`func (o *GetSettings200ResponseDataInner) SetNameIdFormat(v string)`

SetNameIdFormat sets NameIdFormat field to given value.

### HasNameIdFormat

`func (o *GetSettings200ResponseDataInner) HasNameIdFormat() bool`

HasNameIdFormat returns a boolean if a field has been set.

### GetSigningEnabled

`func (o *GetSettings200ResponseDataInner) GetSigningEnabled() bool`

GetSigningEnabled returns the SigningEnabled field if non-nil, zero value otherwise.

### GetSigningEnabledOk

`func (o *GetSettings200ResponseDataInner) GetSigningEnabledOk() (*bool, bool)`

GetSigningEnabledOk returns a tuple with the SigningEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningEnabled

`func (o *GetSettings200ResponseDataInner) SetSigningEnabled(v bool)`

SetSigningEnabled sets SigningEnabled field to given value.

### HasSigningEnabled

`func (o *GetSettings200ResponseDataInner) HasSigningEnabled() bool`

HasSigningEnabled returns a boolean if a field has been set.

### GetEncryptionEnabled

`func (o *GetSettings200ResponseDataInner) GetEncryptionEnabled() bool`

GetEncryptionEnabled returns the EncryptionEnabled field if non-nil, zero value otherwise.

### GetEncryptionEnabledOk

`func (o *GetSettings200ResponseDataInner) GetEncryptionEnabledOk() (*bool, bool)`

GetEncryptionEnabledOk returns a tuple with the EncryptionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionEnabled

`func (o *GetSettings200ResponseDataInner) SetEncryptionEnabled(v bool)`

SetEncryptionEnabled sets EncryptionEnabled field to given value.

### HasEncryptionEnabled

`func (o *GetSettings200ResponseDataInner) HasEncryptionEnabled() bool`

HasEncryptionEnabled returns a boolean if a field has been set.

### GetPkcs12

`func (o *GetSettings200ResponseDataInner) GetPkcs12() string`

GetPkcs12 returns the Pkcs12 field if non-nil, zero value otherwise.

### GetPkcs12Ok

`func (o *GetSettings200ResponseDataInner) GetPkcs12Ok() (*string, bool)`

GetPkcs12Ok returns a tuple with the Pkcs12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkcs12

`func (o *GetSettings200ResponseDataInner) SetPkcs12(v string)`

SetPkcs12 sets Pkcs12 field to given value.

### HasPkcs12

`func (o *GetSettings200ResponseDataInner) HasPkcs12() bool`

HasPkcs12 returns a boolean if a field has been set.

### GetPassword

`func (o *GetSettings200ResponseDataInner) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *GetSettings200ResponseDataInner) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *GetSettings200ResponseDataInner) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *GetSettings200ResponseDataInner) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetOrganizationName

`func (o *GetSettings200ResponseDataInner) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *GetSettings200ResponseDataInner) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *GetSettings200ResponseDataInner) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *GetSettings200ResponseDataInner) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### GetOrganizationUrl

`func (o *GetSettings200ResponseDataInner) GetOrganizationUrl() string`

GetOrganizationUrl returns the OrganizationUrl field if non-nil, zero value otherwise.

### GetOrganizationUrlOk

`func (o *GetSettings200ResponseDataInner) GetOrganizationUrlOk() (*string, bool)`

GetOrganizationUrlOk returns a tuple with the OrganizationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationUrl

`func (o *GetSettings200ResponseDataInner) SetOrganizationUrl(v string)`

SetOrganizationUrl sets OrganizationUrl field to given value.

### HasOrganizationUrl

`func (o *GetSettings200ResponseDataInner) HasOrganizationUrl() bool`

HasOrganizationUrl returns a boolean if a field has been set.

### GetContactName

`func (o *GetSettings200ResponseDataInner) GetContactName() string`

GetContactName returns the ContactName field if non-nil, zero value otherwise.

### GetContactNameOk

`func (o *GetSettings200ResponseDataInner) GetContactNameOk() (*string, bool)`

GetContactNameOk returns a tuple with the ContactName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactName

`func (o *GetSettings200ResponseDataInner) SetContactName(v string)`

SetContactName sets ContactName field to given value.

### HasContactName

`func (o *GetSettings200ResponseDataInner) HasContactName() bool`

HasContactName returns a boolean if a field has been set.

### GetContactEmail

`func (o *GetSettings200ResponseDataInner) GetContactEmail() string`

GetContactEmail returns the ContactEmail field if non-nil, zero value otherwise.

### GetContactEmailOk

`func (o *GetSettings200ResponseDataInner) GetContactEmailOk() (*string, bool)`

GetContactEmailOk returns a tuple with the ContactEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactEmail

`func (o *GetSettings200ResponseDataInner) SetContactEmail(v string)`

SetContactEmail sets ContactEmail field to given value.

### HasContactEmail

`func (o *GetSettings200ResponseDataInner) HasContactEmail() bool`

HasContactEmail returns a boolean if a field has been set.

### GetSamlIdentityProviderEnabled

`func (o *GetSettings200ResponseDataInner) GetSamlIdentityProviderEnabled() bool`

GetSamlIdentityProviderEnabled returns the SamlIdentityProviderEnabled field if non-nil, zero value otherwise.

### GetSamlIdentityProviderEnabledOk

`func (o *GetSettings200ResponseDataInner) GetSamlIdentityProviderEnabledOk() (*bool, bool)`

GetSamlIdentityProviderEnabledOk returns a tuple with the SamlIdentityProviderEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlIdentityProviderEnabled

`func (o *GetSettings200ResponseDataInner) SetSamlIdentityProviderEnabled(v bool)`

SetSamlIdentityProviderEnabled sets SamlIdentityProviderEnabled field to given value.

### HasSamlIdentityProviderEnabled

`func (o *GetSettings200ResponseDataInner) HasSamlIdentityProviderEnabled() bool`

HasSamlIdentityProviderEnabled returns a boolean if a field has been set.

### GetNonSsoAuthenticatorCount

`func (o *GetSettings200ResponseDataInner) GetNonSsoAuthenticatorCount() int32`

GetNonSsoAuthenticatorCount returns the NonSsoAuthenticatorCount field if non-nil, zero value otherwise.

### GetNonSsoAuthenticatorCountOk

`func (o *GetSettings200ResponseDataInner) GetNonSsoAuthenticatorCountOk() (*int32, bool)`

GetNonSsoAuthenticatorCountOk returns a tuple with the NonSsoAuthenticatorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonSsoAuthenticatorCount

`func (o *GetSettings200ResponseDataInner) SetNonSsoAuthenticatorCount(v int32)`

SetNonSsoAuthenticatorCount sets NonSsoAuthenticatorCount field to given value.

### HasNonSsoAuthenticatorCount

`func (o *GetSettings200ResponseDataInner) HasNonSsoAuthenticatorCount() bool`

HasNonSsoAuthenticatorCount returns a boolean if a field has been set.

### GetNonSsoGroupCount

`func (o *GetSettings200ResponseDataInner) GetNonSsoGroupCount() int32`

GetNonSsoGroupCount returns the NonSsoGroupCount field if non-nil, zero value otherwise.

### GetNonSsoGroupCountOk

`func (o *GetSettings200ResponseDataInner) GetNonSsoGroupCountOk() (*int32, bool)`

GetNonSsoGroupCountOk returns a tuple with the NonSsoGroupCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonSsoGroupCount

`func (o *GetSettings200ResponseDataInner) SetNonSsoGroupCount(v int32)`

SetNonSsoGroupCount sets NonSsoGroupCount field to given value.

### HasNonSsoGroupCount

`func (o *GetSettings200ResponseDataInner) HasNonSsoGroupCount() bool`

HasNonSsoGroupCount returns a boolean if a field has been set.

### GetLocalAdminUserCount

`func (o *GetSettings200ResponseDataInner) GetLocalAdminUserCount() int32`

GetLocalAdminUserCount returns the LocalAdminUserCount field if non-nil, zero value otherwise.

### GetLocalAdminUserCountOk

`func (o *GetSettings200ResponseDataInner) GetLocalAdminUserCountOk() (*int32, bool)`

GetLocalAdminUserCountOk returns a tuple with the LocalAdminUserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAdminUserCount

`func (o *GetSettings200ResponseDataInner) SetLocalAdminUserCount(v int32)`

SetLocalAdminUserCount sets LocalAdminUserCount field to given value.

### HasLocalAdminUserCount

`func (o *GetSettings200ResponseDataInner) HasLocalAdminUserCount() bool`

HasLocalAdminUserCount returns a boolean if a field has been set.

### GetLocalNonAdminUserCount

`func (o *GetSettings200ResponseDataInner) GetLocalNonAdminUserCount() int32`

GetLocalNonAdminUserCount returns the LocalNonAdminUserCount field if non-nil, zero value otherwise.

### GetLocalNonAdminUserCountOk

`func (o *GetSettings200ResponseDataInner) GetLocalNonAdminUserCountOk() (*int32, bool)`

GetLocalNonAdminUserCountOk returns a tuple with the LocalNonAdminUserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalNonAdminUserCount

`func (o *GetSettings200ResponseDataInner) SetLocalNonAdminUserCount(v int32)`

SetLocalNonAdminUserCount sets LocalNonAdminUserCount field to given value.

### HasLocalNonAdminUserCount

`func (o *GetSettings200ResponseDataInner) HasLocalNonAdminUserCount() bool`

HasLocalNonAdminUserCount returns a boolean if a field has been set.

### GetHostname

`func (o *GetSettings200ResponseDataInner) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *GetSettings200ResponseDataInner) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *GetSettings200ResponseDataInner) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *GetSettings200ResponseDataInner) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetVersion

`func (o *GetSettings200ResponseDataInner) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetSettings200ResponseDataInner) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetSettings200ResponseDataInner) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *GetSettings200ResponseDataInner) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetAddress

`func (o *GetSettings200ResponseDataInner) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *GetSettings200ResponseDataInner) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *GetSettings200ResponseDataInner) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *GetSettings200ResponseDataInner) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetInterfaceRedundancyEnabled

`func (o *GetSettings200ResponseDataInner) GetInterfaceRedundancyEnabled() bool`

GetInterfaceRedundancyEnabled returns the InterfaceRedundancyEnabled field if non-nil, zero value otherwise.

### GetInterfaceRedundancyEnabledOk

`func (o *GetSettings200ResponseDataInner) GetInterfaceRedundancyEnabledOk() (*bool, bool)`

GetInterfaceRedundancyEnabledOk returns a tuple with the InterfaceRedundancyEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceRedundancyEnabled

`func (o *GetSettings200ResponseDataInner) SetInterfaceRedundancyEnabled(v bool)`

SetInterfaceRedundancyEnabled sets InterfaceRedundancyEnabled field to given value.

### HasInterfaceRedundancyEnabled

`func (o *GetSettings200ResponseDataInner) HasInterfaceRedundancyEnabled() bool`

HasInterfaceRedundancyEnabled returns a boolean if a field has been set.

### GetActiveSessions

`func (o *GetSettings200ResponseDataInner) GetActiveSessions() int32`

GetActiveSessions returns the ActiveSessions field if non-nil, zero value otherwise.

### GetActiveSessionsOk

`func (o *GetSettings200ResponseDataInner) GetActiveSessionsOk() (*int32, bool)`

GetActiveSessionsOk returns a tuple with the ActiveSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveSessions

`func (o *GetSettings200ResponseDataInner) SetActiveSessions(v int32)`

SetActiveSessions sets ActiveSessions field to given value.

### HasActiveSessions

`func (o *GetSettings200ResponseDataInner) HasActiveSessions() bool`

HasActiveSessions returns a boolean if a field has been set.

### GetLinks

`func (o *GetSettings200ResponseDataInner) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetSettings200ResponseDataInner) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetSettings200ResponseDataInner) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetSettings200ResponseDataInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


