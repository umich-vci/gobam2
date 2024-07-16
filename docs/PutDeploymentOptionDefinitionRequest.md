# PutDeploymentOptionDefinitionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | **string** | The type of deployment option definition. | 
**Name** | **string** | The name of the deployment option definition. | 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**AssignableTypes** | Pointer to **[]string** |  | [optional] [readonly] 
**Code** | Pointer to **int32** | The ID of the vendor profile option definition. | [optional] 
**DisplayName** | **string** | The display name for the vendor profile option definition. | 
**Description** | **string** | The description of the information passed by the option definition. | 
**ValueFormat** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPutDeploymentOptionDefinitionRequest

`func NewPutDeploymentOptionDefinitionRequest(type_ string, name string, displayName string, description string, ) *PutDeploymentOptionDefinitionRequest`

NewPutDeploymentOptionDefinitionRequest instantiates a new PutDeploymentOptionDefinitionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutDeploymentOptionDefinitionRequestWithDefaults

`func NewPutDeploymentOptionDefinitionRequestWithDefaults() *PutDeploymentOptionDefinitionRequest`

NewPutDeploymentOptionDefinitionRequestWithDefaults instantiates a new PutDeploymentOptionDefinitionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PutDeploymentOptionDefinitionRequest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PutDeploymentOptionDefinitionRequest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PutDeploymentOptionDefinitionRequest) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *PutDeploymentOptionDefinitionRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *PutDeploymentOptionDefinitionRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PutDeploymentOptionDefinitionRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PutDeploymentOptionDefinitionRequest) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *PutDeploymentOptionDefinitionRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PutDeploymentOptionDefinitionRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PutDeploymentOptionDefinitionRequest) SetName(v string)`

SetName sets Name field to given value.


### GetUserDefinedFields

`func (o *PutDeploymentOptionDefinitionRequest) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *PutDeploymentOptionDefinitionRequest) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *PutDeploymentOptionDefinitionRequest) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *PutDeploymentOptionDefinitionRequest) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetAssignableTypes

`func (o *PutDeploymentOptionDefinitionRequest) GetAssignableTypes() []string`

GetAssignableTypes returns the AssignableTypes field if non-nil, zero value otherwise.

### GetAssignableTypesOk

`func (o *PutDeploymentOptionDefinitionRequest) GetAssignableTypesOk() (*[]string, bool)`

GetAssignableTypesOk returns a tuple with the AssignableTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignableTypes

`func (o *PutDeploymentOptionDefinitionRequest) SetAssignableTypes(v []string)`

SetAssignableTypes sets AssignableTypes field to given value.

### HasAssignableTypes

`func (o *PutDeploymentOptionDefinitionRequest) HasAssignableTypes() bool`

HasAssignableTypes returns a boolean if a field has been set.

### GetCode

`func (o *PutDeploymentOptionDefinitionRequest) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *PutDeploymentOptionDefinitionRequest) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *PutDeploymentOptionDefinitionRequest) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *PutDeploymentOptionDefinitionRequest) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDisplayName

`func (o *PutDeploymentOptionDefinitionRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PutDeploymentOptionDefinitionRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PutDeploymentOptionDefinitionRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetDescription

`func (o *PutDeploymentOptionDefinitionRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PutDeploymentOptionDefinitionRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PutDeploymentOptionDefinitionRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetValueFormat

`func (o *PutDeploymentOptionDefinitionRequest) GetValueFormat() map[string]interface{}`

GetValueFormat returns the ValueFormat field if non-nil, zero value otherwise.

### GetValueFormatOk

`func (o *PutDeploymentOptionDefinitionRequest) GetValueFormatOk() (*map[string]interface{}, bool)`

GetValueFormatOk returns a tuple with the ValueFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueFormat

`func (o *PutDeploymentOptionDefinitionRequest) SetValueFormat(v map[string]interface{})`

SetValueFormat sets ValueFormat field to given value.

### HasValueFormat

`func (o *PutDeploymentOptionDefinitionRequest) HasValueFormat() bool`

HasValueFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


