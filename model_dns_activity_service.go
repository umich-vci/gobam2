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

// checks if the DNSActivityService type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DNSActivityService{}

// DNSActivityService struct for DNSActivityService
type DNSActivityService struct {
	Service
	// The resource type.
	Type *string `json:"type,omitempty"`
	// Indicates whether the health telemetry service is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// The maximum number of events to be stored in the memory buffer
	MaximumBufferedEvents *int64 `json:"maximumBufferedEvents,omitempty"`
	Destination *SinkBean `json:"destination,omitempty"`
	Filters []DnsActivityFilterBean `json:"filters,omitempty"`
}

// NewDNSActivityService instantiates a new DNSActivityService object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDNSActivityService() *DNSActivityService {
	this := DNSActivityService{}
	return &this
}

// NewDNSActivityServiceWithDefaults instantiates a new DNSActivityService object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDNSActivityServiceWithDefaults() *DNSActivityService {
	this := DNSActivityService{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DNSActivityService) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNSActivityService) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DNSActivityService) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DNSActivityService) SetType(v string) {
	o.Type = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *DNSActivityService) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNSActivityService) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *DNSActivityService) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *DNSActivityService) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetMaximumBufferedEvents returns the MaximumBufferedEvents field value if set, zero value otherwise.
func (o *DNSActivityService) GetMaximumBufferedEvents() int64 {
	if o == nil || IsNil(o.MaximumBufferedEvents) {
		var ret int64
		return ret
	}
	return *o.MaximumBufferedEvents
}

// GetMaximumBufferedEventsOk returns a tuple with the MaximumBufferedEvents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNSActivityService) GetMaximumBufferedEventsOk() (*int64, bool) {
	if o == nil || IsNil(o.MaximumBufferedEvents) {
		return nil, false
	}
	return o.MaximumBufferedEvents, true
}

// HasMaximumBufferedEvents returns a boolean if a field has been set.
func (o *DNSActivityService) HasMaximumBufferedEvents() bool {
	if o != nil && !IsNil(o.MaximumBufferedEvents) {
		return true
	}

	return false
}

// SetMaximumBufferedEvents gets a reference to the given int64 and assigns it to the MaximumBufferedEvents field.
func (o *DNSActivityService) SetMaximumBufferedEvents(v int64) {
	o.MaximumBufferedEvents = &v
}

// GetDestination returns the Destination field value if set, zero value otherwise.
func (o *DNSActivityService) GetDestination() SinkBean {
	if o == nil || IsNil(o.Destination) {
		var ret SinkBean
		return ret
	}
	return *o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNSActivityService) GetDestinationOk() (*SinkBean, bool) {
	if o == nil || IsNil(o.Destination) {
		return nil, false
	}
	return o.Destination, true
}

// HasDestination returns a boolean if a field has been set.
func (o *DNSActivityService) HasDestination() bool {
	if o != nil && !IsNil(o.Destination) {
		return true
	}

	return false
}

// SetDestination gets a reference to the given SinkBean and assigns it to the Destination field.
func (o *DNSActivityService) SetDestination(v SinkBean) {
	o.Destination = &v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *DNSActivityService) GetFilters() []DnsActivityFilterBean {
	if o == nil || IsNil(o.Filters) {
		var ret []DnsActivityFilterBean
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNSActivityService) GetFiltersOk() ([]DnsActivityFilterBean, bool) {
	if o == nil || IsNil(o.Filters) {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *DNSActivityService) HasFilters() bool {
	if o != nil && !IsNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given []DnsActivityFilterBean and assigns it to the Filters field.
func (o *DNSActivityService) SetFilters(v []DnsActivityFilterBean) {
	o.Filters = v
}

func (o DNSActivityService) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DNSActivityService) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Filters) {
		toSerialize["filters"] = o.Filters
	}
	return toSerialize, nil
}

type NullableDNSActivityService struct {
	value *DNSActivityService
	isSet bool
}

func (v NullableDNSActivityService) Get() *DNSActivityService {
	return v.value
}

func (v *NullableDNSActivityService) Set(val *DNSActivityService) {
	v.value = val
	v.isSet = true
}

func (v NullableDNSActivityService) IsSet() bool {
	return v.isSet
}

func (v *NullableDNSActivityService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDNSActivityService(val *DNSActivityService) *NullableDNSActivityService {
	return &NullableDNSActivityService{value: val, isSet: true}
}

func (v NullableDNSActivityService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDNSActivityService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


