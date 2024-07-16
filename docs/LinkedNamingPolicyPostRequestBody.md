# LinkedNamingPolicyPostRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The resource identifier. | 
**Type** | Pointer to **string** | The resource type. | [optional] 
**InheritanceEnabled** | Pointer to **bool** |  | [optional] 
**AppliedTypes** | Pointer to **[]string** |  | [optional] 

## Methods

### NewLinkedNamingPolicyPostRequestBody

`func NewLinkedNamingPolicyPostRequestBody(id int64, ) *LinkedNamingPolicyPostRequestBody`

NewLinkedNamingPolicyPostRequestBody instantiates a new LinkedNamingPolicyPostRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkedNamingPolicyPostRequestBodyWithDefaults

`func NewLinkedNamingPolicyPostRequestBodyWithDefaults() *LinkedNamingPolicyPostRequestBody`

NewLinkedNamingPolicyPostRequestBodyWithDefaults instantiates a new LinkedNamingPolicyPostRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LinkedNamingPolicyPostRequestBody) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LinkedNamingPolicyPostRequestBody) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LinkedNamingPolicyPostRequestBody) SetId(v int64)`

SetId sets Id field to given value.


### GetType

`func (o *LinkedNamingPolicyPostRequestBody) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LinkedNamingPolicyPostRequestBody) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LinkedNamingPolicyPostRequestBody) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *LinkedNamingPolicyPostRequestBody) HasType() bool`

HasType returns a boolean if a field has been set.

### GetInheritanceEnabled

`func (o *LinkedNamingPolicyPostRequestBody) GetInheritanceEnabled() bool`

GetInheritanceEnabled returns the InheritanceEnabled field if non-nil, zero value otherwise.

### GetInheritanceEnabledOk

`func (o *LinkedNamingPolicyPostRequestBody) GetInheritanceEnabledOk() (*bool, bool)`

GetInheritanceEnabledOk returns a tuple with the InheritanceEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritanceEnabled

`func (o *LinkedNamingPolicyPostRequestBody) SetInheritanceEnabled(v bool)`

SetInheritanceEnabled sets InheritanceEnabled field to given value.

### HasInheritanceEnabled

`func (o *LinkedNamingPolicyPostRequestBody) HasInheritanceEnabled() bool`

HasInheritanceEnabled returns a boolean if a field has been set.

### GetAppliedTypes

`func (o *LinkedNamingPolicyPostRequestBody) GetAppliedTypes() []string`

GetAppliedTypes returns the AppliedTypes field if non-nil, zero value otherwise.

### GetAppliedTypesOk

`func (o *LinkedNamingPolicyPostRequestBody) GetAppliedTypesOk() (*[]string, bool)`

GetAppliedTypesOk returns a tuple with the AppliedTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedTypes

`func (o *LinkedNamingPolicyPostRequestBody) SetAppliedTypes(v []string)`

SetAppliedTypes sets AppliedTypes field to given value.

### HasAppliedTypes

`func (o *LinkedNamingPolicyPostRequestBody) HasAppliedTypes() bool`

HasAppliedTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


