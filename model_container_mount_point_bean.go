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

// checks if the ContainerMountPointBean type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContainerMountPointBean{}

// ContainerMountPointBean The directories to bind mount to the Gateway container.
type ContainerMountPointBean struct {
	// The DNS/DHCP Server directory that is bind mounted to the container.
	Source *string `json:"source,omitempty" validate:"regexp=^.*\\\\S+.*$"`
	// The container directory that is mapped to the DNS/DHCP Server directory.
	ContainerPath *string `json:"containerPath,omitempty" validate:"regexp=^.*\\\\S+.*$"`
}

// NewContainerMountPointBean instantiates a new ContainerMountPointBean object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerMountPointBean() *ContainerMountPointBean {
	this := ContainerMountPointBean{}
	return &this
}

// NewContainerMountPointBeanWithDefaults instantiates a new ContainerMountPointBean object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerMountPointBeanWithDefaults() *ContainerMountPointBean {
	this := ContainerMountPointBean{}
	return &this
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *ContainerMountPointBean) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerMountPointBean) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *ContainerMountPointBean) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *ContainerMountPointBean) SetSource(v string) {
	o.Source = &v
}

// GetContainerPath returns the ContainerPath field value if set, zero value otherwise.
func (o *ContainerMountPointBean) GetContainerPath() string {
	if o == nil || IsNil(o.ContainerPath) {
		var ret string
		return ret
	}
	return *o.ContainerPath
}

// GetContainerPathOk returns a tuple with the ContainerPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerMountPointBean) GetContainerPathOk() (*string, bool) {
	if o == nil || IsNil(o.ContainerPath) {
		return nil, false
	}
	return o.ContainerPath, true
}

// HasContainerPath returns a boolean if a field has been set.
func (o *ContainerMountPointBean) HasContainerPath() bool {
	if o != nil && !IsNil(o.ContainerPath) {
		return true
	}

	return false
}

// SetContainerPath gets a reference to the given string and assigns it to the ContainerPath field.
func (o *ContainerMountPointBean) SetContainerPath(v string) {
	o.ContainerPath = &v
}

func (o ContainerMountPointBean) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContainerMountPointBean) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.ContainerPath) {
		toSerialize["containerPath"] = o.ContainerPath
	}
	return toSerialize, nil
}

type NullableContainerMountPointBean struct {
	value *ContainerMountPointBean
	isSet bool
}

func (v NullableContainerMountPointBean) Get() *ContainerMountPointBean {
	return v.value
}

func (v *NullableContainerMountPointBean) Set(val *ContainerMountPointBean) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerMountPointBean) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerMountPointBean) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerMountPointBean(val *ContainerMountPointBean) *NullableContainerMountPointBean {
	return &NullableContainerMountPointBean{value: val, isSet: true}
}

func (v NullableContainerMountPointBean) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerMountPointBean) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


