/*
BlueCat Address Manager 9.5 RESTful v2 API

The **BlueCat Address Manager 9.5 RESTful v2 API** is a new RESTful API for Address Manager that presents Address Manager objects as resources. Each resource has a unique endpoint, and related resources are grouped in collections. To fetch an individual resource, a `GET` request is sent to the resource's endpoint. To fetch all resources in a collection, a `GET` request is sent to the collection's endpoint.  The RESTful v2 API is [hypermedia driven](https://en.wikipedia.org/wiki/HATEOAS) and uses the [HAL](https://en.wikipedia.org/wiki/Hypertext_Application_Language) format for representing links. When navigating through the API, you can use those links to navigate to related resources or child resources of the requested endpoint. The API supports the following media types for most endpoints:  `application/hal+json`, `application/json`, and `text/csv`.  For authentication, the API supports both `Basic` and `Bearer` HTTP authentication schemes.  To get started, create a session and receive credentials for `Basic` authentication by sending a `POST` request to the [/sessions](#/Admin%20Resources/postSession) endpoint, with a message body containing the user's credentials:  ```{\"username\":\"apiuser\", \"password\":\"apipass\"}```  The response will contain an `apiToken` field that can be entered with the username in the Swagger UI **Authorize** dialog. The response will also contain a `basicAuthenticationCredentials` field containing a base64 encoding of the requester's username and API token delimited by a colon. This string can be injected directly into request `Authorization` headers in the following format:  ```Authorization: Basic YXBpOlQ0OExOT3VIRGhDcnVBNEo1bGFES3JuS3hTZC9QK3pjczZXTzBJMDY=```  For full details on API format and authentication methods, refer to the Address Manager RESTful v2 API Guide on the BlueCat Documentation Portal.

API version: 9.5.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gobam2

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the HostRecordAllOfViews type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HostRecordAllOfViews{}

// HostRecordAllOfViews struct for HostRecordAllOfViews
type HostRecordAllOfViews struct {
	// The resource identifier.
	Id int64 `json:"id"`
	// The resource type.
	Type *string `json:"type,omitempty"`
	// The name of the view.
	Name *string `json:"name,omitempty" validate:"regexp=^.*\\\\S+.*$"`
	// User-defined fields set for the resource.
	UserDefinedFields *map[string]string `json:"userDefinedFields,omitempty"`
	Configuration *InlinedConfiguration `json:"configuration,omitempty"`
	// Indicates whether DNS redirection has been enabled for use with the BlueCat Device Registration Portal.
	DeviceRegistrationEnabled *bool `json:"deviceRegistrationEnabled,omitempty"`
	// The IPv4 address of the BlueCat Device Registration Portal.
	DeviceRegistrationPortalAddress *string `json:"deviceRegistrationPortalAddress,omitempty"`
}

type _HostRecordAllOfViews HostRecordAllOfViews

// NewHostRecordAllOfViews instantiates a new HostRecordAllOfViews object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHostRecordAllOfViews(id int64) *HostRecordAllOfViews {
	this := HostRecordAllOfViews{}
	this.Id = id
	return &this
}

// NewHostRecordAllOfViewsWithDefaults instantiates a new HostRecordAllOfViews object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHostRecordAllOfViewsWithDefaults() *HostRecordAllOfViews {
	this := HostRecordAllOfViews{}
	return &this
}

// GetId returns the Id field value
func (o *HostRecordAllOfViews) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *HostRecordAllOfViews) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *HostRecordAllOfViews) SetId(v int64) {
	o.Id = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *HostRecordAllOfViews) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostRecordAllOfViews) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *HostRecordAllOfViews) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *HostRecordAllOfViews) SetType(v string) {
	o.Type = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *HostRecordAllOfViews) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostRecordAllOfViews) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *HostRecordAllOfViews) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *HostRecordAllOfViews) SetName(v string) {
	o.Name = &v
}

// GetUserDefinedFields returns the UserDefinedFields field value if set, zero value otherwise.
func (o *HostRecordAllOfViews) GetUserDefinedFields() map[string]string {
	if o == nil || IsNil(o.UserDefinedFields) {
		var ret map[string]string
		return ret
	}
	return *o.UserDefinedFields
}

// GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostRecordAllOfViews) GetUserDefinedFieldsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.UserDefinedFields) {
		return nil, false
	}
	return o.UserDefinedFields, true
}

// HasUserDefinedFields returns a boolean if a field has been set.
func (o *HostRecordAllOfViews) HasUserDefinedFields() bool {
	if o != nil && !IsNil(o.UserDefinedFields) {
		return true
	}

	return false
}

// SetUserDefinedFields gets a reference to the given map[string]string and assigns it to the UserDefinedFields field.
func (o *HostRecordAllOfViews) SetUserDefinedFields(v map[string]string) {
	o.UserDefinedFields = &v
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *HostRecordAllOfViews) GetConfiguration() InlinedConfiguration {
	if o == nil || IsNil(o.Configuration) {
		var ret InlinedConfiguration
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostRecordAllOfViews) GetConfigurationOk() (*InlinedConfiguration, bool) {
	if o == nil || IsNil(o.Configuration) {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *HostRecordAllOfViews) HasConfiguration() bool {
	if o != nil && !IsNil(o.Configuration) {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given InlinedConfiguration and assigns it to the Configuration field.
func (o *HostRecordAllOfViews) SetConfiguration(v InlinedConfiguration) {
	o.Configuration = &v
}

// GetDeviceRegistrationEnabled returns the DeviceRegistrationEnabled field value if set, zero value otherwise.
func (o *HostRecordAllOfViews) GetDeviceRegistrationEnabled() bool {
	if o == nil || IsNil(o.DeviceRegistrationEnabled) {
		var ret bool
		return ret
	}
	return *o.DeviceRegistrationEnabled
}

// GetDeviceRegistrationEnabledOk returns a tuple with the DeviceRegistrationEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostRecordAllOfViews) GetDeviceRegistrationEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.DeviceRegistrationEnabled) {
		return nil, false
	}
	return o.DeviceRegistrationEnabled, true
}

// HasDeviceRegistrationEnabled returns a boolean if a field has been set.
func (o *HostRecordAllOfViews) HasDeviceRegistrationEnabled() bool {
	if o != nil && !IsNil(o.DeviceRegistrationEnabled) {
		return true
	}

	return false
}

// SetDeviceRegistrationEnabled gets a reference to the given bool and assigns it to the DeviceRegistrationEnabled field.
func (o *HostRecordAllOfViews) SetDeviceRegistrationEnabled(v bool) {
	o.DeviceRegistrationEnabled = &v
}

// GetDeviceRegistrationPortalAddress returns the DeviceRegistrationPortalAddress field value if set, zero value otherwise.
func (o *HostRecordAllOfViews) GetDeviceRegistrationPortalAddress() string {
	if o == nil || IsNil(o.DeviceRegistrationPortalAddress) {
		var ret string
		return ret
	}
	return *o.DeviceRegistrationPortalAddress
}

// GetDeviceRegistrationPortalAddressOk returns a tuple with the DeviceRegistrationPortalAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostRecordAllOfViews) GetDeviceRegistrationPortalAddressOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceRegistrationPortalAddress) {
		return nil, false
	}
	return o.DeviceRegistrationPortalAddress, true
}

// HasDeviceRegistrationPortalAddress returns a boolean if a field has been set.
func (o *HostRecordAllOfViews) HasDeviceRegistrationPortalAddress() bool {
	if o != nil && !IsNil(o.DeviceRegistrationPortalAddress) {
		return true
	}

	return false
}

// SetDeviceRegistrationPortalAddress gets a reference to the given string and assigns it to the DeviceRegistrationPortalAddress field.
func (o *HostRecordAllOfViews) SetDeviceRegistrationPortalAddress(v string) {
	o.DeviceRegistrationPortalAddress = &v
}

func (o HostRecordAllOfViews) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HostRecordAllOfViews) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.UserDefinedFields) {
		toSerialize["userDefinedFields"] = o.UserDefinedFields
	}
	if !IsNil(o.Configuration) {
		toSerialize["configuration"] = o.Configuration
	}
	if !IsNil(o.DeviceRegistrationEnabled) {
		toSerialize["deviceRegistrationEnabled"] = o.DeviceRegistrationEnabled
	}
	if !IsNil(o.DeviceRegistrationPortalAddress) {
		toSerialize["deviceRegistrationPortalAddress"] = o.DeviceRegistrationPortalAddress
	}
	return toSerialize, nil
}

func (o *HostRecordAllOfViews) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varHostRecordAllOfViews := _HostRecordAllOfViews{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varHostRecordAllOfViews)

	if err != nil {
		return err
	}

	*o = HostRecordAllOfViews(varHostRecordAllOfViews)

	return err
}

type NullableHostRecordAllOfViews struct {
	value *HostRecordAllOfViews
	isSet bool
}

func (v NullableHostRecordAllOfViews) Get() *HostRecordAllOfViews {
	return v.value
}

func (v *NullableHostRecordAllOfViews) Set(val *HostRecordAllOfViews) {
	v.value = val
	v.isSet = true
}

func (v NullableHostRecordAllOfViews) IsSet() bool {
	return v.isSet
}

func (v *NullableHostRecordAllOfViews) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHostRecordAllOfViews(val *HostRecordAllOfViews) *NullableHostRecordAllOfViews {
	return &NullableHostRecordAllOfViews{value: val, isSet: true}
}

func (v NullableHostRecordAllOfViews) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHostRecordAllOfViews) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


