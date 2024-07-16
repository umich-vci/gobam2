# GatewayContainerBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the Gateway instance deployed to the DNS/DHCP Server. | [optional] [readonly] 
**Image** | Pointer to **string** | The Gateway image deployed on the DNS/DHCP Server. | [optional] 
**EnvironmentVariables** | Pointer to [**[]NameValuePairBean**](NameValuePairBean.md) | The environment variables to set for the container during startup. | [optional] 
**PortMappings** | Pointer to [**[]ContainerPortMappingBean**](ContainerPortMappingBean.md) | The list of ports to map between the DNS/DHCP Server service host and the Gateway container. | [optional] 
**MountPoints** | Pointer to [**[]ContainerMountPointBean**](ContainerMountPointBean.md) | The directories to bind mount to the Gateway container. | [optional] 

## Methods

### NewGatewayContainerBean

`func NewGatewayContainerBean() *GatewayContainerBean`

NewGatewayContainerBean instantiates a new GatewayContainerBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayContainerBeanWithDefaults

`func NewGatewayContainerBeanWithDefaults() *GatewayContainerBean`

NewGatewayContainerBeanWithDefaults instantiates a new GatewayContainerBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GatewayContainerBean) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayContainerBean) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayContainerBean) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GatewayContainerBean) HasName() bool`

HasName returns a boolean if a field has been set.

### GetImage

`func (o *GatewayContainerBean) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *GatewayContainerBean) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *GatewayContainerBean) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *GatewayContainerBean) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetEnvironmentVariables

`func (o *GatewayContainerBean) GetEnvironmentVariables() []NameValuePairBean`

GetEnvironmentVariables returns the EnvironmentVariables field if non-nil, zero value otherwise.

### GetEnvironmentVariablesOk

`func (o *GatewayContainerBean) GetEnvironmentVariablesOk() (*[]NameValuePairBean, bool)`

GetEnvironmentVariablesOk returns a tuple with the EnvironmentVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentVariables

`func (o *GatewayContainerBean) SetEnvironmentVariables(v []NameValuePairBean)`

SetEnvironmentVariables sets EnvironmentVariables field to given value.

### HasEnvironmentVariables

`func (o *GatewayContainerBean) HasEnvironmentVariables() bool`

HasEnvironmentVariables returns a boolean if a field has been set.

### GetPortMappings

`func (o *GatewayContainerBean) GetPortMappings() []ContainerPortMappingBean`

GetPortMappings returns the PortMappings field if non-nil, zero value otherwise.

### GetPortMappingsOk

`func (o *GatewayContainerBean) GetPortMappingsOk() (*[]ContainerPortMappingBean, bool)`

GetPortMappingsOk returns a tuple with the PortMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortMappings

`func (o *GatewayContainerBean) SetPortMappings(v []ContainerPortMappingBean)`

SetPortMappings sets PortMappings field to given value.

### HasPortMappings

`func (o *GatewayContainerBean) HasPortMappings() bool`

HasPortMappings returns a boolean if a field has been set.

### GetMountPoints

`func (o *GatewayContainerBean) GetMountPoints() []ContainerMountPointBean`

GetMountPoints returns the MountPoints field if non-nil, zero value otherwise.

### GetMountPointsOk

`func (o *GatewayContainerBean) GetMountPointsOk() (*[]ContainerMountPointBean, bool)`

GetMountPointsOk returns a tuple with the MountPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountPoints

`func (o *GatewayContainerBean) SetMountPoints(v []ContainerMountPointBean)`

SetMountPoints sets MountPoints field to given value.

### HasMountPoints

`func (o *GatewayContainerBean) HasMountPoints() bool`

HasMountPoints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


