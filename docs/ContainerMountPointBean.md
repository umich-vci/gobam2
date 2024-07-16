# ContainerMountPointBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | Pointer to **string** | The DNS/DHCP Server directory that is bind mounted to the container. | [optional] 
**ContainerPath** | Pointer to **string** | The container directory that is mapped to the DNS/DHCP Server directory. | [optional] 

## Methods

### NewContainerMountPointBean

`func NewContainerMountPointBean() *ContainerMountPointBean`

NewContainerMountPointBean instantiates a new ContainerMountPointBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerMountPointBeanWithDefaults

`func NewContainerMountPointBeanWithDefaults() *ContainerMountPointBean`

NewContainerMountPointBeanWithDefaults instantiates a new ContainerMountPointBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *ContainerMountPointBean) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ContainerMountPointBean) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ContainerMountPointBean) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ContainerMountPointBean) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetContainerPath

`func (o *ContainerMountPointBean) GetContainerPath() string`

GetContainerPath returns the ContainerPath field if non-nil, zero value otherwise.

### GetContainerPathOk

`func (o *ContainerMountPointBean) GetContainerPathOk() (*string, bool)`

GetContainerPathOk returns a tuple with the ContainerPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerPath

`func (o *ContainerMountPointBean) SetContainerPath(v string)`

SetContainerPath sets ContainerPath field to given value.

### HasContainerPath

`func (o *ContainerMountPointBean) HasContainerPath() bool`

HasContainerPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


