# SystemSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The resource type. | [optional] 
**Hostname** | Pointer to **string** | The hostname of the Address Manager server. | [optional] [readonly] 
**Version** | Pointer to **string** | The Address Manager server version. | [optional] [readonly] 
**Address** | Pointer to **string** | The management IP address of the Address Manager server. | [optional] [readonly] 
**InterfaceRedundancyEnabled** | Pointer to **bool** | Indicates whether Address Manager Active/Backup (eth0/eth1) bonding mode is active. | [optional] 
**ActiveSessions** | Pointer to **int32** | The current number of active Address Manager sessions. | [optional] [readonly] 

## Methods

### NewSystemSettings

`func NewSystemSettings() *SystemSettings`

NewSystemSettings instantiates a new SystemSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemSettingsWithDefaults

`func NewSystemSettingsWithDefaults() *SystemSettings`

NewSystemSettingsWithDefaults instantiates a new SystemSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SystemSettings) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SystemSettings) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SystemSettings) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SystemSettings) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHostname

`func (o *SystemSettings) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *SystemSettings) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *SystemSettings) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *SystemSettings) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetVersion

`func (o *SystemSettings) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SystemSettings) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SystemSettings) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SystemSettings) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetAddress

`func (o *SystemSettings) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *SystemSettings) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *SystemSettings) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *SystemSettings) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetInterfaceRedundancyEnabled

`func (o *SystemSettings) GetInterfaceRedundancyEnabled() bool`

GetInterfaceRedundancyEnabled returns the InterfaceRedundancyEnabled field if non-nil, zero value otherwise.

### GetInterfaceRedundancyEnabledOk

`func (o *SystemSettings) GetInterfaceRedundancyEnabledOk() (*bool, bool)`

GetInterfaceRedundancyEnabledOk returns a tuple with the InterfaceRedundancyEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceRedundancyEnabled

`func (o *SystemSettings) SetInterfaceRedundancyEnabled(v bool)`

SetInterfaceRedundancyEnabled sets InterfaceRedundancyEnabled field to given value.

### HasInterfaceRedundancyEnabled

`func (o *SystemSettings) HasInterfaceRedundancyEnabled() bool`

HasInterfaceRedundancyEnabled returns a boolean if a field has been set.

### GetActiveSessions

`func (o *SystemSettings) GetActiveSessions() int32`

GetActiveSessions returns the ActiveSessions field if non-nil, zero value otherwise.

### GetActiveSessionsOk

`func (o *SystemSettings) GetActiveSessionsOk() (*int32, bool)`

GetActiveSessionsOk returns a tuple with the ActiveSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveSessions

`func (o *SystemSettings) SetActiveSessions(v int32)`

SetActiveSessions sets ActiveSessions field to given value.

### HasActiveSessions

`func (o *SystemSettings) HasActiveSessions() bool`

HasActiveSessions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


