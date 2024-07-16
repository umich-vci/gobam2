# GetCollectionRestrictedRanges200Response1DataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The type of IP block. | [optional] 
**Name** | Pointer to **string** | The name of the resource. | [optional] 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**Configuration** | Pointer to [**InlinedConfiguration**](InlinedConfiguration.md) |  | [optional] [readonly] 
**Range** | Pointer to **string** | The address range for the IP block. | [optional] 
**Location** | Pointer to [**InlinedLocation**](InlinedLocation.md) |  | [optional] 
**InheritedFields** | Pointer to **[]string** |  | [optional] [readonly] 
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

### NewGetCollectionRestrictedRanges200Response1DataInner

`func NewGetCollectionRestrictedRanges200Response1DataInner() *GetCollectionRestrictedRanges200Response1DataInner`

NewGetCollectionRestrictedRanges200Response1DataInner instantiates a new GetCollectionRestrictedRanges200Response1DataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCollectionRestrictedRanges200Response1DataInnerWithDefaults

`func NewGetCollectionRestrictedRanges200Response1DataInnerWithDefaults() *GetCollectionRestrictedRanges200Response1DataInner`

NewGetCollectionRestrictedRanges200Response1DataInnerWithDefaults instantiates a new GetCollectionRestrictedRanges200Response1DataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetCollectionRestrictedRanges200Response1DataInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetCollectionRestrictedRanges200Response1DataInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetCollectionRestrictedRanges200Response1DataInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetCollectionRestrictedRanges200Response1DataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GetCollectionRestrictedRanges200Response1DataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetCollectionRestrictedRanges200Response1DataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetCollectionRestrictedRanges200Response1DataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetCollectionRestrictedRanges200Response1DataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *GetCollectionRestrictedRanges200Response1DataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetCollectionRestrictedRanges200Response1DataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetCollectionRestrictedRanges200Response1DataInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetCollectionRestrictedRanges200Response1DataInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *GetCollectionRestrictedRanges200Response1DataInner) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *GetCollectionRestrictedRanges200Response1DataInner) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *GetCollectionRestrictedRanges200Response1DataInner) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *GetCollectionRestrictedRanges200Response1DataInner) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetConfiguration

`func (o *GetCollectionRestrictedRanges200Response1DataInner) GetConfiguration() InlinedConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *GetCollectionRestrictedRanges200Response1DataInner) GetConfigurationOk() (*InlinedConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *GetCollectionRestrictedRanges200Response1DataInner) SetConfiguration(v InlinedConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *GetCollectionRestrictedRanges200Response1DataInner) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetRange

`func (o *GetCollectionRestrictedRanges200Response1DataInner) GetRange() string`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *GetCollectionRestrictedRanges200Response1DataInner) GetRangeOk() (*string, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *GetCollectionRestrictedRanges200Response1DataInner) SetRange(v string)`

SetRange sets Range field to given value.

### HasRange

`func (o *GetCollectionRestrictedRanges200Response1DataInner) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetLocation

`func (o *GetCollectionRestrictedRanges200Response1DataInner) GetLocation() InlinedLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *GetCollectionRestrictedRanges200Response1DataInner) GetLocationOk() (*InlinedLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *GetCollectionRestrictedRanges200Response1DataInner) SetLocation(v InlinedLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *GetCollectionRestrictedRanges200Response1DataInner) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetInheritedFields

`func (o *GetCollectionRestrictedRanges200Response1DataInner) GetInheritedFields() []string`

GetInheritedFields returns the InheritedFields field if non-nil, zero value otherwise.

### GetInheritedFieldsOk

`func (o *GetCollectionRestrictedRanges200Response1DataInner) GetInheritedFieldsOk() (*[]string, bool)`

GetInheritedFieldsOk returns a tuple with the InheritedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritedFields

`func (o *GetCollectionRestrictedRanges200Response1DataInner) SetInheritedFields(v []string)`

SetInheritedFields sets InheritedFields field to given value.

