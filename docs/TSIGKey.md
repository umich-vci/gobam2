# TSIGKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | Pointer to **string** | The name of the resource. | [optional] 
**Length** | Pointer to **int32** | The key length, in bits. | [optional] 
**Algorithm** | Pointer to **string** | The algorithm used to sign the key. | [optional] 

## Methods

### NewTSIGKey

`func NewTSIGKey() *TSIGKey`

NewTSIGKey instantiates a new TSIGKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTSIGKeyWithDefaults

`func NewTSIGKeyWithDefaults() *TSIGKey`

NewTSIGKeyWithDefaults instantiates a new TSIGKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TSIGKey) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TSIGKey) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TSIGKey) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TSIGKey) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *TSIGKey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TSIGKey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TSIGKey) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TSIGKey) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLength

`func (o *TSIGKey) GetLength() int32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *TSIGKey) GetLengthOk() (*int32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *TSIGKey) SetLength(v int32)`

SetLength sets Length field to given value.

### HasLength

`func (o *TSIGKey) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetAlgorithm

`func (o *TSIGKey) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *TSIGKey) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *TSIGKey) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.

### HasAlgorithm

`func (o *TSIGKey) HasAlgorithm() bool`

HasAlgorithm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


