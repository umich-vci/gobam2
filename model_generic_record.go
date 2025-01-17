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

// checks if the GenericRecord type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GenericRecord{}

// GenericRecord struct for GenericRecord
type GenericRecord struct {
	ResourceRecord
	// The resource type.
	Type *string `json:"type,omitempty"`
	// The resource record type.
	RecordType *string `json:"recordType,omitempty"`
	// The resource record data.
	Rdata *string `json:"rdata,omitempty" validate:"regexp=^.*\\\\S+.*$"`
}

// NewGenericRecord instantiates a new GenericRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenericRecord() *GenericRecord {
	this := GenericRecord{}
	return &this
}

// NewGenericRecordWithDefaults instantiates a new GenericRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenericRecordWithDefaults() *GenericRecord {
	this := GenericRecord{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GenericRecord) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericRecord) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GenericRecord) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GenericRecord) SetType(v string) {
	o.Type = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *GenericRecord) GetRecordType() string {
	if o == nil || IsNil(o.RecordType) {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericRecord) GetRecordTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RecordType) {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *GenericRecord) HasRecordType() bool {
	if o != nil && !IsNil(o.RecordType) {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *GenericRecord) SetRecordType(v string) {
	o.RecordType = &v
}

// GetRdata returns the Rdata field value if set, zero value otherwise.
func (o *GenericRecord) GetRdata() string {
	if o == nil || IsNil(o.Rdata) {
		var ret string
		return ret
	}
	return *o.Rdata
}

// GetRdataOk returns a tuple with the Rdata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericRecord) GetRdataOk() (*string, bool) {
	if o == nil || IsNil(o.Rdata) {
		return nil, false
	}
	return o.Rdata, true
}

// HasRdata returns a boolean if a field has been set.
func (o *GenericRecord) HasRdata() bool {
	if o != nil && !IsNil(o.Rdata) {
		return true
	}

	return false
}

// SetRdata gets a reference to the given string and assigns it to the Rdata field.
func (o *GenericRecord) SetRdata(v string) {
	o.Rdata = &v
}

func (o GenericRecord) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GenericRecord) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.RecordType) {
		toSerialize["recordType"] = o.RecordType
	}
	if !IsNil(o.Rdata) {
		toSerialize["rdata"] = o.Rdata
	}
	return toSerialize, nil
}

type NullableGenericRecord struct {
	value *GenericRecord
	isSet bool
}

func (v NullableGenericRecord) Get() *GenericRecord {
	return v.value
}

func (v *NullableGenericRecord) Set(val *GenericRecord) {
	v.value = val
	v.isSet = true
}

func (v NullableGenericRecord) IsSet() bool {
	return v.isSet
}

func (v *NullableGenericRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenericRecord(val *GenericRecord) *NullableGenericRecord {
	return &NullableGenericRecord{value: val, isSet: true}
}

func (v NullableGenericRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenericRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


