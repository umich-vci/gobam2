# GetServerServices200Response1DataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**ServerVersion** | Pointer to **string** |  | [optional] [readonly] 
**Enabled** | Pointer to **bool** | Indicates whether the SSH service is enabled. | [optional] 
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
**ManualOverrideEnabled** | Pointer to **bool** | Indicates whether the syslog redirection service configuration has been manually overridden. | [optional] [readonly] 
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

## Methods

### NewGetServerServices200Response1DataInner

`func NewGetServerServices200Response1DataInner() *GetServerServices200Response1DataInner`

NewGetServerServices200Response1DataInner instantiates a new GetServerServices200Response1DataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetServerServices200Response1DataInnerWithDefaults

`func NewGetServerServices200Response1DataInnerWithDefaults() *GetServerServices200Response1DataInner`

NewGetServerServices200Response1DataInnerWithDefaults instantiates a new GetServerServices200Response1DataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetServerServices200Response1DataInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetServerServices200Response1DataInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetServerServices200Response1DataInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetServerServices200Response1DataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GetServerServices200Response1DataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetServerServices200Response1DataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetServerServices200Response1DataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetServerServices200Response1DataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetServerVersion

`func (o *GetServerServices200Response1DataInner) GetServerVersion() string`

GetServerVersion returns the ServerVersion field if non-nil, zero value otherwise.

### GetServerVersionOk

`func (o *GetServerServices200Response1DataInner) GetServerVersionOk() (*string, bool)`

GetServerVersionOk returns a tuple with the ServerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerVersion

`func (o *GetServerServices200Response1DataInner) SetServerVersion(v string)`

SetServerVersion sets ServerVersion field to given value.

### HasServerVersion

`func (o *GetServerServices200Response1DataInner) HasServerVersion() bool`

HasServerVersion returns a boolean if a field has been set.

### GetEnabled

`func (o *GetServerServices200Response1DataInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetServerServices200Response1DataInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetServerServices200Response1DataInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetServerServices200Response1DataInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMaximumBufferedEvents

`func (o *GetServerServices200Response1DataInner) GetMaximumBufferedEvents() int64`

GetMaximumBufferedEvents returns the MaximumBufferedEvents field if non-nil, zero value otherwise.

### GetMaximumBufferedEventsOk

`func (o *GetServerServices200Response1DataInner) GetMaximumBufferedEventsOk() (*int64, bool)`

GetMaximumBufferedEventsOk returns a tuple with the MaximumBufferedEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumBufferedEvents

`func (o *GetServerServices200Response1DataInner) SetMaximumBufferedEvents(v int64)`

SetMaximumBufferedEvents sets MaximumBufferedEvents field to given value.

### HasMaximumBufferedEvents

`func (o *GetServerServices200Response1DataInner) HasMaximumBufferedEvents() bool`

HasMaximumBufferedEvents returns a boolean if a field has been set.

### GetDestination

`func (o *GetServerServices200Response1DataInner) GetDestination() SinkBean`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *GetServerServices200Response1DataInner) GetDestinationOk() (*SinkBean, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *GetServerServices200Response1DataInner) SetDestination(v SinkBean)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *GetServerServices200Response1DataInner) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetDhcpv4Enabled

`func (o *GetServerServices200Response1DataInner) GetDhcpv4Enabled() bool`

GetDhcpv4Enabled returns the Dhcpv4Enabled field if non-nil, zero value otherwise.

### GetDhcpv4EnabledOk

`func (o *GetServerServices200Response1DataInner) GetDhcpv4EnabledOk() (*bool, bool)`

GetDhcpv4EnabledOk returns a tuple with the Dhcpv4Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpv4Enabled

`func (o *GetServerServices200Response1DataInner) SetDhcpv4Enabled(v bool)`

SetDhcpv4Enabled sets Dhcpv4Enabled field to given value.

### HasDhcpv4Enabled

`func (o *GetServerServices200Response1DataInner) HasDhcpv4Enabled() bool`

HasDhcpv4Enabled returns a boolean if a field has been set.

### GetDhcpv6Enabled

`func (o *GetServerServices200Response1DataInner) GetDhcpv6Enabled() bool`

GetDhcpv6Enabled returns the Dhcpv6Enabled field if non-nil, zero value otherwise.

### GetDhcpv6EnabledOk

`func (o *GetServerServices200Response1DataInner) GetDhcpv6EnabledOk() (*bool, bool)`

GetDhcpv6EnabledOk returns a tuple with the Dhcpv6Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpv6Enabled

`func (o *GetServerServices200Response1DataInner) SetDhcpv6Enabled(v bool)`

SetDhcpv6Enabled sets Dhcpv6Enabled field to given value.

### HasDhcpv6Enabled

`func (o *GetServerServices200Response1DataInner) HasDhcpv6Enabled() bool`

HasDhcpv6Enabled returns a boolean if a field has been set.

### GetFilters

`func (o *GetServerServices200Response1DataInner) GetFilters() []DnsActivityFilterBean`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *GetServerServices200Response1DataInner) GetFiltersOk() (*[]DnsActivityFilterBean, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *GetServerServices200Response1DataInner) SetFilters(v []DnsActivityFilterBean)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *GetServerServices200Response1DataInner) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetServers

`func (o *GetServerServices200Response1DataInner) GetServers() []SyslogRedirectionServerBean`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *GetServerServices200Response1DataInner) GetServersOk() (*[]SyslogRedirectionServerBean, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *GetServerServices200Response1DataInner) SetServers(v []SyslogRedirectionServerBean)`

