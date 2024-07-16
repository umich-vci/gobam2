# InlinedKerberosServicePrincipal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | Pointer to **string** | The name for the Kerberos service principal defined in the User Logon name field in Windows configuration section. | [optional] 

## Methods

### NewInlinedKerberosServicePrincipal

`func NewInlinedKerberosServicePrincipal() *InlinedKerberosServicePrincipal`

NewInlinedKerberosServicePrincipal instantiates a new InlinedKerberosServicePrincipal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlinedKerberosServicePrincipalWithDefaults

`func NewInlinedKerberosServicePrincipalWithDefaults() *InlinedKerberosServicePrincipal`

NewInlinedKerberosServicePrincipalWithDefaults instantiates a new InlinedKerberosServicePrincipal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlinedKerberosServicePrincipal) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlinedKerberosServicePrincipal) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlinedKerberosServicePrincipal) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *InlinedKerberosServicePrincipal) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *InlinedKerberosServicePrincipal) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlinedKerberosServicePrincipal) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlinedKerberosServicePrincipal) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InlinedKerberosServicePrincipal) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *InlinedKerberosServicePrincipal) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlinedKerberosServicePrincipal) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlinedKerberosServicePrincipal) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlinedKerberosServicePrincipal) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


