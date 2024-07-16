# TacacsPlusAuthenticationBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Indicates whether TACACS+ authentication is enabled. | [optional] 
**Server** | Pointer to **string** | The hostname or IP address of the TACACS+ server used for authentication. | [optional] 
**SharedSecret** | Pointer to **string** | The shared secret used to encrypt and decrypt the patckets between the client and the server. | [optional] 
**Users** | Pointer to [**[]TacacsUserBean**](TacacsUserBean.md) |  | [optional] 
**Groups** | Pointer to [**[]TacacsGroupBean**](TacacsGroupBean.md) |  | [optional] 

## Methods

### NewTacacsPlusAuthenticationBean

`func NewTacacsPlusAuthenticationBean() *TacacsPlusAuthenticationBean`

NewTacacsPlusAuthenticationBean instantiates a new TacacsPlusAuthenticationBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTacacsPlusAuthenticationBeanWithDefaults

`func NewTacacsPlusAuthenticationBeanWithDefaults() *TacacsPlusAuthenticationBean`

NewTacacsPlusAuthenticationBeanWithDefaults instantiates a new TacacsPlusAuthenticationBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *TacacsPlusAuthenticationBean) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *TacacsPlusAuthenticationBean) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *TacacsPlusAuthenticationBean) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *TacacsPlusAuthenticationBean) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetServer

`func (o *TacacsPlusAuthenticationBean) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *TacacsPlusAuthenticationBean) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *TacacsPlusAuthenticationBean) SetServer(v string)`

SetServer sets Server field to given value.

### HasServer

`func (o *TacacsPlusAuthenticationBean) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetSharedSecret

`func (o *TacacsPlusAuthenticationBean) GetSharedSecret() string`

GetSharedSecret returns the SharedSecret field if non-nil, zero value otherwise.

### GetSharedSecretOk

`func (o *TacacsPlusAuthenticationBean) GetSharedSecretOk() (*string, bool)`

GetSharedSecretOk returns a tuple with the SharedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSecret

`func (o *TacacsPlusAuthenticationBean) SetSharedSecret(v string)`

SetSharedSecret sets SharedSecret field to given value.

### HasSharedSecret

`func (o *TacacsPlusAuthenticationBean) HasSharedSecret() bool`

HasSharedSecret returns a boolean if a field has been set.

### GetUsers

`func (o *TacacsPlusAuthenticationBean) GetUsers() []TacacsUserBean`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *TacacsPlusAuthenticationBean) GetUsersOk() (*[]TacacsUserBean, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *TacacsPlusAuthenticationBean) SetUsers(v []TacacsUserBean)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *TacacsPlusAuthenticationBean) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetGroups

`func (o *TacacsPlusAuthenticationBean) GetGroups() []TacacsGroupBean`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *TacacsPlusAuthenticationBean) GetGroupsOk() (*[]TacacsGroupBean, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *TacacsPlusAuthenticationBean) SetGroups(v []TacacsGroupBean)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *TacacsPlusAuthenticationBean) HasGroups() bool`

HasGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


