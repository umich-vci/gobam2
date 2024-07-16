# GetTemplateApplications200Response1DataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**State** | Pointer to **string** | The state of the template application operation. | [optional] [readonly] 
**StartDateTime** | Pointer to **time.Time** | The date and time that the template application operation started. | [optional] [readonly] 
**CompletionDateTime** | Pointer to **time.Time** | The date and time that the template application operation completed. | [optional] [readonly] 
**ConflictResolutionStrategy** | Pointer to **string** | The method used to resolve conflicts between template items and the resource the template is applied to. | [optional] 

## Methods

### NewGetTemplateApplications200Response1DataInner

`func NewGetTemplateApplications200Response1DataInner() *GetTemplateApplications200Response1DataInner`

NewGetTemplateApplications200Response1DataInner instantiates a new GetTemplateApplications200Response1DataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTemplateApplications200Response1DataInnerWithDefaults

`func NewGetTemplateApplications200Response1DataInnerWithDefaults() *GetTemplateApplications200Response1DataInner`

NewGetTemplateApplications200Response1DataInnerWithDefaults instantiates a new GetTemplateApplications200Response1DataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetTemplateApplications200Response1DataInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetTemplateApplications200Response1DataInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetTemplateApplications200Response1DataInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetTemplateApplications200Response1DataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GetTemplateApplications200Response1DataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetTemplateApplications200Response1DataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetTemplateApplications200Response1DataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetTemplateApplications200Response1DataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetState

`func (o *GetTemplateApplications200Response1DataInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetTemplateApplications200Response1DataInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetTemplateApplications200Response1DataInner) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GetTemplateApplications200Response1DataInner) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStartDateTime

`func (o *GetTemplateApplications200Response1DataInner) GetStartDateTime() time.Time`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *GetTemplateApplications200Response1DataInner) GetStartDateTimeOk() (*time.Time, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *GetTemplateApplications200Response1DataInner) SetStartDateTime(v time.Time)`

SetStartDateTime sets StartDateTime field to given value.

### HasStartDateTime

`func (o *GetTemplateApplications200Response1DataInner) HasStartDateTime() bool`

HasStartDateTime returns a boolean if a field has been set.

### GetCompletionDateTime

`func (o *GetTemplateApplications200Response1DataInner) GetCompletionDateTime() time.Time`

GetCompletionDateTime returns the CompletionDateTime field if non-nil, zero value otherwise.

### GetCompletionDateTimeOk

`func (o *GetTemplateApplications200Response1DataInner) GetCompletionDateTimeOk() (*time.Time, bool)`

GetCompletionDateTimeOk returns a tuple with the CompletionDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionDateTime

`func (o *GetTemplateApplications200Response1DataInner) SetCompletionDateTime(v time.Time)`

SetCompletionDateTime sets CompletionDateTime field to given value.

### HasCompletionDateTime

`func (o *GetTemplateApplications200Response1DataInner) HasCompletionDateTime() bool`

HasCompletionDateTime returns a boolean if a field has been set.

### GetConflictResolutionStrategy

`func (o *GetTemplateApplications200Response1DataInner) GetConflictResolutionStrategy() string`

GetConflictResolutionStrategy returns the ConflictResolutionStrategy field if non-nil, zero value otherwise.

### GetConflictResolutionStrategyOk

`func (o *GetTemplateApplications200Response1DataInner) GetConflictResolutionStrategyOk() (*string, bool)`

GetConflictResolutionStrategyOk returns a tuple with the ConflictResolutionStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflictResolutionStrategy

`func (o *GetTemplateApplications200Response1DataInner) SetConflictResolutionStrategy(v string)`

SetConflictResolutionStrategy sets ConflictResolutionStrategy field to given value.

### HasConflictResolutionStrategy

`func (o *GetTemplateApplications200Response1DataInner) HasConflictResolutionStrategy() bool`

HasConflictResolutionStrategy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


