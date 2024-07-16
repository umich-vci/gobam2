# GetServerServices200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**ServerVersion** | Pointer to **string** |  | [optional] [readonly] 
**Enabled** | Pointer to **bool** | Indicates whether the SSH service is enabled. | [optional] 
**ManualOverrideEnabled** | Pointer to **bool** | Indicates whether the syslog redirection service configuration has been manually overridden. | [optional] [readonly] 
**Addresses** | Pointer to **[]string** | The Anycast IP addresses. | [optional] 
**Routing** | Pointer to [**RoutingProtocolBean**](RoutingProtocolBean.md) |  | [optional] 
**MaximumBufferedEvents** | Pointer to **int64** | The maximum number of events to be stored in the memory buffer | [optional] 
**Destination** | Pointer to [**SinkBean**](SinkBean.md) |  | [optional] 
**Dhcpv4Enabled** | Pointer to **bool** |  | [optional] 
**Dhcpv6Enabled** | Pointer to **bool** |  | [optional] 
**Filters** | Pointer to [**[]DnsActivityFilterBean**](DnsActivityFilterBean.md) |  | [optional] 
**Servers** | Pointer to [**[]SyslogRedirectionServerBean**](SyslogRedirectionServerBean.md) | The list of servers to redirect syslog messages to. | [optional] 
**PollingInterval** | Pointer to **string** | The SNMP polling interval, in ISO 8601 duration format, that determines the frequency that SNMP values are refreshed on the system. | [optional] 
**PingAllowed** | Pointer to **bool** | Indicates whether the server responds to the ping command. | [optional] 
**State** | Pointer to **string** | The status of the Gateway service running on the DNS/DHCP Server. | [optional] 
**Container** | Pointer to [**GatewayContainerBean**](GatewayContainerBean.md) |  | [optional] 
**RepositoryUsername** | Pointer to **string** | The docker username used to pull the Gateway image from the container repository. | [optional] 
**RepositoryPassword** | Pointer to **string** | The docker password used to pull the image from a private repository. | [optional] 
**DeleteImageOnDisable** | Pointer to **bool** | Indicates whether the Gateway image is removed upon disabling the Gateway service. | [optional] 
**DeleteMountPointsOnDisable** | Pointer to **bool** | Indicates whether the local volumes and the mounted data and log directories are removed upon disabling the Gateway service. | [optional] 
**DedicatedManagementEnabled** | Pointer to **bool** |  | [optional] 
**Interfaces** | Pointer to [**[]InterfaceBean**](InterfaceBean.md) |  | [optional] 
**Routes** | Pointer to [**[]RouteBean**](RouteBean.md) |  | [optional] 
**ClientId** | Pointer to **string** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**ContactName** | Pointer to **string** | The system name to be reported through SNMP. | [optional] 
**ContactEmail** | Pointer to **string** | The email address for the system contact to be reported through SNMP. | [optional] 
**Location** | Pointer to **string** | The description of the system&#39;s location to be reported through SNMP. | [optional] 
**Description** | Pointer to **string** | The description of the system to be reported through SNMP. | [optional] 
**LogLevel** | Pointer to **string** | The logging level for SNMP events. | [optional] 
**Snmpv1** | Pointer to [**Snmpv1Bean**](Snmpv1Bean.md) |  | [optional] 
**Snmpv2c** | Pointer to [**Snmpv2cBean**](Snmpv2cBean.md) |  | [optional] 
**Snmpv3** | Pointer to [**Snmpv3Bean**](Snmpv3Bean.md) |  | [optional] 
**TrapServers** | Pointer to [**[]SnmpTrapServerBean**](SnmpTrapServerBean.md) |  | [optional] 
**TacacsPlusAuthentication** | Pointer to [**TacacsPlusAuthenticationBean**](TacacsPlusAuthenticationBean.md) |  | [optional] 
**Iso8601TimestampsEnabled** | Pointer to **bool** | Indicates whether the ISO 8601 date time format is used when logging timestamps. | [optional] 
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 

## Methods

### NewGetServerServices200ResponseDataInner

`func NewGetServerServices200ResponseDataInner() *GetServerServices200ResponseDataInner`

NewGetServerServices200ResponseDataInner instantiates a new GetServerServices200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetServerServices200ResponseDataInnerWithDefaults

`func NewGetServerServices200ResponseDataInnerWithDefaults() *GetServerServices200ResponseDataInner`

