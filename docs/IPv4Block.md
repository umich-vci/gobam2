# IPv4Block

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of IP block. | [optional] 
**Template** | Pointer to [**InlinedIPv4Template**](InlinedIPv4Template.md) |  | [optional] 
**DefaultView** | Pointer to [**InlinedView**](InlinedView.md) |  | [optional] 
**DuplicateHostnamesAllowed** | Pointer to **bool** | Indicates whether duplicate hostnames are allowed in the block. If enabled, duplicate hostnames will be allowed. If not set explicitly, this value will be inherited from the parent block. | [optional] 
**PingBeforeAssignEnabled** | Pointer to **bool** | If set to true, Address Manager will ping IP addresses before assigning them. If set to false, Address Manager will assign addresses without checking their availability. | [optional] 
**DefaultZonesInherited** | Pointer to **bool** | If set to true, the IPv4 block will inherit the parent resource&#39;s default domains configuration if possible. | [optional] 
**DefaultZones** | Pointer to [**[]IPv4BlockAllOfDefaultZones**](IPv4BlockAllOfDefaultZones.md) |  | [optional] 
**RestrictedZonesInherited** | Pointer to **bool** | If set to true, the IPv4 block will inherit the parent resource&#39;s DNS restriction configuration if possible. | [optional] 
**RestrictedZones** | Pointer to [**[]IPv4BlockAllOfRestrictedZones**](IPv4BlockAllOfRestrictedZones.md) |  | [optional] 
**LowWaterMark** | Pointer to **int32** | A DHCP alert is triggered when usage falls below this percentage (when too few addresses are in use). | [optional] 
**HighWaterMark** | Pointer to **int32** | A DHCP alert is triggered when usage exceeds this percentage (when too many addresses are in use). | [optional] 
**ReverseZoneSigned** | Pointer to **bool** | If set to true, the reverse zone associated with the IPv4 block will be signed according to the DNSSEC signing policy in reverseZoneSigningPolicy. | [optional] 
**ReverseZoneSigningPolicy** | Pointer to [**InlinedDNSSECSigningPolicy**](InlinedDNSSECSigningPolicy.md) |  | [optional] 
**UsagePercentage** | Pointer to **map[string]int32** | Percentage values indicating the current amount of allocated and unallocated IPv4 addresses within the block. | [optional] [readonly] 

## Methods

### NewIPv4Block

`func NewIPv4Block() *IPv4Block`

NewIPv4Block instantiates a new IPv4Block object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPv4BlockWithDefaults

`func NewIPv4BlockWithDefaults() *IPv4Block`

NewIPv4BlockWithDefaults instantiates a new IPv4Block object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *IPv4Block) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IPv4Block) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IPv4Block) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IPv4Block) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTemplate

`func (o *IPv4Block) GetTemplate() InlinedIPv4Template`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *IPv4Block) GetTemplateOk() (*InlinedIPv4Template, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *IPv4Block) SetTemplate(v InlinedIPv4Template)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *IPv4Block) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetDefaultView

`func (o *IPv4Block) GetDefaultView() InlinedView`

GetDefaultView returns the DefaultView field if non-nil, zero value otherwise.

### GetDefaultViewOk

`func (o *IPv4Block) GetDefaultViewOk() (*InlinedView, bool)`

GetDefaultViewOk returns a tuple with the DefaultView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultView

`func (o *IPv4Block) SetDefaultView(v InlinedView)`

SetDefaultView sets DefaultView field to given value.

### HasDefaultView

`func (o *IPv4Block) HasDefaultView() bool`

HasDefaultView returns a boolean if a field has been set.

### GetDuplicateHostnamesAllowed

`func (o *IPv4Block) GetDuplicateHostnamesAllowed() bool`

GetDuplicateHostnamesAllowed returns the DuplicateHostnamesAllowed field if non-nil, zero value otherwise.

### GetDuplicateHostnamesAllowedOk

`func (o *IPv4Block) GetDuplicateHostnamesAllowedOk() (*bool, bool)`

GetDuplicateHostnamesAllowedOk returns a tuple with the DuplicateHostnamesAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplicateHostnamesAllowed

`func (o *IPv4Block) SetDuplicateHostnamesAllowed(v bool)`

SetDuplicateHostnamesAllowed sets DuplicateHostnamesAllowed field to given value.

### HasDuplicateHostnamesAllowed

`func (o *IPv4Block) HasDuplicateHostnamesAllowed() bool`

HasDuplicateHostnamesAllowed returns a boolean if a field has been set.

### GetPingBeforeAssignEnabled

`func (o *IPv4Block) GetPingBeforeAssignEnabled() bool`

GetPingBeforeAssignEnabled returns the PingBeforeAssignEnabled field if non-nil, zero value otherwise.

### GetPingBeforeAssignEnabledOk

`func (o *IPv4Block) GetPingBeforeAssignEnabledOk() (*bool, bool)`

GetPingBeforeAssignEnabledOk returns a tuple with the PingBeforeAssignEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingBeforeAssignEnabled

`func (o *IPv4Block) SetPingBeforeAssignEnabled(v bool)`

SetPingBeforeAssignEnabled sets PingBeforeAssignEnabled field to given value.

### HasPingBeforeAssignEnabled

`func (o *IPv4Block) HasPingBeforeAssignEnabled() bool`

HasPingBeforeAssignEnabled returns a boolean if a field has been set.

### GetDefaultZonesInherited

`func (o *IPv4Block) GetDefaultZonesInherited() bool`

GetDefaultZonesInherited returns the DefaultZonesInherited field if non-nil, zero value otherwise.

### GetDefaultZonesInheritedOk

`func (o *IPv4Block) GetDefaultZonesInheritedOk() (*bool, bool)`

GetDefaultZonesInheritedOk returns a tuple with the DefaultZonesInherited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultZonesInherited

`func (o *IPv4Block) SetDefaultZonesInherited(v bool)`

SetDefaultZonesInherited sets DefaultZonesInherited field to given value.

### HasDefaultZonesInherited

`func (o *IPv4Block) HasDefaultZonesInherited() bool`

HasDefaultZonesInherited returns a boolean if a field has been set.

### GetDefaultZones

`func (o *IPv4Block) GetDefaultZones() []IPv4BlockAllOfDefaultZones`

GetDefaultZones returns the DefaultZones field if non-nil, zero value otherwise.

### GetDefaultZonesOk

`func (o *IPv4Block) GetDefaultZonesOk() (*[]IPv4BlockAllOfDefaultZones, bool)`

GetDefaultZonesOk returns a tuple with the DefaultZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultZones

`func (o *IPv4Block) SetDefaultZones(v []IPv4BlockAllOfDefaultZones)`

SetDefaultZones sets DefaultZones field to given value.

### HasDefaultZones

`func (o *IPv4Block) HasDefaultZones() bool`

HasDefaultZones returns a boolean if a field has been set.

### GetRestrictedZonesInherited

`func (o *IPv4Block) GetRestrictedZonesInherited() bool`

GetRestrictedZonesInherited returns the RestrictedZonesInherited field if non-nil, zero value otherwise.

### GetRestrictedZonesInheritedOk

`func (o *IPv4Block) GetRestrictedZonesInheritedOk() (*bool, bool)`

GetRestrictedZonesInheritedOk returns a tuple with the RestrictedZonesInherited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedZonesInherited

`func (o *IPv4Block) SetRestrictedZonesInherited(v bool)`

SetRestrictedZonesInherited sets RestrictedZonesInherited field to given value.

### HasRestrictedZonesInherited

`func (o *IPv4Block) HasRestrictedZonesInherited() bool`

HasRestrictedZonesInherited returns a boolean if a field has been set.

### GetRestrictedZones

`func (o *IPv4Block) GetRestrictedZones() []IPv4BlockAllOfRestrictedZones`

GetRestrictedZones returns the RestrictedZones field if non-nil, zero value otherwise.

### GetRestrictedZonesOk

`func (o *IPv4Block) GetRestrictedZonesOk() (*[]IPv4BlockAllOfRestrictedZones, bool)`

GetRestrictedZonesOk returns a tuple with the RestrictedZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedZones

