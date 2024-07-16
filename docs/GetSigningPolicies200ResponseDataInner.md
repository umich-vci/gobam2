# GetSigningPolicies200ResponseDataInner

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
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 

## Methods

### NewGetSigningPolicies200ResponseDataInner

`func NewGetSigningPolicies200ResponseDataInner() *GetSigningPolicies200ResponseDataInner`

NewGetSigningPolicies200ResponseDataInner instantiates a new GetSigningPolicies200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSigningPolicies200ResponseDataInnerWithDefaults

`func NewGetSigningPolicies200ResponseDataInnerWithDefaults() *GetSigningPolicies200ResponseDataInner`

NewGetSigningPolicies200ResponseDataInnerWithDefaults instantiates a new GetSigningPolicies200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetSigningPolicies200ResponseDataInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetSigningPolicies200ResponseDataInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetSigningPolicies200ResponseDataInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetSigningPolicies200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GetSigningPolicies200ResponseDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetSigningPolicies200ResponseDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetSigningPolicies200ResponseDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetSigningPolicies200ResponseDataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *GetSigningPolicies200ResponseDataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetSigningPolicies200ResponseDataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetSigningPolicies200ResponseDataInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetSigningPolicies200ResponseDataInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *GetSigningPolicies200ResponseDataInner) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *GetSigningPolicies200ResponseDataInner) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *GetSigningPolicies200ResponseDataInner) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *GetSigningPolicies200ResponseDataInner) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetKeyProvider

`func (o *GetSigningPolicies200ResponseDataInner) GetKeyProvider() string`

GetKeyProvider returns the KeyProvider field if non-nil, zero value otherwise.

### GetKeyProviderOk

`func (o *GetSigningPolicies200ResponseDataInner) GetKeyProviderOk() (*string, bool)`

GetKeyProviderOk returns a tuple with the KeyProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyProvider

`func (o *GetSigningPolicies200ResponseDataInner) SetKeyProvider(v string)`

SetKeyProvider sets KeyProvider field to given value.

### HasKeyProvider

`func (o *GetSigningPolicies200ResponseDataInner) HasKeyProvider() bool`

HasKeyProvider returns a boolean if a field has been set.

### GetSignatureDigestAlgorithm

`func (o *GetSigningPolicies200ResponseDataInner) GetSignatureDigestAlgorithm() string`

GetSignatureDigestAlgorithm returns the SignatureDigestAlgorithm field if non-nil, zero value otherwise.

### GetSignatureDigestAlgorithmOk

`func (o *GetSigningPolicies200ResponseDataInner) GetSignatureDigestAlgorithmOk() (*string, bool)`

GetSignatureDigestAlgorithmOk returns a tuple with the SignatureDigestAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureDigestAlgorithm

`func (o *GetSigningPolicies200ResponseDataInner) SetSignatureDigestAlgorithm(v string)`

SetSignatureDigestAlgorithm sets SignatureDigestAlgorithm field to given value.

### HasSignatureDigestAlgorithm

`func (o *GetSigningPolicies200ResponseDataInner) HasSignatureDigestAlgorithm() bool`

HasSignatureDigestAlgorithm returns a boolean if a field has been set.

### GetSignatureValidityPeriod

`func (o *GetSigningPolicies200ResponseDataInner) GetSignatureValidityPeriod() string`

GetSignatureValidityPeriod returns the SignatureValidityPeriod field if non-nil, zero value otherwise.

### GetSignatureValidityPeriodOk

`func (o *GetSigningPolicies200ResponseDataInner) GetSignatureValidityPeriodOk() (*string, bool)`

GetSignatureValidityPeriodOk returns a tuple with the SignatureValidityPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureValidityPeriod

`func (o *GetSigningPolicies200ResponseDataInner) SetSignatureValidityPeriod(v string)`

SetSignatureValidityPeriod sets SignatureValidityPeriod field to given value.

### HasSignatureValidityPeriod

`func (o *GetSigningPolicies200ResponseDataInner) HasSignatureValidityPeriod() bool`

HasSignatureValidityPeriod returns a boolean if a field has been set.

### GetSignatureResigningPeriod

`func (o *GetSigningPolicies200ResponseDataInner) GetSignatureResigningPeriod() string`

GetSignatureResigningPeriod returns the SignatureResigningPeriod field if non-nil, zero value otherwise.

### GetSignatureResigningPeriodOk

