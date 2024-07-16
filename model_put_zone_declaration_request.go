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

// PutZoneDeclarationRequest - struct for PutZoneDeclarationRequest
type PutZoneDeclarationRequest struct {
	DHCPForwardZonePutRequestBody *DHCPForwardZonePutRequestBody
	DHCPReverseZonePutRequestBody *DHCPReverseZonePutRequestBody
}

// DHCPForwardZonePutRequestBodyAsPutZoneDeclarationRequest is a convenience function that returns DHCPForwardZonePutRequestBody wrapped in PutZoneDeclarationRequest
func DHCPForwardZonePutRequestBodyAsPutZoneDeclarationRequest(v *DHCPForwardZonePutRequestBody) PutZoneDeclarationRequest {
	return PutZoneDeclarationRequest{
		DHCPForwardZonePutRequestBody: v,
	}
}

// DHCPReverseZonePutRequestBodyAsPutZoneDeclarationRequest is a convenience function that returns DHCPReverseZonePutRequestBody wrapped in PutZoneDeclarationRequest
func DHCPReverseZonePutRequestBodyAsPutZoneDeclarationRequest(v *DHCPReverseZonePutRequestBody) PutZoneDeclarationRequest {
	return PutZoneDeclarationRequest{
		DHCPReverseZonePutRequestBody: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *PutZoneDeclarationRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into DHCPForwardZonePutRequestBody
	err = newStrictDecoder(data).Decode(&dst.DHCPForwardZonePutRequestBody)
	if err == nil {
		jsonDHCPForwardZonePutRequestBody, _ := json.Marshal(dst.DHCPForwardZonePutRequestBody)
		if string(jsonDHCPForwardZonePutRequestBody) == "{}" { // empty struct
			dst.DHCPForwardZonePutRequestBody = nil
		} else {
			if err = validator.Validate(dst.DHCPForwardZonePutRequestBody); err != nil {
				dst.DHCPForwardZonePutRequestBody = nil
			} else {
				match++
			}
		}
	} else {
		dst.DHCPForwardZonePutRequestBody = nil
	}

	// try to unmarshal data into DHCPReverseZonePutRequestBody
	err = newStrictDecoder(data).Decode(&dst.DHCPReverseZonePutRequestBody)
	if err == nil {
		jsonDHCPReverseZonePutRequestBody, _ := json.Marshal(dst.DHCPReverseZonePutRequestBody)
		if string(jsonDHCPReverseZonePutRequestBody) == "{}" { // empty struct
			dst.DHCPReverseZonePutRequestBody = nil
		} else {
			if err = validator.Validate(dst.DHCPReverseZonePutRequestBody); err != nil {
				dst.DHCPReverseZonePutRequestBody = nil
			} else {
				match++
			}
		}
	} else {
		dst.DHCPReverseZonePutRequestBody = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.DHCPForwardZonePutRequestBody = nil
		dst.DHCPReverseZonePutRequestBody = nil

		return fmt.Errorf("data matches more than one schema in oneOf(PutZoneDeclarationRequest)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(PutZoneDeclarationRequest)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PutZoneDeclarationRequest) MarshalJSON() ([]byte, error) {
	if src.DHCPForwardZonePutRequestBody != nil {
		return json.Marshal(&src.DHCPForwardZonePutRequestBody)
	}

	if src.DHCPReverseZonePutRequestBody != nil {
		return json.Marshal(&src.DHCPReverseZonePutRequestBody)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PutZoneDeclarationRequest) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.DHCPForwardZonePutRequestBody != nil {
		return obj.DHCPForwardZonePutRequestBody
	}

	if obj.DHCPReverseZonePutRequestBody != nil {
		return obj.DHCPReverseZonePutRequestBody
	}

	// all schemas are nil
	return nil
}

type NullablePutZoneDeclarationRequest struct {
	value *PutZoneDeclarationRequest
	isSet bool
}

func (v NullablePutZoneDeclarationRequest) Get() *PutZoneDeclarationRequest {
	return v.value
}

func (v *NullablePutZoneDeclarationRequest) Set(val *PutZoneDeclarationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePutZoneDeclarationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePutZoneDeclarationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutZoneDeclarationRequest(val *PutZoneDeclarationRequest) *NullablePutZoneDeclarationRequest {
	return &NullablePutZoneDeclarationRequest{value: val, isSet: true}
}

func (v NullablePutZoneDeclarationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutZoneDeclarationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


