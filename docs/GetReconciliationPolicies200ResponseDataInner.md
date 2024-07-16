# GetReconciliationPolicies200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | Pointer to **string** | The name of the resource. | [optional] 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**Configuration** | Pointer to [**InlinedConfiguration**](InlinedConfiguration.md) |  | [optional] [readonly] 
**Enabled** | Pointer to **bool** | Indicates whether the reconciliation policy will run at the scheduled time. | [optional] 
**DiscoveryMethod** | Pointer to **string** | The discovery method for the reconciliation policy. | [optional] 
**SeedRouterAddresses** | Pointer to **[]string** | The IP address or addresses of the router or layer 3 switch where the network discovery operation is to start. | [optional] 
**Snmp** | Pointer to [**SnmpBean**](SnmpBean.md) |  | [optional] 
**SkipFqdnEnabled** | Pointer to **bool** | Indicates whether the Address Manager discovery engine will skip FQDN and reverse DNS lookups. | [optional] 
**DnsServers** | Pointer to **[]string** | The IP address or addresses for the DNS server(s) that the discovery engine will use to perform FQDN and reverse DNS lookups. | [optional] 
**BlackHoleVlan** | Pointer to **int32** | The VLAN ID for the black hole VLAN. The black hole VLAN is used as a default VLAN for all unused ports. | [optional] 
**TrunkDefaultVlan** | Pointer to **int32** | The VLAN ID assigned to a trunk as a native/default VLAN to protect controlled traffic from being spoofed. | [optional] 
**StartDateTime** | Pointer to **time.Time** | The start time for the reconciliation policy discovery. | [optional] 
**Frequency** | Pointer to **string** | The recurring frequency for reconciliation policy discovery. | [optional] 
**NetworkBoundaries** | Pointer to **[]string** | The IPv4 range(s) in CIDR format for the reconciliation policy to search. | [optional] 
**PingSweepNetworkGaps** | Pointer to **[]string** | The IPv4 address ranges in CIDR format for which ping sweep sends an ICMP echo request. | [optional] 
**Overrides** | Pointer to **[]string** | The IPv4 addresses and ranges for the reconciliation policy to ignore. The list of overrides can contain single addresses, ranges in CIDR format (e.g. 10.0.0.0/16), or ranges in address range format (e.g. 10.0.0.0-10.0.255.255). | [optional] 
**ReclaimAcceptanceAge** | Pointer to **string** | The age parameter for automatic reconciliation of reclaimable addresses. Reconciliation will be applied to reclaimable IP addresses older than this value. Reclaimable addresses are addresses that exist in Address Manager but are not found on the physical network. | [optional] 
**UnknownAcceptanceAge** | Pointer to **string** | The age parameter for automatic reconciliation of unknown addresses. Reconciliation will be applied to unknown IP addresses older than this value. Unknown addresses are addresses that exist on the physical network but do not exist in Address Manager. | [optional] 
**MismatchAcceptanceAge** | Pointer to **string** | The age parameter for automatic reconciliation of mismatched addresses. Reconciliation will be applied to mismatched IP addresses older than this value. Mismatched addresses are addresses that exist in both Address Manager and the physical network, but do not match in either the MAC address, DNS host name info, VLAN info, or connect switch port info. | [optional] 
**View** | Pointer to [**InlinedView**](InlinedView.md) |  | [optional] 
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 

## Methods

### NewGetReconciliationPolicies200ResponseDataInner

`func NewGetReconciliationPolicies200ResponseDataInner() *GetReconciliationPolicies200ResponseDataInner`

NewGetReconciliationPolicies200ResponseDataInner instantiates a new GetReconciliationPolicies200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetReconciliationPolicies200ResponseDataInnerWithDefaults

`func NewGetReconciliationPolicies200ResponseDataInnerWithDefaults() *GetReconciliationPolicies200ResponseDataInner`

NewGetReconciliationPolicies200ResponseDataInnerWithDefaults instantiates a new GetReconciliationPolicies200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetReconciliationPolicies200ResponseDataInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetReconciliationPolicies200ResponseDataInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetReconciliationPolicies200ResponseDataInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetReconciliationPolicies200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GetReconciliationPolicies200ResponseDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetReconciliationPolicies200ResponseDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetReconciliationPolicies200ResponseDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetReconciliationPolicies200ResponseDataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *GetReconciliationPolicies200ResponseDataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetReconciliationPolicies200ResponseDataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetReconciliationPolicies200ResponseDataInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetReconciliationPolicies200ResponseDataInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *GetReconciliationPolicies200ResponseDataInner) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *GetReconciliationPolicies200ResponseDataInner) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *GetReconciliationPolicies200ResponseDataInner) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *GetReconciliationPolicies200ResponseDataInner) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetConfiguration

