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

// checks if the IPv4BlockSplit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IPv4BlockSplit{}

// IPv4BlockSplit struct for IPv4BlockSplit
type IPv4BlockSplit struct {
	Split
	// The resource type.
	Type *string `json:"type,omitempty"`
	Block *InlinedIPv4Block `json:"block,omitempty"`
	// The IP address of the point at which you would like to split the block.
	SplitPointAddress *string `json:"splitPointAddress,omitempty"`
}

// NewIPv4BlockSplit instantiates a new IPv4BlockSplit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIPv4BlockSplit() *IPv4BlockSplit {
	this := IPv4BlockSplit{}
	return &this
}

// NewIPv4BlockSplitWithDefaults instantiates a new IPv4BlockSplit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIPv4BlockSplitWithDefaults() *IPv4BlockSplit {
	this := IPv4BlockSplit{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IPv4BlockSplit) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPv4BlockSplit) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IPv4BlockSplit) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *IPv4BlockSplit) SetType(v string) {
	o.Type = &v
}

// GetBlock returns the Block field value if set, zero value otherwise.
func (o *IPv4BlockSplit) GetBlock() InlinedIPv4Block {
	if o == nil || IsNil(o.Block) {
		var ret InlinedIPv4Block
		return ret
	}
	return *o.Block
}

// GetBlockOk returns a tuple with the Block field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPv4BlockSplit) GetBlockOk() (*InlinedIPv4Block, bool) {
	if o == nil || IsNil(o.Block) {
		return nil, false
	}
	return o.Block, true
}

// HasBlock returns a boolean if a field has been set.
func (o *IPv4BlockSplit) HasBlock() bool {
	if o != nil && !IsNil(o.Block) {
		return true
	}

	return false
}

// SetBlock gets a reference to the given InlinedIPv4Block and assigns it to the Block field.
func (o *IPv4BlockSplit) SetBlock(v InlinedIPv4Block) {
	o.Block = &v
}

// GetSplitPointAddress returns the SplitPointAddress field value if set, zero value otherwise.
func (o *IPv4BlockSplit) GetSplitPointAddress() string {
	if o == nil || IsNil(o.SplitPointAddress) {
		var ret string
		return ret
	}
	return *o.SplitPointAddress
}

// GetSplitPointAddressOk returns a tuple with the SplitPointAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPv4BlockSplit) GetSplitPointAddressOk() (*string, bool) {
	if o == nil || IsNil(o.SplitPointAddress) {
		return nil, false
	}
	return o.SplitPointAddress, true
}

// HasSplitPointAddress returns a boolean if a field has been set.
func (o *IPv4BlockSplit) HasSplitPointAddress() bool {
	if o != nil && !IsNil(o.SplitPointAddress) {
		return true
	}

	return false
}

// SetSplitPointAddress gets a reference to the given string and assigns it to the SplitPointAddress field.
func (o *IPv4BlockSplit) SetSplitPointAddress(v string) {
	o.SplitPointAddress = &v
}

func (o IPv4BlockSplit) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IPv4BlockSplit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Block) {
		toSerialize["block"] = o.Block
	}
	if !IsNil(o.SplitPointAddress) {
		toSerialize["splitPointAddress"] = o.SplitPointAddress
	}
	return toSerialize, nil
}

type NullableIPv4BlockSplit struct {
	value *IPv4BlockSplit
	isSet bool
}

func (v NullableIPv4BlockSplit) Get() *IPv4BlockSplit {
	return v.value
}

func (v *NullableIPv4BlockSplit) Set(val *IPv4BlockSplit) {
	v.value = val
	v.isSet = true
}

func (v NullableIPv4BlockSplit) IsSet() bool {
	return v.isSet
}

func (v *NullableIPv4BlockSplit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIPv4BlockSplit(val *IPv4BlockSplit) *NullableIPv4BlockSplit {
	return &NullableIPv4BlockSplit{value: val, isSet: true}
}

func (v NullableIPv4BlockSplit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIPv4BlockSplit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


