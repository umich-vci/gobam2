# AddressManagerMonitoringBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Indicates whether Address Manager monitoring service is enabled. | [optional] 
**Address** | Pointer to **string** | The IP address or fully-qualified domain name of the monitoring server. | [optional] 
**PollingInterval** | Pointer to **string** | The time between polls performed by the monitoring service to Address Manager. | [optional] 
**Snmp** | Pointer to [**SnmpBean**](SnmpBean.md) |  | [optional] 

## Methods

### NewAddressManagerMonitoringBean

`func NewAddressManagerMonitoringBean() *AddressManagerMonitoringBean`

NewAddressManagerMonitoringBean instantiates a new AddressManagerMonitoringBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressManagerMonitoringBeanWithDefaults

`func NewAddressManagerMonitoringBeanWithDefaults() *AddressManagerMonitoringBean`

NewAddressManagerMonitoringBeanWithDefaults instantiates a new AddressManagerMonitoringBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *AddressManagerMonitoringBean) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddressManagerMonitoringBean) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddressManagerMonitoringBean) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AddressManagerMonitoringBean) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAddress

`func (o *AddressManagerMonitoringBean) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AddressManagerMonitoringBean) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AddressManagerMonitoringBean) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *AddressManagerMonitoringBean) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetPollingInterval

`func (o *AddressManagerMonitoringBean) GetPollingInterval() string`

GetPollingInterval returns the PollingInterval field if non-nil, zero value otherwise.

### GetPollingIntervalOk

`func (o *AddressManagerMonitoringBean) GetPollingIntervalOk() (*string, bool)`

GetPollingIntervalOk returns a tuple with the PollingInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollingInterval

`func (o *AddressManagerMonitoringBean) SetPollingInterval(v string)`

SetPollingInterval sets PollingInterval field to given value.

### HasPollingInterval

`func (o *AddressManagerMonitoringBean) HasPollingInterval() bool`

HasPollingInterval returns a boolean if a field has been set.

### GetSnmp

`func (o *AddressManagerMonitoringBean) GetSnmp() SnmpBean`

GetSnmp returns the Snmp field if non-nil, zero value otherwise.

### GetSnmpOk

`func (o *AddressManagerMonitoringBean) GetSnmpOk() (*SnmpBean, bool)`

GetSnmpOk returns a tuple with the Snmp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmp

`func (o *AddressManagerMonitoringBean) SetSnmp(v SnmpBean)`

SetSnmp sets Snmp field to given value.

### HasSnmp

`func (o *AddressManagerMonitoringBean) HasSnmp() bool`

HasSnmp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


