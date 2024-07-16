# LinkedNamingPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**InheritanceEnabled** | Pointer to **bool** |  | [optional] 
**AppliedTypes** | Pointer to **[]string** |  | [optional] 

## Methods

### NewLinkedNamingPolicy

`func NewLinkedNamingPolicy() *LinkedNamingPolicy`

NewLinkedNamingPolicy instantiates a new LinkedNamingPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkedNamingPolicyWithDefaults

`func NewLinkedNamingPolicyWithDefaults() *LinkedNamingPolicy`

NewLinkedNamingPolicyWithDefaults instantiates a new LinkedNamingPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LinkedNamingPolicy) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LinkedNamingPolicy) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LinkedNamingPolicy) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *LinkedNamingPolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *LinkedNamingPolicy) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LinkedNamingPolicy) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LinkedNamingPolicy) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *LinkedNamingPolicy) HasType() bool`

HasType returns a boolean if a field has been set.

### GetInheritanceEnabled

`func (o *LinkedNamingPolicy) GetInheritanceEnabled() bool`

GetInheritanceEnabled returns the InheritanceEnabled field if non-nil, zero value otherwise.

### GetInheritanceEnabledOk

`func (o *LinkedNamingPolicy) GetInheritanceEnabledOk() (*bool, bool)`

GetInheritanceEnabledOk returns a tuple with the InheritanceEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritanceEnabled

`func (o *LinkedNamingPolicy) SetInheritanceEnabled(v bool)`

SetInheritanceEnabled sets InheritanceEnabled field to given value.

### HasInheritanceEnabled

`func (o *LinkedNamingPolicy) HasInheritanceEnabled() bool`

HasInheritanceEnabled returns a boolean if a field has been set.

### GetAppliedTypes

`func (o *LinkedNamingPolicy) GetAppliedTypes() []string`

GetAppliedTypes returns the AppliedTypes field if non-nil, zero value otherwise.

### GetAppliedTypesOk

`func (o *LinkedNamingPolicy) GetAppliedTypesOk() (*[]string, bool)`

GetAppliedTypesOk returns a tuple with the AppliedTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedTypes

`func (o *LinkedNamingPolicy) SetAppliedTypes(v []string)`

SetAppliedTypes sets AppliedTypes field to given value.

### HasAppliedTypes

`func (o *LinkedNamingPolicy) HasAppliedTypes() bool`

HasAppliedTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


