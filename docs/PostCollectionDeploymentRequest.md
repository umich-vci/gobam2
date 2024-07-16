# PostCollectionDeploymentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | **string** | The resource type. | 
**State** | Pointer to **string** | The current state of the deployment. | [optional] 
**Status** | Pointer to **string** | The current status of the deployment. | [optional] [readonly] 
**Message** | Pointer to **string** | The expanded status message for the deployment. | [optional] [readonly] 
**PercentComplete** | Pointer to **int32** | A percentage value representing the current progress of the deployment. | [optional] [readonly] 
**StartDateTime** | Pointer to **time.Time** | The start time of the deployment. | [optional] [readonly] 
**CompletionDateTime** | Pointer to **time.Time** | The completion time for the deployment. | [optional] [readonly] 
**User** | Pointer to [**User**](User.md) |  | [optional] [readonly] 
**Method** | Pointer to **string** | The method that triggered the deployment. | [optional] [readonly] 
**Service** | **string** | The deployment service type. | 

## Methods

### NewPostCollectionDeploymentRequest

`func NewPostCollectionDeploymentRequest(type_ string, service string, ) *PostCollectionDeploymentRequest`

NewPostCollectionDeploymentRequest instantiates a new PostCollectionDeploymentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostCollectionDeploymentRequestWithDefaults

`func NewPostCollectionDeploymentRequestWithDefaults() *PostCollectionDeploymentRequest`

NewPostCollectionDeploymentRequestWithDefaults instantiates a new PostCollectionDeploymentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PostCollectionDeploymentRequest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostCollectionDeploymentRequest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostCollectionDeploymentRequest) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *PostCollectionDeploymentRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *PostCollectionDeploymentRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PostCollectionDeploymentRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PostCollectionDeploymentRequest) SetType(v string)`

SetType sets Type field to given value.


### GetState

`func (o *PostCollectionDeploymentRequest) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PostCollectionDeploymentRequest) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PostCollectionDeploymentRequest) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *PostCollectionDeploymentRequest) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStatus

`func (o *PostCollectionDeploymentRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PostCollectionDeploymentRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PostCollectionDeploymentRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PostCollectionDeploymentRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *PostCollectionDeploymentRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PostCollectionDeploymentRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PostCollectionDeploymentRequest) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *PostCollectionDeploymentRequest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetPercentComplete

`func (o *PostCollectionDeploymentRequest) GetPercentComplete() int32`

GetPercentComplete returns the PercentComplete field if non-nil, zero value otherwise.

### GetPercentCompleteOk

`func (o *PostCollectionDeploymentRequest) GetPercentCompleteOk() (*int32, bool)`

GetPercentCompleteOk returns a tuple with the PercentComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentComplete

`func (o *PostCollectionDeploymentRequest) SetPercentComplete(v int32)`

SetPercentComplete sets PercentComplete field to given value.

### HasPercentComplete

`func (o *PostCollectionDeploymentRequest) HasPercentComplete() bool`

HasPercentComplete returns a boolean if a field has been set.

### GetStartDateTime

`func (o *PostCollectionDeploymentRequest) GetStartDateTime() time.Time`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *PostCollectionDeploymentRequest) GetStartDateTimeOk() (*time.Time, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *PostCollectionDeploymentRequest) SetStartDateTime(v time.Time)`

SetStartDateTime sets StartDateTime field to given value.

### HasStartDateTime

`func (o *PostCollectionDeploymentRequest) HasStartDateTime() bool`

HasStartDateTime returns a boolean if a field has been set.

### GetCompletionDateTime

`func (o *PostCollectionDeploymentRequest) GetCompletionDateTime() time.Time`

GetCompletionDateTime returns the CompletionDateTime field if non-nil, zero value otherwise.

### GetCompletionDateTimeOk

`func (o *PostCollectionDeploymentRequest) GetCompletionDateTimeOk() (*time.Time, bool)`

GetCompletionDateTimeOk returns a tuple with the CompletionDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionDateTime

`func (o *PostCollectionDeploymentRequest) SetCompletionDateTime(v time.Time)`

SetCompletionDateTime sets CompletionDateTime field to given value.

### HasCompletionDateTime

`func (o *PostCollectionDeploymentRequest) HasCompletionDateTime() bool`

HasCompletionDateTime returns a boolean if a field has been set.

### GetUser

`func (o *PostCollectionDeploymentRequest) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *PostCollectionDeploymentRequest) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *PostCollectionDeploymentRequest) SetUser(v User)`

SetUser sets User field to given value.

### HasUser

`func (o *PostCollectionDeploymentRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetMethod

`func (o *PostCollectionDeploymentRequest) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *PostCollectionDeploymentRequest) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *PostCollectionDeploymentRequest) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *PostCollectionDeploymentRequest) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetService

`func (o *PostCollectionDeploymentRequest) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *PostCollectionDeploymentRequest) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *PostCollectionDeploymentRequest) SetService(v string)`

SetService sets Service field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


