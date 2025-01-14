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

// checks if the MonitoringSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MonitoringSettings{}

// MonitoringSettings struct for MonitoringSettings
type MonitoringSettings struct {
	Settings
	// The resource type.
	Type *string `json:"type,omitempty"`
	AddressManagerMonitoring *AddressManagerMonitoringBean `json:"addressManagerMonitoring,omitempty"`
	ServerMonitoring *ServerMonitoringBean `json:"serverMonitoring,omitempty"`
	F5ServerMonitoring *F5ServerMonitoringBean `json:"f5ServerMonitoring,omitempty"`
}

// NewMonitoringSettings instantiates a new MonitoringSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitoringSettings() *MonitoringSettings {
	this := MonitoringSettings{}
	return &this
}

// NewMonitoringSettingsWithDefaults instantiates a new MonitoringSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitoringSettingsWithDefaults() *MonitoringSettings {
	this := MonitoringSettings{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MonitoringSettings) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringSettings) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MonitoringSettings) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *MonitoringSettings) SetType(v string) {
	o.Type = &v
}

// GetAddressManagerMonitoring returns the AddressManagerMonitoring field value if set, zero value otherwise.
func (o *MonitoringSettings) GetAddressManagerMonitoring() AddressManagerMonitoringBean {
	if o == nil || IsNil(o.AddressManagerMonitoring) {
		var ret AddressManagerMonitoringBean
		return ret
	}
	return *o.AddressManagerMonitoring
}

// GetAddressManagerMonitoringOk returns a tuple with the AddressManagerMonitoring field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringSettings) GetAddressManagerMonitoringOk() (*AddressManagerMonitoringBean, bool) {
	if o == nil || IsNil(o.AddressManagerMonitoring) {
		return nil, false
	}
	return o.AddressManagerMonitoring, true
}

// HasAddressManagerMonitoring returns a boolean if a field has been set.
func (o *MonitoringSettings) HasAddressManagerMonitoring() bool {
	if o != nil && !IsNil(o.AddressManagerMonitoring) {
		return true
	}

	return false
}

// SetAddressManagerMonitoring gets a reference to the given AddressManagerMonitoringBean and assigns it to the AddressManagerMonitoring field.
func (o *MonitoringSettings) SetAddressManagerMonitoring(v AddressManagerMonitoringBean) {
	o.AddressManagerMonitoring = &v
}

// GetServerMonitoring returns the ServerMonitoring field value if set, zero value otherwise.
func (o *MonitoringSettings) GetServerMonitoring() ServerMonitoringBean {
	if o == nil || IsNil(o.ServerMonitoring) {
		var ret ServerMonitoringBean
		return ret
	}
	return *o.ServerMonitoring
}

// GetServerMonitoringOk returns a tuple with the ServerMonitoring field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringSettings) GetServerMonitoringOk() (*ServerMonitoringBean, bool) {
	if o == nil || IsNil(o.ServerMonitoring) {
		return nil, false
	}
	return o.ServerMonitoring, true
}

// HasServerMonitoring returns a boolean if a field has been set.
func (o *MonitoringSettings) HasServerMonitoring() bool {
	if o != nil && !IsNil(o.ServerMonitoring) {
		return true
	}

	return false
}

// SetServerMonitoring gets a reference to the given ServerMonitoringBean and assigns it to the ServerMonitoring field.
func (o *MonitoringSettings) SetServerMonitoring(v ServerMonitoringBean) {
	o.ServerMonitoring = &v
}

// GetF5ServerMonitoring returns the F5ServerMonitoring field value if set, zero value otherwise.
func (o *MonitoringSettings) GetF5ServerMonitoring() F5ServerMonitoringBean {
	if o == nil || IsNil(o.F5ServerMonitoring) {
		var ret F5ServerMonitoringBean
		return ret
	}
	return *o.F5ServerMonitoring
}

// GetF5ServerMonitoringOk returns a tuple with the F5ServerMonitoring field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringSettings) GetF5ServerMonitoringOk() (*F5ServerMonitoringBean, bool) {
	if o == nil || IsNil(o.F5ServerMonitoring) {
		return nil, false
	}
	return o.F5ServerMonitoring, true
}

// HasF5ServerMonitoring returns a boolean if a field has been set.
func (o *MonitoringSettings) HasF5ServerMonitoring() bool {
	if o != nil && !IsNil(o.F5ServerMonitoring) {
		return true
	}

	return false
}

// SetF5ServerMonitoring gets a reference to the given F5ServerMonitoringBean and assigns it to the F5ServerMonitoring field.
func (o *MonitoringSettings) SetF5ServerMonitoring(v F5ServerMonitoringBean) {
	o.F5ServerMonitoring = &v
}

func (o MonitoringSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MonitoringSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.AddressManagerMonitoring) {
		toSerialize["addressManagerMonitoring"] = o.AddressManagerMonitoring
	}
	if !IsNil(o.ServerMonitoring) {
		toSerialize["serverMonitoring"] = o.ServerMonitoring
	}
	if !IsNil(o.F5ServerMonitoring) {
		toSerialize["f5ServerMonitoring"] = o.F5ServerMonitoring
	}
	return toSerialize, nil
}

type NullableMonitoringSettings struct {
	value *MonitoringSettings
	isSet bool
}

func (v NullableMonitoringSettings) Get() *MonitoringSettings {
	return v.value
}

func (v *NullableMonitoringSettings) Set(val *MonitoringSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitoringSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitoringSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitoringSettings(val *MonitoringSettings) *NullableMonitoringSettings {
	return &NullableMonitoringSettings{value: val, isSet: true}
}

func (v NullableMonitoringSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitoringSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


