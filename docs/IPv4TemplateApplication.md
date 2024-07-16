# IPv4TemplateApplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The resource type. | [optional] 
**DeploymentOptionsApplied** | Pointer to **bool** | Indicates whether deployment options set in the template are applied. | [optional] 
**GatewayItemApplied** | Pointer to **bool** | Indicates whether the template settings for the gateway address are applied. | [optional] 
**ReservedRangeItemsApplied** | Pointer to **bool** | Indicates whether the reserved address ranges set in the template are applied. | [optional] 
**DhcpRangeItemsApplied** | Pointer to **bool** | Indicates whether the DHCP reserved address ranges set in the template are applied. | [optional] 
**IpGroupItemsApplied** | Pointer to **bool** | Indicates whether the IP groups set in the template are applied. | [optional] 
**OrphanedAddressConversionStrategy** | Pointer to **string** | Determines whether orphaned DHCP allocated IPv4 addresses are converted into another assignment type and the type of host record assigned to the orphaned record. | [optional] 

## Methods

### NewIPv4TemplateApplication

`func NewIPv4TemplateApplication() *IPv4TemplateApplication`

NewIPv4TemplateApplication instantiates a new IPv4TemplateApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPv4TemplateApplicationWithDefaults

`func NewIPv4TemplateApplicationWithDefaults() *IPv4TemplateApplication`

NewIPv4TemplateApplicationWithDefaults instantiates a new IPv4TemplateApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *IPv4TemplateApplication) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IPv4TemplateApplication) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IPv4TemplateApplication) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IPv4TemplateApplication) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDeploymentOptionsApplied

`func (o *IPv4TemplateApplication) GetDeploymentOptionsApplied() bool`

GetDeploymentOptionsApplied returns the DeploymentOptionsApplied field if non-nil, zero value otherwise.

### GetDeploymentOptionsAppliedOk

`func (o *IPv4TemplateApplication) GetDeploymentOptionsAppliedOk() (*bool, bool)`

GetDeploymentOptionsAppliedOk returns a tuple with the DeploymentOptionsApplied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentOptionsApplied

`func (o *IPv4TemplateApplication) SetDeploymentOptionsApplied(v bool)`

SetDeploymentOptionsApplied sets DeploymentOptionsApplied field to given value.

### HasDeploymentOptionsApplied

`func (o *IPv4TemplateApplication) HasDeploymentOptionsApplied() bool`

HasDeploymentOptionsApplied returns a boolean if a field has been set.

### GetGatewayItemApplied

`func (o *IPv4TemplateApplication) GetGatewayItemApplied() bool`

GetGatewayItemApplied returns the GatewayItemApplied field if non-nil, zero value otherwise.

### GetGatewayItemAppliedOk

`func (o *IPv4TemplateApplication) GetGatewayItemAppliedOk() (*bool, bool)`

GetGatewayItemAppliedOk returns a tuple with the GatewayItemApplied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayItemApplied

`func (o *IPv4TemplateApplication) SetGatewayItemApplied(v bool)`

SetGatewayItemApplied sets GatewayItemApplied field to given value.

### HasGatewayItemApplied

`func (o *IPv4TemplateApplication) HasGatewayItemApplied() bool`

HasGatewayItemApplied returns a boolean if a field has been set.

### GetReservedRangeItemsApplied

`func (o *IPv4TemplateApplication) GetReservedRangeItemsApplied() bool`

GetReservedRangeItemsApplied returns the ReservedRangeItemsApplied field if non-nil, zero value otherwise.

### GetReservedRangeItemsAppliedOk

`func (o *IPv4TemplateApplication) GetReservedRangeItemsAppliedOk() (*bool, bool)`

GetReservedRangeItemsAppliedOk returns a tuple with the ReservedRangeItemsApplied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedRangeItemsApplied

`func (o *IPv4TemplateApplication) SetReservedRangeItemsApplied(v bool)`

SetReservedRangeItemsApplied sets ReservedRangeItemsApplied field to given value.

### HasReservedRangeItemsApplied

`func (o *IPv4TemplateApplication) HasReservedRangeItemsApplied() bool`

HasReservedRangeItemsApplied returns a boolean if a field has been set.

### GetDhcpRangeItemsApplied

`func (o *IPv4TemplateApplication) GetDhcpRangeItemsApplied() bool`

GetDhcpRangeItemsApplied returns the DhcpRangeItemsApplied field if non-nil, zero value otherwise.

### GetDhcpRangeItemsAppliedOk

`func (o *IPv4TemplateApplication) GetDhcpRangeItemsAppliedOk() (*bool, bool)`

GetDhcpRangeItemsAppliedOk returns a tuple with the DhcpRangeItemsApplied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpRangeItemsApplied

`func (o *IPv4TemplateApplication) SetDhcpRangeItemsApplied(v bool)`

SetDhcpRangeItemsApplied sets DhcpRangeItemsApplied field to given value.

### HasDhcpRangeItemsApplied

`func (o *IPv4TemplateApplication) HasDhcpRangeItemsApplied() bool`

HasDhcpRangeItemsApplied returns a boolean if a field has been set.

### GetIpGroupItemsApplied

`func (o *IPv4TemplateApplication) GetIpGroupItemsApplied() bool`

GetIpGroupItemsApplied returns the IpGroupItemsApplied field if non-nil, zero value otherwise.

### GetIpGroupItemsAppliedOk

`func (o *IPv4TemplateApplication) GetIpGroupItemsAppliedOk() (*bool, bool)`

GetIpGroupItemsAppliedOk returns a tuple with the IpGroupItemsApplied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpGroupItemsApplied

`func (o *IPv4TemplateApplication) SetIpGroupItemsApplied(v bool)`

SetIpGroupItemsApplied sets IpGroupItemsApplied field to given value.

### HasIpGroupItemsApplied

`func (o *IPv4TemplateApplication) HasIpGroupItemsApplied() bool`

HasIpGroupItemsApplied returns a boolean if a field has been set.

### GetOrphanedAddressConversionStrategy

`func (o *IPv4TemplateApplication) GetOrphanedAddressConversionStrategy() string`

GetOrphanedAddressConversionStrategy returns the OrphanedAddressConversionStrategy field if non-nil, zero value otherwise.

### GetOrphanedAddressConversionStrategyOk

`func (o *IPv4TemplateApplication) GetOrphanedAddressConversionStrategyOk() (*string, bool)`

GetOrphanedAddressConversionStrategyOk returns a tuple with the OrphanedAddressConversionStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrphanedAddressConversionStrategy

`func (o *IPv4TemplateApplication) SetOrphanedAddressConversionStrategy(v string)`

SetOrphanedAddressConversionStrategy sets OrphanedAddressConversionStrategy field to given value.

### HasOrphanedAddressConversionStrategy

`func (o *IPv4TemplateApplication) HasOrphanedAddressConversionStrategy() bool`

HasOrphanedAddressConversionStrategy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


