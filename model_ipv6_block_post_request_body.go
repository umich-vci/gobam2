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

// checks if the IPv6BlockPostRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IPv6BlockPostRequestBody{}

// IPv6BlockPostRequestBody struct for IPv6BlockPostRequestBody
type IPv6BlockPostRequestBody struct {
	IPv6Block
}

type _IPv6BlockPostRequestBody IPv6BlockPostRequestBody

// NewIPv6BlockPostRequestBody instantiates a new IPv6BlockPostRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIPv6BlockPostRequestBody(type_ string, range_ string) *IPv6BlockPostRequestBody {
	this := IPv6BlockPostRequestBody{}
	this.Type = type_
	this.Range = range_
	return &this
}

// NewIPv6BlockPostRequestBodyWithDefaults instantiates a new IPv6BlockPostRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIPv6BlockPostRequestBodyWithDefaults() *IPv6BlockPostRequestBody {
	this := IPv6BlockPostRequestBody{}
	return &this
}

func (o IPv6BlockPostRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IPv6BlockPostRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedIPv6Block, errIPv6Block := json.Marshal(o.IPv6Block)
	if errIPv6Block != nil {
		return map[string]interface{}{}, errIPv6Block
	}
	errIPv6Block = json.Unmarshal([]byte(serializedIPv6Block), &toSerialize)
	if errIPv6Block != nil {
		return map[string]interface{}{}, errIPv6Block
	}
	return toSerialize, nil
}

func (o *IPv6BlockPostRequestBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"range",
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

	varIPv6BlockPostRequestBody := _IPv6BlockPostRequestBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varIPv6BlockPostRequestBody)

	if err != nil {
		return err
	}

	*o = IPv6BlockPostRequestBody(varIPv6BlockPostRequestBody)

	return err
}

type NullableIPv6BlockPostRequestBody struct {
	value *IPv6BlockPostRequestBody
	isSet bool
}

func (v NullableIPv6BlockPostRequestBody) Get() *IPv6BlockPostRequestBody {
	return v.value
}

func (v *NullableIPv6BlockPostRequestBody) Set(val *IPv6BlockPostRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableIPv6BlockPostRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableIPv6BlockPostRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIPv6BlockPostRequestBody(val *IPv6BlockPostRequestBody) *NullableIPv6BlockPostRequestBody {
	return &NullableIPv6BlockPostRequestBody{value: val, isSet: true}
}

func (v NullableIPv6BlockPostRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIPv6BlockPostRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


