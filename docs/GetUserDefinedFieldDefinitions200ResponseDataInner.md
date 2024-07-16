# GetUserDefinedFieldDefinitions200ResponseDataInner

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
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 

## Methods

### NewGetUserDefinedFieldDefinitions200ResponseDataInner

`func NewGetUserDefinedFieldDefinitions200ResponseDataInner() *GetUserDefinedFieldDefinitions200ResponseDataInner`

NewGetUserDefinedFieldDefinitions200ResponseDataInner instantiates a new GetUserDefinedFieldDefinitions200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserDefinedFieldDefinitions200ResponseDataInnerWithDefaults

`func NewGetUserDefinedFieldDefinitions200ResponseDataInnerWithDefaults() *GetUserDefinedFieldDefinitions200ResponseDataInner`

NewGetUserDefinedFieldDefinitions200ResponseDataInnerWithDefaults instantiates a new GetUserDefinedFieldDefinitions200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetUserDefinedFieldDefinitions200ResponseDataInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetUserDefinedFieldDefinitions200ResponseDataInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetUserDefinedFieldDefinitions200ResponseDataInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetUserDefinedFieldDefinitions200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GetUserDefinedFieldDefinitions200ResponseDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetUserDefinedFieldDefinitions200ResponseDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetUserDefinedFieldDefinitions200ResponseDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetUserDefinedFieldDefinitions200ResponseDataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *GetUserDefinedFieldDefinitions200ResponseDataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetUserDefinedFieldDefinitions200ResponseDataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetUserDefinedFieldDefinitions200ResponseDataInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetUserDefinedFieldDefinitions200ResponseDataInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisplayName

`func (o *GetUserDefinedFieldDefinitions200ResponseDataInner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GetUserDefinedFieldDefinitions200ResponseDataInner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GetUserDefinedFieldDefinitions200ResponseDataInner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *GetUserDefinedFieldDefinitions200ResponseDataInner) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetResourceType

`func (o *GetUserDefinedFieldDefinitions200ResponseDataInner) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *GetUserDefinedFieldDefinitions200ResponseDataInner) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *GetUserDefinedFieldDefinitions200ResponseDataInner) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *GetUserDefinedFieldDefinitions200ResponseDataInner) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetDataType

`func (o *GetUserDefinedFieldDefinitions200ResponseDataInner) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *GetUserDefinedFieldDefinitions200ResponseDataInner) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *GetUserDefinedFieldDefinitions200ResponseDataInner) SetDataType(v string)`

SetDataType sets DataType field to given value.

### HasDataType

`func (o *GetUserDefinedFieldDefinitions200ResponseDataInner) HasDataType() bool`

HasDataType returns a boolean if a field has been set.

### GetDefaultValue

`func (o *GetUserDefinedFieldDefinitions200ResponseDataInner) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *GetUserDefinedFieldDefinitions200ResponseDataInner) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *GetUserDefinedFieldDefinitions200ResponseDataInner) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *GetUserDefinedFieldDefinitions200ResponseDataInner) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetRequired

`func (o *GetUserDefinedFieldDefinitions200ResponseDataInner) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *GetUserDefinedFieldDefinitions200ResponseDataInner) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *GetUserDefinedFieldDefinitions200ResponseDataInner) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *GetUserDefinedFieldDefinitions200ResponseDataInner) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetPredefinedValues

`func (o *GetUserDefinedFieldDefinitions200ResponseDataInner) GetPredefinedValues() []string`

GetPredefinedValues returns the PredefinedValues field if non-nil, zero value otherwise.

### GetPredefinedValuesOk

`func (o *GetUserDefinedFieldDefinitions200ResponseDataInner) GetPredefinedValuesOk() (*[]string, bool)`

GetPredefinedValuesOk returns a tuple with the PredefinedValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredefinedValues

`func (o *GetUserDefinedFieldDefinitions200ResponseDataInner) SetPredefinedValues(v []string)`

SetPredefinedValues sets PredefinedValues field to given value.

### HasPredefinedValues

`func (o *GetUserDefinedFieldDefinitions200ResponseDataInner) HasPredefinedValues() bool`

HasPredefinedValues returns a boolean if a field has been set.

### GetValidationProperties

`func (o *GetUserDefinedFieldDefinitions200ResponseDataInner) GetValidationProperties() UserDefinedFieldDefinitionValidationProperties`

GetValidationProperties returns the ValidationProperties field if non-nil, zero value otherwise.

### GetValidationPropertiesOk

`func (o *GetUserDefinedFieldDefinitions200ResponseDataInner) GetValidationPropertiesOk() (*UserDefinedFieldDefinitionValidationProperties, bool)`

GetValidationPropertiesOk returns a tuple with the ValidationProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationProperties

`func (o *GetUserDefinedFieldDefinitions200ResponseDataInner) SetValidationProperties(v UserDefinedFieldDefinitionValidationProperties)`

SetValidationProperties sets ValidationProperties field to given value.

### HasValidationProperties

`func (o *GetUserDefinedFieldDefinitions200ResponseDataInner) HasValidationProperties() bool`

HasValidationProperties returns a boolean if a field has been set.

### GetLinks

`func (o *GetUserDefinedFieldDefinitions200ResponseDataInner) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetUserDefinedFieldDefinitions200ResponseDataInner) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetUserDefinedFieldDefinitions200ResponseDataInner) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetUserDefinedFieldDefinitions200ResponseDataInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


