# NamingPolicyIncrementalValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The resource type. | [optional] [readonly] 
**IncrementalRole** | Pointer to **string** | The incremental counter type. The value can be Counter to make the value a sequential counter or Unique Name to use the value to ensure that the names are unique. When you select Counter, the value starts at a specified value and increments each time the naming policy creates a name. When you select Unique Name, the value increments only to ensure that generated names are unique. | [optional] 
**IncrementType** | Pointer to **string** | The number system for the incremental value. | [optional] 
**PaddingType** | Pointer to **string** | The padding that is added to the incremental value. Simple padding pads the incremental value with a fixed number of leading zeros. Global padding pads the incremental value with leading zeros to make the entire name generated by the policy a specific length. | [optional] 
**SequenceStart** | Pointer to **int32** | The starting value for the incremental value. | [optional] 
**SequenceIncrement** | Pointer to **int32** | The amount by which to increment the value each time it&#39;s used. | [optional] 
**PaddingLength** | Pointer to **int32** | The length of the padding. For simple padding, the value determines how many digits are used for the incremental value. For global padding, the value determines the overall length of the name generated by the naming policy. | [optional] 
**MissingValueReuseEnabled** | Pointer to **bool** | Determines whether the naming policy reuses numeric values if they&#39;re available. | [optional] 

## Methods

### NewNamingPolicyIncrementalValue

`func NewNamingPolicyIncrementalValue() *NamingPolicyIncrementalValue`

NewNamingPolicyIncrementalValue instantiates a new NamingPolicyIncrementalValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNamingPolicyIncrementalValueWithDefaults

`func NewNamingPolicyIncrementalValueWithDefaults() *NamingPolicyIncrementalValue`

NewNamingPolicyIncrementalValueWithDefaults instantiates a new NamingPolicyIncrementalValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *NamingPolicyIncrementalValue) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NamingPolicyIncrementalValue) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NamingPolicyIncrementalValue) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NamingPolicyIncrementalValue) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIncrementalRole

`func (o *NamingPolicyIncrementalValue) GetIncrementalRole() string`

GetIncrementalRole returns the IncrementalRole field if non-nil, zero value otherwise.

### GetIncrementalRoleOk

`func (o *NamingPolicyIncrementalValue) GetIncrementalRoleOk() (*string, bool)`

GetIncrementalRoleOk returns a tuple with the IncrementalRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncrementalRole

`func (o *NamingPolicyIncrementalValue) SetIncrementalRole(v string)`

SetIncrementalRole sets IncrementalRole field to given value.

### HasIncrementalRole

`func (o *NamingPolicyIncrementalValue) HasIncrementalRole() bool`

HasIncrementalRole returns a boolean if a field has been set.

### GetIncrementType

`func (o *NamingPolicyIncrementalValue) GetIncrementType() string`

GetIncrementType returns the IncrementType field if non-nil, zero value otherwise.

### GetIncrementTypeOk

`func (o *NamingPolicyIncrementalValue) GetIncrementTypeOk() (*string, bool)`

GetIncrementTypeOk returns a tuple with the IncrementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncrementType

`func (o *NamingPolicyIncrementalValue) SetIncrementType(v string)`

SetIncrementType sets IncrementType field to given value.

### HasIncrementType

`func (o *NamingPolicyIncrementalValue) HasIncrementType() bool`

HasIncrementType returns a boolean if a field has been set.

### GetPaddingType

`func (o *NamingPolicyIncrementalValue) GetPaddingType() string`

GetPaddingType returns the PaddingType field if non-nil, zero value otherwise.

### GetPaddingTypeOk

`func (o *NamingPolicyIncrementalValue) GetPaddingTypeOk() (*string, bool)`

GetPaddingTypeOk returns a tuple with the PaddingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaddingType

`func (o *NamingPolicyIncrementalValue) SetPaddingType(v string)`

SetPaddingType sets PaddingType field to given value.

### HasPaddingType

`func (o *NamingPolicyIncrementalValue) HasPaddingType() bool`

HasPaddingType returns a boolean if a field has been set.

### GetSequenceStart

`func (o *NamingPolicyIncrementalValue) GetSequenceStart() int32`

GetSequenceStart returns the SequenceStart field if non-nil, zero value otherwise.

### GetSequenceStartOk

`func (o *NamingPolicyIncrementalValue) GetSequenceStartOk() (*int32, bool)`

GetSequenceStartOk returns a tuple with the SequenceStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceStart

`func (o *NamingPolicyIncrementalValue) SetSequenceStart(v int32)`

SetSequenceStart sets SequenceStart field to given value.

### HasSequenceStart

`func (o *NamingPolicyIncrementalValue) HasSequenceStart() bool`

HasSequenceStart returns a boolean if a field has been set.

### GetSequenceIncrement

`func (o *NamingPolicyIncrementalValue) GetSequenceIncrement() int32`

GetSequenceIncrement returns the SequenceIncrement field if non-nil, zero value otherwise.

### GetSequenceIncrementOk

`func (o *NamingPolicyIncrementalValue) GetSequenceIncrementOk() (*int32, bool)`

GetSequenceIncrementOk returns a tuple with the SequenceIncrement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceIncrement

`func (o *NamingPolicyIncrementalValue) SetSequenceIncrement(v int32)`

SetSequenceIncrement sets SequenceIncrement field to given value.

### HasSequenceIncrement

`func (o *NamingPolicyIncrementalValue) HasSequenceIncrement() bool`

HasSequenceIncrement returns a boolean if a field has been set.

### GetPaddingLength

`func (o *NamingPolicyIncrementalValue) GetPaddingLength() int32`

GetPaddingLength returns the PaddingLength field if non-nil, zero value otherwise.

### GetPaddingLengthOk

`func (o *NamingPolicyIncrementalValue) GetPaddingLengthOk() (*int32, bool)`

GetPaddingLengthOk returns a tuple with the PaddingLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaddingLength

`func (o *NamingPolicyIncrementalValue) SetPaddingLength(v int32)`

SetPaddingLength sets PaddingLength field to given value.

### HasPaddingLength

`func (o *NamingPolicyIncrementalValue) HasPaddingLength() bool`

HasPaddingLength returns a boolean if a field has been set.

### GetMissingValueReuseEnabled

`func (o *NamingPolicyIncrementalValue) GetMissingValueReuseEnabled() bool`

GetMissingValueReuseEnabled returns the MissingValueReuseEnabled field if non-nil, zero value otherwise.

### GetMissingValueReuseEnabledOk

`func (o *NamingPolicyIncrementalValue) GetMissingValueReuseEnabledOk() (*bool, bool)`

GetMissingValueReuseEnabledOk returns a tuple with the MissingValueReuseEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingValueReuseEnabled

`func (o *NamingPolicyIncrementalValue) SetMissingValueReuseEnabled(v bool)`

SetMissingValueReuseEnabled sets MissingValueReuseEnabled field to given value.

### HasMissingValueReuseEnabled

`func (o *NamingPolicyIncrementalValue) HasMissingValueReuseEnabled() bool`

HasMissingValueReuseEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


