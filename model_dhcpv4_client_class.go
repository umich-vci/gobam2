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

// checks if the DHCPv4ClientClass type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DHCPv4ClientClass{}

// DHCPv4ClientClass struct for DHCPv4ClientClass
type DHCPv4ClientClass struct {
	// The resource identifier.
	Id *int64 `json:"id,omitempty"`
	// The resource type.
	Type *string `json:"type,omitempty"`
	// The name of the DHCP client class.
	Name *string `json:"name,omitempty" validate:"regexp=^.*\\\\S+.*$"`
	// User-defined fields set for the resource.
	UserDefinedFields *map[string]string `json:"userDefinedFields,omitempty"`
	Configuration *InlinedConfiguration `json:"configuration,omitempty"`
	// The description for the DHCP client class.
	Description *string `json:"description,omitempty"`
	// The DHCP option to match.
	Option *string `json:"option,omitempty" validate:"regexp=^.*\\\\S+.*$"`
	// The offset that determines where to start the check for matching values. The match check is applied offset bytes from the beginning of the identifier. Offset is not configurable for custom match statements.
	MatchOffset *int32 `json:"matchOffset,omitempty"`
	// The length determines the portion of the identifier where the check for matching values will be applied. The match check is applied offset bytes from the beginning of the identifier, continuing for length bytes. Length is not configurable for custom match statements.
	MatchLength *int32 `json:"matchLength,omitempty"`
	// The data expression for custom match statements, or boolean expression for custom match if statements. The expression must match the supported syntax for the ISC's DHCP daemon.
	MatchExpression *string `json:"matchExpression,omitempty" validate:"regexp=^.*\\\\S+.*$"`
}

// NewDHCPv4ClientClass instantiates a new DHCPv4ClientClass object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDHCPv4ClientClass() *DHCPv4ClientClass {
	this := DHCPv4ClientClass{}
	return &this
}

// NewDHCPv4ClientClassWithDefaults instantiates a new DHCPv4ClientClass object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDHCPv4ClientClassWithDefaults() *DHCPv4ClientClass {
	this := DHCPv4ClientClass{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DHCPv4ClientClass) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DHCPv4ClientClass) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DHCPv4ClientClass) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *DHCPv4ClientClass) SetId(v int64) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DHCPv4ClientClass) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DHCPv4ClientClass) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DHCPv4ClientClass) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DHCPv4ClientClass) SetType(v string) {
	o.Type = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DHCPv4ClientClass) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DHCPv4ClientClass) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DHCPv4ClientClass) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DHCPv4ClientClass) SetName(v string) {
	o.Name = &v
}

// GetUserDefinedFields returns the UserDefinedFields field value if set, zero value otherwise.
func (o *DHCPv4ClientClass) GetUserDefinedFields() map[string]string {
	if o == nil || IsNil(o.UserDefinedFields) {
		var ret map[string]string
		return ret
	}
	return *o.UserDefinedFields
}

// GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DHCPv4ClientClass) GetUserDefinedFieldsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.UserDefinedFields) {
		return nil, false
	}
	return o.UserDefinedFields, true
}

// HasUserDefinedFields returns a boolean if a field has been set.
func (o *DHCPv4ClientClass) HasUserDefinedFields() bool {
	if o != nil && !IsNil(o.UserDefinedFields) {
		return true
	}

	return false
}

// SetUserDefinedFields gets a reference to the given map[string]string and assigns it to the UserDefinedFields field.
func (o *DHCPv4ClientClass) SetUserDefinedFields(v map[string]string) {
	o.UserDefinedFields = &v
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *DHCPv4ClientClass) GetConfiguration() InlinedConfiguration {
	if o == nil || IsNil(o.Configuration) {
		var ret InlinedConfiguration
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DHCPv4ClientClass) GetConfigurationOk() (*InlinedConfiguration, bool) {
	if o == nil || IsNil(o.Configuration) {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *DHCPv4ClientClass) HasConfiguration() bool {
	if o != nil && !IsNil(o.Configuration) {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given InlinedConfiguration and assigns it to the Configuration field.
func (o *DHCPv4ClientClass) SetConfiguration(v InlinedConfiguration) {
	o.Configuration = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DHCPv4ClientClass) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DHCPv4ClientClass) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DHCPv4ClientClass) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DHCPv4ClientClass) SetDescription(v string) {
	o.Description = &v
}

// GetOption returns the Option field value if set, zero value otherwise.
func (o *DHCPv4ClientClass) GetOption() string {
	if o == nil || IsNil(o.Option) {
		var ret string
		return ret
	}
	return *o.Option
}

// GetOptionOk returns a tuple with the Option field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DHCPv4ClientClass) GetOptionOk() (*string, bool) {
	if o == nil || IsNil(o.Option) {
		return nil, false
	}
	return o.Option, true
}

