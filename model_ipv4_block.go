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

// checks if the IPv4Block type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IPv4Block{}

// IPv4Block struct for IPv4Block
type IPv4Block struct {
	Block
	// The type of IP block.
	Type *string `json:"type,omitempty"`
	Template *InlinedIPv4Template `json:"template,omitempty"`
	DefaultView *InlinedView `json:"defaultView,omitempty"`
	// Indicates whether duplicate hostnames are allowed in the block. If enabled, duplicate hostnames will be allowed. If not set explicitly, this value will be inherited from the parent block.
	DuplicateHostnamesAllowed *bool `json:"duplicateHostnamesAllowed,omitempty"`
	// If set to true, Address Manager will ping IP addresses before assigning them. If set to false, Address Manager will assign addresses without checking their availability.
	PingBeforeAssignEnabled *bool `json:"pingBeforeAssignEnabled,omitempty"`
	// If set to true, the IPv4 block will inherit the parent resource's default domains configuration if possible.
	DefaultZonesInherited *bool `json:"defaultZonesInherited,omitempty"`
	DefaultZones []IPv4BlockAllOfDefaultZones `json:"defaultZones,omitempty"`
	// If set to true, the IPv4 block will inherit the parent resource's DNS restriction configuration if possible.
	RestrictedZonesInherited *bool `json:"restrictedZonesInherited,omitempty"`
	RestrictedZones []IPv4BlockAllOfRestrictedZones `json:"restrictedZones,omitempty"`
	// A DHCP alert is triggered when usage falls below this percentage (when too few addresses are in use).
	LowWaterMark *int32 `json:"lowWaterMark,omitempty"`
	// A DHCP alert is triggered when usage exceeds this percentage (when too many addresses are in use).
	HighWaterMark *int32 `json:"highWaterMark,omitempty"`
	// If set to true, the reverse zone associated with the IPv4 block will be signed according to the DNSSEC signing policy in reverseZoneSigningPolicy.
	ReverseZoneSigned *bool `json:"reverseZoneSigned,omitempty"`
	ReverseZoneSigningPolicy *InlinedDNSSECSigningPolicy `json:"reverseZoneSigningPolicy,omitempty"`
	// Percentage values indicating the current amount of allocated and unallocated IPv4 addresses within the block.
	UsagePercentage *map[string]int32 `json:"usagePercentage,omitempty"`
}

// NewIPv4Block instantiates a new IPv4Block object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIPv4Block() *IPv4Block {
	this := IPv4Block{}
	return &this
}

// NewIPv4BlockWithDefaults instantiates a new IPv4Block object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIPv4BlockWithDefaults() *IPv4Block {
	this := IPv4Block{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IPv4Block) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPv4Block) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IPv4Block) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *IPv4Block) SetType(v string) {
	o.Type = &v
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *IPv4Block) GetTemplate() InlinedIPv4Template {
	if o == nil || IsNil(o.Template) {
		var ret InlinedIPv4Template
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPv4Block) GetTemplateOk() (*InlinedIPv4Template, bool) {
	if o == nil || IsNil(o.Template) {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *IPv4Block) HasTemplate() bool {
	if o != nil && !IsNil(o.Template) {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given InlinedIPv4Template and assigns it to the Template field.
func (o *IPv4Block) SetTemplate(v InlinedIPv4Template) {
	o.Template = &v
}

// GetDefaultView returns the DefaultView field value if set, zero value otherwise.
func (o *IPv4Block) GetDefaultView() InlinedView {
	if o == nil || IsNil(o.DefaultView) {
		var ret InlinedView
		return ret
	}
	return *o.DefaultView
}

// GetDefaultViewOk returns a tuple with the DefaultView field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPv4Block) GetDefaultViewOk() (*InlinedView, bool) {
	if o == nil || IsNil(o.DefaultView) {
		return nil, false
	}
	return o.DefaultView, true
}

// HasDefaultView returns a boolean if a field has been set.
func (o *IPv4Block) HasDefaultView() bool {
	if o != nil && !IsNil(o.DefaultView) {
		return true
	}

	return false
}

// SetDefaultView gets a reference to the given InlinedView and assigns it to the DefaultView field.
func (o *IPv4Block) SetDefaultView(v InlinedView) {
	o.DefaultView = &v
}

// GetDuplicateHostnamesAllowed returns the DuplicateHostnamesAllowed field value if set, zero value otherwise.
func (o *IPv4Block) GetDuplicateHostnamesAllowed() bool {
	if o == nil || IsNil(o.DuplicateHostnamesAllowed) {
		var ret bool
		return ret
	}
	return *o.DuplicateHostnamesAllowed
}

// GetDuplicateHostnamesAllowedOk returns a tuple with the DuplicateHostnamesAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPv4Block) GetDuplicateHostnamesAllowedOk() (*bool, bool) {
	if o == nil || IsNil(o.DuplicateHostnamesAllowed) {
		return nil, false
	}
	return o.DuplicateHostnamesAllowed, true
}

// HasDuplicateHostnamesAllowed returns a boolean if a field has been set.
func (o *IPv4Block) HasDuplicateHostnamesAllowed() bool {
	if o != nil && !IsNil(o.DuplicateHostnamesAllowed) {
		return true
	}

	return false
}

// SetDuplicateHostnamesAllowed gets a reference to the given bool and assigns it to the DuplicateHostnamesAllowed field.
func (o *IPv4Block) SetDuplicateHostnamesAllowed(v bool) {
	o.DuplicateHostnamesAllowed = &v
}

// GetPingBeforeAssignEnabled returns the PingBeforeAssignEnabled field value if set, zero value otherwise.
func (o *IPv4Block) GetPingBeforeAssignEnabled() bool {
	if o == nil || IsNil(o.PingBeforeAssignEnabled) {
		var ret bool
		return ret
	}
	return *o.PingBeforeAssignEnabled
}

// GetPingBeforeAssignEnabledOk returns a tuple with the PingBeforeAssignEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPv4Block) GetPingBeforeAssignEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.PingBeforeAssignEnabled) {
		return nil, false
	}
	return o.PingBeforeAssignEnabled, true
}

