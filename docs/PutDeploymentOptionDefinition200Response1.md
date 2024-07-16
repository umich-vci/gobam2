# PutDeploymentOptionDefinition200Response1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The type of deployment option definition. | [optional] 
**Name** | Pointer to **string** | The name of the deployment option definition. | [optional] 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**AssignableTypes** | Pointer to **[]string** |  | [optional] [readonly] 
**Code** | Pointer to **int32** | The ID of the vendor profile option definition. | [optional] 
**DisplayName** | Pointer to **string** | The display name for the vendor profile option definition. | [optional] 
**Description** | Pointer to **string** | The description of the information passed by the option definition. | [optional] 
**ValueFormat** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPutDeploymentOptionDefinition200Response1

`func NewPutDeploymentOptionDefinition200Response1() *PutDeploymentOptionDefinition200Response1`

NewPutDeploymentOptionDefinition200Response1 instantiates a new PutDeploymentOptionDefinition200Response1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutDeploymentOptionDefinition200Response1WithDefaults

`func NewPutDeploymentOptionDefinition200Response1WithDefaults() *PutDeploymentOptionDefinition200Response1`

NewPutDeploymentOptionDefinition200Response1WithDefaults instantiates a new PutDeploymentOptionDefinition200Response1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PutDeploymentOptionDefinition200Response1) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PutDeploymentOptionDefinition200Response1) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PutDeploymentOptionDefinition200Response1) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *PutDeploymentOptionDefinition200Response1) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *PutDeploymentOptionDefinition200Response1) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PutDeploymentOptionDefinition200Response1) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PutDeploymentOptionDefinition200Response1) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PutDeploymentOptionDefinition200Response1) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *PutDeploymentOptionDefinition200Response1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PutDeploymentOptionDefinition200Response1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PutDeploymentOptionDefinition200Response1) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PutDeploymentOptionDefinition200Response1) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *PutDeploymentOptionDefinition200Response1) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *PutDeploymentOptionDefinition200Response1) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *PutDeploymentOptionDefinition200Response1) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *PutDeploymentOptionDefinition200Response1) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetAssignableTypes

`func (o *PutDeploymentOptionDefinition200Response1) GetAssignableTypes() []string`

GetAssignableTypes returns the AssignableTypes field if non-nil, zero value otherwise.

### GetAssignableTypesOk

`func (o *PutDeploymentOptionDefinition200Response1) GetAssignableTypesOk() (*[]string, bool)`

GetAssignableTypesOk returns a tuple with the AssignableTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignableTypes

`func (o *PutDeploymentOptionDefinition200Response1) SetAssignableTypes(v []string)`

SetAssignableTypes sets AssignableTypes field to given value.

### HasAssignableTypes

`func (o *PutDeploymentOptionDefinition200Response1) HasAssignableTypes() bool`

HasAssignableTypes returns a boolean if a field has been set.

### GetCode

`func (o *PutDeploymentOptionDefinition200Response1) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *PutDeploymentOptionDefinition200Response1) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *PutDeploymentOptionDefinition200Response1) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *PutDeploymentOptionDefinition200Response1) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDisplayName

`func (o *PutDeploymentOptionDefinition200Response1) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PutDeploymentOptionDefinition200Response1) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PutDeploymentOptionDefinition200Response1) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *PutDeploymentOptionDefinition200Response1) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *PutDeploymentOptionDefinition200Response1) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PutDeploymentOptionDefinition200Response1) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PutDeploymentOptionDefinition200Response1) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PutDeploymentOptionDefinition200Response1) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetValueFormat

`func (o *PutDeploymentOptionDefinition200Response1) GetValueFormat() map[string]interface{}`

GetValueFormat returns the ValueFormat field if non-nil, zero value otherwise.

### GetValueFormatOk

`func (o *PutDeploymentOptionDefinition200Response1) GetValueFormatOk() (*map[string]interface{}, bool)`

GetValueFormatOk returns a tuple with the ValueFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueFormat

`func (o *PutDeploymentOptionDefinition200Response1) SetValueFormat(v map[string]interface{})`

SetValueFormat sets ValueFormat field to given value.

### HasValueFormat

`func (o *PutDeploymentOptionDefinition200Response1) HasValueFormat() bool`

HasValueFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