### HasInheritedFields

`func (o *GetCollectionRestrictedRanges200Response1DataInner) HasInheritedFields() bool`

HasInheritedFields returns a boolean if a field has been set.

### GetTemplate

`func (o *GetCollectionRestrictedRanges200Response1DataInner) GetTemplate() InlinedIPv4Template`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *GetCollectionRestrictedRanges200Response1DataInner) GetTemplateOk() (*InlinedIPv4Template, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *GetCollectionRestrictedRanges200Response1DataInner) SetTemplate(v InlinedIPv4Template)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *GetCollectionRestrictedRanges200Response1DataInner) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetDefaultView

`func (o *GetCollectionRestrictedRanges200Response1DataInner) GetDefaultView() InlinedView`

GetDefaultView returns the DefaultView field if non-nil, zero value otherwise.

### GetDefaultViewOk

`func (o *GetCollectionRestrictedRanges200Response1DataInner) GetDefaultViewOk() (*InlinedView, bool)`

GetDefaultViewOk returns a tuple with the DefaultView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultView

`func (o *GetCollectionRestrictedRanges200Response1DataInner) SetDefaultView(v InlinedView)`

SetDefaultView sets DefaultView field to given value.

### HasDefaultView

`func (o *GetCollectionRestrictedRanges200Response1DataInner) HasDefaultView() bool`

HasDefaultView returns a boolean if a field has been set.

### GetDuplicateHostnamesAllowed

`func (o *GetCollectionRestrictedRanges200Response1DataInner) GetDuplicateHostnamesAllowed() bool`

GetDuplicateHostnamesAllowed returns the DuplicateHostnamesAllowed field if non-nil, zero value otherwise.

### GetDuplicateHostnamesAllowedOk

`func (o *GetCollectionRestrictedRanges200Response1DataInner) GetDuplicateHostnamesAllowedOk() (*bool, bool)`

GetDuplicateHostnamesAllowedOk returns a tuple with the DuplicateHostnamesAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplicateHostnamesAllowed

`func (o *GetCollectionRestrictedRanges200Response1DataInner) SetDuplicateHostnamesAllowed(v bool)`

SetDuplicateHostnamesAllowed sets DuplicateHostnamesAllowed field to given value.

### HasDuplicateHostnamesAllowed

`func (o *GetCollectionRestrictedRanges200Response1DataInner) HasDuplicateHostnamesAllowed() bool`

HasDuplicateHostnamesAllowed returns a boolean if a field has been set.

### GetPingBeforeAssignEnabled

`func (o *GetCollectionRestrictedRanges200Response1DataInner) GetPingBeforeAssignEnabled() bool`

GetPingBeforeAssignEnabled returns the PingBeforeAssignEnabled field if non-nil, zero value otherwise.

### GetPingBeforeAssignEnabledOk

`func (o *GetCollectionRestrictedRanges200Response1DataInner) GetPingBeforeAssignEnabledOk() (*bool, bool)`

GetPingBeforeAssignEnabledOk returns a tuple with the PingBeforeAssignEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingBeforeAssignEnabled

`func (o *GetCollectionRestrictedRanges200Response1DataInner) SetPingBeforeAssignEnabled(v bool)`

SetPingBeforeAssignEnabled sets PingBeforeAssignEnabled field to given value.

### HasPingBeforeAssignEnabled

`func (o *GetCollectionRestrictedRanges200Response1DataInner) HasPingBeforeAssignEnabled() bool`

HasPingBeforeAssignEnabled returns a boolean if a field has been set.

### GetDefaultZonesInherited

`func (o *GetCollectionRestrictedRanges200Response1DataInner) GetDefaultZonesInherited() bool`

GetDefaultZonesInherited returns the DefaultZonesInherited field if non-nil, zero value otherwise.

### GetDefaultZonesInheritedOk

`func (o *GetCollectionRestrictedRanges200Response1DataInner) GetDefaultZonesInheritedOk() (*bool, bool)`

