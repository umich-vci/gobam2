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

// checks if the InterfacesService type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InterfacesService{}

// InterfacesService struct for InterfacesService
type InterfacesService struct {
	Service
	// The resource type.
	Type *string `json:"type,omitempty"`
	DedicatedManagementEnabled *bool `json:"dedicatedManagementEnabled,omitempty"`
	Interfaces []InterfaceBean `json:"interfaces,omitempty"`
	Routes []RouteBean `json:"routes,omitempty"`
}

// NewInterfacesService instantiates a new InterfacesService object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInterfacesService() *InterfacesService {
	this := InterfacesService{}
	return &this
}

// NewInterfacesServiceWithDefaults instantiates a new InterfacesService object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInterfacesServiceWithDefaults() *InterfacesService {
	this := InterfacesService{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InterfacesService) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfacesService) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InterfacesService) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InterfacesService) SetType(v string) {
	o.Type = &v
}

// GetDedicatedManagementEnabled returns the DedicatedManagementEnabled field value if set, zero value otherwise.
func (o *InterfacesService) GetDedicatedManagementEnabled() bool {
	if o == nil || IsNil(o.DedicatedManagementEnabled) {
		var ret bool
		return ret
	}
	return *o.DedicatedManagementEnabled
}

// GetDedicatedManagementEnabledOk returns a tuple with the DedicatedManagementEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfacesService) GetDedicatedManagementEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.DedicatedManagementEnabled) {
		return nil, false
	}
	return o.DedicatedManagementEnabled, true
}

// HasDedicatedManagementEnabled returns a boolean if a field has been set.
func (o *InterfacesService) HasDedicatedManagementEnabled() bool {
	if o != nil && !IsNil(o.DedicatedManagementEnabled) {
		return true
	}

	return false
}

// SetDedicatedManagementEnabled gets a reference to the given bool and assigns it to the DedicatedManagementEnabled field.
func (o *InterfacesService) SetDedicatedManagementEnabled(v bool) {
	o.DedicatedManagementEnabled = &v
}

// GetInterfaces returns the Interfaces field value if set, zero value otherwise.
func (o *InterfacesService) GetInterfaces() []InterfaceBean {
	if o == nil || IsNil(o.Interfaces) {
		var ret []InterfaceBean
		return ret
	}
	return o.Interfaces
}

// GetInterfacesOk returns a tuple with the Interfaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfacesService) GetInterfacesOk() ([]InterfaceBean, bool) {
	if o == nil || IsNil(o.Interfaces) {
		return nil, false
	}
	return o.Interfaces, true
}

// HasInterfaces returns a boolean if a field has been set.
func (o *InterfacesService) HasInterfaces() bool {
	if o != nil && !IsNil(o.Interfaces) {
		return true
	}

	return false
}

// SetInterfaces gets a reference to the given []InterfaceBean and assigns it to the Interfaces field.
func (o *InterfacesService) SetInterfaces(v []InterfaceBean) {
	o.Interfaces = v
}

// GetRoutes returns the Routes field value if set, zero value otherwise.
func (o *InterfacesService) GetRoutes() []RouteBean {
	if o == nil || IsNil(o.Routes) {
		var ret []RouteBean
		return ret
	}
	return o.Routes
}

// GetRoutesOk returns a tuple with the Routes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfacesService) GetRoutesOk() ([]RouteBean, bool) {
	if o == nil || IsNil(o.Routes) {
		return nil, false
	}
	return o.Routes, true
}

// HasRoutes returns a boolean if a field has been set.
func (o *InterfacesService) HasRoutes() bool {
	if o != nil && !IsNil(o.Routes) {
		return true
	}

	return false
}

// SetRoutes gets a reference to the given []RouteBean and assigns it to the Routes field.
func (o *InterfacesService) SetRoutes(v []RouteBean) {
	o.Routes = v
}

func (o InterfacesService) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InterfacesService) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.DedicatedManagementEnabled) {
		toSerialize["dedicatedManagementEnabled"] = o.DedicatedManagementEnabled
	}
	if !IsNil(o.Interfaces) {
		toSerialize["interfaces"] = o.Interfaces
	}
	if !IsNil(o.Routes) {
		toSerialize["routes"] = o.Routes
	}
	return toSerialize, nil
}

type NullableInterfacesService struct {
	value *InterfacesService
	isSet bool
}

func (v NullableInterfacesService) Get() *InterfacesService {
	return v.value
}

func (v *NullableInterfacesService) Set(val *InterfacesService) {
	v.value = val
	v.isSet = true
}

func (v NullableInterfacesService) IsSet() bool {
	return v.isSet
}

func (v *NullableInterfacesService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterfacesService(val *InterfacesService) *NullableInterfacesService {
	return &NullableInterfacesService{value: val, isSet: true}
}

func (v NullableInterfacesService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterfacesService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


