# NTPService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The resource type. | [optional] 
**Enabled** | Pointer to **bool** | Indicates whether the NTP service is enabled. | [optional] 
**ManualOverrideEnabled** | Pointer to **bool** | Indicates whether the NTP service configuration has been manually overridden. | [optional] [readonly] 
**Servers** | Pointer to [**[]NtpServerBean**](NtpServerBean.md) | The list of NTP servers to synchronize with. | [optional] 

## Methods

### NewNTPService

`func NewNTPService() *NTPService`

NewNTPService instantiates a new NTPService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNTPServiceWithDefaults

`func NewNTPServiceWithDefaults() *NTPService`

NewNTPServiceWithDefaults instantiates a new NTPService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *NTPService) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NTPService) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NTPService) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NTPService) HasType() bool`

HasType returns a boolean if a field has been set.

### GetEnabled

`func (o *NTPService) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *NTPService) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *NTPService) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *NTPService) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetManualOverrideEnabled

`func (o *NTPService) GetManualOverrideEnabled() bool`

GetManualOverrideEnabled returns the ManualOverrideEnabled field if non-nil, zero value otherwise.

### GetManualOverrideEnabledOk

`func (o *NTPService) GetManualOverrideEnabledOk() (*bool, bool)`

GetManualOverrideEnabledOk returns a tuple with the ManualOverrideEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualOverrideEnabled

`func (o *NTPService) SetManualOverrideEnabled(v bool)`

SetManualOverrideEnabled sets ManualOverrideEnabled field to given value.

### HasManualOverrideEnabled

`func (o *NTPService) HasManualOverrideEnabled() bool`

HasManualOverrideEnabled returns a boolean if a field has been set.

### GetServers

`func (o *NTPService) GetServers() []NtpServerBean`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *NTPService) GetServersOk() (*[]NtpServerBean, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *NTPService) SetServers(v []NtpServerBean)`

SetServers sets Servers field to given value.

### HasServers

`func (o *NTPService) HasServers() bool`

HasServers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


