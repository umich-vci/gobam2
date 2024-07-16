# DhcpZone

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

## Methods

### NewDhcpZone

`func NewDhcpZone() *DhcpZone`

NewDhcpZone instantiates a new DhcpZone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDhcpZoneWithDefaults

`func NewDhcpZoneWithDefaults() *DhcpZone`

NewDhcpZoneWithDefaults instantiates a new DhcpZone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DhcpZone) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DhcpZone) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DhcpZone) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DhcpZone) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *DhcpZone) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DhcpZone) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DhcpZone) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DhcpZone) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *DhcpZone) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DhcpZone) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DhcpZone) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DhcpZone) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *DhcpZone) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *DhcpZone) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *DhcpZone) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *DhcpZone) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetConfiguration

`func (o *DhcpZone) GetConfiguration() InlinedConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *DhcpZone) GetConfigurationOk() (*InlinedConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *DhcpZone) SetConfiguration(v InlinedConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *DhcpZone) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetPrimaryDnsServerAddress

`func (o *DhcpZone) GetPrimaryDnsServerAddress() string`

GetPrimaryDnsServerAddress returns the PrimaryDnsServerAddress field if non-nil, zero value otherwise.

### GetPrimaryDnsServerAddressOk

`func (o *DhcpZone) GetPrimaryDnsServerAddressOk() (*string, bool)`

GetPrimaryDnsServerAddressOk returns a tuple with the PrimaryDnsServerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryDnsServerAddress

`func (o *DhcpZone) SetPrimaryDnsServerAddress(v string)`

SetPrimaryDnsServerAddress sets PrimaryDnsServerAddress field to given value.

### HasPrimaryDnsServerAddress

`func (o *DhcpZone) HasPrimaryDnsServerAddress() bool`

HasPrimaryDnsServerAddress returns a boolean if a field has been set.

### GetSecondaryDnsServerAddress

`func (o *DhcpZone) GetSecondaryDnsServerAddress() string`

GetSecondaryDnsServerAddress returns the SecondaryDnsServerAddress field if non-nil, zero value otherwise.

### GetSecondaryDnsServerAddressOk

`func (o *DhcpZone) GetSecondaryDnsServerAddressOk() (*string, bool)`

GetSecondaryDnsServerAddressOk returns a tuple with the SecondaryDnsServerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryDnsServerAddress

`func (o *DhcpZone) SetSecondaryDnsServerAddress(v string)`

SetSecondaryDnsServerAddress sets SecondaryDnsServerAddress field to given value.

### HasSecondaryDnsServerAddress

`func (o *DhcpZone) HasSecondaryDnsServerAddress() bool`

HasSecondaryDnsServerAddress returns a boolean if a field has been set.

### GetDynamicUpdateSigningKeyType

`func (o *DhcpZone) GetDynamicUpdateSigningKeyType() string`

GetDynamicUpdateSigningKeyType returns the DynamicUpdateSigningKeyType field if non-nil, zero value otherwise.

### GetDynamicUpdateSigningKeyTypeOk

`func (o *DhcpZone) GetDynamicUpdateSigningKeyTypeOk() (*string, bool)`

GetDynamicUpdateSigningKeyTypeOk returns a tuple with the DynamicUpdateSigningKeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicUpdateSigningKeyType

`func (o *DhcpZone) SetDynamicUpdateSigningKeyType(v string)`

SetDynamicUpdateSigningKeyType sets DynamicUpdateSigningKeyType field to given value.

### HasDynamicUpdateSigningKeyType

`func (o *DhcpZone) HasDynamicUpdateSigningKeyType() bool`

HasDynamicUpdateSigningKeyType returns a boolean if a field has been set.

### GetDynamicUpdateSigningKey

`func (o *DhcpZone) GetDynamicUpdateSigningKey() InlinedSigningKey`

GetDynamicUpdateSigningKey returns the DynamicUpdateSigningKey field if non-nil, zero value otherwise.

### GetDynamicUpdateSigningKeyOk

`func (o *DhcpZone) GetDynamicUpdateSigningKeyOk() (*InlinedSigningKey, bool)`

GetDynamicUpdateSigningKeyOk returns a tuple with the DynamicUpdateSigningKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicUpdateSigningKey

`func (o *DhcpZone) SetDynamicUpdateSigningKey(v InlinedSigningKey)`

SetDynamicUpdateSigningKey sets DynamicUpdateSigningKey field to given value.

### HasDynamicUpdateSigningKey

`func (o *DhcpZone) HasDynamicUpdateSigningKey() bool`

HasDynamicUpdateSigningKey returns a boolean if a field has been set.

### GetDynamicUpdateSigningKeyRealm

`func (o *DhcpZone) GetDynamicUpdateSigningKeyRealm() InlinedKerberosRealm`

GetDynamicUpdateSigningKeyRealm returns the DynamicUpdateSigningKeyRealm field if non-nil, zero value otherwise.

### GetDynamicUpdateSigningKeyRealmOk

`func (o *DhcpZone) GetDynamicUpdateSigningKeyRealmOk() (*InlinedKerberosRealm, bool)`

GetDynamicUpdateSigningKeyRealmOk returns a tuple with the DynamicUpdateSigningKeyRealm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicUpdateSigningKeyRealm

`func (o *DhcpZone) SetDynamicUpdateSigningKeyRealm(v InlinedKerberosRealm)`

SetDynamicUpdateSigningKeyRealm sets DynamicUpdateSigningKeyRealm field to given value.

### HasDynamicUpdateSigningKeyRealm

`func (o *DhcpZone) HasDynamicUpdateSigningKeyRealm() bool`

HasDynamicUpdateSigningKeyRealm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


