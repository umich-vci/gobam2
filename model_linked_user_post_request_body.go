/*
BlueCat Address Manager 9.5 RESTful v2 API

The **BlueCat Address Manager 9.5 RESTful v2 API** is a new RESTful API for Address Manager that presents Address Manager objects as resources. Each resource has a unique endpoint, and related resources are grouped in collections. To fetch an individual resource, a `GET` request is sent to the resource's endpoint. To fetch all resources in a collection, a `GET` request is sent to the collection's endpoint.  The RESTful v2 API is [hypermedia driven](https://en.wikipedia.org/wiki/HATEOAS) and uses the [HAL](https://en.wikipedia.org/wiki/Hypertext_Application_Language) format for representing links. When navigating through the API, you can use those links to navigate to related resources or child resources of the requested endpoint. The API supports the following media types for most endpoints:  `application/hal+json`, `application/json`, and `text/csv`.  For authentication, the API supports both `Basic` and `Bearer` HTTP authentication schemes.  To get started, create a session and receive credentials for `Basic` authentication by sending a `POST` request to the [/sessions](#/Admin%20Resources/postSession) endpoint, with a message body containing the user's credentials:  ```{\"username\":\"apiuser\", \"password\":\"apipass\"}```  The response will contain an `apiToken` field that can be entered with the username in the Swagger UI **Authorize** dialog. The response will also contain a `basicAuthenticationCredentials` field containing a base64 encoding of the requester's username and API token delimited by a colon. This string can be injected directly into request `Authorization` headers in the following format:  ```Authorization: Basic YXBpOlQ0OExOT3VIRGhDcnVBNEo1bGFES3JuS3hTZC9QK3pjczZXTzBJMDY=```  For full details on API format and authentication methods, refer to the Address Manager RESTful v2 API Guide on the BlueCat Documentation Portal.

API version: 9.5.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gobam2

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the LinkedUserPostRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LinkedUserPostRequestBody{}

// LinkedUserPostRequestBody struct for LinkedUserPostRequestBody
type LinkedUserPostRequestBody struct {
	// The resource identifier.
	Id int64 `json:"id"`
	// The resource type.
	Type *string `json:"type,omitempty"`
}

type _LinkedUserPostRequestBody LinkedUserPostRequestBody

// NewLinkedUserPostRequestBody instantiates a new LinkedUserPostRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkedUserPostRequestBody(id int64) *LinkedUserPostRequestBody {
	this := LinkedUserPostRequestBody{}
	this.Id = id
	return &this
}

// NewLinkedUserPostRequestBodyWithDefaults instantiates a new LinkedUserPostRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkedUserPostRequestBodyWithDefaults() *LinkedUserPostRequestBody {
	this := LinkedUserPostRequestBody{}
	return &this
}

// GetId returns the Id field value
func (o *LinkedUserPostRequestBody) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LinkedUserPostRequestBody) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *LinkedUserPostRequestBody) SetId(v int64) {
	o.Id = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *LinkedUserPostRequestBody) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkedUserPostRequestBody) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *LinkedUserPostRequestBody) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *LinkedUserPostRequestBody) SetType(v string) {
	o.Type = &v
}

func (o LinkedUserPostRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LinkedUserPostRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

func (o *LinkedUserPostRequestBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varLinkedUserPostRequestBody := _LinkedUserPostRequestBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLinkedUserPostRequestBody)

	if err != nil {
		return err
	}

	*o = LinkedUserPostRequestBody(varLinkedUserPostRequestBody)

	return err
}

type NullableLinkedUserPostRequestBody struct {
	value *LinkedUserPostRequestBody
	isSet bool
}

func (v NullableLinkedUserPostRequestBody) Get() *LinkedUserPostRequestBody {
	return v.value
}

func (v *NullableLinkedUserPostRequestBody) Set(val *LinkedUserPostRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkedUserPostRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkedUserPostRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkedUserPostRequestBody(val *LinkedUserPostRequestBody) *NullableLinkedUserPostRequestBody {
	return &NullableLinkedUserPostRequestBody{value: val, isSet: true}
}

func (v NullableLinkedUserPostRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkedUserPostRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


