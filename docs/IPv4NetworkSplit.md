# IPv4NetworkSplit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The resource type. | [optional] 
**Network** | Pointer to [**InlinedIPv4Network**](InlinedIPv4Network.md) |  | [optional] 
**NetworkCount** | Pointer to **int32** | The number of resultant sub-networks after the split. | [optional] 
**GatewayPreserved** | Pointer to **bool** | Indicates whether the gateway address of the original network is kept. | [optional] 
**DefaultGatewayAssigned** | Pointer to **bool** | Indicates whether the first IP address after the network ID is assigned as the default gateway address. | [optional] 
**ConflictResolutionStrategy** | Pointer to **string** | The method used to resolve conflicts between template items and the resource the template is applied to. | [optional] 
**Template** | Pointer to [**InlinedIPv4Template**](InlinedIPv4Template.md) |  | [optional] 

## Methods

### NewIPv4NetworkSplit

`func NewIPv4NetworkSplit() *IPv4NetworkSplit`

NewIPv4NetworkSplit instantiates a new IPv4NetworkSplit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPv4NetworkSplitWithDefaults

`func NewIPv4NetworkSplitWithDefaults() *IPv4NetworkSplit`

NewIPv4NetworkSplitWithDefaults instantiates a new IPv4NetworkSplit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *IPv4NetworkSplit) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IPv4NetworkSplit) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IPv4NetworkSplit) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IPv4NetworkSplit) HasType() bool`

HasType returns a boolean if a field has been set.

### GetNetwork

`func (o *IPv4NetworkSplit) GetNetwork() InlinedIPv4Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *IPv4NetworkSplit) GetNetworkOk() (*InlinedIPv4Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *IPv4NetworkSplit) SetNetwork(v InlinedIPv4Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *IPv4NetworkSplit) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetNetworkCount

`func (o *IPv4NetworkSplit) GetNetworkCount() int32`

GetNetworkCount returns the NetworkCount field if non-nil, zero value otherwise.

### GetNetworkCountOk

`func (o *IPv4NetworkSplit) GetNetworkCountOk() (*int32, bool)`

GetNetworkCountOk returns a tuple with the NetworkCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkCount

`func (o *IPv4NetworkSplit) SetNetworkCount(v int32)`

SetNetworkCount sets NetworkCount field to given value.

### HasNetworkCount

`func (o *IPv4NetworkSplit) HasNetworkCount() bool`

HasNetworkCount returns a boolean if a field has been set.

### GetGatewayPreserved

`func (o *IPv4NetworkSplit) GetGatewayPreserved() bool`

GetGatewayPreserved returns the GatewayPreserved field if non-nil, zero value otherwise.

### GetGatewayPreservedOk

`func (o *IPv4NetworkSplit) GetGatewayPreservedOk() (*bool, bool)`

GetGatewayPreservedOk returns a tuple with the GatewayPreserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayPreserved

`func (o *IPv4NetworkSplit) SetGatewayPreserved(v bool)`

SetGatewayPreserved sets GatewayPreserved field to given value.

### HasGatewayPreserved

`func (o *IPv4NetworkSplit) HasGatewayPreserved() bool`

HasGatewayPreserved returns a boolean if a field has been set.

### GetDefaultGatewayAssigned

`func (o *IPv4NetworkSplit) GetDefaultGatewayAssigned() bool`

GetDefaultGatewayAssigned returns the DefaultGatewayAssigned field if non-nil, zero value otherwise.

### GetDefaultGatewayAssignedOk

`func (o *IPv4NetworkSplit) GetDefaultGatewayAssignedOk() (*bool, bool)`

GetDefaultGatewayAssignedOk returns a tuple with the DefaultGatewayAssigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultGatewayAssigned

`func (o *IPv4NetworkSplit) SetDefaultGatewayAssigned(v bool)`

SetDefaultGatewayAssigned sets DefaultGatewayAssigned field to given value.

### HasDefaultGatewayAssigned

`func (o *IPv4NetworkSplit) HasDefaultGatewayAssigned() bool`

HasDefaultGatewayAssigned returns a boolean if a field has been set.

### GetConflictResolutionStrategy

`func (o *IPv4NetworkSplit) GetConflictResolutionStrategy() string`

GetConflictResolutionStrategy returns the ConflictResolutionStrategy field if non-nil, zero value otherwise.

### GetConflictResolutionStrategyOk

`func (o *IPv4NetworkSplit) GetConflictResolutionStrategyOk() (*string, bool)`

GetConflictResolutionStrategyOk returns a tuple with the ConflictResolutionStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflictResolutionStrategy

`func (o *IPv4NetworkSplit) SetConflictResolutionStrategy(v string)`

SetConflictResolutionStrategy sets ConflictResolutionStrategy field to given value.

### HasConflictResolutionStrategy

`func (o *IPv4NetworkSplit) HasConflictResolutionStrategy() bool`

HasConflictResolutionStrategy returns a boolean if a field has been set.

### GetTemplate

`func (o *IPv4NetworkSplit) GetTemplate() InlinedIPv4Template`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *IPv4NetworkSplit) GetTemplateOk() (*InlinedIPv4Template, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *IPv4NetworkSplit) SetTemplate(v InlinedIPv4Template)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *IPv4NetworkSplit) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


