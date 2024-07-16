# PostCollectionDeployment201Response1

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
**Service** | Pointer to **string** | The deployment service type. | [optional] 

## Methods

### NewPostCollectionDeployment201Response1

`func NewPostCollectionDeployment201Response1() *PostCollectionDeployment201Response1`

NewPostCollectionDeployment201Response1 instantiates a new PostCollectionDeployment201Response1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostCollectionDeployment201Response1WithDefaults

`func NewPostCollectionDeployment201Response1WithDefaults() *PostCollectionDeployment201Response1`

NewPostCollectionDeployment201Response1WithDefaults instantiates a new PostCollectionDeployment201Response1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PostCollectionDeployment201Response1) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostCollectionDeployment201Response1) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostCollectionDeployment201Response1) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *PostCollectionDeployment201Response1) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *PostCollectionDeployment201Response1) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PostCollectionDeployment201Response1) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PostCollectionDeployment201Response1) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PostCollectionDeployment201Response1) HasType() bool`

HasType returns a boolean if a field has been set.

### GetState

`func (o *PostCollectionDeployment201Response1) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PostCollectionDeployment201Response1) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PostCollectionDeployment201Response1) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *PostCollectionDeployment201Response1) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStatus

`func (o *PostCollectionDeployment201Response1) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PostCollectionDeployment201Response1) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PostCollectionDeployment201Response1) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PostCollectionDeployment201Response1) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *PostCollectionDeployment201Response1) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PostCollectionDeployment201Response1) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PostCollectionDeployment201Response1) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *PostCollectionDeployment201Response1) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetPercentComplete

`func (o *PostCollectionDeployment201Response1) GetPercentComplete() int32`

GetPercentComplete returns the PercentComplete field if non-nil, zero value otherwise.

### GetPercentCompleteOk

`func (o *PostCollectionDeployment201Response1) GetPercentCompleteOk() (*int32, bool)`

GetPercentCompleteOk returns a tuple with the PercentComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentComplete

`func (o *PostCollectionDeployment201Response1) SetPercentComplete(v int32)`

SetPercentComplete sets PercentComplete field to given value.

### HasPercentComplete

`func (o *PostCollectionDeployment201Response1) HasPercentComplete() bool`

HasPercentComplete returns a boolean if a field has been set.

### GetStartDateTime

`func (o *PostCollectionDeployment201Response1) GetStartDateTime() time.Time`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *PostCollectionDeployment201Response1) GetStartDateTimeOk() (*time.Time, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *PostCollectionDeployment201Response1) SetStartDateTime(v time.Time)`

SetStartDateTime sets StartDateTime field to given value.

### HasStartDateTime

`func (o *PostCollectionDeployment201Response1) HasStartDateTime() bool`

HasStartDateTime returns a boolean if a field has been set.

### GetCompletionDateTime

`func (o *PostCollectionDeployment201Response1) GetCompletionDateTime() time.Time`

GetCompletionDateTime returns the CompletionDateTime field if non-nil, zero value otherwise.

### GetCompletionDateTimeOk

`func (o *PostCollectionDeployment201Response1) GetCompletionDateTimeOk() (*time.Time, bool)`

GetCompletionDateTimeOk returns a tuple with the CompletionDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionDateTime

`func (o *PostCollectionDeployment201Response1) SetCompletionDateTime(v time.Time)`

SetCompletionDateTime sets CompletionDateTime field to given value.

### HasCompletionDateTime

`func (o *PostCollectionDeployment201Response1) HasCompletionDateTime() bool`

HasCompletionDateTime returns a boolean if a field has been set.

### GetUser

`func (o *PostCollectionDeployment201Response1) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *PostCollectionDeployment201Response1) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *PostCollectionDeployment201Response1) SetUser(v User)`

SetUser sets User field to given value.

### HasUser

`func (o *PostCollectionDeployment201Response1) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetMethod

`func (o *PostCollectionDeployment201Response1) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *PostCollectionDeployment201Response1) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *PostCollectionDeployment201Response1) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *PostCollectionDeployment201Response1) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetService

`func (o *PostCollectionDeployment201Response1) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *PostCollectionDeployment201Response1) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *PostCollectionDeployment201Response1) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *PostCollectionDeployment201Response1) HasService() bool`

HasService returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


