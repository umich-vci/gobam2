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

// checks if the UserDefinedFieldDefinitionPostRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserDefinedFieldDefinitionPostRequestBody{}

// UserDefinedFieldDefinitionPostRequestBody struct for UserDefinedFieldDefinitionPostRequestBody
type UserDefinedFieldDefinitionPostRequestBody struct {
	// The resource identifier.
	Id *int64 `json:"id,omitempty"`
	// The resource type.
	Type *string `json:"type,omitempty"`
	// The name of the user defined field.
	Name string `json:"name" validate:"regexp=^.*\\\\S+.*$"`
	// The display name of the user defined field.
	DisplayName string `json:"displayName" validate:"regexp=^.*\\\\S+.*$"`
	// The resource type that the user defined field is applied to.
	ResourceType string `json:"resourceType"`
	// The data type of the user defined field.
	DataType string `json:"dataType"`
	// The default value of the user defined field.
	DefaultValue *string `json:"defaultValue,omitempty"`
	// Indicates whether the user defined field is required or optional for the resource.
	Required *bool `json:"required,omitempty"`
	PredefinedValues []string `json:"predefinedValues,omitempty"`
	ValidationProperties *UserDefinedFieldDefinitionValidationProperties `json:"validationProperties,omitempty"`
}

type _UserDefinedFieldDefinitionPostRequestBody UserDefinedFieldDefinitionPostRequestBody

// NewUserDefinedFieldDefinitionPostRequestBody instantiates a new UserDefinedFieldDefinitionPostRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserDefinedFieldDefinitionPostRequestBody(name string, displayName string, resourceType string, dataType string) *UserDefinedFieldDefinitionPostRequestBody {
	this := UserDefinedFieldDefinitionPostRequestBody{}
	this.Name = name
	this.DisplayName = displayName
	this.ResourceType = resourceType
	this.DataType = dataType
	return &this
}

// NewUserDefinedFieldDefinitionPostRequestBodyWithDefaults instantiates a new UserDefinedFieldDefinitionPostRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserDefinedFieldDefinitionPostRequestBodyWithDefaults() *UserDefinedFieldDefinitionPostRequestBody {
	this := UserDefinedFieldDefinitionPostRequestBody{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UserDefinedFieldDefinitionPostRequestBody) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDefinedFieldDefinitionPostRequestBody) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UserDefinedFieldDefinitionPostRequestBody) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *UserDefinedFieldDefinitionPostRequestBody) SetId(v int64) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UserDefinedFieldDefinitionPostRequestBody) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDefinedFieldDefinitionPostRequestBody) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UserDefinedFieldDefinitionPostRequestBody) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *UserDefinedFieldDefinitionPostRequestBody) SetType(v string) {
	o.Type = &v
}

// GetName returns the Name field value
func (o *UserDefinedFieldDefinitionPostRequestBody) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UserDefinedFieldDefinitionPostRequestBody) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UserDefinedFieldDefinitionPostRequestBody) SetName(v string) {
	o.Name = v
}

// GetDisplayName returns the DisplayName field value
func (o *UserDefinedFieldDefinitionPostRequestBody) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *UserDefinedFieldDefinitionPostRequestBody) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *UserDefinedFieldDefinitionPostRequestBody) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetResourceType returns the ResourceType field value
func (o *UserDefinedFieldDefinitionPostRequestBody) GetResourceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *UserDefinedFieldDefinitionPostRequestBody) GetResourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value
func (o *UserDefinedFieldDefinitionPostRequestBody) SetResourceType(v string) {
	o.ResourceType = v
}

// GetDataType returns the DataType field value
func (o *UserDefinedFieldDefinitionPostRequestBody) GetDataType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DataType
}

// GetDataTypeOk returns a tuple with the DataType field value
// and a boolean to check if the value has been set.
func (o *UserDefinedFieldDefinitionPostRequestBody) GetDataTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataType, true
}

// SetDataType sets field value
func (o *UserDefinedFieldDefinitionPostRequestBody) SetDataType(v string) {
	o.DataType = v
}

// GetDefaultValue returns the DefaultValue field value if set, zero value otherwise.
func (o *UserDefinedFieldDefinitionPostRequestBody) GetDefaultValue() string {
	if o == nil || IsNil(o.DefaultValue) {
		var ret string
		return ret
	}
	return *o.DefaultValue
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDefinedFieldDefinitionPostRequestBody) GetDefaultValueOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultValue) {
		return nil, false
	}
	return o.DefaultValue, true
}

// HasDefaultValue returns a boolean if a field has been set.
func (o *UserDefinedFieldDefinitionPostRequestBody) HasDefaultValue() bool {
	if o != nil && !IsNil(o.DefaultValue) {
		return true
	}

	return false
}

