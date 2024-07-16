# RADIUSAuthenticator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of authenticator. | [optional] 
**Hostname** | Pointer to **string** | The fully qualified domain name or IP address for the authenticator. | [optional] 
**Timeout** | Pointer to **string** | The value that overrides the timeout setting used for authentication requests sent to the RADIUS server. | [optional] 
**Port** | Pointer to **int32** | The port number used for authenticating users against the RADIUS server. | [optional] 
**Retries** | Pointer to **int32** | The number of times Address Manager attempts to retransmit a failed authentication request sent to the RADIUS server. | [optional] 
**SharedSecret** | Pointer to **string** | The shared secret used to encrypt and decrypt packets between the client and the server. | [optional] 
**AuthenticationProtocol** | Pointer to **string** | The authentication protocol. | [optional] 

## Methods

### NewRADIUSAuthenticator

`func NewRADIUSAuthenticator() *RADIUSAuthenticator`

NewRADIUSAuthenticator instantiates a new RADIUSAuthenticator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRADIUSAuthenticatorWithDefaults

`func NewRADIUSAuthenticatorWithDefaults() *RADIUSAuthenticator`

NewRADIUSAuthenticatorWithDefaults instantiates a new RADIUSAuthenticator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RADIUSAuthenticator) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RADIUSAuthenticator) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RADIUSAuthenticator) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RADIUSAuthenticator) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHostname

`func (o *RADIUSAuthenticator) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *RADIUSAuthenticator) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *RADIUSAuthenticator) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *RADIUSAuthenticator) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetTimeout

`func (o *RADIUSAuthenticator) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *RADIUSAuthenticator) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *RADIUSAuthenticator) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *RADIUSAuthenticator) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetPort

`func (o *RADIUSAuthenticator) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *RADIUSAuthenticator) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *RADIUSAuthenticator) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *RADIUSAuthenticator) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetRetries

`func (o *RADIUSAuthenticator) GetRetries() int32`

GetRetries returns the Retries field if non-nil, zero value otherwise.

### GetRetriesOk

`func (o *RADIUSAuthenticator) GetRetriesOk() (*int32, bool)`

GetRetriesOk returns a tuple with the Retries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetries

`func (o *RADIUSAuthenticator) SetRetries(v int32)`

SetRetries sets Retries field to given value.

### HasRetries

`func (o *RADIUSAuthenticator) HasRetries() bool`

HasRetries returns a boolean if a field has been set.

### GetSharedSecret

`func (o *RADIUSAuthenticator) GetSharedSecret() string`

GetSharedSecret returns the SharedSecret field if non-nil, zero value otherwise.

### GetSharedSecretOk

`func (o *RADIUSAuthenticator) GetSharedSecretOk() (*string, bool)`

GetSharedSecretOk returns a tuple with the SharedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSecret

`func (o *RADIUSAuthenticator) SetSharedSecret(v string)`

SetSharedSecret sets SharedSecret field to given value.

### HasSharedSecret

`func (o *RADIUSAuthenticator) HasSharedSecret() bool`

HasSharedSecret returns a boolean if a field has been set.

### GetAuthenticationProtocol

`func (o *RADIUSAuthenticator) GetAuthenticationProtocol() string`

GetAuthenticationProtocol returns the AuthenticationProtocol field if non-nil, zero value otherwise.

### GetAuthenticationProtocolOk

`func (o *RADIUSAuthenticator) GetAuthenticationProtocolOk() (*string, bool)`

GetAuthenticationProtocolOk returns a tuple with the AuthenticationProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationProtocol

`func (o *RADIUSAuthenticator) SetAuthenticationProtocol(v string)`

SetAuthenticationProtocol sets AuthenticationProtocol field to given value.

### HasAuthenticationProtocol

`func (o *RADIUSAuthenticator) HasAuthenticationProtocol() bool`

HasAuthenticationProtocol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


