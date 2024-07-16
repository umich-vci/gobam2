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

// checks if the Address type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Address{}

// Address struct for Address
type Address struct {
	// The resource identifier.
	Id *int64 `json:"id,omitempty"`
	// The type of IP address.
	Type *string `json:"type,omitempty"`
	// The name of the resource.
	Name *string `json:"name,omitempty"`
	// User-defined fields set for the resource.
	UserDefinedFields *map[string]string `json:"userDefinedFields,omitempty"`
	Configuration *InlinedConfiguration `json:"configuration,omitempty"`
	// The state of the IP address.
	State *string `json:"state,omitempty"`
	Address *string `json:"address,omitempty"`
	MacAddress *AddressMacAddress `json:"macAddress,omitempty"`
	Device *InlinedDevice `json:"device,omitempty"`
	Location *InlinedLocation `json:"location,omitempty"`
	ResourceRecords []ResourceRecord `json:"resourceRecords,omitempty"`
	// The timestamp detailing when the lease for the IP address was offered.
	LeaseDateTime *time.Time `json:"leaseDateTime,omitempty"`
	// The timestamp detailing when the lease for the IP address expires.
	LeaseExpirationDateTime *time.Time `json:"leaseExpirationDateTime,omitempty"`
	// The ID of the relay agent.
	RemoteId *string `json:"remoteId,omitempty"`
}

// NewAddress instantiates a new Address object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddress() *Address {
	this := Address{}
	return &this
}

// NewAddressWithDefaults instantiates a new Address object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressWithDefaults() *Address {
	this := Address{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Address) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Address) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *Address) SetId(v int64) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Address) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Address) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Address) SetType(v string) {
	o.Type = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Address) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Address) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Address) SetName(v string) {
	o.Name = &v
}

// GetUserDefinedFields returns the UserDefinedFields field value if set, zero value otherwise.
func (o *Address) GetUserDefinedFields() map[string]string {
	if o == nil || IsNil(o.UserDefinedFields) {
		var ret map[string]string
		return ret
	}
	return *o.UserDefinedFields
}

// GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetUserDefinedFieldsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.UserDefinedFields) {
		return nil, false
	}
	return o.UserDefinedFields, true
}

// HasUserDefinedFields returns a boolean if a field has been set.
func (o *Address) HasUserDefinedFields() bool {
	if o != nil && !IsNil(o.UserDefinedFields) {
		return true
	}

	return false
}

// SetUserDefinedFields gets a reference to the given map[string]string and assigns it to the UserDefinedFields field.
func (o *Address) SetUserDefinedFields(v map[string]string) {
	o.UserDefinedFields = &v
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *Address) GetConfiguration() InlinedConfiguration {
	if o == nil || IsNil(o.Configuration) {
		var ret InlinedConfiguration
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetConfigurationOk() (*InlinedConfiguration, bool) {
	if o == nil || IsNil(o.Configuration) {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *Address) HasConfiguration() bool {
	if o != nil && !IsNil(o.Configuration) {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given InlinedConfiguration and assigns it to the Configuration field.
func (o *Address) SetConfiguration(v InlinedConfiguration) {
	o.Configuration = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *Address) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *Address) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *Address) SetState(v string) {
	o.State = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *Address) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *Address) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *Address) SetAddress(v string) {
	o.Address = &v
}

// GetMacAddress returns the MacAddress field value if set, zero value otherwise.
func (o *Address) GetMacAddress() AddressMacAddress {
	if o == nil || IsNil(o.MacAddress) {
		var ret AddressMacAddress
		return ret
	}
	return *o.MacAddress
}

// GetMacAddressOk returns a tuple with the MacAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetMacAddressOk() (*AddressMacAddress, bool) {
	if o == nil || IsNil(o.MacAddress) {
		return nil, false
	}
	return o.MacAddress, true
}

// HasMacAddress returns a boolean if a field has been set.
func (o *Address) HasMacAddress() bool {
	if o != nil && !IsNil(o.MacAddress) {
		return true
	}

	return false
}