`func (o *GetSigningPolicies200ResponseDataInner) GetSignatureResigningPeriodOk() (*string, bool)`

GetSignatureResigningPeriodOk returns a tuple with the SignatureResigningPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureResigningPeriod

`func (o *GetSigningPolicies200ResponseDataInner) SetSignatureResigningPeriod(v string)`

SetSignatureResigningPeriod sets SignatureResigningPeriod field to given value.

### HasSignatureResigningPeriod

`func (o *GetSigningPolicies200ResponseDataInner) HasSignatureResigningPeriod() bool`

HasSignatureResigningPeriod returns a boolean if a field has been set.

### GetZoneSigningKeyAlgorithm

`func (o *GetSigningPolicies200ResponseDataInner) GetZoneSigningKeyAlgorithm() string`

GetZoneSigningKeyAlgorithm returns the ZoneSigningKeyAlgorithm field if non-nil, zero value otherwise.

### GetZoneSigningKeyAlgorithmOk

`func (o *GetSigningPolicies200ResponseDataInner) GetZoneSigningKeyAlgorithmOk() (*string, bool)`

GetZoneSigningKeyAlgorithmOk returns a tuple with the ZoneSigningKeyAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneSigningKeyAlgorithm

`func (o *GetSigningPolicies200ResponseDataInner) SetZoneSigningKeyAlgorithm(v string)`

SetZoneSigningKeyAlgorithm sets ZoneSigningKeyAlgorithm field to given value.

### HasZoneSigningKeyAlgorithm

`func (o *GetSigningPolicies200ResponseDataInner) HasZoneSigningKeyAlgorithm() bool`

HasZoneSigningKeyAlgorithm returns a boolean if a field has been set.

### GetZoneSigningKeyLength

`func (o *GetSigningPolicies200ResponseDataInner) GetZoneSigningKeyLength() int32`

GetZoneSigningKeyLength returns the ZoneSigningKeyLength field if non-nil, zero value otherwise.

### GetZoneSigningKeyLengthOk

`func (o *GetSigningPolicies200ResponseDataInner) GetZoneSigningKeyLengthOk() (*int32, bool)`

GetZoneSigningKeyLengthOk returns a tuple with the ZoneSigningKeyLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneSigningKeyLength

`func (o *GetSigningPolicies200ResponseDataInner) SetZoneSigningKeyLength(v int32)`

SetZoneSigningKeyLength sets ZoneSigningKeyLength field to given value.

### HasZoneSigningKeyLength

`func (o *GetSigningPolicies200ResponseDataInner) HasZoneSigningKeyLength() bool`

HasZoneSigningKeyLength returns a boolean if a field has been set.

### GetZoneSigningKeyOverrideTtl

`func (o *GetSigningPolicies200ResponseDataInner) GetZoneSigningKeyOverrideTtl() string`

GetZoneSigningKeyOverrideTtl returns the ZoneSigningKeyOverrideTtl field if non-nil, zero value otherwise.

### GetZoneSigningKeyOverrideTtlOk

`func (o *GetSigningPolicies200ResponseDataInner) GetZoneSigningKeyOverrideTtlOk() (*string, bool)`

GetZoneSigningKeyOverrideTtlOk returns a tuple with the ZoneSigningKeyOverrideTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneSigningKeyOverrideTtl

`func (o *GetSigningPolicies200ResponseDataInner) SetZoneSigningKeyOverrideTtl(v string)`

SetZoneSigningKeyOverrideTtl sets ZoneSigningKeyOverrideTtl field to given value.

### HasZoneSigningKeyOverrideTtl

`func (o *GetSigningPolicies200ResponseDataInner) HasZoneSigningKeyOverrideTtl() bool`

HasZoneSigningKeyOverrideTtl returns a boolean if a field has been set.

### GetZoneSigningKeyValidityPeriod

`func (o *GetSigningPolicies200ResponseDataInner) GetZoneSigningKeyValidityPeriod() string`

GetZoneSigningKeyValidityPeriod returns the ZoneSigningKeyValidityPeriod field if non-nil, zero value otherwise.

### GetZoneSigningKeyValidityPeriodOk

`func (o *GetSigningPolicies200ResponseDataInner) GetZoneSigningKeyValidityPeriodOk() (*string, bool)`

GetZoneSigningKeyValidityPeriodOk returns a tuple with the ZoneSigningKeyValidityPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneSigningKeyValidityPeriod

`func (o *GetSigningPolicies200ResponseDataInner) SetZoneSigningKeyValidityPeriod(v string)`

