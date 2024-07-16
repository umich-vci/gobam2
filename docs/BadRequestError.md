# BadRequestError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **int32** |  | [optional] [readonly] 
**Reason** | Pointer to **string** |  | [optional] [readonly] 
**Code** | Pointer to **string** |  | [optional] [readonly] 
**Message** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewBadRequestError

`func NewBadRequestError() *BadRequestError`

NewBadRequestError instantiates a new BadRequestError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBadRequestErrorWithDefaults

`func NewBadRequestErrorWithDefaults() *BadRequestError`

NewBadRequestErrorWithDefaults instantiates a new BadRequestError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *BadRequestError) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BadRequestError) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BadRequestError) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BadRequestError) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetReason

`func (o *BadRequestError) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *BadRequestError) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *BadRequestError) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *BadRequestError) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetCode

`func (o *BadRequestError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *BadRequestError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *BadRequestError) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *BadRequestError) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *BadRequestError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *BadRequestError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *BadRequestError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *BadRequestError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


