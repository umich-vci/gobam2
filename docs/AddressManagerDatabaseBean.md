# AddressManagerDatabaseBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**State** | Pointer to **string** | The current replication state of the Address Manager database. | [optional] 
**EntityCount** | Pointer to **int32** | The number of entities in the database. | [optional] [readonly] 
**Size** | Pointer to **int64** | The size of the database, in bytes. | [optional] [readonly] 
**LastBackupDateTime** | Pointer to **time.Time** | The timestamp of the last database backup. | [optional] [readonly] 
**ActiveSessions** | Pointer to **int32** | The number of users logged in to Address Manager. | [optional] [readonly] 
**ReplicationEnabled** | Pointer to **bool** | Indicates whether database replication is enabled. | [optional] 
**ReplicationRole** | Pointer to **string** | The current role of the server in database replication. | [optional] [readonly] 
**ReplicationStatus** | Pointer to **string** | The current status of the database replication service. | [optional] [readonly] 
**CompressionEnabled** | Pointer to **bool** | Indicates whether compression is enabled with database replication. | [optional] 
**LatencyWarningThreshold** | Pointer to **string** | The warning threshold latency of replication, in ISO-8601 format. When the latency between the Primary and Standby servers reaches the configured threshold, Address Manager logs a warning in syslog and sends an event to warn you that the database of the Standby server is behind the Primary server by the configured threshold or greater. | [optional] 
**LatencyCriticalThreshold** | Pointer to **string** | The critical threshold latency of replication, in ISO-8601 format. When the latency between the Primary and Standby servers reaches the configured threshold, Address Manager logs a warning in syslog and sends an event to warn you that the database of the Standby server is behind the Primary server by the configured threshold or greater. | [optional] 
**Servers** | Pointer to [**[]AddressManagerServer**](AddressManagerServer.md) |  | [optional] 

## Methods

### NewAddressManagerDatabaseBean

`func NewAddressManagerDatabaseBean() *AddressManagerDatabaseBean`

NewAddressManagerDatabaseBean instantiates a new AddressManagerDatabaseBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressManagerDatabaseBeanWithDefaults

`func NewAddressManagerDatabaseBeanWithDefaults() *AddressManagerDatabaseBean`

NewAddressManagerDatabaseBeanWithDefaults instantiates a new AddressManagerDatabaseBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddressManagerDatabaseBean) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddressManagerDatabaseBean) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddressManagerDatabaseBean) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AddressManagerDatabaseBean) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *AddressManagerDatabaseBean) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddressManagerDatabaseBean) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddressManagerDatabaseBean) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AddressManagerDatabaseBean) HasType() bool`

HasType returns a boolean if a field has been set.

### GetState

`func (o *AddressManagerDatabaseBean) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AddressManagerDatabaseBean) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AddressManagerDatabaseBean) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *AddressManagerDatabaseBean) HasState() bool`

HasState returns a boolean if a field has been set.

### GetEntityCount

`func (o *AddressManagerDatabaseBean) GetEntityCount() int32`

GetEntityCount returns the EntityCount field if non-nil, zero value otherwise.

### GetEntityCountOk

`func (o *AddressManagerDatabaseBean) GetEntityCountOk() (*int32, bool)`

GetEntityCountOk returns a tuple with the EntityCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityCount

`func (o *AddressManagerDatabaseBean) SetEntityCount(v int32)`

SetEntityCount sets EntityCount field to given value.

### HasEntityCount

`func (o *AddressManagerDatabaseBean) HasEntityCount() bool`

HasEntityCount returns a boolean if a field has been set.

### GetSize

`func (o *AddressManagerDatabaseBean) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *AddressManagerDatabaseBean) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *AddressManagerDatabaseBean) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *AddressManagerDatabaseBean) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetLastBackupDateTime

`func (o *AddressManagerDatabaseBean) GetLastBackupDateTime() time.Time`

GetLastBackupDateTime returns the LastBackupDateTime field if non-nil, zero value otherwise.

### GetLastBackupDateTimeOk

`func (o *AddressManagerDatabaseBean) GetLastBackupDateTimeOk() (*time.Time, bool)`

GetLastBackupDateTimeOk returns a tuple with the LastBackupDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBackupDateTime

`func (o *AddressManagerDatabaseBean) SetLastBackupDateTime(v time.Time)`

SetLastBackupDateTime sets LastBackupDateTime field to given value.

### HasLastBackupDateTime

`func (o *AddressManagerDatabaseBean) HasLastBackupDateTime() bool`

HasLastBackupDateTime returns a boolean if a field has been set.

### GetActiveSessions

`func (o *AddressManagerDatabaseBean) GetActiveSessions() int32`

GetActiveSessions returns the ActiveSessions field if non-nil, zero value otherwise.

### GetActiveSessionsOk

`func (o *AddressManagerDatabaseBean) GetActiveSessionsOk() (*int32, bool)`

GetActiveSessionsOk returns a tuple with the ActiveSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveSessions

`func (o *AddressManagerDatabaseBean) SetActiveSessions(v int32)`

SetActiveSessions sets ActiveSessions field to given value.

### HasActiveSessions

`func (o *AddressManagerDatabaseBean) HasActiveSessions() bool`