`func (o *GetReconciliationPolicies200ResponseDataInner) GetConfiguration() InlinedConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *GetReconciliationPolicies200ResponseDataInner) GetConfigurationOk() (*InlinedConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *GetReconciliationPolicies200ResponseDataInner) SetConfiguration(v InlinedConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *GetReconciliationPolicies200ResponseDataInner) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetEnabled

`func (o *GetReconciliationPolicies200ResponseDataInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetReconciliationPolicies200ResponseDataInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetReconciliationPolicies200ResponseDataInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetReconciliationPolicies200ResponseDataInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDiscoveryMethod

`func (o *GetReconciliationPolicies200ResponseDataInner) GetDiscoveryMethod() string`

GetDiscoveryMethod returns the DiscoveryMethod field if non-nil, zero value otherwise.

### GetDiscoveryMethodOk

`func (o *GetReconciliationPolicies200ResponseDataInner) GetDiscoveryMethodOk() (*string, bool)`

GetDiscoveryMethodOk returns a tuple with the DiscoveryMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryMethod

`func (o *GetReconciliationPolicies200ResponseDataInner) SetDiscoveryMethod(v string)`

SetDiscoveryMethod sets DiscoveryMethod field to given value.

### HasDiscoveryMethod

`func (o *GetReconciliationPolicies200ResponseDataInner) HasDiscoveryMethod() bool`

HasDiscoveryMethod returns a boolean if a field has been set.

### GetSeedRouterAddresses

`func (o *GetReconciliationPolicies200ResponseDataInner) GetSeedRouterAddresses() []string`

GetSeedRouterAddresses returns the SeedRouterAddresses field if non-nil, zero value otherwise.

### GetSeedRouterAddressesOk

`func (o *GetReconciliationPolicies200ResponseDataInner) GetSeedRouterAddressesOk() (*[]string, bool)`

GetSeedRouterAddressesOk returns a tuple with the SeedRouterAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeedRouterAddresses

`func (o *GetReconciliationPolicies200ResponseDataInner) SetSeedRouterAddresses(v []string)`

SetSeedRouterAddresses sets SeedRouterAddresses field to given value.

### HasSeedRouterAddresses

`func (o *GetReconciliationPolicies200ResponseDataInner) HasSeedRouterAddresses() bool`

HasSeedRouterAddresses returns a boolean if a field has been set.

### GetSnmp

`func (o *GetReconciliationPolicies200ResponseDataInner) GetSnmp() SnmpBean`

GetSnmp returns the Snmp field if non-nil, zero value otherwise.

### GetSnmpOk

`func (o *GetReconciliationPolicies200ResponseDataInner) GetSnmpOk() (*SnmpBean, bool)`

GetSnmpOk returns a tuple with the Snmp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmp

`func (o *GetReconciliationPolicies200ResponseDataInner) SetSnmp(v SnmpBean)`

SetSnmp sets Snmp field to given value.

### HasSnmp

`func (o *GetReconciliationPolicies200ResponseDataInner) HasSnmp() bool`

HasSnmp returns a boolean if a field has been set.

### GetSkipFqdnEnabled

`func (o *GetReconciliationPolicies200ResponseDataInner) GetSkipFqdnEnabled() bool`

GetSkipFqdnEnabled returns the SkipFqdnEnabled field if non-nil, zero value otherwise.

### GetSkipFqdnEnabledOk

`func (o *GetReconciliationPolicies200ResponseDataInner) GetSkipFqdnEnabledOk() (*bool, bool)`

GetSkipFqdnEnabledOk returns a tuple with the SkipFqdnEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipFqdnEnabled

`func (o *GetReconciliationPolicies200ResponseDataInner) SetSkipFqdnEnabled(v bool)`

SetSkipFqdnEnabled sets SkipFqdnEnabled field to given value.

### HasSkipFqdnEnabled

`func (o *GetReconciliationPolicies200ResponseDataInner) HasSkipFqdnEnabled() bool`

HasSkipFqdnEnabled returns a boolean if a field has been set.

### GetDnsServers

`func (o *GetReconciliationPolicies200ResponseDataInner) GetDnsServers() []string`

GetDnsServers returns the DnsServers field if non-nil, zero value otherwise.

### GetDnsServersOk

`func (o *GetReconciliationPolicies200ResponseDataInner) GetDnsServersOk() (*[]string, bool)`

GetDnsServersOk returns a tuple with the DnsServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServers

`func (o *GetReconciliationPolicies200ResponseDataInner) SetDnsServers(v []string)`

SetDnsServers sets DnsServers field to given value.

### HasDnsServers

`func (o *GetReconciliationPolicies200ResponseDataInner) HasDnsServers() bool`

HasDnsServers returns a boolean if a field has been set.

### GetBlackHoleVlan

`func (o *GetReconciliationPolicies200ResponseDataInner) GetBlackHoleVlan() int32`

GetBlackHoleVlan returns the BlackHoleVlan field if non-nil, zero value otherwise.

### GetBlackHoleVlanOk

`func (o *GetReconciliationPolicies200ResponseDataInner) GetBlackHoleVlanOk() (*int32, bool)`

GetBlackHoleVlanOk returns a tuple with the BlackHoleVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlackHoleVlan

`func (o *GetReconciliationPolicies200ResponseDataInner) SetBlackHoleVlan(v int32)`

SetBlackHoleVlan sets BlackHoleVlan field to given value.

### HasBlackHoleVlan

`func (o *GetReconciliationPolicies200ResponseDataInner) HasBlackHoleVlan() bool`

HasBlackHoleVlan returns a boolean if a field has been set.

### GetTrunkDefaultVlan

`func (o *GetReconciliationPolicies200ResponseDataInner) GetTrunkDefaultVlan() int32`

GetTrunkDefaultVlan returns the TrunkDefaultVlan field if non-nil, zero value otherwise.

### GetTrunkDefaultVlanOk

`func (o *GetReconciliationPolicies200ResponseDataInner) GetTrunkDefaultVlanOk() (*int32, bool)`

GetTrunkDefaultVlanOk returns a tuple with the TrunkDefaultVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrunkDefaultVlan

`func (o *GetReconciliationPolicies200ResponseDataInner) SetTrunkDefaultVlan(v int32)`

SetTrunkDefaultVlan sets TrunkDefaultVlan field to given value.

### HasTrunkDefaultVlan

`func (o *GetReconciliationPolicies200ResponseDataInner) HasTrunkDefaultVlan() bool`

HasTrunkDefaultVlan returns a boolean if a field has been set.

### GetStartDateTime

`func (o *GetReconciliationPolicies200ResponseDataInner) GetStartDateTime() time.Time`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *GetReconciliationPolicies200ResponseDataInner) GetStartDateTimeOk() (*time.Time, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *GetReconciliationPolicies200ResponseDataInner) SetStartDateTime(v time.Time)`

SetStartDateTime sets StartDateTime field to given value.

### HasStartDateTime

`func (o *GetReconciliationPolicies200ResponseDataInner) HasStartDateTime() bool`

HasStartDateTime returns a boolean if a field has been set.

### GetFrequency

`func (o *GetReconciliationPolicies200ResponseDataInner) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *GetReconciliationPolicies200ResponseDataInner) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *GetReconciliationPolicies200ResponseDataInner) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *GetReconciliationPolicies200ResponseDataInner) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### GetNetworkBoundaries

