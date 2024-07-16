# GetDeploymentOptionDefinitions200ResponseDataInner

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
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 

## Methods

### NewGetDeploymentOptionDefinitions200ResponseDataInner

`func NewGetDeploymentOptionDefinitions200ResponseDataInner() *GetDeploymentOptionDefinitions200ResponseDataInner`

NewGetDeploymentOptionDefinitions200ResponseDataInner instantiates a new GetDeploymentOptionDefinitions200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDeploymentOptionDefinitions200ResponseDataInnerWithDefaults

`func NewGetDeploymentOptionDefinitions200ResponseDataInnerWithDefaults() *GetDeploymentOptionDefinitions200ResponseDataInner`

NewGetDeploymentOptionDefinitions200ResponseDataInnerWithDefaults instantiates a new GetDeploymentOptionDefinitions200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetDeploymentOptionDefinitions200ResponseDataInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetDeploymentOptionDefinitions200ResponseDataInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetDeploymentOptionDefinitions200ResponseDataInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetDeploymentOptionDefinitions200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GetDeploymentOptionDefinitions200ResponseDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetDeploymentOptionDefinitions200ResponseDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetDeploymentOptionDefinitions200ResponseDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetDeploymentOptionDefinitions200ResponseDataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *GetDeploymentOptionDefinitions200ResponseDataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetDeploymentOptionDefinitions200ResponseDataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetDeploymentOptionDefinitions200ResponseDataInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetDeploymentOptionDefinitions200ResponseDataInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *GetDeploymentOptionDefinitions200ResponseDataInner) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *GetDeploymentOptionDefinitions200ResponseDataInner) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *GetDeploymentOptionDefinitions200ResponseDataInner) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *GetDeploymentOptionDefinitions200ResponseDataInner) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetAssignableTypes

`func (o *GetDeploymentOptionDefinitions200ResponseDataInner) GetAssignableTypes() []string`

GetAssignableTypes returns the AssignableTypes field if non-nil, zero value otherwise.

### GetAssignableTypesOk

`func (o *GetDeploymentOptionDefinitions200ResponseDataInner) GetAssignableTypesOk() (*[]string, bool)`

GetAssignableTypesOk returns a tuple with the AssignableTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignableTypes

`func (o *GetDeploymentOptionDefinitions200ResponseDataInner) SetAssignableTypes(v []string)`

SetAssignableTypes sets AssignableTypes field to given value.

### HasAssignableTypes

`func (o *GetDeploymentOptionDefinitions200ResponseDataInner) HasAssignableTypes() bool`

HasAssignableTypes returns a boolean if a field has been set.

### GetCode

`func (o *GetDeploymentOptionDefinitions200ResponseDataInner) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetDeploymentOptionDefinitions200ResponseDataInner) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetDeploymentOptionDefinitions200ResponseDataInner) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetDeploymentOptionDefinitions200ResponseDataInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDisplayName

`func (o *GetDeploymentOptionDefinitions200ResponseDataInner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GetDeploymentOptionDefinitions200ResponseDataInner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GetDeploymentOptionDefinitions200ResponseDataInner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *GetDeploymentOptionDefinitions200ResponseDataInner) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *GetDeploymentOptionDefinitions200ResponseDataInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetDeploymentOptionDefinitions200ResponseDataInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetDeploymentOptionDefinitions200ResponseDataInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetDeploymentOptionDefinitions200ResponseDataInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetValueFormat

`func (o *GetDeploymentOptionDefinitions200ResponseDataInner) GetValueFormat() map[string]interface{}`

GetValueFormat returns the ValueFormat field if non-nil, zero value otherwise.

### GetValueFormatOk

`func (o *GetDeploymentOptionDefinitions200ResponseDataInner) GetValueFormatOk() (*map[string]interface{}, bool)`

GetValueFormatOk returns a tuple with the ValueFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueFormat

`func (o *GetDeploymentOptionDefinitions200ResponseDataInner) SetValueFormat(v map[string]interface{})`

SetValueFormat sets ValueFormat field to given value.

### HasValueFormat

`func (o *GetDeploymentOptionDefinitions200ResponseDataInner) HasValueFormat() bool`

HasValueFormat returns a boolean if a field has been set.

### GetLinks

`func (o *GetDeploymentOptionDefinitions200ResponseDataInner) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetDeploymentOptionDefinitions200ResponseDataInner) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetDeploymentOptionDefinitions200ResponseDataInner) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetDeploymentOptionDefinitions200ResponseDataInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


