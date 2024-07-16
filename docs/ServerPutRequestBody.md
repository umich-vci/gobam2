# ServerPutRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | **string** | The resource type. | 
**Name** | **string** | The name of the resource. | 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**Configuration** | Pointer to [**InlinedConfiguration**](InlinedConfiguration.md) |  | [optional] [readonly] 
**Profile** | **string** | The profile of the server. | 
**State** | Pointer to **string** | The current state of the server, indicating whether the server is enabled or disabled. | [optional] 
**Username** | Pointer to **string** | The username used to authenticate with the server. | [optional] 
**Password** | Pointer to **string** | The password used to authenticate with the server. | [optional] 
**ServerGroup** | Pointer to [**InlinedServerGroup**](InlinedServerGroup.md) |  | [optional] 
**Connected** | Pointer to **bool** | Indicates whether the server is connected to Address Manager. | [optional] 
**MonitoringEnabled** | Pointer to **bool** | Indicates whether monitoring service is enabled on the server. | [optional] 
**Snmp** | Pointer to [**SnmpBean**](SnmpBean.md) |  | [optional] 
**Interfaces** | Pointer to [**[]NetworkInterface**](NetworkInterface.md) | The list of network interfaces of the server. | [optional] 
**Location** | Pointer to [**InlinedLocation**](InlinedLocation.md) |  | [optional] 
**Statistics** | Pointer to **map[string]map[string]interface{}** | Displays statistics information collected from the monitoring service. | [optional] [readonly] 
**OverrideConfigurationLevelDhcpValidation** | **bool** | Indicates whether DHCP deployment validation settings configured at the configuration level are overridden at the server level. | 
**OverrideConfigurationLevelDnsValidation** | **bool** | Indicates whether DNS deployment validation settings configured at the configuration level are overridden at the server level. | 
**DhcpConfigurationValidationEnabled** | **bool** | Indicates whether the syntax of the dhcpd.conf file is validated prior to deployment from Address Manager. | 
**DnsConfigurationValidationEnabled** | **bool** | Indicates whether the syntax of the named.conf file is validated prior to deployment from Address Manager. | 
**DnsZoneValidationEnabled** | **bool** | Indicates whether the syntax of each DNS zone file is validated prior to deployment from Address Manager. | 
**CheckIntegrityValidation** | **string** | The method for which the syntax checks of the DNS zone file is checked. | 
**CheckMxCnameValidation** | **string** | Checks if MX records point to a CNAME record rather than an A or AAAA, and determines how Address Manager handles conditions found by the check. | 
**CheckMxValidation** | **string** | Checks if MX records point to an IP address rather than an A or AAAA, and determines how Address Manager handles conditions found by the check. | 
**CheckNamesValidation** | **string** | Checks the names within the DNS zone files and determines how Address Manager handles conditions found by the check. | 
**CheckNsValidation** | **string** | Checks if NS records point to an IP address rather than an A or AAAA, and determines how Address Manager handles conditions found by the check. | 
**CheckSrvCnameValidation** | **string** | Checks if SRV records point to a CNAME record rather than an A or AAAA, and determines how Address Manager handles conditions found by the check. | 
**CheckWildcardValidation** | **string** | Checks for wildcards in zone names that don&#39;t appear as the last segment of a zone name, and determines how Address Manager handles conditions found by the check. | 
**PrivateAddress** | Pointer to **string** | The private IP address of the server. | [optional] 
**EncryptedNotificationsEnabled** | Pointer to **bool** | Indicates whether notifications are encrypted between Address Manager and the DNS/DHCP Server | [optional] 
**ManagementUrl** | Pointer to **string** | Specifies the management URL for an F5 LTM or GTM server. | [optional] 
**SelfIpAddress** | Pointer to **string** | Specifies the self URL for an F5 GTM server. | [optional] 
**HaBackboneEnabled** | **bool** | Indicates whether a backbone is enabled between nodes of a high-availability pair. | 
**HaPingAddress** | Pointer to **string** | Sets the ping address of the high-availability pair. | [optional] 
**DhcpServicePrincipal** | Pointer to [**InlinedKerberosServicePrincipal**](InlinedKerberosServicePrincipal.md) |  | [optional] 
**DnsServicePrincipal** | Pointer to [**InlinedKerberosServicePrincipal**](InlinedKerberosServicePrincipal.md) |  | [optional] 
**DedicatedManagementEnabled** | **bool** | Indicates whether dedicated management is enabled on the server. | 
**HaRole** | Pointer to **string** |  | [optional] [readonly] 
**HaPeerConnectionState** | Pointer to **string** |  | [optional] [readonly] 
**HaDiskState** | Pointer to **string** |  | [optional] [readonly] 
**HsmSupportEnabled** | **bool** | Indicates whether HSM is enabled on the server. | 
**InterfaceRedundancyEnabled** | **bool** | Indicates whether interface redundancy is enabled on the server. | 
**InheritedFields** | Pointer to **[]string** |  | [optional] [readonly] 

