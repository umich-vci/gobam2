# IPv4ReconciliationPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The resource type. | [optional] 
**DiscoveryMethod** | Pointer to **string** | The discovery method for the reconciliation policy. | [optional] 
**NetworkBoundaries** | Pointer to **[]string** | The IPv4 range(s) in CIDR format for the reconciliation policy to search. | [optional] 
**PingSweepNetworkGaps** | Pointer to **[]string** | The IPv4 address ranges in CIDR format for which ping sweep sends an ICMP echo request. | [optional] 
**Overrides** | Pointer to **[]string** | The IPv4 addresses and ranges for the reconciliation policy to ignore. The list of overrides can contain single addresses, ranges in CIDR format (e.g. 10.0.0.0/16), or ranges in address range format (e.g. 10.0.0.0-10.0.255.255). | [optional] 
**ReclaimAcceptanceAge** | Pointer to **string** | The age parameter for automatic reconciliation of reclaimable addresses. Reconciliation will be applied to reclaimable IP addresses older than this value. Reclaimable addresses are addresses that exist in Address Manager but are not found on the physical network. | [optional] 
**UnknownAcceptanceAge** | Pointer to **string** | The age parameter for automatic reconciliation of unknown addresses. Reconciliation will be applied to unknown IP addresses older than this value. Unknown addresses are addresses that exist on the physical network but do not exist in Address Manager. | [optional] 
**MismatchAcceptanceAge** | Pointer to **string** | The age parameter for automatic reconciliation of mismatched addresses. Reconciliation will be applied to mismatched IP addresses older than this value. Mismatched addresses are addresses that exist in both Address Manager and the physical network, but do not match in either the MAC address, DNS host name info, VLAN info, or connect switch port info. | [optional] 
**View** | Pointer to [**InlinedView**](InlinedView.md) |  | [optional] 

## Methods

### NewIPv4ReconciliationPolicy

`func NewIPv4ReconciliationPolicy() *IPv4ReconciliationPolicy`

NewIPv4ReconciliationPolicy instantiates a new IPv4ReconciliationPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPv4ReconciliationPolicyWithDefaults

`func NewIPv4ReconciliationPolicyWithDefaults() *IPv4ReconciliationPolicy`

NewIPv4ReconciliationPolicyWithDefaults instantiates a new IPv4ReconciliationPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *IPv4ReconciliationPolicy) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IPv4ReconciliationPolicy) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IPv4ReconciliationPolicy) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IPv4ReconciliationPolicy) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDiscoveryMethod

`func (o *IPv4ReconciliationPolicy) GetDiscoveryMethod() string`

GetDiscoveryMethod returns the DiscoveryMethod field if non-nil, zero value otherwise.

### GetDiscoveryMethodOk

`func (o *IPv4ReconciliationPolicy) GetDiscoveryMethodOk() (*string, bool)`

GetDiscoveryMethodOk returns a tuple with the DiscoveryMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryMethod

`func (o *IPv4ReconciliationPolicy) SetDiscoveryMethod(v string)`

SetDiscoveryMethod sets DiscoveryMethod field to given value.

### HasDiscoveryMethod

`func (o *IPv4ReconciliationPolicy) HasDiscoveryMethod() bool`

HasDiscoveryMethod returns a boolean if a field has been set.

### GetNetworkBoundaries

`func (o *IPv4ReconciliationPolicy) GetNetworkBoundaries() []string`

GetNetworkBoundaries returns the NetworkBoundaries field if non-nil, zero value otherwise.

### GetNetworkBoundariesOk

`func (o *IPv4ReconciliationPolicy) GetNetworkBoundariesOk() (*[]string, bool)`

GetNetworkBoundariesOk returns a tuple with the NetworkBoundaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkBoundaries

`func (o *IPv4ReconciliationPolicy) SetNetworkBoundaries(v []string)`

SetNetworkBoundaries sets NetworkBoundaries field to given value.

### HasNetworkBoundaries

`func (o *IPv4ReconciliationPolicy) HasNetworkBoundaries() bool`

HasNetworkBoundaries returns a boolean if a field has been set.

### GetPingSweepNetworkGaps

`func (o *IPv4ReconciliationPolicy) GetPingSweepNetworkGaps() []string`

GetPingSweepNetworkGaps returns the PingSweepNetworkGaps field if non-nil, zero value otherwise.

### GetPingSweepNetworkGapsOk

