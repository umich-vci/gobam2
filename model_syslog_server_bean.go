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

// checks if the SyslogServerBean type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SyslogServerBean{}

// SyslogServerBean struct for SyslogServerBean
type SyslogServerBean struct {
	SyslogRedirectionServerBean
	// The port used to connect to the syslog server.
	Port *int32 `json:"port,omitempty"`
	// The syslog protocol used to format syslog messages.
	SyslogProtocol *string `json:"syslogProtocol,omitempty"`
	// The severity level of messages that are sent to the syslog server.
	SeverityLevel *string `json:"severityLevel,omitempty"`
	Transport *SyslogTransportBean `json:"transport,omitempty"`
	// Indicates whether the ISO 8601 timestamp format is used for syslog messages.
	Iso8601TimestampsEnabled *bool `json:"iso8601TimestampsEnabled,omitempty"`
	// The services for which syslog messages are generated.
	ServiceTypes []string `json:"serviceTypes,omitempty"`
}

// NewSyslogServerBean instantiates a new SyslogServerBean object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyslogServerBean() *SyslogServerBean {
	this := SyslogServerBean{}
	return &this
}

// NewSyslogServerBeanWithDefaults instantiates a new SyslogServerBean object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyslogServerBeanWithDefaults() *SyslogServerBean {
	this := SyslogServerBean{}
	return &this
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *SyslogServerBean) GetPort() int32 {
	if o == nil || IsNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyslogServerBean) GetPortOk() (*int32, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *SyslogServerBean) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *SyslogServerBean) SetPort(v int32) {
	o.Port = &v
}

// GetSyslogProtocol returns the SyslogProtocol field value if set, zero value otherwise.
func (o *SyslogServerBean) GetSyslogProtocol() string {
	if o == nil || IsNil(o.SyslogProtocol) {
		var ret string
		return ret
	}
	return *o.SyslogProtocol
}

// GetSyslogProtocolOk returns a tuple with the SyslogProtocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyslogServerBean) GetSyslogProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.SyslogProtocol) {
		return nil, false
	}
	return o.SyslogProtocol, true
}

// HasSyslogProtocol returns a boolean if a field has been set.
func (o *SyslogServerBean) HasSyslogProtocol() bool {
	if o != nil && !IsNil(o.SyslogProtocol) {
		return true
	}

	return false
}

// SetSyslogProtocol gets a reference to the given string and assigns it to the SyslogProtocol field.
func (o *SyslogServerBean) SetSyslogProtocol(v string) {
	o.SyslogProtocol = &v
}

// GetSeverityLevel returns the SeverityLevel field value if set, zero value otherwise.
func (o *SyslogServerBean) GetSeverityLevel() string {
	if o == nil || IsNil(o.SeverityLevel) {
		var ret string
		return ret
	}
	return *o.SeverityLevel
}

// GetSeverityLevelOk returns a tuple with the SeverityLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyslogServerBean) GetSeverityLevelOk() (*string, bool) {
	if o == nil || IsNil(o.SeverityLevel) {
		return nil, false
	}
	return o.SeverityLevel, true
}

// HasSeverityLevel returns a boolean if a field has been set.
func (o *SyslogServerBean) HasSeverityLevel() bool {
	if o != nil && !IsNil(o.SeverityLevel) {
		return true
	}

	return false
}

// SetSeverityLevel gets a reference to the given string and assigns it to the SeverityLevel field.
func (o *SyslogServerBean) SetSeverityLevel(v string) {
	o.SeverityLevel = &v
}

// GetTransport returns the Transport field value if set, zero value otherwise.
func (o *SyslogServerBean) GetTransport() SyslogTransportBean {
	if o == nil || IsNil(o.Transport) {
		var ret SyslogTransportBean
		return ret
	}
	return *o.Transport
}

// GetTransportOk returns a tuple with the Transport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyslogServerBean) GetTransportOk() (*SyslogTransportBean, bool) {
	if o == nil || IsNil(o.Transport) {
		return nil, false
	}
	return o.Transport, true
}

