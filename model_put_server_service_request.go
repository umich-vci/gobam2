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

// PutServerServiceRequest - struct for PutServerServiceRequest
type PutServerServiceRequest struct {
	AnycastServicePutRequestBody *AnycastServicePutRequestBody
	DHCPActivityServicePutRequestBody *DHCPActivityServicePutRequestBody
	DHCPStatisticsServicePutRequestBody *DHCPStatisticsServicePutRequestBody
	DNSActivityServicePutRequestBody *DNSActivityServicePutRequestBody
	DNSResolverServicePutRequestBody *DNSResolverServicePutRequestBody
	DNSStatisticsServicePutRequestBody *DNSStatisticsServicePutRequestBody
	FirewallServicePutRequestBody *FirewallServicePutRequestBody
	GatewayServicePutRequestBody *GatewayServicePutRequestBody
	InterfacesServicePutRequestBody *InterfacesServicePutRequestBody
	LicenseServicePutRequestBody *LicenseServicePutRequestBody
	NTPServicePutRequestBody *NTPServicePutRequestBody
	SNMPServicePutRequestBody *SNMPServicePutRequestBody
	SSHServicePutRequestBody *SSHServicePutRequestBody
	SyslogServicePutRequestBody *SyslogServicePutRequestBody
}

// AnycastServicePutRequestBodyAsPutServerServiceRequest is a convenience function that returns AnycastServicePutRequestBody wrapped in PutServerServiceRequest
func AnycastServicePutRequestBodyAsPutServerServiceRequest(v *AnycastServicePutRequestBody) PutServerServiceRequest {
	return PutServerServiceRequest{
		AnycastServicePutRequestBody: v,
	}
}

// DHCPActivityServicePutRequestBodyAsPutServerServiceRequest is a convenience function that returns DHCPActivityServicePutRequestBody wrapped in PutServerServiceRequest
func DHCPActivityServicePutRequestBodyAsPutServerServiceRequest(v *DHCPActivityServicePutRequestBody) PutServerServiceRequest {
	return PutServerServiceRequest{
		DHCPActivityServicePutRequestBody: v,
	}
}

// DHCPStatisticsServicePutRequestBodyAsPutServerServiceRequest is a convenience function that returns DHCPStatisticsServicePutRequestBody wrapped in PutServerServiceRequest
func DHCPStatisticsServicePutRequestBodyAsPutServerServiceRequest(v *DHCPStatisticsServicePutRequestBody) PutServerServiceRequest {
	return PutServerServiceRequest{
		DHCPStatisticsServicePutRequestBody: v,
	}
}

// DNSActivityServicePutRequestBodyAsPutServerServiceRequest is a convenience function that returns DNSActivityServicePutRequestBody wrapped in PutServerServiceRequest
func DNSActivityServicePutRequestBodyAsPutServerServiceRequest(v *DNSActivityServicePutRequestBody) PutServerServiceRequest {
	return PutServerServiceRequest{
		DNSActivityServicePutRequestBody: v,
	}
}

// DNSResolverServicePutRequestBodyAsPutServerServiceRequest is a convenience function that returns DNSResolverServicePutRequestBody wrapped in PutServerServiceRequest
func DNSResolverServicePutRequestBodyAsPutServerServiceRequest(v *DNSResolverServicePutRequestBody) PutServerServiceRequest {
	return PutServerServiceRequest{
		DNSResolverServicePutRequestBody: v,
	}
}

// DNSStatisticsServicePutRequestBodyAsPutServerServiceRequest is a convenience function that returns DNSStatisticsServicePutRequestBody wrapped in PutServerServiceRequest
func DNSStatisticsServicePutRequestBodyAsPutServerServiceRequest(v *DNSStatisticsServicePutRequestBody) PutServerServiceRequest {
	return PutServerServiceRequest{
		DNSStatisticsServicePutRequestBody: v,
	}
}

// FirewallServicePutRequestBodyAsPutServerServiceRequest is a convenience function that returns FirewallServicePutRequestBody wrapped in PutServerServiceRequest
func FirewallServicePutRequestBodyAsPutServerServiceRequest(v *FirewallServicePutRequestBody) PutServerServiceRequest {
	return PutServerServiceRequest{
		FirewallServicePutRequestBody: v,
	}
}

