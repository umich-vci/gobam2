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

// checks if the UserGroupPostRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserGroupPostRequestBody{}

// UserGroupPostRequestBody struct for UserGroupPostRequestBody
type UserGroupPostRequestBody struct {
	// The resource identifier.
	Id *int64 `json:"id,omitempty"`
	// The resource type.
	Type *string `json:"type,omitempty"`
	// The name for the user, LDAP, TACACS+, or SSO group.
	Name string `json:"name" validate:"regexp=^.*\\\\S+.*$"`
	// User-defined fields set for the resource.
	UserDefinedFields *map[string]string `json:"userDefinedFields,omitempty"`
	GroupType *string `json:"groupType,omitempty"`
	// The description of the SSO group.
	Description *string `json:"description,omitempty" validate:"regexp=^.*\\\\S+.*$"`
	// Indicates whether the users within the group have administrative user privileges.
	AdministratorPrivilege *bool `json:"administratorPrivilege,omitempty"`
	Authenticator *InlinedAuthenticator `json:"authenticator,omitempty"`
}

type _UserGroupPostRequestBody UserGroupPostRequestBody

// NewUserGroupPostRequestBody instantiates a new UserGroupPostRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserGroupPostRequestBody(name string) *UserGroupPostRequestBody {
	this := UserGroupPostRequestBody{}
	this.Name = name
	return &this
}

// NewUserGroupPostRequestBodyWithDefaults instantiates a new UserGroupPostRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserGroupPostRequestBodyWithDefaults() *UserGroupPostRequestBody {
	this := UserGroupPostRequestBody{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UserGroupPostRequestBody) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserGroupPostRequestBody) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UserGroupPostRequestBody) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *UserGroupPostRequestBody) SetId(v int64) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UserGroupPostRequestBody) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserGroupPostRequestBody) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UserGroupPostRequestBody) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *UserGroupPostRequestBody) SetType(v string) {
	o.Type = &v
}

// GetName returns the Name field value
func (o *UserGroupPostRequestBody) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UserGroupPostRequestBody) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UserGroupPostRequestBody) SetName(v string) {
	o.Name = v
}

// GetUserDefinedFields returns the UserDefinedFields field value if set, zero value otherwise.
func (o *UserGroupPostRequestBody) GetUserDefinedFields() map[string]string {
	if o == nil || IsNil(o.UserDefinedFields) {
		var ret map[string]string
		return ret
	}
	return *o.UserDefinedFields
}

// GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserGroupPostRequestBody) GetUserDefinedFieldsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.UserDefinedFields) {
		return nil, false
	}
	return o.UserDefinedFields, true
}

// HasUserDefinedFields returns a boolean if a field has been set.
func (o *UserGroupPostRequestBody) HasUserDefinedFields() bool {
	if o != nil && !IsNil(o.UserDefinedFields) {
		return true
	}

	return false
}

// SetUserDefinedFields gets a reference to the given map[string]string and assigns it to the UserDefinedFields field.
func (o *UserGroupPostRequestBody) SetUserDefinedFields(v map[string]string) {
	o.UserDefinedFields = &v
}

// GetGroupType returns the GroupType field value if set, zero value otherwise.
func (o *UserGroupPostRequestBody) GetGroupType() string {
	if o == nil || IsNil(o.GroupType) {
		var ret string
		return ret
	}
	return *o.GroupType
}

// GetGroupTypeOk returns a tuple with the GroupType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserGroupPostRequestBody) GetGroupTypeOk() (*string, bool) {
	if o == nil || IsNil(o.GroupType) {
		return nil, false
	}
	return o.GroupType, true
}

// HasGroupType returns a boolean if a field has been set.
func (o *UserGroupPostRequestBody) HasGroupType() bool {
	if o != nil && !IsNil(o.GroupType) {
		return true
	}

	return false
}

