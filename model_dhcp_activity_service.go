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

// checks if the DHCPActivityService type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DHCPActivityService{}

// DHCPActivityService struct for DHCPActivityService
type DHCPActivityService struct {
	Service
	// The resource type.
	Type *string `json:"type,omitempty"`
	// Indicates whether the health telemetry service is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// The maximum number of events to be stored in the memory buffer
	MaximumBufferedEvents *int64 `json:"maximumBufferedEvents,omitempty"`
	Destination *SinkBean `json:"destination,omitempty"`
	Dhcpv4Enabled *bool `json:"dhcpv4Enabled,omitempty"`
	Dhcpv6Enabled *bool `json:"dhcpv6Enabled,omitempty"`
}

// NewDHCPActivityService instantiates a new DHCPActivityService object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDHCPActivityService() *DHCPActivityService {
	this := DHCPActivityService{}
	return &this
}

// NewDHCPActivityServiceWithDefaults instantiates a new DHCPActivityService object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDHCPActivityServiceWithDefaults() *DHCPActivityService {
	this := DHCPActivityService{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DHCPActivityService) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DHCPActivityService) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DHCPActivityService) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DHCPActivityService) SetType(v string) {
	o.Type = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *DHCPActivityService) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DHCPActivityService) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *DHCPActivityService) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *DHCPActivityService) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetMaximumBufferedEvents returns the MaximumBufferedEvents field value if set, zero value otherwise.
func (o *DHCPActivityService) GetMaximumBufferedEvents() int64 {
	if o == nil || IsNil(o.MaximumBufferedEvents) {
		var ret int64
		return ret
	}
	return *o.MaximumBufferedEvents
}

// GetMaximumBufferedEventsOk returns a tuple with the MaximumBufferedEvents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DHCPActivityService) GetMaximumBufferedEventsOk() (*int64, bool) {
	if o == nil || IsNil(o.MaximumBufferedEvents) {
		return nil, false
	}
	return o.MaximumBufferedEvents, true
}

// HasMaximumBufferedEvents returns a boolean if a field has been set.
func (o *DHCPActivityService) HasMaximumBufferedEvents() bool {
	if o != nil && !IsNil(o.MaximumBufferedEvents) {
		return true
	}

	return false
}

// SetMaximumBufferedEvents gets a reference to the given int64 and assigns it to the MaximumBufferedEvents field.
func (o *DHCPActivityService) SetMaximumBufferedEvents(v int64) {
	o.MaximumBufferedEvents = &v
}

// GetDestination returns the Destination field value if set, zero value otherwise.
func (o *DHCPActivityService) GetDestination() SinkBean {
	if o == nil || IsNil(o.Destination) {
		var ret SinkBean
		return ret
	}
	return *o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DHCPActivityService) GetDestinationOk() (*SinkBean, bool) {
	if o == nil || IsNil(o.Destination) {
		return nil, false
	}
	return o.Destination, true
}

// HasDestination returns a boolean if a field has been set.
func (o *DHCPActivityService) HasDestination() bool {
	if o != nil && !IsNil(o.Destination) {
		return true
	}

	return false
}

// SetDestination gets a reference to the given SinkBean and assigns it to the Destination field.
func (o *DHCPActivityService) SetDestination(v SinkBean) {
	o.Destination = &v
}

// GetDhcpv4Enabled returns the Dhcpv4Enabled field value if set, zero value otherwise.
func (o *DHCPActivityService) GetDhcpv4Enabled() bool {
	if o == nil || IsNil(o.Dhcpv4Enabled) {
		var ret bool
		return ret
	}
	return *o.Dhcpv4Enabled
}

// GetDhcpv4EnabledOk returns a tuple with the Dhcpv4Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DHCPActivityService) GetDhcpv4EnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Dhcpv4Enabled) {
		return nil, false
	}
	return o.Dhcpv4Enabled, true
}

// HasDhcpv4Enabled returns a boolean if a field has been set.
func (o *DHCPActivityService) HasDhcpv4Enabled() bool {
	if o != nil && !IsNil(o.Dhcpv4Enabled) {
		return true
	}

	return false
}

// SetDhcpv4Enabled gets a reference to the given bool and assigns it to the Dhcpv4Enabled field.
func (o *DHCPActivityService) SetDhcpv4Enabled(v bool) {
	o.Dhcpv4Enabled = &v
}

// GetDhcpv6Enabled returns the Dhcpv6Enabled field value if set, zero value otherwise.
func (o *DHCPActivityService) GetDhcpv6Enabled() bool {
	if o == nil || IsNil(o.Dhcpv6Enabled) {
		var ret bool
		return ret
	}
	return *o.Dhcpv6Enabled
}

// GetDhcpv6EnabledOk returns a tuple with the Dhcpv6Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DHCPActivityService) GetDhcpv6EnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Dhcpv6Enabled) {
		return nil, false
	}
	return o.Dhcpv6Enabled, true
}

// HasDhcpv6Enabled returns a boolean if a field has been set.
func (o *DHCPActivityService) HasDhcpv6Enabled() bool {
	if o != nil && !IsNil(o.Dhcpv6Enabled) {
		return true
	}

	return false
}

// SetDhcpv6Enabled gets a reference to the given bool and assigns it to the Dhcpv6Enabled field.
func (o *DHCPActivityService) SetDhcpv6Enabled(v bool) {
	o.Dhcpv6Enabled = &v
}

func (o DHCPActivityService) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DHCPActivityService) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.MaximumBufferedEvents) {
		toSerialize["maximumBufferedEvents"] = o.MaximumBufferedEvents
	}
	if !IsNil(o.Destination) {
		toSerialize["destination"] = o.Destination
	}
	if !IsNil(o.Dhcpv4Enabled) {
		toSerialize["dhcpv4Enabled"] = o.Dhcpv4Enabled
	}
	if !IsNil(o.Dhcpv6Enabled) {
		toSerialize["dhcpv6Enabled"] = o.Dhcpv6Enabled
	}
	return toSerialize, nil
}

type NullableDHCPActivityService struct {
	value *DHCPActivityService
	isSet bool
}

func (v NullableDHCPActivityService) Get() *DHCPActivityService {
	return v.value
}

func (v *NullableDHCPActivityService) Set(val *DHCPActivityService) {
	v.value = val
	v.isSet = true
}

func (v NullableDHCPActivityService) IsSet() bool {
	return v.isSet
}

func (v *NullableDHCPActivityService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDHCPActivityService(val *DHCPActivityService) *NullableDHCPActivityService {
	return &NullableDHCPActivityService{value: val, isSet: true}
}

func (v NullableDHCPActivityService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDHCPActivityService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


