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

// checks if the TFTPDeploymentRole type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TFTPDeploymentRole{}

// TFTPDeploymentRole struct for TFTPDeploymentRole
type TFTPDeploymentRole struct {
	DeploymentRole
	// The resource type.
	Type *string `json:"type,omitempty"`
	RoleType *string `json:"roleType,omitempty"`
	Server *InlinedServer `json:"server,omitempty"`
}

// NewTFTPDeploymentRole instantiates a new TFTPDeploymentRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTFTPDeploymentRole() *TFTPDeploymentRole {
	this := TFTPDeploymentRole{}
	return &this
}

// NewTFTPDeploymentRoleWithDefaults instantiates a new TFTPDeploymentRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTFTPDeploymentRoleWithDefaults() *TFTPDeploymentRole {
	this := TFTPDeploymentRole{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TFTPDeploymentRole) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TFTPDeploymentRole) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TFTPDeploymentRole) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TFTPDeploymentRole) SetType(v string) {
	o.Type = &v
}

// GetRoleType returns the RoleType field value if set, zero value otherwise.
func (o *TFTPDeploymentRole) GetRoleType() string {
	if o == nil || IsNil(o.RoleType) {
		var ret string
		return ret
	}
	return *o.RoleType
}

// GetRoleTypeOk returns a tuple with the RoleType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TFTPDeploymentRole) GetRoleTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RoleType) {
		return nil, false
	}
	return o.RoleType, true
}

// HasRoleType returns a boolean if a field has been set.
func (o *TFTPDeploymentRole) HasRoleType() bool {
	if o != nil && !IsNil(o.RoleType) {
		return true
	}

	return false
}

// SetRoleType gets a reference to the given string and assigns it to the RoleType field.
func (o *TFTPDeploymentRole) SetRoleType(v string) {
	o.RoleType = &v
}

// GetServer returns the Server field value if set, zero value otherwise.
func (o *TFTPDeploymentRole) GetServer() InlinedServer {
	if o == nil || IsNil(o.Server) {
		var ret InlinedServer
		return ret
	}
	return *o.Server
}

// GetServerOk returns a tuple with the Server field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TFTPDeploymentRole) GetServerOk() (*InlinedServer, bool) {
	if o == nil || IsNil(o.Server) {
		return nil, false
	}
	return o.Server, true
}

// HasServer returns a boolean if a field has been set.
func (o *TFTPDeploymentRole) HasServer() bool {
	if o != nil && !IsNil(o.Server) {
		return true
	}

	return false
}

// SetServer gets a reference to the given InlinedServer and assigns it to the Server field.
func (o *TFTPDeploymentRole) SetServer(v InlinedServer) {
	o.Server = &v
}

func (o TFTPDeploymentRole) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TFTPDeploymentRole) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.RoleType) {
		toSerialize["roleType"] = o.RoleType
	}
	if !IsNil(o.Server) {
		toSerialize["server"] = o.Server
	}
	return toSerialize, nil
}

type NullableTFTPDeploymentRole struct {
	value *TFTPDeploymentRole
	isSet bool
}

func (v NullableTFTPDeploymentRole) Get() *TFTPDeploymentRole {
	return v.value
}

func (v *NullableTFTPDeploymentRole) Set(val *TFTPDeploymentRole) {
	v.value = val
	v.isSet = true
}

func (v NullableTFTPDeploymentRole) IsSet() bool {
	return v.isSet
}

func (v *NullableTFTPDeploymentRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTFTPDeploymentRole(val *TFTPDeploymentRole) *NullableTFTPDeploymentRole {
	return &NullableTFTPDeploymentRole{value: val, isSet: true}
}

func (v NullableTFTPDeploymentRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTFTPDeploymentRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


