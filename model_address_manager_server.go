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

// checks if the AddressManagerServer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddressManagerServer{}

// AddressManagerServer struct for AddressManagerServer
type AddressManagerServer struct {
	// The resource identifier.
	Id *int64 `json:"id,omitempty"`
	// The resource type.
	Type *string `json:"type,omitempty"`
	// The name of the Standby Address Manager server.
	Name *string `json:"name,omitempty"`
	// The IP address of the Address Manager server.
	Address *string `json:"address,omitempty"`
	// The current replication state of the Address Manager server.
	State *string `json:"state,omitempty"`
	// The current status of the database replication service.
	ReplicationStatus *string `json:"replicationStatus,omitempty"`
	// The completion status of database replication enrollment.
	EnrollmentPercentComplete *int32 `json:"enrollmentPercentComplete,omitempty"`
	// The replication latency between the Standby and Primary server.
	ReplicationLatency *string `json:"replicationLatency,omitempty"`
}

// NewAddressManagerServer instantiates a new AddressManagerServer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressManagerServer() *AddressManagerServer {
	this := AddressManagerServer{}
	return &this
}

// NewAddressManagerServerWithDefaults instantiates a new AddressManagerServer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressManagerServerWithDefaults() *AddressManagerServer {
	this := AddressManagerServer{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AddressManagerServer) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressManagerServer) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AddressManagerServer) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *AddressManagerServer) SetId(v int64) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AddressManagerServer) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressManagerServer) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AddressManagerServer) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AddressManagerServer) SetType(v string) {
	o.Type = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AddressManagerServer) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressManagerServer) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AddressManagerServer) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AddressManagerServer) SetName(v string) {
	o.Name = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *AddressManagerServer) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressManagerServer) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *AddressManagerServer) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *AddressManagerServer) SetAddress(v string) {
	o.Address = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *AddressManagerServer) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressManagerServer) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *AddressManagerServer) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *AddressManagerServer) SetState(v string) {
	o.State = &v
}

// GetReplicationStatus returns the ReplicationStatus field value if set, zero value otherwise.
func (o *AddressManagerServer) GetReplicationStatus() string {
	if o == nil || IsNil(o.ReplicationStatus) {
		var ret string
		return ret
	}
	return *o.ReplicationStatus
}

// GetReplicationStatusOk returns a tuple with the ReplicationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressManagerServer) GetReplicationStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ReplicationStatus) {
		return nil, false
	}
	return o.ReplicationStatus, true
}

// HasReplicationStatus returns a boolean if a field has been set.
func (o *AddressManagerServer) HasReplicationStatus() bool {
	if o != nil && !IsNil(o.ReplicationStatus) {
		return true
	}

	return false
}

// SetReplicationStatus gets a reference to the given string and assigns it to the ReplicationStatus field.
func (o *AddressManagerServer) SetReplicationStatus(v string) {
	o.ReplicationStatus = &v
}

// GetEnrollmentPercentComplete returns the EnrollmentPercentComplete field value if set, zero value otherwise.
func (o *AddressManagerServer) GetEnrollmentPercentComplete() int32 {
	if o == nil || IsNil(o.EnrollmentPercentComplete) {
		var ret int32
		return ret
	}
	return *o.EnrollmentPercentComplete
}

// GetEnrollmentPercentCompleteOk returns a tuple with the EnrollmentPercentComplete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressManagerServer) GetEnrollmentPercentCompleteOk() (*int32, bool) {
	if o == nil || IsNil(o.EnrollmentPercentComplete) {
		return nil, false
	}
	return o.EnrollmentPercentComplete, true
}

// HasEnrollmentPercentComplete returns a boolean if a field has been set.
func (o *AddressManagerServer) HasEnrollmentPercentComplete() bool {
	if o != nil && !IsNil(o.EnrollmentPercentComplete) {
		return true
	}

	return false
}

// SetEnrollmentPercentComplete gets a reference to the given int32 and assigns it to the EnrollmentPercentComplete field.
func (o *AddressManagerServer) SetEnrollmentPercentComplete(v int32) {
	o.EnrollmentPercentComplete = &v
}

// GetReplicationLatency returns the ReplicationLatency field value if set, zero value otherwise.
func (o *AddressManagerServer) GetReplicationLatency() string {
	if o == nil || IsNil(o.ReplicationLatency) {
		var ret string
		return ret
	}
	return *o.ReplicationLatency
}

// GetReplicationLatencyOk returns a tuple with the ReplicationLatency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressManagerServer) GetReplicationLatencyOk() (*string, bool) {
	if o == nil || IsNil(o.ReplicationLatency) {
		return nil, false
	}
	return o.ReplicationLatency, true
}

// HasReplicationLatency returns a boolean if a field has been set.
func (o *AddressManagerServer) HasReplicationLatency() bool {
	if o != nil && !IsNil(o.ReplicationLatency) {
		return true
	}

	return false
}

// SetReplicationLatency gets a reference to the given string and assigns it to the ReplicationLatency field.
func (o *AddressManagerServer) SetReplicationLatency(v string) {
	o.ReplicationLatency = &v
}

func (o AddressManagerServer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressManagerServer) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.ReplicationStatus) {
		toSerialize["replicationStatus"] = o.ReplicationStatus
	}
	if !IsNil(o.EnrollmentPercentComplete) {
		toSerialize["enrollmentPercentComplete"] = o.EnrollmentPercentComplete
	}
	if !IsNil(o.ReplicationLatency) {
		toSerialize["replicationLatency"] = o.ReplicationLatency
	}
	return toSerialize, nil
}

type NullableAddressManagerServer struct {
	value *AddressManagerServer
	isSet bool
}

func (v NullableAddressManagerServer) Get() *AddressManagerServer {
	return v.value
}

func (v *NullableAddressManagerServer) Set(val *AddressManagerServer) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressManagerServer) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressManagerServer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressManagerServer(val *AddressManagerServer) *NullableAddressManagerServer {
	return &NullableAddressManagerServer{value: val, isSet: true}
}

func (v NullableAddressManagerServer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressManagerServer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


