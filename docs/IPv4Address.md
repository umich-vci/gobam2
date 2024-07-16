# IPv4Address

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of IP address. | [optional] 
**State** | Pointer to **string** | The state of the IP address. | [optional] 
**IpGroup** | Pointer to [**InlinedIPv4Group**](InlinedIPv4Group.md) |  | [optional] [readonly] 
**Template** | Pointer to [**InlinedIPv4Template**](InlinedIPv4Template.md) |  | [optional] 
**CircuitId** | Pointer to **string** | The ID of the circuit from which the request came. | [optional] [readonly] 
**RouterPortInfo** | Pointer to **string** | Connected router port information for the IPv4 address. | [optional] [readonly] 
**SwitchPortInfo** | Pointer to **string** | Connected switch port information for the IPv4 address. | [optional] [readonly] 
**VlanInfo** | Pointer to **string** | VLAN information for the IPv4 address. | [optional] [readonly] 
**VendorClassIdentifier** | Pointer to **string** | An identifier sent by the DHCP client software running on a device. | [optional] [readonly] 
**ParameterRequestList** | Pointer to **[]int32** | The list of parameters that the device requested from the DHCP server. | [optional] [readonly] 
**ClientIdentifier** | Pointer to [**InlinedDHCPClientIdentifier**](InlinedDHCPClientIdentifier.md) |  | [optional] 

## Methods

### NewIPv4Address

`func NewIPv4Address() *IPv4Address`

NewIPv4Address instantiates a new IPv4Address object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPv4AddressWithDefaults

`func NewIPv4AddressWithDefaults() *IPv4Address`

NewIPv4AddressWithDefaults instantiates a new IPv4Address object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *IPv4Address) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IPv4Address) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IPv4Address) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IPv4Address) HasType() bool`

HasType returns a boolean if a field has been set.

### GetState

`func (o *IPv4Address) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *IPv4Address) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *IPv4Address) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *IPv4Address) HasState() bool`

HasState returns a boolean if a field has been set.

### GetIpGroup

`func (o *IPv4Address) GetIpGroup() InlinedIPv4Group`

GetIpGroup returns the IpGroup field if non-nil, zero value otherwise.

### GetIpGroupOk

`func (o *IPv4Address) GetIpGroupOk() (*InlinedIPv4Group, bool)`

GetIpGroupOk returns a tuple with the IpGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpGroup

`func (o *IPv4Address) SetIpGroup(v InlinedIPv4Group)`

SetIpGroup sets IpGroup field to given value.

### HasIpGroup

`func (o *IPv4Address) HasIpGroup() bool`

HasIpGroup returns a boolean if a field has been set.

### GetTemplate

`func (o *IPv4Address) GetTemplate() InlinedIPv4Template`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *IPv4Address) GetTemplateOk() (*InlinedIPv4Template, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *IPv4Address) SetTemplate(v InlinedIPv4Template)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *IPv4Address) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetCircuitId

`func (o *IPv4Address) GetCircuitId() string`

GetCircuitId returns the CircuitId field if non-nil, zero value otherwise.

### GetCircuitIdOk

`func (o *IPv4Address) GetCircuitIdOk() (*string, bool)`

GetCircuitIdOk returns a tuple with the CircuitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuitId

`func (o *IPv4Address) SetCircuitId(v string)`

SetCircuitId sets CircuitId field to given value.

### HasCircuitId

`func (o *IPv4Address) HasCircuitId() bool`

HasCircuitId returns a boolean if a field has been set.

### GetRouterPortInfo

`func (o *IPv4Address) GetRouterPortInfo() string`

GetRouterPortInfo returns the RouterPortInfo field if non-nil, zero value otherwise.

### GetRouterPortInfoOk

`func (o *IPv4Address) GetRouterPortInfoOk() (*string, bool)`

GetRouterPortInfoOk returns a tuple with the RouterPortInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouterPortInfo

`func (o *IPv4Address) SetRouterPortInfo(v string)`

SetRouterPortInfo sets RouterPortInfo field to given value.

### HasRouterPortInfo

`func (o *IPv4Address) HasRouterPortInfo() bool`

HasRouterPortInfo returns a boolean if a field has been set.

### GetSwitchPortInfo

`func (o *IPv4Address) GetSwitchPortInfo() string`

GetSwitchPortInfo returns the SwitchPortInfo field if non-nil, zero value otherwise.

### GetSwitchPortInfoOk

`func (o *IPv4Address) GetSwitchPortInfoOk() (*string, bool)`

GetSwitchPortInfoOk returns a tuple with the SwitchPortInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchPortInfo

`func (o *IPv4Address) SetSwitchPortInfo(v string)`

SetSwitchPortInfo sets SwitchPortInfo field to given value.

### HasSwitchPortInfo

`func (o *IPv4Address) HasSwitchPortInfo() bool`

HasSwitchPortInfo returns a boolean if a field has been set.

### GetVlanInfo

`func (o *IPv4Address) GetVlanInfo() string`

GetVlanInfo returns the VlanInfo field if non-nil, zero value otherwise.

### GetVlanInfoOk

`func (o *IPv4Address) GetVlanInfoOk() (*string, bool)`

GetVlanInfoOk returns a tuple with the VlanInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanInfo

`func (o *IPv4Address) SetVlanInfo(v string)`

SetVlanInfo sets VlanInfo field to given value.

### HasVlanInfo

`func (o *IPv4Address) HasVlanInfo() bool`

HasVlanInfo returns a boolean if a field has been set.

### GetVendorClassIdentifier

`func (o *IPv4Address) GetVendorClassIdentifier() string`

GetVendorClassIdentifier returns the VendorClassIdentifier field if non-nil, zero value otherwise.

### GetVendorClassIdentifierOk

`func (o *IPv4Address) GetVendorClassIdentifierOk() (*string, bool)`

GetVendorClassIdentifierOk returns a tuple with the VendorClassIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorClassIdentifier

`func (o *IPv4Address) SetVendorClassIdentifier(v string)`

SetVendorClassIdentifier sets VendorClassIdentifier field to given value.

### HasVendorClassIdentifier

`func (o *IPv4Address) HasVendorClassIdentifier() bool`

HasVendorClassIdentifier returns a boolean if a field has been set.

### GetParameterRequestList

`func (o *IPv4Address) GetParameterRequestList() []int32`

GetParameterRequestList returns the ParameterRequestList field if non-nil, zero value otherwise.

### GetParameterRequestListOk

`func (o *IPv4Address) GetParameterRequestListOk() (*[]int32, bool)`

GetParameterRequestListOk returns a tuple with the ParameterRequestList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterRequestList

`func (o *IPv4Address) SetParameterRequestList(v []int32)`

SetParameterRequestList sets ParameterRequestList field to given value.

### HasParameterRequestList

`func (o *IPv4Address) HasParameterRequestList() bool`

HasParameterRequestList returns a boolean if a field has been set.

### GetClientIdentifier

`func (o *IPv4Address) GetClientIdentifier() InlinedDHCPClientIdentifier`

GetClientIdentifier returns the ClientIdentifier field if non-nil, zero value otherwise.

### GetClientIdentifierOk

`func (o *IPv4Address) GetClientIdentifierOk() (*InlinedDHCPClientIdentifier, bool)`

GetClientIdentifierOk returns a tuple with the ClientIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIdentifier

`func (o *IPv4Address) SetClientIdentifier(v InlinedDHCPClientIdentifier)`

SetClientIdentifier sets ClientIdentifier field to given value.

### HasClientIdentifier

`func (o *IPv4Address) HasClientIdentifier() bool`

HasClientIdentifier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


