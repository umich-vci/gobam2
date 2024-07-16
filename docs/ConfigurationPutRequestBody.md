# ConfigurationPutRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | **string** | The name of the resource. | 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**ConfigurationGroup** | Pointer to **string** | The configuration group that the configuration is part of. | [optional] 
**Description** | Pointer to **string** | The description of the configuration. | [optional] 
**Snmp** | Pointer to [**SnmpBean**](SnmpBean.md) |  | [optional] 
**ServerMonitoringEnabled** | **bool** | Indicates whether monitoring service is enabled for the configuration. | 
**DhcpConfigurationValidationEnabled** | **bool** | Indicates whether the syntax of the dhcpd.conf file is validated before data is deployed from Address Manager. | 
**DnsConfigurationValidationEnabled** | **bool** | Indicates whether the syntax of the named.conf file is validated before data is deployed from Address Manager. | 
**DnsZoneValidationEnabled** | **bool** | Indicates whether the syntax of each DNS zone file is validated before data is deployed from Address Manager. | 
**CheckIntegrityValidation** | **string** | Indicates the type of DNS zone validation check performed on the configuration. | 
**CheckMxCnameValidation** | **string** | This option checks that MX records point to a CNAME record rather than an A or AAAA record. This is equivalent to setting the -M switch for the named-checkzone tool. | 
**CheckMxValidation** | **string** | This option checks that MX records point to an IP address rather than an A or AAAA record. This is equivalent to setting the -m switch for the named-checkzone tool. | 
**CheckNamesValidation** | **string** | This option checks that A, AAAA, and MX record names are legal hostnames. It also checks that domain names in the RDATA of NS, SOA, and MX records are legal. This is equivalent to setting the -k switch for the named-checkzone tool. | 
**CheckNsValidation** | **string** | This option checks that NS records point to an IP address rather than an A or AAAA record. This is equivalent to setting the -n switch for the named-checkzone tool. | 
**CheckSrvCnameValidation** | **string** | This option checks that SRV records point to a CNAME record rather than A or AAAA record. This is equivalent to setting the -S switch for the named-checkzone tool. | 
**CheckWildcardValidation** | **string** | This option checks for wildcards in zone names that don&#39;t appear as the left-most segment of a zone name: for example, mail.*.example.com. Non-terminal wildcards are permissible, but you may want to be alerted to their presence. This is equivalent to setting the -W switch for the named-checkzone tool. | 
**DnsOptionInheritanceEnabled** | **bool** | Controls whether DNS options are inherited by child zones within a configuration. When DNS option inheritance is disabled, DNS options that have been configured on a zone aren&#39;t inherited by the child zone. In the reverse space, DNS options that have been configured on a block aren&#39;t inherited by the child block or network when DNS option inheritance is disabled. | 
**DnsFeedEnabled** | **bool** | Defines whether the BlueCat Security Feed is enabled. | 
**DnsFeedLicense** | Pointer to **string** | The BlueCat Security Feed license. | [optional] 
**BlockUsageCalculation** | **string** | Defines the IP space use statistics to keep track of the amount of space available in your block. | 
**SharedNetworkTagGroup** | Pointer to [**InlinedTagGroup**](InlinedTagGroup.md) |  | [optional] 
**KeyAutoRegenerationEnabled** | **bool** | Indicates whether Address Manager will regenerate ZSKs and KSKs. When a zone is signed with a DNSSEC signing policy, Address Manager enables this option automatically. | 

## Methods

### NewConfigurationPutRequestBody

`func NewConfigurationPutRequestBody(name string, serverMonitoringEnabled bool, dhcpConfigurationValidationEnabled bool, dnsConfigurationValidationEnabled bool, dnsZoneValidationEnabled bool, checkIntegrityValidation string, checkMxCnameValidation string, checkMxValidation string, checkNamesValidation string, checkNsValidation string, checkSrvCnameValidation string, checkWildcardValidation string, dnsOptionInheritanceEnabled bool, dnsFeedEnabled bool, blockUsageCalculation string, keyAutoRegenerationEnabled bool, ) *ConfigurationPutRequestBody`

NewConfigurationPutRequestBody instantiates a new ConfigurationPutRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigurationPutRequestBodyWithDefaults

`func NewConfigurationPutRequestBodyWithDefaults() *ConfigurationPutRequestBody`

NewConfigurationPutRequestBodyWithDefaults instantiates a new ConfigurationPutRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConfigurationPutRequestBody) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConfigurationPutRequestBody) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConfigurationPutRequestBody) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ConfigurationPutRequestBody) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ConfigurationPutRequestBody) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConfigurationPutRequestBody) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConfigurationPutRequestBody) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ConfigurationPutRequestBody) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *ConfigurationPutRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConfigurationPutRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConfigurationPutRequestBody) SetName(v string)`

SetName sets Name field to given value.


### GetUserDefinedFields

`func (o *ConfigurationPutRequestBody) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *ConfigurationPutRequestBody) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *ConfigurationPutRequestBody) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *ConfigurationPutRequestBody) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetConfigurationGroup

`func (o *ConfigurationPutRequestBody) GetConfigurationGroup() string`

GetConfigurationGroup returns the ConfigurationGroup field if non-nil, zero value otherwise.

### GetConfigurationGroupOk

`func (o *ConfigurationPutRequestBody) GetConfigurationGroupOk() (*string, bool)`

GetConfigurationGroupOk returns a tuple with the ConfigurationGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationGroup

`func (o *ConfigurationPutRequestBody) SetConfigurationGroup(v string)`

SetConfigurationGroup sets ConfigurationGroup field to given value.

### HasConfigurationGroup

`func (o *ConfigurationPutRequestBody) HasConfigurationGroup() bool`

HasConfigurationGroup returns a boolean if a field has been set.

### GetDescription

`func (o *ConfigurationPutRequestBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConfigurationPutRequestBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConfigurationPutRequestBody) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConfigurationPutRequestBody) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSnmp

`func (o *ConfigurationPutRequestBody) GetSnmp() SnmpBean`

GetSnmp returns the Snmp field if non-nil, zero value otherwise.

### GetSnmpOk

`func (o *ConfigurationPutRequestBody) GetSnmpOk() (*SnmpBean, bool)`

GetSnmpOk returns a tuple with the Snmp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmp

`func (o *ConfigurationPutRequestBody) SetSnmp(v SnmpBean)`

SetSnmp sets Snmp field to given value.

### HasSnmp

`func (o *ConfigurationPutRequestBody) HasSnmp() bool`

HasSnmp returns a boolean if a field has been set.

### GetServerMonitoringEnabled

`func (o *ConfigurationPutRequestBody) GetServerMonitoringEnabled() bool`

GetServerMonitoringEnabled returns the ServerMonitoringEnabled field if non-nil, zero value otherwise.

### GetServerMonitoringEnabledOk

`func (o *ConfigurationPutRequestBody) GetServerMonitoringEnabledOk() (*bool, bool)`

GetServerMonitoringEnabledOk returns a tuple with the ServerMonitoringEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerMonitoringEnabled

`func (o *ConfigurationPutRequestBody) SetServerMonitoringEnabled(v bool)`

SetServerMonitoringEnabled sets ServerMonitoringEnabled field to given value.


### GetDhcpConfigurationValidationEnabled

`func (o *ConfigurationPutRequestBody) GetDhcpConfigurationValidationEnabled() bool`

GetDhcpConfigurationValidationEnabled returns the DhcpConfigurationValidationEnabled field if non-nil, zero value otherwise.

### GetDhcpConfigurationValidationEnabledOk

`func (o *ConfigurationPutRequestBody) GetDhcpConfigurationValidationEnabledOk() (*bool, bool)`

GetDhcpConfigurationValidationEnabledOk returns a tuple with the DhcpConfigurationValidationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpConfigurationValidationEnabled

`func (o *ConfigurationPutRequestBody) SetDhcpConfigurationValidationEnabled(v bool)`

SetDhcpConfigurationValidationEnabled sets DhcpConfigurationValidationEnabled field to given value.


### GetDnsConfigurationValidationEnabled

`func (o *ConfigurationPutRequestBody) GetDnsConfigurationValidationEnabled() bool`

GetDnsConfigurationValidationEnabled returns the DnsConfigurationValidationEnabled field if non-nil, zero value otherwise.

### GetDnsConfigurationValidationEnabledOk

`func (o *ConfigurationPutRequestBody) GetDnsConfigurationValidationEnabledOk() (*bool, bool)`

GetDnsConfigurationValidationEnabledOk returns a tuple with the DnsConfigurationValidationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsConfigurationValidationEnabled

`func (o *ConfigurationPutRequestBody) SetDnsConfigurationValidationEnabled(v bool)`

SetDnsConfigurationValidationEnabled sets DnsConfigurationValidationEnabled field to given value.


### GetDnsZoneValidationEnabled

`func (o *ConfigurationPutRequestBody) GetDnsZoneValidationEnabled() bool`

GetDnsZoneValidationEnabled returns the DnsZoneValidationEnabled field if non-nil, zero value otherwise.

### GetDnsZoneValidationEnabledOk

`func (o *ConfigurationPutRequestBody) GetDnsZoneValidationEnabledOk() (*bool, bool)`

GetDnsZoneValidationEnabledOk returns a tuple with the DnsZoneValidationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsZoneValidationEnabled

`func (o *ConfigurationPutRequestBody) SetDnsZoneValidationEnabled(v bool)`

SetDnsZoneValidationEnabled sets DnsZoneValidationEnabled field to given value.


### GetCheckIntegrityValidation

`func (o *ConfigurationPutRequestBody) GetCheckIntegrityValidation() string`

GetCheckIntegrityValidation returns the CheckIntegrityValidation field if non-nil, zero value otherwise.

### GetCheckIntegrityValidationOk

`func (o *ConfigurationPutRequestBody) GetCheckIntegrityValidationOk() (*string, bool)`

GetCheckIntegrityValidationOk returns a tuple with the CheckIntegrityValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckIntegrityValidation

`func (o *ConfigurationPutRequestBody) SetCheckIntegrityValidation(v string)`

SetCheckIntegrityValidation sets CheckIntegrityValidation field to given value.


### GetCheckMxCnameValidation

`func (o *ConfigurationPutRequestBody) GetCheckMxCnameValidation() string`

GetCheckMxCnameValidation returns the CheckMxCnameValidation field if non-nil, zero value otherwise.

### GetCheckMxCnameValidationOk

`func (o *ConfigurationPutRequestBody) GetCheckMxCnameValidationOk() (*string, bool)`

GetCheckMxCnameValidationOk returns a tuple with the CheckMxCnameValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckMxCnameValidation

`func (o *ConfigurationPutRequestBody) SetCheckMxCnameValidation(v string)`

SetCheckMxCnameValidation sets CheckMxCnameValidation field to given value.


### GetCheckMxValidation

`func (o *ConfigurationPutRequestBody) GetCheckMxValidation() string`

GetCheckMxValidation returns the CheckMxValidation field if non-nil, zero value otherwise.

### GetCheckMxValidationOk

`func (o *ConfigurationPutRequestBody) GetCheckMxValidationOk() (*string, bool)`

GetCheckMxValidationOk returns a tuple with the CheckMxValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckMxValidation

`func (o *ConfigurationPutRequestBody) SetCheckMxValidation(v string)`

SetCheckMxValidation sets CheckMxValidation field to given value.


### GetCheckNamesValidation

`func (o *ConfigurationPutRequestBody) GetCheckNamesValidation() string`

GetCheckNamesValidation returns the CheckNamesValidation field if non-nil, zero value otherwise.

### GetCheckNamesValidationOk

`func (o *ConfigurationPutRequestBody) GetCheckNamesValidationOk() (*string, bool)`

GetCheckNamesValidationOk returns a tuple with the CheckNamesValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckNamesValidation

`func (o *ConfigurationPutRequestBody) SetCheckNamesValidation(v string)`

SetCheckNamesValidation sets CheckNamesValidation field to given value.


### GetCheckNsValidation

`func (o *ConfigurationPutRequestBody) GetCheckNsValidation() string`

GetCheckNsValidation returns the CheckNsValidation field if non-nil, zero value otherwise.

### GetCheckNsValidationOk

`func (o *ConfigurationPutRequestBody) GetCheckNsValidationOk() (*string, bool)`

GetCheckNsValidationOk returns a tuple with the CheckNsValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckNsValidation

`func (o *ConfigurationPutRequestBody) SetCheckNsValidation(v string)`

SetCheckNsValidation sets CheckNsValidation field to given value.


### GetCheckSrvCnameValidation

`func (o *ConfigurationPutRequestBody) GetCheckSrvCnameValidation() string`

GetCheckSrvCnameValidation returns the CheckSrvCnameValidation field if non-nil, zero value otherwise.

### GetCheckSrvCnameValidationOk

`func (o *ConfigurationPutRequestBody) GetCheckSrvCnameValidationOk() (*string, bool)`

GetCheckSrvCnameValidationOk returns a tuple with the CheckSrvCnameValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckSrvCnameValidation

`func (o *ConfigurationPutRequestBody) SetCheckSrvCnameValidation(v string)`

SetCheckSrvCnameValidation sets CheckSrvCnameValidation field to given value.


### GetCheckWildcardValidation

`func (o *ConfigurationPutRequestBody) GetCheckWildcardValidation() string`

GetCheckWildcardValidation returns the CheckWildcardValidation field if non-nil, zero value otherwise.

### GetCheckWildcardValidationOk

`func (o *ConfigurationPutRequestBody) GetCheckWildcardValidationOk() (*string, bool)`

GetCheckWildcardValidationOk returns a tuple with the CheckWildcardValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckWildcardValidation

`func (o *ConfigurationPutRequestBody) SetCheckWildcardValidation(v string)`

SetCheckWildcardValidation sets CheckWildcardValidation field to given value.


### GetDnsOptionInheritanceEnabled

`func (o *ConfigurationPutRequestBody) GetDnsOptionInheritanceEnabled() bool`

GetDnsOptionInheritanceEnabled returns the DnsOptionInheritanceEnabled field if non-nil, zero value otherwise.

### GetDnsOptionInheritanceEnabledOk

`func (o *ConfigurationPutRequestBody) GetDnsOptionInheritanceEnabledOk() (*bool, bool)`

GetDnsOptionInheritanceEnabledOk returns a tuple with the DnsOptionInheritanceEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsOptionInheritanceEnabled

`func (o *ConfigurationPutRequestBody) SetDnsOptionInheritanceEnabled(v bool)`

SetDnsOptionInheritanceEnabled sets DnsOptionInheritanceEnabled field to given value.


### GetDnsFeedEnabled

`func (o *ConfigurationPutRequestBody) GetDnsFeedEnabled() bool`

GetDnsFeedEnabled returns the DnsFeedEnabled field if non-nil, zero value otherwise.

### GetDnsFeedEnabledOk

`func (o *ConfigurationPutRequestBody) GetDnsFeedEnabledOk() (*bool, bool)`

GetDnsFeedEnabledOk returns a tuple with the DnsFeedEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsFeedEnabled

`func (o *ConfigurationPutRequestBody) SetDnsFeedEnabled(v bool)`

SetDnsFeedEnabled sets DnsFeedEnabled field to given value.


### GetDnsFeedLicense

`func (o *ConfigurationPutRequestBody) GetDnsFeedLicense() string`

GetDnsFeedLicense returns the DnsFeedLicense field if non-nil, zero value otherwise.

### GetDnsFeedLicenseOk

`func (o *ConfigurationPutRequestBody) GetDnsFeedLicenseOk() (*string, bool)`

GetDnsFeedLicenseOk returns a tuple with the DnsFeedLicense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsFeedLicense

`func (o *ConfigurationPutRequestBody) SetDnsFeedLicense(v string)`

SetDnsFeedLicense sets DnsFeedLicense field to given value.

### HasDnsFeedLicense

`func (o *ConfigurationPutRequestBody) HasDnsFeedLicense() bool`

HasDnsFeedLicense returns a boolean if a field has been set.

### GetBlockUsageCalculation

`func (o *ConfigurationPutRequestBody) GetBlockUsageCalculation() string`

GetBlockUsageCalculation returns the BlockUsageCalculation field if non-nil, zero value otherwise.

### GetBlockUsageCalculationOk

`func (o *ConfigurationPutRequestBody) GetBlockUsageCalculationOk() (*string, bool)`

GetBlockUsageCalculationOk returns a tuple with the BlockUsageCalculation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockUsageCalculation

`func (o *ConfigurationPutRequestBody) SetBlockUsageCalculation(v string)`

SetBlockUsageCalculation sets BlockUsageCalculation field to given value.


### GetSharedNetworkTagGroup

`func (o *ConfigurationPutRequestBody) GetSharedNetworkTagGroup() InlinedTagGroup`

GetSharedNetworkTagGroup returns the SharedNetworkTagGroup field if non-nil, zero value otherwise.

### GetSharedNetworkTagGroupOk

`func (o *ConfigurationPutRequestBody) GetSharedNetworkTagGroupOk() (*InlinedTagGroup, bool)`

GetSharedNetworkTagGroupOk returns a tuple with the SharedNetworkTagGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedNetworkTagGroup

`func (o *ConfigurationPutRequestBody) SetSharedNetworkTagGroup(v InlinedTagGroup)`

SetSharedNetworkTagGroup sets SharedNetworkTagGroup field to given value.

### HasSharedNetworkTagGroup

`func (o *ConfigurationPutRequestBody) HasSharedNetworkTagGroup() bool`

HasSharedNetworkTagGroup returns a boolean if a field has been set.

### GetKeyAutoRegenerationEnabled

`func (o *ConfigurationPutRequestBody) GetKeyAutoRegenerationEnabled() bool`

GetKeyAutoRegenerationEnabled returns the KeyAutoRegenerationEnabled field if non-nil, zero value otherwise.

### GetKeyAutoRegenerationEnabledOk

`func (o *ConfigurationPutRequestBody) GetKeyAutoRegenerationEnabledOk() (*bool, bool)`

GetKeyAutoRegenerationEnabledOk returns a tuple with the KeyAutoRegenerationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyAutoRegenerationEnabled

`func (o *ConfigurationPutRequestBody) SetKeyAutoRegenerationEnabled(v bool)`

SetKeyAutoRegenerationEnabled sets KeyAutoRegenerationEnabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


