# PostCollectionReconciliationPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | **string** | The resource type. | 
**Name** | Pointer to **string** | The name of the resource. | [optional] 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**Configuration** | Pointer to [**InlinedConfiguration**](InlinedConfiguration.md) |  | [optional] [readonly] 
**Enabled** | Pointer to **bool** | Indicates whether the reconciliation policy will run at the scheduled time. | [optional] 
**DiscoveryMethod** | **string** | The discovery method for the reconciliation policy. | 
**SeedRouterAddresses** | Pointer to **[]string** | The IP address or addresses of the router or layer 3 switch where the network discovery operation is to start. | [optional] 
**Snmp** | Pointer to [**SnmpBean**](SnmpBean.md) |  | [optional] 
**SkipFqdnEnabled** | Pointer to **bool** | Indicates whether the Address Manager discovery engine will skip FQDN and reverse DNS lookups. | [optional] 
**DnsServers** | Pointer to **[]string** | The IP address or addresses for the DNS server(s) that the discovery engine will use to perform FQDN and reverse DNS lookups. | [optional] 
**BlackHoleVlan** | Pointer to **int32** | The VLAN ID for the black hole VLAN. The black hole VLAN is used as a default VLAN for all unused ports. | [optional] 
**TrunkDefaultVlan** | Pointer to **int32** | The VLAN ID assigned to a trunk as a native/default VLAN to protect controlled traffic from being spoofed. | [optional] 
**StartDateTime** | **time.Time** | The start time for the reconciliation policy discovery. | 
**Frequency** | Pointer to **string** | The recurring frequency for reconciliation policy discovery. | [optional] 

## Methods

### NewPostCollectionReconciliationPolicyRequest

`func NewPostCollectionReconciliationPolicyRequest(type_ string, discoveryMethod string, startDateTime time.Time, ) *PostCollectionReconciliationPolicyRequest`

NewPostCollectionReconciliationPolicyRequest instantiates a new PostCollectionReconciliationPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostCollectionReconciliationPolicyRequestWithDefaults

`func NewPostCollectionReconciliationPolicyRequestWithDefaults() *PostCollectionReconciliationPolicyRequest`

NewPostCollectionReconciliationPolicyRequestWithDefaults instantiates a new PostCollectionReconciliationPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PostCollectionReconciliationPolicyRequest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostCollectionReconciliationPolicyRequest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostCollectionReconciliationPolicyRequest) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *PostCollectionReconciliationPolicyRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *PostCollectionReconciliationPolicyRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PostCollectionReconciliationPolicyRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PostCollectionReconciliationPolicyRequest) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *PostCollectionReconciliationPolicyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostCollectionReconciliationPolicyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostCollectionReconciliationPolicyRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PostCollectionReconciliationPolicyRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *PostCollectionReconciliationPolicyRequest) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *PostCollectionReconciliationPolicyRequest) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *PostCollectionReconciliationPolicyRequest) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *PostCollectionReconciliationPolicyRequest) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetConfiguration

`func (o *PostCollectionReconciliationPolicyRequest) GetConfiguration() InlinedConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *PostCollectionReconciliationPolicyRequest) GetConfigurationOk() (*InlinedConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *PostCollectionReconciliationPolicyRequest) SetConfiguration(v InlinedConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *PostCollectionReconciliationPolicyRequest) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetEnabled

`func (o *PostCollectionReconciliationPolicyRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PostCollectionReconciliationPolicyRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PostCollectionReconciliationPolicyRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PostCollectionReconciliationPolicyRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDiscoveryMethod

`func (o *PostCollectionReconciliationPolicyRequest) GetDiscoveryMethod() string`

GetDiscoveryMethod returns the DiscoveryMethod field if non-nil, zero value otherwise.

### GetDiscoveryMethodOk

`func (o *PostCollectionReconciliationPolicyRequest) GetDiscoveryMethodOk() (*string, bool)`

GetDiscoveryMethodOk returns a tuple with the DiscoveryMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryMethod

`func (o *PostCollectionReconciliationPolicyRequest) SetDiscoveryMethod(v string)`

SetDiscoveryMethod sets DiscoveryMethod field to given value.


### GetSeedRouterAddresses

`func (o *PostCollectionReconciliationPolicyRequest) GetSeedRouterAddresses() []string`

GetSeedRouterAddresses returns the SeedRouterAddresses field if non-nil, zero value otherwise.

### GetSeedRouterAddressesOk

`func (o *PostCollectionReconciliationPolicyRequest) GetSeedRouterAddressesOk() (*[]string, bool)`

GetSeedRouterAddressesOk returns a tuple with the SeedRouterAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeedRouterAddresses

`func (o *PostCollectionReconciliationPolicyRequest) SetSeedRouterAddresses(v []string)`

SetSeedRouterAddresses sets SeedRouterAddresses field to given value.

### HasSeedRouterAddresses

`func (o *PostCollectionReconciliationPolicyRequest) HasSeedRouterAddresses() bool`

HasSeedRouterAddresses returns a boolean if a field has been set.

### GetSnmp

`func (o *PostCollectionReconciliationPolicyRequest) GetSnmp() SnmpBean`

GetSnmp returns the Snmp field if non-nil, zero value otherwise.

### GetSnmpOk

`func (o *PostCollectionReconciliationPolicyRequest) GetSnmpOk() (*SnmpBean, bool)`

GetSnmpOk returns a tuple with the Snmp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmp

`func (o *PostCollectionReconciliationPolicyRequest) SetSnmp(v SnmpBean)`

SetSnmp sets Snmp field to given value.

### HasSnmp

`func (o *PostCollectionReconciliationPolicyRequest) HasSnmp() bool`

HasSnmp returns a boolean if a field has been set.

### GetSkipFqdnEnabled

`func (o *PostCollectionReconciliationPolicyRequest) GetSkipFqdnEnabled() bool`

GetSkipFqdnEnabled returns the SkipFqdnEnabled field if non-nil, zero value otherwise.

### GetSkipFqdnEnabledOk

`func (o *PostCollectionReconciliationPolicyRequest) GetSkipFqdnEnabledOk() (*bool, bool)`

GetSkipFqdnEnabledOk returns a tuple with the SkipFqdnEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipFqdnEnabled

`func (o *PostCollectionReconciliationPolicyRequest) SetSkipFqdnEnabled(v bool)`

SetSkipFqdnEnabled sets SkipFqdnEnabled field to given value.

### HasSkipFqdnEnabled

`func (o *PostCollectionReconciliationPolicyRequest) HasSkipFqdnEnabled() bool`

HasSkipFqdnEnabled returns a boolean if a field has been set.

### GetDnsServers

`func (o *PostCollectionReconciliationPolicyRequest) GetDnsServers() []string`

GetDnsServers returns the DnsServers field if non-nil, zero value otherwise.

### GetDnsServersOk

`func (o *PostCollectionReconciliationPolicyRequest) GetDnsServersOk() (*[]string, bool)`

GetDnsServersOk returns a tuple with the DnsServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServers

`func (o *PostCollectionReconciliationPolicyRequest) SetDnsServers(v []string)`

SetDnsServers sets DnsServers field to given value.

### HasDnsServers

`func (o *PostCollectionReconciliationPolicyRequest) HasDnsServers() bool`

HasDnsServers returns a boolean if a field has been set.

### GetBlackHoleVlan

`func (o *PostCollectionReconciliationPolicyRequest) GetBlackHoleVlan() int32`

GetBlackHoleVlan returns the BlackHoleVlan field if non-nil, zero value otherwise.

### GetBlackHoleVlanOk

`func (o *PostCollectionReconciliationPolicyRequest) GetBlackHoleVlanOk() (*int32, bool)`

GetBlackHoleVlanOk returns a tuple with the BlackHoleVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlackHoleVlan

`func (o *PostCollectionReconciliationPolicyRequest) SetBlackHoleVlan(v int32)`

SetBlackHoleVlan sets BlackHoleVlan field to given value.

### HasBlackHoleVlan

`func (o *PostCollectionReconciliationPolicyRequest) HasBlackHoleVlan() bool`

HasBlackHoleVlan returns a boolean if a field has been set.

### GetTrunkDefaultVlan

`func (o *PostCollectionReconciliationPolicyRequest) GetTrunkDefaultVlan() int32`

GetTrunkDefaultVlan returns the TrunkDefaultVlan field if non-nil, zero value otherwise.

### GetTrunkDefaultVlanOk

`func (o *PostCollectionReconciliationPolicyRequest) GetTrunkDefaultVlanOk() (*int32, bool)`

GetTrunkDefaultVlanOk returns a tuple with the TrunkDefaultVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrunkDefaultVlan

`func (o *PostCollectionReconciliationPolicyRequest) SetTrunkDefaultVlan(v int32)`

SetTrunkDefaultVlan sets TrunkDefaultVlan field to given value.

### HasTrunkDefaultVlan

`func (o *PostCollectionReconciliationPolicyRequest) HasTrunkDefaultVlan() bool`

HasTrunkDefaultVlan returns a boolean if a field has been set.

### GetStartDateTime

`func (o *PostCollectionReconciliationPolicyRequest) GetStartDateTime() time.Time`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *PostCollectionReconciliationPolicyRequest) GetStartDateTimeOk() (*time.Time, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *PostCollectionReconciliationPolicyRequest) SetStartDateTime(v time.Time)`

SetStartDateTime sets StartDateTime field to given value.


### GetFrequency

`func (o *PostCollectionReconciliationPolicyRequest) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *PostCollectionReconciliationPolicyRequest) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *PostCollectionReconciliationPolicyRequest) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *PostCollectionReconciliationPolicyRequest) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


