# PutDeploymentOptionRequest

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
**Definition** | Pointer to [**InlinedDHCPVendorOptionDefinition**](InlinedDHCPVendorOptionDefinition.md) |  | [optional] 

## Methods

### NewPutDeploymentOptionRequest

`func NewPutDeploymentOptionRequest(type_ string, value map[string]interface{}, ) *PutDeploymentOptionRequest`

NewPutDeploymentOptionRequest instantiates a new PutDeploymentOptionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutDeploymentOptionRequestWithDefaults

`func NewPutDeploymentOptionRequestWithDefaults() *PutDeploymentOptionRequest`

NewPutDeploymentOptionRequestWithDefaults instantiates a new PutDeploymentOptionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PutDeploymentOptionRequest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PutDeploymentOptionRequest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PutDeploymentOptionRequest) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *PutDeploymentOptionRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *PutDeploymentOptionRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PutDeploymentOptionRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PutDeploymentOptionRequest) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *PutDeploymentOptionRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PutDeploymentOptionRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PutDeploymentOptionRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PutDeploymentOptionRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *PutDeploymentOptionRequest) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *PutDeploymentOptionRequest) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *PutDeploymentOptionRequest) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *PutDeploymentOptionRequest) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetConfiguration

`func (o *PutDeploymentOptionRequest) GetConfiguration() InlinedConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *PutDeploymentOptionRequest) GetConfigurationOk() (*InlinedConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *PutDeploymentOptionRequest) SetConfiguration(v InlinedConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *PutDeploymentOptionRequest) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetServerScope

`func (o *PutDeploymentOptionRequest) GetServerScope() InlinedServerScope`

GetServerScope returns the ServerScope field if non-nil, zero value otherwise.

### GetServerScopeOk

`func (o *PutDeploymentOptionRequest) GetServerScopeOk() (*InlinedServerScope, bool)`

GetServerScopeOk returns a tuple with the ServerScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerScope

`func (o *PutDeploymentOptionRequest) SetServerScope(v InlinedServerScope)`

SetServerScope sets ServerScope field to given value.

### HasServerScope

`func (o *PutDeploymentOptionRequest) HasServerScope() bool`

HasServerScope returns a boolean if a field has been set.

### GetValue

`func (o *PutDeploymentOptionRequest) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PutDeploymentOptionRequest) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PutDeploymentOptionRequest) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.


### GetCode

`func (o *PutDeploymentOptionRequest) GetCode() float32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *PutDeploymentOptionRequest) GetCodeOk() (*float32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *PutDeploymentOptionRequest) SetCode(v float32)`

SetCode sets Code field to given value.

### HasCode

`func (o *PutDeploymentOptionRequest) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDefinition

`func (o *PutDeploymentOptionRequest) GetDefinition() InlinedDHCPVendorOptionDefinition`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *PutDeploymentOptionRequest) GetDefinitionOk() (*InlinedDHCPVendorOptionDefinition, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *PutDeploymentOptionRequest) SetDefinition(v InlinedDHCPVendorOptionDefinition)`

SetDefinition sets Definition field to given value.

### HasDefinition

`func (o *PutDeploymentOptionRequest) HasDefinition() bool`

HasDefinition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


