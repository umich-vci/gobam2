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

// checks if the SyslogService type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SyslogService{}

// SyslogService struct for SyslogService
type SyslogService struct {
	Service
	// The resource type.
	Type *string `json:"type,omitempty"`
	// Indicates whether the syslog redirection service configuration has been manually overridden.
	ManualOverrideEnabled *bool `json:"manualOverrideEnabled,omitempty"`
	// Indicates whether the ISO 8601 date time format is used when logging timestamps.
	Iso8601TimestampsEnabled *bool `json:"iso8601TimestampsEnabled,omitempty"`
	// The list of servers to redirect syslog messages to.
	Servers []SyslogRedirectionServerBean `json:"servers,omitempty"`
}

// NewSyslogService instantiates a new SyslogService object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyslogService() *SyslogService {
	this := SyslogService{}
	return &this
}

// NewSyslogServiceWithDefaults instantiates a new SyslogService object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyslogServiceWithDefaults() *SyslogService {
	this := SyslogService{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SyslogService) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyslogService) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SyslogService) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SyslogService) SetType(v string) {
	o.Type = &v
}

// GetManualOverrideEnabled returns the ManualOverrideEnabled field value if set, zero value otherwise.
func (o *SyslogService) GetManualOverrideEnabled() bool {
	if o == nil || IsNil(o.ManualOverrideEnabled) {
		var ret bool
		return ret
	}
	return *o.ManualOverrideEnabled
}

// GetManualOverrideEnabledOk returns a tuple with the ManualOverrideEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyslogService) GetManualOverrideEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.ManualOverrideEnabled) {
		return nil, false
	}
	return o.ManualOverrideEnabled, true
}

// HasManualOverrideEnabled returns a boolean if a field has been set.
func (o *SyslogService) HasManualOverrideEnabled() bool {
	if o != nil && !IsNil(o.ManualOverrideEnabled) {
		return true
	}

	return false
}

// SetManualOverrideEnabled gets a reference to the given bool and assigns it to the ManualOverrideEnabled field.
func (o *SyslogService) SetManualOverrideEnabled(v bool) {
	o.ManualOverrideEnabled = &v
}

// GetIso8601TimestampsEnabled returns the Iso8601TimestampsEnabled field value if set, zero value otherwise.
func (o *SyslogService) GetIso8601TimestampsEnabled() bool {
	if o == nil || IsNil(o.Iso8601TimestampsEnabled) {
		var ret bool
		return ret
	}
	return *o.Iso8601TimestampsEnabled
}

// GetIso8601TimestampsEnabledOk returns a tuple with the Iso8601TimestampsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyslogService) GetIso8601TimestampsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Iso8601TimestampsEnabled) {
		return nil, false
	}
	return o.Iso8601TimestampsEnabled, true
}

// HasIso8601TimestampsEnabled returns a boolean if a field has been set.
func (o *SyslogService) HasIso8601TimestampsEnabled() bool {
	if o != nil && !IsNil(o.Iso8601TimestampsEnabled) {
		return true
	}

	return false
}

// SetIso8601TimestampsEnabled gets a reference to the given bool and assigns it to the Iso8601TimestampsEnabled field.
func (o *SyslogService) SetIso8601TimestampsEnabled(v bool) {
	o.Iso8601TimestampsEnabled = &v
}

// GetServers returns the Servers field value if set, zero value otherwise.
func (o *SyslogService) GetServers() []SyslogRedirectionServerBean {
	if o == nil || IsNil(o.Servers) {
		var ret []SyslogRedirectionServerBean
		return ret
	}
	return o.Servers
}

// GetServersOk returns a tuple with the Servers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyslogService) GetServersOk() ([]SyslogRedirectionServerBean, bool) {
	if o == nil || IsNil(o.Servers) {
		return nil, false
	}
	return o.Servers, true
}

// HasServers returns a boolean if a field has been set.
func (o *SyslogService) HasServers() bool {
	if o != nil && !IsNil(o.Servers) {
		return true
	}

	return false
}

// SetServers gets a reference to the given []SyslogRedirectionServerBean and assigns it to the Servers field.
func (o *SyslogService) SetServers(v []SyslogRedirectionServerBean) {
	o.Servers = v
}

func (o SyslogService) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SyslogService) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.ManualOverrideEnabled) {
		toSerialize["manualOverrideEnabled"] = o.ManualOverrideEnabled
	}
	if !IsNil(o.Iso8601TimestampsEnabled) {
		toSerialize["iso8601TimestampsEnabled"] = o.Iso8601TimestampsEnabled
	}
	if !IsNil(o.Servers) {
		toSerialize["servers"] = o.Servers
	}
	return toSerialize, nil
}

type NullableSyslogService struct {
	value *SyslogService
	isSet bool
}

func (v NullableSyslogService) Get() *SyslogService {
	return v.value
}

func (v *NullableSyslogService) Set(val *SyslogService) {
	v.value = val
	v.isSet = true
}

func (v NullableSyslogService) IsSet() bool {
	return v.isSet
}

func (v *NullableSyslogService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyslogService(val *SyslogService) *NullableSyslogService {
	return &NullableSyslogService{value: val, isSet: true}
}

func (v NullableSyslogService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyslogService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

