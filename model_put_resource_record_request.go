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

// PutResourceRecordRequest - struct for PutResourceRecordRequest
type PutResourceRecordRequest struct {
	AliasRecordPutRequestBody *AliasRecordPutRequestBody
	ExternalHostRecordPutRequestBody *ExternalHostRecordPutRequestBody
	GenericRecordPutRequestBody *GenericRecordPutRequestBody
	HINFORecordPutRequestBody *HINFORecordPutRequestBody
	HostRecordPutRequestBody *HostRecordPutRequestBody
	MXRecordPutRequestBody *MXRecordPutRequestBody
	NAPTRRecordPutRequestBody *NAPTRRecordPutRequestBody
	SRVRecordPutRequestBody *SRVRecordPutRequestBody
	TXTRecordPutRequestBody *TXTRecordPutRequestBody
}

// AliasRecordPutRequestBodyAsPutResourceRecordRequest is a convenience function that returns AliasRecordPutRequestBody wrapped in PutResourceRecordRequest
func AliasRecordPutRequestBodyAsPutResourceRecordRequest(v *AliasRecordPutRequestBody) PutResourceRecordRequest {
	return PutResourceRecordRequest{
		AliasRecordPutRequestBody: v,
	}
}

// ExternalHostRecordPutRequestBodyAsPutResourceRecordRequest is a convenience function that returns ExternalHostRecordPutRequestBody wrapped in PutResourceRecordRequest
func ExternalHostRecordPutRequestBodyAsPutResourceRecordRequest(v *ExternalHostRecordPutRequestBody) PutResourceRecordRequest {
	return PutResourceRecordRequest{
		ExternalHostRecordPutRequestBody: v,
	}
}

// GenericRecordPutRequestBodyAsPutResourceRecordRequest is a convenience function that returns GenericRecordPutRequestBody wrapped in PutResourceRecordRequest
func GenericRecordPutRequestBodyAsPutResourceRecordRequest(v *GenericRecordPutRequestBody) PutResourceRecordRequest {
	return PutResourceRecordRequest{
		GenericRecordPutRequestBody: v,
	}
}

// HINFORecordPutRequestBodyAsPutResourceRecordRequest is a convenience function that returns HINFORecordPutRequestBody wrapped in PutResourceRecordRequest
func HINFORecordPutRequestBodyAsPutResourceRecordRequest(v *HINFORecordPutRequestBody) PutResourceRecordRequest {
	return PutResourceRecordRequest{
		HINFORecordPutRequestBody: v,
	}
}

// HostRecordPutRequestBodyAsPutResourceRecordRequest is a convenience function that returns HostRecordPutRequestBody wrapped in PutResourceRecordRequest
func HostRecordPutRequestBodyAsPutResourceRecordRequest(v *HostRecordPutRequestBody) PutResourceRecordRequest {
	return PutResourceRecordRequest{
		HostRecordPutRequestBody: v,
	}
}

// MXRecordPutRequestBodyAsPutResourceRecordRequest is a convenience function that returns MXRecordPutRequestBody wrapped in PutResourceRecordRequest
func MXRecordPutRequestBodyAsPutResourceRecordRequest(v *MXRecordPutRequestBody) PutResourceRecordRequest {
	return PutResourceRecordRequest{
		MXRecordPutRequestBody: v,
	}
}

// NAPTRRecordPutRequestBodyAsPutResourceRecordRequest is a convenience function that returns NAPTRRecordPutRequestBody wrapped in PutResourceRecordRequest
func NAPTRRecordPutRequestBodyAsPutResourceRecordRequest(v *NAPTRRecordPutRequestBody) PutResourceRecordRequest {
	return PutResourceRecordRequest{
		NAPTRRecordPutRequestBody: v,
	}
}

// SRVRecordPutRequestBodyAsPutResourceRecordRequest is a convenience function that returns SRVRecordPutRequestBody wrapped in PutResourceRecordRequest
func SRVRecordPutRequestBodyAsPutResourceRecordRequest(v *SRVRecordPutRequestBody) PutResourceRecordRequest {
	return PutResourceRecordRequest{
		SRVRecordPutRequestBody: v,
	}
}