// SetDefaultValue gets a reference to the given string and assigns it to the DefaultValue field.
func (o *UserDefinedFieldDefinitionPostRequestBody) SetDefaultValue(v string) {
	o.DefaultValue = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *UserDefinedFieldDefinitionPostRequestBody) GetRequired() bool {
	if o == nil || IsNil(o.Required) {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDefinedFieldDefinitionPostRequestBody) GetRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.Required) {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *UserDefinedFieldDefinitionPostRequestBody) HasRequired() bool {
	if o != nil && !IsNil(o.Required) {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *UserDefinedFieldDefinitionPostRequestBody) SetRequired(v bool) {
	o.Required = &v
}

// GetPredefinedValues returns the PredefinedValues field value if set, zero value otherwise.
func (o *UserDefinedFieldDefinitionPostRequestBody) GetPredefinedValues() []string {
	if o == nil || IsNil(o.PredefinedValues) {
		var ret []string
		return ret
	}
	return o.PredefinedValues
}

// GetPredefinedValuesOk returns a tuple with the PredefinedValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDefinedFieldDefinitionPostRequestBody) GetPredefinedValuesOk() ([]string, bool) {
	if o == nil || IsNil(o.PredefinedValues) {
		return nil, false
	}
	return o.PredefinedValues, true
}

// HasPredefinedValues returns a boolean if a field has been set.
func (o *UserDefinedFieldDefinitionPostRequestBody) HasPredefinedValues() bool {
	if o != nil && !IsNil(o.PredefinedValues) {
		return true
	}

	return false
}

// SetPredefinedValues gets a reference to the given []string and assigns it to the PredefinedValues field.
func (o *UserDefinedFieldDefinitionPostRequestBody) SetPredefinedValues(v []string) {
	o.PredefinedValues = v
}

// GetValidationProperties returns the ValidationProperties field value if set, zero value otherwise.
func (o *UserDefinedFieldDefinitionPostRequestBody) GetValidationProperties() UserDefinedFieldDefinitionValidationProperties {
	if o == nil || IsNil(o.ValidationProperties) {
		var ret UserDefinedFieldDefinitionValidationProperties
		return ret
	}
	return *o.ValidationProperties
}

// GetValidationPropertiesOk returns a tuple with the ValidationProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDefinedFieldDefinitionPostRequestBody) GetValidationPropertiesOk() (*UserDefinedFieldDefinitionValidationProperties, bool) {
	if o == nil || IsNil(o.ValidationProperties) {
		return nil, false
	}
	return o.ValidationProperties, true
}

// HasValidationProperties returns a boolean if a field has been set.
func (o *UserDefinedFieldDefinitionPostRequestBody) HasValidationProperties() bool {
	if o != nil && !IsNil(o.ValidationProperties) {
		return true
	}

	return false
}

// SetValidationProperties gets a reference to the given UserDefinedFieldDefinitionValidationProperties and assigns it to the ValidationProperties field.
func (o *UserDefinedFieldDefinitionPostRequestBody) SetValidationProperties(v UserDefinedFieldDefinitionValidationProperties) {
	o.ValidationProperties = &v
}

func (o UserDefinedFieldDefinitionPostRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserDefinedFieldDefinitionPostRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	toSerialize["name"] = o.Name
	toSerialize["displayName"] = o.DisplayName
	toSerialize["resourceType"] = o.ResourceType
	toSerialize["dataType"] = o.DataType
	if !IsNil(o.DefaultValue) {
		toSerialize["defaultValue"] = o.DefaultValue
	}
	if !IsNil(o.Required) {
		toSerialize["required"] = o.Required
	}
	if !IsNil(o.PredefinedValues) {
		toSerialize["predefinedValues"] = o.PredefinedValues
	}
	if !IsNil(o.ValidationProperties) {
		toSerialize["validationProperties"] = o.ValidationProperties
	}
	return toSerialize, nil
}

func (o *UserDefinedFieldDefinitionPostRequestBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"displayName",
		"resourceType",
		"dataType",
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

	varUserDefinedFieldDefinitionPostRequestBody := _UserDefinedFieldDefinitionPostRequestBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserDefinedFieldDefinitionPostRequestBody)

	if err != nil {
		return err
	}

	*o = UserDefinedFieldDefinitionPostRequestBody(varUserDefinedFieldDefinitionPostRequestBody)

	return err
}

type NullableUserDefinedFieldDefinitionPostRequestBody struct {
	value *UserDefinedFieldDefinitionPostRequestBody
	isSet bool
}

func (v NullableUserDefinedFieldDefinitionPostRequestBody) Get() *UserDefinedFieldDefinitionPostRequestBody {
	return v.value
}

func (v *NullableUserDefinedFieldDefinitionPostRequestBody) Set(val *UserDefinedFieldDefinitionPostRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableUserDefinedFieldDefinitionPostRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableUserDefinedFieldDefinitionPostRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserDefinedFieldDefinitionPostRequestBody(val *UserDefinedFieldDefinitionPostRequestBody) *NullableUserDefinedFieldDefinitionPostRequestBody {
	return &NullableUserDefinedFieldDefinitionPostRequestBody{value: val, isSet: true}
}

func (v NullableUserDefinedFieldDefinitionPostRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserDefinedFieldDefinitionPostRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


