# GetSplits200Response1DataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Network** | Pointer to [**InlinedIPv4Network**](InlinedIPv4Network.md) |  | [optional] 
**NetworkCount** | Pointer to **int32** | The number of resultant sub-networks after the split. | [optional] 
**GatewayPreserved** | Pointer to **bool** | Indicates whether the gateway address of the original network is kept. | [optional] 
**DefaultGatewayAssigned** | Pointer to **bool** | Indicates whether the first IP address after the network ID is assigned as the default gateway address. | [optional] 
**ConflictResolutionStrategy** | Pointer to **string** | The method used to resolve conflicts between template items and the resource the template is applied to. | [optional] 
**Template** | Pointer to [**InlinedIPv4Template**](InlinedIPv4Template.md) |  | [optional] 

## Methods

### NewGetSplits200Response1DataInner

`func NewGetSplits200Response1DataInner() *GetSplits200Response1DataInner`

NewGetSplits200Response1DataInner instantiates a new GetSplits200Response1DataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSplits200Response1DataInnerWithDefaults

`func NewGetSplits200Response1DataInnerWithDefaults() *GetSplits200Response1DataInner`

NewGetSplits200Response1DataInnerWithDefaults instantiates a new GetSplits200Response1DataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetSplits200Response1DataInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetSplits200Response1DataInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetSplits200Response1DataInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetSplits200Response1DataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GetSplits200Response1DataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetSplits200Response1DataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetSplits200Response1DataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetSplits200Response1DataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetNetwork

`func (o *GetSplits200Response1DataInner) GetNetwork() InlinedIPv4Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *GetSplits200Response1DataInner) GetNetworkOk() (*InlinedIPv4Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *GetSplits200Response1DataInner) SetNetwork(v InlinedIPv4Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *GetSplits200Response1DataInner) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetNetworkCount

`func (o *GetSplits200Response1DataInner) GetNetworkCount() int32`

GetNetworkCount returns the NetworkCount field if non-nil, zero value otherwise.

### GetNetworkCountOk

`func (o *GetSplits200Response1DataInner) GetNetworkCountOk() (*int32, bool)`

GetNetworkCountOk returns a tuple with the NetworkCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkCount

`func (o *GetSplits200Response1DataInner) SetNetworkCount(v int32)`

SetNetworkCount sets NetworkCount field to given value.

### HasNetworkCount

`func (o *GetSplits200Response1DataInner) HasNetworkCount() bool`

HasNetworkCount returns a boolean if a field has been set.

### GetGatewayPreserved

`func (o *GetSplits200Response1DataInner) GetGatewayPreserved() bool`

GetGatewayPreserved returns the GatewayPreserved field if non-nil, zero value otherwise.

### GetGatewayPreservedOk

`func (o *GetSplits200Response1DataInner) GetGatewayPreservedOk() (*bool, bool)`

GetGatewayPreservedOk returns a tuple with the GatewayPreserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayPreserved

`func (o *GetSplits200Response1DataInner) SetGatewayPreserved(v bool)`

SetGatewayPreserved sets GatewayPreserved field to given value.

### HasGatewayPreserved

`func (o *GetSplits200Response1DataInner) HasGatewayPreserved() bool`

HasGatewayPreserved returns a boolean if a field has been set.

### GetDefaultGatewayAssigned

`func (o *GetSplits200Response1DataInner) GetDefaultGatewayAssigned() bool`

GetDefaultGatewayAssigned returns the DefaultGatewayAssigned field if non-nil, zero value otherwise.

### GetDefaultGatewayAssignedOk

`func (o *GetSplits200Response1DataInner) GetDefaultGatewayAssignedOk() (*bool, bool)`

GetDefaultGatewayAssignedOk returns a tuple with the DefaultGatewayAssigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultGatewayAssigned

`func (o *GetSplits200Response1DataInner) SetDefaultGatewayAssigned(v bool)`

SetDefaultGatewayAssigned sets DefaultGatewayAssigned field to given value.

### HasDefaultGatewayAssigned

`func (o *GetSplits200Response1DataInner) HasDefaultGatewayAssigned() bool`

HasDefaultGatewayAssigned returns a boolean if a field has been set.

### GetConflictResolutionStrategy

`func (o *GetSplits200Response1DataInner) GetConflictResolutionStrategy() string`

GetConflictResolutionStrategy returns the ConflictResolutionStrategy field if non-nil, zero value otherwise.

### GetConflictResolutionStrategyOk

`func (o *GetSplits200Response1DataInner) GetConflictResolutionStrategyOk() (*string, bool)`

GetConflictResolutionStrategyOk returns a tuple with the ConflictResolutionStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflictResolutionStrategy

`func (o *GetSplits200Response1DataInner) SetConflictResolutionStrategy(v string)`

SetConflictResolutionStrategy sets ConflictResolutionStrategy field to given value.

### HasConflictResolutionStrategy

`func (o *GetSplits200Response1DataInner) HasConflictResolutionStrategy() bool`

HasConflictResolutionStrategy returns a boolean if a field has been set.

### GetTemplate

`func (o *GetSplits200Response1DataInner) GetTemplate() InlinedIPv4Template`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *GetSplits200Response1DataInner) GetTemplateOk() (*InlinedIPv4Template, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *GetSplits200Response1DataInner) SetTemplate(v InlinedIPv4Template)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *GetSplits200Response1DataInner) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


