# DNSSECSigningPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | Pointer to **string** | The name of the resource. | [optional] 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**KeyProvider** | Pointer to **string** | The system that created the DNSSEC signing key. | [optional] 
**SignatureDigestAlgorithm** | Pointer to **string** | The algorithm used for the Delegation Signer (DS) record. | [optional] 
**SignatureValidityPeriod** | Pointer to **string** | The duration for which the RRSIG resource record is valid. | [optional] 
**SignatureResigningPeriod** | Pointer to **string** | The duration before the end of the signature validity period at which BIND resigns the zone. | [optional] 
**ZoneSigningKeyAlgorithm** | Pointer to **string** | The algorithm for the Zone Signing Key (ZSK). | [optional] 
**ZoneSigningKeyLength** | Pointer to **int32** | The length of the zone signing key in bits. | [optional] 
**ZoneSigningKeyOverrideTtl** | Pointer to **string** | The overridden TTL value of the zone signing key. | [optional] 
**ZoneSigningKeyValidityPeriod** | Pointer to **string** | The duration for which the zone signing key is valid. | [optional] 
**ZoneSigningKeyOverlapPeriod** | Pointer to **string** | The duration before the end of the validity period at which a new key is generated for key rollover. | [optional] 
**ZoneSigningKeyRolloverMethod** | Pointer to **string** | The method to make the new zone signing key available when the key rolls over. | [optional] 
**ZoneSigningKeySigningPeriod** | Pointer to **string** | The duration before the end of the key validity period. During this time, the resource records in the zone are simultaneously signed by the new key and unsigned by the old key. | [optional] 
**ZoneSigningKeyProtectionType** | Pointer to **string** | The zone signing key protection type when Enstrust HSM is configured as the key provider. | [optional] 
**KeySigningKeyAlgorithm** | Pointer to **string** | The algorithm for the key signing key (KSK). | [optional] 
**KeySigningKeyLength** | Pointer to **int32** | The length of the key signing key in bits. | [optional] 
**KeySigningKeyOverrideTtl** | Pointer to **string** | The overridden TTL value of the key signing key. | [optional] 
**KeySigningKeyValidityPeriod** | Pointer to **string** | The duration for which the key signing key is valid. | [optional] 
**KeySigningKeyOverlapPeriod** | Pointer to **string** | The duration before the end of the validity period at which a new key is generated for key rollover. | [optional] 
**KeySigningKeyRolloverMethod** | Pointer to **string** | The method to make the new key signing key available when the key rolls over. | [optional] 
**KeySigningKeySigningPeriod** | Pointer to **string** | The duration before the end of the key validity period. During this time, the resource records in the keys are simultaneously signed by the new key and unsigned by the old key. | [optional] 
**KeySigningKeyProtectionType** | Pointer to **string** | The key signing key protection type when Enstrust HSM is configured as the key provider. | [optional] 

## Methods

### NewDNSSECSigningPolicy

`func NewDNSSECSigningPolicy() *DNSSECSigningPolicy`

NewDNSSECSigningPolicy instantiates a new DNSSECSigningPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDNSSECSigningPolicyWithDefaults

`func NewDNSSECSigningPolicyWithDefaults() *DNSSECSigningPolicy`

NewDNSSECSigningPolicyWithDefaults instantiates a new DNSSECSigningPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DNSSECSigningPolicy) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DNSSECSigningPolicy) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DNSSECSigningPolicy) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DNSSECSigningPolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *DNSSECSigningPolicy) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DNSSECSigningPolicy) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DNSSECSigningPolicy) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DNSSECSigningPolicy) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *DNSSECSigningPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DNSSECSigningPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DNSSECSigningPolicy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DNSSECSigningPolicy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *DNSSECSigningPolicy) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *DNSSECSigningPolicy) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *DNSSECSigningPolicy) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *DNSSECSigningPolicy) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetKeyProvider

`func (o *DNSSECSigningPolicy) GetKeyProvider() string`

GetKeyProvider returns the KeyProvider field if non-nil, zero value otherwise.

### GetKeyProviderOk

`func (o *DNSSECSigningPolicy) GetKeyProviderOk() (*string, bool)`

GetKeyProviderOk returns a tuple with the KeyProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyProvider

`func (o *DNSSECSigningPolicy) SetKeyProvider(v string)`

SetKeyProvider sets KeyProvider field to given value.

### HasKeyProvider

`func (o *DNSSECSigningPolicy) HasKeyProvider() bool`

HasKeyProvider returns a boolean if a field has been set.

### GetSignatureDigestAlgorithm

`func (o *DNSSECSigningPolicy) GetSignatureDigestAlgorithm() string`

