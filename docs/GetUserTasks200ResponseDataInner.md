# GetUserTasks200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | Pointer to **string** | The name or description for the task. | [optional] 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**Priority** | Pointer to **string** | The priority level of the task. | [optional] 
**State** | Pointer to **string** | The current status of the task. | [optional] 
**PercentComplete** | Pointer to **int32** | The current progress percentage of the task. | [optional] 
**StartDateTime** | Pointer to **time.Time** | The start date for the task. | [optional] 
**DueDateTime** | Pointer to **time.Time** | The due date for the task. | [optional] 
**Comment** | Pointer to **string** | Custom comments for the task. | [optional] 
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 

## Methods

### NewGetUserTasks200ResponseDataInner

`func NewGetUserTasks200ResponseDataInner() *GetUserTasks200ResponseDataInner`

NewGetUserTasks200ResponseDataInner instantiates a new GetUserTasks200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserTasks200ResponseDataInnerWithDefaults

`func NewGetUserTasks200ResponseDataInnerWithDefaults() *GetUserTasks200ResponseDataInner`

NewGetUserTasks200ResponseDataInnerWithDefaults instantiates a new GetUserTasks200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetUserTasks200ResponseDataInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetUserTasks200ResponseDataInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetUserTasks200ResponseDataInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetUserTasks200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GetUserTasks200ResponseDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetUserTasks200ResponseDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetUserTasks200ResponseDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetUserTasks200ResponseDataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *GetUserTasks200ResponseDataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetUserTasks200ResponseDataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetUserTasks200ResponseDataInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetUserTasks200ResponseDataInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *GetUserTasks200ResponseDataInner) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *GetUserTasks200ResponseDataInner) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *GetUserTasks200ResponseDataInner) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *GetUserTasks200ResponseDataInner) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetPriority

`func (o *GetUserTasks200ResponseDataInner) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *GetUserTasks200ResponseDataInner) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *GetUserTasks200ResponseDataInner) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *GetUserTasks200ResponseDataInner) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetState

`func (o *GetUserTasks200ResponseDataInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetUserTasks200ResponseDataInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetUserTasks200ResponseDataInner) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GetUserTasks200ResponseDataInner) HasState() bool`

HasState returns a boolean if a field has been set.

### GetPercentComplete

`func (o *GetUserTasks200ResponseDataInner) GetPercentComplete() int32`

GetPercentComplete returns the PercentComplete field if non-nil, zero value otherwise.

### GetPercentCompleteOk

`func (o *GetUserTasks200ResponseDataInner) GetPercentCompleteOk() (*int32, bool)`

GetPercentCompleteOk returns a tuple with the PercentComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentComplete

`func (o *GetUserTasks200ResponseDataInner) SetPercentComplete(v int32)`

SetPercentComplete sets PercentComplete field to given value.

### HasPercentComplete

`func (o *GetUserTasks200ResponseDataInner) HasPercentComplete() bool`

HasPercentComplete returns a boolean if a field has been set.

### GetStartDateTime

`func (o *GetUserTasks200ResponseDataInner) GetStartDateTime() time.Time`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *GetUserTasks200ResponseDataInner) GetStartDateTimeOk() (*time.Time, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *GetUserTasks200ResponseDataInner) SetStartDateTime(v time.Time)`

SetStartDateTime sets StartDateTime field to given value.

### HasStartDateTime

`func (o *GetUserTasks200ResponseDataInner) HasStartDateTime() bool`

HasStartDateTime returns a boolean if a field has been set.

### GetDueDateTime

`func (o *GetUserTasks200ResponseDataInner) GetDueDateTime() time.Time`

GetDueDateTime returns the DueDateTime field if non-nil, zero value otherwise.

### GetDueDateTimeOk

`func (o *GetUserTasks200ResponseDataInner) GetDueDateTimeOk() (*time.Time, bool)`

GetDueDateTimeOk returns a tuple with the DueDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDateTime

`func (o *GetUserTasks200ResponseDataInner) SetDueDateTime(v time.Time)`

SetDueDateTime sets DueDateTime field to given value.

### HasDueDateTime

`func (o *GetUserTasks200ResponseDataInner) HasDueDateTime() bool`

HasDueDateTime returns a boolean if a field has been set.

### GetComment

`func (o *GetUserTasks200ResponseDataInner) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetUserTasks200ResponseDataInner) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetUserTasks200ResponseDataInner) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetUserTasks200ResponseDataInner) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetLinks

`func (o *GetUserTasks200ResponseDataInner) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetUserTasks200ResponseDataInner) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetUserTasks200ResponseDataInner) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetUserTasks200ResponseDataInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


