# FirewallService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The resource type. | [optional] 
**Enabled** | Pointer to **bool** | Indicates whether the firewall service is enabled. | [optional] 
**PingAllowed** | Pointer to **bool** | Indicates whether the server responds to the ping command. | [optional] 

## Methods

### NewFirewallService

`func NewFirewallService() *FirewallService`

NewFirewallService instantiates a new FirewallService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallServiceWithDefaults

`func NewFirewallServiceWithDefaults() *FirewallService`

NewFirewallServiceWithDefaults instantiates a new FirewallService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FirewallService) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FirewallService) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FirewallService) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FirewallService) HasType() bool`

HasType returns a boolean if a field has been set.

### GetEnabled

`func (o *FirewallService) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *FirewallService) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *FirewallService) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *FirewallService) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPingAllowed

`func (o *FirewallService) GetPingAllowed() bool`

GetPingAllowed returns the PingAllowed field if non-nil, zero value otherwise.

### GetPingAllowedOk

`func (o *FirewallService) GetPingAllowedOk() (*bool, bool)`

GetPingAllowedOk returns a tuple with the PingAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingAllowed

`func (o *FirewallService) SetPingAllowed(v bool)`

SetPingAllowed sets PingAllowed field to given value.

### HasPingAllowed

`func (o *FirewallService) HasPingAllowed() bool`

HasPingAllowed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


