# PutZoneDeclarationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | **string** | The resource type. | 
**Name** | Pointer to **string** | The name of the resource. | [optional] [readonly] 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**Configuration** | Pointer to [**InlinedConfiguration**](InlinedConfiguration.md) |  | [optional] [readonly] 
**PrimaryDnsServerAddress** | **string** | The IP address of the zone&#39;s primary DNS server. | 
**SecondaryDnsServerAddress** | Pointer to **string** | The IP address of the zone&#39;s secondary DNS server. | [optional] 
**DynamicUpdateSigningKeyType** | **string** | The type of key used to sign DDNS updates. | 
**DynamicUpdateSigningKey** | Pointer to [**InlinedSigningKey**](InlinedSigningKey.md) |  | [optional] 
**DynamicUpdateSigningKeyRealm** | Pointer to [**InlinedKerberosRealm**](InlinedKerberosRealm.md) |  | [optional] 
**Prefix** | [**DHCPReverseZoneAllOfPrefix**](DHCPReverseZoneAllOfPrefix.md) |  | 

## Methods

### NewPutZoneDeclarationRequest

`func NewPutZoneDeclarationRequest(type_ string, primaryDnsServerAddress string, dynamicUpdateSigningKeyType string, prefix DHCPReverseZoneAllOfPrefix, ) *PutZoneDeclarationRequest`

NewPutZoneDeclarationRequest instantiates a new PutZoneDeclarationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutZoneDeclarationRequestWithDefaults

`func NewPutZoneDeclarationRequestWithDefaults() *PutZoneDeclarationRequest`

NewPutZoneDeclarationRequestWithDefaults instantiates a new PutZoneDeclarationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PutZoneDeclarationRequest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PutZoneDeclarationRequest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PutZoneDeclarationRequest) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *PutZoneDeclarationRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *PutZoneDeclarationRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PutZoneDeclarationRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PutZoneDeclarationRequest) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *PutZoneDeclarationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PutZoneDeclarationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PutZoneDeclarationRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PutZoneDeclarationRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *PutZoneDeclarationRequest) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *PutZoneDeclarationRequest) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *PutZoneDeclarationRequest) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *PutZoneDeclarationRequest) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetConfiguration

`func (o *PutZoneDeclarationRequest) GetConfiguration() InlinedConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *PutZoneDeclarationRequest) GetConfigurationOk() (*InlinedConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *PutZoneDeclarationRequest) SetConfiguration(v InlinedConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *PutZoneDeclarationRequest) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetPrimaryDnsServerAddress

`func (o *PutZoneDeclarationRequest) GetPrimaryDnsServerAddress() string`

GetPrimaryDnsServerAddress returns the PrimaryDnsServerAddress field if non-nil, zero value otherwise.

### GetPrimaryDnsServerAddressOk

`func (o *PutZoneDeclarationRequest) GetPrimaryDnsServerAddressOk() (*string, bool)`

GetPrimaryDnsServerAddressOk returns a tuple with the PrimaryDnsServerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryDnsServerAddress

`func (o *PutZoneDeclarationRequest) SetPrimaryDnsServerAddress(v string)`

SetPrimaryDnsServerAddress sets PrimaryDnsServerAddress field to given value.


### GetSecondaryDnsServerAddress

`func (o *PutZoneDeclarationRequest) GetSecondaryDnsServerAddress() string`

GetSecondaryDnsServerAddress returns the SecondaryDnsServerAddress field if non-nil, zero value otherwise.

### GetSecondaryDnsServerAddressOk

`func (o *PutZoneDeclarationRequest) GetSecondaryDnsServerAddressOk() (*string, bool)`

GetSecondaryDnsServerAddressOk returns a tuple with the SecondaryDnsServerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryDnsServerAddress

`func (o *PutZoneDeclarationRequest) SetSecondaryDnsServerAddress(v string)`

SetSecondaryDnsServerAddress sets SecondaryDnsServerAddress field to given value.

### HasSecondaryDnsServerAddress

`func (o *PutZoneDeclarationRequest) HasSecondaryDnsServerAddress() bool`

HasSecondaryDnsServerAddress returns a boolean if a field has been set.

### GetDynamicUpdateSigningKeyType

`func (o *PutZoneDeclarationRequest) GetDynamicUpdateSigningKeyType() string`

GetDynamicUpdateSigningKeyType returns the DynamicUpdateSigningKeyType field if non-nil, zero value otherwise.

### GetDynamicUpdateSigningKeyTypeOk

`func (o *PutZoneDeclarationRequest) GetDynamicUpdateSigningKeyTypeOk() (*string, bool)`

GetDynamicUpdateSigningKeyTypeOk returns a tuple with the DynamicUpdateSigningKeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicUpdateSigningKeyType

`func (o *PutZoneDeclarationRequest) SetDynamicUpdateSigningKeyType(v string)`

SetDynamicUpdateSigningKeyType sets DynamicUpdateSigningKeyType field to given value.


### GetDynamicUpdateSigningKey

`func (o *PutZoneDeclarationRequest) GetDynamicUpdateSigningKey() InlinedSigningKey`

GetDynamicUpdateSigningKey returns the DynamicUpdateSigningKey field if non-nil, zero value otherwise.

### GetDynamicUpdateSigningKeyOk

`func (o *PutZoneDeclarationRequest) GetDynamicUpdateSigningKeyOk() (*InlinedSigningKey, bool)`

GetDynamicUpdateSigningKeyOk returns a tuple with the DynamicUpdateSigningKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicUpdateSigningKey

`func (o *PutZoneDeclarationRequest) SetDynamicUpdateSigningKey(v InlinedSigningKey)`

SetDynamicUpdateSigningKey sets DynamicUpdateSigningKey field to given value.

### HasDynamicUpdateSigningKey

`func (o *PutZoneDeclarationRequest) HasDynamicUpdateSigningKey() bool`

HasDynamicUpdateSigningKey returns a boolean if a field has been set.

### GetDynamicUpdateSigningKeyRealm

`func (o *PutZoneDeclarationRequest) GetDynamicUpdateSigningKeyRealm() InlinedKerberosRealm`

GetDynamicUpdateSigningKeyRealm returns the DynamicUpdateSigningKeyRealm field if non-nil, zero value otherwise.

### GetDynamicUpdateSigningKeyRealmOk

`func (o *PutZoneDeclarationRequest) GetDynamicUpdateSigningKeyRealmOk() (*InlinedKerberosRealm, bool)`

GetDynamicUpdateSigningKeyRealmOk returns a tuple with the DynamicUpdateSigningKeyRealm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicUpdateSigningKeyRealm

`func (o *PutZoneDeclarationRequest) SetDynamicUpdateSigningKeyRealm(v InlinedKerberosRealm)`

SetDynamicUpdateSigningKeyRealm sets DynamicUpdateSigningKeyRealm field to given value.

### HasDynamicUpdateSigningKeyRealm

`func (o *PutZoneDeclarationRequest) HasDynamicUpdateSigningKeyRealm() bool`

HasDynamicUpdateSigningKeyRealm returns a boolean if a field has been set.

### GetPrefix

`func (o *PutZoneDeclarationRequest) GetPrefix() DHCPReverseZoneAllOfPrefix`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *PutZoneDeclarationRequest) GetPrefixOk() (*DHCPReverseZoneAllOfPrefix, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *PutZoneDeclarationRequest) SetPrefix(v DHCPReverseZoneAllOfPrefix)`

SetPrefix sets Prefix field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


