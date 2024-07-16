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

// checks if the ConfigurationPostRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigurationPostRequestBody{}

// ConfigurationPostRequestBody struct for ConfigurationPostRequestBody
type ConfigurationPostRequestBody struct {
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
	ServerMonitoringEnabled *bool `json:"serverMonitoringEnabled,omitempty"`
	// Indicates whether the syntax of the dhcpd.conf file is validated before data is deployed from Address Manager.
	DhcpConfigurationValidationEnabled *bool `json:"dhcpConfigurationValidationEnabled,omitempty"`
	// Indicates whether the syntax of the named.conf file is validated before data is deployed from Address Manager.
	DnsConfigurationValidationEnabled *bool `json:"dnsConfigurationValidationEnabled,omitempty"`
	// Indicates whether the syntax of each DNS zone file is validated before data is deployed from Address Manager.
	DnsZoneValidationEnabled *bool `json:"dnsZoneValidationEnabled,omitempty"`
	// Indicates the type of DNS zone validation check performed on the configuration.
	CheckIntegrityValidation *string `json:"checkIntegrityValidation,omitempty"`
	// This option checks that MX records point to a CNAME record rather than an A or AAAA record. This is equivalent to setting the -M switch for the named-checkzone tool.
	CheckMxCnameValidation *string `json:"checkMxCnameValidation,omitempty"`
	// This option checks that MX records point to an IP address rather than an A or AAAA record. This is equivalent to setting the -m switch for the named-checkzone tool.
	CheckMxValidation *string `json:"checkMxValidation,omitempty"`
	// This option checks that A, AAAA, and MX record names are legal hostnames. It also checks that domain names in the RDATA of NS, SOA, and MX records are legal. This is equivalent to setting the -k switch for the named-checkzone tool.
	CheckNamesValidation *string `json:"checkNamesValidation,omitempty"`
	// This option checks that NS records point to an IP address rather than an A or AAAA record. This is equivalent to setting the -n switch for the named-checkzone tool.
	CheckNsValidation *string `json:"checkNsValidation,omitempty"`
	// This option checks that SRV records point to a CNAME record rather than A or AAAA record. This is equivalent to setting the -S switch for the named-checkzone tool.
	CheckSrvCnameValidation *string `json:"checkSrvCnameValidation,omitempty"`
	// This option checks for wildcards in zone names that don't appear as the left-most segment of a zone name: for example, mail.*.example.com. Non-terminal wildcards are permissible, but you may want to be alerted to their presence. This is equivalent to setting the -W switch for the named-checkzone tool.
	CheckWildcardValidation *string `json:"checkWildcardValidation,omitempty"`
	// Controls whether DNS options are inherited by child zones within a configuration. When DNS option inheritance is disabled, DNS options that have been configured on a zone aren't inherited by the child zone. In the reverse space, DNS options that have been configured on a block aren't inherited by the child block or network when DNS option inheritance is disabled.
	DnsOptionInheritanceEnabled *bool `json:"dnsOptionInheritanceEnabled,omitempty"`
	// Defines whether the BlueCat Security Feed is enabled.
	DnsFeedEnabled *bool `json:"dnsFeedEnabled,omitempty"`
	// The BlueCat Security Feed license.
	DnsFeedLicense *string `json:"dnsFeedLicense,omitempty" validate:"regexp=^.*\\\\S+.*$"`
	// Defines the IP space use statistics to keep track of the amount of space available in your block.
	BlockUsageCalculation *string `json:"blockUsageCalculation,omitempty"`
	SharedNetworkTagGroup *InlinedTagGroup `json:"sharedNetworkTagGroup,omitempty"`
	// Indicates whether Address Manager will regenerate ZSKs and KSKs. When a zone is signed with a DNSSEC signing policy, Address Manager enables this option automatically.
	KeyAutoRegenerationEnabled *bool `json:"keyAutoRegenerationEnabled,omitempty"`
}

type _ConfigurationPostRequestBody ConfigurationPostRequestBody

