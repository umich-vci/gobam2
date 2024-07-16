# GetDeployments200ResponseDataInner

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
**Resources** | Pointer to [**[]InlinedDeployableResource**](InlinedDeployableResource.md) | A list of resource record and IP address resources to deploy. | [optional] 
**BatchMode** | Pointer to **string** | The method used for batching selective deployment tasks. | [optional] 
**RecordScope** | Pointer to **string** | Defines whether selective deployment tasks include resources that are related to the defined DNS records. | [optional] 
**ContinueOnFailureEnabled** | Pointer to **bool** | Indicates the mode of operation on a failed resource record. If false, the deployment stops when a record fails. If true, the deployment continues when a record fails and the deployment moves to the next record. | [optional] 
**DynamicRecordDeploymentStrategy** | Pointer to **string** | Defines how dynamic records are handled with selective deployment tasks. | [optional] 
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 

## Methods

### NewGetDeployments200ResponseDataInner

`func NewGetDeployments200ResponseDataInner() *GetDeployments200ResponseDataInner`

NewGetDeployments200ResponseDataInner instantiates a new GetDeployments200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDeployments200ResponseDataInnerWithDefaults

`func NewGetDeployments200ResponseDataInnerWithDefaults() *GetDeployments200ResponseDataInner`

NewGetDeployments200ResponseDataInnerWithDefaults instantiates a new GetDeployments200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetDeployments200ResponseDataInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetDeployments200ResponseDataInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetDeployments200ResponseDataInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetDeployments200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GetDeployments200ResponseDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetDeployments200ResponseDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetDeployments200ResponseDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetDeployments200ResponseDataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetState

`func (o *GetDeployments200ResponseDataInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetDeployments200ResponseDataInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetDeployments200ResponseDataInner) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GetDeployments200ResponseDataInner) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStatus

`func (o *GetDeployments200ResponseDataInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetDeployments200ResponseDataInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetDeployments200ResponseDataInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetDeployments200ResponseDataInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *GetDeployments200ResponseDataInner) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetDeployments200ResponseDataInner) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetDeployments200ResponseDataInner) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetDeployments200ResponseDataInner) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetPercentComplete

`func (o *GetDeployments200ResponseDataInner) GetPercentComplete() int32`

GetPercentComplete returns the PercentComplete field if non-nil, zero value otherwise.

### GetPercentCompleteOk

`func (o *GetDeployments200ResponseDataInner) GetPercentCompleteOk() (*int32, bool)`

GetPercentCompleteOk returns a tuple with the PercentComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentComplete

`func (o *GetDeployments200ResponseDataInner) SetPercentComplete(v int32)`

SetPercentComplete sets PercentComplete field to given value.

### HasPercentComplete

`func (o *GetDeployments200ResponseDataInner) HasPercentComplete() bool`

HasPercentComplete returns a boolean if a field has been set.

### GetStartDateTime

`func (o *GetDeployments200ResponseDataInner) GetStartDateTime() time.Time`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *GetDeployments200ResponseDataInner) GetStartDateTimeOk() (*time.Time, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *GetDeployments200ResponseDataInner) SetStartDateTime(v time.Time)`

SetStartDateTime sets StartDateTime field to given value.

### HasStartDateTime

`func (o *GetDeployments200ResponseDataInner) HasStartDateTime() bool`

HasStartDateTime returns a boolean if a field has been set.

### GetCompletionDateTime

`func (o *GetDeployments200ResponseDataInner) GetCompletionDateTime() time.Time`

GetCompletionDateTime returns the CompletionDateTime field if non-nil, zero value otherwise.

### GetCompletionDateTimeOk

`func (o *GetDeployments200ResponseDataInner) GetCompletionDateTimeOk() (*time.Time, bool)`

GetCompletionDateTimeOk returns a tuple with the CompletionDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionDateTime

`func (o *GetDeployments200ResponseDataInner) SetCompletionDateTime(v time.Time)`

SetCompletionDateTime sets CompletionDateTime field to given value.

### HasCompletionDateTime

`func (o *GetDeployments200ResponseDataInner) HasCompletionDateTime() bool`

HasCompletionDateTime returns a boolean if a field has been set.

### GetUser

`func (o *GetDeployments200ResponseDataInner) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *GetDeployments200ResponseDataInner) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *GetDeployments200ResponseDataInner) SetUser(v User)`

SetUser sets User field to given value.

### HasUser

`func (o *GetDeployments200ResponseDataInner) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetMethod

`func (o *GetDeployments200ResponseDataInner) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *GetDeployments200ResponseDataInner) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *GetDeployments200ResponseDataInner) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *GetDeployments200ResponseDataInner) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetService

`func (o *GetDeployments200ResponseDataInner) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *GetDeployments200ResponseDataInner) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *GetDeployments200ResponseDataInner) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *GetDeployments200ResponseDataInner) HasService() bool`

HasService returns a boolean if a field has been set.

### GetResources

`func (o *GetDeployments200ResponseDataInner) GetResources() []InlinedDeployableResource`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *GetDeployments200ResponseDataInner) GetResourcesOk() (*[]InlinedDeployableResource, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *GetDeployments200ResponseDataInner) SetResources(v []InlinedDeployableResource)`

SetResources sets Resources field to given value.

### HasResources

`func (o *GetDeployments200ResponseDataInner) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetBatchMode

`func (o *GetDeployments200ResponseDataInner) GetBatchMode() string`

GetBatchMode returns the BatchMode field if non-nil, zero value otherwise.

### GetBatchModeOk

`func (o *GetDeployments200ResponseDataInner) GetBatchModeOk() (*string, bool)`

GetBatchModeOk returns a tuple with the BatchMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchMode

`func (o *GetDeployments200ResponseDataInner) SetBatchMode(v string)`

SetBatchMode sets BatchMode field to given value.

### HasBatchMode

`func (o *GetDeployments200ResponseDataInner) HasBatchMode() bool`

HasBatchMode returns a boolean if a field has been set.

### GetRecordScope

`func (o *GetDeployments200ResponseDataInner) GetRecordScope() string`

GetRecordScope returns the RecordScope field if non-nil, zero value otherwise.

### GetRecordScopeOk

`func (o *GetDeployments200ResponseDataInner) GetRecordScopeOk() (*string, bool)`

GetRecordScopeOk returns a tuple with the RecordScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordScope

`func (o *GetDeployments200ResponseDataInner) SetRecordScope(v string)`

SetRecordScope sets RecordScope field to given value.

### HasRecordScope

`func (o *GetDeployments200ResponseDataInner) HasRecordScope() bool`

HasRecordScope returns a boolean if a field has been set.

### GetContinueOnFailureEnabled

`func (o *GetDeployments200ResponseDataInner) GetContinueOnFailureEnabled() bool`

GetContinueOnFailureEnabled returns the ContinueOnFailureEnabled field if non-nil, zero value otherwise.

### GetContinueOnFailureEnabledOk

`func (o *GetDeployments200ResponseDataInner) GetContinueOnFailureEnabledOk() (*bool, bool)`

GetContinueOnFailureEnabledOk returns a tuple with the ContinueOnFailureEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnFailureEnabled

`func (o *GetDeployments200ResponseDataInner) SetContinueOnFailureEnabled(v bool)`

SetContinueOnFailureEnabled sets ContinueOnFailureEnabled field to given value.

### HasContinueOnFailureEnabled

`func (o *GetDeployments200ResponseDataInner) HasContinueOnFailureEnabled() bool`

HasContinueOnFailureEnabled returns a boolean if a field has been set.

### GetDynamicRecordDeploymentStrategy

`func (o *GetDeployments200ResponseDataInner) GetDynamicRecordDeploymentStrategy() string`

GetDynamicRecordDeploymentStrategy returns the DynamicRecordDeploymentStrategy field if non-nil, zero value otherwise.

### GetDynamicRecordDeploymentStrategyOk

`func (o *GetDeployments200ResponseDataInner) GetDynamicRecordDeploymentStrategyOk() (*string, bool)`

GetDynamicRecordDeploymentStrategyOk returns a tuple with the DynamicRecordDeploymentStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicRecordDeploymentStrategy

`func (o *GetDeployments200ResponseDataInner) SetDynamicRecordDeploymentStrategy(v string)`

SetDynamicRecordDeploymentStrategy sets DynamicRecordDeploymentStrategy field to given value.

### HasDynamicRecordDeploymentStrategy

`func (o *GetDeployments200ResponseDataInner) HasDynamicRecordDeploymentStrategy() bool`

HasDynamicRecordDeploymentStrategy returns a boolean if a field has been set.

### GetLinks

`func (o *GetDeployments200ResponseDataInner) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetDeployments200ResponseDataInner) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetDeployments200ResponseDataInner) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetDeployments200ResponseDataInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


