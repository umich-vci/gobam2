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

// checks if the SnmpTrapServerBean type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SnmpTrapServerBean{}

// SnmpTrapServerBean struct for SnmpTrapServerBean
type SnmpTrapServerBean struct {
	// Indicates whether the SNMP trap server is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// The IP address of the SNMP trap server.
	Address *string `json:"address,omitempty"`
	// The SNMP trap server port.
	Port *int32 `json:"port,omitempty"`
	Snmp *SnmpBean `json:"snmp,omitempty"`
}

// NewSnmpTrapServerBean instantiates a new SnmpTrapServerBean object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnmpTrapServerBean() *SnmpTrapServerBean {
	this := SnmpTrapServerBean{}
	return &this
}

// NewSnmpTrapServerBeanWithDefaults instantiates a new SnmpTrapServerBean object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnmpTrapServerBeanWithDefaults() *SnmpTrapServerBean {
	this := SnmpTrapServerBean{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *SnmpTrapServerBean) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpTrapServerBean) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *SnmpTrapServerBean) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *SnmpTrapServerBean) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *SnmpTrapServerBean) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpTrapServerBean) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *SnmpTrapServerBean) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *SnmpTrapServerBean) SetAddress(v string) {
	o.Address = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *SnmpTrapServerBean) GetPort() int32 {
	if o == nil || IsNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpTrapServerBean) GetPortOk() (*int32, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *SnmpTrapServerBean) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *SnmpTrapServerBean) SetPort(v int32) {
	o.Port = &v
}

// GetSnmp returns the Snmp field value if set, zero value otherwise.
func (o *SnmpTrapServerBean) GetSnmp() SnmpBean {
	if o == nil || IsNil(o.Snmp) {
		var ret SnmpBean
		return ret
	}
	return *o.Snmp
}

// GetSnmpOk returns a tuple with the Snmp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpTrapServerBean) GetSnmpOk() (*SnmpBean, bool) {
	if o == nil || IsNil(o.Snmp) {
		return nil, false
	}
	return o.Snmp, true
}

// HasSnmp returns a boolean if a field has been set.
func (o *SnmpTrapServerBean) HasSnmp() bool {
	if o != nil && !IsNil(o.Snmp) {
		return true
	}

	return false
}

// SetSnmp gets a reference to the given SnmpBean and assigns it to the Snmp field.
func (o *SnmpTrapServerBean) SetSnmp(v SnmpBean) {
	o.Snmp = &v
}

func (o SnmpTrapServerBean) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SnmpTrapServerBean) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !IsNil(o.Snmp) {
		toSerialize["snmp"] = o.Snmp
	}
	return toSerialize, nil
}

type NullableSnmpTrapServerBean struct {
	value *SnmpTrapServerBean
	isSet bool
}

func (v NullableSnmpTrapServerBean) Get() *SnmpTrapServerBean {
	return v.value
}

func (v *NullableSnmpTrapServerBean) Set(val *SnmpTrapServerBean) {
	v.value = val
	v.isSet = true
}

func (v NullableSnmpTrapServerBean) IsSet() bool {
	return v.isSet
}

func (v *NullableSnmpTrapServerBean) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnmpTrapServerBean(val *SnmpTrapServerBean) *NullableSnmpTrapServerBean {
	return &NullableSnmpTrapServerBean{value: val, isSet: true}
}

func (v NullableSnmpTrapServerBean) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnmpTrapServerBean) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


