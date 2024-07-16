# DNSActivityService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The resource type. | [optional] 
**Enabled** | Pointer to **bool** | Indicates whether the health telemetry service is enabled. | [optional] 
**MaximumBufferedEvents** | Pointer to **int64** | The maximum number of events to be stored in the memory buffer | [optional] 
**Destination** | Pointer to [**SinkBean**](SinkBean.md) |  | [optional] 
**Filters** | Pointer to [**[]DnsActivityFilterBean**](DnsActivityFilterBean.md) |  | [optional] 

## Methods

### NewDNSActivityService

`func NewDNSActivityService() *DNSActivityService`

NewDNSActivityService instantiates a new DNSActivityService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDNSActivityServiceWithDefaults

`func NewDNSActivityServiceWithDefaults() *DNSActivityService`

NewDNSActivityServiceWithDefaults instantiates a new DNSActivityService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DNSActivityService) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DNSActivityService) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DNSActivityService) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DNSActivityService) HasType() bool`

HasType returns a boolean if a field has been set.

### GetEnabled

`func (o *DNSActivityService) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DNSActivityService) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DNSActivityService) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *DNSActivityService) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMaximumBufferedEvents

`func (o *DNSActivityService) GetMaximumBufferedEvents() int64`

GetMaximumBufferedEvents returns the MaximumBufferedEvents field if non-nil, zero value otherwise.

### GetMaximumBufferedEventsOk

`func (o *DNSActivityService) GetMaximumBufferedEventsOk() (*int64, bool)`

GetMaximumBufferedEventsOk returns a tuple with the MaximumBufferedEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumBufferedEvents

`func (o *DNSActivityService) SetMaximumBufferedEvents(v int64)`

SetMaximumBufferedEvents sets MaximumBufferedEvents field to given value.

### HasMaximumBufferedEvents

`func (o *DNSActivityService) HasMaximumBufferedEvents() bool`

HasMaximumBufferedEvents returns a boolean if a field has been set.

### GetDestination

`func (o *DNSActivityService) GetDestination() SinkBean`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *DNSActivityService) GetDestinationOk() (*SinkBean, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *DNSActivityService) SetDestination(v SinkBean)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *DNSActivityService) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetFilters

`func (o *DNSActivityService) GetFilters() []DnsActivityFilterBean`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *DNSActivityService) GetFiltersOk() (*[]DnsActivityFilterBean, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *DNSActivityService) SetFilters(v []DnsActivityFilterBean)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *DNSActivityService) HasFilters() bool`

HasFilters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


