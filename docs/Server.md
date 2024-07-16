# Server

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | Pointer to **string** | The name of the resource. | [optional] 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**Configuration** | Pointer to [**InlinedConfiguration**](InlinedConfiguration.md) |  | [optional] [readonly] 
**Profile** | Pointer to **string** | The profile of the server. | [optional] 
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
**OverrideConfigurationLevelDhcpValidation** | Pointer to **bool** | Indicates whether DHCP deployment validation settings configured at the configuration level are overridden at the server level. | [optional] 
**OverrideConfigurationLevelDnsValidation** | Pointer to **bool** | Indicates whether DNS deployment validation settings configured at the configuration level are overridden at the server level. | [optional] 
**DhcpConfigurationValidationEnabled** | Pointer to **bool** | Indicates whether the syntax of the dhcpd.conf file is validated prior to deployment from Address Manager. | [optional] 
**DnsConfigurationValidationEnabled** | Pointer to **bool** | Indicates whether the syntax of the named.conf file is validated prior to deployment from Address Manager. | [optional] 
**DnsZoneValidationEnabled** | Pointer to **bool** | Indicates whether the syntax of each DNS zone file is validated prior to deployment from Address Manager. | [optional] 
**CheckIntegrityValidation** | Pointer to **string** | The method for which the syntax checks of the DNS zone file is checked. | [optional] 
**CheckMxCnameValidation** | Pointer to **string** | Checks if MX records point to a CNAME record rather than an A or AAAA, and determines how Address Manager handles conditions found by the check. | [optional] 
**CheckMxValidation** | Pointer to **string** | Checks if MX records point to an IP address rather than an A or AAAA, and determines how Address Manager handles conditions found by the check. | [optional] 
**CheckNamesValidation** | Pointer to **string** | Checks the names within the DNS zone files and determines how Address Manager handles conditions found by the check. | [optional] 
**CheckNsValidation** | Pointer to **string** | Checks if NS records point to an IP address rather than an A or AAAA, and determines how Address Manager handles conditions found by the check. | [optional] 
**CheckSrvCnameValidation** | Pointer to **string** | Checks if SRV records point to a CNAME record rather than an A or AAAA, and determines how Address Manager handles conditions found by the check. | [optional] 
**CheckWildcardValidation** | Pointer to **string** | Checks for wildcards in zone names that don&#39;t appear as the last segment of a zone name, and determines how Address Manager handles conditions found by the check. | [optional] 
**PrivateAddress** | Pointer to **string** | The private IP address of the server. | [optional] 
**EncryptedNotificationsEnabled** | Pointer to **bool** | Indicates whether notifications are encrypted between Address Manager and the DNS/DHCP Server | [optional] 
**ManagementUrl** | Pointer to **string** | Specifies the management URL for an F5 LTM or GTM server. | [optional] 
**SelfIpAddress** | Pointer to **string** | Specifies the self URL for an F5 GTM server. | [optional] 
**HaBackboneEnabled** | Pointer to **bool** | Indicates whether a backbone is enabled between nodes of a high-availability pair. | [optional] 
**HaPingAddress** | Pointer to **string** | Sets the ping address of the high-availability pair. | [optional] 
**DhcpServicePrincipal** | Pointer to [**InlinedKerberosServicePrincipal**](InlinedKerberosServicePrincipal.md) |  | [optional] 
**DnsServicePrincipal** | Pointer to [**InlinedKerberosServicePrincipal**](InlinedKerberosServicePrincipal.md) |  | [optional] 
**DedicatedManagementEnabled** | Pointer to **bool** | Indicates whether dedicated management is enabled on the server. | [optional] 
**HaRole** | Pointer to **string** |  | [optional] [readonly] 
**HaPeerConnectionState** | Pointer to **string** |  | [optional] [readonly] 
**HaDiskState** | Pointer to **string** |  | [optional] [readonly] 
**HsmSupportEnabled** | Pointer to **bool** | Indicates whether HSM is enabled on the server. | [optional] 
**InterfaceRedundancyEnabled** | Pointer to **bool** | Indicates whether interface redundancy is enabled on the server. | [optional] 
**InheritedFields** | Pointer to **[]string** |  | [optional] [readonly] 

