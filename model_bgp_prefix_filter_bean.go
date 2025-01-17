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

// checks if the BgpPrefixFilterBean type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BgpPrefixFilterBean{}

// BgpPrefixFilterBean struct for BgpPrefixFilterBean
type BgpPrefixFilterBean struct {
	// The type of BGP prefix filter.
	FilterType *string `json:"filterType,omitempty"`
	// The action to take on the filtered prefixes.
	Action *string `json:"action,omitempty"`
	// The IP address prefix to filter.
	Prefix *string `json:"prefix,omitempty"`
}

// NewBgpPrefixFilterBean instantiates a new BgpPrefixFilterBean object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBgpPrefixFilterBean() *BgpPrefixFilterBean {
	this := BgpPrefixFilterBean{}
	return &this
}

// NewBgpPrefixFilterBeanWithDefaults instantiates a new BgpPrefixFilterBean object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBgpPrefixFilterBeanWithDefaults() *BgpPrefixFilterBean {
	this := BgpPrefixFilterBean{}
	return &this
}

// GetFilterType returns the FilterType field value if set, zero value otherwise.
func (o *BgpPrefixFilterBean) GetFilterType() string {
	if o == nil || IsNil(o.FilterType) {
		var ret string
		return ret
	}
	return *o.FilterType
}

// GetFilterTypeOk returns a tuple with the FilterType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpPrefixFilterBean) GetFilterTypeOk() (*string, bool) {
	if o == nil || IsNil(o.FilterType) {
		return nil, false
	}
	return o.FilterType, true
}

// HasFilterType returns a boolean if a field has been set.
func (o *BgpPrefixFilterBean) HasFilterType() bool {
	if o != nil && !IsNil(o.FilterType) {
		return true
	}

	return false
}

// SetFilterType gets a reference to the given string and assigns it to the FilterType field.
func (o *BgpPrefixFilterBean) SetFilterType(v string) {
	o.FilterType = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *BgpPrefixFilterBean) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpPrefixFilterBean) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *BgpPrefixFilterBean) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *BgpPrefixFilterBean) SetAction(v string) {
	o.Action = &v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *BgpPrefixFilterBean) GetPrefix() string {
	if o == nil || IsNil(o.Prefix) {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpPrefixFilterBean) GetPrefixOk() (*string, bool) {
	if o == nil || IsNil(o.Prefix) {
		return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *BgpPrefixFilterBean) HasPrefix() bool {
	if o != nil && !IsNil(o.Prefix) {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *BgpPrefixFilterBean) SetPrefix(v string) {
	o.Prefix = &v
}

func (o BgpPrefixFilterBean) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BgpPrefixFilterBean) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FilterType) {
		toSerialize["filterType"] = o.FilterType
	}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.Prefix) {
		toSerialize["prefix"] = o.Prefix
	}
	return toSerialize, nil
}

type NullableBgpPrefixFilterBean struct {
	value *BgpPrefixFilterBean
	isSet bool
}

func (v NullableBgpPrefixFilterBean) Get() *BgpPrefixFilterBean {
	return v.value
}

func (v *NullableBgpPrefixFilterBean) Set(val *BgpPrefixFilterBean) {
	v.value = val
	v.isSet = true
}

func (v NullableBgpPrefixFilterBean) IsSet() bool {
	return v.isSet
}

func (v *NullableBgpPrefixFilterBean) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBgpPrefixFilterBean(val *BgpPrefixFilterBean) *NullableBgpPrefixFilterBean {
	return &NullableBgpPrefixFilterBean{value: val, isSet: true}
}

func (v NullableBgpPrefixFilterBean) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBgpPrefixFilterBean) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


