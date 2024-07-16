/*
BlueCat Address Manager 9.5 RESTful v2 API

The **BlueCat Address Manager 9.5 RESTful v2 API** is a new RESTful API for Address Manager that presents Address Manager objects as resources. Each resource has a unique endpoint, and related resources are grouped in collections. To fetch an individual resource, a `GET` request is sent to the resource's endpoint. To fetch all resources in a collection, a `GET` request is sent to the collection's endpoint.  The RESTful v2 API is [hypermedia driven](https://en.wikipedia.org/wiki/HATEOAS) and uses the [HAL](https://en.wikipedia.org/wiki/Hypertext_Application_Language) format for representing links. When navigating through the API, you can use those links to navigate to related resources or child resources of the requested endpoint. The API supports the following media types for most endpoints:  `application/hal+json`, `application/json`, and `text/csv`.  For authentication, the API supports both `Basic` and `Bearer` HTTP authentication schemes.  To get started, create a session and receive credentials for `Basic` authentication by sending a `POST` request to the [/sessions](#/Admin%20Resources/postSession) endpoint, with a message body containing the user's credentials:  ```{\"username\":\"apiuser\", \"password\":\"apipass\"}```  The response will contain an `apiToken` field that can be entered with the username in the Swagger UI **Authorize** dialog. The response will also contain a `basicAuthenticationCredentials` field containing a base64 encoding of the requester's username and API token delimited by a colon. This string can be injected directly into request `Authorization` headers in the following format:  ```Authorization: Basic YXBpOlQ0OExOT3VIRGhDcnVBNEo1bGFES3JuS3hTZC9QK3pjczZXTzBJMDY=```  For full details on API format and authentication methods, refer to the Address Manager RESTful v2 API Guide on the BlueCat Documentation Portal.

API version: 9.5.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gobam2

import (
	"encoding/json"
)

// checks if the GetNetworks200ResponseDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworks200ResponseDataInner{}

// GetNetworks200ResponseDataInner struct for GetNetworks200ResponseDataInner
type GetNetworks200ResponseDataInner struct {
	// The resource identifier.
	Id *int64 `json:"id,omitempty"`
	// The resource type.
	Type *string `json:"type,omitempty"`
	// The name of the resource.
	Name *string `json:"name,omitempty"`
	// User-defined fields set for the resource.
	UserDefinedFields *map[string]string `json:"userDefinedFields,omitempty"`
	Configuration *InlinedConfiguration `json:"configuration,omitempty"`
	// The IP address range for the network. The range can be provided in CIDR (e.g. 10.0.0.0/16) or address range format (e.g. 10.0.0.0-10.0.255.255).
	Range *string `json:"range,omitempty"`
	Location *InlinedLocation `json:"location,omitempty"`
	InheritedFields []string `json:"_inheritedFields,omitempty"`
	// The gateway address for the IPv4 network.
	Gateway *string `json:"gateway,omitempty"`
	Template *InlinedIPv4Template `json:"template,omitempty"`
	// Connected router port information for the IPv4 network.
	RouterPortInfo *string `json:"routerPortInfo,omitempty"`
	DefaultView *InlinedView `json:"defaultView,omitempty"`
	SharedNetworkTag *InlinedTag `json:"sharedNetworkTag,omitempty"`
	// Indicates whether duplicate hostnames are allowed in the network. If enabled, duplicate hostnames will be allowed. If not set explicitly, this value will be inherited from the parent block.
	DuplicateHostnamesAllowed *bool `json:"duplicateHostnamesAllowed,omitempty"`
	// Indicates whether Address Manager will ping IP addresses before assigning them. Unavailable addresses will not be assigned.
	PingBeforeAssignEnabled *bool `json:"pingBeforeAssignEnabled,omitempty"`
	// Indicates whether the IPv4 network will inherit the parent resource's default domains if configured.
	DefaultZonesInherited *bool `json:"defaultZonesInherited,omitempty"`
	DefaultZones []IPv4NetworkAllOfDefaultZones `json:"defaultZones,omitempty"`
	// Indicates whether the IPv4 network will inherit the parent resource's restricted views and zones if configured.
	RestrictedZonesInherited *bool `json:"restrictedZonesInherited,omitempty"`
	RestrictedZones []IPv4NetworkAllOfRestrictedZones `json:"restrictedZones,omitempty"`
	// A DHCP alert is triggered when usage falls below this percentage (when too few addresses are in use).
	LowWaterMark *int32 `json:"lowWaterMark,omitempty"`
	// A DHCP alert is triggered when usage exceeds this percentage (when too many addresses are in use).
	HighWaterMark *int32 `json:"highWaterMark,omitempty"`
	// Indicates whether the reverse zone associated with the IPv4 network will be signed according to the DNSSEC signing policy in reverseZoneSigningPolicy.
	ReverseZoneSigned *bool `json:"reverseZoneSigned,omitempty"`
	ReverseZoneSigningPolicy *InlinedDNSSECSigningPolicy `json:"reverseZoneSigningPolicy,omitempty"`
	// Values indicating the current amount of static, DHCP allocated, DHCP abandoned, DHCP reserved, reserved, gateway, and unassigned IP addresses within the network.
	Usage *map[string]int64 `json:"usage,omitempty"`
	Links *ResourceLinks `json:"_links,omitempty"`
}

// NewGetNetworks200ResponseDataInner instantiates a new GetNetworks200ResponseDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworks200ResponseDataInner() *GetNetworks200ResponseDataInner {
	this := GetNetworks200ResponseDataInner{}
	return &this
}

// NewGetNetworks200ResponseDataInnerWithDefaults instantiates a new GetNetworks200ResponseDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworks200ResponseDataInnerWithDefaults() *GetNetworks200ResponseDataInner {
	this := GetNetworks200ResponseDataInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetNetworks200ResponseDataInner) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworks200ResponseDataInner) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetNetworks200ResponseDataInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *GetNetworks200ResponseDataInner) SetId(v int64) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetNetworks200ResponseDataInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworks200ResponseDataInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetNetworks200ResponseDataInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetNetworks200ResponseDataInner) SetType(v string) {
	o.Type = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetNetworks200ResponseDataInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworks200ResponseDataInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetNetworks200ResponseDataInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetNetworks200ResponseDataInner) SetName(v string) {
	o.Name = &v
}

// GetUserDefinedFields returns the UserDefinedFields field value if set, zero value otherwise.
func (o *GetNetworks200ResponseDataInner) GetUserDefinedFields() map[string]string {
	if o == nil || IsNil(o.UserDefinedFields) {
		var ret map[string]string
		return ret
	}
	return *o.UserDefinedFields
}

// GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworks200ResponseDataInner) GetUserDefinedFieldsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.UserDefinedFields) {
		return nil, false
	}
	return o.UserDefinedFields, true
}

// HasUserDefinedFields returns a boolean if a field has been set.
func (o *GetNetworks200ResponseDataInner) HasUserDefinedFields() bool {
	if o != nil && !IsNil(o.UserDefinedFields) {
		return true
	}

	return false
}

// SetUserDefinedFields gets a reference to the given map[string]string and assigns it to the UserDefinedFields field.
func (o *GetNetworks200ResponseDataInner) SetUserDefinedFields(v map[string]string) {
	o.UserDefinedFields = &v
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *GetNetworks200ResponseDataInner) GetConfiguration() InlinedConfiguration {
	if o == nil || IsNil(o.Configuration) {
		var ret InlinedConfiguration
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworks200ResponseDataInner) GetConfigurationOk() (*InlinedConfiguration, bool) {
	if o == nil || IsNil(o.Configuration) {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *GetNetworks200ResponseDataInner) HasConfiguration() bool {
	if o != nil && !IsNil(o.Configuration) {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given InlinedConfiguration and assigns it to the Configuration field.
func (o *GetNetworks200ResponseDataInner) SetConfiguration(v InlinedConfiguration) {
	o.Configuration = &v
}

// GetRange returns the Range field value if set, zero value otherwise.
func (o *GetNetworks200ResponseDataInner) GetRange() string {
	if o == nil || IsNil(o.Range) {
		var ret string
		return ret
	}
	return *o.Range
}

// GetRangeOk returns a tuple with the Range field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworks200ResponseDataInner) GetRangeOk() (*string, bool) {
	if o == nil || IsNil(o.Range) {
		return nil, false
	}
	return o.Range, true
}

// HasRange returns a boolean if a field has been set.
func (o *GetNetworks200ResponseDataInner) HasRange() bool {
	if o != nil && !IsNil(o.Range) {
		return true
	}

	return false
}

// SetRange gets a reference to the given string and assigns it to the Range field.
func (o *GetNetworks200ResponseDataInner) SetRange(v string) {
	o.Range = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *GetNetworks200ResponseDataInner) GetLocation() InlinedLocation {
	if o == nil || IsNil(o.Location) {
		var ret InlinedLocation
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworks200ResponseDataInner) GetLocationOk() (*InlinedLocation, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *GetNetworks200ResponseDataInner) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given InlinedLocation and assigns it to the Location field.
func (o *GetNetworks200ResponseDataInner) SetLocation(v InlinedLocation) {
	o.Location = &v
}

// GetInheritedFields returns the InheritedFields field value if set, zero value otherwise.
func (o *GetNetworks200ResponseDataInner) GetInheritedFields() []string {
	if o == nil || IsNil(o.InheritedFields) {
		var ret []string
		return ret
	}
	return o.InheritedFields
}

// GetInheritedFieldsOk returns a tuple with the InheritedFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworks200ResponseDataInner) GetInheritedFieldsOk() ([]string, bool) {
	if o == nil || IsNil(o.InheritedFields) {
		return nil, false
	}
	return o.InheritedFields, true
}

// HasInheritedFields returns a boolean if a field has been set.
func (o *GetNetworks200ResponseDataInner) HasInheritedFields() bool {
	if o != nil && !IsNil(o.InheritedFields) {
		return true
	}

	return false
}

// SetInheritedFields gets a reference to the given []string and assigns it to the InheritedFields field.
func (o *GetNetworks200ResponseDataInner) SetInheritedFields(v []string) {
	o.InheritedFields = v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *GetNetworks200ResponseDataInner) GetGateway() string {
	if o == nil || IsNil(o.Gateway) {
		var ret string
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworks200ResponseDataInner) GetGatewayOk() (*string, bool) {
	if o == nil || IsNil(o.Gateway) {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *GetNetworks200ResponseDataInner) HasGateway() bool {
	if o != nil && !IsNil(o.Gateway) {
		return true
	}

	return false
}

// SetGateway gets a reference to the given string and assigns it to the Gateway field.
func (o *GetNetworks200ResponseDataInner) SetGateway(v string) {
	o.Gateway = &v
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *GetNetworks200ResponseDataInner) GetTemplate() InlinedIPv4Template {
	if o == nil || IsNil(o.Template) {
		var ret InlinedIPv4Template
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworks200ResponseDataInner) GetTemplateOk() (*InlinedIPv4Template, bool) {
	if o == nil || IsNil(o.Template) {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *GetNetworks200ResponseDataInner) HasTemplate() bool {
	if o != nil && !IsNil(o.Template) {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given InlinedIPv4Template and assigns it to the Template field.
func (o *GetNetworks200ResponseDataInner) SetTemplate(v InlinedIPv4Template) {
	o.Template = &v
}

// GetRouterPortInfo returns the RouterPortInfo field value if set, zero value otherwise.
func (o *GetNetworks200ResponseDataInner) GetRouterPortInfo() string {
	if o == nil || IsNil(o.RouterPortInfo) {
		var ret string
		return ret
	}
	return *o.RouterPortInfo
}

// GetRouterPortInfoOk returns a tuple with the RouterPortInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworks200ResponseDataInner) GetRouterPortInfoOk() (*string, bool) {
	if o == nil || IsNil(o.RouterPortInfo) {
		return nil, false
	}
	return o.RouterPortInfo, true
}

// HasRouterPortInfo returns a boolean if a field has been set.
func (o *GetNetworks200ResponseDataInner) HasRouterPortInfo() bool {
	if o != nil && !IsNil(o.RouterPortInfo) {
		return true
	}

	return false
}

// SetRouterPortInfo gets a reference to the given string and assigns it to the RouterPortInfo field.
func (o *GetNetworks200ResponseDataInner) SetRouterPortInfo(v string) {
	o.RouterPortInfo = &v
}

// GetDefaultView returns the DefaultView field value if set, zero value otherwise.
func (o *GetNetworks200ResponseDataInner) GetDefaultView() InlinedView {
	if o == nil || IsNil(o.DefaultView) {
		var ret InlinedView
		return ret
	}
	return *o.DefaultView
}

// GetDefaultViewOk returns a tuple with the DefaultView field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworks200ResponseDataInner) GetDefaultViewOk() (*InlinedView, bool) {
	if o == nil || IsNil(o.DefaultView) {
		return nil, false
	}
	return o.DefaultView, true
}

// HasDefaultView returns a boolean if a field has been set.
func (o *GetNetworks200ResponseDataInner) HasDefaultView() bool {
	if o != nil && !IsNil(o.DefaultView) {
		return true
	}

	return false
}

// SetDefaultView gets a reference to the given InlinedView and assigns it to the DefaultView field.
func (o *GetNetworks200ResponseDataInner) SetDefaultView(v InlinedView) {
	o.DefaultView = &v
}

// GetSharedNetworkTag returns the SharedNetworkTag field value if set, zero value otherwise.
func (o *GetNetworks200ResponseDataInner) GetSharedNetworkTag() InlinedTag {
	if o == nil || IsNil(o.SharedNetworkTag) {
		var ret InlinedTag
		return ret
	}
	return *o.SharedNetworkTag
}

// GetSharedNetworkTagOk returns a tuple with the SharedNetworkTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworks200ResponseDataInner) GetSharedNetworkTagOk() (*InlinedTag, bool) {
	if o == nil || IsNil(o.SharedNetworkTag) {
		return nil, false
	}
	return o.SharedNetworkTag, true
}

// HasSharedNetworkTag returns a boolean if a field has been set.
func (o *GetNetworks200ResponseDataInner) HasSharedNetworkTag() bool {
	if o != nil && !IsNil(o.SharedNetworkTag) {
		return true
	}

	return false
}

// SetSharedNetworkTag gets a reference to the given InlinedTag and assigns it to the SharedNetworkTag field.
func (o *GetNetworks200ResponseDataInner) SetSharedNetworkTag(v InlinedTag) {
	o.SharedNetworkTag = &v
}

// GetDuplicateHostnamesAllowed returns the DuplicateHostnamesAllowed field value if set, zero value otherwise.
func (o *GetNetworks200ResponseDataInner) GetDuplicateHostnamesAllowed() bool {
	if o == nil || IsNil(o.DuplicateHostnamesAllowed) {
		var ret bool
		return ret
	}
	return *o.DuplicateHostnamesAllowed
}

// GetDuplicateHostnamesAllowedOk returns a tuple with the DuplicateHostnamesAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworks200ResponseDataInner) GetDuplicateHostnamesAllowedOk() (*bool, bool) {
	if o == nil || IsNil(o.DuplicateHostnamesAllowed) {
		return nil, false
	}
	return o.DuplicateHostnamesAllowed, true
}

// HasDuplicateHostnamesAllowed returns a boolean if a field has been set.
func (o *GetNetworks200ResponseDataInner) HasDuplicateHostnamesAllowed() bool {
	if o != nil && !IsNil(o.DuplicateHostnamesAllowed) {
		return true
	}

	return false
}

// SetDuplicateHostnamesAllowed gets a reference to the given bool and assigns it to the DuplicateHostnamesAllowed field.
func (o *GetNetworks200ResponseDataInner) SetDuplicateHostnamesAllowed(v bool) {
	o.DuplicateHostnamesAllowed = &v
}

// GetPingBeforeAssignEnabled returns the PingBeforeAssignEnabled field value if set, zero value otherwise.
func (o *GetNetworks200ResponseDataInner) GetPingBeforeAssignEnabled() bool {
	if o == nil || IsNil(o.PingBeforeAssignEnabled) {
		var ret bool
		return ret
	}
	return *o.PingBeforeAssignEnabled
}

// GetPingBeforeAssignEnabledOk returns a tuple with the PingBeforeAssignEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworks200ResponseDataInner) GetPingBeforeAssignEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.PingBeforeAssignEnabled) {
		return nil, false
	}
	return o.PingBeforeAssignEnabled, true
}

// HasPingBeforeAssignEnabled returns a boolean if a field has been set.
func (o *GetNetworks200ResponseDataInner) HasPingBeforeAssignEnabled() bool {
	if o != nil && !IsNil(o.PingBeforeAssignEnabled) {
		return true
	}

	return false
}

// SetPingBeforeAssignEnabled gets a reference to the given bool and assigns it to the PingBeforeAssignEnabled field.
func (o *GetNetworks200ResponseDataInner) SetPingBeforeAssignEnabled(v bool) {
	o.PingBeforeAssignEnabled = &v
}

// GetDefaultZonesInherited returns the DefaultZonesInherited field value if set, zero value otherwise.
func (o *GetNetworks200ResponseDataInner) GetDefaultZonesInherited() bool {
	if o == nil || IsNil(o.DefaultZonesInherited) {
		var ret bool
		return ret
	}
	return *o.DefaultZonesInherited
}

// GetDefaultZonesInheritedOk returns a tuple with the DefaultZonesInherited field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworks200ResponseDataInner) GetDefaultZonesInheritedOk() (*bool, bool) {
	if o == nil || IsNil(o.DefaultZonesInherited) {
		return nil, false
	}
	return o.DefaultZonesInherited, true
}

// HasDefaultZonesInherited returns a boolean if a field has been set.
func (o *GetNetworks200ResponseDataInner) HasDefaultZonesInherited() bool {
	if o != nil && !IsNil(o.DefaultZonesInherited) {
		return true
	}

	return false
}

// SetDefaultZonesInherited gets a reference to the given bool and assigns it to the DefaultZonesInherited field.
func (o *GetNetworks200ResponseDataInner) SetDefaultZonesInherited(v bool) {
	o.DefaultZonesInherited = &v
}

// GetDefaultZones returns the DefaultZones field value if set, zero value otherwise.
func (o *GetNetworks200ResponseDataInner) GetDefaultZones() []IPv4NetworkAllOfDefaultZones {
	if o == nil || IsNil(o.DefaultZones) {
		var ret []IPv4NetworkAllOfDefaultZones
		return ret
	}
	return o.DefaultZones
}

// GetDefaultZonesOk returns a tuple with the DefaultZones field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworks200ResponseDataInner) GetDefaultZonesOk() ([]IPv4NetworkAllOfDefaultZones, bool) {
	if o == nil || IsNil(o.DefaultZones) {
		return nil, false
	}
	return o.DefaultZones, true
}

// HasDefaultZones returns a boolean if a field has been set.
func (o *GetNetworks200ResponseDataInner) HasDefaultZones() bool {
	if o != nil && !IsNil(o.DefaultZones) {
		return true
	}

	return false
}

// SetDefaultZones gets a reference to the given []IPv4NetworkAllOfDefaultZones and assigns it to the DefaultZones field.
func (o *GetNetworks200ResponseDataInner) SetDefaultZones(v []IPv4NetworkAllOfDefaultZones) {
	o.DefaultZones = v
}

// GetRestrictedZonesInherited returns the RestrictedZonesInherited field value if set, zero value otherwise.
func (o *GetNetworks200ResponseDataInner) GetRestrictedZonesInherited() bool {
	if o == nil || IsNil(o.RestrictedZonesInherited) {
		var ret bool
		return ret
	}
	return *o.RestrictedZonesInherited
}

// GetRestrictedZonesInheritedOk returns a tuple with the RestrictedZonesInherited field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworks200ResponseDataInner) GetRestrictedZonesInheritedOk() (*bool, bool) {
	if o == nil || IsNil(o.RestrictedZonesInherited) {
		return nil, false
	}
	return o.RestrictedZonesInherited, true
}

// HasRestrictedZonesInherited returns a boolean if a field has been set.
func (o *GetNetworks200ResponseDataInner) HasRestrictedZonesInherited() bool {
	if o != nil && !IsNil(o.RestrictedZonesInherited) {
		return true
	}

	return false
}

// SetRestrictedZonesInherited gets a reference to the given bool and assigns it to the RestrictedZonesInherited field.
func (o *GetNetworks200ResponseDataInner) SetRestrictedZonesInherited(v bool) {
	o.RestrictedZonesInherited = &v
}

// GetRestrictedZones returns the RestrictedZones field value if set, zero value otherwise.
func (o *GetNetworks200ResponseDataInner) GetRestrictedZones() []IPv4NetworkAllOfRestrictedZones {
	if o == nil || IsNil(o.RestrictedZones) {
		var ret []IPv4NetworkAllOfRestrictedZones
		return ret
	}
	return o.RestrictedZones
}

// GetRestrictedZonesOk returns a tuple with the RestrictedZones field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworks200ResponseDataInner) GetRestrictedZonesOk() ([]IPv4NetworkAllOfRestrictedZones, bool) {
	if o == nil || IsNil(o.RestrictedZones) {
		return nil, false
	}
	return o.RestrictedZones, true
}

// HasRestrictedZones returns a boolean if a field has been set.
func (o *GetNetworks200ResponseDataInner) HasRestrictedZones() bool {
	if o != nil && !IsNil(o.RestrictedZones) {
		return true
	}

	return false
}

// SetRestrictedZones gets a reference to the given []IPv4NetworkAllOfRestrictedZones and assigns it to the RestrictedZones field.
func (o *GetNetworks200ResponseDataInner) SetRestrictedZones(v []IPv4NetworkAllOfRestrictedZones) {
	o.RestrictedZones = v
}

// GetLowWaterMark returns the LowWaterMark field value if set, zero value otherwise.
func (o *GetNetworks200ResponseDataInner) GetLowWaterMark() int32 {
	if o == nil || IsNil(o.LowWaterMark) {
		var ret int32
		return ret
	}
	return *o.LowWaterMark
}

// GetLowWaterMarkOk returns a tuple with the LowWaterMark field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworks200ResponseDataInner) GetLowWaterMarkOk() (*int32, bool) {
	if o == nil || IsNil(o.LowWaterMark) {
		return nil, false
	}
	return o.LowWaterMark, true
}

// HasLowWaterMark returns a boolean if a field has been set.
func (o *GetNetworks200ResponseDataInner) HasLowWaterMark() bool {
	if o != nil && !IsNil(o.LowWaterMark) {
		return true
	}

	return false
}

// SetLowWaterMark gets a reference to the given int32 and assigns it to the LowWaterMark field.
func (o *GetNetworks200ResponseDataInner) SetLowWaterMark(v int32) {
	o.LowWaterMark = &v
}

// GetHighWaterMark returns the HighWaterMark field value if set, zero value otherwise.
func (o *GetNetworks200ResponseDataInner) GetHighWaterMark() int32 {
	if o == nil || IsNil(o.HighWaterMark) {
		var ret int32
		return ret
	}
	return *o.HighWaterMark
}

// GetHighWaterMarkOk returns a tuple with the HighWaterMark field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworks200ResponseDataInner) GetHighWaterMarkOk() (*int32, bool) {
	if o == nil || IsNil(o.HighWaterMark) {
		return nil, false
	}
	return o.HighWaterMark, true
}

// HasHighWaterMark returns a boolean if a field has been set.
func (o *GetNetworks200ResponseDataInner) HasHighWaterMark() bool {
	if o != nil && !IsNil(o.HighWaterMark) {
		return true
	}

	return false
}

// SetHighWaterMark gets a reference to the given int32 and assigns it to the HighWaterMark field.
func (o *GetNetworks200ResponseDataInner) SetHighWaterMark(v int32) {
	o.HighWaterMark = &v
}

// GetReverseZoneSigned returns the ReverseZoneSigned field value if set, zero value otherwise.
func (o *GetNetworks200ResponseDataInner) GetReverseZoneSigned() bool {
	if o == nil || IsNil(o.ReverseZoneSigned) {
		var ret bool
		return ret
	}
	return *o.ReverseZoneSigned
}

// GetReverseZoneSignedOk returns a tuple with the ReverseZoneSigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworks200ResponseDataInner) GetReverseZoneSignedOk() (*bool, bool) {
	if o == nil || IsNil(o.ReverseZoneSigned) {
		return nil, false
	}
	return o.ReverseZoneSigned, true
}

// HasReverseZoneSigned returns a boolean if a field has been set.
func (o *GetNetworks200ResponseDataInner) HasReverseZoneSigned() bool {
	if o != nil && !IsNil(o.ReverseZoneSigned) {
		return true
	}

	return false
}

// SetReverseZoneSigned gets a reference to the given bool and assigns it to the ReverseZoneSigned field.
func (o *GetNetworks200ResponseDataInner) SetReverseZoneSigned(v bool) {
	o.ReverseZoneSigned = &v
}

// GetReverseZoneSigningPolicy returns the ReverseZoneSigningPolicy field value if set, zero value otherwise.
func (o *GetNetworks200ResponseDataInner) GetReverseZoneSigningPolicy() InlinedDNSSECSigningPolicy {
	if o == nil || IsNil(o.ReverseZoneSigningPolicy) {
		var ret InlinedDNSSECSigningPolicy
		return ret
	}
	return *o.ReverseZoneSigningPolicy
}

// GetReverseZoneSigningPolicyOk returns a tuple with the ReverseZoneSigningPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworks200ResponseDataInner) GetReverseZoneSigningPolicyOk() (*InlinedDNSSECSigningPolicy, bool) {
	if o == nil || IsNil(o.ReverseZoneSigningPolicy) {
		return nil, false
	}
	return o.ReverseZoneSigningPolicy, true
}

// HasReverseZoneSigningPolicy returns a boolean if a field has been set.
func (o *GetNetworks200ResponseDataInner) HasReverseZoneSigningPolicy() bool {
	if o != nil && !IsNil(o.ReverseZoneSigningPolicy) {
		return true
	}

	return false
}

// SetReverseZoneSigningPolicy gets a reference to the given InlinedDNSSECSigningPolicy and assigns it to the ReverseZoneSigningPolicy field.
func (o *GetNetworks200ResponseDataInner) SetReverseZoneSigningPolicy(v InlinedDNSSECSigningPolicy) {
	o.ReverseZoneSigningPolicy = &v
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *GetNetworks200ResponseDataInner) GetUsage() map[string]int64 {
	if o == nil || IsNil(o.Usage) {
		var ret map[string]int64
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworks200ResponseDataInner) GetUsageOk() (*map[string]int64, bool) {
	if o == nil || IsNil(o.Usage) {
		return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *GetNetworks200ResponseDataInner) HasUsage() bool {
	if o != nil && !IsNil(o.Usage) {
		return true
	}

	return false
}

// SetUsage gets a reference to the given map[string]int64 and assigns it to the Usage field.
func (o *GetNetworks200ResponseDataInner) SetUsage(v map[string]int64) {
	o.Usage = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GetNetworks200ResponseDataInner) GetLinks() ResourceLinks {
	if o == nil || IsNil(o.Links) {
		var ret ResourceLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworks200ResponseDataInner) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GetNetworks200ResponseDataInner) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ResourceLinks and assigns it to the Links field.
func (o *GetNetworks200ResponseDataInner) SetLinks(v ResourceLinks) {
	o.Links = &v
}

func (o GetNetworks200ResponseDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworks200ResponseDataInner) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Range) {
		toSerialize["range"] = o.Range
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !IsNil(o.InheritedFields) {
		toSerialize["_inheritedFields"] = o.InheritedFields
	}
	if !IsNil(o.Gateway) {
		toSerialize["gateway"] = o.Gateway
	}
	if !IsNil(o.Template) {
		toSerialize["template"] = o.Template
	}
	if !IsNil(o.RouterPortInfo) {
		toSerialize["routerPortInfo"] = o.RouterPortInfo
	}
	if !IsNil(o.DefaultView) {
		toSerialize["defaultView"] = o.DefaultView
	}
	if !IsNil(o.SharedNetworkTag) {
		toSerialize["sharedNetworkTag"] = o.SharedNetworkTag
	}
	if !IsNil(o.DuplicateHostnamesAllowed) {
		toSerialize["duplicateHostnamesAllowed"] = o.DuplicateHostnamesAllowed
	}
	if !IsNil(o.PingBeforeAssignEnabled) {
		toSerialize["pingBeforeAssignEnabled"] = o.PingBeforeAssignEnabled
	}
	if !IsNil(o.DefaultZonesInherited) {
		toSerialize["defaultZonesInherited"] = o.DefaultZonesInherited
	}
	if !IsNil(o.DefaultZones) {
		toSerialize["defaultZones"] = o.DefaultZones
	}
	if !IsNil(o.RestrictedZonesInherited) {
		toSerialize["restrictedZonesInherited"] = o.RestrictedZonesInherited
	}
	if !IsNil(o.RestrictedZones) {
		toSerialize["restrictedZones"] = o.RestrictedZones
	}
	if !IsNil(o.LowWaterMark) {
		toSerialize["lowWaterMark"] = o.LowWaterMark
	}
	if !IsNil(o.HighWaterMark) {
		toSerialize["highWaterMark"] = o.HighWaterMark
	}
	if !IsNil(o.ReverseZoneSigned) {
		toSerialize["reverseZoneSigned"] = o.ReverseZoneSigned
	}
	if !IsNil(o.ReverseZoneSigningPolicy) {
		toSerialize["reverseZoneSigningPolicy"] = o.ReverseZoneSigningPolicy
	}
	if !IsNil(o.Usage) {
		toSerialize["usage"] = o.Usage
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableGetNetworks200ResponseDataInner struct {
	value *GetNetworks200ResponseDataInner
	isSet bool
}

func (v NullableGetNetworks200ResponseDataInner) Get() *GetNetworks200ResponseDataInner {
	return v.value
}

func (v *NullableGetNetworks200ResponseDataInner) Set(val *GetNetworks200ResponseDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworks200ResponseDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworks200ResponseDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworks200ResponseDataInner(val *GetNetworks200ResponseDataInner) *NullableGetNetworks200ResponseDataInner {
	return &NullableGetNetworks200ResponseDataInner{value: val, isSet: true}
}

func (v NullableGetNetworks200ResponseDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworks200ResponseDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