SetServers sets Servers field to given value.

### HasServers

`func (o *GetServerServices200Response1DataInner) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetPollingInterval

`func (o *GetServerServices200Response1DataInner) GetPollingInterval() string`

GetPollingInterval returns the PollingInterval field if non-nil, zero value otherwise.

### GetPollingIntervalOk

`func (o *GetServerServices200Response1DataInner) GetPollingIntervalOk() (*string, bool)`

GetPollingIntervalOk returns a tuple with the PollingInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollingInterval

`func (o *GetServerServices200Response1DataInner) SetPollingInterval(v string)`

SetPollingInterval sets PollingInterval field to given value.

### HasPollingInterval

`func (o *GetServerServices200Response1DataInner) HasPollingInterval() bool`

HasPollingInterval returns a boolean if a field has been set.

### GetPingAllowed

`func (o *GetServerServices200Response1DataInner) GetPingAllowed() bool`

GetPingAllowed returns the PingAllowed field if non-nil, zero value otherwise.

### GetPingAllowedOk

`func (o *GetServerServices200Response1DataInner) GetPingAllowedOk() (*bool, bool)`

GetPingAllowedOk returns a tuple with the PingAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingAllowed

`func (o *GetServerServices200Response1DataInner) SetPingAllowed(v bool)`

SetPingAllowed sets PingAllowed field to given value.

### HasPingAllowed

`func (o *GetServerServices200Response1DataInner) HasPingAllowed() bool`

HasPingAllowed returns a boolean if a field has been set.

### GetState

`func (o *GetServerServices200Response1DataInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetServerServices200Response1DataInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetServerServices200Response1DataInner) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GetServerServices200Response1DataInner) HasState() bool`

HasState returns a boolean if a field has been set.

### GetContainer

`func (o *GetServerServices200Response1DataInner) GetContainer() GatewayContainerBean`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *GetServerServices200Response1DataInner) GetContainerOk() (*GatewayContainerBean, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *GetServerServices200Response1DataInner) SetContainer(v GatewayContainerBean)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *GetServerServices200Response1DataInner) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetRepositoryUsername

`func (o *GetServerServices200Response1DataInner) GetRepositoryUsername() string`

GetRepositoryUsername returns the RepositoryUsername field if non-nil, zero value otherwise.

### GetRepositoryUsernameOk

`func (o *GetServerServices200Response1DataInner) GetRepositoryUsernameOk() (*string, bool)`

GetRepositoryUsernameOk returns a tuple with the RepositoryUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryUsername

`func (o *GetServerServices200Response1DataInner) SetRepositoryUsername(v string)`

SetRepositoryUsername sets RepositoryUsername field to given value.

### HasRepositoryUsername

`func (o *GetServerServices200Response1DataInner) HasRepositoryUsername() bool`

HasRepositoryUsername returns a boolean if a field has been set.

### GetRepositoryPassword

`func (o *GetServerServices200Response1DataInner) GetRepositoryPassword() string`

GetRepositoryPassword returns the RepositoryPassword field if non-nil, zero value otherwise.

### GetRepositoryPasswordOk

`func (o *GetServerServices200Response1DataInner) GetRepositoryPasswordOk() (*string, bool)`

GetRepositoryPasswordOk returns a tuple with the RepositoryPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryPassword

`func (o *GetServerServices200Response1DataInner) SetRepositoryPassword(v string)`

SetRepositoryPassword sets RepositoryPassword field to given value.

### HasRepositoryPassword

`func (o *GetServerServices200Response1DataInner) HasRepositoryPassword() bool`

HasRepositoryPassword returns a boolean if a field has been set.

### GetDeleteImageOnDisable

`func (o *GetServerServices200Response1DataInner) GetDeleteImageOnDisable() bool`

GetDeleteImageOnDisable returns the DeleteImageOnDisable field if non-nil, zero value otherwise.

### GetDeleteImageOnDisableOk

`func (o *GetServerServices200Response1DataInner) GetDeleteImageOnDisableOk() (*bool, bool)`

GetDeleteImageOnDisableOk returns a tuple with the DeleteImageOnDisable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteImageOnDisable

`func (o *GetServerServices200Response1DataInner) SetDeleteImageOnDisable(v bool)`

SetDeleteImageOnDisable sets DeleteImageOnDisable field to given value.

### HasDeleteImageOnDisable

`func (o *GetServerServices200Response1DataInner) HasDeleteImageOnDisable() bool`

HasDeleteImageOnDisable returns a boolean if a field has been set.

### GetDeleteMountPointsOnDisable

`func (o *GetServerServices200Response1DataInner) GetDeleteMountPointsOnDisable() bool`

GetDeleteMountPointsOnDisable returns the DeleteMountPointsOnDisable field if non-nil, zero value otherwise.

### GetDeleteMountPointsOnDisableOk

`func (o *GetServerServices200Response1DataInner) GetDeleteMountPointsOnDisableOk() (*bool, bool)`

GetDeleteMountPointsOnDisableOk returns a tuple with the DeleteMountPointsOnDisable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteMountPointsOnDisable

`func (o *GetServerServices200Response1DataInner) SetDeleteMountPointsOnDisable(v bool)`

SetDeleteMountPointsOnDisable sets DeleteMountPointsOnDisable field to given value.

### HasDeleteMountPointsOnDisable

`func (o *GetServerServices200Response1DataInner) HasDeleteMountPointsOnDisable() bool`

HasDeleteMountPointsOnDisable returns a boolean if a field has been set.

### GetDedicatedManagementEnabled

`func (o *GetServerServices200Response1DataInner) GetDedicatedManagementEnabled() bool`

GetDedicatedManagementEnabled returns the DedicatedManagementEnabled field if non-nil, zero value otherwise.

### GetDedicatedManagementEnabledOk

`func (o *GetServerServices200Response1DataInner) GetDedicatedManagementEnabledOk() (*bool, bool)`

GetDedicatedManagementEnabledOk returns a tuple with the DedicatedManagementEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedicatedManagementEnabled

`func (o *GetServerServices200Response1DataInner) SetDedicatedManagementEnabled(v bool)`

SetDedicatedManagementEnabled sets DedicatedManagementEnabled field to given value.

### HasDedicatedManagementEnabled

`func (o *GetServerServices200Response1DataInner) HasDedicatedManagementEnabled() bool`

HasDedicatedManagementEnabled returns a boolean if a field has been set.

### GetInterfaces

`func (o *GetServerServices200Response1DataInner) GetInterfaces() []InterfaceBean`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *GetServerServices200Response1DataInner) GetInterfacesOk() (*[]InterfaceBean, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *GetServerServices200Response1DataInner) SetInterfaces(v []InterfaceBean)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *GetServerServices200Response1DataInner) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.

### GetRoutes

`func (o *GetServerServices200Response1DataInner) GetRoutes() []RouteBean`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *GetServerServices200Response1DataInner) GetRoutesOk() (*[]RouteBean, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *GetServerServices200Response1DataInner) SetRoutes(v []RouteBean)`

