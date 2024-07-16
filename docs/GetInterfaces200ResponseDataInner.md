# GetInterfaces200ResponseDataInner

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
**ServicesIpv4Address** | Pointer to **string** | The IPv4 address for the services interface. On a multi-interface DNS/DHCP Server, the interface is used for all service traffic such as DNS, DHCP, and TFTP services. | [optional] [readonly] 
**ServicesIpv4PrefixLength** | Pointer to **int32** | The IPv4 prefix length for the services interface. | [optional] [readonly] 
**ServicesIpv6Address** | Pointer to **string** | The IPv6 address for the services interface. | [optional] [readonly] 
**ServicesIpv6PrefixLength** | Pointer to **int32** | The IPv6 prefix length for the services interface. | [optional] [readonly] 
**HaBackboneAddress** | Pointer to **string** | The IP address of the backbone connection if configured as a member of a high availability pair. | [optional] [readonly] 
**HaBackbonePrefixLength** | Pointer to **int32** | The subnet prefix length of the high availability backbone. | [optional] [readonly] 
**RedundancyEnabled** | Pointer to **bool** | Indicates whether redundancy is enabled through port bonding. | [optional] [readonly] 
**RedundancyScenario** | Pointer to **string** | The type of redundancy scenario configured. | [optional] 
**PrimaryAddress** | Pointer to **string** | The primary IP address of the published interface. | [optional] 
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 

## Methods

### NewGetInterfaces200ResponseDataInner

`func NewGetInterfaces200ResponseDataInner() *GetInterfaces200ResponseDataInner`

NewGetInterfaces200ResponseDataInner instantiates a new GetInterfaces200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInterfaces200ResponseDataInnerWithDefaults

`func NewGetInterfaces200ResponseDataInnerWithDefaults() *GetInterfaces200ResponseDataInner`

NewGetInterfaces200ResponseDataInnerWithDefaults instantiates a new GetInterfaces200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetInterfaces200ResponseDataInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetInterfaces200ResponseDataInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetInterfaces200ResponseDataInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetInterfaces200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GetInterfaces200ResponseDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetInterfaces200ResponseDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetInterfaces200ResponseDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetInterfaces200ResponseDataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *GetInterfaces200ResponseDataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetInterfaces200ResponseDataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetInterfaces200ResponseDataInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetInterfaces200ResponseDataInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *GetInterfaces200ResponseDataInner) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *GetInterfaces200ResponseDataInner) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *GetInterfaces200ResponseDataInner) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *GetInterfaces200ResponseDataInner) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetConfiguration

`func (o *GetInterfaces200ResponseDataInner) GetConfiguration() InlinedConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *GetInterfaces200ResponseDataInner) GetConfigurationOk() (*InlinedConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *GetInterfaces200ResponseDataInner) SetConfiguration(v InlinedConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *GetInterfaces200ResponseDataInner) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetServer

`func (o *GetInterfaces200ResponseDataInner) GetServer() AbstractServer`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *GetInterfaces200ResponseDataInner) GetServerOk() (*AbstractServer, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *GetInterfaces200ResponseDataInner) SetServer(v AbstractServer)`

SetServer sets Server field to given value.

### HasServer

