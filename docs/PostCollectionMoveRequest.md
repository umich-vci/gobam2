# PostCollectionMoveRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | **string** | The resource type. | 
**AbsoluteName** | **string** | The absolute name of the destination zone resource where the DNS resource is to be moved. | 
**RoleTypes** | **[]string** |  | 
**Destination** | [**InlinedServerInterface**](InlinedServerInterface.md) |  | 

## Methods

### NewPostCollectionMoveRequest

`func NewPostCollectionMoveRequest(type_ string, absoluteName string, roleTypes []string, destination InlinedServerInterface, ) *PostCollectionMoveRequest`

NewPostCollectionMoveRequest instantiates a new PostCollectionMoveRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostCollectionMoveRequestWithDefaults

`func NewPostCollectionMoveRequestWithDefaults() *PostCollectionMoveRequest`

NewPostCollectionMoveRequestWithDefaults instantiates a new PostCollectionMoveRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PostCollectionMoveRequest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostCollectionMoveRequest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostCollectionMoveRequest) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *PostCollectionMoveRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *PostCollectionMoveRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PostCollectionMoveRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PostCollectionMoveRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAbsoluteName

`func (o *PostCollectionMoveRequest) GetAbsoluteName() string`

GetAbsoluteName returns the AbsoluteName field if non-nil, zero value otherwise.

### GetAbsoluteNameOk

`func (o *PostCollectionMoveRequest) GetAbsoluteNameOk() (*string, bool)`

GetAbsoluteNameOk returns a tuple with the AbsoluteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsoluteName

`func (o *PostCollectionMoveRequest) SetAbsoluteName(v string)`

SetAbsoluteName sets AbsoluteName field to given value.


### GetRoleTypes

`func (o *PostCollectionMoveRequest) GetRoleTypes() []string`

GetRoleTypes returns the RoleTypes field if non-nil, zero value otherwise.

### GetRoleTypesOk

`func (o *PostCollectionMoveRequest) GetRoleTypesOk() (*[]string, bool)`

GetRoleTypesOk returns a tuple with the RoleTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleTypes

`func (o *PostCollectionMoveRequest) SetRoleTypes(v []string)`

SetRoleTypes sets RoleTypes field to given value.


### GetDestination

`func (o *PostCollectionMoveRequest) GetDestination() InlinedServerInterface`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *PostCollectionMoveRequest) GetDestinationOk() (*InlinedServerInterface, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *PostCollectionMoveRequest) SetDestination(v InlinedServerInterface)`

SetDestination sets Destination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


