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

// checks if the ContainerPortMappingBean type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContainerPortMappingBean{}

// ContainerPortMappingBean The list of ports to map between the DNS/DHCP Server service host and the Gateway container.
type ContainerPortMappingBean struct {
	// The DNS/DHCP Server port that is mapped to the container port.
	HostPort *int32 `json:"hostPort,omitempty"`
	// The host container port that is mapped to the DNS/DHCP Server port.
	ContainerPort *int32 `json:"containerPort,omitempty"`
}

// NewContainerPortMappingBean instantiates a new ContainerPortMappingBean object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerPortMappingBean() *ContainerPortMappingBean {
	this := ContainerPortMappingBean{}
	return &this
}

// NewContainerPortMappingBeanWithDefaults instantiates a new ContainerPortMappingBean object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerPortMappingBeanWithDefaults() *ContainerPortMappingBean {
	this := ContainerPortMappingBean{}
	return &this
}

// GetHostPort returns the HostPort field value if set, zero value otherwise.
func (o *ContainerPortMappingBean) GetHostPort() int32 {
	if o == nil || IsNil(o.HostPort) {
		var ret int32
		return ret
	}
	return *o.HostPort
}

// GetHostPortOk returns a tuple with the HostPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerPortMappingBean) GetHostPortOk() (*int32, bool) {
	if o == nil || IsNil(o.HostPort) {
		return nil, false
	}
	return o.HostPort, true
}

// HasHostPort returns a boolean if a field has been set.
func (o *ContainerPortMappingBean) HasHostPort() bool {
	if o != nil && !IsNil(o.HostPort) {
		return true
	}

	return false
}

// SetHostPort gets a reference to the given int32 and assigns it to the HostPort field.
func (o *ContainerPortMappingBean) SetHostPort(v int32) {
	o.HostPort = &v
}

// GetContainerPort returns the ContainerPort field value if set, zero value otherwise.
func (o *ContainerPortMappingBean) GetContainerPort() int32 {
	if o == nil || IsNil(o.ContainerPort) {
		var ret int32
		return ret
	}
	return *o.ContainerPort
}

// GetContainerPortOk returns a tuple with the ContainerPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerPortMappingBean) GetContainerPortOk() (*int32, bool) {
	if o == nil || IsNil(o.ContainerPort) {
		return nil, false
	}
	return o.ContainerPort, true
}

// HasContainerPort returns a boolean if a field has been set.
func (o *ContainerPortMappingBean) HasContainerPort() bool {
	if o != nil && !IsNil(o.ContainerPort) {
		return true
	}

	return false
}

// SetContainerPort gets a reference to the given int32 and assigns it to the ContainerPort field.
func (o *ContainerPortMappingBean) SetContainerPort(v int32) {
	o.ContainerPort = &v
}

func (o ContainerPortMappingBean) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContainerPortMappingBean) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HostPort) {
		toSerialize["hostPort"] = o.HostPort
	}
	if !IsNil(o.ContainerPort) {
		toSerialize["containerPort"] = o.ContainerPort
	}
	return toSerialize, nil
}

type NullableContainerPortMappingBean struct {
	value *ContainerPortMappingBean
	isSet bool
}

func (v NullableContainerPortMappingBean) Get() *ContainerPortMappingBean {
	return v.value
}

func (v *NullableContainerPortMappingBean) Set(val *ContainerPortMappingBean) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerPortMappingBean) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerPortMappingBean) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerPortMappingBean(val *ContainerPortMappingBean) *NullableContainerPortMappingBean {
	return &NullableContainerPortMappingBean{value: val, isSet: true}
}

func (v NullableContainerPortMappingBean) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerPortMappingBean) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