## Methods

### NewServer

`func NewServer() *Server`

NewServer instantiates a new Server object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerWithDefaults

`func NewServerWithDefaults() *Server`

NewServerWithDefaults instantiates a new Server object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Server) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Server) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Server) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Server) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *Server) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Server) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Server) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Server) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *Server) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Server) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Server) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Server) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *Server) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *Server) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *Server) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *Server) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetConfiguration

`func (o *Server) GetConfiguration() InlinedConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *Server) GetConfigurationOk() (*InlinedConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *Server) SetConfiguration(v InlinedConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *Server) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetProfile

`func (o *Server) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *Server) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *Server) SetProfile(v string)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *Server) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetState

`func (o *Server) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Server) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Server) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Server) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUsername

`func (o *Server) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *Server) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *Server) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *Server) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *Server) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *Server) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *Server) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *Server) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetServerGroup

`func (o *Server) GetServerGroup() InlinedServerGroup`

GetServerGroup returns the ServerGroup field if non-nil, zero value otherwise.

### GetServerGroupOk

`func (o *Server) GetServerGroupOk() (*InlinedServerGroup, bool)`

GetServerGroupOk returns a tuple with the ServerGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerGroup

`func (o *Server) SetServerGroup(v InlinedServerGroup)`

SetServerGroup sets ServerGroup field to given value.

### HasServerGroup

`func (o *Server) HasServerGroup() bool`

HasServerGroup returns a boolean if a field has been set.

### GetConnected

`func (o *Server) GetConnected() bool`

GetConnected returns the Connected field if non-nil, zero value otherwise.

### GetConnectedOk

`func (o *Server) GetConnectedOk() (*bool, bool)`

GetConnectedOk returns a tuple with the Connected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnected

`func (o *Server) SetConnected(v bool)`

SetConnected sets Connected field to given value.

### HasConnected

`func (o *Server) HasConnected() bool`

HasConnected returns a boolean if a field has been set.

### GetMonitoringEnabled

`func (o *Server) GetMonitoringEnabled() bool`

GetMonitoringEnabled returns the MonitoringEnabled field if non-nil, zero value otherwise.

### GetMonitoringEnabledOk

`func (o *Server) GetMonitoringEnabledOk() (*bool, bool)`

GetMonitoringEnabledOk returns a tuple with the MonitoringEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringEnabled

`func (o *Server) SetMonitoringEnabled(v bool)`

SetMonitoringEnabled sets MonitoringEnabled field to given value.

### HasMonitoringEnabled

`func (o *Server) HasMonitoringEnabled() bool`

HasMonitoringEnabled returns a boolean if a field has been set.

### GetSnmp

`func (o *Server) GetSnmp() SnmpBean`

GetSnmp returns the Snmp field if non-nil, zero value otherwise.

### GetSnmpOk

`func (o *Server) GetSnmpOk() (*SnmpBean, bool)`

GetSnmpOk returns a tuple with the Snmp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmp

`func (o *Server) SetSnmp(v SnmpBean)`

SetSnmp sets Snmp field to given value.

### HasSnmp

`func (o *Server) HasSnmp() bool`

HasSnmp returns a boolean if a field has been set.

### GetInterfaces

`func (o *Server) GetInterfaces() []NetworkInterface`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *Server) GetInterfacesOk() (*[]NetworkInterface, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *Server) SetInterfaces(v []NetworkInterface)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *Server) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.

### GetLocation