// NewConfigurationPostRequestBody instantiates a new ConfigurationPostRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigurationPostRequestBody(name string) *ConfigurationPostRequestBody {
	this := ConfigurationPostRequestBody{}
	this.Name = name
	return &this
}

// NewConfigurationPostRequestBodyWithDefaults instantiates a new ConfigurationPostRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigurationPostRequestBodyWithDefaults() *ConfigurationPostRequestBody {
	this := ConfigurationPostRequestBody{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ConfigurationPostRequestBody) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationPostRequestBody) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ConfigurationPostRequestBody) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ConfigurationPostRequestBody) SetId(v int64) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ConfigurationPostRequestBody) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationPostRequestBody) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ConfigurationPostRequestBody) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ConfigurationPostRequestBody) SetType(v string) {
	o.Type = &v
}

// GetName returns the Name field value
func (o *ConfigurationPostRequestBody) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ConfigurationPostRequestBody) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ConfigurationPostRequestBody) SetName(v string) {
	o.Name = v
}

// GetUserDefinedFields returns the UserDefinedFields field value if set, zero value otherwise.
func (o *ConfigurationPostRequestBody) GetUserDefinedFields() map[string]string {
	if o == nil || IsNil(o.UserDefinedFields) {
		var ret map[string]string
		return ret
	}
	return *o.UserDefinedFields
}

// GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationPostRequestBody) GetUserDefinedFieldsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.UserDefinedFields) {
		return nil, false
	}
	return o.UserDefinedFields, true
}

// HasUserDefinedFields returns a boolean if a field has been set.
func (o *ConfigurationPostRequestBody) HasUserDefinedFields() bool {
	if o != nil && !IsNil(o.UserDefinedFields) {
		return true
	}

	return false
}

// SetUserDefinedFields gets a reference to the given map[string]string and assigns it to the UserDefinedFields field.
func (o *ConfigurationPostRequestBody) SetUserDefinedFields(v map[string]string) {
	o.UserDefinedFields = &v
}

// GetConfigurationGroup returns the ConfigurationGroup field value if set, zero value otherwise.
func (o *ConfigurationPostRequestBody) GetConfigurationGroup() string {
	if o == nil || IsNil(o.ConfigurationGroup) {
		var ret string
		return ret
	}
	return *o.ConfigurationGroup
}

// GetConfigurationGroupOk returns a tuple with the ConfigurationGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationPostRequestBody) GetConfigurationGroupOk() (*string, bool) {
	if o == nil || IsNil(o.ConfigurationGroup) {
		return nil, false
	}
	return o.ConfigurationGroup, true
}

// HasConfigurationGroup returns a boolean if a field has been set.
func (o *ConfigurationPostRequestBody) HasConfigurationGroup() bool {
	if o != nil && !IsNil(o.ConfigurationGroup) {
		return true
	}

	return false
}

// SetConfigurationGroup gets a reference to the given string and assigns it to the ConfigurationGroup field.
func (o *ConfigurationPostRequestBody) SetConfigurationGroup(v string) {
	o.ConfigurationGroup = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ConfigurationPostRequestBody) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationPostRequestBody) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ConfigurationPostRequestBody) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ConfigurationPostRequestBody) SetDescription(v string) {
	o.Description = &v
}

// GetSnmp returns the Snmp field value if set, zero value otherwise.
func (o *ConfigurationPostRequestBody) GetSnmp() SnmpBean {
	if o == nil || IsNil(o.Snmp) {
		var ret SnmpBean
		return ret
	}
	return *o.Snmp
}

// GetSnmpOk returns a tuple with the Snmp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationPostRequestBody) GetSnmpOk() (*SnmpBean, bool) {
	if o == nil || IsNil(o.Snmp) {
		return nil, false
	}
	return o.Snmp, true
}

// HasSnmp returns a boolean if a field has been set.
func (o *ConfigurationPostRequestBody) HasSnmp() bool {
	if o != nil && !IsNil(o.Snmp) {
		return true
	}

	return false
}

