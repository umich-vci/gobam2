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

// PostAuthenticator201Response1 - struct for PostAuthenticator201Response1
type PostAuthenticator201Response1 struct {
	KerberosAuthenticator *KerberosAuthenticator
	LDAPAuthenticator *LDAPAuthenticator
	OAuthAuthorizationServer *OAuthAuthorizationServer
	RADIUSAuthenticator *RADIUSAuthenticator
	SAMLIdentityProvider *SAMLIdentityProvider
	TACACSPlusAuthenticator *TACACSPlusAuthenticator
}

// KerberosAuthenticatorAsPostAuthenticator201Response1 is a convenience function that returns KerberosAuthenticator wrapped in PostAuthenticator201Response1
func KerberosAuthenticatorAsPostAuthenticator201Response1(v *KerberosAuthenticator) PostAuthenticator201Response1 {
	return PostAuthenticator201Response1{
		KerberosAuthenticator: v,
	}
}

// LDAPAuthenticatorAsPostAuthenticator201Response1 is a convenience function that returns LDAPAuthenticator wrapped in PostAuthenticator201Response1
func LDAPAuthenticatorAsPostAuthenticator201Response1(v *LDAPAuthenticator) PostAuthenticator201Response1 {
	return PostAuthenticator201Response1{
		LDAPAuthenticator: v,
	}
}

// OAuthAuthorizationServerAsPostAuthenticator201Response1 is a convenience function that returns OAuthAuthorizationServer wrapped in PostAuthenticator201Response1
func OAuthAuthorizationServerAsPostAuthenticator201Response1(v *OAuthAuthorizationServer) PostAuthenticator201Response1 {
	return PostAuthenticator201Response1{
		OAuthAuthorizationServer: v,
	}
}

// RADIUSAuthenticatorAsPostAuthenticator201Response1 is a convenience function that returns RADIUSAuthenticator wrapped in PostAuthenticator201Response1
func RADIUSAuthenticatorAsPostAuthenticator201Response1(v *RADIUSAuthenticator) PostAuthenticator201Response1 {
	return PostAuthenticator201Response1{
		RADIUSAuthenticator: v,
	}
}

// SAMLIdentityProviderAsPostAuthenticator201Response1 is a convenience function that returns SAMLIdentityProvider wrapped in PostAuthenticator201Response1
func SAMLIdentityProviderAsPostAuthenticator201Response1(v *SAMLIdentityProvider) PostAuthenticator201Response1 {
	return PostAuthenticator201Response1{
		SAMLIdentityProvider: v,
	}
}

