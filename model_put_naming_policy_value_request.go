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

// PutNamingPolicyValueRequest - struct for PutNamingPolicyValueRequest
type PutNamingPolicyValueRequest struct {
	NamingPolicyConnectorValuePutRequestBody *NamingPolicyConnectorValuePutRequestBody
	NamingPolicyIncrementalValuePutRequestBody *NamingPolicyIncrementalValuePutRequestBody
	NamingPolicyIntegerValuePutRequestBody *NamingPolicyIntegerValuePutRequestBody
	NamingPolicyRequiredValuePutRequestBody *NamingPolicyRequiredValuePutRequestBody
	NamingPolicyTextValuePutRequestBody *NamingPolicyTextValuePutRequestBody
}

// NamingPolicyConnectorValuePutRequestBodyAsPutNamingPolicyValueRequest is a convenience function that returns NamingPolicyConnectorValuePutRequestBody wrapped in PutNamingPolicyValueRequest
func NamingPolicyConnectorValuePutRequestBodyAsPutNamingPolicyValueRequest(v *NamingPolicyConnectorValuePutRequestBody) PutNamingPolicyValueRequest {
	return PutNamingPolicyValueRequest{
		NamingPolicyConnectorValuePutRequestBody: v,
	}
}

// NamingPolicyIncrementalValuePutRequestBodyAsPutNamingPolicyValueRequest is a convenience function that returns NamingPolicyIncrementalValuePutRequestBody wrapped in PutNamingPolicyValueRequest
func NamingPolicyIncrementalValuePutRequestBodyAsPutNamingPolicyValueRequest(v *NamingPolicyIncrementalValuePutRequestBody) PutNamingPolicyValueRequest {
	return PutNamingPolicyValueRequest{
		NamingPolicyIncrementalValuePutRequestBody: v,
	}
}

// NamingPolicyIntegerValuePutRequestBodyAsPutNamingPolicyValueRequest is a convenience function that returns NamingPolicyIntegerValuePutRequestBody wrapped in PutNamingPolicyValueRequest
func NamingPolicyIntegerValuePutRequestBodyAsPutNamingPolicyValueRequest(v *NamingPolicyIntegerValuePutRequestBody) PutNamingPolicyValueRequest {
	return PutNamingPolicyValueRequest{
		NamingPolicyIntegerValuePutRequestBody: v,
	}
}

// NamingPolicyRequiredValuePutRequestBodyAsPutNamingPolicyValueRequest is a convenience function that returns NamingPolicyRequiredValuePutRequestBody wrapped in PutNamingPolicyValueRequest
func NamingPolicyRequiredValuePutRequestBodyAsPutNamingPolicyValueRequest(v *NamingPolicyRequiredValuePutRequestBody) PutNamingPolicyValueRequest {
	return PutNamingPolicyValueRequest{
		NamingPolicyRequiredValuePutRequestBody: v,
	}
}

