# NetworkInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The resource type. | [optional] 
**ManagementAddress** | Pointer to **string** | The IP address configured on the eth0 interface in the DNS/DHCP Administration Console. If dedicated management is enabled, The IP address for the dedicated management interface (eth2). On a multi-interface DNS/DHCP Server, the interface is used for all management traffic such as CS, PSM, SNMP, and SSH. | [optional] 
**ServicesIpv4Address** | Pointer to **string** | The IPv4 address for the services interface. On a multi-interface DNS/DHCP Server, the interface is used for all service traffic such as DNS, DHCP, and TFTP services. | [optional] [readonly] 
**ServicesIpv4PrefixLength** | Pointer to **int32** | The IPv4 prefix length for the services interface. | [optional] [readonly] 
**ServicesIpv6Address** | Pointer to **string** | The IPv6 address for the services interface. | [optional] [readonly] 
**ServicesIpv6PrefixLength** | Pointer to **int32** | The IPv6 prefix length for the services interface. | [optional] [readonly] 
**HaBackboneAddress** | Pointer to **string** | The IP address of the backbone connection if configured as a member of a high availability pair. | [optional] [readonly] 
**HaBackbonePrefixLength** | Pointer to **int32** | The subnet prefix length of the high availability backbone. | [optional] [readonly] 
**RedundancyEnabled** | Pointer to **bool** | Indicates whether redundancy is enabled through port bonding. | [optional] [readonly] 
**RedundancyScenario** | Pointer to **string** | The type of redundancy scenario configured. | [optional] 

## Methods

### NewNetworkInterface

`func NewNetworkInterface() *NetworkInterface`

NewNetworkInterface instantiates a new NetworkInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkInterfaceWithDefaults

`func NewNetworkInterfaceWithDefaults() *NetworkInterface`

