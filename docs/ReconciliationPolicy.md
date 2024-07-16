# ReconciliationPolicy

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

## Methods

### NewReconciliationPolicy

`func NewReconciliationPolicy() *ReconciliationPolicy`

NewReconciliationPolicy instantiates a new ReconciliationPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReconciliationPolicyWithDefaults

`func NewReconciliationPolicyWithDefaults() *ReconciliationPolicy`

NewReconciliationPolicyWithDefaults instantiates a new ReconciliationPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReconciliationPolicy) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReconciliationPolicy) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReconciliationPolicy) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ReconciliationPolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ReconciliationPolicy) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ReconciliationPolicy) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ReconciliationPolicy) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ReconciliationPolicy) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *ReconciliationPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReconciliationPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReconciliationPolicy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ReconciliationPolicy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *ReconciliationPolicy) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *ReconciliationPolicy) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *ReconciliationPolicy) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *ReconciliationPolicy) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetConfiguration

`func (o *ReconciliationPolicy) GetConfiguration() InlinedConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *ReconciliationPolicy) GetConfigurationOk() (*InlinedConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *ReconciliationPolicy) SetConfiguration(v InlinedConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *ReconciliationPolicy) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetEnabled

`func (o *ReconciliationPolicy) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ReconciliationPolicy) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ReconciliationPolicy) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ReconciliationPolicy) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDiscoveryMethod

`func (o *ReconciliationPolicy) GetDiscoveryMethod() string`

GetDiscoveryMethod returns the DiscoveryMethod field if non-nil, zero value otherwise.

### GetDiscoveryMethodOk

`func (o *ReconciliationPolicy) GetDiscoveryMethodOk() (*string, bool)`

GetDiscoveryMethodOk returns a tuple with the DiscoveryMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryMethod

`func (o *ReconciliationPolicy) SetDiscoveryMethod(v string)`

SetDiscoveryMethod sets DiscoveryMethod field to given value.

### HasDiscoveryMethod

`func (o *ReconciliationPolicy) HasDiscoveryMethod() bool`

HasDiscoveryMethod returns a boolean if a field has been set.

### GetSeedRouterAddresses

`func (o *ReconciliationPolicy) GetSeedRouterAddresses() []string`

GetSeedRouterAddresses returns the SeedRouterAddresses field if non-nil, zero value otherwise.

### GetSeedRouterAddressesOk

`func (o *ReconciliationPolicy) GetSeedRouterAddressesOk() (*[]string, bool)`

GetSeedRouterAddressesOk returns a tuple with the SeedRouterAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeedRouterAddresses

`func (o *ReconciliationPolicy) SetSeedRouterAddresses(v []string)`

SetSeedRouterAddresses sets SeedRouterAddresses field to given value.

### HasSeedRouterAddresses

`func (o *ReconciliationPolicy) HasSeedRouterAddresses() bool`

HasSeedRouterAddresses returns a boolean if a field has been set.

### GetSnmp

`func (o *ReconciliationPolicy) GetSnmp() SnmpBean`

GetSnmp returns the Snmp field if non-nil, zero value otherwise.

### GetSnmpOk

`func (o *ReconciliationPolicy) GetSnmpOk() (*SnmpBean, bool)`

GetSnmpOk returns a tuple with the Snmp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmp

`func (o *ReconciliationPolicy) SetSnmp(v SnmpBean)`

SetSnmp sets Snmp field to given value.

### HasSnmp

`func (o *ReconciliationPolicy) HasSnmp() bool`

HasSnmp returns a boolean if a field has been set.

### GetSkipFqdnEnabled

`func (o *ReconciliationPolicy) GetSkipFqdnEnabled() bool`

GetSkipFqdnEnabled returns the SkipFqdnEnabled field if non-nil, zero value otherwise.

### GetSkipFqdnEnabledOk

`func (o *ReconciliationPolicy) GetSkipFqdnEnabledOk() (*bool, bool)`

GetSkipFqdnEnabledOk returns a tuple with the SkipFqdnEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipFqdnEnabled

`func (o *ReconciliationPolicy) SetSkipFqdnEnabled(v bool)`

SetSkipFqdnEnabled sets SkipFqdnEnabled field to given value.

### HasSkipFqdnEnabled

`func (o *ReconciliationPolicy) HasSkipFqdnEnabled() bool`

HasSkipFqdnEnabled returns a boolean if a field has been set.

### GetDnsServers

`func (o *ReconciliationPolicy) GetDnsServers() []string`

GetDnsServers returns the DnsServers field if non-nil, zero value otherwise.

### GetDnsServersOk

`func (o *ReconciliationPolicy) GetDnsServersOk() (*[]string, bool)`

GetDnsServersOk returns a tuple with the DnsServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServers

`func (o *ReconciliationPolicy) SetDnsServers(v []string)`

SetDnsServers sets DnsServers field to given value.

### HasDnsServers

`func (o *ReconciliationPolicy) HasDnsServers() bool`

HasDnsServers returns a boolean if a field has been set.

### GetBlackHoleVlan

`func (o *ReconciliationPolicy) GetBlackHoleVlan() int32`

GetBlackHoleVlan returns the BlackHoleVlan field if non-nil, zero value otherwise.

### GetBlackHoleVlanOk

`func (o *ReconciliationPolicy) GetBlackHoleVlanOk() (*int32, bool)`

GetBlackHoleVlanOk returns a tuple with the BlackHoleVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlackHoleVlan

`func (o *ReconciliationPolicy) SetBlackHoleVlan(v int32)`

SetBlackHoleVlan sets BlackHoleVlan field to given value.

### HasBlackHoleVlan

`func (o *ReconciliationPolicy) HasBlackHoleVlan() bool`

HasBlackHoleVlan returns a boolean if a field has been set.

### GetTrunkDefaultVlan

`func (o *ReconciliationPolicy) GetTrunkDefaultVlan() int32`

GetTrunkDefaultVlan returns the TrunkDefaultVlan field if non-nil, zero value otherwise.

### GetTrunkDefaultVlanOk

`func (o *ReconciliationPolicy) GetTrunkDefaultVlanOk() (*int32, bool)`

GetTrunkDefaultVlanOk returns a tuple with the TrunkDefaultVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrunkDefaultVlan

`func (o *ReconciliationPolicy) SetTrunkDefaultVlan(v int32)`

SetTrunkDefaultVlan sets TrunkDefaultVlan field to given value.

### HasTrunkDefaultVlan

`func (o *ReconciliationPolicy) HasTrunkDefaultVlan() bool`

HasTrunkDefaultVlan returns a boolean if a field has been set.

### GetStartDateTime

`func (o *ReconciliationPolicy) GetStartDateTime() time.Time`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *ReconciliationPolicy) GetStartDateTimeOk() (*time.Time, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *ReconciliationPolicy) SetStartDateTime(v time.Time)`

SetStartDateTime sets StartDateTime field to given value.

### HasStartDateTime

`func (o *ReconciliationPolicy) HasStartDateTime() bool`

HasStartDateTime returns a boolean if a field has been set.

### GetFrequency

`func (o *ReconciliationPolicy) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *ReconciliationPolicy) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *ReconciliationPolicy) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *ReconciliationPolicy) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