## Methods

### NewServerPutRequestBody

`func NewServerPutRequestBody(type_ string, name string, profile string, overrideConfigurationLevelDhcpValidation bool, overrideConfigurationLevelDnsValidation bool, dhcpConfigurationValidationEnabled bool, dnsConfigurationValidationEnabled bool, dnsZoneValidationEnabled bool, checkIntegrityValidation string, checkMxCnameValidation string, checkMxValidation string, checkNamesValidation string, checkNsValidation string, checkSrvCnameValidation string, checkWildcardValidation string, haBackboneEnabled bool, dedicatedManagementEnabled bool, hsmSupportEnabled bool, interfaceRedundancyEnabled bool, ) *ServerPutRequestBody`

NewServerPutRequestBody instantiates a new ServerPutRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerPutRequestBodyWithDefaults

`func NewServerPutRequestBodyWithDefaults() *ServerPutRequestBody`

NewServerPutRequestBodyWithDefaults instantiates a new ServerPutRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServerPutRequestBody) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServerPutRequestBody) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServerPutRequestBody) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ServerPutRequestBody) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ServerPutRequestBody) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServerPutRequestBody) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServerPutRequestBody) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *ServerPutRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServerPutRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServerPutRequestBody) SetName(v string)`

SetName sets Name field to given value.


### GetUserDefinedFields

`func (o *ServerPutRequestBody) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *ServerPutRequestBody) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *ServerPutRequestBody) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *ServerPutRequestBody) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetConfiguration

`func (o *ServerPutRequestBody) GetConfiguration() InlinedConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *ServerPutRequestBody) GetConfigurationOk() (*InlinedConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *ServerPutRequestBody) SetConfiguration(v InlinedConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *ServerPutRequestBody) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetProfile

`func (o *ServerPutRequestBody) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *ServerPutRequestBody) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *ServerPutRequestBody) SetProfile(v string)`

SetProfile sets Profile field to given value.


### GetState

`func (o *ServerPutRequestBody) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ServerPutRequestBody) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ServerPutRequestBody) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ServerPutRequestBody) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUsername

`func (o *ServerPutRequestBody) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ServerPutRequestBody) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ServerPutRequestBody) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ServerPutRequestBody) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *ServerPutRequestBody) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ServerPutRequestBody) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ServerPutRequestBody) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ServerPutRequestBody) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetServerGroup

`func (o *ServerPutRequestBody) GetServerGroup() InlinedServerGroup`

GetServerGroup returns the ServerGroup field if non-nil, zero value otherwise.

### GetServerGroupOk

`func (o *ServerPutRequestBody) GetServerGroupOk() (*InlinedServerGroup, bool)`

GetServerGroupOk returns a tuple with the ServerGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerGroup

`func (o *ServerPutRequestBody) SetServerGroup(v InlinedServerGroup)`

SetServerGroup sets ServerGroup field to given value.

### HasServerGroup

`func (o *ServerPutRequestBody) HasServerGroup() bool`

HasServerGroup returns a boolean if a field has been set.

### GetConnected

`func (o *ServerPutRequestBody) GetConnected() bool`

GetConnected returns the Connected field if non-nil, zero value otherwise.

### GetConnectedOk

`func (o *ServerPutRequestBody) GetConnectedOk() (*bool, bool)`

GetConnectedOk returns a tuple with the Connected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnected

`func (o *ServerPutRequestBody) SetConnected(v bool)`

SetConnected sets Connected field to given value.

### HasConnected

`func (o *ServerPutRequestBody) HasConnected() bool`

HasConnected returns a boolean if a field has been set.

### GetMonitoringEnabled

`func (o *ServerPutRequestBody) GetMonitoringEnabled() bool`

GetMonitoringEnabled returns the MonitoringEnabled field if non-nil, zero value otherwise.

### GetMonitoringEnabledOk

`func (o *ServerPutRequestBody) GetMonitoringEnabledOk() (*bool, bool)`

GetMonitoringEnabledOk returns a tuple with the MonitoringEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringEnabled

`func (o *ServerPutRequestBody) SetMonitoringEnabled(v bool)`

SetMonitoringEnabled sets MonitoringEnabled field to given value.

### HasMonitoringEnabled

`func (o *ServerPutRequestBody) HasMonitoringEnabled() bool`

HasMonitoringEnabled returns a boolean if a field has been set.

### GetSnmp

`func (o *ServerPutRequestBody) GetSnmp() SnmpBean`

GetSnmp returns the Snmp field if non-nil, zero value otherwise.

### GetSnmpOk

`func (o *ServerPutRequestBody) GetSnmpOk() (*SnmpBean, bool)`

GetSnmpOk returns a tuple with the Snmp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmp

`func (o *ServerPutRequestBody) SetSnmp(v SnmpBean)`

SetSnmp sets Snmp field to given value.

### HasSnmp

`func (o *ServerPutRequestBody) HasSnmp() bool`

HasSnmp returns a boolean if a field has been set.

### GetInterfaces

`func (o *ServerPutRequestBody) GetInterfaces() []NetworkInterface`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *ServerPutRequestBody) GetInterfacesOk() (*[]NetworkInterface, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *ServerPutRequestBody) SetInterfaces(v []NetworkInterface)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *ServerPutRequestBody) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.

### GetLocation

`func (o *ServerPutRequestBody) GetLocation() InlinedLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ServerPutRequestBody) GetLocationOk() (*InlinedLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ServerPutRequestBody) SetLocation(v InlinedLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ServerPutRequestBody) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetStatistics

`func (o *ServerPutRequestBody) GetStatistics() map[string]map[string]interface{}`

GetStatistics returns the Statistics field if non-nil, zero value otherwise.

### GetStatisticsOk

`func (o *ServerPutRequestBody) GetStatisticsOk() (*map[string]map[string]interface{}, bool)`

GetStatisticsOk returns a tuple with the Statistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatistics

`func (o *ServerPutRequestBody) SetStatistics(v map[string]map[string]interface{})`

SetStatistics sets Statistics field to given value.

### HasStatistics

`func (o *ServerPutRequestBody) HasStatistics() bool`

HasStatistics returns a boolean if a field has been set.

### GetOverrideConfigurationLevelDhcpValidation

`func (o *ServerPutRequestBody) GetOverrideConfigurationLevelDhcpValidation() bool`

GetOverrideConfigurationLevelDhcpValidation returns the OverrideConfigurationLevelDhcpValidation field if non-nil, zero value otherwise.

### GetOverrideConfigurationLevelDhcpValidationOk

`func (o *ServerPutRequestBody) GetOverrideConfigurationLevelDhcpValidationOk() (*bool, bool)`

GetOverrideConfigurationLevelDhcpValidationOk returns a tuple with the OverrideConfigurationLevelDhcpValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideConfigurationLevelDhcpValidation

`func (o *ServerPutRequestBody) SetOverrideConfigurationLevelDhcpValidation(v bool)`

SetOverrideConfigurationLevelDhcpValidation sets OverrideConfigurationLevelDhcpValidation field to given value.


### GetOverrideConfigurationLevelDnsValidation

`func (o *ServerPutRequestBody) GetOverrideConfigurationLevelDnsValidation() bool`

GetOverrideConfigurationLevelDnsValidation returns the OverrideConfigurationLevelDnsValidation field if non-nil, zero value otherwise.

### GetOverrideConfigurationLevelDnsValidationOk

`func (o *ServerPutRequestBody) GetOverrideConfigurationLevelDnsValidationOk() (*bool, bool)`

GetOverrideConfigurationLevelDnsValidationOk returns a tuple with the OverrideConfigurationLevelDnsValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideConfigurationLevelDnsValidation

`func (o *ServerPutRequestBody) SetOverrideConfigurationLevelDnsValidation(v bool)`

SetOverrideConfigurationLevelDnsValidation sets OverrideConfigurationLevelDnsValidation field to given value.


### GetDhcpConfigurationValidationEnabled

`func (o *ServerPutRequestBody) GetDhcpConfigurationValidationEnabled() bool`

