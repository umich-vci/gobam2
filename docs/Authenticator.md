# Authenticator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The type of authenticator. | [optional] 
**Name** | Pointer to **string** | The name for the authenticator. | [optional] 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**SecondaryAuthenticator** | Pointer to [**InlinedAuthenticator**](InlinedAuthenticator.md) |  | [optional] 

## Methods

### NewAuthenticator

`func NewAuthenticator() *Authenticator`

NewAuthenticator instantiates a new Authenticator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticatorWithDefaults

`func NewAuthenticatorWithDefaults() *Authenticator`

NewAuthenticatorWithDefaults instantiates a new Authenticator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Authenticator) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Authenticator) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Authenticator) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Authenticator) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *Authenticator) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Authenticator) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Authenticator) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Authenticator) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *Authenticator) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Authenticator) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Authenticator) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Authenticator) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *Authenticator) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *Authenticator) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *Authenticator) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *Authenticator) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetSecondaryAuthenticator

`func (o *Authenticator) GetSecondaryAuthenticator() InlinedAuthenticator`

GetSecondaryAuthenticator returns the SecondaryAuthenticator field if non-nil, zero value otherwise.

### GetSecondaryAuthenticatorOk

`func (o *Authenticator) GetSecondaryAuthenticatorOk() (*InlinedAuthenticator, bool)`

GetSecondaryAuthenticatorOk returns a tuple with the SecondaryAuthenticator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryAuthenticator

`func (o *Authenticator) SetSecondaryAuthenticator(v InlinedAuthenticator)`

SetSecondaryAuthenticator sets SecondaryAuthenticator field to given value.

### HasSecondaryAuthenticator

`func (o *Authenticator) HasSecondaryAuthenticator() bool`

HasSecondaryAuthenticator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


