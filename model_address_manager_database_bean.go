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

// checks if the AddressManagerDatabaseBean type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddressManagerDatabaseBean{}

// AddressManagerDatabaseBean struct for AddressManagerDatabaseBean
type AddressManagerDatabaseBean struct {
	// The resource identifier.
	Id *int64 `json:"id,omitempty"`
	// The resource type.
	Type *string `json:"type,omitempty"`
	// The current replication state of the Address Manager database.
	State *string `json:"state,omitempty"`
	// The number of entities in the database.
	EntityCount *int32 `json:"entityCount,omitempty"`
	// The size of the database, in bytes.
	Size *int64 `json:"size,omitempty"`
	// The timestamp of the last database backup.
	LastBackupDateTime *time.Time `json:"lastBackupDateTime,omitempty"`
	// The number of users logged in to Address Manager.
	ActiveSessions *int32 `json:"activeSessions,omitempty"`
	// Indicates whether database replication is enabled.
	ReplicationEnabled *bool `json:"replicationEnabled,omitempty"`
	// The current role of the server in database replication.
	ReplicationRole *string `json:"replicationRole,omitempty"`
	// The current status of the database replication service.
	ReplicationStatus *string `json:"replicationStatus,omitempty"`
	// Indicates whether compression is enabled with database replication.
	CompressionEnabled *bool `json:"compressionEnabled,omitempty"`
	// The warning threshold latency of replication, in ISO-8601 format. When the latency between the Primary and Standby servers reaches the configured threshold, Address Manager logs a warning in syslog and sends an event to warn you that the database of the Standby server is behind the Primary server by the configured threshold or greater.
	LatencyWarningThreshold *string `json:"latencyWarningThreshold,omitempty"`
	// The critical threshold latency of replication, in ISO-8601 format. When the latency between the Primary and Standby servers reaches the configured threshold, Address Manager logs a warning in syslog and sends an event to warn you that the database of the Standby server is behind the Primary server by the configured threshold or greater.
	LatencyCriticalThreshold *string `json:"latencyCriticalThreshold,omitempty"`
	Servers []AddressManagerServer `json:"servers,omitempty"`
}

// NewAddressManagerDatabaseBean instantiates a new AddressManagerDatabaseBean object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressManagerDatabaseBean() *AddressManagerDatabaseBean {
	this := AddressManagerDatabaseBean{}
	return &this
}

// NewAddressManagerDatabaseBeanWithDefaults instantiates a new AddressManagerDatabaseBean object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressManagerDatabaseBeanWithDefaults() *AddressManagerDatabaseBean {
	this := AddressManagerDatabaseBean{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AddressManagerDatabaseBean) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressManagerDatabaseBean) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AddressManagerDatabaseBean) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *AddressManagerDatabaseBean) SetId(v int64) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AddressManagerDatabaseBean) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressManagerDatabaseBean) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AddressManagerDatabaseBean) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AddressManagerDatabaseBean) SetType(v string) {
	o.Type = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *AddressManagerDatabaseBean) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressManagerDatabaseBean) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *AddressManagerDatabaseBean) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *AddressManagerDatabaseBean) SetState(v string) {
	o.State = &v
}

// GetEntityCount returns the EntityCount field value if set, zero value otherwise.
func (o *AddressManagerDatabaseBean) GetEntityCount() int32 {
	if o == nil || IsNil(o.EntityCount) {
		var ret int32
		return ret
	}
	return *o.EntityCount
}

// GetEntityCountOk returns a tuple with the EntityCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressManagerDatabaseBean) GetEntityCountOk() (*int32, bool) {
	if o == nil || IsNil(o.EntityCount) {
		return nil, false
	}
	return o.EntityCount, true
}

// HasEntityCount returns a boolean if a field has been set.
func (o *AddressManagerDatabaseBean) HasEntityCount() bool {
	if o != nil && !IsNil(o.EntityCount) {
		return true
	}

	return false
}