SetRoutes sets Routes field to given value.

### HasRoutes

`func (o *GetServerServices200Response1DataInner) HasRoutes() bool`

HasRoutes returns a boolean if a field has been set.

### GetClientId

`func (o *GetServerServices200Response1DataInner) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *GetServerServices200Response1DataInner) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *GetServerServices200Response1DataInner) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *GetServerServices200Response1DataInner) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetKey

`func (o *GetServerServices200Response1DataInner) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *GetServerServices200Response1DataInner) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *GetServerServices200Response1DataInner) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *GetServerServices200Response1DataInner) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetManualOverrideEnabled

`func (o *GetServerServices200Response1DataInner) GetManualOverrideEnabled() bool`

GetManualOverrideEnabled returns the ManualOverrideEnabled field if non-nil, zero value otherwise.

### GetManualOverrideEnabledOk

`func (o *GetServerServices200Response1DataInner) GetManualOverrideEnabledOk() (*bool, bool)`

GetManualOverrideEnabledOk returns a tuple with the ManualOverrideEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualOverrideEnabled

`func (o *GetServerServices200Response1DataInner) SetManualOverrideEnabled(v bool)`

SetManualOverrideEnabled sets ManualOverrideEnabled field to given value.

### HasManualOverrideEnabled

`func (o *GetServerServices200Response1DataInner) HasManualOverrideEnabled() bool`

