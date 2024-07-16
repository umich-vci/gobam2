# ServerPostRequestBody

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
**Interfaces** | [**[]NetworkInterface**](NetworkInterface.md) | The list of network interfaces of the server. | 
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

### NewServerPostRequestBody

`func NewServerPostRequestBody(type_ string, name string, profile string, interfaces []NetworkInterface, ) *ServerPostRequestBody`

NewServerPostRequestBody instantiates a new ServerPostRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerPostRequestBodyWithDefaults

`func NewServerPostRequestBodyWithDefaults() *ServerPostRequestBody`

NewServerPostRequestBodyWithDefaults instantiates a new ServerPostRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServerPostRequestBody) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServerPostRequestBody) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServerPostRequestBody) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ServerPostRequestBody) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ServerPostRequestBody) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServerPostRequestBody) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServerPostRequestBody) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *ServerPostRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServerPostRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServerPostRequestBody) SetName(v string)`

SetName sets Name field to given value.


### GetUserDefinedFields

`func (o *ServerPostRequestBody) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *ServerPostRequestBody) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *ServerPostRequestBody) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *ServerPostRequestBody) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetConfiguration

`func (o *ServerPostRequestBody) GetConfiguration() InlinedConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *ServerPostRequestBody) GetConfigurationOk() (*InlinedConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *ServerPostRequestBody) SetConfiguration(v InlinedConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *ServerPostRequestBody) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetProfile

`func (o *ServerPostRequestBody) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *ServerPostRequestBody) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *ServerPostRequestBody) SetProfile(v string)`

SetProfile sets Profile field to given value.


### GetState

`func (o *ServerPostRequestBody) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ServerPostRequestBody) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ServerPostRequestBody) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ServerPostRequestBody) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUsername

`func (o *ServerPostRequestBody) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ServerPostRequestBody) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ServerPostRequestBody) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ServerPostRequestBody) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *ServerPostRequestBody) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ServerPostRequestBody) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ServerPostRequestBody) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ServerPostRequestBody) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetServerGroup

`func (o *ServerPostRequestBody) GetServerGroup() InlinedServerGroup`

GetServerGroup returns the ServerGroup field if non-nil, zero value otherwise.

### GetServerGroupOk

`func (o *ServerPostRequestBody) GetServerGroupOk() (*InlinedServerGroup, bool)`

GetServerGroupOk returns a tuple with the ServerGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerGroup

`func (o *ServerPostRequestBody) SetServerGroup(v InlinedServerGroup)`

SetServerGroup sets ServerGroup field to given value.

### HasServerGroup

`func (o *ServerPostRequestBody) HasServerGroup() bool`

HasServerGroup returns a boolean if a field has been set.

### GetConnected

`func (o *ServerPostRequestBody) GetConnected() bool`

GetConnected returns the Connected field if non-nil, zero value otherwise.

### GetConnectedOk

`func (o *ServerPostRequestBody) GetConnectedOk() (*bool, bool)`

GetConnectedOk returns a tuple with the Connected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnected

`func (o *ServerPostRequestBody) SetConnected(v bool)`

SetConnected sets Connected field to given value.

### HasConnected

`func (o *ServerPostRequestBody) HasConnected() bool`

HasConnected returns a boolean if a field has been set.

### GetMonitoringEnabled

`func (o *ServerPostRequestBody) GetMonitoringEnabled() bool`

GetMonitoringEnabled returns the MonitoringEnabled field if non-nil, zero value otherwise.

### GetMonitoringEnabledOk

`func (o *ServerPostRequestBody) GetMonitoringEnabledOk() (*bool, bool)`

GetMonitoringEnabledOk returns a tuple with the MonitoringEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringEnabled

`func (o *ServerPostRequestBody) SetMonitoringEnabled(v bool)`

SetMonitoringEnabled sets MonitoringEnabled field to given value.

### HasMonitoringEnabled

`func (o *ServerPostRequestBody) HasMonitoringEnabled() bool`

HasMonitoringEnabled returns a boolean if a field has been set.

### GetSnmp

`func (o *ServerPostRequestBody) GetSnmp() SnmpBean`

GetSnmp returns the Snmp field if non-nil, zero value otherwise.

### GetSnmpOk

`func (o *ServerPostRequestBody) GetSnmpOk() (*SnmpBean, bool)`

GetSnmpOk returns a tuple with the Snmp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmp

`func (o *ServerPostRequestBody) SetSnmp(v SnmpBean)`

SetSnmp sets Snmp field to given value.

### HasSnmp

`func (o *ServerPostRequestBody) HasSnmp() bool`

HasSnmp returns a boolean if a field has been set.

### GetInterfaces

`func (o *ServerPostRequestBody) GetInterfaces() []NetworkInterface`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *ServerPostRequestBody) GetInterfacesOk() (*[]NetworkInterface, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *ServerPostRequestBody) SetInterfaces(v []NetworkInterface)`