// SetSnmp gets a reference to the given SnmpBean and assigns it to the Snmp field.
func (o *ConfigurationPostRequestBody) SetSnmp(v SnmpBean) {
	o.Snmp = &v
}

// GetServerMonitoringEnabled returns the ServerMonitoringEnabled field value if set, zero value otherwise.
func (o *ConfigurationPostRequestBody) GetServerMonitoringEnabled() bool {
	if o == nil || IsNil(o.ServerMonitoringEnabled) {
		var ret bool
		return ret
	}
	return *o.ServerMonitoringEnabled
}

// GetServerMonitoringEnabledOk returns a tuple with the ServerMonitoringEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationPostRequestBody) GetServerMonitoringEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.ServerMonitoringEnabled) {
		return nil, false
	}
	return o.ServerMonitoringEnabled, true
}

// HasServerMonitoringEnabled returns a boolean if a field has been set.
func (o *ConfigurationPostRequestBody) HasServerMonitoringEnabled() bool {
	if o != nil && !IsNil(o.ServerMonitoringEnabled) {
		return true
	}

	return false
}

// SetServerMonitoringEnabled gets a reference to the given bool and assigns it to the ServerMonitoringEnabled field.
func (o *ConfigurationPostRequestBody) SetServerMonitoringEnabled(v bool) {
	o.ServerMonitoringEnabled = &v
}

// GetDhcpConfigurationValidationEnabled returns the DhcpConfigurationValidationEnabled field value if set, zero value otherwise.
func (o *ConfigurationPostRequestBody) GetDhcpConfigurationValidationEnabled() bool {
	if o == nil || IsNil(o.DhcpConfigurationValidationEnabled) {
		var ret bool
		return ret
	}
	return *o.DhcpConfigurationValidationEnabled
}

// GetDhcpConfigurationValidationEnabledOk returns a tuple with the DhcpConfigurationValidationEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationPostRequestBody) GetDhcpConfigurationValidationEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.DhcpConfigurationValidationEnabled) {
		return nil, false
	}
	return o.DhcpConfigurationValidationEnabled, true
}

// HasDhcpConfigurationValidationEnabled returns a boolean if a field has been set.
func (o *ConfigurationPostRequestBody) HasDhcpConfigurationValidationEnabled() bool {
	if o != nil && !IsNil(o.DhcpConfigurationValidationEnabled) {
		return true
	}

	return false
}

// SetDhcpConfigurationValidationEnabled gets a reference to the given bool and assigns it to the DhcpConfigurationValidationEnabled field.
func (o *ConfigurationPostRequestBody) SetDhcpConfigurationValidationEnabled(v bool) {
	o.DhcpConfigurationValidationEnabled = &v
}

// GetDnsConfigurationValidationEnabled returns the DnsConfigurationValidationEnabled field value if set, zero value otherwise.
func (o *ConfigurationPostRequestBody) GetDnsConfigurationValidationEnabled() bool {
	if o == nil || IsNil(o.DnsConfigurationValidationEnabled) {
		var ret bool
		return ret
	}
	return *o.DnsConfigurationValidationEnabled
}

// GetDnsConfigurationValidationEnabledOk returns a tuple with the DnsConfigurationValidationEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationPostRequestBody) GetDnsConfigurationValidationEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.DnsConfigurationValidationEnabled) {
		return nil, false
	}
	return o.DnsConfigurationValidationEnabled, true
}

// HasDnsConfigurationValidationEnabled returns a boolean if a field has been set.
func (o *ConfigurationPostRequestBody) HasDnsConfigurationValidationEnabled() bool {
	if o != nil && !IsNil(o.DnsConfigurationValidationEnabled) {
		return true
	}

	return false
}

// SetDnsConfigurationValidationEnabled gets a reference to the given bool and assigns it to the DnsConfigurationValidationEnabled field.
func (o *ConfigurationPostRequestBody) SetDnsConfigurationValidationEnabled(v bool) {
	o.DnsConfigurationValidationEnabled = &v
}

