/*
BlueCat Address Manager 9.5 RESTful v2 API

The **BlueCat Address Manager 9.5 RESTful v2 API** is a new RESTful API for Address Manager that presents Address Manager objects as resources. Each resource has a unique endpoint, and related resources are grouped in collections. To fetch an individual resource, a `GET` request is sent to the resource's endpoint. To fetch all resources in a collection, a `GET` request is sent to the collection's endpoint.  The RESTful v2 API is [hypermedia driven](https://en.wikipedia.org/wiki/HATEOAS) and uses the [HAL](https://en.wikipedia.org/wiki/Hypertext_Application_Language) format for representing links. When navigating through the API, you can use those links to navigate to related resources or child resources of the requested endpoint. The API supports the following media types for most endpoints:  `application/hal+json`, `application/json`, and `text/csv`.  For authentication, the API supports both `Basic` and `Bearer` HTTP authentication schemes.  To get started, create a session and receive credentials for `Basic` authentication by sending a `POST` request to the [/sessions](#/Admin%20Resources/postSession) endpoint, with a message body containing the user's credentials:  ```{\"username\":\"apiuser\", \"password\":\"apipass\"}```  The response will contain an `apiToken` field that can be entered with the username in the Swagger UI **Authorize** dialog. The response will also contain a `basicAuthenticationCredentials` field containing a base64 encoding of the requester's username and API token delimited by a colon. This string can be injected directly into request `Authorization` headers in the following format:  ```Authorization: Basic YXBpOlQ0OExOT3VIRGhDcnVBNEo1bGFES3JuS3hTZC9QK3pjczZXTzBJMDY=```  For full details on API format and authentication methods, refer to the Address Manager RESTful v2 API Guide on the BlueCat Documentation Portal.

API version: 9.5.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gobam2

import (
	"encoding/json"
)

// checks if the AnycastService type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnycastService{}

// AnycastService struct for AnycastService
type AnycastService struct {
	Service
	// The resource type.
	Type *string `json:"type,omitempty"`
	// Indicates whether the Anycast service is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// Indicates whether the Anycast service configuration has been manually overridden.
	ManualOverrideEnabled *bool `json:"manualOverrideEnabled,omitempty"`
	// The Anycast IP addresses.
	Addresses []string `json:"addresses,omitempty"`
	Routing *RoutingProtocolBean `json:"routing,omitempty"`
}

// NewAnycastService instantiates a new AnycastService object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnycastService() *AnycastService {
	this := AnycastService{}
	return &this
}

// NewAnycastServiceWithDefaults instantiates a new AnycastService object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnycastServiceWithDefaults() *AnycastService {
	this := AnycastService{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AnycastService) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnycastService) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AnycastService) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AnycastService) SetType(v string) {
	o.Type = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *AnycastService) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnycastService) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *AnycastService) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *AnycastService) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetManualOverrideEnabled returns the ManualOverrideEnabled field value if set, zero value otherwise.
func (o *AnycastService) GetManualOverrideEnabled() bool {
	if o == nil || IsNil(o.ManualOverrideEnabled) {
		var ret bool
		return ret
	}
	return *o.ManualOverrideEnabled
}

// GetManualOverrideEnabledOk returns a tuple with the ManualOverrideEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnycastService) GetManualOverrideEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.ManualOverrideEnabled) {
		return nil, false
	}
	return o.ManualOverrideEnabled, true
}

// HasManualOverrideEnabled returns a boolean if a field has been set.
func (o *AnycastService) HasManualOverrideEnabled() bool {
	if o != nil && !IsNil(o.ManualOverrideEnabled) {
		return true
	}

	return false
}

// SetManualOverrideEnabled gets a reference to the given bool and assigns it to the ManualOverrideEnabled field.
func (o *AnycastService) SetManualOverrideEnabled(v bool) {
	o.ManualOverrideEnabled = &v
}

// GetAddresses returns the Addresses field value if set, zero value otherwise.
func (o *AnycastService) GetAddresses() []string {
	if o == nil || IsNil(o.Addresses) {
		var ret []string
		return ret
	}
	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnycastService) GetAddressesOk() ([]string, bool) {
	if o == nil || IsNil(o.Addresses) {
		return nil, false
	}
	return o.Addresses, true
}

// HasAddresses returns a boolean if a field has been set.
func (o *AnycastService) HasAddresses() bool {
	if o != nil && !IsNil(o.Addresses) {
		return true
	}

	return false
}

// SetAddresses gets a reference to the given []string and assigns it to the Addresses field.
func (o *AnycastService) SetAddresses(v []string) {
	o.Addresses = v
}

// GetRouting returns the Routing field value if set, zero value otherwise.
func (o *AnycastService) GetRouting() RoutingProtocolBean {
	if o == nil || IsNil(o.Routing) {
		var ret RoutingProtocolBean
		return ret
	}
	return *o.Routing
}

// GetRoutingOk returns a tuple with the Routing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnycastService) GetRoutingOk() (*RoutingProtocolBean, bool) {
	if o == nil || IsNil(o.Routing) {
		return nil, false
	}
	return o.Routing, true
}

// HasRouting returns a boolean if a field has been set.
func (o *AnycastService) HasRouting() bool {
	if o != nil && !IsNil(o.Routing) {
		return true
	}

	return false
}

// SetRouting gets a reference to the given RoutingProtocolBean and assigns it to the Routing field.
func (o *AnycastService) SetRouting(v RoutingProtocolBean) {
	o.Routing = &v
}

func (o AnycastService) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnycastService) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.ManualOverrideEnabled) {
		toSerialize["manualOverrideEnabled"] = o.ManualOverrideEnabled
	}
	if !IsNil(o.Addresses) {
		toSerialize["addresses"] = o.Addresses
	}
	if !IsNil(o.Routing) {
		toSerialize["routing"] = o.Routing
	}
	return toSerialize, nil
}

type NullableAnycastService struct {
	value *AnycastService
	isSet bool
}

func (v NullableAnycastService) Get() *AnycastService {
	return v.value
}

func (v *NullableAnycastService) Set(val *AnycastService) {
	v.value = val
	v.isSet = true
}

func (v NullableAnycastService) IsSet() bool {
	return v.isSet
}

func (v *NullableAnycastService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnycastService(val *AnycastService) *NullableAnycastService {
	return &NullableAnycastService{value: val, isSet: true}
}

func (v NullableAnycastService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnycastService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


