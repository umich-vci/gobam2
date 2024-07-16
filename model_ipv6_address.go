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

// checks if the IPv6Address type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IPv6Address{}

// IPv6Address struct for IPv6Address
type IPv6Address struct {
	Address
	// The type of IP address.
	Type *string `json:"type,omitempty"`
	// The state of the IP address.
	State *string `json:"state,omitempty"`
	// The type of DHCP reservation.
	ReservedUsing *string `json:"reservedUsing,omitempty"`
	// The Identity Association Identifier for the network interface of the client system.
	IdentityAssociationIdentifier *string `json:"identityAssociationIdentifier,omitempty"`
	ClientIdentifier *InlinedDHCPUniqueIdentifier `json:"clientIdentifier,omitempty"`
	// The ID of the interface from which the request came.
	InterfaceId *string `json:"interfaceId,omitempty"`
}

// NewIPv6Address instantiates a new IPv6Address object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIPv6Address() *IPv6Address {
	this := IPv6Address{}
	return &this
}

// NewIPv6AddressWithDefaults instantiates a new IPv6Address object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIPv6AddressWithDefaults() *IPv6Address {
	this := IPv6Address{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IPv6Address) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPv6Address) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IPv6Address) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *IPv6Address) SetType(v string) {
	o.Type = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *IPv6Address) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPv6Address) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *IPv6Address) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *IPv6Address) SetState(v string) {
	o.State = &v
}

// GetReservedUsing returns the ReservedUsing field value if set, zero value otherwise.
func (o *IPv6Address) GetReservedUsing() string {
	if o == nil || IsNil(o.ReservedUsing) {
		var ret string
		return ret
	}
	return *o.ReservedUsing
}

// GetReservedUsingOk returns a tuple with the ReservedUsing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPv6Address) GetReservedUsingOk() (*string, bool) {
	if o == nil || IsNil(o.ReservedUsing) {
		return nil, false
	}
	return o.ReservedUsing, true
}

// HasReservedUsing returns a boolean if a field has been set.
func (o *IPv6Address) HasReservedUsing() bool {
	if o != nil && !IsNil(o.ReservedUsing) {
		return true
	}

	return false
}

// SetReservedUsing gets a reference to the given string and assigns it to the ReservedUsing field.
func (o *IPv6Address) SetReservedUsing(v string) {
	o.ReservedUsing = &v
}

// GetIdentityAssociationIdentifier returns the IdentityAssociationIdentifier field value if set, zero value otherwise.
func (o *IPv6Address) GetIdentityAssociationIdentifier() string {
	if o == nil || IsNil(o.IdentityAssociationIdentifier) {
		var ret string
		return ret
	}
	return *o.IdentityAssociationIdentifier
}

// GetIdentityAssociationIdentifierOk returns a tuple with the IdentityAssociationIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPv6Address) GetIdentityAssociationIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.IdentityAssociationIdentifier) {
		return nil, false
	}
	return o.IdentityAssociationIdentifier, true
}

// HasIdentityAssociationIdentifier returns a boolean if a field has been set.
func (o *IPv6Address) HasIdentityAssociationIdentifier() bool {
	if o != nil && !IsNil(o.IdentityAssociationIdentifier) {
		return true
	}

	return false
}

// SetIdentityAssociationIdentifier gets a reference to the given string and assigns it to the IdentityAssociationIdentifier field.
func (o *IPv6Address) SetIdentityAssociationIdentifier(v string) {
	o.IdentityAssociationIdentifier = &v
}

// GetClientIdentifier returns the ClientIdentifier field value if set, zero value otherwise.
func (o *IPv6Address) GetClientIdentifier() InlinedDHCPUniqueIdentifier {
	if o == nil || IsNil(o.ClientIdentifier) {
		var ret InlinedDHCPUniqueIdentifier
		return ret
	}
	return *o.ClientIdentifier
}

// GetClientIdentifierOk returns a tuple with the ClientIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPv6Address) GetClientIdentifierOk() (*InlinedDHCPUniqueIdentifier, bool) {
	if o == nil || IsNil(o.ClientIdentifier) {
		return nil, false
	}
	return o.ClientIdentifier, true
}

// HasClientIdentifier returns a boolean if a field has been set.
func (o *IPv6Address) HasClientIdentifier() bool {
	if o != nil && !IsNil(o.ClientIdentifier) {
		return true
	}

	return false
}

// SetClientIdentifier gets a reference to the given InlinedDHCPUniqueIdentifier and assigns it to the ClientIdentifier field.
func (o *IPv6Address) SetClientIdentifier(v InlinedDHCPUniqueIdentifier) {
	o.ClientIdentifier = &v
}

// GetInterfaceId returns the InterfaceId field value if set, zero value otherwise.
func (o *IPv6Address) GetInterfaceId() string {
	if o == nil || IsNil(o.InterfaceId) {
		var ret string
		return ret
	}
	return *o.InterfaceId
}

// GetInterfaceIdOk returns a tuple with the InterfaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPv6Address) GetInterfaceIdOk() (*string, bool) {
	if o == nil || IsNil(o.InterfaceId) {
		return nil, false
	}
	return o.InterfaceId, true
}

// HasInterfaceId returns a boolean if a field has been set.
func (o *IPv6Address) HasInterfaceId() bool {
	if o != nil && !IsNil(o.InterfaceId) {
		return true
	}

	return false
}

// SetInterfaceId gets a reference to the given string and assigns it to the InterfaceId field.
func (o *IPv6Address) SetInterfaceId(v string) {
	o.InterfaceId = &v
}

func (o IPv6Address) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IPv6Address) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.ReservedUsing) {
		toSerialize["reservedUsing"] = o.ReservedUsing
	}
	if !IsNil(o.IdentityAssociationIdentifier) {
		toSerialize["identityAssociationIdentifier"] = o.IdentityAssociationIdentifier
	}
	if !IsNil(o.ClientIdentifier) {
		toSerialize["clientIdentifier"] = o.ClientIdentifier
	}
	if !IsNil(o.InterfaceId) {
		toSerialize["interfaceId"] = o.InterfaceId
	}
	return toSerialize, nil
}

type NullableIPv6Address struct {
	value *IPv6Address
	isSet bool
}

func (v NullableIPv6Address) Get() *IPv6Address {
	return v.value
}

func (v *NullableIPv6Address) Set(val *IPv6Address) {
	v.value = val
	v.isSet = true
}

func (v NullableIPv6Address) IsSet() bool {
	return v.isSet
}

func (v *NullableIPv6Address) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIPv6Address(val *IPv6Address) *NullableIPv6Address {
	return &NullableIPv6Address{value: val, isSet: true}
}

func (v NullableIPv6Address) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIPv6Address) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