NewNetworkInterfaceWithDefaults instantiates a new NetworkInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *NetworkInterface) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NetworkInterface) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NetworkInterface) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NetworkInterface) HasType() bool`

HasType returns a boolean if a field has been set.

### GetManagementAddress

`func (o *NetworkInterface) GetManagementAddress() string`

GetManagementAddress returns the ManagementAddress field if non-nil, zero value otherwise.

### GetManagementAddressOk

`func (o *NetworkInterface) GetManagementAddressOk() (*string, bool)`

GetManagementAddressOk returns a tuple with the ManagementAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementAddress

`func (o *NetworkInterface) SetManagementAddress(v string)`

SetManagementAddress sets ManagementAddress field to given value.

### HasManagementAddress

`func (o *NetworkInterface) HasManagementAddress() bool`

HasManagementAddress returns a boolean if a field has been set.

### GetServicesIpv4Address

`func (o *NetworkInterface) GetServicesIpv4Address() string`

GetServicesIpv4Address returns the ServicesIpv4Address field if non-nil, zero value otherwise.

### GetServicesIpv4AddressOk

`func (o *NetworkInterface) GetServicesIpv4AddressOk() (*string, bool)`

GetServicesIpv4AddressOk returns a tuple with the ServicesIpv4Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicesIpv4Address

`func (o *NetworkInterface) SetServicesIpv4Address(v string)`

SetServicesIpv4Address sets ServicesIpv4Address field to given value.

### HasServicesIpv4Address

`func (o *NetworkInterface) HasServicesIpv4Address() bool`

HasServicesIpv4Address returns a boolean if a field has been set.

### GetServicesIpv4PrefixLength

`func (o *NetworkInterface) GetServicesIpv4PrefixLength() int32`

GetServicesIpv4PrefixLength returns the ServicesIpv4PrefixLength field if non-nil, zero value otherwise.

### GetServicesIpv4PrefixLengthOk

`func (o *NetworkInterface) GetServicesIpv4PrefixLengthOk() (*int32, bool)`

GetServicesIpv4PrefixLengthOk returns a tuple with the ServicesIpv4PrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicesIpv4PrefixLength

`func (o *NetworkInterface) SetServicesIpv4PrefixLength(v int32)`

SetServicesIpv4PrefixLength sets ServicesIpv4PrefixLength field to given value.

### HasServicesIpv4PrefixLength

`func (o *NetworkInterface) HasServicesIpv4PrefixLength() bool`

HasServicesIpv4PrefixLength returns a boolean if a field has been set.

### GetServicesIpv6Address

`func (o *NetworkInterface) GetServicesIpv6Address() string`

GetServicesIpv6Address returns the ServicesIpv6Address field if non-nil, zero value otherwise.

### GetServicesIpv6AddressOk

`func (o *NetworkInterface) GetServicesIpv6AddressOk() (*string, bool)`

GetServicesIpv6AddressOk returns a tuple with the ServicesIpv6Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicesIpv6Address

`func (o *NetworkInterface) SetServicesIpv6Address(v string)`

SetServicesIpv6Address sets ServicesIpv6Address field to given value.

### HasServicesIpv6Address

`func (o *NetworkInterface) HasServicesIpv6Address() bool`

HasServicesIpv6Address returns a boolean if a field has been set.

### GetServicesIpv6PrefixLength

`func (o *NetworkInterface) GetServicesIpv6PrefixLength() int32`

GetServicesIpv6PrefixLength returns the ServicesIpv6PrefixLength field if non-nil, zero value otherwise.

### GetServicesIpv6PrefixLengthOk

`func (o *NetworkInterface) GetServicesIpv6PrefixLengthOk() (*int32, bool)`

GetServicesIpv6PrefixLengthOk returns a tuple with the ServicesIpv6PrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicesIpv6PrefixLength

`func (o *NetworkInterface) SetServicesIpv6PrefixLength(v int32)`

SetServicesIpv6PrefixLength sets ServicesIpv6PrefixLength field to given value.

### HasServicesIpv6PrefixLength

`func (o *NetworkInterface) HasServicesIpv6PrefixLength() bool`

HasServicesIpv6PrefixLength returns a boolean if a field has been set.

### GetHaBackboneAddress

`func (o *NetworkInterface) GetHaBackboneAddress() string`

GetHaBackboneAddress returns the HaBackboneAddress field if non-nil, zero value otherwise.

### GetHaBackboneAddressOk

`func (o *NetworkInterface) GetHaBackboneAddressOk() (*string, bool)`

GetHaBackboneAddressOk returns a tuple with the HaBackboneAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaBackboneAddress

`func (o *NetworkInterface) SetHaBackboneAddress(v string)`

SetHaBackboneAddress sets HaBackboneAddress field to given value.

### HasHaBackboneAddress

`func (o *NetworkInterface) HasHaBackboneAddress() bool`

HasHaBackboneAddress returns a boolean if a field has been set.

### GetHaBackbonePrefixLength

`func (o *NetworkInterface) GetHaBackbonePrefixLength() int32`

GetHaBackbonePrefixLength returns the HaBackbonePrefixLength field if non-nil, zero value otherwise.

### GetHaBackbonePrefixLengthOk

`func (o *NetworkInterface) GetHaBackbonePrefixLengthOk() (*int32, bool)`

GetHaBackbonePrefixLengthOk returns a tuple with the HaBackbonePrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaBackbonePrefixLength

`func (o *NetworkInterface) SetHaBackbonePrefixLength(v int32)`

SetHaBackbonePrefixLength sets HaBackbonePrefixLength field to given value.

### HasHaBackbonePrefixLength

`func (o *NetworkInterface) HasHaBackbonePrefixLength() bool`

HasHaBackbonePrefixLength returns a boolean if a field has been set.

### GetRedundancyEnabled

`func (o *NetworkInterface) GetRedundancyEnabled() bool`

GetRedundancyEnabled returns the RedundancyEnabled field if non-nil, zero value otherwise.

### GetRedundancyEnabledOk

`func (o *NetworkInterface) GetRedundancyEnabledOk() (*bool, bool)`

GetRedundancyEnabledOk returns a tuple with the RedundancyEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancyEnabled

`func (o *NetworkInterface) SetRedundancyEnabled(v bool)`

SetRedundancyEnabled sets RedundancyEnabled field to given value.

### HasRedundancyEnabled

`func (o *NetworkInterface) HasRedundancyEnabled() bool`

HasRedundancyEnabled returns a boolean if a field has been set.

### GetRedundancyScenario

`func (o *NetworkInterface) GetRedundancyScenario() string`

GetRedundancyScenario returns the RedundancyScenario field if non-nil, zero value otherwise.

### GetRedundancyScenarioOk

`func (o *NetworkInterface) GetRedundancyScenarioOk() (*string, bool)`

GetRedundancyScenarioOk returns a tuple with the RedundancyScenario field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancyScenario

`func (o *NetworkInterface) SetRedundancyScenario(v string)`

SetRedundancyScenario sets RedundancyScenario field to given value.

### HasRedundancyScenario

`func (o *NetworkInterface) HasRedundancyScenario() bool`

HasRedundancyScenario returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