// NamingPolicyTextValuePutRequestBodyAsPutNamingPolicyValueRequest is a convenience function that returns NamingPolicyTextValuePutRequestBody wrapped in PutNamingPolicyValueRequest
func NamingPolicyTextValuePutRequestBodyAsPutNamingPolicyValueRequest(v *NamingPolicyTextValuePutRequestBody) PutNamingPolicyValueRequest {
	return PutNamingPolicyValueRequest{
		NamingPolicyTextValuePutRequestBody: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *PutNamingPolicyValueRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into NamingPolicyConnectorValuePutRequestBody
	err = newStrictDecoder(data).Decode(&dst.NamingPolicyConnectorValuePutRequestBody)
	if err == nil {
		jsonNamingPolicyConnectorValuePutRequestBody, _ := json.Marshal(dst.NamingPolicyConnectorValuePutRequestBody)
		if string(jsonNamingPolicyConnectorValuePutRequestBody) == "{}" { // empty struct
			dst.NamingPolicyConnectorValuePutRequestBody = nil
		} else {
			if err = validator.Validate(dst.NamingPolicyConnectorValuePutRequestBody); err != nil {
				dst.NamingPolicyConnectorValuePutRequestBody = nil
			} else {
				match++
			}
		}
	} else {
		dst.NamingPolicyConnectorValuePutRequestBody = nil
	}

	// try to unmarshal data into NamingPolicyIncrementalValuePutRequestBody
	err = newStrictDecoder(data).Decode(&dst.NamingPolicyIncrementalValuePutRequestBody)
	if err == nil {
		jsonNamingPolicyIncrementalValuePutRequestBody, _ := json.Marshal(dst.NamingPolicyIncrementalValuePutRequestBody)
		if string(jsonNamingPolicyIncrementalValuePutRequestBody) == "{}" { // empty struct
			dst.NamingPolicyIncrementalValuePutRequestBody = nil
		} else {
			if err = validator.Validate(dst.NamingPolicyIncrementalValuePutRequestBody); err != nil {
				dst.NamingPolicyIncrementalValuePutRequestBody = nil
			} else {
				match++
			}
		}
	} else {
		dst.NamingPolicyIncrementalValuePutRequestBody = nil
	}

	// try to unmarshal data into NamingPolicyIntegerValuePutRequestBody
	err = newStrictDecoder(data).Decode(&dst.NamingPolicyIntegerValuePutRequestBody)
	if err == nil {
		jsonNamingPolicyIntegerValuePutRequestBody, _ := json.Marshal(dst.NamingPolicyIntegerValuePutRequestBody)
		if string(jsonNamingPolicyIntegerValuePutRequestBody) == "{}" { // empty struct
			dst.NamingPolicyIntegerValuePutRequestBody = nil
		} else {
			if err = validator.Validate(dst.NamingPolicyIntegerValuePutRequestBody); err != nil {
				dst.NamingPolicyIntegerValuePutRequestBody = nil
			} else {
				match++
			}
		}
	} else {
		dst.NamingPolicyIntegerValuePutRequestBody = nil
	}

	// try to unmarshal data into NamingPolicyRequiredValuePutRequestBody
	err = newStrictDecoder(data).Decode(&dst.NamingPolicyRequiredValuePutRequestBody)
	if err == nil {
		jsonNamingPolicyRequiredValuePutRequestBody, _ := json.Marshal(dst.NamingPolicyRequiredValuePutRequestBody)
		if string(jsonNamingPolicyRequiredValuePutRequestBody) == "{}" { // empty struct
			dst.NamingPolicyRequiredValuePutRequestBody = nil
		} else {
			if err = validator.Validate(dst.NamingPolicyRequiredValuePutRequestBody); err != nil {
				dst.NamingPolicyRequiredValuePutRequestBody = nil
			} else {
				match++
			}
		}
	} else {
		dst.NamingPolicyRequiredValuePutRequestBody = nil
	}

	// try to unmarshal data into NamingPolicyTextValuePutRequestBody
	err = newStrictDecoder(data).Decode(&dst.NamingPolicyTextValuePutRequestBody)
	if err == nil {
		jsonNamingPolicyTextValuePutRequestBody, _ := json.Marshal(dst.NamingPolicyTextValuePutRequestBody)
		if string(jsonNamingPolicyTextValuePutRequestBody) == "{}" { // empty struct
			dst.NamingPolicyTextValuePutRequestBody = nil
		} else {
			if err = validator.Validate(dst.NamingPolicyTextValuePutRequestBody); err != nil {
				dst.NamingPolicyTextValuePutRequestBody = nil
			} else {
				match++
			}
		}
	} else {
		dst.NamingPolicyTextValuePutRequestBody = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.NamingPolicyConnectorValuePutRequestBody = nil
		dst.NamingPolicyIncrementalValuePutRequestBody = nil
		dst.NamingPolicyIntegerValuePutRequestBody = nil
		dst.NamingPolicyRequiredValuePutRequestBody = nil
		dst.NamingPolicyTextValuePutRequestBody = nil

		return fmt.Errorf("data matches more than one schema in oneOf(PutNamingPolicyValueRequest)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(PutNamingPolicyValueRequest)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PutNamingPolicyValueRequest) MarshalJSON() ([]byte, error) {
	if src.NamingPolicyConnectorValuePutRequestBody != nil {
		return json.Marshal(&src.NamingPolicyConnectorValuePutRequestBody)
	}

	if src.NamingPolicyIncrementalValuePutRequestBody != nil {
		return json.Marshal(&src.NamingPolicyIncrementalValuePutRequestBody)
	}

	if src.NamingPolicyIntegerValuePutRequestBody != nil {
		return json.Marshal(&src.NamingPolicyIntegerValuePutRequestBody)
	}

	if src.NamingPolicyRequiredValuePutRequestBody != nil {
		return json.Marshal(&src.NamingPolicyRequiredValuePutRequestBody)
	}

	if src.NamingPolicyTextValuePutRequestBody != nil {
		return json.Marshal(&src.NamingPolicyTextValuePutRequestBody)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PutNamingPolicyValueRequest) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.NamingPolicyConnectorValuePutRequestBody != nil {
		return obj.NamingPolicyConnectorValuePutRequestBody
	}

	if obj.NamingPolicyIncrementalValuePutRequestBody != nil {
		return obj.NamingPolicyIncrementalValuePutRequestBody
	}

	if obj.NamingPolicyIntegerValuePutRequestBody != nil {
		return obj.NamingPolicyIntegerValuePutRequestBody
	}

	if obj.NamingPolicyRequiredValuePutRequestBody != nil {
		return obj.NamingPolicyRequiredValuePutRequestBody
	}

	if obj.NamingPolicyTextValuePutRequestBody != nil {
		return obj.NamingPolicyTextValuePutRequestBody
	}

	// all schemas are nil
	return nil
}

type NullablePutNamingPolicyValueRequest struct {
	value *PutNamingPolicyValueRequest
	isSet bool
}

func (v NullablePutNamingPolicyValueRequest) Get() *PutNamingPolicyValueRequest {
	return v.value
}

func (v *NullablePutNamingPolicyValueRequest) Set(val *PutNamingPolicyValueRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePutNamingPolicyValueRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePutNamingPolicyValueRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutNamingPolicyValueRequest(val *PutNamingPolicyValueRequest) *NullablePutNamingPolicyValueRequest {
	return &NullablePutNamingPolicyValueRequest{value: val, isSet: true}
}

func (v NullablePutNamingPolicyValueRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutNamingPolicyValueRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


