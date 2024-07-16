# LinkedUserGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 

## Methods

### NewLinkedUserGroup

`func NewLinkedUserGroup() *LinkedUserGroup`

NewLinkedUserGroup instantiates a new LinkedUserGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkedUserGroupWithDefaults

`func NewLinkedUserGroupWithDefaults() *LinkedUserGroup`

NewLinkedUserGroupWithDefaults instantiates a new LinkedUserGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LinkedUserGroup) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LinkedUserGroup) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LinkedUserGroup) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *LinkedUserGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *LinkedUserGroup) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LinkedUserGroup) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LinkedUserGroup) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *LinkedUserGroup) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


