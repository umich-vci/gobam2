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

// GetInterfaces200Response1DataInner - struct for GetInterfaces200Response1DataInner
type GetInterfaces200Response1DataInner struct {
	NetworkInterface *NetworkInterface
	PublishedInterface *PublishedInterface
}

// NetworkInterfaceAsGetInterfaces200Response1DataInner is a convenience function that returns NetworkInterface wrapped in GetInterfaces200Response1DataInner
func NetworkInterfaceAsGetInterfaces200Response1DataInner(v *NetworkInterface) GetInterfaces200Response1DataInner {
	return GetInterfaces200Response1DataInner{
		NetworkInterface: v,
	}
}

// PublishedInterfaceAsGetInterfaces200Response1DataInner is a convenience function that returns PublishedInterface wrapped in GetInterfaces200Response1DataInner
func PublishedInterfaceAsGetInterfaces200Response1DataInner(v *PublishedInterface) GetInterfaces200Response1DataInner {
	return GetInterfaces200Response1DataInner{
		PublishedInterface: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetInterfaces200Response1DataInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into NetworkInterface
	err = newStrictDecoder(data).Decode(&dst.NetworkInterface)
	if err == nil {
		jsonNetworkInterface, _ := json.Marshal(dst.NetworkInterface)
		if string(jsonNetworkInterface) == "{}" { // empty struct
			dst.NetworkInterface = nil
		} else {
			if err = validator.Validate(dst.NetworkInterface); err != nil {
				dst.NetworkInterface = nil
			} else {
				match++
			}
		}
	} else {
		dst.NetworkInterface = nil
	}

	// try to unmarshal data into PublishedInterface
	err = newStrictDecoder(data).Decode(&dst.PublishedInterface)
	if err == nil {
		jsonPublishedInterface, _ := json.Marshal(dst.PublishedInterface)
		if string(jsonPublishedInterface) == "{}" { // empty struct
			dst.PublishedInterface = nil
		} else {
			if err = validator.Validate(dst.PublishedInterface); err != nil {
				dst.PublishedInterface = nil
			} else {
				match++
			}
		}
	} else {
		dst.PublishedInterface = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.NetworkInterface = nil
		dst.PublishedInterface = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetInterfaces200Response1DataInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetInterfaces200Response1DataInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetInterfaces200Response1DataInner) MarshalJSON() ([]byte, error) {
	if src.NetworkInterface != nil {
		return json.Marshal(&src.NetworkInterface)
	}

	if src.PublishedInterface != nil {
		return json.Marshal(&src.PublishedInterface)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetInterfaces200Response1DataInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.NetworkInterface != nil {
		return obj.NetworkInterface
	}

	if obj.PublishedInterface != nil {
		return obj.PublishedInterface
	}

	// all schemas are nil
	return nil
}

type NullableGetInterfaces200Response1DataInner struct {
	value *GetInterfaces200Response1DataInner
	isSet bool
}

func (v NullableGetInterfaces200Response1DataInner) Get() *GetInterfaces200Response1DataInner {
	return v.value
}

func (v *NullableGetInterfaces200Response1DataInner) Set(val *GetInterfaces200Response1DataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetInterfaces200Response1DataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetInterfaces200Response1DataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetInterfaces200Response1DataInner(val *GetInterfaces200Response1DataInner) *NullableGetInterfaces200Response1DataInner {
	return &NullableGetInterfaces200Response1DataInner{value: val, isSet: true}
}

func (v NullableGetInterfaces200Response1DataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetInterfaces200Response1DataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

