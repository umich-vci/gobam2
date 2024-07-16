# DeploymentRoleMove

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The resource type. | [optional] 
**RoleTypes** | Pointer to **[]string** |  | [optional] 
**Destination** | Pointer to [**InlinedServerInterface**](InlinedServerInterface.md) |  | [optional] 

## Methods

### NewDeploymentRoleMove

`func NewDeploymentRoleMove() *DeploymentRoleMove`

NewDeploymentRoleMove instantiates a new DeploymentRoleMove object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentRoleMoveWithDefaults

`func NewDeploymentRoleMoveWithDefaults() *DeploymentRoleMove`

NewDeploymentRoleMoveWithDefaults instantiates a new DeploymentRoleMove object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DeploymentRoleMove) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DeploymentRoleMove) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DeploymentRoleMove) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DeploymentRoleMove) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRoleTypes

`func (o *DeploymentRoleMove) GetRoleTypes() []string`

GetRoleTypes returns the RoleTypes field if non-nil, zero value otherwise.

### GetRoleTypesOk

`func (o *DeploymentRoleMove) GetRoleTypesOk() (*[]string, bool)`

GetRoleTypesOk returns a tuple with the RoleTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleTypes

`func (o *DeploymentRoleMove) SetRoleTypes(v []string)`

SetRoleTypes sets RoleTypes field to given value.

### HasRoleTypes

`func (o *DeploymentRoleMove) HasRoleTypes() bool`

HasRoleTypes returns a boolean if a field has been set.

### GetDestination

`func (o *DeploymentRoleMove) GetDestination() InlinedServerInterface`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *DeploymentRoleMove) GetDestinationOk() (*InlinedServerInterface, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *DeploymentRoleMove) SetDestination(v InlinedServerInterface)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *DeploymentRoleMove) HasDestination() bool`

HasDestination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


