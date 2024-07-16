# PatchUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | Pointer to **string** | The password of the user. | [optional] 

## Methods

### NewPatchUserRequest

`func NewPatchUserRequest() *PatchUserRequest`

NewPatchUserRequest instantiates a new PatchUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchUserRequestWithDefaults

`func NewPatchUserRequestWithDefaults() *PatchUserRequest`

NewPatchUserRequestWithDefaults instantiates a new PatchUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *PatchUserRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PatchUserRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PatchUserRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *PatchUserRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


