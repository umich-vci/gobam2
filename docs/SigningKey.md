# SigningKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | Pointer to **string** | The name of the resource. | [optional] 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**Configuration** | Pointer to [**InlinedConfiguration**](InlinedConfiguration.md) |  | [optional] [readonly] 
**Length** | Pointer to **int32** | The key length, in bits. | [optional] 
**State** | Pointer to **string** | The state of the key. | [optional] 
**KeyGenerationDateTime** | Pointer to **time.Time** | The date and time that they key was generated. | [optional] [readonly] 
**PrivateKey** | Pointer to **string** | The private key. | [optional] 

## Methods

### NewSigningKey

`func NewSigningKey() *SigningKey`

NewSigningKey instantiates a new SigningKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSigningKeyWithDefaults

`func NewSigningKeyWithDefaults() *SigningKey`

NewSigningKeyWithDefaults instantiates a new SigningKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SigningKey) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SigningKey) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SigningKey) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *SigningKey) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *SigningKey) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SigningKey) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SigningKey) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SigningKey) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *SigningKey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SigningKey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SigningKey) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SigningKey) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *SigningKey) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *SigningKey) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *SigningKey) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *SigningKey) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetConfiguration

`func (o *SigningKey) GetConfiguration() InlinedConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *SigningKey) GetConfigurationOk() (*InlinedConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *SigningKey) SetConfiguration(v InlinedConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *SigningKey) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetLength

`func (o *SigningKey) GetLength() int32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *SigningKey) GetLengthOk() (*int32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *SigningKey) SetLength(v int32)`

SetLength sets Length field to given value.

### HasLength

`func (o *SigningKey) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetState

`func (o *SigningKey) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SigningKey) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SigningKey) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *SigningKey) HasState() bool`

HasState returns a boolean if a field has been set.

### GetKeyGenerationDateTime

`func (o *SigningKey) GetKeyGenerationDateTime() time.Time`

GetKeyGenerationDateTime returns the KeyGenerationDateTime field if non-nil, zero value otherwise.

### GetKeyGenerationDateTimeOk

`func (o *SigningKey) GetKeyGenerationDateTimeOk() (*time.Time, bool)`

GetKeyGenerationDateTimeOk returns a tuple with the KeyGenerationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyGenerationDateTime

`func (o *SigningKey) SetKeyGenerationDateTime(v time.Time)`

SetKeyGenerationDateTime sets KeyGenerationDateTime field to given value.

### HasKeyGenerationDateTime

`func (o *SigningKey) HasKeyGenerationDateTime() bool`

HasKeyGenerationDateTime returns a boolean if a field has been set.

### GetPrivateKey

`func (o *SigningKey) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *SigningKey) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *SigningKey) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *SigningKey) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


