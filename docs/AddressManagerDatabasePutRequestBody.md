# AddressManagerDatabasePutRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | **string** | The resource type. | 
**State** | Pointer to **string** | The current replication state of the Address Manager database. | [optional] 
**EntityCount** | Pointer to **int32** | The number of entities in the database. | [optional] [readonly] 
**Size** | Pointer to **int64** | The size of the database, in bytes. | [optional] [readonly] 
**LastBackupDateTime** | Pointer to **time.Time** | The timestamp of the last database backup. | [optional] [readonly] 
**ActiveSessions** | Pointer to **int32** | The number of users logged in to Address Manager. | [optional] [readonly] 
**ReplicationEnabled** | **bool** | Indicates whether database replication is enabled. | 
**ReplicationRole** | Pointer to **string** | The current role of the server in database replication. | [optional] [readonly] 
**ReplicationStatus** | Pointer to **string** | The current status of the database replication service. | [optional] [readonly] 
**CompressionEnabled** | Pointer to **bool** | Indicates whether compression is enabled with database replication. | [optional] 
**LatencyWarningThreshold** | Pointer to **string** | The warning threshold latency of replication, in ISO-8601 format. When the latency between the Primary and Standby servers reaches the configured threshold, Address Manager logs a warning in syslog and sends an event to warn you that the database of the Standby server is behind the Primary server by the configured threshold or greater. | [optional] 
**LatencyCriticalThreshold** | Pointer to **string** | The critical threshold latency of replication, in ISO-8601 format. When the latency between the Primary and Standby servers reaches the configured threshold, Address Manager logs a warning in syslog and sends an event to warn you that the database of the Standby server is behind the Primary server by the configured threshold or greater. | [optional] 
**Servers** | Pointer to [**[]AddressManagerServer**](AddressManagerServer.md) |  | [optional] 

## Methods

### NewAddressManagerDatabasePutRequestBody

`func NewAddressManagerDatabasePutRequestBody(type_ string, replicationEnabled bool, ) *AddressManagerDatabasePutRequestBody`

NewAddressManagerDatabasePutRequestBody instantiates a new AddressManagerDatabasePutRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressManagerDatabasePutRequestBodyWithDefaults

`func NewAddressManagerDatabasePutRequestBodyWithDefaults() *AddressManagerDatabasePutRequestBody`

NewAddressManagerDatabasePutRequestBodyWithDefaults instantiates a new AddressManagerDatabasePutRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddressManagerDatabasePutRequestBody) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddressManagerDatabasePutRequestBody) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddressManagerDatabasePutRequestBody) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AddressManagerDatabasePutRequestBody) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *AddressManagerDatabasePutRequestBody) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddressManagerDatabasePutRequestBody) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddressManagerDatabasePutRequestBody) SetType(v string)`

SetType sets Type field to given value.


### GetState

`func (o *AddressManagerDatabasePutRequestBody) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AddressManagerDatabasePutRequestBody) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AddressManagerDatabasePutRequestBody) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *AddressManagerDatabasePutRequestBody) HasState() bool`

HasState returns a boolean if a field has been set.

### GetEntityCount

`func (o *AddressManagerDatabasePutRequestBody) GetEntityCount() int32`

GetEntityCount returns the EntityCount field if non-nil, zero value otherwise.

### GetEntityCountOk

`func (o *AddressManagerDatabasePutRequestBody) GetEntityCountOk() (*int32, bool)`

GetEntityCountOk returns a tuple with the EntityCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityCount

`func (o *AddressManagerDatabasePutRequestBody) SetEntityCount(v int32)`

SetEntityCount sets EntityCount field to given value.

### HasEntityCount

`func (o *AddressManagerDatabasePutRequestBody) HasEntityCount() bool`

HasEntityCount returns a boolean if a field has been set.

### GetSize

`func (o *AddressManagerDatabasePutRequestBody) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *AddressManagerDatabasePutRequestBody) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *AddressManagerDatabasePutRequestBody) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *AddressManagerDatabasePutRequestBody) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetLastBackupDateTime

`func (o *AddressManagerDatabasePutRequestBody) GetLastBackupDateTime() time.Time`

GetLastBackupDateTime returns the LastBackupDateTime field if non-nil, zero value otherwise.

### GetLastBackupDateTimeOk

`func (o *AddressManagerDatabasePutRequestBody) GetLastBackupDateTimeOk() (*time.Time, bool)`

GetLastBackupDateTimeOk returns a tuple with the LastBackupDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBackupDateTime

`func (o *AddressManagerDatabasePutRequestBody) SetLastBackupDateTime(v time.Time)`

SetLastBackupDateTime sets LastBackupDateTime field to given value.

### HasLastBackupDateTime

`func (o *AddressManagerDatabasePutRequestBody) HasLastBackupDateTime() bool`

HasLastBackupDateTime returns a boolean if a field has been set.

### GetActiveSessions

`func (o *AddressManagerDatabasePutRequestBody) GetActiveSessions() int32`

GetActiveSessions returns the ActiveSessions field if non-nil, zero value otherwise.

### GetActiveSessionsOk

`func (o *AddressManagerDatabasePutRequestBody) GetActiveSessionsOk() (*int32, bool)`

GetActiveSessionsOk returns a tuple with the ActiveSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveSessions

`func (o *AddressManagerDatabasePutRequestBody) SetActiveSessions(v int32)`

SetActiveSessions sets ActiveSessions field to given value.