// SetEntityCount gets a reference to the given int32 and assigns it to the EntityCount field.
func (o *AddressManagerDatabaseBean) SetEntityCount(v int32) {
	o.EntityCount = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *AddressManagerDatabaseBean) GetSize() int64 {
	if o == nil || IsNil(o.Size) {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressManagerDatabaseBean) GetSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *AddressManagerDatabaseBean) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *AddressManagerDatabaseBean) SetSize(v int64) {
	o.Size = &v
}

// GetLastBackupDateTime returns the LastBackupDateTime field value if set, zero value otherwise.
func (o *AddressManagerDatabaseBean) GetLastBackupDateTime() time.Time {
	if o == nil || IsNil(o.LastBackupDateTime) {
		var ret time.Time
		return ret
	}
	return *o.LastBackupDateTime
}

// GetLastBackupDateTimeOk returns a tuple with the LastBackupDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressManagerDatabaseBean) GetLastBackupDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastBackupDateTime) {
		return nil, false
	}
	return o.LastBackupDateTime, true
}

// HasLastBackupDateTime returns a boolean if a field has been set.
func (o *AddressManagerDatabaseBean) HasLastBackupDateTime() bool {
	if o != nil && !IsNil(o.LastBackupDateTime) {
		return true
	}

	return false
}

// SetLastBackupDateTime gets a reference to the given time.Time and assigns it to the LastBackupDateTime field.
func (o *AddressManagerDatabaseBean) SetLastBackupDateTime(v time.Time) {
	o.LastBackupDateTime = &v
}

// GetActiveSessions returns the ActiveSessions field value if set, zero value otherwise.
func (o *AddressManagerDatabaseBean) GetActiveSessions() int32 {
	if o == nil || IsNil(o.ActiveSessions) {
		var ret int32
		return ret
	}
	return *o.ActiveSessions
}

// GetActiveSessionsOk returns a tuple with the ActiveSessions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressManagerDatabaseBean) GetActiveSessionsOk() (*int32, bool) {
	if o == nil || IsNil(o.ActiveSessions) {
		return nil, false
	}
	return o.ActiveSessions, true
}

// HasActiveSessions returns a boolean if a field has been set.
func (o *AddressManagerDatabaseBean) HasActiveSessions() bool {
	if o != nil && !IsNil(o.ActiveSessions) {
		return true
	}

	return false
}

// SetActiveSessions gets a reference to the given int32 and assigns it to the ActiveSessions field.
func (o *AddressManagerDatabaseBean) SetActiveSessions(v int32) {
	o.ActiveSessions = &v
}

// GetReplicationEnabled returns the ReplicationEnabled field value if set, zero value otherwise.
func (o *AddressManagerDatabaseBean) GetReplicationEnabled() bool {
	if o == nil || IsNil(o.ReplicationEnabled) {
		var ret bool
		return ret
	}
	return *o.ReplicationEnabled
}

// GetReplicationEnabledOk returns a tuple with the ReplicationEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressManagerDatabaseBean) GetReplicationEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.ReplicationEnabled) {
		return nil, false
	}
	return o.ReplicationEnabled, true
}

// HasReplicationEnabled returns a boolean if a field has been set.
func (o *AddressManagerDatabaseBean) HasReplicationEnabled() bool {
	if o != nil && !IsNil(o.ReplicationEnabled) {
		return true
	}

	return false
}

// SetReplicationEnabled gets a reference to the given bool and assigns it to the ReplicationEnabled field.
func (o *AddressManagerDatabaseBean) SetReplicationEnabled(v bool) {
	o.ReplicationEnabled = &v
}

// GetReplicationRole returns the ReplicationRole field value if set, zero value otherwise.
func (o *AddressManagerDatabaseBean) GetReplicationRole() string {
	if o == nil || IsNil(o.ReplicationRole) {
		var ret string
		return ret
	}
	return *o.ReplicationRole
}

// GetReplicationRoleOk returns a tuple with the ReplicationRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressManagerDatabaseBean) GetReplicationRoleOk() (*string, bool) {
	if o == nil || IsNil(o.ReplicationRole) {
		return nil, false
	}
	return o.ReplicationRole, true
}

