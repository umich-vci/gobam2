# PutServerServiceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | **string** | The resource type. | 
**ServerVersion** | Pointer to **string** |  | [optional] [readonly] 
**Enabled** | **bool** | Indicates whether the SSH service is enabled. | 
**MaximumBufferedEvents** | **int64** | The maximum number of events to be stored in the memory buffer | 
**Destination** | Pointer to [**SinkBean**](SinkBean.md) |  | [optional] 
**Dhcpv4Enabled** | **bool** |  | 
**Dhcpv6Enabled** | **bool** |  | 
**Filters** | [**[]DnsActivityFilterBean**](DnsActivityFilterBean.md) |  | 
**Servers** | [**[]SyslogRedirectionServerBean**](SyslogRedirectionServerBean.md) | The list of servers to redirect syslog messages to. | 
**PollingInterval** | **string** | The SNMP polling interval, in ISO 8601 duration format, that determines the frequency that SNMP values are refreshed on the system. | 
**PingAllowed** | **bool** | Indicates whether the server responds to the ping command. | 
**State** | Pointer to **string** | The status of the Gateway service running on the DNS/DHCP Server. | [optional] 
**Container** | Pointer to [**GatewayContainerBean**](GatewayContainerBean.md) |  | [optional] 
**RepositoryUsername** | Pointer to **string** | The docker username used to pull the Gateway image from the container repository. | [optional] 
**RepositoryPassword** | Pointer to **string** | The docker password used to pull the image from a private repository. | [optional] 
**DeleteImageOnDisable** | Pointer to **bool** | Indicates whether the Gateway image is removed upon disabling the Gateway service. | [optional] 
**DeleteMountPointsOnDisable** | Pointer to **bool** | Indicates whether the local volumes and the mounted data and log directories are removed upon disabling the Gateway service. | [optional] 
**DedicatedManagementEnabled** | **bool** |  | 
**Interfaces** | [**[]InterfaceBean**](InterfaceBean.md) |  | 
**Routes** | [**[]RouteBean**](RouteBean.md) |  | 
**ClientId** | **string** |  | 
**Key** | **string** |  | 
**ManualOverrideEnabled** | Pointer to **bool** | Indicates whether the syslog redirection service configuration has been manually overridden. | [optional] [readonly] 
**ContactName** | **string** | The system name to be reported through SNMP. | 
**ContactEmail** | **string** | The email address for the system contact to be reported through SNMP. | 
**Location** | Pointer to **string** | The description of the system&#39;s location to be reported through SNMP. | [optional] 
**Description** | **string** | The description of the system to be reported through SNMP. | 
**LogLevel** | **string** | The logging level for SNMP events. | 
**Snmpv1** | [**Snmpv1Bean**](Snmpv1Bean.md) |  | 
**Snmpv2c** | [**Snmpv2cBean**](Snmpv2cBean.md) |  | 
**Snmpv3** | [**Snmpv3Bean**](Snmpv3Bean.md) |  | 
**TrapServers** | [**[]SnmpTrapServerBean**](SnmpTrapServerBean.md) |  | 
**TacacsPlusAuthentication** | [**TacacsPlusAuthenticationBean**](TacacsPlusAuthenticationBean.md) |  | 
**Iso8601TimestampsEnabled** | Pointer to **bool** | Indicates whether the ISO 8601 date time format is used when logging timestamps. | [optional] 

## Methods

### NewPutServerServiceRequest

`func NewPutServerServiceRequest(type_ string, enabled bool, maximumBufferedEvents int64, dhcpv4Enabled bool, dhcpv6Enabled bool, filters []DnsActivityFilterBean, servers []SyslogRedirectionServerBean, pollingInterval string, pingAllowed bool, dedicatedManagementEnabled bool, interfaces []InterfaceBean, routes []RouteBean, clientId string, key string, contactName string, contactEmail string, description string, logLevel string, snmpv1 Snmpv1Bean, snmpv2c Snmpv2cBean, snmpv3 Snmpv3Bean, trapServers []SnmpTrapServerBean, tacacsPlusAuthentication TacacsPlusAuthenticationBean, ) *PutServerServiceRequest`

NewPutServerServiceRequest instantiates a new PutServerServiceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutServerServiceRequestWithDefaults

`func NewPutServerServiceRequestWithDefaults() *PutServerServiceRequest`

NewPutServerServiceRequestWithDefaults instantiates a new PutServerServiceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PutServerServiceRequest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PutServerServiceRequest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PutServerServiceRequest) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *PutServerServiceRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *PutServerServiceRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PutServerServiceRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PutServerServiceRequest) SetType(v string)`

SetType sets Type field to given value.


### GetServerVersion

`func (o *PutServerServiceRequest) GetServerVersion() string`

GetServerVersion returns the ServerVersion field if non-nil, zero value otherwise.

### GetServerVersionOk

`func (o *PutServerServiceRequest) GetServerVersionOk() (*string, bool)`

GetServerVersionOk returns a tuple with the ServerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerVersion

`func (o *PutServerServiceRequest) SetServerVersion(v string)`

SetServerVersion sets ServerVersion field to given value.

### HasServerVersion

`func (o *PutServerServiceRequest) HasServerVersion() bool`

HasServerVersion returns a boolean if a field has been set.

### GetEnabled

`func (o *PutServerServiceRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PutServerServiceRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PutServerServiceRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMaximumBufferedEvents

`func (o *PutServerServiceRequest) GetMaximumBufferedEvents() int64`

GetMaximumBufferedEvents returns the MaximumBufferedEvents field if non-nil, zero value otherwise.

### GetMaximumBufferedEventsOk

`func (o *PutServerServiceRequest) GetMaximumBufferedEventsOk() (*int64, bool)`

GetMaximumBufferedEventsOk returns a tuple with the MaximumBufferedEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumBufferedEvents

`func (o *PutServerServiceRequest) SetMaximumBufferedEvents(v int64)`

SetMaximumBufferedEvents sets MaximumBufferedEvents field to given value.


### GetDestination

`func (o *PutServerServiceRequest) GetDestination() SinkBean`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *PutServerServiceRequest) GetDestinationOk() (*SinkBean, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *PutServerServiceRequest) SetDestination(v SinkBean)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *PutServerServiceRequest) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetDhcpv4Enabled

`func (o *PutServerServiceRequest) GetDhcpv4Enabled() bool`

GetDhcpv4Enabled returns the Dhcpv4Enabled field if non-nil, zero value otherwise.

### GetDhcpv4EnabledOk

`func (o *PutServerServiceRequest) GetDhcpv4EnabledOk() (*bool, bool)`

GetDhcpv4EnabledOk returns a tuple with the Dhcpv4Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpv4Enabled

`func (o *PutServerServiceRequest) SetDhcpv4Enabled(v bool)`

SetDhcpv4Enabled sets Dhcpv4Enabled field to given value.


### GetDhcpv6Enabled

`func (o *PutServerServiceRequest) GetDhcpv6Enabled() bool`

GetDhcpv6Enabled returns the Dhcpv6Enabled field if non-nil, zero value otherwise.

### GetDhcpv6EnabledOk

`func (o *PutServerServiceRequest) GetDhcpv6EnabledOk() (*bool, bool)`

GetDhcpv6EnabledOk returns a tuple with the Dhcpv6Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpv6Enabled

`func (o *PutServerServiceRequest) SetDhcpv6Enabled(v bool)`

SetDhcpv6Enabled sets Dhcpv6Enabled field to given value.


### GetFilters

`func (o *PutServerServiceRequest) GetFilters() []DnsActivityFilterBean`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *PutServerServiceRequest) GetFiltersOk() (*[]DnsActivityFilterBean, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *PutServerServiceRequest) SetFilters(v []DnsActivityFilterBean)`

SetFilters sets Filters field to given value.


### GetServers

`func (o *PutServerServiceRequest) GetServers() []SyslogRedirectionServerBean`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *PutServerServiceRequest) GetServersOk() (*[]SyslogRedirectionServerBean, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *PutServerServiceRequest) SetServers(v []SyslogRedirectionServerBean)`

SetServers sets Servers field to given value.


### GetPollingInterval

`func (o *PutServerServiceRequest) GetPollingInterval() string`

GetPollingInterval returns the PollingInterval field if non-nil, zero value otherwise.

### GetPollingIntervalOk

