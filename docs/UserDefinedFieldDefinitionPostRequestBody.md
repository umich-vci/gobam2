# UserDefinedFieldDefinitionPostRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | **string** | The name of the user defined field. | 
**DisplayName** | **string** | The display name of the user defined field. | 
**ResourceType** | **string** | The resource type that the user defined field is applied to. | 
**DataType** | **string** | The data type of the user defined field. | 
**DefaultValue** | Pointer to **string** | The default value of the user defined field. | [optional] 
**Required** | Pointer to **bool** | Indicates whether the user defined field is required or optional for the resource. | [optional] 
**PredefinedValues** | Pointer to **[]string** |  | [optional] 
**ValidationProperties** | Pointer to [**UserDefinedFieldDefinitionValidationProperties**](UserDefinedFieldDefinitionValidationProperties.md) |  | [optional] 

## Methods

### NewUserDefinedFieldDefinitionPostRequestBody

`func NewUserDefinedFieldDefinitionPostRequestBody(name string, displayName string, resourceType string, dataType string, ) *UserDefinedFieldDefinitionPostRequestBody`

NewUserDefinedFieldDefinitionPostRequestBody instantiates a new UserDefinedFieldDefinitionPostRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserDefinedFieldDefinitionPostRequestBodyWithDefaults

`func NewUserDefinedFieldDefinitionPostRequestBodyWithDefaults() *UserDefinedFieldDefinitionPostRequestBody`

NewUserDefinedFieldDefinitionPostRequestBodyWithDefaults instantiates a new UserDefinedFieldDefinitionPostRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserDefinedFieldDefinitionPostRequestBody) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserDefinedFieldDefinitionPostRequestBody) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserDefinedFieldDefinitionPostRequestBody) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UserDefinedFieldDefinitionPostRequestBody) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *UserDefinedFieldDefinitionPostRequestBody) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserDefinedFieldDefinitionPostRequestBody) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserDefinedFieldDefinitionPostRequestBody) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UserDefinedFieldDefinitionPostRequestBody) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *UserDefinedFieldDefinitionPostRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserDefinedFieldDefinitionPostRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserDefinedFieldDefinitionPostRequestBody) SetName(v string)`

SetName sets Name field to given value.


### GetDisplayName

`func (o *UserDefinedFieldDefinitionPostRequestBody) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UserDefinedFieldDefinitionPostRequestBody) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UserDefinedFieldDefinitionPostRequestBody) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetResourceType

`func (o *UserDefinedFieldDefinitionPostRequestBody) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *UserDefinedFieldDefinitionPostRequestBody) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *UserDefinedFieldDefinitionPostRequestBody) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.


### GetDataType

`func (o *UserDefinedFieldDefinitionPostRequestBody) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *UserDefinedFieldDefinitionPostRequestBody) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *UserDefinedFieldDefinitionPostRequestBody) SetDataType(v string)`

SetDataType sets DataType field to given value.


### GetDefaultValue

`func (o *UserDefinedFieldDefinitionPostRequestBody) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *UserDefinedFieldDefinitionPostRequestBody) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *UserDefinedFieldDefinitionPostRequestBody) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *UserDefinedFieldDefinitionPostRequestBody) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetRequired

`func (o *UserDefinedFieldDefinitionPostRequestBody) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *UserDefinedFieldDefinitionPostRequestBody) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *UserDefinedFieldDefinitionPostRequestBody) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *UserDefinedFieldDefinitionPostRequestBody) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetPredefinedValues

`func (o *UserDefinedFieldDefinitionPostRequestBody) GetPredefinedValues() []string`

GetPredefinedValues returns the PredefinedValues field if non-nil, zero value otherwise.

### GetPredefinedValuesOk

`func (o *UserDefinedFieldDefinitionPostRequestBody) GetPredefinedValuesOk() (*[]string, bool)`

GetPredefinedValuesOk returns a tuple with the PredefinedValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredefinedValues

`func (o *UserDefinedFieldDefinitionPostRequestBody) SetPredefinedValues(v []string)`

SetPredefinedValues sets PredefinedValues field to given value.

### HasPredefinedValues

`func (o *UserDefinedFieldDefinitionPostRequestBody) HasPredefinedValues() bool`

HasPredefinedValues returns a boolean if a field has been set.

### GetValidationProperties

`func (o *UserDefinedFieldDefinitionPostRequestBody) GetValidationProperties() UserDefinedFieldDefinitionValidationProperties`

GetValidationProperties returns the ValidationProperties field if non-nil, zero value otherwise.

### GetValidationPropertiesOk

`func (o *UserDefinedFieldDefinitionPostRequestBody) GetValidationPropertiesOk() (*UserDefinedFieldDefinitionValidationProperties, bool)`

GetValidationPropertiesOk returns a tuple with the ValidationProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationProperties

`func (o *UserDefinedFieldDefinitionPostRequestBody) SetValidationProperties(v UserDefinedFieldDefinitionValidationProperties)`

SetValidationProperties sets ValidationProperties field to given value.

### HasValidationProperties

`func (o *UserDefinedFieldDefinitionPostRequestBody) HasValidationProperties() bool`

HasValidationProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


