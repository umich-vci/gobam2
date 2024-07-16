# Snmpv3Bean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecurityLevel** | Pointer to **string** | The SNMP security level. | [optional] 
**Usernames** | Pointer to **[]string** |  | [optional] 
**Context** | Pointer to **string** |  | [optional] 
**AuthenticationProtocol** | Pointer to **string** | The authentication protocol to use with SNMP version 3. | [optional] 
**AuthenticationPassphrase** | Pointer to **string** | The user authentication password of the SNMP user. | [optional] 
**PrivacyProtocol** | Pointer to **string** | The user authentication protocol used to encrypt the SNMP messages. | [optional] 
**PrivacyPassphrase** | Pointer to **string** | The privacy authentication password used to encrypt the SNMP data. | [optional] 

## Methods

### NewSnmpv3Bean

`func NewSnmpv3Bean() *Snmpv3Bean`

NewSnmpv3Bean instantiates a new Snmpv3Bean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnmpv3BeanWithDefaults

`func NewSnmpv3BeanWithDefaults() *Snmpv3Bean`

NewSnmpv3BeanWithDefaults instantiates a new Snmpv3Bean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecurityLevel

`func (o *Snmpv3Bean) GetSecurityLevel() string`

GetSecurityLevel returns the SecurityLevel field if non-nil, zero value otherwise.

### GetSecurityLevelOk

`func (o *Snmpv3Bean) GetSecurityLevelOk() (*string, bool)`

GetSecurityLevelOk returns a tuple with the SecurityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityLevel

`func (o *Snmpv3Bean) SetSecurityLevel(v string)`

SetSecurityLevel sets SecurityLevel field to given value.

### HasSecurityLevel

`func (o *Snmpv3Bean) HasSecurityLevel() bool`

HasSecurityLevel returns a boolean if a field has been set.

### GetUsernames

`func (o *Snmpv3Bean) GetUsernames() []string`

GetUsernames returns the Usernames field if non-nil, zero value otherwise.

### GetUsernamesOk

`func (o *Snmpv3Bean) GetUsernamesOk() (*[]string, bool)`

GetUsernamesOk returns a tuple with the Usernames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernames

`func (o *Snmpv3Bean) SetUsernames(v []string)`

SetUsernames sets Usernames field to given value.

### HasUsernames

`func (o *Snmpv3Bean) HasUsernames() bool`

HasUsernames returns a boolean if a field has been set.

### GetContext

`func (o *Snmpv3Bean) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *Snmpv3Bean) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *Snmpv3Bean) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *Snmpv3Bean) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetAuthenticationProtocol

`func (o *Snmpv3Bean) GetAuthenticationProtocol() string`

GetAuthenticationProtocol returns the AuthenticationProtocol field if non-nil, zero value otherwise.

### GetAuthenticationProtocolOk

`func (o *Snmpv3Bean) GetAuthenticationProtocolOk() (*string, bool)`

GetAuthenticationProtocolOk returns a tuple with the AuthenticationProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationProtocol

`func (o *Snmpv3Bean) SetAuthenticationProtocol(v string)`

SetAuthenticationProtocol sets AuthenticationProtocol field to given value.

### HasAuthenticationProtocol

`func (o *Snmpv3Bean) HasAuthenticationProtocol() bool`

HasAuthenticationProtocol returns a boolean if a field has been set.

### GetAuthenticationPassphrase

`func (o *Snmpv3Bean) GetAuthenticationPassphrase() string`

GetAuthenticationPassphrase returns the AuthenticationPassphrase field if non-nil, zero value otherwise.

### GetAuthenticationPassphraseOk

`func (o *Snmpv3Bean) GetAuthenticationPassphraseOk() (*string, bool)`

GetAuthenticationPassphraseOk returns a tuple with the AuthenticationPassphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationPassphrase

`func (o *Snmpv3Bean) SetAuthenticationPassphrase(v string)`

SetAuthenticationPassphrase sets AuthenticationPassphrase field to given value.

### HasAuthenticationPassphrase

`func (o *Snmpv3Bean) HasAuthenticationPassphrase() bool`

HasAuthenticationPassphrase returns a boolean if a field has been set.

### GetPrivacyProtocol

`func (o *Snmpv3Bean) GetPrivacyProtocol() string`

GetPrivacyProtocol returns the PrivacyProtocol field if non-nil, zero value otherwise.

### GetPrivacyProtocolOk

`func (o *Snmpv3Bean) GetPrivacyProtocolOk() (*string, bool)`

GetPrivacyProtocolOk returns a tuple with the PrivacyProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyProtocol

`func (o *Snmpv3Bean) SetPrivacyProtocol(v string)`

SetPrivacyProtocol sets PrivacyProtocol field to given value.

### HasPrivacyProtocol

`func (o *Snmpv3Bean) HasPrivacyProtocol() bool`

HasPrivacyProtocol returns a boolean if a field has been set.

### GetPrivacyPassphrase

`func (o *Snmpv3Bean) GetPrivacyPassphrase() string`

GetPrivacyPassphrase returns the PrivacyPassphrase field if non-nil, zero value otherwise.

### GetPrivacyPassphraseOk

`func (o *Snmpv3Bean) GetPrivacyPassphraseOk() (*string, bool)`

GetPrivacyPassphraseOk returns a tuple with the PrivacyPassphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyPassphrase

`func (o *Snmpv3Bean) SetPrivacyPassphrase(v string)`

SetPrivacyPassphrase sets PrivacyPassphrase field to given value.

### HasPrivacyPassphrase

`func (o *Snmpv3Bean) HasPrivacyPassphrase() bool`

HasPrivacyPassphrase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


