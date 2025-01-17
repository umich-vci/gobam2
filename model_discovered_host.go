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

// checks if the DiscoveredHost type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DiscoveredHost{}

// DiscoveredHost struct for DiscoveredHost
type DiscoveredHost struct {
	DiscoveredDevice
	// The resource type.
	Type *string `json:"type,omitempty"`
	Device *DiscoveredDevice `json:"device,omitempty"`
	Address *string `json:"address,omitempty"`
	MacAddress *string `json:"macAddress,omitempty"`
	Interface *DiscoveredInterface `json:"interface,omitempty"`
}

// NewDiscoveredHost instantiates a new DiscoveredHost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiscoveredHost() *DiscoveredHost {
	this := DiscoveredHost{}
	return &this
}

// NewDiscoveredHostWithDefaults instantiates a new DiscoveredHost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiscoveredHostWithDefaults() *DiscoveredHost {
	this := DiscoveredHost{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DiscoveredHost) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveredHost) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DiscoveredHost) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DiscoveredHost) SetType(v string) {
	o.Type = &v
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *DiscoveredHost) GetDevice() DiscoveredDevice {
	if o == nil || IsNil(o.Device) {
		var ret DiscoveredDevice
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveredHost) GetDeviceOk() (*DiscoveredDevice, bool) {
	if o == nil || IsNil(o.Device) {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *DiscoveredHost) HasDevice() bool {
	if o != nil && !IsNil(o.Device) {
		return true
	}

	return false
}

// SetDevice gets a reference to the given DiscoveredDevice and assigns it to the Device field.
func (o *DiscoveredHost) SetDevice(v DiscoveredDevice) {
	o.Device = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *DiscoveredHost) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveredHost) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *DiscoveredHost) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *DiscoveredHost) SetAddress(v string) {
	o.Address = &v
}

// GetMacAddress returns the MacAddress field value if set, zero value otherwise.
func (o *DiscoveredHost) GetMacAddress() string {
	if o == nil || IsNil(o.MacAddress) {
		var ret string
		return ret
	}
	return *o.MacAddress
}

// GetMacAddressOk returns a tuple with the MacAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveredHost) GetMacAddressOk() (*string, bool) {
	if o == nil || IsNil(o.MacAddress) {
		return nil, false
	}
	return o.MacAddress, true
}

// HasMacAddress returns a boolean if a field has been set.
func (o *DiscoveredHost) HasMacAddress() bool {
	if o != nil && !IsNil(o.MacAddress) {
		return true
	}

	return false
}

// SetMacAddress gets a reference to the given string and assigns it to the MacAddress field.
func (o *DiscoveredHost) SetMacAddress(v string) {
	o.MacAddress = &v
}

// GetInterface returns the Interface field value if set, zero value otherwise.
func (o *DiscoveredHost) GetInterface() DiscoveredInterface {
	if o == nil || IsNil(o.Interface) {
		var ret DiscoveredInterface
		return ret
	}
	return *o.Interface
}

// GetInterfaceOk returns a tuple with the Interface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveredHost) GetInterfaceOk() (*DiscoveredInterface, bool) {
	if o == nil || IsNil(o.Interface) {
		return nil, false
	}
	return o.Interface, true
}

// HasInterface returns a boolean if a field has been set.
func (o *DiscoveredHost) HasInterface() bool {
	if o != nil && !IsNil(o.Interface) {
		return true
	}

	return false
}

// SetInterface gets a reference to the given DiscoveredInterface and assigns it to the Interface field.
func (o *DiscoveredHost) SetInterface(v DiscoveredInterface) {
	o.Interface = &v
}

func (o DiscoveredHost) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiscoveredHost) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Device) {
		toSerialize["device"] = o.Device
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.MacAddress) {
		toSerialize["macAddress"] = o.MacAddress
	}
	if !IsNil(o.Interface) {
		toSerialize["interface"] = o.Interface
	}
	return toSerialize, nil
}

type NullableDiscoveredHost struct {
	value *DiscoveredHost
	isSet bool
}

func (v NullableDiscoveredHost) Get() *DiscoveredHost {
	return v.value
}

func (v *NullableDiscoveredHost) Set(val *DiscoveredHost) {
	v.value = val
	v.isSet = true
}

func (v NullableDiscoveredHost) IsSet() bool {
	return v.isSet
}

func (v *NullableDiscoveredHost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiscoveredHost(val *DiscoveredHost) *NullableDiscoveredHost {
	return &NullableDiscoveredHost{value: val, isSet: true}
}

func (v NullableDiscoveredHost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiscoveredHost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


