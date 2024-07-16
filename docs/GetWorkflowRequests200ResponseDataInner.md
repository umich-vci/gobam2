# GetWorkflowRequests200ResponseDataInner

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
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 

## Methods

### NewGetWorkflowRequests200ResponseDataInner

`func NewGetWorkflowRequests200ResponseDataInner() *GetWorkflowRequests200ResponseDataInner`

NewGetWorkflowRequests200ResponseDataInner instantiates a new GetWorkflowRequests200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetWorkflowRequests200ResponseDataInnerWithDefaults

`func NewGetWorkflowRequests200ResponseDataInnerWithDefaults() *GetWorkflowRequests200ResponseDataInner`

NewGetWorkflowRequests200ResponseDataInnerWithDefaults instantiates a new GetWorkflowRequests200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetWorkflowRequests200ResponseDataInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetWorkflowRequests200ResponseDataInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetWorkflowRequests200ResponseDataInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetWorkflowRequests200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GetWorkflowRequests200ResponseDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetWorkflowRequests200ResponseDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetWorkflowRequests200ResponseDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetWorkflowRequests200ResponseDataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetState

`func (o *GetWorkflowRequests200ResponseDataInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetWorkflowRequests200ResponseDataInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetWorkflowRequests200ResponseDataInner) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GetWorkflowRequests200ResponseDataInner) HasState() bool`

HasState returns a boolean if a field has been set.

### GetRequestType

`func (o *GetWorkflowRequests200ResponseDataInner) GetRequestType() string`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *GetWorkflowRequests200ResponseDataInner) GetRequestTypeOk() (*string, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *GetWorkflowRequests200ResponseDataInner) SetRequestType(v string)`

SetRequestType sets RequestType field to given value.

### HasRequestType

`func (o *GetWorkflowRequests200ResponseDataInner) HasRequestType() bool`

HasRequestType returns a boolean if a field has been set.

### GetCreator

`func (o *GetWorkflowRequests200ResponseDataInner) GetCreator() User`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *GetWorkflowRequests200ResponseDataInner) GetCreatorOk() (*User, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *GetWorkflowRequests200ResponseDataInner) SetCreator(v User)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *GetWorkflowRequests200ResponseDataInner) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetResource

`func (o *GetWorkflowRequests200ResponseDataInner) GetResource() Resource`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *GetWorkflowRequests200ResponseDataInner) GetResourceOk() (*Resource, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *GetWorkflowRequests200ResponseDataInner) SetResource(v Resource)`

SetResource sets Resource field to given value.

### HasResource

`func (o *GetWorkflowRequests200ResponseDataInner) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetDependentResource

`func (o *GetWorkflowRequests200ResponseDataInner) GetDependentResource() Resource`

GetDependentResource returns the DependentResource field if non-nil, zero value otherwise.

### GetDependentResourceOk

`func (o *GetWorkflowRequests200ResponseDataInner) GetDependentResourceOk() (*Resource, bool)`

GetDependentResourceOk returns a tuple with the DependentResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependentResource

`func (o *GetWorkflowRequests200ResponseDataInner) SetDependentResource(v Resource)`

SetDependentResource sets DependentResource field to given value.

### HasDependentResource

`func (o *GetWorkflowRequests200ResponseDataInner) HasDependentResource() bool`

HasDependentResource returns a boolean if a field has been set.

### GetModifier

`func (o *GetWorkflowRequests200ResponseDataInner) GetModifier() User`

GetModifier returns the Modifier field if non-nil, zero value otherwise.

### GetModifierOk

`func (o *GetWorkflowRequests200ResponseDataInner) GetModifierOk() (*User, bool)`

GetModifierOk returns a tuple with the Modifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifier

`func (o *GetWorkflowRequests200ResponseDataInner) SetModifier(v User)`

SetModifier sets Modifier field to given value.

### HasModifier

`func (o *GetWorkflowRequests200ResponseDataInner) HasModifier() bool`

HasModifier returns a boolean if a field has been set.

### GetCreationDateTime

`func (o *GetWorkflowRequests200ResponseDataInner) GetCreationDateTime() time.Time`

GetCreationDateTime returns the CreationDateTime field if non-nil, zero value otherwise.

### GetCreationDateTimeOk

`func (o *GetWorkflowRequests200ResponseDataInner) GetCreationDateTimeOk() (*time.Time, bool)`

GetCreationDateTimeOk returns a tuple with the CreationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDateTime

`func (o *GetWorkflowRequests200ResponseDataInner) SetCreationDateTime(v time.Time)`

SetCreationDateTime sets CreationDateTime field to given value.

### HasCreationDateTime

`func (o *GetWorkflowRequests200ResponseDataInner) HasCreationDateTime() bool`

HasCreationDateTime returns a boolean if a field has been set.

### GetModificationDateTime

`func (o *GetWorkflowRequests200ResponseDataInner) GetModificationDateTime() time.Time`

GetModificationDateTime returns the ModificationDateTime field if non-nil, zero value otherwise.

### GetModificationDateTimeOk

`func (o *GetWorkflowRequests200ResponseDataInner) GetModificationDateTimeOk() (*time.Time, bool)`

GetModificationDateTimeOk returns a tuple with the ModificationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModificationDateTime

`func (o *GetWorkflowRequests200ResponseDataInner) SetModificationDateTime(v time.Time)`

SetModificationDateTime sets ModificationDateTime field to given value.

### HasModificationDateTime

`func (o *GetWorkflowRequests200ResponseDataInner) HasModificationDateTime() bool`

HasModificationDateTime returns a boolean if a field has been set.

### GetComment

`func (o *GetWorkflowRequests200ResponseDataInner) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetWorkflowRequests200ResponseDataInner) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetWorkflowRequests200ResponseDataInner) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetWorkflowRequests200ResponseDataInner) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetLinks

`func (o *GetWorkflowRequests200ResponseDataInner) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetWorkflowRequests200ResponseDataInner) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetWorkflowRequests200ResponseDataInner) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetWorkflowRequests200ResponseDataInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


