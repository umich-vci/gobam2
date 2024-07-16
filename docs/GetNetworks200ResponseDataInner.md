# GetNetworks200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | Pointer to **string** | The name of the resource. | [optional] 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**Configuration** | Pointer to [**InlinedConfiguration**](InlinedConfiguration.md) |  | [optional] [readonly] 
**Range** | Pointer to **string** | The IP address range for the network. The range can be provided in CIDR (e.g. 10.0.0.0/16) or address range format (e.g. 10.0.0.0-10.0.255.255). | [optional] 
**Location** | Pointer to [**InlinedLocation**](InlinedLocation.md) |  | [optional] 
**InheritedFields** | Pointer to **[]string** |  | [optional] [readonly] 
**Gateway** | Pointer to **string** | The gateway address for the IPv4 network. | [optional] 
**Template** | Pointer to [**InlinedIPv4Template**](InlinedIPv4Template.md) |  | [optional] 
**RouterPortInfo** | Pointer to **string** | Connected router port information for the IPv4 network. | [optional] [readonly] 
**DefaultView** | Pointer to [**InlinedView**](InlinedView.md) |  | [optional] 
**SharedNetworkTag** | Pointer to [**InlinedTag**](InlinedTag.md) |  | [optional] 
**DuplicateHostnamesAllowed** | Pointer to **bool** | Indicates whether duplicate hostnames are allowed in the network. If enabled, duplicate hostnames will be allowed. If not set explicitly, this value will be inherited from the parent block. | [optional] 
**PingBeforeAssignEnabled** | Pointer to **bool** | Indicates whether Address Manager will ping IP addresses before assigning them. Unavailable addresses will not be assigned. | [optional] 
**DefaultZonesInherited** | Pointer to **bool** | Indicates whether the IPv4 network will inherit the parent resource&#39;s default domains if configured. | [optional] 
**DefaultZones** | Pointer to [**[]IPv4NetworkAllOfDefaultZones**](IPv4NetworkAllOfDefaultZones.md) |  | [optional] 
**RestrictedZonesInherited** | Pointer to **bool** | Indicates whether the IPv4 network will inherit the parent resource&#39;s restricted views and zones if configured. | [optional] 
**RestrictedZones** | Pointer to [**[]IPv4NetworkAllOfRestrictedZones**](IPv4NetworkAllOfRestrictedZones.md) |  | [optional] 
**LowWaterMark** | Pointer to **int32** | A DHCP alert is triggered when usage falls below this percentage (when too few addresses are in use). | [optional] 
**HighWaterMark** | Pointer to **int32** | A DHCP alert is triggered when usage exceeds this percentage (when too many addresses are in use). | [optional] 
**ReverseZoneSigned** | Pointer to **bool** | Indicates whether the reverse zone associated with the IPv4 network will be signed according to the DNSSEC signing policy in reverseZoneSigningPolicy. | [optional] 
**ReverseZoneSigningPolicy** | Pointer to [**InlinedDNSSECSigningPolicy**](InlinedDNSSECSigningPolicy.md) |  | [optional] 
**Usage** | Pointer to **map[string]int64** | Values indicating the current amount of static, DHCP allocated, DHCP abandoned, DHCP reserved, reserved, gateway, and unassigned IP addresses within the network. | [optional] [readonly] 
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 

## Methods

### NewGetNetworks200ResponseDataInner

`func NewGetNetworks200ResponseDataInner() *GetNetworks200ResponseDataInner`

NewGetNetworks200ResponseDataInner instantiates a new GetNetworks200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworks200ResponseDataInnerWithDefaults

`func NewGetNetworks200ResponseDataInnerWithDefaults() *GetNetworks200ResponseDataInner`

NewGetNetworks200ResponseDataInnerWithDefaults instantiates a new GetNetworks200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetNetworks200ResponseDataInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetNetworks200ResponseDataInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetNetworks200ResponseDataInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetNetworks200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GetNetworks200ResponseDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetNetworks200ResponseDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetNetworks200ResponseDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetNetworks200ResponseDataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *GetNetworks200ResponseDataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetNetworks200ResponseDataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetNetworks200ResponseDataInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetNetworks200ResponseDataInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *GetNetworks200ResponseDataInner) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *GetNetworks200ResponseDataInner) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *GetNetworks200ResponseDataInner) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *GetNetworks200ResponseDataInner) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetConfiguration

`func (o *GetNetworks200ResponseDataInner) GetConfiguration() InlinedConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *GetNetworks200ResponseDataInner) GetConfigurationOk() (*InlinedConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *GetNetworks200ResponseDataInner) SetConfiguration(v InlinedConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *GetNetworks200ResponseDataInner) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetRange

`func (o *GetNetworks200ResponseDataInner) GetRange() string`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *GetNetworks200ResponseDataInner) GetRangeOk() (*string, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *GetNetworks200ResponseDataInner) SetRange(v string)`

SetRange sets Range field to given value.

### HasRange

`func (o *GetNetworks200ResponseDataInner) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetLocation

`func (o *GetNetworks200ResponseDataInner) GetLocation() InlinedLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *GetNetworks200ResponseDataInner) GetLocationOk() (*InlinedLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *GetNetworks200ResponseDataInner) SetLocation(v InlinedLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *GetNetworks200ResponseDataInner) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetInheritedFields

`func (o *GetNetworks200ResponseDataInner) GetInheritedFields() []string`

GetInheritedFields returns the InheritedFields field if non-nil, zero value otherwise.

### GetInheritedFieldsOk

`func (o *GetNetworks200ResponseDataInner) GetInheritedFieldsOk() (*[]string, bool)`

GetInheritedFieldsOk returns a tuple with the InheritedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritedFields

`func (o *GetNetworks200ResponseDataInner) SetInheritedFields(v []string)`

SetInheritedFields sets InheritedFields field to given value.

### HasInheritedFields

`func (o *GetNetworks200ResponseDataInner) HasInheritedFields() bool`

HasInheritedFields returns a boolean if a field has been set.

### GetGateway

