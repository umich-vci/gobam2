# GetZoneDeclarations200Response1DataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | Pointer to **string** | The name of the resource. | [optional] [readonly] 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**Configuration** | Pointer to [**InlinedConfiguration**](InlinedConfiguration.md) |  | [optional] [readonly] 
**PrimaryDnsServerAddress** | Pointer to **string** | The IP address of the zone&#39;s primary DNS server. | [optional] 
**SecondaryDnsServerAddress** | Pointer to **string** | The IP address of the zone&#39;s secondary DNS server. | [optional] 
**DynamicUpdateSigningKeyType** | Pointer to **string** | The type of key used to sign DDNS updates. | [optional] 
**DynamicUpdateSigningKey** | Pointer to [**InlinedSigningKey**](InlinedSigningKey.md) |  | [optional] 
**DynamicUpdateSigningKeyRealm** | Pointer to [**InlinedKerberosRealm**](InlinedKerberosRealm.md) |  | [optional] 
**Prefix** | Pointer to [**DHCPReverseZoneAllOfPrefix**](DHCPReverseZoneAllOfPrefix.md) |  | [optional] 

## Methods

### NewGetZoneDeclarations200Response1DataInner

`func NewGetZoneDeclarations200Response1DataInner() *GetZoneDeclarations200Response1DataInner`

NewGetZoneDeclarations200Response1DataInner instantiates a new GetZoneDeclarations200Response1DataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetZoneDeclarations200Response1DataInnerWithDefaults

`func NewGetZoneDeclarations200Response1DataInnerWithDefaults() *GetZoneDeclarations200Response1DataInner`

NewGetZoneDeclarations200Response1DataInnerWithDefaults instantiates a new GetZoneDeclarations200Response1DataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetZoneDeclarations200Response1DataInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetZoneDeclarations200Response1DataInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetZoneDeclarations200Response1DataInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetZoneDeclarations200Response1DataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GetZoneDeclarations200Response1DataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetZoneDeclarations200Response1DataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetZoneDeclarations200Response1DataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetZoneDeclarations200Response1DataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *GetZoneDeclarations200Response1DataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetZoneDeclarations200Response1DataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetZoneDeclarations200Response1DataInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetZoneDeclarations200Response1DataInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *GetZoneDeclarations200Response1DataInner) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *GetZoneDeclarations200Response1DataInner) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *GetZoneDeclarations200Response1DataInner) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *GetZoneDeclarations200Response1DataInner) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetConfiguration

`func (o *GetZoneDeclarations200Response1DataInner) GetConfiguration() InlinedConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *GetZoneDeclarations200Response1DataInner) GetConfigurationOk() (*InlinedConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *GetZoneDeclarations200Response1DataInner) SetConfiguration(v InlinedConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *GetZoneDeclarations200Response1DataInner) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetPrimaryDnsServerAddress

`func (o *GetZoneDeclarations200Response1DataInner) GetPrimaryDnsServerAddress() string`

GetPrimaryDnsServerAddress returns the PrimaryDnsServerAddress field if non-nil, zero value otherwise.

### GetPrimaryDnsServerAddressOk

`func (o *GetZoneDeclarations200Response1DataInner) GetPrimaryDnsServerAddressOk() (*string, bool)`

GetPrimaryDnsServerAddressOk returns a tuple with the PrimaryDnsServerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryDnsServerAddress

`func (o *GetZoneDeclarations200Response1DataInner) SetPrimaryDnsServerAddress(v string)`

SetPrimaryDnsServerAddress sets PrimaryDnsServerAddress field to given value.

### HasPrimaryDnsServerAddress

`func (o *GetZoneDeclarations200Response1DataInner) HasPrimaryDnsServerAddress() bool`

HasPrimaryDnsServerAddress returns a boolean if a field has been set.

### GetSecondaryDnsServerAddress

`func (o *GetZoneDeclarations200Response1DataInner) GetSecondaryDnsServerAddress() string`

GetSecondaryDnsServerAddress returns the SecondaryDnsServerAddress field if non-nil, zero value otherwise.

### GetSecondaryDnsServerAddressOk

`func (o *GetZoneDeclarations200Response1DataInner) GetSecondaryDnsServerAddressOk() (*string, bool)`

GetSecondaryDnsServerAddressOk returns a tuple with the SecondaryDnsServerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryDnsServerAddress

`func (o *GetZoneDeclarations200Response1DataInner) SetSecondaryDnsServerAddress(v string)`

SetSecondaryDnsServerAddress sets SecondaryDnsServerAddress field to given value.

### HasSecondaryDnsServerAddress

`func (o *GetZoneDeclarations200Response1DataInner) HasSecondaryDnsServerAddress() bool`

HasSecondaryDnsServerAddress returns a boolean if a field has been set.

### GetDynamicUpdateSigningKeyType

`func (o *GetZoneDeclarations200Response1DataInner) GetDynamicUpdateSigningKeyType() string`

GetDynamicUpdateSigningKeyType returns the DynamicUpdateSigningKeyType field if non-nil, zero value otherwise.

### GetDynamicUpdateSigningKeyTypeOk

`func (o *GetZoneDeclarations200Response1DataInner) GetDynamicUpdateSigningKeyTypeOk() (*string, bool)`

GetDynamicUpdateSigningKeyTypeOk returns a tuple with the DynamicUpdateSigningKeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicUpdateSigningKeyType

`func (o *GetZoneDeclarations200Response1DataInner) SetDynamicUpdateSigningKeyType(v string)`

SetDynamicUpdateSigningKeyType sets DynamicUpdateSigningKeyType field to given value.

### HasDynamicUpdateSigningKeyType

`func (o *GetZoneDeclarations200Response1DataInner) HasDynamicUpdateSigningKeyType() bool`

HasDynamicUpdateSigningKeyType returns a boolean if a field has been set.

### GetDynamicUpdateSigningKey

`func (o *GetZoneDeclarations200Response1DataInner) GetDynamicUpdateSigningKey() InlinedSigningKey`

GetDynamicUpdateSigningKey returns the DynamicUpdateSigningKey field if non-nil, zero value otherwise.

### GetDynamicUpdateSigningKeyOk

`func (o *GetZoneDeclarations200Response1DataInner) GetDynamicUpdateSigningKeyOk() (*InlinedSigningKey, bool)`

GetDynamicUpdateSigningKeyOk returns a tuple with the DynamicUpdateSigningKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicUpdateSigningKey

`func (o *GetZoneDeclarations200Response1DataInner) SetDynamicUpdateSigningKey(v InlinedSigningKey)`

SetDynamicUpdateSigningKey sets DynamicUpdateSigningKey field to given value.

### HasDynamicUpdateSigningKey

`func (o *GetZoneDeclarations200Response1DataInner) HasDynamicUpdateSigningKey() bool`

HasDynamicUpdateSigningKey returns a boolean if a field has been set.

### GetDynamicUpdateSigningKeyRealm

`func (o *GetZoneDeclarations200Response1DataInner) GetDynamicUpdateSigningKeyRealm() InlinedKerberosRealm`

GetDynamicUpdateSigningKeyRealm returns the DynamicUpdateSigningKeyRealm field if non-nil, zero value otherwise.

### GetDynamicUpdateSigningKeyRealmOk

`func (o *GetZoneDeclarations200Response1DataInner) GetDynamicUpdateSigningKeyRealmOk() (*InlinedKerberosRealm, bool)`

GetDynamicUpdateSigningKeyRealmOk returns a tuple with the DynamicUpdateSigningKeyRealm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicUpdateSigningKeyRealm

`func (o *GetZoneDeclarations200Response1DataInner) SetDynamicUpdateSigningKeyRealm(v InlinedKerberosRealm)`

SetDynamicUpdateSigningKeyRealm sets DynamicUpdateSigningKeyRealm field to given value.

### HasDynamicUpdateSigningKeyRealm

`func (o *GetZoneDeclarations200Response1DataInner) HasDynamicUpdateSigningKeyRealm() bool`

HasDynamicUpdateSigningKeyRealm returns a boolean if a field has been set.

### GetPrefix

`func (o *GetZoneDeclarations200Response1DataInner) GetPrefix() DHCPReverseZoneAllOfPrefix`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *GetZoneDeclarations200Response1DataInner) GetPrefixOk() (*DHCPReverseZoneAllOfPrefix, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *GetZoneDeclarations200Response1DataInner) SetPrefix(v DHCPReverseZoneAllOfPrefix)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *GetZoneDeclarations200Response1DataInner) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