// SetMacAddress gets a reference to the given AddressMacAddress and assigns it to the MacAddress field.
func (o *Address) SetMacAddress(v AddressMacAddress) {
	o.MacAddress = &v
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *Address) GetDevice() InlinedDevice {
	if o == nil || IsNil(o.Device) {
		var ret InlinedDevice
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetDeviceOk() (*InlinedDevice, bool) {
	if o == nil || IsNil(o.Device) {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *Address) HasDevice() bool {
	if o != nil && !IsNil(o.Device) {
		return true
	}

	return false
}

// SetDevice gets a reference to the given InlinedDevice and assigns it to the Device field.
func (o *Address) SetDevice(v InlinedDevice) {
	o.Device = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *Address) GetLocation() InlinedLocation {
	if o == nil || IsNil(o.Location) {
		var ret InlinedLocation
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetLocationOk() (*InlinedLocation, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *Address) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given InlinedLocation and assigns it to the Location field.
func (o *Address) SetLocation(v InlinedLocation) {
	o.Location = &v
}

// GetResourceRecords returns the ResourceRecords field value if set, zero value otherwise.
func (o *Address) GetResourceRecords() []ResourceRecord {
	if o == nil || IsNil(o.ResourceRecords) {
		var ret []ResourceRecord
		return ret
	}
	return o.ResourceRecords
}

// GetResourceRecordsOk returns a tuple with the ResourceRecords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetResourceRecordsOk() ([]ResourceRecord, bool) {
	if o == nil || IsNil(o.ResourceRecords) {
		return nil, false
	}
	return o.ResourceRecords, true
}

// HasResourceRecords returns a boolean if a field has been set.
func (o *Address) HasResourceRecords() bool {
	if o != nil && !IsNil(o.ResourceRecords) {
		return true
	}

	return false
}

// SetResourceRecords gets a reference to the given []ResourceRecord and assigns it to the ResourceRecords field.
func (o *Address) SetResourceRecords(v []ResourceRecord) {
	o.ResourceRecords = v
}

// GetLeaseDateTime returns the LeaseDateTime field value if set, zero value otherwise.
func (o *Address) GetLeaseDateTime() time.Time {
	if o == nil || IsNil(o.LeaseDateTime) {
		var ret time.Time
		return ret
	}
	return *o.LeaseDateTime
}

// GetLeaseDateTimeOk returns a tuple with the LeaseDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetLeaseDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LeaseDateTime) {
		return nil, false
	}
	return o.LeaseDateTime, true
}

// HasLeaseDateTime returns a boolean if a field has been set.
func (o *Address) HasLeaseDateTime() bool {
	if o != nil && !IsNil(o.LeaseDateTime) {
		return true
	}

	return false
}

// SetLeaseDateTime gets a reference to the given time.Time and assigns it to the LeaseDateTime field.
func (o *Address) SetLeaseDateTime(v time.Time) {
	o.LeaseDateTime = &v
}

// GetLeaseExpirationDateTime returns the LeaseExpirationDateTime field value if set, zero value otherwise.
func (o *Address) GetLeaseExpirationDateTime() time.Time {
	if o == nil || IsNil(o.LeaseExpirationDateTime) {
		var ret time.Time
		return ret
	}
	return *o.LeaseExpirationDateTime
}

// GetLeaseExpirationDateTimeOk returns a tuple with the LeaseExpirationDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetLeaseExpirationDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LeaseExpirationDateTime) {
		return nil, false
	}
	return o.LeaseExpirationDateTime, true
}

// HasLeaseExpirationDateTime returns a boolean if a field has been set.
func (o *Address) HasLeaseExpirationDateTime() bool {
	if o != nil && !IsNil(o.LeaseExpirationDateTime) {
		return true
	}

	return false
}

// SetLeaseExpirationDateTime gets a reference to the given time.Time and assigns it to the LeaseExpirationDateTime field.
func (o *Address) SetLeaseExpirationDateTime(v time.Time) {
	o.LeaseExpirationDateTime = &v
}

// GetRemoteId returns the RemoteId field value if set, zero value otherwise.
func (o *Address) GetRemoteId() string {
	if o == nil || IsNil(o.RemoteId) {
		var ret string
		return ret
	}
	return *o.RemoteId
}

// GetRemoteIdOk returns a tuple with the RemoteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetRemoteIdOk() (*string, bool) {
	if o == nil || IsNil(o.RemoteId) {
		return nil, false
	}
	return o.RemoteId, true
}

// HasRemoteId returns a boolean if a field has been set.
func (o *Address) HasRemoteId() bool {
	if o != nil && !IsNil(o.RemoteId) {
		return true
	}

	return false
}

// SetRemoteId gets a reference to the given string and assigns it to the RemoteId field.
func (o *Address) SetRemoteId(v string) {
	o.RemoteId = &v
}

func (o Address) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Address) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
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
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.MacAddress) {
		toSerialize["macAddress"] = o.MacAddress
	}
	if !IsNil(o.Device) {
		toSerialize["device"] = o.Device
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !IsNil(o.ResourceRecords) {
		toSerialize["resourceRecords"] = o.ResourceRecords
	}
	if !IsNil(o.LeaseDateTime) {
		toSerialize["leaseDateTime"] = o.LeaseDateTime
	}
	if !IsNil(o.LeaseExpirationDateTime) {
		toSerialize["leaseExpirationDateTime"] = o.LeaseExpirationDateTime
	}
	if !IsNil(o.RemoteId) {
		toSerialize["remoteId"] = o.RemoteId
	}
	return toSerialize, nil
}

type NullableAddress struct {
	value *Address
	isSet bool
}

func (v NullableAddress) Get() *Address {
	return v.value
}

func (v *NullableAddress) Set(val *Address) {
	v.value = val
	v.isSet = true
}

func (v NullableAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddress(val *Address) *NullableAddress {
	return &NullableAddress{value: val, isSet: true}
}

func (v NullableAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

