# DeploymentOptionDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The type of deployment option definition. | [optional] 
**Name** | Pointer to **string** | The name of the deployment option definition. | [optional] 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**AssignableTypes** | Pointer to **[]string** |  | [optional] [readonly] 

## Methods

### NewDeploymentOptionDefinition

`func NewDeploymentOptionDefinition() *DeploymentOptionDefinition`

NewDeploymentOptionDefinition instantiates a new DeploymentOptionDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentOptionDefinitionWithDefaults

`func NewDeploymentOptionDefinitionWithDefaults() *DeploymentOptionDefinition`

NewDeploymentOptionDefinitionWithDefaults instantiates a new DeploymentOptionDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeploymentOptionDefinition) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeploymentOptionDefinition) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeploymentOptionDefinition) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DeploymentOptionDefinition) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *DeploymentOptionDefinition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DeploymentOptionDefinition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DeploymentOptionDefinition) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DeploymentOptionDefinition) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *DeploymentOptionDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeploymentOptionDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeploymentOptionDefinition) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeploymentOptionDefinition) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *DeploymentOptionDefinition) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *DeploymentOptionDefinition) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *DeploymentOptionDefinition) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *DeploymentOptionDefinition) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetAssignableTypes

`func (o *DeploymentOptionDefinition) GetAssignableTypes() []string`

GetAssignableTypes returns the AssignableTypes field if non-nil, zero value otherwise.

### GetAssignableTypesOk

`func (o *DeploymentOptionDefinition) GetAssignableTypesOk() (*[]string, bool)`

GetAssignableTypesOk returns a tuple with the AssignableTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignableTypes

`func (o *DeploymentOptionDefinition) SetAssignableTypes(v []string)`

SetAssignableTypes sets AssignableTypes field to given value.

### HasAssignableTypes

`func (o *DeploymentOptionDefinition) HasAssignableTypes() bool`

HasAssignableTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


