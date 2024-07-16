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

// checks if the Ipv4TemplateItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Ipv4TemplateItem{}

// Ipv4TemplateItem struct for Ipv4TemplateItem
type Ipv4TemplateItem struct {
	// The resource identifier.
	Id *int64 `json:"id,omitempty"`
	// The resource type.
	Type *string `json:"type,omitempty"`
	// The start offset of the IPv4 template item.
	StartOffset *int64 `json:"startOffset,omitempty"`
	// The name of the template item.
	Name *string `json:"name,omitempty"`
}

// NewIpv4TemplateItem instantiates a new Ipv4TemplateItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpv4TemplateItem() *Ipv4TemplateItem {
	this := Ipv4TemplateItem{}
	return &this
}

// NewIpv4TemplateItemWithDefaults instantiates a new Ipv4TemplateItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpv4TemplateItemWithDefaults() *Ipv4TemplateItem {
	this := Ipv4TemplateItem{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Ipv4TemplateItem) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ipv4TemplateItem) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Ipv4TemplateItem) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *Ipv4TemplateItem) SetId(v int64) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Ipv4TemplateItem) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ipv4TemplateItem) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Ipv4TemplateItem) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Ipv4TemplateItem) SetType(v string) {
	o.Type = &v
}

// GetStartOffset returns the StartOffset field value if set, zero value otherwise.
func (o *Ipv4TemplateItem) GetStartOffset() int64 {
	if o == nil || IsNil(o.StartOffset) {
		var ret int64
		return ret
	}
	return *o.StartOffset
}

// GetStartOffsetOk returns a tuple with the StartOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ipv4TemplateItem) GetStartOffsetOk() (*int64, bool) {
	if o == nil || IsNil(o.StartOffset) {
		return nil, false
	}
	return o.StartOffset, true
}

// HasStartOffset returns a boolean if a field has been set.
func (o *Ipv4TemplateItem) HasStartOffset() bool {
	if o != nil && !IsNil(o.StartOffset) {
		return true
	}

	return false
}

// SetStartOffset gets a reference to the given int64 and assigns it to the StartOffset field.
func (o *Ipv4TemplateItem) SetStartOffset(v int64) {
	o.StartOffset = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Ipv4TemplateItem) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ipv4TemplateItem) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Ipv4TemplateItem) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Ipv4TemplateItem) SetName(v string) {
	o.Name = &v
}

func (o Ipv4TemplateItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Ipv4TemplateItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.StartOffset) {
		toSerialize["startOffset"] = o.StartOffset
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableIpv4TemplateItem struct {
	value *Ipv4TemplateItem
	isSet bool
}

func (v NullableIpv4TemplateItem) Get() *Ipv4TemplateItem {
	return v.value
}

func (v *NullableIpv4TemplateItem) Set(val *Ipv4TemplateItem) {
	v.value = val
	v.isSet = true
}

func (v NullableIpv4TemplateItem) IsSet() bool {
	return v.isSet
}

func (v *NullableIpv4TemplateItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpv4TemplateItem(val *Ipv4TemplateItem) *NullableIpv4TemplateItem {
	return &NullableIpv4TemplateItem{value: val, isSet: true}
}

func (v NullableIpv4TemplateItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpv4TemplateItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


