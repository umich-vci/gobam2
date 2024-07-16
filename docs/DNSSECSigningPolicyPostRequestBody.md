# DNSSECSigningPolicyPostRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | **string** | The name of the resource. | 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**KeyProvider** | **string** | The system that created the DNSSEC signing key. | 
**SignatureDigestAlgorithm** | **string** | The algorithm used for the Delegation Signer (DS) record. | 
**SignatureValidityPeriod** | **string** | The duration for which the RRSIG resource record is valid. | 
**SignatureResigningPeriod** | **string** | The duration before the end of the signature validity period at which BIND resigns the zone. | 
**ZoneSigningKeyAlgorithm** | **string** | The algorithm for the Zone Signing Key (ZSK). | 
**ZoneSigningKeyLength** | **int32** | The length of the zone signing key in bits. | 
**ZoneSigningKeyOverrideTtl** | Pointer to **string** | The overridden TTL value of the zone signing key. | [optional] 
**ZoneSigningKeyValidityPeriod** | **string** | The duration for which the zone signing key is valid. | 
**ZoneSigningKeyOverlapPeriod** | **string** | The duration before the end of the validity period at which a new key is generated for key rollover. | 
**ZoneSigningKeyRolloverMethod** | **string** | The method to make the new zone signing key available when the key rolls over. | 
**ZoneSigningKeySigningPeriod** | Pointer to **string** | The duration before the end of the key validity period. During this time, the resource records in the zone are simultaneously signed by the new key and unsigned by the old key. | [optional] 
**ZoneSigningKeyProtectionType** | Pointer to **string** | The zone signing key protection type when Enstrust HSM is configured as the key provider. | [optional] 
**KeySigningKeyAlgorithm** | **string** | The algorithm for the key signing key (KSK). | 
**KeySigningKeyLength** | **int32** | The length of the key signing key in bits. | 
**KeySigningKeyOverrideTtl** | Pointer to **string** | The overridden TTL value of the key signing key. | [optional] 
**KeySigningKeyValidityPeriod** | **string** | The duration for which the key signing key is valid. | 
**KeySigningKeyOverlapPeriod** | **string** | The duration before the end of the validity period at which a new key is generated for key rollover. | 
**KeySigningKeyRolloverMethod** | **string** | The method to make the new key signing key available when the key rolls over. | 
**KeySigningKeySigningPeriod** | Pointer to **string** | The duration before the end of the key validity period. During this time, the resource records in the keys are simultaneously signed by the new key and unsigned by the old key. | [optional] 
**KeySigningKeyProtectionType** | Pointer to **string** | The key signing key protection type when Enstrust HSM is configured as the key provider. | [optional] 

## Methods

### NewDNSSECSigningPolicyPostRequestBody

`func NewDNSSECSigningPolicyPostRequestBody(name string, keyProvider string, signatureDigestAlgorithm string, signatureValidityPeriod string, signatureResigningPeriod string, zoneSigningKeyAlgorithm string, zoneSigningKeyLength int32, zoneSigningKeyValidityPeriod string, zoneSigningKeyOverlapPeriod string, zoneSigningKeyRolloverMethod string, keySigningKeyAlgorithm string, keySigningKeyLength int32, keySigningKeyValidityPeriod string, keySigningKeyOverlapPeriod string, keySigningKeyRolloverMethod string, ) *DNSSECSigningPolicyPostRequestBody`

NewDNSSECSigningPolicyPostRequestBody instantiates a new DNSSECSigningPolicyPostRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDNSSECSigningPolicyPostRequestBodyWithDefaults

`func NewDNSSECSigningPolicyPostRequestBodyWithDefaults() *DNSSECSigningPolicyPostRequestBody`

NewDNSSECSigningPolicyPostRequestBodyWithDefaults instantiates a new DNSSECSigningPolicyPostRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DNSSECSigningPolicyPostRequestBody) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DNSSECSigningPolicyPostRequestBody) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DNSSECSigningPolicyPostRequestBody) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DNSSECSigningPolicyPostRequestBody) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *DNSSECSigningPolicyPostRequestBody) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DNSSECSigningPolicyPostRequestBody) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DNSSECSigningPolicyPostRequestBody) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DNSSECSigningPolicyPostRequestBody) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *DNSSECSigningPolicyPostRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DNSSECSigningPolicyPostRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DNSSECSigningPolicyPostRequestBody) SetName(v string)`

SetName sets Name field to given value.


### GetUserDefinedFields

`func (o *DNSSECSigningPolicyPostRequestBody) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *DNSSECSigningPolicyPostRequestBody) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *DNSSECSigningPolicyPostRequestBody) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *DNSSECSigningPolicyPostRequestBody) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetKeyProvider

