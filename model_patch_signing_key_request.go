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

// checks if the PatchSigningKeyRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchSigningKeyRequest{}

// PatchSigningKeyRequest struct for PatchSigningKeyRequest
type PatchSigningKeyRequest struct {
	// The state of the key.
	State *string `json:"state,omitempty"`
}

// NewPatchSigningKeyRequest instantiates a new PatchSigningKeyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchSigningKeyRequest() *PatchSigningKeyRequest {
	this := PatchSigningKeyRequest{}
	return &this
}

// NewPatchSigningKeyRequestWithDefaults instantiates a new PatchSigningKeyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchSigningKeyRequestWithDefaults() *PatchSigningKeyRequest {
	this := PatchSigningKeyRequest{}
	return &this
}

// GetState returns the State field value if set, zero value otherwise.
func (o *PatchSigningKeyRequest) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchSigningKeyRequest) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *PatchSigningKeyRequest) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *PatchSigningKeyRequest) SetState(v string) {
	o.State = &v
}

func (o PatchSigningKeyRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchSigningKeyRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	return toSerialize, nil
}

type NullablePatchSigningKeyRequest struct {
	value *PatchSigningKeyRequest
	isSet bool
}

func (v NullablePatchSigningKeyRequest) Get() *PatchSigningKeyRequest {
	return v.value
}

func (v *NullablePatchSigningKeyRequest) Set(val *PatchSigningKeyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchSigningKeyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchSigningKeyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchSigningKeyRequest(val *PatchSigningKeyRequest) *NullablePatchSigningKeyRequest {
	return &NullablePatchSigningKeyRequest{value: val, isSet: true}
}

func (v NullablePatchSigningKeyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchSigningKeyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