// SetGroupType gets a reference to the given string and assigns it to the GroupType field.
func (o *UserGroupPostRequestBody) SetGroupType(v string) {
	o.GroupType = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UserGroupPostRequestBody) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserGroupPostRequestBody) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UserGroupPostRequestBody) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UserGroupPostRequestBody) SetDescription(v string) {
	o.Description = &v
}

// GetAdministratorPrivilege returns the AdministratorPrivilege field value if set, zero value otherwise.
func (o *UserGroupPostRequestBody) GetAdministratorPrivilege() bool {
	if o == nil || IsNil(o.AdministratorPrivilege) {
		var ret bool
		return ret
	}
	return *o.AdministratorPrivilege
}

// GetAdministratorPrivilegeOk returns a tuple with the AdministratorPrivilege field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserGroupPostRequestBody) GetAdministratorPrivilegeOk() (*bool, bool) {
	if o == nil || IsNil(o.AdministratorPrivilege) {
		return nil, false
	}
	return o.AdministratorPrivilege, true
}

// HasAdministratorPrivilege returns a boolean if a field has been set.
func (o *UserGroupPostRequestBody) HasAdministratorPrivilege() bool {
	if o != nil && !IsNil(o.AdministratorPrivilege) {
		return true
	}

	return false
}

// SetAdministratorPrivilege gets a reference to the given bool and assigns it to the AdministratorPrivilege field.
func (o *UserGroupPostRequestBody) SetAdministratorPrivilege(v bool) {
	o.AdministratorPrivilege = &v
}

// GetAuthenticator returns the Authenticator field value if set, zero value otherwise.
func (o *UserGroupPostRequestBody) GetAuthenticator() InlinedAuthenticator {
	if o == nil || IsNil(o.Authenticator) {
		var ret InlinedAuthenticator
		return ret
	}
	return *o.Authenticator
}

// GetAuthenticatorOk returns a tuple with the Authenticator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserGroupPostRequestBody) GetAuthenticatorOk() (*InlinedAuthenticator, bool) {
	if o == nil || IsNil(o.Authenticator) {
		return nil, false
	}
	return o.Authenticator, true
}

// HasAuthenticator returns a boolean if a field has been set.
func (o *UserGroupPostRequestBody) HasAuthenticator() bool {
	if o != nil && !IsNil(o.Authenticator) {
		return true
	}

	return false
}

// SetAuthenticator gets a reference to the given InlinedAuthenticator and assigns it to the Authenticator field.
func (o *UserGroupPostRequestBody) SetAuthenticator(v InlinedAuthenticator) {
	o.Authenticator = &v
}

func (o UserGroupPostRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserGroupPostRequestBody) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.GroupType) {
		toSerialize["groupType"] = o.GroupType
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.AdministratorPrivilege) {
		toSerialize["administratorPrivilege"] = o.AdministratorPrivilege
	}
	if !IsNil(o.Authenticator) {
		toSerialize["authenticator"] = o.Authenticator
	}
	return toSerialize, nil
}

func (o *UserGroupPostRequestBody) UnmarshalJSON(data []byte) (err error) {
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

	varUserGroupPostRequestBody := _UserGroupPostRequestBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserGroupPostRequestBody)

	if err != nil {
		return err
	}

	*o = UserGroupPostRequestBody(varUserGroupPostRequestBody)

	return err
}

type NullableUserGroupPostRequestBody struct {
	value *UserGroupPostRequestBody
	isSet bool
}

func (v NullableUserGroupPostRequestBody) Get() *UserGroupPostRequestBody {
	return v.value
}

func (v *NullableUserGroupPostRequestBody) Set(val *UserGroupPostRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableUserGroupPostRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableUserGroupPostRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserGroupPostRequestBody(val *UserGroupPostRequestBody) *NullableUserGroupPostRequestBody {
	return &NullableUserGroupPostRequestBody{value: val, isSet: true}
}

func (v NullableUserGroupPostRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserGroupPostRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

