# GetResponsePolicyImports200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The type of import | [optional] [readonly] 
**FileInfo** | Pointer to **interface{}** |  | [optional] 
**FileData** | Pointer to **interface{}** |  | [optional] 
**Name** | Pointer to **string** | The name of the import | [optional] [readonly] 
**ContentType** | Pointer to **string** | The content type of the import data | [optional] [readonly] 
**Size** | Pointer to **int64** | The size of the import, in bytes | [optional] [readonly] 
**State** | Pointer to **string** | The current state of the import | [optional] [readonly] 
**CreationDateTime** | Pointer to **time.Time** | The date and time the import resource was created | [optional] [readonly] 
**StartDateTime** | Pointer to **time.Time** | The date and time the import was started | [optional] [readonly] 
**CompletionDateTime** | Pointer to **time.Time** | The date and time the import completed | [optional] [readonly] 
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 

## Methods

### NewGetResponsePolicyImports200ResponseDataInner

`func NewGetResponsePolicyImports200ResponseDataInner() *GetResponsePolicyImports200ResponseDataInner`

NewGetResponsePolicyImports200ResponseDataInner instantiates a new GetResponsePolicyImports200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetResponsePolicyImports200ResponseDataInnerWithDefaults

`func NewGetResponsePolicyImports200ResponseDataInnerWithDefaults() *GetResponsePolicyImports200ResponseDataInner`

NewGetResponsePolicyImports200ResponseDataInnerWithDefaults instantiates a new GetResponsePolicyImports200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetResponsePolicyImports200ResponseDataInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetResponsePolicyImports200ResponseDataInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetResponsePolicyImports200ResponseDataInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetResponsePolicyImports200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GetResponsePolicyImports200ResponseDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetResponsePolicyImports200ResponseDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetResponsePolicyImports200ResponseDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetResponsePolicyImports200ResponseDataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFileInfo

`func (o *GetResponsePolicyImports200ResponseDataInner) GetFileInfo() interface{}`

GetFileInfo returns the FileInfo field if non-nil, zero value otherwise.

### GetFileInfoOk

`func (o *GetResponsePolicyImports200ResponseDataInner) GetFileInfoOk() (*interface{}, bool)`

GetFileInfoOk returns a tuple with the FileInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileInfo

`func (o *GetResponsePolicyImports200ResponseDataInner) SetFileInfo(v interface{})`

SetFileInfo sets FileInfo field to given value.

### HasFileInfo

`func (o *GetResponsePolicyImports200ResponseDataInner) HasFileInfo() bool`

HasFileInfo returns a boolean if a field has been set.

### SetFileInfoNil

`func (o *GetResponsePolicyImports200ResponseDataInner) SetFileInfoNil(b bool)`

 SetFileInfoNil sets the value for FileInfo to be an explicit nil

### UnsetFileInfo
`func (o *GetResponsePolicyImports200ResponseDataInner) UnsetFileInfo()`

UnsetFileInfo ensures that no value is present for FileInfo, not even an explicit nil
### GetFileData

`func (o *GetResponsePolicyImports200ResponseDataInner) GetFileData() interface{}`

GetFileData returns the FileData field if non-nil, zero value otherwise.

### GetFileDataOk

`func (o *GetResponsePolicyImports200ResponseDataInner) GetFileDataOk() (*interface{}, bool)`

GetFileDataOk returns a tuple with the FileData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileData

`func (o *GetResponsePolicyImports200ResponseDataInner) SetFileData(v interface{})`

SetFileData sets FileData field to given value.

### HasFileData

`func (o *GetResponsePolicyImports200ResponseDataInner) HasFileData() bool`

HasFileData returns a boolean if a field has been set.

### SetFileDataNil

`func (o *GetResponsePolicyImports200ResponseDataInner) SetFileDataNil(b bool)`

 SetFileDataNil sets the value for FileData to be an explicit nil

### UnsetFileData
`func (o *GetResponsePolicyImports200ResponseDataInner) UnsetFileData()`

UnsetFileData ensures that no value is present for FileData, not even an explicit nil
### GetName

`func (o *GetResponsePolicyImports200ResponseDataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetResponsePolicyImports200ResponseDataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetResponsePolicyImports200ResponseDataInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetResponsePolicyImports200ResponseDataInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetContentType

`func (o *GetResponsePolicyImports200ResponseDataInner) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *GetResponsePolicyImports200ResponseDataInner) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *GetResponsePolicyImports200ResponseDataInner) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *GetResponsePolicyImports200ResponseDataInner) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetSize

`func (o *GetResponsePolicyImports200ResponseDataInner) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *GetResponsePolicyImports200ResponseDataInner) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *GetResponsePolicyImports200ResponseDataInner) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *GetResponsePolicyImports200ResponseDataInner) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetState

`func (o *GetResponsePolicyImports200ResponseDataInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetResponsePolicyImports200ResponseDataInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetResponsePolicyImports200ResponseDataInner) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GetResponsePolicyImports200ResponseDataInner) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCreationDateTime

`func (o *GetResponsePolicyImports200ResponseDataInner) GetCreationDateTime() time.Time`

GetCreationDateTime returns the CreationDateTime field if non-nil, zero value otherwise.

### GetCreationDateTimeOk

`func (o *GetResponsePolicyImports200ResponseDataInner) GetCreationDateTimeOk() (*time.Time, bool)`

GetCreationDateTimeOk returns a tuple with the CreationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDateTime

`func (o *GetResponsePolicyImports200ResponseDataInner) SetCreationDateTime(v time.Time)`

SetCreationDateTime sets CreationDateTime field to given value.

### HasCreationDateTime

`func (o *GetResponsePolicyImports200ResponseDataInner) HasCreationDateTime() bool`

HasCreationDateTime returns a boolean if a field has been set.

### GetStartDateTime

`func (o *GetResponsePolicyImports200ResponseDataInner) GetStartDateTime() time.Time`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *GetResponsePolicyImports200ResponseDataInner) GetStartDateTimeOk() (*time.Time, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *GetResponsePolicyImports200ResponseDataInner) SetStartDateTime(v time.Time)`

SetStartDateTime sets StartDateTime field to given value.

### HasStartDateTime

`func (o *GetResponsePolicyImports200ResponseDataInner) HasStartDateTime() bool`

HasStartDateTime returns a boolean if a field has been set.

### GetCompletionDateTime

`func (o *GetResponsePolicyImports200ResponseDataInner) GetCompletionDateTime() time.Time`

GetCompletionDateTime returns the CompletionDateTime field if non-nil, zero value otherwise.

### GetCompletionDateTimeOk

`func (o *GetResponsePolicyImports200ResponseDataInner) GetCompletionDateTimeOk() (*time.Time, bool)`

GetCompletionDateTimeOk returns a tuple with the CompletionDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionDateTime

`func (o *GetResponsePolicyImports200ResponseDataInner) SetCompletionDateTime(v time.Time)`

SetCompletionDateTime sets CompletionDateTime field to given value.

### HasCompletionDateTime

`func (o *GetResponsePolicyImports200ResponseDataInner) HasCompletionDateTime() bool`

HasCompletionDateTime returns a boolean if a field has been set.

### GetLinks

`func (o *GetResponsePolicyImports200ResponseDataInner) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetResponsePolicyImports200ResponseDataInner) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetResponsePolicyImports200ResponseDataInner) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetResponsePolicyImports200ResponseDataInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


