# GetCollectionServers200ResponseDataInner

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
**ActiveServer** | Pointer to [**InlinedServer**](InlinedServer.md) |  | [optional] 
**PassiveServer** | Pointer to [**InlinedServer**](InlinedServer.md) |  | [optional] 
**DedicatedManagementEnabled** | Pointer to **bool** | Indicates whether dedicated management is enabled on the server. | [optional] 
**HaRole** | Pointer to **string** |  | [optional] [readonly] 
**HaPeerConnectionState** | Pointer to **string** |  | [optional] [readonly] 
**HaDiskState** | Pointer to **string** |  | [optional] [readonly] 
**HsmSupportEnabled** | Pointer to **bool** | Indicates whether HSM is enabled on the server. | [optional] 
**InterfaceRedundancyEnabled** | Pointer to **bool** | Indicates whether interface redundancy is enabled on the server. | [optional] 
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 

## Methods

### NewGetCollectionServers200ResponseDataInner

`func NewGetCollectionServers200ResponseDataInner() *GetCollectionServers200ResponseDataInner`

NewGetCollectionServers200ResponseDataInner instantiates a new GetCollectionServers200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCollectionServers200ResponseDataInnerWithDefaults

`func NewGetCollectionServers200ResponseDataInnerWithDefaults() *GetCollectionServers200ResponseDataInner`

NewGetCollectionServers200ResponseDataInnerWithDefaults instantiates a new GetCollectionServers200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetCollectionServers200ResponseDataInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetCollectionServers200ResponseDataInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetCollectionServers200ResponseDataInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetCollectionServers200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GetCollectionServers200ResponseDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetCollectionServers200ResponseDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetCollectionServers200ResponseDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetCollectionServers200ResponseDataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *GetCollectionServers200ResponseDataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetCollectionServers200ResponseDataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetCollectionServers200ResponseDataInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetCollectionServers200ResponseDataInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *GetCollectionServers200ResponseDataInner) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *GetCollectionServers200ResponseDataInner) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *GetCollectionServers200ResponseDataInner) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *GetCollectionServers200ResponseDataInner) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetConfiguration

`func (o *GetCollectionServers200ResponseDataInner) GetConfiguration() InlinedConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *GetCollectionServers200ResponseDataInner) GetConfigurationOk() (*InlinedConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *GetCollectionServers200ResponseDataInner) SetConfiguration(v InlinedConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *GetCollectionServers200ResponseDataInner) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetProfile

`func (o *GetCollectionServers200ResponseDataInner) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *GetCollectionServers200ResponseDataInner) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *GetCollectionServers200ResponseDataInner) SetProfile(v string)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *GetCollectionServers200ResponseDataInner) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetState

`func (o *GetCollectionServers200ResponseDataInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetCollectionServers200ResponseDataInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetCollectionServers200ResponseDataInner) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GetCollectionServers200ResponseDataInner) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUsername

`func (o *GetCollectionServers200ResponseDataInner) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GetCollectionServers200ResponseDataInner) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GetCollectionServers200ResponseDataInner) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *GetCollectionServers200ResponseDataInner) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *GetCollectionServers200ResponseDataInner) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *GetCollectionServers200ResponseDataInner) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *GetCollectionServers200ResponseDataInner) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *GetCollectionServers200ResponseDataInner) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetServerGroup

`func (o *GetCollectionServers200ResponseDataInner) GetServerGroup() InlinedServerGroup`

GetServerGroup returns the ServerGroup field if non-nil, zero value otherwise.

### GetServerGroupOk

`func (o *GetCollectionServers200ResponseDataInner) GetServerGroupOk() (*InlinedServerGroup, bool)`

GetServerGroupOk returns a tuple with the ServerGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerGroup

`func (o *GetCollectionServers200ResponseDataInner) SetServerGroup(v InlinedServerGroup)`

SetServerGroup sets ServerGroup field to given value.

### HasServerGroup

`func (o *GetCollectionServers200ResponseDataInner) HasServerGroup() bool`

HasServerGroup returns a boolean if a field has been set.

### GetConnected

`func (o *GetCollectionServers200ResponseDataInner) GetConnected() bool`

GetConnected returns the Connected field if non-nil, zero value otherwise.

### GetConnectedOk

