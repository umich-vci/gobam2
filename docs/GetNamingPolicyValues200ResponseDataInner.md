# GetNamingPolicyValues200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the naming policy value. | [optional] 
**Connector** | Pointer to **string** | A character used to separate elements in a filename; typically, - (hyphen) or _ (underscore) are used. | [optional] 
**IncrementalRole** | Pointer to **string** | The incremental counter type. The value can be Counter to make the value a sequential counter or Unique Name to use the value to ensure that the names are unique. When you select Counter, the value starts at a specified value and increments each time the naming policy creates a name. When you select Unique Name, the value increments only to ensure that generated names are unique. | [optional] 
**IncrementType** | Pointer to **string** | The number system for the incremental value. | [optional] 
**PaddingType** | Pointer to **string** | The padding that is added to the incremental value. Simple padding pads the incremental value with a fixed number of leading zeros. Global padding pads the incremental value with leading zeros to make the entire name generated by the policy a specific length. | [optional] 
**SequenceStart** | Pointer to **int32** | The starting value for the incremental value. | [optional] 
**SequenceIncrement** | Pointer to **int32** | The amount by which to increment the value each time it&#39;s used. | [optional] 
**PaddingLength** | Pointer to **int32** | The length of the padding. For simple padding, the value determines how many digits are used for the incremental value. For global padding, the value determines the overall length of the name generated by the naming policy. | [optional] 
**MissingValueReuseEnabled** | Pointer to **bool** | Determines whether the naming policy reuses numeric values if they&#39;re available. | [optional] 
**Min** | Pointer to **int32** | The start value for the numeric range. | [optional] 
**Max** | Pointer to **int32** | The end value for the numeric range. A value of 0 indicates that the range is unbounded. | [optional] 
**Elements** | Pointer to [**[]NameValuePairBean**](NameValuePairBean.md) |  | [optional] 
**TextType** | Pointer to **string** |  | [optional] 
**MinLength** | Pointer to **int32** | The shortest string allowed in the name. When set of 0, the text string is optional. | [optional] 
**MaxLength** | Pointer to **int32** | The longest string allowed in the name. When set of 0, the range is unbound. | [optional] 
**RegularExpression** | Pointer to **string** | The regular expression to restrict the text string. | [optional] 
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 

## Methods

### NewGetNamingPolicyValues200ResponseDataInner

`func NewGetNamingPolicyValues200ResponseDataInner() *GetNamingPolicyValues200ResponseDataInner`

NewGetNamingPolicyValues200ResponseDataInner instantiates a new GetNamingPolicyValues200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNamingPolicyValues200ResponseDataInnerWithDefaults

`func NewGetNamingPolicyValues200ResponseDataInnerWithDefaults() *GetNamingPolicyValues200ResponseDataInner`

NewGetNamingPolicyValues200ResponseDataInnerWithDefaults instantiates a new GetNamingPolicyValues200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetNamingPolicyValues200ResponseDataInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetNamingPolicyValues200ResponseDataInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetNamingPolicyValues200ResponseDataInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetNamingPolicyValues200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GetNamingPolicyValues200ResponseDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetNamingPolicyValues200ResponseDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetNamingPolicyValues200ResponseDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetNamingPolicyValues200ResponseDataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *GetNamingPolicyValues200ResponseDataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetNamingPolicyValues200ResponseDataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetNamingPolicyValues200ResponseDataInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetNamingPolicyValues200ResponseDataInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetConnector

`func (o *GetNamingPolicyValues200ResponseDataInner) GetConnector() string`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *GetNamingPolicyValues200ResponseDataInner) GetConnectorOk() (*string, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *GetNamingPolicyValues200ResponseDataInner) SetConnector(v string)`

SetConnector sets Connector field to given value.

### HasConnector

`func (o *GetNamingPolicyValues200ResponseDataInner) HasConnector() bool`

HasConnector returns a boolean if a field has been set.

### GetIncrementalRole

`func (o *GetNamingPolicyValues200ResponseDataInner) GetIncrementalRole() string`

GetIncrementalRole returns the IncrementalRole field if non-nil, zero value otherwise.

### GetIncrementalRoleOk

`func (o *GetNamingPolicyValues200ResponseDataInner) GetIncrementalRoleOk() (*string, bool)`

GetIncrementalRoleOk returns a tuple with the IncrementalRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncrementalRole

`func (o *GetNamingPolicyValues200ResponseDataInner) SetIncrementalRole(v string)`

SetIncrementalRole sets IncrementalRole field to given value.

### HasIncrementalRole

`func (o *GetNamingPolicyValues200ResponseDataInner) HasIncrementalRole() bool`

HasIncrementalRole returns a boolean if a field has been set.

### GetIncrementType

`func (o *GetNamingPolicyValues200ResponseDataInner) GetIncrementType() string`

GetIncrementType returns the IncrementType field if non-nil, zero value otherwise.

### GetIncrementTypeOk

`func (o *GetNamingPolicyValues200ResponseDataInner) GetIncrementTypeOk() (*string, bool)`

GetIncrementTypeOk returns a tuple with the IncrementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncrementType

`func (o *GetNamingPolicyValues200ResponseDataInner) SetIncrementType(v string)`

SetIncrementType sets IncrementType field to given value.

### HasIncrementType

`func (o *GetNamingPolicyValues200ResponseDataInner) HasIncrementType() bool`

HasIncrementType returns a boolean if a field has been set.

### GetPaddingType

`func (o *GetNamingPolicyValues200ResponseDataInner) GetPaddingType() string`

GetPaddingType returns the PaddingType field if non-nil, zero value otherwise.

### GetPaddingTypeOk

`func (o *GetNamingPolicyValues200ResponseDataInner) GetPaddingTypeOk() (*string, bool)`

GetPaddingTypeOk returns a tuple with the PaddingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaddingType

`func (o *GetNamingPolicyValues200ResponseDataInner) SetPaddingType(v string)`

SetPaddingType sets PaddingType field to given value.

### HasPaddingType

`func (o *GetNamingPolicyValues200ResponseDataInner) HasPaddingType() bool`

HasPaddingType returns a boolean if a field has been set.

### GetSequenceStart

`func (o *GetNamingPolicyValues200ResponseDataInner) GetSequenceStart() int32`

GetSequenceStart returns the SequenceStart field if non-nil, zero value otherwise.

### GetSequenceStartOk

`func (o *GetNamingPolicyValues200ResponseDataInner) GetSequenceStartOk() (*int32, bool)`

GetSequenceStartOk returns a tuple with the SequenceStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceStart

`func (o *GetNamingPolicyValues200ResponseDataInner) SetSequenceStart(v int32)`

SetSequenceStart sets SequenceStart field to given value.

### HasSequenceStart

`func (o *GetNamingPolicyValues200ResponseDataInner) HasSequenceStart() bool`

HasSequenceStart returns a boolean if a field has been set.

### GetSequenceIncrement

`func (o *GetNamingPolicyValues200ResponseDataInner) GetSequenceIncrement() int32`

GetSequenceIncrement returns the SequenceIncrement field if non-nil, zero value otherwise.

### GetSequenceIncrementOk

`func (o *GetNamingPolicyValues200ResponseDataInner) GetSequenceIncrementOk() (*int32, bool)`

GetSequenceIncrementOk returns a tuple with the SequenceIncrement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceIncrement

`func (o *GetNamingPolicyValues200ResponseDataInner) SetSequenceIncrement(v int32)`

SetSequenceIncrement sets SequenceIncrement field to given value.

### HasSequenceIncrement

`func (o *GetNamingPolicyValues200ResponseDataInner) HasSequenceIncrement() bool`

HasSequenceIncrement returns a boolean if a field has been set.

### GetPaddingLength

`func (o *GetNamingPolicyValues200ResponseDataInner) GetPaddingLength() int32`

GetPaddingLength returns the PaddingLength field if non-nil, zero value otherwise.

### GetPaddingLengthOk

`func (o *GetNamingPolicyValues200ResponseDataInner) GetPaddingLengthOk() (*int32, bool)`

GetPaddingLengthOk returns a tuple with the PaddingLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaddingLength

`func (o *GetNamingPolicyValues200ResponseDataInner) SetPaddingLength(v int32)`

SetPaddingLength sets PaddingLength field to given value.

### HasPaddingLength

`func (o *GetNamingPolicyValues200ResponseDataInner) HasPaddingLength() bool`

HasPaddingLength returns a boolean if a field has been set.

### GetMissingValueReuseEnabled

`func (o *GetNamingPolicyValues200ResponseDataInner) GetMissingValueReuseEnabled() bool`

GetMissingValueReuseEnabled returns the MissingValueReuseEnabled field if non-nil, zero value otherwise.

### GetMissingValueReuseEnabledOk

`func (o *GetNamingPolicyValues200ResponseDataInner) GetMissingValueReuseEnabledOk() (*bool, bool)`

GetMissingValueReuseEnabledOk returns a tuple with the MissingValueReuseEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingValueReuseEnabled

`func (o *GetNamingPolicyValues200ResponseDataInner) SetMissingValueReuseEnabled(v bool)`

SetMissingValueReuseEnabled sets MissingValueReuseEnabled field to given value.

### HasMissingValueReuseEnabled

`func (o *GetNamingPolicyValues200ResponseDataInner) HasMissingValueReuseEnabled() bool`

HasMissingValueReuseEnabled returns a boolean if a field has been set.

### GetMin

`func (o *GetNamingPolicyValues200ResponseDataInner) GetMin() int32`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *GetNamingPolicyValues200ResponseDataInner) GetMinOk() (*int32, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *GetNamingPolicyValues200ResponseDataInner) SetMin(v int32)`

