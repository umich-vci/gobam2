# KerberosServicePrincipalPutRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | **string** | The name for the Kerberos service principal defined in the User Logon name field in Windows configuration section. | 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**KeyVersionNumber** | **int32** | The msDS-KeyVersionNumber attribute value displayed in ADSI Edit on the Windows DC for the principal’s Kerberos key. | 
**Password** | Pointer to **string** | The principal’s Kerberos password. This is the AD user account password created on Windows DC. | [optional] 

## Methods

### NewKerberosServicePrincipalPutRequestBody

`func NewKerberosServicePrincipalPutRequestBody(name string, keyVersionNumber int32, ) *KerberosServicePrincipalPutRequestBody`

NewKerberosServicePrincipalPutRequestBody instantiates a new KerberosServicePrincipalPutRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKerberosServicePrincipalPutRequestBodyWithDefaults

`func NewKerberosServicePrincipalPutRequestBodyWithDefaults() *KerberosServicePrincipalPutRequestBody`

NewKerberosServicePrincipalPutRequestBodyWithDefaults instantiates a new KerberosServicePrincipalPutRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KerberosServicePrincipalPutRequestBody) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KerberosServicePrincipalPutRequestBody) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KerberosServicePrincipalPutRequestBody) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *KerberosServicePrincipalPutRequestBody) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *KerberosServicePrincipalPutRequestBody) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KerberosServicePrincipalPutRequestBody) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KerberosServicePrincipalPutRequestBody) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *KerberosServicePrincipalPutRequestBody) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *KerberosServicePrincipalPutRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KerberosServicePrincipalPutRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KerberosServicePrincipalPutRequestBody) SetName(v string)`

SetName sets Name field to given value.


### GetUserDefinedFields

`func (o *KerberosServicePrincipalPutRequestBody) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *KerberosServicePrincipalPutRequestBody) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *KerberosServicePrincipalPutRequestBody) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *KerberosServicePrincipalPutRequestBody) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetKeyVersionNumber

`func (o *KerberosServicePrincipalPutRequestBody) GetKeyVersionNumber() int32`

GetKeyVersionNumber returns the KeyVersionNumber field if non-nil, zero value otherwise.

### GetKeyVersionNumberOk

`func (o *KerberosServicePrincipalPutRequestBody) GetKeyVersionNumberOk() (*int32, bool)`

GetKeyVersionNumberOk returns a tuple with the KeyVersionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyVersionNumber

`func (o *KerberosServicePrincipalPutRequestBody) SetKeyVersionNumber(v int32)`

SetKeyVersionNumber sets KeyVersionNumber field to given value.


### GetPassword

`func (o *KerberosServicePrincipalPutRequestBody) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *KerberosServicePrincipalPutRequestBody) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *KerberosServicePrincipalPutRequestBody) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *KerberosServicePrincipalPutRequestBody) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


