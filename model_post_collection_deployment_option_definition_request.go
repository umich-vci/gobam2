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

// PostCollectionDeploymentOptionDefinitionRequest - struct for PostCollectionDeploymentOptionDefinitionRequest
type PostCollectionDeploymentOptionDefinitionRequest struct {
	DHCPVendorOptionDefinitionPostRequestBody *DHCPVendorOptionDefinitionPostRequestBody
	DHCPv4CustomClientOptionDefinitionPostRequestBody *DHCPv4CustomClientOptionDefinitionPostRequestBody
}

// DHCPVendorOptionDefinitionPostRequestBodyAsPostCollectionDeploymentOptionDefinitionRequest is a convenience function that returns DHCPVendorOptionDefinitionPostRequestBody wrapped in PostCollectionDeploymentOptionDefinitionRequest
func DHCPVendorOptionDefinitionPostRequestBodyAsPostCollectionDeploymentOptionDefinitionRequest(v *DHCPVendorOptionDefinitionPostRequestBody) PostCollectionDeploymentOptionDefinitionRequest {
	return PostCollectionDeploymentOptionDefinitionRequest{
		DHCPVendorOptionDefinitionPostRequestBody: v,
	}
}

// DHCPv4CustomClientOptionDefinitionPostRequestBodyAsPostCollectionDeploymentOptionDefinitionRequest is a convenience function that returns DHCPv4CustomClientOptionDefinitionPostRequestBody wrapped in PostCollectionDeploymentOptionDefinitionRequest
func DHCPv4CustomClientOptionDefinitionPostRequestBodyAsPostCollectionDeploymentOptionDefinitionRequest(v *DHCPv4CustomClientOptionDefinitionPostRequestBody) PostCollectionDeploymentOptionDefinitionRequest {
	return PostCollectionDeploymentOptionDefinitionRequest{
		DHCPv4CustomClientOptionDefinitionPostRequestBody: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *PostCollectionDeploymentOptionDefinitionRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into DHCPVendorOptionDefinitionPostRequestBody
	err = newStrictDecoder(data).Decode(&dst.DHCPVendorOptionDefinitionPostRequestBody)
	if err == nil {
		jsonDHCPVendorOptionDefinitionPostRequestBody, _ := json.Marshal(dst.DHCPVendorOptionDefinitionPostRequestBody)
		if string(jsonDHCPVendorOptionDefinitionPostRequestBody) == "{}" { // empty struct
			dst.DHCPVendorOptionDefinitionPostRequestBody = nil
		} else {
			if err = validator.Validate(dst.DHCPVendorOptionDefinitionPostRequestBody); err != nil {
				dst.DHCPVendorOptionDefinitionPostRequestBody = nil
			} else {
				match++
			}
		}
	} else {
		dst.DHCPVendorOptionDefinitionPostRequestBody = nil
	}

	// try to unmarshal data into DHCPv4CustomClientOptionDefinitionPostRequestBody
	err = newStrictDecoder(data).Decode(&dst.DHCPv4CustomClientOptionDefinitionPostRequestBody)
	if err == nil {
		jsonDHCPv4CustomClientOptionDefinitionPostRequestBody, _ := json.Marshal(dst.DHCPv4CustomClientOptionDefinitionPostRequestBody)
		if string(jsonDHCPv4CustomClientOptionDefinitionPostRequestBody) == "{}" { // empty struct
			dst.DHCPv4CustomClientOptionDefinitionPostRequestBody = nil
		} else {
			if err = validator.Validate(dst.DHCPv4CustomClientOptionDefinitionPostRequestBody); err != nil {
				dst.DHCPv4CustomClientOptionDefinitionPostRequestBody = nil
			} else {
				match++
			}
		}
	} else {
		dst.DHCPv4CustomClientOptionDefinitionPostRequestBody = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.DHCPVendorOptionDefinitionPostRequestBody = nil
		dst.DHCPv4CustomClientOptionDefinitionPostRequestBody = nil

		return fmt.Errorf("data matches more than one schema in oneOf(PostCollectionDeploymentOptionDefinitionRequest)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(PostCollectionDeploymentOptionDefinitionRequest)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PostCollectionDeploymentOptionDefinitionRequest) MarshalJSON() ([]byte, error) {
	if src.DHCPVendorOptionDefinitionPostRequestBody != nil {
		return json.Marshal(&src.DHCPVendorOptionDefinitionPostRequestBody)
	}

	if src.DHCPv4CustomClientOptionDefinitionPostRequestBody != nil {
		return json.Marshal(&src.DHCPv4CustomClientOptionDefinitionPostRequestBody)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PostCollectionDeploymentOptionDefinitionRequest) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.DHCPVendorOptionDefinitionPostRequestBody != nil {
		return obj.DHCPVendorOptionDefinitionPostRequestBody
	}

	if obj.DHCPv4CustomClientOptionDefinitionPostRequestBody != nil {
		return obj.DHCPv4CustomClientOptionDefinitionPostRequestBody
	}

	// all schemas are nil
	return nil
}

type NullablePostCollectionDeploymentOptionDefinitionRequest struct {
	value *PostCollectionDeploymentOptionDefinitionRequest
	isSet bool
}

func (v NullablePostCollectionDeploymentOptionDefinitionRequest) Get() *PostCollectionDeploymentOptionDefinitionRequest {
	return v.value
}

func (v *NullablePostCollectionDeploymentOptionDefinitionRequest) Set(val *PostCollectionDeploymentOptionDefinitionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostCollectionDeploymentOptionDefinitionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostCollectionDeploymentOptionDefinitionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostCollectionDeploymentOptionDefinitionRequest(val *PostCollectionDeploymentOptionDefinitionRequest) *NullablePostCollectionDeploymentOptionDefinitionRequest {
	return &NullablePostCollectionDeploymentOptionDefinitionRequest{value: val, isSet: true}
}

func (v NullablePostCollectionDeploymentOptionDefinitionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostCollectionDeploymentOptionDefinitionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


