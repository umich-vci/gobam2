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

// PostCollectionRestrictedZoneRequest - struct for PostCollectionRestrictedZoneRequest
type PostCollectionRestrictedZoneRequest struct {
	LinkedViewPostRequestBody *LinkedViewPostRequestBody
	LinkedZonePostRequestBody *LinkedZonePostRequestBody
}

// LinkedViewPostRequestBodyAsPostCollectionRestrictedZoneRequest is a convenience function that returns LinkedViewPostRequestBody wrapped in PostCollectionRestrictedZoneRequest
func LinkedViewPostRequestBodyAsPostCollectionRestrictedZoneRequest(v *LinkedViewPostRequestBody) PostCollectionRestrictedZoneRequest {
	return PostCollectionRestrictedZoneRequest{
		LinkedViewPostRequestBody: v,
	}
}

// LinkedZonePostRequestBodyAsPostCollectionRestrictedZoneRequest is a convenience function that returns LinkedZonePostRequestBody wrapped in PostCollectionRestrictedZoneRequest
func LinkedZonePostRequestBodyAsPostCollectionRestrictedZoneRequest(v *LinkedZonePostRequestBody) PostCollectionRestrictedZoneRequest {
	return PostCollectionRestrictedZoneRequest{
		LinkedZonePostRequestBody: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *PostCollectionRestrictedZoneRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into LinkedViewPostRequestBody
	err = newStrictDecoder(data).Decode(&dst.LinkedViewPostRequestBody)
	if err == nil {
		jsonLinkedViewPostRequestBody, _ := json.Marshal(dst.LinkedViewPostRequestBody)
		if string(jsonLinkedViewPostRequestBody) == "{}" { // empty struct
			dst.LinkedViewPostRequestBody = nil
		} else {
			if err = validator.Validate(dst.LinkedViewPostRequestBody); err != nil {
				dst.LinkedViewPostRequestBody = nil
			} else {
				match++
			}
		}
	} else {
		dst.LinkedViewPostRequestBody = nil
	}

	// try to unmarshal data into LinkedZonePostRequestBody
	err = newStrictDecoder(data).Decode(&dst.LinkedZonePostRequestBody)
	if err == nil {
		jsonLinkedZonePostRequestBody, _ := json.Marshal(dst.LinkedZonePostRequestBody)
		if string(jsonLinkedZonePostRequestBody) == "{}" { // empty struct
			dst.LinkedZonePostRequestBody = nil
		} else {
			if err = validator.Validate(dst.LinkedZonePostRequestBody); err != nil {
				dst.LinkedZonePostRequestBody = nil
			} else {
				match++
			}
		}
	} else {
		dst.LinkedZonePostRequestBody = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.LinkedViewPostRequestBody = nil
		dst.LinkedZonePostRequestBody = nil

		return fmt.Errorf("data matches more than one schema in oneOf(PostCollectionRestrictedZoneRequest)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(PostCollectionRestrictedZoneRequest)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PostCollectionRestrictedZoneRequest) MarshalJSON() ([]byte, error) {
	if src.LinkedViewPostRequestBody != nil {
		return json.Marshal(&src.LinkedViewPostRequestBody)
	}

	if src.LinkedZonePostRequestBody != nil {
		return json.Marshal(&src.LinkedZonePostRequestBody)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PostCollectionRestrictedZoneRequest) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.LinkedViewPostRequestBody != nil {
		return obj.LinkedViewPostRequestBody
	}

	if obj.LinkedZonePostRequestBody != nil {
		return obj.LinkedZonePostRequestBody
	}

	// all schemas are nil
	return nil
}

type NullablePostCollectionRestrictedZoneRequest struct {
	value *PostCollectionRestrictedZoneRequest
	isSet bool
}

func (v NullablePostCollectionRestrictedZoneRequest) Get() *PostCollectionRestrictedZoneRequest {
	return v.value
}

func (v *NullablePostCollectionRestrictedZoneRequest) Set(val *PostCollectionRestrictedZoneRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostCollectionRestrictedZoneRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostCollectionRestrictedZoneRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostCollectionRestrictedZoneRequest(val *PostCollectionRestrictedZoneRequest) *NullablePostCollectionRestrictedZoneRequest {
	return &NullablePostCollectionRestrictedZoneRequest{value: val, isSet: true}
}

func (v NullablePostCollectionRestrictedZoneRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostCollectionRestrictedZoneRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