`func (o *GetReconciliationPolicies200ResponseDataInner) GetNetworkBoundaries() []string`

GetNetworkBoundaries returns the NetworkBoundaries field if non-nil, zero value otherwise.

### GetNetworkBoundariesOk

`func (o *GetReconciliationPolicies200ResponseDataInner) GetNetworkBoundariesOk() (*[]string, bool)`

GetNetworkBoundariesOk returns a tuple with the NetworkBoundaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkBoundaries

`func (o *GetReconciliationPolicies200ResponseDataInner) SetNetworkBoundaries(v []string)`

SetNetworkBoundaries sets NetworkBoundaries field to given value.

### HasNetworkBoundaries

`func (o *GetReconciliationPolicies200ResponseDataInner) HasNetworkBoundaries() bool`

HasNetworkBoundaries returns a boolean if a field has been set.

### GetPingSweepNetworkGaps

`func (o *GetReconciliationPolicies200ResponseDataInner) GetPingSweepNetworkGaps() []string`

GetPingSweepNetworkGaps returns the PingSweepNetworkGaps field if non-nil, zero value otherwise.

### GetPingSweepNetworkGapsOk

`func (o *GetReconciliationPolicies200ResponseDataInner) GetPingSweepNetworkGapsOk() (*[]string, bool)`

GetPingSweepNetworkGapsOk returns a tuple with the PingSweepNetworkGaps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingSweepNetworkGaps

`func (o *GetReconciliationPolicies200ResponseDataInner) SetPingSweepNetworkGaps(v []string)`

SetPingSweepNetworkGaps sets PingSweepNetworkGaps field to given value.