// GetDnsZoneValidationEnabled returns the DnsZoneValidationEnabled field value if set, zero value otherwise.
func (o *ConfigurationPostRequestBody) GetDnsZoneValidationEnabled() bool {
	if o == nil || IsNil(o.DnsZoneValidationEnabled) {
		var ret bool
		return ret
	}
	return *o.DnsZoneValidationEnabled
}

// GetDnsZoneValidationEnabledOk returns a tuple with the DnsZoneValidationEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationPostRequestBody) GetDnsZoneValidationEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.DnsZoneValidationEnabled) {
		return nil, false
	}
	return o.DnsZoneValidationEnabled, true
}

// HasDnsZoneValidationEnabled returns a boolean if a field has been set.
func (o *ConfigurationPostRequestBody) HasDnsZoneValidationEnabled() bool {
	if o != nil && !IsNil(o.DnsZoneValidationEnabled) {
		return true
	}

	return false
}

// SetDnsZoneValidationEnabled gets a reference to the given bool and assigns it to the DnsZoneValidationEnabled field.
func (o *ConfigurationPostRequestBody) SetDnsZoneValidationEnabled(v bool) {
	o.DnsZoneValidationEnabled = &v
}

// GetCheckIntegrityValidation returns the CheckIntegrityValidation field value if set, zero value otherwise.
func (o *ConfigurationPostRequestBody) GetCheckIntegrityValidation() string {
	if o == nil || IsNil(o.CheckIntegrityValidation) {
		var ret string
		return ret
	}
	return *o.CheckIntegrityValidation
}

// GetCheckIntegrityValidationOk returns a tuple with the CheckIntegrityValidation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationPostRequestBody) GetCheckIntegrityValidationOk() (*string, bool) {
	if o == nil || IsNil(o.CheckIntegrityValidation) {
		return nil, false
	}
	return o.CheckIntegrityValidation, true
}

// HasCheckIntegrityValidation returns a boolean if a field has been set.
func (o *ConfigurationPostRequestBody) HasCheckIntegrityValidation() bool {
	if o != nil && !IsNil(o.CheckIntegrityValidation) {
		return true
	}

	return false
}

// SetCheckIntegrityValidation gets a reference to the given string and assigns it to the CheckIntegrityValidation field.
func (o *ConfigurationPostRequestBody) SetCheckIntegrityValidation(v string) {
	o.CheckIntegrityValidation = &v
}

// GetCheckMxCnameValidation returns the CheckMxCnameValidation field value if set, zero value otherwise.
func (o *ConfigurationPostRequestBody) GetCheckMxCnameValidation() string {
	if o == nil || IsNil(o.CheckMxCnameValidation) {
		var ret string
		return ret
	}
	return *o.CheckMxCnameValidation
}

// GetCheckMxCnameValidationOk returns a tuple with the CheckMxCnameValidation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationPostRequestBody) GetCheckMxCnameValidationOk() (*string, bool) {
	if o == nil || IsNil(o.CheckMxCnameValidation) {
		return nil, false
	}
	return o.CheckMxCnameValidation, true
}

// HasCheckMxCnameValidation returns a boolean if a field has been set.
func (o *ConfigurationPostRequestBody) HasCheckMxCnameValidation() bool {
	if o != nil && !IsNil(o.CheckMxCnameValidation) {
		return true
	}

	return false
}

// SetCheckMxCnameValidation gets a reference to the given string and assigns it to the CheckMxCnameValidation field.
func (o *ConfigurationPostRequestBody) SetCheckMxCnameValidation(v string) {
	o.CheckMxCnameValidation = &v
}

// GetCheckMxValidation returns the CheckMxValidation field value if set, zero value otherwise.
func (o *ConfigurationPostRequestBody) GetCheckMxValidation() string {
	if o == nil || IsNil(o.CheckMxValidation) {
		var ret string
		return ret
	}
	return *o.CheckMxValidation
}

// GetCheckMxValidationOk returns a tuple with the CheckMxValidation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationPostRequestBody) GetCheckMxValidationOk() (*string, bool) {
	if o == nil || IsNil(o.CheckMxValidation) {
		return nil, false
	}
	return o.CheckMxValidation, true
}