SetInterfaces sets Interfaces field to given value.


### GetLocation

`func (o *ServerPostRequestBody) GetLocation() InlinedLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ServerPostRequestBody) GetLocationOk() (*InlinedLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ServerPostRequestBody) SetLocation(v InlinedLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ServerPostRequestBody) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetStatistics

`func (o *ServerPostRequestBody) GetStatistics() map[string]map[string]interface{}`

GetStatistics returns the Statistics field if non-nil, zero value otherwise.

### GetStatisticsOk

`func (o *ServerPostRequestBody) GetStatisticsOk() (*map[string]map[string]interface{}, bool)`

GetStatisticsOk returns a tuple with the Statistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatistics

`func (o *ServerPostRequestBody) SetStatistics(v map[string]map[string]interface{})`

SetStatistics sets Statistics field to given value.

### HasStatistics

`func (o *ServerPostRequestBody) HasStatistics() bool`

HasStatistics returns a boolean if a field has been set.

### GetOverrideConfigurationLevelDhcpValidation

`func (o *ServerPostRequestBody) GetOverrideConfigurationLevelDhcpValidation() bool`

GetOverrideConfigurationLevelDhcpValidation returns the OverrideConfigurationLevelDhcpValidation field if non-nil, zero value otherwise.

### GetOverrideConfigurationLevelDhcpValidationOk

`func (o *ServerPostRequestBody) GetOverrideConfigurationLevelDhcpValidationOk() (*bool, bool)`

GetOverrideConfigurationLevelDhcpValidationOk returns a tuple with the OverrideConfigurationLevelDhcpValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideConfigurationLevelDhcpValidation

`func (o *ServerPostRequestBody) SetOverrideConfigurationLevelDhcpValidation(v bool)`

SetOverrideConfigurationLevelDhcpValidation sets OverrideConfigurationLevelDhcpValidation field to given value.

### HasOverrideConfigurationLevelDhcpValidation

`func (o *ServerPostRequestBody) HasOverrideConfigurationLevelDhcpValidation() bool`

HasOverrideConfigurationLevelDhcpValidation returns a boolean if a field has been set.

### GetOverrideConfigurationLevelDnsValidation

`func (o *ServerPostRequestBody) GetOverrideConfigurationLevelDnsValidation() bool`

GetOverrideConfigurationLevelDnsValidation returns the OverrideConfigurationLevelDnsValidation field if non-nil, zero value otherwise.

### GetOverrideConfigurationLevelDnsValidationOk

`func (o *ServerPostRequestBody) GetOverrideConfigurationLevelDnsValidationOk() (*bool, bool)`

GetOverrideConfigurationLevelDnsValidationOk returns a tuple with the OverrideConfigurationLevelDnsValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideConfigurationLevelDnsValidation

`func (o *ServerPostRequestBody) SetOverrideConfigurationLevelDnsValidation(v bool)`

SetOverrideConfigurationLevelDnsValidation sets OverrideConfigurationLevelDnsValidation field to given value.

### HasOverrideConfigurationLevelDnsValidation

`func (o *ServerPostRequestBody) HasOverrideConfigurationLevelDnsValidation() bool`

HasOverrideConfigurationLevelDnsValidation returns a boolean if a field has been set.

### GetDhcpConfigurationValidationEnabled

`func (o *ServerPostRequestBody) GetDhcpConfigurationValidationEnabled() bool`

GetDhcpConfigurationValidationEnabled returns the DhcpConfigurationValidationEnabled field if non-nil, zero value otherwise.

### GetDhcpConfigurationValidationEnabledOk

`func (o *ServerPostRequestBody) GetDhcpConfigurationValidationEnabledOk() (*bool, bool)`

GetDhcpConfigurationValidationEnabledOk returns a tuple with the DhcpConfigurationValidationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpConfigurationValidationEnabled

`func (o *ServerPostRequestBody) SetDhcpConfigurationValidationEnabled(v bool)`

SetDhcpConfigurationValidationEnabled sets DhcpConfigurationValidationEnabled field to given value.

### HasDhcpConfigurationValidationEnabled

`func (o *ServerPostRequestBody) HasDhcpConfigurationValidationEnabled() bool`

HasDhcpConfigurationValidationEnabled returns a boolean if a field has been set.

### GetDnsConfigurationValidationEnabled

`func (o *ServerPostRequestBody) GetDnsConfigurationValidationEnabled() bool`

GetDnsConfigurationValidationEnabled returns the DnsConfigurationValidationEnabled field if non-nil, zero value otherwise.

### GetDnsConfigurationValidationEnabledOk

`func (o *ServerPostRequestBody) GetDnsConfigurationValidationEnabledOk() (*bool, bool)`

GetDnsConfigurationValidationEnabledOk returns a tuple with the DnsConfigurationValidationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsConfigurationValidationEnabled

`func (o *ServerPostRequestBody) SetDnsConfigurationValidationEnabled(v bool)`

SetDnsConfigurationValidationEnabled sets DnsConfigurationValidationEnabled field to given value.

### HasDnsConfigurationValidationEnabled

`func (o *ServerPostRequestBody) HasDnsConfigurationValidationEnabled() bool`

HasDnsConfigurationValidationEnabled returns a boolean if a field has been set.

### GetDnsZoneValidationEnabled

`func (o *ServerPostRequestBody) GetDnsZoneValidationEnabled() bool`

GetDnsZoneValidationEnabled returns the DnsZoneValidationEnabled field if non-nil, zero value otherwise.

### GetDnsZoneValidationEnabledOk

`func (o *ServerPostRequestBody) GetDnsZoneValidationEnabledOk() (*bool, bool)`

GetDnsZoneValidationEnabledOk returns a tuple with the DnsZoneValidationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsZoneValidationEnabled

`func (o *ServerPostRequestBody) SetDnsZoneValidationEnabled(v bool)`

SetDnsZoneValidationEnabled sets DnsZoneValidationEnabled field to given value.

### HasDnsZoneValidationEnabled

`func (o *ServerPostRequestBody) HasDnsZoneValidationEnabled() bool`

HasDnsZoneValidationEnabled returns a boolean if a field has been set.

### GetCheckIntegrityValidation

`func (o *ServerPostRequestBody) GetCheckIntegrityValidation() string`

GetCheckIntegrityValidation returns the CheckIntegrityValidation field if non-nil, zero value otherwise.

### GetCheckIntegrityValidationOk

`func (o *ServerPostRequestBody) GetCheckIntegrityValidationOk() (*string, bool)`

GetCheckIntegrityValidationOk returns a tuple with the CheckIntegrityValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckIntegrityValidation

`func (o *ServerPostRequestBody) SetCheckIntegrityValidation(v string)`

SetCheckIntegrityValidation sets CheckIntegrityValidation field to given value.

### HasCheckIntegrityValidation

`func (o *ServerPostRequestBody) HasCheckIntegrityValidation() bool`

HasCheckIntegrityValidation returns a boolean if a field has been set.

### GetCheckMxCnameValidation

`func (o *ServerPostRequestBody) GetCheckMxCnameValidation() string`

GetCheckMxCnameValidation returns the CheckMxCnameValidation field if non-nil, zero value otherwise.

### GetCheckMxCnameValidationOk

`func (o *ServerPostRequestBody) GetCheckMxCnameValidationOk() (*string, bool)`

GetCheckMxCnameValidationOk returns a tuple with the CheckMxCnameValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckMxCnameValidation

`func (o *ServerPostRequestBody) SetCheckMxCnameValidation(v string)`

SetCheckMxCnameValidation sets CheckMxCnameValidation field to given value.

### HasCheckMxCnameValidation

`func (o *ServerPostRequestBody) HasCheckMxCnameValidation() bool`

HasCheckMxCnameValidation returns a boolean if a field has been set.

### GetCheckMxValidation

`func (o *ServerPostRequestBody) GetCheckMxValidation() string`

GetCheckMxValidation returns the CheckMxValidation field if non-nil, zero value otherwise.

### GetCheckMxValidationOk

`func (o *ServerPostRequestBody) GetCheckMxValidationOk() (*string, bool)`

GetCheckMxValidationOk returns a tuple with the CheckMxValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckMxValidation

`func (o *ServerPostRequestBody) SetCheckMxValidation(v string)`

SetCheckMxValidation sets CheckMxValidation field to given value.

### HasCheckMxValidation

`func (o *ServerPostRequestBody) HasCheckMxValidation() bool`

HasCheckMxValidation returns a boolean if a field has been set.

### GetCheckNamesValidation

`func (o *ServerPostRequestBody) GetCheckNamesValidation() string`

GetCheckNamesValidation returns the CheckNamesValidation field if non-nil, zero value otherwise.

### GetCheckNamesValidationOk

`func (o *ServerPostRequestBody) GetCheckNamesValidationOk() (*string, bool)`

GetCheckNamesValidationOk returns a tuple with the CheckNamesValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckNamesValidation

`func (o *ServerPostRequestBody) SetCheckNamesValidation(v string)`

SetCheckNamesValidation sets CheckNamesValidation field to given value.

### HasCheckNamesValidation

`func (o *ServerPostRequestBody) HasCheckNamesValidation() bool`

HasCheckNamesValidation returns a boolean if a field has been set.

### GetCheckNsValidation

`func (o *ServerPostRequestBody) GetCheckNsValidation() string`

GetCheckNsValidation returns the CheckNsValidation field if non-nil, zero value otherwise.

### GetCheckNsValidationOk

`func (o *ServerPostRequestBody) GetCheckNsValidationOk() (*string, bool)`

GetCheckNsValidationOk returns a tuple with the CheckNsValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckNsValidation

`func (o *ServerPostRequestBody) SetCheckNsValidation(v string)`

SetCheckNsValidation sets CheckNsValidation field to given value.

### HasCheckNsValidation

`func (o *ServerPostRequestBody) HasCheckNsValidation() bool`

HasCheckNsValidation returns a boolean if a field has been set.

### GetCheckSrvCnameValidation

`func (o *ServerPostRequestBody) GetCheckSrvCnameValidation() string`

GetCheckSrvCnameValidation returns the CheckSrvCnameValidation field if non-nil, zero value otherwise.

### GetCheckSrvCnameValidationOk

`func (o *ServerPostRequestBody) GetCheckSrvCnameValidationOk() (*string, bool)`

GetCheckSrvCnameValidationOk returns a tuple with the CheckSrvCnameValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckSrvCnameValidation

`func (o *ServerPostRequestBody) SetCheckSrvCnameValidation(v string)`

SetCheckSrvCnameValidation sets CheckSrvCnameValidation field to given value.

### HasCheckSrvCnameValidation

`func (o *ServerPostRequestBody) HasCheckSrvCnameValidation() bool`

HasCheckSrvCnameValidation returns a boolean if a field has been set.

### GetCheckWildcardValidation

`func (o *ServerPostRequestBody) GetCheckWildcardValidation() string`

GetCheckWildcardValidation returns the CheckWildcardValidation field if non-nil, zero value otherwise.

### GetCheckWildcardValidationOk

`func (o *ServerPostRequestBody) GetCheckWildcardValidationOk() (*string, bool)`

GetCheckWildcardValidationOk returns a tuple with the CheckWildcardValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckWildcardValidation

`func (o *ServerPostRequestBody) SetCheckWildcardValidation(v string)`

SetCheckWildcardValidation sets CheckWildcardValidation field to given value.

### HasCheckWildcardValidation

`func (o *ServerPostRequestBody) HasCheckWildcardValidation() bool`

HasCheckWildcardValidation returns a boolean if a field has been set.

### GetPrivateAddress

`func (o *ServerPostRequestBody) GetPrivateAddress() string`

GetPrivateAddress returns the PrivateAddress field if non-nil, zero value otherwise.

### GetPrivateAddressOk

`func (o *ServerPostRequestBody) GetPrivateAddressOk() (*string, bool)`

GetPrivateAddressOk returns a tuple with the PrivateAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateAddress

`func (o *ServerPostRequestBody) SetPrivateAddress(v string)`

SetPrivateAddress sets PrivateAddress field to given value.

### HasPrivateAddress

`func (o *ServerPostRequestBody) HasPrivateAddress() bool`

HasPrivateAddress returns a boolean if a field has been set.

### GetEncryptedNotificationsEnabled

`func (o *ServerPostRequestBody) GetEncryptedNotificationsEnabled() bool`

GetEncryptedNotificationsEnabled returns the EncryptedNotificationsEnabled field if non-nil, zero value otherwise.

### GetEncryptedNotificationsEnabledOk

`func (o *ServerPostRequestBody) GetEncryptedNotificationsEnabledOk() (*bool, bool)`

GetEncryptedNotificationsEnabledOk returns a tuple with the EncryptedNotificationsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedNotificationsEnabled

`func (o *ServerPostRequestBody) SetEncryptedNotificationsEnabled(v bool)`

SetEncryptedNotificationsEnabled sets EncryptedNotificationsEnabled field to given value.

### HasEncryptedNotificationsEnabled

`func (o *ServerPostRequestBody) HasEncryptedNotificationsEnabled() bool`

HasEncryptedNotificationsEnabled returns a boolean if a field has been set.

### GetManagementUrl

`func (o *ServerPostRequestBody) GetManagementUrl() string`

GetManagementUrl returns the ManagementUrl field if non-nil, zero value otherwise.

### GetManagementUrlOk

`func (o *ServerPostRequestBody) GetManagementUrlOk() (*string, bool)`

GetManagementUrlOk returns a tuple with the ManagementUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementUrl

`func (o *ServerPostRequestBody) SetManagementUrl(v string)`

SetManagementUrl sets ManagementUrl field to given value.

### HasManagementUrl

`func (o *ServerPostRequestBody) HasManagementUrl() bool`

HasManagementUrl returns a boolean if a field has been set.

### GetSelfIpAddress

`func (o *ServerPostRequestBody) GetSelfIpAddress() string`

GetSelfIpAddress returns the SelfIpAddress field if non-nil, zero value otherwise.

### GetSelfIpAddressOk

`func (o *ServerPostRequestBody) GetSelfIpAddressOk() (*string, bool)`

GetSelfIpAddressOk returns a tuple with the SelfIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfIpAddress

`func (o *ServerPostRequestBody) SetSelfIpAddress(v string)`

SetSelfIpAddress sets SelfIpAddress field to given value.

### HasSelfIpAddress

`func (o *ServerPostRequestBody) HasSelfIpAddress() bool`

HasSelfIpAddress returns a boolean if a field has been set.

### GetHaBackboneEnabled

`func (o *ServerPostRequestBody) GetHaBackboneEnabled() bool`

GetHaBackboneEnabled returns the HaBackboneEnabled field if non-nil, zero value otherwise.

### GetHaBackboneEnabledOk

`func (o *ServerPostRequestBody) GetHaBackboneEnabledOk() (*bool, bool)`

GetHaBackboneEnabledOk returns a tuple with the HaBackboneEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaBackboneEnabled

`func (o *ServerPostRequestBody) SetHaBackboneEnabled(v bool)`

SetHaBackboneEnabled sets HaBackboneEnabled field to given value.

### HasHaBackboneEnabled

`func (o *ServerPostRequestBody) HasHaBackboneEnabled() bool`

HasHaBackboneEnabled returns a boolean if a field has been set.

### GetHaPingAddress

`func (o *ServerPostRequestBody) GetHaPingAddress() string`

GetHaPingAddress returns the HaPingAddress field if non-nil, zero value otherwise.

### GetHaPingAddressOk

`func (o *ServerPostRequestBody) GetHaPingAddressOk() (*string, bool)`

GetHaPingAddressOk returns a tuple with the HaPingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaPingAddress

`func (o *ServerPostRequestBody) SetHaPingAddress(v string)`

SetHaPingAddress sets HaPingAddress field to given value.

### HasHaPingAddress

`func (o *ServerPostRequestBody) HasHaPingAddress() bool`

HasHaPingAddress returns a boolean if a field has been set.

### GetDhcpServicePrincipal

`func (o *ServerPostRequestBody) GetDhcpServicePrincipal() InlinedKerberosServicePrincipal`

GetDhcpServicePrincipal returns the DhcpServicePrincipal field if non-nil, zero value otherwise.

### GetDhcpServicePrincipalOk

`func (o *ServerPostRequestBody) GetDhcpServicePrincipalOk() (*InlinedKerberosServicePrincipal, bool)`

GetDhcpServicePrincipalOk returns a tuple with the DhcpServicePrincipal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpServicePrincipal

`func (o *ServerPostRequestBody) SetDhcpServicePrincipal(v InlinedKerberosServicePrincipal)`

SetDhcpServicePrincipal sets DhcpServicePrincipal field to given value.

### HasDhcpServicePrincipal

`func (o *ServerPostRequestBody) HasDhcpServicePrincipal() bool`

HasDhcpServicePrincipal returns a boolean if a field has been set.

### GetDnsServicePrincipal

`func (o *ServerPostRequestBody) GetDnsServicePrincipal() InlinedKerberosServicePrincipal`

GetDnsServicePrincipal returns the DnsServicePrincipal field if non-nil, zero value otherwise.

### GetDnsServicePrincipalOk

`func (o *ServerPostRequestBody) GetDnsServicePrincipalOk() (*InlinedKerberosServicePrincipal, bool)`

GetDnsServicePrincipalOk returns a tuple with the DnsServicePrincipal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServicePrincipal

`func (o *ServerPostRequestBody) SetDnsServicePrincipal(v InlinedKerberosServicePrincipal)`

SetDnsServicePrincipal sets DnsServicePrincipal field to given value.

### HasDnsServicePrincipal

`func (o *ServerPostRequestBody) HasDnsServicePrincipal() bool`

HasDnsServicePrincipal returns a boolean if a field has been set.

### GetDedicatedManagementEnabled

`func (o *ServerPostRequestBody) GetDedicatedManagementEnabled() bool`

GetDedicatedManagementEnabled returns the DedicatedManagementEnabled field if non-nil, zero value otherwise.

### GetDedicatedManagementEnabledOk

`func (o *ServerPostRequestBody) GetDedicatedManagementEnabledOk() (*bool, bool)`

GetDedicatedManagementEnabledOk returns a tuple with the DedicatedManagementEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedicatedManagementEnabled

`func (o *ServerPostRequestBody) SetDedicatedManagementEnabled(v bool)`

SetDedicatedManagementEnabled sets DedicatedManagementEnabled field to given value.

### HasDedicatedManagementEnabled

`func (o *ServerPostRequestBody) HasDedicatedManagementEnabled() bool`

HasDedicatedManagementEnabled returns a boolean if a field has been set.

### GetHaRole

`func (o *ServerPostRequestBody) GetHaRole() string`

GetHaRole returns the HaRole field if non-nil, zero value otherwise.

### GetHaRoleOk

`func (o *ServerPostRequestBody) GetHaRoleOk() (*string, bool)`

GetHaRoleOk returns a tuple with the HaRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaRole

`func (o *ServerPostRequestBody) SetHaRole(v string)`

SetHaRole sets HaRole field to given value.

### HasHaRole

`func (o *ServerPostRequestBody) HasHaRole() bool`

HasHaRole returns a boolean if a field has been set.

### GetHaPeerConnectionState

`func (o *ServerPostRequestBody) GetHaPeerConnectionState() string`

GetHaPeerConnectionState returns the HaPeerConnectionState field if non-nil, zero value otherwise.

### GetHaPeerConnectionStateOk

`func (o *ServerPostRequestBody) GetHaPeerConnectionStateOk() (*string, bool)`

GetHaPeerConnectionStateOk returns a tuple with the HaPeerConnectionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaPeerConnectionState

`func (o *ServerPostRequestBody) SetHaPeerConnectionState(v string)`

SetHaPeerConnectionState sets HaPeerConnectionState field to given value.

### HasHaPeerConnectionState

`func (o *ServerPostRequestBody) HasHaPeerConnectionState() bool`

HasHaPeerConnectionState returns a boolean if a field has been set.

### GetHaDiskState

`func (o *ServerPostRequestBody) GetHaDiskState() string`

GetHaDiskState returns the HaDiskState field if non-nil, zero value otherwise.

### GetHaDiskStateOk

`func (o *ServerPostRequestBody) GetHaDiskStateOk() (*string, bool)`

GetHaDiskStateOk returns a tuple with the HaDiskState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaDiskState

`func (o *ServerPostRequestBody) SetHaDiskState(v string)`

SetHaDiskState sets HaDiskState field to given value.

### HasHaDiskState

`func (o *ServerPostRequestBody) HasHaDiskState() bool`

HasHaDiskState returns a boolean if a field has been set.

### GetHsmSupportEnabled

`func (o *ServerPostRequestBody) GetHsmSupportEnabled() bool`

GetHsmSupportEnabled returns the HsmSupportEnabled field if non-nil, zero value otherwise.

### GetHsmSupportEnabledOk

`func (o *ServerPostRequestBody) GetHsmSupportEnabledOk() (*bool, bool)`

GetHsmSupportEnabledOk returns a tuple with the HsmSupportEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHsmSupportEnabled

`func (o *ServerPostRequestBody) SetHsmSupportEnabled(v bool)`

SetHsmSupportEnabled sets HsmSupportEnabled field to given value.

### HasHsmSupportEnabled

`func (o *ServerPostRequestBody) HasHsmSupportEnabled() bool`

HasHsmSupportEnabled returns a boolean if a field has been set.

### GetInterfaceRedundancyEnabled

`func (o *ServerPostRequestBody) GetInterfaceRedundancyEnabled() bool`

GetInterfaceRedundancyEnabled returns the InterfaceRedundancyEnabled field if non-nil, zero value otherwise.

### GetInterfaceRedundancyEnabledOk

`func (o *ServerPostRequestBody) GetInterfaceRedundancyEnabledOk() (*bool, bool)`

GetInterfaceRedundancyEnabledOk returns a tuple with the InterfaceRedundancyEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceRedundancyEnabled

`func (o *ServerPostRequestBody) SetInterfaceRedundancyEnabled(v bool)`

SetInterfaceRedundancyEnabled sets InterfaceRedundancyEnabled field to given value.

### HasInterfaceRedundancyEnabled

`func (o *ServerPostRequestBody) HasInterfaceRedundancyEnabled() bool`

HasInterfaceRedundancyEnabled returns a boolean if a field has been set.

### GetInheritedFields

`func (o *ServerPostRequestBody) GetInheritedFields() []string`

GetInheritedFields returns the InheritedFields field if non-nil, zero value otherwise.

### GetInheritedFieldsOk

`func (o *ServerPostRequestBody) GetInheritedFieldsOk() (*[]string, bool)`

GetInheritedFieldsOk returns a tuple with the InheritedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritedFields

`func (o *ServerPostRequestBody) SetInheritedFields(v []string)`

SetInheritedFields sets InheritedFields field to given value.

### HasInheritedFields

`func (o *ServerPostRequestBody) HasInheritedFields() bool`

HasInheritedFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


