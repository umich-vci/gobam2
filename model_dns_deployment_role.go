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

// checks if the DNSDeploymentRole type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DNSDeploymentRole{}

// DNSDeploymentRole struct for DNSDeploymentRole
type DNSDeploymentRole struct {
	DeploymentRole
	// The resource type.
	Type *string `json:"type,omitempty"`
	// The type of DNS deployment role. Roles set to NONE are not deployed, instead clearing data from the server they are applied to.
	RoleType *string `json:"roleType,omitempty"`
	View *InlinedView `json:"view,omitempty"`
	// The time-to-live (TTL) value for the name server and glue records deployed via the deployment role.
	NsRecordTtl *int64 `json:"nsRecordTtl,omitempty"`
	// The server interfaces to which the DHCP deployment role is assigned.
	Interfaces []InlinedDnsRoleServerInterface `json:"interfaces,omitempty"`
}

// NewDNSDeploymentRole instantiates a new DNSDeploymentRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDNSDeploymentRole() *DNSDeploymentRole {
	this := DNSDeploymentRole{}
	return &this
}

// NewDNSDeploymentRoleWithDefaults instantiates a new DNSDeploymentRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDNSDeploymentRoleWithDefaults() *DNSDeploymentRole {
	this := DNSDeploymentRole{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DNSDeploymentRole) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNSDeploymentRole) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DNSDeploymentRole) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DNSDeploymentRole) SetType(v string) {
	o.Type = &v
}

// GetRoleType returns the RoleType field value if set, zero value otherwise.
func (o *DNSDeploymentRole) GetRoleType() string {
	if o == nil || IsNil(o.RoleType) {
		var ret string
		return ret
	}
	return *o.RoleType
}

// GetRoleTypeOk returns a tuple with the RoleType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNSDeploymentRole) GetRoleTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RoleType) {
		return nil, false
	}
	return o.RoleType, true
}

// HasRoleType returns a boolean if a field has been set.
func (o *DNSDeploymentRole) HasRoleType() bool {
	if o != nil && !IsNil(o.RoleType) {
		return true
	}

	return false
}

// SetRoleType gets a reference to the given string and assigns it to the RoleType field.
func (o *DNSDeploymentRole) SetRoleType(v string) {
	o.RoleType = &v
}

// GetView returns the View field value if set, zero value otherwise.
func (o *DNSDeploymentRole) GetView() InlinedView {
	if o == nil || IsNil(o.View) {
		var ret InlinedView
		return ret
	}
	return *o.View
}

// GetViewOk returns a tuple with the View field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNSDeploymentRole) GetViewOk() (*InlinedView, bool) {
	if o == nil || IsNil(o.View) {
		return nil, false
	}
	return o.View, true
}

// HasView returns a boolean if a field has been set.
func (o *DNSDeploymentRole) HasView() bool {
	if o != nil && !IsNil(o.View) {
		return true
	}

	return false
}

// SetView gets a reference to the given InlinedView and assigns it to the View field.
func (o *DNSDeploymentRole) SetView(v InlinedView) {
	o.View = &v
}

// GetNsRecordTtl returns the NsRecordTtl field value if set, zero value otherwise.
func (o *DNSDeploymentRole) GetNsRecordTtl() int64 {
	if o == nil || IsNil(o.NsRecordTtl) {
		var ret int64
		return ret
	}
	return *o.NsRecordTtl
}

// GetNsRecordTtlOk returns a tuple with the NsRecordTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNSDeploymentRole) GetNsRecordTtlOk() (*int64, bool) {
	if o == nil || IsNil(o.NsRecordTtl) {
		return nil, false
	}
	return o.NsRecordTtl, true
}

// HasNsRecordTtl returns a boolean if a field has been set.
func (o *DNSDeploymentRole) HasNsRecordTtl() bool {
	if o != nil && !IsNil(o.NsRecordTtl) {
		return true
	}

	return false
}

// SetNsRecordTtl gets a reference to the given int64 and assigns it to the NsRecordTtl field.
func (o *DNSDeploymentRole) SetNsRecordTtl(v int64) {
	o.NsRecordTtl = &v
}

// GetInterfaces returns the Interfaces field value if set, zero value otherwise.
func (o *DNSDeploymentRole) GetInterfaces() []InlinedDnsRoleServerInterface {
	if o == nil || IsNil(o.Interfaces) {
		var ret []InlinedDnsRoleServerInterface
		return ret
	}
	return o.Interfaces
}

// GetInterfacesOk returns a tuple with the Interfaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNSDeploymentRole) GetInterfacesOk() ([]InlinedDnsRoleServerInterface, bool) {
	if o == nil || IsNil(o.Interfaces) {
		return nil, false
	}
	return o.Interfaces, true
}

// HasInterfaces returns a boolean if a field has been set.
func (o *DNSDeploymentRole) HasInterfaces() bool {
	if o != nil && !IsNil(o.Interfaces) {
		return true
	}

	return false
}

// SetInterfaces gets a reference to the given []InlinedDnsRoleServerInterface and assigns it to the Interfaces field.
func (o *DNSDeploymentRole) SetInterfaces(v []InlinedDnsRoleServerInterface) {
	o.Interfaces = v
}

func (o DNSDeploymentRole) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DNSDeploymentRole) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.RoleType) {
		toSerialize["roleType"] = o.RoleType
	}
	if !IsNil(o.View) {
		toSerialize["view"] = o.View
	}
	if !IsNil(o.NsRecordTtl) {
		toSerialize["nsRecordTtl"] = o.NsRecordTtl
	}
	if !IsNil(o.Interfaces) {
		toSerialize["interfaces"] = o.Interfaces
	}
	return toSerialize, nil
}

type NullableDNSDeploymentRole struct {
	value *DNSDeploymentRole
	isSet bool
}

func (v NullableDNSDeploymentRole) Get() *DNSDeploymentRole {
	return v.value
}

func (v *NullableDNSDeploymentRole) Set(val *DNSDeploymentRole) {
	v.value = val
	v.isSet = true
}

func (v NullableDNSDeploymentRole) IsSet() bool {
	return v.isSet
}

func (v *NullableDNSDeploymentRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDNSDeploymentRole(val *DNSDeploymentRole) *NullableDNSDeploymentRole {
	return &NullableDNSDeploymentRole{value: val, isSet: true}
}

func (v NullableDNSDeploymentRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDNSDeploymentRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


