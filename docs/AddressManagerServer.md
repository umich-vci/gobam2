# AddressManagerServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | Pointer to **string** | The name of the Standby Address Manager server. | [optional] [readonly] 
**Address** | Pointer to **string** | The IP address of the Address Manager server. | [optional] 
**State** | Pointer to **string** | The current replication state of the Address Manager server. | [optional] 
**ReplicationStatus** | Pointer to **string** | The current status of the database replication service. | [optional] [readonly] 
**EnrollmentPercentComplete** | Pointer to **int32** | The completion status of database replication enrollment. | [optional] [readonly] 
**ReplicationLatency** | Pointer to **string** | The replication latency between the Standby and Primary server. | [optional] [readonly] 

## Methods

### NewAddressManagerServer

`func NewAddressManagerServer() *AddressManagerServer`

NewAddressManagerServer instantiates a new AddressManagerServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressManagerServerWithDefaults

`func NewAddressManagerServerWithDefaults() *AddressManagerServer`

NewAddressManagerServerWithDefaults instantiates a new AddressManagerServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddressManagerServer) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddressManagerServer) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddressManagerServer) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AddressManagerServer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *AddressManagerServer) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddressManagerServer) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddressManagerServer) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AddressManagerServer) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *AddressManagerServer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddressManagerServer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddressManagerServer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddressManagerServer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAddress

`func (o *AddressManagerServer) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AddressManagerServer) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AddressManagerServer) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *AddressManagerServer) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetState

`func (o *AddressManagerServer) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AddressManagerServer) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AddressManagerServer) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *AddressManagerServer) HasState() bool`

HasState returns a boolean if a field has been set.

### GetReplicationStatus

`func (o *AddressManagerServer) GetReplicationStatus() string`

GetReplicationStatus returns the ReplicationStatus field if non-nil, zero value otherwise.

### GetReplicationStatusOk

`func (o *AddressManagerServer) GetReplicationStatusOk() (*string, bool)`

GetReplicationStatusOk returns a tuple with the ReplicationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationStatus

`func (o *AddressManagerServer) SetReplicationStatus(v string)`

SetReplicationStatus sets ReplicationStatus field to given value.

### HasReplicationStatus

`func (o *AddressManagerServer) HasReplicationStatus() bool`

HasReplicationStatus returns a boolean if a field has been set.

### GetEnrollmentPercentComplete

`func (o *AddressManagerServer) GetEnrollmentPercentComplete() int32`

GetEnrollmentPercentComplete returns the EnrollmentPercentComplete field if non-nil, zero value otherwise.

### GetEnrollmentPercentCompleteOk

`func (o *AddressManagerServer) GetEnrollmentPercentCompleteOk() (*int32, bool)`

GetEnrollmentPercentCompleteOk returns a tuple with the EnrollmentPercentComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentPercentComplete

`func (o *AddressManagerServer) SetEnrollmentPercentComplete(v int32)`

SetEnrollmentPercentComplete sets EnrollmentPercentComplete field to given value.

### HasEnrollmentPercentComplete

`func (o *AddressManagerServer) HasEnrollmentPercentComplete() bool`

HasEnrollmentPercentComplete returns a boolean if a field has been set.

### GetReplicationLatency

`func (o *AddressManagerServer) GetReplicationLatency() string`

GetReplicationLatency returns the ReplicationLatency field if non-nil, zero value otherwise.

### GetReplicationLatencyOk

`func (o *AddressManagerServer) GetReplicationLatencyOk() (*string, bool)`

GetReplicationLatencyOk returns a tuple with the ReplicationLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationLatency

`func (o *AddressManagerServer) SetReplicationLatency(v string)`

SetReplicationLatency sets ReplicationLatency field to given value.

### HasReplicationLatency

`func (o *AddressManagerServer) HasReplicationLatency() bool`

HasReplicationLatency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


