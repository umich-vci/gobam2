# GlobalSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The resource type. | [optional] 
**MandatoryCommentsEnabled** | Pointer to **bool** | Indicates whether users must provide a change control comment when creating, editing, or deleting resources. | [optional] 
**RememberLastAddressEnabled** | Pointer to **bool** | Indicates whether the Address Manager UI Quick Actions widget will automatically provide the subsequent IPv4 address when adding host records. | [optional] 
**SessionTimeout** | Pointer to **string** | The maximum time period of user inactivity before a session is terminated. | [optional] 
**DisclaimerEnabled** | Pointer to **bool** | Indicates whether a disclaimer will display on the Address Manager UI login page with the contents of disclaimerText. | [optional] 
**DisclaimerText** | Pointer to **string** | The disclaimer text or HTML for display on the Address Manager UI login page. If adding or editing disclaimer HTML through JSON, ensure that reserved characters are escaped. | [optional] 
**CustomReverseZoneFormatsAllowed** | Pointer to **bool** | Indicates whether users can set a custom reverse zone name format in DNS deployment options. | [optional] 

## Methods

### NewGlobalSettings

`func NewGlobalSettings() *GlobalSettings`

NewGlobalSettings instantiates a new GlobalSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalSettingsWithDefaults

`func NewGlobalSettingsWithDefaults() *GlobalSettings`

NewGlobalSettingsWithDefaults instantiates a new GlobalSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GlobalSettings) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GlobalSettings) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GlobalSettings) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GlobalSettings) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMandatoryCommentsEnabled

`func (o *GlobalSettings) GetMandatoryCommentsEnabled() bool`

GetMandatoryCommentsEnabled returns the MandatoryCommentsEnabled field if non-nil, zero value otherwise.

### GetMandatoryCommentsEnabledOk

`func (o *GlobalSettings) GetMandatoryCommentsEnabledOk() (*bool, bool)`

GetMandatoryCommentsEnabledOk returns a tuple with the MandatoryCommentsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandatoryCommentsEnabled

`func (o *GlobalSettings) SetMandatoryCommentsEnabled(v bool)`

SetMandatoryCommentsEnabled sets MandatoryCommentsEnabled field to given value.

### HasMandatoryCommentsEnabled

`func (o *GlobalSettings) HasMandatoryCommentsEnabled() bool`

HasMandatoryCommentsEnabled returns a boolean if a field has been set.

### GetRememberLastAddressEnabled

`func (o *GlobalSettings) GetRememberLastAddressEnabled() bool`

GetRememberLastAddressEnabled returns the RememberLastAddressEnabled field if non-nil, zero value otherwise.

### GetRememberLastAddressEnabledOk

`func (o *GlobalSettings) GetRememberLastAddressEnabledOk() (*bool, bool)`

GetRememberLastAddressEnabledOk returns a tuple with the RememberLastAddressEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRememberLastAddressEnabled

`func (o *GlobalSettings) SetRememberLastAddressEnabled(v bool)`

SetRememberLastAddressEnabled sets RememberLastAddressEnabled field to given value.

### HasRememberLastAddressEnabled

`func (o *GlobalSettings) HasRememberLastAddressEnabled() bool`

HasRememberLastAddressEnabled returns a boolean if a field has been set.

### GetSessionTimeout

`func (o *GlobalSettings) GetSessionTimeout() string`

GetSessionTimeout returns the SessionTimeout field if non-nil, zero value otherwise.

### GetSessionTimeoutOk

`func (o *GlobalSettings) GetSessionTimeoutOk() (*string, bool)`

GetSessionTimeoutOk returns a tuple with the SessionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionTimeout

`func (o *GlobalSettings) SetSessionTimeout(v string)`

SetSessionTimeout sets SessionTimeout field to given value.

### HasSessionTimeout

`func (o *GlobalSettings) HasSessionTimeout() bool`

HasSessionTimeout returns a boolean if a field has been set.

### GetDisclaimerEnabled

`func (o *GlobalSettings) GetDisclaimerEnabled() bool`

GetDisclaimerEnabled returns the DisclaimerEnabled field if non-nil, zero value otherwise.

### GetDisclaimerEnabledOk

`func (o *GlobalSettings) GetDisclaimerEnabledOk() (*bool, bool)`

GetDisclaimerEnabledOk returns a tuple with the DisclaimerEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisclaimerEnabled

`func (o *GlobalSettings) SetDisclaimerEnabled(v bool)`

SetDisclaimerEnabled sets DisclaimerEnabled field to given value.

### HasDisclaimerEnabled

`func (o *GlobalSettings) HasDisclaimerEnabled() bool`

HasDisclaimerEnabled returns a boolean if a field has been set.

### GetDisclaimerText

`func (o *GlobalSettings) GetDisclaimerText() string`

GetDisclaimerText returns the DisclaimerText field if non-nil, zero value otherwise.

### GetDisclaimerTextOk

`func (o *GlobalSettings) GetDisclaimerTextOk() (*string, bool)`

GetDisclaimerTextOk returns a tuple with the DisclaimerText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisclaimerText

`func (o *GlobalSettings) SetDisclaimerText(v string)`

SetDisclaimerText sets DisclaimerText field to given value.

### HasDisclaimerText

`func (o *GlobalSettings) HasDisclaimerText() bool`

HasDisclaimerText returns a boolean if a field has been set.

### GetCustomReverseZoneFormatsAllowed

`func (o *GlobalSettings) GetCustomReverseZoneFormatsAllowed() bool`

GetCustomReverseZoneFormatsAllowed returns the CustomReverseZoneFormatsAllowed field if non-nil, zero value otherwise.

### GetCustomReverseZoneFormatsAllowedOk

`func (o *GlobalSettings) GetCustomReverseZoneFormatsAllowedOk() (*bool, bool)`

GetCustomReverseZoneFormatsAllowedOk returns a tuple with the CustomReverseZoneFormatsAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomReverseZoneFormatsAllowed

`func (o *GlobalSettings) SetCustomReverseZoneFormatsAllowed(v bool)`

SetCustomReverseZoneFormatsAllowed sets CustomReverseZoneFormatsAllowed field to given value.

### HasCustomReverseZoneFormatsAllowed

`func (o *GlobalSettings) HasCustomReverseZoneFormatsAllowed() bool`

HasCustomReverseZoneFormatsAllowed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


