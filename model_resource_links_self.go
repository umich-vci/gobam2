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

// checks if the ResourceLinksSelf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceLinksSelf{}

// ResourceLinksSelf struct for ResourceLinksSelf
type ResourceLinksSelf struct {
	Href *string `json:"href,omitempty"`
}

// NewResourceLinksSelf instantiates a new ResourceLinksSelf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceLinksSelf() *ResourceLinksSelf {
	this := ResourceLinksSelf{}
	return &this
}

// NewResourceLinksSelfWithDefaults instantiates a new ResourceLinksSelf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceLinksSelfWithDefaults() *ResourceLinksSelf {
	this := ResourceLinksSelf{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *ResourceLinksSelf) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceLinksSelf) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *ResourceLinksSelf) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *ResourceLinksSelf) SetHref(v string) {
	o.Href = &v
}

func (o ResourceLinksSelf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceLinksSelf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	return toSerialize, nil
}

type NullableResourceLinksSelf struct {
	value *ResourceLinksSelf
	isSet bool
}

func (v NullableResourceLinksSelf) Get() *ResourceLinksSelf {
	return v.value
}

func (v *NullableResourceLinksSelf) Set(val *ResourceLinksSelf) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceLinksSelf) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceLinksSelf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceLinksSelf(val *ResourceLinksSelf) *NullableResourceLinksSelf {
	return &NullableResourceLinksSelf{value: val, isSet: true}
}

func (v NullableResourceLinksSelf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceLinksSelf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


