# NamingPolicyTextValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The resource type. | [optional] [readonly] 
**TextType** | Pointer to **string** |  | [optional] 
**MinLength** | Pointer to **int32** | The shortest string allowed in the name. When set of 0, the text string is optional. | [optional] 
**MaxLength** | Pointer to **int32** | The longest string allowed in the name. When set of 0, the range is unbound. | [optional] 
**RegularExpression** | Pointer to **string** | The regular expression to restrict the text string. | [optional] 

## Methods

### NewNamingPolicyTextValue

`func NewNamingPolicyTextValue() *NamingPolicyTextValue`

NewNamingPolicyTextValue instantiates a new NamingPolicyTextValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNamingPolicyTextValueWithDefaults

`func NewNamingPolicyTextValueWithDefaults() *NamingPolicyTextValue`

NewNamingPolicyTextValueWithDefaults instantiates a new NamingPolicyTextValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *NamingPolicyTextValue) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NamingPolicyTextValue) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NamingPolicyTextValue) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NamingPolicyTextValue) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTextType

`func (o *NamingPolicyTextValue) GetTextType() string`

GetTextType returns the TextType field if non-nil, zero value otherwise.

### GetTextTypeOk

`func (o *NamingPolicyTextValue) GetTextTypeOk() (*string, bool)`

GetTextTypeOk returns a tuple with the TextType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextType

`func (o *NamingPolicyTextValue) SetTextType(v string)`

SetTextType sets TextType field to given value.

### HasTextType

`func (o *NamingPolicyTextValue) HasTextType() bool`

HasTextType returns a boolean if a field has been set.

### GetMinLength

`func (o *NamingPolicyTextValue) GetMinLength() int32`

GetMinLength returns the MinLength field if non-nil, zero value otherwise.

### GetMinLengthOk

`func (o *NamingPolicyTextValue) GetMinLengthOk() (*int32, bool)`

GetMinLengthOk returns a tuple with the MinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLength

`func (o *NamingPolicyTextValue) SetMinLength(v int32)`

SetMinLength sets MinLength field to given value.

### HasMinLength

`func (o *NamingPolicyTextValue) HasMinLength() bool`

HasMinLength returns a boolean if a field has been set.

### GetMaxLength

`func (o *NamingPolicyTextValue) GetMaxLength() int32`

GetMaxLength returns the MaxLength field if non-nil, zero value otherwise.

### GetMaxLengthOk

`func (o *NamingPolicyTextValue) GetMaxLengthOk() (*int32, bool)`

GetMaxLengthOk returns a tuple with the MaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLength

`func (o *NamingPolicyTextValue) SetMaxLength(v int32)`

SetMaxLength sets MaxLength field to given value.

### HasMaxLength

`func (o *NamingPolicyTextValue) HasMaxLength() bool`

HasMaxLength returns a boolean if a field has been set.

### GetRegularExpression

`func (o *NamingPolicyTextValue) GetRegularExpression() string`

GetRegularExpression returns the RegularExpression field if non-nil, zero value otherwise.

### GetRegularExpressionOk

`func (o *NamingPolicyTextValue) GetRegularExpressionOk() (*string, bool)`

GetRegularExpressionOk returns a tuple with the RegularExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegularExpression

`func (o *NamingPolicyTextValue) SetRegularExpression(v string)`

SetRegularExpression sets RegularExpression field to given value.

### HasRegularExpression

`func (o *NamingPolicyTextValue) HasRegularExpression() bool`

HasRegularExpression returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