SetZoneSigningKeyValidityPeriod sets ZoneSigningKeyValidityPeriod field to given value.

### HasZoneSigningKeyValidityPeriod

`func (o *GetSigningPolicies200ResponseDataInner) HasZoneSigningKeyValidityPeriod() bool`

HasZoneSigningKeyValidityPeriod returns a boolean if a field has been set.

### GetZoneSigningKeyOverlapPeriod

`func (o *GetSigningPolicies200ResponseDataInner) GetZoneSigningKeyOverlapPeriod() string`

GetZoneSigningKeyOverlapPeriod returns the ZoneSigningKeyOverlapPeriod field if non-nil, zero value otherwise.

### GetZoneSigningKeyOverlapPeriodOk

`func (o *GetSigningPolicies200ResponseDataInner) GetZoneSigningKeyOverlapPeriodOk() (*string, bool)`

GetZoneSigningKeyOverlapPeriodOk returns a tuple with the ZoneSigningKeyOverlapPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneSigningKeyOverlapPeriod

`func (o *GetSigningPolicies200ResponseDataInner) SetZoneSigningKeyOverlapPeriod(v string)`

SetZoneSigningKeyOverlapPeriod sets ZoneSigningKeyOverlapPeriod field to given value.

### HasZoneSigningKeyOverlapPeriod

`func (o *GetSigningPolicies200ResponseDataInner) HasZoneSigningKeyOverlapPeriod() bool`

HasZoneSigningKeyOverlapPeriod returns a boolean if a field has been set.

### GetZoneSigningKeyRolloverMethod

`func (o *GetSigningPolicies200ResponseDataInner) GetZoneSigningKeyRolloverMethod() string`

GetZoneSigningKeyRolloverMethod returns the ZoneSigningKeyRolloverMethod field if non-nil, zero value otherwise.

### GetZoneSigningKeyRolloverMethodOk

`func (o *GetSigningPolicies200ResponseDataInner) GetZoneSigningKeyRolloverMethodOk() (*string, bool)`

GetZoneSigningKeyRolloverMethodOk returns a tuple with the ZoneSigningKeyRolloverMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneSigningKeyRolloverMethod

`func (o *GetSigningPolicies200ResponseDataInner) SetZoneSigningKeyRolloverMethod(v string)`

SetZoneSigningKeyRolloverMethod sets ZoneSigningKeyRolloverMethod field to given value.

### HasZoneSigningKeyRolloverMethod

`func (o *GetSigningPolicies200ResponseDataInner) HasZoneSigningKeyRolloverMethod() bool`

HasZoneSigningKeyRolloverMethod returns a boolean if a field has been set.

### GetZoneSigningKeySigningPeriod

`func (o *GetSigningPolicies200ResponseDataInner) GetZoneSigningKeySigningPeriod() string`

GetZoneSigningKeySigningPeriod returns the ZoneSigningKeySigningPeriod field if non-nil, zero value otherwise.

### GetZoneSigningKeySigningPeriodOk

`func (o *GetSigningPolicies200ResponseDataInner) GetZoneSigningKeySigningPeriodOk() (*string, bool)`

GetZoneSigningKeySigningPeriodOk returns a tuple with the ZoneSigningKeySigningPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneSigningKeySigningPeriod

`func (o *GetSigningPolicies200ResponseDataInner) SetZoneSigningKeySigningPeriod(v string)`

SetZoneSigningKeySigningPeriod sets ZoneSigningKeySigningPeriod field to given value.

### HasZoneSigningKeySigningPeriod

`func (o *GetSigningPolicies200ResponseDataInner) HasZoneSigningKeySigningPeriod() bool`

HasZoneSigningKeySigningPeriod returns a boolean if a field has been set.

### GetZoneSigningKeyProtectionType

`func (o *GetSigningPolicies200ResponseDataInner) GetZoneSigningKeyProtectionType() string`

GetZoneSigningKeyProtectionType returns the ZoneSigningKeyProtectionType field if non-nil, zero value otherwise.

### GetZoneSigningKeyProtectionTypeOk

`func (o *GetSigningPolicies200ResponseDataInner) GetZoneSigningKeyProtectionTypeOk() (*string, bool)`

GetZoneSigningKeyProtectionTypeOk returns a tuple with the ZoneSigningKeyProtectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneSigningKeyProtectionType

`func (o *GetSigningPolicies200ResponseDataInner) SetZoneSigningKeyProtectionType(v string)`

