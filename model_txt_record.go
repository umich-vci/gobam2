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

// checks if the TXTRecord type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TXTRecord{}

// TXTRecord struct for TXTRecord
type TXTRecord struct {
	ResourceRecord
	// The resource type.
	Type *string `json:"type,omitempty"`
	// The text data within the TXT resource record.
	Text *string `json:"text,omitempty" validate:"regexp=^.*\\\\S+.*$"`
}

// NewTXTRecord instantiates a new TXTRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTXTRecord() *TXTRecord {
	this := TXTRecord{}
	return &this
}

// NewTXTRecordWithDefaults instantiates a new TXTRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTXTRecordWithDefaults() *TXTRecord {
	this := TXTRecord{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TXTRecord) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TXTRecord) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TXTRecord) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TXTRecord) SetType(v string) {
	o.Type = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *TXTRecord) GetText() string {
	if o == nil || IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TXTRecord) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *TXTRecord) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *TXTRecord) SetText(v string) {
	o.Text = &v
}

func (o TXTRecord) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TXTRecord) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	return toSerialize, nil
}

type NullableTXTRecord struct {
	value *TXTRecord
	isSet bool
}

func (v NullableTXTRecord) Get() *TXTRecord {
	return v.value
}

func (v *NullableTXTRecord) Set(val *TXTRecord) {
	v.value = val
	v.isSet = true
}

func (v NullableTXTRecord) IsSet() bool {
	return v.isSet
}

func (v *NullableTXTRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTXTRecord(val *TXTRecord) *NullableTXTRecord {
	return &NullableTXTRecord{value: val, isSet: true}
}

func (v NullableTXTRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTXTRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


