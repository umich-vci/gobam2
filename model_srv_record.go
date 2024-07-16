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

// checks if the SRVRecord type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SRVRecord{}

// SRVRecord struct for SRVRecord
type SRVRecord struct {
	ResourceRecord
	// The resource type.
	Type *string `json:"type,omitempty"`
	LinkedRecord *SRVRecordAllOfLinkedRecord `json:"linkedRecord,omitempty"`
	// The priority assigned to the service. A lower value indicates higher priority.
	Priority *int32 `json:"priority,omitempty"`
	// The weight assigned to the service. The weight is referenced when services have the same priority.
	Weight *int32 `json:"weight,omitempty"`
	// The port on which the service is located.
	Port *int32 `json:"port,omitempty"`
}

// NewSRVRecord instantiates a new SRVRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSRVRecord() *SRVRecord {
	this := SRVRecord{}
	return &this
}

// NewSRVRecordWithDefaults instantiates a new SRVRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSRVRecordWithDefaults() *SRVRecord {
	this := SRVRecord{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SRVRecord) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SRVRecord) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SRVRecord) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SRVRecord) SetType(v string) {
	o.Type = &v
}

// GetLinkedRecord returns the LinkedRecord field value if set, zero value otherwise.
func (o *SRVRecord) GetLinkedRecord() SRVRecordAllOfLinkedRecord {
	if o == nil || IsNil(o.LinkedRecord) {
		var ret SRVRecordAllOfLinkedRecord
		return ret
	}
	return *o.LinkedRecord
}

// GetLinkedRecordOk returns a tuple with the LinkedRecord field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SRVRecord) GetLinkedRecordOk() (*SRVRecordAllOfLinkedRecord, bool) {
	if o == nil || IsNil(o.LinkedRecord) {
		return nil, false
	}
	return o.LinkedRecord, true
}

// HasLinkedRecord returns a boolean if a field has been set.
func (o *SRVRecord) HasLinkedRecord() bool {
	if o != nil && !IsNil(o.LinkedRecord) {
		return true
	}

	return false
}

// SetLinkedRecord gets a reference to the given SRVRecordAllOfLinkedRecord and assigns it to the LinkedRecord field.
func (o *SRVRecord) SetLinkedRecord(v SRVRecordAllOfLinkedRecord) {
	o.LinkedRecord = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *SRVRecord) GetPriority() int32 {
	if o == nil || IsNil(o.Priority) {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SRVRecord) GetPriorityOk() (*int32, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *SRVRecord) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *SRVRecord) SetPriority(v int32) {
	o.Priority = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *SRVRecord) GetWeight() int32 {
	if o == nil || IsNil(o.Weight) {
		var ret int32
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SRVRecord) GetWeightOk() (*int32, bool) {
	if o == nil || IsNil(o.Weight) {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *SRVRecord) HasWeight() bool {
	if o != nil && !IsNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given int32 and assigns it to the Weight field.
func (o *SRVRecord) SetWeight(v int32) {
	o.Weight = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *SRVRecord) GetPort() int32 {
	if o == nil || IsNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SRVRecord) GetPortOk() (*int32, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *SRVRecord) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *SRVRecord) SetPort(v int32) {
	o.Port = &v
}

func (o SRVRecord) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SRVRecord) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.LinkedRecord) {
		toSerialize["linkedRecord"] = o.LinkedRecord
	}
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !IsNil(o.Weight) {
		toSerialize["weight"] = o.Weight
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	return toSerialize, nil
}

type NullableSRVRecord struct {
	value *SRVRecord
	isSet bool
}

func (v NullableSRVRecord) Get() *SRVRecord {
	return v.value
}

func (v *NullableSRVRecord) Set(val *SRVRecord) {
	v.value = val
	v.isSet = true
}

func (v NullableSRVRecord) IsSet() bool {
	return v.isSet
}

func (v *NullableSRVRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSRVRecord(val *SRVRecord) *NullableSRVRecord {
	return &NullableSRVRecord{value: val, isSet: true}
}

func (v NullableSRVRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSRVRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