// TXTRecordPutRequestBodyAsPutResourceRecordRequest is a convenience function that returns TXTRecordPutRequestBody wrapped in PutResourceRecordRequest
func TXTRecordPutRequestBodyAsPutResourceRecordRequest(v *TXTRecordPutRequestBody) PutResourceRecordRequest {
	return PutResourceRecordRequest{
		TXTRecordPutRequestBody: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *PutResourceRecordRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AliasRecordPutRequestBody
	err = newStrictDecoder(data).Decode(&dst.AliasRecordPutRequestBody)
	if err == nil {
		jsonAliasRecordPutRequestBody, _ := json.Marshal(dst.AliasRecordPutRequestBody)
		if string(jsonAliasRecordPutRequestBody) == "{}" { // empty struct
			dst.AliasRecordPutRequestBody = nil
		} else {
			if err = validator.Validate(dst.AliasRecordPutRequestBody); err != nil {
				dst.AliasRecordPutRequestBody = nil
			} else {
				match++
			}
		}
	} else {
		dst.AliasRecordPutRequestBody = nil
	}

	// try to unmarshal data into ExternalHostRecordPutRequestBody
	err = newStrictDecoder(data).Decode(&dst.ExternalHostRecordPutRequestBody)
	if err == nil {
		jsonExternalHostRecordPutRequestBody, _ := json.Marshal(dst.ExternalHostRecordPutRequestBody)
		if string(jsonExternalHostRecordPutRequestBody) == "{}" { // empty struct
			dst.ExternalHostRecordPutRequestBody = nil
		} else {
			if err = validator.Validate(dst.ExternalHostRecordPutRequestBody); err != nil {
				dst.ExternalHostRecordPutRequestBody = nil
			} else {
				match++
			}
		}
	} else {
		dst.ExternalHostRecordPutRequestBody = nil
	}

	// try to unmarshal data into GenericRecordPutRequestBody
	err = newStrictDecoder(data).Decode(&dst.GenericRecordPutRequestBody)
	if err == nil {
		jsonGenericRecordPutRequestBody, _ := json.Marshal(dst.GenericRecordPutRequestBody)
		if string(jsonGenericRecordPutRequestBody) == "{}" { // empty struct
			dst.GenericRecordPutRequestBody = nil
		} else {
			if err = validator.Validate(dst.GenericRecordPutRequestBody); err != nil {
				dst.GenericRecordPutRequestBody = nil
			} else {
				match++
			}
		}
	} else {
		dst.GenericRecordPutRequestBody = nil
	}

	// try to unmarshal data into HINFORecordPutRequestBody
	err = newStrictDecoder(data).Decode(&dst.HINFORecordPutRequestBody)
	if err == nil {
		jsonHINFORecordPutRequestBody, _ := json.Marshal(dst.HINFORecordPutRequestBody)
		if string(jsonHINFORecordPutRequestBody) == "{}" { // empty struct
			dst.HINFORecordPutRequestBody = nil
		} else {
			if err = validator.Validate(dst.HINFORecordPutRequestBody); err != nil {
				dst.HINFORecordPutRequestBody = nil
			} else {
				match++
			}
		}
	} else {
		dst.HINFORecordPutRequestBody = nil
	}

	// try to unmarshal data into HostRecordPutRequestBody
	err = newStrictDecoder(data).Decode(&dst.HostRecordPutRequestBody)
	if err == nil {
		jsonHostRecordPutRequestBody, _ := json.Marshal(dst.HostRecordPutRequestBody)
		if string(jsonHostRecordPutRequestBody) == "{}" { // empty struct
			dst.HostRecordPutRequestBody = nil
		} else {
			if err = validator.Validate(dst.HostRecordPutRequestBody); err != nil {
				dst.HostRecordPutRequestBody = nil
			} else {
				match++
			}
		}
	} else {
		dst.HostRecordPutRequestBody = nil
	}

	// try to unmarshal data into MXRecordPutRequestBody
	err = newStrictDecoder(data).Decode(&dst.MXRecordPutRequestBody)
	if err == nil {
		jsonMXRecordPutRequestBody, _ := json.Marshal(dst.MXRecordPutRequestBody)
		if string(jsonMXRecordPutRequestBody) == "{}" { // empty struct
			dst.MXRecordPutRequestBody = nil
		} else {
			if err = validator.Validate(dst.MXRecordPutRequestBody); err != nil {
				dst.MXRecordPutRequestBody = nil
			} else {
				match++
			}
		}
	} else {
		dst.MXRecordPutRequestBody = nil
	}

	// try to unmarshal data into NAPTRRecordPutRequestBody
	err = newStrictDecoder(data).Decode(&dst.NAPTRRecordPutRequestBody)
	if err == nil {
		jsonNAPTRRecordPutRequestBody, _ := json.Marshal(dst.NAPTRRecordPutRequestBody)
		if string(jsonNAPTRRecordPutRequestBody) == "{}" { // empty struct
			dst.NAPTRRecordPutRequestBody = nil
		} else {
			if err = validator.Validate(dst.NAPTRRecordPutRequestBody); err != nil {
				dst.NAPTRRecordPutRequestBody = nil
			} else {
				match++
			}
		}
	} else {
		dst.NAPTRRecordPutRequestBody = nil
	}

	// try to unmarshal data into SRVRecordPutRequestBody
	err = newStrictDecoder(data).Decode(&dst.SRVRecordPutRequestBody)
	if err == nil {
		jsonSRVRecordPutRequestBody, _ := json.Marshal(dst.SRVRecordPutRequestBody)
		if string(jsonSRVRecordPutRequestBody) == "{}" { // empty struct
			dst.SRVRecordPutRequestBody = nil
		} else {
			if err = validator.Validate(dst.SRVRecordPutRequestBody); err != nil {
				dst.SRVRecordPutRequestBody = nil
			} else {
				match++
			}
		}
	} else {
		dst.SRVRecordPutRequestBody = nil
	}

	// try to unmarshal data into TXTRecordPutRequestBody
	err = newStrictDecoder(data).Decode(&dst.TXTRecordPutRequestBody)
	if err == nil {
		jsonTXTRecordPutRequestBody, _ := json.Marshal(dst.TXTRecordPutRequestBody)
		if string(jsonTXTRecordPutRequestBody) == "{}" { // empty struct
			dst.TXTRecordPutRequestBody = nil
		} else {
			if err = validator.Validate(dst.TXTRecordPutRequestBody); err != nil {
				dst.TXTRecordPutRequestBody = nil
			} else {
				match++
			}
		}
	} else {
		dst.TXTRecordPutRequestBody = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AliasRecordPutRequestBody = nil
		dst.ExternalHostRecordPutRequestBody = nil
		dst.GenericRecordPutRequestBody = nil
		dst.HINFORecordPutRequestBody = nil
		dst.HostRecordPutRequestBody = nil
		dst.MXRecordPutRequestBody = nil
		dst.NAPTRRecordPutRequestBody = nil
		dst.SRVRecordPutRequestBody = nil
		dst.TXTRecordPutRequestBody = nil

		return fmt.Errorf("data matches more than one schema in oneOf(PutResourceRecordRequest)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(PutResourceRecordRequest)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PutResourceRecordRequest) MarshalJSON() ([]byte, error) {
	if src.AliasRecordPutRequestBody != nil {
		return json.Marshal(&src.AliasRecordPutRequestBody)
	}

	if src.ExternalHostRecordPutRequestBody != nil {
		return json.Marshal(&src.ExternalHostRecordPutRequestBody)
	}

	if src.GenericRecordPutRequestBody != nil {
		return json.Marshal(&src.GenericRecordPutRequestBody)
	}

	if src.HINFORecordPutRequestBody != nil {
		return json.Marshal(&src.HINFORecordPutRequestBody)
	}

	if src.HostRecordPutRequestBody != nil {
		return json.Marshal(&src.HostRecordPutRequestBody)
	}

	if src.MXRecordPutRequestBody != nil {
		return json.Marshal(&src.MXRecordPutRequestBody)
	}

	if src.NAPTRRecordPutRequestBody != nil {
		return json.Marshal(&src.NAPTRRecordPutRequestBody)
	}

	if src.SRVRecordPutRequestBody != nil {
		return json.Marshal(&src.SRVRecordPutRequestBody)
	}

	if src.TXTRecordPutRequestBody != nil {
		return json.Marshal(&src.TXTRecordPutRequestBody)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PutResourceRecordRequest) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.AliasRecordPutRequestBody != nil {
		return obj.AliasRecordPutRequestBody
	}

	if obj.ExternalHostRecordPutRequestBody != nil {
		return obj.ExternalHostRecordPutRequestBody
	}

	if obj.GenericRecordPutRequestBody != nil {
		return obj.GenericRecordPutRequestBody
	}

	if obj.HINFORecordPutRequestBody != nil {
		return obj.HINFORecordPutRequestBody
	}

	if obj.HostRecordPutRequestBody != nil {
		return obj.HostRecordPutRequestBody
	}

	if obj.MXRecordPutRequestBody != nil {
		return obj.MXRecordPutRequestBody
	}

	if obj.NAPTRRecordPutRequestBody != nil {
		return obj.NAPTRRecordPutRequestBody
	}

	if obj.SRVRecordPutRequestBody != nil {
		return obj.SRVRecordPutRequestBody
	}

	if obj.TXTRecordPutRequestBody != nil {
		return obj.TXTRecordPutRequestBody
	}

	// all schemas are nil
	return nil
}

type NullablePutResourceRecordRequest struct {
	value *PutResourceRecordRequest
	isSet bool
}

func (v NullablePutResourceRecordRequest) Get() *PutResourceRecordRequest {
	return v.value
}

func (v *NullablePutResourceRecordRequest) Set(val *PutResourceRecordRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePutResourceRecordRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePutResourceRecordRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutResourceRecordRequest(val *PutResourceRecordRequest) *NullablePutResourceRecordRequest {
	return &NullablePutResourceRecordRequest{value: val, isSet: true}
}

func (v NullablePutResourceRecordRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutResourceRecordRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


