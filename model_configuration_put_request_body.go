/*
BlueCat Address Manager 9.5 RESTful v2 API

The **BlueCat Address Manager 9.5 RESTful v2 API** is a new RESTful API for Address Manager that presents Address Manager objects as resources. Each resource has a unique endpoint, and related resources are grouped in collections. To fetch an individual resource, a `GET` request is sent to the resource's endpoint. To fetch all resources in a collection, a `GET` request is sent to the collection's endpoint.  The RESTful v2 API is [hypermedia driven](https://en.wikipedia.org/wiki/HATEOAS) and uses the [HAL](https://en.wikipedia.org/wiki/Hypertext_Application_Language) format for representing links. When navigating through the API, you can use those links to navigate to related resources or child resources of the requested endpoint. The API supports the following media types for most endpoints:  `application/hal+json`, `application/json`, and `text/csv`.  For authentication, the API supports both `Basic` and `Bearer` HTTP authentication schemes.  To get started, create a session and receive credentials for `Basic` authentication by sending a `POST` request to the [/sessions](#/Admin%20Resources/postSession) endpoint, with a message body containing the user's credentials:  ```{\"username\":\"apiuser\", \"password\":\"apipass\"}```  The response will contain an `apiToken` field that can be entered with the username in the Swagger UI **Authorize** dialog. The response will also contain a `basicAuthenticationCredentials` field containing a base64 encoding of the requester's username and API token delimited by a colon. This string can be injected directly into request `Authorization` headers in the following format:  ```Authorization: Basic YXBpOlQ0OExOT3VIRGhDcnVBNEo1bGFES3JuS3hTZC9QK3pjczZXTzBJMDY=```  For full details on API format and authentication methods, refer to the Address Manager RESTful v2 API Guide on the BlueCat Documentation Portal.

API version: 9.5.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gobam2

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ConfigurationPutRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigurationPutRequestBody{}

// ConfigurationPutRequestBody struct for ConfigurationPutRequestBody
type ConfigurationPutRequestBody struct {
	// The resource identifier.
	Id *int64 `json:"id,omitempty"`
	// The resource type.
	Type *string `json:"type,omitempty"`
	// The name of the resource.
	Name string `json:"name" validate:"regexp=^.*\\\\S+.*$"`
	// User-defined fields set for the resource.
	UserDefinedFields *map[string]string `json:"userDefinedFields,omitempty"`
	// The configuration group that the configuration is part of.
	ConfigurationGroup *string `json:"configurationGroup,omitempty"`
	// The description of the configuration.
	Description *string `json:"description,omitempty"`
	Snmp *SnmpBean `json:"snmp,omitempty"`
	// Indicates whether monitoring service is enabled for the configuration.
	ServerMonitoringEnabled bool `json:"serverMonitoringEnabled"`
	// Indicates whether the syntax of the dhcpd.conf file is validated before data is deployed from Address Manager.
	DhcpConfigurationValidationEnabled bool `json:"dhcpConfigurationValidationEnabled"`
	// Indicates whether the syntax of the named.conf file is validated before data is deployed from Address Manager.
	DnsConfigurationValidationEnabled bool `json:"dnsConfigurationValidationEnabled"`
	// Indicates whether the syntax of each DNS zone file is validated before data is deployed from Address Manager.
	DnsZoneValidationEnabled bool `json:"dnsZoneValidationEnabled"`
	// Indicates the type of DNS zone validation check performed on the configuration.
	CheckIntegrityValidation string `json:"checkIntegrityValidation"`
	// This option checks that MX records point to a CNAME record rather than an A or AAAA record. This is equivalent to setting the -M switch for the named-checkzone tool.
	CheckMxCnameValidation string `json:"checkMxCnameValidation"`
	// This option checks that MX records point to an IP address rather than an A or AAAA record. This is equivalent to setting the -m switch for the named-checkzone tool.
	CheckMxValidation string `json:"checkMxValidation"`
	// This option checks that A, AAAA, and MX record names are legal hostnames. It also checks that domain names in the RDATA of NS, SOA, and MX records are legal. This is equivalent to setting the -k switch for the named-checkzone tool.
	CheckNamesValidation string `json:"checkNamesValidation"`
	// This option checks that NS records point to an IP address rather than an A or AAAA record. This is equivalent to setting the -n switch for the named-checkzone tool.
	CheckNsValidation string `json:"checkNsValidation"`
	// This option checks that SRV records point to a CNAME record rather than A or AAAA record. This is equivalent to setting the -S switch for the named-checkzone tool.
	CheckSrvCnameValidation string `json:"checkSrvCnameValidation"`
	// This option checks for wildcards in zone names that don't appear as the left-most segment of a zone name: for example, mail.*.example.com. Non-terminal wildcards are permissible, but you may want to be alerted to their presence. This is equivalent to setting the -W switch for the named-checkzone tool.
	CheckWildcardValidation string `json:"checkWildcardValidation"`
	// Controls whether DNS options are inherited by child zones within a configuration. When DNS option inheritance is disabled, DNS options that have been configured on a zone aren't inherited by the child zone. In the reverse space, DNS options that have been configured on a block aren't inherited by the child block or network when DNS option inheritance is disabled.
	DnsOptionInheritanceEnabled bool `json:"dnsOptionInheritanceEnabled"`
	// Defines whether the BlueCat Security Feed is enabled.
	DnsFeedEnabled bool `json:"dnsFeedEnabled"`
	// The BlueCat Security Feed license.
	DnsFeedLicense *string `json:"dnsFeedLicense,omitempty" validate:"regexp=^.*\\\\S+.*$"`
	// Defines the IP space use statistics to keep track of the amount of space available in your block.
	BlockUsageCalculation string `json:"blockUsageCalculation"`
	SharedNetworkTagGroup *InlinedTagGroup `json:"sharedNetworkTagGroup,omitempty"`
	// Indicates whether Address Manager will regenerate ZSKs and KSKs. When a zone is signed with a DNSSEC signing policy, Address Manager enables this option automatically.
	KeyAutoRegenerationEnabled bool `json:"keyAutoRegenerationEnabled"`
}

type _ConfigurationPutRequestBody ConfigurationPutRequestBody

// NewConfigurationPutRequestBody instantiates a new ConfigurationPutRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigurationPutRequestBody(name string, serverMonitoringEnabled bool, dhcpConfigurationValidationEnabled bool, dnsConfigurationValidationEnabled bool, dnsZoneValidationEnabled bool, checkIntegrityValidation string, checkMxCnameValidation string, checkMxValidation string, checkNamesValidation string, checkNsValidation string, checkSrvCnameValidation string, checkWildcardValidation string, dnsOptionInheritanceEnabled bool, dnsFeedEnabled bool, blockUsageCalculation string, keyAutoRegenerationEnabled bool) *ConfigurationPutRequestBody {
	this := ConfigurationPutRequestBody{}
	this.Name = name
	this.ServerMonitoringEnabled = serverMonitoringEnabled
	this.DhcpConfigurationValidationEnabled = dhcpConfigurationValidationEnabled
	this.DnsConfigurationValidationEnabled = dnsConfigurationValidationEnabled
	this.DnsZoneValidationEnabled = dnsZoneValidationEnabled
	this.CheckIntegrityValidation = checkIntegrityValidation
	this.CheckMxCnameValidation = checkMxCnameValidation
	this.CheckMxValidation = checkMxValidation
	this.CheckNamesValidation = checkNamesValidation
	this.CheckNsValidation = checkNsValidation
	this.CheckSrvCnameValidation = checkSrvCnameValidation
	this.CheckWildcardValidation = checkWildcardValidation
	this.DnsOptionInheritanceEnabled = dnsOptionInheritanceEnabled
	this.DnsFeedEnabled = dnsFeedEnabled
	this.BlockUsageCalculation = blockUsageCalculation
	this.KeyAutoRegenerationEnabled = keyAutoRegenerationEnabled
	return &this
}

// NewConfigurationPutRequestBodyWithDefaults instantiates a new ConfigurationPutRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigurationPutRequestBodyWithDefaults() *ConfigurationPutRequestBody {
	this := ConfigurationPutRequestBody{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ConfigurationPutRequestBody) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationPutRequestBody) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ConfigurationPutRequestBody) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ConfigurationPutRequestBody) SetId(v int64) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ConfigurationPutRequestBody) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationPutRequestBody) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ConfigurationPutRequestBody) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ConfigurationPutRequestBody) SetType(v string) {
	o.Type = &v
}

// GetName returns the Name field value
func (o *ConfigurationPutRequestBody) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ConfigurationPutRequestBody) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ConfigurationPutRequestBody) SetName(v string) {
	o.Name = v
}

// GetUserDefinedFields returns the UserDefinedFields field value if set, zero value otherwise.
func (o *ConfigurationPutRequestBody) GetUserDefinedFields() map[string]string {
	if o == nil || IsNil(o.UserDefinedFields) {
		var ret map[string]string
		return ret
	}
	return *o.UserDefinedFields
}

// GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationPutRequestBody) GetUserDefinedFieldsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.UserDefinedFields) {
		return nil, false
	}
	return o.UserDefinedFields, true
}

// HasUserDefinedFields returns a boolean if a field has been set.
func (o *ConfigurationPutRequestBody) HasUserDefinedFields() bool {
	if o != nil && !IsNil(o.UserDefinedFields) {
		return true
	}

	return false
}

// SetUserDefinedFields gets a reference to the given map[string]string and assigns it to the UserDefinedFields field.
func (o *ConfigurationPutRequestBody) SetUserDefinedFields(v map[string]string) {
	o.UserDefinedFields = &v
}

// GetConfigurationGroup returns the ConfigurationGroup field value if set, zero value otherwise.
func (o *ConfigurationPutRequestBody) GetConfigurationGroup() string {
	if o == nil || IsNil(o.ConfigurationGroup) {
		var ret string
		return ret
	}
	return *o.ConfigurationGroup
}

// GetConfigurationGroupOk returns a tuple with the ConfigurationGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationPutRequestBody) GetConfigurationGroupOk() (*string, bool) {
	if o == nil || IsNil(o.ConfigurationGroup) {
		return nil, false
	}
	return o.ConfigurationGroup, true
}

// HasConfigurationGroup returns a boolean if a field has been set.
func (o *ConfigurationPutRequestBody) HasConfigurationGroup() bool {
	if o != nil && !IsNil(o.ConfigurationGroup) {
		return true
	}

	return false
}

// SetConfigurationGroup gets a reference to the given string and assigns it to the ConfigurationGroup field.
func (o *ConfigurationPutRequestBody) SetConfigurationGroup(v string) {
	o.ConfigurationGroup = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ConfigurationPutRequestBody) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationPutRequestBody) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ConfigurationPutRequestBody) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ConfigurationPutRequestBody) SetDescription(v string) {
	o.Description = &v
}

// GetSnmp returns the Snmp field value if set, zero value otherwise.
func (o *ConfigurationPutRequestBody) GetSnmp() SnmpBean {
	if o == nil || IsNil(o.Snmp) {
		var ret SnmpBean
		return ret
	}
	return *o.Snmp
}

// GetSnmpOk returns a tuple with the Snmp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationPutRequestBody) GetSnmpOk() (*SnmpBean, bool) {
	if o == nil || IsNil(o.Snmp) {
		return nil, false
	}
	return o.Snmp, true
}

// HasSnmp returns a boolean if a field has been set.
func (o *ConfigurationPutRequestBody) HasSnmp() bool {
	if o != nil && !IsNil(o.Snmp) {
		return true
	}

	return false
}

// SetSnmp gets a reference to the given SnmpBean and assigns it to the Snmp field.
func (o *ConfigurationPutRequestBody) SetSnmp(v SnmpBean) {
	o.Snmp = &v
}

// GetServerMonitoringEnabled returns the ServerMonitoringEnabled field value
func (o *ConfigurationPutRequestBody) GetServerMonitoringEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ServerMonitoringEnabled
}

// GetServerMonitoringEnabledOk returns a tuple with the ServerMonitoringEnabled field value
// and a boolean to check if the value has been set.
func (o *ConfigurationPutRequestBody) GetServerMonitoringEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerMonitoringEnabled, true
}

// SetServerMonitoringEnabled sets field value
func (o *ConfigurationPutRequestBody) SetServerMonitoringEnabled(v bool) {
	o.ServerMonitoringEnabled = v
}

// GetDhcpConfigurationValidationEnabled returns the DhcpConfigurationValidationEnabled field value
func (o *ConfigurationPutRequestBody) GetDhcpConfigurationValidationEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DhcpConfigurationValidationEnabled
}

// GetDhcpConfigurationValidationEnabledOk returns a tuple with the DhcpConfigurationValidationEnabled field value
// and a boolean to check if the value has been set.
func (o *ConfigurationPutRequestBody) GetDhcpConfigurationValidationEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DhcpConfigurationValidationEnabled, true
}

// SetDhcpConfigurationValidationEnabled sets field value
func (o *ConfigurationPutRequestBody) SetDhcpConfigurationValidationEnabled(v bool) {
	o.DhcpConfigurationValidationEnabled = v
}

// GetDnsConfigurationValidationEnabled returns the DnsConfigurationValidationEnabled field value
func (o *ConfigurationPutRequestBody) GetDnsConfigurationValidationEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DnsConfigurationValidationEnabled
}

// GetDnsConfigurationValidationEnabledOk returns a tuple with the DnsConfigurationValidationEnabled field value
// and a boolean to check if the value has been set.
func (o *ConfigurationPutRequestBody) GetDnsConfigurationValidationEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DnsConfigurationValidationEnabled, true
}

// SetDnsConfigurationValidationEnabled sets field value
func (o *ConfigurationPutRequestBody) SetDnsConfigurationValidationEnabled(v bool) {
	o.DnsConfigurationValidationEnabled = v
}

// GetDnsZoneValidationEnabled returns the DnsZoneValidationEnabled field value
func (o *ConfigurationPutRequestBody) GetDnsZoneValidationEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DnsZoneValidationEnabled
}

// GetDnsZoneValidationEnabledOk returns a tuple with the DnsZoneValidationEnabled field value
// and a boolean to check if the value has been set.
func (o *ConfigurationPutRequestBody) GetDnsZoneValidationEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DnsZoneValidationEnabled, true
}

// SetDnsZoneValidationEnabled sets field value
func (o *ConfigurationPutRequestBody) SetDnsZoneValidationEnabled(v bool) {
	o.DnsZoneValidationEnabled = v
}

// GetCheckIntegrityValidation returns the CheckIntegrityValidation field value
func (o *ConfigurationPutRequestBody) GetCheckIntegrityValidation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CheckIntegrityValidation
}

// GetCheckIntegrityValidationOk returns a tuple with the CheckIntegrityValidation field value
// and a boolean to check if the value has been set.
func (o *ConfigurationPutRequestBody) GetCheckIntegrityValidationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CheckIntegrityValidation, true
}

// SetCheckIntegrityValidation sets field value
func (o *ConfigurationPutRequestBody) SetCheckIntegrityValidation(v string) {
	o.CheckIntegrityValidation = v
}

// GetCheckMxCnameValidation returns the CheckMxCnameValidation field value
func (o *ConfigurationPutRequestBody) GetCheckMxCnameValidation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CheckMxCnameValidation
}

// GetCheckMxCnameValidationOk returns a tuple with the CheckMxCnameValidation field value
// and a boolean to check if the value has been set.
func (o *ConfigurationPutRequestBody) GetCheckMxCnameValidationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CheckMxCnameValidation, true
}

// SetCheckMxCnameValidation sets field value
func (o *ConfigurationPutRequestBody) SetCheckMxCnameValidation(v string) {
	o.CheckMxCnameValidation = v
}

// GetCheckMxValidation returns the CheckMxValidation field value
func (o *ConfigurationPutRequestBody) GetCheckMxValidation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CheckMxValidation
}

// GetCheckMxValidationOk returns a tuple with the CheckMxValidation field value
// and a boolean to check if the value has been set.
func (o *ConfigurationPutRequestBody) GetCheckMxValidationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CheckMxValidation, true
}

// SetCheckMxValidation sets field value
func (o *ConfigurationPutRequestBody) SetCheckMxValidation(v string) {
	o.CheckMxValidation = v
}

// GetCheckNamesValidation returns the CheckNamesValidation field value
func (o *ConfigurationPutRequestBody) GetCheckNamesValidation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CheckNamesValidation
}

// GetCheckNamesValidationOk returns a tuple with the CheckNamesValidation field value
// and a boolean to check if the value has been set.
func (o *ConfigurationPutRequestBody) GetCheckNamesValidationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CheckNamesValidation, true
}

// SetCheckNamesValidation sets field value
func (o *ConfigurationPutRequestBody) SetCheckNamesValidation(v string) {
	o.CheckNamesValidation = v
}

// GetCheckNsValidation returns the CheckNsValidation field value
func (o *ConfigurationPutRequestBody) GetCheckNsValidation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CheckNsValidation
}

// GetCheckNsValidationOk returns a tuple with the CheckNsValidation field value
// and a boolean to check if the value has been set.
func (o *ConfigurationPutRequestBody) GetCheckNsValidationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CheckNsValidation, true
}

// SetCheckNsValidation sets field value
func (o *ConfigurationPutRequestBody) SetCheckNsValidation(v string) {
	o.CheckNsValidation = v
}

// GetCheckSrvCnameValidation returns the CheckSrvCnameValidation field value
func (o *ConfigurationPutRequestBody) GetCheckSrvCnameValidation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CheckSrvCnameValidation
}

// GetCheckSrvCnameValidationOk returns a tuple with the CheckSrvCnameValidation field value
// and a boolean to check if the value has been set.
func (o *ConfigurationPutRequestBody) GetCheckSrvCnameValidationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CheckSrvCnameValidation, true
}

// SetCheckSrvCnameValidation sets field value
func (o *ConfigurationPutRequestBody) SetCheckSrvCnameValidation(v string) {
	o.CheckSrvCnameValidation = v
}

// GetCheckWildcardValidation returns the CheckWildcardValidation field value
func (o *ConfigurationPutRequestBody) GetCheckWildcardValidation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CheckWildcardValidation
}

// GetCheckWildcardValidationOk returns a tuple with the CheckWildcardValidation field value
// and a boolean to check if the value has been set.
func (o *ConfigurationPutRequestBody) GetCheckWildcardValidationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CheckWildcardValidation, true
}

// SetCheckWildcardValidation sets field value
func (o *ConfigurationPutRequestBody) SetCheckWildcardValidation(v string) {
	o.CheckWildcardValidation = v
}

// GetDnsOptionInheritanceEnabled returns the DnsOptionInheritanceEnabled field value
func (o *ConfigurationPutRequestBody) GetDnsOptionInheritanceEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DnsOptionInheritanceEnabled
}

// GetDnsOptionInheritanceEnabledOk returns a tuple with the DnsOptionInheritanceEnabled field value
// and a boolean to check if the value has been set.
func (o *ConfigurationPutRequestBody) GetDnsOptionInheritanceEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DnsOptionInheritanceEnabled, true
}

// SetDnsOptionInheritanceEnabled sets field value
func (o *ConfigurationPutRequestBody) SetDnsOptionInheritanceEnabled(v bool) {
	o.DnsOptionInheritanceEnabled = v
}

// GetDnsFeedEnabled returns the DnsFeedEnabled field value
func (o *ConfigurationPutRequestBody) GetDnsFeedEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DnsFeedEnabled
}

// GetDnsFeedEnabledOk returns a tuple with the DnsFeedEnabled field value
// and a boolean to check if the value has been set.
func (o *ConfigurationPutRequestBody) GetDnsFeedEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DnsFeedEnabled, true
}

// SetDnsFeedEnabled sets field value
func (o *ConfigurationPutRequestBody) SetDnsFeedEnabled(v bool) {
	o.DnsFeedEnabled = v
}

// GetDnsFeedLicense returns the DnsFeedLicense field value if set, zero value otherwise.
func (o *ConfigurationPutRequestBody) GetDnsFeedLicense() string {
	if o == nil || IsNil(o.DnsFeedLicense) {
		var ret string
		return ret
	}
	return *o.DnsFeedLicense
}

// GetDnsFeedLicenseOk returns a tuple with the DnsFeedLicense field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationPutRequestBody) GetDnsFeedLicenseOk() (*string, bool) {
	if o == nil || IsNil(o.DnsFeedLicense) {
		return nil, false
	}
	return o.DnsFeedLicense, true
}

// HasDnsFeedLicense returns a boolean if a field has been set.
func (o *ConfigurationPutRequestBody) HasDnsFeedLicense() bool {
	if o != nil && !IsNil(o.DnsFeedLicense) {
		return true
	}

	return false
}

// SetDnsFeedLicense gets a reference to the given string and assigns it to the DnsFeedLicense field.
func (o *ConfigurationPutRequestBody) SetDnsFeedLicense(v string) {
	o.DnsFeedLicense = &v
}

// GetBlockUsageCalculation returns the BlockUsageCalculation field value
func (o *ConfigurationPutRequestBody) GetBlockUsageCalculation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BlockUsageCalculation
}

// GetBlockUsageCalculationOk returns a tuple with the BlockUsageCalculation field value
// and a boolean to check if the value has been set.
func (o *ConfigurationPutRequestBody) GetBlockUsageCalculationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockUsageCalculation, true
}

// SetBlockUsageCalculation sets field value
func (o *ConfigurationPutRequestBody) SetBlockUsageCalculation(v string) {
	o.BlockUsageCalculation = v
}

// GetSharedNetworkTagGroup returns the SharedNetworkTagGroup field value if set, zero value otherwise.
func (o *ConfigurationPutRequestBody) GetSharedNetworkTagGroup() InlinedTagGroup {
	if o == nil || IsNil(o.SharedNetworkTagGroup) {
		var ret InlinedTagGroup
		return ret
	}
	return *o.SharedNetworkTagGroup
}

// GetSharedNetworkTagGroupOk returns a tuple with the SharedNetworkTagGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationPutRequestBody) GetSharedNetworkTagGroupOk() (*InlinedTagGroup, bool) {
	if o == nil || IsNil(o.SharedNetworkTagGroup) {
		return nil, false
	}
	return o.SharedNetworkTagGroup, true
}

// HasSharedNetworkTagGroup returns a boolean if a field has been set.
func (o *ConfigurationPutRequestBody) HasSharedNetworkTagGroup() bool {
	if o != nil && !IsNil(o.SharedNetworkTagGroup) {
		return true
	}

	return false
}

// SetSharedNetworkTagGroup gets a reference to the given InlinedTagGroup and assigns it to the SharedNetworkTagGroup field.
func (o *ConfigurationPutRequestBody) SetSharedNetworkTagGroup(v InlinedTagGroup) {
	o.SharedNetworkTagGroup = &v
}

// GetKeyAutoRegenerationEnabled returns the KeyAutoRegenerationEnabled field value
func (o *ConfigurationPutRequestBody) GetKeyAutoRegenerationEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.KeyAutoRegenerationEnabled
}

// GetKeyAutoRegenerationEnabledOk returns a tuple with the KeyAutoRegenerationEnabled field value
// and a boolean to check if the value has been set.
func (o *ConfigurationPutRequestBody) GetKeyAutoRegenerationEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeyAutoRegenerationEnabled, true
}

// SetKeyAutoRegenerationEnabled sets field value
func (o *ConfigurationPutRequestBody) SetKeyAutoRegenerationEnabled(v bool) {
	o.KeyAutoRegenerationEnabled = v
}

func (o ConfigurationPutRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigurationPutRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.UserDefinedFields) {
		toSerialize["userDefinedFields"] = o.UserDefinedFields
	}
	if !IsNil(o.ConfigurationGroup) {
		toSerialize["configurationGroup"] = o.ConfigurationGroup
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Snmp) {
		toSerialize["snmp"] = o.Snmp
	}
	toSerialize["serverMonitoringEnabled"] = o.ServerMonitoringEnabled
	toSerialize["dhcpConfigurationValidationEnabled"] = o.DhcpConfigurationValidationEnabled
	toSerialize["dnsConfigurationValidationEnabled"] = o.DnsConfigurationValidationEnabled
	toSerialize["dnsZoneValidationEnabled"] = o.DnsZoneValidationEnabled
	toSerialize["checkIntegrityValidation"] = o.CheckIntegrityValidation
	toSerialize["checkMxCnameValidation"] = o.CheckMxCnameValidation
	toSerialize["checkMxValidation"] = o.CheckMxValidation
	toSerialize["checkNamesValidation"] = o.CheckNamesValidation
	toSerialize["checkNsValidation"] = o.CheckNsValidation
	toSerialize["checkSrvCnameValidation"] = o.CheckSrvCnameValidation
	toSerialize["checkWildcardValidation"] = o.CheckWildcardValidation
	toSerialize["dnsOptionInheritanceEnabled"] = o.DnsOptionInheritanceEnabled
	toSerialize["dnsFeedEnabled"] = o.DnsFeedEnabled
	if !IsNil(o.DnsFeedLicense) {
		toSerialize["dnsFeedLicense"] = o.DnsFeedLicense
	}
	toSerialize["blockUsageCalculation"] = o.BlockUsageCalculation
	if !IsNil(o.SharedNetworkTagGroup) {
		toSerialize["sharedNetworkTagGroup"] = o.SharedNetworkTagGroup
	}
	toSerialize["keyAutoRegenerationEnabled"] = o.KeyAutoRegenerationEnabled
	return toSerialize, nil
}

func (o *ConfigurationPutRequestBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"serverMonitoringEnabled",
		"dhcpConfigurationValidationEnabled",
		"dnsConfigurationValidationEnabled",
		"dnsZoneValidationEnabled",
		"checkIntegrityValidation",
		"checkMxCnameValidation",
		"checkMxValidation",
		"checkNamesValidation",
		"checkNsValidation",
		"checkSrvCnameValidation",
		"checkWildcardValidation",
		"dnsOptionInheritanceEnabled",
		"dnsFeedEnabled",
		"blockUsageCalculation",
		"keyAutoRegenerationEnabled",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varConfigurationPutRequestBody := _ConfigurationPutRequestBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varConfigurationPutRequestBody)

	if err != nil {
		return err
	}

	*o = ConfigurationPutRequestBody(varConfigurationPutRequestBody)

	return err
}

type NullableConfigurationPutRequestBody struct {
	value *ConfigurationPutRequestBody
	isSet bool
}

func (v NullableConfigurationPutRequestBody) Get() *ConfigurationPutRequestBody {
	return v.value
}

func (v *NullableConfigurationPutRequestBody) Set(val *ConfigurationPutRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigurationPutRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigurationPutRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigurationPutRequestBody(val *ConfigurationPutRequestBody) *NullableConfigurationPutRequestBody {
	return &NullableConfigurationPutRequestBody{value: val, isSet: true}
}

func (v NullableConfigurationPutRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigurationPutRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


