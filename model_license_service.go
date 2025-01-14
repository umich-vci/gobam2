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

// checks if the LicenseService type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LicenseService{}

// LicenseService struct for LicenseService
type LicenseService struct {
	Service
	// The resource type.
	Type *string `json:"type,omitempty"`
	ClientId *string `json:"clientId,omitempty" validate:"regexp=^.*\\\\S+.*$"`
	Key *string `json:"key,omitempty" validate:"regexp=^.*\\\\S+.*$"`
}

// NewLicenseService instantiates a new LicenseService object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLicenseService() *LicenseService {
	this := LicenseService{}
	return &this
}

// NewLicenseServiceWithDefaults instantiates a new LicenseService object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLicenseServiceWithDefaults() *LicenseService {
	this := LicenseService{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *LicenseService) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseService) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *LicenseService) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *LicenseService) SetType(v string) {
	o.Type = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *LicenseService) GetClientId() string {
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseService) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *LicenseService) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *LicenseService) SetClientId(v string) {
	o.ClientId = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *LicenseService) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseService) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *LicenseService) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *LicenseService) SetKey(v string) {
	o.Key = &v
}

func (o LicenseService) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LicenseService) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.ClientId) {
		toSerialize["clientId"] = o.ClientId
	}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	return toSerialize, nil
}

type NullableLicenseService struct {
	value *LicenseService
	isSet bool
}

func (v NullableLicenseService) Get() *LicenseService {
	return v.value
}

func (v *NullableLicenseService) Set(val *LicenseService) {
	v.value = val
	v.isSet = true
}

func (v NullableLicenseService) IsSet() bool {
	return v.isSet
}

func (v *NullableLicenseService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicenseService(val *LicenseService) *NullableLicenseService {
	return &NullableLicenseService{value: val, isSet: true}
}

func (v NullableLicenseService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicenseService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


