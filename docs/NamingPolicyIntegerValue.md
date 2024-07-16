# NamingPolicyIntegerValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The resource type. | [optional] [readonly] 
**Min** | Pointer to **int32** | The start value for the numeric range. | [optional] 
**Max** | Pointer to **int32** | The end value for the numeric range. A value of 0 indicates that the range is unbounded. | [optional] 

## Methods

### NewNamingPolicyIntegerValue

`func NewNamingPolicyIntegerValue() *NamingPolicyIntegerValue`

NewNamingPolicyIntegerValue instantiates a new NamingPolicyIntegerValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNamingPolicyIntegerValueWithDefaults

`func NewNamingPolicyIntegerValueWithDefaults() *NamingPolicyIntegerValue`

NewNamingPolicyIntegerValueWithDefaults instantiates a new NamingPolicyIntegerValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *NamingPolicyIntegerValue) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NamingPolicyIntegerValue) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NamingPolicyIntegerValue) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NamingPolicyIntegerValue) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMin

`func (o *NamingPolicyIntegerValue) GetMin() int32`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *NamingPolicyIntegerValue) GetMinOk() (*int32, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *NamingPolicyIntegerValue) SetMin(v int32)`

SetMin sets Min field to given value.

### HasMin

`func (o *NamingPolicyIntegerValue) HasMin() bool`

HasMin returns a boolean if a field has been set.

### GetMax

`func (o *NamingPolicyIntegerValue) GetMax() int32`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *NamingPolicyIntegerValue) GetMaxOk() (*int32, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *NamingPolicyIntegerValue) SetMax(v int32)`

SetMax sets Max field to given value.

### HasMax

`func (o *NamingPolicyIntegerValue) HasMax() bool`

HasMax returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