HasManualOverrideEnabled returns a boolean if a field has been set.

### GetContactName

`func (o *GetServerServices200Response1DataInner) GetContactName() string`

GetContactName returns the ContactName field if non-nil, zero value otherwise.

### GetContactNameOk

`func (o *GetServerServices200Response1DataInner) GetContactNameOk() (*string, bool)`

GetContactNameOk returns a tuple with the ContactName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactName

`func (o *GetServerServices200Response1DataInner) SetContactName(v string)`

SetContactName sets ContactName field to given value.

### HasContactName

`func (o *GetServerServices200Response1DataInner) HasContactName() bool`

HasContactName returns a boolean if a field has been set.

### GetContactEmail

`func (o *GetServerServices200Response1DataInner) GetContactEmail() string`

GetContactEmail returns the ContactEmail field if non-nil, zero value otherwise.

### GetContactEmailOk

`func (o *GetServerServices200Response1DataInner) GetContactEmailOk() (*string, bool)`

GetContactEmailOk returns a tuple with the ContactEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactEmail

`func (o *GetServerServices200Response1DataInner) SetContactEmail(v string)`

SetContactEmail sets ContactEmail field to given value.

### HasContactEmail

`func (o *GetServerServices200Response1DataInner) HasContactEmail() bool`

HasContactEmail returns a boolean if a field has been set.

### GetLocation