// HasTransport returns a boolean if a field has been set.
func (o *SyslogServerBean) HasTransport() bool {
	if o != nil && !IsNil(o.Transport) {
		return true
	}

	return false
}

// SetTransport gets a reference to the given SyslogTransportBean and assigns it to the Transport field.
func (o *SyslogServerBean) SetTransport(v SyslogTransportBean) {
	o.Transport = &v
}

// GetIso8601TimestampsEnabled returns the Iso8601TimestampsEnabled field value if set, zero value otherwise.
func (o *SyslogServerBean) GetIso8601TimestampsEnabled() bool {
	if o == nil || IsNil(o.Iso8601TimestampsEnabled) {
		var ret bool
		return ret
	}
	return *o.Iso8601TimestampsEnabled
}

// GetIso8601TimestampsEnabledOk returns a tuple with the Iso8601TimestampsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyslogServerBean) GetIso8601TimestampsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Iso8601TimestampsEnabled) {
		return nil, false
	}
	return o.Iso8601TimestampsEnabled, true
}

// HasIso8601TimestampsEnabled returns a boolean if a field has been set.
func (o *SyslogServerBean) HasIso8601TimestampsEnabled() bool {
	if o != nil && !IsNil(o.Iso8601TimestampsEnabled) {
		return true
	}

	return false
}

// SetIso8601TimestampsEnabled gets a reference to the given bool and assigns it to the Iso8601TimestampsEnabled field.
func (o *SyslogServerBean) SetIso8601TimestampsEnabled(v bool) {
	o.Iso8601TimestampsEnabled = &v
}

// GetServiceTypes returns the ServiceTypes field value if set, zero value otherwise.
func (o *SyslogServerBean) GetServiceTypes() []string {
	if o == nil || IsNil(o.ServiceTypes) {
		var ret []string
		return ret
	}
	return o.ServiceTypes
}

// GetServiceTypesOk returns a tuple with the ServiceTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyslogServerBean) GetServiceTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.ServiceTypes) {
		return nil, false
	}
	return o.ServiceTypes, true
}

// HasServiceTypes returns a boolean if a field has been set.
func (o *SyslogServerBean) HasServiceTypes() bool {
	if o != nil && !IsNil(o.ServiceTypes) {
		return true
	}

	return false
}

// SetServiceTypes gets a reference to the given []string and assigns it to the ServiceTypes field.
func (o *SyslogServerBean) SetServiceTypes(v []string) {
	o.ServiceTypes = v
}

func (o SyslogServerBean) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SyslogServerBean) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !IsNil(o.SyslogProtocol) {
		toSerialize["syslogProtocol"] = o.SyslogProtocol
	}
	if !IsNil(o.SeverityLevel) {
		toSerialize["severityLevel"] = o.SeverityLevel
	}
	if !IsNil(o.Transport) {
		toSerialize["transport"] = o.Transport
	}
	if !IsNil(o.Iso8601TimestampsEnabled) {
		toSerialize["iso8601TimestampsEnabled"] = o.Iso8601TimestampsEnabled
	}
	if !IsNil(o.ServiceTypes) {
		toSerialize["serviceTypes"] = o.ServiceTypes
	}
	return toSerialize, nil
}

type NullableSyslogServerBean struct {
	value *SyslogServerBean
	isSet bool
}

func (v NullableSyslogServerBean) Get() *SyslogServerBean {
	return v.value
}

func (v *NullableSyslogServerBean) Set(val *SyslogServerBean) {
	v.value = val
	v.isSet = true
}

func (v NullableSyslogServerBean) IsSet() bool {
	return v.isSet
}

func (v *NullableSyslogServerBean) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyslogServerBean(val *SyslogServerBean) *NullableSyslogServerBean {
	return &NullableSyslogServerBean{value: val, isSet: true}
}

func (v NullableSyslogServerBean) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyslogServerBean) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


