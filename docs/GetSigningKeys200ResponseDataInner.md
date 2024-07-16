# GetSigningKeys200ResponseDataInner

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
**Algorithm** | Pointer to **string** | The algorithm used to sign the key. | [optional] 
**KeyStartDateTime** | Pointer to **time.Time** | The date and time for the beginning of the key&#39;s validity period. | [optional] 
**KeyExpirationDateTime** | Pointer to **time.Time** | The date and time at which the key expires. | [optional] 
**Ttl** | Pointer to **string** | The TTL (time to live) for the key if an override TTL is specified when the key is created. | [optional] 
**KeyTag** | Pointer to **int32** | The key tag value for the key. The key tag is used during DNSSEC validation and when signing and re-signing zones. | [optional] 
**PublicKey** | Pointer to **string** | The public key text. | [optional] 
**Flags** | Pointer to **int32** | The flag used to verify the DNS record signature for resource records. | [optional] 
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 

## Methods

### NewGetSigningKeys200ResponseDataInner

`func NewGetSigningKeys200ResponseDataInner() *GetSigningKeys200ResponseDataInner`

NewGetSigningKeys200ResponseDataInner instantiates a new GetSigningKeys200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSigningKeys200ResponseDataInnerWithDefaults

`func NewGetSigningKeys200ResponseDataInnerWithDefaults() *GetSigningKeys200ResponseDataInner`

NewGetSigningKeys200ResponseDataInnerWithDefaults instantiates a new GetSigningKeys200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetSigningKeys200ResponseDataInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetSigningKeys200ResponseDataInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetSigningKeys200ResponseDataInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetSigningKeys200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GetSigningKeys200ResponseDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetSigningKeys200ResponseDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetSigningKeys200ResponseDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetSigningKeys200ResponseDataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *GetSigningKeys200ResponseDataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetSigningKeys200ResponseDataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetSigningKeys200ResponseDataInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetSigningKeys200ResponseDataInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *GetSigningKeys200ResponseDataInner) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *GetSigningKeys200ResponseDataInner) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *GetSigningKeys200ResponseDataInner) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *GetSigningKeys200ResponseDataInner) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetConfiguration

`func (o *GetSigningKeys200ResponseDataInner) GetConfiguration() InlinedConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *GetSigningKeys200ResponseDataInner) GetConfigurationOk() (*InlinedConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *GetSigningKeys200ResponseDataInner) SetConfiguration(v InlinedConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *GetSigningKeys200ResponseDataInner) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetLength

`func (o *GetSigningKeys200ResponseDataInner) GetLength() int32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *GetSigningKeys200ResponseDataInner) GetLengthOk() (*int32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *GetSigningKeys200ResponseDataInner) SetLength(v int32)`

SetLength sets Length field to given value.

### HasLength

`func (o *GetSigningKeys200ResponseDataInner) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetState

`func (o *GetSigningKeys200ResponseDataInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetSigningKeys200ResponseDataInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetSigningKeys200ResponseDataInner) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GetSigningKeys200ResponseDataInner) HasState() bool`

HasState returns a boolean if a field has been set.

### GetKeyGenerationDateTime

`func (o *GetSigningKeys200ResponseDataInner) GetKeyGenerationDateTime() time.Time`

GetKeyGenerationDateTime returns the KeyGenerationDateTime field if non-nil, zero value otherwise.

### GetKeyGenerationDateTimeOk

`func (o *GetSigningKeys200ResponseDataInner) GetKeyGenerationDateTimeOk() (*time.Time, bool)`

GetKeyGenerationDateTimeOk returns a tuple with the KeyGenerationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyGenerationDateTime

`func (o *GetSigningKeys200ResponseDataInner) SetKeyGenerationDateTime(v time.Time)`

SetKeyGenerationDateTime sets KeyGenerationDateTime field to given value.

### HasKeyGenerationDateTime

`func (o *GetSigningKeys200ResponseDataInner) HasKeyGenerationDateTime() bool`

HasKeyGenerationDateTime returns a boolean if a field has been set.

### GetPrivateKey

