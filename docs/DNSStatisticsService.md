# DNSStatisticsService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The resource type. | [optional] 
**Enabled** | Pointer to **bool** | Indicates whether the health telemetry service is enabled. | [optional] 
**MaximumBufferedEvents** | Pointer to **int64** | The maximum number of events to be stored in the memory buffer | [optional] 
**Destination** | Pointer to [**SinkBean**](SinkBean.md) |  | [optional] 
**PollingInterval** | Pointer to **string** |  | [optional] 

## Methods

### NewDNSStatisticsService

`func NewDNSStatisticsService() *DNSStatisticsService`

NewDNSStatisticsService instantiates a new DNSStatisticsService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDNSStatisticsServiceWithDefaults

`func NewDNSStatisticsServiceWithDefaults() *DNSStatisticsService`

NewDNSStatisticsServiceWithDefaults instantiates a new DNSStatisticsService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DNSStatisticsService) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DNSStatisticsService) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DNSStatisticsService) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DNSStatisticsService) HasType() bool`

HasType returns a boolean if a field has been set.

### GetEnabled

`func (o *DNSStatisticsService) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DNSStatisticsService) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DNSStatisticsService) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *DNSStatisticsService) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMaximumBufferedEvents

`func (o *DNSStatisticsService) GetMaximumBufferedEvents() int64`

GetMaximumBufferedEvents returns the MaximumBufferedEvents field if non-nil, zero value otherwise.

### GetMaximumBufferedEventsOk

`func (o *DNSStatisticsService) GetMaximumBufferedEventsOk() (*int64, bool)`

GetMaximumBufferedEventsOk returns a tuple with the MaximumBufferedEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumBufferedEvents

`func (o *DNSStatisticsService) SetMaximumBufferedEvents(v int64)`

SetMaximumBufferedEvents sets MaximumBufferedEvents field to given value.

### HasMaximumBufferedEvents

`func (o *DNSStatisticsService) HasMaximumBufferedEvents() bool`

HasMaximumBufferedEvents returns a boolean if a field has been set.

### GetDestination

`func (o *DNSStatisticsService) GetDestination() SinkBean`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *DNSStatisticsService) GetDestinationOk() (*SinkBean, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *DNSStatisticsService) SetDestination(v SinkBean)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *DNSStatisticsService) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetPollingInterval

`func (o *DNSStatisticsService) GetPollingInterval() string`

GetPollingInterval returns the PollingInterval field if non-nil, zero value otherwise.

### GetPollingIntervalOk

`func (o *DNSStatisticsService) GetPollingIntervalOk() (*string, bool)`

GetPollingIntervalOk returns a tuple with the PollingInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollingInterval

`func (o *DNSStatisticsService) SetPollingInterval(v string)`

SetPollingInterval sets PollingInterval field to given value.

### HasPollingInterval

`func (o *DNSStatisticsService) HasPollingInterval() bool`

HasPollingInterval returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


