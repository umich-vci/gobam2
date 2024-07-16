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

// GetSigningKeys200Response1DataInner - struct for GetSigningKeys200Response1DataInner
type GetSigningKeys200Response1DataInner struct {
	KeySigningKey *KeySigningKey
	TSIGKey *TSIGKey
	ZoneSigningKey *ZoneSigningKey
}

// KeySigningKeyAsGetSigningKeys200Response1DataInner is a convenience function that returns KeySigningKey wrapped in GetSigningKeys200Response1DataInner
func KeySigningKeyAsGetSigningKeys200Response1DataInner(v *KeySigningKey) GetSigningKeys200Response1DataInner {
	return GetSigningKeys200Response1DataInner{
		KeySigningKey: v,
	}
}

// TSIGKeyAsGetSigningKeys200Response1DataInner is a convenience function that returns TSIGKey wrapped in GetSigningKeys200Response1DataInner
func TSIGKeyAsGetSigningKeys200Response1DataInner(v *TSIGKey) GetSigningKeys200Response1DataInner {
	return GetSigningKeys200Response1DataInner{
		TSIGKey: v,
	}
}

// ZoneSigningKeyAsGetSigningKeys200Response1DataInner is a convenience function that returns ZoneSigningKey wrapped in GetSigningKeys200Response1DataInner
func ZoneSigningKeyAsGetSigningKeys200Response1DataInner(v *ZoneSigningKey) GetSigningKeys200Response1DataInner {
	return GetSigningKeys200Response1DataInner{
		ZoneSigningKey: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetSigningKeys200Response1DataInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into KeySigningKey
	err = newStrictDecoder(data).Decode(&dst.KeySigningKey)
	if err == nil {
		jsonKeySigningKey, _ := json.Marshal(dst.KeySigningKey)
		if string(jsonKeySigningKey) == "{}" { // empty struct
			dst.KeySigningKey = nil
		} else {
			if err = validator.Validate(dst.KeySigningKey); err != nil {
				dst.KeySigningKey = nil
			} else {
				match++
			}
		}
	} else {
		dst.KeySigningKey = nil
	}

	// try to unmarshal data into TSIGKey
	err = newStrictDecoder(data).Decode(&dst.TSIGKey)
	if err == nil {
		jsonTSIGKey, _ := json.Marshal(dst.TSIGKey)
		if string(jsonTSIGKey) == "{}" { // empty struct
			dst.TSIGKey = nil
		} else {
			if err = validator.Validate(dst.TSIGKey); err != nil {
				dst.TSIGKey = nil
			} else {
				match++
			}
		}
	} else {
		dst.TSIGKey = nil
	}

	// try to unmarshal data into ZoneSigningKey
	err = newStrictDecoder(data).Decode(&dst.ZoneSigningKey)
	if err == nil {
		jsonZoneSigningKey, _ := json.Marshal(dst.ZoneSigningKey)
		if string(jsonZoneSigningKey) == "{}" { // empty struct
			dst.ZoneSigningKey = nil
		} else {
			if err = validator.Validate(dst.ZoneSigningKey); err != nil {
				dst.ZoneSigningKey = nil
			} else {
				match++
			}
		}
	} else {
		dst.ZoneSigningKey = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.KeySigningKey = nil
		dst.TSIGKey = nil
		dst.ZoneSigningKey = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetSigningKeys200Response1DataInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetSigningKeys200Response1DataInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetSigningKeys200Response1DataInner) MarshalJSON() ([]byte, error) {
	if src.KeySigningKey != nil {
		return json.Marshal(&src.KeySigningKey)
	}

	if src.TSIGKey != nil {
		return json.Marshal(&src.TSIGKey)
	}

	if src.ZoneSigningKey != nil {
		return json.Marshal(&src.ZoneSigningKey)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetSigningKeys200Response1DataInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.KeySigningKey != nil {
		return obj.KeySigningKey
	}

	if obj.TSIGKey != nil {
		return obj.TSIGKey
	}

	if obj.ZoneSigningKey != nil {
		return obj.ZoneSigningKey
	}

	// all schemas are nil
	return nil
}

type NullableGetSigningKeys200Response1DataInner struct {
	value *GetSigningKeys200Response1DataInner
	isSet bool
}

func (v NullableGetSigningKeys200Response1DataInner) Get() *GetSigningKeys200Response1DataInner {
	return v.value
}

func (v *NullableGetSigningKeys200Response1DataInner) Set(val *GetSigningKeys200Response1DataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSigningKeys200Response1DataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSigningKeys200Response1DataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSigningKeys200Response1DataInner(val *GetSigningKeys200Response1DataInner) *NullableGetSigningKeys200Response1DataInner {
	return &NullableGetSigningKeys200Response1DataInner{value: val, isSet: true}
}

func (v NullableGetSigningKeys200Response1DataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSigningKeys200Response1DataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