### HasPingSweepNetworkGaps

`func (o *GetReconciliationPolicies200ResponseDataInner) HasPingSweepNetworkGaps() bool`

HasPingSweepNetworkGaps returns a boolean if a field has been set.

### GetOverrides

`func (o *GetReconciliationPolicies200ResponseDataInner) GetOverrides() []string`

GetOverrides returns the Overrides field if non-nil, zero value otherwise.

### GetOverridesOk

`func (o *GetReconciliationPolicies200ResponseDataInner) GetOverridesOk() (*[]string, bool)`

GetOverridesOk returns a tuple with the Overrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrides

`func (o *GetReconciliationPolicies200ResponseDataInner) SetOverrides(v []string)`

SetOverrides sets Overrides field to given value.

### HasOverrides

`func (o *GetReconciliationPolicies200ResponseDataInner) HasOverrides() bool`

HasOverrides returns a boolean if a field has been set.

### GetReclaimAcceptanceAge

`func (o *GetReconciliationPolicies200ResponseDataInner) GetReclaimAcceptanceAge() string`

GetReclaimAcceptanceAge returns the ReclaimAcceptanceAge field if non-nil, zero value otherwise.

### GetReclaimAcceptanceAgeOk

`func (o *GetReconciliationPolicies200ResponseDataInner) GetReclaimAcceptanceAgeOk() (*string, bool)`

GetReclaimAcceptanceAgeOk returns a tuple with the ReclaimAcceptanceAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReclaimAcceptanceAge

`func (o *GetReconciliationPolicies200ResponseDataInner) SetReclaimAcceptanceAge(v string)`

SetReclaimAcceptanceAge sets ReclaimAcceptanceAge field to given value.

### HasReclaimAcceptanceAge

`func (o *GetReconciliationPolicies200ResponseDataInner) HasReclaimAcceptanceAge() bool`

HasReclaimAcceptanceAge returns a boolean if a field has been set.

### GetUnknownAcceptanceAge

`func (o *GetReconciliationPolicies200ResponseDataInner) GetUnknownAcceptanceAge() string`

GetUnknownAcceptanceAge returns the UnknownAcceptanceAge field if non-nil, zero value otherwise.

### GetUnknownAcceptanceAgeOk

`func (o *GetReconciliationPolicies200ResponseDataInner) GetUnknownAcceptanceAgeOk() (*string, bool)`

GetUnknownAcceptanceAgeOk returns a tuple with the UnknownAcceptanceAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownAcceptanceAge

`func (o *GetReconciliationPolicies200ResponseDataInner) SetUnknownAcceptanceAge(v string)`

SetUnknownAcceptanceAge sets UnknownAcceptanceAge field to given value.

### HasUnknownAcceptanceAge

`func (o *GetReconciliationPolicies200ResponseDataInner) HasUnknownAcceptanceAge() bool`

HasUnknownAcceptanceAge returns a boolean if a field has been set.

### GetMismatchAcceptanceAge

`func (o *GetReconciliationPolicies200ResponseDataInner) GetMismatchAcceptanceAge() string`

GetMismatchAcceptanceAge returns the MismatchAcceptanceAge field if non-nil, zero value otherwise.

### GetMismatchAcceptanceAgeOk

`func (o *GetReconciliationPolicies200ResponseDataInner) GetMismatchAcceptanceAgeOk() (*string, bool)`

GetMismatchAcceptanceAgeOk returns a tuple with the MismatchAcceptanceAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMismatchAcceptanceAge

`func (o *GetReconciliationPolicies200ResponseDataInner) SetMismatchAcceptanceAge(v string)`

SetMismatchAcceptanceAge sets MismatchAcceptanceAge field to given value.

### HasMismatchAcceptanceAge

`func (o *GetReconciliationPolicies200ResponseDataInner) HasMismatchAcceptanceAge() bool`

HasMismatchAcceptanceAge returns a boolean if a field has been set.

### GetView

`func (o *GetReconciliationPolicies200ResponseDataInner) GetView() InlinedView`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *GetReconciliationPolicies200ResponseDataInner) GetViewOk() (*InlinedView, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *GetReconciliationPolicies200ResponseDataInner) SetView(v InlinedView)`

SetView sets View field to given value.

### HasView

`func (o *GetReconciliationPolicies200ResponseDataInner) HasView() bool`

HasView returns a boolean if a field has been set.

### GetLinks

`func (o *GetReconciliationPolicies200ResponseDataInner) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetReconciliationPolicies200ResponseDataInner) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetReconciliationPolicies200ResponseDataInner) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetReconciliationPolicies200ResponseDataInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


