# GetSplits200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Block** | Pointer to [**InlinedIPv4Block**](InlinedIPv4Block.md) |  | [optional] 
**SplitPointAddress** | Pointer to **string** | The IP address of the point at which you would like to split the block. | [optional] 
**Network** | Pointer to [**InlinedIPv4Network**](InlinedIPv4Network.md) |  | [optional] 
**NetworkCount** | Pointer to **int32** | The number of resultant sub-networks after the split. | [optional] 
**GatewayPreserved** | Pointer to **bool** | Indicates whether the gateway address of the original network is kept. | [optional] 
**DefaultGatewayAssigned** | Pointer to **bool** | Indicates whether the first IP address after the network ID is assigned as the default gateway address. | [optional] 
**ConflictResolutionStrategy** | Pointer to **string** | The method used to resolve conflicts between template items and the resource the template is applied to. | [optional] 
**Template** | Pointer to [**InlinedIPv4Template**](InlinedIPv4Template.md) |  | [optional] 
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 

## Methods

### NewGetSplits200ResponseDataInner

`func NewGetSplits200ResponseDataInner() *GetSplits200ResponseDataInner`

NewGetSplits200ResponseDataInner instantiates a new GetSplits200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSplits200ResponseDataInnerWithDefaults

`func NewGetSplits200ResponseDataInnerWithDefaults() *GetSplits200ResponseDataInner`

NewGetSplits200ResponseDataInnerWithDefaults instantiates a new GetSplits200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetSplits200ResponseDataInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetSplits200ResponseDataInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetSplits200ResponseDataInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetSplits200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GetSplits200ResponseDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetSplits200ResponseDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetSplits200ResponseDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetSplits200ResponseDataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetBlock

`func (o *GetSplits200ResponseDataInner) GetBlock() InlinedIPv4Block`

GetBlock returns the Block field if non-nil, zero value otherwise.

### GetBlockOk

`func (o *GetSplits200ResponseDataInner) GetBlockOk() (*InlinedIPv4Block, bool)`

GetBlockOk returns a tuple with the Block field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlock

`func (o *GetSplits200ResponseDataInner) SetBlock(v InlinedIPv4Block)`

SetBlock sets Block field to given value.

### HasBlock

`func (o *GetSplits200ResponseDataInner) HasBlock() bool`

HasBlock returns a boolean if a field has been set.

### GetSplitPointAddress

`func (o *GetSplits200ResponseDataInner) GetSplitPointAddress() string`

GetSplitPointAddress returns the SplitPointAddress field if non-nil, zero value otherwise.

### GetSplitPointAddressOk

`func (o *GetSplits200ResponseDataInner) GetSplitPointAddressOk() (*string, bool)`

GetSplitPointAddressOk returns a tuple with the SplitPointAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitPointAddress

`func (o *GetSplits200ResponseDataInner) SetSplitPointAddress(v string)`

SetSplitPointAddress sets SplitPointAddress field to given value.

### HasSplitPointAddress

`func (o *GetSplits200ResponseDataInner) HasSplitPointAddress() bool`

HasSplitPointAddress returns a boolean if a field has been set.

### GetNetwork

`func (o *GetSplits200ResponseDataInner) GetNetwork() InlinedIPv4Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *GetSplits200ResponseDataInner) GetNetworkOk() (*InlinedIPv4Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *GetSplits200ResponseDataInner) SetNetwork(v InlinedIPv4Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *GetSplits200ResponseDataInner) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetNetworkCount

`func (o *GetSplits200ResponseDataInner) GetNetworkCount() int32`

GetNetworkCount returns the NetworkCount field if non-nil, zero value otherwise.

### GetNetworkCountOk

`func (o *GetSplits200ResponseDataInner) GetNetworkCountOk() (*int32, bool)`

GetNetworkCountOk returns a tuple with the NetworkCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkCount

`func (o *GetSplits200ResponseDataInner) SetNetworkCount(v int32)`

SetNetworkCount sets NetworkCount field to given value.

### HasNetworkCount

`func (o *GetSplits200ResponseDataInner) HasNetworkCount() bool`

HasNetworkCount returns a boolean if a field has been set.

### GetGatewayPreserved

`func (o *GetSplits200ResponseDataInner) GetGatewayPreserved() bool`

GetGatewayPreserved returns the GatewayPreserved field if non-nil, zero value otherwise.

### GetGatewayPreservedOk

`func (o *GetSplits200ResponseDataInner) GetGatewayPreservedOk() (*bool, bool)`

GetGatewayPreservedOk returns a tuple with the GatewayPreserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayPreserved

`func (o *GetSplits200ResponseDataInner) SetGatewayPreserved(v bool)`

SetGatewayPreserved sets GatewayPreserved field to given value.

### HasGatewayPreserved

`func (o *GetSplits200ResponseDataInner) HasGatewayPreserved() bool`

HasGatewayPreserved returns a boolean if a field has been set.

### GetDefaultGatewayAssigned

`func (o *GetSplits200ResponseDataInner) GetDefaultGatewayAssigned() bool`

GetDefaultGatewayAssigned returns the DefaultGatewayAssigned field if non-nil, zero value otherwise.

### GetDefaultGatewayAssignedOk

`func (o *GetSplits200ResponseDataInner) GetDefaultGatewayAssignedOk() (*bool, bool)`

GetDefaultGatewayAssignedOk returns a tuple with the DefaultGatewayAssigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultGatewayAssigned

`func (o *GetSplits200ResponseDataInner) SetDefaultGatewayAssigned(v bool)`

SetDefaultGatewayAssigned sets DefaultGatewayAssigned field to given value.

### HasDefaultGatewayAssigned

`func (o *GetSplits200ResponseDataInner) HasDefaultGatewayAssigned() bool`

HasDefaultGatewayAssigned returns a boolean if a field has been set.

### GetConflictResolutionStrategy

`func (o *GetSplits200ResponseDataInner) GetConflictResolutionStrategy() string`

GetConflictResolutionStrategy returns the ConflictResolutionStrategy field if non-nil, zero value otherwise.

### GetConflictResolutionStrategyOk

`func (o *GetSplits200ResponseDataInner) GetConflictResolutionStrategyOk() (*string, bool)`

GetConflictResolutionStrategyOk returns a tuple with the ConflictResolutionStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflictResolutionStrategy

`func (o *GetSplits200ResponseDataInner) SetConflictResolutionStrategy(v string)`

SetConflictResolutionStrategy sets ConflictResolutionStrategy field to given value.

### HasConflictResolutionStrategy

`func (o *GetSplits200ResponseDataInner) HasConflictResolutionStrategy() bool`

HasConflictResolutionStrategy returns a boolean if a field has been set.

### GetTemplate

`func (o *GetSplits200ResponseDataInner) GetTemplate() InlinedIPv4Template`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *GetSplits200ResponseDataInner) GetTemplateOk() (*InlinedIPv4Template, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *GetSplits200ResponseDataInner) SetTemplate(v InlinedIPv4Template)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *GetSplits200ResponseDataInner) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetLinks

`func (o *GetSplits200ResponseDataInner) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetSplits200ResponseDataInner) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetSplits200ResponseDataInner) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetSplits200ResponseDataInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