HasActiveSessions returns a boolean if a field has been set.

### GetReplicationEnabled

`func (o *AddressManagerDatabaseBean) GetReplicationEnabled() bool`

GetReplicationEnabled returns the ReplicationEnabled field if non-nil, zero value otherwise.

### GetReplicationEnabledOk

`func (o *AddressManagerDatabaseBean) GetReplicationEnabledOk() (*bool, bool)`

GetReplicationEnabledOk returns a tuple with the ReplicationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationEnabled

`func (o *AddressManagerDatabaseBean) SetReplicationEnabled(v bool)`

SetReplicationEnabled sets ReplicationEnabled field to given value.

### HasReplicationEnabled

`func (o *AddressManagerDatabaseBean) HasReplicationEnabled() bool`

HasReplicationEnabled returns a boolean if a field has been set.

### GetReplicationRole

`func (o *AddressManagerDatabaseBean) GetReplicationRole() string`

GetReplicationRole returns the ReplicationRole field if non-nil, zero value otherwise.

### GetReplicationRoleOk

`func (o *AddressManagerDatabaseBean) GetReplicationRoleOk() (*string, bool)`

GetReplicationRoleOk returns a tuple with the ReplicationRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationRole

`func (o *AddressManagerDatabaseBean) SetReplicationRole(v string)`

SetReplicationRole sets ReplicationRole field to given value.

### HasReplicationRole

`func (o *AddressManagerDatabaseBean) HasReplicationRole() bool`

HasReplicationRole returns a boolean if a field has been set.

### GetReplicationStatus

`func (o *AddressManagerDatabaseBean) GetReplicationStatus() string`

GetReplicationStatus returns the ReplicationStatus field if non-nil, zero value otherwise.

### GetReplicationStatusOk

`func (o *AddressManagerDatabaseBean) GetReplicationStatusOk() (*string, bool)`

GetReplicationStatusOk returns a tuple with the ReplicationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationStatus

`func (o *AddressManagerDatabaseBean) SetReplicationStatus(v string)`

SetReplicationStatus sets ReplicationStatus field to given value.

### HasReplicationStatus

`func (o *AddressManagerDatabaseBean) HasReplicationStatus() bool`

HasReplicationStatus returns a boolean if a field has been set.

### GetCompressionEnabled

`func (o *AddressManagerDatabaseBean) GetCompressionEnabled() bool`

GetCompressionEnabled returns the CompressionEnabled field if non-nil, zero value otherwise.

### GetCompressionEnabledOk

`func (o *AddressManagerDatabaseBean) GetCompressionEnabledOk() (*bool, bool)`

GetCompressionEnabledOk returns a tuple with the CompressionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressionEnabled

`func (o *AddressManagerDatabaseBean) SetCompressionEnabled(v bool)`

SetCompressionEnabled sets CompressionEnabled field to given value.

### HasCompressionEnabled

`func (o *AddressManagerDatabaseBean) HasCompressionEnabled() bool`

HasCompressionEnabled returns a boolean if a field has been set.

### GetLatencyWarningThreshold

`func (o *AddressManagerDatabaseBean) GetLatencyWarningThreshold() string`

GetLatencyWarningThreshold returns the LatencyWarningThreshold field if non-nil, zero value otherwise.

### GetLatencyWarningThresholdOk

`func (o *AddressManagerDatabaseBean) GetLatencyWarningThresholdOk() (*string, bool)`

GetLatencyWarningThresholdOk returns a tuple with the LatencyWarningThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencyWarningThreshold

`func (o *AddressManagerDatabaseBean) SetLatencyWarningThreshold(v string)`

SetLatencyWarningThreshold sets LatencyWarningThreshold field to given value.

### HasLatencyWarningThreshold

`func (o *AddressManagerDatabaseBean) HasLatencyWarningThreshold() bool`

HasLatencyWarningThreshold returns a boolean if a field has been set.

### GetLatencyCriticalThreshold

`func (o *AddressManagerDatabaseBean) GetLatencyCriticalThreshold() string`

GetLatencyCriticalThreshold returns the LatencyCriticalThreshold field if non-nil, zero value otherwise.

### GetLatencyCriticalThresholdOk

`func (o *AddressManagerDatabaseBean) GetLatencyCriticalThresholdOk() (*string, bool)`

GetLatencyCriticalThresholdOk returns a tuple with the LatencyCriticalThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencyCriticalThreshold

`func (o *AddressManagerDatabaseBean) SetLatencyCriticalThreshold(v string)`

SetLatencyCriticalThreshold sets LatencyCriticalThreshold field to given value.

### HasLatencyCriticalThreshold

`func (o *AddressManagerDatabaseBean) HasLatencyCriticalThreshold() bool`

HasLatencyCriticalThreshold returns a boolean if a field has been set.

### GetServers

`func (o *AddressManagerDatabaseBean) GetServers() []AddressManagerServer`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *AddressManagerDatabaseBean) GetServersOk() (*[]AddressManagerServer, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *AddressManagerDatabaseBean) SetServers(v []AddressManagerServer)`

SetServers sets Servers field to given value.

### HasServers

`func (o *AddressManagerDatabaseBean) HasServers() bool`

HasServers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


