# TaskPostRequestBody

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

### NewTaskPostRequestBody

`func NewTaskPostRequestBody(name string, priority string, state string, percentComplete int32, ) *TaskPostRequestBody`

NewTaskPostRequestBody instantiates a new TaskPostRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskPostRequestBodyWithDefaults

`func NewTaskPostRequestBodyWithDefaults() *TaskPostRequestBody`

NewTaskPostRequestBodyWithDefaults instantiates a new TaskPostRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TaskPostRequestBody) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaskPostRequestBody) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaskPostRequestBody) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *TaskPostRequestBody) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *TaskPostRequestBody) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TaskPostRequestBody) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TaskPostRequestBody) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TaskPostRequestBody) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *TaskPostRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TaskPostRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TaskPostRequestBody) SetName(v string)`

SetName sets Name field to given value.


### GetUserDefinedFields

`func (o *TaskPostRequestBody) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *TaskPostRequestBody) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *TaskPostRequestBody) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *TaskPostRequestBody) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetPriority

`func (o *TaskPostRequestBody) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *TaskPostRequestBody) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *TaskPostRequestBody) SetPriority(v string)`

SetPriority sets Priority field to given value.


### GetState

`func (o *TaskPostRequestBody) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TaskPostRequestBody) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TaskPostRequestBody) SetState(v string)`

SetState sets State field to given value.


### GetPercentComplete

`func (o *TaskPostRequestBody) GetPercentComplete() int32`

GetPercentComplete returns the PercentComplete field if non-nil, zero value otherwise.

### GetPercentCompleteOk

`func (o *TaskPostRequestBody) GetPercentCompleteOk() (*int32, bool)`

GetPercentCompleteOk returns a tuple with the PercentComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentComplete

`func (o *TaskPostRequestBody) SetPercentComplete(v int32)`

SetPercentComplete sets PercentComplete field to given value.


### GetStartDateTime

`func (o *TaskPostRequestBody) GetStartDateTime() time.Time`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *TaskPostRequestBody) GetStartDateTimeOk() (*time.Time, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *TaskPostRequestBody) SetStartDateTime(v time.Time)`

SetStartDateTime sets StartDateTime field to given value.

### HasStartDateTime

`func (o *TaskPostRequestBody) HasStartDateTime() bool`

HasStartDateTime returns a boolean if a field has been set.

### GetDueDateTime

`func (o *TaskPostRequestBody) GetDueDateTime() time.Time`

GetDueDateTime returns the DueDateTime field if non-nil, zero value otherwise.

### GetDueDateTimeOk

`func (o *TaskPostRequestBody) GetDueDateTimeOk() (*time.Time, bool)`

GetDueDateTimeOk returns a tuple with the DueDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDateTime

`func (o *TaskPostRequestBody) SetDueDateTime(v time.Time)`

SetDueDateTime sets DueDateTime field to given value.

### HasDueDateTime

`func (o *TaskPostRequestBody) HasDueDateTime() bool`

HasDueDateTime returns a boolean if a field has been set.

### GetComment

`func (o *TaskPostRequestBody) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *TaskPostRequestBody) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *TaskPostRequestBody) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *TaskPostRequestBody) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


