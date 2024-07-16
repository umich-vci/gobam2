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

// PutTemplateItemRequest - struct for PutTemplateItemRequest
type PutTemplateItemRequest struct {
	IPv4DHCPRangeTemplateItemPutRequestBody *IPv4DHCPRangeTemplateItemPutRequestBody
	IPv4GatewayTemplateItemPutRequestBody *IPv4GatewayTemplateItemPutRequestBody
	IPv4GroupTemplateItemPutRequestBody *IPv4GroupTemplateItemPutRequestBody
	IPv4ReservedRangeTemplateItemPutRequestBody *IPv4ReservedRangeTemplateItemPutRequestBody
}

// IPv4DHCPRangeTemplateItemPutRequestBodyAsPutTemplateItemRequest is a convenience function that returns IPv4DHCPRangeTemplateItemPutRequestBody wrapped in PutTemplateItemRequest
func IPv4DHCPRangeTemplateItemPutRequestBodyAsPutTemplateItemRequest(v *IPv4DHCPRangeTemplateItemPutRequestBody) PutTemplateItemRequest {
	return PutTemplateItemRequest{
		IPv4DHCPRangeTemplateItemPutRequestBody: v,
	}
}

// IPv4GatewayTemplateItemPutRequestBodyAsPutTemplateItemRequest is a convenience function that returns IPv4GatewayTemplateItemPutRequestBody wrapped in PutTemplateItemRequest
func IPv4GatewayTemplateItemPutRequestBodyAsPutTemplateItemRequest(v *IPv4GatewayTemplateItemPutRequestBody) PutTemplateItemRequest {
	return PutTemplateItemRequest{
		IPv4GatewayTemplateItemPutRequestBody: v,
	}
}

// IPv4GroupTemplateItemPutRequestBodyAsPutTemplateItemRequest is a convenience function that returns IPv4GroupTemplateItemPutRequestBody wrapped in PutTemplateItemRequest
func IPv4GroupTemplateItemPutRequestBodyAsPutTemplateItemRequest(v *IPv4GroupTemplateItemPutRequestBody) PutTemplateItemRequest {
	return PutTemplateItemRequest{
		IPv4GroupTemplateItemPutRequestBody: v,
	}
}