SetZoneSigningKeyProtectionType sets ZoneSigningKeyProtectionType field to given value.

### HasZoneSigningKeyProtectionType

`func (o *GetSigningPolicies200ResponseDataInner) HasZoneSigningKeyProtectionType() bool`

HasZoneSigningKeyProtectionType returns a boolean if a field has been set.

### GetKeySigningKeyAlgorithm

`func (o *GetSigningPolicies200ResponseDataInner) GetKeySigningKeyAlgorithm() string`

GetKeySigningKeyAlgorithm returns the KeySigningKeyAlgorithm field if non-nil, zero value otherwise.

### GetKeySigningKeyAlgorithmOk

`func (o *GetSigningPolicies200ResponseDataInner) GetKeySigningKeyAlgorithmOk() (*string, bool)`

GetKeySigningKeyAlgorithmOk returns a tuple with the KeySigningKeyAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySigningKeyAlgorithm

`func (o *GetSigningPolicies200ResponseDataInner) SetKeySigningKeyAlgorithm(v string)`

SetKeySigningKeyAlgorithm sets KeySigningKeyAlgorithm field to given value.

### HasKeySigningKeyAlgorithm

`func (o *GetSigningPolicies200ResponseDataInner) HasKeySigningKeyAlgorithm() bool`

HasKeySigningKeyAlgorithm returns a boolean if a field has been set.

### GetKeySigningKeyLength

`func (o *GetSigningPolicies200ResponseDataInner) GetKeySigningKeyLength() int32`

GetKeySigningKeyLength returns the KeySigningKeyLength field if non-nil, zero value otherwise.

### GetKeySigningKeyLengthOk

`func (o *GetSigningPolicies200ResponseDataInner) GetKeySigningKeyLengthOk() (*int32, bool)`

GetKeySigningKeyLengthOk returns a tuple with the KeySigningKeyLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySigningKeyLength

`func (o *GetSigningPolicies200ResponseDataInner) SetKeySigningKeyLength(v int32)`

SetKeySigningKeyLength sets KeySigningKeyLength field to given value.

### HasKeySigningKeyLength

`func (o *GetSigningPolicies200ResponseDataInner) HasKeySigningKeyLength() bool`

HasKeySigningKeyLength returns a boolean if a field has been set.

### GetKeySigningKeyOverrideTtl

`func (o *GetSigningPolicies200ResponseDataInner) GetKeySigningKeyOverrideTtl() string`

GetKeySigningKeyOverrideTtl returns the KeySigningKeyOverrideTtl field if non-nil, zero value otherwise.

### GetKeySigningKeyOverrideTtlOk

`func (o *GetSigningPolicies200ResponseDataInner) GetKeySigningKeyOverrideTtlOk() (*string, bool)`

GetKeySigningKeyOverrideTtlOk returns a tuple with the KeySigningKeyOverrideTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySigningKeyOverrideTtl

`func (o *GetSigningPolicies200ResponseDataInner) SetKeySigningKeyOverrideTtl(v string)`

SetKeySigningKeyOverrideTtl sets KeySigningKeyOverrideTtl field to given value.

### HasKeySigningKeyOverrideTtl

`func (o *GetSigningPolicies200ResponseDataInner) HasKeySigningKeyOverrideTtl() bool`

HasKeySigningKeyOverrideTtl returns a boolean if a field has been set.

### GetKeySigningKeyValidityPeriod

`func (o *GetSigningPolicies200ResponseDataInner) GetKeySigningKeyValidityPeriod() string`

GetKeySigningKeyValidityPeriod returns the KeySigningKeyValidityPeriod field if non-nil, zero value otherwise.

### GetKeySigningKeyValidityPeriodOk

`func (o *GetSigningPolicies200ResponseDataInner) GetKeySigningKeyValidityPeriodOk() (*string, bool)`

GetKeySigningKeyValidityPeriodOk returns a tuple with the KeySigningKeyValidityPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySigningKeyValidityPeriod

`func (o *GetSigningPolicies200ResponseDataInner) SetKeySigningKeyValidityPeriod(v string)`

SetKeySigningKeyValidityPeriod sets KeySigningKeyValidityPeriod field to given value.

### HasKeySigningKeyValidityPeriod

`func (o *GetSigningPolicies200ResponseDataInner) HasKeySigningKeyValidityPeriod() bool`

HasKeySigningKeyValidityPeriod returns a boolean if a field has been set.

