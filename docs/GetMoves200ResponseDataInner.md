# GetMoves200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Address** | Pointer to **string** | The IP address of the destination resource where the IP resource is to be moved. | [optional] 
**MacAddress** | Pointer to [**IPMoveAllOfMacAddress**](IPMoveAllOfMacAddress.md) |  | [optional] 
**ServersUpdated** | Pointer to **bool** |  | [optional] 
**AbsoluteName** | Pointer to **string** | The absolute name of the destination zone resource where the DNS resource is to be moved. | [optional] 
**RoleTypes** | Pointer to **[]string** |  | [optional] 
**Destination** | Pointer to [**InlinedServerInterface**](InlinedServerInterface.md) |  | [optional] 
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 

## Methods

### NewGetMoves200ResponseDataInner

`func NewGetMoves200ResponseDataInner() *GetMoves200ResponseDataInner`

NewGetMoves200ResponseDataInner instantiates a new GetMoves200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMoves200ResponseDataInnerWithDefaults

`func NewGetMoves200ResponseDataInnerWithDefaults() *GetMoves200ResponseDataInner`

NewGetMoves200ResponseDataInnerWithDefaults instantiates a new GetMoves200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetMoves200ResponseDataInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetMoves200ResponseDataInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetMoves200ResponseDataInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetMoves200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GetMoves200ResponseDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetMoves200ResponseDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetMoves200ResponseDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetMoves200ResponseDataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAddress

`func (o *GetMoves200ResponseDataInner) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *GetMoves200ResponseDataInner) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *GetMoves200ResponseDataInner) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *GetMoves200ResponseDataInner) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetMacAddress

`func (o *GetMoves200ResponseDataInner) GetMacAddress() IPMoveAllOfMacAddress`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *GetMoves200ResponseDataInner) GetMacAddressOk() (*IPMoveAllOfMacAddress, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *GetMoves200ResponseDataInner) SetMacAddress(v IPMoveAllOfMacAddress)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *GetMoves200ResponseDataInner) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetServersUpdated

`func (o *GetMoves200ResponseDataInner) GetServersUpdated() bool`

GetServersUpdated returns the ServersUpdated field if non-nil, zero value otherwise.

### GetServersUpdatedOk

`func (o *GetMoves200ResponseDataInner) GetServersUpdatedOk() (*bool, bool)`

GetServersUpdatedOk returns a tuple with the ServersUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServersUpdated

`func (o *GetMoves200ResponseDataInner) SetServersUpdated(v bool)`

SetServersUpdated sets ServersUpdated field to given value.

### HasServersUpdated

`func (o *GetMoves200ResponseDataInner) HasServersUpdated() bool`

HasServersUpdated returns a boolean if a field has been set.

### GetAbsoluteName

`func (o *GetMoves200ResponseDataInner) GetAbsoluteName() string`

GetAbsoluteName returns the AbsoluteName field if non-nil, zero value otherwise.

### GetAbsoluteNameOk

`func (o *GetMoves200ResponseDataInner) GetAbsoluteNameOk() (*string, bool)`

GetAbsoluteNameOk returns a tuple with the AbsoluteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsoluteName

`func (o *GetMoves200ResponseDataInner) SetAbsoluteName(v string)`

SetAbsoluteName sets AbsoluteName field to given value.

### HasAbsoluteName

`func (o *GetMoves200ResponseDataInner) HasAbsoluteName() bool`

HasAbsoluteName returns a boolean if a field has been set.

### GetRoleTypes

`func (o *GetMoves200ResponseDataInner) GetRoleTypes() []string`

GetRoleTypes returns the RoleTypes field if non-nil, zero value otherwise.

### GetRoleTypesOk

`func (o *GetMoves200ResponseDataInner) GetRoleTypesOk() (*[]string, bool)`

GetRoleTypesOk returns a tuple with the RoleTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleTypes

`func (o *GetMoves200ResponseDataInner) SetRoleTypes(v []string)`

SetRoleTypes sets RoleTypes field to given value.

### HasRoleTypes

`func (o *GetMoves200ResponseDataInner) HasRoleTypes() bool`

HasRoleTypes returns a boolean if a field has been set.

### GetDestination

`func (o *GetMoves200ResponseDataInner) GetDestination() InlinedServerInterface`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *GetMoves200ResponseDataInner) GetDestinationOk() (*InlinedServerInterface, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *GetMoves200ResponseDataInner) SetDestination(v InlinedServerInterface)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *GetMoves200ResponseDataInner) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetLinks

`func (o *GetMoves200ResponseDataInner) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetMoves200ResponseDataInner) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetMoves200ResponseDataInner) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetMoves200ResponseDataInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


