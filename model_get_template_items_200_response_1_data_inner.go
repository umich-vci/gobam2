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

// GetTemplateItems200Response1DataInner - struct for GetTemplateItems200Response1DataInner
type GetTemplateItems200Response1DataInner struct {
	IPv4DHCPRangeTemplateItem *IPv4DHCPRangeTemplateItem
	IPv4GatewayTemplateItem *IPv4GatewayTemplateItem
	IPv4GroupTemplateItem *IPv4GroupTemplateItem
	IPv4ReservedRangeTemplateItem *IPv4ReservedRangeTemplateItem
}

// IPv4DHCPRangeTemplateItemAsGetTemplateItems200Response1DataInner is a convenience function that returns IPv4DHCPRangeTemplateItem wrapped in GetTemplateItems200Response1DataInner
func IPv4DHCPRangeTemplateItemAsGetTemplateItems200Response1DataInner(v *IPv4DHCPRangeTemplateItem) GetTemplateItems200Response1DataInner {
	return GetTemplateItems200Response1DataInner{
		IPv4DHCPRangeTemplateItem: v,
	}
}

// IPv4GatewayTemplateItemAsGetTemplateItems200Response1DataInner is a convenience function that returns IPv4GatewayTemplateItem wrapped in GetTemplateItems200Response1DataInner
func IPv4GatewayTemplateItemAsGetTemplateItems200Response1DataInner(v *IPv4GatewayTemplateItem) GetTemplateItems200Response1DataInner {
	return GetTemplateItems200Response1DataInner{
		IPv4GatewayTemplateItem: v,
	}
}

// IPv4GroupTemplateItemAsGetTemplateItems200Response1DataInner is a convenience function that returns IPv4GroupTemplateItem wrapped in GetTemplateItems200Response1DataInner
func IPv4GroupTemplateItemAsGetTemplateItems200Response1DataInner(v *IPv4GroupTemplateItem) GetTemplateItems200Response1DataInner {
	return GetTemplateItems200Response1DataInner{
		IPv4GroupTemplateItem: v,
	}
}