GetDefaultZonesInheritedOk returns a tuple with the DefaultZonesInherited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultZonesInherited

`func (o *GetCollectionRestrictedRanges200Response1DataInner) SetDefaultZonesInherited(v bool)`

SetDefaultZonesInherited sets DefaultZonesInherited field to given value.

### HasDefaultZonesInherited

`func (o *GetCollectionRestrictedRanges200Response1DataInner) HasDefaultZonesInherited() bool`

HasDefaultZonesInherited returns a boolean if a field has been set.

### GetDefaultZones

`func (o *GetCollectionRestrictedRanges200Response1DataInner) GetDefaultZones() []IPv4BlockAllOfDefaultZones`

GetDefaultZones returns the DefaultZones field if non-nil, zero value otherwise.

### GetDefaultZonesOk

`func (o *GetCollectionRestrictedRanges200Response1DataInner) GetDefaultZonesOk() (*[]IPv4BlockAllOfDefaultZones, bool)`

GetDefaultZonesOk returns a tuple with the DefaultZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultZones

`func (o *GetCollectionRestrictedRanges200Response1DataInner) SetDefaultZones(v []IPv4BlockAllOfDefaultZones)`

SetDefaultZones sets DefaultZones field to given value.

### HasDefaultZones

`func (o *GetCollectionRestrictedRanges200Response1DataInner) HasDefaultZones() bool`

HasDefaultZones returns a boolean if a field has been set.

### GetRestrictedZonesInherited

`func (o *GetCollectionRestrictedRanges200Response1DataInner) GetRestrictedZonesInherited() bool`

GetRestrictedZonesInherited returns the RestrictedZonesInherited field if non-nil, zero value otherwise.

### GetRestrictedZonesInheritedOk

`func (o *GetCollectionRestrictedRanges200Response1DataInner) GetRestrictedZonesInheritedOk() (*bool, bool)`

GetRestrictedZonesInheritedOk returns a tuple with the RestrictedZonesInherited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedZonesInherited

`func (o *GetCollectionRestrictedRanges200Response1DataInner) SetRestrictedZonesInherited(v bool)`

SetRestrictedZonesInherited sets RestrictedZonesInherited field to given value.

### HasRestrictedZonesInherited

`func (o *GetCollectionRestrictedRanges200Response1DataInner) HasRestrictedZonesInherited() bool`

HasRestrictedZonesInherited returns a boolean if a field has been set.

### GetRestrictedZones

`func (o *GetCollectionRestrictedRanges200Response1DataInner) GetRestrictedZones() []IPv4BlockAllOfRestrictedZones`

GetRestrictedZones returns the RestrictedZones field if non-nil, zero value otherwise.

### GetRestrictedZonesOk

`func (o *GetCollectionRestrictedRanges200Response1DataInner) GetRestrictedZonesOk() (*[]IPv4BlockAllOfRestrictedZones, bool)`

GetRestrictedZonesOk returns a tuple with the RestrictedZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedZones

`func (o *GetCollectionRestrictedRanges200Response1DataInner) SetRestrictedZones(v []IPv4BlockAllOfRestrictedZones)`

SetRestrictedZones sets RestrictedZones field to given value.

### HasRestrictedZones

`func (o *GetCollectionRestrictedRanges200Response1DataInner) HasRestrictedZones() bool`

HasRestrictedZones returns a boolean if a field has been set.

### GetLowWaterMark

`func (o *GetCollectionRestrictedRanges200Response1DataInner) GetLowWaterMark() int32`

GetLowWaterMark returns the LowWaterMark field if non-nil, zero value otherwise.

### GetLowWaterMarkOk

`func (o *GetCollectionRestrictedRanges200Response1DataInner) GetLowWaterMarkOk() (*int32, bool)`

GetLowWaterMarkOk returns a tuple with the LowWaterMark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowWaterMark

`func (o *GetCollectionRestrictedRanges200Response1DataInner) SetLowWaterMark(v int32)`

