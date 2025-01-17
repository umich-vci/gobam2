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

// checks if the GetDeploymentOptions200ResponseDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDeploymentOptions200ResponseDataInner{}

// GetDeploymentOptions200ResponseDataInner struct for GetDeploymentOptions200ResponseDataInner
type GetDeploymentOptions200ResponseDataInner struct {
	// The resource identifier.
	Id *int64 `json:"id,omitempty"`
	// The resource type.
	Type *string `json:"type,omitempty"`
	// The name of the resource.
	Name *string `json:"name,omitempty" validate:"regexp=^.*\\\\S+.*$"`
	// User-defined fields set for the resource.
	UserDefinedFields *map[string]string `json:"userDefinedFields,omitempty"`
	Configuration *InlinedConfiguration `json:"configuration,omitempty"`
	ServerScope *InlinedServerScope `json:"serverScope,omitempty"`
	Value map[string]interface{} `json:"value,omitempty"`
	// The vendor option code.
	Code *float32 `json:"code,omitempty"`
	Definition *InlinedDHCPVendorOptionDefinition `json:"definition,omitempty"`
	Links *ResourceLinks `json:"_links,omitempty"`
}

// NewGetDeploymentOptions200ResponseDataInner instantiates a new GetDeploymentOptions200ResponseDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDeploymentOptions200ResponseDataInner() *GetDeploymentOptions200ResponseDataInner {
	this := GetDeploymentOptions200ResponseDataInner{}
	return &this
}

// NewGetDeploymentOptions200ResponseDataInnerWithDefaults instantiates a new GetDeploymentOptions200ResponseDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDeploymentOptions200ResponseDataInnerWithDefaults() *GetDeploymentOptions200ResponseDataInner {
	this := GetDeploymentOptions200ResponseDataInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetDeploymentOptions200ResponseDataInner) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeploymentOptions200ResponseDataInner) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetDeploymentOptions200ResponseDataInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *GetDeploymentOptions200ResponseDataInner) SetId(v int64) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetDeploymentOptions200ResponseDataInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeploymentOptions200ResponseDataInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetDeploymentOptions200ResponseDataInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetDeploymentOptions200ResponseDataInner) SetType(v string) {
	o.Type = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetDeploymentOptions200ResponseDataInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeploymentOptions200ResponseDataInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetDeploymentOptions200ResponseDataInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetDeploymentOptions200ResponseDataInner) SetName(v string) {
	o.Name = &v
}

// GetUserDefinedFields returns the UserDefinedFields field value if set, zero value otherwise.
func (o *GetDeploymentOptions200ResponseDataInner) GetUserDefinedFields() map[string]string {
	if o == nil || IsNil(o.UserDefinedFields) {
		var ret map[string]string
		return ret
	}
	return *o.UserDefinedFields
}

// GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeploymentOptions200ResponseDataInner) GetUserDefinedFieldsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.UserDefinedFields) {
		return nil, false
	}
	return o.UserDefinedFields, true
}

// HasUserDefinedFields returns a boolean if a field has been set.
func (o *GetDeploymentOptions200ResponseDataInner) HasUserDefinedFields() bool {
	if o != nil && !IsNil(o.UserDefinedFields) {
		return true
	}

	return false
}

