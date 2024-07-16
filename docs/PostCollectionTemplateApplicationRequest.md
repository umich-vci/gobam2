# PostCollectionTemplateApplicationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | **string** | The resource type. | 
**State** | Pointer to **string** | The state of the template application operation. | [optional] [readonly] 
**StartDateTime** | Pointer to **time.Time** | The date and time that the template application operation started. | [optional] [readonly] 
**CompletionDateTime** | Pointer to **time.Time** | The date and time that the template application operation completed. | [optional] [readonly] 
**ConflictResolutionStrategy** | Pointer to **string** | The method used to resolve conflicts between template items and the resource the template is applied to. | [optional] 

## Methods

### NewPostCollectionTemplateApplicationRequest

`func NewPostCollectionTemplateApplicationRequest(type_ string, ) *PostCollectionTemplateApplicationRequest`

NewPostCollectionTemplateApplicationRequest instantiates a new PostCollectionTemplateApplicationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostCollectionTemplateApplicationRequestWithDefaults

`func NewPostCollectionTemplateApplicationRequestWithDefaults() *PostCollectionTemplateApplicationRequest`

NewPostCollectionTemplateApplicationRequestWithDefaults instantiates a new PostCollectionTemplateApplicationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PostCollectionTemplateApplicationRequest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostCollectionTemplateApplicationRequest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostCollectionTemplateApplicationRequest) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *PostCollectionTemplateApplicationRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *PostCollectionTemplateApplicationRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PostCollectionTemplateApplicationRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PostCollectionTemplateApplicationRequest) SetType(v string)`

SetType sets Type field to given value.


### GetState

`func (o *PostCollectionTemplateApplicationRequest) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PostCollectionTemplateApplicationRequest) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PostCollectionTemplateApplicationRequest) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *PostCollectionTemplateApplicationRequest) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStartDateTime

`func (o *PostCollectionTemplateApplicationRequest) GetStartDateTime() time.Time`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *PostCollectionTemplateApplicationRequest) GetStartDateTimeOk() (*time.Time, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *PostCollectionTemplateApplicationRequest) SetStartDateTime(v time.Time)`

SetStartDateTime sets StartDateTime field to given value.

### HasStartDateTime

`func (o *PostCollectionTemplateApplicationRequest) HasStartDateTime() bool`

HasStartDateTime returns a boolean if a field has been set.

### GetCompletionDateTime

`func (o *PostCollectionTemplateApplicationRequest) GetCompletionDateTime() time.Time`

GetCompletionDateTime returns the CompletionDateTime field if non-nil, zero value otherwise.

### GetCompletionDateTimeOk

`func (o *PostCollectionTemplateApplicationRequest) GetCompletionDateTimeOk() (*time.Time, bool)`

GetCompletionDateTimeOk returns a tuple with the CompletionDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionDateTime

`func (o *PostCollectionTemplateApplicationRequest) SetCompletionDateTime(v time.Time)`

SetCompletionDateTime sets CompletionDateTime field to given value.

### HasCompletionDateTime

`func (o *PostCollectionTemplateApplicationRequest) HasCompletionDateTime() bool`

HasCompletionDateTime returns a boolean if a field has been set.

### GetConflictResolutionStrategy

`func (o *PostCollectionTemplateApplicationRequest) GetConflictResolutionStrategy() string`

GetConflictResolutionStrategy returns the ConflictResolutionStrategy field if non-nil, zero value otherwise.

### GetConflictResolutionStrategyOk

`func (o *PostCollectionTemplateApplicationRequest) GetConflictResolutionStrategyOk() (*string, bool)`

GetConflictResolutionStrategyOk returns a tuple with the ConflictResolutionStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflictResolutionStrategy

`func (o *PostCollectionTemplateApplicationRequest) SetConflictResolutionStrategy(v string)`

SetConflictResolutionStrategy sets ConflictResolutionStrategy field to given value.

### HasConflictResolutionStrategy

`func (o *PostCollectionTemplateApplicationRequest) HasConflictResolutionStrategy() bool`

HasConflictResolutionStrategy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