// HasPingBeforeAssignEnabled returns a boolean if a field has been set.
func (o *IPv4Block) HasPingBeforeAssignEnabled() bool {
	if o != nil && !IsNil(o.PingBeforeAssignEnabled) {
		return true
	}

	return false
}

// SetPingBeforeAssignEnabled gets a reference to the given bool and assigns it to the PingBeforeAssignEnabled field.
func (o *IPv4Block) SetPingBeforeAssignEnabled(v bool) {
	o.PingBeforeAssignEnabled = &v
}

// GetDefaultZonesInherited returns the DefaultZonesInherited field value if set, zero value otherwise.
func (o *IPv4Block) GetDefaultZonesInherited() bool {
	if o == nil || IsNil(o.DefaultZonesInherited) {
		var ret bool
		return ret
	}
	return *o.DefaultZonesInherited
}

// GetDefaultZonesInheritedOk returns a tuple with the DefaultZonesInherited field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPv4Block) GetDefaultZonesInheritedOk() (*bool, bool) {
	if o == nil || IsNil(o.DefaultZonesInherited) {
		return nil, false
	}
	return o.DefaultZonesInherited, true
}

// HasDefaultZonesInherited returns a boolean if a field has been set.
func (o *IPv4Block) HasDefaultZonesInherited() bool {
	if o != nil && !IsNil(o.DefaultZonesInherited) {
		return true
	}

	return false
}

// SetDefaultZonesInherited gets a reference to the given bool and assigns it to the DefaultZonesInherited field.
func (o *IPv4Block) SetDefaultZonesInherited(v bool) {
	o.DefaultZonesInherited = &v
}

// GetDefaultZones returns the DefaultZones field value if set, zero value otherwise.
func (o *IPv4Block) GetDefaultZones() []IPv4BlockAllOfDefaultZones {
	if o == nil || IsNil(o.DefaultZones) {
		var ret []IPv4BlockAllOfDefaultZones
		return ret
	}
	return o.DefaultZones
}

// GetDefaultZonesOk returns a tuple with the DefaultZones field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPv4Block) GetDefaultZonesOk() ([]IPv4BlockAllOfDefaultZones, bool) {
	if o == nil || IsNil(o.DefaultZones) {
		return nil, false
	}
	return o.DefaultZones, true
}

// HasDefaultZones returns a boolean if a field has been set.
func (o *IPv4Block) HasDefaultZones() bool {
	if o != nil && !IsNil(o.DefaultZones) {
		return true
	}

	return false
}

// SetDefaultZones gets a reference to the given []IPv4BlockAllOfDefaultZones and assigns it to the DefaultZones field.
func (o *IPv4Block) SetDefaultZones(v []IPv4BlockAllOfDefaultZones) {
	o.DefaultZones = v
}