GetDhcpConfigurationValidationEnabled returns the DhcpConfigurationValidationEnabled field if non-nil, zero value otherwise.

### GetDhcpConfigurationValidationEnabledOk

`func (o *ServerPutRequestBody) GetDhcpConfigurationValidationEnabledOk() (*bool, bool)`

GetDhcpConfigurationValidationEnabledOk returns a tuple with the DhcpConfigurationValidationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpConfigurationValidationEnabled

`func (o *ServerPutRequestBody) SetDhcpConfigurationValidationEnabled(v bool)`

SetDhcpConfigurationValidationEnabled sets DhcpConfigurationValidationEnabled field to given value.


### GetDnsConfigurationValidationEnabled

`func (o *ServerPutRequestBody) GetDnsConfigurationValidationEnabled() bool`

GetDnsConfigurationValidationEnabled returns the DnsConfigurationValidationEnabled field if non-nil, zero value otherwise.

### GetDnsConfigurationValidationEnabledOk

`func (o *ServerPutRequestBody) GetDnsConfigurationValidationEnabledOk() (*bool, bool)`

GetDnsConfigurationValidationEnabledOk returns a tuple with the DnsConfigurationValidationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsConfigurationValidationEnabled

`func (o *ServerPutRequestBody) SetDnsConfigurationValidationEnabled(v bool)`

SetDnsConfigurationValidationEnabled sets DnsConfigurationValidationEnabled field to given value.


### GetDnsZoneValidationEnabled

`func (o *ServerPutRequestBody) GetDnsZoneValidationEnabled() bool`

GetDnsZoneValidationEnabled returns the DnsZoneValidationEnabled field if non-nil, zero value otherwise.

### GetDnsZoneValidationEnabledOk

`func (o *ServerPutRequestBody) GetDnsZoneValidationEnabledOk() (*bool, bool)`

GetDnsZoneValidationEnabledOk returns a tuple with the DnsZoneValidationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsZoneValidationEnabled

`func (o *ServerPutRequestBody) SetDnsZoneValidationEnabled(v bool)`

SetDnsZoneValidationEnabled sets DnsZoneValidationEnabled field to given value.


### GetCheckIntegrityValidation

`func (o *ServerPutRequestBody) GetCheckIntegrityValidation() string`

GetCheckIntegrityValidation returns the CheckIntegrityValidation field if non-nil, zero value otherwise.

### GetCheckIntegrityValidationOk

`func (o *ServerPutRequestBody) GetCheckIntegrityValidationOk() (*string, bool)`

GetCheckIntegrityValidationOk returns a tuple with the CheckIntegrityValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckIntegrityValidation

`func (o *ServerPutRequestBody) SetCheckIntegrityValidation(v string)`

SetCheckIntegrityValidation sets CheckIntegrityValidation field to given value.


### GetCheckMxCnameValidation

`func (o *ServerPutRequestBody) GetCheckMxCnameValidation() string`

GetCheckMxCnameValidation returns the CheckMxCnameValidation field if non-nil, zero value otherwise.

### GetCheckMxCnameValidationOk

`func (o *ServerPutRequestBody) GetCheckMxCnameValidationOk() (*string, bool)`

GetCheckMxCnameValidationOk returns a tuple with the CheckMxCnameValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckMxCnameValidation

`func (o *ServerPutRequestBody) SetCheckMxCnameValidation(v string)`

SetCheckMxCnameValidation sets CheckMxCnameValidation field to given value.


### GetCheckMxValidation

`func (o *ServerPutRequestBody) GetCheckMxValidation() string`

GetCheckMxValidation returns the CheckMxValidation field if non-nil, zero value otherwise.

### GetCheckMxValidationOk

`func (o *ServerPutRequestBody) GetCheckMxValidationOk() (*string, bool)`

GetCheckMxValidationOk returns a tuple with the CheckMxValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckMxValidation

`func (o *ServerPutRequestBody) SetCheckMxValidation(v string)`

SetCheckMxValidation sets CheckMxValidation field to given value.


### GetCheckNamesValidation

`func (o *ServerPutRequestBody) GetCheckNamesValidation() string`

GetCheckNamesValidation returns the CheckNamesValidation field if non-nil, zero value otherwise.

### GetCheckNamesValidationOk

`func (o *ServerPutRequestBody) GetCheckNamesValidationOk() (*string, bool)`

