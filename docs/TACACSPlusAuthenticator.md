# TACACSPlusAuthenticator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of authenticator. | [optional] 
**Hostname** | Pointer to **string** | The fully qualified domain name or IP address for the authenticator. | [optional] 
**Timeout** | Pointer to **string** | The value that overrides the timeout setting used for authentication requests sent to the TACACS+ server. | [optional] 
**Port** | Pointer to **int32** | The TCP port number used for communication between the client and server. | [optional] 
**SharedSecret** | Pointer to **string** | The shared secret used to encrypt and decrypt packets between the client and the server. | [optional] 
**AuthenticationProtocol** | Pointer to **string** | The authentication protocol. | [optional] 
**GroupAttribute** | Pointer to **string** | The special attribute used for the custom service in the TACACS+ server. This attribute is used to get the group name defined in the TACACS+ server. | [optional] 
**Attributes** | Pointer to [**[]NameValuePairBean**](NameValuePairBean.md) |  | [optional] 

## Methods

### NewTACACSPlusAuthenticator

`func NewTACACSPlusAuthenticator() *TACACSPlusAuthenticator`

NewTACACSPlusAuthenticator instantiates a new TACACSPlusAuthenticator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTACACSPlusAuthenticatorWithDefaults

`func NewTACACSPlusAuthenticatorWithDefaults() *TACACSPlusAuthenticator`

NewTACACSPlusAuthenticatorWithDefaults instantiates a new TACACSPlusAuthenticator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TACACSPlusAuthenticator) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TACACSPlusAuthenticator) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TACACSPlusAuthenticator) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TACACSPlusAuthenticator) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHostname

`func (o *TACACSPlusAuthenticator) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *TACACSPlusAuthenticator) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *TACACSPlusAuthenticator) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *TACACSPlusAuthenticator) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetTimeout

`func (o *TACACSPlusAuthenticator) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *TACACSPlusAuthenticator) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *TACACSPlusAuthenticator) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *TACACSPlusAuthenticator) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetPort

`func (o *TACACSPlusAuthenticator) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *TACACSPlusAuthenticator) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *TACACSPlusAuthenticator) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *TACACSPlusAuthenticator) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetSharedSecret

`func (o *TACACSPlusAuthenticator) GetSharedSecret() string`

GetSharedSecret returns the SharedSecret field if non-nil, zero value otherwise.

### GetSharedSecretOk

`func (o *TACACSPlusAuthenticator) GetSharedSecretOk() (*string, bool)`

GetSharedSecretOk returns a tuple with the SharedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSecret

`func (o *TACACSPlusAuthenticator) SetSharedSecret(v string)`

SetSharedSecret sets SharedSecret field to given value.

### HasSharedSecret

`func (o *TACACSPlusAuthenticator) HasSharedSecret() bool`

HasSharedSecret returns a boolean if a field has been set.

### GetAuthenticationProtocol

`func (o *TACACSPlusAuthenticator) GetAuthenticationProtocol() string`

GetAuthenticationProtocol returns the AuthenticationProtocol field if non-nil, zero value otherwise.

### GetAuthenticationProtocolOk

`func (o *TACACSPlusAuthenticator) GetAuthenticationProtocolOk() (*string, bool)`

GetAuthenticationProtocolOk returns a tuple with the AuthenticationProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationProtocol

`func (o *TACACSPlusAuthenticator) SetAuthenticationProtocol(v string)`

SetAuthenticationProtocol sets AuthenticationProtocol field to given value.

### HasAuthenticationProtocol

`func (o *TACACSPlusAuthenticator) HasAuthenticationProtocol() bool`

HasAuthenticationProtocol returns a boolean if a field has been set.

### GetGroupAttribute

`func (o *TACACSPlusAuthenticator) GetGroupAttribute() string`

GetGroupAttribute returns the GroupAttribute field if non-nil, zero value otherwise.

### GetGroupAttributeOk

`func (o *TACACSPlusAuthenticator) GetGroupAttributeOk() (*string, bool)`

GetGroupAttributeOk returns a tuple with the GroupAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupAttribute

`func (o *TACACSPlusAuthenticator) SetGroupAttribute(v string)`

SetGroupAttribute sets GroupAttribute field to given value.

### HasGroupAttribute

`func (o *TACACSPlusAuthenticator) HasGroupAttribute() bool`

HasGroupAttribute returns a boolean if a field has been set.

### GetAttributes

`func (o *TACACSPlusAuthenticator) GetAttributes() []NameValuePairBean`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TACACSPlusAuthenticator) GetAttributesOk() (*[]NameValuePairBean, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TACACSPlusAuthenticator) SetAttributes(v []NameValuePairBean)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *TACACSPlusAuthenticator) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