`func (o *GetNetworks200ResponseDataInner) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *GetNetworks200ResponseDataInner) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *GetNetworks200ResponseDataInner) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *GetNetworks200ResponseDataInner) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetTemplate

`func (o *GetNetworks200ResponseDataInner) GetTemplate() InlinedIPv4Template`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *GetNetworks200ResponseDataInner) GetTemplateOk() (*InlinedIPv4Template, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *GetNetworks200ResponseDataInner) SetTemplate(v InlinedIPv4Template)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *GetNetworks200ResponseDataInner) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetRouterPortInfo

`func (o *GetNetworks200ResponseDataInner) GetRouterPortInfo() string`

GetRouterPortInfo returns the RouterPortInfo field if non-nil, zero value otherwise.

### GetRouterPortInfoOk

`func (o *GetNetworks200ResponseDataInner) GetRouterPortInfoOk() (*string, bool)`

GetRouterPortInfoOk returns a tuple with the RouterPortInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouterPortInfo

`func (o *GetNetworks200ResponseDataInner) SetRouterPortInfo(v string)`

SetRouterPortInfo sets RouterPortInfo field to given value.

### HasRouterPortInfo

`func (o *GetNetworks200ResponseDataInner) HasRouterPortInfo() bool`

HasRouterPortInfo returns a boolean if a field has been set.

### GetDefaultView

`func (o *GetNetworks200ResponseDataInner) GetDefaultView() InlinedView`

GetDefaultView returns the DefaultView field if non-nil, zero value otherwise.

### GetDefaultViewOk

`func (o *GetNetworks200ResponseDataInner) GetDefaultViewOk() (*InlinedView, bool)`

GetDefaultViewOk returns a tuple with the DefaultView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultView

`func (o *GetNetworks200ResponseDataInner) SetDefaultView(v InlinedView)`

SetDefaultView sets DefaultView field to given value.

### HasDefaultView

`func (o *GetNetworks200ResponseDataInner) HasDefaultView() bool`

HasDefaultView returns a boolean if a field has been set.

### GetSharedNetworkTag

`func (o *GetNetworks200ResponseDataInner) GetSharedNetworkTag() InlinedTag`

GetSharedNetworkTag returns the SharedNetworkTag field if non-nil, zero value otherwise.

### GetSharedNetworkTagOk

`func (o *GetNetworks200ResponseDataInner) GetSharedNetworkTagOk() (*InlinedTag, bool)`

GetSharedNetworkTagOk returns a tuple with the SharedNetworkTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedNetworkTag

`func (o *GetNetworks200ResponseDataInner) SetSharedNetworkTag(v InlinedTag)`

SetSharedNetworkTag sets SharedNetworkTag field to given value.

### HasSharedNetworkTag

`func (o *GetNetworks200ResponseDataInner) HasSharedNetworkTag() bool`

HasSharedNetworkTag returns a boolean if a field has been set.

### GetDuplicateHostnamesAllowed

`func (o *GetNetworks200ResponseDataInner) GetDuplicateHostnamesAllowed() bool`

GetDuplicateHostnamesAllowed returns the DuplicateHostnamesAllowed field if non-nil, zero value otherwise.

### GetDuplicateHostnamesAllowedOk

`func (o *GetNetworks200ResponseDataInner) GetDuplicateHostnamesAllowedOk() (*bool, bool)`

GetDuplicateHostnamesAllowedOk returns a tuple with the DuplicateHostnamesAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplicateHostnamesAllowed

`func (o *GetNetworks200ResponseDataInner) SetDuplicateHostnamesAllowed(v bool)`

SetDuplicateHostnamesAllowed sets DuplicateHostnamesAllowed field to given value.

### HasDuplicateHostnamesAllowed

`func (o *GetNetworks200ResponseDataInner) HasDuplicateHostnamesAllowed() bool`

HasDuplicateHostnamesAllowed returns a boolean if a field has been set.

### GetPingBeforeAssignEnabled

`func (o *GetNetworks200ResponseDataInner) GetPingBeforeAssignEnabled() bool`

GetPingBeforeAssignEnabled returns the PingBeforeAssignEnabled field if non-nil, zero value otherwise.

### GetPingBeforeAssignEnabledOk

`func (o *GetNetworks200ResponseDataInner) GetPingBeforeAssignEnabledOk() (*bool, bool)`

GetPingBeforeAssignEnabledOk returns a tuple with the PingBeforeAssignEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingBeforeAssignEnabled

`func (o *GetNetworks200ResponseDataInner) SetPingBeforeAssignEnabled(v bool)`

SetPingBeforeAssignEnabled sets PingBeforeAssignEnabled field to given value.

### HasPingBeforeAssignEnabled

`func (o *GetNetworks200ResponseDataInner) HasPingBeforeAssignEnabled() bool`

HasPingBeforeAssignEnabled returns a boolean if a field has been set.

### GetDefaultZonesInherited

`func (o *GetNetworks200ResponseDataInner) GetDefaultZonesInherited() bool`

GetDefaultZonesInherited returns the DefaultZonesInherited field if non-nil, zero value otherwise.

### GetDefaultZonesInheritedOk

`func (o *GetNetworks200ResponseDataInner) GetDefaultZonesInheritedOk() (*bool, bool)`

GetDefaultZonesInheritedOk returns a tuple with the DefaultZonesInherited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultZonesInherited

`func (o *GetNetworks200ResponseDataInner) SetDefaultZonesInherited(v bool)`

SetDefaultZonesInherited sets DefaultZonesInherited field to given value.

### HasDefaultZonesInherited

`func (o *GetNetworks200ResponseDataInner) HasDefaultZonesInherited() bool`

HasDefaultZonesInherited returns a boolean if a field has been set.

### GetDefaultZones

`func (o *GetNetworks200ResponseDataInner) GetDefaultZones() []IPv4NetworkAllOfDefaultZones`

GetDefaultZones returns the DefaultZones field if non-nil, zero value otherwise.

### GetDefaultZonesOk

`func (o *GetNetworks200ResponseDataInner) GetDefaultZonesOk() (*[]IPv4NetworkAllOfDefaultZones, bool)`

GetDefaultZonesOk returns a tuple with the DefaultZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultZones

`func (o *GetNetworks200ResponseDataInner) SetDefaultZones(v []IPv4NetworkAllOfDefaultZones)`

SetDefaultZones sets DefaultZones field to given value.

### HasDefaultZones

`func (o *GetNetworks200ResponseDataInner) HasDefaultZones() bool`

HasDefaultZones returns a boolean if a field has been set.

### GetRestrictedZonesInherited

`func (o *GetNetworks200ResponseDataInner) GetRestrictedZonesInherited() bool`

GetRestrictedZonesInherited returns the RestrictedZonesInherited field if non-nil, zero value otherwise.

### GetRestrictedZonesInheritedOk

`func (o *GetNetworks200ResponseDataInner) GetRestrictedZonesInheritedOk() (*bool, bool)`

GetRestrictedZonesInheritedOk returns a tuple with the RestrictedZonesInherited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedZonesInherited

`func (o *GetNetworks200ResponseDataInner) SetRestrictedZonesInherited(v bool)`

SetRestrictedZonesInherited sets RestrictedZonesInherited field to given value.

### HasRestrictedZonesInherited

`func (o *GetNetworks200ResponseDataInner) HasRestrictedZonesInherited() bool`

HasRestrictedZonesInherited returns a boolean if a field has been set.

### GetRestrictedZones

`func (o *GetNetworks200ResponseDataInner) GetRestrictedZones() []IPv4NetworkAllOfRestrictedZones`

GetRestrictedZones returns the RestrictedZones field if non-nil, zero value otherwise.

### GetRestrictedZonesOk

`func (o *GetNetworks200ResponseDataInner) GetRestrictedZonesOk() (*[]IPv4NetworkAllOfRestrictedZones, bool)`

GetRestrictedZonesOk returns a tuple with the RestrictedZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedZones

`func (o *GetNetworks200ResponseDataInner) SetRestrictedZones(v []IPv4NetworkAllOfRestrictedZones)`

SetRestrictedZones sets RestrictedZones field to given value.

### HasRestrictedZones

`func (o *GetNetworks200ResponseDataInner) HasRestrictedZones() bool`

HasRestrictedZones returns a boolean if a field has been set.

### GetLowWaterMark

`func (o *GetNetworks200ResponseDataInner) GetLowWaterMark() int32`

GetLowWaterMark returns the LowWaterMark field if non-nil, zero value otherwise.

### GetLowWaterMarkOk

`func (o *GetNetworks200ResponseDataInner) GetLowWaterMarkOk() (*int32, bool)`

GetLowWaterMarkOk returns a tuple with the LowWaterMark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowWaterMark

`func (o *GetNetworks200ResponseDataInner) SetLowWaterMark(v int32)`

SetLowWaterMark sets LowWaterMark field to given value.

### HasLowWaterMark

`func (o *GetNetworks200ResponseDataInner) HasLowWaterMark() bool`

HasLowWaterMark returns a boolean if a field has been set.

### GetHighWaterMark

`func (o *GetNetworks200ResponseDataInner) GetHighWaterMark() int32`

GetHighWaterMark returns the HighWaterMark field if non-nil, zero value otherwise.

### GetHighWaterMarkOk

`func (o *GetNetworks200ResponseDataInner) GetHighWaterMarkOk() (*int32, bool)`

GetHighWaterMarkOk returns a tuple with the HighWaterMark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighWaterMark

`func (o *GetNetworks200ResponseDataInner) SetHighWaterMark(v int32)`

SetHighWaterMark sets HighWaterMark field to given value.

### HasHighWaterMark

`func (o *GetNetworks200ResponseDataInner) HasHighWaterMark() bool`

HasHighWaterMark returns a boolean if a field has been set.

### GetReverseZoneSigned

`func (o *GetNetworks200ResponseDataInner) GetReverseZoneSigned() bool`

GetReverseZoneSigned returns the ReverseZoneSigned field if non-nil, zero value otherwise.

### GetReverseZoneSignedOk

`func (o *GetNetworks200ResponseDataInner) GetReverseZoneSignedOk() (*bool, bool)`

GetReverseZoneSignedOk returns a tuple with the ReverseZoneSigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReverseZoneSigned

`func (o *GetNetworks200ResponseDataInner) SetReverseZoneSigned(v bool)`

SetReverseZoneSigned sets ReverseZoneSigned field to given value.

### HasReverseZoneSigned

`func (o *GetNetworks200ResponseDataInner) HasReverseZoneSigned() bool`

HasReverseZoneSigned returns a boolean if a field has been set.

### GetReverseZoneSigningPolicy

`func (o *GetNetworks200ResponseDataInner) GetReverseZoneSigningPolicy() InlinedDNSSECSigningPolicy`

GetReverseZoneSigningPolicy returns the ReverseZoneSigningPolicy field if non-nil, zero value otherwise.

### GetReverseZoneSigningPolicyOk

`func (o *GetNetworks200ResponseDataInner) GetReverseZoneSigningPolicyOk() (*InlinedDNSSECSigningPolicy, bool)`

GetReverseZoneSigningPolicyOk returns a tuple with the ReverseZoneSigningPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReverseZoneSigningPolicy

`func (o *GetNetworks200ResponseDataInner) SetReverseZoneSigningPolicy(v InlinedDNSSECSigningPolicy)`

SetReverseZoneSigningPolicy sets ReverseZoneSigningPolicy field to given value.

### HasReverseZoneSigningPolicy

`func (o *GetNetworks200ResponseDataInner) HasReverseZoneSigningPolicy() bool`

HasReverseZoneSigningPolicy returns a boolean if a field has been set.

### GetUsage

`func (o *GetNetworks200ResponseDataInner) GetUsage() map[string]int64`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *GetNetworks200ResponseDataInner) GetUsageOk() (*map[string]int64, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *GetNetworks200ResponseDataInner) SetUsage(v map[string]int64)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *GetNetworks200ResponseDataInner) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetLinks

`func (o *GetNetworks200ResponseDataInner) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetNetworks200ResponseDataInner) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetNetworks200ResponseDataInner) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetNetworks200ResponseDataInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


