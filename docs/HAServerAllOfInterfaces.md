# HAServerAllOfInterfaces

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | Pointer to **string** | The name of the resource. | [optional] 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**Configuration** | Pointer to [**InlinedConfiguration**](InlinedConfiguration.md) |  | [optional] [readonly] 
**Server** | Pointer to [**AbstractServer**](AbstractServer.md) |  | [optional] [readonly] 
**SecondaryAddress** | Pointer to **string** | The secondary IP address of the interface. | [optional] 
**ManagementAddress** | Pointer to **string** | The IP address configured on the eth0 interface in the DNS/DHCP Administration Console. If dedicated management is enabled, The IP address for the dedicated management interface (eth2). On a multi-interface DNS/DHCP Server, the interface is used for all management traffic such as CS, PSM, SNMP, and SSH. | [optional] 
**ServicesIpv4Address** | **string** | The IPv4 address for the services interface. On a multi-interface DNS/DHCP Server, the interface is used for all service traffic such as DNS, DHCP, and TFTP services. | [readonly] 
**ServicesIpv4PrefixLength** | Pointer to **int32** | The IPv4 prefix length for the services interface. | [optional] [readonly] 
**ServicesIpv6Address** | **string** | The IPv6 address for the services interface. | [readonly] 
**ServicesIpv6PrefixLength** | Pointer to **int32** | The IPv6 prefix length for the services interface. | [optional] [readonly] 
**HaBackboneAddress** | Pointer to **string** | The IP address of the backbone connection if configured as a member of a high availability pair. | [optional] [readonly] 
**HaBackbonePrefixLength** | Pointer to **int32** | The subnet prefix length of the high availability backbone. | [optional] [readonly] 
**RedundancyEnabled** | Pointer to **bool** | Indicates whether redundancy is enabled through port bonding. | [optional] [readonly] 
**RedundancyScenario** | Pointer to **string** | The type of redundancy scenario configured. | [optional] 

## Methods

### NewHAServerAllOfInterfaces

`func NewHAServerAllOfInterfaces(servicesIpv4Address string, servicesIpv6Address string, ) *HAServerAllOfInterfaces`

NewHAServerAllOfInterfaces instantiates a new HAServerAllOfInterfaces object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHAServerAllOfInterfacesWithDefaults

`func NewHAServerAllOfInterfacesWithDefaults() *HAServerAllOfInterfaces`

NewHAServerAllOfInterfacesWithDefaults instantiates a new HAServerAllOfInterfaces object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HAServerAllOfInterfaces) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HAServerAllOfInterfaces) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HAServerAllOfInterfaces) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *HAServerAllOfInterfaces) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *HAServerAllOfInterfaces) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HAServerAllOfInterfaces) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HAServerAllOfInterfaces) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *HAServerAllOfInterfaces) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *HAServerAllOfInterfaces) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HAServerAllOfInterfaces) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HAServerAllOfInterfaces) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HAServerAllOfInterfaces) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *HAServerAllOfInterfaces) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *HAServerAllOfInterfaces) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *HAServerAllOfInterfaces) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *HAServerAllOfInterfaces) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetConfiguration

`func (o *HAServerAllOfInterfaces) GetConfiguration() InlinedConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *HAServerAllOfInterfaces) GetConfigurationOk() (*InlinedConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *HAServerAllOfInterfaces) SetConfiguration(v InlinedConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *HAServerAllOfInterfaces) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetServer

`func (o *HAServerAllOfInterfaces) GetServer() AbstractServer`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *HAServerAllOfInterfaces) GetServerOk() (*AbstractServer, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *HAServerAllOfInterfaces) SetServer(v AbstractServer)`

SetServer sets Server field to given value.

### HasServer

