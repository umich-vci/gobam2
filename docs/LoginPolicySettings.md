# LoginPolicySettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The resource type. | [optional] 
**Enabled** | Pointer to **bool** | Indicates whether a custom login policy is enabled. | [optional] 
**FailureLimit** | Pointer to **int32** | The limit for consecutive failed login attempts that results in a login delay. | [optional] 
**FailureLimitPeriod** | Pointer to **string** | The timespan in which consecutive failed login attempts count towards the failure limit. | [optional] 
**DelayDuration** | Pointer to **string** | The amount of time that a user must wait before they can attempt to log in after exceeding the login policy failure conditions. | [optional] 

## Methods

### NewLoginPolicySettings

`func NewLoginPolicySettings() *LoginPolicySettings`

NewLoginPolicySettings instantiates a new LoginPolicySettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoginPolicySettingsWithDefaults

`func NewLoginPolicySettingsWithDefaults() *LoginPolicySettings`

NewLoginPolicySettingsWithDefaults instantiates a new LoginPolicySettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *LoginPolicySettings) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LoginPolicySettings) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LoginPolicySettings) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *LoginPolicySettings) HasType() bool`

HasType returns a boolean if a field has been set.

### GetEnabled

`func (o *LoginPolicySettings) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *LoginPolicySettings) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *LoginPolicySettings) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *LoginPolicySettings) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFailureLimit

`func (o *LoginPolicySettings) GetFailureLimit() int32`

GetFailureLimit returns the FailureLimit field if non-nil, zero value otherwise.

### GetFailureLimitOk

`func (o *LoginPolicySettings) GetFailureLimitOk() (*int32, bool)`

GetFailureLimitOk returns a tuple with the FailureLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureLimit

`func (o *LoginPolicySettings) SetFailureLimit(v int32)`

SetFailureLimit sets FailureLimit field to given value.

### HasFailureLimit

`func (o *LoginPolicySettings) HasFailureLimit() bool`

HasFailureLimit returns a boolean if a field has been set.

### GetFailureLimitPeriod

`func (o *LoginPolicySettings) GetFailureLimitPeriod() string`

GetFailureLimitPeriod returns the FailureLimitPeriod field if non-nil, zero value otherwise.

### GetFailureLimitPeriodOk

`func (o *LoginPolicySettings) GetFailureLimitPeriodOk() (*string, bool)`

GetFailureLimitPeriodOk returns a tuple with the FailureLimitPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureLimitPeriod

`func (o *LoginPolicySettings) SetFailureLimitPeriod(v string)`

SetFailureLimitPeriod sets FailureLimitPeriod field to given value.

### HasFailureLimitPeriod

`func (o *LoginPolicySettings) HasFailureLimitPeriod() bool`

HasFailureLimitPeriod returns a boolean if a field has been set.

### GetDelayDuration

`func (o *LoginPolicySettings) GetDelayDuration() string`

GetDelayDuration returns the DelayDuration field if non-nil, zero value otherwise.

### GetDelayDurationOk

`func (o *LoginPolicySettings) GetDelayDurationOk() (*string, bool)`

GetDelayDurationOk returns a tuple with the DelayDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayDuration

`func (o *LoginPolicySettings) SetDelayDuration(v string)`

SetDelayDuration sets DelayDuration field to given value.

### HasDelayDuration

`func (o *LoginPolicySettings) HasDelayDuration() bool`

HasDelayDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


