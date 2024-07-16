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

// checks if the TcpSyslogTransportBean type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TcpSyslogTransportBean{}

// TcpSyslogTransportBean struct for TcpSyslogTransportBean
type TcpSyslogTransportBean struct {
	SyslogTransportBean
}

// NewTcpSyslogTransportBean instantiates a new TcpSyslogTransportBean object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTcpSyslogTransportBean() *TcpSyslogTransportBean {
	this := TcpSyslogTransportBean{}
	return &this
}

// NewTcpSyslogTransportBeanWithDefaults instantiates a new TcpSyslogTransportBean object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTcpSyslogTransportBeanWithDefaults() *TcpSyslogTransportBean {
	this := TcpSyslogTransportBean{}
	return &this
}

func (o TcpSyslogTransportBean) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TcpSyslogTransportBean) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

type NullableTcpSyslogTransportBean struct {
	value *TcpSyslogTransportBean
	isSet bool
}

func (v NullableTcpSyslogTransportBean) Get() *TcpSyslogTransportBean {
	return v.value
}

func (v *NullableTcpSyslogTransportBean) Set(val *TcpSyslogTransportBean) {
	v.value = val
	v.isSet = true
}

func (v NullableTcpSyslogTransportBean) IsSet() bool {
	return v.isSet
}

func (v *NullableTcpSyslogTransportBean) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTcpSyslogTransportBean(val *TcpSyslogTransportBean) *NullableTcpSyslogTransportBean {
	return &NullableTcpSyslogTransportBean{value: val, isSet: true}
}

func (v NullableTcpSyslogTransportBean) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTcpSyslogTransportBean) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


