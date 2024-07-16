# AddressManagerAuthenticator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of authenticator. | [optional] 
**Hostname** | Pointer to **string** | The fully qualified domain name or IP address for the authenticator. | [optional] 

## Methods

### NewAddressManagerAuthenticator

`func NewAddressManagerAuthenticator() *AddressManagerAuthenticator`

NewAddressManagerAuthenticator instantiates a new AddressManagerAuthenticator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressManagerAuthenticatorWithDefaults

`func NewAddressManagerAuthenticatorWithDefaults() *AddressManagerAuthenticator`

NewAddressManagerAuthenticatorWithDefaults instantiates a new AddressManagerAuthenticator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AddressManagerAuthenticator) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddressManagerAuthenticator) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddressManagerAuthenticator) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AddressManagerAuthenticator) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHostname

`func (o *AddressManagerAuthenticator) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *AddressManagerAuthenticator) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *AddressManagerAuthenticator) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *AddressManagerAuthenticator) HasHostname() bool`

HasHostname returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


