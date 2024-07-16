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

// GetAddresses200Response1DataInner - struct for GetAddresses200Response1DataInner
type GetAddresses200Response1DataInner struct {
	IPv4Address *IPv4Address
	IPv6Address *IPv6Address
}

// IPv4AddressAsGetAddresses200Response1DataInner is a convenience function that returns IPv4Address wrapped in GetAddresses200Response1DataInner
func IPv4AddressAsGetAddresses200Response1DataInner(v *IPv4Address) GetAddresses200Response1DataInner {
	return GetAddresses200Response1DataInner{
		IPv4Address: v,
	}
}

// IPv6AddressAsGetAddresses200Response1DataInner is a convenience function that returns IPv6Address wrapped in GetAddresses200Response1DataInner
func IPv6AddressAsGetAddresses200Response1DataInner(v *IPv6Address) GetAddresses200Response1DataInner {
	return GetAddresses200Response1DataInner{
		IPv6Address: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetAddresses200Response1DataInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into IPv4Address
	err = newStrictDecoder(data).Decode(&dst.IPv4Address)
	if err == nil {
		jsonIPv4Address, _ := json.Marshal(dst.IPv4Address)
		if string(jsonIPv4Address) == "{}" { // empty struct
			dst.IPv4Address = nil
		} else {
			if err = validator.Validate(dst.IPv4Address); err != nil {
				dst.IPv4Address = nil
			} else {
				match++
			}
		}
	} else {
		dst.IPv4Address = nil
	}

	// try to unmarshal data into IPv6Address
	err = newStrictDecoder(data).Decode(&dst.IPv6Address)
	if err == nil {
		jsonIPv6Address, _ := json.Marshal(dst.IPv6Address)
		if string(jsonIPv6Address) == "{}" { // empty struct
			dst.IPv6Address = nil
		} else {
			if err = validator.Validate(dst.IPv6Address); err != nil {
				dst.IPv6Address = nil
			} else {
				match++
			}
		}
	} else {
		dst.IPv6Address = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.IPv4Address = nil
		dst.IPv6Address = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetAddresses200Response1DataInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetAddresses200Response1DataInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetAddresses200Response1DataInner) MarshalJSON() ([]byte, error) {
	if src.IPv4Address != nil {
		return json.Marshal(&src.IPv4Address)
	}

	if src.IPv6Address != nil {
		return json.Marshal(&src.IPv6Address)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetAddresses200Response1DataInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.IPv4Address != nil {
		return obj.IPv4Address
	}

	if obj.IPv6Address != nil {
		return obj.IPv6Address
	}

	// all schemas are nil
	return nil
}

type NullableGetAddresses200Response1DataInner struct {
	value *GetAddresses200Response1DataInner
	isSet bool
}

func (v NullableGetAddresses200Response1DataInner) Get() *GetAddresses200Response1DataInner {
	return v.value
}

func (v *NullableGetAddresses200Response1DataInner) Set(val *GetAddresses200Response1DataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAddresses200Response1DataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAddresses200Response1DataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAddresses200Response1DataInner(val *GetAddresses200Response1DataInner) *NullableGetAddresses200Response1DataInner {
	return &NullableGetAddresses200Response1DataInner{value: val, isSet: true}
}

func (v NullableGetAddresses200Response1DataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAddresses200Response1DataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