`func (o *GetCollectionServers200ResponseDataInner) GetConnectedOk() (*bool, bool)`

GetConnectedOk returns a tuple with the Connected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnected

`func (o *GetCollectionServers200ResponseDataInner) SetConnected(v bool)`

SetConnected sets Connected field to given value.

### HasConnected

`func (o *GetCollectionServers200ResponseDataInner) HasConnected() bool`

HasConnected returns a boolean if a field has been set.

### GetMonitoringEnabled

`func (o *GetCollectionServers200ResponseDataInner) GetMonitoringEnabled() bool`

GetMonitoringEnabled returns the MonitoringEnabled field if non-nil, zero value otherwise.

### GetMonitoringEnabledOk

`func (o *GetCollectionServers200ResponseDataInner) GetMonitoringEnabledOk() (*bool, bool)`

GetMonitoringEnabledOk returns a tuple with the MonitoringEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringEnabled

`func (o *GetCollectionServers200ResponseDataInner) SetMonitoringEnabled(v bool)`

SetMonitoringEnabled sets MonitoringEnabled field to given value.

### HasMonitoringEnabled

`func (o *GetCollectionServers200ResponseDataInner) HasMonitoringEnabled() bool`

HasMonitoringEnabled returns a boolean if a field has been set.

### GetSnmp

`func (o *GetCollectionServers200ResponseDataInner) GetSnmp() SnmpBean`

GetSnmp returns the Snmp field if non-nil, zero value otherwise.

### GetSnmpOk

`func (o *GetCollectionServers200ResponseDataInner) GetSnmpOk() (*SnmpBean, bool)`

GetSnmpOk returns a tuple with the Snmp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmp

`func (o *GetCollectionServers200ResponseDataInner) SetSnmp(v SnmpBean)`

SetSnmp sets Snmp field to given value.

### HasSnmp

`func (o *GetCollectionServers200ResponseDataInner) HasSnmp() bool`

HasSnmp returns a boolean if a field has been set.

### GetInterfaces

`func (o *GetCollectionServers200ResponseDataInner) GetInterfaces() []NetworkInterface`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *GetCollectionServers200ResponseDataInner) GetInterfacesOk() (*[]NetworkInterface, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *GetCollectionServers200ResponseDataInner) SetInterfaces(v []NetworkInterface)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *GetCollectionServers200ResponseDataInner) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.

### GetLocation

