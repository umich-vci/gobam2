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

// PostCollectionDeploymentRoleRequest - struct for PostCollectionDeploymentRoleRequest
type PostCollectionDeploymentRoleRequest struct {
	DHCPDeploymentRolePostRequestBody *DHCPDeploymentRolePostRequestBody
	DNSDeploymentRolePostRequestBody *DNSDeploymentRolePostRequestBody
	TFTPDeploymentRolePostRequestBody *TFTPDeploymentRolePostRequestBody
}

// DHCPDeploymentRolePostRequestBodyAsPostCollectionDeploymentRoleRequest is a convenience function that returns DHCPDeploymentRolePostRequestBody wrapped in PostCollectionDeploymentRoleRequest
func DHCPDeploymentRolePostRequestBodyAsPostCollectionDeploymentRoleRequest(v *DHCPDeploymentRolePostRequestBody) PostCollectionDeploymentRoleRequest {
	return PostCollectionDeploymentRoleRequest{
		DHCPDeploymentRolePostRequestBody: v,
	}
}

// DNSDeploymentRolePostRequestBodyAsPostCollectionDeploymentRoleRequest is a convenience function that returns DNSDeploymentRolePostRequestBody wrapped in PostCollectionDeploymentRoleRequest
func DNSDeploymentRolePostRequestBodyAsPostCollectionDeploymentRoleRequest(v *DNSDeploymentRolePostRequestBody) PostCollectionDeploymentRoleRequest {
	return PostCollectionDeploymentRoleRequest{
		DNSDeploymentRolePostRequestBody: v,
	}
}

// TFTPDeploymentRolePostRequestBodyAsPostCollectionDeploymentRoleRequest is a convenience function that returns TFTPDeploymentRolePostRequestBody wrapped in PostCollectionDeploymentRoleRequest
func TFTPDeploymentRolePostRequestBodyAsPostCollectionDeploymentRoleRequest(v *TFTPDeploymentRolePostRequestBody) PostCollectionDeploymentRoleRequest {
	return PostCollectionDeploymentRoleRequest{
		TFTPDeploymentRolePostRequestBody: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *PostCollectionDeploymentRoleRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into DHCPDeploymentRolePostRequestBody
	err = newStrictDecoder(data).Decode(&dst.DHCPDeploymentRolePostRequestBody)
	if err == nil {
		jsonDHCPDeploymentRolePostRequestBody, _ := json.Marshal(dst.DHCPDeploymentRolePostRequestBody)
		if string(jsonDHCPDeploymentRolePostRequestBody) == "{}" { // empty struct
			dst.DHCPDeploymentRolePostRequestBody = nil
		} else {
			if err = validator.Validate(dst.DHCPDeploymentRolePostRequestBody); err != nil {
				dst.DHCPDeploymentRolePostRequestBody = nil
			} else {
				match++
			}
		}
	} else {
		dst.DHCPDeploymentRolePostRequestBody = nil
	}

	// try to unmarshal data into DNSDeploymentRolePostRequestBody
	err = newStrictDecoder(data).Decode(&dst.DNSDeploymentRolePostRequestBody)
	if err == nil {
		jsonDNSDeploymentRolePostRequestBody, _ := json.Marshal(dst.DNSDeploymentRolePostRequestBody)
		if string(jsonDNSDeploymentRolePostRequestBody) == "{}" { // empty struct
			dst.DNSDeploymentRolePostRequestBody = nil
		} else {
			if err = validator.Validate(dst.DNSDeploymentRolePostRequestBody); err != nil {
				dst.DNSDeploymentRolePostRequestBody = nil
			} else {
				match++
			}
		}
	} else {
		dst.DNSDeploymentRolePostRequestBody = nil
	}

	// try to unmarshal data into TFTPDeploymentRolePostRequestBody
	err = newStrictDecoder(data).Decode(&dst.TFTPDeploymentRolePostRequestBody)
	if err == nil {
		jsonTFTPDeploymentRolePostRequestBody, _ := json.Marshal(dst.TFTPDeploymentRolePostRequestBody)
		if string(jsonTFTPDeploymentRolePostRequestBody) == "{}" { // empty struct
			dst.TFTPDeploymentRolePostRequestBody = nil
		} else {
			if err = validator.Validate(dst.TFTPDeploymentRolePostRequestBody); err != nil {
				dst.TFTPDeploymentRolePostRequestBody = nil
			} else {
				match++
			}
		}
	} else {
		dst.TFTPDeploymentRolePostRequestBody = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.DHCPDeploymentRolePostRequestBody = nil
		dst.DNSDeploymentRolePostRequestBody = nil
		dst.TFTPDeploymentRolePostRequestBody = nil

		return fmt.Errorf("data matches more than one schema in oneOf(PostCollectionDeploymentRoleRequest)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(PostCollectionDeploymentRoleRequest)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PostCollectionDeploymentRoleRequest) MarshalJSON() ([]byte, error) {
	if src.DHCPDeploymentRolePostRequestBody != nil {
		return json.Marshal(&src.DHCPDeploymentRolePostRequestBody)
	}

	if src.DNSDeploymentRolePostRequestBody != nil {
		return json.Marshal(&src.DNSDeploymentRolePostRequestBody)
	}

	if src.TFTPDeploymentRolePostRequestBody != nil {
		return json.Marshal(&src.TFTPDeploymentRolePostRequestBody)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PostCollectionDeploymentRoleRequest) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.DHCPDeploymentRolePostRequestBody != nil {
		return obj.DHCPDeploymentRolePostRequestBody
	}

	if obj.DNSDeploymentRolePostRequestBody != nil {
		return obj.DNSDeploymentRolePostRequestBody
	}

	if obj.TFTPDeploymentRolePostRequestBody != nil {
		return obj.TFTPDeploymentRolePostRequestBody
	}

	// all schemas are nil
	return nil
}

type NullablePostCollectionDeploymentRoleRequest struct {
	value *PostCollectionDeploymentRoleRequest
	isSet bool
}

func (v NullablePostCollectionDeploymentRoleRequest) Get() *PostCollectionDeploymentRoleRequest {
	return v.value
}

func (v *NullablePostCollectionDeploymentRoleRequest) Set(val *PostCollectionDeploymentRoleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostCollectionDeploymentRoleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostCollectionDeploymentRoleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostCollectionDeploymentRoleRequest(val *PostCollectionDeploymentRoleRequest) *NullablePostCollectionDeploymentRoleRequest {
	return &NullablePostCollectionDeploymentRoleRequest{value: val, isSet: true}
}

func (v NullablePostCollectionDeploymentRoleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostCollectionDeploymentRoleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