GetCheckNamesValidationOk returns a tuple with the CheckNamesValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckNamesValidation

`func (o *ServerPutRequestBody) SetCheckNamesValidation(v string)`

SetCheckNamesValidation sets CheckNamesValidation field to given value.


### GetCheckNsValidation

`func (o *ServerPutRequestBody) GetCheckNsValidation() string`

GetCheckNsValidation returns the CheckNsValidation field if non-nil, zero value otherwise.

### GetCheckNsValidationOk

`func (o *ServerPutRequestBody) GetCheckNsValidationOk() (*string, bool)`

GetCheckNsValidationOk returns a tuple with the CheckNsValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckNsValidation

`func (o *ServerPutRequestBody) SetCheckNsValidation(v string)`

SetCheckNsValidation sets CheckNsValidation field to given value.


### GetCheckSrvCnameValidation

`func (o *ServerPutRequestBody) GetCheckSrvCnameValidation() string`

GetCheckSrvCnameValidation returns the CheckSrvCnameValidation field if non-nil, zero value otherwise.

### GetCheckSrvCnameValidationOk

`func (o *ServerPutRequestBody) GetCheckSrvCnameValidationOk() (*string, bool)`

GetCheckSrvCnameValidationOk returns a tuple with the CheckSrvCnameValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckSrvCnameValidation

`func (o *ServerPutRequestBody) SetCheckSrvCnameValidation(v string)`

SetCheckSrvCnameValidation sets CheckSrvCnameValidation field to given value.


### GetCheckWildcardValidation

`func (o *ServerPutRequestBody) GetCheckWildcardValidation() string`

GetCheckWildcardValidation returns the CheckWildcardValidation field if non-nil, zero value otherwise.

### GetCheckWildcardValidationOk

`func (o *ServerPutRequestBody) GetCheckWildcardValidationOk() (*string, bool)`

GetCheckWildcardValidationOk returns a tuple with the CheckWildcardValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckWildcardValidation

`func (o *ServerPutRequestBody) SetCheckWildcardValidation(v string)`

SetCheckWildcardValidation sets CheckWildcardValidation field to given value.


### GetPrivateAddress

`func (o *ServerPutRequestBody) GetPrivateAddress() string`

GetPrivateAddress returns the PrivateAddress field if non-nil, zero value otherwise.

### GetPrivateAddressOk

`func (o *ServerPutRequestBody) GetPrivateAddressOk() (*string, bool)`

GetPrivateAddressOk returns a tuple with the PrivateAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateAddress

`func (o *ServerPutRequestBody) SetPrivateAddress(v string)`

SetPrivateAddress sets PrivateAddress field to given value.

### HasPrivateAddress

`func (o *ServerPutRequestBody) HasPrivateAddress() bool`

HasPrivateAddress returns a boolean if a field has been set.

### GetEncryptedNotificationsEnabled

`func (o *ServerPutRequestBody) GetEncryptedNotificationsEnabled() bool`

GetEncryptedNotificationsEnabled returns the EncryptedNotificationsEnabled field if non-nil, zero value otherwise.

### GetEncryptedNotificationsEnabledOk

`func (o *ServerPutRequestBody) GetEncryptedNotificationsEnabledOk() (*bool, bool)`

GetEncryptedNotificationsEnabledOk returns a tuple with the EncryptedNotificationsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedNotificationsEnabled

`func (o *ServerPutRequestBody) SetEncryptedNotificationsEnabled(v bool)`

SetEncryptedNotificationsEnabled sets EncryptedNotificationsEnabled field to given value.

### HasEncryptedNotificationsEnabled

`func (o *ServerPutRequestBody) HasEncryptedNotificationsEnabled() bool`

HasEncryptedNotificationsEnabled returns a boolean if a field has been set.

### GetManagementUrl

`func (o *ServerPutRequestBody) GetManagementUrl() string`

GetManagementUrl returns the ManagementUrl field if non-nil, zero value otherwise.

### GetManagementUrlOk

`func (o *ServerPutRequestBody) GetManagementUrlOk() (*string, bool)`

GetManagementUrlOk returns a tuple with the ManagementUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementUrl

`func (o *ServerPutRequestBody) SetManagementUrl(v string)`

SetManagementUrl sets ManagementUrl field to given value.

### HasManagementUrl

`func (o *ServerPutRequestBody) HasManagementUrl() bool`

