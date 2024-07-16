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

// GetLeases200Response1DataInner - struct for GetLeases200Response1DataInner
type GetLeases200Response1DataInner struct {
	DHCPv4Lease *DHCPv4Lease
	DHCPv6Lease *DHCPv6Lease
}

// DHCPv4LeaseAsGetLeases200Response1DataInner is a convenience function that returns DHCPv4Lease wrapped in GetLeases200Response1DataInner
func DHCPv4LeaseAsGetLeases200Response1DataInner(v *DHCPv4Lease) GetLeases200Response1DataInner {
	return GetLeases200Response1DataInner{
		DHCPv4Lease: v,
	}
}

// DHCPv6LeaseAsGetLeases200Response1DataInner is a convenience function that returns DHCPv6Lease wrapped in GetLeases200Response1DataInner
func DHCPv6LeaseAsGetLeases200Response1DataInner(v *DHCPv6Lease) GetLeases200Response1DataInner {
	return GetLeases200Response1DataInner{
		DHCPv6Lease: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetLeases200Response1DataInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into DHCPv4Lease
	err = newStrictDecoder(data).Decode(&dst.DHCPv4Lease)
	if err == nil {
		jsonDHCPv4Lease, _ := json.Marshal(dst.DHCPv4Lease)
		if string(jsonDHCPv4Lease) == "{}" { // empty struct
			dst.DHCPv4Lease = nil
		} else {
			if err = validator.Validate(dst.DHCPv4Lease); err != nil {
				dst.DHCPv4Lease = nil
			} else {
				match++
			}
		}
	} else {
		dst.DHCPv4Lease = nil
	}

	// try to unmarshal data into DHCPv6Lease
	err = newStrictDecoder(data).Decode(&dst.DHCPv6Lease)
	if err == nil {
		jsonDHCPv6Lease, _ := json.Marshal(dst.DHCPv6Lease)
		if string(jsonDHCPv6Lease) == "{}" { // empty struct
			dst.DHCPv6Lease = nil
		} else {
			if err = validator.Validate(dst.DHCPv6Lease); err != nil {
				dst.DHCPv6Lease = nil
			} else {
				match++
			}
		}
	} else {
		dst.DHCPv6Lease = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.DHCPv4Lease = nil
		dst.DHCPv6Lease = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetLeases200Response1DataInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetLeases200Response1DataInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetLeases200Response1DataInner) MarshalJSON() ([]byte, error) {
	if src.DHCPv4Lease != nil {
		return json.Marshal(&src.DHCPv4Lease)
	}

	if src.DHCPv6Lease != nil {
		return json.Marshal(&src.DHCPv6Lease)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetLeases200Response1DataInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.DHCPv4Lease != nil {
		return obj.DHCPv4Lease
	}

	if obj.DHCPv6Lease != nil {
		return obj.DHCPv6Lease
	}

	// all schemas are nil
	return nil
}

type NullableGetLeases200Response1DataInner struct {
	value *GetLeases200Response1DataInner
	isSet bool
}

func (v NullableGetLeases200Response1DataInner) Get() *GetLeases200Response1DataInner {
	return v.value
}

func (v *NullableGetLeases200Response1DataInner) Set(val *GetLeases200Response1DataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLeases200Response1DataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLeases200Response1DataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLeases200Response1DataInner(val *GetLeases200Response1DataInner) *NullableGetLeases200Response1DataInner {
	return &NullableGetLeases200Response1DataInner{value: val, isSet: true}
}

func (v NullableGetLeases200Response1DataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLeases200Response1DataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


