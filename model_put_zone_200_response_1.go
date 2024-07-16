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

// PutZone200Response1 - struct for PutZone200Response1
type PutZone200Response1 struct {
	ENUMNumberZone *ENUMNumberZone
	ENUMZone *ENUMZone
	InternalRootZone *InternalRootZone
	ResponsePolicyZone *ResponsePolicyZone
	Zone *Zone
}

// ENUMNumberZoneAsPutZone200Response1 is a convenience function that returns ENUMNumberZone wrapped in PutZone200Response1
func ENUMNumberZoneAsPutZone200Response1(v *ENUMNumberZone) PutZone200Response1 {
	return PutZone200Response1{
		ENUMNumberZone: v,
	}
}

// ENUMZoneAsPutZone200Response1 is a convenience function that returns ENUMZone wrapped in PutZone200Response1
func ENUMZoneAsPutZone200Response1(v *ENUMZone) PutZone200Response1 {
	return PutZone200Response1{
		ENUMZone: v,
	}
}

// InternalRootZoneAsPutZone200Response1 is a convenience function that returns InternalRootZone wrapped in PutZone200Response1
func InternalRootZoneAsPutZone200Response1(v *InternalRootZone) PutZone200Response1 {
	return PutZone200Response1{
		InternalRootZone: v,
	}
}

// ResponsePolicyZoneAsPutZone200Response1 is a convenience function that returns ResponsePolicyZone wrapped in PutZone200Response1
func ResponsePolicyZoneAsPutZone200Response1(v *ResponsePolicyZone) PutZone200Response1 {
	return PutZone200Response1{
		ResponsePolicyZone: v,
	}
}

// ZoneAsPutZone200Response1 is a convenience function that returns Zone wrapped in PutZone200Response1
func ZoneAsPutZone200Response1(v *Zone) PutZone200Response1 {
	return PutZone200Response1{
		Zone: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *PutZone200Response1) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ENUMNumberZone
	err = newStrictDecoder(data).Decode(&dst.ENUMNumberZone)
	if err == nil {
		jsonENUMNumberZone, _ := json.Marshal(dst.ENUMNumberZone)
		if string(jsonENUMNumberZone) == "{}" { // empty struct
			dst.ENUMNumberZone = nil
		} else {
			if err = validator.Validate(dst.ENUMNumberZone); err != nil {
				dst.ENUMNumberZone = nil
			} else {
				match++
			}
		}
	} else {
		dst.ENUMNumberZone = nil
	}

	// try to unmarshal data into ENUMZone
	err = newStrictDecoder(data).Decode(&dst.ENUMZone)
	if err == nil {
		jsonENUMZone, _ := json.Marshal(dst.ENUMZone)
		if string(jsonENUMZone) == "{}" { // empty struct
			dst.ENUMZone = nil
		} else {
			if err = validator.Validate(dst.ENUMZone); err != nil {
				dst.ENUMZone = nil
			} else {
				match++
			}
		}
	} else {
		dst.ENUMZone = nil
	}

	// try to unmarshal data into InternalRootZone
	err = newStrictDecoder(data).Decode(&dst.InternalRootZone)
	if err == nil {
		jsonInternalRootZone, _ := json.Marshal(dst.InternalRootZone)
		if string(jsonInternalRootZone) == "{}" { // empty struct
			dst.InternalRootZone = nil
		} else {
			if err = validator.Validate(dst.InternalRootZone); err != nil {
				dst.InternalRootZone = nil
			} else {
				match++
			}
		}
	} else {
		dst.InternalRootZone = nil
	}

	// try to unmarshal data into ResponsePolicyZone
	err = newStrictDecoder(data).Decode(&dst.ResponsePolicyZone)
	if err == nil {
		jsonResponsePolicyZone, _ := json.Marshal(dst.ResponsePolicyZone)
		if string(jsonResponsePolicyZone) == "{}" { // empty struct
			dst.ResponsePolicyZone = nil
		} else {
			if err = validator.Validate(dst.ResponsePolicyZone); err != nil {
				dst.ResponsePolicyZone = nil
			} else {
				match++
			}
		}
	} else {
		dst.ResponsePolicyZone = nil
	}

	// try to unmarshal data into Zone
	err = newStrictDecoder(data).Decode(&dst.Zone)
	if err == nil {
		jsonZone, _ := json.Marshal(dst.Zone)
		if string(jsonZone) == "{}" { // empty struct
			dst.Zone = nil
		} else {
			if err = validator.Validate(dst.Zone); err != nil {
				dst.Zone = nil
			} else {
				match++
			}
		}
	} else {
		dst.Zone = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ENUMNumberZone = nil
		dst.ENUMZone = nil
		dst.InternalRootZone = nil
		dst.ResponsePolicyZone = nil
		dst.Zone = nil

		return fmt.Errorf("data matches more than one schema in oneOf(PutZone200Response1)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(PutZone200Response1)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PutZone200Response1) MarshalJSON() ([]byte, error) {
	if src.ENUMNumberZone != nil {
		return json.Marshal(&src.ENUMNumberZone)
	}

	if src.ENUMZone != nil {
		return json.Marshal(&src.ENUMZone)
	}

	if src.InternalRootZone != nil {
		return json.Marshal(&src.InternalRootZone)
	}

	if src.ResponsePolicyZone != nil {
		return json.Marshal(&src.ResponsePolicyZone)
	}

	if src.Zone != nil {
		return json.Marshal(&src.Zone)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PutZone200Response1) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ENUMNumberZone != nil {
		return obj.ENUMNumberZone
	}

	if obj.ENUMZone != nil {
		return obj.ENUMZone
	}

	if obj.InternalRootZone != nil {
		return obj.InternalRootZone
	}

	if obj.ResponsePolicyZone != nil {
		return obj.ResponsePolicyZone
	}

	if obj.Zone != nil {
		return obj.Zone
	}

	// all schemas are nil
	return nil
}

type NullablePutZone200Response1 struct {
	value *PutZone200Response1
	isSet bool
}

func (v NullablePutZone200Response1) Get() *PutZone200Response1 {
	return v.value
}

func (v *NullablePutZone200Response1) Set(val *PutZone200Response1) {
	v.value = val
	v.isSet = true
}

func (v NullablePutZone200Response1) IsSet() bool {
	return v.isSet
}

func (v *NullablePutZone200Response1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutZone200Response1(val *PutZone200Response1) *NullablePutZone200Response1 {
	return &NullablePutZone200Response1{value: val, isSet: true}
}

func (v NullablePutZone200Response1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutZone200Response1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