GetSignatureDigestAlgorithm returns the SignatureDigestAlgorithm field if non-nil, zero value otherwise.

### GetSignatureDigestAlgorithmOk

`func (o *DNSSECSigningPolicy) GetSignatureDigestAlgorithmOk() (*string, bool)`

GetSignatureDigestAlgorithmOk returns a tuple with the SignatureDigestAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureDigestAlgorithm

`func (o *DNSSECSigningPolicy) SetSignatureDigestAlgorithm(v string)`

SetSignatureDigestAlgorithm sets SignatureDigestAlgorithm field to given value.

### HasSignatureDigestAlgorithm

`func (o *DNSSECSigningPolicy) HasSignatureDigestAlgorithm() bool`

HasSignatureDigestAlgorithm returns a boolean if a field has been set.

### GetSignatureValidityPeriod

`func (o *DNSSECSigningPolicy) GetSignatureValidityPeriod() string`

GetSignatureValidityPeriod returns the SignatureValidityPeriod field if non-nil, zero value otherwise.

### GetSignatureValidityPeriodOk

`func (o *DNSSECSigningPolicy) GetSignatureValidityPeriodOk() (*string, bool)`

GetSignatureValidityPeriodOk returns a tuple with the SignatureValidityPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureValidityPeriod

`func (o *DNSSECSigningPolicy) SetSignatureValidityPeriod(v string)`

SetSignatureValidityPeriod sets SignatureValidityPeriod field to given value.

### HasSignatureValidityPeriod

`func (o *DNSSECSigningPolicy) HasSignatureValidityPeriod() bool`

HasSignatureValidityPeriod returns a boolean if a field has been set.

### GetSignatureResigningPeriod

`func (o *DNSSECSigningPolicy) GetSignatureResigningPeriod() string`

GetSignatureResigningPeriod returns the SignatureResigningPeriod field if non-nil, zero value otherwise.

### GetSignatureResigningPeriodOk

`func (o *DNSSECSigningPolicy) GetSignatureResigningPeriodOk() (*string, bool)`

GetSignatureResigningPeriodOk returns a tuple with the SignatureResigningPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureResigningPeriod

`func (o *DNSSECSigningPolicy) SetSignatureResigningPeriod(v string)`

SetSignatureResigningPeriod sets SignatureResigningPeriod field to given value.

### HasSignatureResigningPeriod

`func (o *DNSSECSigningPolicy) HasSignatureResigningPeriod() bool`

HasSignatureResigningPeriod returns a boolean if a field has been set.

### GetZoneSigningKeyAlgorithm

`func (o *DNSSECSigningPolicy) GetZoneSigningKeyAlgorithm() string`

GetZoneSigningKeyAlgorithm returns the ZoneSigningKeyAlgorithm field if non-nil, zero value otherwise.

### GetZoneSigningKeyAlgorithmOk

`func (o *DNSSECSigningPolicy) GetZoneSigningKeyAlgorithmOk() (*string, bool)`

GetZoneSigningKeyAlgorithmOk returns a tuple with the ZoneSigningKeyAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneSigningKeyAlgorithm

`func (o *DNSSECSigningPolicy) SetZoneSigningKeyAlgorithm(v string)`

SetZoneSigningKeyAlgorithm sets ZoneSigningKeyAlgorithm field to given value.

### HasZoneSigningKeyAlgorithm

`func (o *DNSSECSigningPolicy) HasZoneSigningKeyAlgorithm() bool`

HasZoneSigningKeyAlgorithm returns a boolean if a field has been set.

### GetZoneSigningKeyLength

`func (o *DNSSECSigningPolicy) GetZoneSigningKeyLength() int32`

GetZoneSigningKeyLength returns the ZoneSigningKeyLength field if non-nil, zero value otherwise.

### GetZoneSigningKeyLengthOk

`func (o *DNSSECSigningPolicy) GetZoneSigningKeyLengthOk() (*int32, bool)`

GetZoneSigningKeyLengthOk returns a tuple with the ZoneSigningKeyLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneSigningKeyLength

`func (o *DNSSECSigningPolicy) SetZoneSigningKeyLength(v int32)`

SetZoneSigningKeyLength sets ZoneSigningKeyLength field to given value.

### HasZoneSigningKeyLength

`func (o *DNSSECSigningPolicy) HasZoneSigningKeyLength() bool`

HasZoneSigningKeyLength returns a boolean if a field has been set.

### GetZoneSigningKeyOverrideTtl

`func (o *DNSSECSigningPolicy) GetZoneSigningKeyOverrideTtl() string`