HasManagementUrl returns a boolean if a field has been set.

### GetSelfIpAddress

`func (o *ServerPutRequestBody) GetSelfIpAddress() string`

GetSelfIpAddress returns the SelfIpAddress field if non-nil, zero value otherwise.

### GetSelfIpAddressOk

`func (o *ServerPutRequestBody) GetSelfIpAddressOk() (*string, bool)`

GetSelfIpAddressOk returns a tuple with the SelfIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfIpAddress

`func (o *ServerPutRequestBody) SetSelfIpAddress(v string)`

SetSelfIpAddress sets SelfIpAddress field to given value.

### HasSelfIpAddress

`func (o *ServerPutRequestBody) HasSelfIpAddress() bool`

HasSelfIpAddress returns a boolean if a field has been set.

### GetHaBackboneEnabled

`func (o *ServerPutRequestBody) GetHaBackboneEnabled() bool`

GetHaBackboneEnabled returns the HaBackboneEnabled field if non-nil, zero value otherwise.

### GetHaBackboneEnabledOk

`func (o *ServerPutRequestBody) GetHaBackboneEnabledOk() (*bool, bool)`

GetHaBackboneEnabledOk returns a tuple with the HaBackboneEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaBackboneEnabled

`func (o *ServerPutRequestBody) SetHaBackboneEnabled(v bool)`

SetHaBackboneEnabled sets HaBackboneEnabled field to given value.


### GetHaPingAddress

`func (o *ServerPutRequestBody) GetHaPingAddress() string`

GetHaPingAddress returns the HaPingAddress field if non-nil, zero value otherwise.

### GetHaPingAddressOk

`func (o *ServerPutRequestBody) GetHaPingAddressOk() (*string, bool)`

GetHaPingAddressOk returns a tuple with the HaPingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaPingAddress

`func (o *ServerPutRequestBody) SetHaPingAddress(v string)`

SetHaPingAddress sets HaPingAddress field to given value.

### HasHaPingAddress

`func (o *ServerPutRequestBody) HasHaPingAddress() bool`

HasHaPingAddress returns a boolean if a field has been set.

### GetDhcpServicePrincipal

`func (o *ServerPutRequestBody) GetDhcpServicePrincipal() InlinedKerberosServicePrincipal`

GetDhcpServicePrincipal returns the DhcpServicePrincipal field if non-nil, zero value otherwise.

### GetDhcpServicePrincipalOk

`func (o *ServerPutRequestBody) GetDhcpServicePrincipalOk() (*InlinedKerberosServicePrincipal, bool)`

GetDhcpServicePrincipalOk returns a tuple with the DhcpServicePrincipal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpServicePrincipal

`func (o *ServerPutRequestBody) SetDhcpServicePrincipal(v InlinedKerberosServicePrincipal)`

SetDhcpServicePrincipal sets DhcpServicePrincipal field to given value.

### HasDhcpServicePrincipal

`func (o *ServerPutRequestBody) HasDhcpServicePrincipal() bool`

HasDhcpServicePrincipal returns a boolean if a field has been set.

### GetDnsServicePrincipal

`func (o *ServerPutRequestBody) GetDnsServicePrincipal() InlinedKerberosServicePrincipal`

GetDnsServicePrincipal returns the DnsServicePrincipal field if non-nil, zero value otherwise.

### GetDnsServicePrincipalOk

`func (o *ServerPutRequestBody) GetDnsServicePrincipalOk() (*InlinedKerberosServicePrincipal, bool)`

GetDnsServicePrincipalOk returns a tuple with the DnsServicePrincipal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServicePrincipal

`func (o *ServerPutRequestBody) SetDnsServicePrincipal(v InlinedKerberosServicePrincipal)`

SetDnsServicePrincipal sets DnsServicePrincipal field to given value.

### HasDnsServicePrincipal

`func (o *ServerPutRequestBody) HasDnsServicePrincipal() bool`

HasDnsServicePrincipal returns a boolean if a field has been set.

### GetDedicatedManagementEnabled

`func (o *ServerPutRequestBody) GetDedicatedManagementEnabled() bool`

GetDedicatedManagementEnabled returns the DedicatedManagementEnabled field if non-nil, zero value otherwise.

### GetDedicatedManagementEnabledOk

`func (o *ServerPutRequestBody) GetDedicatedManagementEnabledOk() (*bool, bool)`

