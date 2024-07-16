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

// GetDeploymentOptions200Response1DataInner - struct for GetDeploymentOptions200Response1DataInner
type GetDeploymentOptions200Response1DataInner struct {
	DHCPVendorOption *DHCPVendorOption
	DHCPv4ClientOption *DHCPv4ClientOption
	DHCPv4RawOption *DHCPv4RawOption
	DHCPv4ServiceOption *DHCPv4ServiceOption
	DHCPv6ClientOption *DHCPv6ClientOption
	DHCPv6RawOption *DHCPv6RawOption
	DHCPv6ServiceOption *DHCPv6ServiceOption
	DNSOption *DNSOption
	DNSRawOption *DNSRawOption
	StartOfAuthority *StartOfAuthority
}

// DHCPVendorOptionAsGetDeploymentOptions200Response1DataInner is a convenience function that returns DHCPVendorOption wrapped in GetDeploymentOptions200Response1DataInner
func DHCPVendorOptionAsGetDeploymentOptions200Response1DataInner(v *DHCPVendorOption) GetDeploymentOptions200Response1DataInner {
	return GetDeploymentOptions200Response1DataInner{
		DHCPVendorOption: v,
	}
}

// DHCPv4ClientOptionAsGetDeploymentOptions200Response1DataInner is a convenience function that returns DHCPv4ClientOption wrapped in GetDeploymentOptions200Response1DataInner
func DHCPv4ClientOptionAsGetDeploymentOptions200Response1DataInner(v *DHCPv4ClientOption) GetDeploymentOptions200Response1DataInner {
	return GetDeploymentOptions200Response1DataInner{
		DHCPv4ClientOption: v,
	}
}

// DHCPv4RawOptionAsGetDeploymentOptions200Response1DataInner is a convenience function that returns DHCPv4RawOption wrapped in GetDeploymentOptions200Response1DataInner
func DHCPv4RawOptionAsGetDeploymentOptions200Response1DataInner(v *DHCPv4RawOption) GetDeploymentOptions200Response1DataInner {
	return GetDeploymentOptions200Response1DataInner{
		DHCPv4RawOption: v,
	}
}

// DHCPv4ServiceOptionAsGetDeploymentOptions200Response1DataInner is a convenience function that returns DHCPv4ServiceOption wrapped in GetDeploymentOptions200Response1DataInner
func DHCPv4ServiceOptionAsGetDeploymentOptions200Response1DataInner(v *DHCPv4ServiceOption) GetDeploymentOptions200Response1DataInner {
	return GetDeploymentOptions200Response1DataInner{
		DHCPv4ServiceOption: v,
	}
}

// DHCPv6ClientOptionAsGetDeploymentOptions200Response1DataInner is a convenience function that returns DHCPv6ClientOption wrapped in GetDeploymentOptions200Response1DataInner
func DHCPv6ClientOptionAsGetDeploymentOptions200Response1DataInner(v *DHCPv6ClientOption) GetDeploymentOptions200Response1DataInner {
	return GetDeploymentOptions200Response1DataInner{
		DHCPv6ClientOption: v,
	}
}

// DHCPv6RawOptionAsGetDeploymentOptions200Response1DataInner is a convenience function that returns DHCPv6RawOption wrapped in GetDeploymentOptions200Response1DataInner
func DHCPv6RawOptionAsGetDeploymentOptions200Response1DataInner(v *DHCPv6RawOption) GetDeploymentOptions200Response1DataInner {
	return GetDeploymentOptions200Response1DataInner{
		DHCPv6RawOption: v,
	}
}

// DHCPv6ServiceOptionAsGetDeploymentOptions200Response1DataInner is a convenience function that returns DHCPv6ServiceOption wrapped in GetDeploymentOptions200Response1DataInner
func DHCPv6ServiceOptionAsGetDeploymentOptions200Response1DataInner(v *DHCPv6ServiceOption) GetDeploymentOptions200Response1DataInner {
	return GetDeploymentOptions200Response1DataInner{
		DHCPv6ServiceOption: v,
	}
}

// DNSOptionAsGetDeploymentOptions200Response1DataInner is a convenience function that returns DNSOption wrapped in GetDeploymentOptions200Response1DataInner
func DNSOptionAsGetDeploymentOptions200Response1DataInner(v *DNSOption) GetDeploymentOptions200Response1DataInner {
	return GetDeploymentOptions200Response1DataInner{
		DNSOption: v,
	}
}

// DNSRawOptionAsGetDeploymentOptions200Response1DataInner is a convenience function that returns DNSRawOption wrapped in GetDeploymentOptions200Response1DataInner
func DNSRawOptionAsGetDeploymentOptions200Response1DataInner(v *DNSRawOption) GetDeploymentOptions200Response1DataInner {
	return GetDeploymentOptions200Response1DataInner{
		DNSRawOption: v,
	}
}

