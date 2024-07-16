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

// checks if the IPMove type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IPMove{}

// IPMove struct for IPMove
type IPMove struct {
	Move
	// The resource type.
	Type *string `json:"type,omitempty"`
	// The IP address of the destination resource where the IP resource is to be moved.
	Address *string `json:"address,omitempty"`
	MacAddress *IPMoveAllOfMacAddress `json:"macAddress,omitempty"`
	ServersUpdated *bool `json:"serversUpdated,omitempty"`
}

// NewIPMove instantiates a new IPMove object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIPMove() *IPMove {
	this := IPMove{}
	return &this
}

// NewIPMoveWithDefaults instantiates a new IPMove object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIPMoveWithDefaults() *IPMove {
	this := IPMove{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IPMove) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPMove) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IPMove) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *IPMove) SetType(v string) {
	o.Type = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *IPMove) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPMove) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *IPMove) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *IPMove) SetAddress(v string) {
	o.Address = &v
}

// GetMacAddress returns the MacAddress field value if set, zero value otherwise.
func (o *IPMove) GetMacAddress() IPMoveAllOfMacAddress {
	if o == nil || IsNil(o.MacAddress) {
		var ret IPMoveAllOfMacAddress
		return ret
	}
	return *o.MacAddress
}

// GetMacAddressOk returns a tuple with the MacAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPMove) GetMacAddressOk() (*IPMoveAllOfMacAddress, bool) {
	if o == nil || IsNil(o.MacAddress) {
		return nil, false
	}
	return o.MacAddress, true
}

// HasMacAddress returns a boolean if a field has been set.
func (o *IPMove) HasMacAddress() bool {
	if o != nil && !IsNil(o.MacAddress) {
		return true
	}

	return false
}

// SetMacAddress gets a reference to the given IPMoveAllOfMacAddress and assigns it to the MacAddress field.
func (o *IPMove) SetMacAddress(v IPMoveAllOfMacAddress) {
	o.MacAddress = &v
}

// GetServersUpdated returns the ServersUpdated field value if set, zero value otherwise.
func (o *IPMove) GetServersUpdated() bool {
	if o == nil || IsNil(o.ServersUpdated) {
		var ret bool
		return ret
	}
	return *o.ServersUpdated
}

// GetServersUpdatedOk returns a tuple with the ServersUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPMove) GetServersUpdatedOk() (*bool, bool) {
	if o == nil || IsNil(o.ServersUpdated) {
		return nil, false
	}
	return o.ServersUpdated, true
}

// HasServersUpdated returns a boolean if a field has been set.
func (o *IPMove) HasServersUpdated() bool {
	if o != nil && !IsNil(o.ServersUpdated) {
		return true
	}

	return false
}

// SetServersUpdated gets a reference to the given bool and assigns it to the ServersUpdated field.
func (o *IPMove) SetServersUpdated(v bool) {
	o.ServersUpdated = &v
}

func (o IPMove) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IPMove) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.MacAddress) {
		toSerialize["macAddress"] = o.MacAddress
	}
	if !IsNil(o.ServersUpdated) {
		toSerialize["serversUpdated"] = o.ServersUpdated
	}
	return toSerialize, nil
}

type NullableIPMove struct {
	value *IPMove
	isSet bool
}

func (v NullableIPMove) Get() *IPMove {
	return v.value
}

func (v *NullableIPMove) Set(val *IPMove) {
	v.value = val
	v.isSet = true
}

func (v NullableIPMove) IsSet() bool {
	return v.isSet
}

func (v *NullableIPMove) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIPMove(val *IPMove) *NullableIPMove {
	return &NullableIPMove{value: val, isSet: true}
}

func (v NullableIPMove) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIPMove) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


