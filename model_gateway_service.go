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

// checks if the GatewayService type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GatewayService{}

// GatewayService struct for GatewayService
type GatewayService struct {
	Service
	// The resource type.
	Type *string `json:"type,omitempty"`
	// Indicates whether the Gateway service is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// The status of the Gateway service running on the DNS/DHCP Server.
	State *string `json:"state,omitempty"`
	Container *GatewayContainerBean `json:"container,omitempty"`
	// The docker username used to pull the Gateway image from the container repository.
	RepositoryUsername *string `json:"repositoryUsername,omitempty" validate:"regexp=^.*\\\\S+.*$"`
	// The docker password used to pull the image from a private repository.
	RepositoryPassword *string `json:"repositoryPassword,omitempty" validate:"regexp=^.*\\\\S+.*$"`
	// Indicates whether the Gateway image is removed upon disabling the Gateway service.
	DeleteImageOnDisable *bool `json:"deleteImageOnDisable,omitempty"`
	// Indicates whether the local volumes and the mounted data and log directories are removed upon disabling the Gateway service.
	DeleteMountPointsOnDisable *bool `json:"deleteMountPointsOnDisable,omitempty"`
}

// NewGatewayService instantiates a new GatewayService object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayService() *GatewayService {
	this := GatewayService{}
	return &this
}

// NewGatewayServiceWithDefaults instantiates a new GatewayService object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayServiceWithDefaults() *GatewayService {
	this := GatewayService{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GatewayService) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayService) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GatewayService) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GatewayService) SetType(v string) {
	o.Type = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *GatewayService) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayService) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *GatewayService) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *GatewayService) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *GatewayService) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayService) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *GatewayService) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *GatewayService) SetState(v string) {
	o.State = &v
}

// GetContainer returns the Container field value if set, zero value otherwise.
func (o *GatewayService) GetContainer() GatewayContainerBean {
	if o == nil || IsNil(o.Container) {
		var ret GatewayContainerBean
		return ret
	}
	return *o.Container
}

// GetContainerOk returns a tuple with the Container field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayService) GetContainerOk() (*GatewayContainerBean, bool) {
	if o == nil || IsNil(o.Container) {
		return nil, false
	}
	return o.Container, true
}

// HasContainer returns a boolean if a field has been set.
func (o *GatewayService) HasContainer() bool {
	if o != nil && !IsNil(o.Container) {
		return true
	}

	return false
}

// SetContainer gets a reference to the given GatewayContainerBean and assigns it to the Container field.
func (o *GatewayService) SetContainer(v GatewayContainerBean) {
	o.Container = &v
}

// GetRepositoryUsername returns the RepositoryUsername field value if set, zero value otherwise.
func (o *GatewayService) GetRepositoryUsername() string {
	if o == nil || IsNil(o.RepositoryUsername) {
		var ret string
		return ret
	}
	return *o.RepositoryUsername
}

// GetRepositoryUsernameOk returns a tuple with the RepositoryUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayService) GetRepositoryUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.RepositoryUsername) {
		return nil, false
	}
	return o.RepositoryUsername, true
}

// HasRepositoryUsername returns a boolean if a field has been set.
func (o *GatewayService) HasRepositoryUsername() bool {
	if o != nil && !IsNil(o.RepositoryUsername) {
		return true
	}

	return false
}

// SetRepositoryUsername gets a reference to the given string and assigns it to the RepositoryUsername field.
func (o *GatewayService) SetRepositoryUsername(v string) {
	o.RepositoryUsername = &v
}

// GetRepositoryPassword returns the RepositoryPassword field value if set, zero value otherwise.
func (o *GatewayService) GetRepositoryPassword() string {
	if o == nil || IsNil(o.RepositoryPassword) {
		var ret string
		return ret
	}
	return *o.RepositoryPassword
}

// GetRepositoryPasswordOk returns a tuple with the RepositoryPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayService) GetRepositoryPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.RepositoryPassword) {
		return nil, false
	}
	return o.RepositoryPassword, true
}