// StartOfAuthorityAsGetDeploymentOptions200Response1DataInner is a convenience function that returns StartOfAuthority wrapped in GetDeploymentOptions200Response1DataInner
func StartOfAuthorityAsGetDeploymentOptions200Response1DataInner(v *StartOfAuthority) GetDeploymentOptions200Response1DataInner {
	return GetDeploymentOptions200Response1DataInner{
		StartOfAuthority: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetDeploymentOptions200Response1DataInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into DHCPVendorOption
	err = newStrictDecoder(data).Decode(&dst.DHCPVendorOption)
	if err == nil {
		jsonDHCPVendorOption, _ := json.Marshal(dst.DHCPVendorOption)
		if string(jsonDHCPVendorOption) == "{}" { // empty struct
			dst.DHCPVendorOption = nil
		} else {
			if err = validator.Validate(dst.DHCPVendorOption); err != nil {
				dst.DHCPVendorOption = nil
			} else {
				match++
			}
		}
	} else {
		dst.DHCPVendorOption = nil
	}

	// try to unmarshal data into DHCPv4ClientOption
	err = newStrictDecoder(data).Decode(&dst.DHCPv4ClientOption)
	if err == nil {
		jsonDHCPv4ClientOption, _ := json.Marshal(dst.DHCPv4ClientOption)
		if string(jsonDHCPv4ClientOption) == "{}" { // empty struct
			dst.DHCPv4ClientOption = nil
		} else {
			if err = validator.Validate(dst.DHCPv4ClientOption); err != nil {
				dst.DHCPv4ClientOption = nil
			} else {
				match++
			}
		}
	} else {
		dst.DHCPv4ClientOption = nil
	}

	// try to unmarshal data into DHCPv4RawOption
	err = newStrictDecoder(data).Decode(&dst.DHCPv4RawOption)
	if err == nil {
		jsonDHCPv4RawOption, _ := json.Marshal(dst.DHCPv4RawOption)
		if string(jsonDHCPv4RawOption) == "{}" { // empty struct
			dst.DHCPv4RawOption = nil
		} else {
			if err = validator.Validate(dst.DHCPv4RawOption); err != nil {
				dst.DHCPv4RawOption = nil
			} else {
				match++
			}
		}
	} else {
		dst.DHCPv4RawOption = nil
	}

	// try to unmarshal data into DHCPv4ServiceOption
	err = newStrictDecoder(data).Decode(&dst.DHCPv4ServiceOption)
	if err == nil {
		jsonDHCPv4ServiceOption, _ := json.Marshal(dst.DHCPv4ServiceOption)
		if string(jsonDHCPv4ServiceOption) == "{}" { // empty struct
			dst.DHCPv4ServiceOption = nil
		} else {
			if err = validator.Validate(dst.DHCPv4ServiceOption); err != nil {
				dst.DHCPv4ServiceOption = nil
			} else {
				match++
			}
		}
	} else {
		dst.DHCPv4ServiceOption = nil
	}

	// try to unmarshal data into DHCPv6ClientOption
	err = newStrictDecoder(data).Decode(&dst.DHCPv6ClientOption)
	if err == nil {
		jsonDHCPv6ClientOption, _ := json.Marshal(dst.DHCPv6ClientOption)
		if string(jsonDHCPv6ClientOption) == "{}" { // empty struct
			dst.DHCPv6ClientOption = nil
		} else {
			if err = validator.Validate(dst.DHCPv6ClientOption); err != nil {
				dst.DHCPv6ClientOption = nil
			} else {
				match++
			}
		}
	} else {
		dst.DHCPv6ClientOption = nil
	}

	// try to unmarshal data into DHCPv6RawOption
	err = newStrictDecoder(data).Decode(&dst.DHCPv6RawOption)
	if err == nil {
		jsonDHCPv6RawOption, _ := json.Marshal(dst.DHCPv6RawOption)
		if string(jsonDHCPv6RawOption) == "{}" { // empty struct
			dst.DHCPv6RawOption = nil
		} else {
			if err = validator.Validate(dst.DHCPv6RawOption); err != nil {
				dst.DHCPv6RawOption = nil
			} else {
				match++
			}
		}
	} else {
		dst.DHCPv6RawOption = nil
	}

	// try to unmarshal data into DHCPv6ServiceOption
	err = newStrictDecoder(data).Decode(&dst.DHCPv6ServiceOption)
	if err == nil {
		jsonDHCPv6ServiceOption, _ := json.Marshal(dst.DHCPv6ServiceOption)
		if string(jsonDHCPv6ServiceOption) == "{}" { // empty struct
			dst.DHCPv6ServiceOption = nil
		} else {
			if err = validator.Validate(dst.DHCPv6ServiceOption); err != nil {
				dst.DHCPv6ServiceOption = nil
			} else {
				match++
			}
		}
	} else {
		dst.DHCPv6ServiceOption = nil
	}

	// try to unmarshal data into DNSOption
	err = newStrictDecoder(data).Decode(&dst.DNSOption)
	if err == nil {
		jsonDNSOption, _ := json.Marshal(dst.DNSOption)
		if string(jsonDNSOption) == "{}" { // empty struct
			dst.DNSOption = nil
		} else {
			if err = validator.Validate(dst.DNSOption); err != nil {
				dst.DNSOption = nil
			} else {
				match++
			}
		}
	} else {
		dst.DNSOption = nil
	}

	// try to unmarshal data into DNSRawOption
	err = newStrictDecoder(data).Decode(&dst.DNSRawOption)
	if err == nil {
		jsonDNSRawOption, _ := json.Marshal(dst.DNSRawOption)
		if string(jsonDNSRawOption) == "{}" { // empty struct
			dst.DNSRawOption = nil
		} else {
			if err = validator.Validate(dst.DNSRawOption); err != nil {
				dst.DNSRawOption = nil
			} else {
				match++
			}
		}
	} else {
		dst.DNSRawOption = nil
	}

	// try to unmarshal data into StartOfAuthority
	err = newStrictDecoder(data).Decode(&dst.StartOfAuthority)
	if err == nil {
		jsonStartOfAuthority, _ := json.Marshal(dst.StartOfAuthority)
		if string(jsonStartOfAuthority) == "{}" { // empty struct
			dst.StartOfAuthority = nil
		} else {
			if err = validator.Validate(dst.StartOfAuthority); err != nil {
				dst.StartOfAuthority = nil
			} else {
				match++
			}
		}
	} else {
		dst.StartOfAuthority = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.DHCPVendorOption = nil
		dst.DHCPv4ClientOption = nil
		dst.DHCPv4RawOption = nil
		dst.DHCPv4ServiceOption = nil
		dst.DHCPv6ClientOption = nil
		dst.DHCPv6RawOption = nil
		dst.DHCPv6ServiceOption = nil
		dst.DNSOption = nil
		dst.DNSRawOption = nil
		dst.StartOfAuthority = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetDeploymentOptions200Response1DataInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetDeploymentOptions200Response1DataInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetDeploymentOptions200Response1DataInner) MarshalJSON() ([]byte, error) {
	if src.DHCPVendorOption != nil {
		return json.Marshal(&src.DHCPVendorOption)
	}

	if src.DHCPv4ClientOption != nil {
		return json.Marshal(&src.DHCPv4ClientOption)
	}

	if src.DHCPv4RawOption != nil {
		return json.Marshal(&src.DHCPv4RawOption)
	}

	if src.DHCPv4ServiceOption != nil {
		return json.Marshal(&src.DHCPv4ServiceOption)
	}

	if src.DHCPv6ClientOption != nil {
		return json.Marshal(&src.DHCPv6ClientOption)
	}

	if src.DHCPv6RawOption != nil {
		return json.Marshal(&src.DHCPv6RawOption)
	}

	if src.DHCPv6ServiceOption != nil {
		return json.Marshal(&src.DHCPv6ServiceOption)
	}

	if src.DNSOption != nil {
		return json.Marshal(&src.DNSOption)
	}

	if src.DNSRawOption != nil {
		return json.Marshal(&src.DNSRawOption)
	}

	if src.StartOfAuthority != nil {
		return json.Marshal(&src.StartOfAuthority)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetDeploymentOptions200Response1DataInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.DHCPVendorOption != nil {
		return obj.DHCPVendorOption
	}

	if obj.DHCPv4ClientOption != nil {
		return obj.DHCPv4ClientOption
	}

	if obj.DHCPv4RawOption != nil {
		return obj.DHCPv4RawOption
	}

	if obj.DHCPv4ServiceOption != nil {
		return obj.DHCPv4ServiceOption
	}

	if obj.DHCPv6ClientOption != nil {
		return obj.DHCPv6ClientOption
	}

	if obj.DHCPv6RawOption != nil {
		return obj.DHCPv6RawOption
	}

	if obj.DHCPv6ServiceOption != nil {
		return obj.DHCPv6ServiceOption
	}

	if obj.DNSOption != nil {
		return obj.DNSOption
	}

	if obj.DNSRawOption != nil {
		return obj.DNSRawOption
	}

	if obj.StartOfAuthority != nil {
		return obj.StartOfAuthority
	}

	// all schemas are nil
	return nil
}

type NullableGetDeploymentOptions200Response1DataInner struct {
	value *GetDeploymentOptions200Response1DataInner
	isSet bool
}

func (v NullableGetDeploymentOptions200Response1DataInner) Get() *GetDeploymentOptions200Response1DataInner {
	return v.value
}

func (v *NullableGetDeploymentOptions200Response1DataInner) Set(val *GetDeploymentOptions200Response1DataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDeploymentOptions200Response1DataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDeploymentOptions200Response1DataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDeploymentOptions200Response1DataInner(val *GetDeploymentOptions200Response1DataInner) *NullableGetDeploymentOptions200Response1DataInner {
	return &NullableGetDeploymentOptions200Response1DataInner{value: val, isSet: true}
}

func (v NullableGetDeploymentOptions200Response1DataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDeploymentOptions200Response1DataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