`func (o *GetCollectionServers200ResponseDataInner) GetLocation() InlinedLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *GetCollectionServers200ResponseDataInner) GetLocationOk() (*InlinedLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *GetCollectionServers200ResponseDataInner) SetLocation(v InlinedLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *GetCollectionServers200ResponseDataInner) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetStatistics

`func (o *GetCollectionServers200ResponseDataInner) GetStatistics() map[string]map[string]interface{}`

GetStatistics returns the Statistics field if non-nil, zero value otherwise.

### GetStatisticsOk

`func (o *GetCollectionServers200ResponseDataInner) GetStatisticsOk() (*map[string]map[string]interface{}, bool)`

GetStatisticsOk returns a tuple with the Statistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatistics

`func (o *GetCollectionServers200ResponseDataInner) SetStatistics(v map[string]map[string]interface{})`

SetStatistics sets Statistics field to given value.

### HasStatistics

`func (o *GetCollectionServers200ResponseDataInner) HasStatistics() bool`

HasStatistics returns a boolean if a field has been set.

### GetOverrideConfigurationLevelDhcpValidation

`func (o *GetCollectionServers200ResponseDataInner) GetOverrideConfigurationLevelDhcpValidation() bool`

GetOverrideConfigurationLevelDhcpValidation returns the OverrideConfigurationLevelDhcpValidation field if non-nil, zero value otherwise.

### GetOverrideConfigurationLevelDhcpValidationOk

`func (o *GetCollectionServers200ResponseDataInner) GetOverrideConfigurationLevelDhcpValidationOk() (*bool, bool)`

GetOverrideConfigurationLevelDhcpValidationOk returns a tuple with the OverrideConfigurationLevelDhcpValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideConfigurationLevelDhcpValidation

`func (o *GetCollectionServers200ResponseDataInner) SetOverrideConfigurationLevelDhcpValidation(v bool)`

SetOverrideConfigurationLevelDhcpValidation sets OverrideConfigurationLevelDhcpValidation field to given value.

### HasOverrideConfigurationLevelDhcpValidation

`func (o *GetCollectionServers200ResponseDataInner) HasOverrideConfigurationLevelDhcpValidation() bool`

HasOverrideConfigurationLevelDhcpValidation returns a boolean if a field has been set.

### GetOverrideConfigurationLevelDnsValidation

`func (o *GetCollectionServers200ResponseDataInner) GetOverrideConfigurationLevelDnsValidation() bool`

GetOverrideConfigurationLevelDnsValidation returns the OverrideConfigurationLevelDnsValidation field if non-nil, zero value otherwise.

### GetOverrideConfigurationLevelDnsValidationOk

`func (o *GetCollectionServers200ResponseDataInner) GetOverrideConfigurationLevelDnsValidationOk() (*bool, bool)`

GetOverrideConfigurationLevelDnsValidationOk returns a tuple with the OverrideConfigurationLevelDnsValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideConfigurationLevelDnsValidation

`func (o *GetCollectionServers200ResponseDataInner) SetOverrideConfigurationLevelDnsValidation(v bool)`

SetOverrideConfigurationLevelDnsValidation sets OverrideConfigurationLevelDnsValidation field to given value.

### HasOverrideConfigurationLevelDnsValidation

`func (o *GetCollectionServers200ResponseDataInner) HasOverrideConfigurationLevelDnsValidation() bool`

HasOverrideConfigurationLevelDnsValidation returns a boolean if a field has been set.

### GetDhcpConfigurationValidationEnabled

`func (o *GetCollectionServers200ResponseDataInner) GetDhcpConfigurationValidationEnabled() bool`

GetDhcpConfigurationValidationEnabled returns the DhcpConfigurationValidationEnabled field if non-nil, zero value otherwise.

### GetDhcpConfigurationValidationEnabledOk

`func (o *GetCollectionServers200ResponseDataInner) GetDhcpConfigurationValidationEnabledOk() (*bool, bool)`

GetDhcpConfigurationValidationEnabledOk returns a tuple with the DhcpConfigurationValidationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpConfigurationValidationEnabled

`func (o *GetCollectionServers200ResponseDataInner) SetDhcpConfigurationValidationEnabled(v bool)`

SetDhcpConfigurationValidationEnabled sets DhcpConfigurationValidationEnabled field to given value.

### HasDhcpConfigurationValidationEnabled

`func (o *GetCollectionServers200ResponseDataInner) HasDhcpConfigurationValidationEnabled() bool`

HasDhcpConfigurationValidationEnabled returns a boolean if a field has been set.

### GetDnsConfigurationValidationEnabled

`func (o *GetCollectionServers200ResponseDataInner) GetDnsConfigurationValidationEnabled() bool`

GetDnsConfigurationValidationEnabled returns the DnsConfigurationValidationEnabled field if non-nil, zero value otherwise.

### GetDnsConfigurationValidationEnabledOk

`func (o *GetCollectionServers200ResponseDataInner) GetDnsConfigurationValidationEnabledOk() (*bool, bool)`

GetDnsConfigurationValidationEnabledOk returns a tuple with the DnsConfigurationValidationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsConfigurationValidationEnabled

`func (o *GetCollectionServers200ResponseDataInner) SetDnsConfigurationValidationEnabled(v bool)`

SetDnsConfigurationValidationEnabled sets DnsConfigurationValidationEnabled field to given value.

### HasDnsConfigurationValidationEnabled

`func (o *GetCollectionServers200ResponseDataInner) HasDnsConfigurationValidationEnabled() bool`

HasDnsConfigurationValidationEnabled returns a boolean if a field has been set.

### GetDnsZoneValidationEnabled

`func (o *GetCollectionServers200ResponseDataInner) GetDnsZoneValidationEnabled() bool`

GetDnsZoneValidationEnabled returns the DnsZoneValidationEnabled field if non-nil, zero value otherwise.

### GetDnsZoneValidationEnabledOk

`func (o *GetCollectionServers200ResponseDataInner) GetDnsZoneValidationEnabledOk() (*bool, bool)`

GetDnsZoneValidationEnabledOk returns a tuple with the DnsZoneValidationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsZoneValidationEnabled

`func (o *GetCollectionServers200ResponseDataInner) SetDnsZoneValidationEnabled(v bool)`

SetDnsZoneValidationEnabled sets DnsZoneValidationEnabled field to given value.

### HasDnsZoneValidationEnabled

`func (o *GetCollectionServers200ResponseDataInner) HasDnsZoneValidationEnabled() bool`

HasDnsZoneValidationEnabled returns a boolean if a field has been set.

### GetCheckIntegrityValidation

`func (o *GetCollectionServers200ResponseDataInner) GetCheckIntegrityValidation() string`

GetCheckIntegrityValidation returns the CheckIntegrityValidation field if non-nil, zero value otherwise.

### GetCheckIntegrityValidationOk

`func (o *GetCollectionServers200ResponseDataInner) GetCheckIntegrityValidationOk() (*string, bool)`

GetCheckIntegrityValidationOk returns a tuple with the CheckIntegrityValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckIntegrityValidation

`func (o *GetCollectionServers200ResponseDataInner) SetCheckIntegrityValidation(v string)`

SetCheckIntegrityValidation sets CheckIntegrityValidation field to given value.

### HasCheckIntegrityValidation

`func (o *GetCollectionServers200ResponseDataInner) HasCheckIntegrityValidation() bool`

HasCheckIntegrityValidation returns a boolean if a field has been set.

### GetCheckMxCnameValidation

`func (o *GetCollectionServers200ResponseDataInner) GetCheckMxCnameValidation() string`

GetCheckMxCnameValidation returns the CheckMxCnameValidation field if non-nil, zero value otherwise.

### GetCheckMxCnameValidationOk

`func (o *GetCollectionServers200ResponseDataInner) GetCheckMxCnameValidationOk() (*string, bool)`

GetCheckMxCnameValidationOk returns a tuple with the CheckMxCnameValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckMxCnameValidation

`func (o *GetCollectionServers200ResponseDataInner) SetCheckMxCnameValidation(v string)`

SetCheckMxCnameValidation sets CheckMxCnameValidation field to given value.

### HasCheckMxCnameValidation

`func (o *GetCollectionServers200ResponseDataInner) HasCheckMxCnameValidation() bool`

HasCheckMxCnameValidation returns a boolean if a field has been set.

### GetCheckMxValidation

`func (o *GetCollectionServers200ResponseDataInner) GetCheckMxValidation() string`

GetCheckMxValidation returns the CheckMxValidation field if non-nil, zero value otherwise.

### GetCheckMxValidationOk

`func (o *GetCollectionServers200ResponseDataInner) GetCheckMxValidationOk() (*string, bool)`

GetCheckMxValidationOk returns a tuple with the CheckMxValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckMxValidation

`func (o *GetCollectionServers200ResponseDataInner) SetCheckMxValidation(v string)`

SetCheckMxValidation sets CheckMxValidation field to given value.

### HasCheckMxValidation

`func (o *GetCollectionServers200ResponseDataInner) HasCheckMxValidation() bool`

HasCheckMxValidation returns a boolean if a field has been set.

### GetCheckNamesValidation

`func (o *GetCollectionServers200ResponseDataInner) GetCheckNamesValidation() string`

GetCheckNamesValidation returns the CheckNamesValidation field if non-nil, zero value otherwise.

### GetCheckNamesValidationOk

`func (o *GetCollectionServers200ResponseDataInner) GetCheckNamesValidationOk() (*string, bool)`

GetCheckNamesValidationOk returns a tuple with the CheckNamesValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckNamesValidation

`func (o *GetCollectionServers200ResponseDataInner) SetCheckNamesValidation(v string)`

SetCheckNamesValidation sets CheckNamesValidation field to given value.

### HasCheckNamesValidation

`func (o *GetCollectionServers200ResponseDataInner) HasCheckNamesValidation() bool`

HasCheckNamesValidation returns a boolean if a field has been set.

### GetCheckNsValidation

`func (o *GetCollectionServers200ResponseDataInner) GetCheckNsValidation() string`

GetCheckNsValidation returns the CheckNsValidation field if non-nil, zero value otherwise.

### GetCheckNsValidationOk

`func (o *GetCollectionServers200ResponseDataInner) GetCheckNsValidationOk() (*string, bool)`

GetCheckNsValidationOk returns a tuple with the CheckNsValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckNsValidation

`func (o *GetCollectionServers200ResponseDataInner) SetCheckNsValidation(v string)`

SetCheckNsValidation sets CheckNsValidation field to given value.

### HasCheckNsValidation

`func (o *GetCollectionServers200ResponseDataInner) HasCheckNsValidation() bool`

HasCheckNsValidation returns a boolean if a field has been set.

### GetCheckSrvCnameValidation

`func (o *GetCollectionServers200ResponseDataInner) GetCheckSrvCnameValidation() string`

GetCheckSrvCnameValidation returns the CheckSrvCnameValidation field if non-nil, zero value otherwise.

### GetCheckSrvCnameValidationOk

`func (o *GetCollectionServers200ResponseDataInner) GetCheckSrvCnameValidationOk() (*string, bool)`

GetCheckSrvCnameValidationOk returns a tuple with the CheckSrvCnameValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckSrvCnameValidation

`func (o *GetCollectionServers200ResponseDataInner) SetCheckSrvCnameValidation(v string)`

SetCheckSrvCnameValidation sets CheckSrvCnameValidation field to given value.

### HasCheckSrvCnameValidation

`func (o *GetCollectionServers200ResponseDataInner) HasCheckSrvCnameValidation() bool`

HasCheckSrvCnameValidation returns a boolean if a field has been set.

### GetCheckWildcardValidation

`func (o *GetCollectionServers200ResponseDataInner) GetCheckWildcardValidation() string`

GetCheckWildcardValidation returns the CheckWildcardValidation field if non-nil, zero value otherwise.

### GetCheckWildcardValidationOk

`func (o *GetCollectionServers200ResponseDataInner) GetCheckWildcardValidationOk() (*string, bool)`

GetCheckWildcardValidationOk returns a tuple with the CheckWildcardValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckWildcardValidation

`func (o *GetCollectionServers200ResponseDataInner) SetCheckWildcardValidation(v string)`

SetCheckWildcardValidation sets CheckWildcardValidation field to given value.

### HasCheckWildcardValidation

`func (o *GetCollectionServers200ResponseDataInner) HasCheckWildcardValidation() bool`

HasCheckWildcardValidation returns a boolean if a field has been set.

### GetPrivateAddress

`func (o *GetCollectionServers200ResponseDataInner) GetPrivateAddress() string`

GetPrivateAddress returns the PrivateAddress field if non-nil, zero value otherwise.

### GetPrivateAddressOk

`func (o *GetCollectionServers200ResponseDataInner) GetPrivateAddressOk() (*string, bool)`

GetPrivateAddressOk returns a tuple with the PrivateAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateAddress

`func (o *GetCollectionServers200ResponseDataInner) SetPrivateAddress(v string)`

SetPrivateAddress sets PrivateAddress field to given value.

### HasPrivateAddress

`func (o *GetCollectionServers200ResponseDataInner) HasPrivateAddress() bool`

HasPrivateAddress returns a boolean if a field has been set.

### GetEncryptedNotificationsEnabled

`func (o *GetCollectionServers200ResponseDataInner) GetEncryptedNotificationsEnabled() bool`

GetEncryptedNotificationsEnabled returns the EncryptedNotificationsEnabled field if non-nil, zero value otherwise.

### GetEncryptedNotificationsEnabledOk

`func (o *GetCollectionServers200ResponseDataInner) GetEncryptedNotificationsEnabledOk() (*bool, bool)`

GetEncryptedNotificationsEnabledOk returns a tuple with the EncryptedNotificationsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedNotificationsEnabled

`func (o *GetCollectionServers200ResponseDataInner) SetEncryptedNotificationsEnabled(v bool)`

SetEncryptedNotificationsEnabled sets EncryptedNotificationsEnabled field to given value.

### HasEncryptedNotificationsEnabled

`func (o *GetCollectionServers200ResponseDataInner) HasEncryptedNotificationsEnabled() bool`

HasEncryptedNotificationsEnabled returns a boolean if a field has been set.

### GetManagementUrl

`func (o *GetCollectionServers200ResponseDataInner) GetManagementUrl() string`

GetManagementUrl returns the ManagementUrl field if non-nil, zero value otherwise.

### GetManagementUrlOk

`func (o *GetCollectionServers200ResponseDataInner) GetManagementUrlOk() (*string, bool)`

GetManagementUrlOk returns a tuple with the ManagementUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementUrl

`func (o *GetCollectionServers200ResponseDataInner) SetManagementUrl(v string)`

SetManagementUrl sets ManagementUrl field to given value.

### HasManagementUrl

`func (o *GetCollectionServers200ResponseDataInner) HasManagementUrl() bool`

HasManagementUrl returns a boolean if a field has been set.

### GetSelfIpAddress

`func (o *GetCollectionServers200ResponseDataInner) GetSelfIpAddress() string`

GetSelfIpAddress returns the SelfIpAddress field if non-nil, zero value otherwise.

### GetSelfIpAddressOk

`func (o *GetCollectionServers200ResponseDataInner) GetSelfIpAddressOk() (*string, bool)`

GetSelfIpAddressOk returns a tuple with the SelfIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfIpAddress

`func (o *GetCollectionServers200ResponseDataInner) SetSelfIpAddress(v string)`

SetSelfIpAddress sets SelfIpAddress field to given value.

### HasSelfIpAddress

`func (o *GetCollectionServers200ResponseDataInner) HasSelfIpAddress() bool`

HasSelfIpAddress returns a boolean if a field has been set.

### GetHaBackboneEnabled

`func (o *GetCollectionServers200ResponseDataInner) GetHaBackboneEnabled() bool`

GetHaBackboneEnabled returns the HaBackboneEnabled field if non-nil, zero value otherwise.

### GetHaBackboneEnabledOk

`func (o *GetCollectionServers200ResponseDataInner) GetHaBackboneEnabledOk() (*bool, bool)`

GetHaBackboneEnabledOk returns a tuple with the HaBackboneEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaBackboneEnabled

`func (o *GetCollectionServers200ResponseDataInner) SetHaBackboneEnabled(v bool)`

SetHaBackboneEnabled sets HaBackboneEnabled field to given value.

### HasHaBackboneEnabled

`func (o *GetCollectionServers200ResponseDataInner) HasHaBackboneEnabled() bool`

HasHaBackboneEnabled returns a boolean if a field has been set.

### GetHaPingAddress

`func (o *GetCollectionServers200ResponseDataInner) GetHaPingAddress() string`

GetHaPingAddress returns the HaPingAddress field if non-nil, zero value otherwise.

### GetHaPingAddressOk

`func (o *GetCollectionServers200ResponseDataInner) GetHaPingAddressOk() (*string, bool)`

GetHaPingAddressOk returns a tuple with the HaPingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaPingAddress

`func (o *GetCollectionServers200ResponseDataInner) SetHaPingAddress(v string)`

SetHaPingAddress sets HaPingAddress field to given value.

### HasHaPingAddress

`func (o *GetCollectionServers200ResponseDataInner) HasHaPingAddress() bool`

HasHaPingAddress returns a boolean if a field has been set.

### GetDhcpServicePrincipal

`func (o *GetCollectionServers200ResponseDataInner) GetDhcpServicePrincipal() InlinedKerberosServicePrincipal`

GetDhcpServicePrincipal returns the DhcpServicePrincipal field if non-nil, zero value otherwise.

### GetDhcpServicePrincipalOk

`func (o *GetCollectionServers200ResponseDataInner) GetDhcpServicePrincipalOk() (*InlinedKerberosServicePrincipal, bool)`

GetDhcpServicePrincipalOk returns a tuple with the DhcpServicePrincipal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpServicePrincipal

`func (o *GetCollectionServers200ResponseDataInner) SetDhcpServicePrincipal(v InlinedKerberosServicePrincipal)`

SetDhcpServicePrincipal sets DhcpServicePrincipal field to given value.

### HasDhcpServicePrincipal

`func (o *GetCollectionServers200ResponseDataInner) HasDhcpServicePrincipal() bool`

HasDhcpServicePrincipal returns a boolean if a field has been set.

### GetDnsServicePrincipal

`func (o *GetCollectionServers200ResponseDataInner) GetDnsServicePrincipal() InlinedKerberosServicePrincipal`

GetDnsServicePrincipal returns the DnsServicePrincipal field if non-nil, zero value otherwise.

### GetDnsServicePrincipalOk

`func (o *GetCollectionServers200ResponseDataInner) GetDnsServicePrincipalOk() (*InlinedKerberosServicePrincipal, bool)`

GetDnsServicePrincipalOk returns a tuple with the DnsServicePrincipal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServicePrincipal

`func (o *GetCollectionServers200ResponseDataInner) SetDnsServicePrincipal(v InlinedKerberosServicePrincipal)`

SetDnsServicePrincipal sets DnsServicePrincipal field to given value.

### HasDnsServicePrincipal

`func (o *GetCollectionServers200ResponseDataInner) HasDnsServicePrincipal() bool`

HasDnsServicePrincipal returns a boolean if a field has been set.

### GetInheritedFields

`func (o *GetCollectionServers200ResponseDataInner) GetInheritedFields() []string`

GetInheritedFields returns the InheritedFields field if non-nil, zero value otherwise.

### GetInheritedFieldsOk

`func (o *GetCollectionServers200ResponseDataInner) GetInheritedFieldsOk() (*[]string, bool)`

GetInheritedFieldsOk returns a tuple with the InheritedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritedFields

`func (o *GetCollectionServers200ResponseDataInner) SetInheritedFields(v []string)`

SetInheritedFields sets InheritedFields field to given value.

### HasInheritedFields

`func (o *GetCollectionServers200ResponseDataInner) HasInheritedFields() bool`

HasInheritedFields returns a boolean if a field has been set.

### GetActiveServer

`func (o *GetCollectionServers200ResponseDataInner) GetActiveServer() InlinedServer`

GetActiveServer returns the ActiveServer field if non-nil, zero value otherwise.

### GetActiveServerOk

`func (o *GetCollectionServers200ResponseDataInner) GetActiveServerOk() (*InlinedServer, bool)`

GetActiveServerOk returns a tuple with the ActiveServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveServer

`func (o *GetCollectionServers200ResponseDataInner) SetActiveServer(v InlinedServer)`

SetActiveServer sets ActiveServer field to given value.

### HasActiveServer

`func (o *GetCollectionServers200ResponseDataInner) HasActiveServer() bool`

HasActiveServer returns a boolean if a field has been set.

### GetPassiveServer

`func (o *GetCollectionServers200ResponseDataInner) GetPassiveServer() InlinedServer`

GetPassiveServer returns the PassiveServer field if non-nil, zero value otherwise.

### GetPassiveServerOk

`func (o *GetCollectionServers200ResponseDataInner) GetPassiveServerOk() (*InlinedServer, bool)`

GetPassiveServerOk returns a tuple with the PassiveServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassiveServer

`func (o *GetCollectionServers200ResponseDataInner) SetPassiveServer(v InlinedServer)`

SetPassiveServer sets PassiveServer field to given value.

### HasPassiveServer

`func (o *GetCollectionServers200ResponseDataInner) HasPassiveServer() bool`

HasPassiveServer returns a boolean if a field has been set.

### GetDedicatedManagementEnabled

`func (o *GetCollectionServers200ResponseDataInner) GetDedicatedManagementEnabled() bool`

GetDedicatedManagementEnabled returns the DedicatedManagementEnabled field if non-nil, zero value otherwise.

### GetDedicatedManagementEnabledOk

`func (o *GetCollectionServers200ResponseDataInner) GetDedicatedManagementEnabledOk() (*bool, bool)`

GetDedicatedManagementEnabledOk returns a tuple with the DedicatedManagementEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedicatedManagementEnabled

`func (o *GetCollectionServers200ResponseDataInner) SetDedicatedManagementEnabled(v bool)`

SetDedicatedManagementEnabled sets DedicatedManagementEnabled field to given value.

### HasDedicatedManagementEnabled

`func (o *GetCollectionServers200ResponseDataInner) HasDedicatedManagementEnabled() bool`

HasDedicatedManagementEnabled returns a boolean if a field has been set.

### GetHaRole

`func (o *GetCollectionServers200ResponseDataInner) GetHaRole() string`

GetHaRole returns the HaRole field if non-nil, zero value otherwise.

### GetHaRoleOk

`func (o *GetCollectionServers200ResponseDataInner) GetHaRoleOk() (*string, bool)`

GetHaRoleOk returns a tuple with the HaRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaRole

`func (o *GetCollectionServers200ResponseDataInner) SetHaRole(v string)`

SetHaRole sets HaRole field to given value.

### HasHaRole

`func (o *GetCollectionServers200ResponseDataInner) HasHaRole() bool`

HasHaRole returns a boolean if a field has been set.

### GetHaPeerConnectionState

`func (o *GetCollectionServers200ResponseDataInner) GetHaPeerConnectionState() string`

GetHaPeerConnectionState returns the HaPeerConnectionState field if non-nil, zero value otherwise.

### GetHaPeerConnectionStateOk

`func (o *GetCollectionServers200ResponseDataInner) GetHaPeerConnectionStateOk() (*string, bool)`

GetHaPeerConnectionStateOk returns a tuple with the HaPeerConnectionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaPeerConnectionState

`func (o *GetCollectionServers200ResponseDataInner) SetHaPeerConnectionState(v string)`

SetHaPeerConnectionState sets HaPeerConnectionState field to given value.

### HasHaPeerConnectionState

`func (o *GetCollectionServers200ResponseDataInner) HasHaPeerConnectionState() bool`

HasHaPeerConnectionState returns a boolean if a field has been set.

### GetHaDiskState

`func (o *GetCollectionServers200ResponseDataInner) GetHaDiskState() string`

GetHaDiskState returns the HaDiskState field if non-nil, zero value otherwise.

### GetHaDiskStateOk

`func (o *GetCollectionServers200ResponseDataInner) GetHaDiskStateOk() (*string, bool)`

GetHaDiskStateOk returns a tuple with the HaDiskState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaDiskState

`func (o *GetCollectionServers200ResponseDataInner) SetHaDiskState(v string)`

SetHaDiskState sets HaDiskState field to given value.

### HasHaDiskState

`func (o *GetCollectionServers200ResponseDataInner) HasHaDiskState() bool`

HasHaDiskState returns a boolean if a field has been set.

### GetHsmSupportEnabled

`func (o *GetCollectionServers200ResponseDataInner) GetHsmSupportEnabled() bool`

GetHsmSupportEnabled returns the HsmSupportEnabled field if non-nil, zero value otherwise.

### GetHsmSupportEnabledOk

`func (o *GetCollectionServers200ResponseDataInner) GetHsmSupportEnabledOk() (*bool, bool)`

GetHsmSupportEnabledOk returns a tuple with the HsmSupportEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHsmSupportEnabled

`func (o *GetCollectionServers200ResponseDataInner) SetHsmSupportEnabled(v bool)`

SetHsmSupportEnabled sets HsmSupportEnabled field to given value.

### HasHsmSupportEnabled

`func (o *GetCollectionServers200ResponseDataInner) HasHsmSupportEnabled() bool`

HasHsmSupportEnabled returns a boolean if a field has been set.

### GetInterfaceRedundancyEnabled

`func (o *GetCollectionServers200ResponseDataInner) GetInterfaceRedundancyEnabled() bool`

GetInterfaceRedundancyEnabled returns the InterfaceRedundancyEnabled field if non-nil, zero value otherwise.

### GetInterfaceRedundancyEnabledOk

`func (o *GetCollectionServers200ResponseDataInner) GetInterfaceRedundancyEnabledOk() (*bool, bool)`

GetInterfaceRedundancyEnabledOk returns a tuple with the InterfaceRedundancyEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceRedundancyEnabled

`func (o *GetCollectionServers200ResponseDataInner) SetInterfaceRedundancyEnabled(v bool)`

SetInterfaceRedundancyEnabled sets InterfaceRedundancyEnabled field to given value.

### HasInterfaceRedundancyEnabled

`func (o *GetCollectionServers200ResponseDataInner) HasInterfaceRedundancyEnabled() bool`

HasInterfaceRedundancyEnabled returns a boolean if a field has been set.

### GetLinks

`func (o *GetCollectionServers200ResponseDataInner) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetCollectionServers200ResponseDataInner) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetCollectionServers200ResponseDataInner) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetCollectionServers200ResponseDataInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

