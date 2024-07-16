# AbstractServer

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
**InheritedFields** | Pointer to **[]string** |  | [optional] [readonly] 

## Methods

### NewAbstractServer

`func NewAbstractServer() *AbstractServer`

NewAbstractServer instantiates a new AbstractServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAbstractServerWithDefaults

`func NewAbstractServerWithDefaults() *AbstractServer`

NewAbstractServerWithDefaults instantiates a new AbstractServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AbstractServer) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AbstractServer) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AbstractServer) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AbstractServer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *AbstractServer) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AbstractServer) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AbstractServer) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AbstractServer) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *AbstractServer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AbstractServer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AbstractServer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AbstractServer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *AbstractServer) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *AbstractServer) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *AbstractServer) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *AbstractServer) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetConfiguration

`func (o *AbstractServer) GetConfiguration() InlinedConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *AbstractServer) GetConfigurationOk() (*InlinedConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *AbstractServer) SetConfiguration(v InlinedConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *AbstractServer) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetProfile

`func (o *AbstractServer) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *AbstractServer) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *AbstractServer) SetProfile(v string)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *AbstractServer) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetState

`func (o *AbstractServer) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AbstractServer) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AbstractServer) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *AbstractServer) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUsername

`func (o *AbstractServer) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AbstractServer) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AbstractServer) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *AbstractServer) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *AbstractServer) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AbstractServer) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AbstractServer) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *AbstractServer) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetServerGroup

`func (o *AbstractServer) GetServerGroup() InlinedServerGroup`

GetServerGroup returns the ServerGroup field if non-nil, zero value otherwise.

### GetServerGroupOk

`func (o *AbstractServer) GetServerGroupOk() (*InlinedServerGroup, bool)`

GetServerGroupOk returns a tuple with the ServerGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerGroup

`func (o *AbstractServer) SetServerGroup(v InlinedServerGroup)`

SetServerGroup sets ServerGroup field to given value.

### HasServerGroup

`func (o *AbstractServer) HasServerGroup() bool`

HasServerGroup returns a boolean if a field has been set.

### GetConnected

`func (o *AbstractServer) GetConnected() bool`

GetConnected returns the Connected field if non-nil, zero value otherwise.

### GetConnectedOk

`func (o *AbstractServer) GetConnectedOk() (*bool, bool)`

GetConnectedOk returns a tuple with the Connected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnected

`func (o *AbstractServer) SetConnected(v bool)`

SetConnected sets Connected field to given value.

### HasConnected

`func (o *AbstractServer) HasConnected() bool`

HasConnected returns a boolean if a field has been set.

### GetMonitoringEnabled

`func (o *AbstractServer) GetMonitoringEnabled() bool`

GetMonitoringEnabled returns the MonitoringEnabled field if non-nil, zero value otherwise.

### GetMonitoringEnabledOk

`func (o *AbstractServer) GetMonitoringEnabledOk() (*bool, bool)`

GetMonitoringEnabledOk returns a tuple with the MonitoringEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringEnabled

`func (o *AbstractServer) SetMonitoringEnabled(v bool)`

SetMonitoringEnabled sets MonitoringEnabled field to given value.

### HasMonitoringEnabled

`func (o *AbstractServer) HasMonitoringEnabled() bool`

HasMonitoringEnabled returns a boolean if a field has been set.

### GetSnmp

`func (o *AbstractServer) GetSnmp() SnmpBean`

GetSnmp returns the Snmp field if non-nil, zero value otherwise.

### GetSnmpOk

`func (o *AbstractServer) GetSnmpOk() (*SnmpBean, bool)`

GetSnmpOk returns a tuple with the Snmp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmp

`func (o *AbstractServer) SetSnmp(v SnmpBean)`

SetSnmp sets Snmp field to given value.

### HasSnmp

`func (o *AbstractServer) HasSnmp() bool`

HasSnmp returns a boolean if a field has been set.

### GetInterfaces

`func (o *AbstractServer) GetInterfaces() []NetworkInterface`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *AbstractServer) GetInterfacesOk() (*[]NetworkInterface, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *AbstractServer) SetInterfaces(v []NetworkInterface)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *AbstractServer) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.

### GetLocation