// IPv4ReservedRangeTemplateItemAsGetTemplateItems200Response1DataInner is a convenience function that returns IPv4ReservedRangeTemplateItem wrapped in GetTemplateItems200Response1DataInner
func IPv4ReservedRangeTemplateItemAsGetTemplateItems200Response1DataInner(v *IPv4ReservedRangeTemplateItem) GetTemplateItems200Response1DataInner {
	return GetTemplateItems200Response1DataInner{
		IPv4ReservedRangeTemplateItem: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetTemplateItems200Response1DataInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into IPv4DHCPRangeTemplateItem
	err = newStrictDecoder(data).Decode(&dst.IPv4DHCPRangeTemplateItem)
	if err == nil {
		jsonIPv4DHCPRangeTemplateItem, _ := json.Marshal(dst.IPv4DHCPRangeTemplateItem)
		if string(jsonIPv4DHCPRangeTemplateItem) == "{}" { // empty struct
			dst.IPv4DHCPRangeTemplateItem = nil
		} else {
			if err = validator.Validate(dst.IPv4DHCPRangeTemplateItem); err != nil {
				dst.IPv4DHCPRangeTemplateItem = nil
			} else {
				match++
			}
		}
	} else {
		dst.IPv4DHCPRangeTemplateItem = nil
	}

	// try to unmarshal data into IPv4GatewayTemplateItem
	err = newStrictDecoder(data).Decode(&dst.IPv4GatewayTemplateItem)
	if err == nil {
		jsonIPv4GatewayTemplateItem, _ := json.Marshal(dst.IPv4GatewayTemplateItem)
		if string(jsonIPv4GatewayTemplateItem) == "{}" { // empty struct
			dst.IPv4GatewayTemplateItem = nil
		} else {
			if err = validator.Validate(dst.IPv4GatewayTemplateItem); err != nil {
				dst.IPv4GatewayTemplateItem = nil
			} else {
				match++
			}
		}
	} else {
		dst.IPv4GatewayTemplateItem = nil
	}

	// try to unmarshal data into IPv4GroupTemplateItem
	err = newStrictDecoder(data).Decode(&dst.IPv4GroupTemplateItem)
	if err == nil {
		jsonIPv4GroupTemplateItem, _ := json.Marshal(dst.IPv4GroupTemplateItem)
		if string(jsonIPv4GroupTemplateItem) == "{}" { // empty struct
			dst.IPv4GroupTemplateItem = nil
		} else {
			if err = validator.Validate(dst.IPv4GroupTemplateItem); err != nil {
				dst.IPv4GroupTemplateItem = nil
			} else {
				match++
			}
		}
	} else {
		dst.IPv4GroupTemplateItem = nil
	}

	// try to unmarshal data into IPv4ReservedRangeTemplateItem
	err = newStrictDecoder(data).Decode(&dst.IPv4ReservedRangeTemplateItem)
	if err == nil {
		jsonIPv4ReservedRangeTemplateItem, _ := json.Marshal(dst.IPv4ReservedRangeTemplateItem)
		if string(jsonIPv4ReservedRangeTemplateItem) == "{}" { // empty struct
			dst.IPv4ReservedRangeTemplateItem = nil
		} else {
			if err = validator.Validate(dst.IPv4ReservedRangeTemplateItem); err != nil {
				dst.IPv4ReservedRangeTemplateItem = nil
			} else {
				match++
			}
		}
	} else {
		dst.IPv4ReservedRangeTemplateItem = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.IPv4DHCPRangeTemplateItem = nil
		dst.IPv4GatewayTemplateItem = nil
		dst.IPv4GroupTemplateItem = nil
		dst.IPv4ReservedRangeTemplateItem = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetTemplateItems200Response1DataInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetTemplateItems200Response1DataInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetTemplateItems200Response1DataInner) MarshalJSON() ([]byte, error) {
	if src.IPv4DHCPRangeTemplateItem != nil {
		return json.Marshal(&src.IPv4DHCPRangeTemplateItem)
	}

	if src.IPv4GatewayTemplateItem != nil {
		return json.Marshal(&src.IPv4GatewayTemplateItem)
	}

	if src.IPv4GroupTemplateItem != nil {
		return json.Marshal(&src.IPv4GroupTemplateItem)
	}

	if src.IPv4ReservedRangeTemplateItem != nil {
		return json.Marshal(&src.IPv4ReservedRangeTemplateItem)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetTemplateItems200Response1DataInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.IPv4DHCPRangeTemplateItem != nil {
		return obj.IPv4DHCPRangeTemplateItem
	}

	if obj.IPv4GatewayTemplateItem != nil {
		return obj.IPv4GatewayTemplateItem
	}

	if obj.IPv4GroupTemplateItem != nil {
		return obj.IPv4GroupTemplateItem
	}

	if obj.IPv4ReservedRangeTemplateItem != nil {
		return obj.IPv4ReservedRangeTemplateItem
	}

	// all schemas are nil
	return nil
}

type NullableGetTemplateItems200Response1DataInner struct {
	value *GetTemplateItems200Response1DataInner
	isSet bool
}

func (v NullableGetTemplateItems200Response1DataInner) Get() *GetTemplateItems200Response1DataInner {
	return v.value
}

func (v *NullableGetTemplateItems200Response1DataInner) Set(val *GetTemplateItems200Response1DataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTemplateItems200Response1DataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTemplateItems200Response1DataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTemplateItems200Response1DataInner(val *GetTemplateItems200Response1DataInner) *NullableGetTemplateItems200Response1DataInner {
	return &NullableGetTemplateItems200Response1DataInner{value: val, isSet: true}
}

func (v NullableGetTemplateItems200Response1DataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTemplateItems200Response1DataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


