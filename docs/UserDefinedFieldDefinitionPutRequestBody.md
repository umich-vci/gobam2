# UserDefinedFieldDefinitionPutRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | Pointer to **string** | The name of the user defined field. | [optional] 
**DisplayName** | **string** | The display name of the user defined field. | 
**ResourceType** | **string** | The resource type that the user defined field is applied to. | 
**DataType** | Pointer to **string** | The data type of the user defined field. | [optional] 
**DefaultValue** | Pointer to **string** | The default value of the user defined field. | [optional] 
**Required** | Pointer to **bool** | Indicates whether the user defined field is required or optional for the resource. | [optional] 
**PredefinedValues** | Pointer to **[]string** |  | [optional] 
**ValidationProperties** | Pointer to [**UserDefinedFieldDefinitionValidationProperties**](UserDefinedFieldDefinitionValidationProperties.md) |  | [optional] 

## Methods

### NewUserDefinedFieldDefinitionPutRequestBody

`func NewUserDefinedFieldDefinitionPutRequestBody(displayName string, resourceType string, ) *UserDefinedFieldDefinitionPutRequestBody`

NewUserDefinedFieldDefinitionPutRequestBody instantiates a new UserDefinedFieldDefinitionPutRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserDefinedFieldDefinitionPutRequestBodyWithDefaults

`func NewUserDefinedFieldDefinitionPutRequestBodyWithDefaults() *UserDefinedFieldDefinitionPutRequestBody`

NewUserDefinedFieldDefinitionPutRequestBodyWithDefaults instantiates a new UserDefinedFieldDefinitionPutRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserDefinedFieldDefinitionPutRequestBody) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserDefinedFieldDefinitionPutRequestBody) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserDefinedFieldDefinitionPutRequestBody) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UserDefinedFieldDefinitionPutRequestBody) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *UserDefinedFieldDefinitionPutRequestBody) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserDefinedFieldDefinitionPutRequestBody) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserDefinedFieldDefinitionPutRequestBody) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UserDefinedFieldDefinitionPutRequestBody) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *UserDefinedFieldDefinitionPutRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserDefinedFieldDefinitionPutRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserDefinedFieldDefinitionPutRequestBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserDefinedFieldDefinitionPutRequestBody) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisplayName

`func (o *UserDefinedFieldDefinitionPutRequestBody) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UserDefinedFieldDefinitionPutRequestBody) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UserDefinedFieldDefinitionPutRequestBody) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetResourceType

`func (o *UserDefinedFieldDefinitionPutRequestBody) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *UserDefinedFieldDefinitionPutRequestBody) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *UserDefinedFieldDefinitionPutRequestBody) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.


### GetDataType

`func (o *UserDefinedFieldDefinitionPutRequestBody) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *UserDefinedFieldDefinitionPutRequestBody) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *UserDefinedFieldDefinitionPutRequestBody) SetDataType(v string)`

SetDataType sets DataType field to given value.

### HasDataType

`func (o *UserDefinedFieldDefinitionPutRequestBody) HasDataType() bool`

HasDataType returns a boolean if a field has been set.

### GetDefaultValue

`func (o *UserDefinedFieldDefinitionPutRequestBody) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *UserDefinedFieldDefinitionPutRequestBody) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *UserDefinedFieldDefinitionPutRequestBody) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *UserDefinedFieldDefinitionPutRequestBody) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetRequired

`func (o *UserDefinedFieldDefinitionPutRequestBody) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *UserDefinedFieldDefinitionPutRequestBody) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *UserDefinedFieldDefinitionPutRequestBody) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *UserDefinedFieldDefinitionPutRequestBody) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetPredefinedValues

`func (o *UserDefinedFieldDefinitionPutRequestBody) GetPredefinedValues() []string`

GetPredefinedValues returns the PredefinedValues field if non-nil, zero value otherwise.

### GetPredefinedValuesOk

`func (o *UserDefinedFieldDefinitionPutRequestBody) GetPredefinedValuesOk() (*[]string, bool)`

GetPredefinedValuesOk returns a tuple with the PredefinedValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredefinedValues

`func (o *UserDefinedFieldDefinitionPutRequestBody) SetPredefinedValues(v []string)`

SetPredefinedValues sets PredefinedValues field to given value.

### HasPredefinedValues

`func (o *UserDefinedFieldDefinitionPutRequestBody) HasPredefinedValues() bool`

HasPredefinedValues returns a boolean if a field has been set.

### GetValidationProperties

`func (o *UserDefinedFieldDefinitionPutRequestBody) GetValidationProperties() UserDefinedFieldDefinitionValidationProperties`

GetValidationProperties returns the ValidationProperties field if non-nil, zero value otherwise.

### GetValidationPropertiesOk

`func (o *UserDefinedFieldDefinitionPutRequestBody) GetValidationPropertiesOk() (*UserDefinedFieldDefinitionValidationProperties, bool)`

GetValidationPropertiesOk returns a tuple with the ValidationProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationProperties

`func (o *UserDefinedFieldDefinitionPutRequestBody) SetValidationProperties(v UserDefinedFieldDefinitionValidationProperties)`

SetValidationProperties sets ValidationProperties field to given value.

### HasValidationProperties

`func (o *UserDefinedFieldDefinitionPutRequestBody) HasValidationProperties() bool`

HasValidationProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


