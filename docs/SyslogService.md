# SyslogService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The resource type. | [optional] 
**ManualOverrideEnabled** | Pointer to **bool** | Indicates whether the syslog redirection service configuration has been manually overridden. | [optional] [readonly] 
**Iso8601TimestampsEnabled** | Pointer to **bool** | Indicates whether the ISO 8601 date time format is used when logging timestamps. | [optional] 
**Servers** | Pointer to [**[]SyslogRedirectionServerBean**](SyslogRedirectionServerBean.md) | The list of servers to redirect syslog messages to. | [optional] 

## Methods

### NewSyslogService

`func NewSyslogService() *SyslogService`

NewSyslogService instantiates a new SyslogService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyslogServiceWithDefaults

`func NewSyslogServiceWithDefaults() *SyslogService`

NewSyslogServiceWithDefaults instantiates a new SyslogService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SyslogService) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SyslogService) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SyslogService) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SyslogService) HasType() bool`

HasType returns a boolean if a field has been set.

### GetManualOverrideEnabled

`func (o *SyslogService) GetManualOverrideEnabled() bool`

GetManualOverrideEnabled returns the ManualOverrideEnabled field if non-nil, zero value otherwise.

### GetManualOverrideEnabledOk

`func (o *SyslogService) GetManualOverrideEnabledOk() (*bool, bool)`

GetManualOverrideEnabledOk returns a tuple with the ManualOverrideEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualOverrideEnabled

`func (o *SyslogService) SetManualOverrideEnabled(v bool)`

SetManualOverrideEnabled sets ManualOverrideEnabled field to given value.

### HasManualOverrideEnabled

`func (o *SyslogService) HasManualOverrideEnabled() bool`

HasManualOverrideEnabled returns a boolean if a field has been set.

### GetIso8601TimestampsEnabled

`func (o *SyslogService) GetIso8601TimestampsEnabled() bool`

GetIso8601TimestampsEnabled returns the Iso8601TimestampsEnabled field if non-nil, zero value otherwise.

### GetIso8601TimestampsEnabledOk

`func (o *SyslogService) GetIso8601TimestampsEnabledOk() (*bool, bool)`

GetIso8601TimestampsEnabledOk returns a tuple with the Iso8601TimestampsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIso8601TimestampsEnabled

`func (o *SyslogService) SetIso8601TimestampsEnabled(v bool)`

SetIso8601TimestampsEnabled sets Iso8601TimestampsEnabled field to given value.

### HasIso8601TimestampsEnabled

`func (o *SyslogService) HasIso8601TimestampsEnabled() bool`

HasIso8601TimestampsEnabled returns a boolean if a field has been set.

### GetServers

`func (o *SyslogService) GetServers() []SyslogRedirectionServerBean`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *SyslogService) GetServersOk() (*[]SyslogRedirectionServerBean, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *SyslogService) SetServers(v []SyslogRedirectionServerBean)`

SetServers sets Servers field to given value.

### HasServers

`func (o *SyslogService) HasServers() bool`

HasServers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