// HasRepositoryPassword returns a boolean if a field has been set.
func (o *GatewayService) HasRepositoryPassword() bool {
	if o != nil && !IsNil(o.RepositoryPassword) {
		return true
	}

	return false
}

// SetRepositoryPassword gets a reference to the given string and assigns it to the RepositoryPassword field.
func (o *GatewayService) SetRepositoryPassword(v string) {
	o.RepositoryPassword = &v
}

// GetDeleteImageOnDisable returns the DeleteImageOnDisable field value if set, zero value otherwise.
func (o *GatewayService) GetDeleteImageOnDisable() bool {
	if o == nil || IsNil(o.DeleteImageOnDisable) {
		var ret bool
		return ret
	}
	return *o.DeleteImageOnDisable
}

// GetDeleteImageOnDisableOk returns a tuple with the DeleteImageOnDisable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayService) GetDeleteImageOnDisableOk() (*bool, bool) {
	if o == nil || IsNil(o.DeleteImageOnDisable) {
		return nil, false
	}
	return o.DeleteImageOnDisable, true
}

// HasDeleteImageOnDisable returns a boolean if a field has been set.
func (o *GatewayService) HasDeleteImageOnDisable() bool {
	if o != nil && !IsNil(o.DeleteImageOnDisable) {
		return true
	}

	return false
}

// SetDeleteImageOnDisable gets a reference to the given bool and assigns it to the DeleteImageOnDisable field.
func (o *GatewayService) SetDeleteImageOnDisable(v bool) {
	o.DeleteImageOnDisable = &v
}

// GetDeleteMountPointsOnDisable returns the DeleteMountPointsOnDisable field value if set, zero value otherwise.
func (o *GatewayService) GetDeleteMountPointsOnDisable() bool {
	if o == nil || IsNil(o.DeleteMountPointsOnDisable) {
		var ret bool
		return ret
	}
	return *o.DeleteMountPointsOnDisable
}

// GetDeleteMountPointsOnDisableOk returns a tuple with the DeleteMountPointsOnDisable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayService) GetDeleteMountPointsOnDisableOk() (*bool, bool) {
	if o == nil || IsNil(o.DeleteMountPointsOnDisable) {
		return nil, false
	}
	return o.DeleteMountPointsOnDisable, true
}

// HasDeleteMountPointsOnDisable returns a boolean if a field has been set.
func (o *GatewayService) HasDeleteMountPointsOnDisable() bool {
	if o != nil && !IsNil(o.DeleteMountPointsOnDisable) {
		return true
	}

	return false
}

// SetDeleteMountPointsOnDisable gets a reference to the given bool and assigns it to the DeleteMountPointsOnDisable field.
func (o *GatewayService) SetDeleteMountPointsOnDisable(v bool) {
	o.DeleteMountPointsOnDisable = &v
}

func (o GatewayService) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GatewayService) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Container) {
		toSerialize["container"] = o.Container
	}
	if !IsNil(o.RepositoryUsername) {
		toSerialize["repositoryUsername"] = o.RepositoryUsername
	}
	if !IsNil(o.RepositoryPassword) {
		toSerialize["repositoryPassword"] = o.RepositoryPassword
	}
	if !IsNil(o.DeleteImageOnDisable) {
		toSerialize["deleteImageOnDisable"] = o.DeleteImageOnDisable
	}
	if !IsNil(o.DeleteMountPointsOnDisable) {
		toSerialize["deleteMountPointsOnDisable"] = o.DeleteMountPointsOnDisable
	}
	return toSerialize, nil
}

type NullableGatewayService struct {
	value *GatewayService
	isSet bool
}

func (v NullableGatewayService) Get() *GatewayService {
	return v.value
}

func (v *NullableGatewayService) Set(val *GatewayService) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayService) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayService(val *GatewayService) *NullableGatewayService {
	return &NullableGatewayService{value: val, isSet: true}
}

func (v NullableGatewayService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


