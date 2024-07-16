# GetAddresses200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The type of IP address. | [optional] 
**Name** | Pointer to **string** | The name of the resource. | [optional] 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**Configuration** | Pointer to [**InlinedConfiguration**](InlinedConfiguration.md) |  | [optional] [readonly] 
**State** | Pointer to **string** | The state of the IP address. | [optional] 
**Address** | Pointer to **string** |  | [optional] 
**MacAddress** | Pointer to [**AddressMacAddress**](AddressMacAddress.md) |  | [optional] 
**Device** | Pointer to [**InlinedDevice**](InlinedDevice.md) |  | [optional] 
**Location** | Pointer to [**InlinedLocation**](InlinedLocation.md) |  | [optional] 
**ResourceRecords** | Pointer to [**[]ResourceRecord**](ResourceRecord.md) |  | [optional] 
**LeaseDateTime** | Pointer to **time.Time** | The timestamp detailing when the lease for the IP address was offered. | [optional] [readonly] 
**LeaseExpirationDateTime** | Pointer to **time.Time** | The timestamp detailing when the lease for the IP address expires. | [optional] [readonly] 
**RemoteId** | Pointer to **string** | The ID of the relay agent. | [optional] [readonly] 
**IpGroup** | Pointer to [**InlinedIPv4Group**](InlinedIPv4Group.md) |  | [optional] [readonly] 
**Template** | Pointer to [**InlinedIPv4Template**](InlinedIPv4Template.md) |  | [optional] 
**CircuitId** | Pointer to **string** | The ID of the circuit from which the request came. | [optional] [readonly] 
**RouterPortInfo** | Pointer to **string** | Connected router port information for the IPv4 address. | [optional] [readonly] 
**SwitchPortInfo** | Pointer to **string** | Connected switch port information for the IPv4 address. | [optional] [readonly] 
**VlanInfo** | Pointer to **string** | VLAN information for the IPv4 address. | [optional] [readonly] 
**VendorClassIdentifier** | Pointer to **string** | An identifier sent by the DHCP client software running on a device. | [optional] [readonly] 
**ParameterRequestList** | Pointer to **[]int32** | The list of parameters that the device requested from the DHCP server. | [optional] [readonly] 
**ClientIdentifier** | Pointer to [**InlinedDHCPUniqueIdentifier**](InlinedDHCPUniqueIdentifier.md) |  | [optional] 
**ReservedUsing** | Pointer to **string** | The type of DHCP reservation. | [optional] 
**IdentityAssociationIdentifier** | Pointer to **string** | The Identity Association Identifier for the network interface of the client system. | [optional] [readonly] 
**InterfaceId** | Pointer to **string** | The ID of the interface from which the request came. | [optional] [readonly] 
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 

## Methods

### NewGetAddresses200ResponseDataInner

`func NewGetAddresses200ResponseDataInner() *GetAddresses200ResponseDataInner`

NewGetAddresses200ResponseDataInner instantiates a new GetAddresses200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAddresses200ResponseDataInnerWithDefaults

`func NewGetAddresses200ResponseDataInnerWithDefaults() *GetAddresses200ResponseDataInner`

NewGetAddresses200ResponseDataInnerWithDefaults instantiates a new GetAddresses200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetAddresses200ResponseDataInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetAddresses200ResponseDataInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetAddresses200ResponseDataInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetAddresses200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GetAddresses200ResponseDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetAddresses200ResponseDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetAddresses200ResponseDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetAddresses200ResponseDataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *GetAddresses200ResponseDataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetAddresses200ResponseDataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetAddresses200ResponseDataInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetAddresses200ResponseDataInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *GetAddresses200ResponseDataInner) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *GetAddresses200ResponseDataInner) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *GetAddresses200ResponseDataInner) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *GetAddresses200ResponseDataInner) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetConfiguration

`func (o *GetAddresses200ResponseDataInner) GetConfiguration() InlinedConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *GetAddresses200ResponseDataInner) GetConfigurationOk() (*InlinedConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *GetAddresses200ResponseDataInner) SetConfiguration(v InlinedConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *GetAddresses200ResponseDataInner) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetState

`func (o *GetAddresses200ResponseDataInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetAddresses200ResponseDataInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetAddresses200ResponseDataInner) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GetAddresses200ResponseDataInner) HasState() bool`

HasState returns a boolean if a field has been set.

### GetAddress

`func (o *GetAddresses200ResponseDataInner) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *GetAddresses200ResponseDataInner) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *GetAddresses200ResponseDataInner) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *GetAddresses200ResponseDataInner) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetMacAddress

`func (o *GetAddresses200ResponseDataInner) GetMacAddress() AddressMacAddress`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *GetAddresses200ResponseDataInner) GetMacAddressOk() (*AddressMacAddress, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *GetAddresses200ResponseDataInner) SetMacAddress(v AddressMacAddress)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *GetAddresses200ResponseDataInner) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetDevice

`func (o *GetAddresses200ResponseDataInner) GetDevice() InlinedDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *GetAddresses200ResponseDataInner) GetDeviceOk() (*InlinedDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *GetAddresses200ResponseDataInner) SetDevice(v InlinedDevice)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *GetAddresses200ResponseDataInner) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetLocation

`func (o *GetAddresses200ResponseDataInner) GetLocation() InlinedLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *GetAddresses200ResponseDataInner) GetLocationOk() (*InlinedLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *GetAddresses200ResponseDataInner) SetLocation(v InlinedLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *GetAddresses200ResponseDataInner) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetResourceRecords

`func (o *GetAddresses200ResponseDataInner) GetResourceRecords() []ResourceRecord`

GetResourceRecords returns the ResourceRecords field if non-nil, zero value otherwise.

### GetResourceRecordsOk

`func (o *GetAddresses200ResponseDataInner) GetResourceRecordsOk() (*[]ResourceRecord, bool)`

GetResourceRecordsOk returns a tuple with the ResourceRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceRecords

`func (o *GetAddresses200ResponseDataInner) SetResourceRecords(v []ResourceRecord)`

SetResourceRecords sets ResourceRecords field to given value.

### HasResourceRecords

`func (o *GetAddresses200ResponseDataInner) HasResourceRecords() bool`

HasResourceRecords returns a boolean if a field has been set.

### GetLeaseDateTime

`func (o *GetAddresses200ResponseDataInner) GetLeaseDateTime() time.Time`

GetLeaseDateTime returns the LeaseDateTime field if non-nil, zero value otherwise.

### GetLeaseDateTimeOk

`func (o *GetAddresses200ResponseDataInner) GetLeaseDateTimeOk() (*time.Time, bool)`

GetLeaseDateTimeOk returns a tuple with the LeaseDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseDateTime

`func (o *GetAddresses200ResponseDataInner) SetLeaseDateTime(v time.Time)`

SetLeaseDateTime sets LeaseDateTime field to given value.

### HasLeaseDateTime

`func (o *GetAddresses200ResponseDataInner) HasLeaseDateTime() bool`

HasLeaseDateTime returns a boolean if a field has been set.

### GetLeaseExpirationDateTime

`func (o *GetAddresses200ResponseDataInner) GetLeaseExpirationDateTime() time.Time`

GetLeaseExpirationDateTime returns the LeaseExpirationDateTime field if non-nil, zero value otherwise.

### GetLeaseExpirationDateTimeOk

`func (o *GetAddresses200ResponseDataInner) GetLeaseExpirationDateTimeOk() (*time.Time, bool)`

GetLeaseExpirationDateTimeOk returns a tuple with the LeaseExpirationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseExpirationDateTime

`func (o *GetAddresses200ResponseDataInner) SetLeaseExpirationDateTime(v time.Time)`

SetLeaseExpirationDateTime sets LeaseExpirationDateTime field to given value.

### HasLeaseExpirationDateTime

`func (o *GetAddresses200ResponseDataInner) HasLeaseExpirationDateTime() bool`

HasLeaseExpirationDateTime returns a boolean if a field has been set.

### GetRemoteId