// SetUserDefinedFields gets a reference to the given map[string]string and assigns it to the UserDefinedFields field.
func (o *GetDeploymentOptions200ResponseDataInner) SetUserDefinedFields(v map[string]string) {
	o.UserDefinedFields = &v
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *GetDeploymentOptions200ResponseDataInner) GetConfiguration() InlinedConfiguration {
	if o == nil || IsNil(o.Configuration) {
		var ret InlinedConfiguration
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeploymentOptions200ResponseDataInner) GetConfigurationOk() (*InlinedConfiguration, bool) {
	if o == nil || IsNil(o.Configuration) {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *GetDeploymentOptions200ResponseDataInner) HasConfiguration() bool {
	if o != nil && !IsNil(o.Configuration) {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given InlinedConfiguration and assigns it to the Configuration field.
func (o *GetDeploymentOptions200ResponseDataInner) SetConfiguration(v InlinedConfiguration) {
	o.Configuration = &v
}

// GetServerScope returns the ServerScope field value if set, zero value otherwise.
func (o *GetDeploymentOptions200ResponseDataInner) GetServerScope() InlinedServerScope {
	if o == nil || IsNil(o.ServerScope) {
		var ret InlinedServerScope
		return ret
	}
	return *o.ServerScope
}

// GetServerScopeOk returns a tuple with the ServerScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeploymentOptions200ResponseDataInner) GetServerScopeOk() (*InlinedServerScope, bool) {
	if o == nil || IsNil(o.ServerScope) {
		return nil, false
	}
	return o.ServerScope, true
}

// HasServerScope returns a boolean if a field has been set.
func (o *GetDeploymentOptions200ResponseDataInner) HasServerScope() bool {
	if o != nil && !IsNil(o.ServerScope) {
		return true
	}

	return false
}

// SetServerScope gets a reference to the given InlinedServerScope and assigns it to the ServerScope field.
func (o *GetDeploymentOptions200ResponseDataInner) SetServerScope(v InlinedServerScope) {
	o.ServerScope = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *GetDeploymentOptions200ResponseDataInner) GetValue() map[string]interface{} {
	if o == nil || IsNil(o.Value) {
		var ret map[string]interface{}
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeploymentOptions200ResponseDataInner) GetValueOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Value) {
		return map[string]interface{}{}, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *GetDeploymentOptions200ResponseDataInner) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given map[string]interface{} and assigns it to the Value field.
func (o *GetDeploymentOptions200ResponseDataInner) SetValue(v map[string]interface{}) {
	o.Value = v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *GetDeploymentOptions200ResponseDataInner) GetCode() float32 {
	if o == nil || IsNil(o.Code) {
		var ret float32
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeploymentOptions200ResponseDataInner) GetCodeOk() (*float32, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *GetDeploymentOptions200ResponseDataInner) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given float32 and assigns it to the Code field.
func (o *GetDeploymentOptions200ResponseDataInner) SetCode(v float32) {
	o.Code = &v
}

// GetDefinition returns the Definition field value if set, zero value otherwise.
func (o *GetDeploymentOptions200ResponseDataInner) GetDefinition() InlinedDHCPVendorOptionDefinition {
	if o == nil || IsNil(o.Definition) {
		var ret InlinedDHCPVendorOptionDefinition
		return ret
	}
	return *o.Definition
}

// GetDefinitionOk returns a tuple with the Definition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeploymentOptions200ResponseDataInner) GetDefinitionOk() (*InlinedDHCPVendorOptionDefinition, bool) {
	if o == nil || IsNil(o.Definition) {
		return nil, false
	}
	return o.Definition, true
}

// HasDefinition returns a boolean if a field has been set.
func (o *GetDeploymentOptions200ResponseDataInner) HasDefinition() bool {
	if o != nil && !IsNil(o.Definition) {
		return true
	}

	return false
}

// SetDefinition gets a reference to the given InlinedDHCPVendorOptionDefinition and assigns it to the Definition field.
func (o *GetDeploymentOptions200ResponseDataInner) SetDefinition(v InlinedDHCPVendorOptionDefinition) {
	o.Definition = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GetDeploymentOptions200ResponseDataInner) GetLinks() ResourceLinks {
	if o == nil || IsNil(o.Links) {
		var ret ResourceLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeploymentOptions200ResponseDataInner) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GetDeploymentOptions200ResponseDataInner) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ResourceLinks and assigns it to the Links field.
func (o *GetDeploymentOptions200ResponseDataInner) SetLinks(v ResourceLinks) {
	o.Links = &v
}

func (o GetDeploymentOptions200ResponseDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDeploymentOptions200ResponseDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.UserDefinedFields) {
		toSerialize["userDefinedFields"] = o.UserDefinedFields
	}
	if !IsNil(o.Configuration) {
		toSerialize["configuration"] = o.Configuration
	}
	if !IsNil(o.ServerScope) {
		toSerialize["serverScope"] = o.ServerScope
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Definition) {
		toSerialize["definition"] = o.Definition
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableGetDeploymentOptions200ResponseDataInner struct {
	value *GetDeploymentOptions200ResponseDataInner
	isSet bool
}

func (v NullableGetDeploymentOptions200ResponseDataInner) Get() *GetDeploymentOptions200ResponseDataInner {
	return v.value
}

func (v *NullableGetDeploymentOptions200ResponseDataInner) Set(val *GetDeploymentOptions200ResponseDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDeploymentOptions200ResponseDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDeploymentOptions200ResponseDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDeploymentOptions200ResponseDataInner(val *GetDeploymentOptions200ResponseDataInner) *NullableGetDeploymentOptions200ResponseDataInner {
	return &NullableGetDeploymentOptions200ResponseDataInner{value: val, isSet: true}
}

func (v NullableGetDeploymentOptions200ResponseDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDeploymentOptions200ResponseDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


