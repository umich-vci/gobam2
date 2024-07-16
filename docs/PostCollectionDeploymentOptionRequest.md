# PostCollectionDeploymentOptionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | **string** | The resource type. | 
**Name** | Pointer to **string** | The name of the resource. | [optional] 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**Configuration** | Pointer to [**InlinedConfiguration**](InlinedConfiguration.md) |  | [optional] [readonly] 
**ServerScope** | Pointer to [**InlinedServerScope**](InlinedServerScope.md) |  | [optional] 
**Value** | **map[string]interface{}** |  | 
**Code** | Pointer to **float32** | The vendor option code. | [optional] [readonly] 
**Definition** | [**InlinedDHCPVendorOptionDefinition**](InlinedDHCPVendorOptionDefinition.md) |  | 

## Methods

### NewPostCollectionDeploymentOptionRequest

`func NewPostCollectionDeploymentOptionRequest(type_ string, value map[string]interface{}, definition InlinedDHCPVendorOptionDefinition, ) *PostCollectionDeploymentOptionRequest`

NewPostCollectionDeploymentOptionRequest instantiates a new PostCollectionDeploymentOptionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostCollectionDeploymentOptionRequestWithDefaults

`func NewPostCollectionDeploymentOptionRequestWithDefaults() *PostCollectionDeploymentOptionRequest`

NewPostCollectionDeploymentOptionRequestWithDefaults instantiates a new PostCollectionDeploymentOptionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PostCollectionDeploymentOptionRequest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostCollectionDeploymentOptionRequest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostCollectionDeploymentOptionRequest) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *PostCollectionDeploymentOptionRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *PostCollectionDeploymentOptionRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PostCollectionDeploymentOptionRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PostCollectionDeploymentOptionRequest) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *PostCollectionDeploymentOptionRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostCollectionDeploymentOptionRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostCollectionDeploymentOptionRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PostCollectionDeploymentOptionRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *PostCollectionDeploymentOptionRequest) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *PostCollectionDeploymentOptionRequest) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *PostCollectionDeploymentOptionRequest) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *PostCollectionDeploymentOptionRequest) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetConfiguration

`func (o *PostCollectionDeploymentOptionRequest) GetConfiguration() InlinedConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *PostCollectionDeploymentOptionRequest) GetConfigurationOk() (*InlinedConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *PostCollectionDeploymentOptionRequest) SetConfiguration(v InlinedConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *PostCollectionDeploymentOptionRequest) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetServerScope

`func (o *PostCollectionDeploymentOptionRequest) GetServerScope() InlinedServerScope`

GetServerScope returns the ServerScope field if non-nil, zero value otherwise.

### GetServerScopeOk

`func (o *PostCollectionDeploymentOptionRequest) GetServerScopeOk() (*InlinedServerScope, bool)`

GetServerScopeOk returns a tuple with the ServerScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerScope

`func (o *PostCollectionDeploymentOptionRequest) SetServerScope(v InlinedServerScope)`

SetServerScope sets ServerScope field to given value.

### HasServerScope

`func (o *PostCollectionDeploymentOptionRequest) HasServerScope() bool`

HasServerScope returns a boolean if a field has been set.

### GetValue

`func (o *PostCollectionDeploymentOptionRequest) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PostCollectionDeploymentOptionRequest) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PostCollectionDeploymentOptionRequest) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.


### GetCode

`func (o *PostCollectionDeploymentOptionRequest) GetCode() float32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *PostCollectionDeploymentOptionRequest) GetCodeOk() (*float32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *PostCollectionDeploymentOptionRequest) SetCode(v float32)`

SetCode sets Code field to given value.

### HasCode

`func (o *PostCollectionDeploymentOptionRequest) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDefinition

`func (o *PostCollectionDeploymentOptionRequest) GetDefinition() InlinedDHCPVendorOptionDefinition`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *PostCollectionDeploymentOptionRequest) GetDefinitionOk() (*InlinedDHCPVendorOptionDefinition, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *PostCollectionDeploymentOptionRequest) SetDefinition(v InlinedDHCPVendorOptionDefinition)`

SetDefinition sets Definition field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