// HasCheckMxValidation returns a boolean if a field has been set.
func (o *ConfigurationPostRequestBody) HasCheckMxValidation() bool {
	if o != nil && !IsNil(o.CheckMxValidation) {
		return true
	}

	return false
}

// SetCheckMxValidation gets a reference to the given string and assigns it to the CheckMxValidation field.
func (o *ConfigurationPostRequestBody) SetCheckMxValidation(v string) {
	o.CheckMxValidation = &v
}

// GetCheckNamesValidation returns the CheckNamesValidation field value if set, zero value otherwise.
func (o *ConfigurationPostRequestBody) GetCheckNamesValidation() string {
	if o == nil || IsNil(o.CheckNamesValidation) {
		var ret string
		return ret
	}
	return *o.CheckNamesValidation
}

// GetCheckNamesValidationOk returns a tuple with the CheckNamesValidation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationPostRequestBody) GetCheckNamesValidationOk() (*string, bool) {
	if o == nil || IsNil(o.CheckNamesValidation) {
		return nil, false
	}
	return o.CheckNamesValidation, true
}

// HasCheckNamesValidation returns a boolean if a field has been set.
func (o *ConfigurationPostRequestBody) HasCheckNamesValidation() bool {
	if o != nil && !IsNil(o.CheckNamesValidation) {
		return true
	}

	return false
}

// SetCheckNamesValidation gets a reference to the given string and assigns it to the CheckNamesValidation field.
func (o *ConfigurationPostRequestBody) SetCheckNamesValidation(v string) {
	o.CheckNamesValidation = &v
}

// GetCheckNsValidation returns the CheckNsValidation field value if set, zero value otherwise.
func (o *ConfigurationPostRequestBody) GetCheckNsValidation() string {
	if o == nil || IsNil(o.CheckNsValidation) {
		var ret string
		return ret
	}
	return *o.CheckNsValidation
}

// GetCheckNsValidationOk returns a tuple with the CheckNsValidation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationPostRequestBody) GetCheckNsValidationOk() (*string, bool) {
	if o == nil || IsNil(o.CheckNsValidation) {
		return nil, false
	}
	return o.CheckNsValidation, true
}

// HasCheckNsValidation returns a boolean if a field has been set.
func (o *ConfigurationPostRequestBody) HasCheckNsValidation() bool {
	if o != nil && !IsNil(o.CheckNsValidation) {
		return true
	}

	return false
}

// SetCheckNsValidation gets a reference to the given string and assigns it to the CheckNsValidation field.
func (o *ConfigurationPostRequestBody) SetCheckNsValidation(v string) {
	o.CheckNsValidation = &v
}

// GetCheckSrvCnameValidation returns the CheckSrvCnameValidation field value if set, zero value otherwise.
func (o *ConfigurationPostRequestBody) GetCheckSrvCnameValidation() string {
	if o == nil || IsNil(o.CheckSrvCnameValidation) {
		var ret string
		return ret
	}
	return *o.CheckSrvCnameValidation
}

// GetCheckSrvCnameValidationOk returns a tuple with the CheckSrvCnameValidation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationPostRequestBody) GetCheckSrvCnameValidationOk() (*string, bool) {
	if o == nil || IsNil(o.CheckSrvCnameValidation) {
		return nil, false
	}
	return o.CheckSrvCnameValidation, true
}

// HasCheckSrvCnameValidation returns a boolean if a field has been set.
func (o *ConfigurationPostRequestBody) HasCheckSrvCnameValidation() bool {
	if o != nil && !IsNil(o.CheckSrvCnameValidation) {
		return true
	}

	return false
}

// SetCheckSrvCnameValidation gets a reference to the given string and assigns it to the CheckSrvCnameValidation field.
func (o *ConfigurationPostRequestBody) SetCheckSrvCnameValidation(v string) {
	o.CheckSrvCnameValidation = &v
}