`func (o *Server) GetLocation() InlinedLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Server) GetLocationOk() (*InlinedLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Server) SetLocation(v InlinedLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *Server) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetStatistics

`func (o *Server) GetStatistics() map[string]map[string]interface{}`

GetStatistics returns the Statistics field if non-nil, zero value otherwise.

### GetStatisticsOk

`func (o *Server) GetStatisticsOk() (*map[string]map[string]interface{}, bool)`

GetStatisticsOk returns a tuple with the Statistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatistics

`func (o *Server) SetStatistics(v map[string]map[string]interface{})`

SetStatistics sets Statistics field to given value.

### HasStatistics

`func (o *Server) HasStatistics() bool`

HasStatistics returns a boolean if a field has been set.

### GetOverrideConfigurationLevelDhcpValidation

`func (o *Server) GetOverrideConfigurationLevelDhcpValidation() bool`

GetOverrideConfigurationLevelDhcpValidation returns the OverrideConfigurationLevelDhcpValidation field if non-nil, zero value otherwise.

### GetOverrideConfigurationLevelDhcpValidationOk

`func (o *Server) GetOverrideConfigurationLevelDhcpValidationOk() (*bool, bool)`

GetOverrideConfigurationLevelDhcpValidationOk returns a tuple with the OverrideConfigurationLevelDhcpValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideConfigurationLevelDhcpValidation

`func (o *Server) SetOverrideConfigurationLevelDhcpValidation(v bool)`

SetOverrideConfigurationLevelDhcpValidation sets OverrideConfigurationLevelDhcpValidation field to given value.

### HasOverrideConfigurationLevelDhcpValidation

`func (o *Server) HasOverrideConfigurationLevelDhcpValidation() bool`

HasOverrideConfigurationLevelDhcpValidation returns a boolean if a field has been set.

### GetOverrideConfigurationLevelDnsValidation

`func (o *Server) GetOverrideConfigurationLevelDnsValidation() bool`

GetOverrideConfigurationLevelDnsValidation returns the OverrideConfigurationLevelDnsValidation field if non-nil, zero value otherwise.

### GetOverrideConfigurationLevelDnsValidationOk

`func (o *Server) GetOverrideConfigurationLevelDnsValidationOk() (*bool, bool)`

GetOverrideConfigurationLevelDnsValidationOk returns a tuple with the OverrideConfigurationLevelDnsValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideConfigurationLevelDnsValidation

`func (o *Server) SetOverrideConfigurationLevelDnsValidation(v bool)`

SetOverrideConfigurationLevelDnsValidation sets OverrideConfigurationLevelDnsValidation field to given value.

### HasOverrideConfigurationLevelDnsValidation

`func (o *Server) HasOverrideConfigurationLevelDnsValidation() bool`

HasOverrideConfigurationLevelDnsValidation returns a boolean if a field has been set.

### GetDhcpConfigurationValidationEnabled

`func (o *Server) GetDhcpConfigurationValidationEnabled() bool`

GetDhcpConfigurationValidationEnabled returns the DhcpConfigurationValidationEnabled field if non-nil, zero value otherwise.

### GetDhcpConfigurationValidationEnabledOk

`func (o *Server) GetDhcpConfigurationValidationEnabledOk() (*bool, bool)`

GetDhcpConfigurationValidationEnabledOk returns a tuple with the DhcpConfigurationValidationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpConfigurationValidationEnabled

`func (o *Server) SetDhcpConfigurationValidationEnabled(v bool)`

SetDhcpConfigurationValidationEnabled sets DhcpConfigurationValidationEnabled field to given value.

### HasDhcpConfigurationValidationEnabled

`func (o *Server) HasDhcpConfigurationValidationEnabled() bool`

HasDhcpConfigurationValidationEnabled returns a boolean if a field has been set.

### GetDnsConfigurationValidationEnabled

`func (o *Server) GetDnsConfigurationValidationEnabled() bool`

GetDnsConfigurationValidationEnabled returns the DnsConfigurationValidationEnabled field if non-nil, zero value otherwise.

### GetDnsConfigurationValidationEnabledOk

`func (o *Server) GetDnsConfigurationValidationEnabledOk() (*bool, bool)`

GetDnsConfigurationValidationEnabledOk returns a tuple with the DnsConfigurationValidationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsConfigurationValidationEnabled

`func (o *Server) SetDnsConfigurationValidationEnabled(v bool)`

SetDnsConfigurationValidationEnabled sets DnsConfigurationValidationEnabled field to given value.

### HasDnsConfigurationValidationEnabled

`func (o *Server) HasDnsConfigurationValidationEnabled() bool`

HasDnsConfigurationValidationEnabled returns a boolean if a field has been set.

### GetDnsZoneValidationEnabled

`func (o *Server) GetDnsZoneValidationEnabled() bool`

GetDnsZoneValidationEnabled returns the DnsZoneValidationEnabled field if non-nil, zero value otherwise.

### GetDnsZoneValidationEnabledOk

`func (o *Server) GetDnsZoneValidationEnabledOk() (*bool, bool)`

GetDnsZoneValidationEnabledOk returns a tuple with the DnsZoneValidationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsZoneValidationEnabled

`func (o *Server) SetDnsZoneValidationEnabled(v bool)`

SetDnsZoneValidationEnabled sets DnsZoneValidationEnabled field to given value.

### HasDnsZoneValidationEnabled

`func (o *Server) HasDnsZoneValidationEnabled() bool`

HasDnsZoneValidationEnabled returns a boolean if a field has been set.

### GetCheckIntegrityValidation

`func (o *Server) GetCheckIntegrityValidation() string`

GetCheckIntegrityValidation returns the CheckIntegrityValidation field if non-nil, zero value otherwise.

### GetCheckIntegrityValidationOk

`func (o *Server) GetCheckIntegrityValidationOk() (*string, bool)`

GetCheckIntegrityValidationOk returns a tuple with the CheckIntegrityValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckIntegrityValidation

`func (o *Server) SetCheckIntegrityValidation(v string)`

SetCheckIntegrityValidation sets CheckIntegrityValidation field to given value.

### HasCheckIntegrityValidation

`func (o *Server) HasCheckIntegrityValidation() bool`

HasCheckIntegrityValidation returns a boolean if a field has been set.

### GetCheckMxCnameValidation

`func (o *Server) GetCheckMxCnameValidation() string`

GetCheckMxCnameValidation returns the CheckMxCnameValidation field if non-nil, zero value otherwise.

### GetCheckMxCnameValidationOk

`func (o *Server) GetCheckMxCnameValidationOk() (*string, bool)`

GetCheckMxCnameValidationOk returns a tuple with the CheckMxCnameValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckMxCnameValidation

`func (o *Server) SetCheckMxCnameValidation(v string)`

SetCheckMxCnameValidation sets CheckMxCnameValidation field to given value.

### HasCheckMxCnameValidation

`func (o *Server) HasCheckMxCnameValidation() bool`

HasCheckMxCnameValidation returns a boolean if a field has been set.

### GetCheckMxValidation

`func (o *Server) GetCheckMxValidation() string`

GetCheckMxValidation returns the CheckMxValidation field if non-nil, zero value otherwise.

### GetCheckMxValidationOk

`func (o *Server) GetCheckMxValidationOk() (*string, bool)`

GetCheckMxValidationOk returns a tuple with the CheckMxValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckMxValidation

`func (o *Server) SetCheckMxValidation(v string)`

SetCheckMxValidation sets CheckMxValidation field to given value.

### HasCheckMxValidation

`func (o *Server) HasCheckMxValidation() bool`

HasCheckMxValidation returns a boolean if a field has been set.

### GetCheckNamesValidation

`func (o *Server) GetCheckNamesValidation() string`

GetCheckNamesValidation returns the CheckNamesValidation field if non-nil, zero value otherwise.

### GetCheckNamesValidationOk

`func (o *Server) GetCheckNamesValidationOk() (*string, bool)`

GetCheckNamesValidationOk returns a tuple with the CheckNamesValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckNamesValidation

`func (o *Server) SetCheckNamesValidation(v string)`

SetCheckNamesValidation sets CheckNamesValidation field to given value.

### HasCheckNamesValidation

`func (o *Server) HasCheckNamesValidation() bool`

HasCheckNamesValidation returns a boolean if a field has been set.

### GetCheckNsValidation

`func (o *Server) GetCheckNsValidation() string`

GetCheckNsValidation returns the CheckNsValidation field if non-nil, zero value otherwise.

### GetCheckNsValidationOk

`func (o *Server) GetCheckNsValidationOk() (*string, bool)`

GetCheckNsValidationOk returns a tuple with the CheckNsValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckNsValidation

`func (o *Server) SetCheckNsValidation(v string)`

SetCheckNsValidation sets CheckNsValidation field to given value.

### HasCheckNsValidation

`func (o *Server) HasCheckNsValidation() bool`

HasCheckNsValidation returns a boolean if a field has been set.

### GetCheckSrvCnameValidation

`func (o *Server) GetCheckSrvCnameValidation() string`

GetCheckSrvCnameValidation returns the CheckSrvCnameValidation field if non-nil, zero value otherwise.

### GetCheckSrvCnameValidationOk

`func (o *Server) GetCheckSrvCnameValidationOk() (*string, bool)`

GetCheckSrvCnameValidationOk returns a tuple with the CheckSrvCnameValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckSrvCnameValidation

`func (o *Server) SetCheckSrvCnameValidation(v string)`

SetCheckSrvCnameValidation sets CheckSrvCnameValidation field to given value.

### HasCheckSrvCnameValidation

`func (o *Server) HasCheckSrvCnameValidation() bool`

HasCheckSrvCnameValidation returns a boolean if a field has been set.

### GetCheckWildcardValidation

`func (o *Server) GetCheckWildcardValidation() string`

GetCheckWildcardValidation returns the CheckWildcardValidation field if non-nil, zero value otherwise.

### GetCheckWildcardValidationOk

`func (o *Server) GetCheckWildcardValidationOk() (*string, bool)`

GetCheckWildcardValidationOk returns a tuple with the CheckWildcardValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckWildcardValidation

`func (o *Server) SetCheckWildcardValidation(v string)`

SetCheckWildcardValidation sets CheckWildcardValidation field to given value.

### HasCheckWildcardValidation

`func (o *Server) HasCheckWildcardValidation() bool`

HasCheckWildcardValidation returns a boolean if a field has been set.

### GetPrivateAddress

`func (o *Server) GetPrivateAddress() string`

GetPrivateAddress returns the PrivateAddress field if non-nil, zero value otherwise.

### GetPrivateAddressOk

`func (o *Server) GetPrivateAddressOk() (*string, bool)`

GetPrivateAddressOk returns a tuple with the PrivateAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateAddress

`func (o *Server) SetPrivateAddress(v string)`

SetPrivateAddress sets PrivateAddress field to given value.

### HasPrivateAddress

`func (o *Server) HasPrivateAddress() bool`

HasPrivateAddress returns a boolean if a field has been set.

### GetEncryptedNotificationsEnabled

`func (o *Server) GetEncryptedNotificationsEnabled() bool`

GetEncryptedNotificationsEnabled returns the EncryptedNotificationsEnabled field if non-nil, zero value otherwise.

### GetEncryptedNotificationsEnabledOk

`func (o *Server) GetEncryptedNotificationsEnabledOk() (*bool, bool)`

GetEncryptedNotificationsEnabledOk returns a tuple with the EncryptedNotificationsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedNotificationsEnabled

`func (o *Server) SetEncryptedNotificationsEnabled(v bool)`

SetEncryptedNotificationsEnabled sets EncryptedNotificationsEnabled field to given value.

### HasEncryptedNotificationsEnabled

`func (o *Server) HasEncryptedNotificationsEnabled() bool`

HasEncryptedNotificationsEnabled returns a boolean if a field has been set.

### GetManagementUrl

`func (o *Server) GetManagementUrl() string`

GetManagementUrl returns the ManagementUrl field if non-nil, zero value otherwise.

### GetManagementUrlOk

`func (o *Server) GetManagementUrlOk() (*string, bool)`

GetManagementUrlOk returns a tuple with the ManagementUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementUrl

`func (o *Server) SetManagementUrl(v string)`

SetManagementUrl sets ManagementUrl field to given value.

### HasManagementUrl

`func (o *Server) HasManagementUrl() bool`

HasManagementUrl returns a boolean if a field has been set.

### GetSelfIpAddress

`func (o *Server) GetSelfIpAddress() string`

GetSelfIpAddress returns the SelfIpAddress field if non-nil, zero value otherwise.

### GetSelfIpAddressOk

`func (o *Server) GetSelfIpAddressOk() (*string, bool)`

GetSelfIpAddressOk returns a tuple with the SelfIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfIpAddress

`func (o *Server) SetSelfIpAddress(v string)`

SetSelfIpAddress sets SelfIpAddress field to given value.

### HasSelfIpAddress

`func (o *Server) HasSelfIpAddress() bool`

HasSelfIpAddress returns a boolean if a field has been set.

### GetHaBackboneEnabled

`func (o *Server) GetHaBackboneEnabled() bool`

GetHaBackboneEnabled returns the HaBackboneEnabled field if non-nil, zero value otherwise.

### GetHaBackboneEnabledOk

`func (o *Server) GetHaBackboneEnabledOk() (*bool, bool)`

GetHaBackboneEnabledOk returns a tuple with the HaBackboneEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaBackboneEnabled

`func (o *Server) SetHaBackboneEnabled(v bool)`

SetHaBackboneEnabled sets HaBackboneEnabled field to given value.

### HasHaBackboneEnabled

`func (o *Server) HasHaBackboneEnabled() bool`

HasHaBackboneEnabled returns a boolean if a field has been set.

### GetHaPingAddress

`func (o *Server) GetHaPingAddress() string`

GetHaPingAddress returns the HaPingAddress field if non-nil, zero value otherwise.

### GetHaPingAddressOk

`func (o *Server) GetHaPingAddressOk() (*string, bool)`

GetHaPingAddressOk returns a tuple with the HaPingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaPingAddress

`func (o *Server) SetHaPingAddress(v string)`

SetHaPingAddress sets HaPingAddress field to given value.

### HasHaPingAddress

`func (o *Server) HasHaPingAddress() bool`

HasHaPingAddress returns a boolean if a field has been set.

### GetDhcpServicePrincipal

`func (o *Server) GetDhcpServicePrincipal() InlinedKerberosServicePrincipal`

GetDhcpServicePrincipal returns the DhcpServicePrincipal field if non-nil, zero value otherwise.

### GetDhcpServicePrincipalOk

`func (o *Server) GetDhcpServicePrincipalOk() (*InlinedKerberosServicePrincipal, bool)`

GetDhcpServicePrincipalOk returns a tuple with the DhcpServicePrincipal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpServicePrincipal

`func (o *Server) SetDhcpServicePrincipal(v InlinedKerberosServicePrincipal)`

SetDhcpServicePrincipal sets DhcpServicePrincipal field to given value.

### HasDhcpServicePrincipal

`func (o *Server) HasDhcpServicePrincipal() bool`

HasDhcpServicePrincipal returns a boolean if a field has been set.

### GetDnsServicePrincipal

`func (o *Server) GetDnsServicePrincipal() InlinedKerberosServicePrincipal`

GetDnsServicePrincipal returns the DnsServicePrincipal field if non-nil, zero value otherwise.

### GetDnsServicePrincipalOk

`func (o *Server) GetDnsServicePrincipalOk() (*InlinedKerberosServicePrincipal, bool)`

GetDnsServicePrincipalOk returns a tuple with the DnsServicePrincipal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServicePrincipal

`func (o *Server) SetDnsServicePrincipal(v InlinedKerberosServicePrincipal)`

SetDnsServicePrincipal sets DnsServicePrincipal field to given value.

### HasDnsServicePrincipal

`func (o *Server) HasDnsServicePrincipal() bool`

HasDnsServicePrincipal returns a boolean if a field has been set.

### GetDedicatedManagementEnabled

`func (o *Server) GetDedicatedManagementEnabled() bool`

GetDedicatedManagementEnabled returns the DedicatedManagementEnabled field if non-nil, zero value otherwise.

### GetDedicatedManagementEnabledOk

`func (o *Server) GetDedicatedManagementEnabledOk() (*bool, bool)`

GetDedicatedManagementEnabledOk returns a tuple with the DedicatedManagementEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedicatedManagementEnabled

`func (o *Server) SetDedicatedManagementEnabled(v bool)`

SetDedicatedManagementEnabled sets DedicatedManagementEnabled field to given value.

### HasDedicatedManagementEnabled

`func (o *Server) HasDedicatedManagementEnabled() bool`

HasDedicatedManagementEnabled returns a boolean if a field has been set.

### GetHaRole

`func (o *Server) GetHaRole() string`

GetHaRole returns the HaRole field if non-nil, zero value otherwise.

### GetHaRoleOk

`func (o *Server) GetHaRoleOk() (*string, bool)`

GetHaRoleOk returns a tuple with the HaRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaRole

`func (o *Server) SetHaRole(v string)`

SetHaRole sets HaRole field to given value.

### HasHaRole

`func (o *Server) HasHaRole() bool`

HasHaRole returns a boolean if a field has been set.

### GetHaPeerConnectionState

`func (o *Server) GetHaPeerConnectionState() string`

GetHaPeerConnectionState returns the HaPeerConnectionState field if non-nil, zero value otherwise.

### GetHaPeerConnectionStateOk

`func (o *Server) GetHaPeerConnectionStateOk() (*string, bool)`

GetHaPeerConnectionStateOk returns a tuple with the HaPeerConnectionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaPeerConnectionState

`func (o *Server) SetHaPeerConnectionState(v string)`

SetHaPeerConnectionState sets HaPeerConnectionState field to given value.

### HasHaPeerConnectionState

`func (o *Server) HasHaPeerConnectionState() bool`

HasHaPeerConnectionState returns a boolean if a field has been set.

### GetHaDiskState

`func (o *Server) GetHaDiskState() string`

GetHaDiskState returns the HaDiskState field if non-nil, zero value otherwise.

### GetHaDiskStateOk

`func (o *Server) GetHaDiskStateOk() (*string, bool)`

GetHaDiskStateOk returns a tuple with the HaDiskState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaDiskState

`func (o *Server) SetHaDiskState(v string)`

SetHaDiskState sets HaDiskState field to given value.

### HasHaDiskState

`func (o *Server) HasHaDiskState() bool`

HasHaDiskState returns a boolean if a field has been set.

### GetHsmSupportEnabled

`func (o *Server) GetHsmSupportEnabled() bool`

GetHsmSupportEnabled returns the HsmSupportEnabled field if non-nil, zero value otherwise.

### GetHsmSupportEnabledOk

`func (o *Server) GetHsmSupportEnabledOk() (*bool, bool)`

GetHsmSupportEnabledOk returns a tuple with the HsmSupportEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHsmSupportEnabled

`func (o *Server) SetHsmSupportEnabled(v bool)`

SetHsmSupportEnabled sets HsmSupportEnabled field to given value.

### HasHsmSupportEnabled

`func (o *Server) HasHsmSupportEnabled() bool`

HasHsmSupportEnabled returns a boolean if a field has been set.

### GetInterfaceRedundancyEnabled

`func (o *Server) GetInterfaceRedundancyEnabled() bool`

GetInterfaceRedundancyEnabled returns the InterfaceRedundancyEnabled field if non-nil, zero value otherwise.

### GetInterfaceRedundancyEnabledOk

`func (o *Server) GetInterfaceRedundancyEnabledOk() (*bool, bool)`

GetInterfaceRedundancyEnabledOk returns a tuple with the InterfaceRedundancyEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceRedundancyEnabled

`func (o *Server) SetInterfaceRedundancyEnabled(v bool)`

SetInterfaceRedundancyEnabled sets InterfaceRedundancyEnabled field to given value.

### HasInterfaceRedundancyEnabled

`func (o *Server) HasInterfaceRedundancyEnabled() bool`

HasInterfaceRedundancyEnabled returns a boolean if a field has been set.

### GetInheritedFields

`func (o *Server) GetInheritedFields() []string`

GetInheritedFields returns the InheritedFields field if non-nil, zero value otherwise.

### GetInheritedFieldsOk

`func (o *Server) GetInheritedFieldsOk() (*[]string, bool)`

GetInheritedFieldsOk returns a tuple with the InheritedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritedFields

`func (o *Server) SetInheritedFields(v []string)`

SetInheritedFields sets InheritedFields field to given value.

### HasInheritedFields

`func (o *Server) HasInheritedFields() bool`

HasInheritedFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


