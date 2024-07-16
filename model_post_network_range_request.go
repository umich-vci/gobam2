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

// PostNetworkRangeRequest - struct for PostNetworkRangeRequest
type PostNetworkRangeRequest struct {
	IPv4DHCPRangePostRequestBody *IPv4DHCPRangePostRequestBody
	IPv6DHCPRangePostRequestBody *IPv6DHCPRangePostRequestBody
}

// IPv4DHCPRangePostRequestBodyAsPostNetworkRangeRequest is a convenience function that returns IPv4DHCPRangePostRequestBody wrapped in PostNetworkRangeRequest
func IPv4DHCPRangePostRequestBodyAsPostNetworkRangeRequest(v *IPv4DHCPRangePostRequestBody) PostNetworkRangeRequest {
	return PostNetworkRangeRequest{
		IPv4DHCPRangePostRequestBody: v,
	}
}

// IPv6DHCPRangePostRequestBodyAsPostNetworkRangeRequest is a convenience function that returns IPv6DHCPRangePostRequestBody wrapped in PostNetworkRangeRequest
func IPv6DHCPRangePostRequestBodyAsPostNetworkRangeRequest(v *IPv6DHCPRangePostRequestBody) PostNetworkRangeRequest {
	return PostNetworkRangeRequest{
		IPv6DHCPRangePostRequestBody: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *PostNetworkRangeRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into IPv4DHCPRangePostRequestBody
	err = newStrictDecoder(data).Decode(&dst.IPv4DHCPRangePostRequestBody)
	if err == nil {
		jsonIPv4DHCPRangePostRequestBody, _ := json.Marshal(dst.IPv4DHCPRangePostRequestBody)
		if string(jsonIPv4DHCPRangePostRequestBody) == "{}" { // empty struct
			dst.IPv4DHCPRangePostRequestBody = nil
		} else {
			if err = validator.Validate(dst.IPv4DHCPRangePostRequestBody); err != nil {
				dst.IPv4DHCPRangePostRequestBody = nil
			} else {
				match++
			}
		}
	} else {
		dst.IPv4DHCPRangePostRequestBody = nil
	}

	// try to unmarshal data into IPv6DHCPRangePostRequestBody
	err = newStrictDecoder(data).Decode(&dst.IPv6DHCPRangePostRequestBody)
	if err == nil {
		jsonIPv6DHCPRangePostRequestBody, _ := json.Marshal(dst.IPv6DHCPRangePostRequestBody)
		if string(jsonIPv6DHCPRangePostRequestBody) == "{}" { // empty struct
			dst.IPv6DHCPRangePostRequestBody = nil
		} else {
			if err = validator.Validate(dst.IPv6DHCPRangePostRequestBody); err != nil {
				dst.IPv6DHCPRangePostRequestBody = nil
			} else {
				match++
			}
		}
	} else {
		dst.IPv6DHCPRangePostRequestBody = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.IPv4DHCPRangePostRequestBody = nil
		dst.IPv6DHCPRangePostRequestBody = nil

		return fmt.Errorf("data matches more than one schema in oneOf(PostNetworkRangeRequest)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(PostNetworkRangeRequest)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PostNetworkRangeRequest) MarshalJSON() ([]byte, error) {
	if src.IPv4DHCPRangePostRequestBody != nil {
		return json.Marshal(&src.IPv4DHCPRangePostRequestBody)
	}

	if src.IPv6DHCPRangePostRequestBody != nil {
		return json.Marshal(&src.IPv6DHCPRangePostRequestBody)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PostNetworkRangeRequest) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.IPv4DHCPRangePostRequestBody != nil {
		return obj.IPv4DHCPRangePostRequestBody
	}

	if obj.IPv6DHCPRangePostRequestBody != nil {
		return obj.IPv6DHCPRangePostRequestBody
	}

	// all schemas are nil
	return nil
}

type NullablePostNetworkRangeRequest struct {
	value *PostNetworkRangeRequest
	isSet bool
}

func (v NullablePostNetworkRangeRequest) Get() *PostNetworkRangeRequest {
	return v.value
}

func (v *NullablePostNetworkRangeRequest) Set(val *PostNetworkRangeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostNetworkRangeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostNetworkRangeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostNetworkRangeRequest(val *PostNetworkRangeRequest) *NullablePostNetworkRangeRequest {
	return &NullablePostNetworkRangeRequest{value: val, isSet: true}
}

func (v NullablePostNetworkRangeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostNetworkRangeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


