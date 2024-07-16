# TacacsGroupBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the TACACS+ group. | [optional] 
**Executables** | Pointer to **[]string** |  | [optional] 

## Methods

### NewTacacsGroupBean

`func NewTacacsGroupBean() *TacacsGroupBean`

NewTacacsGroupBean instantiates a new TacacsGroupBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTacacsGroupBeanWithDefaults

`func NewTacacsGroupBeanWithDefaults() *TacacsGroupBean`

NewTacacsGroupBeanWithDefaults instantiates a new TacacsGroupBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TacacsGroupBean) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TacacsGroupBean) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TacacsGroupBean) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TacacsGroupBean) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExecutables

`func (o *TacacsGroupBean) GetExecutables() []string`

GetExecutables returns the Executables field if non-nil, zero value otherwise.

### GetExecutablesOk

`func (o *TacacsGroupBean) GetExecutablesOk() (*[]string, bool)`

GetExecutablesOk returns a tuple with the Executables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutables

`func (o *TacacsGroupBean) SetExecutables(v []string)`

SetExecutables sets Executables field to given value.

### HasExecutables

`func (o *TacacsGroupBean) HasExecutables() bool`

HasExecutables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


