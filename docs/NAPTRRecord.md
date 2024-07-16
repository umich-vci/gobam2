# NAPTRRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The resource type. | [optional] 
**Order** | Pointer to **int32** | The order in which NAPTR records are to be processed. | [optional] 
**Preference** | Pointer to **int32** | The preference assigned to the NAPTR record. The preference is referenced when NAPTR records have the same order. | [optional] 
**Service** | Pointer to **string** | The service or protocol to which the NAPTR record resolves. | [optional] 
**RegularExpression** | Pointer to **string** | The regular expression used to transform client data. | [optional] 
**Replacement** | Pointer to **string** | The replacement string used to replace client data. | [optional] 
**Flags** | Pointer to **string** | The control flags in the NAPTR record. Control flags influence how the replacement functions are performed when processing the record. | [optional] 

## Methods

### NewNAPTRRecord

`func NewNAPTRRecord() *NAPTRRecord`

NewNAPTRRecord instantiates a new NAPTRRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNAPTRRecordWithDefaults

`func NewNAPTRRecordWithDefaults() *NAPTRRecord`

NewNAPTRRecordWithDefaults instantiates a new NAPTRRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *NAPTRRecord) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NAPTRRecord) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NAPTRRecord) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NAPTRRecord) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOrder

`func (o *NAPTRRecord) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *NAPTRRecord) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *NAPTRRecord) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *NAPTRRecord) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetPreference

`func (o *NAPTRRecord) GetPreference() int32`

GetPreference returns the Preference field if non-nil, zero value otherwise.

### GetPreferenceOk

`func (o *NAPTRRecord) GetPreferenceOk() (*int32, bool)`

GetPreferenceOk returns a tuple with the Preference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreference

`func (o *NAPTRRecord) SetPreference(v int32)`

SetPreference sets Preference field to given value.

### HasPreference

`func (o *NAPTRRecord) HasPreference() bool`

HasPreference returns a boolean if a field has been set.

### GetService

`func (o *NAPTRRecord) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *NAPTRRecord) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *NAPTRRecord) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *NAPTRRecord) HasService() bool`

HasService returns a boolean if a field has been set.

### GetRegularExpression

`func (o *NAPTRRecord) GetRegularExpression() string`

GetRegularExpression returns the RegularExpression field if non-nil, zero value otherwise.

### GetRegularExpressionOk

`func (o *NAPTRRecord) GetRegularExpressionOk() (*string, bool)`

GetRegularExpressionOk returns a tuple with the RegularExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegularExpression

`func (o *NAPTRRecord) SetRegularExpression(v string)`

SetRegularExpression sets RegularExpression field to given value.

### HasRegularExpression

`func (o *NAPTRRecord) HasRegularExpression() bool`

HasRegularExpression returns a boolean if a field has been set.

### GetReplacement

`func (o *NAPTRRecord) GetReplacement() string`

GetReplacement returns the Replacement field if non-nil, zero value otherwise.

### GetReplacementOk

`func (o *NAPTRRecord) GetReplacementOk() (*string, bool)`

GetReplacementOk returns a tuple with the Replacement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacement

`func (o *NAPTRRecord) SetReplacement(v string)`

SetReplacement sets Replacement field to given value.

### HasReplacement

`func (o *NAPTRRecord) HasReplacement() bool`

HasReplacement returns a boolean if a field has been set.

### GetFlags

`func (o *NAPTRRecord) GetFlags() string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *NAPTRRecord) GetFlagsOk() (*string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *NAPTRRecord) SetFlags(v string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *NAPTRRecord) HasFlags() bool`

HasFlags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