SetMin sets Min field to given value.

### HasMin

`func (o *GetNamingPolicyValues200ResponseDataInner) HasMin() bool`

HasMin returns a boolean if a field has been set.

### GetMax

`func (o *GetNamingPolicyValues200ResponseDataInner) GetMax() int32`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *GetNamingPolicyValues200ResponseDataInner) GetMaxOk() (*int32, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *GetNamingPolicyValues200ResponseDataInner) SetMax(v int32)`

SetMax sets Max field to given value.

### HasMax

`func (o *GetNamingPolicyValues200ResponseDataInner) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetElements

`func (o *GetNamingPolicyValues200ResponseDataInner) GetElements() []NameValuePairBean`

GetElements returns the Elements field if non-nil, zero value otherwise.

### GetElementsOk

`func (o *GetNamingPolicyValues200ResponseDataInner) GetElementsOk() (*[]NameValuePairBean, bool)`

GetElementsOk returns a tuple with the Elements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElements

`func (o *GetNamingPolicyValues200ResponseDataInner) SetElements(v []NameValuePairBean)`

SetElements sets Elements field to given value.

### HasElements

`func (o *GetNamingPolicyValues200ResponseDataInner) HasElements() bool`

HasElements returns a boolean if a field has been set.

### GetTextType

`func (o *GetNamingPolicyValues200ResponseDataInner) GetTextType() string`

GetTextType returns the TextType field if non-nil, zero value otherwise.

### GetTextTypeOk

`func (o *GetNamingPolicyValues200ResponseDataInner) GetTextTypeOk() (*string, bool)`

GetTextTypeOk returns a tuple with the TextType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextType

`func (o *GetNamingPolicyValues200ResponseDataInner) SetTextType(v string)`

SetTextType sets TextType field to given value.

### HasTextType

`func (o *GetNamingPolicyValues200ResponseDataInner) HasTextType() bool`

HasTextType returns a boolean if a field has been set.

### GetMinLength

`func (o *GetNamingPolicyValues200ResponseDataInner) GetMinLength() int32`

GetMinLength returns the MinLength field if non-nil, zero value otherwise.

### GetMinLengthOk

`func (o *GetNamingPolicyValues200ResponseDataInner) GetMinLengthOk() (*int32, bool)`

GetMinLengthOk returns a tuple with the MinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLength

`func (o *GetNamingPolicyValues200ResponseDataInner) SetMinLength(v int32)`

SetMinLength sets MinLength field to given value.

### HasMinLength

`func (o *GetNamingPolicyValues200ResponseDataInner) HasMinLength() bool`

HasMinLength returns a boolean if a field has been set.

### GetMaxLength

`func (o *GetNamingPolicyValues200ResponseDataInner) GetMaxLength() int32`

GetMaxLength returns the MaxLength field if non-nil, zero value otherwise.

### GetMaxLengthOk

`func (o *GetNamingPolicyValues200ResponseDataInner) GetMaxLengthOk() (*int32, bool)`

GetMaxLengthOk returns a tuple with the MaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLength

`func (o *GetNamingPolicyValues200ResponseDataInner) SetMaxLength(v int32)`

SetMaxLength sets MaxLength field to given value.

### HasMaxLength

`func (o *GetNamingPolicyValues200ResponseDataInner) HasMaxLength() bool`

HasMaxLength returns a boolean if a field has been set.

### GetRegularExpression

`func (o *GetNamingPolicyValues200ResponseDataInner) GetRegularExpression() string`

GetRegularExpression returns the RegularExpression field if non-nil, zero value otherwise.

### GetRegularExpressionOk

`func (o *GetNamingPolicyValues200ResponseDataInner) GetRegularExpressionOk() (*string, bool)`

GetRegularExpressionOk returns a tuple with the RegularExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegularExpression

`func (o *GetNamingPolicyValues200ResponseDataInner) SetRegularExpression(v string)`

SetRegularExpression sets RegularExpression field to given value.

### HasRegularExpression

`func (o *GetNamingPolicyValues200ResponseDataInner) HasRegularExpression() bool`

HasRegularExpression returns a boolean if a field has been set.

### GetLinks

`func (o *GetNamingPolicyValues200ResponseDataInner) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetNamingPolicyValues200ResponseDataInner) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetNamingPolicyValues200ResponseDataInner) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetNamingPolicyValues200ResponseDataInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

