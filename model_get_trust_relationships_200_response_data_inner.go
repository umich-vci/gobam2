/*
BlueCat Address Manager 9.5 RESTful v2 API

The **BlueCat Address Manager 9.5 RESTful v2 API** is a new RESTful API for Address Manager that presents Address Manager objects as resources. Each resource has a unique endpoint, and related resources are grouped in collections. To fetch an individual resource, a `GET` request is sent to the resource's endpoint. To fetch all resources in a collection, a `GET` request is sent to the collection's endpoint.  The RESTful v2 API is [hypermedia driven](https://en.wikipedia.org/wiki/HATEOAS) and uses the [HAL](https://en.wikipedia.org/wiki/Hypertext_Application_Language) format for representing links. When navigating through the API, you can use those links to navigate to related resources or child resources of the requested endpoint. The API supports the following media types for most endpoints:  `application/hal+json`, `application/json`, and `text/csv`.  For authentication, the API supports both `Basic` and `Bearer` HTTP authentication schemes.  To get started, create a session and receive credentials for `Basic` authentication by sending a `POST` request to the [/sessions](#/Admin%20Resources/postSession) endpoint, with a message body containing the user's credentials:  ```{\"username\":\"apiuser\", \"password\":\"apipass\"}```  The response will contain an `apiToken` field that can be entered with the username in the Swagger UI **Authorize** dialog. The response will also contain a `basicAuthenticationCredentials` field containing a base64 encoding of the requester's username and API token delimited by a colon. This string can be injected directly into request `Authorization` headers in the following format:  ```Authorization: Basic YXBpOlQ0OExOT3VIRGhDcnVBNEo1bGFES3JuS3hTZC9QK3pjczZXTzBJMDY=```  For full details on API format and authentication methods, refer to the Address Manager RESTful v2 API Guide on the BlueCat Documentation Portal.

API version: 9.5.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gobam2

import (
	"encoding/json"
	"time"
)

// checks if the GetTrustRelationships200ResponseDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTrustRelationships200ResponseDataInner{}

// GetTrustRelationships200ResponseDataInner struct for GetTrustRelationships200ResponseDataInner
type GetTrustRelationships200ResponseDataInner struct {
	// The resource identifier.
	Id *int64 `json:"id,omitempty"`
	// The resource type.
	Type *string `json:"type,omitempty"`
	// The IP address of the remote Address Manager server that is part of the trust relationship.
	Address *string `json:"address,omitempty"`
	// The API username of the remote Address Manager server that is used to establish the trust relationship.
	Username *string `json:"username,omitempty" validate:"regexp=^.*\\\\S+.*$"`
	// The password of the API user that is used to establish the trust relationship with the remote Address Manager server.
	Password *string `json:"password,omitempty" validate:"regexp=^.*\\\\S+.*$"`
	// The timestamp of when the SSH key was created for exchange with the remote Address Manager server to establish the trust relationship.
	KeyGenerationDateTime *time.Time `json:"keyGenerationDateTime,omitempty"`
	// The state of the trust relationship.
	State *string `json:"state,omitempty"`
	Links *ResourceLinks `json:"_links,omitempty"`
}

// NewGetTrustRelationships200ResponseDataInner instantiates a new GetTrustRelationships200ResponseDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTrustRelationships200ResponseDataInner() *GetTrustRelationships200ResponseDataInner {
	this := GetTrustRelationships200ResponseDataInner{}
	return &this
}

// NewGetTrustRelationships200ResponseDataInnerWithDefaults instantiates a new GetTrustRelationships200ResponseDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTrustRelationships200ResponseDataInnerWithDefaults() *GetTrustRelationships200ResponseDataInner {
	this := GetTrustRelationships200ResponseDataInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetTrustRelationships200ResponseDataInner) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTrustRelationships200ResponseDataInner) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetTrustRelationships200ResponseDataInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *GetTrustRelationships200ResponseDataInner) SetId(v int64) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetTrustRelationships200ResponseDataInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTrustRelationships200ResponseDataInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetTrustRelationships200ResponseDataInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetTrustRelationships200ResponseDataInner) SetType(v string) {
	o.Type = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *GetTrustRelationships200ResponseDataInner) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTrustRelationships200ResponseDataInner) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *GetTrustRelationships200ResponseDataInner) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *GetTrustRelationships200ResponseDataInner) SetAddress(v string) {
	o.Address = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *GetTrustRelationships200ResponseDataInner) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTrustRelationships200ResponseDataInner) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *GetTrustRelationships200ResponseDataInner) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *GetTrustRelationships200ResponseDataInner) SetUsername(v string) {
	o.Username = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *GetTrustRelationships200ResponseDataInner) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTrustRelationships200ResponseDataInner) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *GetTrustRelationships200ResponseDataInner) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *GetTrustRelationships200ResponseDataInner) SetPassword(v string) {
	o.Password = &v
}

// GetKeyGenerationDateTime returns the KeyGenerationDateTime field value if set, zero value otherwise.
func (o *GetTrustRelationships200ResponseDataInner) GetKeyGenerationDateTime() time.Time {
	if o == nil || IsNil(o.KeyGenerationDateTime) {
		var ret time.Time
		return ret
	}
	return *o.KeyGenerationDateTime
}

// GetKeyGenerationDateTimeOk returns a tuple with the KeyGenerationDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTrustRelationships200ResponseDataInner) GetKeyGenerationDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.KeyGenerationDateTime) {
		return nil, false
	}
	return o.KeyGenerationDateTime, true
}

// HasKeyGenerationDateTime returns a boolean if a field has been set.
func (o *GetTrustRelationships200ResponseDataInner) HasKeyGenerationDateTime() bool {
	if o != nil && !IsNil(o.KeyGenerationDateTime) {
		return true
	}

	return false
}

// SetKeyGenerationDateTime gets a reference to the given time.Time and assigns it to the KeyGenerationDateTime field.
func (o *GetTrustRelationships200ResponseDataInner) SetKeyGenerationDateTime(v time.Time) {
	o.KeyGenerationDateTime = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *GetTrustRelationships200ResponseDataInner) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTrustRelationships200ResponseDataInner) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *GetTrustRelationships200ResponseDataInner) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *GetTrustRelationships200ResponseDataInner) SetState(v string) {
	o.State = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GetTrustRelationships200ResponseDataInner) GetLinks() ResourceLinks {
	if o == nil || IsNil(o.Links) {
		var ret ResourceLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTrustRelationships200ResponseDataInner) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GetTrustRelationships200ResponseDataInner) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ResourceLinks and assigns it to the Links field.
func (o *GetTrustRelationships200ResponseDataInner) SetLinks(v ResourceLinks) {
	o.Links = &v
}

func (o GetTrustRelationships200ResponseDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTrustRelationships200ResponseDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.KeyGenerationDateTime) {
		toSerialize["keyGenerationDateTime"] = o.KeyGenerationDateTime
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableGetTrustRelationships200ResponseDataInner struct {
	value *GetTrustRelationships200ResponseDataInner
	isSet bool
}

func (v NullableGetTrustRelationships200ResponseDataInner) Get() *GetTrustRelationships200ResponseDataInner {
	return v.value
}

func (v *NullableGetTrustRelationships200ResponseDataInner) Set(val *GetTrustRelationships200ResponseDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTrustRelationships200ResponseDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTrustRelationships200ResponseDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTrustRelationships200ResponseDataInner(val *GetTrustRelationships200ResponseDataInner) *NullableGetTrustRelationships200ResponseDataInner {
	return &NullableGetTrustRelationships200ResponseDataInner{value: val, isSet: true}
}

func (v NullableGetTrustRelationships200ResponseDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTrustRelationships200ResponseDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


