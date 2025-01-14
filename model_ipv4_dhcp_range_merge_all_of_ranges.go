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

// checks if the IPv4DHCPRangeMergeAllOfRanges type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IPv4DHCPRangeMergeAllOfRanges{}

// IPv4DHCPRangeMergeAllOfRanges struct for IPv4DHCPRangeMergeAllOfRanges
type IPv4DHCPRangeMergeAllOfRanges struct {
	// The resource identifier.
	Id int64 `json:"id"`
	// The resource type.
	Type *string `json:"type,omitempty"`
	// The name of the resource.
	Name *string `json:"name,omitempty"`
}

type _IPv4DHCPRangeMergeAllOfRanges IPv4DHCPRangeMergeAllOfRanges

// NewIPv4DHCPRangeMergeAllOfRanges instantiates a new IPv4DHCPRangeMergeAllOfRanges object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIPv4DHCPRangeMergeAllOfRanges(id int64) *IPv4DHCPRangeMergeAllOfRanges {
	this := IPv4DHCPRangeMergeAllOfRanges{}
	this.Id = id
	return &this
}

// NewIPv4DHCPRangeMergeAllOfRangesWithDefaults instantiates a new IPv4DHCPRangeMergeAllOfRanges object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIPv4DHCPRangeMergeAllOfRangesWithDefaults() *IPv4DHCPRangeMergeAllOfRanges {
	this := IPv4DHCPRangeMergeAllOfRanges{}
	return &this
}

// GetId returns the Id field value
func (o *IPv4DHCPRangeMergeAllOfRanges) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *IPv4DHCPRangeMergeAllOfRanges) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *IPv4DHCPRangeMergeAllOfRanges) SetId(v int64) {
	o.Id = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IPv4DHCPRangeMergeAllOfRanges) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPv4DHCPRangeMergeAllOfRanges) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IPv4DHCPRangeMergeAllOfRanges) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *IPv4DHCPRangeMergeAllOfRanges) SetType(v string) {
	o.Type = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IPv4DHCPRangeMergeAllOfRanges) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPv4DHCPRangeMergeAllOfRanges) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IPv4DHCPRangeMergeAllOfRanges) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IPv4DHCPRangeMergeAllOfRanges) SetName(v string) {
	o.Name = &v
}

func (o IPv4DHCPRangeMergeAllOfRanges) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IPv4DHCPRangeMergeAllOfRanges) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

func (o *IPv4DHCPRangeMergeAllOfRanges) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varIPv4DHCPRangeMergeAllOfRanges := _IPv4DHCPRangeMergeAllOfRanges{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varIPv4DHCPRangeMergeAllOfRanges)

	if err != nil {
		return err
	}

	*o = IPv4DHCPRangeMergeAllOfRanges(varIPv4DHCPRangeMergeAllOfRanges)

	return err
}

type NullableIPv4DHCPRangeMergeAllOfRanges struct {
	value *IPv4DHCPRangeMergeAllOfRanges
	isSet bool
}

func (v NullableIPv4DHCPRangeMergeAllOfRanges) Get() *IPv4DHCPRangeMergeAllOfRanges {
	return v.value
}

func (v *NullableIPv4DHCPRangeMergeAllOfRanges) Set(val *IPv4DHCPRangeMergeAllOfRanges) {
	v.value = val
	v.isSet = true
}

func (v NullableIPv4DHCPRangeMergeAllOfRanges) IsSet() bool {
	return v.isSet
}

func (v *NullableIPv4DHCPRangeMergeAllOfRanges) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIPv4DHCPRangeMergeAllOfRanges(val *IPv4DHCPRangeMergeAllOfRanges) *NullableIPv4DHCPRangeMergeAllOfRanges {
	return &NullableIPv4DHCPRangeMergeAllOfRanges{value: val, isSet: true}
}

func (v NullableIPv4DHCPRangeMergeAllOfRanges) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIPv4DHCPRangeMergeAllOfRanges) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


