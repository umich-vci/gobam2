# ConfigurationPostRequestBody

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
**ServerMonitoringEnabled** | Pointer to **bool** | Indicates whether monitoring service is enabled for the configuration. | [optional] 
**DhcpConfigurationValidationEnabled** | Pointer to **bool** | Indicates whether the syntax of the dhcpd.conf file is validated before data is deployed from Address Manager. | [optional] 
**DnsConfigurationValidationEnabled** | Pointer to **bool** | Indicates whether the syntax of the named.conf file is validated before data is deployed from Address Manager. | [optional] 
**DnsZoneValidationEnabled** | Pointer to **bool** | Indicates whether the syntax of each DNS zone file is validated before data is deployed from Address Manager. | [optional] 
**CheckIntegrityValidation** | Pointer to **string** | Indicates the type of DNS zone validation check performed on the configuration. | [optional] 
**CheckMxCnameValidation** | Pointer to **string** | This option checks that MX records point to a CNAME record rather than an A or AAAA record. This is equivalent to setting the -M switch for the named-checkzone tool. | [optional] 
**CheckMxValidation** | Pointer to **string** | This option checks that MX records point to an IP address rather than an A or AAAA record. This is equivalent to setting the -m switch for the named-checkzone tool. | [optional] 
**CheckNamesValidation** | Pointer to **string** | This option checks that A, AAAA, and MX record names are legal hostnames. It also checks that domain names in the RDATA of NS, SOA, and MX records are legal. This is equivalent to setting the -k switch for the named-checkzone tool. | [optional] 
**CheckNsValidation** | Pointer to **string** | This option checks that NS records point to an IP address rather than an A or AAAA record. This is equivalent to setting the -n switch for the named-checkzone tool. | [optional] 
**CheckSrvCnameValidation** | Pointer to **string** | This option checks that SRV records point to a CNAME record rather than A or AAAA record. This is equivalent to setting the -S switch for the named-checkzone tool. | [optional] 
**CheckWildcardValidation** | Pointer to **string** | This option checks for wildcards in zone names that don&#39;t appear as the left-most segment of a zone name: for example, mail.*.example.com. Non-terminal wildcards are permissible, but you may want to be alerted to their presence. This is equivalent to setting the -W switch for the named-checkzone tool. | [optional] 
**DnsOptionInheritanceEnabled** | Pointer to **bool** | Controls whether DNS options are inherited by child zones within a configuration. When DNS option inheritance is disabled, DNS options that have been configured on a zone aren&#39;t inherited by the child zone. In the reverse space, DNS options that have been configured on a block aren&#39;t inherited by the child block or network when DNS option inheritance is disabled. | [optional] 
**DnsFeedEnabled** | Pointer to **bool** | Defines whether the BlueCat Security Feed is enabled. | [optional] 
**DnsFeedLicense** | Pointer to **string** | The BlueCat Security Feed license. | [optional] 
**BlockUsageCalculation** | Pointer to **string** | Defines the IP space use statistics to keep track of the amount of space available in your block. | [optional] 
**SharedNetworkTagGroup** | Pointer to [**InlinedTagGroup**](InlinedTagGroup.md) |  | [optional] 
**KeyAutoRegenerationEnabled** | Pointer to **bool** | Indicates whether Address Manager will regenerate ZSKs and KSKs. When a zone is signed with a DNSSEC signing policy, Address Manager enables this option automatically. | [optional] 

## Methods

### NewConfigurationPostRequestBody

`func NewConfigurationPostRequestBody(name string, ) *ConfigurationPostRequestBody`

NewConfigurationPostRequestBody instantiates a new ConfigurationPostRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigurationPostRequestBodyWithDefaults

`func NewConfigurationPostRequestBodyWithDefaults() *ConfigurationPostRequestBody`

NewConfigurationPostRequestBodyWithDefaults instantiates a new ConfigurationPostRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConfigurationPostRequestBody) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConfigurationPostRequestBody) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConfigurationPostRequestBody) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ConfigurationPostRequestBody) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ConfigurationPostRequestBody) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConfigurationPostRequestBody) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConfigurationPostRequestBody) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ConfigurationPostRequestBody) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *ConfigurationPostRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConfigurationPostRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConfigurationPostRequestBody) SetName(v string)`

SetName sets Name field to given value.


### GetUserDefinedFields

`func (o *ConfigurationPostRequestBody) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *ConfigurationPostRequestBody) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *ConfigurationPostRequestBody) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *ConfigurationPostRequestBody) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetConfigurationGroup

`func (o *ConfigurationPostRequestBody) GetConfigurationGroup() string`

GetConfigurationGroup returns the ConfigurationGroup field if non-nil, zero value otherwise.

### GetConfigurationGroupOk

`func (o *ConfigurationPostRequestBody) GetConfigurationGroupOk() (*string, bool)`

GetConfigurationGroupOk returns a tuple with the ConfigurationGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationGroup

`func (o *ConfigurationPostRequestBody) SetConfigurationGroup(v string)`

SetConfigurationGroup sets ConfigurationGroup field to given value.

### HasConfigurationGroup

`func (o *ConfigurationPostRequestBody) HasConfigurationGroup() bool`

HasConfigurationGroup returns a boolean if a field has been set.

### GetDescription

`func (o *ConfigurationPostRequestBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConfigurationPostRequestBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConfigurationPostRequestBody) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConfigurationPostRequestBody) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSnmp

`func (o *ConfigurationPostRequestBody) GetSnmp() SnmpBean`

GetSnmp returns the Snmp field if non-nil, zero value otherwise.

### GetSnmpOk

`func (o *ConfigurationPostRequestBody) GetSnmpOk() (*SnmpBean, bool)`

GetSnmpOk returns a tuple with the Snmp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmp

`func (o *ConfigurationPostRequestBody) SetSnmp(v SnmpBean)`

SetSnmp sets Snmp field to given value.

### HasSnmp

`func (o *ConfigurationPostRequestBody) HasSnmp() bool`

HasSnmp returns a boolean if a field has been set.

### GetServerMonitoringEnabled

`func (o *ConfigurationPostRequestBody) GetServerMonitoringEnabled() bool`

GetServerMonitoringEnabled returns the ServerMonitoringEnabled field if non-nil, zero value otherwise.

### GetServerMonitoringEnabledOk

`func (o *ConfigurationPostRequestBody) GetServerMonitoringEnabledOk() (*bool, bool)`

GetServerMonitoringEnabledOk returns a tuple with the ServerMonitoringEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerMonitoringEnabled

`func (o *ConfigurationPostRequestBody) SetServerMonitoringEnabled(v bool)`

SetServerMonitoringEnabled sets ServerMonitoringEnabled field to given value.

### HasServerMonitoringEnabled

`func (o *ConfigurationPostRequestBody) HasServerMonitoringEnabled() bool`

HasServerMonitoringEnabled returns a boolean if a field has been set.

### GetDhcpConfigurationValidationEnabled

`func (o *ConfigurationPostRequestBody) GetDhcpConfigurationValidationEnabled() bool`

GetDhcpConfigurationValidationEnabled returns the DhcpConfigurationValidationEnabled field if non-nil, zero value otherwise.

### GetDhcpConfigurationValidationEnabledOk

`func (o *ConfigurationPostRequestBody) GetDhcpConfigurationValidationEnabledOk() (*bool, bool)`

GetDhcpConfigurationValidationEnabledOk returns a tuple with the DhcpConfigurationValidationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpConfigurationValidationEnabled

`func (o *ConfigurationPostRequestBody) SetDhcpConfigurationValidationEnabled(v bool)`

SetDhcpConfigurationValidationEnabled sets DhcpConfigurationValidationEnabled field to given value.

### HasDhcpConfigurationValidationEnabled

`func (o *ConfigurationPostRequestBody) HasDhcpConfigurationValidationEnabled() bool`

HasDhcpConfigurationValidationEnabled returns a boolean if a field has been set.

### GetDnsConfigurationValidationEnabled

`func (o *ConfigurationPostRequestBody) GetDnsConfigurationValidationEnabled() bool`

GetDnsConfigurationValidationEnabled returns the DnsConfigurationValidationEnabled field if non-nil, zero value otherwise.

### GetDnsConfigurationValidationEnabledOk

`func (o *ConfigurationPostRequestBody) GetDnsConfigurationValidationEnabledOk() (*bool, bool)`

GetDnsConfigurationValidationEnabledOk returns a tuple with the DnsConfigurationValidationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsConfigurationValidationEnabled

`func (o *ConfigurationPostRequestBody) SetDnsConfigurationValidationEnabled(v bool)`

SetDnsConfigurationValidationEnabled sets DnsConfigurationValidationEnabled field to given value.

### HasDnsConfigurationValidationEnabled

`func (o *ConfigurationPostRequestBody) HasDnsConfigurationValidationEnabled() bool`

HasDnsConfigurationValidationEnabled returns a boolean if a field has been set.

### GetDnsZoneValidationEnabled

`func (o *ConfigurationPostRequestBody) GetDnsZoneValidationEnabled() bool`

GetDnsZoneValidationEnabled returns the DnsZoneValidationEnabled field if non-nil, zero value otherwise.

### GetDnsZoneValidationEnabledOk

`func (o *ConfigurationPostRequestBody) GetDnsZoneValidationEnabledOk() (*bool, bool)`

GetDnsZoneValidationEnabledOk returns a tuple with the DnsZoneValidationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsZoneValidationEnabled

`func (o *ConfigurationPostRequestBody) SetDnsZoneValidationEnabled(v bool)`

SetDnsZoneValidationEnabled sets DnsZoneValidationEnabled field to given value.

### HasDnsZoneValidationEnabled

`func (o *ConfigurationPostRequestBody) HasDnsZoneValidationEnabled() bool`

HasDnsZoneValidationEnabled returns a boolean if a field has been set.

### GetCheckIntegrityValidation

`func (o *ConfigurationPostRequestBody) GetCheckIntegrityValidation() string`

GetCheckIntegrityValidation returns the CheckIntegrityValidation field if non-nil, zero value otherwise.

### GetCheckIntegrityValidationOk

`func (o *ConfigurationPostRequestBody) GetCheckIntegrityValidationOk() (*string, bool)`

GetCheckIntegrityValidationOk returns a tuple with the CheckIntegrityValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckIntegrityValidation

`func (o *ConfigurationPostRequestBody) SetCheckIntegrityValidation(v string)`

SetCheckIntegrityValidation sets CheckIntegrityValidation field to given value.

### HasCheckIntegrityValidation

`func (o *ConfigurationPostRequestBody) HasCheckIntegrityValidation() bool`

HasCheckIntegrityValidation returns a boolean if a field has been set.

### GetCheckMxCnameValidation

`func (o *ConfigurationPostRequestBody) GetCheckMxCnameValidation() string`

GetCheckMxCnameValidation returns the CheckMxCnameValidation field if non-nil, zero value otherwise.

### GetCheckMxCnameValidationOk

`func (o *ConfigurationPostRequestBody) GetCheckMxCnameValidationOk() (*string, bool)`

GetCheckMxCnameValidationOk returns a tuple with the CheckMxCnameValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckMxCnameValidation

`func (o *ConfigurationPostRequestBody) SetCheckMxCnameValidation(v string)`

SetCheckMxCnameValidation sets CheckMxCnameValidation field to given value.

### HasCheckMxCnameValidation

`func (o *ConfigurationPostRequestBody) HasCheckMxCnameValidation() bool`

HasCheckMxCnameValidation returns a boolean if a field has been set.

### GetCheckMxValidation

`func (o *ConfigurationPostRequestBody) GetCheckMxValidation() string`

GetCheckMxValidation returns the CheckMxValidation field if non-nil, zero value otherwise.

### GetCheckMxValidationOk

`func (o *ConfigurationPostRequestBody) GetCheckMxValidationOk() (*string, bool)`

GetCheckMxValidationOk returns a tuple with the CheckMxValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckMxValidation

`func (o *ConfigurationPostRequestBody) SetCheckMxValidation(v string)`

SetCheckMxValidation sets CheckMxValidation field to given value.

### HasCheckMxValidation

`func (o *ConfigurationPostRequestBody) HasCheckMxValidation() bool`

HasCheckMxValidation returns a boolean if a field has been set.

### GetCheckNamesValidation

`func (o *ConfigurationPostRequestBody) GetCheckNamesValidation() string`

GetCheckNamesValidation returns the CheckNamesValidation field if non-nil, zero value otherwise.

### GetCheckNamesValidationOk

`func (o *ConfigurationPostRequestBody) GetCheckNamesValidationOk() (*string, bool)`

GetCheckNamesValidationOk returns a tuple with the CheckNamesValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckNamesValidation

`func (o *ConfigurationPostRequestBody) SetCheckNamesValidation(v string)`

SetCheckNamesValidation sets CheckNamesValidation field to given value.

### HasCheckNamesValidation

`func (o *ConfigurationPostRequestBody) HasCheckNamesValidation() bool`

HasCheckNamesValidation returns a boolean if a field has been set.

### GetCheckNsValidation

`func (o *ConfigurationPostRequestBody) GetCheckNsValidation() string`

GetCheckNsValidation returns the CheckNsValidation field if non-nil, zero value otherwise.

### GetCheckNsValidationOk

`func (o *ConfigurationPostRequestBody) GetCheckNsValidationOk() (*string, bool)`

GetCheckNsValidationOk returns a tuple with the CheckNsValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckNsValidation

`func (o *ConfigurationPostRequestBody) SetCheckNsValidation(v string)`

SetCheckNsValidation sets CheckNsValidation field to given value.

### HasCheckNsValidation

`func (o *ConfigurationPostRequestBody) HasCheckNsValidation() bool`

HasCheckNsValidation returns a boolean if a field has been set.

### GetCheckSrvCnameValidation

`func (o *ConfigurationPostRequestBody) GetCheckSrvCnameValidation() string`

GetCheckSrvCnameValidation returns the CheckSrvCnameValidation field if non-nil, zero value otherwise.

### GetCheckSrvCnameValidationOk

`func (o *ConfigurationPostRequestBody) GetCheckSrvCnameValidationOk() (*string, bool)`

GetCheckSrvCnameValidationOk returns a tuple with the CheckSrvCnameValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckSrvCnameValidation

`func (o *ConfigurationPostRequestBody) SetCheckSrvCnameValidation(v string)`

SetCheckSrvCnameValidation sets CheckSrvCnameValidation field to given value.

### HasCheckSrvCnameValidation

`func (o *ConfigurationPostRequestBody) HasCheckSrvCnameValidation() bool`

HasCheckSrvCnameValidation returns a boolean if a field has been set.

### GetCheckWildcardValidation

`func (o *ConfigurationPostRequestBody) GetCheckWildcardValidation() string`

GetCheckWildcardValidation returns the CheckWildcardValidation field if non-nil, zero value otherwise.

### GetCheckWildcardValidationOk

`func (o *ConfigurationPostRequestBody) GetCheckWildcardValidationOk() (*string, bool)`

GetCheckWildcardValidationOk returns a tuple with the CheckWildcardValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckWildcardValidation

`func (o *ConfigurationPostRequestBody) SetCheckWildcardValidation(v string)`

SetCheckWildcardValidation sets CheckWildcardValidation field to given value.

### HasCheckWildcardValidation

`func (o *ConfigurationPostRequestBody) HasCheckWildcardValidation() bool`

HasCheckWildcardValidation returns a boolean if a field has been set.

### GetDnsOptionInheritanceEnabled

`func (o *ConfigurationPostRequestBody) GetDnsOptionInheritanceEnabled() bool`

GetDnsOptionInheritanceEnabled returns the DnsOptionInheritanceEnabled field if non-nil, zero value otherwise.

### GetDnsOptionInheritanceEnabledOk

`func (o *ConfigurationPostRequestBody) GetDnsOptionInheritanceEnabledOk() (*bool, bool)`

GetDnsOptionInheritanceEnabledOk returns a tuple with the DnsOptionInheritanceEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsOptionInheritanceEnabled

`func (o *ConfigurationPostRequestBody) SetDnsOptionInheritanceEnabled(v bool)`

SetDnsOptionInheritanceEnabled sets DnsOptionInheritanceEnabled field to given value.

### HasDnsOptionInheritanceEnabled

`func (o *ConfigurationPostRequestBody) HasDnsOptionInheritanceEnabled() bool`

HasDnsOptionInheritanceEnabled returns a boolean if a field has been set.

### GetDnsFeedEnabled

`func (o *ConfigurationPostRequestBody) GetDnsFeedEnabled() bool`

GetDnsFeedEnabled returns the DnsFeedEnabled field if non-nil, zero value otherwise.

### GetDnsFeedEnabledOk

`func (o *ConfigurationPostRequestBody) GetDnsFeedEnabledOk() (*bool, bool)`

GetDnsFeedEnabledOk returns a tuple with the DnsFeedEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsFeedEnabled

`func (o *ConfigurationPostRequestBody) SetDnsFeedEnabled(v bool)`

SetDnsFeedEnabled sets DnsFeedEnabled field to given value.

### HasDnsFeedEnabled

`func (o *ConfigurationPostRequestBody) HasDnsFeedEnabled() bool`

HasDnsFeedEnabled returns a boolean if a field has been set.

### GetDnsFeedLicense

`func (o *ConfigurationPostRequestBody) GetDnsFeedLicense() string`

GetDnsFeedLicense returns the DnsFeedLicense field if non-nil, zero value otherwise.

### GetDnsFeedLicenseOk

`func (o *ConfigurationPostRequestBody) GetDnsFeedLicenseOk() (*string, bool)`

GetDnsFeedLicenseOk returns a tuple with the DnsFeedLicense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsFeedLicense

`func (o *ConfigurationPostRequestBody) SetDnsFeedLicense(v string)`

SetDnsFeedLicense sets DnsFeedLicense field to given value.

### HasDnsFeedLicense

`func (o *ConfigurationPostRequestBody) HasDnsFeedLicense() bool`

HasDnsFeedLicense returns a boolean if a field has been set.

### GetBlockUsageCalculation

`func (o *ConfigurationPostRequestBody) GetBlockUsageCalculation() string`

GetBlockUsageCalculation returns the BlockUsageCalculation field if non-nil, zero value otherwise.

### GetBlockUsageCalculationOk

`func (o *ConfigurationPostRequestBody) GetBlockUsageCalculationOk() (*string, bool)`

GetBlockUsageCalculationOk returns a tuple with the BlockUsageCalculation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockUsageCalculation

`func (o *ConfigurationPostRequestBody) SetBlockUsageCalculation(v string)`

SetBlockUsageCalculation sets BlockUsageCalculation field to given value.

### HasBlockUsageCalculation

`func (o *ConfigurationPostRequestBody) HasBlockUsageCalculation() bool`

HasBlockUsageCalculation returns a boolean if a field has been set.

### GetSharedNetworkTagGroup

`func (o *ConfigurationPostRequestBody) GetSharedNetworkTagGroup() InlinedTagGroup`

GetSharedNetworkTagGroup returns the SharedNetworkTagGroup field if non-nil, zero value otherwise.

### GetSharedNetworkTagGroupOk

`func (o *ConfigurationPostRequestBody) GetSharedNetworkTagGroupOk() (*InlinedTagGroup, bool)`

GetSharedNetworkTagGroupOk returns a tuple with the SharedNetworkTagGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedNetworkTagGroup

`func (o *ConfigurationPostRequestBody) SetSharedNetworkTagGroup(v InlinedTagGroup)`

SetSharedNetworkTagGroup sets SharedNetworkTagGroup field to given value.

### HasSharedNetworkTagGroup

`func (o *ConfigurationPostRequestBody) HasSharedNetworkTagGroup() bool`

HasSharedNetworkTagGroup returns a boolean if a field has been set.

### GetKeyAutoRegenerationEnabled

`func (o *ConfigurationPostRequestBody) GetKeyAutoRegenerationEnabled() bool`

GetKeyAutoRegenerationEnabled returns the KeyAutoRegenerationEnabled field if non-nil, zero value otherwise.

### GetKeyAutoRegenerationEnabledOk

`func (o *ConfigurationPostRequestBody) GetKeyAutoRegenerationEnabledOk() (*bool, bool)`

GetKeyAutoRegenerationEnabledOk returns a tuple with the KeyAutoRegenerationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyAutoRegenerationEnabled

`func (o *ConfigurationPostRequestBody) SetKeyAutoRegenerationEnabled(v bool)`

SetKeyAutoRegenerationEnabled sets KeyAutoRegenerationEnabled field to given value.

### HasKeyAutoRegenerationEnabled

`func (o *ConfigurationPostRequestBody) HasKeyAutoRegenerationEnabled() bool`

HasKeyAutoRegenerationEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


