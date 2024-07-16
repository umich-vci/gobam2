/*
BlueCat Address Manager 9.5 RESTful v2 API

The **BlueCat Address Manager 9.5 RESTful v2 API** is a new RESTful API for Address Manager that presents Address Manager objects as resources. Each resource has a unique endpoint, and related resources are grouped in collections. To fetch an individual resource, a `GET` request is sent to the resource's endpoint. To fetch all resources in a collection, a `GET` request is sent to the collection's endpoint.  The RESTful v2 API is [hypermedia driven](https://en.wikipedia.org/wiki/HATEOAS) and uses the [HAL](https://en.wikipedia.org/wiki/Hypertext_Application_Language) format for representing links. When navigating through the API, you can use those links to navigate to related resources or child resources of the requested endpoint. The API supports the following media types for most endpoints:  `application/hal+json`, `application/json`, and `text/csv`.  For authentication, the API supports both `Basic` and `Bearer` HTTP authentication schemes.  To get started, create a session and receive credentials for `Basic` authentication by sending a `POST` request to the [/sessions](#/Admin%20Resources/postSession) endpoint, with a message body containing the user's credentials:  ```{\"username\":\"apiuser\", \"password\":\"apipass\"}```  The response will contain an `apiToken` field that can be entered with the username in the Swagger UI **Authorize** dialog. The response will also contain a `basicAuthenticationCredentials` field containing a base64 encoding of the requester's username and API token delimited by a colon. This string can be injected directly into request `Authorization` headers in the following format:  ```Authorization: Basic YXBpOlQ0OExOT3VIRGhDcnVBNEo1bGFES3JuS3hTZC9QK3pjczZXTzBJMDY=```  For full details on API format and authentication methods, refer to the Address Manager RESTful v2 API Guide on the BlueCat Documentation Portal.

API version: 9.5.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gobam2

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the NamingPolicyPostRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NamingPolicyPostRequestBody{}

// NamingPolicyPostRequestBody struct for NamingPolicyPostRequestBody
type NamingPolicyPostRequestBody struct {
	// The resource identifier.
	Id *int64 `json:"id,omitempty"`
	// The resource type.
	Type *string `json:"type,omitempty"`
	// The name of the naming policy.
	Name string `json:"name" validate:"regexp=^.*\\\\S+.*$"`
	// User-defined fields set for the resource.
	UserDefinedFields *map[string]string `json:"userDefinedFields,omitempty"`
	PolicyValues []NamingPolicyValue `json:"policyValues,omitempty"`
	PolicyRestrictions []NamingPolicyRestriction `json:"policyRestrictions,omitempty"`
}

type _NamingPolicyPostRequestBody NamingPolicyPostRequestBody

// NewNamingPolicyPostRequestBody instantiates a new NamingPolicyPostRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNamingPolicyPostRequestBody(name string) *NamingPolicyPostRequestBody {
	this := NamingPolicyPostRequestBody{}
	this.Name = name
	return &this
}

// NewNamingPolicyPostRequestBodyWithDefaults instantiates a new NamingPolicyPostRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNamingPolicyPostRequestBodyWithDefaults() *NamingPolicyPostRequestBody {
	this := NamingPolicyPostRequestBody{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NamingPolicyPostRequestBody) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NamingPolicyPostRequestBody) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NamingPolicyPostRequestBody) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *NamingPolicyPostRequestBody) SetId(v int64) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *NamingPolicyPostRequestBody) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NamingPolicyPostRequestBody) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *NamingPolicyPostRequestBody) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *NamingPolicyPostRequestBody) SetType(v string) {
	o.Type = &v
}

// GetName returns the Name field value
func (o *NamingPolicyPostRequestBody) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NamingPolicyPostRequestBody) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NamingPolicyPostRequestBody) SetName(v string) {
	o.Name = v
}

// GetUserDefinedFields returns the UserDefinedFields field value if set, zero value otherwise.
func (o *NamingPolicyPostRequestBody) GetUserDefinedFields() map[string]string {
	if o == nil || IsNil(o.UserDefinedFields) {
		var ret map[string]string
		return ret
	}
	return *o.UserDefinedFields
}

// GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NamingPolicyPostRequestBody) GetUserDefinedFieldsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.UserDefinedFields) {
		return nil, false
	}
	return o.UserDefinedFields, true
}

// HasUserDefinedFields returns a boolean if a field has been set.
func (o *NamingPolicyPostRequestBody) HasUserDefinedFields() bool {
	if o != nil && !IsNil(o.UserDefinedFields) {
		return true
	}

	return false
}

// SetUserDefinedFields gets a reference to the given map[string]string and assigns it to the UserDefinedFields field.
func (o *NamingPolicyPostRequestBody) SetUserDefinedFields(v map[string]string) {
	o.UserDefinedFields = &v
}

// GetPolicyValues returns the PolicyValues field value if set, zero value otherwise.
func (o *NamingPolicyPostRequestBody) GetPolicyValues() []NamingPolicyValue {
	if o == nil || IsNil(o.PolicyValues) {
		var ret []NamingPolicyValue
		return ret
	}
	return o.PolicyValues
}

// GetPolicyValuesOk returns a tuple with the PolicyValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NamingPolicyPostRequestBody) GetPolicyValuesOk() ([]NamingPolicyValue, bool) {
	if o == nil || IsNil(o.PolicyValues) {
		return nil, false
	}
	return o.PolicyValues, true
}

// HasPolicyValues returns a boolean if a field has been set.
func (o *NamingPolicyPostRequestBody) HasPolicyValues() bool {
	if o != nil && !IsNil(o.PolicyValues) {
		return true
	}

	return false
}

// SetPolicyValues gets a reference to the given []NamingPolicyValue and assigns it to the PolicyValues field.
func (o *NamingPolicyPostRequestBody) SetPolicyValues(v []NamingPolicyValue) {
	o.PolicyValues = v
}

// GetPolicyRestrictions returns the PolicyRestrictions field value if set, zero value otherwise.
func (o *NamingPolicyPostRequestBody) GetPolicyRestrictions() []NamingPolicyRestriction {
	if o == nil || IsNil(o.PolicyRestrictions) {
		var ret []NamingPolicyRestriction
		return ret
	}
	return o.PolicyRestrictions
}

// GetPolicyRestrictionsOk returns a tuple with the PolicyRestrictions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NamingPolicyPostRequestBody) GetPolicyRestrictionsOk() ([]NamingPolicyRestriction, bool) {
	if o == nil || IsNil(o.PolicyRestrictions) {
		return nil, false
	}
	return o.PolicyRestrictions, true
}

// HasPolicyRestrictions returns a boolean if a field has been set.
func (o *NamingPolicyPostRequestBody) HasPolicyRestrictions() bool {
	if o != nil && !IsNil(o.PolicyRestrictions) {
		return true
	}

	return false
}

// SetPolicyRestrictions gets a reference to the given []NamingPolicyRestriction and assigns it to the PolicyRestrictions field.
func (o *NamingPolicyPostRequestBody) SetPolicyRestrictions(v []NamingPolicyRestriction) {
	o.PolicyRestrictions = v
}

func (o NamingPolicyPostRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NamingPolicyPostRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.UserDefinedFields) {
		toSerialize["userDefinedFields"] = o.UserDefinedFields
	}
	if !IsNil(o.PolicyValues) {
		toSerialize["policyValues"] = o.PolicyValues
	}
	if !IsNil(o.PolicyRestrictions) {
		toSerialize["policyRestrictions"] = o.PolicyRestrictions
	}
	return toSerialize, nil
}

func (o *NamingPolicyPostRequestBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varNamingPolicyPostRequestBody := _NamingPolicyPostRequestBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNamingPolicyPostRequestBody)

	if err != nil {
		return err
	}

	*o = NamingPolicyPostRequestBody(varNamingPolicyPostRequestBody)

	return err
}

type NullableNamingPolicyPostRequestBody struct {
	value *NamingPolicyPostRequestBody
	isSet bool
}

func (v NullableNamingPolicyPostRequestBody) Get() *NamingPolicyPostRequestBody {
	return v.value
}

func (v *NullableNamingPolicyPostRequestBody) Set(val *NamingPolicyPostRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableNamingPolicyPostRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableNamingPolicyPostRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNamingPolicyPostRequestBody(val *NamingPolicyPostRequestBody) *NullableNamingPolicyPostRequestBody {
	return &NullableNamingPolicyPostRequestBody{value: val, isSet: true}
}

func (v NullableNamingPolicyPostRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNamingPolicyPostRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


