# ServerMonitoringBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Indicates whether DNS/DHCP Server monitoring service is enabled. | [optional] 
**PollingInterval** | Pointer to **string** | The time between polls performed by the monitoring service to the DNS/DHCP Server. | [optional] 
**StatusRefreshInterval** | Pointer to **string** | The interval at which Address Manager updates monitoring information in the Address Manager user interface. | [optional] 
**FailureDetectionThreshold** | Pointer to **int64** | The number of consecutive failures permitted before a problem affects the monitoring statistics. | [optional] 

## Methods

### NewServerMonitoringBean

`func NewServerMonitoringBean() *ServerMonitoringBean`

NewServerMonitoringBean instantiates a new ServerMonitoringBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerMonitoringBeanWithDefaults

`func NewServerMonitoringBeanWithDefaults() *ServerMonitoringBean`

NewServerMonitoringBeanWithDefaults instantiates a new ServerMonitoringBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *ServerMonitoringBean) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ServerMonitoringBean) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ServerMonitoringBean) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ServerMonitoringBean) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPollingInterval

`func (o *ServerMonitoringBean) GetPollingInterval() string`

GetPollingInterval returns the PollingInterval field if non-nil, zero value otherwise.

### GetPollingIntervalOk

`func (o *ServerMonitoringBean) GetPollingIntervalOk() (*string, bool)`

GetPollingIntervalOk returns a tuple with the PollingInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollingInterval

`func (o *ServerMonitoringBean) SetPollingInterval(v string)`

SetPollingInterval sets PollingInterval field to given value.

### HasPollingInterval

`func (o *ServerMonitoringBean) HasPollingInterval() bool`

HasPollingInterval returns a boolean if a field has been set.

### GetStatusRefreshInterval

`func (o *ServerMonitoringBean) GetStatusRefreshInterval() string`

GetStatusRefreshInterval returns the StatusRefreshInterval field if non-nil, zero value otherwise.

### GetStatusRefreshIntervalOk

`func (o *ServerMonitoringBean) GetStatusRefreshIntervalOk() (*string, bool)`

GetStatusRefreshIntervalOk returns a tuple with the StatusRefreshInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusRefreshInterval

`func (o *ServerMonitoringBean) SetStatusRefreshInterval(v string)`

SetStatusRefreshInterval sets StatusRefreshInterval field to given value.

### HasStatusRefreshInterval

`func (o *ServerMonitoringBean) HasStatusRefreshInterval() bool`

HasStatusRefreshInterval returns a boolean if a field has been set.

### GetFailureDetectionThreshold

`func (o *ServerMonitoringBean) GetFailureDetectionThreshold() int64`

GetFailureDetectionThreshold returns the FailureDetectionThreshold field if non-nil, zero value otherwise.

### GetFailureDetectionThresholdOk

`func (o *ServerMonitoringBean) GetFailureDetectionThresholdOk() (*int64, bool)`

GetFailureDetectionThresholdOk returns a tuple with the FailureDetectionThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureDetectionThreshold

`func (o *ServerMonitoringBean) SetFailureDetectionThreshold(v int64)`

SetFailureDetectionThreshold sets FailureDetectionThreshold field to given value.

### HasFailureDetectionThreshold

`func (o *ServerMonitoringBean) HasFailureDetectionThreshold() bool`

HasFailureDetectionThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