NewGetServerServices200ResponseDataInnerWithDefaults instantiates a new GetServerServices200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetServerServices200ResponseDataInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetServerServices200ResponseDataInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetServerServices200ResponseDataInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetServerServices200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GetServerServices200ResponseDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetServerServices200ResponseDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetServerServices200ResponseDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetServerServices200ResponseDataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetServerVersion

`func (o *GetServerServices200ResponseDataInner) GetServerVersion() string`

GetServerVersion returns the ServerVersion field if non-nil, zero value otherwise.

### GetServerVersionOk

`func (o *GetServerServices200ResponseDataInner) GetServerVersionOk() (*string, bool)`

GetServerVersionOk returns a tuple with the ServerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerVersion

`func (o *GetServerServices200ResponseDataInner) SetServerVersion(v string)`

SetServerVersion sets ServerVersion field to given value.

### HasServerVersion

`func (o *GetServerServices200ResponseDataInner) HasServerVersion() bool`

HasServerVersion returns a boolean if a field has been set.

### GetEnabled

`func (o *GetServerServices200ResponseDataInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetServerServices200ResponseDataInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetServerServices200ResponseDataInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetServerServices200ResponseDataInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetManualOverrideEnabled

`func (o *GetServerServices200ResponseDataInner) GetManualOverrideEnabled() bool`

GetManualOverrideEnabled returns the ManualOverrideEnabled field if non-nil, zero value otherwise.

### GetManualOverrideEnabledOk

`func (o *GetServerServices200ResponseDataInner) GetManualOverrideEnabledOk() (*bool, bool)`

GetManualOverrideEnabledOk returns a tuple with the ManualOverrideEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualOverrideEnabled

`func (o *GetServerServices200ResponseDataInner) SetManualOverrideEnabled(v bool)`

SetManualOverrideEnabled sets ManualOverrideEnabled field to given value.

### HasManualOverrideEnabled

`func (o *GetServerServices200ResponseDataInner) HasManualOverrideEnabled() bool`

HasManualOverrideEnabled returns a boolean if a field has been set.

### GetAddresses

`func (o *GetServerServices200ResponseDataInner) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *GetServerServices200ResponseDataInner) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *GetServerServices200ResponseDataInner) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *GetServerServices200ResponseDataInner) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetRouting

`func (o *GetServerServices200ResponseDataInner) GetRouting() RoutingProtocolBean`

GetRouting returns the Routing field if non-nil, zero value otherwise.

### GetRoutingOk

`func (o *GetServerServices200ResponseDataInner) GetRoutingOk() (*RoutingProtocolBean, bool)`

GetRoutingOk returns a tuple with the Routing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouting

`func (o *GetServerServices200ResponseDataInner) SetRouting(v RoutingProtocolBean)`

SetRouting sets Routing field to given value.

### HasRouting

`func (o *GetServerServices200ResponseDataInner) HasRouting() bool`

HasRouting returns a boolean if a field has been set.

### GetMaximumBufferedEvents

`func (o *GetServerServices200ResponseDataInner) GetMaximumBufferedEvents() int64`

GetMaximumBufferedEvents returns the MaximumBufferedEvents field if non-nil, zero value otherwise.

### GetMaximumBufferedEventsOk

`func (o *GetServerServices200ResponseDataInner) GetMaximumBufferedEventsOk() (*int64, bool)`

GetMaximumBufferedEventsOk returns a tuple with the MaximumBufferedEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumBufferedEvents

`func (o *GetServerServices200ResponseDataInner) SetMaximumBufferedEvents(v int64)`

SetMaximumBufferedEvents sets MaximumBufferedEvents field to given value.

### HasMaximumBufferedEvents

`func (o *GetServerServices200ResponseDataInner) HasMaximumBufferedEvents() bool`

HasMaximumBufferedEvents returns a boolean if a field has been set.

### GetDestination

`func (o *GetServerServices200ResponseDataInner) GetDestination() SinkBean`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *GetServerServices200ResponseDataInner) GetDestinationOk() (*SinkBean, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *GetServerServices200ResponseDataInner) SetDestination(v SinkBean)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *GetServerServices200ResponseDataInner) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetDhcpv4Enabled

`func (o *GetServerServices200ResponseDataInner) GetDhcpv4Enabled() bool`

GetDhcpv4Enabled returns the Dhcpv4Enabled field if non-nil, zero value otherwise.

### GetDhcpv4EnabledOk