// GatewayServicePutRequestBodyAsPutServerServiceRequest is a convenience function that returns GatewayServicePutRequestBody wrapped in PutServerServiceRequest
func GatewayServicePutRequestBodyAsPutServerServiceRequest(v *GatewayServicePutRequestBody) PutServerServiceRequest {
	return PutServerServiceRequest{
		GatewayServicePutRequestBody: v,
	}
}

// InterfacesServicePutRequestBodyAsPutServerServiceRequest is a convenience function that returns InterfacesServicePutRequestBody wrapped in PutServerServiceRequest
func InterfacesServicePutRequestBodyAsPutServerServiceRequest(v *InterfacesServicePutRequestBody) PutServerServiceRequest {
	return PutServerServiceRequest{
		InterfacesServicePutRequestBody: v,
	}
}

// LicenseServicePutRequestBodyAsPutServerServiceRequest is a convenience function that returns LicenseServicePutRequestBody wrapped in PutServerServiceRequest
func LicenseServicePutRequestBodyAsPutServerServiceRequest(v *LicenseServicePutRequestBody) PutServerServiceRequest {
	return PutServerServiceRequest{
		LicenseServicePutRequestBody: v,
	}
}

// NTPServicePutRequestBodyAsPutServerServiceRequest is a convenience function that returns NTPServicePutRequestBody wrapped in PutServerServiceRequest
func NTPServicePutRequestBodyAsPutServerServiceRequest(v *NTPServicePutRequestBody) PutServerServiceRequest {
	return PutServerServiceRequest{
		NTPServicePutRequestBody: v,
	}
}

// SNMPServicePutRequestBodyAsPutServerServiceRequest is a convenience function that returns SNMPServicePutRequestBody wrapped in PutServerServiceRequest
func SNMPServicePutRequestBodyAsPutServerServiceRequest(v *SNMPServicePutRequestBody) PutServerServiceRequest {
	return PutServerServiceRequest{
		SNMPServicePutRequestBody: v,
	}
}

// SSHServicePutRequestBodyAsPutServerServiceRequest is a convenience function that returns SSHServicePutRequestBody wrapped in PutServerServiceRequest
func SSHServicePutRequestBodyAsPutServerServiceRequest(v *SSHServicePutRequestBody) PutServerServiceRequest {
	return PutServerServiceRequest{
		SSHServicePutRequestBody: v,
	}
}

