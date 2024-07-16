# SSOEnforcementSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The resource type. | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**SamlIdentityProviderEnabled** | Pointer to **bool** |  | [optional] [readonly] 
**NonSsoAuthenticatorCount** | Pointer to **int32** |  | [optional] [readonly] 
**NonSsoGroupCount** | Pointer to **int32** |  | [optional] [readonly] 
**LocalAdminUserCount** | Pointer to **int32** |  | [optional] [readonly] 
**LocalNonAdminUserCount** | Pointer to **int32** |  | [optional] [readonly] 

## Methods

### NewSSOEnforcementSettings

`func NewSSOEnforcementSettings() *SSOEnforcementSettings`

NewSSOEnforcementSettings instantiates a new SSOEnforcementSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSSOEnforcementSettingsWithDefaults

`func NewSSOEnforcementSettingsWithDefaults() *SSOEnforcementSettings`

NewSSOEnforcementSettingsWithDefaults instantiates a new SSOEnforcementSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SSOEnforcementSettings) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SSOEnforcementSettings) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SSOEnforcementSettings) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SSOEnforcementSettings) HasType() bool`

HasType returns a boolean if a field has been set.

### GetEnabled

`func (o *SSOEnforcementSettings) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SSOEnforcementSettings) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SSOEnforcementSettings) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SSOEnforcementSettings) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetSamlIdentityProviderEnabled

`func (o *SSOEnforcementSettings) GetSamlIdentityProviderEnabled() bool`

GetSamlIdentityProviderEnabled returns the SamlIdentityProviderEnabled field if non-nil, zero value otherwise.

### GetSamlIdentityProviderEnabledOk

`func (o *SSOEnforcementSettings) GetSamlIdentityProviderEnabledOk() (*bool, bool)`

GetSamlIdentityProviderEnabledOk returns a tuple with the SamlIdentityProviderEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlIdentityProviderEnabled

`func (o *SSOEnforcementSettings) SetSamlIdentityProviderEnabled(v bool)`

SetSamlIdentityProviderEnabled sets SamlIdentityProviderEnabled field to given value.

### HasSamlIdentityProviderEnabled

`func (o *SSOEnforcementSettings) HasSamlIdentityProviderEnabled() bool`

HasSamlIdentityProviderEnabled returns a boolean if a field has been set.

### GetNonSsoAuthenticatorCount

`func (o *SSOEnforcementSettings) GetNonSsoAuthenticatorCount() int32`

GetNonSsoAuthenticatorCount returns the NonSsoAuthenticatorCount field if non-nil, zero value otherwise.

### GetNonSsoAuthenticatorCountOk

`func (o *SSOEnforcementSettings) GetNonSsoAuthenticatorCountOk() (*int32, bool)`

GetNonSsoAuthenticatorCountOk returns a tuple with the NonSsoAuthenticatorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonSsoAuthenticatorCount

`func (o *SSOEnforcementSettings) SetNonSsoAuthenticatorCount(v int32)`

SetNonSsoAuthenticatorCount sets NonSsoAuthenticatorCount field to given value.

### HasNonSsoAuthenticatorCount

`func (o *SSOEnforcementSettings) HasNonSsoAuthenticatorCount() bool`

HasNonSsoAuthenticatorCount returns a boolean if a field has been set.

### GetNonSsoGroupCount

`func (o *SSOEnforcementSettings) GetNonSsoGroupCount() int32`

GetNonSsoGroupCount returns the NonSsoGroupCount field if non-nil, zero value otherwise.

### GetNonSsoGroupCountOk

`func (o *SSOEnforcementSettings) GetNonSsoGroupCountOk() (*int32, bool)`

GetNonSsoGroupCountOk returns a tuple with the NonSsoGroupCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonSsoGroupCount

`func (o *SSOEnforcementSettings) SetNonSsoGroupCount(v int32)`

SetNonSsoGroupCount sets NonSsoGroupCount field to given value.

### HasNonSsoGroupCount

`func (o *SSOEnforcementSettings) HasNonSsoGroupCount() bool`

HasNonSsoGroupCount returns a boolean if a field has been set.

### GetLocalAdminUserCount

`func (o *SSOEnforcementSettings) GetLocalAdminUserCount() int32`

GetLocalAdminUserCount returns the LocalAdminUserCount field if non-nil, zero value otherwise.

### GetLocalAdminUserCountOk

`func (o *SSOEnforcementSettings) GetLocalAdminUserCountOk() (*int32, bool)`

GetLocalAdminUserCountOk returns a tuple with the LocalAdminUserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAdminUserCount

`func (o *SSOEnforcementSettings) SetLocalAdminUserCount(v int32)`

SetLocalAdminUserCount sets LocalAdminUserCount field to given value.

### HasLocalAdminUserCount

`func (o *SSOEnforcementSettings) HasLocalAdminUserCount() bool`

HasLocalAdminUserCount returns a boolean if a field has been set.

### GetLocalNonAdminUserCount

`func (o *SSOEnforcementSettings) GetLocalNonAdminUserCount() int32`

GetLocalNonAdminUserCount returns the LocalNonAdminUserCount field if non-nil, zero value otherwise.

### GetLocalNonAdminUserCountOk

`func (o *SSOEnforcementSettings) GetLocalNonAdminUserCountOk() (*int32, bool)`

GetLocalNonAdminUserCountOk returns a tuple with the LocalNonAdminUserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalNonAdminUserCount

`func (o *SSOEnforcementSettings) SetLocalNonAdminUserCount(v int32)`

SetLocalNonAdminUserCount sets LocalNonAdminUserCount field to given value.

### HasLocalNonAdminUserCount

`func (o *SSOEnforcementSettings) HasLocalNonAdminUserCount() bool`

HasLocalNonAdminUserCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