`func (o *GetSigningKeys200ResponseDataInner) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *GetSigningKeys200ResponseDataInner) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *GetSigningKeys200ResponseDataInner) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *GetSigningKeys200ResponseDataInner) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetAlgorithm

`func (o *GetSigningKeys200ResponseDataInner) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *GetSigningKeys200ResponseDataInner) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *GetSigningKeys200ResponseDataInner) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.

### HasAlgorithm

`func (o *GetSigningKeys200ResponseDataInner) HasAlgorithm() bool`

HasAlgorithm returns a boolean if a field has been set.

### GetKeyStartDateTime

`func (o *GetSigningKeys200ResponseDataInner) GetKeyStartDateTime() time.Time`

GetKeyStartDateTime returns the KeyStartDateTime field if non-nil, zero value otherwise.

### GetKeyStartDateTimeOk

`func (o *GetSigningKeys200ResponseDataInner) GetKeyStartDateTimeOk() (*time.Time, bool)`

GetKeyStartDateTimeOk returns a tuple with the KeyStartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyStartDateTime

`func (o *GetSigningKeys200ResponseDataInner) SetKeyStartDateTime(v time.Time)`

SetKeyStartDateTime sets KeyStartDateTime field to given value.

### HasKeyStartDateTime

`func (o *GetSigningKeys200ResponseDataInner) HasKeyStartDateTime() bool`

HasKeyStartDateTime returns a boolean if a field has been set.

### GetKeyExpirationDateTime

`func (o *GetSigningKeys200ResponseDataInner) GetKeyExpirationDateTime() time.Time`

GetKeyExpirationDateTime returns the KeyExpirationDateTime field if non-nil, zero value otherwise.

### GetKeyExpirationDateTimeOk

`func (o *GetSigningKeys200ResponseDataInner) GetKeyExpirationDateTimeOk() (*time.Time, bool)`

GetKeyExpirationDateTimeOk returns a tuple with the KeyExpirationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyExpirationDateTime

`func (o *GetSigningKeys200ResponseDataInner) SetKeyExpirationDateTime(v time.Time)`

SetKeyExpirationDateTime sets KeyExpirationDateTime field to given value.

### HasKeyExpirationDateTime

`func (o *GetSigningKeys200ResponseDataInner) HasKeyExpirationDateTime() bool`

HasKeyExpirationDateTime returns a boolean if a field has been set.

### GetTtl

`func (o *GetSigningKeys200ResponseDataInner) GetTtl() string`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *GetSigningKeys200ResponseDataInner) GetTtlOk() (*string, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *GetSigningKeys200ResponseDataInner) SetTtl(v string)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *GetSigningKeys200ResponseDataInner) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetKeyTag

`func (o *GetSigningKeys200ResponseDataInner) GetKeyTag() int32`

GetKeyTag returns the KeyTag field if non-nil, zero value otherwise.

### GetKeyTagOk

`func (o *GetSigningKeys200ResponseDataInner) GetKeyTagOk() (*int32, bool)`

GetKeyTagOk returns a tuple with the KeyTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyTag

`func (o *GetSigningKeys200ResponseDataInner) SetKeyTag(v int32)`

SetKeyTag sets KeyTag field to given value.

### HasKeyTag

`func (o *GetSigningKeys200ResponseDataInner) HasKeyTag() bool`

HasKeyTag returns a boolean if a field has been set.

### GetPublicKey

`func (o *GetSigningKeys200ResponseDataInner) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *GetSigningKeys200ResponseDataInner) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *GetSigningKeys200ResponseDataInner) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *GetSigningKeys200ResponseDataInner) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetFlags

`func (o *GetSigningKeys200ResponseDataInner) GetFlags() int32`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *GetSigningKeys200ResponseDataInner) GetFlagsOk() (*int32, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *GetSigningKeys200ResponseDataInner) SetFlags(v int32)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *GetSigningKeys200ResponseDataInner) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetLinks

`func (o *GetSigningKeys200ResponseDataInner) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetSigningKeys200ResponseDataInner) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetSigningKeys200ResponseDataInner) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetSigningKeys200ResponseDataInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