// SyslogServicePutRequestBodyAsPutServerServiceRequest is a convenience function that returns SyslogServicePutRequestBody wrapped in PutServerServiceRequest
func SyslogServicePutRequestBodyAsPutServerServiceRequest(v *SyslogServicePutRequestBody) PutServerServiceRequest {
	return PutServerServiceRequest{
		SyslogServicePutRequestBody: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *PutServerServiceRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AnycastServicePutRequestBody
	err = newStrictDecoder(data).Decode(&dst.AnycastServicePutRequestBody)
	if err == nil {
		jsonAnycastServicePutRequestBody, _ := json.Marshal(dst.AnycastServicePutRequestBody)
		if string(jsonAnycastServicePutRequestBody) == "{}" { // empty struct
			dst.AnycastServicePutRequestBody = nil
		} else {
			if err = validator.Validate(dst.AnycastServicePutRequestBody); err != nil {
				dst.AnycastServicePutRequestBody = nil
			} else {
				match++
			}
		}
	} else {
		dst.AnycastServicePutRequestBody = nil
	}

	// try to unmarshal data into DHCPActivityServicePutRequestBody
	err = newStrictDecoder(data).Decode(&dst.DHCPActivityServicePutRequestBody)
	if err == nil {
		jsonDHCPActivityServicePutRequestBody, _ := json.Marshal(dst.DHCPActivityServicePutRequestBody)
		if string(jsonDHCPActivityServicePutRequestBody) == "{}" { // empty struct
			dst.DHCPActivityServicePutRequestBody = nil
		} else {
			if err = validator.Validate(dst.DHCPActivityServicePutRequestBody); err != nil {
				dst.DHCPActivityServicePutRequestBody = nil
			} else {
				match++
			}
		}
	} else {
		dst.DHCPActivityServicePutRequestBody = nil
	}

	// try to unmarshal data into DHCPStatisticsServicePutRequestBody
	err = newStrictDecoder(data).Decode(&dst.DHCPStatisticsServicePutRequestBody)
	if err == nil {
		jsonDHCPStatisticsServicePutRequestBody, _ := json.Marshal(dst.DHCPStatisticsServicePutRequestBody)
		if string(jsonDHCPStatisticsServicePutRequestBody) == "{}" { // empty struct
			dst.DHCPStatisticsServicePutRequestBody = nil
		} else {
			if err = validator.Validate(dst.DHCPStatisticsServicePutRequestBody); err != nil {
				dst.DHCPStatisticsServicePutRequestBody = nil
			} else {
				match++
			}
		}
	} else {
		dst.DHCPStatisticsServicePutRequestBody = nil
	}

	// try to unmarshal data into DNSActivityServicePutRequestBody
	err = newStrictDecoder(data).Decode(&dst.DNSActivityServicePutRequestBody)
	if err == nil {
		jsonDNSActivityServicePutRequestBody, _ := json.Marshal(dst.DNSActivityServicePutRequestBody)
		if string(jsonDNSActivityServicePutRequestBody) == "{}" { // empty struct
			dst.DNSActivityServicePutRequestBody = nil
		} else {
			if err = validator.Validate(dst.DNSActivityServicePutRequestBody); err != nil {
				dst.DNSActivityServicePutRequestBody = nil
			} else {
				match++
			}
		}
	} else {
		dst.DNSActivityServicePutRequestBody = nil
	}

	// try to unmarshal data into DNSResolverServicePutRequestBody
	err = newStrictDecoder(data).Decode(&dst.DNSResolverServicePutRequestBody)
	if err == nil {
		jsonDNSResolverServicePutRequestBody, _ := json.Marshal(dst.DNSResolverServicePutRequestBody)
		if string(jsonDNSResolverServicePutRequestBody) == "{}" { // empty struct
			dst.DNSResolverServicePutRequestBody = nil
		} else {
			if err = validator.Validate(dst.DNSResolverServicePutRequestBody); err != nil {
				dst.DNSResolverServicePutRequestBody = nil
			} else {
				match++
			}
		}
	} else {
		dst.DNSResolverServicePutRequestBody = nil
	}

	// try to unmarshal data into DNSStatisticsServicePutRequestBody
	err = newStrictDecoder(data).Decode(&dst.DNSStatisticsServicePutRequestBody)
	if err == nil {
		jsonDNSStatisticsServicePutRequestBody, _ := json.Marshal(dst.DNSStatisticsServicePutRequestBody)
		if string(jsonDNSStatisticsServicePutRequestBody) == "{}" { // empty struct
			dst.DNSStatisticsServicePutRequestBody = nil
		} else {
			if err = validator.Validate(dst.DNSStatisticsServicePutRequestBody); err != nil {
				dst.DNSStatisticsServicePutRequestBody = nil
			} else {
				match++
			}
		}
	} else {
		dst.DNSStatisticsServicePutRequestBody = nil
	}

	// try to unmarshal data into FirewallServicePutRequestBody
	err = newStrictDecoder(data).Decode(&dst.FirewallServicePutRequestBody)
	if err == nil {
		jsonFirewallServicePutRequestBody, _ := json.Marshal(dst.FirewallServicePutRequestBody)
		if string(jsonFirewallServicePutRequestBody) == "{}" { // empty struct
			dst.FirewallServicePutRequestBody = nil
		} else {
			if err = validator.Validate(dst.FirewallServicePutRequestBody); err != nil {
				dst.FirewallServicePutRequestBody = nil
			} else {
				match++
			}
		}
	} else {
		dst.FirewallServicePutRequestBody = nil
	}

	// try to unmarshal data into GatewayServicePutRequestBody
	err = newStrictDecoder(data).Decode(&dst.GatewayServicePutRequestBody)
	if err == nil {
		jsonGatewayServicePutRequestBody, _ := json.Marshal(dst.GatewayServicePutRequestBody)
		if string(jsonGatewayServicePutRequestBody) == "{}" { // empty struct
			dst.GatewayServicePutRequestBody = nil
		} else {
			if err = validator.Validate(dst.GatewayServicePutRequestBody); err != nil {
				dst.GatewayServicePutRequestBody = nil
			} else {
				match++
			}
		}
	} else {
		dst.GatewayServicePutRequestBody = nil
	}

	// try to unmarshal data into InterfacesServicePutRequestBody
	err = newStrictDecoder(data).Decode(&dst.InterfacesServicePutRequestBody)
	if err == nil {
		jsonInterfacesServicePutRequestBody, _ := json.Marshal(dst.InterfacesServicePutRequestBody)
		if string(jsonInterfacesServicePutRequestBody) == "{}" { // empty struct
			dst.InterfacesServicePutRequestBody = nil
		} else {
			if err = validator.Validate(dst.InterfacesServicePutRequestBody); err != nil {
				dst.InterfacesServicePutRequestBody = nil
			} else {
				match++
			}
		}
	} else {
		dst.InterfacesServicePutRequestBody = nil
	}

	// try to unmarshal data into LicenseServicePutRequestBody
	err = newStrictDecoder(data).Decode(&dst.LicenseServicePutRequestBody)
	if err == nil {
		jsonLicenseServicePutRequestBody, _ := json.Marshal(dst.LicenseServicePutRequestBody)
		if string(jsonLicenseServicePutRequestBody) == "{}" { // empty struct
			dst.LicenseServicePutRequestBody = nil
		} else {
			if err = validator.Validate(dst.LicenseServicePutRequestBody); err != nil {
				dst.LicenseServicePutRequestBody = nil
			} else {
				match++
			}
		}
	} else {
		dst.LicenseServicePutRequestBody = nil
	}

	// try to unmarshal data into NTPServicePutRequestBody
	err = newStrictDecoder(data).Decode(&dst.NTPServicePutRequestBody)
	if err == nil {
		jsonNTPServicePutRequestBody, _ := json.Marshal(dst.NTPServicePutRequestBody)
		if string(jsonNTPServicePutRequestBody) == "{}" { // empty struct
			dst.NTPServicePutRequestBody = nil
		} else {
			if err = validator.Validate(dst.NTPServicePutRequestBody); err != nil {
				dst.NTPServicePutRequestBody = nil
			} else {
				match++
			}
		}
	} else {
		dst.NTPServicePutRequestBody = nil
	}

	// try to unmarshal data into SNMPServicePutRequestBody
	err = newStrictDecoder(data).Decode(&dst.SNMPServicePutRequestBody)
	if err == nil {
		jsonSNMPServicePutRequestBody, _ := json.Marshal(dst.SNMPServicePutRequestBody)
		if string(jsonSNMPServicePutRequestBody) == "{}" { // empty struct
			dst.SNMPServicePutRequestBody = nil
		} else {
			if err = validator.Validate(dst.SNMPServicePutRequestBody); err != nil {
				dst.SNMPServicePutRequestBody = nil
			} else {
				match++
			}
		}
	} else {
		dst.SNMPServicePutRequestBody = nil
	}

	// try to unmarshal data into SSHServicePutRequestBody
	err = newStrictDecoder(data).Decode(&dst.SSHServicePutRequestBody)
	if err == nil {
		jsonSSHServicePutRequestBody, _ := json.Marshal(dst.SSHServicePutRequestBody)
		if string(jsonSSHServicePutRequestBody) == "{}" { // empty struct
			dst.SSHServicePutRequestBody = nil
		} else {
			if err = validator.Validate(dst.SSHServicePutRequestBody); err != nil {
				dst.SSHServicePutRequestBody = nil
			} else {
				match++
			}
		}
	} else {
		dst.SSHServicePutRequestBody = nil
	}

	// try to unmarshal data into SyslogServicePutRequestBody
	err = newStrictDecoder(data).Decode(&dst.SyslogServicePutRequestBody)
	if err == nil {
		jsonSyslogServicePutRequestBody, _ := json.Marshal(dst.SyslogServicePutRequestBody)
		if string(jsonSyslogServicePutRequestBody) == "{}" { // empty struct
			dst.SyslogServicePutRequestBody = nil
		} else {
			if err = validator.Validate(dst.SyslogServicePutRequestBody); err != nil {
				dst.SyslogServicePutRequestBody = nil
			} else {
				match++
			}
		}
	} else {
		dst.SyslogServicePutRequestBody = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AnycastServicePutRequestBody = nil
		dst.DHCPActivityServicePutRequestBody = nil
		dst.DHCPStatisticsServicePutRequestBody = nil
		dst.DNSActivityServicePutRequestBody = nil
		dst.DNSResolverServicePutRequestBody = nil
		dst.DNSStatisticsServicePutRequestBody = nil
		dst.FirewallServicePutRequestBody = nil
		dst.GatewayServicePutRequestBody = nil
		dst.InterfacesServicePutRequestBody = nil
		dst.LicenseServicePutRequestBody = nil
		dst.NTPServicePutRequestBody = nil
		dst.SNMPServicePutRequestBody = nil
		dst.SSHServicePutRequestBody = nil
		dst.SyslogServicePutRequestBody = nil

		return fmt.Errorf("data matches more than one schema in oneOf(PutServerServiceRequest)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(PutServerServiceRequest)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PutServerServiceRequest) MarshalJSON() ([]byte, error) {
	if src.AnycastServicePutRequestBody != nil {
		return json.Marshal(&src.AnycastServicePutRequestBody)
	}

	if src.DHCPActivityServicePutRequestBody != nil {
		return json.Marshal(&src.DHCPActivityServicePutRequestBody)
	}

	if src.DHCPStatisticsServicePutRequestBody != nil {
		return json.Marshal(&src.DHCPStatisticsServicePutRequestBody)
	}

	if src.DNSActivityServicePutRequestBody != nil {
		return json.Marshal(&src.DNSActivityServicePutRequestBody)
	}

	if src.DNSResolverServicePutRequestBody != nil {
		return json.Marshal(&src.DNSResolverServicePutRequestBody)
	}

	if src.DNSStatisticsServicePutRequestBody != nil {
		return json.Marshal(&src.DNSStatisticsServicePutRequestBody)
	}

	if src.FirewallServicePutRequestBody != nil {
		return json.Marshal(&src.FirewallServicePutRequestBody)
	}

	if src.GatewayServicePutRequestBody != nil {
		return json.Marshal(&src.GatewayServicePutRequestBody)
	}

	if src.InterfacesServicePutRequestBody != nil {
		return json.Marshal(&src.InterfacesServicePutRequestBody)
	}

	if src.LicenseServicePutRequestBody != nil {
		return json.Marshal(&src.LicenseServicePutRequestBody)
	}

	if src.NTPServicePutRequestBody != nil {
		return json.Marshal(&src.NTPServicePutRequestBody)
	}

	if src.SNMPServicePutRequestBody != nil {
		return json.Marshal(&src.SNMPServicePutRequestBody)
	}

	if src.SSHServicePutRequestBody != nil {
		return json.Marshal(&src.SSHServicePutRequestBody)
	}

	if src.SyslogServicePutRequestBody != nil {
		return json.Marshal(&src.SyslogServicePutRequestBody)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PutServerServiceRequest) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.AnycastServicePutRequestBody != nil {
		return obj.AnycastServicePutRequestBody
	}

	if obj.DHCPActivityServicePutRequestBody != nil {
		return obj.DHCPActivityServicePutRequestBody
	}

	if obj.DHCPStatisticsServicePutRequestBody != nil {
		return obj.DHCPStatisticsServicePutRequestBody
	}

	if obj.DNSActivityServicePutRequestBody != nil {
		return obj.DNSActivityServicePutRequestBody
	}

	if obj.DNSResolverServicePutRequestBody != nil {
		return obj.DNSResolverServicePutRequestBody
	}

	if obj.DNSStatisticsServicePutRequestBody != nil {
		return obj.DNSStatisticsServicePutRequestBody
	}

	if obj.FirewallServicePutRequestBody != nil {
		return obj.FirewallServicePutRequestBody
	}

	if obj.GatewayServicePutRequestBody != nil {
		return obj.GatewayServicePutRequestBody
	}

	if obj.InterfacesServicePutRequestBody != nil {
		return obj.InterfacesServicePutRequestBody
	}

	if obj.LicenseServicePutRequestBody != nil {
		return obj.LicenseServicePutRequestBody
	}

	if obj.NTPServicePutRequestBody != nil {
		return obj.NTPServicePutRequestBody
	}

	if obj.SNMPServicePutRequestBody != nil {
		return obj.SNMPServicePutRequestBody
	}

	if obj.SSHServicePutRequestBody != nil {
		return obj.SSHServicePutRequestBody
	}

	if obj.SyslogServicePutRequestBody != nil {
		return obj.SyslogServicePutRequestBody
	}

	// all schemas are nil
	return nil
}

type NullablePutServerServiceRequest struct {
	value *PutServerServiceRequest
	isSet bool
}

func (v NullablePutServerServiceRequest) Get() *PutServerServiceRequest {
	return v.value
}

func (v *NullablePutServerServiceRequest) Set(val *PutServerServiceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePutServerServiceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePutServerServiceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutServerServiceRequest(val *PutServerServiceRequest) *NullablePutServerServiceRequest {
	return &NullablePutServerServiceRequest{value: val, isSet: true}
}

func (v NullablePutServerServiceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutServerServiceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


