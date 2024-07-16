/*
BlueCat Address Manager 9.5 RESTful v2 API

The **BlueCat Address Manager 9.5 RESTful v2 API** is a new RESTful API for Address Manager that presents Address Manager objects as resources. Each resource has a unique endpoint, and related resources are grouped in collections. To fetch an individual resource, a `GET` request is sent to the resource's endpoint. To fetch all resources in a collection, a `GET` request is sent to the collection's endpoint.  The RESTful v2 API is [hypermedia driven](https://en.wikipedia.org/wiki/HATEOAS) and uses the [HAL](https://en.wikipedia.org/wiki/Hypertext_Application_Language) format for representing links. When navigating through the API, you can use those links to navigate to related resources or child resources of the requested endpoint. The API supports the following media types for most endpoints:  `application/hal+json`, `application/json`, and `text/csv`.  For authentication, the API supports both `Basic` and `Bearer` HTTP authentication schemes.  To get started, create a session and receive credentials for `Basic` authentication by sending a `POST` request to the [/sessions](#/Admin%20Resources/postSession) endpoint, with a message body containing the user's credentials:  ```{\"username\":\"apiuser\", \"password\":\"apipass\"}```  The response will contain an `apiToken` field that can be entered with the username in the Swagger UI **Authorize** dialog. The response will also contain a `basicAuthenticationCredentials` field containing a base64 encoding of the requester's username and API token delimited by a colon. This string can be injected directly into request `Authorization` headers in the following format:  ```Authorization: Basic YXBpOlQ0OExOT3VIRGhDcnVBNEo1bGFES3JuS3hTZC9QK3pjczZXTzBJMDY=```  For full details on API format and authentication methods, refer to the Address Manager RESTful v2 API Guide on the BlueCat Documentation Portal.

API version: 9.5.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gobam2

import (
	"encoding/json"
	"gopkg.in/validator.v2"
	"fmt"
)

// PutRangeRequest - struct for PutRangeRequest
type PutRangeRequest struct {
	IPv4DHCPRangePutRequestBody *IPv4DHCPRangePutRequestBody
	IPv6DHCPRangePutRequestBody *IPv6DHCPRangePutRequestBody
}

// IPv4DHCPRangePutRequestBodyAsPutRangeRequest is a convenience function that returns IPv4DHCPRangePutRequestBody wrapped in PutRangeRequest
func IPv4DHCPRangePutRequestBodyAsPutRangeRequest(v *IPv4DHCPRangePutRequestBody) PutRangeRequest {
	return PutRangeRequest{
		IPv4DHCPRangePutRequestBody: v,
	}
}

// IPv6DHCPRangePutRequestBodyAsPutRangeRequest is a convenience function that returns IPv6DHCPRangePutRequestBody wrapped in PutRangeRequest
func IPv6DHCPRangePutRequestBodyAsPutRangeRequest(v *IPv6DHCPRangePutRequestBody) PutRangeRequest {
	return PutRangeRequest{
		IPv6DHCPRangePutRequestBody: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *PutRangeRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into IPv4DHCPRangePutRequestBody
	err = newStrictDecoder(data).Decode(&dst.IPv4DHCPRangePutRequestBody)
	if err == nil {
		jsonIPv4DHCPRangePutRequestBody, _ := json.Marshal(dst.IPv4DHCPRangePutRequestBody)
		if string(jsonIPv4DHCPRangePutRequestBody) == "{}" { // empty struct
			dst.IPv4DHCPRangePutRequestBody = nil
		} else {
			if err = validator.Validate(dst.IPv4DHCPRangePutRequestBody); err != nil {
				dst.IPv4DHCPRangePutRequestBody = nil
			} else {
				match++
			}
		}
	} else {
		dst.IPv4DHCPRangePutRequestBody = nil
	}

	// try to unmarshal data into IPv6DHCPRangePutRequestBody
	err = newStrictDecoder(data).Decode(&dst.IPv6DHCPRangePutRequestBody)
	if err == nil {
		jsonIPv6DHCPRangePutRequestBody, _ := json.Marshal(dst.IPv6DHCPRangePutRequestBody)
		if string(jsonIPv6DHCPRangePutRequestBody) == "{}" { // empty struct
			dst.IPv6DHCPRangePutRequestBody = nil
		} else {
			if err = validator.Validate(dst.IPv6DHCPRangePutRequestBody); err != nil {
				dst.IPv6DHCPRangePutRequestBody = nil
			} else {
				match++
			}
		}
	} else {
		dst.IPv6DHCPRangePutRequestBody = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.IPv4DHCPRangePutRequestBody = nil
		dst.IPv6DHCPRangePutRequestBody = nil

		return fmt.Errorf("data matches more than one schema in oneOf(PutRangeRequest)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(PutRangeRequest)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PutRangeRequest) MarshalJSON() ([]byte, error) {
	if src.IPv4DHCPRangePutRequestBody != nil {
		return json.Marshal(&src.IPv4DHCPRangePutRequestBody)
	}

	if src.IPv6DHCPRangePutRequestBody != nil {
		return json.Marshal(&src.IPv6DHCPRangePutRequestBody)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PutRangeRequest) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.IPv4DHCPRangePutRequestBody != nil {
		return obj.IPv4DHCPRangePutRequestBody
	}

	if obj.IPv6DHCPRangePutRequestBody != nil {
		return obj.IPv6DHCPRangePutRequestBody
	}

	// all schemas are nil
	return nil
}

type NullablePutRangeRequest struct {
	value *PutRangeRequest
	isSet bool
}

func (v NullablePutRangeRequest) Get() *PutRangeRequest {
	return v.value
}

func (v *NullablePutRangeRequest) Set(val *PutRangeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePutRangeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePutRangeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutRangeRequest(val *PutRangeRequest) *NullablePutRangeRequest {
	return &NullablePutRangeRequest{value: val, isSet: true}
}

func (v NullablePutRangeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutRangeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