// IPv4ReservedRangeTemplateItemPutRequestBodyAsPutTemplateItemRequest is a convenience function that returns IPv4ReservedRangeTemplateItemPutRequestBody wrapped in PutTemplateItemRequest
func IPv4ReservedRangeTemplateItemPutRequestBodyAsPutTemplateItemRequest(v *IPv4ReservedRangeTemplateItemPutRequestBody) PutTemplateItemRequest {
	return PutTemplateItemRequest{
		IPv4ReservedRangeTemplateItemPutRequestBody: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *PutTemplateItemRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into IPv4DHCPRangeTemplateItemPutRequestBody
	err = newStrictDecoder(data).Decode(&dst.IPv4DHCPRangeTemplateItemPutRequestBody)
	if err == nil {
		jsonIPv4DHCPRangeTemplateItemPutRequestBody, _ := json.Marshal(dst.IPv4DHCPRangeTemplateItemPutRequestBody)
		if string(jsonIPv4DHCPRangeTemplateItemPutRequestBody) == "{}" { // empty struct
			dst.IPv4DHCPRangeTemplateItemPutRequestBody = nil
		} else {
			if err = validator.Validate(dst.IPv4DHCPRangeTemplateItemPutRequestBody); err != nil {
				dst.IPv4DHCPRangeTemplateItemPutRequestBody = nil
			} else {
				match++
			}
		}
	} else {
		dst.IPv4DHCPRangeTemplateItemPutRequestBody = nil
	}

	// try to unmarshal data into IPv4GatewayTemplateItemPutRequestBody
	err = newStrictDecoder(data).Decode(&dst.IPv4GatewayTemplateItemPutRequestBody)
	if err == nil {
		jsonIPv4GatewayTemplateItemPutRequestBody, _ := json.Marshal(dst.IPv4GatewayTemplateItemPutRequestBody)
		if string(jsonIPv4GatewayTemplateItemPutRequestBody) == "{}" { // empty struct
			dst.IPv4GatewayTemplateItemPutRequestBody = nil
		} else {
			if err = validator.Validate(dst.IPv4GatewayTemplateItemPutRequestBody); err != nil {
				dst.IPv4GatewayTemplateItemPutRequestBody = nil
			} else {
				match++
			}
		}
	} else {
		dst.IPv4GatewayTemplateItemPutRequestBody = nil
	}

	// try to unmarshal data into IPv4GroupTemplateItemPutRequestBody
	err = newStrictDecoder(data).Decode(&dst.IPv4GroupTemplateItemPutRequestBody)
	if err == nil {
		jsonIPv4GroupTemplateItemPutRequestBody, _ := json.Marshal(dst.IPv4GroupTemplateItemPutRequestBody)
		if string(jsonIPv4GroupTemplateItemPutRequestBody) == "{}" { // empty struct
			dst.IPv4GroupTemplateItemPutRequestBody = nil
		} else {
			if err = validator.Validate(dst.IPv4GroupTemplateItemPutRequestBody); err != nil {
				dst.IPv4GroupTemplateItemPutRequestBody = nil
			} else {
				match++
			}
		}
	} else {
		dst.IPv4GroupTemplateItemPutRequestBody = nil
	}

	// try to unmarshal data into IPv4ReservedRangeTemplateItemPutRequestBody
	err = newStrictDecoder(data).Decode(&dst.IPv4ReservedRangeTemplateItemPutRequestBody)
	if err == nil {
		jsonIPv4ReservedRangeTemplateItemPutRequestBody, _ := json.Marshal(dst.IPv4ReservedRangeTemplateItemPutRequestBody)
		if string(jsonIPv4ReservedRangeTemplateItemPutRequestBody) == "{}" { // empty struct
			dst.IPv4ReservedRangeTemplateItemPutRequestBody = nil
		} else {
			if err = validator.Validate(dst.IPv4ReservedRangeTemplateItemPutRequestBody); err != nil {
				dst.IPv4ReservedRangeTemplateItemPutRequestBody = nil
			} else {
				match++
			}
		}
	} else {
		dst.IPv4ReservedRangeTemplateItemPutRequestBody = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.IPv4DHCPRangeTemplateItemPutRequestBody = nil
		dst.IPv4GatewayTemplateItemPutRequestBody = nil
		dst.IPv4GroupTemplateItemPutRequestBody = nil
		dst.IPv4ReservedRangeTemplateItemPutRequestBody = nil

		return fmt.Errorf("data matches more than one schema in oneOf(PutTemplateItemRequest)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(PutTemplateItemRequest)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PutTemplateItemRequest) MarshalJSON() ([]byte, error) {
	if src.IPv4DHCPRangeTemplateItemPutRequestBody != nil {
		return json.Marshal(&src.IPv4DHCPRangeTemplateItemPutRequestBody)
	}

	if src.IPv4GatewayTemplateItemPutRequestBody != nil {
		return json.Marshal(&src.IPv4GatewayTemplateItemPutRequestBody)
	}

	if src.IPv4GroupTemplateItemPutRequestBody != nil {
		return json.Marshal(&src.IPv4GroupTemplateItemPutRequestBody)
	}

	if src.IPv4ReservedRangeTemplateItemPutRequestBody != nil {
		return json.Marshal(&src.IPv4ReservedRangeTemplateItemPutRequestBody)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PutTemplateItemRequest) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.IPv4DHCPRangeTemplateItemPutRequestBody != nil {
		return obj.IPv4DHCPRangeTemplateItemPutRequestBody
	}

	if obj.IPv4GatewayTemplateItemPutRequestBody != nil {
		return obj.IPv4GatewayTemplateItemPutRequestBody
	}

	if obj.IPv4GroupTemplateItemPutRequestBody != nil {
		return obj.IPv4GroupTemplateItemPutRequestBody
	}

	if obj.IPv4ReservedRangeTemplateItemPutRequestBody != nil {
		return obj.IPv4ReservedRangeTemplateItemPutRequestBody
	}

	// all schemas are nil
	return nil
}

type NullablePutTemplateItemRequest struct {
	value *PutTemplateItemRequest
	isSet bool
}

func (v NullablePutTemplateItemRequest) Get() *PutTemplateItemRequest {
	return v.value
}

func (v *NullablePutTemplateItemRequest) Set(val *PutTemplateItemRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePutTemplateItemRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePutTemplateItemRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutTemplateItemRequest(val *PutTemplateItemRequest) *NullablePutTemplateItemRequest {
	return &NullablePutTemplateItemRequest{value: val, isSet: true}
}

func (v NullablePutTemplateItemRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutTemplateItemRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


