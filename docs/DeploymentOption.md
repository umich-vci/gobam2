# DeploymentOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | Pointer to **string** | The name of the resource. | [optional] 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**Configuration** | Pointer to [**InlinedConfiguration**](InlinedConfiguration.md) |  | [optional] [readonly] 
**ServerScope** | Pointer to [**InlinedServerScope**](InlinedServerScope.md) |  | [optional] 

## Methods

### NewDeploymentOption

`func NewDeploymentOption() *DeploymentOption`

NewDeploymentOption instantiates a new DeploymentOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentOptionWithDefaults

`func NewDeploymentOptionWithDefaults() *DeploymentOption`

NewDeploymentOptionWithDefaults instantiates a new DeploymentOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeploymentOption) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeploymentOption) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeploymentOption) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DeploymentOption) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *DeploymentOption) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DeploymentOption) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DeploymentOption) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DeploymentOption) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *DeploymentOption) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeploymentOption) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeploymentOption) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeploymentOption) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *DeploymentOption) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *DeploymentOption) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *DeploymentOption) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *DeploymentOption) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetConfiguration

`func (o *DeploymentOption) GetConfiguration() InlinedConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *DeploymentOption) GetConfigurationOk() (*InlinedConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *DeploymentOption) SetConfiguration(v InlinedConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *DeploymentOption) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetServerScope

`func (o *DeploymentOption) GetServerScope() InlinedServerScope`

GetServerScope returns the ServerScope field if non-nil, zero value otherwise.

### GetServerScopeOk

`func (o *DeploymentOption) GetServerScopeOk() (*InlinedServerScope, bool)`

GetServerScopeOk returns a tuple with the ServerScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerScope

`func (o *DeploymentOption) SetServerScope(v InlinedServerScope)`

SetServerScope sets ServerScope field to given value.

### HasServerScope

`func (o *DeploymentOption) HasServerScope() bool`

HasServerScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