`func (o *GetServerServices200Response1DataInner) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *GetServerServices200Response1DataInner) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *GetServerServices200Response1DataInner) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *GetServerServices200Response1DataInner) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetDescription

`func (o *GetServerServices200Response1DataInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetServerServices200Response1DataInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetServerServices200Response1DataInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetServerServices200Response1DataInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLogLevel

`func (o *GetServerServices200Response1DataInner) GetLogLevel() string`

GetLogLevel returns the LogLevel field if non-nil, zero value otherwise.

### GetLogLevelOk

`func (o *GetServerServices200Response1DataInner) GetLogLevelOk() (*string, bool)`

GetLogLevelOk returns a tuple with the LogLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogLevel

`func (o *GetServerServices200Response1DataInner) SetLogLevel(v string)`

SetLogLevel sets LogLevel field to given value.

### HasLogLevel

`func (o *GetServerServices200Response1DataInner) HasLogLevel() bool`

HasLogLevel returns a boolean if a field has been set.

### GetSnmpv1

`func (o *GetServerServices200Response1DataInner) GetSnmpv1() Snmpv1Bean`

GetSnmpv1 returns the Snmpv1 field if non-nil, zero value otherwise.

### GetSnmpv1Ok

`func (o *GetServerServices200Response1DataInner) GetSnmpv1Ok() (*Snmpv1Bean, bool)`

GetSnmpv1Ok returns a tuple with the Snmpv1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpv1

`func (o *GetServerServices200Response1DataInner) SetSnmpv1(v Snmpv1Bean)`

SetSnmpv1 sets Snmpv1 field to given value.

### HasSnmpv1

`func (o *GetServerServices200Response1DataInner) HasSnmpv1() bool`

HasSnmpv1 returns a boolean if a field has been set.

### GetSnmpv2c

`func (o *GetServerServices200Response1DataInner) GetSnmpv2c() Snmpv2cBean`

GetSnmpv2c returns the Snmpv2c field if non-nil, zero value otherwise.

### GetSnmpv2cOk

`func (o *GetServerServices200Response1DataInner) GetSnmpv2cOk() (*Snmpv2cBean, bool)`

GetSnmpv2cOk returns a tuple with the Snmpv2c field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpv2c

`func (o *GetServerServices200Response1DataInner) SetSnmpv2c(v Snmpv2cBean)`

SetSnmpv2c sets Snmpv2c field to given value.

### HasSnmpv2c

`func (o *GetServerServices200Response1DataInner) HasSnmpv2c() bool`

HasSnmpv2c returns a boolean if a field has been set.

### GetSnmpv3

`func (o *GetServerServices200Response1DataInner) GetSnmpv3() Snmpv3Bean`

GetSnmpv3 returns the Snmpv3 field if non-nil, zero value otherwise.

### GetSnmpv3Ok

`func (o *GetServerServices200Response1DataInner) GetSnmpv3Ok() (*Snmpv3Bean, bool)`

GetSnmpv3Ok returns a tuple with the Snmpv3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpv3

`func (o *GetServerServices200Response1DataInner) SetSnmpv3(v Snmpv3Bean)`

SetSnmpv3 sets Snmpv3 field to given value.

### HasSnmpv3

`func (o *GetServerServices200Response1DataInner) HasSnmpv3() bool`

HasSnmpv3 returns a boolean if a field has been set.

### GetTrapServers

`func (o *GetServerServices200Response1DataInner) GetTrapServers() []SnmpTrapServerBean`

GetTrapServers returns the TrapServers field if non-nil, zero value otherwise.

### GetTrapServersOk

`func (o *GetServerServices200Response1DataInner) GetTrapServersOk() (*[]SnmpTrapServerBean, bool)`

GetTrapServersOk returns a tuple with the TrapServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrapServers

`func (o *GetServerServices200Response1DataInner) SetTrapServers(v []SnmpTrapServerBean)`

SetTrapServers sets TrapServers field to given value.

### HasTrapServers

`func (o *GetServerServices200Response1DataInner) HasTrapServers() bool`

HasTrapServers returns a boolean if a field has been set.

### GetTacacsPlusAuthentication

`func (o *GetServerServices200Response1DataInner) GetTacacsPlusAuthentication() TacacsPlusAuthenticationBean`

GetTacacsPlusAuthentication returns the TacacsPlusAuthentication field if non-nil, zero value otherwise.

### GetTacacsPlusAuthenticationOk

`func (o *GetServerServices200Response1DataInner) GetTacacsPlusAuthenticationOk() (*TacacsPlusAuthenticationBean, bool)`

GetTacacsPlusAuthenticationOk returns a tuple with the TacacsPlusAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTacacsPlusAuthentication

`func (o *GetServerServices200Response1DataInner) SetTacacsPlusAuthentication(v TacacsPlusAuthenticationBean)`

SetTacacsPlusAuthentication sets TacacsPlusAuthentication field to given value.

### HasTacacsPlusAuthentication

`func (o *GetServerServices200Response1DataInner) HasTacacsPlusAuthentication() bool`

HasTacacsPlusAuthentication returns a boolean if a field has been set.

### GetIso8601TimestampsEnabled

`func (o *GetServerServices200Response1DataInner) GetIso8601TimestampsEnabled() bool`

GetIso8601TimestampsEnabled returns the Iso8601TimestampsEnabled field if non-nil, zero value otherwise.

### GetIso8601TimestampsEnabledOk

`func (o *GetServerServices200Response1DataInner) GetIso8601TimestampsEnabledOk() (*bool, bool)`

GetIso8601TimestampsEnabledOk returns a tuple with the Iso8601TimestampsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIso8601TimestampsEnabled

`func (o *GetServerServices200Response1DataInner) SetIso8601TimestampsEnabled(v bool)`

SetIso8601TimestampsEnabled sets Iso8601TimestampsEnabled field to given value.

### HasIso8601TimestampsEnabled

`func (o *GetServerServices200Response1DataInner) HasIso8601TimestampsEnabled() bool`

HasIso8601TimestampsEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


