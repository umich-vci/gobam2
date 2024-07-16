# DHCPDeploymentRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The resource type. | [optional] 
**RoleType** | Pointer to **string** |  | [optional] 
**Interfaces** | Pointer to [**[]InlinedDhcpRoleServerInterface**](InlinedDhcpRoleServerInterface.md) | The list of server interfaces for the role. | [optional] 

## Methods

### NewDHCPDeploymentRole

`func NewDHCPDeploymentRole() *DHCPDeploymentRole`

NewDHCPDeploymentRole instantiates a new DHCPDeploymentRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDHCPDeploymentRoleWithDefaults

`func NewDHCPDeploymentRoleWithDefaults() *DHCPDeploymentRole`

NewDHCPDeploymentRoleWithDefaults instantiates a new DHCPDeploymentRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DHCPDeploymentRole) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DHCPDeploymentRole) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DHCPDeploymentRole) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DHCPDeploymentRole) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRoleType

`func (o *DHCPDeploymentRole) GetRoleType() string`

GetRoleType returns the RoleType field if non-nil, zero value otherwise.

### GetRoleTypeOk

`func (o *DHCPDeploymentRole) GetRoleTypeOk() (*string, bool)`

GetRoleTypeOk returns a tuple with the RoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleType

`func (o *DHCPDeploymentRole) SetRoleType(v string)`

SetRoleType sets RoleType field to given value.

### HasRoleType

`func (o *DHCPDeploymentRole) HasRoleType() bool`

HasRoleType returns a boolean if a field has been set.

### GetInterfaces

`func (o *DHCPDeploymentRole) GetInterfaces() []InlinedDhcpRoleServerInterface`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *DHCPDeploymentRole) GetInterfacesOk() (*[]InlinedDhcpRoleServerInterface, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *DHCPDeploymentRole) SetInterfaces(v []InlinedDhcpRoleServerInterface)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *DHCPDeploymentRole) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