`func (o *GetServerServices200ResponseDataInner) GetDhcpv4EnabledOk() (*bool, bool)`

GetDhcpv4EnabledOk returns a tuple with the Dhcpv4Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpv4Enabled

`func (o *GetServerServices200ResponseDataInner) SetDhcpv4Enabled(v bool)`

SetDhcpv4Enabled sets Dhcpv4Enabled field to given value.

### HasDhcpv4Enabled

`func (o *GetServerServices200ResponseDataInner) HasDhcpv4Enabled() bool`

HasDhcpv4Enabled returns a boolean if a field has been set.

### GetDhcpv6Enabled

`func (o *GetServerServices200ResponseDataInner) GetDhcpv6Enabled() bool`

GetDhcpv6Enabled returns the Dhcpv6Enabled field if non-nil, zero value otherwise.

### GetDhcpv6EnabledOk

`func (o *GetServerServices200ResponseDataInner) GetDhcpv6EnabledOk() (*bool, bool)`

GetDhcpv6EnabledOk returns a tuple with the Dhcpv6Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpv6Enabled

`func (o *GetServerServices200ResponseDataInner) SetDhcpv6Enabled(v bool)`

SetDhcpv6Enabled sets Dhcpv6Enabled field to given value.

### HasDhcpv6Enabled

`func (o *GetServerServices200ResponseDataInner) HasDhcpv6Enabled() bool`

HasDhcpv6Enabled returns a boolean if a field has been set.

### GetFilters

`func (o *GetServerServices200ResponseDataInner) GetFilters() []DnsActivityFilterBean`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *GetServerServices200ResponseDataInner) GetFiltersOk() (*[]DnsActivityFilterBean, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *GetServerServices200ResponseDataInner) SetFilters(v []DnsActivityFilterBean)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *GetServerServices200ResponseDataInner) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetServers

`func (o *GetServerServices200ResponseDataInner) GetServers() []SyslogRedirectionServerBean`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *GetServerServices200ResponseDataInner) GetServersOk() (*[]SyslogRedirectionServerBean, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *GetServerServices200ResponseDataInner) SetServers(v []SyslogRedirectionServerBean)`

SetServers sets Servers field to given value.

### HasServers

`func (o *GetServerServices200ResponseDataInner) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetPollingInterval

`func (o *GetServerServices200ResponseDataInner) GetPollingInterval() string`

GetPollingInterval returns the PollingInterval field if non-nil, zero value otherwise.

### GetPollingIntervalOk

`func (o *GetServerServices200ResponseDataInner) GetPollingIntervalOk() (*string, bool)`

GetPollingIntervalOk returns a tuple with the PollingInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollingInterval

`func (o *GetServerServices200ResponseDataInner) SetPollingInterval(v string)`

SetPollingInterval sets PollingInterval field to given value.

### HasPollingInterval

`func (o *GetServerServices200ResponseDataInner) HasPollingInterval() bool`

HasPollingInterval returns a boolean if a field has been set.

### GetPingAllowed

`func (o *GetServerServices200ResponseDataInner) GetPingAllowed() bool`

GetPingAllowed returns the PingAllowed field if non-nil, zero value otherwise.

### GetPingAllowedOk

`func (o *GetServerServices200ResponseDataInner) GetPingAllowedOk() (*bool, bool)`

GetPingAllowedOk returns a tuple with the PingAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingAllowed

`func (o *GetServerServices200ResponseDataInner) SetPingAllowed(v bool)`

SetPingAllowed sets PingAllowed field to given value.

### HasPingAllowed

`func (o *GetServerServices200ResponseDataInner) HasPingAllowed() bool`

HasPingAllowed returns a boolean if a field has been set.

### GetState

`func (o *GetServerServices200ResponseDataInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetServerServices200ResponseDataInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetServerServices200ResponseDataInner) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GetServerServices200ResponseDataInner) HasState() bool`

HasState returns a boolean if a field has been set.

### GetContainer

`func (o *GetServerServices200ResponseDataInner) GetContainer() GatewayContainerBean`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *GetServerServices200ResponseDataInner) GetContainerOk() (*GatewayContainerBean, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *GetServerServices200ResponseDataInner) SetContainer(v GatewayContainerBean)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *GetServerServices200ResponseDataInner) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetRepositoryUsername

`func (o *GetServerServices200ResponseDataInner) GetRepositoryUsername() string`

