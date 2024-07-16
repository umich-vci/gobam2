# PostCollectionDeploymentOptionDefinitionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | **string** | The type of deployment option definition. | 
**Name** | **string** | The name of the deployment option definition. | 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**AssignableTypes** | Pointer to **[]string** |  | [optional] [readonly] 
**Code** | **int32** | The ID of the vendor profile option definition. | 
**DisplayName** | **string** | The display name for the vendor profile option definition. | 
**Description** | **string** | The description of the information passed by the option definition. | 
**ValueFormat** | **map[string]interface{}** |  | 

## Methods

### NewPostCollectionDeploymentOptionDefinitionRequest

`func NewPostCollectionDeploymentOptionDefinitionRequest(type_ string, name string, code int32, displayName string, description string, valueFormat map[string]interface{}, ) *PostCollectionDeploymentOptionDefinitionRequest`

NewPostCollectionDeploymentOptionDefinitionRequest instantiates a new PostCollectionDeploymentOptionDefinitionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostCollectionDeploymentOptionDefinitionRequestWithDefaults

`func NewPostCollectionDeploymentOptionDefinitionRequestWithDefaults() *PostCollectionDeploymentOptionDefinitionRequest`

NewPostCollectionDeploymentOptionDefinitionRequestWithDefaults instantiates a new PostCollectionDeploymentOptionDefinitionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PostCollectionDeploymentOptionDefinitionRequest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostCollectionDeploymentOptionDefinitionRequest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostCollectionDeploymentOptionDefinitionRequest) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *PostCollectionDeploymentOptionDefinitionRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *PostCollectionDeploymentOptionDefinitionRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PostCollectionDeploymentOptionDefinitionRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PostCollectionDeploymentOptionDefinitionRequest) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *PostCollectionDeploymentOptionDefinitionRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostCollectionDeploymentOptionDefinitionRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostCollectionDeploymentOptionDefinitionRequest) SetName(v string)`

SetName sets Name field to given value.


### GetUserDefinedFields

`func (o *PostCollectionDeploymentOptionDefinitionRequest) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *PostCollectionDeploymentOptionDefinitionRequest) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *PostCollectionDeploymentOptionDefinitionRequest) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *PostCollectionDeploymentOptionDefinitionRequest) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetAssignableTypes

`func (o *PostCollectionDeploymentOptionDefinitionRequest) GetAssignableTypes() []string`

GetAssignableTypes returns the AssignableTypes field if non-nil, zero value otherwise.

### GetAssignableTypesOk

`func (o *PostCollectionDeploymentOptionDefinitionRequest) GetAssignableTypesOk() (*[]string, bool)`

GetAssignableTypesOk returns a tuple with the AssignableTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignableTypes

`func (o *PostCollectionDeploymentOptionDefinitionRequest) SetAssignableTypes(v []string)`

SetAssignableTypes sets AssignableTypes field to given value.

### HasAssignableTypes

`func (o *PostCollectionDeploymentOptionDefinitionRequest) HasAssignableTypes() bool`

HasAssignableTypes returns a boolean if a field has been set.

### GetCode

`func (o *PostCollectionDeploymentOptionDefinitionRequest) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *PostCollectionDeploymentOptionDefinitionRequest) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *PostCollectionDeploymentOptionDefinitionRequest) SetCode(v int32)`

SetCode sets Code field to given value.


### GetDisplayName

`func (o *PostCollectionDeploymentOptionDefinitionRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PostCollectionDeploymentOptionDefinitionRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PostCollectionDeploymentOptionDefinitionRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetDescription

`func (o *PostCollectionDeploymentOptionDefinitionRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PostCollectionDeploymentOptionDefinitionRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PostCollectionDeploymentOptionDefinitionRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetValueFormat

`func (o *PostCollectionDeploymentOptionDefinitionRequest) GetValueFormat() map[string]interface{}`

GetValueFormat returns the ValueFormat field if non-nil, zero value otherwise.

### GetValueFormatOk

`func (o *PostCollectionDeploymentOptionDefinitionRequest) GetValueFormatOk() (*map[string]interface{}, bool)`

GetValueFormatOk returns a tuple with the ValueFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueFormat

`func (o *PostCollectionDeploymentOptionDefinitionRequest) SetValueFormat(v map[string]interface{})`

SetValueFormat sets ValueFormat field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


