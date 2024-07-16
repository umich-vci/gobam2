# KerberosAuthenticator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of authenticator. | [optional] 
**Hostname** | Pointer to **string** | The fully qualified domain name or IP address for the authenticator. | [optional] 
**Realm** | Pointer to **string** | The administrative domain for the Kerberos server. | [optional] 

## Methods

### NewKerberosAuthenticator

`func NewKerberosAuthenticator() *KerberosAuthenticator`

NewKerberosAuthenticator instantiates a new KerberosAuthenticator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKerberosAuthenticatorWithDefaults

`func NewKerberosAuthenticatorWithDefaults() *KerberosAuthenticator`

NewKerberosAuthenticatorWithDefaults instantiates a new KerberosAuthenticator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *KerberosAuthenticator) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KerberosAuthenticator) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KerberosAuthenticator) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *KerberosAuthenticator) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHostname

`func (o *KerberosAuthenticator) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *KerberosAuthenticator) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *KerberosAuthenticator) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *KerberosAuthenticator) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetRealm

`func (o *KerberosAuthenticator) GetRealm() string`

GetRealm returns the Realm field if non-nil, zero value otherwise.

### GetRealmOk

`func (o *KerberosAuthenticator) GetRealmOk() (*string, bool)`

GetRealmOk returns a tuple with the Realm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealm

`func (o *KerberosAuthenticator) SetRealm(v string)`

SetRealm sets Realm field to given value.

### HasRealm

`func (o *KerberosAuthenticator) HasRealm() bool`

HasRealm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