SetLowWaterMark sets LowWaterMark field to given value.

### HasLowWaterMark

`func (o *GetCollectionRestrictedRanges200Response1DataInner) HasLowWaterMark() bool`

HasLowWaterMark returns a boolean if a field has been set.

### GetHighWaterMark

`func (o *GetCollectionRestrictedRanges200Response1DataInner) GetHighWaterMark() int32`

GetHighWaterMark returns the HighWaterMark field if non-nil, zero value otherwise.

### GetHighWaterMarkOk

`func (o *GetCollectionRestrictedRanges200Response1DataInner) GetHighWaterMarkOk() (*int32, bool)`

GetHighWaterMarkOk returns a tuple with the HighWaterMark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighWaterMark

`func (o *GetCollectionRestrictedRanges200Response1DataInner) SetHighWaterMark(v int32)`

SetHighWaterMark sets HighWaterMark field to given value.

### HasHighWaterMark

`func (o *GetCollectionRestrictedRanges200Response1DataInner) HasHighWaterMark() bool`

HasHighWaterMark returns a boolean if a field has been set.

### GetReverseZoneSigned

`func (o *GetCollectionRestrictedRanges200Response1DataInner) GetReverseZoneSigned() bool`

GetReverseZoneSigned returns the ReverseZoneSigned field if non-nil, zero value otherwise.

### GetReverseZoneSignedOk

`func (o *GetCollectionRestrictedRanges200Response1DataInner) GetReverseZoneSignedOk() (*bool, bool)`

GetReverseZoneSignedOk returns a tuple with the ReverseZoneSigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReverseZoneSigned

`func (o *GetCollectionRestrictedRanges200Response1DataInner) SetReverseZoneSigned(v bool)`

SetReverseZoneSigned sets ReverseZoneSigned field to given value.

### HasReverseZoneSigned

`func (o *GetCollectionRestrictedRanges200Response1DataInner) HasReverseZoneSigned() bool`

HasReverseZoneSigned returns a boolean if a field has been set.

### GetReverseZoneSigningPolicy

`func (o *GetCollectionRestrictedRanges200Response1DataInner) GetReverseZoneSigningPolicy() InlinedDNSSECSigningPolicy`

GetReverseZoneSigningPolicy returns the ReverseZoneSigningPolicy field if non-nil, zero value otherwise.

### GetReverseZoneSigningPolicyOk

`func (o *GetCollectionRestrictedRanges200Response1DataInner) GetReverseZoneSigningPolicyOk() (*InlinedDNSSECSigningPolicy, bool)`

GetReverseZoneSigningPolicyOk returns a tuple with the ReverseZoneSigningPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReverseZoneSigningPolicy

`func (o *GetCollectionRestrictedRanges200Response1DataInner) SetReverseZoneSigningPolicy(v InlinedDNSSECSigningPolicy)`

SetReverseZoneSigningPolicy sets ReverseZoneSigningPolicy field to given value.

### HasReverseZoneSigningPolicy

`func (o *GetCollectionRestrictedRanges200Response1DataInner) HasReverseZoneSigningPolicy() bool`

HasReverseZoneSigningPolicy returns a boolean if a field has been set.

### GetUsagePercentage

`func (o *GetCollectionRestrictedRanges200Response1DataInner) GetUsagePercentage() map[string]int32`

GetUsagePercentage returns the UsagePercentage field if non-nil, zero value otherwise.

### GetUsagePercentageOk

`func (o *GetCollectionRestrictedRanges200Response1DataInner) GetUsagePercentageOk() (*map[string]int32, bool)`

GetUsagePercentageOk returns a tuple with the UsagePercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsagePercentage

`func (o *GetCollectionRestrictedRanges200Response1DataInner) SetUsagePercentage(v map[string]int32)`

SetUsagePercentage sets UsagePercentage field to given value.

### HasUsagePercentage

`func (o *GetCollectionRestrictedRanges200Response1DataInner) HasUsagePercentage() bool`

HasUsagePercentage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