`func (o *GetAddresses200ResponseDataInner) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *GetAddresses200ResponseDataInner) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *GetAddresses200ResponseDataInner) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.

### HasRemoteId

`func (o *GetAddresses200ResponseDataInner) HasRemoteId() bool`

HasRemoteId returns a boolean if a field has been set.

### GetIpGroup

`func (o *GetAddresses200ResponseDataInner) GetIpGroup() InlinedIPv4Group`

GetIpGroup returns the IpGroup field if non-nil, zero value otherwise.

### GetIpGroupOk

`func (o *GetAddresses200ResponseDataInner) GetIpGroupOk() (*InlinedIPv4Group, bool)`

GetIpGroupOk returns a tuple with the IpGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpGroup

`func (o *GetAddresses200ResponseDataInner) SetIpGroup(v InlinedIPv4Group)`

SetIpGroup sets IpGroup field to given value.

### HasIpGroup

`func (o *GetAddresses200ResponseDataInner) HasIpGroup() bool`

HasIpGroup returns a boolean if a field has been set.

### GetTemplate

`func (o *GetAddresses200ResponseDataInner) GetTemplate() InlinedIPv4Template`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *GetAddresses200ResponseDataInner) GetTemplateOk() (*InlinedIPv4Template, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *GetAddresses200ResponseDataInner) SetTemplate(v InlinedIPv4Template)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *GetAddresses200ResponseDataInner) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetCircuitId

`func (o *GetAddresses200ResponseDataInner) GetCircuitId() string`

GetCircuitId returns the CircuitId field if non-nil, zero value otherwise.

### GetCircuitIdOk

`func (o *GetAddresses200ResponseDataInner) GetCircuitIdOk() (*string, bool)`

GetCircuitIdOk returns a tuple with the CircuitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuitId

`func (o *GetAddresses200ResponseDataInner) SetCircuitId(v string)`

SetCircuitId sets CircuitId field to given value.

### HasCircuitId

`func (o *GetAddresses200ResponseDataInner) HasCircuitId() bool`

HasCircuitId returns a boolean if a field has been set.

### GetRouterPortInfo

`func (o *GetAddresses200ResponseDataInner) GetRouterPortInfo() string`

GetRouterPortInfo returns the RouterPortInfo field if non-nil, zero value otherwise.

### GetRouterPortInfoOk

`func (o *GetAddresses200ResponseDataInner) GetRouterPortInfoOk() (*string, bool)`

GetRouterPortInfoOk returns a tuple with the RouterPortInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouterPortInfo

`func (o *GetAddresses200ResponseDataInner) SetRouterPortInfo(v string)`

SetRouterPortInfo sets RouterPortInfo field to given value.

### HasRouterPortInfo

`func (o *GetAddresses200ResponseDataInner) HasRouterPortInfo() bool`

HasRouterPortInfo returns a boolean if a field has been set.

### GetSwitchPortInfo

`func (o *GetAddresses200ResponseDataInner) GetSwitchPortInfo() string`

GetSwitchPortInfo returns the SwitchPortInfo field if non-nil, zero value otherwise.

### GetSwitchPortInfoOk

`func (o *GetAddresses200ResponseDataInner) GetSwitchPortInfoOk() (*string, bool)`

GetSwitchPortInfoOk returns a tuple with the SwitchPortInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchPortInfo

`func (o *GetAddresses200ResponseDataInner) SetSwitchPortInfo(v string)`

SetSwitchPortInfo sets SwitchPortInfo field to given value.

### HasSwitchPortInfo

`func (o *GetAddresses200ResponseDataInner) HasSwitchPortInfo() bool`

HasSwitchPortInfo returns a boolean if a field has been set.

### GetVlanInfo

`func (o *GetAddresses200ResponseDataInner) GetVlanInfo() string`

GetVlanInfo returns the VlanInfo field if non-nil, zero value otherwise.

### GetVlanInfoOk

`func (o *GetAddresses200ResponseDataInner) GetVlanInfoOk() (*string, bool)`

GetVlanInfoOk returns a tuple with the VlanInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanInfo

`func (o *GetAddresses200ResponseDataInner) SetVlanInfo(v string)`

SetVlanInfo sets VlanInfo field to given value.

### HasVlanInfo

`func (o *GetAddresses200ResponseDataInner) HasVlanInfo() bool`

HasVlanInfo returns a boolean if a field has been set.

### GetVendorClassIdentifier

`func (o *GetAddresses200ResponseDataInner) GetVendorClassIdentifier() string`

GetVendorClassIdentifier returns the VendorClassIdentifier field if non-nil, zero value otherwise.

### GetVendorClassIdentifierOk

`func (o *GetAddresses200ResponseDataInner) GetVendorClassIdentifierOk() (*string, bool)`

GetVendorClassIdentifierOk returns a tuple with the VendorClassIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorClassIdentifier

`func (o *GetAddresses200ResponseDataInner) SetVendorClassIdentifier(v string)`

SetVendorClassIdentifier sets VendorClassIdentifier field to given value.

### HasVendorClassIdentifier

`func (o *GetAddresses200ResponseDataInner) HasVendorClassIdentifier() bool`

HasVendorClassIdentifier returns a boolean if a field has been set.

### GetParameterRequestList

`func (o *GetAddresses200ResponseDataInner) GetParameterRequestList() []int32`

GetParameterRequestList returns the ParameterRequestList field if non-nil, zero value otherwise.

### GetParameterRequestListOk

`func (o *GetAddresses200ResponseDataInner) GetParameterRequestListOk() (*[]int32, bool)`

GetParameterRequestListOk returns a tuple with the ParameterRequestList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterRequestList

`func (o *GetAddresses200ResponseDataInner) SetParameterRequestList(v []int32)`

SetParameterRequestList sets ParameterRequestList field to given value.

### HasParameterRequestList

`func (o *GetAddresses200ResponseDataInner) HasParameterRequestList() bool`

HasParameterRequestList returns a boolean if a field has been set.

### GetClientIdentifier

`func (o *GetAddresses200ResponseDataInner) GetClientIdentifier() InlinedDHCPUniqueIdentifier`

GetClientIdentifier returns the ClientIdentifier field if non-nil, zero value otherwise.

### GetClientIdentifierOk

`func (o *GetAddresses200ResponseDataInner) GetClientIdentifierOk() (*InlinedDHCPUniqueIdentifier, bool)`

GetClientIdentifierOk returns a tuple with the ClientIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIdentifier

`func (o *GetAddresses200ResponseDataInner) SetClientIdentifier(v InlinedDHCPUniqueIdentifier)`

SetClientIdentifier sets ClientIdentifier field to given value.

### HasClientIdentifier

`func (o *GetAddresses200ResponseDataInner) HasClientIdentifier() bool`

HasClientIdentifier returns a boolean if a field has been set.

### GetReservedUsing

`func (o *GetAddresses200ResponseDataInner) GetReservedUsing() string`

GetReservedUsing returns the ReservedUsing field if non-nil, zero value otherwise.

### GetReservedUsingOk

`func (o *GetAddresses200ResponseDataInner) GetReservedUsingOk() (*string, bool)`

GetReservedUsingOk returns a tuple with the ReservedUsing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedUsing

`func (o *GetAddresses200ResponseDataInner) SetReservedUsing(v string)`

SetReservedUsing sets ReservedUsing field to given value.

### HasReservedUsing

`func (o *GetAddresses200ResponseDataInner) HasReservedUsing() bool`

HasReservedUsing returns a boolean if a field has been set.

### GetIdentityAssociationIdentifier

`func (o *GetAddresses200ResponseDataInner) GetIdentityAssociationIdentifier() string`

GetIdentityAssociationIdentifier returns the IdentityAssociationIdentifier field if non-nil, zero value otherwise.

### GetIdentityAssociationIdentifierOk

`func (o *GetAddresses200ResponseDataInner) GetIdentityAssociationIdentifierOk() (*string, bool)`

GetIdentityAssociationIdentifierOk returns a tuple with the IdentityAssociationIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityAssociationIdentifier

`func (o *GetAddresses200ResponseDataInner) SetIdentityAssociationIdentifier(v string)`

SetIdentityAssociationIdentifier sets IdentityAssociationIdentifier field to given value.

### HasIdentityAssociationIdentifier

`func (o *GetAddresses200ResponseDataInner) HasIdentityAssociationIdentifier() bool`

HasIdentityAssociationIdentifier returns a boolean if a field has been set.

### GetInterfaceId

`func (o *GetAddresses200ResponseDataInner) GetInterfaceId() string`

GetInterfaceId returns the InterfaceId field if non-nil, zero value otherwise.

### GetInterfaceIdOk

`func (o *GetAddresses200ResponseDataInner) GetInterfaceIdOk() (*string, bool)`

GetInterfaceIdOk returns a tuple with the InterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceId

`func (o *GetAddresses200ResponseDataInner) SetInterfaceId(v string)`

SetInterfaceId sets InterfaceId field to given value.

### HasInterfaceId

`func (o *GetAddresses200ResponseDataInner) HasInterfaceId() bool`

HasInterfaceId returns a boolean if a field has been set.

### GetLinks

`func (o *GetAddresses200ResponseDataInner) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetAddresses200ResponseDataInner) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetAddresses200ResponseDataInner) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetAddresses200ResponseDataInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