`func (o *DNSSECSigningPolicyPostRequestBody) GetKeyProvider() string`

GetKeyProvider returns the KeyProvider field if non-nil, zero value otherwise.

### GetKeyProviderOk

`func (o *DNSSECSigningPolicyPostRequestBody) GetKeyProviderOk() (*string, bool)`

GetKeyProviderOk returns a tuple with the KeyProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyProvider

`func (o *DNSSECSigningPolicyPostRequestBody) SetKeyProvider(v string)`

SetKeyProvider sets KeyProvider field to given value.


### GetSignatureDigestAlgorithm

`func (o *DNSSECSigningPolicyPostRequestBody) GetSignatureDigestAlgorithm() string`

GetSignatureDigestAlgorithm returns the SignatureDigestAlgorithm field if non-nil, zero value otherwise.

### GetSignatureDigestAlgorithmOk

`func (o *DNSSECSigningPolicyPostRequestBody) GetSignatureDigestAlgorithmOk() (*string, bool)`

GetSignatureDigestAlgorithmOk returns a tuple with the SignatureDigestAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureDigestAlgorithm

`func (o *DNSSECSigningPolicyPostRequestBody) SetSignatureDigestAlgorithm(v string)`

SetSignatureDigestAlgorithm sets SignatureDigestAlgorithm field to given value.


### GetSignatureValidityPeriod

`func (o *DNSSECSigningPolicyPostRequestBody) GetSignatureValidityPeriod() string`

GetSignatureValidityPeriod returns the SignatureValidityPeriod field if non-nil, zero value otherwise.

### GetSignatureValidityPeriodOk

`func (o *DNSSECSigningPolicyPostRequestBody) GetSignatureValidityPeriodOk() (*string, bool)`

GetSignatureValidityPeriodOk returns a tuple with the SignatureValidityPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureValidityPeriod

`func (o *DNSSECSigningPolicyPostRequestBody) SetSignatureValidityPeriod(v string)`

SetSignatureValidityPeriod sets SignatureValidityPeriod field to given value.


### GetSignatureResigningPeriod

`func (o *DNSSECSigningPolicyPostRequestBody) GetSignatureResigningPeriod() string`

GetSignatureResigningPeriod returns the SignatureResigningPeriod field if non-nil, zero value otherwise.

### GetSignatureResigningPeriodOk

`func (o *DNSSECSigningPolicyPostRequestBody) GetSignatureResigningPeriodOk() (*string, bool)`

GetSignatureResigningPeriodOk returns a tuple with the SignatureResigningPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureResigningPeriod

`func (o *DNSSECSigningPolicyPostRequestBody) SetSignatureResigningPeriod(v string)`

SetSignatureResigningPeriod sets SignatureResigningPeriod field to given value.


### GetZoneSigningKeyAlgorithm

`func (o *DNSSECSigningPolicyPostRequestBody) GetZoneSigningKeyAlgorithm() string`

GetZoneSigningKeyAlgorithm returns the ZoneSigningKeyAlgorithm field if non-nil, zero value otherwise.

### GetZoneSigningKeyAlgorithmOk

`func (o *DNSSECSigningPolicyPostRequestBody) GetZoneSigningKeyAlgorithmOk() (*string, bool)`

GetZoneSigningKeyAlgorithmOk returns a tuple with the ZoneSigningKeyAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneSigningKeyAlgorithm

`func (o *DNSSECSigningPolicyPostRequestBody) SetZoneSigningKeyAlgorithm(v string)`

SetZoneSigningKeyAlgorithm sets ZoneSigningKeyAlgorithm field to given value.


### GetZoneSigningKeyLength

`func (o *DNSSECSigningPolicyPostRequestBody) GetZoneSigningKeyLength() int32`

GetZoneSigningKeyLength returns the ZoneSigningKeyLength field if non-nil, zero value otherwise.

### GetZoneSigningKeyLengthOk

`func (o *DNSSECSigningPolicyPostRequestBody) GetZoneSigningKeyLengthOk() (*int32, bool)`

GetZoneSigningKeyLengthOk returns a tuple with the ZoneSigningKeyLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneSigningKeyLength

`func (o *DNSSECSigningPolicyPostRequestBody) SetZoneSigningKeyLength(v int32)`

SetZoneSigningKeyLength sets ZoneSigningKeyLength field to given value.


### GetZoneSigningKeyOverrideTtl

`func (o *DNSSECSigningPolicyPostRequestBody) GetZoneSigningKeyOverrideTtl() string`

GetZoneSigningKeyOverrideTtl returns the ZoneSigningKeyOverrideTtl field if non-nil, zero value otherwise.

### GetZoneSigningKeyOverrideTtlOk

`func (o *DNSSECSigningPolicyPostRequestBody) GetZoneSigningKeyOverrideTtlOk() (*string, bool)`

GetZoneSigningKeyOverrideTtlOk returns a tuple with the ZoneSigningKeyOverrideTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneSigningKeyOverrideTtl

`func (o *DNSSECSigningPolicyPostRequestBody) SetZoneSigningKeyOverrideTtl(v string)`

SetZoneSigningKeyOverrideTtl sets ZoneSigningKeyOverrideTtl field to given value.

### HasZoneSigningKeyOverrideTtl

`func (o *DNSSECSigningPolicyPostRequestBody) HasZoneSigningKeyOverrideTtl() bool`

HasZoneSigningKeyOverrideTtl returns a boolean if a field has been set.

### GetZoneSigningKeyValidityPeriod

`func (o *DNSSECSigningPolicyPostRequestBody) GetZoneSigningKeyValidityPeriod() string`

GetZoneSigningKeyValidityPeriod returns the ZoneSigningKeyValidityPeriod field if non-nil, zero value otherwise.

### GetZoneSigningKeyValidityPeriodOk

`func (o *DNSSECSigningPolicyPostRequestBody) GetZoneSigningKeyValidityPeriodOk() (*string, bool)`

GetZoneSigningKeyValidityPeriodOk returns a tuple with the ZoneSigningKeyValidityPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneSigningKeyValidityPeriod

`func (o *DNSSECSigningPolicyPostRequestBody) SetZoneSigningKeyValidityPeriod(v string)`

SetZoneSigningKeyValidityPeriod sets ZoneSigningKeyValidityPeriod field to given value.


### GetZoneSigningKeyOverlapPeriod

`func (o *DNSSECSigningPolicyPostRequestBody) GetZoneSigningKeyOverlapPeriod() string`

GetZoneSigningKeyOverlapPeriod returns the ZoneSigningKeyOverlapPeriod field if non-nil, zero value otherwise.

### GetZoneSigningKeyOverlapPeriodOk

`func (o *DNSSECSigningPolicyPostRequestBody) GetZoneSigningKeyOverlapPeriodOk() (*string, bool)`

GetZoneSigningKeyOverlapPeriodOk returns a tuple with the ZoneSigningKeyOverlapPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneSigningKeyOverlapPeriod

`func (o *DNSSECSigningPolicyPostRequestBody) SetZoneSigningKeyOverlapPeriod(v string)`

SetZoneSigningKeyOverlapPeriod sets ZoneSigningKeyOverlapPeriod field to given value.


### GetZoneSigningKeyRolloverMethod

`func (o *DNSSECSigningPolicyPostRequestBody) GetZoneSigningKeyRolloverMethod() string`

GetZoneSigningKeyRolloverMethod returns the ZoneSigningKeyRolloverMethod field if non-nil, zero value otherwise.

### GetZoneSigningKeyRolloverMethodOk

`func (o *DNSSECSigningPolicyPostRequestBody) GetZoneSigningKeyRolloverMethodOk() (*string, bool)`

GetZoneSigningKeyRolloverMethodOk returns a tuple with the ZoneSigningKeyRolloverMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneSigningKeyRolloverMethod

`func (o *DNSSECSigningPolicyPostRequestBody) SetZoneSigningKeyRolloverMethod(v string)`

SetZoneSigningKeyRolloverMethod sets ZoneSigningKeyRolloverMethod field to given value.


### GetZoneSigningKeySigningPeriod

`func (o *DNSSECSigningPolicyPostRequestBody) GetZoneSigningKeySigningPeriod() string`

GetZoneSigningKeySigningPeriod returns the ZoneSigningKeySigningPeriod field if non-nil, zero value otherwise.

### GetZoneSigningKeySigningPeriodOk

`func (o *DNSSECSigningPolicyPostRequestBody) GetZoneSigningKeySigningPeriodOk() (*string, bool)`