// GetRestrictedZonesInherited returns the RestrictedZonesInherited field value if set, zero value otherwise.
func (o *IPv4Block) GetRestrictedZonesInherited() bool {
	if o == nil || IsNil(o.RestrictedZonesInherited) {
		var ret bool
		return ret
	}
	return *o.RestrictedZonesInherited
}

// GetRestrictedZonesInheritedOk returns a tuple with the RestrictedZonesInherited field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPv4Block) GetRestrictedZonesInheritedOk() (*bool, bool) {
	if o == nil || IsNil(o.RestrictedZonesInherited) {
		return nil, false
	}
	return o.RestrictedZonesInherited, true
}

// HasRestrictedZonesInherited returns a boolean if a field has been set.
func (o *IPv4Block) HasRestrictedZonesInherited() bool {
	if o != nil && !IsNil(o.RestrictedZonesInherited) {
		return true
	}

	return false
}

// SetRestrictedZonesInherited gets a reference to the given bool and assigns it to the RestrictedZonesInherited field.
func (o *IPv4Block) SetRestrictedZonesInherited(v bool) {
	o.RestrictedZonesInherited = &v
}

// GetRestrictedZones returns the RestrictedZones field value if set, zero value otherwise.
func (o *IPv4Block) GetRestrictedZones() []IPv4BlockAllOfRestrictedZones {
	if o == nil || IsNil(o.RestrictedZones) {
		var ret []IPv4BlockAllOfRestrictedZones
		return ret
	}
	return o.RestrictedZones
}

// GetRestrictedZonesOk returns a tuple with the RestrictedZones field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPv4Block) GetRestrictedZonesOk() ([]IPv4BlockAllOfRestrictedZones, bool) {
	if o == nil || IsNil(o.RestrictedZones) {
		return nil, false
	}
	return o.RestrictedZones, true
}

// HasRestrictedZones returns a boolean if a field has been set.
func (o *IPv4Block) HasRestrictedZones() bool {
	if o != nil && !IsNil(o.RestrictedZones) {
		return true
	}

	return false
}

// SetRestrictedZones gets a reference to the given []IPv4BlockAllOfRestrictedZones and assigns it to the RestrictedZones field.
func (o *IPv4Block) SetRestrictedZones(v []IPv4BlockAllOfRestrictedZones) {
	o.RestrictedZones = v
}

// GetLowWaterMark returns the LowWaterMark field value if set, zero value otherwise.
func (o *IPv4Block) GetLowWaterMark() int32 {
	if o == nil || IsNil(o.LowWaterMark) {
		var ret int32
		return ret
	}
	return *o.LowWaterMark
}

// GetLowWaterMarkOk returns a tuple with the LowWaterMark field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPv4Block) GetLowWaterMarkOk() (*int32, bool) {
	if o == nil || IsNil(o.LowWaterMark) {
		return nil, false
	}
	return o.LowWaterMark, true
}

// HasLowWaterMark returns a boolean if a field has been set.
func (o *IPv4Block) HasLowWaterMark() bool {
	if o != nil && !IsNil(o.LowWaterMark) {
		return true
	}

	return false
}

// SetLowWaterMark gets a reference to the given int32 and assigns it to the LowWaterMark field.
func (o *IPv4Block) SetLowWaterMark(v int32) {
	o.LowWaterMark = &v
}

// GetHighWaterMark returns the HighWaterMark field value if set, zero value otherwise.
func (o *IPv4Block) GetHighWaterMark() int32 {
	if o == nil || IsNil(o.HighWaterMark) {
		var ret int32
		return ret
	}
	return *o.HighWaterMark
}

// GetHighWaterMarkOk returns a tuple with the HighWaterMark field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPv4Block) GetHighWaterMarkOk() (*int32, bool) {
	if o == nil || IsNil(o.HighWaterMark) {
		return nil, false
	}
	return o.HighWaterMark, true
}

// HasHighWaterMark returns a boolean if a field has been set.
func (o *IPv4Block) HasHighWaterMark() bool {
	if o != nil && !IsNil(o.HighWaterMark) {
		return true
	}

	return false
}

// SetHighWaterMark gets a reference to the given int32 and assigns it to the HighWaterMark field.
func (o *IPv4Block) SetHighWaterMark(v int32) {
	o.HighWaterMark = &v
}

