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

// checks if the RADIUSAuthenticator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RADIUSAuthenticator{}

// RADIUSAuthenticator struct for RADIUSAuthenticator
type RADIUSAuthenticator struct {
	Authenticator
	// The type of authenticator.
	Type *string `json:"type,omitempty"`
	// The fully qualified domain name or IP address for the authenticator.
	Hostname *string `json:"hostname,omitempty" validate:"regexp=^.*\\\\S+.*$"`
	// The value that overrides the timeout setting used for authentication requests sent to the RADIUS server.
	Timeout *string `json:"timeout,omitempty"`
	// The port number used for authenticating users against the RADIUS server.
	Port *int32 `json:"port,omitempty"`
	// The number of times Address Manager attempts to retransmit a failed authentication request sent to the RADIUS server.
	Retries *int32 `json:"retries,omitempty"`
	// The shared secret used to encrypt and decrypt packets between the client and the server.
	SharedSecret *string `json:"sharedSecret,omitempty" validate:"regexp=^.*\\\\S+.*$"`
	// The authentication protocol.
	AuthenticationProtocol *string `json:"authenticationProtocol,omitempty"`
}

// NewRADIUSAuthenticator instantiates a new RADIUSAuthenticator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRADIUSAuthenticator() *RADIUSAuthenticator {
	this := RADIUSAuthenticator{}
	return &this
}

// NewRADIUSAuthenticatorWithDefaults instantiates a new RADIUSAuthenticator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRADIUSAuthenticatorWithDefaults() *RADIUSAuthenticator {
	this := RADIUSAuthenticator{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RADIUSAuthenticator) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RADIUSAuthenticator) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RADIUSAuthenticator) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *RADIUSAuthenticator) SetType(v string) {
	o.Type = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *RADIUSAuthenticator) GetHostname() string {
	if o == nil || IsNil(o.Hostname) {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RADIUSAuthenticator) GetHostnameOk() (*string, bool) {
	if o == nil || IsNil(o.Hostname) {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *RADIUSAuthenticator) HasHostname() bool {
	if o != nil && !IsNil(o.Hostname) {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *RADIUSAuthenticator) SetHostname(v string) {
	o.Hostname = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *RADIUSAuthenticator) GetTimeout() string {
	if o == nil || IsNil(o.Timeout) {
		var ret string
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RADIUSAuthenticator) GetTimeoutOk() (*string, bool) {
	if o == nil || IsNil(o.Timeout) {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *RADIUSAuthenticator) HasTimeout() bool {
	if o != nil && !IsNil(o.Timeout) {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given string and assigns it to the Timeout field.
func (o *RADIUSAuthenticator) SetTimeout(v string) {
	o.Timeout = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *RADIUSAuthenticator) GetPort() int32 {
	if o == nil || IsNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RADIUSAuthenticator) GetPortOk() (*int32, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *RADIUSAuthenticator) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *RADIUSAuthenticator) SetPort(v int32) {
	o.Port = &v
}

// GetRetries returns the Retries field value if set, zero value otherwise.
func (o *RADIUSAuthenticator) GetRetries() int32 {
	if o == nil || IsNil(o.Retries) {
		var ret int32
		return ret
	}
	return *o.Retries
}

// GetRetriesOk returns a tuple with the Retries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RADIUSAuthenticator) GetRetriesOk() (*int32, bool) {
	if o == nil || IsNil(o.Retries) {
		return nil, false
	}
	return o.Retries, true
}

// HasRetries returns a boolean if a field has been set.
func (o *RADIUSAuthenticator) HasRetries() bool {
	if o != nil && !IsNil(o.Retries) {
		return true
	}

	return false
}

// SetRetries gets a reference to the given int32 and assigns it to the Retries field.
func (o *RADIUSAuthenticator) SetRetries(v int32) {
	o.Retries = &v
}

// GetSharedSecret returns the SharedSecret field value if set, zero value otherwise.
func (o *RADIUSAuthenticator) GetSharedSecret() string {
	if o == nil || IsNil(o.SharedSecret) {
		var ret string
		return ret
	}
	return *o.SharedSecret
}

// GetSharedSecretOk returns a tuple with the SharedSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RADIUSAuthenticator) GetSharedSecretOk() (*string, bool) {
	if o == nil || IsNil(o.SharedSecret) {
		return nil, false
	}
	return o.SharedSecret, true
}

// HasSharedSecret returns a boolean if a field has been set.
func (o *RADIUSAuthenticator) HasSharedSecret() bool {
	if o != nil && !IsNil(o.SharedSecret) {
		return true
	}

	return false
}

// SetSharedSecret gets a reference to the given string and assigns it to the SharedSecret field.
func (o *RADIUSAuthenticator) SetSharedSecret(v string) {
	o.SharedSecret = &v
}

// GetAuthenticationProtocol returns the AuthenticationProtocol field value if set, zero value otherwise.
func (o *RADIUSAuthenticator) GetAuthenticationProtocol() string {
	if o == nil || IsNil(o.AuthenticationProtocol) {
		var ret string
		return ret
	}
	return *o.AuthenticationProtocol
}

// GetAuthenticationProtocolOk returns a tuple with the AuthenticationProtocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RADIUSAuthenticator) GetAuthenticationProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.AuthenticationProtocol) {
		return nil, false
	}
	return o.AuthenticationProtocol, true
}

// HasAuthenticationProtocol returns a boolean if a field has been set.
func (o *RADIUSAuthenticator) HasAuthenticationProtocol() bool {
	if o != nil && !IsNil(o.AuthenticationProtocol) {
		return true
	}

	return false
}

// SetAuthenticationProtocol gets a reference to the given string and assigns it to the AuthenticationProtocol field.
func (o *RADIUSAuthenticator) SetAuthenticationProtocol(v string) {
	o.AuthenticationProtocol = &v
}

func (o RADIUSAuthenticator) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RADIUSAuthenticator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Hostname) {
		toSerialize["hostname"] = o.Hostname
	}
	if !IsNil(o.Timeout) {
		toSerialize["timeout"] = o.Timeout
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !IsNil(o.Retries) {
		toSerialize["retries"] = o.Retries
	}
	if !IsNil(o.SharedSecret) {
		toSerialize["sharedSecret"] = o.SharedSecret
	}
	if !IsNil(o.AuthenticationProtocol) {
		toSerialize["authenticationProtocol"] = o.AuthenticationProtocol
	}
	return toSerialize, nil
}

type NullableRADIUSAuthenticator struct {
	value *RADIUSAuthenticator
	isSet bool
}

func (v NullableRADIUSAuthenticator) Get() *RADIUSAuthenticator {
	return v.value
}

func (v *NullableRADIUSAuthenticator) Set(val *RADIUSAuthenticator) {
	v.value = val
	v.isSet = true
}

func (v NullableRADIUSAuthenticator) IsSet() bool {
	return v.isSet
}

func (v *NullableRADIUSAuthenticator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRADIUSAuthenticator(val *RADIUSAuthenticator) *NullableRADIUSAuthenticator {
	return &NullableRADIUSAuthenticator{value: val, isSet: true}
}

func (v NullableRADIUSAuthenticator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRADIUSAuthenticator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