`func (o *HAServerAllOfInterfaces) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetSecondaryAddress

`func (o *HAServerAllOfInterfaces) GetSecondaryAddress() string`

GetSecondaryAddress returns the SecondaryAddress field if non-nil, zero value otherwise.

### GetSecondaryAddressOk

`func (o *HAServerAllOfInterfaces) GetSecondaryAddressOk() (*string, bool)`

GetSecondaryAddressOk returns a tuple with the SecondaryAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryAddress

`func (o *HAServerAllOfInterfaces) SetSecondaryAddress(v string)`

SetSecondaryAddress sets SecondaryAddress field to given value.

### HasSecondaryAddress

`func (o *HAServerAllOfInterfaces) HasSecondaryAddress() bool`

HasSecondaryAddress returns a boolean if a field has been set.

### GetManagementAddress

`func (o *HAServerAllOfInterfaces) GetManagementAddress() string`

GetManagementAddress returns the ManagementAddress field if non-nil, zero value otherwise.

### GetManagementAddressOk

`func (o *HAServerAllOfInterfaces) GetManagementAddressOk() (*string, bool)`

GetManagementAddressOk returns a tuple with the ManagementAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementAddress

`func (o *HAServerAllOfInterfaces) SetManagementAddress(v string)`

SetManagementAddress sets ManagementAddress field to given value.

### HasManagementAddress

`func (o *HAServerAllOfInterfaces) HasManagementAddress() bool`

HasManagementAddress returns a boolean if a field has been set.

### GetServicesIpv4Address

`func (o *HAServerAllOfInterfaces) GetServicesIpv4Address() string`

GetServicesIpv4Address returns the ServicesIpv4Address field if non-nil, zero value otherwise.

### GetServicesIpv4AddressOk

`func (o *HAServerAllOfInterfaces) GetServicesIpv4AddressOk() (*string, bool)`

GetServicesIpv4AddressOk returns a tuple with the ServicesIpv4Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicesIpv4Address

`func (o *HAServerAllOfInterfaces) SetServicesIpv4Address(v string)`

SetServicesIpv4Address sets ServicesIpv4Address field to given value.


### GetServicesIpv4PrefixLength

`func (o *HAServerAllOfInterfaces) GetServicesIpv4PrefixLength() int32`

GetServicesIpv4PrefixLength returns the ServicesIpv4PrefixLength field if non-nil, zero value otherwise.

### GetServicesIpv4PrefixLengthOk

`func (o *HAServerAllOfInterfaces) GetServicesIpv4PrefixLengthOk() (*int32, bool)`

GetServicesIpv4PrefixLengthOk returns a tuple with the ServicesIpv4PrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicesIpv4PrefixLength

`func (o *HAServerAllOfInterfaces) SetServicesIpv4PrefixLength(v int32)`

SetServicesIpv4PrefixLength sets ServicesIpv4PrefixLength field to given value.

### HasServicesIpv4PrefixLength

`func (o *HAServerAllOfInterfaces) HasServicesIpv4PrefixLength() bool`

HasServicesIpv4PrefixLength returns a boolean if a field has been set.

### GetServicesIpv6Address

`func (o *HAServerAllOfInterfaces) GetServicesIpv6Address() string`

GetServicesIpv6Address returns the ServicesIpv6Address field if non-nil, zero value otherwise.

### GetServicesIpv6AddressOk

`func (o *HAServerAllOfInterfaces) GetServicesIpv6AddressOk() (*string, bool)`

GetServicesIpv6AddressOk returns a tuple with the ServicesIpv6Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicesIpv6Address

`func (o *HAServerAllOfInterfaces) SetServicesIpv6Address(v string)`

SetServicesIpv6Address sets ServicesIpv6Address field to given value.


### GetServicesIpv6PrefixLength

`func (o *HAServerAllOfInterfaces) GetServicesIpv6PrefixLength() int32`

GetServicesIpv6PrefixLength returns the ServicesIpv6PrefixLength field if non-nil, zero value otherwise.

### GetServicesIpv6PrefixLengthOk

`func (o *HAServerAllOfInterfaces) GetServicesIpv6PrefixLengthOk() (*int32, bool)`

GetServicesIpv6PrefixLengthOk returns a tuple with the ServicesIpv6PrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicesIpv6PrefixLength

`func (o *HAServerAllOfInterfaces) SetServicesIpv6PrefixLength(v int32)`

SetServicesIpv6PrefixLength sets ServicesIpv6PrefixLength field to given value.

### HasServicesIpv6PrefixLength

`func (o *HAServerAllOfInterfaces) HasServicesIpv6PrefixLength() bool`

HasServicesIpv6PrefixLength returns a boolean if a field has been set.

### GetHaBackboneAddress

`func (o *HAServerAllOfInterfaces) GetHaBackboneAddress() string`

GetHaBackboneAddress returns the HaBackboneAddress field if non-nil, zero value otherwise.

### GetHaBackboneAddressOk

`func (o *HAServerAllOfInterfaces) GetHaBackboneAddressOk() (*string, bool)`

GetHaBackboneAddressOk returns a tuple with the HaBackboneAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaBackboneAddress

`func (o *HAServerAllOfInterfaces) SetHaBackboneAddress(v string)`

SetHaBackboneAddress sets HaBackboneAddress field to given value.

### HasHaBackboneAddress

`func (o *HAServerAllOfInterfaces) HasHaBackboneAddress() bool`

HasHaBackboneAddress returns a boolean if a field has been set.

### GetHaBackbonePrefixLength

`func (o *HAServerAllOfInterfaces) GetHaBackbonePrefixLength() int32`

GetHaBackbonePrefixLength returns the HaBackbonePrefixLength field if non-nil, zero value otherwise.

### GetHaBackbonePrefixLengthOk

`func (o *HAServerAllOfInterfaces) GetHaBackbonePrefixLengthOk() (*int32, bool)`

GetHaBackbonePrefixLengthOk returns a tuple with the HaBackbonePrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaBackbonePrefixLength

`func (o *HAServerAllOfInterfaces) SetHaBackbonePrefixLength(v int32)`

SetHaBackbonePrefixLength sets HaBackbonePrefixLength field to given value.

### HasHaBackbonePrefixLength

`func (o *HAServerAllOfInterfaces) HasHaBackbonePrefixLength() bool`

HasHaBackbonePrefixLength returns a boolean if a field has been set.

### GetRedundancyEnabled

`func (o *HAServerAllOfInterfaces) GetRedundancyEnabled() bool`

GetRedundancyEnabled returns the RedundancyEnabled field if non-nil, zero value otherwise.

### GetRedundancyEnabledOk

`func (o *HAServerAllOfInterfaces) GetRedundancyEnabledOk() (*bool, bool)`

GetRedundancyEnabledOk returns a tuple with the RedundancyEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancyEnabled

`func (o *HAServerAllOfInterfaces) SetRedundancyEnabled(v bool)`

SetRedundancyEnabled sets RedundancyEnabled field to given value.

### HasRedundancyEnabled

`func (o *HAServerAllOfInterfaces) HasRedundancyEnabled() bool`

HasRedundancyEnabled returns a boolean if a field has been set.

### GetRedundancyScenario

`func (o *HAServerAllOfInterfaces) GetRedundancyScenario() string`

GetRedundancyScenario returns the RedundancyScenario field if non-nil, zero value otherwise.

### GetRedundancyScenarioOk

`func (o *HAServerAllOfInterfaces) GetRedundancyScenarioOk() (*string, bool)`

GetRedundancyScenarioOk returns a tuple with the RedundancyScenario field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancyScenario

`func (o *HAServerAllOfInterfaces) SetRedundancyScenario(v string)`

SetRedundancyScenario sets RedundancyScenario field to given value.

### HasRedundancyScenario

`func (o *HAServerAllOfInterfaces) HasRedundancyScenario() bool`

HasRedundancyScenario returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


