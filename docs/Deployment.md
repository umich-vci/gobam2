# Deployment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**State** | Pointer to **string** | The current state of the deployment. | [optional] 
**Status** | Pointer to **string** | The current status of the deployment. | [optional] [readonly] 
**Message** | Pointer to **string** | The expanded status message for the deployment. | [optional] [readonly] 
**PercentComplete** | Pointer to **int32** | A percentage value representing the current progress of the deployment. | [optional] [readonly] 
**StartDateTime** | Pointer to **time.Time** | The start time of the deployment. | [optional] [readonly] 
**CompletionDateTime** | Pointer to **time.Time** | The completion time for the deployment. | [optional] [readonly] 
**User** | Pointer to [**User**](User.md) |  | [optional] [readonly] 
**Method** | Pointer to **string** | The method that triggered the deployment. | [optional] [readonly] 

## Methods

### NewDeployment

`func NewDeployment() *Deployment`

NewDeployment instantiates a new Deployment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentWithDefaults

`func NewDeploymentWithDefaults() *Deployment`

NewDeploymentWithDefaults instantiates a new Deployment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Deployment) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Deployment) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Deployment) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Deployment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *Deployment) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Deployment) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Deployment) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Deployment) HasType() bool`

HasType returns a boolean if a field has been set.

### GetState

`func (o *Deployment) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Deployment) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Deployment) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Deployment) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStatus

`func (o *Deployment) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Deployment) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Deployment) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Deployment) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *Deployment) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Deployment) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Deployment) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *Deployment) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetPercentComplete

`func (o *Deployment) GetPercentComplete() int32`

GetPercentComplete returns the PercentComplete field if non-nil, zero value otherwise.

### GetPercentCompleteOk

`func (o *Deployment) GetPercentCompleteOk() (*int32, bool)`

GetPercentCompleteOk returns a tuple with the PercentComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentComplete

`func (o *Deployment) SetPercentComplete(v int32)`

SetPercentComplete sets PercentComplete field to given value.

### HasPercentComplete

`func (o *Deployment) HasPercentComplete() bool`

HasPercentComplete returns a boolean if a field has been set.

### GetStartDateTime

`func (o *Deployment) GetStartDateTime() time.Time`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *Deployment) GetStartDateTimeOk() (*time.Time, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *Deployment) SetStartDateTime(v time.Time)`

SetStartDateTime sets StartDateTime field to given value.

### HasStartDateTime

`func (o *Deployment) HasStartDateTime() bool`

HasStartDateTime returns a boolean if a field has been set.

### GetCompletionDateTime

`func (o *Deployment) GetCompletionDateTime() time.Time`

GetCompletionDateTime returns the CompletionDateTime field if non-nil, zero value otherwise.

### GetCompletionDateTimeOk

`func (o *Deployment) GetCompletionDateTimeOk() (*time.Time, bool)`

GetCompletionDateTimeOk returns a tuple with the CompletionDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionDateTime

`func (o *Deployment) SetCompletionDateTime(v time.Time)`

SetCompletionDateTime sets CompletionDateTime field to given value.

### HasCompletionDateTime

`func (o *Deployment) HasCompletionDateTime() bool`

HasCompletionDateTime returns a boolean if a field has been set.

### GetUser

`func (o *Deployment) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Deployment) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Deployment) SetUser(v User)`

SetUser sets User field to given value.

### HasUser

`func (o *Deployment) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetMethod

`func (o *Deployment) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *Deployment) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *Deployment) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *Deployment) HasMethod() bool`

HasMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