// GetReverseZoneSigned returns the ReverseZoneSigned field value if set, zero value otherwise.
func (o *IPv4Block) GetReverseZoneSigned() bool {
	if o == nil || IsNil(o.ReverseZoneSigned) {
		var ret bool
		return ret
	}
	return *o.ReverseZoneSigned
}

// GetReverseZoneSignedOk returns a tuple with the ReverseZoneSigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPv4Block) GetReverseZoneSignedOk() (*bool, bool) {
	if o == nil || IsNil(o.ReverseZoneSigned) {
		return nil, false
	}
	return o.ReverseZoneSigned, true
}

// HasReverseZoneSigned returns a boolean if a field has been set.
func (o *IPv4Block) HasReverseZoneSigned() bool {
	if o != nil && !IsNil(o.ReverseZoneSigned) {
		return true
	}

	return false
}

// SetReverseZoneSigned gets a reference to the given bool and assigns it to the ReverseZoneSigned field.
func (o *IPv4Block) SetReverseZoneSigned(v bool) {
	o.ReverseZoneSigned = &v
}

// GetReverseZoneSigningPolicy returns the ReverseZoneSigningPolicy field value if set, zero value otherwise.
func (o *IPv4Block) GetReverseZoneSigningPolicy() InlinedDNSSECSigningPolicy {
	if o == nil || IsNil(o.ReverseZoneSigningPolicy) {
		var ret InlinedDNSSECSigningPolicy
		return ret
	}
	return *o.ReverseZoneSigningPolicy
}

// GetReverseZoneSigningPolicyOk returns a tuple with the ReverseZoneSigningPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPv4Block) GetReverseZoneSigningPolicyOk() (*InlinedDNSSECSigningPolicy, bool) {
	if o == nil || IsNil(o.ReverseZoneSigningPolicy) {
		return nil, false
	}
	return o.ReverseZoneSigningPolicy, true
}

// HasReverseZoneSigningPolicy returns a boolean if a field has been set.
func (o *IPv4Block) HasReverseZoneSigningPolicy() bool {
	if o != nil && !IsNil(o.ReverseZoneSigningPolicy) {
		return true
	}

	return false
}

// SetReverseZoneSigningPolicy gets a reference to the given InlinedDNSSECSigningPolicy and assigns it to the ReverseZoneSigningPolicy field.
func (o *IPv4Block) SetReverseZoneSigningPolicy(v InlinedDNSSECSigningPolicy) {
	o.ReverseZoneSigningPolicy = &v
}

// GetUsagePercentage returns the UsagePercentage field value if set, zero value otherwise.
func (o *IPv4Block) GetUsagePercentage() map[string]int32 {
	if o == nil || IsNil(o.UsagePercentage) {
		var ret map[string]int32
		return ret
	}
	return *o.UsagePercentage
}

// GetUsagePercentageOk returns a tuple with the UsagePercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPv4Block) GetUsagePercentageOk() (*map[string]int32, bool) {
	if o == nil || IsNil(o.UsagePercentage) {
		return nil, false
	}
	return o.UsagePercentage, true
}

// HasUsagePercentage returns a boolean if a field has been set.
func (o *IPv4Block) HasUsagePercentage() bool {
	if o != nil && !IsNil(o.UsagePercentage) {
		return true
	}

	return false
}

// SetUsagePercentage gets a reference to the given map[string]int32 and assigns it to the UsagePercentage field.
func (o *IPv4Block) SetUsagePercentage(v map[string]int32) {
	o.UsagePercentage = &v
}

func (o IPv4Block) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IPv4Block) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Template) {
		toSerialize["template"] = o.Template
	}
	if !IsNil(o.DefaultView) {
		toSerialize["defaultView"] = o.DefaultView
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
	if !IsNil(o.UsagePercentage) {
		toSerialize["usagePercentage"] = o.UsagePercentage
	}
	return toSerialize, nil
}

type NullableIPv4Block struct {
	value *IPv4Block
	isSet bool
}

func (v NullableIPv4Block) Get() *IPv4Block {
	return v.value
}

func (v *NullableIPv4Block) Set(val *IPv4Block) {
	v.value = val
	v.isSet = true
}

func (v NullableIPv4Block) IsSet() bool {
	return v.isSet
}

func (v *NullableIPv4Block) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIPv4Block(val *IPv4Block) *NullableIPv4Block {
	return &NullableIPv4Block{value: val, isSet: true}
}

func (v NullableIPv4Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIPv4Block) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


