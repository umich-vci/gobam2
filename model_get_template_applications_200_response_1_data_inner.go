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

// GetTemplateApplications200Response1DataInner - struct for GetTemplateApplications200Response1DataInner
type GetTemplateApplications200Response1DataInner struct {
	IPv4TemplateApplication *IPv4TemplateApplication
	ZoneTemplateApplication *ZoneTemplateApplication
}

// IPv4TemplateApplicationAsGetTemplateApplications200Response1DataInner is a convenience function that returns IPv4TemplateApplication wrapped in GetTemplateApplications200Response1DataInner
func IPv4TemplateApplicationAsGetTemplateApplications200Response1DataInner(v *IPv4TemplateApplication) GetTemplateApplications200Response1DataInner {
	return GetTemplateApplications200Response1DataInner{
		IPv4TemplateApplication: v,
	}
}

// ZoneTemplateApplicationAsGetTemplateApplications200Response1DataInner is a convenience function that returns ZoneTemplateApplication wrapped in GetTemplateApplications200Response1DataInner
func ZoneTemplateApplicationAsGetTemplateApplications200Response1DataInner(v *ZoneTemplateApplication) GetTemplateApplications200Response1DataInner {
	return GetTemplateApplications200Response1DataInner{
		ZoneTemplateApplication: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetTemplateApplications200Response1DataInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into IPv4TemplateApplication
	err = newStrictDecoder(data).Decode(&dst.IPv4TemplateApplication)
	if err == nil {
		jsonIPv4TemplateApplication, _ := json.Marshal(dst.IPv4TemplateApplication)
		if string(jsonIPv4TemplateApplication) == "{}" { // empty struct
			dst.IPv4TemplateApplication = nil
		} else {
			if err = validator.Validate(dst.IPv4TemplateApplication); err != nil {
				dst.IPv4TemplateApplication = nil
			} else {
				match++
			}
		}
	} else {
		dst.IPv4TemplateApplication = nil
	}

	// try to unmarshal data into ZoneTemplateApplication
	err = newStrictDecoder(data).Decode(&dst.ZoneTemplateApplication)
	if err == nil {
		jsonZoneTemplateApplication, _ := json.Marshal(dst.ZoneTemplateApplication)
		if string(jsonZoneTemplateApplication) == "{}" { // empty struct
			dst.ZoneTemplateApplication = nil
		} else {
			if err = validator.Validate(dst.ZoneTemplateApplication); err != nil {
				dst.ZoneTemplateApplication = nil
			} else {
				match++
			}
		}
	} else {
		dst.ZoneTemplateApplication = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.IPv4TemplateApplication = nil
		dst.ZoneTemplateApplication = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetTemplateApplications200Response1DataInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetTemplateApplications200Response1DataInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetTemplateApplications200Response1DataInner) MarshalJSON() ([]byte, error) {
	if src.IPv4TemplateApplication != nil {
		return json.Marshal(&src.IPv4TemplateApplication)
	}

	if src.ZoneTemplateApplication != nil {
		return json.Marshal(&src.ZoneTemplateApplication)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetTemplateApplications200Response1DataInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.IPv4TemplateApplication != nil {
		return obj.IPv4TemplateApplication
	}

	if obj.ZoneTemplateApplication != nil {
		return obj.ZoneTemplateApplication
	}

	// all schemas are nil
	return nil
}

type NullableGetTemplateApplications200Response1DataInner struct {
	value *GetTemplateApplications200Response1DataInner
	isSet bool
}

func (v NullableGetTemplateApplications200Response1DataInner) Get() *GetTemplateApplications200Response1DataInner {
	return v.value
}

func (v *NullableGetTemplateApplications200Response1DataInner) Set(val *GetTemplateApplications200Response1DataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTemplateApplications200Response1DataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTemplateApplications200Response1DataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTemplateApplications200Response1DataInner(val *GetTemplateApplications200Response1DataInner) *NullableGetTemplateApplications200Response1DataInner {
	return &NullableGetTemplateApplications200Response1DataInner{value: val, isSet: true}
}

func (v NullableGetTemplateApplications200Response1DataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTemplateApplications200Response1DataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

