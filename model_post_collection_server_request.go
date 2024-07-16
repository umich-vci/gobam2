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

// PostCollectionServerRequest - struct for PostCollectionServerRequest
type PostCollectionServerRequest struct {
	LinkedHAServerPostRequestBody *LinkedHAServerPostRequestBody
	LinkedServerPostRequestBody *LinkedServerPostRequestBody
}

// LinkedHAServerPostRequestBodyAsPostCollectionServerRequest is a convenience function that returns LinkedHAServerPostRequestBody wrapped in PostCollectionServerRequest
func LinkedHAServerPostRequestBodyAsPostCollectionServerRequest(v *LinkedHAServerPostRequestBody) PostCollectionServerRequest {
	return PostCollectionServerRequest{
		LinkedHAServerPostRequestBody: v,
	}
}

// LinkedServerPostRequestBodyAsPostCollectionServerRequest is a convenience function that returns LinkedServerPostRequestBody wrapped in PostCollectionServerRequest
func LinkedServerPostRequestBodyAsPostCollectionServerRequest(v *LinkedServerPostRequestBody) PostCollectionServerRequest {
	return PostCollectionServerRequest{
		LinkedServerPostRequestBody: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *PostCollectionServerRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into LinkedHAServerPostRequestBody
	err = newStrictDecoder(data).Decode(&dst.LinkedHAServerPostRequestBody)
	if err == nil {
		jsonLinkedHAServerPostRequestBody, _ := json.Marshal(dst.LinkedHAServerPostRequestBody)
		if string(jsonLinkedHAServerPostRequestBody) == "{}" { // empty struct
			dst.LinkedHAServerPostRequestBody = nil
		} else {
			if err = validator.Validate(dst.LinkedHAServerPostRequestBody); err != nil {
				dst.LinkedHAServerPostRequestBody = nil
			} else {
				match++
			}
		}
	} else {
		dst.LinkedHAServerPostRequestBody = nil
	}

	// try to unmarshal data into LinkedServerPostRequestBody
	err = newStrictDecoder(data).Decode(&dst.LinkedServerPostRequestBody)
	if err == nil {
		jsonLinkedServerPostRequestBody, _ := json.Marshal(dst.LinkedServerPostRequestBody)
		if string(jsonLinkedServerPostRequestBody) == "{}" { // empty struct
			dst.LinkedServerPostRequestBody = nil
		} else {
			if err = validator.Validate(dst.LinkedServerPostRequestBody); err != nil {
				dst.LinkedServerPostRequestBody = nil
			} else {
				match++
			}
		}
	} else {
		dst.LinkedServerPostRequestBody = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.LinkedHAServerPostRequestBody = nil
		dst.LinkedServerPostRequestBody = nil

		return fmt.Errorf("data matches more than one schema in oneOf(PostCollectionServerRequest)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(PostCollectionServerRequest)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PostCollectionServerRequest) MarshalJSON() ([]byte, error) {
	if src.LinkedHAServerPostRequestBody != nil {
		return json.Marshal(&src.LinkedHAServerPostRequestBody)
	}

	if src.LinkedServerPostRequestBody != nil {
		return json.Marshal(&src.LinkedServerPostRequestBody)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PostCollectionServerRequest) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.LinkedHAServerPostRequestBody != nil {
		return obj.LinkedHAServerPostRequestBody
	}

	if obj.LinkedServerPostRequestBody != nil {
		return obj.LinkedServerPostRequestBody
	}

	// all schemas are nil
	return nil
}

type NullablePostCollectionServerRequest struct {
	value *PostCollectionServerRequest
	isSet bool
}

func (v NullablePostCollectionServerRequest) Get() *PostCollectionServerRequest {
	return v.value
}

func (v *NullablePostCollectionServerRequest) Set(val *PostCollectionServerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostCollectionServerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostCollectionServerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostCollectionServerRequest(val *PostCollectionServerRequest) *NullablePostCollectionServerRequest {
	return &NullablePostCollectionServerRequest{value: val, isSet: true}
}

func (v NullablePostCollectionServerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostCollectionServerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

