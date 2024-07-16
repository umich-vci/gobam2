# TacacsUserBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the TACACS+ user. | [optional] 
**MemberOf** | Pointer to **string** | The name of the TACACS+ group that the user is a member of. | [optional] 
**Executables** | Pointer to **[]string** |  | [optional] 

## Methods

### NewTacacsUserBean

`func NewTacacsUserBean() *TacacsUserBean`

NewTacacsUserBean instantiates a new TacacsUserBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTacacsUserBeanWithDefaults

`func NewTacacsUserBeanWithDefaults() *TacacsUserBean`

NewTacacsUserBeanWithDefaults instantiates a new TacacsUserBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TacacsUserBean) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TacacsUserBean) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TacacsUserBean) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TacacsUserBean) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMemberOf

`func (o *TacacsUserBean) GetMemberOf() string`

GetMemberOf returns the MemberOf field if non-nil, zero value otherwise.

### GetMemberOfOk

`func (o *TacacsUserBean) GetMemberOfOk() (*string, bool)`

GetMemberOfOk returns a tuple with the MemberOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberOf

`func (o *TacacsUserBean) SetMemberOf(v string)`

SetMemberOf sets MemberOf field to given value.

### HasMemberOf

`func (o *TacacsUserBean) HasMemberOf() bool`

HasMemberOf returns a boolean if a field has been set.

### GetExecutables

`func (o *TacacsUserBean) GetExecutables() []string`

GetExecutables returns the Executables field if non-nil, zero value otherwise.

### GetExecutablesOk

`func (o *TacacsUserBean) GetExecutablesOk() (*[]string, bool)`

GetExecutablesOk returns a tuple with the Executables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutables

`func (o *TacacsUserBean) SetExecutables(v []string)`

SetExecutables sets Executables field to given value.

### HasExecutables

`func (o *TacacsUserBean) HasExecutables() bool`

HasExecutables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


