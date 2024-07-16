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

// checks if the GetResponsePolicyImports200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetResponsePolicyImports200Response{}

// GetResponsePolicyImports200Response struct for GetResponsePolicyImports200Response
type GetResponsePolicyImports200Response struct {
	Links *CollectionHalLinks `json:"_links,omitempty"`
	Count *int32 `json:"count,omitempty"`
	TotalCount *int32 `json:"totalCount,omitempty"`
	Data []GetResponsePolicyImports200ResponseDataInner `json:"data,omitempty"`
}

// NewGetResponsePolicyImports200Response instantiates a new GetResponsePolicyImports200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetResponsePolicyImports200Response() *GetResponsePolicyImports200Response {
	this := GetResponsePolicyImports200Response{}
	return &this
}

// NewGetResponsePolicyImports200ResponseWithDefaults instantiates a new GetResponsePolicyImports200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetResponsePolicyImports200ResponseWithDefaults() *GetResponsePolicyImports200Response {
	this := GetResponsePolicyImports200Response{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GetResponsePolicyImports200Response) GetLinks() CollectionHalLinks {
	if o == nil || IsNil(o.Links) {
		var ret CollectionHalLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetResponsePolicyImports200Response) GetLinksOk() (*CollectionHalLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GetResponsePolicyImports200Response) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given CollectionHalLinks and assigns it to the Links field.
func (o *GetResponsePolicyImports200Response) SetLinks(v CollectionHalLinks) {
	o.Links = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *GetResponsePolicyImports200Response) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetResponsePolicyImports200Response) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *GetResponsePolicyImports200Response) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *GetResponsePolicyImports200Response) SetCount(v int32) {
	o.Count = &v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *GetResponsePolicyImports200Response) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetResponsePolicyImports200Response) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *GetResponsePolicyImports200Response) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *GetResponsePolicyImports200Response) SetTotalCount(v int32) {
	o.TotalCount = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetResponsePolicyImports200Response) GetData() []GetResponsePolicyImports200ResponseDataInner {
	if o == nil || IsNil(o.Data) {
		var ret []GetResponsePolicyImports200ResponseDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetResponsePolicyImports200Response) GetDataOk() ([]GetResponsePolicyImports200ResponseDataInner, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetResponsePolicyImports200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []GetResponsePolicyImports200ResponseDataInner and assigns it to the Data field.
func (o *GetResponsePolicyImports200Response) SetData(v []GetResponsePolicyImports200ResponseDataInner) {
	o.Data = v
}

func (o GetResponsePolicyImports200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetResponsePolicyImports200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
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

type NullableGetResponsePolicyImports200Response struct {
	value *GetResponsePolicyImports200Response
	isSet bool
}

func (v NullableGetResponsePolicyImports200Response) Get() *GetResponsePolicyImports200Response {
	return v.value
}

func (v *NullableGetResponsePolicyImports200Response) Set(val *GetResponsePolicyImports200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetResponsePolicyImports200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetResponsePolicyImports200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetResponsePolicyImports200Response(val *GetResponsePolicyImports200Response) *NullableGetResponsePolicyImports200Response {
	return &NullableGetResponsePolicyImports200Response{value: val, isSet: true}
}

func (v NullableGetResponsePolicyImports200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetResponsePolicyImports200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


