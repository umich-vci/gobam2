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

// PutSettingRequest - struct for PutSettingRequest
type PutSettingRequest struct {
	AuditDataSettingsPutRequestBody *AuditDataSettingsPutRequestBody
	GlobalSettingsPutRequestBody *GlobalSettingsPutRequestBody
	LogLevelSettingsPutRequestBody *LogLevelSettingsPutRequestBody
	LoginPolicySettingsPutRequestBody *LoginPolicySettingsPutRequestBody
	MonitoringSettingsPutRequestBody *MonitoringSettingsPutRequestBody
	PasswordPolicySettingsPutRequestBody *PasswordPolicySettingsPutRequestBody
	SAMLServiceProviderSettingsPutRequestBody *SAMLServiceProviderSettingsPutRequestBody
	SSOEnforcementSettingsPutRequestBody *SSOEnforcementSettingsPutRequestBody
	SystemSettingsPutRequestBody *SystemSettingsPutRequestBody
}

// AuditDataSettingsPutRequestBodyAsPutSettingRequest is a convenience function that returns AuditDataSettingsPutRequestBody wrapped in PutSettingRequest
func AuditDataSettingsPutRequestBodyAsPutSettingRequest(v *AuditDataSettingsPutRequestBody) PutSettingRequest {
	return PutSettingRequest{
		AuditDataSettingsPutRequestBody: v,
	}
}

// GlobalSettingsPutRequestBodyAsPutSettingRequest is a convenience function that returns GlobalSettingsPutRequestBody wrapped in PutSettingRequest
func GlobalSettingsPutRequestBodyAsPutSettingRequest(v *GlobalSettingsPutRequestBody) PutSettingRequest {
	return PutSettingRequest{
		GlobalSettingsPutRequestBody: v,
	}
}

// LogLevelSettingsPutRequestBodyAsPutSettingRequest is a convenience function that returns LogLevelSettingsPutRequestBody wrapped in PutSettingRequest
func LogLevelSettingsPutRequestBodyAsPutSettingRequest(v *LogLevelSettingsPutRequestBody) PutSettingRequest {
	return PutSettingRequest{
		LogLevelSettingsPutRequestBody: v,
	}
}

// LoginPolicySettingsPutRequestBodyAsPutSettingRequest is a convenience function that returns LoginPolicySettingsPutRequestBody wrapped in PutSettingRequest
func LoginPolicySettingsPutRequestBodyAsPutSettingRequest(v *LoginPolicySettingsPutRequestBody) PutSettingRequest {
	return PutSettingRequest{
		LoginPolicySettingsPutRequestBody: v,
	}
}

// MonitoringSettingsPutRequestBodyAsPutSettingRequest is a convenience function that returns MonitoringSettingsPutRequestBody wrapped in PutSettingRequest
func MonitoringSettingsPutRequestBodyAsPutSettingRequest(v *MonitoringSettingsPutRequestBody) PutSettingRequest {
	return PutSettingRequest{
		MonitoringSettingsPutRequestBody: v,
	}
}

// PasswordPolicySettingsPutRequestBodyAsPutSettingRequest is a convenience function that returns PasswordPolicySettingsPutRequestBody wrapped in PutSettingRequest
func PasswordPolicySettingsPutRequestBodyAsPutSettingRequest(v *PasswordPolicySettingsPutRequestBody) PutSettingRequest {
	return PutSettingRequest{
		PasswordPolicySettingsPutRequestBody: v,
	}
}

// SAMLServiceProviderSettingsPutRequestBodyAsPutSettingRequest is a convenience function that returns SAMLServiceProviderSettingsPutRequestBody wrapped in PutSettingRequest
func SAMLServiceProviderSettingsPutRequestBodyAsPutSettingRequest(v *SAMLServiceProviderSettingsPutRequestBody) PutSettingRequest {
	return PutSettingRequest{
		SAMLServiceProviderSettingsPutRequestBody: v,
	}
}

