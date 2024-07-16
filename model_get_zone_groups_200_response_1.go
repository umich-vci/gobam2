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

// checks if the GetZoneGroups200Response1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetZoneGroups200Response1{}

// GetZoneGroups200Response1 struct for GetZoneGroups200Response1
type GetZoneGroups200Response1 struct {
	Count *int32 `json:"count,omitempty"`
	TotalCount *int32 `json:"totalCount,omitempty"`
	Data []DHCPZoneGroup `json:"data,omitempty"`
}

// NewGetZoneGroups200Response1 instantiates a new GetZoneGroups200Response1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetZoneGroups200Response1() *GetZoneGroups200Response1 {
	this := GetZoneGroups200Response1{}
	return &this
}

// NewGetZoneGroups200Response1WithDefaults instantiates a new GetZoneGroups200Response1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetZoneGroups200Response1WithDefaults() *GetZoneGroups200Response1 {
	this := GetZoneGroups200Response1{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *GetZoneGroups200Response1) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetZoneGroups200Response1) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *GetZoneGroups200Response1) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *GetZoneGroups200Response1) SetCount(v int32) {
	o.Count = &v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *GetZoneGroups200Response1) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetZoneGroups200Response1) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *GetZoneGroups200Response1) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *GetZoneGroups200Response1) SetTotalCount(v int32) {
	o.TotalCount = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetZoneGroups200Response1) GetData() []DHCPZoneGroup {
	if o == nil || IsNil(o.Data) {
		var ret []DHCPZoneGroup
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetZoneGroups200Response1) GetDataOk() ([]DHCPZoneGroup, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetZoneGroups200Response1) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []DHCPZoneGroup and assigns it to the Data field.
func (o *GetZoneGroups200Response1) SetData(v []DHCPZoneGroup) {
	o.Data = v
}

func (o GetZoneGroups200Response1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetZoneGroups200Response1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGetZoneGroups200Response1 struct {
	value *GetZoneGroups200Response1
	isSet bool
}

func (v NullableGetZoneGroups200Response1) Get() *GetZoneGroups200Response1 {
	return v.value
}

func (v *NullableGetZoneGroups200Response1) Set(val *GetZoneGroups200Response1) {
	v.value = val
	v.isSet = true
}

func (v NullableGetZoneGroups200Response1) IsSet() bool {
	return v.isSet
}

func (v *NullableGetZoneGroups200Response1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetZoneGroups200Response1(val *GetZoneGroups200Response1) *NullableGetZoneGroups200Response1 {
	return &NullableGetZoneGroups200Response1{value: val, isSet: true}
}

func (v NullableGetZoneGroups200Response1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetZoneGroups200Response1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


