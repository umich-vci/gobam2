# ZoneSigningKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The resource type. | [optional] 
**Algorithm** | Pointer to **string** | The algorithm used to sign the key. | [optional] 
**KeyStartDateTime** | Pointer to **time.Time** | The date and time for the beginning of the key&#39;s validity period. | [optional] 
**KeyExpirationDateTime** | Pointer to **time.Time** | The date and time at which the key expires. | [optional] 
**Ttl** | Pointer to **string** | The TTL (time to live) for the key if an override TTL is specified when the key is created. | [optional] 
**KeyTag** | Pointer to **int32** | The key tag value for the key. The key tag is used during DNSSEC validation and when signing and re-signing zones. | [optional] 
**PublicKey** | Pointer to **string** | The public key text. | [optional] 
**Flags** | Pointer to **int32** | The flag used to verify the DNS record signature for resource records. | [optional] 

## Methods

### NewZoneSigningKey

`func NewZoneSigningKey() *ZoneSigningKey`

NewZoneSigningKey instantiates a new ZoneSigningKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneSigningKeyWithDefaults

`func NewZoneSigningKeyWithDefaults() *ZoneSigningKey`

NewZoneSigningKeyWithDefaults instantiates a new ZoneSigningKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ZoneSigningKey) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ZoneSigningKey) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ZoneSigningKey) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ZoneSigningKey) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAlgorithm

`func (o *ZoneSigningKey) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *ZoneSigningKey) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *ZoneSigningKey) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.

### HasAlgorithm

`func (o *ZoneSigningKey) HasAlgorithm() bool`

HasAlgorithm returns a boolean if a field has been set.

### GetKeyStartDateTime

`func (o *ZoneSigningKey) GetKeyStartDateTime() time.Time`

GetKeyStartDateTime returns the KeyStartDateTime field if non-nil, zero value otherwise.

### GetKeyStartDateTimeOk

`func (o *ZoneSigningKey) GetKeyStartDateTimeOk() (*time.Time, bool)`

GetKeyStartDateTimeOk returns a tuple with the KeyStartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyStartDateTime

`func (o *ZoneSigningKey) SetKeyStartDateTime(v time.Time)`

SetKeyStartDateTime sets KeyStartDateTime field to given value.

### HasKeyStartDateTime

`func (o *ZoneSigningKey) HasKeyStartDateTime() bool`

HasKeyStartDateTime returns a boolean if a field has been set.

### GetKeyExpirationDateTime

`func (o *ZoneSigningKey) GetKeyExpirationDateTime() time.Time`

GetKeyExpirationDateTime returns the KeyExpirationDateTime field if non-nil, zero value otherwise.

### GetKeyExpirationDateTimeOk

`func (o *ZoneSigningKey) GetKeyExpirationDateTimeOk() (*time.Time, bool)`

GetKeyExpirationDateTimeOk returns a tuple with the KeyExpirationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyExpirationDateTime

`func (o *ZoneSigningKey) SetKeyExpirationDateTime(v time.Time)`

SetKeyExpirationDateTime sets KeyExpirationDateTime field to given value.

### HasKeyExpirationDateTime

`func (o *ZoneSigningKey) HasKeyExpirationDateTime() bool`

HasKeyExpirationDateTime returns a boolean if a field has been set.

### GetTtl

`func (o *ZoneSigningKey) GetTtl() string`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *ZoneSigningKey) GetTtlOk() (*string, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *ZoneSigningKey) SetTtl(v string)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *ZoneSigningKey) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetKeyTag

`func (o *ZoneSigningKey) GetKeyTag() int32`

GetKeyTag returns the KeyTag field if non-nil, zero value otherwise.

### GetKeyTagOk

`func (o *ZoneSigningKey) GetKeyTagOk() (*int32, bool)`

GetKeyTagOk returns a tuple with the KeyTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyTag

`func (o *ZoneSigningKey) SetKeyTag(v int32)`

SetKeyTag sets KeyTag field to given value.

### HasKeyTag

`func (o *ZoneSigningKey) HasKeyTag() bool`

HasKeyTag returns a boolean if a field has been set.

### GetPublicKey

`func (o *ZoneSigningKey) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *ZoneSigningKey) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *ZoneSigningKey) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *ZoneSigningKey) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetFlags

`func (o *ZoneSigningKey) GetFlags() int32`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *ZoneSigningKey) GetFlagsOk() (*int32, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *ZoneSigningKey) SetFlags(v int32)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *ZoneSigningKey) HasFlags() bool`

HasFlags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


