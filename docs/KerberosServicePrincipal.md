# KerberosServicePrincipal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | Pointer to **string** | The name for the Kerberos service principal defined in the User Logon name field in Windows configuration section. | [optional] 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**KeyVersionNumber** | Pointer to **int32** | The msDS-KeyVersionNumber attribute value displayed in ADSI Edit on the Windows DC for the principal’s Kerberos key. | [optional] 
**Password** | Pointer to **string** | The principal’s Kerberos password. This is the AD user account password created on Windows DC. | [optional] 

## Methods

### NewKerberosServicePrincipal

`func NewKerberosServicePrincipal() *KerberosServicePrincipal`

NewKerberosServicePrincipal instantiates a new KerberosServicePrincipal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKerberosServicePrincipalWithDefaults

`func NewKerberosServicePrincipalWithDefaults() *KerberosServicePrincipal`

NewKerberosServicePrincipalWithDefaults instantiates a new KerberosServicePrincipal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KerberosServicePrincipal) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KerberosServicePrincipal) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KerberosServicePrincipal) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *KerberosServicePrincipal) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *KerberosServicePrincipal) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KerberosServicePrincipal) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KerberosServicePrincipal) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *KerberosServicePrincipal) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *KerberosServicePrincipal) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KerberosServicePrincipal) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KerberosServicePrincipal) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KerberosServicePrincipal) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *KerberosServicePrincipal) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *KerberosServicePrincipal) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *KerberosServicePrincipal) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *KerberosServicePrincipal) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetKeyVersionNumber

`func (o *KerberosServicePrincipal) GetKeyVersionNumber() int32`

GetKeyVersionNumber returns the KeyVersionNumber field if non-nil, zero value otherwise.

### GetKeyVersionNumberOk

`func (o *KerberosServicePrincipal) GetKeyVersionNumberOk() (*int32, bool)`

GetKeyVersionNumberOk returns a tuple with the KeyVersionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyVersionNumber

`func (o *KerberosServicePrincipal) SetKeyVersionNumber(v int32)`

SetKeyVersionNumber sets KeyVersionNumber field to given value.

### HasKeyVersionNumber

`func (o *KerberosServicePrincipal) HasKeyVersionNumber() bool`

HasKeyVersionNumber returns a boolean if a field has been set.

### GetPassword

`func (o *KerberosServicePrincipal) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *KerberosServicePrincipal) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *KerberosServicePrincipal) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *KerberosServicePrincipal) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