// SSOEnforcementSettingsPutRequestBodyAsPutSettingRequest is a convenience function that returns SSOEnforcementSettingsPutRequestBody wrapped in PutSettingRequest
func SSOEnforcementSettingsPutRequestBodyAsPutSettingRequest(v *SSOEnforcementSettingsPutRequestBody) PutSettingRequest {
	return PutSettingRequest{
		SSOEnforcementSettingsPutRequestBody: v,
	}
}

// SystemSettingsPutRequestBodyAsPutSettingRequest is a convenience function that returns SystemSettingsPutRequestBody wrapped in PutSettingRequest
func SystemSettingsPutRequestBodyAsPutSettingRequest(v *SystemSettingsPutRequestBody) PutSettingRequest {
	return PutSettingRequest{
		SystemSettingsPutRequestBody: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *PutSettingRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AuditDataSettingsPutRequestBody
	err = newStrictDecoder(data).Decode(&dst.AuditDataSettingsPutRequestBody)
	if err == nil {
		jsonAuditDataSettingsPutRequestBody, _ := json.Marshal(dst.AuditDataSettingsPutRequestBody)
		if string(jsonAuditDataSettingsPutRequestBody) == "{}" { // empty struct
			dst.AuditDataSettingsPutRequestBody = nil
		} else {
			if err = validator.Validate(dst.AuditDataSettingsPutRequestBody); err != nil {
				dst.AuditDataSettingsPutRequestBody = nil
			} else {
				match++
			}
		}
	} else {
		dst.AuditDataSettingsPutRequestBody = nil
	}

	// try to unmarshal data into GlobalSettingsPutRequestBody
	err = newStrictDecoder(data).Decode(&dst.GlobalSettingsPutRequestBody)
	if err == nil {
		jsonGlobalSettingsPutRequestBody, _ := json.Marshal(dst.GlobalSettingsPutRequestBody)
		if string(jsonGlobalSettingsPutRequestBody) == "{}" { // empty struct
			dst.GlobalSettingsPutRequestBody = nil
		} else {
			if err = validator.Validate(dst.GlobalSettingsPutRequestBody); err != nil {
				dst.GlobalSettingsPutRequestBody = nil
			} else {
				match++
			}
		}
	} else {
		dst.GlobalSettingsPutRequestBody = nil
	}

	// try to unmarshal data into LogLevelSettingsPutRequestBody
	err = newStrictDecoder(data).Decode(&dst.LogLevelSettingsPutRequestBody)
	if err == nil {
		jsonLogLevelSettingsPutRequestBody, _ := json.Marshal(dst.LogLevelSettingsPutRequestBody)
		if string(jsonLogLevelSettingsPutRequestBody) == "{}" { // empty struct
			dst.LogLevelSettingsPutRequestBody = nil
		} else {
			if err = validator.Validate(dst.LogLevelSettingsPutRequestBody); err != nil {
				dst.LogLevelSettingsPutRequestBody = nil
			} else {
				match++
			}
		}
	} else {
		dst.LogLevelSettingsPutRequestBody = nil
	}

	// try to unmarshal data into LoginPolicySettingsPutRequestBody
	err = newStrictDecoder(data).Decode(&dst.LoginPolicySettingsPutRequestBody)
	if err == nil {
		jsonLoginPolicySettingsPutRequestBody, _ := json.Marshal(dst.LoginPolicySettingsPutRequestBody)
		if string(jsonLoginPolicySettingsPutRequestBody) == "{}" { // empty struct
			dst.LoginPolicySettingsPutRequestBody = nil
		} else {
			if err = validator.Validate(dst.LoginPolicySettingsPutRequestBody); err != nil {
				dst.LoginPolicySettingsPutRequestBody = nil
			} else {
				match++
			}
		}
	} else {
		dst.LoginPolicySettingsPutRequestBody = nil
	}

	// try to unmarshal data into MonitoringSettingsPutRequestBody
	err = newStrictDecoder(data).Decode(&dst.MonitoringSettingsPutRequestBody)
	if err == nil {
		jsonMonitoringSettingsPutRequestBody, _ := json.Marshal(dst.MonitoringSettingsPutRequestBody)
		if string(jsonMonitoringSettingsPutRequestBody) == "{}" { // empty struct
			dst.MonitoringSettingsPutRequestBody = nil
		} else {
			if err = validator.Validate(dst.MonitoringSettingsPutRequestBody); err != nil {
				dst.MonitoringSettingsPutRequestBody = nil
			} else {
				match++
			}
		}
	} else {
		dst.MonitoringSettingsPutRequestBody = nil
	}

	// try to unmarshal data into PasswordPolicySettingsPutRequestBody
	err = newStrictDecoder(data).Decode(&dst.PasswordPolicySettingsPutRequestBody)
	if err == nil {
		jsonPasswordPolicySettingsPutRequestBody, _ := json.Marshal(dst.PasswordPolicySettingsPutRequestBody)
		if string(jsonPasswordPolicySettingsPutRequestBody) == "{}" { // empty struct
			dst.PasswordPolicySettingsPutRequestBody = nil
		} else {
			if err = validator.Validate(dst.PasswordPolicySettingsPutRequestBody); err != nil {
				dst.PasswordPolicySettingsPutRequestBody = nil
			} else {
				match++
			}
		}
	} else {
		dst.PasswordPolicySettingsPutRequestBody = nil
	}

	// try to unmarshal data into SAMLServiceProviderSettingsPutRequestBody
	err = newStrictDecoder(data).Decode(&dst.SAMLServiceProviderSettingsPutRequestBody)
	if err == nil {
		jsonSAMLServiceProviderSettingsPutRequestBody, _ := json.Marshal(dst.SAMLServiceProviderSettingsPutRequestBody)
		if string(jsonSAMLServiceProviderSettingsPutRequestBody) == "{}" { // empty struct
			dst.SAMLServiceProviderSettingsPutRequestBody = nil
		} else {
			if err = validator.Validate(dst.SAMLServiceProviderSettingsPutRequestBody); err != nil {
				dst.SAMLServiceProviderSettingsPutRequestBody = nil
			} else {
				match++
			}
		}
	} else {
		dst.SAMLServiceProviderSettingsPutRequestBody = nil
	}

	// try to unmarshal data into SSOEnforcementSettingsPutRequestBody
	err = newStrictDecoder(data).Decode(&dst.SSOEnforcementSettingsPutRequestBody)
	if err == nil {
		jsonSSOEnforcementSettingsPutRequestBody, _ := json.Marshal(dst.SSOEnforcementSettingsPutRequestBody)
		if string(jsonSSOEnforcementSettingsPutRequestBody) == "{}" { // empty struct
			dst.SSOEnforcementSettingsPutRequestBody = nil
		} else {
			if err = validator.Validate(dst.SSOEnforcementSettingsPutRequestBody); err != nil {
				dst.SSOEnforcementSettingsPutRequestBody = nil
			} else {
				match++
			}
		}
	} else {
		dst.SSOEnforcementSettingsPutRequestBody = nil
	}

	// try to unmarshal data into SystemSettingsPutRequestBody
	err = newStrictDecoder(data).Decode(&dst.SystemSettingsPutRequestBody)
	if err == nil {
		jsonSystemSettingsPutRequestBody, _ := json.Marshal(dst.SystemSettingsPutRequestBody)
		if string(jsonSystemSettingsPutRequestBody) == "{}" { // empty struct
			dst.SystemSettingsPutRequestBody = nil
		} else {
			if err = validator.Validate(dst.SystemSettingsPutRequestBody); err != nil {
				dst.SystemSettingsPutRequestBody = nil
			} else {
				match++
			}
		}
	} else {
		dst.SystemSettingsPutRequestBody = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AuditDataSettingsPutRequestBody = nil
		dst.GlobalSettingsPutRequestBody = nil
		dst.LogLevelSettingsPutRequestBody = nil
		dst.LoginPolicySettingsPutRequestBody = nil
		dst.MonitoringSettingsPutRequestBody = nil
		dst.PasswordPolicySettingsPutRequestBody = nil
		dst.SAMLServiceProviderSettingsPutRequestBody = nil
		dst.SSOEnforcementSettingsPutRequestBody = nil
		dst.SystemSettingsPutRequestBody = nil

		return fmt.Errorf("data matches more than one schema in oneOf(PutSettingRequest)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(PutSettingRequest)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PutSettingRequest) MarshalJSON() ([]byte, error) {
	if src.AuditDataSettingsPutRequestBody != nil {
		return json.Marshal(&src.AuditDataSettingsPutRequestBody)
	}

	if src.GlobalSettingsPutRequestBody != nil {
		return json.Marshal(&src.GlobalSettingsPutRequestBody)
	}

	if src.LogLevelSettingsPutRequestBody != nil {
		return json.Marshal(&src.LogLevelSettingsPutRequestBody)
	}

	if src.LoginPolicySettingsPutRequestBody != nil {
		return json.Marshal(&src.LoginPolicySettingsPutRequestBody)
	}

	if src.MonitoringSettingsPutRequestBody != nil {
		return json.Marshal(&src.MonitoringSettingsPutRequestBody)
	}

	if src.PasswordPolicySettingsPutRequestBody != nil {
		return json.Marshal(&src.PasswordPolicySettingsPutRequestBody)
	}

	if src.SAMLServiceProviderSettingsPutRequestBody != nil {
		return json.Marshal(&src.SAMLServiceProviderSettingsPutRequestBody)
	}

	if src.SSOEnforcementSettingsPutRequestBody != nil {
		return json.Marshal(&src.SSOEnforcementSettingsPutRequestBody)
	}

	if src.SystemSettingsPutRequestBody != nil {
		return json.Marshal(&src.SystemSettingsPutRequestBody)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PutSettingRequest) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.AuditDataSettingsPutRequestBody != nil {
		return obj.AuditDataSettingsPutRequestBody
	}

	if obj.GlobalSettingsPutRequestBody != nil {
		return obj.GlobalSettingsPutRequestBody
	}

	if obj.LogLevelSettingsPutRequestBody != nil {
		return obj.LogLevelSettingsPutRequestBody
	}

	if obj.LoginPolicySettingsPutRequestBody != nil {
		return obj.LoginPolicySettingsPutRequestBody
	}

	if obj.MonitoringSettingsPutRequestBody != nil {
		return obj.MonitoringSettingsPutRequestBody
	}

	if obj.PasswordPolicySettingsPutRequestBody != nil {
		return obj.PasswordPolicySettingsPutRequestBody
	}

	if obj.SAMLServiceProviderSettingsPutRequestBody != nil {
		return obj.SAMLServiceProviderSettingsPutRequestBody
	}

	if obj.SSOEnforcementSettingsPutRequestBody != nil {
		return obj.SSOEnforcementSettingsPutRequestBody
	}

	if obj.SystemSettingsPutRequestBody != nil {
		return obj.SystemSettingsPutRequestBody
	}

	// all schemas are nil
	return nil
}

type NullablePutSettingRequest struct {
	value *PutSettingRequest
	isSet bool
}

func (v NullablePutSettingRequest) Get() *PutSettingRequest {
	return v.value
}

func (v *NullablePutSettingRequest) Set(val *PutSettingRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePutSettingRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePutSettingRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutSettingRequest(val *PutSettingRequest) *NullablePutSettingRequest {
	return &NullablePutSettingRequest{value: val, isSet: true}
}

func (v NullablePutSettingRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutSettingRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


