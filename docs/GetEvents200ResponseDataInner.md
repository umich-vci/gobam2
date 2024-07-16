# GetEvents200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**EventType** | Pointer to **string** | The log level of the event. | [optional] 
**Category** | Pointer to **string** | The classification of the event, such as an APPLICATION or DEPLOYMENT event. | [optional] 
**Subcategory** | Pointer to **string** | The subcategory classification of the event. | [optional] 
**Source** | Pointer to [**Resource**](Resource.md) |  | [optional] 
**SourceName** | Pointer to **string** | The name of the software component involved in the event. | [optional] 
**Message** | Pointer to **string** | The message describing the event. | [optional] 
**DateTime** | Pointer to **time.Time** | The date and time that the event occurred. | [optional] 
**Details** | Pointer to **map[string]interface{}** |  | [optional] 
**TaskId** | Pointer to **string** | The ID of the task performed. | [optional] 
**Payload** | Pointer to **string** | The payload associated with the task. | [optional] 
**User** | Pointer to [**User**](User.md) |  | [optional] 
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 

## Methods

### NewGetEvents200ResponseDataInner

`func NewGetEvents200ResponseDataInner() *GetEvents200ResponseDataInner`

NewGetEvents200ResponseDataInner instantiates a new GetEvents200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEvents200ResponseDataInnerWithDefaults

`func NewGetEvents200ResponseDataInnerWithDefaults() *GetEvents200ResponseDataInner`

NewGetEvents200ResponseDataInnerWithDefaults instantiates a new GetEvents200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetEvents200ResponseDataInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetEvents200ResponseDataInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetEvents200ResponseDataInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetEvents200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GetEvents200ResponseDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetEvents200ResponseDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetEvents200ResponseDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetEvents200ResponseDataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetEventType

`func (o *GetEvents200ResponseDataInner) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *GetEvents200ResponseDataInner) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *GetEvents200ResponseDataInner) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *GetEvents200ResponseDataInner) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetCategory

`func (o *GetEvents200ResponseDataInner) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *GetEvents200ResponseDataInner) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *GetEvents200ResponseDataInner) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *GetEvents200ResponseDataInner) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetSubcategory

`func (o *GetEvents200ResponseDataInner) GetSubcategory() string`

GetSubcategory returns the Subcategory field if non-nil, zero value otherwise.

### GetSubcategoryOk

`func (o *GetEvents200ResponseDataInner) GetSubcategoryOk() (*string, bool)`

GetSubcategoryOk returns a tuple with the Subcategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubcategory

`func (o *GetEvents200ResponseDataInner) SetSubcategory(v string)`

SetSubcategory sets Subcategory field to given value.

### HasSubcategory

`func (o *GetEvents200ResponseDataInner) HasSubcategory() bool`

HasSubcategory returns a boolean if a field has been set.

### GetSource

`func (o *GetEvents200ResponseDataInner) GetSource() Resource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GetEvents200ResponseDataInner) GetSourceOk() (*Resource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GetEvents200ResponseDataInner) SetSource(v Resource)`

SetSource sets Source field to given value.

### HasSource

`func (o *GetEvents200ResponseDataInner) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSourceName

`func (o *GetEvents200ResponseDataInner) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *GetEvents200ResponseDataInner) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *GetEvents200ResponseDataInner) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.

### HasSourceName

`func (o *GetEvents200ResponseDataInner) HasSourceName() bool`

HasSourceName returns a boolean if a field has been set.

### GetMessage

`func (o *GetEvents200ResponseDataInner) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetEvents200ResponseDataInner) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetEvents200ResponseDataInner) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetEvents200ResponseDataInner) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetDateTime

`func (o *GetEvents200ResponseDataInner) GetDateTime() time.Time`

GetDateTime returns the DateTime field if non-nil, zero value otherwise.

### GetDateTimeOk

`func (o *GetEvents200ResponseDataInner) GetDateTimeOk() (*time.Time, bool)`

GetDateTimeOk returns a tuple with the DateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTime

`func (o *GetEvents200ResponseDataInner) SetDateTime(v time.Time)`

SetDateTime sets DateTime field to given value.

### HasDateTime

`func (o *GetEvents200ResponseDataInner) HasDateTime() bool`

HasDateTime returns a boolean if a field has been set.

### GetDetails

`func (o *GetEvents200ResponseDataInner) GetDetails() map[string]interface{}`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *GetEvents200ResponseDataInner) GetDetailsOk() (*map[string]interface{}, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *GetEvents200ResponseDataInner) SetDetails(v map[string]interface{})`

SetDetails sets Details field to given value.

### HasDetails

`func (o *GetEvents200ResponseDataInner) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetTaskId

`func (o *GetEvents200ResponseDataInner) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *GetEvents200ResponseDataInner) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *GetEvents200ResponseDataInner) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *GetEvents200ResponseDataInner) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### GetPayload

`func (o *GetEvents200ResponseDataInner) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *GetEvents200ResponseDataInner) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *GetEvents200ResponseDataInner) SetPayload(v string)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *GetEvents200ResponseDataInner) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetUser

`func (o *GetEvents200ResponseDataInner) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *GetEvents200ResponseDataInner) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *GetEvents200ResponseDataInner) SetUser(v User)`

SetUser sets User field to given value.

### HasUser

`func (o *GetEvents200ResponseDataInner) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetLinks

`func (o *GetEvents200ResponseDataInner) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetEvents200ResponseDataInner) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetEvents200ResponseDataInner) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetEvents200ResponseDataInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


