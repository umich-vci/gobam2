# AddressManagerServerPostRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | Pointer to **string** | The name of the Standby Address Manager server. | [optional] [readonly] 
**Address** | **string** | The IP address of the Address Manager server. | 
**State** | Pointer to **string** | The current replication state of the Address Manager server. | [optional] 
**ReplicationStatus** | Pointer to **string** | The current status of the database replication service. | [optional] [readonly] 
**EnrollmentPercentComplete** | Pointer to **int32** | The completion status of database replication enrollment. | [optional] [readonly] 
**ReplicationLatency** | Pointer to **string** | The replication latency between the Standby and Primary server. | [optional] [readonly] 

## Methods

### NewAddressManagerServerPostRequestBody

`func NewAddressManagerServerPostRequestBody(address string, ) *AddressManagerServerPostRequestBody`

NewAddressManagerServerPostRequestBody instantiates a new AddressManagerServerPostRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressManagerServerPostRequestBodyWithDefaults

`func NewAddressManagerServerPostRequestBodyWithDefaults() *AddressManagerServerPostRequestBody`

NewAddressManagerServerPostRequestBodyWithDefaults instantiates a new AddressManagerServerPostRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddressManagerServerPostRequestBody) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddressManagerServerPostRequestBody) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddressManagerServerPostRequestBody) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AddressManagerServerPostRequestBody) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *AddressManagerServerPostRequestBody) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddressManagerServerPostRequestBody) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddressManagerServerPostRequestBody) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AddressManagerServerPostRequestBody) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *AddressManagerServerPostRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddressManagerServerPostRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddressManagerServerPostRequestBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddressManagerServerPostRequestBody) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAddress

`func (o *AddressManagerServerPostRequestBody) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AddressManagerServerPostRequestBody) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AddressManagerServerPostRequestBody) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetState

`func (o *AddressManagerServerPostRequestBody) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AddressManagerServerPostRequestBody) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AddressManagerServerPostRequestBody) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *AddressManagerServerPostRequestBody) HasState() bool`

HasState returns a boolean if a field has been set.

### GetReplicationStatus

`func (o *AddressManagerServerPostRequestBody) GetReplicationStatus() string`

GetReplicationStatus returns the ReplicationStatus field if non-nil, zero value otherwise.

### GetReplicationStatusOk

`func (o *AddressManagerServerPostRequestBody) GetReplicationStatusOk() (*string, bool)`

GetReplicationStatusOk returns a tuple with the ReplicationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationStatus

`func (o *AddressManagerServerPostRequestBody) SetReplicationStatus(v string)`

SetReplicationStatus sets ReplicationStatus field to given value.

### HasReplicationStatus

`func (o *AddressManagerServerPostRequestBody) HasReplicationStatus() bool`

HasReplicationStatus returns a boolean if a field has been set.

### GetEnrollmentPercentComplete

`func (o *AddressManagerServerPostRequestBody) GetEnrollmentPercentComplete() int32`

GetEnrollmentPercentComplete returns the EnrollmentPercentComplete field if non-nil, zero value otherwise.

### GetEnrollmentPercentCompleteOk

`func (o *AddressManagerServerPostRequestBody) GetEnrollmentPercentCompleteOk() (*int32, bool)`

GetEnrollmentPercentCompleteOk returns a tuple with the EnrollmentPercentComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentPercentComplete

`func (o *AddressManagerServerPostRequestBody) SetEnrollmentPercentComplete(v int32)`

SetEnrollmentPercentComplete sets EnrollmentPercentComplete field to given value.

### HasEnrollmentPercentComplete

`func (o *AddressManagerServerPostRequestBody) HasEnrollmentPercentComplete() bool`

HasEnrollmentPercentComplete returns a boolean if a field has been set.

### GetReplicationLatency

`func (o *AddressManagerServerPostRequestBody) GetReplicationLatency() string`

GetReplicationLatency returns the ReplicationLatency field if non-nil, zero value otherwise.

### GetReplicationLatencyOk

`func (o *AddressManagerServerPostRequestBody) GetReplicationLatencyOk() (*string, bool)`

GetReplicationLatencyOk returns a tuple with the ReplicationLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationLatency

`func (o *AddressManagerServerPostRequestBody) SetReplicationLatency(v string)`

SetReplicationLatency sets ReplicationLatency field to given value.

### HasReplicationLatency

`func (o *AddressManagerServerPostRequestBody) HasReplicationLatency() bool`

HasReplicationLatency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


