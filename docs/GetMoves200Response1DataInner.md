# GetMoves200Response1DataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**AbsoluteName** | Pointer to **string** | The absolute name of the destination zone resource where the DNS resource is to be moved. | [optional] 
**RoleTypes** | Pointer to **[]string** |  | [optional] 
**Destination** | Pointer to [**InlinedServerInterface**](InlinedServerInterface.md) |  | [optional] 

## Methods

### NewGetMoves200Response1DataInner

`func NewGetMoves200Response1DataInner() *GetMoves200Response1DataInner`

NewGetMoves200Response1DataInner instantiates a new GetMoves200Response1DataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMoves200Response1DataInnerWithDefaults

`func NewGetMoves200Response1DataInnerWithDefaults() *GetMoves200Response1DataInner`

NewGetMoves200Response1DataInnerWithDefaults instantiates a new GetMoves200Response1DataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetMoves200Response1DataInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetMoves200Response1DataInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetMoves200Response1DataInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetMoves200Response1DataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GetMoves200Response1DataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetMoves200Response1DataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetMoves200Response1DataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetMoves200Response1DataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAbsoluteName

`func (o *GetMoves200Response1DataInner) GetAbsoluteName() string`

GetAbsoluteName returns the AbsoluteName field if non-nil, zero value otherwise.

### GetAbsoluteNameOk

`func (o *GetMoves200Response1DataInner) GetAbsoluteNameOk() (*string, bool)`

GetAbsoluteNameOk returns a tuple with the AbsoluteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsoluteName

`func (o *GetMoves200Response1DataInner) SetAbsoluteName(v string)`

SetAbsoluteName sets AbsoluteName field to given value.

### HasAbsoluteName

`func (o *GetMoves200Response1DataInner) HasAbsoluteName() bool`

HasAbsoluteName returns a boolean if a field has been set.

### GetRoleTypes

`func (o *GetMoves200Response1DataInner) GetRoleTypes() []string`

GetRoleTypes returns the RoleTypes field if non-nil, zero value otherwise.

### GetRoleTypesOk

`func (o *GetMoves200Response1DataInner) GetRoleTypesOk() (*[]string, bool)`

GetRoleTypesOk returns a tuple with the RoleTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleTypes

`func (o *GetMoves200Response1DataInner) SetRoleTypes(v []string)`

SetRoleTypes sets RoleTypes field to given value.

### HasRoleTypes

`func (o *GetMoves200Response1DataInner) HasRoleTypes() bool`

HasRoleTypes returns a boolean if a field has been set.

### GetDestination

`func (o *GetMoves200Response1DataInner) GetDestination() InlinedServerInterface`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *GetMoves200Response1DataInner) GetDestinationOk() (*InlinedServerInterface, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *GetMoves200Response1DataInner) SetDestination(v InlinedServerInterface)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *GetMoves200Response1DataInner) HasDestination() bool`

HasDestination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


