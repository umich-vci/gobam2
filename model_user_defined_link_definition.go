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

// checks if the UserDefinedLinkDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserDefinedLinkDefinition{}

// UserDefinedLinkDefinition struct for UserDefinedLinkDefinition
type UserDefinedLinkDefinition struct {
	// The resource identifier.
	Id *int64 `json:"id,omitempty"`
	// The resource type.
	Type *string `json:"type,omitempty"`
	// The name of the user-defined link definition.
	Name *string `json:"name,omitempty" validate:"regexp=^.*\\\\S+.*$"`
	// The display name of the user-defined link definition.
	DisplayName *string `json:"displayName,omitempty" validate:"regexp=^.*\\\\S+.*$"`
	// The description of the user-defined link definition.
	Description *string `json:"description,omitempty"`
	// The source entity types of the user-defined link definition.
	SourceTypes []string `json:"sourceTypes,omitempty"`
	// The destination entity types of the user-defined link definition.
	DestinationTypes []string `json:"destinationTypes,omitempty"`
}

// NewUserDefinedLinkDefinition instantiates a new UserDefinedLinkDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserDefinedLinkDefinition() *UserDefinedLinkDefinition {
	this := UserDefinedLinkDefinition{}
	return &this
}

// NewUserDefinedLinkDefinitionWithDefaults instantiates a new UserDefinedLinkDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserDefinedLinkDefinitionWithDefaults() *UserDefinedLinkDefinition {
	this := UserDefinedLinkDefinition{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UserDefinedLinkDefinition) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDefinedLinkDefinition) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UserDefinedLinkDefinition) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *UserDefinedLinkDefinition) SetId(v int64) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UserDefinedLinkDefinition) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDefinedLinkDefinition) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UserDefinedLinkDefinition) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *UserDefinedLinkDefinition) SetType(v string) {
	o.Type = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UserDefinedLinkDefinition) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDefinedLinkDefinition) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UserDefinedLinkDefinition) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UserDefinedLinkDefinition) SetName(v string) {
	o.Name = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *UserDefinedLinkDefinition) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDefinedLinkDefinition) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *UserDefinedLinkDefinition) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *UserDefinedLinkDefinition) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UserDefinedLinkDefinition) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDefinedLinkDefinition) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UserDefinedLinkDefinition) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UserDefinedLinkDefinition) SetDescription(v string) {
	o.Description = &v
}

// GetSourceTypes returns the SourceTypes field value if set, zero value otherwise.
func (o *UserDefinedLinkDefinition) GetSourceTypes() []string {
	if o == nil || IsNil(o.SourceTypes) {
		var ret []string
		return ret
	}
	return o.SourceTypes
}

// GetSourceTypesOk returns a tuple with the SourceTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDefinedLinkDefinition) GetSourceTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.SourceTypes) {
		return nil, false
	}
	return o.SourceTypes, true
}

// HasSourceTypes returns a boolean if a field has been set.
func (o *UserDefinedLinkDefinition) HasSourceTypes() bool {
	if o != nil && !IsNil(o.SourceTypes) {
		return true
	}

	return false
}

// SetSourceTypes gets a reference to the given []string and assigns it to the SourceTypes field.
func (o *UserDefinedLinkDefinition) SetSourceTypes(v []string) {
	o.SourceTypes = v
}

// GetDestinationTypes returns the DestinationTypes field value if set, zero value otherwise.
func (o *UserDefinedLinkDefinition) GetDestinationTypes() []string {
	if o == nil || IsNil(o.DestinationTypes) {
		var ret []string
		return ret
	}
	return o.DestinationTypes
}

// GetDestinationTypesOk returns a tuple with the DestinationTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDefinedLinkDefinition) GetDestinationTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.DestinationTypes) {
		return nil, false
	}
	return o.DestinationTypes, true
}

// HasDestinationTypes returns a boolean if a field has been set.
func (o *UserDefinedLinkDefinition) HasDestinationTypes() bool {
	if o != nil && !IsNil(o.DestinationTypes) {
		return true
	}

	return false
}

// SetDestinationTypes gets a reference to the given []string and assigns it to the DestinationTypes field.
func (o *UserDefinedLinkDefinition) SetDestinationTypes(v []string) {
	o.DestinationTypes = v
}

func (o UserDefinedLinkDefinition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserDefinedLinkDefinition) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.SourceTypes) {
		toSerialize["sourceTypes"] = o.SourceTypes
	}
	if !IsNil(o.DestinationTypes) {
		toSerialize["destinationTypes"] = o.DestinationTypes
	}
	return toSerialize, nil
}

type NullableUserDefinedLinkDefinition struct {
	value *UserDefinedLinkDefinition
	isSet bool
}

func (v NullableUserDefinedLinkDefinition) Get() *UserDefinedLinkDefinition {
	return v.value
}

func (v *NullableUserDefinedLinkDefinition) Set(val *UserDefinedLinkDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableUserDefinedLinkDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableUserDefinedLinkDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserDefinedLinkDefinition(val *UserDefinedLinkDefinition) *NullableUserDefinedLinkDefinition {
	return &NullableUserDefinedLinkDefinition{value: val, isSet: true}
}

func (v NullableUserDefinedLinkDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserDefinedLinkDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


