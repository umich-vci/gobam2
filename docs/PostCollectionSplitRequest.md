# PostCollectionSplitRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | **string** | The resource type. | 
**Network** | [**InlinedIPv4Network**](InlinedIPv4Network.md) |  | 
**NetworkCount** | **int32** | The number of resultant sub-networks after the split. | 
**GatewayPreserved** | Pointer to **bool** | Indicates whether the gateway address of the original network is kept. | [optional] 
**DefaultGatewayAssigned** | Pointer to **bool** | Indicates whether the first IP address after the network ID is assigned as the default gateway address. | [optional] 
**ConflictResolutionStrategy** | Pointer to **string** | The method used to resolve conflicts between template items and the resource the template is applied to. | [optional] 
**Template** | Pointer to [**InlinedIPv4Template**](InlinedIPv4Template.md) |  | [optional] 

## Methods

### NewPostCollectionSplitRequest

`func NewPostCollectionSplitRequest(type_ string, network InlinedIPv4Network, networkCount int32, ) *PostCollectionSplitRequest`

NewPostCollectionSplitRequest instantiates a new PostCollectionSplitRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostCollectionSplitRequestWithDefaults

`func NewPostCollectionSplitRequestWithDefaults() *PostCollectionSplitRequest`

NewPostCollectionSplitRequestWithDefaults instantiates a new PostCollectionSplitRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PostCollectionSplitRequest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostCollectionSplitRequest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostCollectionSplitRequest) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *PostCollectionSplitRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *PostCollectionSplitRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PostCollectionSplitRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PostCollectionSplitRequest) SetType(v string)`

SetType sets Type field to given value.


### GetNetwork

`func (o *PostCollectionSplitRequest) GetNetwork() InlinedIPv4Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *PostCollectionSplitRequest) GetNetworkOk() (*InlinedIPv4Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *PostCollectionSplitRequest) SetNetwork(v InlinedIPv4Network)`

SetNetwork sets Network field to given value.


### GetNetworkCount

`func (o *PostCollectionSplitRequest) GetNetworkCount() int32`

GetNetworkCount returns the NetworkCount field if non-nil, zero value otherwise.

### GetNetworkCountOk

`func (o *PostCollectionSplitRequest) GetNetworkCountOk() (*int32, bool)`

GetNetworkCountOk returns a tuple with the NetworkCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkCount

`func (o *PostCollectionSplitRequest) SetNetworkCount(v int32)`

SetNetworkCount sets NetworkCount field to given value.


### GetGatewayPreserved

`func (o *PostCollectionSplitRequest) GetGatewayPreserved() bool`

GetGatewayPreserved returns the GatewayPreserved field if non-nil, zero value otherwise.

### GetGatewayPreservedOk

`func (o *PostCollectionSplitRequest) GetGatewayPreservedOk() (*bool, bool)`

GetGatewayPreservedOk returns a tuple with the GatewayPreserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayPreserved

`func (o *PostCollectionSplitRequest) SetGatewayPreserved(v bool)`

SetGatewayPreserved sets GatewayPreserved field to given value.

### HasGatewayPreserved

`func (o *PostCollectionSplitRequest) HasGatewayPreserved() bool`

HasGatewayPreserved returns a boolean if a field has been set.

### GetDefaultGatewayAssigned

`func (o *PostCollectionSplitRequest) GetDefaultGatewayAssigned() bool`

GetDefaultGatewayAssigned returns the DefaultGatewayAssigned field if non-nil, zero value otherwise.

### GetDefaultGatewayAssignedOk

`func (o *PostCollectionSplitRequest) GetDefaultGatewayAssignedOk() (*bool, bool)`

GetDefaultGatewayAssignedOk returns a tuple with the DefaultGatewayAssigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultGatewayAssigned

`func (o *PostCollectionSplitRequest) SetDefaultGatewayAssigned(v bool)`

SetDefaultGatewayAssigned sets DefaultGatewayAssigned field to given value.

### HasDefaultGatewayAssigned

`func (o *PostCollectionSplitRequest) HasDefaultGatewayAssigned() bool`

HasDefaultGatewayAssigned returns a boolean if a field has been set.

### GetConflictResolutionStrategy

`func (o *PostCollectionSplitRequest) GetConflictResolutionStrategy() string`

GetConflictResolutionStrategy returns the ConflictResolutionStrategy field if non-nil, zero value otherwise.

### GetConflictResolutionStrategyOk

`func (o *PostCollectionSplitRequest) GetConflictResolutionStrategyOk() (*string, bool)`

GetConflictResolutionStrategyOk returns a tuple with the ConflictResolutionStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflictResolutionStrategy

`func (o *PostCollectionSplitRequest) SetConflictResolutionStrategy(v string)`

SetConflictResolutionStrategy sets ConflictResolutionStrategy field to given value.

### HasConflictResolutionStrategy

`func (o *PostCollectionSplitRequest) HasConflictResolutionStrategy() bool`

HasConflictResolutionStrategy returns a boolean if a field has been set.

### GetTemplate

`func (o *PostCollectionSplitRequest) GetTemplate() InlinedIPv4Template`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *PostCollectionSplitRequest) GetTemplateOk() (*InlinedIPv4Template, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *PostCollectionSplitRequest) SetTemplate(v InlinedIPv4Template)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *PostCollectionSplitRequest) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


