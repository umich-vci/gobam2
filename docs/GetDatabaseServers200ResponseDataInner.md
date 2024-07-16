# GetDatabaseServers200ResponseDataInner

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
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 

## Methods

### NewGetDatabaseServers200ResponseDataInner

`func NewGetDatabaseServers200ResponseDataInner() *GetDatabaseServers200ResponseDataInner`

NewGetDatabaseServers200ResponseDataInner instantiates a new GetDatabaseServers200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDatabaseServers200ResponseDataInnerWithDefaults

`func NewGetDatabaseServers200ResponseDataInnerWithDefaults() *GetDatabaseServers200ResponseDataInner`

NewGetDatabaseServers200ResponseDataInnerWithDefaults instantiates a new GetDatabaseServers200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetDatabaseServers200ResponseDataInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetDatabaseServers200ResponseDataInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetDatabaseServers200ResponseDataInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetDatabaseServers200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GetDatabaseServers200ResponseDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetDatabaseServers200ResponseDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetDatabaseServers200ResponseDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetDatabaseServers200ResponseDataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *GetDatabaseServers200ResponseDataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetDatabaseServers200ResponseDataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetDatabaseServers200ResponseDataInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetDatabaseServers200ResponseDataInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAddress

`func (o *GetDatabaseServers200ResponseDataInner) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *GetDatabaseServers200ResponseDataInner) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *GetDatabaseServers200ResponseDataInner) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *GetDatabaseServers200ResponseDataInner) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetState

`func (o *GetDatabaseServers200ResponseDataInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetDatabaseServers200ResponseDataInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetDatabaseServers200ResponseDataInner) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GetDatabaseServers200ResponseDataInner) HasState() bool`

HasState returns a boolean if a field has been set.

### GetReplicationStatus

`func (o *GetDatabaseServers200ResponseDataInner) GetReplicationStatus() string`

GetReplicationStatus returns the ReplicationStatus field if non-nil, zero value otherwise.

### GetReplicationStatusOk

`func (o *GetDatabaseServers200ResponseDataInner) GetReplicationStatusOk() (*string, bool)`

GetReplicationStatusOk returns a tuple with the ReplicationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationStatus

`func (o *GetDatabaseServers200ResponseDataInner) SetReplicationStatus(v string)`

SetReplicationStatus sets ReplicationStatus field to given value.

### HasReplicationStatus

`func (o *GetDatabaseServers200ResponseDataInner) HasReplicationStatus() bool`

HasReplicationStatus returns a boolean if a field has been set.

### GetEnrollmentPercentComplete

`func (o *GetDatabaseServers200ResponseDataInner) GetEnrollmentPercentComplete() int32`

GetEnrollmentPercentComplete returns the EnrollmentPercentComplete field if non-nil, zero value otherwise.

### GetEnrollmentPercentCompleteOk

`func (o *GetDatabaseServers200ResponseDataInner) GetEnrollmentPercentCompleteOk() (*int32, bool)`

GetEnrollmentPercentCompleteOk returns a tuple with the EnrollmentPercentComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentPercentComplete

`func (o *GetDatabaseServers200ResponseDataInner) SetEnrollmentPercentComplete(v int32)`

SetEnrollmentPercentComplete sets EnrollmentPercentComplete field to given value.

### HasEnrollmentPercentComplete

`func (o *GetDatabaseServers200ResponseDataInner) HasEnrollmentPercentComplete() bool`

HasEnrollmentPercentComplete returns a boolean if a field has been set.

### GetReplicationLatency

`func (o *GetDatabaseServers200ResponseDataInner) GetReplicationLatency() string`

GetReplicationLatency returns the ReplicationLatency field if non-nil, zero value otherwise.

### GetReplicationLatencyOk

`func (o *GetDatabaseServers200ResponseDataInner) GetReplicationLatencyOk() (*string, bool)`

GetReplicationLatencyOk returns a tuple with the ReplicationLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationLatency

`func (o *GetDatabaseServers200ResponseDataInner) SetReplicationLatency(v string)`

SetReplicationLatency sets ReplicationLatency field to given value.

### HasReplicationLatency

`func (o *GetDatabaseServers200ResponseDataInner) HasReplicationLatency() bool`

HasReplicationLatency returns a boolean if a field has been set.

### GetLinks

`func (o *GetDatabaseServers200ResponseDataInner) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetDatabaseServers200ResponseDataInner) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetDatabaseServers200ResponseDataInner) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetDatabaseServers200ResponseDataInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


