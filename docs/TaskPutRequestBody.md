# TaskPutRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | **string** | The name or description for the task. | 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**Priority** | **string** | The priority level of the task. | 
**State** | **string** | The current status of the task. | 
**PercentComplete** | **int32** | The current progress percentage of the task. | 
**StartDateTime** | Pointer to **time.Time** | The start date for the task. | [optional] 
**DueDateTime** | Pointer to **time.Time** | The due date for the task. | [optional] 
**Comment** | Pointer to **string** | Custom comments for the task. | [optional] 

## Methods

### NewTaskPutRequestBody

`func NewTaskPutRequestBody(name string, priority string, state string, percentComplete int32, ) *TaskPutRequestBody`

NewTaskPutRequestBody instantiates a new TaskPutRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskPutRequestBodyWithDefaults

`func NewTaskPutRequestBodyWithDefaults() *TaskPutRequestBody`

NewTaskPutRequestBodyWithDefaults instantiates a new TaskPutRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TaskPutRequestBody) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaskPutRequestBody) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaskPutRequestBody) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *TaskPutRequestBody) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *TaskPutRequestBody) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TaskPutRequestBody) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TaskPutRequestBody) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TaskPutRequestBody) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *TaskPutRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TaskPutRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TaskPutRequestBody) SetName(v string)`

SetName sets Name field to given value.


### GetUserDefinedFields

`func (o *TaskPutRequestBody) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *TaskPutRequestBody) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *TaskPutRequestBody) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *TaskPutRequestBody) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetPriority

`func (o *TaskPutRequestBody) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *TaskPutRequestBody) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *TaskPutRequestBody) SetPriority(v string)`

SetPriority sets Priority field to given value.


### GetState

`func (o *TaskPutRequestBody) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TaskPutRequestBody) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TaskPutRequestBody) SetState(v string)`

SetState sets State field to given value.


### GetPercentComplete

`func (o *TaskPutRequestBody) GetPercentComplete() int32`

GetPercentComplete returns the PercentComplete field if non-nil, zero value otherwise.

### GetPercentCompleteOk

`func (o *TaskPutRequestBody) GetPercentCompleteOk() (*int32, bool)`

GetPercentCompleteOk returns a tuple with the PercentComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentComplete

`func (o *TaskPutRequestBody) SetPercentComplete(v int32)`

SetPercentComplete sets PercentComplete field to given value.


### GetStartDateTime

`func (o *TaskPutRequestBody) GetStartDateTime() time.Time`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *TaskPutRequestBody) GetStartDateTimeOk() (*time.Time, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *TaskPutRequestBody) SetStartDateTime(v time.Time)`

SetStartDateTime sets StartDateTime field to given value.

### HasStartDateTime

`func (o *TaskPutRequestBody) HasStartDateTime() bool`

HasStartDateTime returns a boolean if a field has been set.

### GetDueDateTime

`func (o *TaskPutRequestBody) GetDueDateTime() time.Time`

GetDueDateTime returns the DueDateTime field if non-nil, zero value otherwise.

### GetDueDateTimeOk

`func (o *TaskPutRequestBody) GetDueDateTimeOk() (*time.Time, bool)`

GetDueDateTimeOk returns a tuple with the DueDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDateTime

`func (o *TaskPutRequestBody) SetDueDateTime(v time.Time)`

SetDueDateTime sets DueDateTime field to given value.

### HasDueDateTime

`func (o *TaskPutRequestBody) HasDueDateTime() bool`

HasDueDateTime returns a boolean if a field has been set.

### GetComment

`func (o *TaskPutRequestBody) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *TaskPutRequestBody) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *TaskPutRequestBody) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *TaskPutRequestBody) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