GetZoneSigningKeySigningPeriodOk returns a tuple with the ZoneSigningKeySigningPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneSigningKeySigningPeriod

`func (o *DNSSECSigningPolicyPostRequestBody) SetZoneSigningKeySigningPeriod(v string)`

SetZoneSigningKeySigningPeriod sets ZoneSigningKeySigningPeriod field to given value.

### HasZoneSigningKeySigningPeriod

`func (o *DNSSECSigningPolicyPostRequestBody) HasZoneSigningKeySigningPeriod() bool`

HasZoneSigningKeySigningPeriod returns a boolean if a field has been set.

### GetZoneSigningKeyProtectionType

`func (o *DNSSECSigningPolicyPostRequestBody) GetZoneSigningKeyProtectionType() string`

GetZoneSigningKeyProtectionType returns the ZoneSigningKeyProtectionType field if non-nil, zero value otherwise.

### GetZoneSigningKeyProtectionTypeOk

`func (o *DNSSECSigningPolicyPostRequestBody) GetZoneSigningKeyProtectionTypeOk() (*string, bool)`

GetZoneSigningKeyProtectionTypeOk returns a tuple with the ZoneSigningKeyProtectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneSigningKeyProtectionType

`func (o *DNSSECSigningPolicyPostRequestBody) SetZoneSigningKeyProtectionType(v string)`

SetZoneSigningKeyProtectionType sets ZoneSigningKeyProtectionType field to given value.

### HasZoneSigningKeyProtectionType

`func (o *DNSSECSigningPolicyPostRequestBody) HasZoneSigningKeyProtectionType() bool`

HasZoneSigningKeyProtectionType returns a boolean if a field has been set.

### GetKeySigningKeyAlgorithm

`func (o *DNSSECSigningPolicyPostRequestBody) GetKeySigningKeyAlgorithm() string`

GetKeySigningKeyAlgorithm returns the KeySigningKeyAlgorithm field if non-nil, zero value otherwise.

### GetKeySigningKeyAlgorithmOk

`func (o *DNSSECSigningPolicyPostRequestBody) GetKeySigningKeyAlgorithmOk() (*string, bool)`

GetKeySigningKeyAlgorithmOk returns a tuple with the KeySigningKeyAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySigningKeyAlgorithm

`func (o *DNSSECSigningPolicyPostRequestBody) SetKeySigningKeyAlgorithm(v string)`

SetKeySigningKeyAlgorithm sets KeySigningKeyAlgorithm field to given value.


### GetKeySigningKeyLength

`func (o *DNSSECSigningPolicyPostRequestBody) GetKeySigningKeyLength() int32`

GetKeySigningKeyLength returns the KeySigningKeyLength field if non-nil, zero value otherwise.

### GetKeySigningKeyLengthOk

`func (o *DNSSECSigningPolicyPostRequestBody) GetKeySigningKeyLengthOk() (*int32, bool)`

GetKeySigningKeyLengthOk returns a tuple with the KeySigningKeyLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySigningKeyLength

`func (o *DNSSECSigningPolicyPostRequestBody) SetKeySigningKeyLength(v int32)`

SetKeySigningKeyLength sets KeySigningKeyLength field to given value.


### GetKeySigningKeyOverrideTtl

`func (o *DNSSECSigningPolicyPostRequestBody) GetKeySigningKeyOverrideTtl() string`

GetKeySigningKeyOverrideTtl returns the KeySigningKeyOverrideTtl field if non-nil, zero value otherwise.

### GetKeySigningKeyOverrideTtlOk

`func (o *DNSSECSigningPolicyPostRequestBody) GetKeySigningKeyOverrideTtlOk() (*string, bool)`

GetKeySigningKeyOverrideTtlOk returns a tuple with the KeySigningKeyOverrideTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySigningKeyOverrideTtl

`func (o *DNSSECSigningPolicyPostRequestBody) SetKeySigningKeyOverrideTtl(v string)`

SetKeySigningKeyOverrideTtl sets KeySigningKeyOverrideTtl field to given value.

### HasKeySigningKeyOverrideTtl

`func (o *DNSSECSigningPolicyPostRequestBody) HasKeySigningKeyOverrideTtl() bool`

HasKeySigningKeyOverrideTtl returns a boolean if a field has been set.

### GetKeySigningKeyValidityPeriod

`func (o *DNSSECSigningPolicyPostRequestBody) GetKeySigningKeyValidityPeriod() string`

