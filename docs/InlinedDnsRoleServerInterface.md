# InlinedDnsRoleServerInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | Pointer to **string** | The name of the resource. | [optional] [readonly] 
**DeploymentRoleInterfaceType** | Pointer to **string** | The type of the deployment role server interface | [optional] 
**DeploymentRoleInterfaceProtocol** | Pointer to **string** | The internet protocol used for DNS zone transfers. | [optional] 

## Methods

### NewInlinedDnsRoleServerInterface

`func NewInlinedDnsRoleServerInterface() *InlinedDnsRoleServerInterface`

NewInlinedDnsRoleServerInterface instantiates a new InlinedDnsRoleServerInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlinedDnsRoleServerInterfaceWithDefaults

`func NewInlinedDnsRoleServerInterfaceWithDefaults() *InlinedDnsRoleServerInterface`

NewInlinedDnsRoleServerInterfaceWithDefaults instantiates a new InlinedDnsRoleServerInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlinedDnsRoleServerInterface) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlinedDnsRoleServerInterface) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlinedDnsRoleServerInterface) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *InlinedDnsRoleServerInterface) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *InlinedDnsRoleServerInterface) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlinedDnsRoleServerInterface) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlinedDnsRoleServerInterface) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InlinedDnsRoleServerInterface) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *InlinedDnsRoleServerInterface) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlinedDnsRoleServerInterface) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlinedDnsRoleServerInterface) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlinedDnsRoleServerInterface) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDeploymentRoleInterfaceType

`func (o *InlinedDnsRoleServerInterface) GetDeploymentRoleInterfaceType() string`

GetDeploymentRoleInterfaceType returns the DeploymentRoleInterfaceType field if non-nil, zero value otherwise.

### GetDeploymentRoleInterfaceTypeOk

`func (o *InlinedDnsRoleServerInterface) GetDeploymentRoleInterfaceTypeOk() (*string, bool)`

GetDeploymentRoleInterfaceTypeOk returns a tuple with the DeploymentRoleInterfaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentRoleInterfaceType

`func (o *InlinedDnsRoleServerInterface) SetDeploymentRoleInterfaceType(v string)`

SetDeploymentRoleInterfaceType sets DeploymentRoleInterfaceType field to given value.

### HasDeploymentRoleInterfaceType

`func (o *InlinedDnsRoleServerInterface) HasDeploymentRoleInterfaceType() bool`

HasDeploymentRoleInterfaceType returns a boolean if a field has been set.

### GetDeploymentRoleInterfaceProtocol

`func (o *InlinedDnsRoleServerInterface) GetDeploymentRoleInterfaceProtocol() string`

GetDeploymentRoleInterfaceProtocol returns the DeploymentRoleInterfaceProtocol field if non-nil, zero value otherwise.

### GetDeploymentRoleInterfaceProtocolOk

`func (o *InlinedDnsRoleServerInterface) GetDeploymentRoleInterfaceProtocolOk() (*string, bool)`

GetDeploymentRoleInterfaceProtocolOk returns a tuple with the DeploymentRoleInterfaceProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentRoleInterfaceProtocol

`func (o *InlinedDnsRoleServerInterface) SetDeploymentRoleInterfaceProtocol(v string)`

SetDeploymentRoleInterfaceProtocol sets DeploymentRoleInterfaceProtocol field to given value.

### HasDeploymentRoleInterfaceProtocol

`func (o *InlinedDnsRoleServerInterface) HasDeploymentRoleInterfaceProtocol() bool`

HasDeploymentRoleInterfaceProtocol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