`func (o *IPv4Block) SetRestrictedZones(v []IPv4BlockAllOfRestrictedZones)`

SetRestrictedZones sets RestrictedZones field to given value.

### HasRestrictedZones

`func (o *IPv4Block) HasRestrictedZones() bool`

HasRestrictedZones returns a boolean if a field has been set.

### GetLowWaterMark

`func (o *IPv4Block) GetLowWaterMark() int32`

GetLowWaterMark returns the LowWaterMark field if non-nil, zero value otherwise.

### GetLowWaterMarkOk

`func (o *IPv4Block) GetLowWaterMarkOk() (*int32, bool)`

GetLowWaterMarkOk returns a tuple with the LowWaterMark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowWaterMark

`func (o *IPv4Block) SetLowWaterMark(v int32)`

SetLowWaterMark sets LowWaterMark field to given value.

### HasLowWaterMark

`func (o *IPv4Block) HasLowWaterMark() bool`

HasLowWaterMark returns a boolean if a field has been set.

### GetHighWaterMark

`func (o *IPv4Block) GetHighWaterMark() int32`

GetHighWaterMark returns the HighWaterMark field if non-nil, zero value otherwise.

### GetHighWaterMarkOk

`func (o *IPv4Block) GetHighWaterMarkOk() (*int32, bool)`

GetHighWaterMarkOk returns a tuple with the HighWaterMark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighWaterMark

`func (o *IPv4Block) SetHighWaterMark(v int32)`

SetHighWaterMark sets HighWaterMark field to given value.

### HasHighWaterMark

`func (o *IPv4Block) HasHighWaterMark() bool`

HasHighWaterMark returns a boolean if a field has been set.

### GetReverseZoneSigned

`func (o *IPv4Block) GetReverseZoneSigned() bool`

GetReverseZoneSigned returns the ReverseZoneSigned field if non-nil, zero value otherwise.

### GetReverseZoneSignedOk

`func (o *IPv4Block) GetReverseZoneSignedOk() (*bool, bool)`

GetReverseZoneSignedOk returns a tuple with the ReverseZoneSigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReverseZoneSigned

`func (o *IPv4Block) SetReverseZoneSigned(v bool)`

SetReverseZoneSigned sets ReverseZoneSigned field to given value.

### HasReverseZoneSigned

`func (o *IPv4Block) HasReverseZoneSigned() bool`

HasReverseZoneSigned returns a boolean if a field has been set.

### GetReverseZoneSigningPolicy

`func (o *IPv4Block) GetReverseZoneSigningPolicy() InlinedDNSSECSigningPolicy`

GetReverseZoneSigningPolicy returns the ReverseZoneSigningPolicy field if non-nil, zero value otherwise.

### GetReverseZoneSigningPolicyOk

`func (o *IPv4Block) GetReverseZoneSigningPolicyOk() (*InlinedDNSSECSigningPolicy, bool)`

GetReverseZoneSigningPolicyOk returns a tuple with the ReverseZoneSigningPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReverseZoneSigningPolicy

`func (o *IPv4Block) SetReverseZoneSigningPolicy(v InlinedDNSSECSigningPolicy)`

SetReverseZoneSigningPolicy sets ReverseZoneSigningPolicy field to given value.

### HasReverseZoneSigningPolicy

`func (o *IPv4Block) HasReverseZoneSigningPolicy() bool`

HasReverseZoneSigningPolicy returns a boolean if a field has been set.

### GetUsagePercentage

`func (o *IPv4Block) GetUsagePercentage() map[string]int32`

GetUsagePercentage returns the UsagePercentage field if non-nil, zero value otherwise.

### GetUsagePercentageOk

`func (o *IPv4Block) GetUsagePercentageOk() (*map[string]int32, bool)`

GetUsagePercentageOk returns a tuple with the UsagePercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsagePercentage

`func (o *IPv4Block) SetUsagePercentage(v map[string]int32)`

SetUsagePercentage sets UsagePercentage field to given value.

### HasUsagePercentage

`func (o *IPv4Block) HasUsagePercentage() bool`

HasUsagePercentage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