GetZoneSigningKeyOverrideTtl returns the ZoneSigningKeyOverrideTtl field if non-nil, zero value otherwise.

### GetZoneSigningKeyOverrideTtlOk

`func (o *DNSSECSigningPolicy) GetZoneSigningKeyOverrideTtlOk() (*string, bool)`

GetZoneSigningKeyOverrideTtlOk returns a tuple with the ZoneSigningKeyOverrideTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneSigningKeyOverrideTtl

`func (o *DNSSECSigningPolicy) SetZoneSigningKeyOverrideTtl(v string)`

SetZoneSigningKeyOverrideTtl sets ZoneSigningKeyOverrideTtl field to given value.

### HasZoneSigningKeyOverrideTtl

`func (o *DNSSECSigningPolicy) HasZoneSigningKeyOverrideTtl() bool`

HasZoneSigningKeyOverrideTtl returns a boolean if a field has been set.

### GetZoneSigningKeyValidityPeriod

`func (o *DNSSECSigningPolicy) GetZoneSigningKeyValidityPeriod() string`

GetZoneSigningKeyValidityPeriod returns the ZoneSigningKeyValidityPeriod field if non-nil, zero value otherwise.

### GetZoneSigningKeyValidityPeriodOk

`func (o *DNSSECSigningPolicy) GetZoneSigningKeyValidityPeriodOk() (*string, bool)`

GetZoneSigningKeyValidityPeriodOk returns a tuple with the ZoneSigningKeyValidityPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneSigningKeyValidityPeriod

`func (o *DNSSECSigningPolicy) SetZoneSigningKeyValidityPeriod(v string)`

SetZoneSigningKeyValidityPeriod sets ZoneSigningKeyValidityPeriod field to given value.

### HasZoneSigningKeyValidityPeriod

`func (o *DNSSECSigningPolicy) HasZoneSigningKeyValidityPeriod() bool`

HasZoneSigningKeyValidityPeriod returns a boolean if a field has been set.

### GetZoneSigningKeyOverlapPeriod

`func (o *DNSSECSigningPolicy) GetZoneSigningKeyOverlapPeriod() string`

GetZoneSigningKeyOverlapPeriod returns the ZoneSigningKeyOverlapPeriod field if non-nil, zero value otherwise.

### GetZoneSigningKeyOverlapPeriodOk

`func (o *DNSSECSigningPolicy) GetZoneSigningKeyOverlapPeriodOk() (*string, bool)`

GetZoneSigningKeyOverlapPeriodOk returns a tuple with the ZoneSigningKeyOverlapPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneSigningKeyOverlapPeriod

`func (o *DNSSECSigningPolicy) SetZoneSigningKeyOverlapPeriod(v string)`

SetZoneSigningKeyOverlapPeriod sets ZoneSigningKeyOverlapPeriod field to given value.

### HasZoneSigningKeyOverlapPeriod

`func (o *DNSSECSigningPolicy) HasZoneSigningKeyOverlapPeriod() bool`

HasZoneSigningKeyOverlapPeriod returns a boolean if a field has been set.

### GetZoneSigningKeyRolloverMethod

`func (o *DNSSECSigningPolicy) GetZoneSigningKeyRolloverMethod() string`

GetZoneSigningKeyRolloverMethod returns the ZoneSigningKeyRolloverMethod field if non-nil, zero value otherwise.

### GetZoneSigningKeyRolloverMethodOk

`func (o *DNSSECSigningPolicy) GetZoneSigningKeyRolloverMethodOk() (*string, bool)`

GetZoneSigningKeyRolloverMethodOk returns a tuple with the ZoneSigningKeyRolloverMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneSigningKeyRolloverMethod

`func (o *DNSSECSigningPolicy) SetZoneSigningKeyRolloverMethod(v string)`

SetZoneSigningKeyRolloverMethod sets ZoneSigningKeyRolloverMethod field to given value.

### HasZoneSigningKeyRolloverMethod

`func (o *DNSSECSigningPolicy) HasZoneSigningKeyRolloverMethod() bool`

HasZoneSigningKeyRolloverMethod returns a boolean if a field has been set.

### GetZoneSigningKeySigningPeriod

`func (o *DNSSECSigningPolicy) GetZoneSigningKeySigningPeriod() string`

GetZoneSigningKeySigningPeriod returns the ZoneSigningKeySigningPeriod field if non-nil, zero value otherwise.

### GetZoneSigningKeySigningPeriodOk

`func (o *DNSSECSigningPolicy) GetZoneSigningKeySigningPeriodOk() (*string, bool)`

GetZoneSigningKeySigningPeriodOk returns a tuple with the ZoneSigningKeySigningPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneSigningKeySigningPeriod

`func (o *DNSSECSigningPolicy) SetZoneSigningKeySigningPeriod(v string)`

SetZoneSigningKeySigningPeriod sets ZoneSigningKeySigningPeriod field to given value.

### HasZoneSigningKeySigningPeriod

`func (o *DNSSECSigningPolicy) HasZoneSigningKeySigningPeriod() bool`

HasZoneSigningKeySigningPeriod returns a boolean if a field has been set.

### GetZoneSigningKeyProtectionType

`func (o *DNSSECSigningPolicy) GetZoneSigningKeyProtectionType() string`

GetZoneSigningKeyProtectionType returns the ZoneSigningKeyProtectionType field if non-nil, zero value otherwise.

### GetZoneSigningKeyProtectionTypeOk

`func (o *DNSSECSigningPolicy) GetZoneSigningKeyProtectionTypeOk() (*string, bool)`

GetZoneSigningKeyProtectionTypeOk returns a tuple with the ZoneSigningKeyProtectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneSigningKeyProtectionType

`func (o *DNSSECSigningPolicy) SetZoneSigningKeyProtectionType(v string)`

SetZoneSigningKeyProtectionType sets ZoneSigningKeyProtectionType field to given value.

### HasZoneSigningKeyProtectionType

`func (o *DNSSECSigningPolicy) HasZoneSigningKeyProtectionType() bool`

HasZoneSigningKeyProtectionType returns a boolean if a field has been set.

### GetKeySigningKeyAlgorithm

`func (o *DNSSECSigningPolicy) GetKeySigningKeyAlgorithm() string`

GetKeySigningKeyAlgorithm returns the KeySigningKeyAlgorithm field if non-nil, zero value otherwise.

### GetKeySigningKeyAlgorithmOk

`func (o *DNSSECSigningPolicy) GetKeySigningKeyAlgorithmOk() (*string, bool)`

GetKeySigningKeyAlgorithmOk returns a tuple with the KeySigningKeyAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySigningKeyAlgorithm

`func (o *DNSSECSigningPolicy) SetKeySigningKeyAlgorithm(v string)`

SetKeySigningKeyAlgorithm sets KeySigningKeyAlgorithm field to given value.

### HasKeySigningKeyAlgorithm

`func (o *DNSSECSigningPolicy) HasKeySigningKeyAlgorithm() bool`

HasKeySigningKeyAlgorithm returns a boolean if a field has been set.

### GetKeySigningKeyLength

`func (o *DNSSECSigningPolicy) GetKeySigningKeyLength() int32`

GetKeySigningKeyLength returns the KeySigningKeyLength field if non-nil, zero value otherwise.

### GetKeySigningKeyLengthOk

`func (o *DNSSECSigningPolicy) GetKeySigningKeyLengthOk() (*int32, bool)`

GetKeySigningKeyLengthOk returns a tuple with the KeySigningKeyLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySigningKeyLength

`func (o *DNSSECSigningPolicy) SetKeySigningKeyLength(v int32)`

SetKeySigningKeyLength sets KeySigningKeyLength field to given value.

### HasKeySigningKeyLength

`func (o *DNSSECSigningPolicy) HasKeySigningKeyLength() bool`

HasKeySigningKeyLength returns a boolean if a field has been set.

### GetKeySigningKeyOverrideTtl

`func (o *DNSSECSigningPolicy) GetKeySigningKeyOverrideTtl() string`

GetKeySigningKeyOverrideTtl returns the KeySigningKeyOverrideTtl field if non-nil, zero value otherwise.

### GetKeySigningKeyOverrideTtlOk

`func (o *DNSSECSigningPolicy) GetKeySigningKeyOverrideTtlOk() (*string, bool)`

GetKeySigningKeyOverrideTtlOk returns a tuple with the KeySigningKeyOverrideTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySigningKeyOverrideTtl

`func (o *DNSSECSigningPolicy) SetKeySigningKeyOverrideTtl(v string)`

SetKeySigningKeyOverrideTtl sets KeySigningKeyOverrideTtl field to given value.

### HasKeySigningKeyOverrideTtl

`func (o *DNSSECSigningPolicy) HasKeySigningKeyOverrideTtl() bool`

HasKeySigningKeyOverrideTtl returns a boolean if a field has been set.

### GetKeySigningKeyValidityPeriod

`func (o *DNSSECSigningPolicy) GetKeySigningKeyValidityPeriod() string`

GetKeySigningKeyValidityPeriod returns the KeySigningKeyValidityPeriod field if non-nil, zero value otherwise.

### GetKeySigningKeyValidityPeriodOk