### HasActiveSessions

`func (o *AddressManagerDatabasePutRequestBody) HasActiveSessions() bool`

HasActiveSessions returns a boolean if a field has been set.

### GetReplicationEnabled

`func (o *AddressManagerDatabasePutRequestBody) GetReplicationEnabled() bool`

GetReplicationEnabled returns the ReplicationEnabled field if non-nil, zero value otherwise.

### GetReplicationEnabledOk

`func (o *AddressManagerDatabasePutRequestBody) GetReplicationEnabledOk() (*bool, bool)`

GetReplicationEnabledOk returns a tuple with the ReplicationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationEnabled

`func (o *AddressManagerDatabasePutRequestBody) SetReplicationEnabled(v bool)`

SetReplicationEnabled sets ReplicationEnabled field to given value.


### GetReplicationRole

`func (o *AddressManagerDatabasePutRequestBody) GetReplicationRole() string`

GetReplicationRole returns the ReplicationRole field if non-nil, zero value otherwise.

### GetReplicationRoleOk

`func (o *AddressManagerDatabasePutRequestBody) GetReplicationRoleOk() (*string, bool)`

GetReplicationRoleOk returns a tuple with the ReplicationRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationRole

`func (o *AddressManagerDatabasePutRequestBody) SetReplicationRole(v string)`

SetReplicationRole sets ReplicationRole field to given value.

### HasReplicationRole

`func (o *AddressManagerDatabasePutRequestBody) HasReplicationRole() bool`

HasReplicationRole returns a boolean if a field has been set.

### GetReplicationStatus

`func (o *AddressManagerDatabasePutRequestBody) GetReplicationStatus() string`

GetReplicationStatus returns the ReplicationStatus field if non-nil, zero value otherwise.

### GetReplicationStatusOk

`func (o *AddressManagerDatabasePutRequestBody) GetReplicationStatusOk() (*string, bool)`

GetReplicationStatusOk returns a tuple with the ReplicationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationStatus

`func (o *AddressManagerDatabasePutRequestBody) SetReplicationStatus(v string)`

SetReplicationStatus sets ReplicationStatus field to given value.

### HasReplicationStatus

`func (o *AddressManagerDatabasePutRequestBody) HasReplicationStatus() bool`

HasReplicationStatus returns a boolean if a field has been set.

### GetCompressionEnabled

`func (o *AddressManagerDatabasePutRequestBody) GetCompressionEnabled() bool`

GetCompressionEnabled returns the CompressionEnabled field if non-nil, zero value otherwise.

### GetCompressionEnabledOk

`func (o *AddressManagerDatabasePutRequestBody) GetCompressionEnabledOk() (*bool, bool)`

GetCompressionEnabledOk returns a tuple with the CompressionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressionEnabled

`func (o *AddressManagerDatabasePutRequestBody) SetCompressionEnabled(v bool)`

SetCompressionEnabled sets CompressionEnabled field to given value.

### HasCompressionEnabled

`func (o *AddressManagerDatabasePutRequestBody) HasCompressionEnabled() bool`

HasCompressionEnabled returns a boolean if a field has been set.

### GetLatencyWarningThreshold

`func (o *AddressManagerDatabasePutRequestBody) GetLatencyWarningThreshold() string`

GetLatencyWarningThreshold returns the LatencyWarningThreshold field if non-nil, zero value otherwise.

### GetLatencyWarningThresholdOk

`func (o *AddressManagerDatabasePutRequestBody) GetLatencyWarningThresholdOk() (*string, bool)`

GetLatencyWarningThresholdOk returns a tuple with the LatencyWarningThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencyWarningThreshold

`func (o *AddressManagerDatabasePutRequestBody) SetLatencyWarningThreshold(v string)`

SetLatencyWarningThreshold sets LatencyWarningThreshold field to given value.

### HasLatencyWarningThreshold

`func (o *AddressManagerDatabasePutRequestBody) HasLatencyWarningThreshold() bool`

HasLatencyWarningThreshold returns a boolean if a field has been set.

### GetLatencyCriticalThreshold

`func (o *AddressManagerDatabasePutRequestBody) GetLatencyCriticalThreshold() string`

GetLatencyCriticalThreshold returns the LatencyCriticalThreshold field if non-nil, zero value otherwise.

### GetLatencyCriticalThresholdOk

`func (o *AddressManagerDatabasePutRequestBody) GetLatencyCriticalThresholdOk() (*string, bool)`

GetLatencyCriticalThresholdOk returns a tuple with the LatencyCriticalThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencyCriticalThreshold

`func (o *AddressManagerDatabasePutRequestBody) SetLatencyCriticalThreshold(v string)`

SetLatencyCriticalThreshold sets LatencyCriticalThreshold field to given value.

### HasLatencyCriticalThreshold

`func (o *AddressManagerDatabasePutRequestBody) HasLatencyCriticalThreshold() bool`

HasLatencyCriticalThreshold returns a boolean if a field has been set.

### GetServers

`func (o *AddressManagerDatabasePutRequestBody) GetServers() []AddressManagerServer`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *AddressManagerDatabasePutRequestBody) GetServersOk() (*[]AddressManagerServer, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *AddressManagerDatabasePutRequestBody) SetServers(v []AddressManagerServer)`

SetServers sets Servers field to given value.

### HasServers

`func (o *AddressManagerDatabasePutRequestBody) HasServers() bool`

HasServers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


