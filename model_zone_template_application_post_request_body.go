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

// checks if the ZoneTemplateApplicationPostRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ZoneTemplateApplicationPostRequestBody{}

// ZoneTemplateApplicationPostRequestBody struct for ZoneTemplateApplicationPostRequestBody
type ZoneTemplateApplicationPostRequestBody struct {
	ZoneTemplateApplication
}

type _ZoneTemplateApplicationPostRequestBody ZoneTemplateApplicationPostRequestBody

// NewZoneTemplateApplicationPostRequestBody instantiates a new ZoneTemplateApplicationPostRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZoneTemplateApplicationPostRequestBody(type_ string) *ZoneTemplateApplicationPostRequestBody {
	this := ZoneTemplateApplicationPostRequestBody{}
	this.Type = type_
	return &this
}

// NewZoneTemplateApplicationPostRequestBodyWithDefaults instantiates a new ZoneTemplateApplicationPostRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZoneTemplateApplicationPostRequestBodyWithDefaults() *ZoneTemplateApplicationPostRequestBody {
	this := ZoneTemplateApplicationPostRequestBody{}
	return &this
}

func (o ZoneTemplateApplicationPostRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ZoneTemplateApplicationPostRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedZoneTemplateApplication, errZoneTemplateApplication := json.Marshal(o.ZoneTemplateApplication)
	if errZoneTemplateApplication != nil {
		return map[string]interface{}{}, errZoneTemplateApplication
	}
	errZoneTemplateApplication = json.Unmarshal([]byte(serializedZoneTemplateApplication), &toSerialize)
	if errZoneTemplateApplication != nil {
		return map[string]interface{}{}, errZoneTemplateApplication
	}
	return toSerialize, nil
}

func (o *ZoneTemplateApplicationPostRequestBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
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

	varZoneTemplateApplicationPostRequestBody := _ZoneTemplateApplicationPostRequestBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varZoneTemplateApplicationPostRequestBody)

	if err != nil {
		return err
	}

	*o = ZoneTemplateApplicationPostRequestBody(varZoneTemplateApplicationPostRequestBody)

	return err
}

type NullableZoneTemplateApplicationPostRequestBody struct {
	value *ZoneTemplateApplicationPostRequestBody
	isSet bool
}

func (v NullableZoneTemplateApplicationPostRequestBody) Get() *ZoneTemplateApplicationPostRequestBody {
	return v.value
}

func (v *NullableZoneTemplateApplicationPostRequestBody) Set(val *ZoneTemplateApplicationPostRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableZoneTemplateApplicationPostRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableZoneTemplateApplicationPostRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZoneTemplateApplicationPostRequestBody(val *ZoneTemplateApplicationPostRequestBody) *NullableZoneTemplateApplicationPostRequestBody {
	return &NullableZoneTemplateApplicationPostRequestBody{value: val, isSet: true}
}

func (v NullableZoneTemplateApplicationPostRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZoneTemplateApplicationPostRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