// HasReplicationRole returns a boolean if a field has been set.
func (o *AddressManagerDatabaseBean) HasReplicationRole() bool {
	if o != nil && !IsNil(o.ReplicationRole) {
		return true
	}

	return false
}

// SetReplicationRole gets a reference to the given string and assigns it to the ReplicationRole field.
func (o *AddressManagerDatabaseBean) SetReplicationRole(v string) {
	o.ReplicationRole = &v
}

// GetReplicationStatus returns the ReplicationStatus field value if set, zero value otherwise.
func (o *AddressManagerDatabaseBean) GetReplicationStatus() string {
	if o == nil || IsNil(o.ReplicationStatus) {
		var ret string
		return ret
	}
	return *o.ReplicationStatus
}

// GetReplicationStatusOk returns a tuple with the ReplicationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressManagerDatabaseBean) GetReplicationStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ReplicationStatus) {
		return nil, false
	}
	return o.ReplicationStatus, true
}

// HasReplicationStatus returns a boolean if a field has been set.
func (o *AddressManagerDatabaseBean) HasReplicationStatus() bool {
	if o != nil && !IsNil(o.ReplicationStatus) {
		return true
	}

	return false
}

// SetReplicationStatus gets a reference to the given string and assigns it to the ReplicationStatus field.
func (o *AddressManagerDatabaseBean) SetReplicationStatus(v string) {
	o.ReplicationStatus = &v
}

// GetCompressionEnabled returns the CompressionEnabled field value if set, zero value otherwise.
func (o *AddressManagerDatabaseBean) GetCompressionEnabled() bool {
	if o == nil || IsNil(o.CompressionEnabled) {
		var ret bool
		return ret
	}
	return *o.CompressionEnabled
}

// GetCompressionEnabledOk returns a tuple with the CompressionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressManagerDatabaseBean) GetCompressionEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.CompressionEnabled) {
		return nil, false
	}
	return o.CompressionEnabled, true
}

// HasCompressionEnabled returns a boolean if a field has been set.
func (o *AddressManagerDatabaseBean) HasCompressionEnabled() bool {
	if o != nil && !IsNil(o.CompressionEnabled) {
		return true
	}

	return false
}

// SetCompressionEnabled gets a reference to the given bool and assigns it to the CompressionEnabled field.
func (o *AddressManagerDatabaseBean) SetCompressionEnabled(v bool) {
	o.CompressionEnabled = &v
}

// GetLatencyWarningThreshold returns the LatencyWarningThreshold field value if set, zero value otherwise.
func (o *AddressManagerDatabaseBean) GetLatencyWarningThreshold() string {
	if o == nil || IsNil(o.LatencyWarningThreshold) {
		var ret string
		return ret
	}
	return *o.LatencyWarningThreshold
}

// GetLatencyWarningThresholdOk returns a tuple with the LatencyWarningThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressManagerDatabaseBean) GetLatencyWarningThresholdOk() (*string, bool) {
	if o == nil || IsNil(o.LatencyWarningThreshold) {
		return nil, false
	}
	return o.LatencyWarningThreshold, true
}

// HasLatencyWarningThreshold returns a boolean if a field has been set.
func (o *AddressManagerDatabaseBean) HasLatencyWarningThreshold() bool {
	if o != nil && !IsNil(o.LatencyWarningThreshold) {
		return true
	}

	return false
}

// SetLatencyWarningThreshold gets a reference to the given string and assigns it to the LatencyWarningThreshold field.
func (o *AddressManagerDatabaseBean) SetLatencyWarningThreshold(v string) {
	o.LatencyWarningThreshold = &v
}

// GetLatencyCriticalThreshold returns the LatencyCriticalThreshold field value if set, zero value otherwise.
func (o *AddressManagerDatabaseBean) GetLatencyCriticalThreshold() string {
	if o == nil || IsNil(o.LatencyCriticalThreshold) {
		var ret string
		return ret
	}
	return *o.LatencyCriticalThreshold
}

