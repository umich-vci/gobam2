# Configuration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | Pointer to **string** | The name of the resource. | [optional] 
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

### NewConfiguration

`func NewConfiguration() *Configuration`

NewConfiguration instantiates a new Configuration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigurationWithDefaults

`func NewConfigurationWithDefaults() *Configuration`

NewConfigurationWithDefaults instantiates a new Configuration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Configuration) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Configuration) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Configuration) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Configuration) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *Configuration) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Configuration) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Configuration) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Configuration) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *Configuration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Configuration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Configuration) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Configuration) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *Configuration) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *Configuration) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *Configuration) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *Configuration) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetConfigurationGroup

`func (o *Configuration) GetConfigurationGroup() string`

GetConfigurationGroup returns the ConfigurationGroup field if non-nil, zero value otherwise.

### GetConfigurationGroupOk

`func (o *Configuration) GetConfigurationGroupOk() (*string, bool)`

GetConfigurationGroupOk returns a tuple with the ConfigurationGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationGroup

`func (o *Configuration) SetConfigurationGroup(v string)`

SetConfigurationGroup sets ConfigurationGroup field to given value.

### HasConfigurationGroup

`func (o *Configuration) HasConfigurationGroup() bool`

HasConfigurationGroup returns a boolean if a field has been set.

### GetDescription

