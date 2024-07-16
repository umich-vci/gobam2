# UserDefinedFieldDefinitionValidationProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Min** | Pointer to **int64** | The minimum value | [optional] 
**Max** | Pointer to **int64** | The maximum value | [optional] 
**MinLength** | Pointer to **int32** | The minimum length | [optional] 
**MaxLength** | Pointer to **int32** | The maximum length | [optional] 
**Pattern** | Pointer to **string** | The pattern to match against. The wildcard characters supported are &#39;*&#39; and &#39;?&#39;. | [optional] 

## Methods

### NewUserDefinedFieldDefinitionValidationProperties

`func NewUserDefinedFieldDefinitionValidationProperties() *UserDefinedFieldDefinitionValidationProperties`

NewUserDefinedFieldDefinitionValidationProperties instantiates a new UserDefinedFieldDefinitionValidationProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserDefinedFieldDefinitionValidationPropertiesWithDefaults

`func NewUserDefinedFieldDefinitionValidationPropertiesWithDefaults() *UserDefinedFieldDefinitionValidationProperties`

NewUserDefinedFieldDefinitionValidationPropertiesWithDefaults instantiates a new UserDefinedFieldDefinitionValidationProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMin

`func (o *UserDefinedFieldDefinitionValidationProperties) GetMin() int64`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *UserDefinedFieldDefinitionValidationProperties) GetMinOk() (*int64, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *UserDefinedFieldDefinitionValidationProperties) SetMin(v int64)`

SetMin sets Min field to given value.

### HasMin

`func (o *UserDefinedFieldDefinitionValidationProperties) HasMin() bool`

HasMin returns a boolean if a field has been set.

### GetMax

`func (o *UserDefinedFieldDefinitionValidationProperties) GetMax() int64`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *UserDefinedFieldDefinitionValidationProperties) GetMaxOk() (*int64, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *UserDefinedFieldDefinitionValidationProperties) SetMax(v int64)`

SetMax sets Max field to given value.

### HasMax

`func (o *UserDefinedFieldDefinitionValidationProperties) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetMinLength

`func (o *UserDefinedFieldDefinitionValidationProperties) GetMinLength() int32`

GetMinLength returns the MinLength field if non-nil, zero value otherwise.

### GetMinLengthOk

`func (o *UserDefinedFieldDefinitionValidationProperties) GetMinLengthOk() (*int32, bool)`

GetMinLengthOk returns a tuple with the MinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLength

`func (o *UserDefinedFieldDefinitionValidationProperties) SetMinLength(v int32)`

SetMinLength sets MinLength field to given value.

### HasMinLength

`func (o *UserDefinedFieldDefinitionValidationProperties) HasMinLength() bool`

HasMinLength returns a boolean if a field has been set.

### GetMaxLength

`func (o *UserDefinedFieldDefinitionValidationProperties) GetMaxLength() int32`

GetMaxLength returns the MaxLength field if non-nil, zero value otherwise.

### GetMaxLengthOk

`func (o *UserDefinedFieldDefinitionValidationProperties) GetMaxLengthOk() (*int32, bool)`

GetMaxLengthOk returns a tuple with the MaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLength

`func (o *UserDefinedFieldDefinitionValidationProperties) SetMaxLength(v int32)`

SetMaxLength sets MaxLength field to given value.

### HasMaxLength

`func (o *UserDefinedFieldDefinitionValidationProperties) HasMaxLength() bool`

HasMaxLength returns a boolean if a field has been set.

### GetPattern

`func (o *UserDefinedFieldDefinitionValidationProperties) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *UserDefinedFieldDefinitionValidationProperties) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *UserDefinedFieldDefinitionValidationProperties) SetPattern(v string)`

SetPattern sets Pattern field to given value.

### HasPattern

`func (o *UserDefinedFieldDefinitionValidationProperties) HasPattern() bool`

HasPattern returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


