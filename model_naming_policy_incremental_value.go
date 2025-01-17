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

// checks if the NamingPolicyIncrementalValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NamingPolicyIncrementalValue{}

// NamingPolicyIncrementalValue struct for NamingPolicyIncrementalValue
type NamingPolicyIncrementalValue struct {
	NamingPolicyValue
	// The resource type.
	Type *string `json:"type,omitempty"`
	// The incremental counter type. The value can be Counter to make the value a sequential counter or Unique Name to use the value to ensure that the names are unique. When you select Counter, the value starts at a specified value and increments each time the naming policy creates a name. When you select Unique Name, the value increments only to ensure that generated names are unique.
	IncrementalRole *string `json:"incrementalRole,omitempty"`
	// The number system for the incremental value.
	IncrementType *string `json:"incrementType,omitempty"`
	// The padding that is added to the incremental value. Simple padding pads the incremental value with a fixed number of leading zeros. Global padding pads the incremental value with leading zeros to make the entire name generated by the policy a specific length.
	PaddingType *string `json:"paddingType,omitempty"`
	// The starting value for the incremental value.
	SequenceStart *int32 `json:"sequenceStart,omitempty"`
	// The amount by which to increment the value each time it's used.
	SequenceIncrement *int32 `json:"sequenceIncrement,omitempty"`
	// The length of the padding. For simple padding, the value determines how many digits are used for the incremental value. For global padding, the value determines the overall length of the name generated by the naming policy.
	PaddingLength *int32 `json:"paddingLength,omitempty"`
	// Determines whether the naming policy reuses numeric values if they're available.
	MissingValueReuseEnabled *bool `json:"missingValueReuseEnabled,omitempty"`
}

// NewNamingPolicyIncrementalValue instantiates a new NamingPolicyIncrementalValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNamingPolicyIncrementalValue() *NamingPolicyIncrementalValue {
	this := NamingPolicyIncrementalValue{}
	return &this
}

// NewNamingPolicyIncrementalValueWithDefaults instantiates a new NamingPolicyIncrementalValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNamingPolicyIncrementalValueWithDefaults() *NamingPolicyIncrementalValue {
	this := NamingPolicyIncrementalValue{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *NamingPolicyIncrementalValue) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NamingPolicyIncrementalValue) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *NamingPolicyIncrementalValue) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *NamingPolicyIncrementalValue) SetType(v string) {
	o.Type = &v
}

// GetIncrementalRole returns the IncrementalRole field value if set, zero value otherwise.
func (o *NamingPolicyIncrementalValue) GetIncrementalRole() string {
	if o == nil || IsNil(o.IncrementalRole) {
		var ret string
		return ret
	}
	return *o.IncrementalRole
}

// GetIncrementalRoleOk returns a tuple with the IncrementalRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NamingPolicyIncrementalValue) GetIncrementalRoleOk() (*string, bool) {
	if o == nil || IsNil(o.IncrementalRole) {
		return nil, false
	}
	return o.IncrementalRole, true
}

// HasIncrementalRole returns a boolean if a field has been set.
func (o *NamingPolicyIncrementalValue) HasIncrementalRole() bool {
	if o != nil && !IsNil(o.IncrementalRole) {
		return true
	}

	return false
}

// SetIncrementalRole gets a reference to the given string and assigns it to the IncrementalRole field.
func (o *NamingPolicyIncrementalValue) SetIncrementalRole(v string) {
	o.IncrementalRole = &v
}

// GetIncrementType returns the IncrementType field value if set, zero value otherwise.
func (o *NamingPolicyIncrementalValue) GetIncrementType() string {
	if o == nil || IsNil(o.IncrementType) {
		var ret string
		return ret
	}
	return *o.IncrementType
}

// GetIncrementTypeOk returns a tuple with the IncrementType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NamingPolicyIncrementalValue) GetIncrementTypeOk() (*string, bool) {
	if o == nil || IsNil(o.IncrementType) {
		return nil, false
	}
	return o.IncrementType, true
}

// HasIncrementType returns a boolean if a field has been set.
func (o *NamingPolicyIncrementalValue) HasIncrementType() bool {
	if o != nil && !IsNil(o.IncrementType) {
		return true
	}

	return false
}

// SetIncrementType gets a reference to the given string and assigns it to the IncrementType field.
func (o *NamingPolicyIncrementalValue) SetIncrementType(v string) {
	o.IncrementType = &v
}

// GetPaddingType returns the PaddingType field value if set, zero value otherwise.
func (o *NamingPolicyIncrementalValue) GetPaddingType() string {
	if o == nil || IsNil(o.PaddingType) {
		var ret string
		return ret
	}
	return *o.PaddingType
}

// GetPaddingTypeOk returns a tuple with the PaddingType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NamingPolicyIncrementalValue) GetPaddingTypeOk() (*string, bool) {
	if o == nil || IsNil(o.PaddingType) {
		return nil, false
	}
	return o.PaddingType, true
}

// HasPaddingType returns a boolean if a field has been set.
func (o *NamingPolicyIncrementalValue) HasPaddingType() bool {
	if o != nil && !IsNil(o.PaddingType) {
		return true
	}

	return false
}

// SetPaddingType gets a reference to the given string and assigns it to the PaddingType field.
func (o *NamingPolicyIncrementalValue) SetPaddingType(v string) {
	o.PaddingType = &v
}

// GetSequenceStart returns the SequenceStart field value if set, zero value otherwise.
func (o *NamingPolicyIncrementalValue) GetSequenceStart() int32 {
	if o == nil || IsNil(o.SequenceStart) {
		var ret int32
		return ret
	}
	return *o.SequenceStart
}

