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

// checks if the ACLPostRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ACLPostRequestBody{}

// ACLPostRequestBody struct for ACLPostRequestBody
type ACLPostRequestBody struct {
	// The resource identifier.
	Id *int64 `json:"id,omitempty"`
	// The resource type.
	Type *string `json:"type,omitempty"`
	// The name of the resource.
	Name string `json:"name" validate:"regexp=^.*\\\\S+.*$"`
	// User-defined fields set for the resource.
	UserDefinedFields *map[string]string `json:"userDefinedFields,omitempty"`
	Configuration *InlinedConfiguration `json:"configuration,omitempty"`
	// Indicates whether the access control list is predefined by the system (true) or user-created (false).
	Predefined *bool `json:"predefined,omitempty"`
	MatchElements []string `json:"matchElements"`
}

type _ACLPostRequestBody ACLPostRequestBody

// NewACLPostRequestBody instantiates a new ACLPostRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewACLPostRequestBody(name string, matchElements []string) *ACLPostRequestBody {
	this := ACLPostRequestBody{}
	this.Name = name
	this.MatchElements = matchElements
	return &this
}

// NewACLPostRequestBodyWithDefaults instantiates a new ACLPostRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewACLPostRequestBodyWithDefaults() *ACLPostRequestBody {
	this := ACLPostRequestBody{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ACLPostRequestBody) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACLPostRequestBody) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ACLPostRequestBody) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ACLPostRequestBody) SetId(v int64) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ACLPostRequestBody) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACLPostRequestBody) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ACLPostRequestBody) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ACLPostRequestBody) SetType(v string) {
	o.Type = &v
}

// GetName returns the Name field value
func (o *ACLPostRequestBody) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ACLPostRequestBody) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ACLPostRequestBody) SetName(v string) {
	o.Name = v
}

// GetUserDefinedFields returns the UserDefinedFields field value if set, zero value otherwise.
func (o *ACLPostRequestBody) GetUserDefinedFields() map[string]string {
	if o == nil || IsNil(o.UserDefinedFields) {
		var ret map[string]string
		return ret
	}
	return *o.UserDefinedFields
}

// GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACLPostRequestBody) GetUserDefinedFieldsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.UserDefinedFields) {
		return nil, false
	}
	return o.UserDefinedFields, true
}

// HasUserDefinedFields returns a boolean if a field has been set.
func (o *ACLPostRequestBody) HasUserDefinedFields() bool {
	if o != nil && !IsNil(o.UserDefinedFields) {
		return true
	}

	return false
}

// SetUserDefinedFields gets a reference to the given map[string]string and assigns it to the UserDefinedFields field.
func (o *ACLPostRequestBody) SetUserDefinedFields(v map[string]string) {
	o.UserDefinedFields = &v
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *ACLPostRequestBody) GetConfiguration() InlinedConfiguration {
	if o == nil || IsNil(o.Configuration) {
		var ret InlinedConfiguration
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACLPostRequestBody) GetConfigurationOk() (*InlinedConfiguration, bool) {
	if o == nil || IsNil(o.Configuration) {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *ACLPostRequestBody) HasConfiguration() bool {
	if o != nil && !IsNil(o.Configuration) {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given InlinedConfiguration and assigns it to the Configuration field.
func (o *ACLPostRequestBody) SetConfiguration(v InlinedConfiguration) {
	o.Configuration = &v
}

// GetPredefined returns the Predefined field value if set, zero value otherwise.
func (o *ACLPostRequestBody) GetPredefined() bool {
	if o == nil || IsNil(o.Predefined) {
		var ret bool
		return ret
	}
	return *o.Predefined
}

// GetPredefinedOk returns a tuple with the Predefined field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACLPostRequestBody) GetPredefinedOk() (*bool, bool) {
	if o == nil || IsNil(o.Predefined) {
		return nil, false
	}
	return o.Predefined, true
}

// HasPredefined returns a boolean if a field has been set.
func (o *ACLPostRequestBody) HasPredefined() bool {
	if o != nil && !IsNil(o.Predefined) {
		return true
	}

	return false
}

// SetPredefined gets a reference to the given bool and assigns it to the Predefined field.
func (o *ACLPostRequestBody) SetPredefined(v bool) {
	o.Predefined = &v
}

// GetMatchElements returns the MatchElements field value
func (o *ACLPostRequestBody) GetMatchElements() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.MatchElements
}

// GetMatchElementsOk returns a tuple with the MatchElements field value
// and a boolean to check if the value has been set.
func (o *ACLPostRequestBody) GetMatchElementsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MatchElements, true
}

// SetMatchElements sets field value
func (o *ACLPostRequestBody) SetMatchElements(v []string) {
	o.MatchElements = v
}

func (o ACLPostRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ACLPostRequestBody) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Configuration) {
		toSerialize["configuration"] = o.Configuration
	}
	if !IsNil(o.Predefined) {
		toSerialize["predefined"] = o.Predefined
	}
	toSerialize["matchElements"] = o.MatchElements
	return toSerialize, nil
}

func (o *ACLPostRequestBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"matchElements",
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

	varACLPostRequestBody := _ACLPostRequestBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varACLPostRequestBody)

	if err != nil {
		return err
	}

	*o = ACLPostRequestBody(varACLPostRequestBody)

	return err
}

type NullableACLPostRequestBody struct {
	value *ACLPostRequestBody
	isSet bool
}

func (v NullableACLPostRequestBody) Get() *ACLPostRequestBody {
	return v.value
}

func (v *NullableACLPostRequestBody) Set(val *ACLPostRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableACLPostRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableACLPostRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableACLPostRequestBody(val *ACLPostRequestBody) *NullableACLPostRequestBody {
	return &NullableACLPostRequestBody{value: val, isSet: true}
}

func (v NullableACLPostRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableACLPostRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