// HasOption returns a boolean if a field has been set.
func (o *DHCPv4ClientClass) HasOption() bool {
	if o != nil && !IsNil(o.Option) {
		return true
	}

	return false
}

// SetOption gets a reference to the given string and assigns it to the Option field.
func (o *DHCPv4ClientClass) SetOption(v string) {
	o.Option = &v
}

// GetMatchOffset returns the MatchOffset field value if set, zero value otherwise.
func (o *DHCPv4ClientClass) GetMatchOffset() int32 {
	if o == nil || IsNil(o.MatchOffset) {
		var ret int32
		return ret
	}
	return *o.MatchOffset
}

// GetMatchOffsetOk returns a tuple with the MatchOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DHCPv4ClientClass) GetMatchOffsetOk() (*int32, bool) {
	if o == nil || IsNil(o.MatchOffset) {
		return nil, false
	}
	return o.MatchOffset, true
}

// HasMatchOffset returns a boolean if a field has been set.
func (o *DHCPv4ClientClass) HasMatchOffset() bool {
	if o != nil && !IsNil(o.MatchOffset) {
		return true
	}

	return false
}

// SetMatchOffset gets a reference to the given int32 and assigns it to the MatchOffset field.
func (o *DHCPv4ClientClass) SetMatchOffset(v int32) {
	o.MatchOffset = &v
}

// GetMatchLength returns the MatchLength field value if set, zero value otherwise.
func (o *DHCPv4ClientClass) GetMatchLength() int32 {
	if o == nil || IsNil(o.MatchLength) {
		var ret int32
		return ret
	}
	return *o.MatchLength
}

// GetMatchLengthOk returns a tuple with the MatchLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DHCPv4ClientClass) GetMatchLengthOk() (*int32, bool) {
	if o == nil || IsNil(o.MatchLength) {
		return nil, false
	}
	return o.MatchLength, true
}

// HasMatchLength returns a boolean if a field has been set.
func (o *DHCPv4ClientClass) HasMatchLength() bool {
	if o != nil && !IsNil(o.MatchLength) {
		return true
	}

	return false
}

// SetMatchLength gets a reference to the given int32 and assigns it to the MatchLength field.
func (o *DHCPv4ClientClass) SetMatchLength(v int32) {
	o.MatchLength = &v
}

// GetMatchExpression returns the MatchExpression field value if set, zero value otherwise.
func (o *DHCPv4ClientClass) GetMatchExpression() string {
	if o == nil || IsNil(o.MatchExpression) {
		var ret string
		return ret
	}
	return *o.MatchExpression
}

// GetMatchExpressionOk returns a tuple with the MatchExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DHCPv4ClientClass) GetMatchExpressionOk() (*string, bool) {
	if o == nil || IsNil(o.MatchExpression) {
		return nil, false
	}
	return o.MatchExpression, true
}

// HasMatchExpression returns a boolean if a field has been set.
func (o *DHCPv4ClientClass) HasMatchExpression() bool {
	if o != nil && !IsNil(o.MatchExpression) {
		return true
	}

	return false
}

// SetMatchExpression gets a reference to the given string and assigns it to the MatchExpression field.
func (o *DHCPv4ClientClass) SetMatchExpression(v string) {
	o.MatchExpression = &v
}

func (o DHCPv4ClientClass) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DHCPv4ClientClass) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Option) {
		toSerialize["option"] = o.Option
	}
	if !IsNil(o.MatchOffset) {
		toSerialize["matchOffset"] = o.MatchOffset
	}
	if !IsNil(o.MatchLength) {
		toSerialize["matchLength"] = o.MatchLength
	}
	if !IsNil(o.MatchExpression) {
		toSerialize["matchExpression"] = o.MatchExpression
	}
	return toSerialize, nil
}

type NullableDHCPv4ClientClass struct {
	value *DHCPv4ClientClass
	isSet bool
}

func (v NullableDHCPv4ClientClass) Get() *DHCPv4ClientClass {
	return v.value
}

func (v *NullableDHCPv4ClientClass) Set(val *DHCPv4ClientClass) {
	v.value = val
	v.isSet = true
}

func (v NullableDHCPv4ClientClass) IsSet() bool {
	return v.isSet
}

func (v *NullableDHCPv4ClientClass) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDHCPv4ClientClass(val *DHCPv4ClientClass) *NullableDHCPv4ClientClass {
	return &NullableDHCPv4ClientClass{value: val, isSet: true}
}

func (v NullableDHCPv4ClientClass) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDHCPv4ClientClass) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