### GetKeySigningKeyOverlapPeriod

`func (o *GetSigningPolicies200ResponseDataInner) GetKeySigningKeyOverlapPeriod() string`

GetKeySigningKeyOverlapPeriod returns the KeySigningKeyOverlapPeriod field if non-nil, zero value otherwise.

### GetKeySigningKeyOverlapPeriodOk

`func (o *GetSigningPolicies200ResponseDataInner) GetKeySigningKeyOverlapPeriodOk() (*string, bool)`

GetKeySigningKeyOverlapPeriodOk returns a tuple with the KeySigningKeyOverlapPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySigningKeyOverlapPeriod

`func (o *GetSigningPolicies200ResponseDataInner) SetKeySigningKeyOverlapPeriod(v string)`

SetKeySigningKeyOverlapPeriod sets KeySigningKeyOverlapPeriod field to given value.

### HasKeySigningKeyOverlapPeriod

`func (o *GetSigningPolicies200ResponseDataInner) HasKeySigningKeyOverlapPeriod() bool`

HasKeySigningKeyOverlapPeriod returns a boolean if a field has been set.

### GetKeySigningKeyRolloverMethod

`func (o *GetSigningPolicies200ResponseDataInner) GetKeySigningKeyRolloverMethod() string`

GetKeySigningKeyRolloverMethod returns the KeySigningKeyRolloverMethod field if non-nil, zero value otherwise.

### GetKeySigningKeyRolloverMethodOk

`func (o *GetSigningPolicies200ResponseDataInner) GetKeySigningKeyRolloverMethodOk() (*string, bool)`

GetKeySigningKeyRolloverMethodOk returns a tuple with the KeySigningKeyRolloverMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySigningKeyRolloverMethod

`func (o *GetSigningPolicies200ResponseDataInner) SetKeySigningKeyRolloverMethod(v string)`

SetKeySigningKeyRolloverMethod sets KeySigningKeyRolloverMethod field to given value.

### HasKeySigningKeyRolloverMethod

`func (o *GetSigningPolicies200ResponseDataInner) HasKeySigningKeyRolloverMethod() bool`

HasKeySigningKeyRolloverMethod returns a boolean if a field has been set.

### GetKeySigningKeySigningPeriod

`func (o *GetSigningPolicies200ResponseDataInner) GetKeySigningKeySigningPeriod() string`

GetKeySigningKeySigningPeriod returns the KeySigningKeySigningPeriod field if non-nil, zero value otherwise.

### GetKeySigningKeySigningPeriodOk

`func (o *GetSigningPolicies200ResponseDataInner) GetKeySigningKeySigningPeriodOk() (*string, bool)`

GetKeySigningKeySigningPeriodOk returns a tuple with the KeySigningKeySigningPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySigningKeySigningPeriod

`func (o *GetSigningPolicies200ResponseDataInner) SetKeySigningKeySigningPeriod(v string)`

SetKeySigningKeySigningPeriod sets KeySigningKeySigningPeriod field to given value.

### HasKeySigningKeySigningPeriod

`func (o *GetSigningPolicies200ResponseDataInner) HasKeySigningKeySigningPeriod() bool`

HasKeySigningKeySigningPeriod returns a boolean if a field has been set.

### GetKeySigningKeyProtectionType

`func (o *GetSigningPolicies200ResponseDataInner) GetKeySigningKeyProtectionType() string`

GetKeySigningKeyProtectionType returns the KeySigningKeyProtectionType field if non-nil, zero value otherwise.

### GetKeySigningKeyProtectionTypeOk

`func (o *GetSigningPolicies200ResponseDataInner) GetKeySigningKeyProtectionTypeOk() (*string, bool)`

GetKeySigningKeyProtectionTypeOk returns a tuple with the KeySigningKeyProtectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySigningKeyProtectionType

`func (o *GetSigningPolicies200ResponseDataInner) SetKeySigningKeyProtectionType(v string)`

SetKeySigningKeyProtectionType sets KeySigningKeyProtectionType field to given value.

### HasKeySigningKeyProtectionType

`func (o *GetSigningPolicies200ResponseDataInner) HasKeySigningKeyProtectionType() bool`

HasKeySigningKeyProtectionType returns a boolean if a field has been set.

### GetLinks

`func (o *GetSigningPolicies200ResponseDataInner) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetSigningPolicies200ResponseDataInner) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetSigningPolicies200ResponseDataInner) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetSigningPolicies200ResponseDataInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


