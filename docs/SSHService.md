# SSHService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The resource type. | [optional] 
**Enabled** | Pointer to **bool** | Indicates whether the SSH service is enabled. | [optional] 
**TacacsPlusAuthentication** | Pointer to [**TacacsPlusAuthenticationBean**](TacacsPlusAuthenticationBean.md) |  | [optional] 

## Methods

### NewSSHService

`func NewSSHService() *SSHService`

NewSSHService instantiates a new SSHService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSSHServiceWithDefaults

`func NewSSHServiceWithDefaults() *SSHService`

NewSSHServiceWithDefaults instantiates a new SSHService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SSHService) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SSHService) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SSHService) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SSHService) HasType() bool`

HasType returns a boolean if a field has been set.

### GetEnabled

`func (o *SSHService) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SSHService) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SSHService) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SSHService) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetTacacsPlusAuthentication

`func (o *SSHService) GetTacacsPlusAuthentication() TacacsPlusAuthenticationBean`

GetTacacsPlusAuthentication returns the TacacsPlusAuthentication field if non-nil, zero value otherwise.

### GetTacacsPlusAuthenticationOk

`func (o *SSHService) GetTacacsPlusAuthenticationOk() (*TacacsPlusAuthenticationBean, bool)`

GetTacacsPlusAuthenticationOk returns a tuple with the TacacsPlusAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTacacsPlusAuthentication

`func (o *SSHService) SetTacacsPlusAuthentication(v TacacsPlusAuthenticationBean)`

SetTacacsPlusAuthentication sets TacacsPlusAuthentication field to given value.

### HasTacacsPlusAuthentication

`func (o *SSHService) HasTacacsPlusAuthentication() bool`

HasTacacsPlusAuthentication returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


