# DHCPActivityService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The resource type. | [optional] 
**Enabled** | Pointer to **bool** | Indicates whether the health telemetry service is enabled. | [optional] 
**MaximumBufferedEvents** | Pointer to **int64** | The maximum number of events to be stored in the memory buffer | [optional] 
**Destination** | Pointer to [**SinkBean**](SinkBean.md) |  | [optional] 
**Dhcpv4Enabled** | Pointer to **bool** |  | [optional] 
**Dhcpv6Enabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewDHCPActivityService

`func NewDHCPActivityService() *DHCPActivityService`

NewDHCPActivityService instantiates a new DHCPActivityService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDHCPActivityServiceWithDefaults

`func NewDHCPActivityServiceWithDefaults() *DHCPActivityService`

NewDHCPActivityServiceWithDefaults instantiates a new DHCPActivityService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DHCPActivityService) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DHCPActivityService) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DHCPActivityService) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DHCPActivityService) HasType() bool`

HasType returns a boolean if a field has been set.

### GetEnabled

`func (o *DHCPActivityService) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DHCPActivityService) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DHCPActivityService) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *DHCPActivityService) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMaximumBufferedEvents

`func (o *DHCPActivityService) GetMaximumBufferedEvents() int64`

GetMaximumBufferedEvents returns the MaximumBufferedEvents field if non-nil, zero value otherwise.

### GetMaximumBufferedEventsOk

`func (o *DHCPActivityService) GetMaximumBufferedEventsOk() (*int64, bool)`

GetMaximumBufferedEventsOk returns a tuple with the MaximumBufferedEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumBufferedEvents

`func (o *DHCPActivityService) SetMaximumBufferedEvents(v int64)`

SetMaximumBufferedEvents sets MaximumBufferedEvents field to given value.

### HasMaximumBufferedEvents

`func (o *DHCPActivityService) HasMaximumBufferedEvents() bool`

HasMaximumBufferedEvents returns a boolean if a field has been set.

### GetDestination

`func (o *DHCPActivityService) GetDestination() SinkBean`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *DHCPActivityService) GetDestinationOk() (*SinkBean, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *DHCPActivityService) SetDestination(v SinkBean)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *DHCPActivityService) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetDhcpv4Enabled

`func (o *DHCPActivityService) GetDhcpv4Enabled() bool`

GetDhcpv4Enabled returns the Dhcpv4Enabled field if non-nil, zero value otherwise.

### GetDhcpv4EnabledOk

`func (o *DHCPActivityService) GetDhcpv4EnabledOk() (*bool, bool)`

GetDhcpv4EnabledOk returns a tuple with the Dhcpv4Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpv4Enabled

`func (o *DHCPActivityService) SetDhcpv4Enabled(v bool)`

SetDhcpv4Enabled sets Dhcpv4Enabled field to given value.

### HasDhcpv4Enabled

`func (o *DHCPActivityService) HasDhcpv4Enabled() bool`

HasDhcpv4Enabled returns a boolean if a field has been set.

### GetDhcpv6Enabled

`func (o *DHCPActivityService) GetDhcpv6Enabled() bool`

GetDhcpv6Enabled returns the Dhcpv6Enabled field if non-nil, zero value otherwise.

### GetDhcpv6EnabledOk

`func (o *DHCPActivityService) GetDhcpv6EnabledOk() (*bool, bool)`

GetDhcpv6EnabledOk returns a tuple with the Dhcpv6Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpv6Enabled

`func (o *DHCPActivityService) SetDhcpv6Enabled(v bool)`

SetDhcpv6Enabled sets Dhcpv6Enabled field to given value.

### HasDhcpv6Enabled

`func (o *DHCPActivityService) HasDhcpv6Enabled() bool`

HasDhcpv6Enabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


