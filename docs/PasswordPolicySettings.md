# PasswordPolicySettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The resource type. | [optional] 
**Enabled** | Pointer to **bool** | Indicates whether a custom password policy is enabled. | [optional] 
**MinLength** | Pointer to **int32** | The minimum length that the password must be. | [optional] 
**MaxLength** | Pointer to **int32** | The maximum length that the password can be. | [optional] 
**DigitRequired** | Pointer to **bool** | Indicates whether the password must contain at least one digit. | [optional] 
**MixedCaseRequired** | Pointer to **bool** | Indicates whether the password must contain mixed case letters. | [optional] 
**SpecialCharacterRequired** | Pointer to **bool** | Indicates whether the password must contain at least one special character. | [optional] 

## Methods

### NewPasswordPolicySettings

`func NewPasswordPolicySettings() *PasswordPolicySettings`

NewPasswordPolicySettings instantiates a new PasswordPolicySettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordPolicySettingsWithDefaults

`func NewPasswordPolicySettingsWithDefaults() *PasswordPolicySettings`

NewPasswordPolicySettingsWithDefaults instantiates a new PasswordPolicySettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PasswordPolicySettings) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PasswordPolicySettings) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PasswordPolicySettings) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PasswordPolicySettings) HasType() bool`

HasType returns a boolean if a field has been set.

### GetEnabled

`func (o *PasswordPolicySettings) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PasswordPolicySettings) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PasswordPolicySettings) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PasswordPolicySettings) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMinLength

`func (o *PasswordPolicySettings) GetMinLength() int32`

GetMinLength returns the MinLength field if non-nil, zero value otherwise.

### GetMinLengthOk

`func (o *PasswordPolicySettings) GetMinLengthOk() (*int32, bool)`

GetMinLengthOk returns a tuple with the MinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLength

`func (o *PasswordPolicySettings) SetMinLength(v int32)`

SetMinLength sets MinLength field to given value.

### HasMinLength

`func (o *PasswordPolicySettings) HasMinLength() bool`

HasMinLength returns a boolean if a field has been set.

### GetMaxLength

`func (o *PasswordPolicySettings) GetMaxLength() int32`

GetMaxLength returns the MaxLength field if non-nil, zero value otherwise.

### GetMaxLengthOk

`func (o *PasswordPolicySettings) GetMaxLengthOk() (*int32, bool)`

GetMaxLengthOk returns a tuple with the MaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLength

`func (o *PasswordPolicySettings) SetMaxLength(v int32)`

SetMaxLength sets MaxLength field to given value.

### HasMaxLength

`func (o *PasswordPolicySettings) HasMaxLength() bool`

HasMaxLength returns a boolean if a field has been set.

### GetDigitRequired

`func (o *PasswordPolicySettings) GetDigitRequired() bool`

GetDigitRequired returns the DigitRequired field if non-nil, zero value otherwise.

### GetDigitRequiredOk

`func (o *PasswordPolicySettings) GetDigitRequiredOk() (*bool, bool)`

GetDigitRequiredOk returns a tuple with the DigitRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitRequired

`func (o *PasswordPolicySettings) SetDigitRequired(v bool)`

SetDigitRequired sets DigitRequired field to given value.

### HasDigitRequired

`func (o *PasswordPolicySettings) HasDigitRequired() bool`

HasDigitRequired returns a boolean if a field has been set.

### GetMixedCaseRequired

`func (o *PasswordPolicySettings) GetMixedCaseRequired() bool`

GetMixedCaseRequired returns the MixedCaseRequired field if non-nil, zero value otherwise.

### GetMixedCaseRequiredOk

`func (o *PasswordPolicySettings) GetMixedCaseRequiredOk() (*bool, bool)`

GetMixedCaseRequiredOk returns a tuple with the MixedCaseRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMixedCaseRequired

`func (o *PasswordPolicySettings) SetMixedCaseRequired(v bool)`

SetMixedCaseRequired sets MixedCaseRequired field to given value.

### HasMixedCaseRequired

`func (o *PasswordPolicySettings) HasMixedCaseRequired() bool`

HasMixedCaseRequired returns a boolean if a field has been set.

### GetSpecialCharacterRequired

`func (o *PasswordPolicySettings) GetSpecialCharacterRequired() bool`

GetSpecialCharacterRequired returns the SpecialCharacterRequired field if non-nil, zero value otherwise.

### GetSpecialCharacterRequiredOk

`func (o *PasswordPolicySettings) GetSpecialCharacterRequiredOk() (*bool, bool)`

GetSpecialCharacterRequiredOk returns a tuple with the SpecialCharacterRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialCharacterRequired

`func (o *PasswordPolicySettings) SetSpecialCharacterRequired(v bool)`

SetSpecialCharacterRequired sets SpecialCharacterRequired field to given value.

### HasSpecialCharacterRequired

`func (o *PasswordPolicySettings) HasSpecialCharacterRequired() bool`

HasSpecialCharacterRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