GetKeySigningKeyValidityPeriod returns the KeySigningKeyValidityPeriod field if non-nil, zero value otherwise.

### GetKeySigningKeyValidityPeriodOk

`func (o *DNSSECSigningPolicyPostRequestBody) GetKeySigningKeyValidityPeriodOk() (*string, bool)`

GetKeySigningKeyValidityPeriodOk returns a tuple with the KeySigningKeyValidityPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySigningKeyValidityPeriod

`func (o *DNSSECSigningPolicyPostRequestBody) SetKeySigningKeyValidityPeriod(v string)`

SetKeySigningKeyValidityPeriod sets KeySigningKeyValidityPeriod field to given value.


### GetKeySigningKeyOverlapPeriod

`func (o *DNSSECSigningPolicyPostRequestBody) GetKeySigningKeyOverlapPeriod() string`

GetKeySigningKeyOverlapPeriod returns the KeySigningKeyOverlapPeriod field if non-nil, zero value otherwise.

### GetKeySigningKeyOverlapPeriodOk

`func (o *DNSSECSigningPolicyPostRequestBody) GetKeySigningKeyOverlapPeriodOk() (*string, bool)`

GetKeySigningKeyOverlapPeriodOk returns a tuple with the KeySigningKeyOverlapPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySigningKeyOverlapPeriod

`func (o *DNSSECSigningPolicyPostRequestBody) SetKeySigningKeyOverlapPeriod(v string)`

SetKeySigningKeyOverlapPeriod sets KeySigningKeyOverlapPeriod field to given value.


### GetKeySigningKeyRolloverMethod

`func (o *DNSSECSigningPolicyPostRequestBody) GetKeySigningKeyRolloverMethod() string`

GetKeySigningKeyRolloverMethod returns the KeySigningKeyRolloverMethod field if non-nil, zero value otherwise.

### GetKeySigningKeyRolloverMethodOk

`func (o *DNSSECSigningPolicyPostRequestBody) GetKeySigningKeyRolloverMethodOk() (*string, bool)`

GetKeySigningKeyRolloverMethodOk returns a tuple with the KeySigningKeyRolloverMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySigningKeyRolloverMethod

`func (o *DNSSECSigningPolicyPostRequestBody) SetKeySigningKeyRolloverMethod(v string)`

SetKeySigningKeyRolloverMethod sets KeySigningKeyRolloverMethod field to given value.


### GetKeySigningKeySigningPeriod

`func (o *DNSSECSigningPolicyPostRequestBody) GetKeySigningKeySigningPeriod() string`

GetKeySigningKeySigningPeriod returns the KeySigningKeySigningPeriod field if non-nil, zero value otherwise.

### GetKeySigningKeySigningPeriodOk

`func (o *DNSSECSigningPolicyPostRequestBody) GetKeySigningKeySigningPeriodOk() (*string, bool)`

GetKeySigningKeySigningPeriodOk returns a tuple with the KeySigningKeySigningPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySigningKeySigningPeriod

`func (o *DNSSECSigningPolicyPostRequestBody) SetKeySigningKeySigningPeriod(v string)`

SetKeySigningKeySigningPeriod sets KeySigningKeySigningPeriod field to given value.

### HasKeySigningKeySigningPeriod

`func (o *DNSSECSigningPolicyPostRequestBody) HasKeySigningKeySigningPeriod() bool`

HasKeySigningKeySigningPeriod returns a boolean if a field has been set.

### GetKeySigningKeyProtectionType

`func (o *DNSSECSigningPolicyPostRequestBody) GetKeySigningKeyProtectionType() string`

GetKeySigningKeyProtectionType returns the KeySigningKeyProtectionType field if non-nil, zero value otherwise.

### GetKeySigningKeyProtectionTypeOk

`func (o *DNSSECSigningPolicyPostRequestBody) GetKeySigningKeyProtectionTypeOk() (*string, bool)`

GetKeySigningKeyProtectionTypeOk returns a tuple with the KeySigningKeyProtectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySigningKeyProtectionType

`func (o *DNSSECSigningPolicyPostRequestBody) SetKeySigningKeyProtectionType(v string)`

SetKeySigningKeyProtectionType sets KeySigningKeyProtectionType field to given value.

### HasKeySigningKeyProtectionType

`func (o *DNSSECSigningPolicyPostRequestBody) HasKeySigningKeyProtectionType() bool`

HasKeySigningKeyProtectionType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


