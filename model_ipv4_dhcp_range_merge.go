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

// checks if the IPv4DHCPRangeMerge type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IPv4DHCPRangeMerge{}

// IPv4DHCPRangeMerge struct for IPv4DHCPRangeMerge
type IPv4DHCPRangeMerge struct {
	Merge
	// The resource type.
	Type *string `json:"type,omitempty"`
	Ranges []IPv4DHCPRangeMergeAllOfRanges `json:"ranges,omitempty"`
	Destination *InlinedIPv4DHCPRange `json:"destination,omitempty"`
}

// NewIPv4DHCPRangeMerge instantiates a new IPv4DHCPRangeMerge object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIPv4DHCPRangeMerge() *IPv4DHCPRangeMerge {
	this := IPv4DHCPRangeMerge{}
	return &this
}

// NewIPv4DHCPRangeMergeWithDefaults instantiates a new IPv4DHCPRangeMerge object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIPv4DHCPRangeMergeWithDefaults() *IPv4DHCPRangeMerge {
	this := IPv4DHCPRangeMerge{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IPv4DHCPRangeMerge) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPv4DHCPRangeMerge) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IPv4DHCPRangeMerge) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *IPv4DHCPRangeMerge) SetType(v string) {
	o.Type = &v
}

// GetRanges returns the Ranges field value if set, zero value otherwise.
func (o *IPv4DHCPRangeMerge) GetRanges() []IPv4DHCPRangeMergeAllOfRanges {
	if o == nil || IsNil(o.Ranges) {
		var ret []IPv4DHCPRangeMergeAllOfRanges
		return ret
	}
	return o.Ranges
}

// GetRangesOk returns a tuple with the Ranges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPv4DHCPRangeMerge) GetRangesOk() ([]IPv4DHCPRangeMergeAllOfRanges, bool) {
	if o == nil || IsNil(o.Ranges) {
		return nil, false
	}
	return o.Ranges, true
}

// HasRanges returns a boolean if a field has been set.
func (o *IPv4DHCPRangeMerge) HasRanges() bool {
	if o != nil && !IsNil(o.Ranges) {
		return true
	}

	return false
}

// SetRanges gets a reference to the given []IPv4DHCPRangeMergeAllOfRanges and assigns it to the Ranges field.
func (o *IPv4DHCPRangeMerge) SetRanges(v []IPv4DHCPRangeMergeAllOfRanges) {
	o.Ranges = v
}

// GetDestination returns the Destination field value if set, zero value otherwise.
func (o *IPv4DHCPRangeMerge) GetDestination() InlinedIPv4DHCPRange {
	if o == nil || IsNil(o.Destination) {
		var ret InlinedIPv4DHCPRange
		return ret
	}
	return *o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPv4DHCPRangeMerge) GetDestinationOk() (*InlinedIPv4DHCPRange, bool) {
	if o == nil || IsNil(o.Destination) {
		return nil, false
	}
	return o.Destination, true
}

// HasDestination returns a boolean if a field has been set.
func (o *IPv4DHCPRangeMerge) HasDestination() bool {
	if o != nil && !IsNil(o.Destination) {
		return true
	}

	return false
}

// SetDestination gets a reference to the given InlinedIPv4DHCPRange and assigns it to the Destination field.
func (o *IPv4DHCPRangeMerge) SetDestination(v InlinedIPv4DHCPRange) {
	o.Destination = &v
}

func (o IPv4DHCPRangeMerge) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IPv4DHCPRangeMerge) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Ranges) {
		toSerialize["ranges"] = o.Ranges
	}
	if !IsNil(o.Destination) {
		toSerialize["destination"] = o.Destination
	}
	return toSerialize, nil
}

type NullableIPv4DHCPRangeMerge struct {
	value *IPv4DHCPRangeMerge
	isSet bool
}

func (v NullableIPv4DHCPRangeMerge) Get() *IPv4DHCPRangeMerge {
	return v.value
}

func (v *NullableIPv4DHCPRangeMerge) Set(val *IPv4DHCPRangeMerge) {
	v.value = val
	v.isSet = true
}

func (v NullableIPv4DHCPRangeMerge) IsSet() bool {
	return v.isSet
}

func (v *NullableIPv4DHCPRangeMerge) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIPv4DHCPRangeMerge(val *IPv4DHCPRangeMerge) *NullableIPv4DHCPRangeMerge {
	return &NullableIPv4DHCPRangeMerge{value: val, isSet: true}
}

func (v NullableIPv4DHCPRangeMerge) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIPv4DHCPRangeMerge) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


