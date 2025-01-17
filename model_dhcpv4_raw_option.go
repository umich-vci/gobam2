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

// checks if the DHCPv4RawOption type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DHCPv4RawOption{}

// DHCPv4RawOption struct for DHCPv4RawOption
type DHCPv4RawOption struct {
	DeploymentOption
	// The resource type.
	Type *string `json:"type,omitempty"`
	// The deployment option value.
	Value *string `json:"value,omitempty" validate:"regexp=^.*\\\\S+.*$"`
}

// NewDHCPv4RawOption instantiates a new DHCPv4RawOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDHCPv4RawOption() *DHCPv4RawOption {
	this := DHCPv4RawOption{}
	return &this
}

// NewDHCPv4RawOptionWithDefaults instantiates a new DHCPv4RawOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDHCPv4RawOptionWithDefaults() *DHCPv4RawOption {
	this := DHCPv4RawOption{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DHCPv4RawOption) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DHCPv4RawOption) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DHCPv4RawOption) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DHCPv4RawOption) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *DHCPv4RawOption) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DHCPv4RawOption) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *DHCPv4RawOption) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *DHCPv4RawOption) SetValue(v string) {
	o.Value = &v
}

func (o DHCPv4RawOption) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DHCPv4RawOption) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableDHCPv4RawOption struct {
	value *DHCPv4RawOption
	isSet bool
}

func (v NullableDHCPv4RawOption) Get() *DHCPv4RawOption {
	return v.value
}

func (v *NullableDHCPv4RawOption) Set(val *DHCPv4RawOption) {
	v.value = val
	v.isSet = true
}

func (v NullableDHCPv4RawOption) IsSet() bool {
	return v.isSet
}

func (v *NullableDHCPv4RawOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDHCPv4RawOption(val *DHCPv4RawOption) *NullableDHCPv4RawOption {
	return &NullableDHCPv4RawOption{value: val, isSet: true}
}

func (v NullableDHCPv4RawOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDHCPv4RawOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


