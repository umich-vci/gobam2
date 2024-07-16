# GetZoneDeclarations200ResponseDataInner

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
**Zone** | Pointer to [**DHCPForwardZoneAllOfZone**](DHCPForwardZoneAllOfZone.md) |  | [optional] 
**Prefix** | Pointer to [**DHCPReverseZoneAllOfPrefix**](DHCPReverseZoneAllOfPrefix.md) |  | [optional] 
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 

## Methods

### NewGetZoneDeclarations200ResponseDataInner

`func NewGetZoneDeclarations200ResponseDataInner() *GetZoneDeclarations200ResponseDataInner`

NewGetZoneDeclarations200ResponseDataInner instantiates a new GetZoneDeclarations200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetZoneDeclarations200ResponseDataInnerWithDefaults

`func NewGetZoneDeclarations200ResponseDataInnerWithDefaults() *GetZoneDeclarations200ResponseDataInner`

NewGetZoneDeclarations200ResponseDataInnerWithDefaults instantiates a new GetZoneDeclarations200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetZoneDeclarations200ResponseDataInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetZoneDeclarations200ResponseDataInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetZoneDeclarations200ResponseDataInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetZoneDeclarations200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GetZoneDeclarations200ResponseDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetZoneDeclarations200ResponseDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetZoneDeclarations200ResponseDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetZoneDeclarations200ResponseDataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *GetZoneDeclarations200ResponseDataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetZoneDeclarations200ResponseDataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetZoneDeclarations200ResponseDataInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetZoneDeclarations200ResponseDataInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *GetZoneDeclarations200ResponseDataInner) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *GetZoneDeclarations200ResponseDataInner) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *GetZoneDeclarations200ResponseDataInner) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *GetZoneDeclarations200ResponseDataInner) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetConfiguration

`func (o *GetZoneDeclarations200ResponseDataInner) GetConfiguration() InlinedConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *GetZoneDeclarations200ResponseDataInner) GetConfigurationOk() (*InlinedConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *GetZoneDeclarations200ResponseDataInner) SetConfiguration(v InlinedConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *GetZoneDeclarations200ResponseDataInner) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetPrimaryDnsServerAddress

`func (o *GetZoneDeclarations200ResponseDataInner) GetPrimaryDnsServerAddress() string`

GetPrimaryDnsServerAddress returns the PrimaryDnsServerAddress field if non-nil, zero value otherwise.

### GetPrimaryDnsServerAddressOk

`func (o *GetZoneDeclarations200ResponseDataInner) GetPrimaryDnsServerAddressOk() (*string, bool)`

GetPrimaryDnsServerAddressOk returns a tuple with the PrimaryDnsServerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryDnsServerAddress

`func (o *GetZoneDeclarations200ResponseDataInner) SetPrimaryDnsServerAddress(v string)`

SetPrimaryDnsServerAddress sets PrimaryDnsServerAddress field to given value.

### HasPrimaryDnsServerAddress

`func (o *GetZoneDeclarations200ResponseDataInner) HasPrimaryDnsServerAddress() bool`

HasPrimaryDnsServerAddress returns a boolean if a field has been set.

### GetSecondaryDnsServerAddress

`func (o *GetZoneDeclarations200ResponseDataInner) GetSecondaryDnsServerAddress() string`

GetSecondaryDnsServerAddress returns the SecondaryDnsServerAddress field if non-nil, zero value otherwise.

### GetSecondaryDnsServerAddressOk

`func (o *GetZoneDeclarations200ResponseDataInner) GetSecondaryDnsServerAddressOk() (*string, bool)`

GetSecondaryDnsServerAddressOk returns a tuple with the SecondaryDnsServerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryDnsServerAddress

`func (o *GetZoneDeclarations200ResponseDataInner) SetSecondaryDnsServerAddress(v string)`

SetSecondaryDnsServerAddress sets SecondaryDnsServerAddress field to given value.

### HasSecondaryDnsServerAddress

`func (o *GetZoneDeclarations200ResponseDataInner) HasSecondaryDnsServerAddress() bool`

HasSecondaryDnsServerAddress returns a boolean if a field has been set.

### GetDynamicUpdateSigningKeyType

`func (o *GetZoneDeclarations200ResponseDataInner) GetDynamicUpdateSigningKeyType() string`

GetDynamicUpdateSigningKeyType returns the DynamicUpdateSigningKeyType field if non-nil, zero value otherwise.

### GetDynamicUpdateSigningKeyTypeOk

`func (o *GetZoneDeclarations200ResponseDataInner) GetDynamicUpdateSigningKeyTypeOk() (*string, bool)`

GetDynamicUpdateSigningKeyTypeOk returns a tuple with the DynamicUpdateSigningKeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicUpdateSigningKeyType

`func (o *GetZoneDeclarations200ResponseDataInner) SetDynamicUpdateSigningKeyType(v string)`

SetDynamicUpdateSigningKeyType sets DynamicUpdateSigningKeyType field to given value.

### HasDynamicUpdateSigningKeyType

`func (o *GetZoneDeclarations200ResponseDataInner) HasDynamicUpdateSigningKeyType() bool`

HasDynamicUpdateSigningKeyType returns a boolean if a field has been set.

### GetDynamicUpdateSigningKey

`func (o *GetZoneDeclarations200ResponseDataInner) GetDynamicUpdateSigningKey() InlinedSigningKey`

GetDynamicUpdateSigningKey returns the DynamicUpdateSigningKey field if non-nil, zero value otherwise.

### GetDynamicUpdateSigningKeyOk

`func (o *GetZoneDeclarations200ResponseDataInner) GetDynamicUpdateSigningKeyOk() (*InlinedSigningKey, bool)`

GetDynamicUpdateSigningKeyOk returns a tuple with the DynamicUpdateSigningKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicUpdateSigningKey

`func (o *GetZoneDeclarations200ResponseDataInner) SetDynamicUpdateSigningKey(v InlinedSigningKey)`

SetDynamicUpdateSigningKey sets DynamicUpdateSigningKey field to given value.

### HasDynamicUpdateSigningKey

`func (o *GetZoneDeclarations200ResponseDataInner) HasDynamicUpdateSigningKey() bool`

HasDynamicUpdateSigningKey returns a boolean if a field has been set.

### GetDynamicUpdateSigningKeyRealm

`func (o *GetZoneDeclarations200ResponseDataInner) GetDynamicUpdateSigningKeyRealm() InlinedKerberosRealm`

GetDynamicUpdateSigningKeyRealm returns the DynamicUpdateSigningKeyRealm field if non-nil, zero value otherwise.

### GetDynamicUpdateSigningKeyRealmOk

`func (o *GetZoneDeclarations200ResponseDataInner) GetDynamicUpdateSigningKeyRealmOk() (*InlinedKerberosRealm, bool)`

GetDynamicUpdateSigningKeyRealmOk returns a tuple with the DynamicUpdateSigningKeyRealm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicUpdateSigningKeyRealm

`func (o *GetZoneDeclarations200ResponseDataInner) SetDynamicUpdateSigningKeyRealm(v InlinedKerberosRealm)`

SetDynamicUpdateSigningKeyRealm sets DynamicUpdateSigningKeyRealm field to given value.

### HasDynamicUpdateSigningKeyRealm

`func (o *GetZoneDeclarations200ResponseDataInner) HasDynamicUpdateSigningKeyRealm() bool`

HasDynamicUpdateSigningKeyRealm returns a boolean if a field has been set.

### GetZone

`func (o *GetZoneDeclarations200ResponseDataInner) GetZone() DHCPForwardZoneAllOfZone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *GetZoneDeclarations200ResponseDataInner) GetZoneOk() (*DHCPForwardZoneAllOfZone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *GetZoneDeclarations200ResponseDataInner) SetZone(v DHCPForwardZoneAllOfZone)`

SetZone sets Zone field to given value.

### HasZone

`func (o *GetZoneDeclarations200ResponseDataInner) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetPrefix

`func (o *GetZoneDeclarations200ResponseDataInner) GetPrefix() DHCPReverseZoneAllOfPrefix`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *GetZoneDeclarations200ResponseDataInner) GetPrefixOk() (*DHCPReverseZoneAllOfPrefix, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *GetZoneDeclarations200ResponseDataInner) SetPrefix(v DHCPReverseZoneAllOfPrefix)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *GetZoneDeclarations200ResponseDataInner) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetLinks

`func (o *GetZoneDeclarations200ResponseDataInner) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetZoneDeclarations200ResponseDataInner) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetZoneDeclarations200ResponseDataInner) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetZoneDeclarations200ResponseDataInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