// GetCheckWildcardValidation returns the CheckWildcardValidation field value if set, zero value otherwise.
func (o *ConfigurationPostRequestBody) GetCheckWildcardValidation() string {
	if o == nil || IsNil(o.CheckWildcardValidation) {
		var ret string
		return ret
	}
	return *o.CheckWildcardValidation
}

// GetCheckWildcardValidationOk returns a tuple with the CheckWildcardValidation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationPostRequestBody) GetCheckWildcardValidationOk() (*string, bool) {
	if o == nil || IsNil(o.CheckWildcardValidation) {
		return nil, false
	}
	return o.CheckWildcardValidation, true
}

// HasCheckWildcardValidation returns a boolean if a field has been set.
func (o *ConfigurationPostRequestBody) HasCheckWildcardValidation() bool {
	if o != nil && !IsNil(o.CheckWildcardValidation) {
		return true
	}

	return false
}

// SetCheckWildcardValidation gets a reference to the given string and assigns it to the CheckWildcardValidation field.
func (o *ConfigurationPostRequestBody) SetCheckWildcardValidation(v string) {
	o.CheckWildcardValidation = &v
}

// GetDnsOptionInheritanceEnabled returns the DnsOptionInheritanceEnabled field value if set, zero value otherwise.
func (o *ConfigurationPostRequestBody) GetDnsOptionInheritanceEnabled() bool {
	if o == nil || IsNil(o.DnsOptionInheritanceEnabled) {
		var ret bool
		return ret
	}
	return *o.DnsOptionInheritanceEnabled
}

// GetDnsOptionInheritanceEnabledOk returns a tuple with the DnsOptionInheritanceEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationPostRequestBody) GetDnsOptionInheritanceEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.DnsOptionInheritanceEnabled) {
		return nil, false
	}
	return o.DnsOptionInheritanceEnabled, true
}

// HasDnsOptionInheritanceEnabled returns a boolean if a field has been set.
func (o *ConfigurationPostRequestBody) HasDnsOptionInheritanceEnabled() bool {
	if o != nil && !IsNil(o.DnsOptionInheritanceEnabled) {
		return true
	}

	return false
}

// SetDnsOptionInheritanceEnabled gets a reference to the given bool and assigns it to the DnsOptionInheritanceEnabled field.
func (o *ConfigurationPostRequestBody) SetDnsOptionInheritanceEnabled(v bool) {
	o.DnsOptionInheritanceEnabled = &v
}

// GetDnsFeedEnabled returns the DnsFeedEnabled field value if set, zero value otherwise.
func (o *ConfigurationPostRequestBody) GetDnsFeedEnabled() bool {
	if o == nil || IsNil(o.DnsFeedEnabled) {
		var ret bool
		return ret
	}
	return *o.DnsFeedEnabled
}

// GetDnsFeedEnabledOk returns a tuple with the DnsFeedEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationPostRequestBody) GetDnsFeedEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.DnsFeedEnabled) {
		return nil, false
	}
	return o.DnsFeedEnabled, true
}

// HasDnsFeedEnabled returns a boolean if a field has been set.
func (o *ConfigurationPostRequestBody) HasDnsFeedEnabled() bool {
	if o != nil && !IsNil(o.DnsFeedEnabled) {
		return true
	}

	return false
}

// SetDnsFeedEnabled gets a reference to the given bool and assigns it to the DnsFeedEnabled field.
func (o *ConfigurationPostRequestBody) SetDnsFeedEnabled(v bool) {
	o.DnsFeedEnabled = &v
}

// GetDnsFeedLicense returns the DnsFeedLicense field value if set, zero value otherwise.
func (o *ConfigurationPostRequestBody) GetDnsFeedLicense() string {
	if o == nil || IsNil(o.DnsFeedLicense) {
		var ret string
		return ret
	}
	return *o.DnsFeedLicense
}

// GetDnsFeedLicenseOk returns a tuple with the DnsFeedLicense field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationPostRequestBody) GetDnsFeedLicenseOk() (*string, bool) {
	if o == nil || IsNil(o.DnsFeedLicense) {
		return nil, false
	}
	return o.DnsFeedLicense, true
}

