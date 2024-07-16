# GatewayService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The resource type. | [optional] 
**Enabled** | Pointer to **bool** | Indicates whether the Gateway service is enabled. | [optional] 
**State** | Pointer to **string** | The status of the Gateway service running on the DNS/DHCP Server. | [optional] 
**Container** | Pointer to [**GatewayContainerBean**](GatewayContainerBean.md) |  | [optional] 
**RepositoryUsername** | Pointer to **string** | The docker username used to pull the Gateway image from the container repository. | [optional] 
**RepositoryPassword** | Pointer to **string** | The docker password used to pull the image from a private repository. | [optional] 
**DeleteImageOnDisable** | Pointer to **bool** | Indicates whether the Gateway image is removed upon disabling the Gateway service. | [optional] 
**DeleteMountPointsOnDisable** | Pointer to **bool** | Indicates whether the local volumes and the mounted data and log directories are removed upon disabling the Gateway service. | [optional] 

## Methods

### NewGatewayService

`func NewGatewayService() *GatewayService`

NewGatewayService instantiates a new GatewayService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayServiceWithDefaults

`func NewGatewayServiceWithDefaults() *GatewayService`

NewGatewayServiceWithDefaults instantiates a new GatewayService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GatewayService) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GatewayService) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GatewayService) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GatewayService) HasType() bool`

HasType returns a boolean if a field has been set.

### GetEnabled

`func (o *GatewayService) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GatewayService) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GatewayService) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GatewayService) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetState

`func (o *GatewayService) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GatewayService) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GatewayService) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GatewayService) HasState() bool`

HasState returns a boolean if a field has been set.

### GetContainer

`func (o *GatewayService) GetContainer() GatewayContainerBean`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *GatewayService) GetContainerOk() (*GatewayContainerBean, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *GatewayService) SetContainer(v GatewayContainerBean)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *GatewayService) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetRepositoryUsername

`func (o *GatewayService) GetRepositoryUsername() string`

GetRepositoryUsername returns the RepositoryUsername field if non-nil, zero value otherwise.

### GetRepositoryUsernameOk

`func (o *GatewayService) GetRepositoryUsernameOk() (*string, bool)`

GetRepositoryUsernameOk returns a tuple with the RepositoryUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryUsername

`func (o *GatewayService) SetRepositoryUsername(v string)`

SetRepositoryUsername sets RepositoryUsername field to given value.

### HasRepositoryUsername

`func (o *GatewayService) HasRepositoryUsername() bool`

HasRepositoryUsername returns a boolean if a field has been set.

### GetRepositoryPassword

`func (o *GatewayService) GetRepositoryPassword() string`

GetRepositoryPassword returns the RepositoryPassword field if non-nil, zero value otherwise.

### GetRepositoryPasswordOk

`func (o *GatewayService) GetRepositoryPasswordOk() (*string, bool)`

GetRepositoryPasswordOk returns a tuple with the RepositoryPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryPassword

`func (o *GatewayService) SetRepositoryPassword(v string)`

SetRepositoryPassword sets RepositoryPassword field to given value.

### HasRepositoryPassword

`func (o *GatewayService) HasRepositoryPassword() bool`

HasRepositoryPassword returns a boolean if a field has been set.

### GetDeleteImageOnDisable

`func (o *GatewayService) GetDeleteImageOnDisable() bool`

GetDeleteImageOnDisable returns the DeleteImageOnDisable field if non-nil, zero value otherwise.

### GetDeleteImageOnDisableOk

`func (o *GatewayService) GetDeleteImageOnDisableOk() (*bool, bool)`

GetDeleteImageOnDisableOk returns a tuple with the DeleteImageOnDisable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteImageOnDisable

`func (o *GatewayService) SetDeleteImageOnDisable(v bool)`

SetDeleteImageOnDisable sets DeleteImageOnDisable field to given value.

### HasDeleteImageOnDisable

`func (o *GatewayService) HasDeleteImageOnDisable() bool`

HasDeleteImageOnDisable returns a boolean if a field has been set.

### GetDeleteMountPointsOnDisable

`func (o *GatewayService) GetDeleteMountPointsOnDisable() bool`

GetDeleteMountPointsOnDisable returns the DeleteMountPointsOnDisable field if non-nil, zero value otherwise.

### GetDeleteMountPointsOnDisableOk

`func (o *GatewayService) GetDeleteMountPointsOnDisableOk() (*bool, bool)`

GetDeleteMountPointsOnDisableOk returns a tuple with the DeleteMountPointsOnDisable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteMountPointsOnDisable

`func (o *GatewayService) SetDeleteMountPointsOnDisable(v bool)`

SetDeleteMountPointsOnDisable sets DeleteMountPointsOnDisable field to given value.

### HasDeleteMountPointsOnDisable

`func (o *GatewayService) HasDeleteMountPointsOnDisable() bool`

HasDeleteMountPointsOnDisable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


