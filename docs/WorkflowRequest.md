# WorkflowRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**State** | Pointer to **string** | The approval state of the workflow request. The state will be PENDING if the workflow request has not yet been processed by an authorized user. | [optional] 
**RequestType** | Pointer to **string** | The requested operation for the workflow request. | [optional] [readonly] 
**Creator** | Pointer to [**User**](User.md) |  | [optional] [readonly] 
**Resource** | Pointer to [**Resource**](Resource.md) |  | [optional] [readonly] 
**DependentResource** | Pointer to [**Resource**](Resource.md) |  | [optional] [readonly] 
**Modifier** | Pointer to [**User**](User.md) |  | [optional] [readonly] 
**CreationDateTime** | Pointer to **time.Time** | The time the workflow request was created. | [optional] [readonly] 
**ModificationDateTime** | Pointer to **time.Time** | The time the workflow request was modified. | [optional] [readonly] 
**Comment** | Pointer to **string** | The change control comment added during creation of the workflow request. | [optional] [readonly] 

## Methods

### NewWorkflowRequest

`func NewWorkflowRequest() *WorkflowRequest`

NewWorkflowRequest instantiates a new WorkflowRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowRequestWithDefaults

`func NewWorkflowRequestWithDefaults() *WorkflowRequest`

NewWorkflowRequestWithDefaults instantiates a new WorkflowRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkflowRequest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowRequest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowRequest) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *WorkflowRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *WorkflowRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WorkflowRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WorkflowRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WorkflowRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetState

`func (o *WorkflowRequest) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *WorkflowRequest) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *WorkflowRequest) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *WorkflowRequest) HasState() bool`

HasState returns a boolean if a field has been set.

### GetRequestType

`func (o *WorkflowRequest) GetRequestType() string`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *WorkflowRequest) GetRequestTypeOk() (*string, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *WorkflowRequest) SetRequestType(v string)`

SetRequestType sets RequestType field to given value.

### HasRequestType

`func (o *WorkflowRequest) HasRequestType() bool`

HasRequestType returns a boolean if a field has been set.

### GetCreator

`func (o *WorkflowRequest) GetCreator() User`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *WorkflowRequest) GetCreatorOk() (*User, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *WorkflowRequest) SetCreator(v User)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *WorkflowRequest) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetResource

`func (o *WorkflowRequest) GetResource() Resource`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *WorkflowRequest) GetResourceOk() (*Resource, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *WorkflowRequest) SetResource(v Resource)`

SetResource sets Resource field to given value.

### HasResource

`func (o *WorkflowRequest) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetDependentResource

`func (o *WorkflowRequest) GetDependentResource() Resource`

GetDependentResource returns the DependentResource field if non-nil, zero value otherwise.

### GetDependentResourceOk

`func (o *WorkflowRequest) GetDependentResourceOk() (*Resource, bool)`

GetDependentResourceOk returns a tuple with the DependentResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependentResource

`func (o *WorkflowRequest) SetDependentResource(v Resource)`

SetDependentResource sets DependentResource field to given value.

### HasDependentResource

`func (o *WorkflowRequest) HasDependentResource() bool`

HasDependentResource returns a boolean if a field has been set.

### GetModifier

`func (o *WorkflowRequest) GetModifier() User`

GetModifier returns the Modifier field if non-nil, zero value otherwise.

### GetModifierOk

`func (o *WorkflowRequest) GetModifierOk() (*User, bool)`

GetModifierOk returns a tuple with the Modifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifier

`func (o *WorkflowRequest) SetModifier(v User)`

SetModifier sets Modifier field to given value.

### HasModifier

`func (o *WorkflowRequest) HasModifier() bool`

HasModifier returns a boolean if a field has been set.

### GetCreationDateTime

`func (o *WorkflowRequest) GetCreationDateTime() time.Time`

GetCreationDateTime returns the CreationDateTime field if non-nil, zero value otherwise.

### GetCreationDateTimeOk

`func (o *WorkflowRequest) GetCreationDateTimeOk() (*time.Time, bool)`

GetCreationDateTimeOk returns a tuple with the CreationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDateTime

`func (o *WorkflowRequest) SetCreationDateTime(v time.Time)`

SetCreationDateTime sets CreationDateTime field to given value.

### HasCreationDateTime

`func (o *WorkflowRequest) HasCreationDateTime() bool`

HasCreationDateTime returns a boolean if a field has been set.

### GetModificationDateTime

`func (o *WorkflowRequest) GetModificationDateTime() time.Time`

GetModificationDateTime returns the ModificationDateTime field if non-nil, zero value otherwise.

### GetModificationDateTimeOk

`func (o *WorkflowRequest) GetModificationDateTimeOk() (*time.Time, bool)`

GetModificationDateTimeOk returns a tuple with the ModificationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModificationDateTime

`func (o *WorkflowRequest) SetModificationDateTime(v time.Time)`

SetModificationDateTime sets ModificationDateTime field to given value.

### HasModificationDateTime

`func (o *WorkflowRequest) HasModificationDateTime() bool`

HasModificationDateTime returns a boolean if a field has been set.

### GetComment

`func (o *WorkflowRequest) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *WorkflowRequest) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *WorkflowRequest) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *WorkflowRequest) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


