# ContainerPortMappingBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostPort** | Pointer to **int32** | The DNS/DHCP Server port that is mapped to the container port. | [optional] 
**ContainerPort** | Pointer to **int32** | The host container port that is mapped to the DNS/DHCP Server port. | [optional] 

## Methods

### NewContainerPortMappingBean

`func NewContainerPortMappingBean() *ContainerPortMappingBean`

NewContainerPortMappingBean instantiates a new ContainerPortMappingBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerPortMappingBeanWithDefaults

`func NewContainerPortMappingBeanWithDefaults() *ContainerPortMappingBean`

NewContainerPortMappingBeanWithDefaults instantiates a new ContainerPortMappingBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostPort

`func (o *ContainerPortMappingBean) GetHostPort() int32`

GetHostPort returns the HostPort field if non-nil, zero value otherwise.

### GetHostPortOk

`func (o *ContainerPortMappingBean) GetHostPortOk() (*int32, bool)`

GetHostPortOk returns a tuple with the HostPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostPort

`func (o *ContainerPortMappingBean) SetHostPort(v int32)`

SetHostPort sets HostPort field to given value.

### HasHostPort

`func (o *ContainerPortMappingBean) HasHostPort() bool`

HasHostPort returns a boolean if a field has been set.

### GetContainerPort

`func (o *ContainerPortMappingBean) GetContainerPort() int32`

GetContainerPort returns the ContainerPort field if non-nil, zero value otherwise.

### GetContainerPortOk

`func (o *ContainerPortMappingBean) GetContainerPortOk() (*int32, bool)`

GetContainerPortOk returns a tuple with the ContainerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerPort

`func (o *ContainerPortMappingBean) SetContainerPort(v int32)`

SetContainerPort sets ContainerPort field to given value.

### HasContainerPort

`func (o *ContainerPortMappingBean) HasContainerPort() bool`

HasContainerPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