// TACACSPlusAuthenticatorAsPostAuthenticator201Response1 is a convenience function that returns TACACSPlusAuthenticator wrapped in PostAuthenticator201Response1
func TACACSPlusAuthenticatorAsPostAuthenticator201Response1(v *TACACSPlusAuthenticator) PostAuthenticator201Response1 {
	return PostAuthenticator201Response1{
		TACACSPlusAuthenticator: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *PostAuthenticator201Response1) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into KerberosAuthenticator
	err = newStrictDecoder(data).Decode(&dst.KerberosAuthenticator)
	if err == nil {
		jsonKerberosAuthenticator, _ := json.Marshal(dst.KerberosAuthenticator)
		if string(jsonKerberosAuthenticator) == "{}" { // empty struct
			dst.KerberosAuthenticator = nil
		} else {
			if err = validator.Validate(dst.KerberosAuthenticator); err != nil {
				dst.KerberosAuthenticator = nil
			} else {
				match++
			}
		}
	} else {
		dst.KerberosAuthenticator = nil
	}

	// try to unmarshal data into LDAPAuthenticator
	err = newStrictDecoder(data).Decode(&dst.LDAPAuthenticator)
	if err == nil {
		jsonLDAPAuthenticator, _ := json.Marshal(dst.LDAPAuthenticator)
		if string(jsonLDAPAuthenticator) == "{}" { // empty struct
			dst.LDAPAuthenticator = nil
		} else {
			if err = validator.Validate(dst.LDAPAuthenticator); err != nil {
				dst.LDAPAuthenticator = nil
			} else {
				match++
			}
		}
	} else {
		dst.LDAPAuthenticator = nil
	}

	// try to unmarshal data into OAuthAuthorizationServer
	err = newStrictDecoder(data).Decode(&dst.OAuthAuthorizationServer)
	if err == nil {
		jsonOAuthAuthorizationServer, _ := json.Marshal(dst.OAuthAuthorizationServer)
		if string(jsonOAuthAuthorizationServer) == "{}" { // empty struct
			dst.OAuthAuthorizationServer = nil
		} else {
			if err = validator.Validate(dst.OAuthAuthorizationServer); err != nil {
				dst.OAuthAuthorizationServer = nil
			} else {
				match++
			}
		}
	} else {
		dst.OAuthAuthorizationServer = nil
	}

	// try to unmarshal data into RADIUSAuthenticator
	err = newStrictDecoder(data).Decode(&dst.RADIUSAuthenticator)
	if err == nil {
		jsonRADIUSAuthenticator, _ := json.Marshal(dst.RADIUSAuthenticator)
		if string(jsonRADIUSAuthenticator) == "{}" { // empty struct
			dst.RADIUSAuthenticator = nil
		} else {
			if err = validator.Validate(dst.RADIUSAuthenticator); err != nil {
				dst.RADIUSAuthenticator = nil
			} else {
				match++
			}
		}
	} else {
		dst.RADIUSAuthenticator = nil
	}

	// try to unmarshal data into SAMLIdentityProvider
	err = newStrictDecoder(data).Decode(&dst.SAMLIdentityProvider)
	if err == nil {
		jsonSAMLIdentityProvider, _ := json.Marshal(dst.SAMLIdentityProvider)
		if string(jsonSAMLIdentityProvider) == "{}" { // empty struct
			dst.SAMLIdentityProvider = nil
		} else {
			if err = validator.Validate(dst.SAMLIdentityProvider); err != nil {
				dst.SAMLIdentityProvider = nil
			} else {
				match++
			}
		}
	} else {
		dst.SAMLIdentityProvider = nil
	}

	// try to unmarshal data into TACACSPlusAuthenticator
	err = newStrictDecoder(data).Decode(&dst.TACACSPlusAuthenticator)
	if err == nil {
		jsonTACACSPlusAuthenticator, _ := json.Marshal(dst.TACACSPlusAuthenticator)
		if string(jsonTACACSPlusAuthenticator) == "{}" { // empty struct
			dst.TACACSPlusAuthenticator = nil
		} else {
			if err = validator.Validate(dst.TACACSPlusAuthenticator); err != nil {
				dst.TACACSPlusAuthenticator = nil
			} else {
				match++
			}
		}
	} else {
		dst.TACACSPlusAuthenticator = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.KerberosAuthenticator = nil
		dst.LDAPAuthenticator = nil
		dst.OAuthAuthorizationServer = nil
		dst.RADIUSAuthenticator = nil
		dst.SAMLIdentityProvider = nil
		dst.TACACSPlusAuthenticator = nil

		return fmt.Errorf("data matches more than one schema in oneOf(PostAuthenticator201Response1)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(PostAuthenticator201Response1)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PostAuthenticator201Response1) MarshalJSON() ([]byte, error) {
	if src.KerberosAuthenticator != nil {
		return json.Marshal(&src.KerberosAuthenticator)
	}

	if src.LDAPAuthenticator != nil {
		return json.Marshal(&src.LDAPAuthenticator)
	}

	if src.OAuthAuthorizationServer != nil {
		return json.Marshal(&src.OAuthAuthorizationServer)
	}

	if src.RADIUSAuthenticator != nil {
		return json.Marshal(&src.RADIUSAuthenticator)
	}

	if src.SAMLIdentityProvider != nil {
		return json.Marshal(&src.SAMLIdentityProvider)
	}

	if src.TACACSPlusAuthenticator != nil {
		return json.Marshal(&src.TACACSPlusAuthenticator)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PostAuthenticator201Response1) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.KerberosAuthenticator != nil {
		return obj.KerberosAuthenticator
	}

	if obj.LDAPAuthenticator != nil {
		return obj.LDAPAuthenticator
	}

	if obj.OAuthAuthorizationServer != nil {
		return obj.OAuthAuthorizationServer
	}

	if obj.RADIUSAuthenticator != nil {
		return obj.RADIUSAuthenticator
	}

	if obj.SAMLIdentityProvider != nil {
		return obj.SAMLIdentityProvider
	}

	if obj.TACACSPlusAuthenticator != nil {
		return obj.TACACSPlusAuthenticator
	}

	// all schemas are nil
	return nil
}

type NullablePostAuthenticator201Response1 struct {
	value *PostAuthenticator201Response1
	isSet bool
}

func (v NullablePostAuthenticator201Response1) Get() *PostAuthenticator201Response1 {
	return v.value
}

func (v *NullablePostAuthenticator201Response1) Set(val *PostAuthenticator201Response1) {
	v.value = val
	v.isSet = true
}

func (v NullablePostAuthenticator201Response1) IsSet() bool {
	return v.isSet
}

func (v *NullablePostAuthenticator201Response1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostAuthenticator201Response1(val *PostAuthenticator201Response1) *NullablePostAuthenticator201Response1 {
	return &NullablePostAuthenticator201Response1{value: val, isSet: true}
}

func (v NullablePostAuthenticator201Response1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostAuthenticator201Response1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