`func (o *Configuration) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Configuration) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Configuration) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Configuration) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSnmp

`func (o *Configuration) GetSnmp() SnmpBean`

GetSnmp returns the Snmp field if non-nil, zero value otherwise.

### GetSnmpOk

`func (o *Configuration) GetSnmpOk() (*SnmpBean, bool)`

GetSnmpOk returns a tuple with the Snmp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmp

`func (o *Configuration) SetSnmp(v SnmpBean)`

SetSnmp sets Snmp field to given value.

### HasSnmp

`func (o *Configuration) HasSnmp() bool`

HasSnmp returns a boolean if a field has been set.

### GetServerMonitoringEnabled

`func (o *Configuration) GetServerMonitoringEnabled() bool`

GetServerMonitoringEnabled returns the ServerMonitoringEnabled field if non-nil, zero value otherwise.

### GetServerMonitoringEnabledOk

`func (o *Configuration) GetServerMonitoringEnabledOk() (*bool, bool)`

GetServerMonitoringEnabledOk returns a tuple with the ServerMonitoringEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerMonitoringEnabled

`func (o *Configuration) SetServerMonitoringEnabled(v bool)`

SetServerMonitoringEnabled sets ServerMonitoringEnabled field to given value.

### HasServerMonitoringEnabled

`func (o *Configuration) HasServerMonitoringEnabled() bool`

HasServerMonitoringEnabled returns a boolean if a field has been set.

### GetDhcpConfigurationValidationEnabled

`func (o *Configuration) GetDhcpConfigurationValidationEnabled() bool`

GetDhcpConfigurationValidationEnabled returns the DhcpConfigurationValidationEnabled field if non-nil, zero value otherwise.

### GetDhcpConfigurationValidationEnabledOk

`func (o *Configuration) GetDhcpConfigurationValidationEnabledOk() (*bool, bool)`

GetDhcpConfigurationValidationEnabledOk returns a tuple with the DhcpConfigurationValidationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpConfigurationValidationEnabled

`func (o *Configuration) SetDhcpConfigurationValidationEnabled(v bool)`

SetDhcpConfigurationValidationEnabled sets DhcpConfigurationValidationEnabled field to given value.

### HasDhcpConfigurationValidationEnabled

`func (o *Configuration) HasDhcpConfigurationValidationEnabled() bool`

HasDhcpConfigurationValidationEnabled returns a boolean if a field has been set.

### GetDnsConfigurationValidationEnabled

`func (o *Configuration) GetDnsConfigurationValidationEnabled() bool`

GetDnsConfigurationValidationEnabled returns the DnsConfigurationValidationEnabled field if non-nil, zero value otherwise.

### GetDnsConfigurationValidationEnabledOk

`func (o *Configuration) GetDnsConfigurationValidationEnabledOk() (*bool, bool)`

GetDnsConfigurationValidationEnabledOk returns a tuple with the DnsConfigurationValidationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsConfigurationValidationEnabled

`func (o *Configuration) SetDnsConfigurationValidationEnabled(v bool)`

SetDnsConfigurationValidationEnabled sets DnsConfigurationValidationEnabled field to given value.

### HasDnsConfigurationValidationEnabled

`func (o *Configuration) HasDnsConfigurationValidationEnabled() bool`

HasDnsConfigurationValidationEnabled returns a boolean if a field has been set.

### GetDnsZoneValidationEnabled

`func (o *Configuration) GetDnsZoneValidationEnabled() bool`

GetDnsZoneValidationEnabled returns the DnsZoneValidationEnabled field if non-nil, zero value otherwise.

### GetDnsZoneValidationEnabledOk

`func (o *Configuration) GetDnsZoneValidationEnabledOk() (*bool, bool)`

GetDnsZoneValidationEnabledOk returns a tuple with the DnsZoneValidationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsZoneValidationEnabled

`func (o *Configuration) SetDnsZoneValidationEnabled(v bool)`

SetDnsZoneValidationEnabled sets DnsZoneValidationEnabled field to given value.

### HasDnsZoneValidationEnabled

`func (o *Configuration) HasDnsZoneValidationEnabled() bool`

HasDnsZoneValidationEnabled returns a boolean if a field has been set.

### GetCheckIntegrityValidation

`func (o *Configuration) GetCheckIntegrityValidation() string`

GetCheckIntegrityValidation returns the CheckIntegrityValidation field if non-nil, zero value otherwise.

### GetCheckIntegrityValidationOk

`func (o *Configuration) GetCheckIntegrityValidationOk() (*string, bool)`

GetCheckIntegrityValidationOk returns a tuple with the CheckIntegrityValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckIntegrityValidation

`func (o *Configuration) SetCheckIntegrityValidation(v string)`

SetCheckIntegrityValidation sets CheckIntegrityValidation field to given value.

### HasCheckIntegrityValidation

`func (o *Configuration) HasCheckIntegrityValidation() bool`

HasCheckIntegrityValidation returns a boolean if a field has been set.

### GetCheckMxCnameValidation

`func (o *Configuration) GetCheckMxCnameValidation() string`

GetCheckMxCnameValidation returns the CheckMxCnameValidation field if non-nil, zero value otherwise.

### GetCheckMxCnameValidationOk

`func (o *Configuration) GetCheckMxCnameValidationOk() (*string, bool)`

GetCheckMxCnameValidationOk returns a tuple with the CheckMxCnameValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckMxCnameValidation

`func (o *Configuration) SetCheckMxCnameValidation(v string)`

SetCheckMxCnameValidation sets CheckMxCnameValidation field to given value.

### HasCheckMxCnameValidation

`func (o *Configuration) HasCheckMxCnameValidation() bool`

HasCheckMxCnameValidation returns a boolean if a field has been set.

### GetCheckMxValidation

`func (o *Configuration) GetCheckMxValidation() string`

GetCheckMxValidation returns the CheckMxValidation field if non-nil, zero value otherwise.

### GetCheckMxValidationOk

`func (o *Configuration) GetCheckMxValidationOk() (*string, bool)`

GetCheckMxValidationOk returns a tuple with the CheckMxValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckMxValidation

`func (o *Configuration) SetCheckMxValidation(v string)`

SetCheckMxValidation sets CheckMxValidation field to given value.

### HasCheckMxValidation

`func (o *Configuration) HasCheckMxValidation() bool`

HasCheckMxValidation returns a boolean if a field has been set.

### GetCheckNamesValidation

`func (o *Configuration) GetCheckNamesValidation() string`

GetCheckNamesValidation returns the CheckNamesValidation field if non-nil, zero value otherwise.

### GetCheckNamesValidationOk

`func (o *Configuration) GetCheckNamesValidationOk() (*string, bool)`

GetCheckNamesValidationOk returns a tuple with the CheckNamesValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckNamesValidation

`func (o *Configuration) SetCheckNamesValidation(v string)`

SetCheckNamesValidation sets CheckNamesValidation field to given value.

### HasCheckNamesValidation

`func (o *Configuration) HasCheckNamesValidation() bool`

HasCheckNamesValidation returns a boolean if a field has been set.

### GetCheckNsValidation

`func (o *Configuration) GetCheckNsValidation() string`

GetCheckNsValidation returns the CheckNsValidation field if non-nil, zero value otherwise.

### GetCheckNsValidationOk

`func (o *Configuration) GetCheckNsValidationOk() (*string, bool)`

GetCheckNsValidationOk returns a tuple with the CheckNsValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckNsValidation

`func (o *Configuration) SetCheckNsValidation(v string)`

SetCheckNsValidation sets CheckNsValidation field to given value.

### HasCheckNsValidation

`func (o *Configuration) HasCheckNsValidation() bool`

HasCheckNsValidation returns a boolean if a field has been set.

### GetCheckSrvCnameValidation

`func (o *Configuration) GetCheckSrvCnameValidation() string`

GetCheckSrvCnameValidation returns the CheckSrvCnameValidation field if non-nil, zero value otherwise.

### GetCheckSrvCnameValidationOk

`func (o *Configuration) GetCheckSrvCnameValidationOk() (*string, bool)`

GetCheckSrvCnameValidationOk returns a tuple with the CheckSrvCnameValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckSrvCnameValidation

`func (o *Configuration) SetCheckSrvCnameValidation(v string)`

SetCheckSrvCnameValidation sets CheckSrvCnameValidation field to given value.

### HasCheckSrvCnameValidation

`func (o *Configuration) HasCheckSrvCnameValidation() bool`

HasCheckSrvCnameValidation returns a boolean if a field has been set.

### GetCheckWildcardValidation

`func (o *Configuration) GetCheckWildcardValidation() string`

GetCheckWildcardValidation returns the CheckWildcardValidation field if non-nil, zero value otherwise.

### GetCheckWildcardValidationOk

`func (o *Configuration) GetCheckWildcardValidationOk() (*string, bool)`

GetCheckWildcardValidationOk returns a tuple with the CheckWildcardValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckWildcardValidation

`func (o *Configuration) SetCheckWildcardValidation(v string)`

SetCheckWildcardValidation sets CheckWildcardValidation field to given value.

### HasCheckWildcardValidation

`func (o *Configuration) HasCheckWildcardValidation() bool`

HasCheckWildcardValidation returns a boolean if a field has been set.

### GetDnsOptionInheritanceEnabled

`func (o *Configuration) GetDnsOptionInheritanceEnabled() bool`

GetDnsOptionInheritanceEnabled returns the DnsOptionInheritanceEnabled field if non-nil, zero value otherwise.

### GetDnsOptionInheritanceEnabledOk

`func (o *Configuration) GetDnsOptionInheritanceEnabledOk() (*bool, bool)`

GetDnsOptionInheritanceEnabledOk returns a tuple with the DnsOptionInheritanceEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsOptionInheritanceEnabled

`func (o *Configuration) SetDnsOptionInheritanceEnabled(v bool)`

SetDnsOptionInheritanceEnabled sets DnsOptionInheritanceEnabled field to given value.

### HasDnsOptionInheritanceEnabled

`func (o *Configuration) HasDnsOptionInheritanceEnabled() bool`

HasDnsOptionInheritanceEnabled returns a boolean if a field has been set.

### GetDnsFeedEnabled

`func (o *Configuration) GetDnsFeedEnabled() bool`

GetDnsFeedEnabled returns the DnsFeedEnabled field if non-nil, zero value otherwise.

### GetDnsFeedEnabledOk

`func (o *Configuration) GetDnsFeedEnabledOk() (*bool, bool)`

GetDnsFeedEnabledOk returns a tuple with the DnsFeedEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsFeedEnabled

`func (o *Configuration) SetDnsFeedEnabled(v bool)`

SetDnsFeedEnabled sets DnsFeedEnabled field to given value.

### HasDnsFeedEnabled

`func (o *Configuration) HasDnsFeedEnabled() bool`

HasDnsFeedEnabled returns a boolean if a field has been set.

### GetDnsFeedLicense

`func (o *Configuration) GetDnsFeedLicense() string`

GetDnsFeedLicense returns the DnsFeedLicense field if non-nil, zero value otherwise.

### GetDnsFeedLicenseOk

`func (o *Configuration) GetDnsFeedLicenseOk() (*string, bool)`

GetDnsFeedLicenseOk returns a tuple with the DnsFeedLicense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsFeedLicense

`func (o *Configuration) SetDnsFeedLicense(v string)`

SetDnsFeedLicense sets DnsFeedLicense field to given value.

### HasDnsFeedLicense

`func (o *Configuration) HasDnsFeedLicense() bool`

HasDnsFeedLicense returns a boolean if a field has been set.

### GetBlockUsageCalculation

`func (o *Configuration) GetBlockUsageCalculation() string`

GetBlockUsageCalculation returns the BlockUsageCalculation field if non-nil, zero value otherwise.

### GetBlockUsageCalculationOk

`func (o *Configuration) GetBlockUsageCalculationOk() (*string, bool)`

GetBlockUsageCalculationOk returns a tuple with the BlockUsageCalculation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockUsageCalculation

`func (o *Configuration) SetBlockUsageCalculation(v string)`

SetBlockUsageCalculation sets BlockUsageCalculation field to given value.

### HasBlockUsageCalculation

`func (o *Configuration) HasBlockUsageCalculation() bool`

HasBlockUsageCalculation returns a boolean if a field has been set.

### GetSharedNetworkTagGroup

`func (o *Configuration) GetSharedNetworkTagGroup() InlinedTagGroup`

GetSharedNetworkTagGroup returns the SharedNetworkTagGroup field if non-nil, zero value otherwise.

### GetSharedNetworkTagGroupOk

`func (o *Configuration) GetSharedNetworkTagGroupOk() (*InlinedTagGroup, bool)`

GetSharedNetworkTagGroupOk returns a tuple with the SharedNetworkTagGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedNetworkTagGroup

`func (o *Configuration) SetSharedNetworkTagGroup(v InlinedTagGroup)`

SetSharedNetworkTagGroup sets SharedNetworkTagGroup field to given value.

### HasSharedNetworkTagGroup

`func (o *Configuration) HasSharedNetworkTagGroup() bool`

HasSharedNetworkTagGroup returns a boolean if a field has been set.

### GetKeyAutoRegenerationEnabled

`func (o *Configuration) GetKeyAutoRegenerationEnabled() bool`

GetKeyAutoRegenerationEnabled returns the KeyAutoRegenerationEnabled field if non-nil, zero value otherwise.

### GetKeyAutoRegenerationEnabledOk

`func (o *Configuration) GetKeyAutoRegenerationEnabledOk() (*bool, bool)`

GetKeyAutoRegenerationEnabledOk returns a tuple with the KeyAutoRegenerationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyAutoRegenerationEnabled

`func (o *Configuration) SetKeyAutoRegenerationEnabled(v bool)`

SetKeyAutoRegenerationEnabled sets KeyAutoRegenerationEnabled field to given value.

### HasKeyAutoRegenerationEnabled

`func (o *Configuration) HasKeyAutoRegenerationEnabled() bool`

HasKeyAutoRegenerationEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


