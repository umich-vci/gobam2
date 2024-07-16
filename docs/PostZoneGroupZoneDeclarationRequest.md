# PostZoneGroupZoneDeclarationRequest

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
**DynamicUpdateSigningKeyType** | Pointer to **string** | The type of key used to sign DDNS updates. | [optional] 
**DynamicUpdateSigningKey** | Pointer to [**InlinedSigningKey**](InlinedSigningKey.md) |  | [optional] 
**DynamicUpdateSigningKeyRealm** | Pointer to [**InlinedKerberosRealm**](InlinedKerberosRealm.md) |  | [optional] 
**Prefix** | [**DHCPReverseZoneAllOfPrefix**](DHCPReverseZoneAllOfPrefix.md) |  | 

## Methods

### NewPostZoneGroupZoneDeclarationRequest

`func NewPostZoneGroupZoneDeclarationRequest(type_ string, primaryDnsServerAddress string, prefix DHCPReverseZoneAllOfPrefix, ) *PostZoneGroupZoneDeclarationRequest`

NewPostZoneGroupZoneDeclarationRequest instantiates a new PostZoneGroupZoneDeclarationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostZoneGroupZoneDeclarationRequestWithDefaults

`func NewPostZoneGroupZoneDeclarationRequestWithDefaults() *PostZoneGroupZoneDeclarationRequest`

NewPostZoneGroupZoneDeclarationRequestWithDefaults instantiates a new PostZoneGroupZoneDeclarationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PostZoneGroupZoneDeclarationRequest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostZoneGroupZoneDeclarationRequest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostZoneGroupZoneDeclarationRequest) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *PostZoneGroupZoneDeclarationRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *PostZoneGroupZoneDeclarationRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PostZoneGroupZoneDeclarationRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PostZoneGroupZoneDeclarationRequest) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *PostZoneGroupZoneDeclarationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostZoneGroupZoneDeclarationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostZoneGroupZoneDeclarationRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PostZoneGroupZoneDeclarationRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *PostZoneGroupZoneDeclarationRequest) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *PostZoneGroupZoneDeclarationRequest) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *PostZoneGroupZoneDeclarationRequest) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *PostZoneGroupZoneDeclarationRequest) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetConfiguration

`func (o *PostZoneGroupZoneDeclarationRequest) GetConfiguration() InlinedConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *PostZoneGroupZoneDeclarationRequest) GetConfigurationOk() (*InlinedConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *PostZoneGroupZoneDeclarationRequest) SetConfiguration(v InlinedConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *PostZoneGroupZoneDeclarationRequest) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetPrimaryDnsServerAddress

`func (o *PostZoneGroupZoneDeclarationRequest) GetPrimaryDnsServerAddress() string`

GetPrimaryDnsServerAddress returns the PrimaryDnsServerAddress field if non-nil, zero value otherwise.

### GetPrimaryDnsServerAddressOk

`func (o *PostZoneGroupZoneDeclarationRequest) GetPrimaryDnsServerAddressOk() (*string, bool)`

GetPrimaryDnsServerAddressOk returns a tuple with the PrimaryDnsServerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryDnsServerAddress

`func (o *PostZoneGroupZoneDeclarationRequest) SetPrimaryDnsServerAddress(v string)`

SetPrimaryDnsServerAddress sets PrimaryDnsServerAddress field to given value.


### GetSecondaryDnsServerAddress

`func (o *PostZoneGroupZoneDeclarationRequest) GetSecondaryDnsServerAddress() string`

GetSecondaryDnsServerAddress returns the SecondaryDnsServerAddress field if non-nil, zero value otherwise.

### GetSecondaryDnsServerAddressOk

`func (o *PostZoneGroupZoneDeclarationRequest) GetSecondaryDnsServerAddressOk() (*string, bool)`

GetSecondaryDnsServerAddressOk returns a tuple with the SecondaryDnsServerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryDnsServerAddress

`func (o *PostZoneGroupZoneDeclarationRequest) SetSecondaryDnsServerAddress(v string)`

SetSecondaryDnsServerAddress sets SecondaryDnsServerAddress field to given value.

### HasSecondaryDnsServerAddress

`func (o *PostZoneGroupZoneDeclarationRequest) HasSecondaryDnsServerAddress() bool`

HasSecondaryDnsServerAddress returns a boolean if a field has been set.

### GetDynamicUpdateSigningKeyType

`func (o *PostZoneGroupZoneDeclarationRequest) GetDynamicUpdateSigningKeyType() string`

GetDynamicUpdateSigningKeyType returns the DynamicUpdateSigningKeyType field if non-nil, zero value otherwise.

### GetDynamicUpdateSigningKeyTypeOk

`func (o *PostZoneGroupZoneDeclarationRequest) GetDynamicUpdateSigningKeyTypeOk() (*string, bool)`

GetDynamicUpdateSigningKeyTypeOk returns a tuple with the DynamicUpdateSigningKeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicUpdateSigningKeyType

`func (o *PostZoneGroupZoneDeclarationRequest) SetDynamicUpdateSigningKeyType(v string)`

SetDynamicUpdateSigningKeyType sets DynamicUpdateSigningKeyType field to given value.

### HasDynamicUpdateSigningKeyType

`func (o *PostZoneGroupZoneDeclarationRequest) HasDynamicUpdateSigningKeyType() bool`

HasDynamicUpdateSigningKeyType returns a boolean if a field has been set.

### GetDynamicUpdateSigningKey

`func (o *PostZoneGroupZoneDeclarationRequest) GetDynamicUpdateSigningKey() InlinedSigningKey`

GetDynamicUpdateSigningKey returns the DynamicUpdateSigningKey field if non-nil, zero value otherwise.

### GetDynamicUpdateSigningKeyOk

`func (o *PostZoneGroupZoneDeclarationRequest) GetDynamicUpdateSigningKeyOk() (*InlinedSigningKey, bool)`

GetDynamicUpdateSigningKeyOk returns a tuple with the DynamicUpdateSigningKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicUpdateSigningKey

`func (o *PostZoneGroupZoneDeclarationRequest) SetDynamicUpdateSigningKey(v InlinedSigningKey)`

SetDynamicUpdateSigningKey sets DynamicUpdateSigningKey field to given value.

### HasDynamicUpdateSigningKey

`func (o *PostZoneGroupZoneDeclarationRequest) HasDynamicUpdateSigningKey() bool`

HasDynamicUpdateSigningKey returns a boolean if a field has been set.

### GetDynamicUpdateSigningKeyRealm

`func (o *PostZoneGroupZoneDeclarationRequest) GetDynamicUpdateSigningKeyRealm() InlinedKerberosRealm`

GetDynamicUpdateSigningKeyRealm returns the DynamicUpdateSigningKeyRealm field if non-nil, zero value otherwise.

### GetDynamicUpdateSigningKeyRealmOk

`func (o *PostZoneGroupZoneDeclarationRequest) GetDynamicUpdateSigningKeyRealmOk() (*InlinedKerberosRealm, bool)`

GetDynamicUpdateSigningKeyRealmOk returns a tuple with the DynamicUpdateSigningKeyRealm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicUpdateSigningKeyRealm

`func (o *PostZoneGroupZoneDeclarationRequest) SetDynamicUpdateSigningKeyRealm(v InlinedKerberosRealm)`

SetDynamicUpdateSigningKeyRealm sets DynamicUpdateSigningKeyRealm field to given value.

### HasDynamicUpdateSigningKeyRealm

`func (o *PostZoneGroupZoneDeclarationRequest) HasDynamicUpdateSigningKeyRealm() bool`

HasDynamicUpdateSigningKeyRealm returns a boolean if a field has been set.

### GetPrefix

`func (o *PostZoneGroupZoneDeclarationRequest) GetPrefix() DHCPReverseZoneAllOfPrefix`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *PostZoneGroupZoneDeclarationRequest) GetPrefixOk() (*DHCPReverseZoneAllOfPrefix, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *PostZoneGroupZoneDeclarationRequest) SetPrefix(v DHCPReverseZoneAllOfPrefix)`

SetPrefix sets Prefix field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