`func (o *PutServerServiceRequest) GetPollingIntervalOk() (*string, bool)`

GetPollingIntervalOk returns a tuple with the PollingInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollingInterval

`func (o *PutServerServiceRequest) SetPollingInterval(v string)`

SetPollingInterval sets PollingInterval field to given value.


### GetPingAllowed

`func (o *PutServerServiceRequest) GetPingAllowed() bool`

GetPingAllowed returns the PingAllowed field if non-nil, zero value otherwise.

### GetPingAllowedOk

`func (o *PutServerServiceRequest) GetPingAllowedOk() (*bool, bool)`

GetPingAllowedOk returns a tuple with the PingAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingAllowed

`func (o *PutServerServiceRequest) SetPingAllowed(v bool)`

SetPingAllowed sets PingAllowed field to given value.


### GetState

`func (o *PutServerServiceRequest) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PutServerServiceRequest) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PutServerServiceRequest) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *PutServerServiceRequest) HasState() bool`

HasState returns a boolean if a field has been set.

### GetContainer

`func (o *PutServerServiceRequest) GetContainer() GatewayContainerBean`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *PutServerServiceRequest) GetContainerOk() (*GatewayContainerBean, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *PutServerServiceRequest) SetContainer(v GatewayContainerBean)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *PutServerServiceRequest) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetRepositoryUsername

`func (o *PutServerServiceRequest) GetRepositoryUsername() string`

GetRepositoryUsername returns the RepositoryUsername field if non-nil, zero value otherwise.

### GetRepositoryUsernameOk

`func (o *PutServerServiceRequest) GetRepositoryUsernameOk() (*string, bool)`

GetRepositoryUsernameOk returns a tuple with the RepositoryUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryUsername

`func (o *PutServerServiceRequest) SetRepositoryUsername(v string)`

SetRepositoryUsername sets RepositoryUsername field to given value.

### HasRepositoryUsername

`func (o *PutServerServiceRequest) HasRepositoryUsername() bool`

HasRepositoryUsername returns a boolean if a field has been set.

### GetRepositoryPassword

`func (o *PutServerServiceRequest) GetRepositoryPassword() string`

GetRepositoryPassword returns the RepositoryPassword field if non-nil, zero value otherwise.

### GetRepositoryPasswordOk

`func (o *PutServerServiceRequest) GetRepositoryPasswordOk() (*string, bool)`

GetRepositoryPasswordOk returns a tuple with the RepositoryPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryPassword

`func (o *PutServerServiceRequest) SetRepositoryPassword(v string)`

SetRepositoryPassword sets RepositoryPassword field to given value.

### HasRepositoryPassword

`func (o *PutServerServiceRequest) HasRepositoryPassword() bool`

HasRepositoryPassword returns a boolean if a field has been set.

### GetDeleteImageOnDisable

`func (o *PutServerServiceRequest) GetDeleteImageOnDisable() bool`

GetDeleteImageOnDisable returns the DeleteImageOnDisable field if non-nil, zero value otherwise.

### GetDeleteImageOnDisableOk

`func (o *PutServerServiceRequest) GetDeleteImageOnDisableOk() (*bool, bool)`

GetDeleteImageOnDisableOk returns a tuple with the DeleteImageOnDisable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteImageOnDisable

`func (o *PutServerServiceRequest) SetDeleteImageOnDisable(v bool)`

SetDeleteImageOnDisable sets DeleteImageOnDisable field to given value.

### HasDeleteImageOnDisable

`func (o *PutServerServiceRequest) HasDeleteImageOnDisable() bool`

HasDeleteImageOnDisable returns a boolean if a field has been set.

### GetDeleteMountPointsOnDisable

`func (o *PutServerServiceRequest) GetDeleteMountPointsOnDisable() bool`

GetDeleteMountPointsOnDisable returns the DeleteMountPointsOnDisable field if non-nil, zero value otherwise.

### GetDeleteMountPointsOnDisableOk

`func (o *PutServerServiceRequest) GetDeleteMountPointsOnDisableOk() (*bool, bool)`

GetDeleteMountPointsOnDisableOk returns a tuple with the DeleteMountPointsOnDisable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteMountPointsOnDisable

`func (o *PutServerServiceRequest) SetDeleteMountPointsOnDisable(v bool)`

SetDeleteMountPointsOnDisable sets DeleteMountPointsOnDisable field to given value.

### HasDeleteMountPointsOnDisable

`func (o *PutServerServiceRequest) HasDeleteMountPointsOnDisable() bool`

HasDeleteMountPointsOnDisable returns a boolean if a field has been set.

### GetDedicatedManagementEnabled

`func (o *PutServerServiceRequest) GetDedicatedManagementEnabled() bool`

GetDedicatedManagementEnabled returns the DedicatedManagementEnabled field if non-nil, zero value otherwise.

### GetDedicatedManagementEnabledOk

`func (o *PutServerServiceRequest) GetDedicatedManagementEnabledOk() (*bool, bool)`

GetDedicatedManagementEnabledOk returns a tuple with the DedicatedManagementEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedicatedManagementEnabled

`func (o *PutServerServiceRequest) SetDedicatedManagementEnabled(v bool)`

SetDedicatedManagementEnabled sets DedicatedManagementEnabled field to given value.


### GetInterfaces

`func (o *PutServerServiceRequest) GetInterfaces() []InterfaceBean`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *PutServerServiceRequest) GetInterfacesOk() (*[]InterfaceBean, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *PutServerServiceRequest) SetInterfaces(v []InterfaceBean)`

SetInterfaces sets Interfaces field to given value.


### GetRoutes

`func (o *PutServerServiceRequest) GetRoutes() []RouteBean`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *PutServerServiceRequest) GetRoutesOk() (*[]RouteBean, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *PutServerServiceRequest) SetRoutes(v []RouteBean)`

SetRoutes sets Routes field to given value.


### GetClientId

`func (o *PutServerServiceRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *PutServerServiceRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *PutServerServiceRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetKey

`func (o *PutServerServiceRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PutServerServiceRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PutServerServiceRequest) SetKey(v string)`

SetKey sets Key field to given value.


### GetManualOverrideEnabled

`func (o *PutServerServiceRequest) GetManualOverrideEnabled() bool`

GetManualOverrideEnabled returns the ManualOverrideEnabled field if non-nil, zero value otherwise.

### GetManualOverrideEnabledOk

`func (o *PutServerServiceRequest) GetManualOverrideEnabledOk() (*bool, bool)`

GetManualOverrideEnabledOk returns a tuple with the ManualOverrideEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualOverrideEnabled

`func (o *PutServerServiceRequest) SetManualOverrideEnabled(v bool)`

SetManualOverrideEnabled sets ManualOverrideEnabled field to given value.

### HasManualOverrideEnabled

`func (o *PutServerServiceRequest) HasManualOverrideEnabled() bool`

HasManualOverrideEnabled returns a boolean if a field has been set.

### GetContactName

`func (o *PutServerServiceRequest) GetContactName() string`

GetContactName returns the ContactName field if non-nil, zero value otherwise.

### GetContactNameOk

`func (o *PutServerServiceRequest) GetContactNameOk() (*string, bool)`

GetContactNameOk returns a tuple with the ContactName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactName

`func (o *PutServerServiceRequest) SetContactName(v string)`

SetContactName sets ContactName field to given value.


### GetContactEmail

`func (o *PutServerServiceRequest) GetContactEmail() string`

GetContactEmail returns the ContactEmail field if non-nil, zero value otherwise.

### GetContactEmailOk

`func (o *PutServerServiceRequest) GetContactEmailOk() (*string, bool)`

GetContactEmailOk returns a tuple with the ContactEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactEmail

`func (o *PutServerServiceRequest) SetContactEmail(v string)`

SetContactEmail sets ContactEmail field to given value.


### GetLocation

`func (o *PutServerServiceRequest) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PutServerServiceRequest) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PutServerServiceRequest) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *PutServerServiceRequest) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetDescription

`func (o *PutServerServiceRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PutServerServiceRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PutServerServiceRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetLogLevel

`func (o *PutServerServiceRequest) GetLogLevel() string`

GetLogLevel returns the LogLevel field if non-nil, zero value otherwise.

### GetLogLevelOk

`func (o *PutServerServiceRequest) GetLogLevelOk() (*string, bool)`

GetLogLevelOk returns a tuple with the LogLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogLevel

`func (o *PutServerServiceRequest) SetLogLevel(v string)`

SetLogLevel sets LogLevel field to given value.


### GetSnmpv1

`func (o *PutServerServiceRequest) GetSnmpv1() Snmpv1Bean`

GetSnmpv1 returns the Snmpv1 field if non-nil, zero value otherwise.

### GetSnmpv1Ok

`func (o *PutServerServiceRequest) GetSnmpv1Ok() (*Snmpv1Bean, bool)`

GetSnmpv1Ok returns a tuple with the Snmpv1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpv1

`func (o *PutServerServiceRequest) SetSnmpv1(v Snmpv1Bean)`

SetSnmpv1 sets Snmpv1 field to given value.


### GetSnmpv2c

`func (o *PutServerServiceRequest) GetSnmpv2c() Snmpv2cBean`

GetSnmpv2c returns the Snmpv2c field if non-nil, zero value otherwise.

### GetSnmpv2cOk

`func (o *PutServerServiceRequest) GetSnmpv2cOk() (*Snmpv2cBean, bool)`

GetSnmpv2cOk returns a tuple with the Snmpv2c field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpv2c

`func (o *PutServerServiceRequest) SetSnmpv2c(v Snmpv2cBean)`

SetSnmpv2c sets Snmpv2c field to given value.


### GetSnmpv3

`func (o *PutServerServiceRequest) GetSnmpv3() Snmpv3Bean`

GetSnmpv3 returns the Snmpv3 field if non-nil, zero value otherwise.

### GetSnmpv3Ok

`func (o *PutServerServiceRequest) GetSnmpv3Ok() (*Snmpv3Bean, bool)`

GetSnmpv3Ok returns a tuple with the Snmpv3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpv3

`func (o *PutServerServiceRequest) SetSnmpv3(v Snmpv3Bean)`

SetSnmpv3 sets Snmpv3 field to given value.


### GetTrapServers

`func (o *PutServerServiceRequest) GetTrapServers() []SnmpTrapServerBean`

GetTrapServers returns the TrapServers field if non-nil, zero value otherwise.

### GetTrapServersOk

`func (o *PutServerServiceRequest) GetTrapServersOk() (*[]SnmpTrapServerBean, bool)`

GetTrapServersOk returns a tuple with the TrapServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrapServers

`func (o *PutServerServiceRequest) SetTrapServers(v []SnmpTrapServerBean)`

SetTrapServers sets TrapServers field to given value.


### GetTacacsPlusAuthentication

`func (o *PutServerServiceRequest) GetTacacsPlusAuthentication() TacacsPlusAuthenticationBean`

GetTacacsPlusAuthentication returns the TacacsPlusAuthentication field if non-nil, zero value otherwise.

### GetTacacsPlusAuthenticationOk

`func (o *PutServerServiceRequest) GetTacacsPlusAuthenticationOk() (*TacacsPlusAuthenticationBean, bool)`

GetTacacsPlusAuthenticationOk returns a tuple with the TacacsPlusAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTacacsPlusAuthentication

`func (o *PutServerServiceRequest) SetTacacsPlusAuthentication(v TacacsPlusAuthenticationBean)`

SetTacacsPlusAuthentication sets TacacsPlusAuthentication field to given value.


### GetIso8601TimestampsEnabled

`func (o *PutServerServiceRequest) GetIso8601TimestampsEnabled() bool`

GetIso8601TimestampsEnabled returns the Iso8601TimestampsEnabled field if non-nil, zero value otherwise.

### GetIso8601TimestampsEnabledOk

`func (o *PutServerServiceRequest) GetIso8601TimestampsEnabledOk() (*bool, bool)`

GetIso8601TimestampsEnabledOk returns a tuple with the Iso8601TimestampsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIso8601TimestampsEnabled

`func (o *PutServerServiceRequest) SetIso8601TimestampsEnabled(v bool)`

SetIso8601TimestampsEnabled sets Iso8601TimestampsEnabled field to given value.

### HasIso8601TimestampsEnabled

`func (o *PutServerServiceRequest) HasIso8601TimestampsEnabled() bool`

HasIso8601TimestampsEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