`func (o *IPv4ReconciliationPolicy) GetPingSweepNetworkGapsOk() (*[]string, bool)`

GetPingSweepNetworkGapsOk returns a tuple with the PingSweepNetworkGaps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingSweepNetworkGaps

`func (o *IPv4ReconciliationPolicy) SetPingSweepNetworkGaps(v []string)`

SetPingSweepNetworkGaps sets PingSweepNetworkGaps field to given value.

### HasPingSweepNetworkGaps

`func (o *IPv4ReconciliationPolicy) HasPingSweepNetworkGaps() bool`

HasPingSweepNetworkGaps returns a boolean if a field has been set.

### GetOverrides

`func (o *IPv4ReconciliationPolicy) GetOverrides() []string`

GetOverrides returns the Overrides field if non-nil, zero value otherwise.

### GetOverridesOk

`func (o *IPv4ReconciliationPolicy) GetOverridesOk() (*[]string, bool)`

GetOverridesOk returns a tuple with the Overrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrides

`func (o *IPv4ReconciliationPolicy) SetOverrides(v []string)`

SetOverrides sets Overrides field to given value.

### HasOverrides

`func (o *IPv4ReconciliationPolicy) HasOverrides() bool`

HasOverrides returns a boolean if a field has been set.

### GetReclaimAcceptanceAge

`func (o *IPv4ReconciliationPolicy) GetReclaimAcceptanceAge() string`

GetReclaimAcceptanceAge returns the ReclaimAcceptanceAge field if non-nil, zero value otherwise.

### GetReclaimAcceptanceAgeOk

`func (o *IPv4ReconciliationPolicy) GetReclaimAcceptanceAgeOk() (*string, bool)`

GetReclaimAcceptanceAgeOk returns a tuple with the ReclaimAcceptanceAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReclaimAcceptanceAge

`func (o *IPv4ReconciliationPolicy) SetReclaimAcceptanceAge(v string)`

SetReclaimAcceptanceAge sets ReclaimAcceptanceAge field to given value.

### HasReclaimAcceptanceAge

`func (o *IPv4ReconciliationPolicy) HasReclaimAcceptanceAge() bool`

HasReclaimAcceptanceAge returns a boolean if a field has been set.

### GetUnknownAcceptanceAge

`func (o *IPv4ReconciliationPolicy) GetUnknownAcceptanceAge() string`

GetUnknownAcceptanceAge returns the UnknownAcceptanceAge field if non-nil, zero value otherwise.

### GetUnknownAcceptanceAgeOk

`func (o *IPv4ReconciliationPolicy) GetUnknownAcceptanceAgeOk() (*string, bool)`

GetUnknownAcceptanceAgeOk returns a tuple with the UnknownAcceptanceAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownAcceptanceAge

`func (o *IPv4ReconciliationPolicy) SetUnknownAcceptanceAge(v string)`

SetUnknownAcceptanceAge sets UnknownAcceptanceAge field to given value.

### HasUnknownAcceptanceAge

`func (o *IPv4ReconciliationPolicy) HasUnknownAcceptanceAge() bool`

HasUnknownAcceptanceAge returns a boolean if a field has been set.

### GetMismatchAcceptanceAge

`func (o *IPv4ReconciliationPolicy) GetMismatchAcceptanceAge() string`

GetMismatchAcceptanceAge returns the MismatchAcceptanceAge field if non-nil, zero value otherwise.

### GetMismatchAcceptanceAgeOk

`func (o *IPv4ReconciliationPolicy) GetMismatchAcceptanceAgeOk() (*string, bool)`

GetMismatchAcceptanceAgeOk returns a tuple with the MismatchAcceptanceAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMismatchAcceptanceAge

`func (o *IPv4ReconciliationPolicy) SetMismatchAcceptanceAge(v string)`

SetMismatchAcceptanceAge sets MismatchAcceptanceAge field to given value.

### HasMismatchAcceptanceAge

`func (o *IPv4ReconciliationPolicy) HasMismatchAcceptanceAge() bool`

HasMismatchAcceptanceAge returns a boolean if a field has been set.

### GetView

`func (o *IPv4ReconciliationPolicy) GetView() InlinedView`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *IPv4ReconciliationPolicy) GetViewOk() (*InlinedView, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *IPv4ReconciliationPolicy) SetView(v InlinedView)`

SetView sets View field to given value.

### HasView

`func (o *IPv4ReconciliationPolicy) HasView() bool`

HasView returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


