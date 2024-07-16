# MonitoringSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The resource type. | [optional] 
**AddressManagerMonitoring** | Pointer to [**AddressManagerMonitoringBean**](AddressManagerMonitoringBean.md) |  | [optional] 
**ServerMonitoring** | Pointer to [**ServerMonitoringBean**](ServerMonitoringBean.md) |  | [optional] 
**F5ServerMonitoring** | Pointer to [**F5ServerMonitoringBean**](F5ServerMonitoringBean.md) |  | [optional] 

## Methods

### NewMonitoringSettings

`func NewMonitoringSettings() *MonitoringSettings`

NewMonitoringSettings instantiates a new MonitoringSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringSettingsWithDefaults

`func NewMonitoringSettingsWithDefaults() *MonitoringSettings`

NewMonitoringSettingsWithDefaults instantiates a new MonitoringSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MonitoringSettings) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MonitoringSettings) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MonitoringSettings) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MonitoringSettings) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAddressManagerMonitoring

`func (o *MonitoringSettings) GetAddressManagerMonitoring() AddressManagerMonitoringBean`

GetAddressManagerMonitoring returns the AddressManagerMonitoring field if non-nil, zero value otherwise.

### GetAddressManagerMonitoringOk

`func (o *MonitoringSettings) GetAddressManagerMonitoringOk() (*AddressManagerMonitoringBean, bool)`

GetAddressManagerMonitoringOk returns a tuple with the AddressManagerMonitoring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressManagerMonitoring

`func (o *MonitoringSettings) SetAddressManagerMonitoring(v AddressManagerMonitoringBean)`

SetAddressManagerMonitoring sets AddressManagerMonitoring field to given value.

### HasAddressManagerMonitoring

`func (o *MonitoringSettings) HasAddressManagerMonitoring() bool`

HasAddressManagerMonitoring returns a boolean if a field has been set.

### GetServerMonitoring

`func (o *MonitoringSettings) GetServerMonitoring() ServerMonitoringBean`

GetServerMonitoring returns the ServerMonitoring field if non-nil, zero value otherwise.

### GetServerMonitoringOk

`func (o *MonitoringSettings) GetServerMonitoringOk() (*ServerMonitoringBean, bool)`

GetServerMonitoringOk returns a tuple with the ServerMonitoring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerMonitoring

`func (o *MonitoringSettings) SetServerMonitoring(v ServerMonitoringBean)`

SetServerMonitoring sets ServerMonitoring field to given value.

### HasServerMonitoring

`func (o *MonitoringSettings) HasServerMonitoring() bool`

HasServerMonitoring returns a boolean if a field has been set.

### GetF5ServerMonitoring

`func (o *MonitoringSettings) GetF5ServerMonitoring() F5ServerMonitoringBean`

GetF5ServerMonitoring returns the F5ServerMonitoring field if non-nil, zero value otherwise.

### GetF5ServerMonitoringOk

`func (o *MonitoringSettings) GetF5ServerMonitoringOk() (*F5ServerMonitoringBean, bool)`

GetF5ServerMonitoringOk returns a tuple with the F5ServerMonitoring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetF5ServerMonitoring

`func (o *MonitoringSettings) SetF5ServerMonitoring(v F5ServerMonitoringBean)`

SetF5ServerMonitoring sets F5ServerMonitoring field to given value.

### HasF5ServerMonitoring

`func (o *MonitoringSettings) HasF5ServerMonitoring() bool`

HasF5ServerMonitoring returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