// HasDnsFeedLicense returns a boolean if a field has been set.
func (o *ConfigurationPostRequestBody) HasDnsFeedLicense() bool {
	if o != nil && !IsNil(o.DnsFeedLicense) {
		return true
	}

	return false
}

// SetDnsFeedLicense gets a reference to the given string and assigns it to the DnsFeedLicense field.
func (o *ConfigurationPostRequestBody) SetDnsFeedLicense(v string) {
	o.DnsFeedLicense = &v
}

// GetBlockUsageCalculation returns the BlockUsageCalculation field value if set, zero value otherwise.
func (o *ConfigurationPostRequestBody) GetBlockUsageCalculation() string {
	if o == nil || IsNil(o.BlockUsageCalculation) {
		var ret string
		return ret
	}
	return *o.BlockUsageCalculation
}

// GetBlockUsageCalculationOk returns a tuple with the BlockUsageCalculation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationPostRequestBody) GetBlockUsageCalculationOk() (*string, bool) {
	if o == nil || IsNil(o.BlockUsageCalculation) {
		return nil, false
	}
	return o.BlockUsageCalculation, true
}

// HasBlockUsageCalculation returns a boolean if a field has been set.
func (o *ConfigurationPostRequestBody) HasBlockUsageCalculation() bool {
	if o != nil && !IsNil(o.BlockUsageCalculation) {
		return true
	}

	return false
}

// SetBlockUsageCalculation gets a reference to the given string and assigns it to the BlockUsageCalculation field.
func (o *ConfigurationPostRequestBody) SetBlockUsageCalculation(v string) {
	o.BlockUsageCalculation = &v
}

// GetSharedNetworkTagGroup returns the SharedNetworkTagGroup field value if set, zero value otherwise.
func (o *ConfigurationPostRequestBody) GetSharedNetworkTagGroup() InlinedTagGroup {
	if o == nil || IsNil(o.SharedNetworkTagGroup) {
		var ret InlinedTagGroup
		return ret
	}
	return *o.SharedNetworkTagGroup
}

// GetSharedNetworkTagGroupOk returns a tuple with the SharedNetworkTagGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationPostRequestBody) GetSharedNetworkTagGroupOk() (*InlinedTagGroup, bool) {
	if o == nil || IsNil(o.SharedNetworkTagGroup) {
		return nil, false
	}
	return o.SharedNetworkTagGroup, true
}

// HasSharedNetworkTagGroup returns a boolean if a field has been set.
func (o *ConfigurationPostRequestBody) HasSharedNetworkTagGroup() bool {
	if o != nil && !IsNil(o.SharedNetworkTagGroup) {
		return true
	}

	return false
}

// SetSharedNetworkTagGroup gets a reference to the given InlinedTagGroup and assigns it to the SharedNetworkTagGroup field.
func (o *ConfigurationPostRequestBody) SetSharedNetworkTagGroup(v InlinedTagGroup) {
	o.SharedNetworkTagGroup = &v
}

// GetKeyAutoRegenerationEnabled returns the KeyAutoRegenerationEnabled field value if set, zero value otherwise.
func (o *ConfigurationPostRequestBody) GetKeyAutoRegenerationEnabled() bool {
	if o == nil || IsNil(o.KeyAutoRegenerationEnabled) {
		var ret bool
		return ret
	}
	return *o.KeyAutoRegenerationEnabled
}

// GetKeyAutoRegenerationEnabledOk returns a tuple with the KeyAutoRegenerationEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationPostRequestBody) GetKeyAutoRegenerationEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.KeyAutoRegenerationEnabled) {
		return nil, false
	}
	return o.KeyAutoRegenerationEnabled, true
}

// HasKeyAutoRegenerationEnabled returns a boolean if a field has been set.
func (o *ConfigurationPostRequestBody) HasKeyAutoRegenerationEnabled() bool {
	if o != nil && !IsNil(o.KeyAutoRegenerationEnabled) {
		return true
	}

	return false
}

