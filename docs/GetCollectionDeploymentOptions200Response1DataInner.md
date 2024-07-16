# GetCollectionDeploymentOptions200Response1DataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | Pointer to **string** | The name of the resource. | [optional] 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**Configuration** | Pointer to [**InlinedConfiguration**](InlinedConfiguration.md) |  | [optional] [readonly] 
**ServerScope** | Pointer to [**InlinedServerScope**](InlinedServerScope.md) |  | [optional] 
**Value** | Pointer to **map[string]interface{}** |  | [optional] 
**Code** | Pointer to **float32** |  | [optional] 
**Definition** | Pointer to [**InlinedStartOfAuthorityDefinition**](InlinedStartOfAuthorityDefinition.md) |  | [optional] 

## Methods

### NewGetCollectionDeploymentOptions200Response1DataInner

`func NewGetCollectionDeploymentOptions200Response1DataInner() *GetCollectionDeploymentOptions200Response1DataInner`

NewGetCollectionDeploymentOptions200Response1DataInner instantiates a new GetCollectionDeploymentOptions200Response1DataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCollectionDeploymentOptions200Response1DataInnerWithDefaults

`func NewGetCollectionDeploymentOptions200Response1DataInnerWithDefaults() *GetCollectionDeploymentOptions200Response1DataInner`

NewGetCollectionDeploymentOptions200Response1DataInnerWithDefaults instantiates a new GetCollectionDeploymentOptions200Response1DataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetCollectionDeploymentOptions200Response1DataInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetCollectionDeploymentOptions200Response1DataInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetCollectionDeploymentOptions200Response1DataInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetCollectionDeploymentOptions200Response1DataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GetCollectionDeploymentOptions200Response1DataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetCollectionDeploymentOptions200Response1DataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetCollectionDeploymentOptions200Response1DataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetCollectionDeploymentOptions200Response1DataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *GetCollectionDeploymentOptions200Response1DataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetCollectionDeploymentOptions200Response1DataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetCollectionDeploymentOptions200Response1DataInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetCollectionDeploymentOptions200Response1DataInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *GetCollectionDeploymentOptions200Response1DataInner) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *GetCollectionDeploymentOptions200Response1DataInner) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *GetCollectionDeploymentOptions200Response1DataInner) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *GetCollectionDeploymentOptions200Response1DataInner) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetConfiguration

`func (o *GetCollectionDeploymentOptions200Response1DataInner) GetConfiguration() InlinedConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *GetCollectionDeploymentOptions200Response1DataInner) GetConfigurationOk() (*InlinedConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *GetCollectionDeploymentOptions200Response1DataInner) SetConfiguration(v InlinedConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *GetCollectionDeploymentOptions200Response1DataInner) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetServerScope

`func (o *GetCollectionDeploymentOptions200Response1DataInner) GetServerScope() InlinedServerScope`

GetServerScope returns the ServerScope field if non-nil, zero value otherwise.

### GetServerScopeOk

`func (o *GetCollectionDeploymentOptions200Response1DataInner) GetServerScopeOk() (*InlinedServerScope, bool)`

GetServerScopeOk returns a tuple with the ServerScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerScope

`func (o *GetCollectionDeploymentOptions200Response1DataInner) SetServerScope(v InlinedServerScope)`

SetServerScope sets ServerScope field to given value.

### HasServerScope

`func (o *GetCollectionDeploymentOptions200Response1DataInner) HasServerScope() bool`

HasServerScope returns a boolean if a field has been set.

### GetValue

`func (o *GetCollectionDeploymentOptions200Response1DataInner) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GetCollectionDeploymentOptions200Response1DataInner) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GetCollectionDeploymentOptions200Response1DataInner) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *GetCollectionDeploymentOptions200Response1DataInner) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetCode

`func (o *GetCollectionDeploymentOptions200Response1DataInner) GetCode() float32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetCollectionDeploymentOptions200Response1DataInner) GetCodeOk() (*float32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetCollectionDeploymentOptions200Response1DataInner) SetCode(v float32)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetCollectionDeploymentOptions200Response1DataInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDefinition

`func (o *GetCollectionDeploymentOptions200Response1DataInner) GetDefinition() InlinedStartOfAuthorityDefinition`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *GetCollectionDeploymentOptions200Response1DataInner) GetDefinitionOk() (*InlinedStartOfAuthorityDefinition, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *GetCollectionDeploymentOptions200Response1DataInner) SetDefinition(v InlinedStartOfAuthorityDefinition)`

SetDefinition sets Definition field to given value.

### HasDefinition

`func (o *GetCollectionDeploymentOptions200Response1DataInner) HasDefinition() bool`

HasDefinition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

