/*
BlueCat Address Manager 9.5 RESTful v2 API

The **BlueCat Address Manager 9.5 RESTful v2 API** is a new RESTful API for Address Manager that presents Address Manager objects as resources. Each resource has a unique endpoint, and related resources are grouped in collections. To fetch an individual resource, a `GET` request is sent to the resource's endpoint. To fetch all resources in a collection, a `GET` request is sent to the collection's endpoint.  The RESTful v2 API is [hypermedia driven](https://en.wikipedia.org/wiki/HATEOAS) and uses the [HAL](https://en.wikipedia.org/wiki/Hypertext_Application_Language) format for representing links. When navigating through the API, you can use those links to navigate to related resources or child resources of the requested endpoint. The API supports the following media types for most endpoints:  `application/hal+json`, `application/json`, and `text/csv`.  For authentication, the API supports both `Basic` and `Bearer` HTTP authentication schemes.  To get started, create a session and receive credentials for `Basic` authentication by sending a `POST` request to the [/sessions](#/Admin%20Resources/postSession) endpoint, with a message body containing the user's credentials:  ```{\"username\":\"apiuser\", \"password\":\"apipass\"}```  The response will contain an `apiToken` field that can be entered with the username in the Swagger UI **Authorize** dialog. The response will also contain a `basicAuthenticationCredentials` field containing a base64 encoding of the requester's username and API token delimited by a colon. This string can be injected directly into request `Authorization` headers in the following format:  ```Authorization: Basic YXBpOlQ0OExOT3VIRGhDcnVBNEo1bGFES3JuS3hTZC9QK3pjczZXTzBJMDY=```  For full details on API format and authentication methods, refer to the Address Manager RESTful v2 API Guide on the BlueCat Documentation Portal.

API version: 9.5.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gobam2

import (
	"encoding/json"
	"time"
)

// checks if the GetReconciliationPolicies200ResponseDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetReconciliationPolicies200ResponseDataInner{}

// GetReconciliationPolicies200ResponseDataInner struct for GetReconciliationPolicies200ResponseDataInner
type GetReconciliationPolicies200ResponseDataInner struct {
	// The resource identifier.
	Id *int64 `json:"id,omitempty"`
	// The resource type.
	Type *string `json:"type,omitempty"`
	// The name of the resource.
	Name *string `json:"name,omitempty"`
	// User-defined fields set for the resource.
	UserDefinedFields *map[string]string `json:"userDefinedFields,omitempty"`
	Configuration *InlinedConfiguration `json:"configuration,omitempty"`
	// Indicates whether the reconciliation policy will run at the scheduled time.
	Enabled *bool `json:"enabled,omitempty"`
	// The discovery method for the reconciliation policy.
	DiscoveryMethod *string `json:"discoveryMethod,omitempty"`
	// The IP address or addresses of the router or layer 3 switch where the network discovery operation is to start.
	SeedRouterAddresses []string `json:"seedRouterAddresses,omitempty"`
	Snmp *SnmpBean `json:"snmp,omitempty"`
	// Indicates whether the Address Manager discovery engine will skip FQDN and reverse DNS lookups.
	SkipFqdnEnabled *bool `json:"skipFqdnEnabled,omitempty"`
	// The IP address or addresses for the DNS server(s) that the discovery engine will use to perform FQDN and reverse DNS lookups.
	DnsServers []string `json:"dnsServers,omitempty"`
	// The VLAN ID for the black hole VLAN. The black hole VLAN is used as a default VLAN for all unused ports.
	BlackHoleVlan *int32 `json:"blackHoleVlan,omitempty"`
	// The VLAN ID assigned to a trunk as a native/default VLAN to protect controlled traffic from being spoofed.
	TrunkDefaultVlan *int32 `json:"trunkDefaultVlan,omitempty"`
	// The start time for the reconciliation policy discovery.
	StartDateTime *time.Time `json:"startDateTime,omitempty"`
	// The recurring frequency for reconciliation policy discovery.
	Frequency *string `json:"frequency,omitempty"`
	// The IPv4 range(s) in CIDR format for the reconciliation policy to search.
	NetworkBoundaries []string `json:"networkBoundaries,omitempty"`
	// The IPv4 address ranges in CIDR format for which ping sweep sends an ICMP echo request.
	PingSweepNetworkGaps []string `json:"pingSweepNetworkGaps,omitempty"`
	// The IPv4 addresses and ranges for the reconciliation policy to ignore. The list of overrides can contain single addresses, ranges in CIDR format (e.g. 10.0.0.0/16), or ranges in address range format (e.g. 10.0.0.0-10.0.255.255).
	Overrides []string `json:"overrides,omitempty"`
	// The age parameter for automatic reconciliation of reclaimable addresses. Reconciliation will be applied to reclaimable IP addresses older than this value. Reclaimable addresses are addresses that exist in Address Manager but are not found on the physical network.
	ReclaimAcceptanceAge *string `json:"reclaimAcceptanceAge,omitempty"`
	// The age parameter for automatic reconciliation of unknown addresses. Reconciliation will be applied to unknown IP addresses older than this value. Unknown addresses are addresses that exist on the physical network but do not exist in Address Manager.
	UnknownAcceptanceAge *string `json:"unknownAcceptanceAge,omitempty"`
	// The age parameter for automatic reconciliation of mismatched addresses. Reconciliation will be applied to mismatched IP addresses older than this value. Mismatched addresses are addresses that exist in both Address Manager and the physical network, but do not match in either the MAC address, DNS host name info, VLAN info, or connect switch port info.
	MismatchAcceptanceAge *string `json:"mismatchAcceptanceAge,omitempty"`
	View *InlinedView `json:"view,omitempty"`
	Links *ResourceLinks `json:"_links,omitempty"`
}

// NewGetReconciliationPolicies200ResponseDataInner instantiates a new GetReconciliationPolicies200ResponseDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetReconciliationPolicies200ResponseDataInner() *GetReconciliationPolicies200ResponseDataInner {
	this := GetReconciliationPolicies200ResponseDataInner{}
	return &this
}

// NewGetReconciliationPolicies200ResponseDataInnerWithDefaults instantiates a new GetReconciliationPolicies200ResponseDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetReconciliationPolicies200ResponseDataInnerWithDefaults() *GetReconciliationPolicies200ResponseDataInner {
	this := GetReconciliationPolicies200ResponseDataInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetReconciliationPolicies200ResponseDataInner) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetReconciliationPolicies200ResponseDataInner) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetReconciliationPolicies200ResponseDataInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *GetReconciliationPolicies200ResponseDataInner) SetId(v int64) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetReconciliationPolicies200ResponseDataInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetReconciliationPolicies200ResponseDataInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetReconciliationPolicies200ResponseDataInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetReconciliationPolicies200ResponseDataInner) SetType(v string) {
	o.Type = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetReconciliationPolicies200ResponseDataInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetReconciliationPolicies200ResponseDataInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetReconciliationPolicies200ResponseDataInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetReconciliationPolicies200ResponseDataInner) SetName(v string) {
	o.Name = &v
}

// GetUserDefinedFields returns the UserDefinedFields field value if set, zero value otherwise.
func (o *GetReconciliationPolicies200ResponseDataInner) GetUserDefinedFields() map[string]string {
	if o == nil || IsNil(o.UserDefinedFields) {
		var ret map[string]string
		return ret
	}
	return *o.UserDefinedFields
}

// GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetReconciliationPolicies200ResponseDataInner) GetUserDefinedFieldsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.UserDefinedFields) {
		return nil, false
	}
	return o.UserDefinedFields, true
}

// HasUserDefinedFields returns a boolean if a field has been set.
func (o *GetReconciliationPolicies200ResponseDataInner) HasUserDefinedFields() bool {
	if o != nil && !IsNil(o.UserDefinedFields) {
		return true
	}

	return false
}

// SetUserDefinedFields gets a reference to the given map[string]string and assigns it to the UserDefinedFields field.
func (o *GetReconciliationPolicies200ResponseDataInner) SetUserDefinedFields(v map[string]string) {
	o.UserDefinedFields = &v
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *GetReconciliationPolicies200ResponseDataInner) GetConfiguration() InlinedConfiguration {
	if o == nil || IsNil(o.Configuration) {
		var ret InlinedConfiguration
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetReconciliationPolicies200ResponseDataInner) GetConfigurationOk() (*InlinedConfiguration, bool) {
	if o == nil || IsNil(o.Configuration) {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *GetReconciliationPolicies200ResponseDataInner) HasConfiguration() bool {
	if o != nil && !IsNil(o.Configuration) {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given InlinedConfiguration and assigns it to the Configuration field.
func (o *GetReconciliationPolicies200ResponseDataInner) SetConfiguration(v InlinedConfiguration) {
	o.Configuration = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *GetReconciliationPolicies200ResponseDataInner) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetReconciliationPolicies200ResponseDataInner) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *GetReconciliationPolicies200ResponseDataInner) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *GetReconciliationPolicies200ResponseDataInner) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetDiscoveryMethod returns the DiscoveryMethod field value if set, zero value otherwise.
func (o *GetReconciliationPolicies200ResponseDataInner) GetDiscoveryMethod() string {
	if o == nil || IsNil(o.DiscoveryMethod) {
		var ret string
		return ret
	}
	return *o.DiscoveryMethod
}

// GetDiscoveryMethodOk returns a tuple with the DiscoveryMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetReconciliationPolicies200ResponseDataInner) GetDiscoveryMethodOk() (*string, bool) {
	if o == nil || IsNil(o.DiscoveryMethod) {
		return nil, false
	}
	return o.DiscoveryMethod, true
}

// HasDiscoveryMethod returns a boolean if a field has been set.
func (o *GetReconciliationPolicies200ResponseDataInner) HasDiscoveryMethod() bool {
	if o != nil && !IsNil(o.DiscoveryMethod) {
		return true
	}

	return false
}

// SetDiscoveryMethod gets a reference to the given string and assigns it to the DiscoveryMethod field.
func (o *GetReconciliationPolicies200ResponseDataInner) SetDiscoveryMethod(v string) {
	o.DiscoveryMethod = &v
}

// GetSeedRouterAddresses returns the SeedRouterAddresses field value if set, zero value otherwise.
func (o *GetReconciliationPolicies200ResponseDataInner) GetSeedRouterAddresses() []string {
	if o == nil || IsNil(o.SeedRouterAddresses) {
		var ret []string
		return ret
	}
	return o.SeedRouterAddresses
}

// GetSeedRouterAddressesOk returns a tuple with the SeedRouterAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetReconciliationPolicies200ResponseDataInner) GetSeedRouterAddressesOk() ([]string, bool) {
	if o == nil || IsNil(o.SeedRouterAddresses) {
		return nil, false
	}
	return o.SeedRouterAddresses, true
}

// HasSeedRouterAddresses returns a boolean if a field has been set.
func (o *GetReconciliationPolicies200ResponseDataInner) HasSeedRouterAddresses() bool {
	if o != nil && !IsNil(o.SeedRouterAddresses) {
		return true
	}

	return false
}

// SetSeedRouterAddresses gets a reference to the given []string and assigns it to the SeedRouterAddresses field.
func (o *GetReconciliationPolicies200ResponseDataInner) SetSeedRouterAddresses(v []string) {
	o.SeedRouterAddresses = v
}

// GetSnmp returns the Snmp field value if set, zero value otherwise.
func (o *GetReconciliationPolicies200ResponseDataInner) GetSnmp() SnmpBean {
	if o == nil || IsNil(o.Snmp) {
		var ret SnmpBean
		return ret
	}
	return *o.Snmp
}

// GetSnmpOk returns a tuple with the Snmp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetReconciliationPolicies200ResponseDataInner) GetSnmpOk() (*SnmpBean, bool) {
	if o == nil || IsNil(o.Snmp) {
		return nil, false
	}
	return o.Snmp, true
}

// HasSnmp returns a boolean if a field has been set.
func (o *GetReconciliationPolicies200ResponseDataInner) HasSnmp() bool {
	if o != nil && !IsNil(o.Snmp) {
		return true
	}

	return false
}

// SetSnmp gets a reference to the given SnmpBean and assigns it to the Snmp field.
func (o *GetReconciliationPolicies200ResponseDataInner) SetSnmp(v SnmpBean) {
	o.Snmp = &v
}

// GetSkipFqdnEnabled returns the SkipFqdnEnabled field value if set, zero value otherwise.
func (o *GetReconciliationPolicies200ResponseDataInner) GetSkipFqdnEnabled() bool {
	if o == nil || IsNil(o.SkipFqdnEnabled) {
		var ret bool
		return ret
	}
	return *o.SkipFqdnEnabled
}

// GetSkipFqdnEnabledOk returns a tuple with the SkipFqdnEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetReconciliationPolicies200ResponseDataInner) GetSkipFqdnEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.SkipFqdnEnabled) {
		return nil, false
	}
	return o.SkipFqdnEnabled, true
}

// HasSkipFqdnEnabled returns a boolean if a field has been set.
func (o *GetReconciliationPolicies200ResponseDataInner) HasSkipFqdnEnabled() bool {
	if o != nil && !IsNil(o.SkipFqdnEnabled) {
		return true
	}

	return false
}

// SetSkipFqdnEnabled gets a reference to the given bool and assigns it to the SkipFqdnEnabled field.
func (o *GetReconciliationPolicies200ResponseDataInner) SetSkipFqdnEnabled(v bool) {
	o.SkipFqdnEnabled = &v
}

// GetDnsServers returns the DnsServers field value if set, zero value otherwise.
func (o *GetReconciliationPolicies200ResponseDataInner) GetDnsServers() []string {
	if o == nil || IsNil(o.DnsServers) {
		var ret []string
		return ret
	}
	return o.DnsServers
}

// GetDnsServersOk returns a tuple with the DnsServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetReconciliationPolicies200ResponseDataInner) GetDnsServersOk() ([]string, bool) {
	if o == nil || IsNil(o.DnsServers) {
		return nil, false
	}
	return o.DnsServers, true
}

// HasDnsServers returns a boolean if a field has been set.
func (o *GetReconciliationPolicies200ResponseDataInner) HasDnsServers() bool {
	if o != nil && !IsNil(o.DnsServers) {
		return true
	}

	return false
}

// SetDnsServers gets a reference to the given []string and assigns it to the DnsServers field.
func (o *GetReconciliationPolicies200ResponseDataInner) SetDnsServers(v []string) {
	o.DnsServers = v
}

// GetBlackHoleVlan returns the BlackHoleVlan field value if set, zero value otherwise.
func (o *GetReconciliationPolicies200ResponseDataInner) GetBlackHoleVlan() int32 {
	if o == nil || IsNil(o.BlackHoleVlan) {
		var ret int32
		return ret
	}
	return *o.BlackHoleVlan
}

// GetBlackHoleVlanOk returns a tuple with the BlackHoleVlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetReconciliationPolicies200ResponseDataInner) GetBlackHoleVlanOk() (*int32, bool) {
	if o == nil || IsNil(o.BlackHoleVlan) {
		return nil, false
	}
	return o.BlackHoleVlan, true
}

// HasBlackHoleVlan returns a boolean if a field has been set.
func (o *GetReconciliationPolicies200ResponseDataInner) HasBlackHoleVlan() bool {
	if o != nil && !IsNil(o.BlackHoleVlan) {
		return true
	}

	return false
}

// SetBlackHoleVlan gets a reference to the given int32 and assigns it to the BlackHoleVlan field.
func (o *GetReconciliationPolicies200ResponseDataInner) SetBlackHoleVlan(v int32) {
	o.BlackHoleVlan = &v
}

// GetTrunkDefaultVlan returns the TrunkDefaultVlan field value if set, zero value otherwise.
func (o *GetReconciliationPolicies200ResponseDataInner) GetTrunkDefaultVlan() int32 {
	if o == nil || IsNil(o.TrunkDefaultVlan) {
		var ret int32
		return ret
	}
	return *o.TrunkDefaultVlan
}

// GetTrunkDefaultVlanOk returns a tuple with the TrunkDefaultVlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetReconciliationPolicies200ResponseDataInner) GetTrunkDefaultVlanOk() (*int32, bool) {
	if o == nil || IsNil(o.TrunkDefaultVlan) {
		return nil, false
	}
	return o.TrunkDefaultVlan, true
}

// HasTrunkDefaultVlan returns a boolean if a field has been set.
func (o *GetReconciliationPolicies200ResponseDataInner) HasTrunkDefaultVlan() bool {
	if o != nil && !IsNil(o.TrunkDefaultVlan) {
		return true
	}

	return false
}

// SetTrunkDefaultVlan gets a reference to the given int32 and assigns it to the TrunkDefaultVlan field.
func (o *GetReconciliationPolicies200ResponseDataInner) SetTrunkDefaultVlan(v int32) {
	o.TrunkDefaultVlan = &v
}

// GetStartDateTime returns the StartDateTime field value if set, zero value otherwise.
func (o *GetReconciliationPolicies200ResponseDataInner) GetStartDateTime() time.Time {
	if o == nil || IsNil(o.StartDateTime) {
		var ret time.Time
		return ret
	}
	return *o.StartDateTime
}

// GetStartDateTimeOk returns a tuple with the StartDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetReconciliationPolicies200ResponseDataInner) GetStartDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartDateTime) {
		return nil, false
	}
	return o.StartDateTime, true
}

// HasStartDateTime returns a boolean if a field has been set.
func (o *GetReconciliationPolicies200ResponseDataInner) HasStartDateTime() bool {
	if o != nil && !IsNil(o.StartDateTime) {
		return true
	}

	return false
}

// SetStartDateTime gets a reference to the given time.Time and assigns it to the StartDateTime field.
func (o *GetReconciliationPolicies200ResponseDataInner) SetStartDateTime(v time.Time) {
	o.StartDateTime = &v
}

// GetFrequency returns the Frequency field value if set, zero value otherwise.
func (o *GetReconciliationPolicies200ResponseDataInner) GetFrequency() string {
	if o == nil || IsNil(o.Frequency) {
		var ret string
		return ret
	}
	return *o.Frequency
}

// GetFrequencyOk returns a tuple with the Frequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetReconciliationPolicies200ResponseDataInner) GetFrequencyOk() (*string, bool) {
	if o == nil || IsNil(o.Frequency) {
		return nil, false
	}
	return o.Frequency, true
}

// HasFrequency returns a boolean if a field has been set.
func (o *GetReconciliationPolicies200ResponseDataInner) HasFrequency() bool {
	if o != nil && !IsNil(o.Frequency) {
		return true
	}

	return false
}

// SetFrequency gets a reference to the given string and assigns it to the Frequency field.
func (o *GetReconciliationPolicies200ResponseDataInner) SetFrequency(v string) {
	o.Frequency = &v
}

// GetNetworkBoundaries returns the NetworkBoundaries field value if set, zero value otherwise.
func (o *GetReconciliationPolicies200ResponseDataInner) GetNetworkBoundaries() []string {
	if o == nil || IsNil(o.NetworkBoundaries) {
		var ret []string
		return ret
	}
	return o.NetworkBoundaries
}

// GetNetworkBoundariesOk returns a tuple with the NetworkBoundaries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetReconciliationPolicies200ResponseDataInner) GetNetworkBoundariesOk() ([]string, bool) {
	if o == nil || IsNil(o.NetworkBoundaries) {
		return nil, false
	}
	return o.NetworkBoundaries, true
}

// HasNetworkBoundaries returns a boolean if a field has been set.
func (o *GetReconciliationPolicies200ResponseDataInner) HasNetworkBoundaries() bool {
	if o != nil && !IsNil(o.NetworkBoundaries) {
		return true
	}

	return false
}

// SetNetworkBoundaries gets a reference to the given []string and assigns it to the NetworkBoundaries field.
func (o *GetReconciliationPolicies200ResponseDataInner) SetNetworkBoundaries(v []string) {
	o.NetworkBoundaries = v
}

// GetPingSweepNetworkGaps returns the PingSweepNetworkGaps field value if set, zero value otherwise.
func (o *GetReconciliationPolicies200ResponseDataInner) GetPingSweepNetworkGaps() []string {
	if o == nil || IsNil(o.PingSweepNetworkGaps) {
		var ret []string
		return ret
	}
	return o.PingSweepNetworkGaps
}

// GetPingSweepNetworkGapsOk returns a tuple with the PingSweepNetworkGaps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetReconciliationPolicies200ResponseDataInner) GetPingSweepNetworkGapsOk() ([]string, bool) {
	if o == nil || IsNil(o.PingSweepNetworkGaps) {
		return nil, false
	}
	return o.PingSweepNetworkGaps, true
}

// HasPingSweepNetworkGaps returns a boolean if a field has been set.
func (o *GetReconciliationPolicies200ResponseDataInner) HasPingSweepNetworkGaps() bool {
	if o != nil && !IsNil(o.PingSweepNetworkGaps) {
		return true
	}

	return false
}

// SetPingSweepNetworkGaps gets a reference to the given []string and assigns it to the PingSweepNetworkGaps field.
func (o *GetReconciliationPolicies200ResponseDataInner) SetPingSweepNetworkGaps(v []string) {
	o.PingSweepNetworkGaps = v
}

// GetOverrides returns the Overrides field value if set, zero value otherwise.
func (o *GetReconciliationPolicies200ResponseDataInner) GetOverrides() []string {
	if o == nil || IsNil(o.Overrides) {
		var ret []string
		return ret
	}
	return o.Overrides
}

// GetOverridesOk returns a tuple with the Overrides field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetReconciliationPolicies200ResponseDataInner) GetOverridesOk() ([]string, bool) {
	if o == nil || IsNil(o.Overrides) {
		return nil, false
	}
	return o.Overrides, true
}

// HasOverrides returns a boolean if a field has been set.
func (o *GetReconciliationPolicies200ResponseDataInner) HasOverrides() bool {
	if o != nil && !IsNil(o.Overrides) {
		return true
	}

	return false
}

// SetOverrides gets a reference to the given []string and assigns it to the Overrides field.
func (o *GetReconciliationPolicies200ResponseDataInner) SetOverrides(v []string) {
	o.Overrides = v
}

// GetReclaimAcceptanceAge returns the ReclaimAcceptanceAge field value if set, zero value otherwise.
func (o *GetReconciliationPolicies200ResponseDataInner) GetReclaimAcceptanceAge() string {
	if o == nil || IsNil(o.ReclaimAcceptanceAge) {
		var ret string
		return ret
	}
	return *o.ReclaimAcceptanceAge
}

// GetReclaimAcceptanceAgeOk returns a tuple with the ReclaimAcceptanceAge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetReconciliationPolicies200ResponseDataInner) GetReclaimAcceptanceAgeOk() (*string, bool) {
	if o == nil || IsNil(o.ReclaimAcceptanceAge) {
		return nil, false
	}
	return o.ReclaimAcceptanceAge, true
}

// HasReclaimAcceptanceAge returns a boolean if a field has been set.
func (o *GetReconciliationPolicies200ResponseDataInner) HasReclaimAcceptanceAge() bool {
	if o != nil && !IsNil(o.ReclaimAcceptanceAge) {
		return true
	}

	return false
}

// SetReclaimAcceptanceAge gets a reference to the given string and assigns it to the ReclaimAcceptanceAge field.
func (o *GetReconciliationPolicies200ResponseDataInner) SetReclaimAcceptanceAge(v string) {
	o.ReclaimAcceptanceAge = &v
}

// GetUnknownAcceptanceAge returns the UnknownAcceptanceAge field value if set, zero value otherwise.
func (o *GetReconciliationPolicies200ResponseDataInner) GetUnknownAcceptanceAge() string {
	if o == nil || IsNil(o.UnknownAcceptanceAge) {
		var ret string
		return ret
	}
	return *o.UnknownAcceptanceAge
}

// GetUnknownAcceptanceAgeOk returns a tuple with the UnknownAcceptanceAge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetReconciliationPolicies200ResponseDataInner) GetUnknownAcceptanceAgeOk() (*string, bool) {
	if o == nil || IsNil(o.UnknownAcceptanceAge) {
		return nil, false
	}
	return o.UnknownAcceptanceAge, true
}

// HasUnknownAcceptanceAge returns a boolean if a field has been set.
func (o *GetReconciliationPolicies200ResponseDataInner) HasUnknownAcceptanceAge() bool {
	if o != nil && !IsNil(o.UnknownAcceptanceAge) {
		return true
	}

	return false
}

// SetUnknownAcceptanceAge gets a reference to the given string and assigns it to the UnknownAcceptanceAge field.
func (o *GetReconciliationPolicies200ResponseDataInner) SetUnknownAcceptanceAge(v string) {
	o.UnknownAcceptanceAge = &v
}

// GetMismatchAcceptanceAge returns the MismatchAcceptanceAge field value if set, zero value otherwise.
func (o *GetReconciliationPolicies200ResponseDataInner) GetMismatchAcceptanceAge() string {
	if o == nil || IsNil(o.MismatchAcceptanceAge) {
		var ret string
		return ret
	}
	return *o.MismatchAcceptanceAge
}

// GetMismatchAcceptanceAgeOk returns a tuple with the MismatchAcceptanceAge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetReconciliationPolicies200ResponseDataInner) GetMismatchAcceptanceAgeOk() (*string, bool) {
	if o == nil || IsNil(o.MismatchAcceptanceAge) {
		return nil, false
	}
	return o.MismatchAcceptanceAge, true
}

// HasMismatchAcceptanceAge returns a boolean if a field has been set.
func (o *GetReconciliationPolicies200ResponseDataInner) HasMismatchAcceptanceAge() bool {
	if o != nil && !IsNil(o.MismatchAcceptanceAge) {
		return true
	}

	return false
}

// SetMismatchAcceptanceAge gets a reference to the given string and assigns it to the MismatchAcceptanceAge field.
func (o *GetReconciliationPolicies200ResponseDataInner) SetMismatchAcceptanceAge(v string) {
	o.MismatchAcceptanceAge = &v
}

// GetView returns the View field value if set, zero value otherwise.
func (o *GetReconciliationPolicies200ResponseDataInner) GetView() InlinedView {
	if o == nil || IsNil(o.View) {
		var ret InlinedView
		return ret
	}
	return *o.View
}

// GetViewOk returns a tuple with the View field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetReconciliationPolicies200ResponseDataInner) GetViewOk() (*InlinedView, bool) {
	if o == nil || IsNil(o.View) {
		return nil, false
	}
	return o.View, true
}

// HasView returns a boolean if a field has been set.
func (o *GetReconciliationPolicies200ResponseDataInner) HasView() bool {
	if o != nil && !IsNil(o.View) {
		return true
	}

	return false
}

// SetView gets a reference to the given InlinedView and assigns it to the View field.
func (o *GetReconciliationPolicies200ResponseDataInner) SetView(v InlinedView) {
	o.View = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GetReconciliationPolicies200ResponseDataInner) GetLinks() ResourceLinks {
	if o == nil || IsNil(o.Links) {
		var ret ResourceLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetReconciliationPolicies200ResponseDataInner) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GetReconciliationPolicies200ResponseDataInner) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ResourceLinks and assigns it to the Links field.
func (o *GetReconciliationPolicies200ResponseDataInner) SetLinks(v ResourceLinks) {
	o.Links = &v
}

func (o GetReconciliationPolicies200ResponseDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetReconciliationPolicies200ResponseDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.UserDefinedFields) {
		toSerialize["userDefinedFields"] = o.UserDefinedFields
	}
	if !IsNil(o.Configuration) {
		toSerialize["configuration"] = o.Configuration
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.DiscoveryMethod) {
		toSerialize["discoveryMethod"] = o.DiscoveryMethod
	}
	if !IsNil(o.SeedRouterAddresses) {
		toSerialize["seedRouterAddresses"] = o.SeedRouterAddresses
	}
	if !IsNil(o.Snmp) {
		toSerialize["snmp"] = o.Snmp
	}
	if !IsNil(o.SkipFqdnEnabled) {
		toSerialize["skipFqdnEnabled"] = o.SkipFqdnEnabled
	}
	if !IsNil(o.DnsServers) {
		toSerialize["dnsServers"] = o.DnsServers
	}
	if !IsNil(o.BlackHoleVlan) {
		toSerialize["blackHoleVlan"] = o.BlackHoleVlan
	}
	if !IsNil(o.TrunkDefaultVlan) {
		toSerialize["trunkDefaultVlan"] = o.TrunkDefaultVlan
	}
	if !IsNil(o.StartDateTime) {
		toSerialize["startDateTime"] = o.StartDateTime
	}
	if !IsNil(o.Frequency) {
		toSerialize["frequency"] = o.Frequency
	}
	if !IsNil(o.NetworkBoundaries) {
		toSerialize["networkBoundaries"] = o.NetworkBoundaries
	}
	if !IsNil(o.PingSweepNetworkGaps) {
		toSerialize["pingSweepNetworkGaps"] = o.PingSweepNetworkGaps
	}
	if !IsNil(o.Overrides) {
		toSerialize["overrides"] = o.Overrides
	}
	if !IsNil(o.ReclaimAcceptanceAge) {
		toSerialize["reclaimAcceptanceAge"] = o.ReclaimAcceptanceAge
	}
	if !IsNil(o.UnknownAcceptanceAge) {
		toSerialize["unknownAcceptanceAge"] = o.UnknownAcceptanceAge
	}
	if !IsNil(o.MismatchAcceptanceAge) {
		toSerialize["mismatchAcceptanceAge"] = o.MismatchAcceptanceAge
	}
	if !IsNil(o.View) {
		toSerialize["view"] = o.View
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableGetReconciliationPolicies200ResponseDataInner struct {
	value *GetReconciliationPolicies200ResponseDataInner
	isSet bool
}

func (v NullableGetReconciliationPolicies200ResponseDataInner) Get() *GetReconciliationPolicies200ResponseDataInner {
	return v.value
}

func (v *NullableGetReconciliationPolicies200ResponseDataInner) Set(val *GetReconciliationPolicies200ResponseDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetReconciliationPolicies200ResponseDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetReconciliationPolicies200ResponseDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetReconciliationPolicies200ResponseDataInner(val *GetReconciliationPolicies200ResponseDataInner) *NullableGetReconciliationPolicies200ResponseDataInner {
	return &NullableGetReconciliationPolicies200ResponseDataInner{value: val, isSet: true}
}

func (v NullableGetReconciliationPolicies200ResponseDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetReconciliationPolicies200ResponseDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