// GetSequenceStartOk returns a tuple with the SequenceStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NamingPolicyIncrementalValue) GetSequenceStartOk() (*int32, bool) {
	if o == nil || IsNil(o.SequenceStart) {
		return nil, false
	}
	return o.SequenceStart, true
}

// HasSequenceStart returns a boolean if a field has been set.
func (o *NamingPolicyIncrementalValue) HasSequenceStart() bool {
	if o != nil && !IsNil(o.SequenceStart) {
		return true
	}

	return false
}

// SetSequenceStart gets a reference to the given int32 and assigns it to the SequenceStart field.
func (o *NamingPolicyIncrementalValue) SetSequenceStart(v int32) {
	o.SequenceStart = &v
}

// GetSequenceIncrement returns the SequenceIncrement field value if set, zero value otherwise.
func (o *NamingPolicyIncrementalValue) GetSequenceIncrement() int32 {
	if o == nil || IsNil(o.SequenceIncrement) {
		var ret int32
		return ret
	}
	return *o.SequenceIncrement
}

// GetSequenceIncrementOk returns a tuple with the SequenceIncrement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NamingPolicyIncrementalValue) GetSequenceIncrementOk() (*int32, bool) {
	if o == nil || IsNil(o.SequenceIncrement) {
		return nil, false
	}
	return o.SequenceIncrement, true
}

// HasSequenceIncrement returns a boolean if a field has been set.
func (o *NamingPolicyIncrementalValue) HasSequenceIncrement() bool {
	if o != nil && !IsNil(o.SequenceIncrement) {
		return true
	}

	return false
}

// SetSequenceIncrement gets a reference to the given int32 and assigns it to the SequenceIncrement field.
func (o *NamingPolicyIncrementalValue) SetSequenceIncrement(v int32) {
	o.SequenceIncrement = &v
}

// GetPaddingLength returns the PaddingLength field value if set, zero value otherwise.
func (o *NamingPolicyIncrementalValue) GetPaddingLength() int32 {
	if o == nil || IsNil(o.PaddingLength) {
		var ret int32
		return ret
	}
	return *o.PaddingLength
}

// GetPaddingLengthOk returns a tuple with the PaddingLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NamingPolicyIncrementalValue) GetPaddingLengthOk() (*int32, bool) {
	if o == nil || IsNil(o.PaddingLength) {
		return nil, false
	}
	return o.PaddingLength, true
}

// HasPaddingLength returns a boolean if a field has been set.
func (o *NamingPolicyIncrementalValue) HasPaddingLength() bool {
	if o != nil && !IsNil(o.PaddingLength) {
		return true
	}

	return false
}

// SetPaddingLength gets a reference to the given int32 and assigns it to the PaddingLength field.
func (o *NamingPolicyIncrementalValue) SetPaddingLength(v int32) {
	o.PaddingLength = &v
}

// GetMissingValueReuseEnabled returns the MissingValueReuseEnabled field value if set, zero value otherwise.
func (o *NamingPolicyIncrementalValue) GetMissingValueReuseEnabled() bool {
	if o == nil || IsNil(o.MissingValueReuseEnabled) {
		var ret bool
		return ret
	}
	return *o.MissingValueReuseEnabled
}

// GetMissingValueReuseEnabledOk returns a tuple with the MissingValueReuseEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NamingPolicyIncrementalValue) GetMissingValueReuseEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.MissingValueReuseEnabled) {
		return nil, false
	}
	return o.MissingValueReuseEnabled, true
}

// HasMissingValueReuseEnabled returns a boolean if a field has been set.
func (o *NamingPolicyIncrementalValue) HasMissingValueReuseEnabled() bool {
	if o != nil && !IsNil(o.MissingValueReuseEnabled) {
		return true
	}

	return false
}

// SetMissingValueReuseEnabled gets a reference to the given bool and assigns it to the MissingValueReuseEnabled field.
func (o *NamingPolicyIncrementalValue) SetMissingValueReuseEnabled(v bool) {
	o.MissingValueReuseEnabled = &v
}

func (o NamingPolicyIncrementalValue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NamingPolicyIncrementalValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.IncrementalRole) {
		toSerialize["incrementalRole"] = o.IncrementalRole
	}
	if !IsNil(o.IncrementType) {
		toSerialize["incrementType"] = o.IncrementType
	}
	if !IsNil(o.PaddingType) {
		toSerialize["paddingType"] = o.PaddingType
	}
	if !IsNil(o.SequenceStart) {
		toSerialize["sequenceStart"] = o.SequenceStart
	}
	if !IsNil(o.SequenceIncrement) {
		toSerialize["sequenceIncrement"] = o.SequenceIncrement
	}
	if !IsNil(o.PaddingLength) {
		toSerialize["paddingLength"] = o.PaddingLength
	}
	if !IsNil(o.MissingValueReuseEnabled) {
		toSerialize["missingValueReuseEnabled"] = o.MissingValueReuseEnabled
	}
	return toSerialize, nil
}

type NullableNamingPolicyIncrementalValue struct {
	value *NamingPolicyIncrementalValue
	isSet bool
}

func (v NullableNamingPolicyIncrementalValue) Get() *NamingPolicyIncrementalValue {
	return v.value
}

func (v *NullableNamingPolicyIncrementalValue) Set(val *NamingPolicyIncrementalValue) {
	v.value = val
	v.isSet = true
}

func (v NullableNamingPolicyIncrementalValue) IsSet() bool {
	return v.isSet
}

func (v *NullableNamingPolicyIncrementalValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNamingPolicyIncrementalValue(val *NamingPolicyIncrementalValue) *NullableNamingPolicyIncrementalValue {
	return &NullableNamingPolicyIncrementalValue{value: val, isSet: true}
}

func (v NullableNamingPolicyIncrementalValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNamingPolicyIncrementalValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


