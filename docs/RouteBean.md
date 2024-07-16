# RouteBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gateway** | Pointer to **string** |  | [optional] 
**Network** | Pointer to **string** |  | [optional] 
**PrefixLength** | Pointer to **int32** |  | [optional] 

## Methods

### NewRouteBean

`func NewRouteBean() *RouteBean`

NewRouteBean instantiates a new RouteBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteBeanWithDefaults

`func NewRouteBeanWithDefaults() *RouteBean`

NewRouteBeanWithDefaults instantiates a new RouteBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGateway

`func (o *RouteBean) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *RouteBean) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *RouteBean) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *RouteBean) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetNetwork

`func (o *RouteBean) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *RouteBean) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *RouteBean) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *RouteBean) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetPrefixLength

`func (o *RouteBean) GetPrefixLength() int32`

GetPrefixLength returns the PrefixLength field if non-nil, zero value otherwise.

### GetPrefixLengthOk

`func (o *RouteBean) GetPrefixLengthOk() (*int32, bool)`

GetPrefixLengthOk returns a tuple with the PrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixLength

`func (o *RouteBean) SetPrefixLength(v int32)`

SetPrefixLength sets PrefixLength field to given value.

### HasPrefixLength

`func (o *RouteBean) HasPrefixLength() bool`

HasPrefixLength returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