// GetLatencyCriticalThresholdOk returns a tuple with the LatencyCriticalThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressManagerDatabaseBean) GetLatencyCriticalThresholdOk() (*string, bool) {
	if o == nil || IsNil(o.LatencyCriticalThreshold) {
		return nil, false
	}
	return o.LatencyCriticalThreshold, true
}

// HasLatencyCriticalThreshold returns a boolean if a field has been set.
func (o *AddressManagerDatabaseBean) HasLatencyCriticalThreshold() bool {
	if o != nil && !IsNil(o.LatencyCriticalThreshold) {
		return true
	}

	return false
}

// SetLatencyCriticalThreshold gets a reference to the given string and assigns it to the LatencyCriticalThreshold field.
func (o *AddressManagerDatabaseBean) SetLatencyCriticalThreshold(v string) {
	o.LatencyCriticalThreshold = &v
}

// GetServers returns the Servers field value if set, zero value otherwise.
func (o *AddressManagerDatabaseBean) GetServers() []AddressManagerServer {
	if o == nil || IsNil(o.Servers) {
		var ret []AddressManagerServer
		return ret
	}
	return o.Servers
}

// GetServersOk returns a tuple with the Servers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressManagerDatabaseBean) GetServersOk() ([]AddressManagerServer, bool) {
	if o == nil || IsNil(o.Servers) {
		return nil, false
	}
	return o.Servers, true
}

// HasServers returns a boolean if a field has been set.
func (o *AddressManagerDatabaseBean) HasServers() bool {
	if o != nil && !IsNil(o.Servers) {
		return true
	}

	return false
}

// SetServers gets a reference to the given []AddressManagerServer and assigns it to the Servers field.
func (o *AddressManagerDatabaseBean) SetServers(v []AddressManagerServer) {
	o.Servers = v
}

func (o AddressManagerDatabaseBean) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressManagerDatabaseBean) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.EntityCount) {
		toSerialize["entityCount"] = o.EntityCount
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !IsNil(o.LastBackupDateTime) {
		toSerialize["lastBackupDateTime"] = o.LastBackupDateTime
	}
	if !IsNil(o.ActiveSessions) {
		toSerialize["activeSessions"] = o.ActiveSessions
	}
	if !IsNil(o.ReplicationEnabled) {
		toSerialize["replicationEnabled"] = o.ReplicationEnabled
	}
	if !IsNil(o.ReplicationRole) {
		toSerialize["replicationRole"] = o.ReplicationRole
	}
	if !IsNil(o.ReplicationStatus) {
		toSerialize["replicationStatus"] = o.ReplicationStatus
	}
	if !IsNil(o.CompressionEnabled) {
		toSerialize["compressionEnabled"] = o.CompressionEnabled
	}
	if !IsNil(o.LatencyWarningThreshold) {
		toSerialize["latencyWarningThreshold"] = o.LatencyWarningThreshold
	}
	if !IsNil(o.LatencyCriticalThreshold) {
		toSerialize["latencyCriticalThreshold"] = o.LatencyCriticalThreshold
	}
	if !IsNil(o.Servers) {
		toSerialize["servers"] = o.Servers
	}
	return toSerialize, nil
}

type NullableAddressManagerDatabaseBean struct {
	value *AddressManagerDatabaseBean
	isSet bool
}

func (v NullableAddressManagerDatabaseBean) Get() *AddressManagerDatabaseBean {
	return v.value
}

func (v *NullableAddressManagerDatabaseBean) Set(val *AddressManagerDatabaseBean) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressManagerDatabaseBean) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressManagerDatabaseBean) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressManagerDatabaseBean(val *AddressManagerDatabaseBean) *NullableAddressManagerDatabaseBean {
	return &NullableAddressManagerDatabaseBean{value: val, isSet: true}
}

func (v NullableAddressManagerDatabaseBean) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressManagerDatabaseBean) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


