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

// checks if the SSOEnforcementSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SSOEnforcementSettings{}

// SSOEnforcementSettings struct for SSOEnforcementSettings
type SSOEnforcementSettings struct {
	Settings
	// The resource type.
	Type *string `json:"type,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	SamlIdentityProviderEnabled *bool `json:"samlIdentityProviderEnabled,omitempty"`
	NonSsoAuthenticatorCount *int32 `json:"nonSsoAuthenticatorCount,omitempty"`
	NonSsoGroupCount *int32 `json:"nonSsoGroupCount,omitempty"`
	LocalAdminUserCount *int32 `json:"localAdminUserCount,omitempty"`
	LocalNonAdminUserCount *int32 `json:"localNonAdminUserCount,omitempty"`
}

// NewSSOEnforcementSettings instantiates a new SSOEnforcementSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSSOEnforcementSettings() *SSOEnforcementSettings {
	this := SSOEnforcementSettings{}
	return &this
}

// NewSSOEnforcementSettingsWithDefaults instantiates a new SSOEnforcementSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSSOEnforcementSettingsWithDefaults() *SSOEnforcementSettings {
	this := SSOEnforcementSettings{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SSOEnforcementSettings) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSOEnforcementSettings) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SSOEnforcementSettings) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SSOEnforcementSettings) SetType(v string) {
	o.Type = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *SSOEnforcementSettings) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSOEnforcementSettings) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *SSOEnforcementSettings) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *SSOEnforcementSettings) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetSamlIdentityProviderEnabled returns the SamlIdentityProviderEnabled field value if set, zero value otherwise.
func (o *SSOEnforcementSettings) GetSamlIdentityProviderEnabled() bool {
	if o == nil || IsNil(o.SamlIdentityProviderEnabled) {
		var ret bool
		return ret
	}
	return *o.SamlIdentityProviderEnabled
}

// GetSamlIdentityProviderEnabledOk returns a tuple with the SamlIdentityProviderEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSOEnforcementSettings) GetSamlIdentityProviderEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.SamlIdentityProviderEnabled) {
		return nil, false
	}
	return o.SamlIdentityProviderEnabled, true
}

// HasSamlIdentityProviderEnabled returns a boolean if a field has been set.
func (o *SSOEnforcementSettings) HasSamlIdentityProviderEnabled() bool {
	if o != nil && !IsNil(o.SamlIdentityProviderEnabled) {
		return true
	}

	return false
}

// SetSamlIdentityProviderEnabled gets a reference to the given bool and assigns it to the SamlIdentityProviderEnabled field.
func (o *SSOEnforcementSettings) SetSamlIdentityProviderEnabled(v bool) {
	o.SamlIdentityProviderEnabled = &v
}

// GetNonSsoAuthenticatorCount returns the NonSsoAuthenticatorCount field value if set, zero value otherwise.
func (o *SSOEnforcementSettings) GetNonSsoAuthenticatorCount() int32 {
	if o == nil || IsNil(o.NonSsoAuthenticatorCount) {
		var ret int32
		return ret
	}
	return *o.NonSsoAuthenticatorCount
}

// GetNonSsoAuthenticatorCountOk returns a tuple with the NonSsoAuthenticatorCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSOEnforcementSettings) GetNonSsoAuthenticatorCountOk() (*int32, bool) {
	if o == nil || IsNil(o.NonSsoAuthenticatorCount) {
		return nil, false
	}
	return o.NonSsoAuthenticatorCount, true
}

// HasNonSsoAuthenticatorCount returns a boolean if a field has been set.
func (o *SSOEnforcementSettings) HasNonSsoAuthenticatorCount() bool {
	if o != nil && !IsNil(o.NonSsoAuthenticatorCount) {
		return true
	}

	return false
}

// SetNonSsoAuthenticatorCount gets a reference to the given int32 and assigns it to the NonSsoAuthenticatorCount field.
func (o *SSOEnforcementSettings) SetNonSsoAuthenticatorCount(v int32) {
	o.NonSsoAuthenticatorCount = &v
}

// GetNonSsoGroupCount returns the NonSsoGroupCount field value if set, zero value otherwise.
func (o *SSOEnforcementSettings) GetNonSsoGroupCount() int32 {
	if o == nil || IsNil(o.NonSsoGroupCount) {
		var ret int32
		return ret
	}
	return *o.NonSsoGroupCount
}

// GetNonSsoGroupCountOk returns a tuple with the NonSsoGroupCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSOEnforcementSettings) GetNonSsoGroupCountOk() (*int32, bool) {
	if o == nil || IsNil(o.NonSsoGroupCount) {
		return nil, false
	}
	return o.NonSsoGroupCount, true
}

// HasNonSsoGroupCount returns a boolean if a field has been set.
func (o *SSOEnforcementSettings) HasNonSsoGroupCount() bool {
	if o != nil && !IsNil(o.NonSsoGroupCount) {
		return true
	}

	return false
}

// SetNonSsoGroupCount gets a reference to the given int32 and assigns it to the NonSsoGroupCount field.
func (o *SSOEnforcementSettings) SetNonSsoGroupCount(v int32) {
	o.NonSsoGroupCount = &v
}

// GetLocalAdminUserCount returns the LocalAdminUserCount field value if set, zero value otherwise.
func (o *SSOEnforcementSettings) GetLocalAdminUserCount() int32 {
	if o == nil || IsNil(o.LocalAdminUserCount) {
		var ret int32
		return ret
	}
	return *o.LocalAdminUserCount
}

// GetLocalAdminUserCountOk returns a tuple with the LocalAdminUserCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSOEnforcementSettings) GetLocalAdminUserCountOk() (*int32, bool) {
	if o == nil || IsNil(o.LocalAdminUserCount) {
		return nil, false
	}
	return o.LocalAdminUserCount, true
}

// HasLocalAdminUserCount returns a boolean if a field has been set.
func (o *SSOEnforcementSettings) HasLocalAdminUserCount() bool {
	if o != nil && !IsNil(o.LocalAdminUserCount) {
		return true
	}

	return false
}

// SetLocalAdminUserCount gets a reference to the given int32 and assigns it to the LocalAdminUserCount field.
func (o *SSOEnforcementSettings) SetLocalAdminUserCount(v int32) {
	o.LocalAdminUserCount = &v
}

// GetLocalNonAdminUserCount returns the LocalNonAdminUserCount field value if set, zero value otherwise.
func (o *SSOEnforcementSettings) GetLocalNonAdminUserCount() int32 {
	if o == nil || IsNil(o.LocalNonAdminUserCount) {
		var ret int32
		return ret
	}
	return *o.LocalNonAdminUserCount
}

// GetLocalNonAdminUserCountOk returns a tuple with the LocalNonAdminUserCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSOEnforcementSettings) GetLocalNonAdminUserCountOk() (*int32, bool) {
	if o == nil || IsNil(o.LocalNonAdminUserCount) {
		return nil, false
	}
	return o.LocalNonAdminUserCount, true
}

// HasLocalNonAdminUserCount returns a boolean if a field has been set.
func (o *SSOEnforcementSettings) HasLocalNonAdminUserCount() bool {
	if o != nil && !IsNil(o.LocalNonAdminUserCount) {
		return true
	}

	return false
}

// SetLocalNonAdminUserCount gets a reference to the given int32 and assigns it to the LocalNonAdminUserCount field.
func (o *SSOEnforcementSettings) SetLocalNonAdminUserCount(v int32) {
	o.LocalNonAdminUserCount = &v
}

func (o SSOEnforcementSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SSOEnforcementSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.SamlIdentityProviderEnabled) {
		toSerialize["samlIdentityProviderEnabled"] = o.SamlIdentityProviderEnabled
	}
	if !IsNil(o.NonSsoAuthenticatorCount) {
		toSerialize["nonSsoAuthenticatorCount"] = o.NonSsoAuthenticatorCount
	}
	if !IsNil(o.NonSsoGroupCount) {
		toSerialize["nonSsoGroupCount"] = o.NonSsoGroupCount
	}
	if !IsNil(o.LocalAdminUserCount) {
		toSerialize["localAdminUserCount"] = o.LocalAdminUserCount
	}
	if !IsNil(o.LocalNonAdminUserCount) {
		toSerialize["localNonAdminUserCount"] = o.LocalNonAdminUserCount
	}
	return toSerialize, nil
}

type NullableSSOEnforcementSettings struct {
	value *SSOEnforcementSettings
	isSet bool
}

func (v NullableSSOEnforcementSettings) Get() *SSOEnforcementSettings {
	return v.value
}

func (v *NullableSSOEnforcementSettings) Set(val *SSOEnforcementSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableSSOEnforcementSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableSSOEnforcementSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSSOEnforcementSettings(val *SSOEnforcementSettings) *NullableSSOEnforcementSettings {
	return &NullableSSOEnforcementSettings{value: val, isSet: true}
}

func (v NullableSSOEnforcementSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSSOEnforcementSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


