# TemplateApplication

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

### NewTemplateApplication

`func NewTemplateApplication() *TemplateApplication`

NewTemplateApplication instantiates a new TemplateApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateApplicationWithDefaults

`func NewTemplateApplicationWithDefaults() *TemplateApplication`

NewTemplateApplicationWithDefaults instantiates a new TemplateApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TemplateApplication) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TemplateApplication) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TemplateApplication) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *TemplateApplication) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *TemplateApplication) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TemplateApplication) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TemplateApplication) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TemplateApplication) HasType() bool`

HasType returns a boolean if a field has been set.

### GetState

`func (o *TemplateApplication) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TemplateApplication) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TemplateApplication) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *TemplateApplication) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStartDateTime

`func (o *TemplateApplication) GetStartDateTime() time.Time`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *TemplateApplication) GetStartDateTimeOk() (*time.Time, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *TemplateApplication) SetStartDateTime(v time.Time)`

SetStartDateTime sets StartDateTime field to given value.

### HasStartDateTime

`func (o *TemplateApplication) HasStartDateTime() bool`

HasStartDateTime returns a boolean if a field has been set.

### GetCompletionDateTime

`func (o *TemplateApplication) GetCompletionDateTime() time.Time`

GetCompletionDateTime returns the CompletionDateTime field if non-nil, zero value otherwise.

### GetCompletionDateTimeOk

`func (o *TemplateApplication) GetCompletionDateTimeOk() (*time.Time, bool)`

GetCompletionDateTimeOk returns a tuple with the CompletionDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionDateTime

`func (o *TemplateApplication) SetCompletionDateTime(v time.Time)`

SetCompletionDateTime sets CompletionDateTime field to given value.

### HasCompletionDateTime

`func (o *TemplateApplication) HasCompletionDateTime() bool`

HasCompletionDateTime returns a boolean if a field has been set.

### GetConflictResolutionStrategy

`func (o *TemplateApplication) GetConflictResolutionStrategy() string`

GetConflictResolutionStrategy returns the ConflictResolutionStrategy field if non-nil, zero value otherwise.

### GetConflictResolutionStrategyOk

`func (o *TemplateApplication) GetConflictResolutionStrategyOk() (*string, bool)`

GetConflictResolutionStrategyOk returns a tuple with the ConflictResolutionStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflictResolutionStrategy

`func (o *TemplateApplication) SetConflictResolutionStrategy(v string)`

SetConflictResolutionStrategy sets ConflictResolutionStrategy field to given value.

### HasConflictResolutionStrategy

`func (o *TemplateApplication) HasConflictResolutionStrategy() bool`

HasConflictResolutionStrategy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


