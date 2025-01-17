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

// checks if the PutMacPool200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PutMacPool200Response{}

// PutMacPool200Response struct for PutMacPool200Response
type PutMacPool200Response struct {
	MACPool
	Links *ResourceLinks `json:"_links,omitempty"`
}

// NewPutMacPool200Response instantiates a new PutMacPool200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutMacPool200Response() *PutMacPool200Response {
	this := PutMacPool200Response{}
	return &this
}

// NewPutMacPool200ResponseWithDefaults instantiates a new PutMacPool200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutMacPool200ResponseWithDefaults() *PutMacPool200Response {
	this := PutMacPool200Response{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PutMacPool200Response) GetLinks() ResourceLinks {
	if o == nil || IsNil(o.Links) {
		var ret ResourceLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutMacPool200Response) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PutMacPool200Response) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ResourceLinks and assigns it to the Links field.
func (o *PutMacPool200Response) SetLinks(v ResourceLinks) {
	o.Links = &v
}

func (o PutMacPool200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PutMacPool200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMACPool, errMACPool := json.Marshal(o.MACPool)
	if errMACPool != nil {
		return map[string]interface{}{}, errMACPool
	}
	errMACPool = json.Unmarshal([]byte(serializedMACPool), &toSerialize)
	if errMACPool != nil {
		return map[string]interface{}{}, errMACPool
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullablePutMacPool200Response struct {
	value *PutMacPool200Response
	isSet bool
}

func (v NullablePutMacPool200Response) Get() *PutMacPool200Response {
	return v.value
}

func (v *NullablePutMacPool200Response) Set(val *PutMacPool200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePutMacPool200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePutMacPool200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutMacPool200Response(val *PutMacPool200Response) *NullablePutMacPool200Response {
	return &NullablePutMacPool200Response{value: val, isSet: true}
}

func (v NullablePutMacPool200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutMacPool200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


