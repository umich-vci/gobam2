# PatchServerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** | The current state of the server, indicating whether the server is enabled or disabled. | [optional] 

## Methods

### NewPatchServerRequest

`func NewPatchServerRequest() *PatchServerRequest`

NewPatchServerRequest instantiates a new PatchServerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchServerRequestWithDefaults

`func NewPatchServerRequestWithDefaults() *PatchServerRequest`

NewPatchServerRequestWithDefaults instantiates a new PatchServerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *PatchServerRequest) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PatchServerRequest) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PatchServerRequest) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *PatchServerRequest) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


