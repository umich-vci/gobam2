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

// checks if the SSHService type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SSHService{}

// SSHService struct for SSHService
type SSHService struct {
	Service
	// The resource type.
	Type *string `json:"type,omitempty"`
	// Indicates whether the SSH service is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	TacacsPlusAuthentication *TacacsPlusAuthenticationBean `json:"tacacsPlusAuthentication,omitempty"`
}

// NewSSHService instantiates a new SSHService object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSSHService() *SSHService {
	this := SSHService{}
	return &this
}

// NewSSHServiceWithDefaults instantiates a new SSHService object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSSHServiceWithDefaults() *SSHService {
	this := SSHService{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SSHService) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSHService) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SSHService) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SSHService) SetType(v string) {
	o.Type = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *SSHService) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSHService) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *SSHService) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *SSHService) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetTacacsPlusAuthentication returns the TacacsPlusAuthentication field value if set, zero value otherwise.
func (o *SSHService) GetTacacsPlusAuthentication() TacacsPlusAuthenticationBean {
	if o == nil || IsNil(o.TacacsPlusAuthentication) {
		var ret TacacsPlusAuthenticationBean
		return ret
	}
	return *o.TacacsPlusAuthentication
}

// GetTacacsPlusAuthenticationOk returns a tuple with the TacacsPlusAuthentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSHService) GetTacacsPlusAuthenticationOk() (*TacacsPlusAuthenticationBean, bool) {
	if o == nil || IsNil(o.TacacsPlusAuthentication) {
		return nil, false
	}
	return o.TacacsPlusAuthentication, true
}

// HasTacacsPlusAuthentication returns a boolean if a field has been set.
func (o *SSHService) HasTacacsPlusAuthentication() bool {
	if o != nil && !IsNil(o.TacacsPlusAuthentication) {
		return true
	}

	return false
}

// SetTacacsPlusAuthentication gets a reference to the given TacacsPlusAuthenticationBean and assigns it to the TacacsPlusAuthentication field.
func (o *SSHService) SetTacacsPlusAuthentication(v TacacsPlusAuthenticationBean) {
	o.TacacsPlusAuthentication = &v
}

func (o SSHService) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SSHService) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.TacacsPlusAuthentication) {
		toSerialize["tacacsPlusAuthentication"] = o.TacacsPlusAuthentication
	}
	return toSerialize, nil
}

type NullableSSHService struct {
	value *SSHService
	isSet bool
}

func (v NullableSSHService) Get() *SSHService {
	return v.value
}

func (v *NullableSSHService) Set(val *SSHService) {
	v.value = val
	v.isSet = true
}

func (v NullableSSHService) IsSet() bool {
	return v.isSet
}

func (v *NullableSSHService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSSHService(val *SSHService) *NullableSSHService {
	return &NullableSSHService{value: val, isSet: true}
}

func (v NullableSSHService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSSHService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

