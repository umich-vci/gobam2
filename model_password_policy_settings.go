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

// checks if the PasswordPolicySettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PasswordPolicySettings{}

// PasswordPolicySettings struct for PasswordPolicySettings
type PasswordPolicySettings struct {
	Settings
	// The resource type.
	Type *string `json:"type,omitempty"`
	// Indicates whether a custom password policy is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// The minimum length that the password must be.
	MinLength *int32 `json:"minLength,omitempty"`
	// The maximum length that the password can be.
	MaxLength *int32 `json:"maxLength,omitempty"`
	// Indicates whether the password must contain at least one digit.
	DigitRequired *bool `json:"digitRequired,omitempty"`
	// Indicates whether the password must contain mixed case letters.
	MixedCaseRequired *bool `json:"mixedCaseRequired,omitempty"`
	// Indicates whether the password must contain at least one special character.
	SpecialCharacterRequired *bool `json:"specialCharacterRequired,omitempty"`
}

// NewPasswordPolicySettings instantiates a new PasswordPolicySettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordPolicySettings() *PasswordPolicySettings {
	this := PasswordPolicySettings{}
	return &this
}

// NewPasswordPolicySettingsWithDefaults instantiates a new PasswordPolicySettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordPolicySettingsWithDefaults() *PasswordPolicySettings {
	this := PasswordPolicySettings{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PasswordPolicySettings) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicySettings) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PasswordPolicySettings) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PasswordPolicySettings) SetType(v string) {
	o.Type = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *PasswordPolicySettings) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicySettings) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *PasswordPolicySettings) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *PasswordPolicySettings) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetMinLength returns the MinLength field value if set, zero value otherwise.
func (o *PasswordPolicySettings) GetMinLength() int32 {
	if o == nil || IsNil(o.MinLength) {
		var ret int32
		return ret
	}
	return *o.MinLength
}

// GetMinLengthOk returns a tuple with the MinLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicySettings) GetMinLengthOk() (*int32, bool) {
	if o == nil || IsNil(o.MinLength) {
		return nil, false
	}
	return o.MinLength, true
}

// HasMinLength returns a boolean if a field has been set.
func (o *PasswordPolicySettings) HasMinLength() bool {
	if o != nil && !IsNil(o.MinLength) {
		return true
	}

	return false
}

// SetMinLength gets a reference to the given int32 and assigns it to the MinLength field.
func (o *PasswordPolicySettings) SetMinLength(v int32) {
	o.MinLength = &v
}

// GetMaxLength returns the MaxLength field value if set, zero value otherwise.
func (o *PasswordPolicySettings) GetMaxLength() int32 {
	if o == nil || IsNil(o.MaxLength) {
		var ret int32
		return ret
	}
	return *o.MaxLength
}

// GetMaxLengthOk returns a tuple with the MaxLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicySettings) GetMaxLengthOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxLength) {
		return nil, false
	}
	return o.MaxLength, true
}

// HasMaxLength returns a boolean if a field has been set.
func (o *PasswordPolicySettings) HasMaxLength() bool {
	if o != nil && !IsNil(o.MaxLength) {
		return true
	}

	return false
}

// SetMaxLength gets a reference to the given int32 and assigns it to the MaxLength field.
func (o *PasswordPolicySettings) SetMaxLength(v int32) {
	o.MaxLength = &v
}

// GetDigitRequired returns the DigitRequired field value if set, zero value otherwise.
func (o *PasswordPolicySettings) GetDigitRequired() bool {
	if o == nil || IsNil(o.DigitRequired) {
		var ret bool
		return ret
	}
	return *o.DigitRequired
}

// GetDigitRequiredOk returns a tuple with the DigitRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicySettings) GetDigitRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.DigitRequired) {
		return nil, false
	}
	return o.DigitRequired, true
}

// HasDigitRequired returns a boolean if a field has been set.
func (o *PasswordPolicySettings) HasDigitRequired() bool {
	if o != nil && !IsNil(o.DigitRequired) {
		return true
	}

	return false
}

// SetDigitRequired gets a reference to the given bool and assigns it to the DigitRequired field.
func (o *PasswordPolicySettings) SetDigitRequired(v bool) {
	o.DigitRequired = &v
}

// GetMixedCaseRequired returns the MixedCaseRequired field value if set, zero value otherwise.
func (o *PasswordPolicySettings) GetMixedCaseRequired() bool {
	if o == nil || IsNil(o.MixedCaseRequired) {
		var ret bool
		return ret
	}
	return *o.MixedCaseRequired
}

// GetMixedCaseRequiredOk returns a tuple with the MixedCaseRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicySettings) GetMixedCaseRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.MixedCaseRequired) {
		return nil, false
	}
	return o.MixedCaseRequired, true
}

// HasMixedCaseRequired returns a boolean if a field has been set.
func (o *PasswordPolicySettings) HasMixedCaseRequired() bool {
	if o != nil && !IsNil(o.MixedCaseRequired) {
		return true
	}

	return false
}

// SetMixedCaseRequired gets a reference to the given bool and assigns it to the MixedCaseRequired field.
func (o *PasswordPolicySettings) SetMixedCaseRequired(v bool) {
	o.MixedCaseRequired = &v
}

// GetSpecialCharacterRequired returns the SpecialCharacterRequired field value if set, zero value otherwise.
func (o *PasswordPolicySettings) GetSpecialCharacterRequired() bool {
	if o == nil || IsNil(o.SpecialCharacterRequired) {
		var ret bool
		return ret
	}
	return *o.SpecialCharacterRequired
}

// GetSpecialCharacterRequiredOk returns a tuple with the SpecialCharacterRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicySettings) GetSpecialCharacterRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.SpecialCharacterRequired) {
		return nil, false
	}
	return o.SpecialCharacterRequired, true
}

// HasSpecialCharacterRequired returns a boolean if a field has been set.
func (o *PasswordPolicySettings) HasSpecialCharacterRequired() bool {
	if o != nil && !IsNil(o.SpecialCharacterRequired) {
		return true
	}

	return false
}

// SetSpecialCharacterRequired gets a reference to the given bool and assigns it to the SpecialCharacterRequired field.
func (o *PasswordPolicySettings) SetSpecialCharacterRequired(v bool) {
	o.SpecialCharacterRequired = &v
}

func (o PasswordPolicySettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PasswordPolicySettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.MinLength) {
		toSerialize["minLength"] = o.MinLength
	}
	if !IsNil(o.MaxLength) {
		toSerialize["maxLength"] = o.MaxLength
	}
	if !IsNil(o.DigitRequired) {
		toSerialize["digitRequired"] = o.DigitRequired
	}
	if !IsNil(o.MixedCaseRequired) {
		toSerialize["mixedCaseRequired"] = o.MixedCaseRequired
	}
	if !IsNil(o.SpecialCharacterRequired) {
		toSerialize["specialCharacterRequired"] = o.SpecialCharacterRequired
	}
	return toSerialize, nil
}

type NullablePasswordPolicySettings struct {
	value *PasswordPolicySettings
	isSet bool
}

func (v NullablePasswordPolicySettings) Get() *PasswordPolicySettings {
	return v.value
}

func (v *NullablePasswordPolicySettings) Set(val *PasswordPolicySettings) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordPolicySettings) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordPolicySettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordPolicySettings(val *PasswordPolicySettings) *NullablePasswordPolicySettings {
	return &NullablePasswordPolicySettings{value: val, isSet: true}
}

func (v NullablePasswordPolicySettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordPolicySettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

