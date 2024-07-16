# NtpServerBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Server** | Pointer to **string** | The FQDN or IP address of the remote NTP server. | [optional] 
**Stratum** | Pointer to **int32** | The stratum value for the NTP server. The value will be associated to an individual NTP server. Stratum values indicate the hierarchy level for the NTP server, which is the number of servers to a reference clock used to avoid synchronization loops by preferring servers with a lower stratum. | [optional] 

## Methods

### NewNtpServerBean

`func NewNtpServerBean() *NtpServerBean`

NewNtpServerBean instantiates a new NtpServerBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNtpServerBeanWithDefaults

`func NewNtpServerBeanWithDefaults() *NtpServerBean`

NewNtpServerBeanWithDefaults instantiates a new NtpServerBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServer

`func (o *NtpServerBean) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *NtpServerBean) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *NtpServerBean) SetServer(v string)`

SetServer sets Server field to given value.

### HasServer

`func (o *NtpServerBean) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetStratum

`func (o *NtpServerBean) GetStratum() int32`

GetStratum returns the Stratum field if non-nil, zero value otherwise.

### GetStratumOk

`func (o *NtpServerBean) GetStratumOk() (*int32, bool)`

GetStratumOk returns a tuple with the Stratum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStratum

`func (o *NtpServerBean) SetStratum(v int32)`

SetStratum sets Stratum field to given value.

### HasStratum

`func (o *NtpServerBean) HasStratum() bool`

HasStratum returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