GetRepositoryUsername returns the RepositoryUsername field if non-nil, zero value otherwise.

### GetRepositoryUsernameOk

`func (o *GetServerServices200ResponseDataInner) GetRepositoryUsernameOk() (*string, bool)`

GetRepositoryUsernameOk returns a tuple with the RepositoryUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryUsername

`func (o *GetServerServices200ResponseDataInner) SetRepositoryUsername(v string)`

SetRepositoryUsername sets RepositoryUsername field to given value.

### HasRepositoryUsername

`func (o *GetServerServices200ResponseDataInner) HasRepositoryUsername() bool`

HasRepositoryUsername returns a boolean if a field has been set.

### GetRepositoryPassword

`func (o *GetServerServices200ResponseDataInner) GetRepositoryPassword() string`

GetRepositoryPassword returns the RepositoryPassword field if non-nil, zero value otherwise.

### GetRepositoryPasswordOk

`func (o *GetServerServices200ResponseDataInner) GetRepositoryPasswordOk() (*string, bool)`

GetRepositoryPasswordOk returns a tuple with the RepositoryPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryPassword

`func (o *GetServerServices200ResponseDataInner) SetRepositoryPassword(v string)`

SetRepositoryPassword sets RepositoryPassword field to given value.

### HasRepositoryPassword

`func (o *GetServerServices200ResponseDataInner) HasRepositoryPassword() bool`

HasRepositoryPassword returns a boolean if a field has been set.

### GetDeleteImageOnDisable

`func (o *GetServerServices200ResponseDataInner) GetDeleteImageOnDisable() bool`

GetDeleteImageOnDisable returns the DeleteImageOnDisable field if non-nil, zero value otherwise.

### GetDeleteImageOnDisableOk

`func (o *GetServerServices200ResponseDataInner) GetDeleteImageOnDisableOk() (*bool, bool)`

GetDeleteImageOnDisableOk returns a tuple with the DeleteImageOnDisable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteImageOnDisable

`func (o *GetServerServices200ResponseDataInner) SetDeleteImageOnDisable(v bool)`

SetDeleteImageOnDisable sets DeleteImageOnDisable field to given value.

### HasDeleteImageOnDisable

`func (o *GetServerServices200ResponseDataInner) HasDeleteImageOnDisable() bool`

HasDeleteImageOnDisable returns a boolean if a field has been set.

### GetDeleteMountPointsOnDisable

`func (o *GetServerServices200ResponseDataInner) GetDeleteMountPointsOnDisable() bool`

GetDeleteMountPointsOnDisable returns the DeleteMountPointsOnDisable field if non-nil, zero value otherwise.

### GetDeleteMountPointsOnDisableOk

`func (o *GetServerServices200ResponseDataInner) GetDeleteMountPointsOnDisableOk() (*bool, bool)`

GetDeleteMountPointsOnDisableOk returns a tuple with the DeleteMountPointsOnDisable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteMountPointsOnDisable

`func (o *GetServerServices200ResponseDataInner) SetDeleteMountPointsOnDisable(v bool)`

SetDeleteMountPointsOnDisable sets DeleteMountPointsOnDisable field to given value.

### HasDeleteMountPointsOnDisable

`func (o *GetServerServices200ResponseDataInner) HasDeleteMountPointsOnDisable() bool`

HasDeleteMountPointsOnDisable returns a boolean if a field has been set.

### GetDedicatedManagementEnabled

`func (o *GetServerServices200ResponseDataInner) GetDedicatedManagementEnabled() bool`

GetDedicatedManagementEnabled returns the DedicatedManagementEnabled field if non-nil, zero value otherwise.

### GetDedicatedManagementEnabledOk

`func (o *GetServerServices200ResponseDataInner) GetDedicatedManagementEnabledOk() (*bool, bool)`

GetDedicatedManagementEnabledOk returns a tuple with the DedicatedManagementEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedicatedManagementEnabled

`func (o *GetServerServices200ResponseDataInner) SetDedicatedManagementEnabled(v bool)`

SetDedicatedManagementEnabled sets DedicatedManagementEnabled field to given value.

### HasDedicatedManagementEnabled

`func (o *GetServerServices200ResponseDataInner) HasDedicatedManagementEnabled() bool`

HasDedicatedManagementEnabled returns a boolean if a field has been set.

### GetInterfaces