`func (o *GetInterfaces200ResponseDataInner) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetSecondaryAddress

`func (o *GetInterfaces200ResponseDataInner) GetSecondaryAddress() string`

GetSecondaryAddress returns the SecondaryAddress field if non-nil, zero value otherwise.

### GetSecondaryAddressOk

`func (o *GetInterfaces200ResponseDataInner) GetSecondaryAddressOk() (*string, bool)`

GetSecondaryAddressOk returns a tuple with the SecondaryAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryAddress

`func (o *GetInterfaces200ResponseDataInner) SetSecondaryAddress(v string)`

SetSecondaryAddress sets SecondaryAddress field to given value.

### HasSecondaryAddress

`func (o *GetInterfaces200ResponseDataInner) HasSecondaryAddress() bool`

HasSecondaryAddress returns a boolean if a field has been set.

### GetManagementAddress

`func (o *GetInterfaces200ResponseDataInner) GetManagementAddress() string`

GetManagementAddress returns the ManagementAddress field if non-nil, zero value otherwise.

### GetManagementAddressOk

`func (o *GetInterfaces200ResponseDataInner) GetManagementAddressOk() (*string, bool)`

GetManagementAddressOk returns a tuple with the ManagementAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementAddress

`func (o *GetInterfaces200ResponseDataInner) SetManagementAddress(v string)`

SetManagementAddress sets ManagementAddress field to given value.

### HasManagementAddress

`func (o *GetInterfaces200ResponseDataInner) HasManagementAddress() bool`

HasManagementAddress returns a boolean if a field has been set.

### GetServicesIpv4Address

`func (o *GetInterfaces200ResponseDataInner) GetServicesIpv4Address() string`

GetServicesIpv4Address returns the ServicesIpv4Address field if non-nil, zero value otherwise.

### GetServicesIpv4AddressOk

`func (o *GetInterfaces200ResponseDataInner) GetServicesIpv4AddressOk() (*string, bool)`

GetServicesIpv4AddressOk returns a tuple with the ServicesIpv4Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicesIpv4Address

`func (o *GetInterfaces200ResponseDataInner) SetServicesIpv4Address(v string)`

SetServicesIpv4Address sets ServicesIpv4Address field to given value.

### HasServicesIpv4Address

`func (o *GetInterfaces200ResponseDataInner) HasServicesIpv4Address() bool`

HasServicesIpv4Address returns a boolean if a field has been set.

### GetServicesIpv4PrefixLength

`func (o *GetInterfaces200ResponseDataInner) GetServicesIpv4PrefixLength() int32`

GetServicesIpv4PrefixLength returns the ServicesIpv4PrefixLength field if non-nil, zero value otherwise.

### GetServicesIpv4PrefixLengthOk

`func (o *GetInterfaces200ResponseDataInner) GetServicesIpv4PrefixLengthOk() (*int32, bool)`

GetServicesIpv4PrefixLengthOk returns a tuple with the ServicesIpv4PrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicesIpv4PrefixLength

`func (o *GetInterfaces200ResponseDataInner) SetServicesIpv4PrefixLength(v int32)`

SetServicesIpv4PrefixLength sets ServicesIpv4PrefixLength field to given value.

### HasServicesIpv4PrefixLength

`func (o *GetInterfaces200ResponseDataInner) HasServicesIpv4PrefixLength() bool`

HasServicesIpv4PrefixLength returns a boolean if a field has been set.

### GetServicesIpv6Address

`func (o *GetInterfaces200ResponseDataInner) GetServicesIpv6Address() string`

GetServicesIpv6Address returns the ServicesIpv6Address field if non-nil, zero value otherwise.

### GetServicesIpv6AddressOk

`func (o *GetInterfaces200ResponseDataInner) GetServicesIpv6AddressOk() (*string, bool)`

GetServicesIpv6AddressOk returns a tuple with the ServicesIpv6Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicesIpv6Address

`func (o *GetInterfaces200ResponseDataInner) SetServicesIpv6Address(v string)`

SetServicesIpv6Address sets ServicesIpv6Address field to given value.

### HasServicesIpv6Address

`func (o *GetInterfaces200ResponseDataInner) HasServicesIpv6Address() bool`

HasServicesIpv6Address returns a boolean if a field has been set.

### GetServicesIpv6PrefixLength

`func (o *GetInterfaces200ResponseDataInner) GetServicesIpv6PrefixLength() int32`

GetServicesIpv6PrefixLength returns the ServicesIpv6PrefixLength field if non-nil, zero value otherwise.

### GetServicesIpv6PrefixLengthOk

`func (o *GetInterfaces200ResponseDataInner) GetServicesIpv6PrefixLengthOk() (*int32, bool)`

GetServicesIpv6PrefixLengthOk returns a tuple with the ServicesIpv6PrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicesIpv6PrefixLength

`func (o *GetInterfaces200ResponseDataInner) SetServicesIpv6PrefixLength(v int32)`

SetServicesIpv6PrefixLength sets ServicesIpv6PrefixLength field to given value.

### HasServicesIpv6PrefixLength

`func (o *GetInterfaces200ResponseDataInner) HasServicesIpv6PrefixLength() bool`

HasServicesIpv6PrefixLength returns a boolean if a field has been set.

### GetHaBackboneAddress

`func (o *GetInterfaces200ResponseDataInner) GetHaBackboneAddress() string`

GetHaBackboneAddress returns the HaBackboneAddress field if non-nil, zero value otherwise.

### GetHaBackboneAddressOk

`func (o *GetInterfaces200ResponseDataInner) GetHaBackboneAddressOk() (*string, bool)`

GetHaBackboneAddressOk returns a tuple with the HaBackboneAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaBackboneAddress

`func (o *GetInterfaces200ResponseDataInner) SetHaBackboneAddress(v string)`

SetHaBackboneAddress sets HaBackboneAddress field to given value.

### HasHaBackboneAddress

`func (o *GetInterfaces200ResponseDataInner) HasHaBackboneAddress() bool`

HasHaBackboneAddress returns a boolean if a field has been set.

### GetHaBackbonePrefixLength

`func (o *GetInterfaces200ResponseDataInner) GetHaBackbonePrefixLength() int32`

GetHaBackbonePrefixLength returns the HaBackbonePrefixLength field if non-nil, zero value otherwise.

### GetHaBackbonePrefixLengthOk

`func (o *GetInterfaces200ResponseDataInner) GetHaBackbonePrefixLengthOk() (*int32, bool)`

GetHaBackbonePrefixLengthOk returns a tuple with the HaBackbonePrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaBackbonePrefixLength

`func (o *GetInterfaces200ResponseDataInner) SetHaBackbonePrefixLength(v int32)`

SetHaBackbonePrefixLength sets HaBackbonePrefixLength field to given value.

### HasHaBackbonePrefixLength

`func (o *GetInterfaces200ResponseDataInner) HasHaBackbonePrefixLength() bool`

HasHaBackbonePrefixLength returns a boolean if a field has been set.

### GetRedundancyEnabled

`func (o *GetInterfaces200ResponseDataInner) GetRedundancyEnabled() bool`

GetRedundancyEnabled returns the RedundancyEnabled field if non-nil, zero value otherwise.

### GetRedundancyEnabledOk

`func (o *GetInterfaces200ResponseDataInner) GetRedundancyEnabledOk() (*bool, bool)`

GetRedundancyEnabledOk returns a tuple with the RedundancyEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancyEnabled

`func (o *GetInterfaces200ResponseDataInner) SetRedundancyEnabled(v bool)`

SetRedundancyEnabled sets RedundancyEnabled field to given value.

### HasRedundancyEnabled

`func (o *GetInterfaces200ResponseDataInner) HasRedundancyEnabled() bool`

HasRedundancyEnabled returns a boolean if a field has been set.

### GetRedundancyScenario

`func (o *GetInterfaces200ResponseDataInner) GetRedundancyScenario() string`

GetRedundancyScenario returns the RedundancyScenario field if non-nil, zero value otherwise.

### GetRedundancyScenarioOk

`func (o *GetInterfaces200ResponseDataInner) GetRedundancyScenarioOk() (*string, bool)`

GetRedundancyScenarioOk returns a tuple with the RedundancyScenario field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancyScenario

`func (o *GetInterfaces200ResponseDataInner) SetRedundancyScenario(v string)`

SetRedundancyScenario sets RedundancyScenario field to given value.

### HasRedundancyScenario

`func (o *GetInterfaces200ResponseDataInner) HasRedundancyScenario() bool`

HasRedundancyScenario returns a boolean if a field has been set.

### GetPrimaryAddress

`func (o *GetInterfaces200ResponseDataInner) GetPrimaryAddress() string`

GetPrimaryAddress returns the PrimaryAddress field if non-nil, zero value otherwise.

### GetPrimaryAddressOk

`func (o *GetInterfaces200ResponseDataInner) GetPrimaryAddressOk() (*string, bool)`

GetPrimaryAddressOk returns a tuple with the PrimaryAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryAddress

`func (o *GetInterfaces200ResponseDataInner) SetPrimaryAddress(v string)`

SetPrimaryAddress sets PrimaryAddress field to given value.

### HasPrimaryAddress

`func (o *GetInterfaces200ResponseDataInner) HasPrimaryAddress() bool`

HasPrimaryAddress returns a boolean if a field has been set.

### GetLinks

`func (o *GetInterfaces200ResponseDataInner) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetInterfaces200ResponseDataInner) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetInterfaces200ResponseDataInner) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetInterfaces200ResponseDataInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


