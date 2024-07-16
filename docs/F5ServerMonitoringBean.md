# F5ServerMonitoringBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Indicates whether BIG-IP DNS/LTM monitoring service is enabled. | [optional] 
**PollingInterval** | Pointer to **string** | The time between polls performed by the monitoring service to the BIG-IP DNS/LTM server. | [optional] 

## Methods

### NewF5ServerMonitoringBean

`func NewF5ServerMonitoringBean() *F5ServerMonitoringBean`

NewF5ServerMonitoringBean instantiates a new F5ServerMonitoringBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewF5ServerMonitoringBeanWithDefaults

`func NewF5ServerMonitoringBeanWithDefaults() *F5ServerMonitoringBean`

NewF5ServerMonitoringBeanWithDefaults instantiates a new F5ServerMonitoringBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *F5ServerMonitoringBean) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *F5ServerMonitoringBean) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *F5ServerMonitoringBean) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *F5ServerMonitoringBean) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPollingInterval

`func (o *F5ServerMonitoringBean) GetPollingInterval() string`

GetPollingInterval returns the PollingInterval field if non-nil, zero value otherwise.

### GetPollingIntervalOk

`func (o *F5ServerMonitoringBean) GetPollingIntervalOk() (*string, bool)`

GetPollingIntervalOk returns a tuple with the PollingInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollingInterval

`func (o *F5ServerMonitoringBean) SetPollingInterval(v string)`

SetPollingInterval sets PollingInterval field to given value.

### HasPollingInterval

`func (o *F5ServerMonitoringBean) HasPollingInterval() bool`

HasPollingInterval returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