`func (o *GetServerServices200ResponseDataInner) GetInterfaces() []InterfaceBean`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *GetServerServices200ResponseDataInner) GetInterfacesOk() (*[]InterfaceBean, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *GetServerServices200ResponseDataInner) SetInterfaces(v []InterfaceBean)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *GetServerServices200ResponseDataInner) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.

### GetRoutes

`func (o *GetServerServices200ResponseDataInner) GetRoutes() []RouteBean`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *GetServerServices200ResponseDataInner) GetRoutesOk() (*[]RouteBean, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *GetServerServices200ResponseDataInner) SetRoutes(v []RouteBean)`

SetRoutes sets Routes field to given value.

### HasRoutes

`func (o *GetServerServices200ResponseDataInner) HasRoutes() bool`

HasRoutes returns a boolean if a field has been set.

### GetClientId

`func (o *GetServerServices200ResponseDataInner) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *GetServerServices200ResponseDataInner) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *GetServerServices200ResponseDataInner) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *GetServerServices200ResponseDataInner) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetKey

`func (o *GetServerServices200ResponseDataInner) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *GetServerServices200ResponseDataInner) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *GetServerServices200ResponseDataInner) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *GetServerServices200ResponseDataInner) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetContactName

`func (o *GetServerServices200ResponseDataInner) GetContactName() string`

GetContactName returns the ContactName field if non-nil, zero value otherwise.

### GetContactNameOk

`func (o *GetServerServices200ResponseDataInner) GetContactNameOk() (*string, bool)`

GetContactNameOk returns a tuple with the ContactName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactName

`func (o *GetServerServices200ResponseDataInner) SetContactName(v string)`

SetContactName sets ContactName field to given value.

### HasContactName

`func (o *GetServerServices200ResponseDataInner) HasContactName() bool`

HasContactName returns a boolean if a field has been set.

### GetContactEmail

`func (o *GetServerServices200ResponseDataInner) GetContactEmail() string`

GetContactEmail returns the ContactEmail field if non-nil, zero value otherwise.

### GetContactEmailOk

`func (o *GetServerServices200ResponseDataInner) GetContactEmailOk() (*string, bool)`

GetContactEmailOk returns a tuple with the ContactEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactEmail

`func (o *GetServerServices200ResponseDataInner) SetContactEmail(v string)`

SetContactEmail sets ContactEmail field to given value.

### HasContactEmail

`func (o *GetServerServices200ResponseDataInner) HasContactEmail() bool`

HasContactEmail returns a boolean if a field has been set.

### GetLocation

`func (o *GetServerServices200ResponseDataInner) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *GetServerServices200ResponseDataInner) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *GetServerServices200ResponseDataInner) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *GetServerServices200ResponseDataInner) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetDescription

`func (o *GetServerServices200ResponseDataInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetServerServices200ResponseDataInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetServerServices200ResponseDataInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetServerServices200ResponseDataInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLogLevel

`func (o *GetServerServices200ResponseDataInner) GetLogLevel() string`

GetLogLevel returns the LogLevel field if non-nil, zero value otherwise.

### GetLogLevelOk

`func (o *GetServerServices200ResponseDataInner) GetLogLevelOk() (*string, bool)`

GetLogLevelOk returns a tuple with the LogLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogLevel

`func (o *GetServerServices200ResponseDataInner) SetLogLevel(v string)`

SetLogLevel sets LogLevel field to given value.

### HasLogLevel

`func (o *GetServerServices200ResponseDataInner) HasLogLevel() bool`

HasLogLevel returns a boolean if a field has been set.

### GetSnmpv1

`func (o *GetServerServices200ResponseDataInner) GetSnmpv1() Snmpv1Bean`

GetSnmpv1 returns the Snmpv1 field if non-nil, zero value otherwise.

### GetSnmpv1Ok

`func (o *GetServerServices200ResponseDataInner) GetSnmpv1Ok() (*Snmpv1Bean, bool)`

GetSnmpv1Ok returns a tuple with the Snmpv1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpv1

`func (o *GetServerServices200ResponseDataInner) SetSnmpv1(v Snmpv1Bean)`

SetSnmpv1 sets Snmpv1 field to given value.

### HasSnmpv1

`func (o *GetServerServices200ResponseDataInner) HasSnmpv1() bool`

HasSnmpv1 returns a boolean if a field has been set.

### GetSnmpv2c

`func (o *GetServerServices200ResponseDataInner) GetSnmpv2c() Snmpv2cBean`

GetSnmpv2c returns the Snmpv2c field if non-nil, zero value otherwise.

### GetSnmpv2cOk

`func (o *GetServerServices200ResponseDataInner) GetSnmpv2cOk() (*Snmpv2cBean, bool)`

GetSnmpv2cOk returns a tuple with the Snmpv2c field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpv2c

`func (o *GetServerServices200ResponseDataInner) SetSnmpv2c(v Snmpv2cBean)`

SetSnmpv2c sets Snmpv2c field to given value.

### HasSnmpv2c

`func (o *GetServerServices200ResponseDataInner) HasSnmpv2c() bool`

HasSnmpv2c returns a boolean if a field has been set.

### GetSnmpv3

`func (o *GetServerServices200ResponseDataInner) GetSnmpv3() Snmpv3Bean`

GetSnmpv3 returns the Snmpv3 field if non-nil, zero value otherwise.

### GetSnmpv3Ok

`func (o *GetServerServices200ResponseDataInner) GetSnmpv3Ok() (*Snmpv3Bean, bool)`

GetSnmpv3Ok returns a tuple with the Snmpv3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpv3

`func (o *GetServerServices200ResponseDataInner) SetSnmpv3(v Snmpv3Bean)`

SetSnmpv3 sets Snmpv3 field to given value.

### HasSnmpv3

`func (o *GetServerServices200ResponseDataInner) HasSnmpv3() bool`

HasSnmpv3 returns a boolean if a field has been set.

### GetTrapServers

`func (o *GetServerServices200ResponseDataInner) GetTrapServers() []SnmpTrapServerBean`

GetTrapServers returns the TrapServers field if non-nil, zero value otherwise.

### GetTrapServersOk

`func (o *GetServerServices200ResponseDataInner) GetTrapServersOk() (*[]SnmpTrapServerBean, bool)`

GetTrapServersOk returns a tuple with the TrapServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrapServers

`func (o *GetServerServices200ResponseDataInner) SetTrapServers(v []SnmpTrapServerBean)`

SetTrapServers sets TrapServers field to given value.

### HasTrapServers

`func (o *GetServerServices200ResponseDataInner) HasTrapServers() bool`

HasTrapServers returns a boolean if a field has been set.

### GetTacacsPlusAuthentication

`func (o *GetServerServices200ResponseDataInner) GetTacacsPlusAuthentication() TacacsPlusAuthenticationBean`

GetTacacsPlusAuthentication returns the TacacsPlusAuthentication field if non-nil, zero value otherwise.

### GetTacacsPlusAuthenticationOk

`func (o *GetServerServices200ResponseDataInner) GetTacacsPlusAuthenticationOk() (*TacacsPlusAuthenticationBean, bool)`

GetTacacsPlusAuthenticationOk returns a tuple with the TacacsPlusAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTacacsPlusAuthentication

`func (o *GetServerServices200ResponseDataInner) SetTacacsPlusAuthentication(v TacacsPlusAuthenticationBean)`

SetTacacsPlusAuthentication sets TacacsPlusAuthentication field to given value.

### HasTacacsPlusAuthentication

`func (o *GetServerServices200ResponseDataInner) HasTacacsPlusAuthentication() bool`

HasTacacsPlusAuthentication returns a boolean if a field has been set.

### GetIso8601TimestampsEnabled

`func (o *GetServerServices200ResponseDataInner) GetIso8601TimestampsEnabled() bool`

GetIso8601TimestampsEnabled returns the Iso8601TimestampsEnabled field if non-nil, zero value otherwise.

### GetIso8601TimestampsEnabledOk

`func (o *GetServerServices200ResponseDataInner) GetIso8601TimestampsEnabledOk() (*bool, bool)`

GetIso8601TimestampsEnabledOk returns a tuple with the Iso8601TimestampsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIso8601TimestampsEnabled

`func (o *GetServerServices200ResponseDataInner) SetIso8601TimestampsEnabled(v bool)`

SetIso8601TimestampsEnabled sets Iso8601TimestampsEnabled field to given value.

### HasIso8601TimestampsEnabled

`func (o *GetServerServices200ResponseDataInner) HasIso8601TimestampsEnabled() bool`

HasIso8601TimestampsEnabled returns a boolean if a field has been set.

### GetLinks

`func (o *GetServerServices200ResponseDataInner) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetServerServices200ResponseDataInner) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetServerServices200ResponseDataInner) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetServerServices200ResponseDataInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


