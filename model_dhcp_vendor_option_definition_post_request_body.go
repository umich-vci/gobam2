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

// checks if the DHCPVendorOptionDefinitionPostRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DHCPVendorOptionDefinitionPostRequestBody{}

// DHCPVendorOptionDefinitionPostRequestBody struct for DHCPVendorOptionDefinitionPostRequestBody
type DHCPVendorOptionDefinitionPostRequestBody struct {
	DHCPVendorOptionDefinition
}

type _DHCPVendorOptionDefinitionPostRequestBody DHCPVendorOptionDefinitionPostRequestBody

// NewDHCPVendorOptionDefinitionPostRequestBody instantiates a new DHCPVendorOptionDefinitionPostRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDHCPVendorOptionDefinitionPostRequestBody(type_ string, name string, code int32, displayName string, description string, valueFormat map[string]interface{}) *DHCPVendorOptionDefinitionPostRequestBody {
	this := DHCPVendorOptionDefinitionPostRequestBody{}
	this.Type = type_
	this.Name = name
	this.Code = code
	this.DisplayName = displayName
	this.Description = description
	this.ValueFormat = valueFormat
	return &this
}

// NewDHCPVendorOptionDefinitionPostRequestBodyWithDefaults instantiates a new DHCPVendorOptionDefinitionPostRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDHCPVendorOptionDefinitionPostRequestBodyWithDefaults() *DHCPVendorOptionDefinitionPostRequestBody {
	this := DHCPVendorOptionDefinitionPostRequestBody{}
	return &this
}

func (o DHCPVendorOptionDefinitionPostRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DHCPVendorOptionDefinitionPostRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedDHCPVendorOptionDefinition, errDHCPVendorOptionDefinition := json.Marshal(o.DHCPVendorOptionDefinition)
	if errDHCPVendorOptionDefinition != nil {
		return map[string]interface{}{}, errDHCPVendorOptionDefinition
	}
	errDHCPVendorOptionDefinition = json.Unmarshal([]byte(serializedDHCPVendorOptionDefinition), &toSerialize)
	if errDHCPVendorOptionDefinition != nil {
		return map[string]interface{}{}, errDHCPVendorOptionDefinition
	}
	return toSerialize, nil
}

func (o *DHCPVendorOptionDefinitionPostRequestBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"name",
		"code",
		"displayName",
		"description",
		"valueFormat",
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

	varDHCPVendorOptionDefinitionPostRequestBody := _DHCPVendorOptionDefinitionPostRequestBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDHCPVendorOptionDefinitionPostRequestBody)

	if err != nil {
		return err
	}

	*o = DHCPVendorOptionDefinitionPostRequestBody(varDHCPVendorOptionDefinitionPostRequestBody)

	return err
}

type NullableDHCPVendorOptionDefinitionPostRequestBody struct {
	value *DHCPVendorOptionDefinitionPostRequestBody
	isSet bool
}

func (v NullableDHCPVendorOptionDefinitionPostRequestBody) Get() *DHCPVendorOptionDefinitionPostRequestBody {
	return v.value
}

func (v *NullableDHCPVendorOptionDefinitionPostRequestBody) Set(val *DHCPVendorOptionDefinitionPostRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableDHCPVendorOptionDefinitionPostRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableDHCPVendorOptionDefinitionPostRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDHCPVendorOptionDefinitionPostRequestBody(val *DHCPVendorOptionDefinitionPostRequestBody) *NullableDHCPVendorOptionDefinitionPostRequestBody {
	return &NullableDHCPVendorOptionDefinitionPostRequestBody{value: val, isSet: true}
}

func (v NullableDHCPVendorOptionDefinitionPostRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDHCPVendorOptionDefinitionPostRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


