# PatchServicePrincipalRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | Pointer to **string** | The principal’s Kerberos password. This is the AD user account password created on Windows DC. | [optional] 

## Methods

### NewPatchServicePrincipalRequest

`func NewPatchServicePrincipalRequest() *PatchServicePrincipalRequest`

NewPatchServicePrincipalRequest instantiates a new PatchServicePrincipalRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchServicePrincipalRequestWithDefaults

`func NewPatchServicePrincipalRequestWithDefaults() *PatchServicePrincipalRequest`

NewPatchServicePrincipalRequestWithDefaults instantiates a new PatchServicePrincipalRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *PatchServicePrincipalRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PatchServicePrincipalRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PatchServicePrincipalRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *PatchServicePrincipalRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


