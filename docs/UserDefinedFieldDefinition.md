# UserDefinedFieldDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | Pointer to **string** | The name of the user defined field. | [optional] 
**DisplayName** | Pointer to **string** | The display name of the user defined field. | [optional] 
**ResourceType** | Pointer to **string** | The resource type that the user defined field is applied to. | [optional] 
**DataType** | Pointer to **string** | The data type of the user defined field. | [optional] 
**DefaultValue** | Pointer to **string** | The default value of the user defined field. | [optional] 
**Required** | Pointer to **bool** | Indicates whether the user defined field is required or optional for the resource. | [optional] 
**PredefinedValues** | Pointer to **[]string** |  | [optional] 
**ValidationProperties** | Pointer to [**UserDefinedFieldDefinitionValidationProperties**](UserDefinedFieldDefinitionValidationProperties.md) |  | [optional] 

## Methods

### NewUserDefinedFieldDefinition

`func NewUserDefinedFieldDefinition() *UserDefinedFieldDefinition`

NewUserDefinedFieldDefinition instantiates a new UserDefinedFieldDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserDefinedFieldDefinitionWithDefaults

`func NewUserDefinedFieldDefinitionWithDefaults() *UserDefinedFieldDefinition`

NewUserDefinedFieldDefinitionWithDefaults instantiates a new UserDefinedFieldDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserDefinedFieldDefinition) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserDefinedFieldDefinition) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserDefinedFieldDefinition) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UserDefinedFieldDefinition) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *UserDefinedFieldDefinition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserDefinedFieldDefinition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserDefinedFieldDefinition) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UserDefinedFieldDefinition) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *UserDefinedFieldDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserDefinedFieldDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserDefinedFieldDefinition) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserDefinedFieldDefinition) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisplayName

`func (o *UserDefinedFieldDefinition) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UserDefinedFieldDefinition) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UserDefinedFieldDefinition) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *UserDefinedFieldDefinition) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetResourceType

`func (o *UserDefinedFieldDefinition) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *UserDefinedFieldDefinition) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *UserDefinedFieldDefinition) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *UserDefinedFieldDefinition) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetDataType

`func (o *UserDefinedFieldDefinition) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *UserDefinedFieldDefinition) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *UserDefinedFieldDefinition) SetDataType(v string)`

SetDataType sets DataType field to given value.

### HasDataType

`func (o *UserDefinedFieldDefinition) HasDataType() bool`

HasDataType returns a boolean if a field has been set.

### GetDefaultValue

`func (o *UserDefinedFieldDefinition) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *UserDefinedFieldDefinition) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *UserDefinedFieldDefinition) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *UserDefinedFieldDefinition) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetRequired

`func (o *UserDefinedFieldDefinition) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *UserDefinedFieldDefinition) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *UserDefinedFieldDefinition) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *UserDefinedFieldDefinition) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetPredefinedValues

`func (o *UserDefinedFieldDefinition) GetPredefinedValues() []string`

GetPredefinedValues returns the PredefinedValues field if non-nil, zero value otherwise.

### GetPredefinedValuesOk

`func (o *UserDefinedFieldDefinition) GetPredefinedValuesOk() (*[]string, bool)`

GetPredefinedValuesOk returns a tuple with the PredefinedValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredefinedValues

`func (o *UserDefinedFieldDefinition) SetPredefinedValues(v []string)`

SetPredefinedValues sets PredefinedValues field to given value.

### HasPredefinedValues

`func (o *UserDefinedFieldDefinition) HasPredefinedValues() bool`

HasPredefinedValues returns a boolean if a field has been set.

### GetValidationProperties

`func (o *UserDefinedFieldDefinition) GetValidationProperties() UserDefinedFieldDefinitionValidationProperties`

GetValidationProperties returns the ValidationProperties field if non-nil, zero value otherwise.

### GetValidationPropertiesOk

`func (o *UserDefinedFieldDefinition) GetValidationPropertiesOk() (*UserDefinedFieldDefinitionValidationProperties, bool)`

GetValidationPropertiesOk returns a tuple with the ValidationProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationProperties

`func (o *UserDefinedFieldDefinition) SetValidationProperties(v UserDefinedFieldDefinitionValidationProperties)`

SetValidationProperties sets ValidationProperties field to given value.

### HasValidationProperties

`func (o *UserDefinedFieldDefinition) HasValidationProperties() bool`

HasValidationProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