// SetKeyAutoRegenerationEnabled gets a reference to the given bool and assigns it to the KeyAutoRegenerationEnabled field.
func (o *ConfigurationPostRequestBody) SetKeyAutoRegenerationEnabled(v bool) {
	o.KeyAutoRegenerationEnabled = &v
}

func (o ConfigurationPostRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigurationPostRequestBody) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.ServerMonitoringEnabled) {
		toSerialize["serverMonitoringEnabled"] = o.ServerMonitoringEnabled
	}
	if !IsNil(o.DhcpConfigurationValidationEnabled) {
		toSerialize["dhcpConfigurationValidationEnabled"] = o.DhcpConfigurationValidationEnabled
	}
	if !IsNil(o.DnsConfigurationValidationEnabled) {
		toSerialize["dnsConfigurationValidationEnabled"] = o.DnsConfigurationValidationEnabled
	}
	if !IsNil(o.DnsZoneValidationEnabled) {
		toSerialize["dnsZoneValidationEnabled"] = o.DnsZoneValidationEnabled
	}
	if !IsNil(o.CheckIntegrityValidation) {
		toSerialize["checkIntegrityValidation"] = o.CheckIntegrityValidation
	}
	if !IsNil(o.CheckMxCnameValidation) {
		toSerialize["checkMxCnameValidation"] = o.CheckMxCnameValidation
	}
	if !IsNil(o.CheckMxValidation) {
		toSerialize["checkMxValidation"] = o.CheckMxValidation
	}
	if !IsNil(o.CheckNamesValidation) {
		toSerialize["checkNamesValidation"] = o.CheckNamesValidation
	}
	if !IsNil(o.CheckNsValidation) {
		toSerialize["checkNsValidation"] = o.CheckNsValidation
	}
	if !IsNil(o.CheckSrvCnameValidation) {
		toSerialize["checkSrvCnameValidation"] = o.CheckSrvCnameValidation
	}
	if !IsNil(o.CheckWildcardValidation) {
		toSerialize["checkWildcardValidation"] = o.CheckWildcardValidation
	}
	if !IsNil(o.DnsOptionInheritanceEnabled) {
		toSerialize["dnsOptionInheritanceEnabled"] = o.DnsOptionInheritanceEnabled
	}
	if !IsNil(o.DnsFeedEnabled) {
		toSerialize["dnsFeedEnabled"] = o.DnsFeedEnabled
	}
	if !IsNil(o.DnsFeedLicense) {
		toSerialize["dnsFeedLicense"] = o.DnsFeedLicense
	}
	if !IsNil(o.BlockUsageCalculation) {
		toSerialize["blockUsageCalculation"] = o.BlockUsageCalculation
	}
	if !IsNil(o.SharedNetworkTagGroup) {
		toSerialize["sharedNetworkTagGroup"] = o.SharedNetworkTagGroup
	}
	if !IsNil(o.KeyAutoRegenerationEnabled) {
		toSerialize["keyAutoRegenerationEnabled"] = o.KeyAutoRegenerationEnabled
	}
	return toSerialize, nil
}

func (o *ConfigurationPostRequestBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varConfigurationPostRequestBody := _ConfigurationPostRequestBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varConfigurationPostRequestBody)

	if err != nil {
		return err
	}

	*o = ConfigurationPostRequestBody(varConfigurationPostRequestBody)

	return err
}

type NullableConfigurationPostRequestBody struct {
	value *ConfigurationPostRequestBody
	isSet bool
}

func (v NullableConfigurationPostRequestBody) Get() *ConfigurationPostRequestBody {
	return v.value
}

func (v *NullableConfigurationPostRequestBody) Set(val *ConfigurationPostRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigurationPostRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigurationPostRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigurationPostRequestBody(val *ConfigurationPostRequestBody) *NullableConfigurationPostRequestBody {
	return &NullableConfigurationPostRequestBody{value: val, isSet: true}
}

func (v NullableConfigurationPostRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigurationPostRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


