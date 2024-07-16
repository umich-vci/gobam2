# ConflictError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **int32** |  | [optional] [readonly] 
**Reason** | Pointer to **string** |  | [optional] [readonly] 
**Code** | Pointer to **string** |  | [optional] [readonly] 
**Message** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewConflictError

`func NewConflictError() *ConflictError`

NewConflictError instantiates a new ConflictError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConflictErrorWithDefaults

`func NewConflictErrorWithDefaults() *ConflictError`

NewConflictErrorWithDefaults instantiates a new ConflictError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ConflictError) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ConflictError) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ConflictError) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ConflictError) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetReason

`func (o *ConflictError) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ConflictError) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ConflictError) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ConflictError) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetCode

`func (o *ConflictError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ConflictError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ConflictError) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ConflictError) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *ConflictError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ConflictError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ConflictError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ConflictError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