`func (o *DNSSECSigningPolicy) GetKeySigningKeyValidityPeriodOk() (*string, bool)`

GetKeySigningKeyValidityPeriodOk returns a tuple with the KeySigningKeyValidityPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySigningKeyValidityPeriod

`func (o *DNSSECSigningPolicy) SetKeySigningKeyValidityPeriod(v string)`

SetKeySigningKeyValidityPeriod sets KeySigningKeyValidityPeriod field to given value.

### HasKeySigningKeyValidityPeriod

`func (o *DNSSECSigningPolicy) HasKeySigningKeyValidityPeriod() bool`

HasKeySigningKeyValidityPeriod returns a boolean if a field has been set.

### GetKeySigningKeyOverlapPeriod

`func (o *DNSSECSigningPolicy) GetKeySigningKeyOverlapPeriod() string`

GetKeySigningKeyOverlapPeriod returns the KeySigningKeyOverlapPeriod field if non-nil, zero value otherwise.

### GetKeySigningKeyOverlapPeriodOk

`func (o *DNSSECSigningPolicy) GetKeySigningKeyOverlapPeriodOk() (*string, bool)`

GetKeySigningKeyOverlapPeriodOk returns a tuple with the KeySigningKeyOverlapPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySigningKeyOverlapPeriod

`func (o *DNSSECSigningPolicy) SetKeySigningKeyOverlapPeriod(v string)`

SetKeySigningKeyOverlapPeriod sets KeySigningKeyOverlapPeriod field to given value.

### HasKeySigningKeyOverlapPeriod

`func (o *DNSSECSigningPolicy) HasKeySigningKeyOverlapPeriod() bool`

HasKeySigningKeyOverlapPeriod returns a boolean if a field has been set.

### GetKeySigningKeyRolloverMethod

`func (o *DNSSECSigningPolicy) GetKeySigningKeyRolloverMethod() string`

GetKeySigningKeyRolloverMethod returns the KeySigningKeyRolloverMethod field if non-nil, zero value otherwise.

### GetKeySigningKeyRolloverMethodOk

`func (o *DNSSECSigningPolicy) GetKeySigningKeyRolloverMethodOk() (*string, bool)`

GetKeySigningKeyRolloverMethodOk returns a tuple with the KeySigningKeyRolloverMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySigningKeyRolloverMethod

`func (o *DNSSECSigningPolicy) SetKeySigningKeyRolloverMethod(v string)`

SetKeySigningKeyRolloverMethod sets KeySigningKeyRolloverMethod field to given value.

### HasKeySigningKeyRolloverMethod

`func (o *DNSSECSigningPolicy) HasKeySigningKeyRolloverMethod() bool`

HasKeySigningKeyRolloverMethod returns a boolean if a field has been set.

### GetKeySigningKeySigningPeriod

`func (o *DNSSECSigningPolicy) GetKeySigningKeySigningPeriod() string`

GetKeySigningKeySigningPeriod returns the KeySigningKeySigningPeriod field if non-nil, zero value otherwise.

### GetKeySigningKeySigningPeriodOk

`func (o *DNSSECSigningPolicy) GetKeySigningKeySigningPeriodOk() (*string, bool)`

GetKeySigningKeySigningPeriodOk returns a tuple with the KeySigningKeySigningPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySigningKeySigningPeriod

`func (o *DNSSECSigningPolicy) SetKeySigningKeySigningPeriod(v string)`

SetKeySigningKeySigningPeriod sets KeySigningKeySigningPeriod field to given value.

### HasKeySigningKeySigningPeriod

`func (o *DNSSECSigningPolicy) HasKeySigningKeySigningPeriod() bool`

HasKeySigningKeySigningPeriod returns a boolean if a field has been set.

### GetKeySigningKeyProtectionType

`func (o *DNSSECSigningPolicy) GetKeySigningKeyProtectionType() string`

GetKeySigningKeyProtectionType returns the KeySigningKeyProtectionType field if non-nil, zero value otherwise.

### GetKeySigningKeyProtectionTypeOk

`func (o *DNSSECSigningPolicy) GetKeySigningKeyProtectionTypeOk() (*string, bool)`

GetKeySigningKeyProtectionTypeOk returns a tuple with the KeySigningKeyProtectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySigningKeyProtectionType

`func (o *DNSSECSigningPolicy) SetKeySigningKeyProtectionType(v string)`

SetKeySigningKeyProtectionType sets KeySigningKeyProtectionType field to given value.

### HasKeySigningKeyProtectionType

`func (o *DNSSECSigningPolicy) HasKeySigningKeyProtectionType() bool`

HasKeySigningKeyProtectionType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