GetDedicatedManagementEnabledOk returns a tuple with the DedicatedManagementEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedicatedManagementEnabled

`func (o *ServerPutRequestBody) SetDedicatedManagementEnabled(v bool)`

SetDedicatedManagementEnabled sets DedicatedManagementEnabled field to given value.


### GetHaRole

`func (o *ServerPutRequestBody) GetHaRole() string`

GetHaRole returns the HaRole field if non-nil, zero value otherwise.

### GetHaRoleOk

`func (o *ServerPutRequestBody) GetHaRoleOk() (*string, bool)`

GetHaRoleOk returns a tuple with the HaRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaRole

`func (o *ServerPutRequestBody) SetHaRole(v string)`

SetHaRole sets HaRole field to given value.

### HasHaRole

`func (o *ServerPutRequestBody) HasHaRole() bool`

HasHaRole returns a boolean if a field has been set.

### GetHaPeerConnectionState

`func (o *ServerPutRequestBody) GetHaPeerConnectionState() string`

GetHaPeerConnectionState returns the HaPeerConnectionState field if non-nil, zero value otherwise.

### GetHaPeerConnectionStateOk

`func (o *ServerPutRequestBody) GetHaPeerConnectionStateOk() (*string, bool)`

GetHaPeerConnectionStateOk returns a tuple with the HaPeerConnectionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaPeerConnectionState

`func (o *ServerPutRequestBody) SetHaPeerConnectionState(v string)`

SetHaPeerConnectionState sets HaPeerConnectionState field to given value.

### HasHaPeerConnectionState

`func (o *ServerPutRequestBody) HasHaPeerConnectionState() bool`

HasHaPeerConnectionState returns a boolean if a field has been set.

### GetHaDiskState

`func (o *ServerPutRequestBody) GetHaDiskState() string`

GetHaDiskState returns the HaDiskState field if non-nil, zero value otherwise.

### GetHaDiskStateOk

`func (o *ServerPutRequestBody) GetHaDiskStateOk() (*string, bool)`

GetHaDiskStateOk returns a tuple with the HaDiskState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaDiskState

`func (o *ServerPutRequestBody) SetHaDiskState(v string)`

SetHaDiskState sets HaDiskState field to given value.

### HasHaDiskState

`func (o *ServerPutRequestBody) HasHaDiskState() bool`

HasHaDiskState returns a boolean if a field has been set.

### GetHsmSupportEnabled

`func (o *ServerPutRequestBody) GetHsmSupportEnabled() bool`

GetHsmSupportEnabled returns the HsmSupportEnabled field if non-nil, zero value otherwise.

### GetHsmSupportEnabledOk

`func (o *ServerPutRequestBody) GetHsmSupportEnabledOk() (*bool, bool)`

GetHsmSupportEnabledOk returns a tuple with the HsmSupportEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHsmSupportEnabled

`func (o *ServerPutRequestBody) SetHsmSupportEnabled(v bool)`

SetHsmSupportEnabled sets HsmSupportEnabled field to given value.


### GetInterfaceRedundancyEnabled

`func (o *ServerPutRequestBody) GetInterfaceRedundancyEnabled() bool`

GetInterfaceRedundancyEnabled returns the InterfaceRedundancyEnabled field if non-nil, zero value otherwise.

### GetInterfaceRedundancyEnabledOk

`func (o *ServerPutRequestBody) GetInterfaceRedundancyEnabledOk() (*bool, bool)`

GetInterfaceRedundancyEnabledOk returns a tuple with the InterfaceRedundancyEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceRedundancyEnabled

`func (o *ServerPutRequestBody) SetInterfaceRedundancyEnabled(v bool)`

SetInterfaceRedundancyEnabled sets InterfaceRedundancyEnabled field to given value.


### GetInheritedFields

`func (o *ServerPutRequestBody) GetInheritedFields() []string`

GetInheritedFields returns the InheritedFields field if non-nil, zero value otherwise.

### GetInheritedFieldsOk

`func (o *ServerPutRequestBody) GetInheritedFieldsOk() (*[]string, bool)`

GetInheritedFieldsOk returns a tuple with the InheritedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritedFields

`func (o *ServerPutRequestBody) SetInheritedFields(v []string)`

SetInheritedFields sets InheritedFields field to given value.

### HasInheritedFields

`func (o *ServerPutRequestBody) HasInheritedFields() bool`

HasInheritedFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


