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

// checks if the IPv4GroupTemplateItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IPv4GroupTemplateItem{}

// IPv4GroupTemplateItem struct for IPv4GroupTemplateItem
type IPv4GroupTemplateItem struct {
	Ipv4TemplateItem
	// The resource type.
	Type *string `json:"type,omitempty"`
	// The name of the IPv4 group.
	Name *string `json:"name,omitempty" validate:"regexp=^.*\\\\S+.*$"`
	RangeDefinitionFormat *string `json:"rangeDefinitionFormat,omitempty"`
	Size *int64 `json:"size,omitempty"`
}

// NewIPv4GroupTemplateItem instantiates a new IPv4GroupTemplateItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIPv4GroupTemplateItem() *IPv4GroupTemplateItem {
	this := IPv4GroupTemplateItem{}
	return &this
}

// NewIPv4GroupTemplateItemWithDefaults instantiates a new IPv4GroupTemplateItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIPv4GroupTemplateItemWithDefaults() *IPv4GroupTemplateItem {
	this := IPv4GroupTemplateItem{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IPv4GroupTemplateItem) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPv4GroupTemplateItem) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IPv4GroupTemplateItem) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *IPv4GroupTemplateItem) SetType(v string) {
	o.Type = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IPv4GroupTemplateItem) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPv4GroupTemplateItem) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IPv4GroupTemplateItem) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IPv4GroupTemplateItem) SetName(v string) {
	o.Name = &v
}

// GetRangeDefinitionFormat returns the RangeDefinitionFormat field value if set, zero value otherwise.
func (o *IPv4GroupTemplateItem) GetRangeDefinitionFormat() string {
	if o == nil || IsNil(o.RangeDefinitionFormat) {
		var ret string
		return ret
	}
	return *o.RangeDefinitionFormat
}

// GetRangeDefinitionFormatOk returns a tuple with the RangeDefinitionFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPv4GroupTemplateItem) GetRangeDefinitionFormatOk() (*string, bool) {
	if o == nil || IsNil(o.RangeDefinitionFormat) {
		return nil, false
	}
	return o.RangeDefinitionFormat, true
}

// HasRangeDefinitionFormat returns a boolean if a field has been set.
func (o *IPv4GroupTemplateItem) HasRangeDefinitionFormat() bool {
	if o != nil && !IsNil(o.RangeDefinitionFormat) {
		return true
	}

	return false
}

// SetRangeDefinitionFormat gets a reference to the given string and assigns it to the RangeDefinitionFormat field.
func (o *IPv4GroupTemplateItem) SetRangeDefinitionFormat(v string) {
	o.RangeDefinitionFormat = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *IPv4GroupTemplateItem) GetSize() int64 {
	if o == nil || IsNil(o.Size) {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPv4GroupTemplateItem) GetSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *IPv4GroupTemplateItem) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *IPv4GroupTemplateItem) SetSize(v int64) {
	o.Size = &v
}

func (o IPv4GroupTemplateItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IPv4GroupTemplateItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.RangeDefinitionFormat) {
		toSerialize["rangeDefinitionFormat"] = o.RangeDefinitionFormat
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	return toSerialize, nil
}

type NullableIPv4GroupTemplateItem struct {
	value *IPv4GroupTemplateItem
	isSet bool
}

func (v NullableIPv4GroupTemplateItem) Get() *IPv4GroupTemplateItem {
	return v.value
}

func (v *NullableIPv4GroupTemplateItem) Set(val *IPv4GroupTemplateItem) {
	v.value = val
	v.isSet = true
}

func (v NullableIPv4GroupTemplateItem) IsSet() bool {
	return v.isSet
}

func (v *NullableIPv4GroupTemplateItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIPv4GroupTemplateItem(val *IPv4GroupTemplateItem) *NullableIPv4GroupTemplateItem {
	return &NullableIPv4GroupTemplateItem{value: val, isSet: true}
}

func (v NullableIPv4GroupTemplateItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIPv4GroupTemplateItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


