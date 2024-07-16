/*
BlueCat Address Manager 9.5 RESTful v2 API

The **BlueCat Address Manager 9.5 RESTful v2 API** is a new RESTful API for Address Manager that presents Address Manager objects as resources. Each resource has a unique endpoint, and related resources are grouped in collections. To fetch an individual resource, a `GET` request is sent to the resource's endpoint. To fetch all resources in a collection, a `GET` request is sent to the collection's endpoint.  The RESTful v2 API is [hypermedia driven](https://en.wikipedia.org/wiki/HATEOAS) and uses the [HAL](https://en.wikipedia.org/wiki/Hypertext_Application_Language) format for representing links. When navigating through the API, you can use those links to navigate to related resources or child resources of the requested endpoint. The API supports the following media types for most endpoints:  `application/hal+json`, `application/json`, and `text/csv`.  For authentication, the API supports both `Basic` and `Bearer` HTTP authentication schemes.  To get started, create a session and receive credentials for `Basic` authentication by sending a `POST` request to the [/sessions](#/Admin%20Resources/postSession) endpoint, with a message body containing the user's credentials:  ```{\"username\":\"apiuser\", \"password\":\"apipass\"}```  The response will contain an `apiToken` field that can be entered with the username in the Swagger UI **Authorize** dialog. The response will also contain a `basicAuthenticationCredentials` field containing a base64 encoding of the requester's username and API token delimited by a colon. This string can be injected directly into request `Authorization` headers in the following format:  ```Authorization: Basic YXBpOlQ0OExOT3VIRGhDcnVBNEo1bGFES3JuS3hTZC9QK3pjczZXTzBJMDY=```  For full details on API format and authentication methods, refer to the Address Manager RESTful v2 API Guide on the BlueCat Documentation Portal.

API version: 9.5.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gobam2

import (
	"encoding/json"
	"fmt"
)

// IPMoveAllOfMacAddress struct for IPMoveAllOfMacAddress
type IPMoveAllOfMacAddress struct {
	InlinedMACAddress *InlinedMACAddress
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *IPMoveAllOfMacAddress) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into InlinedMACAddress
	err = json.Unmarshal(data, &dst.InlinedMACAddress);
	if err == nil {
		jsonInlinedMACAddress, _ := json.Marshal(dst.InlinedMACAddress)
		if string(jsonInlinedMACAddress) == "{}" { // empty struct
			dst.InlinedMACAddress = nil
		} else {
			return nil // data stored in dst.InlinedMACAddress, return on the first match
		}
	} else {
		dst.InlinedMACAddress = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(IPMoveAllOfMacAddress)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *IPMoveAllOfMacAddress) MarshalJSON() ([]byte, error) {
	if src.InlinedMACAddress != nil {
		return json.Marshal(&src.InlinedMACAddress)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableIPMoveAllOfMacAddress struct {
	value *IPMoveAllOfMacAddress
	isSet bool
}

func (v NullableIPMoveAllOfMacAddress) Get() *IPMoveAllOfMacAddress {
	return v.value
}

func (v *NullableIPMoveAllOfMacAddress) Set(val *IPMoveAllOfMacAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableIPMoveAllOfMacAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableIPMoveAllOfMacAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIPMoveAllOfMacAddress(val *IPMoveAllOfMacAddress) *NullableIPMoveAllOfMacAddress {
	return &NullableIPMoveAllOfMacAddress{value: val, isSet: true}
}

func (v NullableIPMoveAllOfMacAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIPMoveAllOfMacAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

