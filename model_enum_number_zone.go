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

// checks if the ENUMNumberZone type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ENUMNumberZone{}

// ENUMNumberZone struct for ENUMNumberZone
type ENUMNumberZone struct {
	AbstractZone
	// The resource type.
	Type *string `json:"type,omitempty"`
	// The name of the resource.
	Name *string `json:"name,omitempty" validate:"regexp=^.*\\\\S+.*$"`
	NumberName *string `json:"numberName,omitempty" validate:"regexp=^.*\\\\S+.*$"`
	AbsoluteName *string `json:"absoluteName,omitempty"`
}

// NewENUMNumberZone instantiates a new ENUMNumberZone object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewENUMNumberZone() *ENUMNumberZone {
	this := ENUMNumberZone{}
	return &this
}

// NewENUMNumberZoneWithDefaults instantiates a new ENUMNumberZone object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewENUMNumberZoneWithDefaults() *ENUMNumberZone {
	this := ENUMNumberZone{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ENUMNumberZone) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ENUMNumberZone) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ENUMNumberZone) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ENUMNumberZone) SetType(v string) {
	o.Type = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ENUMNumberZone) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ENUMNumberZone) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ENUMNumberZone) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ENUMNumberZone) SetName(v string) {
	o.Name = &v
}

// GetNumberName returns the NumberName field value if set, zero value otherwise.
func (o *ENUMNumberZone) GetNumberName() string {
	if o == nil || IsNil(o.NumberName) {
		var ret string
		return ret
	}
	return *o.NumberName
}

// GetNumberNameOk returns a tuple with the NumberName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ENUMNumberZone) GetNumberNameOk() (*string, bool) {
	if o == nil || IsNil(o.NumberName) {
		return nil, false
	}
	return o.NumberName, true
}

// HasNumberName returns a boolean if a field has been set.
func (o *ENUMNumberZone) HasNumberName() bool {
	if o != nil && !IsNil(o.NumberName) {
		return true
	}

	return false
}

// SetNumberName gets a reference to the given string and assigns it to the NumberName field.
func (o *ENUMNumberZone) SetNumberName(v string) {
	o.NumberName = &v
}

// GetAbsoluteName returns the AbsoluteName field value if set, zero value otherwise.
func (o *ENUMNumberZone) GetAbsoluteName() string {
	if o == nil || IsNil(o.AbsoluteName) {
		var ret string
		return ret
	}
	return *o.AbsoluteName
}

// GetAbsoluteNameOk returns a tuple with the AbsoluteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ENUMNumberZone) GetAbsoluteNameOk() (*string, bool) {
	if o == nil || IsNil(o.AbsoluteName) {
		return nil, false
	}
	return o.AbsoluteName, true
}

// HasAbsoluteName returns a boolean if a field has been set.
func (o *ENUMNumberZone) HasAbsoluteName() bool {
	if o != nil && !IsNil(o.AbsoluteName) {
		return true
	}

	return false
}

// SetAbsoluteName gets a reference to the given string and assigns it to the AbsoluteName field.
func (o *ENUMNumberZone) SetAbsoluteName(v string) {
	o.AbsoluteName = &v
}

func (o ENUMNumberZone) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ENUMNumberZone) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.NumberName) {
		toSerialize["numberName"] = o.NumberName
	}
	if !IsNil(o.AbsoluteName) {
		toSerialize["absoluteName"] = o.AbsoluteName
	}
	return toSerialize, nil
}

type NullableENUMNumberZone struct {
	value *ENUMNumberZone
	isSet bool
}

func (v NullableENUMNumberZone) Get() *ENUMNumberZone {
	return v.value
}

func (v *NullableENUMNumberZone) Set(val *ENUMNumberZone) {
	v.value = val
	v.isSet = true
}

func (v NullableENUMNumberZone) IsSet() bool {
	return v.isSet
}

func (v *NullableENUMNumberZone) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableENUMNumberZone(val *ENUMNumberZone) *NullableENUMNumberZone {
	return &NullableENUMNumberZone{value: val, isSet: true}
}

func (v NullableENUMNumberZone) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableENUMNumberZone) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