`func (o *AbstractServer) GetLocation() InlinedLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *AbstractServer) GetLocationOk() (*InlinedLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *AbstractServer) SetLocation(v InlinedLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *AbstractServer) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetStatistics

`func (o *AbstractServer) GetStatistics() map[string]map[string]interface{}`

GetStatistics returns the Statistics field if non-nil, zero value otherwise.

### GetStatisticsOk

`func (o *AbstractServer) GetStatisticsOk() (*map[string]map[string]interface{}, bool)`

GetStatisticsOk returns a tuple with the Statistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatistics

`func (o *AbstractServer) SetStatistics(v map[string]map[string]interface{})`

SetStatistics sets Statistics field to given value.

### HasStatistics

`func (o *AbstractServer) HasStatistics() bool`

HasStatistics returns a boolean if a field has been set.

### GetOverrideConfigurationLevelDhcpValidation

`func (o *AbstractServer) GetOverrideConfigurationLevelDhcpValidation() bool`

GetOverrideConfigurationLevelDhcpValidation returns the OverrideConfigurationLevelDhcpValidation field if non-nil, zero value otherwise.

### GetOverrideConfigurationLevelDhcpValidationOk

`func (o *AbstractServer) GetOverrideConfigurationLevelDhcpValidationOk() (*bool, bool)`

GetOverrideConfigurationLevelDhcpValidationOk returns a tuple with the OverrideConfigurationLevelDhcpValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideConfigurationLevelDhcpValidation

`func (o *AbstractServer) SetOverrideConfigurationLevelDhcpValidation(v bool)`

SetOverrideConfigurationLevelDhcpValidation sets OverrideConfigurationLevelDhcpValidation field to given value.

### HasOverrideConfigurationLevelDhcpValidation

`func (o *AbstractServer) HasOverrideConfigurationLevelDhcpValidation() bool`

HasOverrideConfigurationLevelDhcpValidation returns a boolean if a field has been set.

### GetOverrideConfigurationLevelDnsValidation

`func (o *AbstractServer) GetOverrideConfigurationLevelDnsValidation() bool`

GetOverrideConfigurationLevelDnsValidation returns the OverrideConfigurationLevelDnsValidation field if non-nil, zero value otherwise.

### GetOverrideConfigurationLevelDnsValidationOk

`func (o *AbstractServer) GetOverrideConfigurationLevelDnsValidationOk() (*bool, bool)`

GetOverrideConfigurationLevelDnsValidationOk returns a tuple with the OverrideConfigurationLevelDnsValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideConfigurationLevelDnsValidation

`func (o *AbstractServer) SetOverrideConfigurationLevelDnsValidation(v bool)`

SetOverrideConfigurationLevelDnsValidation sets OverrideConfigurationLevelDnsValidation field to given value.

### HasOverrideConfigurationLevelDnsValidation

`func (o *AbstractServer) HasOverrideConfigurationLevelDnsValidation() bool`

HasOverrideConfigurationLevelDnsValidation returns a boolean if a field has been set.

### GetDhcpConfigurationValidationEnabled

`func (o *AbstractServer) GetDhcpConfigurationValidationEnabled() bool`

GetDhcpConfigurationValidationEnabled returns the DhcpConfigurationValidationEnabled field if non-nil, zero value otherwise.

### GetDhcpConfigurationValidationEnabledOk

`func (o *AbstractServer) GetDhcpConfigurationValidationEnabledOk() (*bool, bool)`

GetDhcpConfigurationValidationEnabledOk returns a tuple with the DhcpConfigurationValidationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpConfigurationValidationEnabled

`func (o *AbstractServer) SetDhcpConfigurationValidationEnabled(v bool)`

SetDhcpConfigurationValidationEnabled sets DhcpConfigurationValidationEnabled field to given value.

### HasDhcpConfigurationValidationEnabled

`func (o *AbstractServer) HasDhcpConfigurationValidationEnabled() bool`

HasDhcpConfigurationValidationEnabled returns a boolean if a field has been set.

### GetDnsConfigurationValidationEnabled

`func (o *AbstractServer) GetDnsConfigurationValidationEnabled() bool`

GetDnsConfigurationValidationEnabled returns the DnsConfigurationValidationEnabled field if non-nil, zero value otherwise.

### GetDnsConfigurationValidationEnabledOk

`func (o *AbstractServer) GetDnsConfigurationValidationEnabledOk() (*bool, bool)`

GetDnsConfigurationValidationEnabledOk returns a tuple with the DnsConfigurationValidationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsConfigurationValidationEnabled

`func (o *AbstractServer) SetDnsConfigurationValidationEnabled(v bool)`

SetDnsConfigurationValidationEnabled sets DnsConfigurationValidationEnabled field to given value.

### HasDnsConfigurationValidationEnabled

`func (o *AbstractServer) HasDnsConfigurationValidationEnabled() bool`

HasDnsConfigurationValidationEnabled returns a boolean if a field has been set.

### GetDnsZoneValidationEnabled

`func (o *AbstractServer) GetDnsZoneValidationEnabled() bool`

GetDnsZoneValidationEnabled returns the DnsZoneValidationEnabled field if non-nil, zero value otherwise.

### GetDnsZoneValidationEnabledOk

`func (o *AbstractServer) GetDnsZoneValidationEnabledOk() (*bool, bool)`

GetDnsZoneValidationEnabledOk returns a tuple with the DnsZoneValidationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsZoneValidationEnabled

`func (o *AbstractServer) SetDnsZoneValidationEnabled(v bool)`

SetDnsZoneValidationEnabled sets DnsZoneValidationEnabled field to given value.

### HasDnsZoneValidationEnabled

`func (o *AbstractServer) HasDnsZoneValidationEnabled() bool`

HasDnsZoneValidationEnabled returns a boolean if a field has been set.

### GetCheckIntegrityValidation

`func (o *AbstractServer) GetCheckIntegrityValidation() string`

GetCheckIntegrityValidation returns the CheckIntegrityValidation field if non-nil, zero value otherwise.

### GetCheckIntegrityValidationOk

`func (o *AbstractServer) GetCheckIntegrityValidationOk() (*string, bool)`

GetCheckIntegrityValidationOk returns a tuple with the CheckIntegrityValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckIntegrityValidation

`func (o *AbstractServer) SetCheckIntegrityValidation(v string)`

SetCheckIntegrityValidation sets CheckIntegrityValidation field to given value.

### HasCheckIntegrityValidation

`func (o *AbstractServer) HasCheckIntegrityValidation() bool`

HasCheckIntegrityValidation returns a boolean if a field has been set.

### GetCheckMxCnameValidation

`func (o *AbstractServer) GetCheckMxCnameValidation() string`

GetCheckMxCnameValidation returns the CheckMxCnameValidation field if non-nil, zero value otherwise.

### GetCheckMxCnameValidationOk

`func (o *AbstractServer) GetCheckMxCnameValidationOk() (*string, bool)`

GetCheckMxCnameValidationOk returns a tuple with the CheckMxCnameValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckMxCnameValidation

`func (o *AbstractServer) SetCheckMxCnameValidation(v string)`

SetCheckMxCnameValidation sets CheckMxCnameValidation field to given value.

### HasCheckMxCnameValidation

`func (o *AbstractServer) HasCheckMxCnameValidation() bool`

HasCheckMxCnameValidation returns a boolean if a field has been set.

### GetCheckMxValidation

`func (o *AbstractServer) GetCheckMxValidation() string`

GetCheckMxValidation returns the CheckMxValidation field if non-nil, zero value otherwise.

### GetCheckMxValidationOk

`func (o *AbstractServer) GetCheckMxValidationOk() (*string, bool)`

GetCheckMxValidationOk returns a tuple with the CheckMxValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckMxValidation

`func (o *AbstractServer) SetCheckMxValidation(v string)`

SetCheckMxValidation sets CheckMxValidation field to given value.

### HasCheckMxValidation

`func (o *AbstractServer) HasCheckMxValidation() bool`

HasCheckMxValidation returns a boolean if a field has been set.

### GetCheckNamesValidation

`func (o *AbstractServer) GetCheckNamesValidation() string`

GetCheckNamesValidation returns the CheckNamesValidation field if non-nil, zero value otherwise.

### GetCheckNamesValidationOk

`func (o *AbstractServer) GetCheckNamesValidationOk() (*string, bool)`

GetCheckNamesValidationOk returns a tuple with the CheckNamesValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckNamesValidation

`func (o *AbstractServer) SetCheckNamesValidation(v string)`

SetCheckNamesValidation sets CheckNamesValidation field to given value.

### HasCheckNamesValidation

`func (o *AbstractServer) HasCheckNamesValidation() bool`

HasCheckNamesValidation returns a boolean if a field has been set.

### GetCheckNsValidation

`func (o *AbstractServer) GetCheckNsValidation() string`

GetCheckNsValidation returns the CheckNsValidation field if non-nil, zero value otherwise.

### GetCheckNsValidationOk

`func (o *AbstractServer) GetCheckNsValidationOk() (*string, bool)`

GetCheckNsValidationOk returns a tuple with the CheckNsValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckNsValidation

`func (o *AbstractServer) SetCheckNsValidation(v string)`

SetCheckNsValidation sets CheckNsValidation field to given value.

### HasCheckNsValidation

`func (o *AbstractServer) HasCheckNsValidation() bool`

HasCheckNsValidation returns a boolean if a field has been set.

### GetCheckSrvCnameValidation

`func (o *AbstractServer) GetCheckSrvCnameValidation() string`

GetCheckSrvCnameValidation returns the CheckSrvCnameValidation field if non-nil, zero value otherwise.

### GetCheckSrvCnameValidationOk

`func (o *AbstractServer) GetCheckSrvCnameValidationOk() (*string, bool)`

GetCheckSrvCnameValidationOk returns a tuple with the CheckSrvCnameValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckSrvCnameValidation

`func (o *AbstractServer) SetCheckSrvCnameValidation(v string)`

SetCheckSrvCnameValidation sets CheckSrvCnameValidation field to given value.

### HasCheckSrvCnameValidation

`func (o *AbstractServer) HasCheckSrvCnameValidation() bool`

HasCheckSrvCnameValidation returns a boolean if a field has been set.

### GetCheckWildcardValidation

`func (o *AbstractServer) GetCheckWildcardValidation() string`

GetCheckWildcardValidation returns the CheckWildcardValidation field if non-nil, zero value otherwise.

### GetCheckWildcardValidationOk

`func (o *AbstractServer) GetCheckWildcardValidationOk() (*string, bool)`

GetCheckWildcardValidationOk returns a tuple with the CheckWildcardValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckWildcardValidation

`func (o *AbstractServer) SetCheckWildcardValidation(v string)`

SetCheckWildcardValidation sets CheckWildcardValidation field to given value.

### HasCheckWildcardValidation

`func (o *AbstractServer) HasCheckWildcardValidation() bool`

HasCheckWildcardValidation returns a boolean if a field has been set.

### GetPrivateAddress

`func (o *AbstractServer) GetPrivateAddress() string`

GetPrivateAddress returns the PrivateAddress field if non-nil, zero value otherwise.

### GetPrivateAddressOk

`func (o *AbstractServer) GetPrivateAddressOk() (*string, bool)`

GetPrivateAddressOk returns a tuple with the PrivateAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateAddress

`func (o *AbstractServer) SetPrivateAddress(v string)`

SetPrivateAddress sets PrivateAddress field to given value.

### HasPrivateAddress

`func (o *AbstractServer) HasPrivateAddress() bool`

HasPrivateAddress returns a boolean if a field has been set.

### GetEncryptedNotificationsEnabled

`func (o *AbstractServer) GetEncryptedNotificationsEnabled() bool`

GetEncryptedNotificationsEnabled returns the EncryptedNotificationsEnabled field if non-nil, zero value otherwise.

### GetEncryptedNotificationsEnabledOk

`func (o *AbstractServer) GetEncryptedNotificationsEnabledOk() (*bool, bool)`

GetEncryptedNotificationsEnabledOk returns a tuple with the EncryptedNotificationsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedNotificationsEnabled

`func (o *AbstractServer) SetEncryptedNotificationsEnabled(v bool)`

SetEncryptedNotificationsEnabled sets EncryptedNotificationsEnabled field to given value.

### HasEncryptedNotificationsEnabled

`func (o *AbstractServer) HasEncryptedNotificationsEnabled() bool`

HasEncryptedNotificationsEnabled returns a boolean if a field has been set.

### GetManagementUrl

`func (o *AbstractServer) GetManagementUrl() string`

GetManagementUrl returns the ManagementUrl field if non-nil, zero value otherwise.

### GetManagementUrlOk

`func (o *AbstractServer) GetManagementUrlOk() (*string, bool)`

GetManagementUrlOk returns a tuple with the ManagementUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementUrl

`func (o *AbstractServer) SetManagementUrl(v string)`

SetManagementUrl sets ManagementUrl field to given value.

### HasManagementUrl

`func (o *AbstractServer) HasManagementUrl() bool`

HasManagementUrl returns a boolean if a field has been set.

### GetSelfIpAddress

`func (o *AbstractServer) GetSelfIpAddress() string`

GetSelfIpAddress returns the SelfIpAddress field if non-nil, zero value otherwise.

### GetSelfIpAddressOk

`func (o *AbstractServer) GetSelfIpAddressOk() (*string, bool)`

GetSelfIpAddressOk returns a tuple with the SelfIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfIpAddress

`func (o *AbstractServer) SetSelfIpAddress(v string)`

SetSelfIpAddress sets SelfIpAddress field to given value.

### HasSelfIpAddress

`func (o *AbstractServer) HasSelfIpAddress() bool`

HasSelfIpAddress returns a boolean if a field has been set.

### GetHaBackboneEnabled

`func (o *AbstractServer) GetHaBackboneEnabled() bool`

GetHaBackboneEnabled returns the HaBackboneEnabled field if non-nil, zero value otherwise.

### GetHaBackboneEnabledOk

`func (o *AbstractServer) GetHaBackboneEnabledOk() (*bool, bool)`

GetHaBackboneEnabledOk returns a tuple with the HaBackboneEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaBackboneEnabled

`func (o *AbstractServer) SetHaBackboneEnabled(v bool)`

SetHaBackboneEnabled sets HaBackboneEnabled field to given value.

### HasHaBackboneEnabled

`func (o *AbstractServer) HasHaBackboneEnabled() bool`

HasHaBackboneEnabled returns a boolean if a field has been set.

### GetHaPingAddress

`func (o *AbstractServer) GetHaPingAddress() string`

GetHaPingAddress returns the HaPingAddress field if non-nil, zero value otherwise.

### GetHaPingAddressOk

`func (o *AbstractServer) GetHaPingAddressOk() (*string, bool)`

GetHaPingAddressOk returns a tuple with the HaPingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaPingAddress

`func (o *AbstractServer) SetHaPingAddress(v string)`

SetHaPingAddress sets HaPingAddress field to given value.

### HasHaPingAddress

`func (o *AbstractServer) HasHaPingAddress() bool`

HasHaPingAddress returns a boolean if a field has been set.

### GetDhcpServicePrincipal

`func (o *AbstractServer) GetDhcpServicePrincipal() InlinedKerberosServicePrincipal`

GetDhcpServicePrincipal returns the DhcpServicePrincipal field if non-nil, zero value otherwise.

### GetDhcpServicePrincipalOk

`func (o *AbstractServer) GetDhcpServicePrincipalOk() (*InlinedKerberosServicePrincipal, bool)`

GetDhcpServicePrincipalOk returns a tuple with the DhcpServicePrincipal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpServicePrincipal

`func (o *AbstractServer) SetDhcpServicePrincipal(v InlinedKerberosServicePrincipal)`

SetDhcpServicePrincipal sets DhcpServicePrincipal field to given value.

### HasDhcpServicePrincipal

`func (o *AbstractServer) HasDhcpServicePrincipal() bool`

HasDhcpServicePrincipal returns a boolean if a field has been set.

### GetDnsServicePrincipal

`func (o *AbstractServer) GetDnsServicePrincipal() InlinedKerberosServicePrincipal`

GetDnsServicePrincipal returns the DnsServicePrincipal field if non-nil, zero value otherwise.

### GetDnsServicePrincipalOk

`func (o *AbstractServer) GetDnsServicePrincipalOk() (*InlinedKerberosServicePrincipal, bool)`

GetDnsServicePrincipalOk returns a tuple with the DnsServicePrincipal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServicePrincipal

`func (o *AbstractServer) SetDnsServicePrincipal(v InlinedKerberosServicePrincipal)`

SetDnsServicePrincipal sets DnsServicePrincipal field to given value.

### HasDnsServicePrincipal

`func (o *AbstractServer) HasDnsServicePrincipal() bool`

HasDnsServicePrincipal returns a boolean if a field has been set.

### GetInheritedFields

`func (o *AbstractServer) GetInheritedFields() []string`

GetInheritedFields returns the InheritedFields field if non-nil, zero value otherwise.

### GetInheritedFieldsOk

`func (o *AbstractServer) GetInheritedFieldsOk() (*[]string, bool)`

GetInheritedFieldsOk returns a tuple with the InheritedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritedFields

`func (o *AbstractServer) SetInheritedFields(v []string)`

SetInheritedFields sets InheritedFields field to given value.

### HasInheritedFields

`func (o *AbstractServer) HasInheritedFields() bool`

HasInheritedFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


